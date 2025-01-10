package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type Cache struct {
	gorm.Model
	CacheKey   string `gorm:"unique;not null"`
	CacheValue string
	Expired    *time.Time
}

type Loan struct {
	gorm.Model
	AleoAddress string
	BscAddress  string
	Status      int // 0 save, 1 confirmed(receive deposit), 2 released usdt, 3 staging, 4 redeemed, 5 cleared
	Stages      int
	PayStages   int
	DayPerStage int
	StartAt     time.Time
	Health      decimal.Decimal
	Rate        decimal.Decimal
	ReleaseRate decimal.Decimal
	Hash        string // first deposit hash
	Type        int
	Email       string
}

// include first and recharge
type Deposit struct {
	gorm.Model
	LoanId      uint
	AleoAddress string
	AleoAmount  decimal.Decimal
	AleoPrice   decimal.Decimal
	UsdtValue   decimal.Decimal
	Hash        string
	At          time.Time
	Status      int // 0 save, 1 confirmed
}

type LoanConfig struct {
	gorm.Model
	Rate            decimal.Decimal
	ReleaseRate     decimal.Decimal
	AvailableStages int
	DayPerStage     int
	AllowTypes      string
	BannerIds       string
	MinLoanAmount   int64
	MaxLoanAmount   int64
	AleoPrice       decimal.Decimal
}

type ImageAssets struct {
	gorm.Model
	Url string
}

//// pay rate and pay back (usdt)
//type PayBack struct {
//	gorm.Model
//	Amount   decimal.Decimal
//	Stage    int
//	Hash     string
//	OnlyRate int
//	At       time.Time
//}
