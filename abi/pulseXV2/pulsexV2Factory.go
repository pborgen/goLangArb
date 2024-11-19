// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pulsexV2Factory

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

// PulsexV2FactoryMetaData contains all meta data concerning the PulsexV2Factory contract.
var PulsexV2FactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"INIT_CODE_PAIR_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"}]",
}

// PulsexV2FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use PulsexV2FactoryMetaData.ABI instead.
var PulsexV2FactoryABI = PulsexV2FactoryMetaData.ABI

// PulsexV2Factory is an auto generated Go binding around an Ethereum contract.
type PulsexV2Factory struct {
	PulsexV2FactoryCaller     // Read-only binding to the contract
	PulsexV2FactoryTransactor // Write-only binding to the contract
	PulsexV2FactoryFilterer   // Log filterer for contract events
}

// PulsexV2FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PulsexV2FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PulsexV2FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PulsexV2FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PulsexV2FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PulsexV2FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PulsexV2FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PulsexV2FactorySession struct {
	Contract     *PulsexV2Factory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PulsexV2FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PulsexV2FactoryCallerSession struct {
	Contract *PulsexV2FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PulsexV2FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PulsexV2FactoryTransactorSession struct {
	Contract     *PulsexV2FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PulsexV2FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PulsexV2FactoryRaw struct {
	Contract *PulsexV2Factory // Generic contract binding to access the raw methods on
}

// PulsexV2FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PulsexV2FactoryCallerRaw struct {
	Contract *PulsexV2FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// PulsexV2FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PulsexV2FactoryTransactorRaw struct {
	Contract *PulsexV2FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPulsexV2Factory creates a new instance of PulsexV2Factory, bound to a specific deployed contract.
func NewPulsexV2Factory(address common.Address, backend bind.ContractBackend) (*PulsexV2Factory, error) {
	contract, err := bindPulsexV2Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PulsexV2Factory{PulsexV2FactoryCaller: PulsexV2FactoryCaller{contract: contract}, PulsexV2FactoryTransactor: PulsexV2FactoryTransactor{contract: contract}, PulsexV2FactoryFilterer: PulsexV2FactoryFilterer{contract: contract}}, nil
}

// NewPulsexV2FactoryCaller creates a new read-only instance of PulsexV2Factory, bound to a specific deployed contract.
func NewPulsexV2FactoryCaller(address common.Address, caller bind.ContractCaller) (*PulsexV2FactoryCaller, error) {
	contract, err := bindPulsexV2Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PulsexV2FactoryCaller{contract: contract}, nil
}

// NewPulsexV2FactoryTransactor creates a new write-only instance of PulsexV2Factory, bound to a specific deployed contract.
func NewPulsexV2FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*PulsexV2FactoryTransactor, error) {
	contract, err := bindPulsexV2Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PulsexV2FactoryTransactor{contract: contract}, nil
}

// NewPulsexV2FactoryFilterer creates a new log filterer instance of PulsexV2Factory, bound to a specific deployed contract.
func NewPulsexV2FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*PulsexV2FactoryFilterer, error) {
	contract, err := bindPulsexV2Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PulsexV2FactoryFilterer{contract: contract}, nil
}

// bindPulsexV2Factory binds a generic wrapper to an already deployed contract.
func bindPulsexV2Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PulsexV2FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PulsexV2Factory *PulsexV2FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PulsexV2Factory.Contract.PulsexV2FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PulsexV2Factory *PulsexV2FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PulsexV2Factory.Contract.PulsexV2FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PulsexV2Factory *PulsexV2FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PulsexV2Factory.Contract.PulsexV2FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PulsexV2Factory *PulsexV2FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PulsexV2Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PulsexV2Factory *PulsexV2FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PulsexV2Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PulsexV2Factory *PulsexV2FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PulsexV2Factory.Contract.contract.Transact(opts, method, params...)
}

// INITCODEPAIRHASH is a free data retrieval call binding the contract method 0x5855a25a.
//
// Solidity: function INIT_CODE_PAIR_HASH() view returns(bytes32)
func (_PulsexV2Factory *PulsexV2FactoryCaller) INITCODEPAIRHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PulsexV2Factory.contract.Call(opts, &out, "INIT_CODE_PAIR_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// INITCODEPAIRHASH is a free data retrieval call binding the contract method 0x5855a25a.
//
// Solidity: function INIT_CODE_PAIR_HASH() view returns(bytes32)
func (_PulsexV2Factory *PulsexV2FactorySession) INITCODEPAIRHASH() ([32]byte, error) {
	return _PulsexV2Factory.Contract.INITCODEPAIRHASH(&_PulsexV2Factory.CallOpts)
}

// INITCODEPAIRHASH is a free data retrieval call binding the contract method 0x5855a25a.
//
// Solidity: function INIT_CODE_PAIR_HASH() view returns(bytes32)
func (_PulsexV2Factory *PulsexV2FactoryCallerSession) INITCODEPAIRHASH() ([32]byte, error) {
	return _PulsexV2Factory.Contract.INITCODEPAIRHASH(&_PulsexV2Factory.CallOpts)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_PulsexV2Factory *PulsexV2FactoryCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PulsexV2Factory.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_PulsexV2Factory *PulsexV2FactorySession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _PulsexV2Factory.Contract.AllPairs(&_PulsexV2Factory.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_PulsexV2Factory *PulsexV2FactoryCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _PulsexV2Factory.Contract.AllPairs(&_PulsexV2Factory.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_PulsexV2Factory *PulsexV2FactoryCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PulsexV2Factory.contract.Call(opts, &out, "allPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_PulsexV2Factory *PulsexV2FactorySession) AllPairsLength() (*big.Int, error) {
	return _PulsexV2Factory.Contract.AllPairsLength(&_PulsexV2Factory.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_PulsexV2Factory *PulsexV2FactoryCallerSession) AllPairsLength() (*big.Int, error) {
	return _PulsexV2Factory.Contract.AllPairsLength(&_PulsexV2Factory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_PulsexV2Factory *PulsexV2FactoryCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PulsexV2Factory.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_PulsexV2Factory *PulsexV2FactorySession) FeeTo() (common.Address, error) {
	return _PulsexV2Factory.Contract.FeeTo(&_PulsexV2Factory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_PulsexV2Factory *PulsexV2FactoryCallerSession) FeeTo() (common.Address, error) {
	return _PulsexV2Factory.Contract.FeeTo(&_PulsexV2Factory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_PulsexV2Factory *PulsexV2FactoryCaller) FeeToSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PulsexV2Factory.contract.Call(opts, &out, "feeToSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_PulsexV2Factory *PulsexV2FactorySession) FeeToSetter() (common.Address, error) {
	return _PulsexV2Factory.Contract.FeeToSetter(&_PulsexV2Factory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_PulsexV2Factory *PulsexV2FactoryCallerSession) FeeToSetter() (common.Address, error) {
	return _PulsexV2Factory.Contract.FeeToSetter(&_PulsexV2Factory.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_PulsexV2Factory *PulsexV2FactoryCaller) GetPair(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _PulsexV2Factory.contract.Call(opts, &out, "getPair", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_PulsexV2Factory *PulsexV2FactorySession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _PulsexV2Factory.Contract.GetPair(&_PulsexV2Factory.CallOpts, arg0, arg1)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_PulsexV2Factory *PulsexV2FactoryCallerSession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _PulsexV2Factory.Contract.GetPair(&_PulsexV2Factory.CallOpts, arg0, arg1)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_PulsexV2Factory *PulsexV2FactoryTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _PulsexV2Factory.contract.Transact(opts, "createPair", tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_PulsexV2Factory *PulsexV2FactorySession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _PulsexV2Factory.Contract.CreatePair(&_PulsexV2Factory.TransactOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_PulsexV2Factory *PulsexV2FactoryTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _PulsexV2Factory.Contract.CreatePair(&_PulsexV2Factory.TransactOpts, tokenA, tokenB)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_PulsexV2Factory *PulsexV2FactoryTransactor) SetFeeTo(opts *bind.TransactOpts, _feeTo common.Address) (*types.Transaction, error) {
	return _PulsexV2Factory.contract.Transact(opts, "setFeeTo", _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_PulsexV2Factory *PulsexV2FactorySession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _PulsexV2Factory.Contract.SetFeeTo(&_PulsexV2Factory.TransactOpts, _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_PulsexV2Factory *PulsexV2FactoryTransactorSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _PulsexV2Factory.Contract.SetFeeTo(&_PulsexV2Factory.TransactOpts, _feeTo)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_PulsexV2Factory *PulsexV2FactoryTransactor) SetFeeToSetter(opts *bind.TransactOpts, _feeToSetter common.Address) (*types.Transaction, error) {
	return _PulsexV2Factory.contract.Transact(opts, "setFeeToSetter", _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_PulsexV2Factory *PulsexV2FactorySession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _PulsexV2Factory.Contract.SetFeeToSetter(&_PulsexV2Factory.TransactOpts, _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_PulsexV2Factory *PulsexV2FactoryTransactorSession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _PulsexV2Factory.Contract.SetFeeToSetter(&_PulsexV2Factory.TransactOpts, _feeToSetter)
}

// PulsexV2FactoryPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the PulsexV2Factory contract.
type PulsexV2FactoryPairCreatedIterator struct {
	Event *PulsexV2FactoryPairCreated // Event containing the contract specifics and raw log

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
func (it *PulsexV2FactoryPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PulsexV2FactoryPairCreated)
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
		it.Event = new(PulsexV2FactoryPairCreated)
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
func (it *PulsexV2FactoryPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PulsexV2FactoryPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PulsexV2FactoryPairCreated represents a PairCreated event raised by the PulsexV2Factory contract.
type PulsexV2FactoryPairCreated struct {
	Token0 common.Address
	Token1 common.Address
	Pair   common.Address
	Arg3   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_PulsexV2Factory *PulsexV2FactoryFilterer) FilterPairCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*PulsexV2FactoryPairCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _PulsexV2Factory.contract.FilterLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &PulsexV2FactoryPairCreatedIterator{contract: _PulsexV2Factory.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_PulsexV2Factory *PulsexV2FactoryFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *PulsexV2FactoryPairCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _PulsexV2Factory.contract.WatchLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PulsexV2FactoryPairCreated)
				if err := _PulsexV2Factory.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_PulsexV2Factory *PulsexV2FactoryFilterer) ParsePairCreated(log types.Log) (*PulsexV2FactoryPairCreated, error) {
	event := new(PulsexV2FactoryPairCreated)
	if err := _PulsexV2Factory.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
