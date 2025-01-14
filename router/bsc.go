package router

import (
	"github.com/gin-gonic/gin"
	"loan-server/common/res"
)

func (myRouter *Router) loadBscRouters(engine *gin.Engine) {
	engine.GET("/bsc/config", bscConfig(myRouter))
	engine.GET("/bsc/provide_record", provideRecord(myRouter))
	engine.GET("/bsc/provide_income", provideIncome(myRouter))
	engine.GET("/bsc/provide_income_withdraw_record", provideIncomeReleaseRecord(myRouter))
	engine.GET("/bsc/loan_list", bscLoanList(myRouter))
}

// config of program
func bscConfig(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		myRouter.BscService.CheckAddresses()
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}

// provideRecord
func provideRecord(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}

// provideIncome
func provideIncome(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}

// provideIncomeReleaseRecord
func provideIncomeReleaseRecord(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}

// bscLoanList
func bscLoanList(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}
