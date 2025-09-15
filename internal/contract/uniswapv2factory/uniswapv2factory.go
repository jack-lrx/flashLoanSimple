// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapv2factory

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

// Uniswapv2factoryMetaData contains all meta data concerning the Uniswapv2factory contract.
var Uniswapv2factoryMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Uniswapv2factoryABI is the input ABI used to generate the binding from.
// Deprecated: Use Uniswapv2factoryMetaData.ABI instead.
var Uniswapv2factoryABI = Uniswapv2factoryMetaData.ABI

// Uniswapv2factory is an auto generated Go binding around an Ethereum contract.
type Uniswapv2factory struct {
	Uniswapv2factoryCaller     // Read-only binding to the contract
	Uniswapv2factoryTransactor // Write-only binding to the contract
	Uniswapv2factoryFilterer   // Log filterer for contract events
}

// Uniswapv2factoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type Uniswapv2factoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv2factoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Uniswapv2factoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv2factoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Uniswapv2factoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv2factorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Uniswapv2factorySession struct {
	Contract     *Uniswapv2factory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Uniswapv2factoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Uniswapv2factoryCallerSession struct {
	Contract *Uniswapv2factoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// Uniswapv2factoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Uniswapv2factoryTransactorSession struct {
	Contract     *Uniswapv2factoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// Uniswapv2factoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type Uniswapv2factoryRaw struct {
	Contract *Uniswapv2factory // Generic contract binding to access the raw methods on
}

// Uniswapv2factoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Uniswapv2factoryCallerRaw struct {
	Contract *Uniswapv2factoryCaller // Generic read-only contract binding to access the raw methods on
}

// Uniswapv2factoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Uniswapv2factoryTransactorRaw struct {
	Contract *Uniswapv2factoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapv2factory creates a new instance of Uniswapv2factory, bound to a specific deployed contract.
func NewUniswapv2factory(address common.Address, backend bind.ContractBackend) (*Uniswapv2factory, error) {
	contract, err := bindUniswapv2factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2factory{Uniswapv2factoryCaller: Uniswapv2factoryCaller{contract: contract}, Uniswapv2factoryTransactor: Uniswapv2factoryTransactor{contract: contract}, Uniswapv2factoryFilterer: Uniswapv2factoryFilterer{contract: contract}}, nil
}

// NewUniswapv2factoryCaller creates a new read-only instance of Uniswapv2factory, bound to a specific deployed contract.
func NewUniswapv2factoryCaller(address common.Address, caller bind.ContractCaller) (*Uniswapv2factoryCaller, error) {
	contract, err := bindUniswapv2factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2factoryCaller{contract: contract}, nil
}

// NewUniswapv2factoryTransactor creates a new write-only instance of Uniswapv2factory, bound to a specific deployed contract.
func NewUniswapv2factoryTransactor(address common.Address, transactor bind.ContractTransactor) (*Uniswapv2factoryTransactor, error) {
	contract, err := bindUniswapv2factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2factoryTransactor{contract: contract}, nil
}

// NewUniswapv2factoryFilterer creates a new log filterer instance of Uniswapv2factory, bound to a specific deployed contract.
func NewUniswapv2factoryFilterer(address common.Address, filterer bind.ContractFilterer) (*Uniswapv2factoryFilterer, error) {
	contract, err := bindUniswapv2factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2factoryFilterer{contract: contract}, nil
}

// bindUniswapv2factory binds a generic wrapper to an already deployed contract.
func bindUniswapv2factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Uniswapv2factoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv2factory *Uniswapv2factoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv2factory.Contract.Uniswapv2factoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv2factory *Uniswapv2factoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv2factory.Contract.Uniswapv2factoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv2factory *Uniswapv2factoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv2factory.Contract.Uniswapv2factoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv2factory *Uniswapv2factoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv2factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv2factory *Uniswapv2factoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv2factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv2factory *Uniswapv2factoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv2factory.Contract.contract.Transact(opts, method, params...)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_Uniswapv2factory *Uniswapv2factoryCaller) GetPair(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv2factory.contract.Call(opts, &out, "getPair", tokenA, tokenB)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_Uniswapv2factory *Uniswapv2factorySession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _Uniswapv2factory.Contract.GetPair(&_Uniswapv2factory.CallOpts, tokenA, tokenB)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_Uniswapv2factory *Uniswapv2factoryCallerSession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _Uniswapv2factory.Contract.GetPair(&_Uniswapv2factory.CallOpts, tokenA, tokenB)
}
