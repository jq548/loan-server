package db

import (
	"loan-server/config"
	"testing"
)

var database, _ = Init(&config.Db{Dsn: "dev:dev123456@tcp(192.168.188.233:3306)/loan?charset=utf8mb4&parseTime=True&loc=UTC"})

func TestMyDb_NewDeposit(t *testing.T) {
	err := database.NewDeposit(
		"aleoAddress1",
		"bscAddress1",
		"email",
		30000,
		3,
		7)
	if err != nil {
		t.Error(err)
	}
}

func TestMyDb_SelectLoan(t *testing.T) {
	loans, err := database.SelectLoan("aleoAddress1")
	if err != nil {
		t.Error(err)
	}
	println(len(loans))
}

func TestMyDb_SelectDepositByAddress(t *testing.T) {
	deposits, err := database.SelectDepositByAddress("aleoAddress1")
	if err != nil {
		t.Error(err)
	}
	println(len(deposits))
}

func TestMyDb_SelectDepositByLoanId(t *testing.T) {
	deposits, err := database.SelectDepositByLoanId(4)
	if err != nil {
		t.Error(err)
	}
	println(len(deposits))
}

func TestMyDb_TotalIncomeLastDay(t *testing.T) {
	amount, err := database.TotalIncomeLastDay(false)
	if err != nil {
		t.Error(err)
	}
	println(amount.String())
}

func TestMyDb_SelectLoanById(t *testing.T) {
	loan, err := database.SelectLoanById(14)
	if err != nil {
		t.Error(err)
	}
	println(loan)
}
