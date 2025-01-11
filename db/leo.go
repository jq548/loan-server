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
)

func (m *MyDb) GetConfig() (model.LoanConfig, error) {
	var config model.LoanConfig
	tx := m.Db.Find(&config)
	if tx.Error != nil {
		return config, tx.Error
	}
	return config, nil
}

func (m *MyDb) GetLeoBlockHeight() (int, error) {

	cache, err := m.FindCacheByKey(consts.LeoBlockHeightKey)
	if err != nil {
		return 0, err
	}
	blockNum, err := strconv.Atoi(cache.CacheValue)
	if err != nil {
		return 0, err
	}
	return blockNum, nil
}

func (m *MyDb) SaveLeoBlockHeight(height int) error {
	cache := &model.Cache{
		CacheKey:   consts.LeoBlockHeightKey,
		CacheValue: strconv.Itoa(height),
	}
	res, err := m.UpdateCache(m.Db, cache)
	if err != nil {
		return err
	}
	if !res {
		return errors2.New(errors2.SystemError)
	}
	return nil
}

//func (m *MyDb) NewLoan(
//	aleoAddress,
//	bscAddress,
//	email string,
//	stages,
//	dayPerStage,
//	type_ int,
//	startAt time.Time,
//	rate,
//	releaseRate float32) error {
//	loan := &model.Loan{
//		AleoAddress: aleoAddress,
//		BscAddress:  bscAddress,
//		Status:      0,
//		Stages:      stages,
//		PayStages:   0,
//		DayPerStage: dayPerStage,
//		StartAt:     startAt,
//		Health:      decimal.NewFromInt(1),
//		Rate:        decimal.NewFromFloat32(rate),
//		ReleaseRate: decimal.NewFromFloat32(releaseRate),
//		Hash:        "",
//		Type:        type_,
//		Email:       email,
//	}
//	tx := m.Db.Create(&loan)
//	if tx.Error != nil {
//		return tx.Error
//	}
//	return nil
//}

func (m *MyDb) NewDeposit(
	aleoAddress, bscAddress, email string,
	aleoAmount int64,
	stages, dayPerStage int) error {
	loan := &model.Loan{
		AleoAddress: aleoAddress,
		BscAddress:  bscAddress,
		Status:      0,
		Email:       email,
		Stages:      stages,
		DayPerStage: dayPerStage,
	}
	tx := m.Db.Create(&loan)
	if tx.Error != nil {
		return tx.Error
	}

	deposit := &model.Deposit{
		LoanId:      loan.ID,
		AleoAddress: aleoAddress,
		AleoAmount:  decimal.NewFromInt(aleoAmount),
		Status:      0,
	}
	tx = m.Db.Create(&deposit)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (m *MyDb) SaveDepositHash(
	hash string,
	id uint,
	at int) error {
	var deposit model.Deposit
	var selector = model.Deposit{}
	selector.ID = id
	tx := m.Db.Where(&selector).Find(&deposit)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
	}
	deposit.Hash = hash
	deposit.Status = 1
	deposit.At = at
	tx = m.Db.Save(&deposit)
	if tx.Error != nil {
		return tx.Error
	}
	var loan model.Loan
	tx = m.Db.Where(&gorm.Model{
		ID: deposit.LoanId,
	}).Find(&loan)
	if loan.Status == 0 {
		loan.Status = 1
		loan.StartAt = at
		tx = m.Db.Save(&loan)
		if tx.Error != nil {
			return tx.Error
		}
	}
	return nil
}

func (m *MyDb) SelectLoan(
	aleoAddress string) ([]model.Loan, error) {
	var loan []model.Loan
	sqls := fmt.Sprintf("SELECT * FROM loan WHERE aleo_address=\"%s\" AND status>0", aleoAddress)
	tx := m.Db.Raw(sqls).Scan(&loan)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
	}
	return loan, nil
}

func (m *MyDb) SelectUnConfirmDepositByAddress(
	aleoAddress string, amount int64) ([]model.Deposit, error) {
	var deposits []model.Deposit
	tx := m.Db.Where(&model.Deposit{
		AleoAddress: aleoAddress,
		Status:      0,
		AleoAmount:  decimal.NewFromInt(amount),
	}).Find(&deposits)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return deposits, nil
		}
	}
	return deposits, nil
}

func (m *MyDb) SelectDepositByAddress(
	aleoAddress string) ([]model.Deposit, error) {
	var deposits []model.Deposit
	sqls := fmt.Sprintf("SELECT * FROM deposit WHERE aleo_address=\"%s\" AND status>0", aleoAddress)
	tx := m.Db.Raw(sqls).Scan(&deposits)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
	}
	return deposits, nil
}

func (m *MyDb) SelectDepositByLoanId(
	loanId int) ([]model.Deposit, error) {
	var deposits []model.Deposit
	sqls := fmt.Sprintf("SELECT * FROM deposit WHERE loan_id=%d AND status>0", loanId)
	tx := m.Db.Raw(sqls).Scan(&deposits)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
	}
	return deposits, nil
}

func (m *MyDb) Payback(
	loanId int,
	hash string,
	at int) error {
	return nil
}

func (m *MyDb) Clear(
	loanId int,
	hash string,
	at int) error {
	return nil
}

func (m *MyDb) IncreaseProviderRewardAmount(
	amount decimal.Decimal,
	address string,
	hash string,
	at int) error {
	return nil
}

func (m *MyDb) ReleaseProviderReward(
	amount decimal.Decimal,
	address string,
	hash string,
	at int) error {
	return nil
}

func (m *MyDb) IncreaseProviderAmount(
	amount decimal.Decimal,
	address string,
	hash string,
	at int) error {
	return nil
}

func (m *MyDb) RetrieveProviderAmount(
	amount decimal.Decimal,
	address string,
	hash string,
	at int) error {
	return nil
}
