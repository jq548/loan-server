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

// 配置信息，参数等
func loanConfig(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}

// 根据aleo计算发放usdt数量
func calculateUsdt(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}

// 保存抵押（发起支付前）
func saveDeposit(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}

// 保存抵押（发起支付后）
func completeDeposit(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}

// 查询抵押
func leoLoanList(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}
