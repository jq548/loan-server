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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"loaner\",\"type\":\"address\"}],\"name\":\"eventClear\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"eventIncreaseLiquidReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"loaner\",\"type\":\"address\"}],\"name\":\"eventNewLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"loaner\",\"type\":\"address\"}],\"name\":\"eventPayBack\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"eventProviderIncrease\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"eventProviderRetrieve\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"eventReleaseLiquidReward\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"loaner\",\"type\":\"address\"}],\"name\":\"addNewLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"}],\"name\":\"clear\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"forward\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"exchangeLpUsdt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"extract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"increaseLiquidReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usdt\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lp\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"loans\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"loaner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"forward\",\"type\":\"bool\"}],\"name\":\"maxExchangeLpUsdt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"}],\"name\":\"payBack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"provideUsdt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"releaseAbleLiquidReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"releaseLiquidReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"retrieveUsdt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lp\",\"type\":\"address\"}],\"name\":\"setLpContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usdt\",\"type\":\"address\"}],\"name\":\"setUsdtContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"transferCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// Addresses is a free data retrieval call binding the contract method 0xedf26d9b.
//
// Solidity: function addresses(uint256 ) view returns(address)
func (_Contract *ContractCaller) Addresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "addresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addresses is a free data retrieval call binding the contract method 0xedf26d9b.
//
// Solidity: function addresses(uint256 ) view returns(address)
func (_Contract *ContractSession) Addresses(arg0 *big.Int) (common.Address, error) {
	return _Contract.Contract.Addresses(&_Contract.CallOpts, arg0)
}

// Addresses is a free data retrieval call binding the contract method 0xedf26d9b.
//
// Solidity: function addresses(uint256 ) view returns(address)
func (_Contract *ContractCallerSession) Addresses(arg0 *big.Int) (common.Address, error) {
	return _Contract.Contract.Addresses(&_Contract.CallOpts, arg0)
}

// Loans is a free data retrieval call binding the contract method 0xe1ec3c68.
//
// Solidity: function loans(uint256 ) view returns(uint256 amount, uint256 duration, uint256 start, uint256 status, address loaner)
func (_Contract *ContractCaller) Loans(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount   *big.Int
	Duration *big.Int
	Start    *big.Int
	Status   *big.Int
	Loaner   common.Address
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "loans", arg0)

	outstruct := new(struct {
		Amount   *big.Int
		Duration *big.Int
		Start    *big.Int
		Status   *big.Int
		Loaner   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Start = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Loaner = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Loans is a free data retrieval call binding the contract method 0xe1ec3c68.
//
// Solidity: function loans(uint256 ) view returns(uint256 amount, uint256 duration, uint256 start, uint256 status, address loaner)
func (_Contract *ContractSession) Loans(arg0 *big.Int) (struct {
	Amount   *big.Int
	Duration *big.Int
	Start    *big.Int
	Status   *big.Int
	Loaner   common.Address
}, error) {
	return _Contract.Contract.Loans(&_Contract.CallOpts, arg0)
}

// Loans is a free data retrieval call binding the contract method 0xe1ec3c68.
//
// Solidity: function loans(uint256 ) view returns(uint256 amount, uint256 duration, uint256 start, uint256 status, address loaner)
func (_Contract *ContractCallerSession) Loans(arg0 *big.Int) (struct {
	Amount   *big.Int
	Duration *big.Int
	Start    *big.Int
	Status   *big.Int
	Loaner   common.Address
}, error) {
	return _Contract.Contract.Loans(&_Contract.CallOpts, arg0)
}

// MaxExchangeLpUsdt is a free data retrieval call binding the contract method 0xca938a62.
//
// Solidity: function maxExchangeLpUsdt(bool forward) view returns(uint256)
func (_Contract *ContractCaller) MaxExchangeLpUsdt(opts *bind.CallOpts, forward bool) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxExchangeLpUsdt", forward)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxExchangeLpUsdt is a free data retrieval call binding the contract method 0xca938a62.
//
// Solidity: function maxExchangeLpUsdt(bool forward) view returns(uint256)
func (_Contract *ContractSession) MaxExchangeLpUsdt(forward bool) (*big.Int, error) {
	return _Contract.Contract.MaxExchangeLpUsdt(&_Contract.CallOpts, forward)
}

// MaxExchangeLpUsdt is a free data retrieval call binding the contract method 0xca938a62.
//
// Solidity: function maxExchangeLpUsdt(bool forward) view returns(uint256)
func (_Contract *ContractCallerSession) MaxExchangeLpUsdt(forward bool) (*big.Int, error) {
	return _Contract.Contract.MaxExchangeLpUsdt(&_Contract.CallOpts, forward)
}

// ReleaseAbleLiquidReward is a free data retrieval call binding the contract method 0xae357555.
//
// Solidity: function releaseAbleLiquidReward(address provider) view returns(uint256)
func (_Contract *ContractCaller) ReleaseAbleLiquidReward(opts *bind.CallOpts, provider common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "releaseAbleLiquidReward", provider)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReleaseAbleLiquidReward is a free data retrieval call binding the contract method 0xae357555.
//
// Solidity: function releaseAbleLiquidReward(address provider) view returns(uint256)
func (_Contract *ContractSession) ReleaseAbleLiquidReward(provider common.Address) (*big.Int, error) {
	return _Contract.Contract.ReleaseAbleLiquidReward(&_Contract.CallOpts, provider)
}

// ReleaseAbleLiquidReward is a free data retrieval call binding the contract method 0xae357555.
//
// Solidity: function releaseAbleLiquidReward(address provider) view returns(uint256)
func (_Contract *ContractCallerSession) ReleaseAbleLiquidReward(provider common.Address) (*big.Int, error) {
	return _Contract.Contract.ReleaseAbleLiquidReward(&_Contract.CallOpts, provider)
}

// AddNewLoan is a paid mutator transaction binding the contract method 0xcc6c2483.
//
// Solidity: function addNewLoan(uint256 id, uint256 amount, uint256 duration, address loaner) returns()
func (_Contract *ContractTransactor) AddNewLoan(opts *bind.TransactOpts, id *big.Int, amount *big.Int, duration *big.Int, loaner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addNewLoan", id, amount, duration, loaner)
}

// AddNewLoan is a paid mutator transaction binding the contract method 0xcc6c2483.
//
// Solidity: function addNewLoan(uint256 id, uint256 amount, uint256 duration, address loaner) returns()
func (_Contract *ContractSession) AddNewLoan(id *big.Int, amount *big.Int, duration *big.Int, loaner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddNewLoan(&_Contract.TransactOpts, id, amount, duration, loaner)
}

// AddNewLoan is a paid mutator transaction binding the contract method 0xcc6c2483.
//
// Solidity: function addNewLoan(uint256 id, uint256 amount, uint256 duration, address loaner) returns()
func (_Contract *ContractTransactorSession) AddNewLoan(id *big.Int, amount *big.Int, duration *big.Int, loaner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddNewLoan(&_Contract.TransactOpts, id, amount, duration, loaner)
}

// Clear is a paid mutator transaction binding the contract method 0xc0fe1af8.
//
// Solidity: function clear(uint256 loanId) returns()
func (_Contract *ContractTransactor) Clear(opts *bind.TransactOpts, loanId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "clear", loanId)
}

// Clear is a paid mutator transaction binding the contract method 0xc0fe1af8.
//
// Solidity: function clear(uint256 loanId) returns()
func (_Contract *ContractSession) Clear(loanId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Clear(&_Contract.TransactOpts, loanId)
}

// Clear is a paid mutator transaction binding the contract method 0xc0fe1af8.
//
// Solidity: function clear(uint256 loanId) returns()
func (_Contract *ContractTransactorSession) Clear(loanId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Clear(&_Contract.TransactOpts, loanId)
}

// ExchangeLpUsdt is a paid mutator transaction binding the contract method 0x5f141556.
//
// Solidity: function exchangeLpUsdt(bool forward, uint256 amount) returns()
func (_Contract *ContractTransactor) ExchangeLpUsdt(opts *bind.TransactOpts, forward bool, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "exchangeLpUsdt", forward, amount)
}

// ExchangeLpUsdt is a paid mutator transaction binding the contract method 0x5f141556.
//
// Solidity: function exchangeLpUsdt(bool forward, uint256 amount) returns()
func (_Contract *ContractSession) ExchangeLpUsdt(forward bool, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ExchangeLpUsdt(&_Contract.TransactOpts, forward, amount)
}

// ExchangeLpUsdt is a paid mutator transaction binding the contract method 0x5f141556.
//
// Solidity: function exchangeLpUsdt(bool forward, uint256 amount) returns()
func (_Contract *ContractTransactorSession) ExchangeLpUsdt(forward bool, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ExchangeLpUsdt(&_Contract.TransactOpts, forward, amount)
}

// Extract is a paid mutator transaction binding the contract method 0xc7a5d285.
//
// Solidity: function extract(address token) returns()
func (_Contract *ContractTransactor) Extract(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "extract", token)
}

// Extract is a paid mutator transaction binding the contract method 0xc7a5d285.
//
// Solidity: function extract(address token) returns()
func (_Contract *ContractSession) Extract(token common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Extract(&_Contract.TransactOpts, token)
}

// Extract is a paid mutator transaction binding the contract method 0xc7a5d285.
//
// Solidity: function extract(address token) returns()
func (_Contract *ContractTransactorSession) Extract(token common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Extract(&_Contract.TransactOpts, token)
}

// IncreaseLiquidReward is a paid mutator transaction binding the contract method 0xbaafdf5d.
//
// Solidity: function increaseLiquidReward(uint256 amount, address provider) returns()
func (_Contract *ContractTransactor) IncreaseLiquidReward(opts *bind.TransactOpts, amount *big.Int, provider common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "increaseLiquidReward", amount, provider)
}

// IncreaseLiquidReward is a paid mutator transaction binding the contract method 0xbaafdf5d.
//
// Solidity: function increaseLiquidReward(uint256 amount, address provider) returns()
func (_Contract *ContractSession) IncreaseLiquidReward(amount *big.Int, provider common.Address) (*types.Transaction, error) {
	return _Contract.Contract.IncreaseLiquidReward(&_Contract.TransactOpts, amount, provider)
}

// IncreaseLiquidReward is a paid mutator transaction binding the contract method 0xbaafdf5d.
//
// Solidity: function increaseLiquidReward(uint256 amount, address provider) returns()
func (_Contract *ContractTransactorSession) IncreaseLiquidReward(amount *big.Int, provider common.Address) (*types.Transaction, error) {
	return _Contract.Contract.IncreaseLiquidReward(&_Contract.TransactOpts, amount, provider)
}

// Init is a paid mutator transaction binding the contract method 0x06552ff3.
//
// Solidity: function init(address owner, address caller, address usdt, address lp) returns()
func (_Contract *ContractTransactor) Init(opts *bind.TransactOpts, owner common.Address, caller common.Address, usdt common.Address, lp common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "init", owner, caller, usdt, lp)
}

// Init is a paid mutator transaction binding the contract method 0x06552ff3.
//
// Solidity: function init(address owner, address caller, address usdt, address lp) returns()
func (_Contract *ContractSession) Init(owner common.Address, caller common.Address, usdt common.Address, lp common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Init(&_Contract.TransactOpts, owner, caller, usdt, lp)
}

// Init is a paid mutator transaction binding the contract method 0x06552ff3.
//
// Solidity: function init(address owner, address caller, address usdt, address lp) returns()
func (_Contract *ContractTransactorSession) Init(owner common.Address, caller common.Address, usdt common.Address, lp common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Init(&_Contract.TransactOpts, owner, caller, usdt, lp)
}

// PayBack is a paid mutator transaction binding the contract method 0x7ccc5d35.
//
// Solidity: function payBack(uint256 loanId) returns()
func (_Contract *ContractTransactor) PayBack(opts *bind.TransactOpts, loanId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "payBack", loanId)
}

// PayBack is a paid mutator transaction binding the contract method 0x7ccc5d35.
//
// Solidity: function payBack(uint256 loanId) returns()
func (_Contract *ContractSession) PayBack(loanId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PayBack(&_Contract.TransactOpts, loanId)
}

// PayBack is a paid mutator transaction binding the contract method 0x7ccc5d35.
//
// Solidity: function payBack(uint256 loanId) returns()
func (_Contract *ContractTransactorSession) PayBack(loanId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PayBack(&_Contract.TransactOpts, loanId)
}

// ProvideUsdt is a paid mutator transaction binding the contract method 0xfa455e09.
//
// Solidity: function provideUsdt(uint256 amount) returns()
func (_Contract *ContractTransactor) ProvideUsdt(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "provideUsdt", amount)
}

// ProvideUsdt is a paid mutator transaction binding the contract method 0xfa455e09.
//
// Solidity: function provideUsdt(uint256 amount) returns()
func (_Contract *ContractSession) ProvideUsdt(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ProvideUsdt(&_Contract.TransactOpts, amount)
}

// ProvideUsdt is a paid mutator transaction binding the contract method 0xfa455e09.
//
// Solidity: function provideUsdt(uint256 amount) returns()
func (_Contract *ContractTransactorSession) ProvideUsdt(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ProvideUsdt(&_Contract.TransactOpts, amount)
}

// ReleaseLiquidReward is a paid mutator transaction binding the contract method 0x19b11619.
//
// Solidity: function releaseLiquidReward() returns()
func (_Contract *ContractTransactor) ReleaseLiquidReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "releaseLiquidReward")
}

// ReleaseLiquidReward is a paid mutator transaction binding the contract method 0x19b11619.
//
// Solidity: function releaseLiquidReward() returns()
func (_Contract *ContractSession) ReleaseLiquidReward() (*types.Transaction, error) {
	return _Contract.Contract.ReleaseLiquidReward(&_Contract.TransactOpts)
}

// ReleaseLiquidReward is a paid mutator transaction binding the contract method 0x19b11619.
//
// Solidity: function releaseLiquidReward() returns()
func (_Contract *ContractTransactorSession) ReleaseLiquidReward() (*types.Transaction, error) {
	return _Contract.Contract.ReleaseLiquidReward(&_Contract.TransactOpts)
}

// RetrieveUsdt is a paid mutator transaction binding the contract method 0xc903cf67.
//
// Solidity: function retrieveUsdt() returns()
func (_Contract *ContractTransactor) RetrieveUsdt(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "retrieveUsdt")
}

// RetrieveUsdt is a paid mutator transaction binding the contract method 0xc903cf67.
//
// Solidity: function retrieveUsdt() returns()
func (_Contract *ContractSession) RetrieveUsdt() (*types.Transaction, error) {
	return _Contract.Contract.RetrieveUsdt(&_Contract.TransactOpts)
}

// RetrieveUsdt is a paid mutator transaction binding the contract method 0xc903cf67.
//
// Solidity: function retrieveUsdt() returns()
func (_Contract *ContractTransactorSession) RetrieveUsdt() (*types.Transaction, error) {
	return _Contract.Contract.RetrieveUsdt(&_Contract.TransactOpts)
}

// SetLpContract is a paid mutator transaction binding the contract method 0x40c0ac2d.
//
// Solidity: function setLpContract(address lp) returns()
func (_Contract *ContractTransactor) SetLpContract(opts *bind.TransactOpts, lp common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setLpContract", lp)
}

// SetLpContract is a paid mutator transaction binding the contract method 0x40c0ac2d.
//
// Solidity: function setLpContract(address lp) returns()
func (_Contract *ContractSession) SetLpContract(lp common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetLpContract(&_Contract.TransactOpts, lp)
}

// SetLpContract is a paid mutator transaction binding the contract method 0x40c0ac2d.
//
// Solidity: function setLpContract(address lp) returns()
func (_Contract *ContractTransactorSession) SetLpContract(lp common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetLpContract(&_Contract.TransactOpts, lp)
}

// SetUsdtContract is a paid mutator transaction binding the contract method 0x576338a5.
//
// Solidity: function setUsdtContract(address usdt) returns()
func (_Contract *ContractTransactor) SetUsdtContract(opts *bind.TransactOpts, usdt common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setUsdtContract", usdt)
}

// SetUsdtContract is a paid mutator transaction binding the contract method 0x576338a5.
//
// Solidity: function setUsdtContract(address usdt) returns()
func (_Contract *ContractSession) SetUsdtContract(usdt common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetUsdtContract(&_Contract.TransactOpts, usdt)
}

// SetUsdtContract is a paid mutator transaction binding the contract method 0x576338a5.
//
// Solidity: function setUsdtContract(address usdt) returns()
func (_Contract *ContractTransactorSession) SetUsdtContract(usdt common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetUsdtContract(&_Contract.TransactOpts, usdt)
}

// TransferCaller is a paid mutator transaction binding the contract method 0x4c2cc01d.
//
// Solidity: function transferCaller(address caller) returns()
func (_Contract *ContractTransactor) TransferCaller(opts *bind.TransactOpts, caller common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferCaller", caller)
}

// TransferCaller is a paid mutator transaction binding the contract method 0x4c2cc01d.
//
// Solidity: function transferCaller(address caller) returns()
func (_Contract *ContractSession) TransferCaller(caller common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferCaller(&_Contract.TransactOpts, caller)
}

// TransferCaller is a paid mutator transaction binding the contract method 0x4c2cc01d.
//
// Solidity: function transferCaller(address caller) returns()
func (_Contract *ContractTransactorSession) TransferCaller(caller common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferCaller(&_Contract.TransactOpts, caller)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(address owner) returns()
func (_Contract *ContractTransactor) TransferOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwner", owner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(address owner) returns()
func (_Contract *ContractSession) TransferOwner(owner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwner(&_Contract.TransactOpts, owner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(address owner) returns()
func (_Contract *ContractTransactorSession) TransferOwner(owner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwner(&_Contract.TransactOpts, owner)
}

// ContractEventClearIterator is returned from FilterEventClear and is used to iterate over the raw logs and unpacked data for EventClear events raised by the Contract contract.
type ContractEventClearIterator struct {
	Event *ContractEventClear // Event containing the contract specifics and raw log

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
func (it *ContractEventClearIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEventClear)
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
		it.Event = new(ContractEventClear)
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
func (it *ContractEventClearIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEventClearIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEventClear represents a EventClear event raised by the Contract contract.
type ContractEventClear struct {
	LoanId *big.Int
	Amount *big.Int
	Loaner common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEventClear is a free log retrieval operation binding the contract event 0x5a322d6a1b1ff3f692c4c5995ba8133271b684bd011bc8e2f711d50db23bfe03.
//
// Solidity: event eventClear(uint256 loanId, uint256 amount, address loaner)
func (_Contract *ContractFilterer) FilterEventClear(opts *bind.FilterOpts) (*ContractEventClearIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "eventClear")
	if err != nil {
		return nil, err
	}
	return &ContractEventClearIterator{contract: _Contract.contract, event: "eventClear", logs: logs, sub: sub}, nil
}

// WatchEventClear is a free log subscription operation binding the contract event 0x5a322d6a1b1ff3f692c4c5995ba8133271b684bd011bc8e2f711d50db23bfe03.
//
// Solidity: event eventClear(uint256 loanId, uint256 amount, address loaner)
func (_Contract *ContractFilterer) WatchEventClear(opts *bind.WatchOpts, sink chan<- *ContractEventClear) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "eventClear")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEventClear)
				if err := _Contract.contract.UnpackLog(event, "eventClear", log); err != nil {
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
func (_Contract *ContractFilterer) ParseEventClear(log types.Log) (*ContractEventClear, error) {
	event := new(ContractEventClear)
	if err := _Contract.contract.UnpackLog(event, "eventClear", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEventIncreaseLiquidRewardIterator is returned from FilterEventIncreaseLiquidReward and is used to iterate over the raw logs and unpacked data for EventIncreaseLiquidReward events raised by the Contract contract.
type ContractEventIncreaseLiquidRewardIterator struct {
	Event *ContractEventIncreaseLiquidReward // Event containing the contract specifics and raw log

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
func (it *ContractEventIncreaseLiquidRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEventIncreaseLiquidReward)
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
		it.Event = new(ContractEventIncreaseLiquidReward)
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
func (it *ContractEventIncreaseLiquidRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEventIncreaseLiquidRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEventIncreaseLiquidReward represents a EventIncreaseLiquidReward event raised by the Contract contract.
type ContractEventIncreaseLiquidReward struct {
	Amount   *big.Int
	Provider common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEventIncreaseLiquidReward is a free log retrieval operation binding the contract event 0x045492b29efb11990e42d3b9ba88978042183c94051194b886f8595139c64db9.
//
// Solidity: event eventIncreaseLiquidReward(uint256 amount, address provider)
func (_Contract *ContractFilterer) FilterEventIncreaseLiquidReward(opts *bind.FilterOpts) (*ContractEventIncreaseLiquidRewardIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "eventIncreaseLiquidReward")
	if err != nil {
		return nil, err
	}
	return &ContractEventIncreaseLiquidRewardIterator{contract: _Contract.contract, event: "eventIncreaseLiquidReward", logs: logs, sub: sub}, nil
}

// WatchEventIncreaseLiquidReward is a free log subscription operation binding the contract event 0x045492b29efb11990e42d3b9ba88978042183c94051194b886f8595139c64db9.
//
// Solidity: event eventIncreaseLiquidReward(uint256 amount, address provider)
func (_Contract *ContractFilterer) WatchEventIncreaseLiquidReward(opts *bind.WatchOpts, sink chan<- *ContractEventIncreaseLiquidReward) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "eventIncreaseLiquidReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEventIncreaseLiquidReward)
				if err := _Contract.contract.UnpackLog(event, "eventIncreaseLiquidReward", log); err != nil {
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
func (_Contract *ContractFilterer) ParseEventIncreaseLiquidReward(log types.Log) (*ContractEventIncreaseLiquidReward, error) {
	event := new(ContractEventIncreaseLiquidReward)
	if err := _Contract.contract.UnpackLog(event, "eventIncreaseLiquidReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEventNewLoanIterator is returned from FilterEventNewLoan and is used to iterate over the raw logs and unpacked data for EventNewLoan events raised by the Contract contract.
type ContractEventNewLoanIterator struct {
	Event *ContractEventNewLoan // Event containing the contract specifics and raw log

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
func (it *ContractEventNewLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEventNewLoan)
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
		it.Event = new(ContractEventNewLoan)
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
func (it *ContractEventNewLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEventNewLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEventNewLoan represents a EventNewLoan event raised by the Contract contract.
type ContractEventNewLoan struct {
	LoanId   *big.Int
	Duration *big.Int
	Start    *big.Int
	Amount   *big.Int
	Loaner   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEventNewLoan is a free log retrieval operation binding the contract event 0xb2990d266ec4e479259ef8c68e87d6c03ab8dbafa4e785e79d2ed1545a383083.
//
// Solidity: event eventNewLoan(uint256 loanId, uint256 duration, uint256 start, uint256 amount, address loaner)
func (_Contract *ContractFilterer) FilterEventNewLoan(opts *bind.FilterOpts) (*ContractEventNewLoanIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "eventNewLoan")
	if err != nil {
		return nil, err
	}
	return &ContractEventNewLoanIterator{contract: _Contract.contract, event: "eventNewLoan", logs: logs, sub: sub}, nil
}

// WatchEventNewLoan is a free log subscription operation binding the contract event 0xb2990d266ec4e479259ef8c68e87d6c03ab8dbafa4e785e79d2ed1545a383083.
//
// Solidity: event eventNewLoan(uint256 loanId, uint256 duration, uint256 start, uint256 amount, address loaner)
func (_Contract *ContractFilterer) WatchEventNewLoan(opts *bind.WatchOpts, sink chan<- *ContractEventNewLoan) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "eventNewLoan")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEventNewLoan)
				if err := _Contract.contract.UnpackLog(event, "eventNewLoan", log); err != nil {
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

// ParseEventNewLoan is a log parse operation binding the contract event 0xb2990d266ec4e479259ef8c68e87d6c03ab8dbafa4e785e79d2ed1545a383083.
//
// Solidity: event eventNewLoan(uint256 loanId, uint256 duration, uint256 start, uint256 amount, address loaner)
func (_Contract *ContractFilterer) ParseEventNewLoan(log types.Log) (*ContractEventNewLoan, error) {
	event := new(ContractEventNewLoan)
	if err := _Contract.contract.UnpackLog(event, "eventNewLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEventPayBackIterator is returned from FilterEventPayBack and is used to iterate over the raw logs and unpacked data for EventPayBack events raised by the Contract contract.
type ContractEventPayBackIterator struct {
	Event *ContractEventPayBack // Event containing the contract specifics and raw log

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
func (it *ContractEventPayBackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEventPayBack)
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
		it.Event = new(ContractEventPayBack)
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
func (it *ContractEventPayBackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEventPayBackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEventPayBack represents a EventPayBack event raised by the Contract contract.
type ContractEventPayBack struct {
	LoanId *big.Int
	Amount *big.Int
	Loaner common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEventPayBack is a free log retrieval operation binding the contract event 0x38dcb8e7ce8c7f182d53142ee0fa94a1778cd54eddcc1209f86469e7d3b48733.
//
// Solidity: event eventPayBack(uint256 loanId, uint256 amount, address loaner)
func (_Contract *ContractFilterer) FilterEventPayBack(opts *bind.FilterOpts) (*ContractEventPayBackIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "eventPayBack")
	if err != nil {
		return nil, err
	}
	return &ContractEventPayBackIterator{contract: _Contract.contract, event: "eventPayBack", logs: logs, sub: sub}, nil
}

// WatchEventPayBack is a free log subscription operation binding the contract event 0x38dcb8e7ce8c7f182d53142ee0fa94a1778cd54eddcc1209f86469e7d3b48733.
//
// Solidity: event eventPayBack(uint256 loanId, uint256 amount, address loaner)
func (_Contract *ContractFilterer) WatchEventPayBack(opts *bind.WatchOpts, sink chan<- *ContractEventPayBack) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "eventPayBack")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEventPayBack)
				if err := _Contract.contract.UnpackLog(event, "eventPayBack", log); err != nil {
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
func (_Contract *ContractFilterer) ParseEventPayBack(log types.Log) (*ContractEventPayBack, error) {
	event := new(ContractEventPayBack)
	if err := _Contract.contract.UnpackLog(event, "eventPayBack", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEventProviderIncreaseIterator is returned from FilterEventProviderIncrease and is used to iterate over the raw logs and unpacked data for EventProviderIncrease events raised by the Contract contract.
type ContractEventProviderIncreaseIterator struct {
	Event *ContractEventProviderIncrease // Event containing the contract specifics and raw log

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
func (it *ContractEventProviderIncreaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEventProviderIncrease)
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
		it.Event = new(ContractEventProviderIncrease)
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
func (it *ContractEventProviderIncreaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEventProviderIncreaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEventProviderIncrease represents a EventProviderIncrease event raised by the Contract contract.
type ContractEventProviderIncrease struct {
	Amount   *big.Int
	Provider common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEventProviderIncrease is a free log retrieval operation binding the contract event 0xc7f079bb1739a7fcb563479a77ec3ff5de5ba875b2b8b44d897abfc3ac58a8ed.
//
// Solidity: event eventProviderIncrease(uint256 amount, address provider)
func (_Contract *ContractFilterer) FilterEventProviderIncrease(opts *bind.FilterOpts) (*ContractEventProviderIncreaseIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "eventProviderIncrease")
	if err != nil {
		return nil, err
	}
	return &ContractEventProviderIncreaseIterator{contract: _Contract.contract, event: "eventProviderIncrease", logs: logs, sub: sub}, nil
}

// WatchEventProviderIncrease is a free log subscription operation binding the contract event 0xc7f079bb1739a7fcb563479a77ec3ff5de5ba875b2b8b44d897abfc3ac58a8ed.
//
// Solidity: event eventProviderIncrease(uint256 amount, address provider)
func (_Contract *ContractFilterer) WatchEventProviderIncrease(opts *bind.WatchOpts, sink chan<- *ContractEventProviderIncrease) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "eventProviderIncrease")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEventProviderIncrease)
				if err := _Contract.contract.UnpackLog(event, "eventProviderIncrease", log); err != nil {
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

// ParseEventProviderIncrease is a log parse operation binding the contract event 0xc7f079bb1739a7fcb563479a77ec3ff5de5ba875b2b8b44d897abfc3ac58a8ed.
//
// Solidity: event eventProviderIncrease(uint256 amount, address provider)
func (_Contract *ContractFilterer) ParseEventProviderIncrease(log types.Log) (*ContractEventProviderIncrease, error) {
	event := new(ContractEventProviderIncrease)
	if err := _Contract.contract.UnpackLog(event, "eventProviderIncrease", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEventProviderRetrieveIterator is returned from FilterEventProviderRetrieve and is used to iterate over the raw logs and unpacked data for EventProviderRetrieve events raised by the Contract contract.
type ContractEventProviderRetrieveIterator struct {
	Event *ContractEventProviderRetrieve // Event containing the contract specifics and raw log

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
func (it *ContractEventProviderRetrieveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEventProviderRetrieve)
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
		it.Event = new(ContractEventProviderRetrieve)
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
func (it *ContractEventProviderRetrieveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEventProviderRetrieveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEventProviderRetrieve represents a EventProviderRetrieve event raised by the Contract contract.
type ContractEventProviderRetrieve struct {
	Amount   *big.Int
	Provider common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEventProviderRetrieve is a free log retrieval operation binding the contract event 0x89fe3f29313aa6c03800bc780dc2b251ec749edd116569f40bd397cf1e8e08c8.
//
// Solidity: event eventProviderRetrieve(uint256 amount, address provider)
func (_Contract *ContractFilterer) FilterEventProviderRetrieve(opts *bind.FilterOpts) (*ContractEventProviderRetrieveIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "eventProviderRetrieve")
	if err != nil {
		return nil, err
	}
	return &ContractEventProviderRetrieveIterator{contract: _Contract.contract, event: "eventProviderRetrieve", logs: logs, sub: sub}, nil
}

// WatchEventProviderRetrieve is a free log subscription operation binding the contract event 0x89fe3f29313aa6c03800bc780dc2b251ec749edd116569f40bd397cf1e8e08c8.
//
// Solidity: event eventProviderRetrieve(uint256 amount, address provider)
func (_Contract *ContractFilterer) WatchEventProviderRetrieve(opts *bind.WatchOpts, sink chan<- *ContractEventProviderRetrieve) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "eventProviderRetrieve")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEventProviderRetrieve)
				if err := _Contract.contract.UnpackLog(event, "eventProviderRetrieve", log); err != nil {
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

// ParseEventProviderRetrieve is a log parse operation binding the contract event 0x89fe3f29313aa6c03800bc780dc2b251ec749edd116569f40bd397cf1e8e08c8.
//
// Solidity: event eventProviderRetrieve(uint256 amount, address provider)
func (_Contract *ContractFilterer) ParseEventProviderRetrieve(log types.Log) (*ContractEventProviderRetrieve, error) {
	event := new(ContractEventProviderRetrieve)
	if err := _Contract.contract.UnpackLog(event, "eventProviderRetrieve", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEventReleaseLiquidRewardIterator is returned from FilterEventReleaseLiquidReward and is used to iterate over the raw logs and unpacked data for EventReleaseLiquidReward events raised by the Contract contract.
type ContractEventReleaseLiquidRewardIterator struct {
	Event *ContractEventReleaseLiquidReward // Event containing the contract specifics and raw log

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
func (it *ContractEventReleaseLiquidRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEventReleaseLiquidReward)
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
		it.Event = new(ContractEventReleaseLiquidReward)
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
func (it *ContractEventReleaseLiquidRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEventReleaseLiquidRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEventReleaseLiquidReward represents a EventReleaseLiquidReward event raised by the Contract contract.
type ContractEventReleaseLiquidReward struct {
	Amount   *big.Int
	Provider common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEventReleaseLiquidReward is a free log retrieval operation binding the contract event 0x467456fb7eb39617bb976fc80dbff34a252f7ef0f1887ebcca77c2caaac5704d.
//
// Solidity: event eventReleaseLiquidReward(uint256 amount, address provider)
func (_Contract *ContractFilterer) FilterEventReleaseLiquidReward(opts *bind.FilterOpts) (*ContractEventReleaseLiquidRewardIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "eventReleaseLiquidReward")
	if err != nil {
		return nil, err
	}
	return &ContractEventReleaseLiquidRewardIterator{contract: _Contract.contract, event: "eventReleaseLiquidReward", logs: logs, sub: sub}, nil
}

// WatchEventReleaseLiquidReward is a free log subscription operation binding the contract event 0x467456fb7eb39617bb976fc80dbff34a252f7ef0f1887ebcca77c2caaac5704d.
//
// Solidity: event eventReleaseLiquidReward(uint256 amount, address provider)
func (_Contract *ContractFilterer) WatchEventReleaseLiquidReward(opts *bind.WatchOpts, sink chan<- *ContractEventReleaseLiquidReward) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "eventReleaseLiquidReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEventReleaseLiquidReward)
				if err := _Contract.contract.UnpackLog(event, "eventReleaseLiquidReward", log); err != nil {
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

// ParseEventReleaseLiquidReward is a log parse operation binding the contract event 0x467456fb7eb39617bb976fc80dbff34a252f7ef0f1887ebcca77c2caaac5704d.
//
// Solidity: event eventReleaseLiquidReward(uint256 amount, address provider)
func (_Contract *ContractFilterer) ParseEventReleaseLiquidReward(log types.Log) (*ContractEventReleaseLiquidReward, error) {
	event := new(ContractEventReleaseLiquidReward)
	if err := _Contract.contract.UnpackLog(event, "eventReleaseLiquidReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
