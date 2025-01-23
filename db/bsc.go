package db

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"loan-server/common/consts"
	errors2 "loan-server/common/errors"
	"loan-server/model"
	"strconv"
	"time"
)

func (m *MyDb) GetBscBlockHeight() (int64, error) {

	cache, err := m.FindCacheByKey(consts.BscBlockHeightKey)
	if err != nil {
		return 0, err
	}
	if cache.CacheValue == "" {
		cache.CacheValue = "0"
	}
	blockNum, err := strconv.ParseInt(cache.CacheValue, 10, 64)
	if err != nil {
		return 0, err
	}
	return blockNum, nil
}

func (m *MyDb) SaveBscBlockHeight(height int) error {
	res, err := m.UpdateCache(consts.BscBlockHeightKey, strconv.Itoa(height))
	if err != nil {
		return err
	}
	if !res {
		return errors2.New(errors2.SystemError)
	}
	return nil
}

func (m *MyDb) SaveDepositOnBscHash(
	hash,
	loaner string,
	loanId,
	duration,
	start,
	at int,
	releaseAmount, interestAmount decimal.Decimal) error {
	var count int64
	tx := m.Db.Model(&model.Loan{}).Where(&model.Loan{
		ReleaseHash: hash,
	}).Count(&count)
	if tx.Error != nil {
		return tx.Error
	}
	if count > 0 {
		return nil
	}
	var loan model.Loan
	var selector = model.Loan{}
	selector.ID = uint(loanId)
	tx = m.Db.Model(&model.Loan{}).Where(&selector).Last(&loan)
	if tx.Error != nil {
		return tx.Error
	}

	loan.BscLoanId = loanId
	loan.ReleaseAt = at
	loan.ReleaseHash = hash
	loan.ReleaseAmount = releaseAmount
	loan.Status = 2

	tx = m.Db.Save(&loan)
	if tx.Error != nil {
		return tx.Error
	}
	tx = m.Db.Create(&model.IncomeRecord{
		Type:       1,
		Amount:     interestAmount,
		At:         start,
		IsNegative: 0,
		SplitDays:  duration / 24 / 3600,
		EndAt:      start + duration,
		Hash:       hash,
	})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (m *MyDb) Payback(
	loanId int,
	hash string,
	at int,
	amount decimal.Decimal) error {
	var count int64
	tx := m.Db.Model(&model.Loan{}).Where(&model.Loan{
		PayBackHash: hash,
	}).Count(&count)
	if tx.Error != nil {
		return tx.Error
	}
	if count > 0 {
		return nil
	}
	var loan model.Loan
	tx = m.Db.Model(&model.Loan{}).Where(&model.Loan{
		BscLoanId: loanId,
	}).Last(&loan)
	if tx.Error != nil {
		return gorm.ErrRecordNotFound
	}

	loan.PayBackAt = at
	loan.PayBackHash = hash
	loan.PayBackAmount = amount
	loan.Status = 4

	tx = m.Db.Save(&loan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (m *MyDb) Clear(
	loanId int,
	hash string,
	at int,
	amount decimal.Decimal) error {
	var count int64
	tx := m.Db.Model(&model.Loan{}).Where(&model.Loan{
		PayBackHash: hash,
	}).Count(&count)
	if tx.Error != nil {
		return tx.Error
	}
	if count > 0 {
		return nil
	}
	var loan model.Loan
	tx = m.Db.Model(&model.Loan{}).Where(&model.Loan{
		BscLoanId: loanId,
	}).Last(&loan)
	if tx.Error != nil {
		return gorm.ErrRecordNotFound
	}

	loan.PayBackAt = at
	loan.PayBackHash = hash
	loan.PayBackAmount = amount
	loan.Status = 4

	tx = m.Db.Save(&loan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (m *MyDb) IncreaseProviderRewardAmount(
	amount decimal.Decimal,
	address string,
	hash string,
	at,
	recordId int) error {

	var rec []model.ProvideRewardRecord
	tx := m.Db.Model(&model.ProvideRewardRecord{}).Where(&model.ProvideRewardRecord{Hash: hash, Type: 0}).Find(&rec)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			record := model.ProvideRewardRecord{
				Type:     0,
				Provider: address,
				Amount:   amount,
				At:       at,
				Hash:     hash,
				RecordId: recordId,
			}
			tx := m.Db.Create(&record)
			if tx.Error != nil {
				return tx.Error
			}
			return nil
		}
		return tx.Error
	}
	return nil
}

func (m *MyDb) ReleaseProviderReward(
	amount decimal.Decimal,
	address string,
	hash string,
	at int,
	fee decimal.Decimal) error {
	var rec []model.ProvideRewardRecord
	tx := m.Db.Model(&model.ProvideRewardRecord{}).Where(&model.ProvideRewardRecord{Hash: hash, Type: 1}).Find(&rec)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			record := model.ProvideRewardRecord{
				Type:     1,
				Provider: address,
				Amount:   amount,
				At:       at,
				Hash:     hash,
			}
			tx := m.Db.Create(&record)
			if tx.Error != nil {
				return tx.Error
			}
		} else {
			return tx.Error
		}
	}
	if len(rec) > 0 {
		return nil
	}
	tx = m.Db.Create(&model.IncomeRecord{
		Type:       2,
		Amount:     fee,
		At:         at,
		IsNegative: 0,
		SplitDays:  0,
		EndAt:      at,
		Hash:       hash,
	})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (m *MyDb) IncreaseProviderAmount(
	amount decimal.Decimal,
	start, duration int,
	address string,
	hash string,
	at, recordId int) error {
	var count int64
	tx := m.Db.Model(&model.ProvideLiquid{}).Where(&model.ProvideLiquid{
		CreateHash: hash,
	}).Count(&count)
	if tx.Error != nil {
		return tx.Error
	}
	if count > 0 {
		return nil
	}
	tx = m.Db.Create(&model.ProvideLiquid{
		Amount:     amount,
		Duration:   duration,
		Start:      start,
		Status:     0,
		Provider:   address,
		CreateAt:   at,
		CreateHash: hash,
		RecordId:   recordId,
	})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (m *MyDb) RetrieveProviderAmount(
	id int,
	hash string,
	at int,
	fee decimal.Decimal) error {
	var count int64
	tx := m.Db.Model(&model.ProvideLiquid{}).Where(&model.ProvideLiquid{
		RetrieveHash: hash,
	}).Count(&count)
	if tx.Error != nil {
		return tx.Error
	}
	if count > 0 {
		return nil
	}
	var pl model.ProvideLiquid
	var selector = model.ProvideLiquid{}
	selector.ID = uint(id)
	tx = m.Db.Model(&model.ProvideLiquid{}).Where(&selector).First(&pl)
	if tx.Error != nil {
		return tx.Error
	}
	pl.Status = 1
	pl.RetrieveAt = at
	pl.RetrieveHash = hash
	tx = m.Db.Save(&pl)
	if tx.Error != nil {
		return tx.Error
	}
	if fee.GreaterThan(decimal.Zero) {
		tx := m.Db.Create(&model.IncomeRecord{
			Type:       3,
			Amount:     fee,
			At:         at,
			IsNegative: 0,
			SplitDays:  0,
			EndAt:      at,
			Hash:       hash,
		})
		if tx.Error != nil {
			return tx.Error
		}
	}
	return nil
}

func (m *MyDb) TotalIncomeLastDay(isPlatform bool) (decimal.Decimal, error) {
	var result decimal.Decimal
	config, err := m.GetConfig()
	if err != nil {
		return result, err
	}
	now := time.Now()
	beginOfToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix()
	beginOfYesterday := beginOfToday - 3600*24
	sqls := fmt.Sprintf("SELECT * from income_record WHERE type=1 AND ((at>%d AND at<%d) OR (end_at>%d AND end_at<%d) OR (at<%d AND end_at>%d))", beginOfYesterday, beginOfToday, beginOfYesterday, beginOfToday, beginOfYesterday, beginOfToday)
	var interestRecords []model.IncomeRecord
	tx := m.Db.Raw(sqls).Scan(&interestRecords)
	if tx.Error != nil {
		return result, tx.Error
	}
	sqls = fmt.Sprintf("SELECT * from income_record WHERE type=4 AND (at>%d AND at<%d)", beginOfYesterday, beginOfToday)
	var clearRecords []model.IncomeRecord
	tx = m.Db.Raw(sqls).Scan(&clearRecords)
	if tx.Error != nil {
		return result, tx.Error
	}
	if isPlatform {
		sqls = fmt.Sprintf("SELECT * from income_record WHERE type=2 AND (at>%d AND at<%d)", beginOfYesterday, beginOfToday)
		var providerWithdrawRecords []model.IncomeRecord
		tx = m.Db.Raw(sqls).Scan(&providerWithdrawRecords)
		if tx.Error != nil {
			return result, tx.Error
		}
		sqls = fmt.Sprintf("SELECT * from income_record WHERE type=3 AND (at>%d AND at<%d)", beginOfYesterday, beginOfToday)
		var providerRetrieveRecords []model.IncomeRecord
		tx = m.Db.Raw(sqls).Scan(&providerRetrieveRecords)
		if tx.Error != nil {
			return result, tx.Error
		}
		for _, record := range interestRecords {
			result = result.Add(record.Amount.Mul(config.PlatformIncomeRate).Div(decimal.NewFromInt(int64(record.SplitDays))))
		}
		for _, record := range clearRecords {
			if record.IsNegative == 1 {
				result = result.Sub(record.Amount.Mul(config.PlatformIncomeRate))
			} else {
				result = result.Add(record.Amount.Mul(config.PlatformIncomeRate))
			}

		}
		for _, record := range providerWithdrawRecords {
			result = result.Add(record.Amount)
		}
		for _, record := range providerRetrieveRecords {
			result = result.Add(record.Amount)
		}
	} else {
		for _, record := range interestRecords {
			result = result.Add(record.Amount.Mul(decimal.NewFromInt(1).Sub(config.PlatformIncomeRate)).Div(decimal.NewFromInt(int64(record.SplitDays))))
		}
		for _, record := range clearRecords {
			if record.IsNegative == 1 {
				result = result.Sub(record.Amount.Mul(decimal.NewFromInt(1).Sub(config.PlatformIncomeRate)))
			} else {
				result = result.Add(record.Amount.Mul(decimal.NewFromInt(1).Sub(config.PlatformIncomeRate)))
			}
		}
	}

	return result.RoundDown(0), nil
}

func (m *MyDb) TotalIncome30Day() (decimal.Decimal, int, error) {
	var result decimal.Decimal
	var days = 0

	now := time.Now()
	beginOfToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix()
	beginOfSearch := beginOfToday - 3600*24*30
	sqls := fmt.Sprintf("SELECT * from provide_reward_record WHERE type=0 AND source_type=0 AND (at>%d AND at<%d)", beginOfSearch, beginOfToday)
	var records []model.ProvideRewardRecord
	tx := m.Db.Raw(sqls).Scan(&records)
	if tx.Error != nil {
		return result, days, tx.Error
	}
	var earliest = records[0].At
	for _, record := range records {
		result = result.Add(record.Amount.Mul(decimal.NewFromInt(1)))
		if earliest > record.At {
			earliest = record.At
		}
	}
	days = int(int(beginOfToday)-earliest) / 3600 / 24

	return result, days, nil
}

func (m *MyDb) ProvideLiquidRecords() ([]model.ProvideLiquid, error) {
	var result []model.ProvideLiquid
	tx := m.Db.Raw(`SELECT * from provide_liquid WHERE status=0`).Scan(&result)
	if tx.Error != nil {
		return result, tx.Error
	}
	return result, nil
}

func (m *MyDb) SelectAllActiveLoans() ([]model.Loan, error) {
	var result []model.Loan
	tx := m.Db.Raw(`SELECT * from loan WHERE status=2 OR status=3`).Scan(&result)
	if tx.Error != nil {
		return result, tx.Error
	}
	return result, nil
}

func (m *MyDb) SelectAllDepositsOfActiveLoans() ([]model.Deposit, error) {
	var result []model.Deposit
	tx := m.Db.Raw(`SELECT deposit.* from deposit LEFT JOIN loan ON deposit.loan_id=loan.id WHERE loan.status=2 OR loan.status=3;`).Scan(&result)
	if tx.Error != nil {
		return result, tx.Error
	}
	return result, nil
}

func (m *MyDb) SelectLoanByETHAddress(ethAddress string) ([]model.Loan, error) {
	var loan []model.Loan
	sqls := fmt.Sprintf("SELECT * FROM loan WHERE bsc_address=\"%s\" AND status>0", ethAddress)
	tx := m.Db.Raw(sqls).Scan(&loan)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
	}
	return loan, nil
}

func (m *MyDb) SaveExchangeLpToUsdtRecord(
	forward bool,
	amount decimal.Decimal,
	address string,
	hash string,
	at int) error {
	var count int64
	tx := m.Db.Model(&model.ExchangeRecord{}).Where(&model.ExchangeRecord{
		Hash: hash,
	}).Count(&count)
	if tx.Error != nil {
		return tx.Error
	}
	if count > 0 {
		return nil
	}
	type_ := 2
	if forward {
		type_ = 1
	}
	tx = m.Db.Create(&model.ExchangeRecord{
		Type:    type_,
		Amount:  amount,
		Address: address,
		At:      at,
		Hash:    hash,
	})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (m *MyDb) SelectExchangeRecordByAddress(address string) ([]model.ExchangeRecord, error) {
	var record []model.ExchangeRecord
	sqls := fmt.Sprintf("SELECT * FROM exchange_record WHERE address=\"%s\"", address)
	tx := m.Db.Raw(sqls).Scan(&record)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
	}
	return record, nil
}

func (m *MyDb) SelectProvideIncome(address string) ([]model.RewardRecordWithProvideInfo, error) {
	var record []model.RewardRecordWithProvideInfo
	sqls := fmt.Sprintf("SELECT provide_reward_record.*,provide_liquid.amount AS income_amount,provide_liquid.duration,provide_liquid.start,provide_liquid.status,provide_liquid.create_at,provide_liquid.create_hash,provide_liquid.retrieve_at,provide_liquid.retrieve_hash FROM provide_reward_record LEFT JOIN provide_liquid ON provide_reward_record.record_id=provide_liquid.record_id WHERE provide_reward_record.provider=\"%s\" AND provide_reward_record.type=0;", address)
	tx := m.Db.Raw(sqls).Scan(&record)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
	}
	return record, nil
}

func (m *MyDb) SelectProvideIncomeWithdrawRecord(address string) ([]model.ProvideRewardRecord, error) {
	var record []model.ProvideRewardRecord
	sqls := fmt.Sprintf("SELECT * FROM provide_reward_record WHERE provider=\"%s\" AND type=1;", address)
	tx := m.Db.Raw(sqls).Scan(&record)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
	}
	return record, nil
}

func (m *MyDb) SelectRecentRateOfProvideLiquidIncome() ([]model.ProvideLiquidIncomeRateYear, error) {
	var record []model.ProvideLiquidIncomeRateYear
	sqls := fmt.Sprintf("SELECT * FROM provide_liquid_income_rate_year ORDER BY at DESC LIMIT 1;")
	tx := m.Db.Raw(sqls).Scan(&record)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
	}
	return record, nil
}

func (m *MyDb) SaveRecentRateOfProvideLiquidIncome(
	rate decimal.Decimal,
	at int) error {
	tx := m.Db.Create(&model.ProvideLiquidIncomeRateYear{
		Rate: rate,
		At:   at,
	})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (m *MyDb) NewIncomeGenerateRecord(
	at, status int,
	hash, ids, addresses, amounts string) error {
	tx := m.Db.Create(&model.IncomeGenerateRecord{
		At:        at,
		Status:    status,
		Hash:      hash,
		Ids:       ids,
		Addresses: addresses,
		Amounts:   amounts,
	})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (m *MyDb) SelectFailedIncomeGenerateRecord() ([]model.IncomeGenerateRecord, error) {
	var results []model.IncomeGenerateRecord
	sqls := fmt.Sprintf("SELECT * FROM income_generate_record WHERE status=2;")
	tx := m.Db.Raw(sqls).Scan(&results)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return results, nil
		}
		return nil, tx.Error
	}
	return results, nil
}

func (m *MyDb) CompleteIncomeGenerateRecord(id int, hash string) error {
	var record model.IncomeGenerateRecord
	tx := m.Db.Model(&model.IncomeGenerateRecord{}).First(&record, "id=?", id)
	if tx.Error != nil {
		return tx.Error
	}
	record.Status = 2
	record.Hash = hash
	tx = m.Db.Save(&record)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
