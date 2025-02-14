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

func (m *MyDb) NewDeposit(
	aleoAddress,
	bscAddress,
	email string,
	aleoAmount int64,
	stages, dayPerStage int,
	releaseRate decimal.Decimal) error {
	loan := &model.Loan{
		AleoAddress:    aleoAddress,
		BscAddress:     bscAddress,
		Status:         0,
		Email:          email,
		Stages:         stages,
		DayPerStage:    dayPerStage,
		InterestAmount: decimal.NewFromInt(0),
		ReleaseRate:    releaseRate,
		DepositAmount:  decimal.NewFromInt(aleoAmount),
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

func (m *MyDb) NewDepositOfLoan(
	loanId int,
	aleoAddress string,
	aleoAmount int64) error {
	deposit := &model.Deposit{
		LoanId:      uint(loanId),
		AleoAddress: aleoAddress,
		AleoAmount:  decimal.NewFromInt(aleoAmount),
		Status:      0,
	}
	tx := m.Db.Create(&deposit)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (m *MyDb) SaveCreateFailed(id, status int) error {
	var loan model.Loan
	tx := m.Db.First(&loan, id)
	if tx.Error != nil {
		return tx.Error
	}
	loan.Status = status
	tx = m.Db.Save(&loan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (m *MyDb) SaveDepositHash(
	hash string,
	depositDbId uint,
	at int,
	loanAmount,
	interestAmount, price, rate decimal.Decimal) (*model.Loan, error) {
	var deposit model.Deposit
	var selector = model.Deposit{}
	selector.ID = depositDbId
	tx := m.Db.Model(&model.Deposit{}).Where(&selector).Find(&deposit)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, tx.Error
	}
	deposit.Hash = hash
	deposit.Status = 1
	deposit.At = at
	deposit.UsdtValue = loanAmount
	deposit.AleoPrice = price
	tx = m.Db.Save(&deposit)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var loan model.Loan
	tx = m.Db.Model(&model.Loan{}).Where(&gorm.Model{
		ID: deposit.LoanId,
	}).Find(&loan)
	if loan.Status == 0 {
		loan.Type = 1
		loan.Rate = rate
		loan.Status = 1
		loan.StartAt = at
		loan.LoanAmount = loanAmount
		loan.ReleaseAmount = loanAmount.Sub(interestAmount)
		loan.ReleaseAt = 0
		loan.ReleaseHash = ""
		loan.InterestAmount = interestAmount
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
		return nil, tx.Error
	}
	return loan, nil
}

func (m *MyDb) SelectUnConfirmDepositByAddress(
	aleoAddress string, amount int64) ([]model.Deposit, error) {
	var deposits []model.Deposit
	sqls := fmt.Sprintf("SELECT * FROM deposit WHERE status=0 AND hash=\"\" AND aleo_address=\"%s\" AND aleo_amount=%d ORDER BY id DESC;", aleoAddress, amount)
	tx := m.Db.Raw(sqls).Scan(&deposits)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return deposits, nil
		}
		return nil, tx.Error
	}
	return deposits, nil
}

func (m *MyDb) SelectDepositByAddress(
	aleoAddress string) ([]model.Deposit, error) {
	var deposits []model.Deposit
	sqls := fmt.Sprintf("SELECT * FROM deposit WHERE aleo_address=\"%s\" AND status>0", aleoAddress)
	tx := m.Db.Raw(sqls).Scan(&deposits)
	if tx.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return deposits, nil
}

func (m *MyDb) SelectDepositByLoanId(
	loanId int) ([]model.Deposit, error) {
	var deposits []model.Deposit
	sqls := fmt.Sprintf("SELECT * FROM deposit WHERE loan_id=%d AND status>0", loanId)
	tx := m.Db.Raw(sqls).Scan(&deposits)
	if tx.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return deposits, nil
}

func (m *MyDb) SaveHealthOfLoan(
	loanId int, health decimal.Decimal) error {
	var loan model.Loan
	tx := m.Db.Model(&model.Loan{}).First(&loan, loanId)
	if tx.Error != nil {
		return tx.Error
	}
	loan.Health = health
	tx = m.Db.Save(&loan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (m *MyDb) SaveStatusOfLoan(
	loanId, status int) error {
	var loan model.Loan
	tx := m.Db.Model(&model.Loan{}).First(&loan, loanId)
	if tx.Error != nil {
		return tx.Error
	}
	loan.Status = status
	tx = m.Db.Save(&loan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (m *MyDb) UpdateReleaseAleoBack(
	loaner,
	hash string,
	amount decimal.Decimal,
	at int) error {
	var count int64
	tx := m.Db.Model(&model.Loan{}).Where(&model.Loan{
		PayBackHash: hash,
	}).Count(&count)
	if tx.Error != nil {
		return tx.Error
	}
	if count == 1 {
		return nil
	}
	var loans []model.Loan
	sqls := fmt.Sprintf("SELECT * FROM loan WHERE pay_back_hash=\"\" AND status=4 AND address=\"%s\"", loaner)
	tx = m.Db.Raw(sqls).Scan(loans)
	if tx.Error != nil {
		return tx.Error
	}
	if len(loans) > 0 {
		loan := loans[0]
		loan.PayBackAt = at
		loan.PayBackHash = hash
		loan.PayBackAmount = amount
		tx = m.Db.Save(&loan)
		if tx.Error != nil {
			return tx.Error
		}
	}
	return nil
}

func (m *MyDb) GetLatestPrice() (float64, error) {
	var record model.LeoPriceRecord
	tx := m.Db.Model(&model.LeoPriceRecord{}).Last(&record)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, tx.Error
	}
	res := record.Price.InexactFloat64()
	return res, nil
}

func (m *MyDb) GetLatestRateOfWeek() ([]model.LeoRateRecord, error) {
	var record []model.LeoRateRecord
	tx := m.Db.Model(&model.LeoRateRecord{}).Order("id DESC").Limit(4).Find(&record)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return record, nil
		}
		return nil, tx.Error
	}
	return record, nil
}

func (m *MyDb) GetBanners() ([]model.ImageAssets, error) {
	var assets []model.ImageAssets
	tx := m.Db.Model(&model.ImageAssets{}).Find(&assets)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return assets, nil
		}
		return nil, tx.Error
	}
	return assets, nil
}

func (m *MyDb) SelectHistoryOfRateOf1Week() ([]model.LeoRateRecord, error) {
	var record []model.LeoRateRecord
	ago := time.Now().Unix() - 3600*24*365
	sqls := fmt.Sprintf("SELECT * FROM leo_rate_record WHERE days=7 AND at>%d;", ago)
	tx := m.Db.Raw(sqls).Scan(&record)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return record, nil
		}
		return record, tx.Error
	}
	return record, nil
}

func (m *MyDb) SaveLatestAleoPrice(price float64) error {
	tx := m.Db.Create(&model.LeoPriceRecord{
		Price: decimal.NewFromFloat(price),
		At:    int(time.Now().Unix()),
	})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (m *MyDb) SaveLatestRate(record model.LeoRateRecord) error {
	tx := m.Db.Create(&record)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (m *MyDb) SelectLoanById(loanId uint) (*model.Loan, error) {
	var loan model.Loan
	var selector = &model.Loan{}
	selector.ID = loanId
	tx := m.Db.Model(&model.Loan{}).Where(selector).First(&loan)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, tx.Error
	}
	return &loan, nil
}
