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
	amount decimal.Decimal) error {
	var loan model.Loan
	tx := m.Db.Where(&model.Loan{
		ReleaseAmount: amount,
		BscAddress:    loaner,
	}).Last(&loan)
	if tx.Error != nil {
		return gorm.ErrRecordNotFound
	}

	loan.BscLoanId = loanId
	loan.ReleaseAt = at
	loan.ReleaseHash = hash

	tx = m.Db.Save(&loan)
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
	var rec []model.ProvideRecord
	tx := m.Db.Where(&model.ProvideRecord{Hash: hash, Provider: address, Amount: amount}).Find(&rec)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			record := model.ProvideRecord{
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
	at int) error {
	var rec []model.ProvideRecord
	tx := m.Db.Where(&model.ProvideRecord{Hash: hash, Provider: address, Amount: amount}).Find(&rec)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			record := model.ProvideRecord{
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
	return nil
}

func (m *MyDb) IncreaseProviderAmount(
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

func (m *MyDb) RetrieveProviderAmount(
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
