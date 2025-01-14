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
	if cache.CacheValue == "" {
		cache.CacheValue = "0"
	}
	blockNum, err := strconv.Atoi(cache.CacheValue)
	if err != nil {
		return 0, err
	}
	return blockNum, nil
}

func (m *MyDb) SaveLeoBlockHeight(height int) error {
	res, err := m.UpdateCache(consts.LeoBlockHeightKey, strconv.Itoa(height))
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
	aleoAddress,
	bscAddress,
	email string,
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
	depositDbId uint,
	at int,
	usdt decimal.Decimal) (*model.Loan, error) {
	var deposit model.Deposit
	var selector = model.Deposit{}
	selector.ID = depositDbId
	tx := m.Db.Where(&selector).Find(&deposit)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
	}
	deposit.Hash = hash
	deposit.Status = 1
	deposit.At = at
	deposit.UsdtValue = usdt
	tx = m.Db.Save(&deposit)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var loan model.Loan
	tx = m.Db.Where(&gorm.Model{
		ID: deposit.LoanId,
	}).Find(&loan)
	if loan.Status == 0 {
		loan.Status = 1
		loan.StartAt = at
		loan.ReleaseAmount = usdt
		loan.ReleaseAt = 0
		loan.ReleaseHash = ""
		tx = m.Db.Save(&loan)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}
	return &loan, nil
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

func (m *MyDb) UpdateReleaseAleoBack(
	loaner,
	hash string,
	amount decimal.Decimal,
	at int) error {
	var loan model.Loan
	tx := m.Db.Where(&model.Loan{
		Status:      4,
		PayBackHash: "",
		AleoAddress: loaner,
	}).Last(&loan)
	if tx.Error != nil {
		return tx.Error
	}
	loan.PayBackAt = at
	loan.PayBackHash = hash
	loan.PayBackAmount = amount
	tx = m.Db.Save(&loan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (m *MyDb) GetLatestPrice() (float64, error) {
	var record model.LeoPriceRecord
	tx := m.Db.Last(&record)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return 0, nil
		}
	}
	res := record.Price.InexactFloat64()
	return res, nil
}

func (m *MyDb) GetLatestRate() (float64, error) {
	var record model.LeoRateRecord
	tx := m.Db.Last(&record)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return 0, nil
		}
	}
	res := record.Rate.InexactFloat64()
	return res, nil
}
