package db

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"loan-server/common/consts"
	"loan-server/model"
	"strconv"
	"time"
)

func (m *MyDb) GetLeoBlockHeight() (int64, error) {

	cache, err := m.FindCacheByKey(consts.LeoBlockHeightKey)
	if err != nil {
		return 0, err
	}
	blockNum, err := strconv.ParseInt(cache.CacheValue, 10, 64)
	if err != nil {
		return 0, err
	}
	return blockNum, nil
}

func (m *MyDb) NewLoan(
	aleoAddress,
	bscAddress,
	email string,
	stages,
	dayPerStage,
	type_ int,
	startAt time.Time,
	rate,
	releaseRate float32) error {
	loan := &model.Loan{
		AleoAddress: aleoAddress,
		BscAddress:  bscAddress,
		Status:      0,
		Stages:      stages,
		PayStages:   0,
		DayPerStage: dayPerStage,
		StartAt:     startAt,
		Health:      decimal.NewFromInt(1),
		Rate:        decimal.NewFromFloat32(rate),
		ReleaseRate: decimal.NewFromFloat32(releaseRate),
		Hash:        "",
		Type:        type_,
		Email:       email,
	}
	tx := m.Db.Create(&loan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (m *MyDb) NewDeposit(
	aleoAddress string,
	loanId int,
	aleoAmount,
	aleoPrice,
	usdtValue decimal.Decimal,
	at time.Time) error {
	deposit := &model.Deposit{
		LoanId:      uint(loanId),
		AleoAddress: aleoAddress,
		AleoAmount:  aleoAmount,
		AleoPrice:   aleoPrice,
		UsdtValue:   usdtValue,
		At:          at,
		Status:      0,
	}
	tx := m.Db.Create(&deposit)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (m *MyDb) SaveLoanHash(
	aleoAddress, hash string) error {
	var loan model.Loan
	tx := m.Db.Where(&model.Loan{
		AleoAddress: aleoAddress,
	}).Order("created_at desc").First(&loan)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
	}
	loan.Hash = hash
	loan.Status = 1
	tx = m.Db.Save(&loan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (m *MyDb) SaveDepositHash(
	aleoAddress, hash string) error {
	var deposit model.Deposit
	tx := m.Db.Where(&model.Deposit{
		AleoAddress: aleoAddress,
	}).Order("created_at desc").First(&deposit)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
	}
	deposit.Hash = hash
	deposit.Status = 1
	tx = m.Db.Save(&deposit)
	if tx.Error != nil {
		return tx.Error
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
