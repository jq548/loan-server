package job

import (
	"encoding/json"
	"github.com/robfig/cron/v3"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"loan-server/common/consts"
	"loan-server/common/utils"
	"loan-server/config"
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
	MailConfig             *config.GoMail
}

func NewJob(
	leoService *service.LeoChainService,
	bscService *service.BscChainService,
	db *db.MyDb,
	platformReceiveAddress string,
	mail *config.GoMail) *Job {
	return &Job{
		LeoService:             leoService,
		BscService:             bscService,
		Db:                     db,
		PlatformReceiveAddress: platformReceiveAddress,
		MailConfig:             mail,
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
	err = job.updateLoanHealth()
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
	var idsInt []int
	var addresses []string
	var amounts []*big.Int
	var amountsStr []string
	platformIncome, err := job.Db.TotalIncomeLastDay(true)
	if err != nil {
		zap.S().Error(err)
	}
	if platformIncome.GreaterThan(decimal.Zero) {
		ids = append(ids, big.NewInt(0))
		idsInt = append(idsInt, 0)
		addresses = append(addresses, job.PlatformReceiveAddress)
		amounts = append(amounts, platformIncome.BigInt())
		amountsStr = append(amountsStr, platformIncome.String())
	}
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
		if amount.GreaterThan(decimal.Zero) {
			ids = append(ids, big.NewInt(int64(p.RecordId)))
			idsInt = append(idsInt, p.RecordId)
			addresses = append(addresses, p.Provider)
			amounts = append(amounts, amount.BigInt())
			amountsStr = append(amountsStr, amount.String())
		}
	}

	if len(ids) > 0 {
		status := 1
		hash, err := job.BscService.IncreaseIncome(ids, addresses, amounts)
		if err != nil {
			zap.S().Error(err)
			status = 2
		}
		idsjson, err := json.Marshal(ids)
		if err != nil {
			zap.S().Error(err)
		}
		addressesjson, err := json.Marshal(addresses)
		if err != nil {
			zap.S().Error(err)
		}
		amountsjson, err := json.Marshal(amountsStr)
		if err != nil {
			zap.S().Error(err)
		}
		err = job.Db.NewIncomeGenerateRecord(
			int(time.Now().Unix()),
			status,
			hash,
			string(idsjson),
			string(addressesjson),
			string(amountsjson))
		if err != nil {
			zap.S().Error(err)
		}
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

func (job *Job) StartCheckExpiredLoans() {
	zap.S().Info("StartCheckExpiredLoans...")
	loans, err := job.Db.FindLoanByStatus(2)
	if err != nil {
		zap.S().Error(err)
		return
	}
	for _, loan := range loans {
		if int64(loan.StartAt+loan.Stages*loan.DayPerStage)+24*3600 < time.Now().Unix() {
			status := loan.Status
			err := job.BscService.ClearLoanInContract(big.NewInt(int64(loan.ID)))
			if err != nil {
				zap.S().Error(err)
				status = 7
			} else {
				status = 5
			}
			err = job.Db.SaveStatusOfLoan(int(loan.ID), status)
			if err != nil {
				zap.S().Error(err)
			}
		}
	}
}

func (job *Job) CheckFailedJobs() {
	zap.S().Info("CheckFailedJobs...")
	// income
	job.checkIncreaseIncome()
	// loan
	job.checkCreateLoan()
	// clear contract
	job.checkClearLoanOnContract()
	// clear sold
	job.checkClearLoanSold()
	// clear finish
	job.checkClearLoanFinish()
}

func (job *Job) updateLoanHealth() error {
	price, err := job.Db.GetLatestPrice()
	if err != nil {
		return err
	}
	loans, err := job.Db.SelectAllActiveLoans()
	if err != nil {
		return err
	}
	for _, loan := range loans {
		orgValue := decimal.Zero
		aleoAmount := decimal.Zero
		deposits, err := job.Db.SelectDepositByLoanId(int(loan.ID))
		if err != nil {
			return err
		}
		for _, deposit := range deposits {
			orgValue = orgValue.Add(deposit.UsdtValue)
			aleoAmount = aleoAmount.Add(deposit.AleoAmount)
		}
		nowValue := decimal.NewFromFloat(price).Mul(aleoAmount.Div(decimal.NewFromInt(1000000))).Mul(decimal.NewFromInt(consts.Wei))
		loan.Health = nowValue.Div(orgValue)
		err = job.Db.SaveHealthOfLoan(int(loan.ID), loan.Health)
		if err != nil {
			return err
		}
		if loan.Health.LessThan(decimal.NewFromFloat(0.8)) && loan.Health.GreaterThan(decimal.NewFromFloat(0.7)) {
			job.sendWarningEmail(loan)
		}
		if loan.Health.LessThan(decimal.NewFromFloat(0.7)) {
			status := 5
			err := job.BscService.ClearLoanInContract(big.NewInt(int64(loan.ID)))
			if err != nil {
				zap.S().Error(err)
				status = 7
			}
			err = job.Db.SaveStatusOfLoan(int(loan.ID), status)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (job *Job) checkIncreaseIncome() {
	failedSaveRecentRateIncome, err := job.Db.SelectFailedIncomeGenerateRecord()
	if err != nil {
		zap.S().Error(err)
	}
	for _, record := range failedSaveRecentRateIncome {
		var ids []*big.Int
		var idsInt []int
		var addresses []string
		var amounts []*big.Int
		var amountsStr []string
		err := json.Unmarshal([]byte(record.Ids), &idsInt)
		if err != nil {
			zap.S().Error(err)
			continue
		}
		for _, id := range idsInt {
			ids = append(ids, big.NewInt(int64(id)))
		}
		err = json.Unmarshal([]byte(record.Addresses), &addresses)
		if err != nil {
			zap.S().Error(err)
			continue
		}
		err = json.Unmarshal([]byte(record.Amounts), &amountsStr)
		if err != nil {
			zap.S().Error(err)
			continue
		}

		for _, amount := range amountsStr {
			tmp := new(big.Int)
			a, success := tmp.SetString(amount, 10)
			if !success {
				continue
			}
			amounts = append(amounts, a)
		}

		hash, err := job.BscService.IncreaseIncome(ids, addresses, amounts)
		if err != nil {
			zap.S().Error(err)
			continue
		}
		err = job.Db.CompleteIncomeGenerateRecord(int(record.ID), hash)
		if err != nil {
			zap.S().Error(err)
			continue
		}
	}
}

func (job *Job) checkCreateLoan() {
	loans, err := job.Db.FindLoanByStatus(6)
	if err != nil {
		zap.S().Error(err)
		return
	}
	for _, loan := range loans {
		err := job.BscService.CreateLoanInContract(
			big.NewInt(int64(loan.ID)),
			loan.LoanAmount.BigInt(),
			big.NewInt(int64(loan.Stages*loan.DayPerStage*24*3600)),
			loan.InterestAmount.BigInt(),
			loan.BscAddress,
			loan.AleoAddress,
			loan.DepositAmount.BigInt(),
			loan.DepositPrice.BigInt(),
		)
		if err != nil {
			zap.S().Error(err)
			err = job.Db.SaveCreateFailed(int(loan.ID), 6)
			if err != nil {
				zap.S().Error(err)
			}
		}
	}
}

func (job *Job) checkClearLoanOnContract() {
	loans, err := job.Db.FindLoanByStatus(7)
	if err != nil {
		zap.S().Error(err)
		return
	}
	for _, loan := range loans {
		status := loan.Status
		err := job.BscService.ClearLoanInContract(big.NewInt(int64(loan.ID)))
		if err != nil {
			zap.S().Error(err)
			status = 7
		} else {
			status = 5
		}
		err = job.Db.SaveStatusOfLoan(int(loan.ID), status)
		if err != nil {
			zap.S().Error(err)
		}
	}
}

func (job *Job) checkClearLoanSold() {
	loans, err := job.Db.FindLoanByStatus(8)
	if err != nil {
		zap.S().Error(err)
		return
	}
	for _, loan := range loans {
		job.BscService.ExecClearSold(loan.ID)
	}
}

func (job *Job) checkClearLoanFinish() {
	loans, err := job.Db.FindLoanByStatus(9)
	if err != nil {
		zap.S().Error(err)
		return
	}
	for _, loan := range loans {
		job.BscService.ExecClearCalculateIncome(loan.ID)
	}
}

func (job *Job) sendWarningEmail(loan model.Loan) {
	if len(loan.Email) == 0 {
		return
	}
	record, err := job.Db.FindSendEmailRecord(loan.ID)
	if err != nil {
		zap.S().Error(err)
		return
	}
	if len(record) == 0 {
		success, err := utils.SendEmail(false, job.MailConfig.Email, job.MailConfig.Account, job.MailConfig.Password, "smtp.gmail.com", 587, loan.Email, "Your loan health is lower than 80%.")
		if err != nil {
			success = false
		}
		err = job.Db.SaveSendEmailRecord(loan.ID, success, loan.Email)
		if err != nil {
			zap.S().Error(err)
			return
		}
	}
}
