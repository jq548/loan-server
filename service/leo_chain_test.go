package service

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"loan-server/common/consts"
	"loan-server/config"
	"loan-server/contract"
	"loan-server/db"
	"math/big"
	"testing"
)

var database, _ = db.Init(&config.Db{Dsn: "dev:dev123456@tcp(192.168.188.233:3306)/loan?charset=utf8mb4&parseTime=True&loc=UTC"})
var leoService = NewLeoChainService(&config.Leo{
	"https://api.explorer.provable.com/v1",
	"testnet",
	"aleo14s8rn8km6uqwatevtk9qyf6vut0vuce398ghklqggmvrfskr7cxsrx5p7a",
	"",
}, database)

var bscService, _ = NewBscChainService(&config.Bsc{
	"https://bsc-testnet-rpc.publicnode.com",
	97,
	"0xd851D918C4970F91453f5Cf50CD59e6f38aE6D5b",
	"0xD856fEc774FA5E7CA8561DE9ef852cb0D94AFE77",
	"",
	"0xe1354798516b08D65160CA5CB2C409b166699013",
	"",
}, database)

func TestSetContractValue(t *testing.T) {
	tokenContract, err := contract.NewLoan(common.HexToAddress(bscService.LoanContractAddress), bscService.EthClient)
	if err != nil {
		t.Error(err)
		return
	}
	opts, err := bscService.GetTransactOpts("c4698e08f86bea243f5c5f08ef37ce883b51437e6a34e7922f25cf55fd32add0")
	if err != nil {
		t.Error(err)
		return
	}
	tx, err := tokenContract.SetParams(opts, big.NewInt(4), decimal.NewFromInt(consts.Wei).Mul(decimal.NewFromInt(10000)).BigInt())
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(tx.Hash().Hex())
}

func TestApprove(t *testing.T) {
	tokenContract, err := contract.NewToken(common.HexToAddress(bscService.UsdtContractAddress), bscService.EthClient)
	if err != nil {
		t.Error(err)
		return
	}
	opts, err := bscService.GetTransactOpts("841e60745df7d8e47526dd9725eec9ad6594863549ab653a57d3e93aa095d99f")
	if err != nil {
		t.Error(err)
		return
	}
	tx, err := tokenContract.Approve(opts, common.HexToAddress(bscService.LoanContractAddress), decimal.NewFromInt(consts.Wei).Mul(decimal.NewFromInt(100)).BigInt())
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(tx.Hash().Hex())
}

func TestProvideLiquid(t *testing.T) {
	loanContract, err := contract.NewLoan(common.HexToAddress(bscService.LoanContractAddress), bscService.EthClient)
	if err != nil {
		t.Error(err)
		return
	}
	opts, err := bscService.GetTransactOpts("841e60745df7d8e47526dd9725eec9ad6594863549ab653a57d3e93aa095d99f")
	if err != nil {
		t.Error(err)
		return
	}
	tx, err := loanContract.ProvideUsdt(opts, decimal.NewFromInt(consts.Wei).Mul(decimal.NewFromInt(100)).BigInt(), big.NewInt(18144000))
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(tx.Hash().Hex())
}
