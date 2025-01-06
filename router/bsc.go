package router

import (
	"github.com/gin-gonic/gin"
	"loan-server/common/res"
)

func (myRouter *Router) loadBscRouters(engine *gin.Engine) {
	engine.GET("/bsc/config", bscConfig(myRouter))
	engine.GET("/bsc/provide_record", provideRecord(myRouter))
	engine.GET("/bsc/provide_income", provideIncome(myRouter))
	engine.GET("/bsc/loan_list", bscLoanList(myRouter))
}

// 配置信息，参数等
func bscConfig(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}

// 流动性资金提供记录
func provideRecord(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}

// 流动性资金收益记录
func provideIncome(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}

// 查询我的代款
func bscLoanList(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}
