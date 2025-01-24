package model

import "github.com/shopspring/decimal"

type BscExchangeRecord struct {
	Type    int    `json:"type"` // 1 lp to usdt, 2 usdt to lp
	Amount  string `json:"amount"`
	Address string `json:"address"`
	At      int    `json:"at"`
	Hash    string `json:"hash"`
}

type BscProvideInfo struct {
	TotalProvide    string             `json:"total_provide"`
	Income30        string             `json:"income_30"`
	IncomeYesterday string             `json:"income_yesterday"`
	ProvideRecord   []BscProvideRecord `json:"provide_record"`
}

type BscProvideRecord struct {
	Days               int             `json:"days"`
	Amount             string          `json:"amount"`
	RateYear           string          `json:"rate_year"`
	TotalIncome        string          `json:"total_income"`
	Duration           int             `json:"duration"`
	Start              int             `json:"start"`
	Status             int             `json:"status"`
	Provider           string          `json:"provider"`
	CreateAt           int             `json:"create_at"`
	CreateHash         string          `json:"create_hash"`
	RetrieveAt         int             `json:"retrieve_at"`
	RetrieveHash       string          `json:"retrieve_hash"`
	RecordId           int             `json:"record_id"`
	YesterdayIncome    string          `json:"yesterday_income"`
	TotalIncomeDec     decimal.Decimal `json:"total_income_dec"`
	YesterdayIncomeDec decimal.Decimal `json:"yesterday_income_dec"`
	IncomeStartDay     string          `json:"income_start_day"`
	IncomeEndDay       string          `json:"income_end_day"`
}

type BscProvideRewardRecord struct {
	Provider string `json:"provider"`
	Amount   string `json:"amount"`
	Hash     string `json:"hash"`
	At       int    `json:"at"`
}

type BscConfig struct {
	WithdrawIncomeFee  string                   `json:"withdraw_income_fee"`
	MinProvideAmount   string                   `json:"min_provide_amount"`
	MaxProvideAmount   string                   `json:"max_provide_amount"`
	WithdrawProvideFee string                   `json:"withdraw_provide_fee"` // withdraw before create + duration
	ProvideLiquid      []BscConfigProvideLiquid `json:"provide_liquid"`
}

type BscConfigProvideLiquid struct {
	Duration         int     `json:"duration"`
	Days             int     `json:"days"`
	EstimateWeekRate float32 `json:"estimate_week_rate"`
}
