package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"loan-server/common/consts"
	"loan-server/common/errors"
	"loan-server/common/res"
	"loan-server/common/utils"
	"loan-server/model"
	"strconv"
	"strings"
)

func (myRouter *Router) loadLeoRouters(engine *gin.Engine) {
	engine.GET("/leo/config", loanConfig(myRouter))
	engine.GET("/leo/calculate_usdt", calculateUsdt(myRouter))
	engine.POST("/leo/save_deposoit", saveDeposit(myRouter))
	engine.GET("/leo/loan_list", leoLoanList(myRouter))
	engine.GET("/leo/overview", overview(myRouter))
}

// loanConfig
func loanConfig(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		config, err := myRouter.mydb.GetConfig()
		if err != nil {
			panic(errors.New(errors.SystemError))
		}
		price, err := myRouter.mydb.GetLatestPrice()
		if err != nil {
			panic(errors.New(errors.SystemError))
		}
		rates, err := myRouter.mydb.GetLatestRateOfWeek()
		if err != nil {
			panic(errors.New(errors.SystemError))
		}
		var bannerList []string
		banners, err := myRouter.mydb.GetBanners()
		if err != nil {
			panic(errors.New(errors.SystemError))
		}
		ids := strings.Split(config.BannerIds, ",")
		for _, banner := range banners {
			for _, id := range ids {
				if strconv.Itoa(int(banner.ID)) == id {
					bannerList = append(bannerList, banner.Url)
					break
				}
			}
		}

		resCfg := model.LeoResConfig{
			Rate:           rates[0].Rate.String(),
			AvailableStage: config.AvailableStages,
			DayPerStage:    config.DayPerStage,
			Price:          strconv.FormatFloat(price, 'f', 6, 64),
			AllowTypes:     "1",
			Banners:        bannerList,
			MinAmount:      config.MinLoanAmount / 1000000,
			MaxAmount:      config.MaxLoanAmount / 1000000,
		}
		success := res.Success(resCfg)
		context.JSON(success.Code, success)
	}
}

// calculateUsdt
func calculateUsdt(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		config, err := myRouter.mydb.GetConfig()
		if err != nil {
			panic(errors.New(errors.SystemError))
		}

		as := context.Query("amount")
		if as == "" {
			panic(errors.New(errors.ParameterError))
		}
		amount, err := strconv.ParseFloat(as, 64)
		intAmount := int64(amount * 1000000)
		if err != nil {
			panic(errors.New(errors.ParameterError))
		}
		if config.MinLoanAmount > intAmount || config.MaxLoanAmount < intAmount {
			panic(errors.New(errors.ParameterError))
		}

		price, err := myRouter.mydb.GetLatestPrice()
		if err != nil {
			panic(errors.New(errors.SystemError))
		}
		rates, err := myRouter.mydb.GetLatestRateOfWeek()
		if err != nil {
			panic(errors.New(errors.SystemError))
		}

		collateralAmount := decimal.NewFromFloat(price).Mul(decimal.NewFromFloat(amount))
		borrowAmount := collateralAmount.Mul(config.ReleaseRate)

		var installments []model.LeoResInstallment
		for i := 1; i <= config.AvailableStages; i++ {

			interestRate := decimal.NewFromInt(0)
			for _, rate := range rates {
				if rate.Days == i*config.DayPerStage {
					interestRate = rate.Rate
					break
				}
			}
			installments = append(installments, model.LeoResInstallment{
				Installments:        i,
				DayPerInstallment:   config.DayPerStage,
				InterestRate:        interestRate.String(),
				InterestInstallment: borrowAmount.Mul(interestRate).String(),
			})
		}
		result := model.LeoResCalculateUsdt{
			BorrowingAmount:  borrowAmount.String(),
			CollateralAmount: collateralAmount.String(),
			CollateralRate:   config.ReleaseRate.String(),
			Installment:      installments,
		}

		//usdt := .Mul(decimal.NewFromInt(1).Sub(decimal.NewFromFloat(rate)))
		success := res.Success(result)
		context.JSON(success.Code, success)
	}
}

/*
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"aleo_address":"aleo1hac8kndgfp7eh545yeu6k2ue32yn3dt7qe5xl54d6lpe7xecyq9qkxc3tx", "aleo_amount":2, "bsc_address":"0x4332B66D46761476B0A50A2F12EE6a17DaCe7247", "email": "1140830756@qq.com", "stages": 2, "day_per_stage": 7, "type": 1, "loan_type": 1}' \
  https://wings.tcds.ltd/leo/save_deposoit
*/
// saveDeposit
func saveDeposit(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		config, err := myRouter.mydb.GetConfig()
		if err != nil {
			panic(errors.New(errors.SystemError))
		}
		var params model.ReqSaveDeposit
		if err := context.ShouldBindJSON(&params); err != nil {
			panic(errors.New(errors.ParameterError))
		}
		if !utils.IsValidLeoAddress(params.AleoAddress) {
			panic(errors.New(errors.ParameterError))
		}
		if params.Type < 0 || params.Type > 2 {
			panic(errors.New(errors.ParameterError))
		}
		intAmount := int64(params.AleoAmount * 1000000)

		if params.Type == 0 {
			if params.LoanType != 1 {
				panic(errors.New(errors.ParameterError))
			}
			if config.MinLoanAmount > intAmount || config.MaxLoanAmount < intAmount {
				panic(errors.New(errors.ParameterError))
			}
			if !utils.IsValidAddress(params.BscAddress) {
				panic(errors.New(errors.ParameterError))
			}
			if !utils.VerifyEmailFormat(params.Email) {
				panic(errors.New(errors.ParameterError))
			}
			if config.DayPerStage != params.DayPerStage {
				panic(errors.New(errors.ParameterError))
			}
			if config.AvailableStages < params.Stages || params.Stages < 0 {
				panic(errors.New(errors.ParameterError))
			}
			err = myRouter.mydb.NewDeposit(
				params.AleoAddress,
				params.BscAddress,
				params.Email,
				intAmount,
				params.Stages,
				params.DayPerStage,
				config.ReleaseRate)
			if err != nil {
				panic(errors.New(errors.SystemError))
			}
		} else {
			err = myRouter.mydb.NewDepositOfLoan(params.LoanId, params.AleoAddress, intAmount)
			if err != nil {
				panic(errors.New(errors.SystemError))
			}
		}
		success := res.Success(map[string]bool{
			"success": true,
		})
		context.JSON(success.Code, success)
	}
}

// leoLoanList
func leoLoanList(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		price, err := myRouter.mydb.GetLatestPrice()
		if err != nil {
			panic(errors.New(errors.SystemError))
		}
		address := context.Query("address")
		if address == "" {

			panic(errors.New(errors.ParameterError))
		}
		if !utils.IsValidLeoAddress(address) {
			panic(errors.New(errors.ParameterError))
		}
		loans, err := myRouter.mydb.SelectLoan(address)
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
			ValueWhenDeposit := decimal.Zero
			ValueCurrent := decimal.Zero
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
				ValueWhenDeposit = ValueWhenDeposit.Add(d.UsdtValue.Div(decimal.NewFromInt(consts.Wei)))
				ValueCurrent = ValueWhenDeposit.Add(d.AleoAmount.Div(decimal.NewFromInt(1000000)).Mul(decimal.NewFromFloat(price)))
			}
			minRecharge := decimal.Zero
			if loan.Health.LessThan(decimal.NewFromInt(1)) {
				minRecharge = ValueWhenDeposit.Sub(ValueCurrent).Div(decimal.NewFromInt(consts.Wei)).Div(decimal.NewFromFloat(price))
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
				DepositAmount:     loan.DepositAmount.Div(decimal.NewFromInt(consts.Wei)).String(),
				DepositPrice:      loan.DepositPrice.String(),
				Contract:          loan.Contract,
				ValueCurrent:      ValueCurrent.String(),
				ValueWhenDeposit:  ValueWhenDeposit.String(),
				MinRecharge:       minRecharge.String(),
			})
		}
		success := res.Success(resLoans)
		context.JSON(success.Code, success)
	}
}

// overview
func overview(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		provideRecords, err := myRouter.mydb.ProvideLiquidRecords()
		if err != nil {
			panic(errors.New(errors.SystemError))
		}
		var totalLiquid decimal.Decimal
		for _, p := range provideRecords {
			totalLiquid = totalLiquid.Add(p.Amount)
		}

		activeLoans, err := myRouter.mydb.SelectAllActiveLoans()
		if err != nil {
			panic(errors.New(errors.SystemError))
		}
		var totalLoaned decimal.Decimal
		for _, a := range activeLoans {
			totalLoaned = totalLoaned.Add(a.ReleaseAmount)
		}
		useRate := decimal.Zero
		if !totalLiquid.Equal(decimal.Zero) {
			useRate = totalLoaned.Div(totalLiquid)
		}

		totalDeposits := decimal.Zero
		deposits, err := myRouter.mydb.SelectAllDepositsOfActiveLoans()
		if err != nil {
			panic(errors.New(errors.SystemError))
		}
		for _, d := range deposits {
			totalDeposits = totalDeposits.Add(d.AleoAmount)
		}

		var historyRate []model.LeoOverViewRateOfWeek
		rates, err := myRouter.mydb.SelectHistoryOfRateOf1Week()
		for _, r := range rates {
			historyRate = append(historyRate, model.LeoOverViewRateOfWeek{
				Rate: r.Rate,
				At:   r.At,
				Days: r.Days,
			})
		}
		// banner
		config, err := myRouter.mydb.GetConfig()
		if err != nil {
			panic(errors.New(errors.SystemError))
		}
		var bannerList []string
		banners, err := myRouter.mydb.GetBanners()
		if err != nil {
			panic(errors.New(errors.SystemError))
		}
		ids := strings.Split(config.BannerIds, ",")
		for _, banner := range banners {
			for _, id := range ids {
				if strconv.Itoa(int(banner.ID)) == id {
					bannerList = append(bannerList, banner.Url)
					break
				}
			}
		}
		// calculate income rate

		incomeRates, err := myRouter.mydb.SelectRecentRateOfProvideLiquidIncome()
		if err != nil {
			panic(errors.New(errors.SystemError))
		}
		rate := decimal.Zero
		if len(incomeRates) > 0 {
			rate = incomeRates[0].Rate
		}
		rate = rate.Div(decimal.NewFromFloat(365))
		estimateRate := rate.Mul(decimal.NewFromFloat(1.6)).RoundDown(6)
		overviewData := model.LeoOverView{
			TotalProvideLiquid:      totalLiquid.Div(decimal.NewFromInt(consts.Wei)).String(),
			TotalLoaned:             totalLoaned.Div(decimal.NewFromInt(consts.Wei)).String(),
			LiquidUsedRate:          useRate.RoundDown(4).String(),
			ProvideLiquidRewardRate: estimateRate.String(),
			TotalDepositAleo:        totalDeposits.Div(decimal.NewFromInt(1000000)).String(),
			HistoryRate:             historyRate,
			Banners:                 bannerList,
		}
		success := res.Success(overviewData)
		context.JSON(success.Code, success)
	}
}
