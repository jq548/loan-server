// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// LoanMetaData contains all meta data concerning the Loan contract.
var LoanMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"loaner\",\"type\":\"address\"}],\"name\":\"eventClear\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"forward\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"eventExchangeLpToUsdt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"eventIncreaseLiquidReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"providers\",\"type\":\"address[]\"}],\"name\":\"eventIncreaseLiquidRewardBath\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loanAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"loaner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"releaseAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestAmount\",\"type\":\"uint256\"}],\"name\":\"eventNewLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"loaner\",\"type\":\"address\"}],\"name\":\"eventPayBack\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"eventProviderAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"eventProviderRedeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"eventReleaseLiquidReward\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"loaner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"interestAmount\",\"type\":\"uint256\"}],\"name\":\"addNewLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"}],\"name\":\"clear\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"forward\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"exchangeLpUsdt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"extract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"increaseLiquidReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"providers\",\"type\":\"address[]\"}],\"name\":\"increaseLiquidRewardBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usdt\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lp\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"liquidProvides\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"liquidReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"loans\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"loanAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"loaner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"releaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"forward\",\"type\":\"bool\"}],\"name\":\"maxExchangeLpUsdt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"params\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"}],\"name\":\"payBack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"provideUsdt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"releaseAbleLiquidReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"releaseLiquidReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"provideId\",\"type\":\"uint256\"}],\"name\":\"retrieveUsdt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lp\",\"type\":\"address\"}],\"name\":\"setLpContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usdt\",\"type\":\"address\"}],\"name\":\"setUsdtContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"transferCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LoanABI is the input ABI used to generate the binding from.
// Deprecated: Use LoanMetaData.ABI instead.
var LoanABI = LoanMetaData.ABI

// Loan is an auto generated Go binding around an Ethereum contract.
type Loan struct {
	LoanCaller     // Read-only binding to the contract
	LoanTransactor // Write-only binding to the contract
	LoanFilterer   // Log filterer for contract events
}

// LoanCaller is an auto generated read-only Go binding around an Ethereum contract.
type LoanCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoanTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LoanTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoanFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LoanFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoanSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LoanSession struct {
	Contract     *Loan             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoanCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LoanCallerSession struct {
	Contract *LoanCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LoanTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LoanTransactorSession struct {
	Contract     *LoanTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoanRaw is an auto generated low-level Go binding around an Ethereum contract.
type LoanRaw struct {
	Contract *Loan // Generic contract binding to access the raw methods on
}

// LoanCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LoanCallerRaw struct {
	Contract *LoanCaller // Generic read-only contract binding to access the raw methods on
}

// LoanTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LoanTransactorRaw struct {
	Contract *LoanTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLoan creates a new instance of Loan, bound to a specific deployed contract.
func NewLoan(address common.Address, backend bind.ContractBackend) (*Loan, error) {
	contract, err := bindLoan(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Loan{LoanCaller: LoanCaller{contract: contract}, LoanTransactor: LoanTransactor{contract: contract}, LoanFilterer: LoanFilterer{contract: contract}}, nil
}

// NewLoanCaller creates a new read-only instance of Loan, bound to a specific deployed contract.
func NewLoanCaller(address common.Address, caller bind.ContractCaller) (*LoanCaller, error) {
	contract, err := bindLoan(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LoanCaller{contract: contract}, nil
}

// NewLoanTransactor creates a new write-only instance of Loan, bound to a specific deployed contract.
func NewLoanTransactor(address common.Address, transactor bind.ContractTransactor) (*LoanTransactor, error) {
	contract, err := bindLoan(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LoanTransactor{contract: contract}, nil
}

// NewLoanFilterer creates a new log filterer instance of Loan, bound to a specific deployed contract.
func NewLoanFilterer(address common.Address, filterer bind.ContractFilterer) (*LoanFilterer, error) {
	contract, err := bindLoan(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LoanFilterer{contract: contract}, nil
}

// bindLoan binds a generic wrapper to an already deployed contract.
func bindLoan(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LoanMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Loan *LoanRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Loan.Contract.LoanCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Loan *LoanRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Loan.Contract.LoanTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Loan *LoanRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Loan.Contract.LoanTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Loan *LoanCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Loan.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Loan *LoanTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Loan.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Loan *LoanTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Loan.Contract.contract.Transact(opts, method, params...)
}

// Addresses is a free data retrieval call binding the contract method 0xedf26d9b.
//
// Solidity: function addresses(uint256 ) view returns(address)
func (_Loan *LoanCaller) Addresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Loan.contract.Call(opts, &out, "addresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addresses is a free data retrieval call binding the contract method 0xedf26d9b.
//
// Solidity: function addresses(uint256 ) view returns(address)
func (_Loan *LoanSession) Addresses(arg0 *big.Int) (common.Address, error) {
	return _Loan.Contract.Addresses(&_Loan.CallOpts, arg0)
}

// Addresses is a free data retrieval call binding the contract method 0xedf26d9b.
//
// Solidity: function addresses(uint256 ) view returns(address)
func (_Loan *LoanCallerSession) Addresses(arg0 *big.Int) (common.Address, error) {
	return _Loan.Contract.Addresses(&_Loan.CallOpts, arg0)
}

// LiquidProvides is a free data retrieval call binding the contract method 0xdf0796c0.
//
// Solidity: function liquidProvides(uint256 ) view returns(uint256 amount, uint256 duration, uint256 start, uint256 status, address provider)
func (_Loan *LoanCaller) LiquidProvides(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount   *big.Int
	Duration *big.Int
	Start    *big.Int
	Status   *big.Int
	Provider common.Address
}, error) {
	var out []interface{}
	err := _Loan.contract.Call(opts, &out, "liquidProvides", arg0)

	outstruct := new(struct {
		Amount   *big.Int
		Duration *big.Int
		Start    *big.Int
		Status   *big.Int
		Provider common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Start = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Provider = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// LiquidProvides is a free data retrieval call binding the contract method 0xdf0796c0.
//
// Solidity: function liquidProvides(uint256 ) view returns(uint256 amount, uint256 duration, uint256 start, uint256 status, address provider)
func (_Loan *LoanSession) LiquidProvides(arg0 *big.Int) (struct {
	Amount   *big.Int
	Duration *big.Int
	Start    *big.Int
	Status   *big.Int
	Provider common.Address
}, error) {
	return _Loan.Contract.LiquidProvides(&_Loan.CallOpts, arg0)
}

// LiquidProvides is a free data retrieval call binding the contract method 0xdf0796c0.
//
// Solidity: function liquidProvides(uint256 ) view returns(uint256 amount, uint256 duration, uint256 start, uint256 status, address provider)
func (_Loan *LoanCallerSession) LiquidProvides(arg0 *big.Int) (struct {
	Amount   *big.Int
	Duration *big.Int
	Start    *big.Int
	Status   *big.Int
	Provider common.Address
}, error) {
	return _Loan.Contract.LiquidProvides(&_Loan.CallOpts, arg0)
}

// LiquidReward is a free data retrieval call binding the contract method 0xa9bc32d7.
//
// Solidity: function liquidReward(address ) view returns(uint256)
func (_Loan *LoanCaller) LiquidReward(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Loan.contract.Call(opts, &out, "liquidReward", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidReward is a free data retrieval call binding the contract method 0xa9bc32d7.
//
// Solidity: function liquidReward(address ) view returns(uint256)
func (_Loan *LoanSession) LiquidReward(arg0 common.Address) (*big.Int, error) {
	return _Loan.Contract.LiquidReward(&_Loan.CallOpts, arg0)
}

// LiquidReward is a free data retrieval call binding the contract method 0xa9bc32d7.
//
// Solidity: function liquidReward(address ) view returns(uint256)
func (_Loan *LoanCallerSession) LiquidReward(arg0 common.Address) (*big.Int, error) {
	return _Loan.Contract.LiquidReward(&_Loan.CallOpts, arg0)
}

// Loans is a free data retrieval call binding the contract method 0xe1ec3c68.
//
// Solidity: function loans(uint256 ) view returns(uint256 loanAmount, uint256 duration, uint256 start, uint256 status, address loaner, uint256 releaseAmount, uint256 interestAmount)
func (_Loan *LoanCaller) Loans(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LoanAmount     *big.Int
	Duration       *big.Int
	Start          *big.Int
	Status         *big.Int
	Loaner         common.Address
	ReleaseAmount  *big.Int
	InterestAmount *big.Int
}, error) {
	var out []interface{}
	err := _Loan.contract.Call(opts, &out, "loans", arg0)

	outstruct := new(struct {
		LoanAmount     *big.Int
		Duration       *big.Int
		Start          *big.Int
		Status         *big.Int
		Loaner         common.Address
		ReleaseAmount  *big.Int
		InterestAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LoanAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Start = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Loaner = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.ReleaseAmount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.InterestAmount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Loans is a free data retrieval call binding the contract method 0xe1ec3c68.
//
// Solidity: function loans(uint256 ) view returns(uint256 loanAmount, uint256 duration, uint256 start, uint256 status, address loaner, uint256 releaseAmount, uint256 interestAmount)
func (_Loan *LoanSession) Loans(arg0 *big.Int) (struct {
	LoanAmount     *big.Int
	Duration       *big.Int
	Start          *big.Int
	Status         *big.Int
	Loaner         common.Address
	ReleaseAmount  *big.Int
	InterestAmount *big.Int
}, error) {
	return _Loan.Contract.Loans(&_Loan.CallOpts, arg0)
}

// Loans is a free data retrieval call binding the contract method 0xe1ec3c68.
//
// Solidity: function loans(uint256 ) view returns(uint256 loanAmount, uint256 duration, uint256 start, uint256 status, address loaner, uint256 releaseAmount, uint256 interestAmount)
func (_Loan *LoanCallerSession) Loans(arg0 *big.Int) (struct {
	LoanAmount     *big.Int
	Duration       *big.Int
	Start          *big.Int
	Status         *big.Int
	Loaner         common.Address
	ReleaseAmount  *big.Int
	InterestAmount *big.Int
}, error) {
	return _Loan.Contract.Loans(&_Loan.CallOpts, arg0)
}

// MaxExchangeLpUsdt is a free data retrieval call binding the contract method 0xca938a62.
//
// Solidity: function maxExchangeLpUsdt(bool forward) view returns(uint256)
func (_Loan *LoanCaller) MaxExchangeLpUsdt(opts *bind.CallOpts, forward bool) (*big.Int, error) {
	var out []interface{}
	err := _Loan.contract.Call(opts, &out, "maxExchangeLpUsdt", forward)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxExchangeLpUsdt is a free data retrieval call binding the contract method 0xca938a62.
//
// Solidity: function maxExchangeLpUsdt(bool forward) view returns(uint256)
func (_Loan *LoanSession) MaxExchangeLpUsdt(forward bool) (*big.Int, error) {
	return _Loan.Contract.MaxExchangeLpUsdt(&_Loan.CallOpts, forward)
}

// MaxExchangeLpUsdt is a free data retrieval call binding the contract method 0xca938a62.
//
// Solidity: function maxExchangeLpUsdt(bool forward) view returns(uint256)
func (_Loan *LoanCallerSession) MaxExchangeLpUsdt(forward bool) (*big.Int, error) {
	return _Loan.Contract.MaxExchangeLpUsdt(&_Loan.CallOpts, forward)
}

// Params is a free data retrieval call binding the contract method 0x9d2f053c.
//
// Solidity: function params(uint256 ) view returns(uint256)
func (_Loan *LoanCaller) Params(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Loan.contract.Call(opts, &out, "params", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Params is a free data retrieval call binding the contract method 0x9d2f053c.
//
// Solidity: function params(uint256 ) view returns(uint256)
func (_Loan *LoanSession) Params(arg0 *big.Int) (*big.Int, error) {
	return _Loan.Contract.Params(&_Loan.CallOpts, arg0)
}

// Params is a free data retrieval call binding the contract method 0x9d2f053c.
//
// Solidity: function params(uint256 ) view returns(uint256)
func (_Loan *LoanCallerSession) Params(arg0 *big.Int) (*big.Int, error) {
	return _Loan.Contract.Params(&_Loan.CallOpts, arg0)
}

// ReleaseAbleLiquidReward is a free data retrieval call binding the contract method 0xae357555.
//
// Solidity: function releaseAbleLiquidReward(address provider) view returns(uint256)
func (_Loan *LoanCaller) ReleaseAbleLiquidReward(opts *bind.CallOpts, provider common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Loan.contract.Call(opts, &out, "releaseAbleLiquidReward", provider)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReleaseAbleLiquidReward is a free data retrieval call binding the contract method 0xae357555.
//
// Solidity: function releaseAbleLiquidReward(address provider) view returns(uint256)
func (_Loan *LoanSession) ReleaseAbleLiquidReward(provider common.Address) (*big.Int, error) {
	return _Loan.Contract.ReleaseAbleLiquidReward(&_Loan.CallOpts, provider)
}

// ReleaseAbleLiquidReward is a free data retrieval call binding the contract method 0xae357555.
//
// Solidity: function releaseAbleLiquidReward(address provider) view returns(uint256)
func (_Loan *LoanCallerSession) ReleaseAbleLiquidReward(provider common.Address) (*big.Int, error) {
	return _Loan.Contract.ReleaseAbleLiquidReward(&_Loan.CallOpts, provider)
}

// AddNewLoan is a paid mutator transaction binding the contract method 0x7535cd4b.
//
// Solidity: function addNewLoan(uint256 id, uint256 amount, uint256 duration, address loaner, uint256 interestAmount) returns()
func (_Loan *LoanTransactor) AddNewLoan(opts *bind.TransactOpts, id *big.Int, amount *big.Int, duration *big.Int, loaner common.Address, interestAmount *big.Int) (*types.Transaction, error) {
	return _Loan.contract.Transact(opts, "addNewLoan", id, amount, duration, loaner, interestAmount)
}

// AddNewLoan is a paid mutator transaction binding the contract method 0x7535cd4b.
//
// Solidity: function addNewLoan(uint256 id, uint256 amount, uint256 duration, address loaner, uint256 interestAmount) returns()
func (_Loan *LoanSession) AddNewLoan(id *big.Int, amount *big.Int, duration *big.Int, loaner common.Address, interestAmount *big.Int) (*types.Transaction, error) {
	return _Loan.Contract.AddNewLoan(&_Loan.TransactOpts, id, amount, duration, loaner, interestAmount)
}

// AddNewLoan is a paid mutator transaction binding the contract method 0x7535cd4b.
//
// Solidity: function addNewLoan(uint256 id, uint256 amount, uint256 duration, address loaner, uint256 interestAmount) returns()
func (_Loan *LoanTransactorSession) AddNewLoan(id *big.Int, amount *big.Int, duration *big.Int, loaner common.Address, interestAmount *big.Int) (*types.Transaction, error) {
	return _Loan.Contract.AddNewLoan(&_Loan.TransactOpts, id, amount, duration, loaner, interestAmount)
}

// Clear is a paid mutator transaction binding the contract method 0xc0fe1af8.
//
// Solidity: function clear(uint256 loanId) returns()
func (_Loan *LoanTransactor) Clear(opts *bind.TransactOpts, loanId *big.Int) (*types.Transaction, error) {
	return _Loan.contract.Transact(opts, "clear", loanId)
}

// Clear is a paid mutator transaction binding the contract method 0xc0fe1af8.
//
// Solidity: function clear(uint256 loanId) returns()
func (_Loan *LoanSession) Clear(loanId *big.Int) (*types.Transaction, error) {
	return _Loan.Contract.Clear(&_Loan.TransactOpts, loanId)
}

// Clear is a paid mutator transaction binding the contract method 0xc0fe1af8.
//
// Solidity: function clear(uint256 loanId) returns()
func (_Loan *LoanTransactorSession) Clear(loanId *big.Int) (*types.Transaction, error) {
	return _Loan.Contract.Clear(&_Loan.TransactOpts, loanId)
}

// ExchangeLpUsdt is a paid mutator transaction binding the contract method 0x5f141556.
//
// Solidity: function exchangeLpUsdt(bool forward, uint256 amount) returns()
func (_Loan *LoanTransactor) ExchangeLpUsdt(opts *bind.TransactOpts, forward bool, amount *big.Int) (*types.Transaction, error) {
	return _Loan.contract.Transact(opts, "exchangeLpUsdt", forward, amount)
}

// ExchangeLpUsdt is a paid mutator transaction binding the contract method 0x5f141556.
//
// Solidity: function exchangeLpUsdt(bool forward, uint256 amount) returns()
func (_Loan *LoanSession) ExchangeLpUsdt(forward bool, amount *big.Int) (*types.Transaction, error) {
	return _Loan.Contract.ExchangeLpUsdt(&_Loan.TransactOpts, forward, amount)
}

// ExchangeLpUsdt is a paid mutator transaction binding the contract method 0x5f141556.
//
// Solidity: function exchangeLpUsdt(bool forward, uint256 amount) returns()
func (_Loan *LoanTransactorSession) ExchangeLpUsdt(forward bool, amount *big.Int) (*types.Transaction, error) {
	return _Loan.Contract.ExchangeLpUsdt(&_Loan.TransactOpts, forward, amount)
}

// Extract is a paid mutator transaction binding the contract method 0xc7a5d285.
//
// Solidity: function extract(address token) returns()
func (_Loan *LoanTransactor) Extract(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Loan.contract.Transact(opts, "extract", token)
}

// Extract is a paid mutator transaction binding the contract method 0xc7a5d285.
//
// Solidity: function extract(address token) returns()
func (_Loan *LoanSession) Extract(token common.Address) (*types.Transaction, error) {
	return _Loan.Contract.Extract(&_Loan.TransactOpts, token)
}

// Extract is a paid mutator transaction binding the contract method 0xc7a5d285.
//
// Solidity: function extract(address token) returns()
func (_Loan *LoanTransactorSession) Extract(token common.Address) (*types.Transaction, error) {
	return _Loan.Contract.Extract(&_Loan.TransactOpts, token)
}

// IncreaseLiquidReward is a paid mutator transaction binding the contract method 0xbaafdf5d.
//
// Solidity: function increaseLiquidReward(uint256 amount, address provider) returns()
func (_Loan *LoanTransactor) IncreaseLiquidReward(opts *bind.TransactOpts, amount *big.Int, provider common.Address) (*types.Transaction, error) {
	return _Loan.contract.Transact(opts, "increaseLiquidReward", amount, provider)
}

// IncreaseLiquidReward is a paid mutator transaction binding the contract method 0xbaafdf5d.
//
// Solidity: function increaseLiquidReward(uint256 amount, address provider) returns()
func (_Loan *LoanSession) IncreaseLiquidReward(amount *big.Int, provider common.Address) (*types.Transaction, error) {
	return _Loan.Contract.IncreaseLiquidReward(&_Loan.TransactOpts, amount, provider)
}

// IncreaseLiquidReward is a paid mutator transaction binding the contract method 0xbaafdf5d.
//
// Solidity: function increaseLiquidReward(uint256 amount, address provider) returns()
func (_Loan *LoanTransactorSession) IncreaseLiquidReward(amount *big.Int, provider common.Address) (*types.Transaction, error) {
	return _Loan.Contract.IncreaseLiquidReward(&_Loan.TransactOpts, amount, provider)
}

// IncreaseLiquidRewardBatch is a paid mutator transaction binding the contract method 0x514e0a85.
//
// Solidity: function increaseLiquidRewardBatch(uint256[] ids, uint256[] amounts, address[] providers) returns()
func (_Loan *LoanTransactor) IncreaseLiquidRewardBatch(opts *bind.TransactOpts, ids []*big.Int, amounts []*big.Int, providers []common.Address) (*types.Transaction, error) {
	return _Loan.contract.Transact(opts, "increaseLiquidRewardBatch", ids, amounts, providers)
}

// IncreaseLiquidRewardBatch is a paid mutator transaction binding the contract method 0x514e0a85.
//
// Solidity: function increaseLiquidRewardBatch(uint256[] ids, uint256[] amounts, address[] providers) returns()
func (_Loan *LoanSession) IncreaseLiquidRewardBatch(ids []*big.Int, amounts []*big.Int, providers []common.Address) (*types.Transaction, error) {
	return _Loan.Contract.IncreaseLiquidRewardBatch(&_Loan.TransactOpts, ids, amounts, providers)
}

// IncreaseLiquidRewardBatch is a paid mutator transaction binding the contract method 0x514e0a85.
//
// Solidity: function increaseLiquidRewardBatch(uint256[] ids, uint256[] amounts, address[] providers) returns()
func (_Loan *LoanTransactorSession) IncreaseLiquidRewardBatch(ids []*big.Int, amounts []*big.Int, providers []common.Address) (*types.Transaction, error) {
	return _Loan.Contract.IncreaseLiquidRewardBatch(&_Loan.TransactOpts, ids, amounts, providers)
}

// Init is a paid mutator transaction binding the contract method 0x06552ff3.
//
// Solidity: function init(address owner, address caller, address usdt, address lp) returns()
func (_Loan *LoanTransactor) Init(opts *bind.TransactOpts, owner common.Address, caller common.Address, usdt common.Address, lp common.Address) (*types.Transaction, error) {
	return _Loan.contract.Transact(opts, "init", owner, caller, usdt, lp)
}

// Init is a paid mutator transaction binding the contract method 0x06552ff3.
//
// Solidity: function init(address owner, address caller, address usdt, address lp) returns()
func (_Loan *LoanSession) Init(owner common.Address, caller common.Address, usdt common.Address, lp common.Address) (*types.Transaction, error) {
	return _Loan.Contract.Init(&_Loan.TransactOpts, owner, caller, usdt, lp)
}

// Init is a paid mutator transaction binding the contract method 0x06552ff3.
//
// Solidity: function init(address owner, address caller, address usdt, address lp) returns()
func (_Loan *LoanTransactorSession) Init(owner common.Address, caller common.Address, usdt common.Address, lp common.Address) (*types.Transaction, error) {
	return _Loan.Contract.Init(&_Loan.TransactOpts, owner, caller, usdt, lp)
}

// PayBack is a paid mutator transaction binding the contract method 0x7ccc5d35.
//
// Solidity: function payBack(uint256 loanId) returns()
func (_Loan *LoanTransactor) PayBack(opts *bind.TransactOpts, loanId *big.Int) (*types.Transaction, error) {
	return _Loan.contract.Transact(opts, "payBack", loanId)
}

// PayBack is a paid mutator transaction binding the contract method 0x7ccc5d35.
//
// Solidity: function payBack(uint256 loanId) returns()
func (_Loan *LoanSession) PayBack(loanId *big.Int) (*types.Transaction, error) {
	return _Loan.Contract.PayBack(&_Loan.TransactOpts, loanId)
}

// PayBack is a paid mutator transaction binding the contract method 0x7ccc5d35.
//
// Solidity: function payBack(uint256 loanId) returns()
func (_Loan *LoanTransactorSession) PayBack(loanId *big.Int) (*types.Transaction, error) {
	return _Loan.Contract.PayBack(&_Loan.TransactOpts, loanId)
}

// ProvideUsdt is a paid mutator transaction binding the contract method 0xff47eeab.
//
// Solidity: function provideUsdt(uint256 amount, uint256 duration) returns()
func (_Loan *LoanTransactor) ProvideUsdt(opts *bind.TransactOpts, amount *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Loan.contract.Transact(opts, "provideUsdt", amount, duration)
}

// ProvideUsdt is a paid mutator transaction binding the contract method 0xff47eeab.
//
// Solidity: function provideUsdt(uint256 amount, uint256 duration) returns()
func (_Loan *LoanSession) ProvideUsdt(amount *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Loan.Contract.ProvideUsdt(&_Loan.TransactOpts, amount, duration)
}

// ProvideUsdt is a paid mutator transaction binding the contract method 0xff47eeab.
//
// Solidity: function provideUsdt(uint256 amount, uint256 duration) returns()
func (_Loan *LoanTransactorSession) ProvideUsdt(amount *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Loan.Contract.ProvideUsdt(&_Loan.TransactOpts, amount, duration)
}

// ReleaseLiquidReward is a paid mutator transaction binding the contract method 0x19b11619.
//
// Solidity: function releaseLiquidReward() returns()
func (_Loan *LoanTransactor) ReleaseLiquidReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Loan.contract.Transact(opts, "releaseLiquidReward")
}

// ReleaseLiquidReward is a paid mutator transaction binding the contract method 0x19b11619.
//
// Solidity: function releaseLiquidReward() returns()
func (_Loan *LoanSession) ReleaseLiquidReward() (*types.Transaction, error) {
	return _Loan.Contract.ReleaseLiquidReward(&_Loan.TransactOpts)
}

// ReleaseLiquidReward is a paid mutator transaction binding the contract method 0x19b11619.
//
// Solidity: function releaseLiquidReward() returns()
func (_Loan *LoanTransactorSession) ReleaseLiquidReward() (*types.Transaction, error) {
	return _Loan.Contract.ReleaseLiquidReward(&_Loan.TransactOpts)
}

// RetrieveUsdt is a paid mutator transaction binding the contract method 0x952932b8.
//
// Solidity: function retrieveUsdt(uint256 provideId) returns()
func (_Loan *LoanTransactor) RetrieveUsdt(opts *bind.TransactOpts, provideId *big.Int) (*types.Transaction, error) {
	return _Loan.contract.Transact(opts, "retrieveUsdt", provideId)
}

// RetrieveUsdt is a paid mutator transaction binding the contract method 0x952932b8.
//
// Solidity: function retrieveUsdt(uint256 provideId) returns()
func (_Loan *LoanSession) RetrieveUsdt(provideId *big.Int) (*types.Transaction, error) {
	return _Loan.Contract.RetrieveUsdt(&_Loan.TransactOpts, provideId)
}

// RetrieveUsdt is a paid mutator transaction binding the contract method 0x952932b8.
//
// Solidity: function retrieveUsdt(uint256 provideId) returns()
func (_Loan *LoanTransactorSession) RetrieveUsdt(provideId *big.Int) (*types.Transaction, error) {
	return _Loan.Contract.RetrieveUsdt(&_Loan.TransactOpts, provideId)
}

// SetLpContract is a paid mutator transaction binding the contract method 0x40c0ac2d.
//
// Solidity: function setLpContract(address lp) returns()
func (_Loan *LoanTransactor) SetLpContract(opts *bind.TransactOpts, lp common.Address) (*types.Transaction, error) {
	return _Loan.contract.Transact(opts, "setLpContract", lp)
}

// SetLpContract is a paid mutator transaction binding the contract method 0x40c0ac2d.
//
// Solidity: function setLpContract(address lp) returns()
func (_Loan *LoanSession) SetLpContract(lp common.Address) (*types.Transaction, error) {
	return _Loan.Contract.SetLpContract(&_Loan.TransactOpts, lp)
}

// SetLpContract is a paid mutator transaction binding the contract method 0x40c0ac2d.
//
// Solidity: function setLpContract(address lp) returns()
func (_Loan *LoanTransactorSession) SetLpContract(lp common.Address) (*types.Transaction, error) {
	return _Loan.Contract.SetLpContract(&_Loan.TransactOpts, lp)
}

// SetParams is a paid mutator transaction binding the contract method 0xc0324c77.
//
// Solidity: function setParams(uint256 key, uint256 value) returns()
func (_Loan *LoanTransactor) SetParams(opts *bind.TransactOpts, key *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Loan.contract.Transact(opts, "setParams", key, value)
}

// SetParams is a paid mutator transaction binding the contract method 0xc0324c77.
//
// Solidity: function setParams(uint256 key, uint256 value) returns()
func (_Loan *LoanSession) SetParams(key *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Loan.Contract.SetParams(&_Loan.TransactOpts, key, value)
}

// SetParams is a paid mutator transaction binding the contract method 0xc0324c77.
//
// Solidity: function setParams(uint256 key, uint256 value) returns()
func (_Loan *LoanTransactorSession) SetParams(key *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Loan.Contract.SetParams(&_Loan.TransactOpts, key, value)
}

// SetUsdtContract is a paid mutator transaction binding the contract method 0x576338a5.
//
// Solidity: function setUsdtContract(address usdt) returns()
func (_Loan *LoanTransactor) SetUsdtContract(opts *bind.TransactOpts, usdt common.Address) (*types.Transaction, error) {
	return _Loan.contract.Transact(opts, "setUsdtContract", usdt)
}

// SetUsdtContract is a paid mutator transaction binding the contract method 0x576338a5.
//
// Solidity: function setUsdtContract(address usdt) returns()
func (_Loan *LoanSession) SetUsdtContract(usdt common.Address) (*types.Transaction, error) {
	return _Loan.Contract.SetUsdtContract(&_Loan.TransactOpts, usdt)
}

// SetUsdtContract is a paid mutator transaction binding the contract method 0x576338a5.
//
// Solidity: function setUsdtContract(address usdt) returns()
func (_Loan *LoanTransactorSession) SetUsdtContract(usdt common.Address) (*types.Transaction, error) {
	return _Loan.Contract.SetUsdtContract(&_Loan.TransactOpts, usdt)
}

// TransferCaller is a paid mutator transaction binding the contract method 0x4c2cc01d.
//
// Solidity: function transferCaller(address caller) returns()
func (_Loan *LoanTransactor) TransferCaller(opts *bind.TransactOpts, caller common.Address) (*types.Transaction, error) {
	return _Loan.contract.Transact(opts, "transferCaller", caller)
}

// TransferCaller is a paid mutator transaction binding the contract method 0x4c2cc01d.
//
// Solidity: function transferCaller(address caller) returns()
func (_Loan *LoanSession) TransferCaller(caller common.Address) (*types.Transaction, error) {
	return _Loan.Contract.TransferCaller(&_Loan.TransactOpts, caller)
}

// TransferCaller is a paid mutator transaction binding the contract method 0x4c2cc01d.
//
// Solidity: function transferCaller(address caller) returns()
func (_Loan *LoanTransactorSession) TransferCaller(caller common.Address) (*types.Transaction, error) {
	return _Loan.Contract.TransferCaller(&_Loan.TransactOpts, caller)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(address owner) returns()
func (_Loan *LoanTransactor) TransferOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _Loan.contract.Transact(opts, "transferOwner", owner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(address owner) returns()
func (_Loan *LoanSession) TransferOwner(owner common.Address) (*types.Transaction, error) {
	return _Loan.Contract.TransferOwner(&_Loan.TransactOpts, owner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(address owner) returns()
func (_Loan *LoanTransactorSession) TransferOwner(owner common.Address) (*types.Transaction, error) {
	return _Loan.Contract.TransferOwner(&_Loan.TransactOpts, owner)
}

// LoanEventClearIterator is returned from FilterEventClear and is used to iterate over the raw logs and unpacked data for EventClear events raised by the Loan contract.
type LoanEventClearIterator struct {
	Event *LoanEventClear // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanEventClearIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanEventClear)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanEventClear)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanEventClearIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanEventClearIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanEventClear represents a EventClear event raised by the Loan contract.
type LoanEventClear struct {
	LoanId *big.Int
	Amount *big.Int
	Loaner common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEventClear is a free log retrieval operation binding the contract event 0x5a322d6a1b1ff3f692c4c5995ba8133271b684bd011bc8e2f711d50db23bfe03.
//
// Solidity: event eventClear(uint256 loanId, uint256 amount, address loaner)
func (_Loan *LoanFilterer) FilterEventClear(opts *bind.FilterOpts) (*LoanEventClearIterator, error) {

	logs, sub, err := _Loan.contract.FilterLogs(opts, "eventClear")
	if err != nil {
		return nil, err
	}
	return &LoanEventClearIterator{contract: _Loan.contract, event: "eventClear", logs: logs, sub: sub}, nil
}

// WatchEventClear is a free log subscription operation binding the contract event 0x5a322d6a1b1ff3f692c4c5995ba8133271b684bd011bc8e2f711d50db23bfe03.
//
// Solidity: event eventClear(uint256 loanId, uint256 amount, address loaner)
func (_Loan *LoanFilterer) WatchEventClear(opts *bind.WatchOpts, sink chan<- *LoanEventClear) (event.Subscription, error) {

	logs, sub, err := _Loan.contract.WatchLogs(opts, "eventClear")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanEventClear)
				if err := _Loan.contract.UnpackLog(event, "eventClear", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEventClear is a log parse operation binding the contract event 0x5a322d6a1b1ff3f692c4c5995ba8133271b684bd011bc8e2f711d50db23bfe03.
//
// Solidity: event eventClear(uint256 loanId, uint256 amount, address loaner)
func (_Loan *LoanFilterer) ParseEventClear(log types.Log) (*LoanEventClear, error) {
	event := new(LoanEventClear)
	if err := _Loan.contract.UnpackLog(event, "eventClear", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanEventExchangeLpToUsdtIterator is returned from FilterEventExchangeLpToUsdt and is used to iterate over the raw logs and unpacked data for EventExchangeLpToUsdt events raised by the Loan contract.
type LoanEventExchangeLpToUsdtIterator struct {
	Event *LoanEventExchangeLpToUsdt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanEventExchangeLpToUsdtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanEventExchangeLpToUsdt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanEventExchangeLpToUsdt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanEventExchangeLpToUsdtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanEventExchangeLpToUsdtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanEventExchangeLpToUsdt represents a EventExchangeLpToUsdt event raised by the Loan contract.
type LoanEventExchangeLpToUsdt struct {
	Forward bool
	Amount  *big.Int
	Caller  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEventExchangeLpToUsdt is a free log retrieval operation binding the contract event 0x1a83ec821f4483d6557a8f7a974b0861f82494f76a8589611ab01a019489b491.
//
// Solidity: event eventExchangeLpToUsdt(bool forward, uint256 amount, address caller)
func (_Loan *LoanFilterer) FilterEventExchangeLpToUsdt(opts *bind.FilterOpts) (*LoanEventExchangeLpToUsdtIterator, error) {

	logs, sub, err := _Loan.contract.FilterLogs(opts, "eventExchangeLpToUsdt")
	if err != nil {
		return nil, err
	}
	return &LoanEventExchangeLpToUsdtIterator{contract: _Loan.contract, event: "eventExchangeLpToUsdt", logs: logs, sub: sub}, nil
}

// WatchEventExchangeLpToUsdt is a free log subscription operation binding the contract event 0x1a83ec821f4483d6557a8f7a974b0861f82494f76a8589611ab01a019489b491.
//
// Solidity: event eventExchangeLpToUsdt(bool forward, uint256 amount, address caller)
func (_Loan *LoanFilterer) WatchEventExchangeLpToUsdt(opts *bind.WatchOpts, sink chan<- *LoanEventExchangeLpToUsdt) (event.Subscription, error) {

	logs, sub, err := _Loan.contract.WatchLogs(opts, "eventExchangeLpToUsdt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanEventExchangeLpToUsdt)
				if err := _Loan.contract.UnpackLog(event, "eventExchangeLpToUsdt", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEventExchangeLpToUsdt is a log parse operation binding the contract event 0x1a83ec821f4483d6557a8f7a974b0861f82494f76a8589611ab01a019489b491.
//
// Solidity: event eventExchangeLpToUsdt(bool forward, uint256 amount, address caller)
func (_Loan *LoanFilterer) ParseEventExchangeLpToUsdt(log types.Log) (*LoanEventExchangeLpToUsdt, error) {
	event := new(LoanEventExchangeLpToUsdt)
	if err := _Loan.contract.UnpackLog(event, "eventExchangeLpToUsdt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanEventIncreaseLiquidRewardIterator is returned from FilterEventIncreaseLiquidReward and is used to iterate over the raw logs and unpacked data for EventIncreaseLiquidReward events raised by the Loan contract.
type LoanEventIncreaseLiquidRewardIterator struct {
	Event *LoanEventIncreaseLiquidReward // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanEventIncreaseLiquidRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanEventIncreaseLiquidReward)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanEventIncreaseLiquidReward)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanEventIncreaseLiquidRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanEventIncreaseLiquidRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanEventIncreaseLiquidReward represents a EventIncreaseLiquidReward event raised by the Loan contract.
type LoanEventIncreaseLiquidReward struct {
	Amount   *big.Int
	Provider common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEventIncreaseLiquidReward is a free log retrieval operation binding the contract event 0x045492b29efb11990e42d3b9ba88978042183c94051194b886f8595139c64db9.
//
// Solidity: event eventIncreaseLiquidReward(uint256 amount, address provider)
func (_Loan *LoanFilterer) FilterEventIncreaseLiquidReward(opts *bind.FilterOpts) (*LoanEventIncreaseLiquidRewardIterator, error) {

	logs, sub, err := _Loan.contract.FilterLogs(opts, "eventIncreaseLiquidReward")
	if err != nil {
		return nil, err
	}
	return &LoanEventIncreaseLiquidRewardIterator{contract: _Loan.contract, event: "eventIncreaseLiquidReward", logs: logs, sub: sub}, nil
}

// WatchEventIncreaseLiquidReward is a free log subscription operation binding the contract event 0x045492b29efb11990e42d3b9ba88978042183c94051194b886f8595139c64db9.
//
// Solidity: event eventIncreaseLiquidReward(uint256 amount, address provider)
func (_Loan *LoanFilterer) WatchEventIncreaseLiquidReward(opts *bind.WatchOpts, sink chan<- *LoanEventIncreaseLiquidReward) (event.Subscription, error) {

	logs, sub, err := _Loan.contract.WatchLogs(opts, "eventIncreaseLiquidReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanEventIncreaseLiquidReward)
				if err := _Loan.contract.UnpackLog(event, "eventIncreaseLiquidReward", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEventIncreaseLiquidReward is a log parse operation binding the contract event 0x045492b29efb11990e42d3b9ba88978042183c94051194b886f8595139c64db9.
//
// Solidity: event eventIncreaseLiquidReward(uint256 amount, address provider)
func (_Loan *LoanFilterer) ParseEventIncreaseLiquidReward(log types.Log) (*LoanEventIncreaseLiquidReward, error) {
	event := new(LoanEventIncreaseLiquidReward)
	if err := _Loan.contract.UnpackLog(event, "eventIncreaseLiquidReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanEventIncreaseLiquidRewardBathIterator is returned from FilterEventIncreaseLiquidRewardBath and is used to iterate over the raw logs and unpacked data for EventIncreaseLiquidRewardBath events raised by the Loan contract.
type LoanEventIncreaseLiquidRewardBathIterator struct {
	Event *LoanEventIncreaseLiquidRewardBath // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanEventIncreaseLiquidRewardBathIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanEventIncreaseLiquidRewardBath)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanEventIncreaseLiquidRewardBath)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanEventIncreaseLiquidRewardBathIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanEventIncreaseLiquidRewardBathIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanEventIncreaseLiquidRewardBath represents a EventIncreaseLiquidRewardBath event raised by the Loan contract.
type LoanEventIncreaseLiquidRewardBath struct {
	Ids       []*big.Int
	Amounts   []*big.Int
	Providers []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEventIncreaseLiquidRewardBath is a free log retrieval operation binding the contract event 0xa16a887e0d16d473c5a8459cbf20c45ef7f0a282e60e60adcacb455ae31ebb62.
//
// Solidity: event eventIncreaseLiquidRewardBath(uint256[] ids, uint256[] amounts, address[] providers)
func (_Loan *LoanFilterer) FilterEventIncreaseLiquidRewardBath(opts *bind.FilterOpts) (*LoanEventIncreaseLiquidRewardBathIterator, error) {

	logs, sub, err := _Loan.contract.FilterLogs(opts, "eventIncreaseLiquidRewardBath")
	if err != nil {
		return nil, err
	}
	return &LoanEventIncreaseLiquidRewardBathIterator{contract: _Loan.contract, event: "eventIncreaseLiquidRewardBath", logs: logs, sub: sub}, nil
}

// WatchEventIncreaseLiquidRewardBath is a free log subscription operation binding the contract event 0xa16a887e0d16d473c5a8459cbf20c45ef7f0a282e60e60adcacb455ae31ebb62.
//
// Solidity: event eventIncreaseLiquidRewardBath(uint256[] ids, uint256[] amounts, address[] providers)
func (_Loan *LoanFilterer) WatchEventIncreaseLiquidRewardBath(opts *bind.WatchOpts, sink chan<- *LoanEventIncreaseLiquidRewardBath) (event.Subscription, error) {

	logs, sub, err := _Loan.contract.WatchLogs(opts, "eventIncreaseLiquidRewardBath")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanEventIncreaseLiquidRewardBath)
				if err := _Loan.contract.UnpackLog(event, "eventIncreaseLiquidRewardBath", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEventIncreaseLiquidRewardBath is a log parse operation binding the contract event 0xa16a887e0d16d473c5a8459cbf20c45ef7f0a282e60e60adcacb455ae31ebb62.
//
// Solidity: event eventIncreaseLiquidRewardBath(uint256[] ids, uint256[] amounts, address[] providers)
func (_Loan *LoanFilterer) ParseEventIncreaseLiquidRewardBath(log types.Log) (*LoanEventIncreaseLiquidRewardBath, error) {
	event := new(LoanEventIncreaseLiquidRewardBath)
	if err := _Loan.contract.UnpackLog(event, "eventIncreaseLiquidRewardBath", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanEventNewLoanIterator is returned from FilterEventNewLoan and is used to iterate over the raw logs and unpacked data for EventNewLoan events raised by the Loan contract.
type LoanEventNewLoanIterator struct {
	Event *LoanEventNewLoan // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanEventNewLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanEventNewLoan)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanEventNewLoan)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanEventNewLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanEventNewLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanEventNewLoan represents a EventNewLoan event raised by the Loan contract.
type LoanEventNewLoan struct {
	LoanId         *big.Int
	Duration       *big.Int
	Start          *big.Int
	LoanAmount     *big.Int
	Loaner         common.Address
	ReleaseAmount  *big.Int
	InterestAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterEventNewLoan is a free log retrieval operation binding the contract event 0x8359c828396108eedea00704782ac2a600d822d6d56312c4e10f62408aedca5d.
//
// Solidity: event eventNewLoan(uint256 loanId, uint256 duration, uint256 start, uint256 loanAmount, address loaner, uint256 releaseAmount, uint256 interestAmount)
func (_Loan *LoanFilterer) FilterEventNewLoan(opts *bind.FilterOpts) (*LoanEventNewLoanIterator, error) {

	logs, sub, err := _Loan.contract.FilterLogs(opts, "eventNewLoan")
	if err != nil {
		return nil, err
	}
	return &LoanEventNewLoanIterator{contract: _Loan.contract, event: "eventNewLoan", logs: logs, sub: sub}, nil
}

// WatchEventNewLoan is a free log subscription operation binding the contract event 0x8359c828396108eedea00704782ac2a600d822d6d56312c4e10f62408aedca5d.
//
// Solidity: event eventNewLoan(uint256 loanId, uint256 duration, uint256 start, uint256 loanAmount, address loaner, uint256 releaseAmount, uint256 interestAmount)
func (_Loan *LoanFilterer) WatchEventNewLoan(opts *bind.WatchOpts, sink chan<- *LoanEventNewLoan) (event.Subscription, error) {

	logs, sub, err := _Loan.contract.WatchLogs(opts, "eventNewLoan")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanEventNewLoan)
				if err := _Loan.contract.UnpackLog(event, "eventNewLoan", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEventNewLoan is a log parse operation binding the contract event 0x8359c828396108eedea00704782ac2a600d822d6d56312c4e10f62408aedca5d.
//
// Solidity: event eventNewLoan(uint256 loanId, uint256 duration, uint256 start, uint256 loanAmount, address loaner, uint256 releaseAmount, uint256 interestAmount)
func (_Loan *LoanFilterer) ParseEventNewLoan(log types.Log) (*LoanEventNewLoan, error) {
	event := new(LoanEventNewLoan)
	if err := _Loan.contract.UnpackLog(event, "eventNewLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanEventPayBackIterator is returned from FilterEventPayBack and is used to iterate over the raw logs and unpacked data for EventPayBack events raised by the Loan contract.
type LoanEventPayBackIterator struct {
	Event *LoanEventPayBack // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanEventPayBackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanEventPayBack)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanEventPayBack)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanEventPayBackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanEventPayBackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanEventPayBack represents a EventPayBack event raised by the Loan contract.
type LoanEventPayBack struct {
	LoanId *big.Int
	Amount *big.Int
	Loaner common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEventPayBack is a free log retrieval operation binding the contract event 0x38dcb8e7ce8c7f182d53142ee0fa94a1778cd54eddcc1209f86469e7d3b48733.
//
// Solidity: event eventPayBack(uint256 loanId, uint256 amount, address loaner)
func (_Loan *LoanFilterer) FilterEventPayBack(opts *bind.FilterOpts) (*LoanEventPayBackIterator, error) {

	logs, sub, err := _Loan.contract.FilterLogs(opts, "eventPayBack")
	if err != nil {
		return nil, err
	}
	return &LoanEventPayBackIterator{contract: _Loan.contract, event: "eventPayBack", logs: logs, sub: sub}, nil
}

// WatchEventPayBack is a free log subscription operation binding the contract event 0x38dcb8e7ce8c7f182d53142ee0fa94a1778cd54eddcc1209f86469e7d3b48733.
//
// Solidity: event eventPayBack(uint256 loanId, uint256 amount, address loaner)
func (_Loan *LoanFilterer) WatchEventPayBack(opts *bind.WatchOpts, sink chan<- *LoanEventPayBack) (event.Subscription, error) {

	logs, sub, err := _Loan.contract.WatchLogs(opts, "eventPayBack")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanEventPayBack)
				if err := _Loan.contract.UnpackLog(event, "eventPayBack", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEventPayBack is a log parse operation binding the contract event 0x38dcb8e7ce8c7f182d53142ee0fa94a1778cd54eddcc1209f86469e7d3b48733.
//
// Solidity: event eventPayBack(uint256 loanId, uint256 amount, address loaner)
func (_Loan *LoanFilterer) ParseEventPayBack(log types.Log) (*LoanEventPayBack, error) {
	event := new(LoanEventPayBack)
	if err := _Loan.contract.UnpackLog(event, "eventPayBack", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanEventProviderAddIterator is returned from FilterEventProviderAdd and is used to iterate over the raw logs and unpacked data for EventProviderAdd events raised by the Loan contract.
type LoanEventProviderAddIterator struct {
	Event *LoanEventProviderAdd // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanEventProviderAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanEventProviderAdd)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanEventProviderAdd)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanEventProviderAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanEventProviderAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanEventProviderAdd represents a EventProviderAdd event raised by the Loan contract.
type LoanEventProviderAdd struct {
	Id       *big.Int
	Duration *big.Int
	Start    *big.Int
	Amount   *big.Int
	Provider common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEventProviderAdd is a free log retrieval operation binding the contract event 0x75780a70131ef5cf8aff25941e13a743681d34e1eb85abde44b32e09280e1fcc.
//
// Solidity: event eventProviderAdd(uint256 id, uint256 duration, uint256 start, uint256 amount, address provider)
func (_Loan *LoanFilterer) FilterEventProviderAdd(opts *bind.FilterOpts) (*LoanEventProviderAddIterator, error) {

	logs, sub, err := _Loan.contract.FilterLogs(opts, "eventProviderAdd")
	if err != nil {
		return nil, err
	}
	return &LoanEventProviderAddIterator{contract: _Loan.contract, event: "eventProviderAdd", logs: logs, sub: sub}, nil
}

// WatchEventProviderAdd is a free log subscription operation binding the contract event 0x75780a70131ef5cf8aff25941e13a743681d34e1eb85abde44b32e09280e1fcc.
//
// Solidity: event eventProviderAdd(uint256 id, uint256 duration, uint256 start, uint256 amount, address provider)
func (_Loan *LoanFilterer) WatchEventProviderAdd(opts *bind.WatchOpts, sink chan<- *LoanEventProviderAdd) (event.Subscription, error) {

	logs, sub, err := _Loan.contract.WatchLogs(opts, "eventProviderAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanEventProviderAdd)
				if err := _Loan.contract.UnpackLog(event, "eventProviderAdd", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEventProviderAdd is a log parse operation binding the contract event 0x75780a70131ef5cf8aff25941e13a743681d34e1eb85abde44b32e09280e1fcc.
//
// Solidity: event eventProviderAdd(uint256 id, uint256 duration, uint256 start, uint256 amount, address provider)
func (_Loan *LoanFilterer) ParseEventProviderAdd(log types.Log) (*LoanEventProviderAdd, error) {
	event := new(LoanEventProviderAdd)
	if err := _Loan.contract.UnpackLog(event, "eventProviderAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanEventProviderRedeemIterator is returned from FilterEventProviderRedeem and is used to iterate over the raw logs and unpacked data for EventProviderRedeem events raised by the Loan contract.
type LoanEventProviderRedeemIterator struct {
	Event *LoanEventProviderRedeem // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanEventProviderRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanEventProviderRedeem)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanEventProviderRedeem)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanEventProviderRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanEventProviderRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanEventProviderRedeem represents a EventProviderRedeem event raised by the Loan contract.
type LoanEventProviderRedeem struct {
	Id       *big.Int
	Amount   *big.Int
	Provider common.Address
	Fee      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEventProviderRedeem is a free log retrieval operation binding the contract event 0x1ec3add8915e0172b379ff4433663fbf4a45d4da64edb3e0402fb369e04b024a.
//
// Solidity: event eventProviderRedeem(uint256 id, uint256 amount, address provider, uint256 fee)
func (_Loan *LoanFilterer) FilterEventProviderRedeem(opts *bind.FilterOpts) (*LoanEventProviderRedeemIterator, error) {

	logs, sub, err := _Loan.contract.FilterLogs(opts, "eventProviderRedeem")
	if err != nil {
		return nil, err
	}
	return &LoanEventProviderRedeemIterator{contract: _Loan.contract, event: "eventProviderRedeem", logs: logs, sub: sub}, nil
}

// WatchEventProviderRedeem is a free log subscription operation binding the contract event 0x1ec3add8915e0172b379ff4433663fbf4a45d4da64edb3e0402fb369e04b024a.
//
// Solidity: event eventProviderRedeem(uint256 id, uint256 amount, address provider, uint256 fee)
func (_Loan *LoanFilterer) WatchEventProviderRedeem(opts *bind.WatchOpts, sink chan<- *LoanEventProviderRedeem) (event.Subscription, error) {

	logs, sub, err := _Loan.contract.WatchLogs(opts, "eventProviderRedeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanEventProviderRedeem)
				if err := _Loan.contract.UnpackLog(event, "eventProviderRedeem", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEventProviderRedeem is a log parse operation binding the contract event 0x1ec3add8915e0172b379ff4433663fbf4a45d4da64edb3e0402fb369e04b024a.
//
// Solidity: event eventProviderRedeem(uint256 id, uint256 amount, address provider, uint256 fee)
func (_Loan *LoanFilterer) ParseEventProviderRedeem(log types.Log) (*LoanEventProviderRedeem, error) {
	event := new(LoanEventProviderRedeem)
	if err := _Loan.contract.UnpackLog(event, "eventProviderRedeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanEventReleaseLiquidRewardIterator is returned from FilterEventReleaseLiquidReward and is used to iterate over the raw logs and unpacked data for EventReleaseLiquidReward events raised by the Loan contract.
type LoanEventReleaseLiquidRewardIterator struct {
	Event *LoanEventReleaseLiquidReward // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanEventReleaseLiquidRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanEventReleaseLiquidReward)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanEventReleaseLiquidReward)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanEventReleaseLiquidRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanEventReleaseLiquidRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanEventReleaseLiquidReward represents a EventReleaseLiquidReward event raised by the Loan contract.
type LoanEventReleaseLiquidReward struct {
	Amount   *big.Int
	Provider common.Address
	Fee      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEventReleaseLiquidReward is a free log retrieval operation binding the contract event 0x9d55a88ba6edf4a14f0ad37d9f0833bb65734beea749cfeff8d52825ffd58ef9.
//
// Solidity: event eventReleaseLiquidReward(uint256 amount, address provider, uint256 fee)
func (_Loan *LoanFilterer) FilterEventReleaseLiquidReward(opts *bind.FilterOpts) (*LoanEventReleaseLiquidRewardIterator, error) {

	logs, sub, err := _Loan.contract.FilterLogs(opts, "eventReleaseLiquidReward")
	if err != nil {
		return nil, err
	}
	return &LoanEventReleaseLiquidRewardIterator{contract: _Loan.contract, event: "eventReleaseLiquidReward", logs: logs, sub: sub}, nil
}

// WatchEventReleaseLiquidReward is a free log subscription operation binding the contract event 0x9d55a88ba6edf4a14f0ad37d9f0833bb65734beea749cfeff8d52825ffd58ef9.
//
// Solidity: event eventReleaseLiquidReward(uint256 amount, address provider, uint256 fee)
func (_Loan *LoanFilterer) WatchEventReleaseLiquidReward(opts *bind.WatchOpts, sink chan<- *LoanEventReleaseLiquidReward) (event.Subscription, error) {

	logs, sub, err := _Loan.contract.WatchLogs(opts, "eventReleaseLiquidReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanEventReleaseLiquidReward)
				if err := _Loan.contract.UnpackLog(event, "eventReleaseLiquidReward", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEventReleaseLiquidReward is a log parse operation binding the contract event 0x9d55a88ba6edf4a14f0ad37d9f0833bb65734beea749cfeff8d52825ffd58ef9.
//
// Solidity: event eventReleaseLiquidReward(uint256 amount, address provider, uint256 fee)
func (_Loan *LoanFilterer) ParseEventReleaseLiquidReward(log types.Log) (*LoanEventReleaseLiquidReward, error) {
	event := new(LoanEventReleaseLiquidReward)
	if err := _Loan.contract.UnpackLog(event, "eventReleaseLiquidReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
