package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Db      Db      `yaml:"db"`
	Leo     Leo     `yaml:"leo"`
	Bsc     Bsc     `yaml:"bsc"`
	Service Service `yaml:"service"`
	Job     Job     `yaml:"job"`
	Log     Log     `yaml:"log"`
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
	Lp           string `yaml:"lp"`            // LP coin contract
	Income       string `yaml:"income"`        // Income coin contract
	LoanContract string `yaml:"loan_contract"` // main contract
	Caller       string `yaml:"caller"`        // caller
}

type Service struct {
	Port int `yaml:"port"`
}

type Job struct {
	AleoPrice string `yaml:"aleo_price"`
	Income    string `yaml:"income"`
}

type Log struct {
	Level string `yaml:"level"`
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
