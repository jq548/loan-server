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
	Health            decimal.Decimal `json:"health"`
	Rate              decimal.Decimal `json:"rate"`
	ReleaseRate       decimal.Decimal `json:"release_rate"`
	Hash              string          `json:"hash"`
	Type              int             `json:"type"`
	Email             string          `json:"email"`
	BscLoanId         int             `json:"bsc_loan_id"`
	ReleaseAt         int             `json:"release_at"`
	ReleaseHash       string          `json:"release_hash"`
	ReleaseAmount     decimal.Decimal `json:"release_amount"`
	PayBackAt         int             `json:"pay_back_at"`
	PayBackHash       string          `json:"pay_back_hash"`
	PayBackAmount     decimal.Decimal `json:"pay_back_amount"`
	ReleaseAleoHash   string          `json:"release_aleo_hash"`
	ReleaseAleoAt     int             `json:"release_aleo_at"`
	ReleaseAleoAmount decimal.Decimal `json:"release_aleo_amount"`
	Deposits          []LeoResDeposit `json:"deposits"`
}

type LeoResDeposit struct {
	ID          int             `json:"id"`
	LoanId      uint            `json:"loan_id"`
	AleoAddress string          `json:"aleo_address"`
	AleoAmount  decimal.Decimal `json:"aleo_amount"`
	AleoPrice   decimal.Decimal `json:"aleo_price"`
	UsdtValue   decimal.Decimal `json:"usdt_value"`
	Hash        string          `json:"hash"`
	At          int             `json:"at"`
	Status      int             `json:"status"`
}
