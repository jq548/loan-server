package job

import (
	"github.com/robfig/cron/v3"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"loan-server/common/consts"
	"loan-server/db"
	"loan-server/model"
	"loan-server/service"
	"math/big"
	"math/rand"
	"time"
)

type Job struct {
	LeoService             *service.LeoChainService
	BscService             *service.BscChainService
	Db                     *db.MyDb
	PlatformReceiveAddress string
}

func NewJob(leoService *service.LeoChainService, bscService *service.BscChainService, db *db.MyDb, platformReceiveAddress string) *Job {
	return &Job{
		LeoService:             leoService,
		BscService:             bscService,
		Db:                     db,
		PlatformReceiveAddress: platformReceiveAddress,
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

	provideRecord, err := job.Db.ProvideLiquidRecords()
	if err != nil {
		zap.S().Error(err)
	}
	var totalLiquid decimal.Decimal
	for _, p := range provideRecord {
		totalLiquid = totalLiquid.Add(p.Amount)
	}

	activeLoans, err := job.Db.SelectAllActiveLoans()
	if err != nil {
		zap.S().Error(err)
	}
	var totalLoaned decimal.Decimal
	for _, a := range activeLoans {
		totalLoaned = totalLoaned.Add(a.ReleaseAmount)
	}

	useRate := totalLoaned.Div(totalLiquid)
	cfg, err := job.Db.GetConfig()
	if err != nil {
		zap.S().Error(err)
	}

	now := time.Now()
	beginOfToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix()

	cycleEnv := []float64{1, 0.95, 0.9, 0.85}
	cycleDays := []int{7, 14, 21, 28}
	for i, _ := range cycleEnv {
		rate := cfg.Rate.Mul(decimal.NewFromFloat(cycleEnv[i])).Mul(decimal.NewFromInt(1).Add(useRate))
		if rate.GreaterThan(decimal.NewFromFloat(0.5)) {
			rate = decimal.NewFromFloat(0.5)
		}
		rate = rate.Div(decimal.NewFromInt(52))
		record := model.LeoRateRecord{
			Rate: rate,
			At:   int(beginOfToday),
			Days: cycleDays[i],
		}
		err := job.Db.SaveLatestRate(record)
		if err != nil {
			zap.S().Error(err)
		}
	}
}

func (job *Job) StartCalculateIncome() {
	zap.S().Info("StartCalculateIncome...")
	var addresses []string
	var amounts []*big.Int
	platformIncome, err := job.Db.TotalIncomeLastDay(true)
	if err != nil {
		zap.S().Error(err)
	}
	addresses = append(addresses, job.PlatformReceiveAddress)
	amounts = append(amounts, platformIncome.Mul(decimal.NewFromInt(consts.Wei)).BigInt())
	providerIncome, err := job.Db.TotalIncomeLastDay(false)
	if err != nil {
		zap.S().Error(err)
	}
	provideRecord, err := job.Db.ProvideLiquidRecords()
	if err != nil {
		zap.S().Error(err)
	}
	var totalLiquid decimal.Decimal
	for _, p := range provideRecord {
		totalLiquid = totalLiquid.Add(p.Amount)
	}
	for _, p := range provideRecord {
		amount := p.Amount.Div(totalLiquid).Mul(providerIncome)
		index := indexOf(addresses, p.Provider)
		if index == -1 {
			addresses = append(addresses, p.Provider)
			amounts = append(amounts, amount.BigInt())
		} else {
			amounts[index] = big.NewInt(0).Add(amounts[index], amount.BigInt())
		}
	}

	_, err = job.BscService.IncreaseIncome(addresses, amounts)
	if err != nil {
		zap.S().Error(err)
	}
}

func indexOf(slice []string, target string) int {
	for i, item := range slice {
		if item == target {
			return i
		}
	}
	return -1
}
