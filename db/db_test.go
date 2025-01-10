package db

import (
	"loan-server/config"
	"testing"
	"time"
)

var database, _ = Init(&config.Db{Dsn: "dev:dev123456@tcp(192.168.188.233:3306)/loan?charset=utf8mb4&parseTime=True&loc=UTC"})

func TestMyDb_NewLoan(t *testing.T) {
	err := database.NewLoan("aleoAddress1", "bscAddress1", "email", 4, 7, 1, time.Now(), 0.2, 0.7)
	if err != nil {
		t.Error(err)
	}
}

func TestMyDb_NewDeposit(t *testing.T) {
	err := database.NewDeposit("aleoAddress1", "bscAddress1", 30000)
	if err != nil {
		t.Error(err)
	}
}

func TestMyDb_SaveLoanHash(t *testing.T) {
	err := database.SaveLoanHash("aleoAddress1", "hash")
	if err != nil {
		t.Error(err)
	}
}

//func TestMyDb_SaveDepositHash(t *testing.T) {
//	err := database.SaveDepositHash("aleoAddress1", "hash")
//	if err != nil {
//		t.Error(err)
//	}
//}

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
