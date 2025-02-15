package service

import (
	"encoding/json"
	errors2 "errors"
	"fmt"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"io/ioutil"
	"loan-server/common/consts"
	"loan-server/config"
	"loan-server/db"
	"loan-server/model"
	"math/big"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// block height
// https://api.explorer.provable.com/v1/testnet/block/height/latest
// transaction in block
// https://api.explorer.provable.com/v1/testnet/blocks?start=4249173&end=4249174
// block info
// https://api.explorer.provable.com/v1/testnet/block/4249173

type LeoChainService struct {
	Rpc                  string
	Net                  string
	Db                   *db.MyDb
	holder               string
	holderPK             string
	lastCheckBlockNumber int
	BscService           *BscChainService
}

func NewLeoChainService(leo *config.Leo, myDb *db.MyDb) *LeoChainService {
	return &LeoChainService{
		Rpc:                  leo.Rpc,
		Net:                  leo.Net,
		holder:               leo.Holder,
		holderPK:             leo.HolderPK,
		Db:                   myDb,
		lastCheckBlockNumber: 0,
	}
}

func (s *LeoChainService) Start() {
	for {
		height, err := s.Db.GetLeoBlockHeight()
		if err != nil {
			zap.S().Error(err)
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
			if blocks < 5 {
				time.Sleep(1 * time.Second)
				continue
			}
			if blocks > 30 {
				blocks = 30
			}

			blockData, err := s.GetBlocks(height, blocks+height-1)
			if err != nil {
				zap.S().Error(err)
				time.Sleep(1 * time.Second)
				continue
			}
			zap.S().Infof("Filtered Leo Blocks from:%v to:%v", height, blocks+height-1)
			for _, block := range blockData {
				transactions := block.Transactions
				for _, transaction := range transactions {
					if transaction.Status == "accepted" && transaction.Type == "execute" {
						for _, transition := range transaction.Transaction.Execution.Transitions {
							if transition.Program == "credits.aleo" && transition.Function == "transfer_public" {
								if len(transition.Outputs) > 0 {
									output := transition.Outputs[0].Value
									output = strings.ReplaceAll(output, "\n", "")
									output = strings.ReplaceAll(output, " ", "")
									if strings.Contains(output, "program_id:credits.aleo") && strings.Contains(output, "function_name:transfer_public") {
										sub := strings.Split(output, "arguments:[")
										if len(sub) == 2 {
											sub1 := strings.ReplaceAll(sub[1], "]}", "")
											values := strings.Split(sub1, ",")
											if len(values) == 3 {
												if values[1] == s.holder || values[0] == s.holder {
													err = s.SaveBlockTransaction(transition.ID, values[0], values[1], values[2], block.Header.Metadata.Timestamp)
													if err != nil {
														zap.S().Error(err)
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
			// error when saving
			if err != nil {
				zap.S().Error(err)
				continue
			}
			err = s.Db.SaveLeoBlockHeight(int(blocks + height))
			if err != nil {
				zap.S().Error(err)
				continue
			}
		} else {
			time.Sleep(3 * time.Second)
		}
	}
}

func (s *LeoChainService) getRequest(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (s *LeoChainService) GetLatestBlockOnChain() error {
	url := fmt.Sprintf("%s/%s/block/height/latest", s.Rpc, s.Net)
	res, err := s.getRequest(url)
	if err != nil {
		return err
	}
	s.lastCheckBlockNumber, err = strconv.Atoi(res)
	if err != nil {
		return err
	}
	return nil
}

func (s *LeoChainService) GetBlocks(from, to int) ([]model.LeoBlock, error) {
	url := fmt.Sprintf("%s/%s/blocks?start=%d&end=%d", s.Rpc, s.Net, from, to)
	res, err := s.getRequest(url)
	if err != nil {
		return nil, err
	}
	var trans []model.LeoBlock
	err = json.Unmarshal([]byte(res), &trans)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, err
	}
	return trans, nil
}

func (s *LeoChainService) SaveBlockTransaction(
	txId,
	sender,
	receiver,
	amount string,
	blockTime int) error {

	as := strings.ReplaceAll(amount, "u64", "")
	parseInt, err := strconv.ParseInt(as, 10, 64)
	if err != nil {
		return err
	}

	if sender == s.holder {
		err := s.Db.UpdateReleaseAleoBack(receiver, txId, decimal.NewFromInt(parseInt), blockTime)
		if err != nil {
			return err
		}
	} else {
		deposit, err := s.Db.SelectUnConfirmDepositByAddress(sender, parseInt)
		if err != nil {
			return err
		}
		if len(deposit) > 0 {
			if deposit[0].Hash == txId {
				return nil
			}
			sloan, err := s.Db.SelectLoanById(deposit[0].LoanId)
			if err != nil {
				return err
			}
			loanAmount, interestAmount, price, rate, err := s.CalculateReleaseUsdt(parseInt, sloan.Stages, sloan.DayPerStage, sloan.ReleaseRate)
			if err != nil {
				return err
			}
			loan, err := s.Db.SaveDepositHash(txId, deposit[0].ID, blockTime, loanAmount, interestAmount, price, rate)
			if err != nil {
				return err
			}
			if loan.Status == 1 && loan.ReleaseHash == "" {
				err = s.BscService.CreateLoanInContract(
					big.NewInt(int64(loan.ID)),
					loanAmount.BigInt(),
					big.NewInt(int64(loan.Stages*loan.DayPerStage*24*3600)),
					interestAmount.BigInt(),
					loan.BscAddress,
					loan.AleoAddress,
					loan.DepositAmount.BigInt(),
					price.BigInt(),
				)
				if err != nil {
					zap.S().Error(err)
					err = s.Db.SaveCreateFailed(int(loan.ID), 6)
					if err != nil {
						return err
					}
				}
			}
			return nil
		}
	}

	return nil
}

func (s *LeoChainService) PayBackLoan(to, amount string) error {
	cmd := exec.Command(
		"node",
		"/Users/jq/Desktop/aleo-project/index.js",
		s.Rpc,
		s.holderPK,
		to,
		amount,
	)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	if !strings.Contains(string(output), "success") {
		return errors2.New(string(output))
	}
	return nil
}

func (s *LeoChainService) CalculateReleaseUsdt(aleoAmount int64, stages, perDay int, releaseRate decimal.Decimal) (decimal.Decimal, decimal.Decimal, decimal.Decimal, decimal.Decimal, error) {
	price, err := s.Db.GetLatestPrice()
	if err != nil {
		return decimal.NewFromInt(0), decimal.NewFromInt(0), decimal.NewFromInt(0), decimal.NewFromInt(0), err
	}
	rates, err := s.Db.GetLatestRateOfWeek()
	if err != nil {
		return decimal.NewFromInt(0), decimal.NewFromInt(0), decimal.NewFromInt(0), decimal.NewFromInt(0), err
	}
	rate := rates[0].Rate
	for _, r := range rates {
		if r.Days == stages*perDay {
			rate = r.Rate
			break
		}
	}
	rate = rate.Mul(decimal.NewFromInt(int64(stages)))
	aleoA := decimal.NewFromInt(aleoAmount).Div(decimal.NewFromInt(1000000))
	usdtAmount := aleoA.Mul(decimal.NewFromFloat(price).Mul(decimal.NewFromInt(consts.Wei))).Mul(releaseRate)
	interestAmount := usdtAmount.Mul(rate)
	return usdtAmount, interestAmount, decimal.NewFromFloat(price), rate, nil
}
