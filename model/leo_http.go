package model

import "github.com/shopspring/decimal"

type LeoResConfig struct {
	Rate           string   `json:"rate"`
	AvailableStage int      `json:"available_stage"`
	DayPerStage    int      `json:"day_per_stage"`
	Price          string   `json:"price"`
	AllowTypes     string   `json:"allow_types"`
	Banners        []string `json:"banners"`
	MinAmount      int64    `json:"min_amount"`
	MaxAmount      int64    `json:"max_amount"`
}

type LeoResLoan struct {
	ID                int             `json:"id"`
	AleoAddress       string          `json:"aleo_address"`
	BscAddress        string          `json:"bsc_address"`
	Status            int             `json:"status"`
	Stages            int             `json:"stages"`
	PayStages         int             `json:"pay_stages"`
	DayPerStage       int             `json:"day_per_stage"`
	StartAt           int             `json:"start_at"`
	Health            string          `json:"health"`
	Rate              string          `json:"rate"`
	ReleaseRate       string          `json:"release_rate"`
	Hash              string          `json:"hash"`
	Type              int             `json:"type"`
	Email             string          `json:"email"`
	BscLoanId         int             `json:"bsc_loan_id"`
	ReleaseAt         int             `json:"release_at"`
	ReleaseHash       string          `json:"release_hash"`
	ReleaseAmount     string          `json:"release_amount"`
	PayBackAt         int             `json:"pay_back_at"`
	PayBackHash       string          `json:"pay_back_hash"`
	PayBackAmount     string          `json:"pay_back_amount"`
	ReleaseAleoHash   string          `json:"release_aleo_hash"`
	ReleaseAleoAt     int             `json:"release_aleo_at"`
	ReleaseAleoAmount string          `json:"release_aleo_amount"`
	Deposits          []LeoResDeposit `json:"deposits"`
	DepositAmount     string          `json:"deposit_amount"`
	DepositPrice      string          `json:"deposit_price"`
	ValueWhenDeposit  string          `json:"value_when_deposit"`
	ValueCurrent      string          `json:"value_current"`
	Contract          string          `json:"contract"`
	MinRecharge       string          `json:"min_recharge"`
}

type LeoResDeposit struct {
	ID          int    `json:"id"`
	LoanId      uint   `json:"loan_id"`
	AleoAddress string `json:"aleo_address"`
	AleoAmount  string `json:"aleo_amount"`
	AleoPrice   string `json:"aleo_price"`
	UsdtValue   string `json:"usdt_value"`
	Hash        string `json:"hash"`
	At          int    `json:"at"`
	Status      int    `json:"status"`
}

type LeoResCalculateUsdt struct {
	BorrowingAmount  string              `json:"borrowing_amount"`  //
	CollateralAmount string              `json:"collateral_amount"` // value of aleo
	CollateralRate   string              `json:"collateral_rate"`   // release_rate
	Installment      []LeoResInstallment `json:"installment"`
}

type LeoResInstallment struct {
	Installments        int    `json:"installments"`
	DayPerInstallment   int    `json:"day_per_installment"`
	InterestRate        string `json:"interest_rate"`
	InterestInstallment string `json:"interest_installment"`
}

type LeoOverView struct {
	TotalProvideLiquid      string                  `json:"total_provide_liquid"`
	TotalLoaned             string                  `json:"total_loaned"`
	LiquidUsedRate          string                  `json:"liquid_used_rate"`
	ProvideLiquidRewardRate string                  `json:"provide_liquid_reward_rate"`
	TotalDepositAleo        string                  `json:"total_deposit_aleo"`
	Banners                 []string                `json:"banners"`
	HistoryRate             []LeoOverViewRateOfWeek `json:"history_rate"`
}

type LeoOverViewRateOfWeek struct {
	Rate decimal.Decimal `json:"rate"`
	At   int             `json:"at"`
	Days int             `json:"days"`
}
