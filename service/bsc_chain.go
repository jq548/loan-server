package service

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"loan-server/common/utils"
	"loan-server/config"
	"loan-server/contract"
	"loan-server/db"
	"math/big"
	"time"
)

type BscChainService struct {
	Rpc                  string
	Db                   *db.MyDb
	LoanContractAddress  string
	UsdtContractAddress  string
	LPContractAddress    string
	lastCheckBlockNumber int64
	EthClient            *ethclient.Client
	blockTimeMap         map[int]int
	CallerPk             string
	LeoService           *LeoChainService
	ChainId              int
}

func NewBscChainService(bsc *config.Bsc, myDb *db.MyDb) (*BscChainService, error) {
	client, err := ethclient.Dial(bsc.Rpc)
	if err != nil {
		return nil, err
	}
	return &BscChainService{
		Rpc:                 bsc.Rpc,
		Db:                  myDb,
		LoanContractAddress: bsc.LoanContract,
		UsdtContractAddress: bsc.Usdt,
		LPContractAddress:   bsc.Lp,
		EthClient:           client,
		blockTimeMap:        make(map[int]int),
		CallerPk:            bsc.Caller,
		ChainId:             bsc.ChainId,
	}, nil
}

func (s *BscChainService) StartFetchEvent() {
	for {
		height, err := s.Db.GetBscBlockHeight()
		if err != nil {
			zap.S().Error(err)
			time.Sleep(1 * time.Second)
			continue
		}
		if s.lastCheckBlockNumber == 0 || s.lastCheckBlockNumber < height+10 {
			err := s.GetLatestBlockOnChain()
			if err != nil {
				zap.S().Error(err)
				time.Sleep(1 * time.Second)
				continue
			}
		}
		if s.lastCheckBlockNumber > height+5 {
			blocks := s.lastCheckBlockNumber - height - 5
			if blocks > 20 {
				blocks = 20
			}
			err := s.FilterLogs(height, blocks+height-1)
			if err != nil {
				zap.S().Error(err)
				continue
			}
			err = s.Db.SaveBscBlockHeight(int(blocks + height))
			if err != nil {
				zap.S().Error(err)
				continue
			}
		} else {
			time.Sleep(3 * time.Second)
		}

	}
}

func (s *BscChainService) GetLatestBlockOnChain() error {
	number, err := s.EthClient.BlockNumber(context.Background())
	if err != nil {
		return err
	}
	s.lastCheckBlockNumber = int64(number)
	return nil
}

func (s *BscChainService) GetCallOpts() *bind.CallOpts {
	return &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
}

func (s *BscChainService) GetTransactOpts(pk string) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return nil, err
	}
	gasPrice, err := s.EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(s.ChainId)))
	if err != nil {
		return nil, err
	}
	transactor.Value = big.NewInt(0)
	transactor.GasPrice = gasPrice
	return transactor, nil
}

func (s *BscChainService) FilterLogs(from, to int64) error {
	filterQuery := ethereum.FilterQuery{
		FromBlock: big.NewInt(from),
		ToBlock:   big.NewInt(to),
		Addresses: []common.Address{
			common.HexToAddress(s.LoanContractAddress),
		},
	}
	logs, err := s.EthClient.FilterLogs(context.Background(), filterQuery)
	if err != nil {
		return err
	}
	zap.S().Infof("Filtered BSC Logs from:%v to:%v", filterQuery.FromBlock, filterQuery.ToBlock)
	filterer, err := contract.NewLoanFilterer(common.HexToAddress(s.LoanContractAddress), s.EthClient)
	if err != nil {
		return err
	}
	for _, log := range logs {
		timestamp, err := s.ReqTimeForBlock(int(from), int(to), int(log.BlockNumber))
		if err != nil {
			return err
		}
		for _, topic := range log.Topics {
			switch topic.Hex() {
			case "0x8359c828396108eedea00704782ac2a600d822d6d56312c4e10f62408aedca5d":
				params, err := filterer.ParseEventNewLoan(log)
				if err != nil {
					return err
				}
				err = s.Db.SaveDepositOnBscHash(
					log.TxHash.Hex(),
					params.Loaner.Hex(),
					int(params.LoanId.Int64()),
					int(params.Duration.Int64()),
					int(params.Start.Int64()),
					timestamp,
					decimal.NewFromBigInt(params.ReleaseAmount, 1),
					decimal.NewFromBigInt(params.InterestAmount, 1))
				if err != nil {
					return err
				}
				//TODO calculate income of providers
			case "0x38dcb8e7ce8c7f182d53142ee0fa94a1778cd54eddcc1209f86469e7d3b48733":
				params, err := filterer.ParseEventPayBack(log)
				if err != nil {
					return err
				}
				err = s.Db.Payback(
					int(params.LoanId.Int64()),
					log.TxHash.Hex(),
					timestamp,
					decimal.NewFromBigInt(params.Amount, 1))
				if err != nil {
					return err
				}
				//TODO transfer aleo back
			case "0x5a322d6a1b1ff3f692c4c5995ba8133271b684bd011bc8e2f711d50db23bfe03":
				params, err := filterer.ParseEventClear(log)
				if err != nil {
					return err
				}
				err = s.Db.Clear(
					int(params.LoanId.Int64()),
					log.TxHash.Hex(),
					timestamp,
					decimal.NewFromBigInt(params.Amount, 1))
				if err != nil {
					return err
				}
			case "0xa16a887e0d16d473c5a8459cbf20c45ef7f0a282e60e60adcacb455ae31ebb62":
				params, err := filterer.ParseEventIncreaseLiquidRewardBath(log)
				if err != nil {
					return err
				}
				for i, id := range params.Ids {
					err = s.Db.IncreaseProviderRewardAmount(
						decimal.NewFromBigInt(params.Amounts[i], 1),
						params.Providers[i].Hex(),
						log.TxHash.Hex(),
						timestamp,
						int(id.Int64()))
					if err != nil {
						return err
					}
				}
			case "0x9d55a88ba6edf4a14f0ad37d9f0833bb65734beea749cfeff8d52825ffd58ef9":
				params, err := filterer.ParseEventReleaseLiquidReward(log)
				if err != nil {
					return err
				}
				err = s.Db.ReleaseProviderReward(
					decimal.NewFromBigInt(params.Amount, 1),
					params.Provider.Hex(),
					log.TxHash.Hex(),
					timestamp,
					decimal.NewFromBigInt(params.Fee, 1))
				if err != nil {
					return err
				}
			case "0x75780a70131ef5cf8aff25941e13a743681d34e1eb85abde44b32e09280e1fcc":
				params, err := filterer.ParseEventProviderAdd(log)
				if err != nil {
					return err
				}
				err = s.Db.IncreaseProviderAmount(
					decimal.NewFromBigInt(params.Amount, 1),
					int(params.Duration.Int64()),
					int(params.Start.Int64()),
					params.Provider.Hex(),
					log.TxHash.Hex(),
					timestamp,
					int(params.Id.Int64()))
				if err != nil {
					return err
				}
			case "0x1ec3add8915e0172b379ff4433663fbf4a45d4da64edb3e0402fb369e04b024a":
				params, err := filterer.ParseEventProviderRedeem(log)
				if err != nil {
					return err
				}
				err = s.Db.RetrieveProviderAmount(
					int(params.Id.Int64()),
					log.TxHash.Hex(),
					timestamp,
					decimal.NewFromBigInt(params.Fee, 1))
				if err != nil {
					return err
				}
			case "0x1a83ec821f4483d6557a8f7a974b0861f82494f76a8589611ab01a019489b491":
				params, err := filterer.ParseEventExchangeLpToUsdt(log)
				if err != nil {
					return err
				}
				err = s.Db.SaveExchangeLpToUsdtRecord(
					params.Forward,
					decimal.NewFromBigInt(params.Amount, 1),
					params.Caller.Hex(),
					log.TxHash.Hex(),
					timestamp)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (s *BscChainService) ReqTimeForBlock(from, to, blockNumber int) (int, error) {
	if from == to && from == blockNumber {
		fromBlock, err := s.EthClient.BlockByNumber(context.Background(), big.NewInt(int64(from)))
		if err != nil {
			return 0, err
		}
		return int(fromBlock.Time()), nil
	}
	if s.blockTimeMap[from] == 0 && s.blockTimeMap[to] == 0 {
		fromBlock, err := s.EthClient.BlockByNumber(context.Background(), big.NewInt(int64(from)))
		if err != nil {
			return 0, err
		}
		s.blockTimeMap[from] = int(fromBlock.Time())
		toBlock, err := s.EthClient.BlockByNumber(context.Background(), big.NewInt(int64(to)))
		if err != nil {
			return 0, err
		}
		s.blockTimeMap[to] = int(toBlock.Time())
	}
	averageForOneBlock := float32(s.blockTimeMap[to]-s.blockTimeMap[from]) / float32(to-from)
	return s.blockTimeMap[from] + int(averageForOneBlock*float32(blockNumber-from)), nil
}

func (s *BscChainService) CreateLoanInContract(
	id,
	amount,
	duration,
	interestAmount *big.Int,
	loaner string) error {
	loanContract, err := contract.NewLoan(common.HexToAddress(s.LoanContractAddress), s.EthClient)
	if err != nil {
		return err
	}
	loan, err := loanContract.Loans(s.GetCallOpts(), id)
	if err != nil {
		return err
	}
	if utils.IsZeroAddress(loan.Loaner) {
		opts, err := s.GetTransactOpts(s.CallerPk)
		if err != nil {
			return err
		}
		_, err = loanContract.AddNewLoan(opts, id, amount, duration, common.HexToAddress(loaner), interestAmount)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *BscChainService) ClearLoanInContract(loanId *big.Int) error {
	loanContract, err := contract.NewLoan(common.HexToAddress(s.LoanContractAddress), s.EthClient)
	if err != nil {
		return err
	}
	opts, err := s.GetTransactOpts(s.CallerPk)
	if err != nil {
		return err
	}
	_, err = loanContract.Clear(opts, loanId)
	if err != nil {
		return err
	}

	return nil
}

func (s *BscChainService) IncreaseLiquidReward(amount *big.Int, provider string) error {
	loanContract, err := contract.NewLoan(common.HexToAddress(s.LoanContractAddress), s.EthClient)
	if err != nil {
		return err
	}
	opts, err := s.GetTransactOpts(s.CallerPk)
	if err != nil {
		return err
	}
	_, err = loanContract.IncreaseLiquidReward(opts, amount, common.HexToAddress(provider))
	if err != nil {
		return err
	}

	return nil
}

func (s *BscChainService) CheckAddresses() error {
	loanContract, err := contract.NewLoan(common.HexToAddress(s.LoanContractAddress), s.EthClient)
	if err != nil {
		return err
	}
	res, err := loanContract.Addresses(s.GetCallOpts(), big.NewInt(0))
	if err != nil {
		return err
	}
	zap.S().Info(res.Hex())

	return nil
}

func (s *BscChainService) IncreaseIncome(ids []*big.Int, addresses []string, amounts []*big.Int) (string, error) {
	var providers []common.Address
	for _, address := range addresses {
		providers = append(providers, common.HexToAddress(address))
	}
	loanContract, err := contract.NewLoan(common.HexToAddress(s.LoanContractAddress), s.EthClient)
	if err != nil {
		return "", err
	}
	opts, err := s.GetTransactOpts(s.CallerPk)
	if err != nil {
		return "", err
	}
	tx, err := loanContract.IncreaseLiquidRewardBatch(opts, ids, amounts, providers)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}
