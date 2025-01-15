package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
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
		resCfg := model.LeoResConfig{
			Rate:           config.Rate.String(),
			AvailableStage: config.AvailableStages,
			DayPerStage:    config.DayPerStage,
			Price:          config.AleoPrice.String(),
			AllowTypes:     "1",
			Banners:        []string{"", ""},
			MinAmount:      config.MinLoanAmount,
			MaxAmount:      config.MaxLoanAmount,
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

		stages := context.Query("stage")
		if stages == "" {
			panic(errors.New(errors.ParameterError))
		}
		stage, err := strconv.Atoi(stages)
		if err != nil {
			panic(errors.New(errors.ParameterError))
		}
		if config.AvailableStages < stage || stage < 0 {
			panic(errors.New(errors.ParameterError))
		}

		dayPerStages := context.Query("day_per_stage")
		if dayPerStages == "" {
			panic(errors.New(errors.ParameterError))
		}
		dayPerStage, err := strconv.Atoi(dayPerStages)
		if err != nil {
			panic(errors.New(errors.ParameterError))
		}
		if config.DayPerStage != dayPerStage {
			panic(errors.New(errors.ParameterError))
		}

		usdt := config.AleoPrice.Mul(decimal.NewFromFloat(amount)).Mul(decimal.NewFromInt(1).Sub(config.Rate))
		success := res.Success(map[string]string{
			"usdt": usdt.String(),
		})
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
					AleoAmount:  d.AleoAmount,
					AleoPrice:   d.AleoPrice,
					UsdtValue:   d.UsdtValue,
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
				Health:            loan.Health,
				Rate:              loan.Rate,
				ReleaseRate:       loan.ReleaseRate,
				Hash:              loan.Hash,
				Type:              loan.Type,
				BscLoanId:         loan.BscLoanId,
				ReleaseAt:         loan.ReleaseAt,
				ReleaseHash:       loan.ReleaseHash,
				ReleaseAmount:     loan.ReleaseAmount,
				PayBackAt:         loan.PayBackAt,
				PayBackHash:       loan.PayBackHash,
				PayBackAmount:     loan.PayBackAmount,
				ReleaseAleoHash:   loan.ReleaseAleoHash,
				ReleaseAleoAt:     loan.ReleaseAleoAt,
				ReleaseAleoAmount: loan.ReleaseAleoAmount,
				Deposits:          deposits,
			})
		}
		success := res.Success(resLoans)
		context.JSON(success.Code, success)
	}
}
