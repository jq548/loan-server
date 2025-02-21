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
	"0x15dD7A4580A0523215bB1ae8B00f1Aa4D93ae308",
	"",
}, database, &config.GateIo{
	"",
	"",
})

func TestSetContractValue(t *testing.T) {
	tokenContract, err := contract.NewLoan(common.HexToAddress(bscService.LoanContractAddress), bscService.EthClient)
	if err != nil {
		t.Error(err)
		return
	}
	opts, err := bscService.GetTransactOpts("")
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
	opts, err := bscService.GetTransactOpts("")
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
	opts, err := bscService.GetTransactOpts("")
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

func TestBscChainService_SoldAleo(t *testing.T) {
	bscService.FetchAleoPrice()
}
