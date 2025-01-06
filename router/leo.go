package router

import (
	"github.com/gin-gonic/gin"
	"loan-server/common/res"
)

func (myRouter *Router) loadLeoRouters(engine *gin.Engine) {
	engine.GET("/leo/config", loanConfig(myRouter))
	engine.GET("/leo/calculate_usdt", calculateUsdt(myRouter))
	engine.POST("/leo/save_deposoit", saveDeposit(myRouter))
	engine.POST("/leo/complete_deposit", completeDeposit(myRouter))
	engine.GET("/leo/loan_list", leoLoanList(myRouter))
}

// loanConfig
func loanConfig(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}

// calculateUsdt
func calculateUsdt(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}

// saveDeposit
func saveDeposit(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}

// completeDeposit
func completeDeposit(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}

// leoLoanList
func leoLoanList(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}
