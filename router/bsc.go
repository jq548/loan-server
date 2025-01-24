package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"loan-server/common/consts"
	"loan-server/common/errors"
	"loan-server/common/res"
	"loan-server/common/utils"
	"loan-server/model"
	"time"
)

func (myRouter *Router) loadBscRouters(engine *gin.Engine) {
	engine.GET("/bsc/config", bscConfig(myRouter))
	engine.GET("/bsc/loan_list", bscLoanList(myRouter))
	engine.GET("/bsc/provide_record", provideRecord(myRouter))
	engine.GET("/bsc/provide_income_withdraw_record", provideIncomeReleaseRecord(myRouter))
	engine.GET("/bsc/exchange_record", exchangeRecord(myRouter))
}

// config of program
func bscConfig(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		var provideConfigs []model.BscConfigProvideLiquid
		days := []int{
			7, 30, 60, 90, 180,
		}
		factor := []float32{
			0.8, 1, 1.2, 1.4, 1.6,
		}
		rates, err := myRouter.mydb.SelectRecentRateOfProvideLiquidIncome()
		if err != nil {
			panic(errors.New(errors.SystemError))
		}
		rate := decimal.Zero
		if len(rates) > 0 {
			rate = rates[0].Rate
		}
		rate = rate.Div(decimal.NewFromFloat(365))
		for i, d := range days {
			provideConfigs = append(provideConfigs, model.BscConfigProvideLiquid{
				Duration:         3600 * 24 * d,
				Days:             d,
				EstimateWeekRate: float32(rate.Mul(decimal.NewFromFloat(float64(factor[i]))).RoundDown(6).InexactFloat64()),
			})
		}
		cfg := model.BscConfig{
			WithdrawIncomeFee:  "1",
			WithdrawProvideFee: "2",
			MinProvideAmount:   "100",
			MaxProvideAmount:   "10000",
			ProvideLiquid:      provideConfigs,
		}
		success := res.Success(cfg)
		context.JSON(success.Code, success)
	}
}
func getFactorOfDay(day int) float64 {
	switch day {
	case 7:
		return 0.8
	case 30:
		return 1
	case 60:
		return 1.2
	case 90:
		return 1.4
	case 180:
		return 1.6
	default:
		return 1
	}
}

func provideRecord(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		address := context.Query("address")
		if address == "" {
			panic(errors.New(errors.ParameterError))
		}
		if !utils.IsValidAddress(address) {
			panic(errors.New(errors.ParameterError))
		}
		rates, err := myRouter.mydb.SelectRecentRateOfProvideLiquidIncome()
		if err != nil {
			panic(errors.New(errors.SystemError))
		}
		rate := decimal.Zero
		if len(rates) > 0 {
			rate = rates[0].Rate
		}

		totalProvide := decimal.NewFromInt(0)
		income30 := decimal.NewFromInt(0)
		incomeYesterday := decimal.NewFromInt(0)
		var records []model.BscProvideRecord
		incomes, err := myRouter.mydb.SelectProvideIncome(address)
		if err != nil {
			panic(errors.New(errors.SystemError))
		}
		now := time.Now()
		beginOfToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix()
		for _, income := range incomes {
			if income.Status == 0 {
				totalProvide = totalProvide.Add(income.Amount)
			}
			if income.CreateAt < int(beginOfToday) && income.CreateAt > int(beginOfToday)-3600*24*30 {
				income30 = income30.Add(income.IncomeAmount)
			}
			if income.CreateAt < int(beginOfToday) && income.CreateAt > int(beginOfToday)-3600*24 {
				incomeYesterday = incomeYesterday.Add(income.Amount)
			}
			index := -1
			for i, record := range records {
				if record.RecordId == income.RecordId {
					index = i
					break
				}
			}
			if index == -1 {
				newRecord := model.BscProvideRecord{
					Days:           income.Duration / 24 / 3600,
					Amount:         income.Amount.Div(decimal.NewFromInt(consts.Wei)).String(),
					RateYear:       decimal.NewFromFloat(getFactorOfDay(income.Duration / 24 / 3600)).Mul(rate).String(),
					TotalIncomeDec: income.IncomeAmount.Div(decimal.NewFromInt(consts.Wei)),
					Duration:       income.Duration,
					Start:          income.Start,
					Status:         income.Status,
					Provider:       income.Provider,
					CreateAt:       income.CreateAt,
					CreateHash:     income.CreateHash,
					RetrieveAt:     income.RetrieveAt,
					RetrieveHash:   income.RetrieveHash,
					RecordId:       income.RecordId,
				}
				if income.CreateAt < int(beginOfToday) && income.CreateAt > int(beginOfToday)-3600*24 {
					newRecord.YesterdayIncomeDec = income.IncomeAmount.Div(decimal.NewFromInt(consts.Wei))
				}
				records = append(records, newRecord)
			} else {
				records[index].TotalIncomeDec = records[index].TotalIncomeDec.Add(income.IncomeAmount.Div(decimal.NewFromInt(consts.Wei)))
				if income.CreateAt < int(beginOfToday) && income.CreateAt > int(beginOfToday)-3600*24 {
					records[index].YesterdayIncomeDec = records[index].YesterdayIncomeDec.Add(income.IncomeAmount.Div(decimal.NewFromInt(consts.Wei)))
				}
			}
		}
		for i, _ := range records {
			records[i].TotalIncome = records[i].TotalIncomeDec.String()
			records[i].YesterdayIncome = records[i].YesterdayIncomeDec.String()
			createAt := time.Unix(int64(records[i].CreateAt), 0)
			records[i].IncomeStartDay = createAt.Format("2006-01-02")
			if records[i].Status == 1 {
				endAt := time.Unix(int64(records[i].RetrieveAt), 0)
				records[i].IncomeEndDay = endAt.Format("2006-01-02")
			} else {
				records[i].IncomeEndDay = "-"
			}
		}
		BscProvideInfo := model.BscProvideInfo{
			TotalProvide:    totalProvide.Div(decimal.NewFromInt(consts.Wei)).String(),
			Income30:        income30.Div(decimal.NewFromInt(consts.Wei)).String(),
			IncomeYesterday: incomeYesterday.Div(decimal.NewFromInt(consts.Wei)).String(),
			ProvideRecord:   records,
		}
		success := res.Success(BscProvideInfo)
		context.JSON(success.Code, success)
	}
}

func provideIncomeReleaseRecord(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		address := context.Query("address")
		if address == "" {
			panic(errors.New(errors.ParameterError))
		}
		if !utils.IsValidAddress(address) {
			panic(errors.New(errors.ParameterError))
		}
		var records []model.BscProvideRewardRecord
		incomes, err := myRouter.mydb.SelectProvideIncomeWithdrawRecord(address)
		if err != nil {
			panic(errors.New(errors.SystemError))
		}
		for _, income := range incomes {
			records = append(records, model.BscProvideRewardRecord{
				Provider: income.Provider,
				Amount:   income.Amount.Div(decimal.NewFromInt(consts.Wei)).String(),
				Hash:     income.Hash,
				At:       income.At,
			})
		}
		success := res.Success(records)
		context.JSON(success.Code, success)
	}
}

func exchangeRecord(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		address := context.Query("address")
		if address == "" {
			panic(errors.New(errors.ParameterError))
		}
		if !utils.IsValidAddress(address) {
			panic(errors.New(errors.ParameterError))
		}
		var result []model.BscExchangeRecord
		records, err := myRouter.mydb.SelectExchangeRecordByAddress(address)
		if err != nil {
			panic(errors.New(errors.SystemError))
		}
		for _, record := range records {
			result = append(result, model.BscExchangeRecord{
				Type:    record.Type,
				Amount:  record.Amount.Div(decimal.NewFromInt(consts.Wei)).String(),
				Address: record.Address,
				At:      record.At,
				Hash:    record.Hash,
			})
		}
		success := res.Success(result)
		context.JSON(success.Code, success)
	}
}

func bscLoanList(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		address := context.Query("address")
		if address == "" {

			panic(errors.New(errors.ParameterError))
		}
		if !utils.IsValidAddress(address) {
			panic(errors.New(errors.ParameterError))
		}
		loans, err := myRouter.mydb.SelectLoanByETHAddress(address)
		if err != nil {
			return
		}
		var resLoans []model.LeoResLoan
		for _, loan := range loans {
			var deposits []model.LeoResDeposit
			ds, err := myRouter.mydb.SelectDepositByLoanId(int(loan.ID))
			if err != nil {
				return
			}
			for _, d := range ds {
				deposits = append(deposits, model.LeoResDeposit{
					ID:          int(d.ID),
					LoanId:      d.LoanId,
					AleoAddress: d.AleoAddress,
					AleoAmount:  d.AleoAmount.String(),
					AleoPrice:   d.AleoPrice.String(),
					UsdtValue:   d.UsdtValue.Div(decimal.NewFromInt(consts.Wei)).String(),
					Hash:        d.Hash,
					At:          d.At,
					Status:      d.Status,
				})
			}
			resLoans = append(resLoans, model.LeoResLoan{
				ID:                int(loan.ID),
				AleoAddress:       loan.AleoAddress,
				BscAddress:        loan.BscAddress,
				Email:             loan.Email,
				Status:            loan.Status,
				Stages:            loan.Stages,
				DayPerStage:       loan.DayPerStage,
				StartAt:           loan.StartAt,
				Health:            loan.Health.String(),
				Rate:              loan.Rate.String(),
				ReleaseRate:       loan.ReleaseRate.String(),
				Hash:              loan.Hash,
				Type:              loan.Type,
				BscLoanId:         loan.BscLoanId,
				ReleaseAt:         loan.ReleaseAt,
				ReleaseHash:       loan.ReleaseHash,
				ReleaseAmount:     loan.ReleaseAmount.Div(decimal.NewFromInt(consts.Wei)).String(),
				PayBackAt:         loan.PayBackAt,
				PayBackHash:       loan.PayBackHash,
				PayBackAmount:     loan.PayBackAmount.String(),
				ReleaseAleoHash:   loan.ReleaseAleoHash,
				ReleaseAleoAt:     loan.ReleaseAleoAt,
				ReleaseAleoAmount: loan.ReleaseAleoAmount.Div(decimal.NewFromInt(consts.Wei)).String(),
				Deposits:          deposits,
			})
		}
		success := res.Success(resLoans)
		context.JSON(success.Code, success)
	}
}
