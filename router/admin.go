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

// admin login
func login(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}

// change admin password
func changePassword(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}

// add admin account
func addAccount(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}

// get config
func getConfig(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}

// save config
func saveConfig(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		success := res.Success("")
		context.JSON(success.Code, success)
	}
}
