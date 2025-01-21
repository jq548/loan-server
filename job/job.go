package job

import (
	"github.com/robfig/cron/v3"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
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
	fakePrice := float64(rand.Int()%500)/20.0 + 100
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
	useRate := decimal.Zero
	if !totalLiquid.Equal(decimal.Zero) {
		useRate = totalLoaned.Div(totalLiquid)
	}
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
	var ids []*big.Int
	var addresses []string
	var amounts []*big.Int
	platformIncome, err := job.Db.TotalIncomeLastDay(true)
	if err != nil {
		zap.S().Error(err)
	}
	ids = append(ids, big.NewInt(0))
	addresses = append(addresses, job.PlatformReceiveAddress)
	amounts = append(amounts, platformIncome.BigInt())
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
		amount := decimal.Zero
		if !totalLiquid.Equal(decimal.Zero) {
			amount = p.Amount.Div(totalLiquid).Mul(providerIncome)
		}
		ids = append(ids, big.NewInt(int64(p.RecordId)))
		addresses = append(addresses, p.Provider)
		amounts = append(amounts, amount.BigInt())
	}

	_, err = job.BscService.IncreaseIncome(ids, addresses, amounts)
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

func (job *Job) StartCalculateIncomeRate() {
	zap.S().Info("StartCalculateIncomeRate...")
	providerIncome, days, err := job.Db.TotalIncome30Day()
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
	rate := decimal.Zero
	if totalLiquid != decimal.Zero {
		rate = providerIncome.Div(totalLiquid).Mul(decimal.NewFromInt(365)).Div(decimal.NewFromInt(int64(days))).Mul(decimal.NewFromInt(100))
	}
	err = job.Db.SaveRecentRateOfProvideLiquidIncome(rate, days)
	if err != nil {
		zap.S().Error(err)
	}
}
