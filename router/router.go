package router

import (
	"github.com/gin-gonic/gin"
	"loan-server/config"
	"loan-server/db"
	"loan-server/service"
)

type Router struct {
	mydb       *db.MyDb
	MyConfig   *config.Config
	LeoService *service.LeoChainService
	BscService *service.BscChainService
}

func NewRouter(
	myDb *db.MyDb,
	myConfig *config.Config,
	leoService *service.LeoChainService,
	bscService *service.BscChainService) (*Router, error) {
	router := Router{
		mydb:       myDb,
		MyConfig:   myConfig,
		LeoService: leoService,
		BscService: bscService,
	}
	return &router, nil
}

func (myRouter *Router) LoadRouters(engine *gin.Engine) {
	myRouter.loadLeoRouters(engine)
	myRouter.loadBscRouters(engine)
	myRouter.loadAdminRouters(engine)
}
