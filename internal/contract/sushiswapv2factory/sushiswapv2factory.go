// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sushiswapv2factory

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

// Sushiswapv2factoryMetaData contains all meta data concerning the Sushiswapv2factory contract.
var Sushiswapv2factoryMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Sushiswapv2factoryABI is the input ABI used to generate the binding from.
// Deprecated: Use Sushiswapv2factoryMetaData.ABI instead.
var Sushiswapv2factoryABI = Sushiswapv2factoryMetaData.ABI

// Sushiswapv2factory is an auto generated Go binding around an Ethereum contract.
type Sushiswapv2factory struct {
	Sushiswapv2factoryCaller     // Read-only binding to the contract
	Sushiswapv2factoryTransactor // Write-only binding to the contract
	Sushiswapv2factoryFilterer   // Log filterer for contract events
}

// Sushiswapv2factoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type Sushiswapv2factoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Sushiswapv2factoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Sushiswapv2factoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Sushiswapv2factoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Sushiswapv2factoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Sushiswapv2factorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Sushiswapv2factorySession struct {
	Contract     *Sushiswapv2factory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Sushiswapv2factoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Sushiswapv2factoryCallerSession struct {
	Contract *Sushiswapv2factoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// Sushiswapv2factoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Sushiswapv2factoryTransactorSession struct {
	Contract     *Sushiswapv2factoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// Sushiswapv2factoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type Sushiswapv2factoryRaw struct {
	Contract *Sushiswapv2factory // Generic contract binding to access the raw methods on
}

// Sushiswapv2factoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Sushiswapv2factoryCallerRaw struct {
	Contract *Sushiswapv2factoryCaller // Generic read-only contract binding to access the raw methods on
}

// Sushiswapv2factoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Sushiswapv2factoryTransactorRaw struct {
	Contract *Sushiswapv2factoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSushiswapv2factory creates a new instance of Sushiswapv2factory, bound to a specific deployed contract.
func NewSushiswapv2factory(address common.Address, backend bind.ContractBackend) (*Sushiswapv2factory, error) {
	contract, err := bindSushiswapv2factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sushiswapv2factory{Sushiswapv2factoryCaller: Sushiswapv2factoryCaller{contract: contract}, Sushiswapv2factoryTransactor: Sushiswapv2factoryTransactor{contract: contract}, Sushiswapv2factoryFilterer: Sushiswapv2factoryFilterer{contract: contract}}, nil
}

// NewSushiswapv2factoryCaller creates a new read-only instance of Sushiswapv2factory, bound to a specific deployed contract.
func NewSushiswapv2factoryCaller(address common.Address, caller bind.ContractCaller) (*Sushiswapv2factoryCaller, error) {
	contract, err := bindSushiswapv2factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Sushiswapv2factoryCaller{contract: contract}, nil
}

// NewSushiswapv2factoryTransactor creates a new write-only instance of Sushiswapv2factory, bound to a specific deployed contract.
func NewSushiswapv2factoryTransactor(address common.Address, transactor bind.ContractTransactor) (*Sushiswapv2factoryTransactor, error) {
	contract, err := bindSushiswapv2factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Sushiswapv2factoryTransactor{contract: contract}, nil
}

// NewSushiswapv2factoryFilterer creates a new log filterer instance of Sushiswapv2factory, bound to a specific deployed contract.
func NewSushiswapv2factoryFilterer(address common.Address, filterer bind.ContractFilterer) (*Sushiswapv2factoryFilterer, error) {
	contract, err := bindSushiswapv2factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Sushiswapv2factoryFilterer{contract: contract}, nil
}

// bindSushiswapv2factory binds a generic wrapper to an already deployed contract.
func bindSushiswapv2factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Sushiswapv2factoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sushiswapv2factory *Sushiswapv2factoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sushiswapv2factory.Contract.Sushiswapv2factoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sushiswapv2factory *Sushiswapv2factoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sushiswapv2factory.Contract.Sushiswapv2factoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sushiswapv2factory *Sushiswapv2factoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sushiswapv2factory.Contract.Sushiswapv2factoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sushiswapv2factory *Sushiswapv2factoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sushiswapv2factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sushiswapv2factory *Sushiswapv2factoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sushiswapv2factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sushiswapv2factory *Sushiswapv2factoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sushiswapv2factory.Contract.contract.Transact(opts, method, params...)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_Sushiswapv2factory *Sushiswapv2factoryCaller) GetPair(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	var out []interface{}
	err := _Sushiswapv2factory.contract.Call(opts, &out, "getPair", tokenA, tokenB)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_Sushiswapv2factory *Sushiswapv2factorySession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _Sushiswapv2factory.Contract.GetPair(&_Sushiswapv2factory.CallOpts, tokenA, tokenB)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_Sushiswapv2factory *Sushiswapv2factoryCallerSession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _Sushiswapv2factory.Contract.GetPair(&_Sushiswapv2factory.CallOpts, tokenA, tokenB)
}
