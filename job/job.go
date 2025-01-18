package job

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"loan-server/db"
	"loan-server/service"
	"math/rand"
)

type Job struct {
	LeoService *service.LeoChainService
	BscService *service.BscChainService
	Db         *db.MyDb
}

func NewJob(leoService *service.LeoChainService, bscService *service.BscChainService, db *db.MyDb) *Job {
	return &Job{
		LeoService: leoService,
		BscService: bscService,
		Db:         db,
	}
}

func (job *Job) StartJob(spec string, fun func()) {
	c := cron.New(cron.WithSeconds())
	_, err := c.AddFunc(spec, fun)
	if err != nil {
		zap.L().Error("Start Job failed")
	} else {
		c.Start()
	}
}

func (job *Job) StartFetchAleoPrice() {
	zap.S().Info("StartFetchAleoPrice...")
	fakePrice := float64(rand.Int()%500)/1000.0 + 0.5
	err := job.Db.SaveLatestAleoPrice(fakePrice)
	if err != nil {
		zap.S().Error(err)
	}
}

func (job *Job) StartCalculateRate() {
	zap.S().Info("StartCalculateRate...")
	fakePrice := float64(rand.Int()%10) / 10000.0
	err := job.Db.SaveLatestRate(fakePrice)
	if err != nil {
		zap.S().Error(err)
	}
}

func (job *Job) StartCalculateIncome() {
	zap.S().Info("StartCalculateIncome...")

}
