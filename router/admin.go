package router

import (
	"github.com/gin-gonic/gin"
	"loan-server/common/res"
)

func (myRouter *Router) loadAdminRouters(engine *gin.Engine) {
	engine.POST("/baked/login", login(myRouter))
	engine.POST("/baked/change_password", changePassword(myRouter))
	engine.POST("/baked/add_account", addAccount(myRouter))
	engine.GET("/baked/login", getConfig(myRouter))
	engine.POST("/baked/save_config", saveConfig(myRouter))
}

// 配置信息，参数等
func login(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}

// 修改密码
func changePassword(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}

// 添加管理账号
func addAccount(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}

// 查询配置参数
func getConfig(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}

// 保存配置参数
func saveConfig(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}
