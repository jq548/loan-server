package db

import (
	"errors"
	"gorm.io/gorm"
	"loan-server/model"
)

func (m *MyDb) SaveSendEmailRecord(loanId uint, success bool, email string) error {
	status := 1
	if success {
		status = 0
	}
	tx := m.Db.Create(&model.SendEmailRecord{
		LoanId: loanId,
		Status: status,
		Email:  email,
	})
	return tx.Error
}

func (m *MyDb) FindSendEmailRecord(loanId uint) ([]model.SendEmailRecord, error) {
	var results []model.SendEmailRecord
	tx := m.Db.Model(&model.SendEmailRecord{}).Where("loan_id=? AND status=1", loanId).Find(&results)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return results, nil
		}
		return nil, tx.Error
	}
	return results, nil
}

func (m *MyDb) FindLoanByStatus(status int) ([]model.Loan, error) {
	var results []model.Loan
	tx := m.Db.Model(&model.Loan{}).Where("status=?", status).Find(&results)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return results, nil
		}
		return nil, tx.Error
	}
	return results, nil
}
