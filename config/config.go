package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Db       Db       `yaml:"db"`
	Leo      Leo      `yaml:"leo"`
	Bsc      Bsc      `yaml:"bsc"`
	Service  Service  `yaml:"service"`
	Job      Job      `yaml:"job"`
	Log      Log      `yaml:"log"`
	Platform Platform `yaml:"platform"`
}

type Db struct {
	Dsn string `yaml:"dsn"`
}

type Leo struct {
	Rpc      string `yaml:"rpc"` // rpc of aleo
	Net      string `yaml:"net"`
	Holder   string `yaml:"holder"` // holder of aleo
	HolderPK string `yaml:"holder_pk"`
}

type Bsc struct {
	Rpc          string `yaml:"rpc"`           // RPC
	ChainId      int    `yaml:"chain_id"`      //
	Usdt         string `yaml:"usdt"`          // USDT coin contract
	Dinar        string `yaml:"dinar"`         // DINAR coin contract
	LoanContract string `yaml:"loan_contract"` // main contract
	Caller       string `yaml:"caller"`        // caller
}

type Service struct {
	Port int `yaml:"port"`
}

type Job struct {
	AleoPrice               string `yaml:"aleo_price"`
	CalculateRate           string `yaml:"calculate_rate"`
	CalculateIncome         string `yaml:"calculate_income"`
	CalculateIncomeYearRate string `yaml:"calculate_income_year_rate"`
	CheckFailedJobs         string `json:"check_failed_jobs"`
	CheckExpiredLoans       string `json:"check_expired_loans"`
}

type Log struct {
	Level string `yaml:"level"`
}

type Platform struct {
	ReceiveAddress string `yaml:"receive_address"`
}

func InitConfig() (*Config, error) {
	bytes, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		return nil, err
	}
	data := &Config{}
	err = yaml.Unmarshal(bytes, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
