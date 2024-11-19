// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simplev2arbabi

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

// Simplev2arbabiMetaData contains all meta data concerning the Simplev2arbabi contract.
var Simplev2arbabiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"routerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sell_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buy_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"_getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approvee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"approveToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken0\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"routerAddress0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"routerAddress1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0Address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1Address\",\"type\":\"address\"}],\"name\":\"makeArbitrageSimple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken0\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"routerAddress0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"routerAddress1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0Address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1Address\",\"type\":\"address\"}],\"name\":\"makeArbitrageSimpleNoCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken0\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"routerAddress0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"routerAddress1\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"pathRouter0\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"pathRouter1\",\"type\":\"address[]\"}],\"name\":\"makeArbitrageSimpleWithPath\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFinalAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"routerAddress0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"routerAddress1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0Address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1Address\",\"type\":\"address\"}],\"name\":\"testArb\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFinalAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"routerAddress0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"routerAddress1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0Address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1Address\",\"type\":\"address\"}],\"name\":\"testArbTwo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Simplev2arbabiABI is the input ABI used to generate the binding from.
// Deprecated: Use Simplev2arbabiMetaData.ABI instead.
var Simplev2arbabiABI = Simplev2arbabiMetaData.ABI

// Simplev2arbabi is an auto generated Go binding around an Ethereum contract.
type Simplev2arbabi struct {
	Simplev2arbabiCaller     // Read-only binding to the contract
	Simplev2arbabiTransactor // Write-only binding to the contract
	Simplev2arbabiFilterer   // Log filterer for contract events
}

// Simplev2arbabiCaller is an auto generated read-only Go binding around an Ethereum contract.
type Simplev2arbabiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Simplev2arbabiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Simplev2arbabiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Simplev2arbabiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Simplev2arbabiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Simplev2arbabiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Simplev2arbabiSession struct {
	Contract     *Simplev2arbabi   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Simplev2arbabiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Simplev2arbabiCallerSession struct {
	Contract *Simplev2arbabiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// Simplev2arbabiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Simplev2arbabiTransactorSession struct {
	Contract     *Simplev2arbabiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Simplev2arbabiRaw is an auto generated low-level Go binding around an Ethereum contract.
type Simplev2arbabiRaw struct {
	Contract *Simplev2arbabi // Generic contract binding to access the raw methods on
}

// Simplev2arbabiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Simplev2arbabiCallerRaw struct {
	Contract *Simplev2arbabiCaller // Generic read-only contract binding to access the raw methods on
}

// Simplev2arbabiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Simplev2arbabiTransactorRaw struct {
	Contract *Simplev2arbabiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimplev2arbabi creates a new instance of Simplev2arbabi, bound to a specific deployed contract.
func NewSimplev2arbabi(address common.Address, backend bind.ContractBackend) (*Simplev2arbabi, error) {
	contract, err := bindSimplev2arbabi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Simplev2arbabi{Simplev2arbabiCaller: Simplev2arbabiCaller{contract: contract}, Simplev2arbabiTransactor: Simplev2arbabiTransactor{contract: contract}, Simplev2arbabiFilterer: Simplev2arbabiFilterer{contract: contract}}, nil
}

// NewSimplev2arbabiCaller creates a new read-only instance of Simplev2arbabi, bound to a specific deployed contract.
func NewSimplev2arbabiCaller(address common.Address, caller bind.ContractCaller) (*Simplev2arbabiCaller, error) {
	contract, err := bindSimplev2arbabi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Simplev2arbabiCaller{contract: contract}, nil
}

// NewSimplev2arbabiTransactor creates a new write-only instance of Simplev2arbabi, bound to a specific deployed contract.
func NewSimplev2arbabiTransactor(address common.Address, transactor bind.ContractTransactor) (*Simplev2arbabiTransactor, error) {
	contract, err := bindSimplev2arbabi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Simplev2arbabiTransactor{contract: contract}, nil
}

// NewSimplev2arbabiFilterer creates a new log filterer instance of Simplev2arbabi, bound to a specific deployed contract.
func NewSimplev2arbabiFilterer(address common.Address, filterer bind.ContractFilterer) (*Simplev2arbabiFilterer, error) {
	contract, err := bindSimplev2arbabi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Simplev2arbabiFilterer{contract: contract}, nil
}

// bindSimplev2arbabi binds a generic wrapper to an already deployed contract.
func bindSimplev2arbabi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Simplev2arbabiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simplev2arbabi *Simplev2arbabiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simplev2arbabi.Contract.Simplev2arbabiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simplev2arbabi *Simplev2arbabiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simplev2arbabi.Contract.Simplev2arbabiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simplev2arbabi *Simplev2arbabiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simplev2arbabi.Contract.Simplev2arbabiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simplev2arbabi *Simplev2arbabiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simplev2arbabi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simplev2arbabi *Simplev2arbabiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simplev2arbabi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simplev2arbabi *Simplev2arbabiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simplev2arbabi.Contract.contract.Transact(opts, method, params...)
}

// GetPrice is a free data retrieval call binding the contract method 0x7937a986.
//
// Solidity: function _getPrice(address routerAddress, address sell_token, address buy_token, uint256 amount) view returns(uint256)
func (_Simplev2arbabi *Simplev2arbabiCaller) GetPrice(opts *bind.CallOpts, routerAddress common.Address, sell_token common.Address, buy_token common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Simplev2arbabi.contract.Call(opts, &out, "_getPrice", routerAddress, sell_token, buy_token, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x7937a986.
//
// Solidity: function _getPrice(address routerAddress, address sell_token, address buy_token, uint256 amount) view returns(uint256)
func (_Simplev2arbabi *Simplev2arbabiSession) GetPrice(routerAddress common.Address, sell_token common.Address, buy_token common.Address, amount *big.Int) (*big.Int, error) {
	return _Simplev2arbabi.Contract.GetPrice(&_Simplev2arbabi.CallOpts, routerAddress, sell_token, buy_token, amount)
}

// GetPrice is a free data retrieval call binding the contract method 0x7937a986.
//
// Solidity: function _getPrice(address routerAddress, address sell_token, address buy_token, uint256 amount) view returns(uint256)
func (_Simplev2arbabi *Simplev2arbabiCallerSession) GetPrice(routerAddress common.Address, sell_token common.Address, buy_token common.Address, amount *big.Int) (*big.Int, error) {
	return _Simplev2arbabi.Contract.GetPrice(&_Simplev2arbabi.CallOpts, routerAddress, sell_token, buy_token, amount)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Simplev2arbabi *Simplev2arbabiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Simplev2arbabi.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Simplev2arbabi *Simplev2arbabiSession) Owner() (common.Address, error) {
	return _Simplev2arbabi.Contract.Owner(&_Simplev2arbabi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Simplev2arbabi *Simplev2arbabiCallerSession) Owner() (common.Address, error) {
	return _Simplev2arbabi.Contract.Owner(&_Simplev2arbabi.CallOpts)
}

// ApproveToken is a paid mutator transaction binding the contract method 0x03105b04.
//
// Solidity: function approveToken(address approvee, address tokenAddress) returns()
func (_Simplev2arbabi *Simplev2arbabiTransactor) ApproveToken(opts *bind.TransactOpts, approvee common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Simplev2arbabi.contract.Transact(opts, "approveToken", approvee, tokenAddress)
}

// ApproveToken is a paid mutator transaction binding the contract method 0x03105b04.
//
// Solidity: function approveToken(address approvee, address tokenAddress) returns()
func (_Simplev2arbabi *Simplev2arbabiSession) ApproveToken(approvee common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Simplev2arbabi.Contract.ApproveToken(&_Simplev2arbabi.TransactOpts, approvee, tokenAddress)
}

// ApproveToken is a paid mutator transaction binding the contract method 0x03105b04.
//
// Solidity: function approveToken(address approvee, address tokenAddress) returns()
func (_Simplev2arbabi *Simplev2arbabiTransactorSession) ApproveToken(approvee common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Simplev2arbabi.Contract.ApproveToken(&_Simplev2arbabi.TransactOpts, approvee, tokenAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address token) returns()
func (_Simplev2arbabi *Simplev2arbabiTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _Simplev2arbabi.contract.Transact(opts, "deposit", amount, token)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address token) returns()
func (_Simplev2arbabi *Simplev2arbabiSession) Deposit(amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _Simplev2arbabi.Contract.Deposit(&_Simplev2arbabi.TransactOpts, amount, token)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address token) returns()
func (_Simplev2arbabi *Simplev2arbabiTransactorSession) Deposit(amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _Simplev2arbabi.Contract.Deposit(&_Simplev2arbabi.TransactOpts, amount, token)
}

// MakeArbitrageSimple is a paid mutator transaction binding the contract method 0xac5580a4.
//
// Solidity: function makeArbitrageSimple(uint256 amountToken0, address routerAddress0, address routerAddress1, address token0Address, address token1Address) returns()
func (_Simplev2arbabi *Simplev2arbabiTransactor) MakeArbitrageSimple(opts *bind.TransactOpts, amountToken0 *big.Int, routerAddress0 common.Address, routerAddress1 common.Address, token0Address common.Address, token1Address common.Address) (*types.Transaction, error) {
	return _Simplev2arbabi.contract.Transact(opts, "makeArbitrageSimple", amountToken0, routerAddress0, routerAddress1, token0Address, token1Address)
}

// MakeArbitrageSimple is a paid mutator transaction binding the contract method 0xac5580a4.
//
// Solidity: function makeArbitrageSimple(uint256 amountToken0, address routerAddress0, address routerAddress1, address token0Address, address token1Address) returns()
func (_Simplev2arbabi *Simplev2arbabiSession) MakeArbitrageSimple(amountToken0 *big.Int, routerAddress0 common.Address, routerAddress1 common.Address, token0Address common.Address, token1Address common.Address) (*types.Transaction, error) {
	return _Simplev2arbabi.Contract.MakeArbitrageSimple(&_Simplev2arbabi.TransactOpts, amountToken0, routerAddress0, routerAddress1, token0Address, token1Address)
}

// MakeArbitrageSimple is a paid mutator transaction binding the contract method 0xac5580a4.
//
// Solidity: function makeArbitrageSimple(uint256 amountToken0, address routerAddress0, address routerAddress1, address token0Address, address token1Address) returns()
func (_Simplev2arbabi *Simplev2arbabiTransactorSession) MakeArbitrageSimple(amountToken0 *big.Int, routerAddress0 common.Address, routerAddress1 common.Address, token0Address common.Address, token1Address common.Address) (*types.Transaction, error) {
	return _Simplev2arbabi.Contract.MakeArbitrageSimple(&_Simplev2arbabi.TransactOpts, amountToken0, routerAddress0, routerAddress1, token0Address, token1Address)
}

// MakeArbitrageSimpleNoCheck is a paid mutator transaction binding the contract method 0x2f5cac73.
//
// Solidity: function makeArbitrageSimpleNoCheck(uint256 amountToken0, address routerAddress0, address routerAddress1, address token0Address, address token1Address) returns()
func (_Simplev2arbabi *Simplev2arbabiTransactor) MakeArbitrageSimpleNoCheck(opts *bind.TransactOpts, amountToken0 *big.Int, routerAddress0 common.Address, routerAddress1 common.Address, token0Address common.Address, token1Address common.Address) (*types.Transaction, error) {
	return _Simplev2arbabi.contract.Transact(opts, "makeArbitrageSimpleNoCheck", amountToken0, routerAddress0, routerAddress1, token0Address, token1Address)
}

// MakeArbitrageSimpleNoCheck is a paid mutator transaction binding the contract method 0x2f5cac73.
//
// Solidity: function makeArbitrageSimpleNoCheck(uint256 amountToken0, address routerAddress0, address routerAddress1, address token0Address, address token1Address) returns()
func (_Simplev2arbabi *Simplev2arbabiSession) MakeArbitrageSimpleNoCheck(amountToken0 *big.Int, routerAddress0 common.Address, routerAddress1 common.Address, token0Address common.Address, token1Address common.Address) (*types.Transaction, error) {
	return _Simplev2arbabi.Contract.MakeArbitrageSimpleNoCheck(&_Simplev2arbabi.TransactOpts, amountToken0, routerAddress0, routerAddress1, token0Address, token1Address)
}

// MakeArbitrageSimpleNoCheck is a paid mutator transaction binding the contract method 0x2f5cac73.
//
// Solidity: function makeArbitrageSimpleNoCheck(uint256 amountToken0, address routerAddress0, address routerAddress1, address token0Address, address token1Address) returns()
func (_Simplev2arbabi *Simplev2arbabiTransactorSession) MakeArbitrageSimpleNoCheck(amountToken0 *big.Int, routerAddress0 common.Address, routerAddress1 common.Address, token0Address common.Address, token1Address common.Address) (*types.Transaction, error) {
	return _Simplev2arbabi.Contract.MakeArbitrageSimpleNoCheck(&_Simplev2arbabi.TransactOpts, amountToken0, routerAddress0, routerAddress1, token0Address, token1Address)
}

// MakeArbitrageSimpleWithPath is a paid mutator transaction binding the contract method 0x5ed5b76c.
//
// Solidity: function makeArbitrageSimpleWithPath(uint256 amountToken0, address routerAddress0, address routerAddress1, address[] pathRouter0, address[] pathRouter1) returns()
func (_Simplev2arbabi *Simplev2arbabiTransactor) MakeArbitrageSimpleWithPath(opts *bind.TransactOpts, amountToken0 *big.Int, routerAddress0 common.Address, routerAddress1 common.Address, pathRouter0 []common.Address, pathRouter1 []common.Address) (*types.Transaction, error) {
	return _Simplev2arbabi.contract.Transact(opts, "makeArbitrageSimpleWithPath", amountToken0, routerAddress0, routerAddress1, pathRouter0, pathRouter1)
}

// MakeArbitrageSimpleWithPath is a paid mutator transaction binding the contract method 0x5ed5b76c.
//
// Solidity: function makeArbitrageSimpleWithPath(uint256 amountToken0, address routerAddress0, address routerAddress1, address[] pathRouter0, address[] pathRouter1) returns()
func (_Simplev2arbabi *Simplev2arbabiSession) MakeArbitrageSimpleWithPath(amountToken0 *big.Int, routerAddress0 common.Address, routerAddress1 common.Address, pathRouter0 []common.Address, pathRouter1 []common.Address) (*types.Transaction, error) {
	return _Simplev2arbabi.Contract.MakeArbitrageSimpleWithPath(&_Simplev2arbabi.TransactOpts, amountToken0, routerAddress0, routerAddress1, pathRouter0, pathRouter1)
}

// MakeArbitrageSimpleWithPath is a paid mutator transaction binding the contract method 0x5ed5b76c.
//
// Solidity: function makeArbitrageSimpleWithPath(uint256 amountToken0, address routerAddress0, address routerAddress1, address[] pathRouter0, address[] pathRouter1) returns()
func (_Simplev2arbabi *Simplev2arbabiTransactorSession) MakeArbitrageSimpleWithPath(amountToken0 *big.Int, routerAddress0 common.Address, routerAddress1 common.Address, pathRouter0 []common.Address, pathRouter1 []common.Address) (*types.Transaction, error) {
	return _Simplev2arbabi.Contract.MakeArbitrageSimpleWithPath(&_Simplev2arbabi.TransactOpts, amountToken0, routerAddress0, routerAddress1, pathRouter0, pathRouter1)
}

// TestArb is a paid mutator transaction binding the contract method 0x0770d34b.
//
// Solidity: function testArb(uint256 amountToken0, uint256 expectedFinalAmount, address routerAddress0, address routerAddress1, address token0Address, address token1Address) returns()
func (_Simplev2arbabi *Simplev2arbabiTransactor) TestArb(opts *bind.TransactOpts, amountToken0 *big.Int, expectedFinalAmount *big.Int, routerAddress0 common.Address, routerAddress1 common.Address, token0Address common.Address, token1Address common.Address) (*types.Transaction, error) {
	return _Simplev2arbabi.contract.Transact(opts, "testArb", amountToken0, expectedFinalAmount, routerAddress0, routerAddress1, token0Address, token1Address)
}

// TestArb is a paid mutator transaction binding the contract method 0x0770d34b.
//
// Solidity: function testArb(uint256 amountToken0, uint256 expectedFinalAmount, address routerAddress0, address routerAddress1, address token0Address, address token1Address) returns()
func (_Simplev2arbabi *Simplev2arbabiSession) TestArb(amountToken0 *big.Int, expectedFinalAmount *big.Int, routerAddress0 common.Address, routerAddress1 common.Address, token0Address common.Address, token1Address common.Address) (*types.Transaction, error) {
	return _Simplev2arbabi.Contract.TestArb(&_Simplev2arbabi.TransactOpts, amountToken0, expectedFinalAmount, routerAddress0, routerAddress1, token0Address, token1Address)
}

// TestArb is a paid mutator transaction binding the contract method 0x0770d34b.
//
// Solidity: function testArb(uint256 amountToken0, uint256 expectedFinalAmount, address routerAddress0, address routerAddress1, address token0Address, address token1Address) returns()
func (_Simplev2arbabi *Simplev2arbabiTransactorSession) TestArb(amountToken0 *big.Int, expectedFinalAmount *big.Int, routerAddress0 common.Address, routerAddress1 common.Address, token0Address common.Address, token1Address common.Address) (*types.Transaction, error) {
	return _Simplev2arbabi.Contract.TestArb(&_Simplev2arbabi.TransactOpts, amountToken0, expectedFinalAmount, routerAddress0, routerAddress1, token0Address, token1Address)
}

// TestArbTwo is a paid mutator transaction binding the contract method 0xbff10f3e.
//
// Solidity: function testArbTwo(uint256 amountToken0, uint256 expectedFinalAmount, address routerAddress0, address routerAddress1, address token0Address, address token1Address) returns()
func (_Simplev2arbabi *Simplev2arbabiTransactor) TestArbTwo(opts *bind.TransactOpts, amountToken0 *big.Int, expectedFinalAmount *big.Int, routerAddress0 common.Address, routerAddress1 common.Address, token0Address common.Address, token1Address common.Address) (*types.Transaction, error) {
	return _Simplev2arbabi.contract.Transact(opts, "testArbTwo", amountToken0, expectedFinalAmount, routerAddress0, routerAddress1, token0Address, token1Address)
}

// TestArbTwo is a paid mutator transaction binding the contract method 0xbff10f3e.
//
// Solidity: function testArbTwo(uint256 amountToken0, uint256 expectedFinalAmount, address routerAddress0, address routerAddress1, address token0Address, address token1Address) returns()
func (_Simplev2arbabi *Simplev2arbabiSession) TestArbTwo(amountToken0 *big.Int, expectedFinalAmount *big.Int, routerAddress0 common.Address, routerAddress1 common.Address, token0Address common.Address, token1Address common.Address) (*types.Transaction, error) {
	return _Simplev2arbabi.Contract.TestArbTwo(&_Simplev2arbabi.TransactOpts, amountToken0, expectedFinalAmount, routerAddress0, routerAddress1, token0Address, token1Address)
}

// TestArbTwo is a paid mutator transaction binding the contract method 0xbff10f3e.
//
// Solidity: function testArbTwo(uint256 amountToken0, uint256 expectedFinalAmount, address routerAddress0, address routerAddress1, address token0Address, address token1Address) returns()
func (_Simplev2arbabi *Simplev2arbabiTransactorSession) TestArbTwo(amountToken0 *big.Int, expectedFinalAmount *big.Int, routerAddress0 common.Address, routerAddress1 common.Address, token0Address common.Address, token1Address common.Address) (*types.Transaction, error) {
	return _Simplev2arbabi.Contract.TestArbTwo(&_Simplev2arbabi.TransactOpts, amountToken0, expectedFinalAmount, routerAddress0, routerAddress1, token0Address, token1Address)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address tokenAddress) returns()
func (_Simplev2arbabi *Simplev2arbabiTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Simplev2arbabi.contract.Transact(opts, "withdraw", amount, tokenAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address tokenAddress) returns()
func (_Simplev2arbabi *Simplev2arbabiSession) Withdraw(amount *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Simplev2arbabi.Contract.Withdraw(&_Simplev2arbabi.TransactOpts, amount, tokenAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address tokenAddress) returns()
func (_Simplev2arbabi *Simplev2arbabiTransactorSession) Withdraw(amount *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Simplev2arbabi.Contract.Withdraw(&_Simplev2arbabi.TransactOpts, amount, tokenAddress)
}
