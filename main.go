package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"loan-server/common/logger"
	"loan-server/config"
	"loan-server/db"
	"loan-server/handler"
	"loan-server/job"
	routers "loan-server/router"
	"loan-server/service"
	"log"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = logger.Init(cfg.Log.Level)
	if err != nil {
		log.Fatal(err)
	}

	myDb, err := db.Init(&cfg.Db)
	if err != nil {
		log.Fatal(err)
	}

	ls := service.NewLeoChainService(&cfg.Leo, myDb)

	bs, err := service.NewBscChainService(&cfg.Bsc, myDb)
	if err != nil {
		log.Fatal(err)
	}
	ls.BscService = bs
	bs.LeoService = ls
	go ls.Start()
	go bs.StartFetchEvent()

	// ---- start job ----
	myJob := job.NewJob(ls, bs, myDb, cfg.Platform.ReceiveAddress)
	myJob.StartJob(cfg.Job.AleoPrice, myJob.StartFetchAleoPrice)
	myJob.StartJob(cfg.Job.CalculateRate, myJob.StartCalculateRate)
	myJob.StartJob(cfg.Job.CalculateIncome, myJob.StartCalculateIncome)
	myJob.StartJob(cfg.Job.CalculateIncome, myJob.StartCalculateIncomeRate)

	ginEngine := gin.Default()
	gin.SetMode(gin.DebugMode)
	ginEngine.NoRoute(handler.HandleNotFound)
	ginEngine.NoMethod(handler.HandleNotFound)
	ginEngine.Use(handler.GinLogger(), handler.GinRecovery(true), handler.Cors())
	// load routers
	if router, err := routers.NewRouter(myDb, cfg, ls, bs); err != nil {
		log.Fatal(err)
	} else {
		router.LoadRouters(ginEngine)
	}

	// run
	err = ginEngine.Run(fmt.Sprintf(":%d", cfg.Service.Port))
	if err != nil {
		log.Fatal(err)
	}
}
