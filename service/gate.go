package service

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"github.com/gateio/gateapi-go/v6"
	"github.com/shopspring/decimal"
	"loan-server/model"
	"time"
)

func (s *BscChainService) SaveSoldAleo(loanId uint) error {
	loan, err := s.Db.SelectLoanById(loanId)
	if err != nil {
		return err
	}
	deposits, err := s.Db.SelectDepositByLoanId(int(loanId))
	if err != nil {
		return err
	}
	aleoAmount := decimal.Zero
	for _, deposit := range deposits {
		aleoAmount = aleoAmount.Add(deposit.AleoAmount)
	}

	tx := s.Db.Db.Create(&model.SoldWithdrawAleo{
		LoanId:      loanId,
		AleoAmount:  aleoAmount,
		AleoAddress: loan.AleoAddress,
		Status:      1,
		CreateAt:    int(time.Now().Unix()),
	})
	if tx.Error != nil {
		return tx.Error
	}
	s.SoldAleo(loanId)
	return nil
}

func (s *BscChainService) SoldAleo(loanId uint) {
	s.placeMarketSellOrder("1", int(loanId))
}

func (s *BscChainService) getGateApiClient() *gateapi.APIClient {
	return gateapi.NewAPIClient(gateapi.NewConfiguration())
}

func (s *BscChainService) getGateApiCtx() context.Context {
	return context.WithValue(context.Background(),
		gateapi.ContextGateAPIV4,
		gateapi.GateAPIV4{
			Key:    s.GateIoConfig.ApiKey,
			Secret: s.GateIoConfig.SecretKey,
		},
	)
}

func (s *BscChainService) placeMarketSellOrder(amount string, loanId int) {

	client := s.getGateApiClient()
	ctx := s.getGateApiCtx()

	order := gateapi.Order{
		Text:         fmt.Sprintf("t-%d", loanId),
		CurrencyPair: "ALEO_USDT",
		Type:         "market",
		Side:         "sell",
		Amount:       amount,
		TimeInForce:  "ioc",
	}
	localVarOptionals := gateapi.CreateOrderOpts{}

	result, response, err := client.SpotApi.CreateOrder(
		ctx,
		order,
		&localVarOptionals)

	if err != nil {
		if e, ok := err.(gateapi.GateAPIError); ok {
			fmt.Printf("gate api error: %s\n", e.Error())
		} else {
			fmt.Printf("generic error: %s\n", err.Error())
		}
	} else {
		fmt.Println(result)
		fmt.Println(response)
	}

}

func (s *BscChainService) getTradeHistory() {

	client := s.getGateApiClient()
	ctx := s.getGateApiCtx()

	localVarOptionals := gateapi.ListOrdersOpts{
		Page:  optional.NewInt32(int32(1)),
		Limit: optional.NewInt32(int32(100)),
		Side:  optional.NewString("sell"),
	}

	result, response, err := client.SpotApi.ListOrders(
		ctx,
		"ALEO_USDT",
		"finished",
		&localVarOptionals)
	if err != nil {
		if e, ok := err.(gateapi.GateAPIError); ok {
			fmt.Printf("gate api error: %s\n", e.Error())
		} else {
			fmt.Printf("generic error: %s\n", err.Error())
		}
	} else {
		fmt.Println(result)
		fmt.Println(response)
	}
}

func (s *BscChainService) withdrawUSDT(address, amount string) {
	client := s.getGateApiClient()
	ctx := s.getGateApiCtx()

	result, response, err := client.WithdrawalApi.Withdraw(ctx, gateapi.LedgerRecord{
		Currency: "USDT",
		Amount:   amount,
		Address:  address,
		Chain:    "BSC",
	})
	if err != nil {
		if e, ok := err.(gateapi.GateAPIError); ok {
			fmt.Printf("gate api error: %s\n", e.Error())
		} else {
			fmt.Printf("generic error: %s\n", err.Error())
		}
	} else {
		fmt.Println(result)
		fmt.Println(response)
	}
}

func (s *BscChainService) checkWithdraw() {
	client := s.getGateApiClient()
	ctx := s.getGateApiCtx()

	localVarOptionals := gateapi.ListWithdrawStatusOpts{
		Currency: optional.NewString("USDT"),
	}

	result, response, err := client.WalletApi.ListWithdrawStatus(ctx, &localVarOptionals)
	if err != nil {
		if e, ok := err.(gateapi.GateAPIError); ok {
			fmt.Printf("gate api error: %s\n", e.Error())
		} else {
			fmt.Printf("generic error: %s\n", err.Error())
		}
	} else {
		fmt.Println(result)
		fmt.Println(response)
	}
}
