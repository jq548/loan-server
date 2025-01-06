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
}

type Db struct {
	Dsn string `yaml:"dsn"`
}

type Leo struct {
	Rpc    string `yaml:"rpc"`    // aleo链上的rpc
	Holder string `yaml:"holder"` // aleo链上保存代币的账户私钥
}

type Bsc struct {
	Rpc          string `yaml:"rpc"`           // RPC地址
	Usdt         string `yaml:"usdt"`          // USDT代币合约地址
	Lp           string `yaml:"lp"`            // LP代币合约地址
	Income       string `yaml:"income"`        // 收益代币合约地址
	LoanContract string `yaml:"loan_contract"` // 合约地址
}

type Service struct {
	Port int `yaml:"port"`
}

type Job struct {
	AleoPrice string `yaml:"aleo_price"`
	Income    string `yaml:"income"`
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
