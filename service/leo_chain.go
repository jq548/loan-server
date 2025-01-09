package service

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"loan-server/config"
	"loan-server/db"
	"loan-server/model"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// block height
// https://api.explorer.provable.com/v1/testnet/block/height/latest
// transaction in block
// https://api.explorer.provable.com/v1/testnet/block/4223631/transactions

type LeoChainService struct {
	Rpc                  string
	Net                  string
	Db                   *db.MyDb
	holder               string
	lastCheckBlockNumber int
}

func NewLeoChainService(leo *config.Leo, myDb *db.MyDb) *LeoChainService {
	return &LeoChainService{
		Rpc:                  leo.Rpc,
		Net:                  leo.Net,
		holder:               leo.Holder,
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
		if s.lastCheckBlockNumber == 0 || s.lastCheckBlockNumber < height+5 {
			err := s.GetLatestBlockOnChain()
			if err != nil {
				zap.S().Error(err)
				time.Sleep(1 * time.Second)
				continue
			}
		}
		if s.lastCheckBlockNumber > height+5 {
			transactions, err := s.GetTransactionsInBlock(height + 1)
			if err != nil {
				zap.S().Error(err)
				time.Sleep(1 * time.Second)
				continue
			}
			currentTime, err := s.CalculateTimeForBlock(height + 1)
			if err != nil {
				zap.S().Error(err)
				time.Sleep(1 * time.Second)
				continue
			}
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
											if values[1] == s.holder {
												err = s.SaveBlockTransaction(transition.ID, values[0], values[2], currentTime)
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
			// error when saving
			if err != nil {
				zap.S().Error(err)
				continue
			}
			err = s.Db.SaveLeoBlockHeight(height + 1)
			if err != nil {
				zap.S().Error(err)
				continue
			}
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

func (s *LeoChainService) GetTransactionsInBlock(blockNumber int) ([]model.LeoTransaction, error) {
	url := fmt.Sprintf("%s/%s/block/%d/transactions", s.Rpc, s.Net, blockNumber)
	res, err := s.getRequest(url)
	if err != nil {
		return nil, err
	}
	//var result []map[string]interface{}
	//err = json.Unmarshal([]byte(res), &result)

	var trans []model.LeoTransaction
	err = json.Unmarshal([]byte(res), &trans)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, err
	}
	return trans, nil
}

func (s *LeoChainService) SaveBlockTransaction(
	txId,
	address,
	amount string,
	time time.Time) error {
	return nil
}

func (s *LeoChainService) CalculateTimeForBlock(blockNumber int) (time.Time, error) {
	return time.Now(), nil
}
