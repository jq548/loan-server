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
	var loan model.Loan
	var selector = model.Loan{}
	selector.ID = uint(loanId)
	tx := m.Db.Where(&selector).Last(&loan)
	if tx.Error != nil {
		return gorm.ErrRecordNotFound
	}

	loan.BscLoanId = loanId
	loan.ReleaseAt = at
	loan.ReleaseHash = hash
	loan.ReleaseAmount = releaseAmount

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
	var loan model.Loan
	tx := m.Db.Where(&model.Loan{
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
	var loan model.Loan
	tx := m.Db.Where(&model.Loan{
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
	at int) error {

	var rec []model.ProvideRewardRecord
	tx := m.Db.Where(&model.ProvideRewardRecord{Hash: hash, Provider: address, Amount: amount}).Find(&rec)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			record := model.ProvideRewardRecord{
				Type:     0,
				Provider: address,
				Amount:   amount,
				At:       at,
				Hash:     hash,
			}
			tx := m.Db.Create(&record)
			if tx.Error != nil {
				return tx.Error
			}
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
	tx := m.Db.Where(&model.ProvideRewardRecord{Hash: hash, Provider: address, Amount: amount}).Find(&rec)
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
		}
		return tx.Error
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
	at int) error {
	tx := m.Db.Create(&model.ProvideLiquid{
		Amount:     amount,
		Duration:   duration,
		Start:      start,
		Status:     0,
		Provider:   address,
		CreateAt:   at,
		CreateHash: hash,
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
	var pl model.ProvideLiquid
	var selector = model.ProvideLiquid{}
	selector.ID = uint(id)
	tx := m.Db.Where(&selector).First(&pl)
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

	return result.RoundDown(6), nil
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
