package service

import (
	"loan-server/config"
	"loan-server/db"
	"strings"
	"testing"
)

var database, _ = db.Init(&config.Db{Dsn: "dev:dev123456@tcp(192.168.188.233:3306)/loan?charset=utf8mb4&parseTime=True&loc=UTC"})
var service = NewLeoChainService(&config.Leo{
	"https://api.explorer.provable.com/v1",
	"testnet",
	"aleo14s8rn8km6uqwatevtk9qyf6vut0vuce398ghklqggmvrfskr7cxsrx5p7a",
}, database)

func TestLeoChainService_GetLatestBlockOnChain(t *testing.T) {
	service.GetLatestBlockOnChain()
}

func TestLeoChainService_GetTransactionsInBlock(t *testing.T) {
	transactions, err := service.GetTransactionsInBlock(4223631)
	if err != nil {
		t.Fatal(err)
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

								}
								for _, value := range values {
									print(value + "\n")
								}
							}
						}
					}
				}
			}
		}
	}
}
