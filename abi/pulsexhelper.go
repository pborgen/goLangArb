// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pulsexhelper

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

// PulsexhelperMetaData contains all meta data concerning the Pulsexhelper contract.
var PulsexhelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"addressValue\",\"type\":\"uint256\"}],\"name\":\"convertUintToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"pairFor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// PulsexhelperABI is the input ABI used to generate the binding from.
// Deprecated: Use PulsexhelperMetaData.ABI instead.
var PulsexhelperABI = PulsexhelperMetaData.ABI

// Pulsexhelper is an auto generated Go binding around an Ethereum contract.
type Pulsexhelper struct {
	PulsexhelperCaller     // Read-only binding to the contract
	PulsexhelperTransactor // Write-only binding to the contract
	PulsexhelperFilterer   // Log filterer for contract events
}

// PulsexhelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type PulsexhelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PulsexhelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PulsexhelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PulsexhelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PulsexhelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PulsexhelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PulsexhelperSession struct {
	Contract     *Pulsexhelper     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PulsexhelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PulsexhelperCallerSession struct {
	Contract *PulsexhelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PulsexhelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PulsexhelperTransactorSession struct {
	Contract     *PulsexhelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PulsexhelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type PulsexhelperRaw struct {
	Contract *Pulsexhelper // Generic contract binding to access the raw methods on
}

// PulsexhelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PulsexhelperCallerRaw struct {
	Contract *PulsexhelperCaller // Generic read-only contract binding to access the raw methods on
}

// PulsexhelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PulsexhelperTransactorRaw struct {
	Contract *PulsexhelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPulsexhelper creates a new instance of Pulsexhelper, bound to a specific deployed contract.
func NewPulsexhelper(address common.Address, backend bind.ContractBackend) (*Pulsexhelper, error) {
	contract, err := bindPulsexhelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pulsexhelper{PulsexhelperCaller: PulsexhelperCaller{contract: contract}, PulsexhelperTransactor: PulsexhelperTransactor{contract: contract}, PulsexhelperFilterer: PulsexhelperFilterer{contract: contract}}, nil
}

// NewPulsexhelperCaller creates a new read-only instance of Pulsexhelper, bound to a specific deployed contract.
func NewPulsexhelperCaller(address common.Address, caller bind.ContractCaller) (*PulsexhelperCaller, error) {
	contract, err := bindPulsexhelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PulsexhelperCaller{contract: contract}, nil
}

// NewPulsexhelperTransactor creates a new write-only instance of Pulsexhelper, bound to a specific deployed contract.
func NewPulsexhelperTransactor(address common.Address, transactor bind.ContractTransactor) (*PulsexhelperTransactor, error) {
	contract, err := bindPulsexhelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PulsexhelperTransactor{contract: contract}, nil
}

// NewPulsexhelperFilterer creates a new log filterer instance of Pulsexhelper, bound to a specific deployed contract.
func NewPulsexhelperFilterer(address common.Address, filterer bind.ContractFilterer) (*PulsexhelperFilterer, error) {
	contract, err := bindPulsexhelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PulsexhelperFilterer{contract: contract}, nil
}

// bindPulsexhelper binds a generic wrapper to an already deployed contract.
func bindPulsexhelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PulsexhelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pulsexhelper *PulsexhelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pulsexhelper.Contract.PulsexhelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pulsexhelper *PulsexhelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pulsexhelper.Contract.PulsexhelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pulsexhelper *PulsexhelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pulsexhelper.Contract.PulsexhelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pulsexhelper *PulsexhelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pulsexhelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pulsexhelper *PulsexhelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pulsexhelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pulsexhelper *PulsexhelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pulsexhelper.Contract.contract.Transact(opts, method, params...)
}

// ConvertUintToAddress is a free data retrieval call binding the contract method 0xf5c2db2b.
//
// Solidity: function convertUintToAddress(uint256 addressValue) pure returns(address)
func (_Pulsexhelper *PulsexhelperCaller) ConvertUintToAddress(opts *bind.CallOpts, addressValue *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Pulsexhelper.contract.Call(opts, &out, "convertUintToAddress", addressValue)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConvertUintToAddress is a free data retrieval call binding the contract method 0xf5c2db2b.
//
// Solidity: function convertUintToAddress(uint256 addressValue) pure returns(address)
func (_Pulsexhelper *PulsexhelperSession) ConvertUintToAddress(addressValue *big.Int) (common.Address, error) {
	return _Pulsexhelper.Contract.ConvertUintToAddress(&_Pulsexhelper.CallOpts, addressValue)
}

// ConvertUintToAddress is a free data retrieval call binding the contract method 0xf5c2db2b.
//
// Solidity: function convertUintToAddress(uint256 addressValue) pure returns(address)
func (_Pulsexhelper *PulsexhelperCallerSession) ConvertUintToAddress(addressValue *big.Int) (common.Address, error) {
	return _Pulsexhelper.Contract.ConvertUintToAddress(&_Pulsexhelper.CallOpts, addressValue)
}

// PairFor is a free data retrieval call binding the contract method 0x6d91c0e2.
//
// Solidity: function pairFor(address factory, address tokenA, address tokenB) pure returns(address pair)
func (_Pulsexhelper *PulsexhelperCaller) PairFor(opts *bind.CallOpts, factory common.Address, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	var out []interface{}
	err := _Pulsexhelper.contract.Call(opts, &out, "pairFor", factory, tokenA, tokenB)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairFor is a free data retrieval call binding the contract method 0x6d91c0e2.
//
// Solidity: function pairFor(address factory, address tokenA, address tokenB) pure returns(address pair)
func (_Pulsexhelper *PulsexhelperSession) PairFor(factory common.Address, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _Pulsexhelper.Contract.PairFor(&_Pulsexhelper.CallOpts, factory, tokenA, tokenB)
}

// PairFor is a free data retrieval call binding the contract method 0x6d91c0e2.
//
// Solidity: function pairFor(address factory, address tokenA, address tokenB) pure returns(address pair)
func (_Pulsexhelper *PulsexhelperCallerSession) PairFor(factory common.Address, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _Pulsexhelper.Contract.PairFor(&_Pulsexhelper.CallOpts, factory, tokenA, tokenB)
}
