package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Cache struct {
	gorm.Model
	CacheKey   string `gorm:"unique;not null"`
	CacheValue string
}

// loan data
type Loan struct {
	gorm.Model
	Contract          string
	AleoAddress       string          //
	BscAddress        string          //
	Status            int             // 0 save, 1 confirmed(receive deposit), 2 released usdt, 3 staging, 4 redeemed, 5 cleared, 6 release failed, 7 clear change state failed(contract), 8 clear sold failed, 9 clear calculate failed(income), 10 redeemed failed(transfer back)
	Stages            int             //
	PayStages         int             //
	DayPerStage       int             //
	StartAt           int             // seconds of time stamp
	Health            decimal.Decimal // health of loan
	Rate              decimal.Decimal //
	ReleaseRate       decimal.Decimal // usdt rate
	Hash              string          // first deposit hash on aleo
	Type              int             // 1
	Email             string          //
	BscLoanId         int             // loan id of contract
	DepositAmount     decimal.Decimal //
	LoanAmount        decimal.Decimal // amount of loan
	ReleaseAt         int             // loan create(release usdt) time
	ReleaseHash       string          // loan create(release usdt) hash
	ReleaseAmount     decimal.Decimal // release usdt amount
	InterestAmount    decimal.Decimal //
	PayBackAt         int             // pay back time
	PayBackHash       string          // pay back hash
	PayBackAmount     decimal.Decimal // pay back usdt amount
	ReleaseAleoHash   string          // pay back release aleo hash
	ReleaseAleoAt     int             // pay back release aleo time
	ReleaseAleoAmount decimal.Decimal // pay back release aleo amount
	ClearSoldAleo     decimal.Decimal //
	ClearSoldUsdt     decimal.Decimal // from exchanges
	ClearSoldAt       int             // from exchanges
	ClearRetrieveUsdt decimal.Decimal // from exchanges
	ClearRetrieveAt   int             // from exchanges
	ClearRetrieveHash string          // from exchanges
}

// deposit data of loan, include first and recharge
type Deposit struct {
	gorm.Model
	LoanId      uint
	AleoAddress string
	AleoAmount  decimal.Decimal
	AleoPrice   decimal.Decimal
	UsdtValue   decimal.Decimal
	Hash        string
	At          int // seconds of time stamp
	Status      int // 0 save, 1 confirmed
}

// config of program
type LoanConfig struct {
	gorm.Model
	Rate               decimal.Decimal
	ReleaseRate        decimal.Decimal
	AvailableStages    int
	DayPerStage        int
	AllowTypes         string
	BannerIds          string
	MinLoanAmount      int64
	MaxLoanAmount      int64
	AleoPrice          decimal.Decimal
	PlatformIncomeRate decimal.Decimal
}

// images
type ImageAssets struct {
	gorm.Model
	Url string
}

// provide liquid data
type ProvideLiquid struct {
	gorm.Model
	Contract     string
	Amount       decimal.Decimal
	Duration     int
	Start        int
	Status       int // 0 normal, 1 retrieve
	Provider     string
	CreateAt     int
	CreateHash   string
	RetrieveFee  decimal.Decimal
	RetrieveAt   int
	RetrieveHash string
	RecordId     int
}

// provide liquid reward data, write by caller, withdraw by provider
type ProvideRewardRecord struct {
	gorm.Model
	Type       int // 0 increase, 1 withdraw
	Provider   string
	Amount     decimal.Decimal
	Hash       string
	At         int
	Fee        decimal.Decimal
	SourceType int // 0 provide(provider), 1 loan(platform), 2 provider withdraw reward fee(platform), 3 provider release fee(platform)
	RecordId   int // type = 0 (0 (platform), 1 (provider))
}

// for select
type RewardRecordWithProvideInfo struct {
	Type         int // 0 increase, 1 withdraw
	Provider     string
	Amount       decimal.Decimal
	IncomeAmount decimal.Decimal
	Hash         string
	At           int
	SourceType   int // 0 provide(provider), 1 loan(platform), 2 provider withdraw reward fee(platform), 3 provider release fee(platform)
	RecordId     int
	Duration     int
	Start        int
	Status       int // 0 normal, 1 retrieve
	CreateAt     int
	CreateHash   string
	RetrieveAt   int
	RetrieveHash string
}

// not used
type ActionRecord struct {
	gorm.Model
	Type       int    // 0 create loan on bsc, 1 clear loan (bsc), 2 clear loan (sold aleo), 3 release aleo back, 4 update reward
	Parameters string // json
	Status     int    // 0 created, 1 processing, 2 complete, 3 failed
	Hash       string // hash of transaction
}

// leo price update record
type LeoPriceRecord struct {
	gorm.Model
	Price decimal.Decimal
	At    int
}

// loan rate update record
type LeoRateRecord struct {
	gorm.Model
	Rate decimal.Decimal
	At   int
	Days int
}

// income create record, before split on morning
type IncomeRecord struct {
	gorm.Model
	Amount       decimal.Decimal
	At           int
	Contract     string
	ContractType int    // 1 loan, 2 provide
	IsNegative   int    // 0 no, 1 yes
	Type         int    // 1 interest(3:7), 2 provider withdraw reward fee(1:0), 3 provider release fee(1:0), 4 clear(3:7)
	SplitDays    int    //
	EndAt        int    //
	Hash         string // hash of create action
}

// exchange lp to usdt record
type ExchangeRecord struct {
	gorm.Model
	Type    int // 1 lp to usdt, 2 usdt to lp
	Amount  decimal.Decimal
	Address string
	At      int
	Hash    string
}

// rate save every day
type ProvideLiquidIncomeRateYear struct {
	gorm.Model
	Rate decimal.Decimal
	At   int
}

type IncomeGenerateRecord struct {
	gorm.Model
	Ids       string
	Addresses string
	Amounts   string
	At        int
	Hash      string
	Status    int // 0 created, 1 transact success, 2 transact failed
}

type SendEmailRecord struct {
	gorm.Model
	LoanId uint
	Email  string
	Status int // 0 success, 1 failed
}
