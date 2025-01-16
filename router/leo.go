package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"loan-server/common/consts"
	"loan-server/common/errors"
	"loan-server/common/res"
	"loan-server/common/utils"
	"loan-server/model"
	"strconv"
)

func (myRouter *Router) loadLeoRouters(engine *gin.Engine) {
	engine.GET("/leo/config", loanConfig(myRouter))
	engine.GET("/leo/calculate_usdt", calculateUsdt(myRouter))
	engine.POST("/leo/save_deposoit", saveDeposit(myRouter))
	engine.GET("/leo/loan_list", leoLoanList(myRouter))
}

// loanConfig
func loanConfig(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		config, err := myRouter.mydb.GetConfig()
		if err != nil {
			panic(errors.New(errors.SystemError))
		}
		price, err := myRouter.mydb.GetLatestPrice()
		if err != nil {
			panic(errors.New(errors.SystemError))
		}
		rate, err := myRouter.mydb.GetLatestRate()
		if err != nil {
			panic(errors.New(errors.SystemError))
		}
		resCfg := model.LeoResConfig{
			Rate:           strconv.FormatFloat(rate, 'f', 6, 64),
			AvailableStage: config.AvailableStages,
			DayPerStage:    config.DayPerStage,
			Price:          strconv.FormatFloat(price, 'f', 6, 64),
			AllowTypes:     "1",
			Banners:        []string{},
			MinAmount:      config.MinLoanAmount / 1000000,
			MaxAmount:      config.MaxLoanAmount / 1000000,
		}
		success := res.Success(resCfg)
		context.JSON(success.Code, success)
	}
}

// calculateUsdt
func calculateUsdt(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		config, err := myRouter.mydb.GetConfig()
		if err != nil {
			panic(errors.New(errors.SystemError))
		}

		as := context.Query("amount")
		if as == "" {
			panic(errors.New(errors.ParameterError))
		}
		amount, err := strconv.ParseFloat(as, 64)
		intAmount := int64(amount * 1000000)
		if err != nil {
			panic(errors.New(errors.ParameterError))
		}
		if config.MinLoanAmount > intAmount || config.MaxLoanAmount < intAmount {
			panic(errors.New(errors.ParameterError))
		}

		price, err := myRouter.mydb.GetLatestPrice()
		if err != nil {
			panic(errors.New(errors.SystemError))
		}
		rate, err := myRouter.mydb.GetLatestRate()
		if err != nil {
			panic(errors.New(errors.SystemError))
		}

		collateralAmount := decimal.NewFromFloat(price).Mul(decimal.NewFromFloat(amount))
		borrowAmount := collateralAmount.Mul(config.ReleaseRate)

		var installments []model.LeoResInstallment
		for i := 1; i <= config.AvailableStages; i++ {
			interestRate := decimal.NewFromFloat(rate).Mul(decimal.NewFromInt(int64(i * config.DayPerStage)))
			installments = append(installments, model.LeoResInstallment{
				Installments:           i,
				DayPerInstallment:      config.DayPerStage,
				InterestRate:           interestRate.String(),
				InterestPerInstallment: borrowAmount.Mul(interestRate).String(),
			})
		}
		result := model.LeoResCalculateUsdt{
			BorrowingAmount:  borrowAmount.String(),
			CollateralAmount: collateralAmount.String(),
			CollateralRate:   config.ReleaseRate.String(),
			Installment:      installments,
		}

		//usdt := .Mul(decimal.NewFromInt(1).Sub(decimal.NewFromFloat(rate)))
		success := res.Success(result)
		context.JSON(success.Code, success)
	}
}

/*
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"aleo_address":"aleo1hac8kndgfp7eh545yeu6k2ue32yn3dt7qe5xl54d6lpe7xecyq9qkxc3tx", "aleo_amount":1000000, "bsc_address":"0x4332B66D46761476B0A50A2F12EE6a17DaCe7247", "email": "1140830756@qq.com", "stages": 2, "day_per_stage": 7, "type": 1}' \
  https://wings.tcds.ltd/leo/save_deposoit
*/
// saveDeposit
func saveDeposit(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		var params model.ReqSaveDeposit
		if err := context.ShouldBindJSON(&params); err != nil {
			panic(errors.New(errors.ParameterError))
		}
		if !utils.IsValidLeoAddress(params.AleoAddress) {
			panic(errors.New(errors.ParameterError))
		}
		if !utils.IsValidAddress(params.BscAddress) {
			panic(errors.New(errors.ParameterError))
		}
		if !utils.VerifyEmailFormat(params.Email) {
			panic(errors.New(errors.ParameterError))
		}
		if params.LoanType != 1 {
			panic(errors.New(errors.ParameterError))
		}
		if params.Type < 0 || params.Type > 2 {
			panic(errors.New(errors.ParameterError))
		}
		config, err := myRouter.mydb.GetConfig()
		if err != nil {
			panic(errors.New(errors.SystemError))
		}
		if config.DayPerStage != params.DayPerStage {
			panic(errors.New(errors.ParameterError))
		}
		if config.AvailableStages < params.Stages || params.Stages < 0 {
			panic(errors.New(errors.ParameterError))
		}
		intAmount := int64(params.AleoAmount * 1000000)
		if config.MinLoanAmount > intAmount || config.MaxLoanAmount < intAmount {
			panic(errors.New(errors.ParameterError))
		}
		err = myRouter.mydb.NewDeposit(
			params.AleoAddress,
			params.BscAddress,
			params.Email,
			intAmount,
			params.Stages,
			params.DayPerStage)
		if err != nil {
			panic(errors.New(errors.SystemError))
		}
		success := res.Success(map[string]bool{
			"success": true,
		})
		context.JSON(success.Code, success)
	}
}

// leoLoanList
func leoLoanList(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		address := context.Query("address")
		if address == "" {

			panic(errors.New(errors.ParameterError))
		}
		if !utils.IsValidLeoAddress(address) {
			panic(errors.New(errors.ParameterError))
		}
		loans, err := myRouter.mydb.SelectLoan(address)
		if err != nil {
			return
		}
		var resLoans []model.LeoResLoan
		for _, loan := range loans {
			var deposits []model.LeoResDeposit
			ds, err := myRouter.mydb.SelectDepositByLoanId(int(loan.ID))
			if err != nil {
				return
			}
			for _, d := range ds {
				deposits = append(deposits, model.LeoResDeposit{
					ID:          int(d.ID),
					LoanId:      d.LoanId,
					AleoAddress: d.AleoAddress,
					AleoAmount:  d.AleoAmount.String(),
					AleoPrice:   d.AleoPrice.String(),
					UsdtValue:   d.UsdtValue.Div(decimal.NewFromInt(consts.Wei)).String(),
					Hash:        d.Hash,
					At:          d.At,
					Status:      d.Status,
				})
			}
			resLoans = append(resLoans, model.LeoResLoan{
				ID:                int(loan.ID),
				AleoAddress:       loan.AleoAddress,
				BscAddress:        loan.BscAddress,
				Email:             loan.Email,
				Status:            loan.Status,
				Stages:            loan.Stages,
				DayPerStage:       loan.DayPerStage,
				StartAt:           loan.StartAt,
				Health:            loan.Health.String(),
				Rate:              loan.Rate.String(),
				ReleaseRate:       loan.ReleaseRate.String(),
				Hash:              loan.Hash,
				Type:              loan.Type,
				BscLoanId:         loan.BscLoanId,
				ReleaseAt:         loan.ReleaseAt,
				ReleaseHash:       loan.ReleaseHash,
				ReleaseAmount:     loan.ReleaseAmount.Div(decimal.NewFromInt(consts.Wei)).String(),
				PayBackAt:         loan.PayBackAt,
				PayBackHash:       loan.PayBackHash,
				PayBackAmount:     loan.PayBackAmount.String(),
				ReleaseAleoHash:   loan.ReleaseAleoHash,
				ReleaseAleoAt:     loan.ReleaseAleoAt,
				ReleaseAleoAmount: loan.ReleaseAleoAmount.Div(decimal.NewFromInt(consts.Wei)).String(),
				Deposits:          deposits,
			})
		}
		success := res.Success(resLoans)
		context.JSON(success.Code, success)
	}
}
