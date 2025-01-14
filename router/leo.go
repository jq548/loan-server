package router

import (
	"github.com/gin-gonic/gin"
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
		rate, _ := config.Rate.BigFloat().Float32()
		releaseRate, _ := config.ReleaseRate.BigFloat().Float32()
		price, _ := config.AleoPrice.BigFloat().Float32()
		resCfg := model.LeoConfig{
			Rate:           rate,
			ReleaseRate:    releaseRate,
			AvailableStage: config.AvailableStages,
			DayPerStage:    config.DayPerStage,
			Price:          price,
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
		as := context.Query("amount")
		if as == "" {
			panic(errors.New(errors.ParameterError))
		}
		amount, err := strconv.Atoi(as)
		if err != nil {
			panic(errors.New(errors.ParameterError))
		}
		config, err := myRouter.mydb.GetConfig()
		if err != nil {
			panic(errors.New(errors.SystemError))
		}
		price, _ := config.AleoPrice.BigFloat().Float64()
		floatAmount := float64(amount) / 1000000
		success := res.Success(price * floatAmount)
		context.JSON(success.Code, success)
	}
}

/*
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"aleo_address":"aleo1hac8kndgfp7eh545yeu6k2ue32yn3dt7qe5xl54d6lpe7xecyq9qkxc3tx", "aleo_amount":1000000, "bsc_address":"0x4332B66D46761476B0A50A2F12EE6a17DaCe7247", "email": "1140830756@qq.com", "stages": 2, "day_per_stage": 7, "type": 1}' \
  http://127.0.0.1:8899/leo/save_deposoit
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
		if config.MinLoanAmount > params.AleoAmount || config.MaxLoanAmount < params.AleoAmount {
			panic(errors.New(errors.ParameterError))
		}
		err = myRouter.mydb.NewDeposit(
			params.AleoAddress,
			params.BscAddress,
			params.Email,
			params.AleoAmount,
			params.Stages,
			params.DayPerStage)
		if err != nil {
			panic(errors.New(errors.SystemError))
		}
		success := res.Success("")
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
		loan, err := myRouter.mydb.SelectLoan(address)
		if err != nil {
			return
		}
		success := res.Success(loan)
		context.JSON(success.Code, success)
	}
}
