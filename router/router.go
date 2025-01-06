package router

import (
	"github.com/gin-gonic/gin"
	"loan-server/config"
	"loan-server/db"
)

type Router struct {
	mydb     *db.MyDb
	MyConfig *config.Config
}

func NewRouter(myDb *db.MyDb, myConfig *config.Config) (*Router, error) {
	router := Router{
		mydb:     myDb,
		MyConfig: myConfig,
	}
	return &router, nil
}

func (myRouter *Router) LoadRouters(engine *gin.Engine) {
	myRouter.loadLeoRouters(engine)
	myRouter.loadBscRouters(engine)
	myRouter.loadAdminRouters(engine)
}
