package db

import (
	"errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"loan-server/common/consts"
	errors2 "loan-server/common/errors"
	"loan-server/model"
	"strconv"
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
