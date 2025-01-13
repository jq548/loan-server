package service

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
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
			if blocks > 100 {
				blocks = 100
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
	filterer, err := contract.NewContractFilterer(common.HexToAddress(s.LoanContractAddress), s.EthClient)
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
			case "0xb2990d266ec4e479259ef8c68e87d6c03ab8dbafa4e785e79d2ed1545a383083":
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
					decimal.NewFromBigInt(params.Amount, 1))
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
			case "0x045492b29efb11990e42d3b9ba88978042183c94051194b886f8595139c64db9":
				params, err := filterer.ParseEventIncreaseLiquidReward(log)
				if err != nil {
					return err
				}
				err = s.Db.IncreaseProviderRewardAmount(
					decimal.NewFromBigInt(params.Amount, 1),
					params.Provider.Hex(),
					log.TxHash.Hex(),
					timestamp)
				if err != nil {
					return err
				}
			case "0x467456fb7eb39617bb976fc80dbff34a252f7ef0f1887ebcca77c2caaac5704d":
				params, err := filterer.ParseEventReleaseLiquidReward(log)
				if err != nil {
					return err
				}
				err = s.Db.ReleaseProviderReward(
					decimal.NewFromBigInt(params.Amount, 1),
					params.Provider.Hex(),
					log.TxHash.Hex(),
					timestamp)
				if err != nil {
					return err
				}
			case "0xc7f079bb1739a7fcb563479a77ec3ff5de5ba875b2b8b44d897abfc3ac58a8ed":
				params, err := filterer.ParseEventProviderIncrease(log)
				if err != nil {
					return err
				}
				err = s.Db.IncreaseProviderAmount(
					decimal.NewFromBigInt(params.Amount, 1),
					params.Provider.Hex(),
					log.TxHash.Hex(),
					timestamp)
				if err != nil {
					return err
				}
			case "0x89fe3f29313aa6c03800bc780dc2b251ec749edd116569f40bd397cf1e8e08c8":
				params, err := filterer.ParseEventProviderRetrieve(log)
				if err != nil {
					return err
				}
				err = s.Db.RetrieveProviderAmount(
					decimal.NewFromBigInt(params.Amount, 1),
					params.Provider.Hex(),
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
