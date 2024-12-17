// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi_pulsex_swap_router

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

// AbiPulsexSwapRouterMetaData contains all meta data concerning the AbiPulsexSwapRouter contract.
var AbiPulsexSwapRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factoryV1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"factoryV2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stableInfo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH9\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"info\",\"type\":\"address\"}],\"name\":\"SetStableSwap\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"swapContracts\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"exactInputStableSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"swapContracts\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"exactOutputStableSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factoryV1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factoryV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"previousBlockhash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"pull\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowed\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowedIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stableSwapInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"swapExactTokensForTokensV1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"swapExactTokensForTokensV2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"swapTokensForExactTokensV1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"swapTokensForExactTokensV2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"sweepTokenWithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"sweepTokenWithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9WithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9WithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"wrapETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// AbiPulsexSwapRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use AbiPulsexSwapRouterMetaData.ABI instead.
var AbiPulsexSwapRouterABI = AbiPulsexSwapRouterMetaData.ABI

// AbiPulsexSwapRouter is an auto generated Go binding around an Ethereum contract.
type AbiPulsexSwapRouter struct {
	AbiPulsexSwapRouterCaller     // Read-only binding to the contract
	AbiPulsexSwapRouterTransactor // Write-only binding to the contract
	AbiPulsexSwapRouterFilterer   // Log filterer for contract events
}

// AbiPulsexSwapRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbiPulsexSwapRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiPulsexSwapRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbiPulsexSwapRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiPulsexSwapRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbiPulsexSwapRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiPulsexSwapRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbiPulsexSwapRouterSession struct {
	Contract     *AbiPulsexSwapRouter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AbiPulsexSwapRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbiPulsexSwapRouterCallerSession struct {
	Contract *AbiPulsexSwapRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// AbiPulsexSwapRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbiPulsexSwapRouterTransactorSession struct {
	Contract     *AbiPulsexSwapRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// AbiPulsexSwapRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbiPulsexSwapRouterRaw struct {
	Contract *AbiPulsexSwapRouter // Generic contract binding to access the raw methods on
}

// AbiPulsexSwapRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbiPulsexSwapRouterCallerRaw struct {
	Contract *AbiPulsexSwapRouterCaller // Generic read-only contract binding to access the raw methods on
}

// AbiPulsexSwapRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbiPulsexSwapRouterTransactorRaw struct {
	Contract *AbiPulsexSwapRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbiPulsexSwapRouter creates a new instance of AbiPulsexSwapRouter, bound to a specific deployed contract.
func NewAbiPulsexSwapRouter(address common.Address, backend bind.ContractBackend) (*AbiPulsexSwapRouter, error) {
	contract, err := bindAbiPulsexSwapRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AbiPulsexSwapRouter{AbiPulsexSwapRouterCaller: AbiPulsexSwapRouterCaller{contract: contract}, AbiPulsexSwapRouterTransactor: AbiPulsexSwapRouterTransactor{contract: contract}, AbiPulsexSwapRouterFilterer: AbiPulsexSwapRouterFilterer{contract: contract}}, nil
}

// NewAbiPulsexSwapRouterCaller creates a new read-only instance of AbiPulsexSwapRouter, bound to a specific deployed contract.
func NewAbiPulsexSwapRouterCaller(address common.Address, caller bind.ContractCaller) (*AbiPulsexSwapRouterCaller, error) {
	contract, err := bindAbiPulsexSwapRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbiPulsexSwapRouterCaller{contract: contract}, nil
}

// NewAbiPulsexSwapRouterTransactor creates a new write-only instance of AbiPulsexSwapRouter, bound to a specific deployed contract.
func NewAbiPulsexSwapRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*AbiPulsexSwapRouterTransactor, error) {
	contract, err := bindAbiPulsexSwapRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbiPulsexSwapRouterTransactor{contract: contract}, nil
}

// NewAbiPulsexSwapRouterFilterer creates a new log filterer instance of AbiPulsexSwapRouter, bound to a specific deployed contract.
func NewAbiPulsexSwapRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*AbiPulsexSwapRouterFilterer, error) {
	contract, err := bindAbiPulsexSwapRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbiPulsexSwapRouterFilterer{contract: contract}, nil
}

// bindAbiPulsexSwapRouter binds a generic wrapper to an already deployed contract.
func bindAbiPulsexSwapRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AbiPulsexSwapRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbiPulsexSwapRouter.Contract.AbiPulsexSwapRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.AbiPulsexSwapRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.AbiPulsexSwapRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbiPulsexSwapRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbiPulsexSwapRouter.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) WETH9() (common.Address, error) {
	return _AbiPulsexSwapRouter.Contract.WETH9(&_AbiPulsexSwapRouter.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterCallerSession) WETH9() (common.Address, error) {
	return _AbiPulsexSwapRouter.Contract.WETH9(&_AbiPulsexSwapRouter.CallOpts)
}

// FactoryV1 is a free data retrieval call binding the contract method 0x3957f453.
//
// Solidity: function factoryV1() view returns(address)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterCaller) FactoryV1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbiPulsexSwapRouter.contract.Call(opts, &out, "factoryV1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryV1 is a free data retrieval call binding the contract method 0x3957f453.
//
// Solidity: function factoryV1() view returns(address)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) FactoryV1() (common.Address, error) {
	return _AbiPulsexSwapRouter.Contract.FactoryV1(&_AbiPulsexSwapRouter.CallOpts)
}

// FactoryV1 is a free data retrieval call binding the contract method 0x3957f453.
//
// Solidity: function factoryV1() view returns(address)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterCallerSession) FactoryV1() (common.Address, error) {
	return _AbiPulsexSwapRouter.Contract.FactoryV1(&_AbiPulsexSwapRouter.CallOpts)
}

// FactoryV2 is a free data retrieval call binding the contract method 0x68e0d4e1.
//
// Solidity: function factoryV2() view returns(address)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterCaller) FactoryV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbiPulsexSwapRouter.contract.Call(opts, &out, "factoryV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryV2 is a free data retrieval call binding the contract method 0x68e0d4e1.
//
// Solidity: function factoryV2() view returns(address)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) FactoryV2() (common.Address, error) {
	return _AbiPulsexSwapRouter.Contract.FactoryV2(&_AbiPulsexSwapRouter.CallOpts)
}

// FactoryV2 is a free data retrieval call binding the contract method 0x68e0d4e1.
//
// Solidity: function factoryV2() view returns(address)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterCallerSession) FactoryV2() (common.Address, error) {
	return _AbiPulsexSwapRouter.Contract.FactoryV2(&_AbiPulsexSwapRouter.CallOpts)
}

// StableSwapInfo is a free data retrieval call binding the contract method 0xb85aa7af.
//
// Solidity: function stableSwapInfo() view returns(address)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterCaller) StableSwapInfo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbiPulsexSwapRouter.contract.Call(opts, &out, "stableSwapInfo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StableSwapInfo is a free data retrieval call binding the contract method 0xb85aa7af.
//
// Solidity: function stableSwapInfo() view returns(address)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) StableSwapInfo() (common.Address, error) {
	return _AbiPulsexSwapRouter.Contract.StableSwapInfo(&_AbiPulsexSwapRouter.CallOpts)
}

// StableSwapInfo is a free data retrieval call binding the contract method 0xb85aa7af.
//
// Solidity: function stableSwapInfo() view returns(address)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterCallerSession) StableSwapInfo() (common.Address, error) {
	return _AbiPulsexSwapRouter.Contract.StableSwapInfo(&_AbiPulsexSwapRouter.CallOpts)
}

// ExactInputStableSwap is a paid mutator transaction binding the contract method 0x4aa94288.
//
// Solidity: function exactInputStableSwap(address[] path, address[] swapContracts, uint256 amountIn, uint256 amountOutMin, address to) payable returns(uint256 amountOut)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactor) ExactInputStableSwap(opts *bind.TransactOpts, path []common.Address, swapContracts []common.Address, amountIn *big.Int, amountOutMin *big.Int, to common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.contract.Transact(opts, "exactInputStableSwap", path, swapContracts, amountIn, amountOutMin, to)
}

// ExactInputStableSwap is a paid mutator transaction binding the contract method 0x4aa94288.
//
// Solidity: function exactInputStableSwap(address[] path, address[] swapContracts, uint256 amountIn, uint256 amountOutMin, address to) payable returns(uint256 amountOut)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) ExactInputStableSwap(path []common.Address, swapContracts []common.Address, amountIn *big.Int, amountOutMin *big.Int, to common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.ExactInputStableSwap(&_AbiPulsexSwapRouter.TransactOpts, path, swapContracts, amountIn, amountOutMin, to)
}

// ExactInputStableSwap is a paid mutator transaction binding the contract method 0x4aa94288.
//
// Solidity: function exactInputStableSwap(address[] path, address[] swapContracts, uint256 amountIn, uint256 amountOutMin, address to) payable returns(uint256 amountOut)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorSession) ExactInputStableSwap(path []common.Address, swapContracts []common.Address, amountIn *big.Int, amountOutMin *big.Int, to common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.ExactInputStableSwap(&_AbiPulsexSwapRouter.TransactOpts, path, swapContracts, amountIn, amountOutMin, to)
}

// ExactOutputStableSwap is a paid mutator transaction binding the contract method 0xc8bb1856.
//
// Solidity: function exactOutputStableSwap(address[] path, address[] swapContracts, uint256 amountOut, uint256 amountInMax, address to) payable returns(uint256 amountIn)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactor) ExactOutputStableSwap(opts *bind.TransactOpts, path []common.Address, swapContracts []common.Address, amountOut *big.Int, amountInMax *big.Int, to common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.contract.Transact(opts, "exactOutputStableSwap", path, swapContracts, amountOut, amountInMax, to)
}

// ExactOutputStableSwap is a paid mutator transaction binding the contract method 0xc8bb1856.
//
// Solidity: function exactOutputStableSwap(address[] path, address[] swapContracts, uint256 amountOut, uint256 amountInMax, address to) payable returns(uint256 amountIn)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) ExactOutputStableSwap(path []common.Address, swapContracts []common.Address, amountOut *big.Int, amountInMax *big.Int, to common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.ExactOutputStableSwap(&_AbiPulsexSwapRouter.TransactOpts, path, swapContracts, amountOut, amountInMax, to)
}

// ExactOutputStableSwap is a paid mutator transaction binding the contract method 0xc8bb1856.
//
// Solidity: function exactOutputStableSwap(address[] path, address[] swapContracts, uint256 amountOut, uint256 amountInMax, address to) payable returns(uint256 amountIn)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorSession) ExactOutputStableSwap(path []common.Address, swapContracts []common.Address, amountOut *big.Int, amountInMax *big.Int, to common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.ExactOutputStableSwap(&_AbiPulsexSwapRouter.TransactOpts, path, swapContracts, amountOut, amountInMax, to)
}

// Multicall is a paid mutator transaction binding the contract method 0x1f0464d1.
//
// Solidity: function multicall(bytes32 previousBlockhash, bytes[] data) payable returns(bytes[])
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactor) Multicall(opts *bind.TransactOpts, previousBlockhash [32]byte, data [][]byte) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.contract.Transact(opts, "multicall", previousBlockhash, data)
}

// Multicall is a paid mutator transaction binding the contract method 0x1f0464d1.
//
// Solidity: function multicall(bytes32 previousBlockhash, bytes[] data) payable returns(bytes[])
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) Multicall(previousBlockhash [32]byte, data [][]byte) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.Multicall(&_AbiPulsexSwapRouter.TransactOpts, previousBlockhash, data)
}

// Multicall is a paid mutator transaction binding the contract method 0x1f0464d1.
//
// Solidity: function multicall(bytes32 previousBlockhash, bytes[] data) payable returns(bytes[])
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorSession) Multicall(previousBlockhash [32]byte, data [][]byte) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.Multicall(&_AbiPulsexSwapRouter.TransactOpts, previousBlockhash, data)
}

// Multicall0 is a paid mutator transaction binding the contract method 0x5ae401dc.
//
// Solidity: function multicall(uint256 deadline, bytes[] data) payable returns(bytes[])
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactor) Multicall0(opts *bind.TransactOpts, deadline *big.Int, data [][]byte) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.contract.Transact(opts, "multicall0", deadline, data)
}

// Multicall0 is a paid mutator transaction binding the contract method 0x5ae401dc.
//
// Solidity: function multicall(uint256 deadline, bytes[] data) payable returns(bytes[])
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) Multicall0(deadline *big.Int, data [][]byte) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.Multicall0(&_AbiPulsexSwapRouter.TransactOpts, deadline, data)
}

// Multicall0 is a paid mutator transaction binding the contract method 0x5ae401dc.
//
// Solidity: function multicall(uint256 deadline, bytes[] data) payable returns(bytes[])
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorSession) Multicall0(deadline *big.Int, data [][]byte) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.Multicall0(&_AbiPulsexSwapRouter.TransactOpts, deadline, data)
}

// Multicall1 is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactor) Multicall1(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.contract.Transact(opts, "multicall1", data)
}

// Multicall1 is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) Multicall1(data [][]byte) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.Multicall1(&_AbiPulsexSwapRouter.TransactOpts, data)
}

// Multicall1 is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorSession) Multicall1(data [][]byte) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.Multicall1(&_AbiPulsexSwapRouter.TransactOpts, data)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address token, uint256 value) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactor) Pull(opts *bind.TransactOpts, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.contract.Transact(opts, "pull", token, value)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address token, uint256 value) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) Pull(token common.Address, value *big.Int) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.Pull(&_AbiPulsexSwapRouter.TransactOpts, token, value)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address token, uint256 value) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorSession) Pull(token common.Address, value *big.Int) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.Pull(&_AbiPulsexSwapRouter.TransactOpts, token, value)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactor) RefundETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.contract.Transact(opts, "refundETH")
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) RefundETH() (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.RefundETH(&_AbiPulsexSwapRouter.TransactOpts)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorSession) RefundETH() (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.RefundETH(&_AbiPulsexSwapRouter.TransactOpts)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactor) SelfPermit(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.contract.Transact(opts, "selfPermit", token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.SelfPermit(&_AbiPulsexSwapRouter.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.SelfPermit(&_AbiPulsexSwapRouter.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactor) SelfPermitAllowed(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.contract.Transact(opts, "selfPermitAllowed", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.SelfPermitAllowed(&_AbiPulsexSwapRouter.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.SelfPermitAllowed(&_AbiPulsexSwapRouter.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactor) SelfPermitAllowedIfNecessary(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.contract.Transact(opts, "selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.SelfPermitAllowedIfNecessary(&_AbiPulsexSwapRouter.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.SelfPermitAllowedIfNecessary(&_AbiPulsexSwapRouter.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactor) SelfPermitIfNecessary(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.contract.Transact(opts, "selfPermitIfNecessary", token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.SelfPermitIfNecessary(&_AbiPulsexSwapRouter.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.SelfPermitIfNecessary(&_AbiPulsexSwapRouter.TransactOpts, token, value, deadline, v, r, s)
}

// SwapExactTokensForTokensV1 is a paid mutator transaction binding the contract method 0x9b5c7a81.
//
// Solidity: function swapExactTokensForTokensV1(uint256 amountIn, uint256 amountOutMin, address[] path, address to) payable returns(uint256 amountOut)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactor) SwapExactTokensForTokensV1(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.contract.Transact(opts, "swapExactTokensForTokensV1", amountIn, amountOutMin, path, to)
}

// SwapExactTokensForTokensV1 is a paid mutator transaction binding the contract method 0x9b5c7a81.
//
// Solidity: function swapExactTokensForTokensV1(uint256 amountIn, uint256 amountOutMin, address[] path, address to) payable returns(uint256 amountOut)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) SwapExactTokensForTokensV1(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.SwapExactTokensForTokensV1(&_AbiPulsexSwapRouter.TransactOpts, amountIn, amountOutMin, path, to)
}

// SwapExactTokensForTokensV1 is a paid mutator transaction binding the contract method 0x9b5c7a81.
//
// Solidity: function swapExactTokensForTokensV1(uint256 amountIn, uint256 amountOutMin, address[] path, address to) payable returns(uint256 amountOut)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorSession) SwapExactTokensForTokensV1(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.SwapExactTokensForTokensV1(&_AbiPulsexSwapRouter.TransactOpts, amountIn, amountOutMin, path, to)
}

// SwapExactTokensForTokensV2 is a paid mutator transaction binding the contract method 0xab0acea4.
//
// Solidity: function swapExactTokensForTokensV2(uint256 amountIn, uint256 amountOutMin, address[] path, address to) payable returns(uint256 amountOut)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactor) SwapExactTokensForTokensV2(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.contract.Transact(opts, "swapExactTokensForTokensV2", amountIn, amountOutMin, path, to)
}

// SwapExactTokensForTokensV2 is a paid mutator transaction binding the contract method 0xab0acea4.
//
// Solidity: function swapExactTokensForTokensV2(uint256 amountIn, uint256 amountOutMin, address[] path, address to) payable returns(uint256 amountOut)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) SwapExactTokensForTokensV2(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.SwapExactTokensForTokensV2(&_AbiPulsexSwapRouter.TransactOpts, amountIn, amountOutMin, path, to)
}

// SwapExactTokensForTokensV2 is a paid mutator transaction binding the contract method 0xab0acea4.
//
// Solidity: function swapExactTokensForTokensV2(uint256 amountIn, uint256 amountOutMin, address[] path, address to) payable returns(uint256 amountOut)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorSession) SwapExactTokensForTokensV2(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.SwapExactTokensForTokensV2(&_AbiPulsexSwapRouter.TransactOpts, amountIn, amountOutMin, path, to)
}

// SwapTokensForExactTokensV1 is a paid mutator transaction binding the contract method 0x1d4389f5.
//
// Solidity: function swapTokensForExactTokensV1(uint256 amountOut, uint256 amountInMax, address[] path, address to) payable returns(uint256 amountIn)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactor) SwapTokensForExactTokensV1(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.contract.Transact(opts, "swapTokensForExactTokensV1", amountOut, amountInMax, path, to)
}

// SwapTokensForExactTokensV1 is a paid mutator transaction binding the contract method 0x1d4389f5.
//
// Solidity: function swapTokensForExactTokensV1(uint256 amountOut, uint256 amountInMax, address[] path, address to) payable returns(uint256 amountIn)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) SwapTokensForExactTokensV1(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.SwapTokensForExactTokensV1(&_AbiPulsexSwapRouter.TransactOpts, amountOut, amountInMax, path, to)
}

// SwapTokensForExactTokensV1 is a paid mutator transaction binding the contract method 0x1d4389f5.
//
// Solidity: function swapTokensForExactTokensV1(uint256 amountOut, uint256 amountInMax, address[] path, address to) payable returns(uint256 amountIn)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorSession) SwapTokensForExactTokensV1(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.SwapTokensForExactTokensV1(&_AbiPulsexSwapRouter.TransactOpts, amountOut, amountInMax, path, to)
}

// SwapTokensForExactTokensV2 is a paid mutator transaction binding the contract method 0xcafe95a0.
//
// Solidity: function swapTokensForExactTokensV2(uint256 amountOut, uint256 amountInMax, address[] path, address to) payable returns(uint256 amountIn)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactor) SwapTokensForExactTokensV2(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.contract.Transact(opts, "swapTokensForExactTokensV2", amountOut, amountInMax, path, to)
}

// SwapTokensForExactTokensV2 is a paid mutator transaction binding the contract method 0xcafe95a0.
//
// Solidity: function swapTokensForExactTokensV2(uint256 amountOut, uint256 amountInMax, address[] path, address to) payable returns(uint256 amountIn)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) SwapTokensForExactTokensV2(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.SwapTokensForExactTokensV2(&_AbiPulsexSwapRouter.TransactOpts, amountOut, amountInMax, path, to)
}

// SwapTokensForExactTokensV2 is a paid mutator transaction binding the contract method 0xcafe95a0.
//
// Solidity: function swapTokensForExactTokensV2(uint256 amountOut, uint256 amountInMax, address[] path, address to) payable returns(uint256 amountIn)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorSession) SwapTokensForExactTokensV2(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.SwapTokensForExactTokensV2(&_AbiPulsexSwapRouter.TransactOpts, amountOut, amountInMax, path, to)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactor) SweepToken(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.contract.Transact(opts, "sweepToken", token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.SweepToken(&_AbiPulsexSwapRouter.TransactOpts, token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.SweepToken(&_AbiPulsexSwapRouter.TransactOpts, token, amountMinimum, recipient)
}

// SweepToken0 is a paid mutator transaction binding the contract method 0xe90a182f.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactor) SweepToken0(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.contract.Transact(opts, "sweepToken0", token, amountMinimum)
}

// SweepToken0 is a paid mutator transaction binding the contract method 0xe90a182f.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) SweepToken0(token common.Address, amountMinimum *big.Int) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.SweepToken0(&_AbiPulsexSwapRouter.TransactOpts, token, amountMinimum)
}

// SweepToken0 is a paid mutator transaction binding the contract method 0xe90a182f.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorSession) SweepToken0(token common.Address, amountMinimum *big.Int) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.SweepToken0(&_AbiPulsexSwapRouter.TransactOpts, token, amountMinimum)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0x3068c554.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactor) SweepTokenWithFee(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.contract.Transact(opts, "sweepTokenWithFee", token, amountMinimum, feeBips, feeRecipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0x3068c554.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) SweepTokenWithFee(token common.Address, amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.SweepTokenWithFee(&_AbiPulsexSwapRouter.TransactOpts, token, amountMinimum, feeBips, feeRecipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0x3068c554.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorSession) SweepTokenWithFee(token common.Address, amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.SweepTokenWithFee(&_AbiPulsexSwapRouter.TransactOpts, token, amountMinimum, feeBips, feeRecipient)
}

// SweepTokenWithFee0 is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactor) SweepTokenWithFee0(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.contract.Transact(opts, "sweepTokenWithFee0", token, amountMinimum, recipient, feeBips, feeRecipient)
}

// SweepTokenWithFee0 is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) SweepTokenWithFee0(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.SweepTokenWithFee0(&_AbiPulsexSwapRouter.TransactOpts, token, amountMinimum, recipient, feeBips, feeRecipient)
}

// SweepTokenWithFee0 is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorSession) SweepTokenWithFee0(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.SweepTokenWithFee0(&_AbiPulsexSwapRouter.TransactOpts, token, amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactor) UnwrapWETH9(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.contract.Transact(opts, "unwrapWETH9", amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.UnwrapWETH9(&_AbiPulsexSwapRouter.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.UnwrapWETH9(&_AbiPulsexSwapRouter.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH90 is a paid mutator transaction binding the contract method 0x49616997.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactor) UnwrapWETH90(opts *bind.TransactOpts, amountMinimum *big.Int) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.contract.Transact(opts, "unwrapWETH90", amountMinimum)
}

// UnwrapWETH90 is a paid mutator transaction binding the contract method 0x49616997.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) UnwrapWETH90(amountMinimum *big.Int) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.UnwrapWETH90(&_AbiPulsexSwapRouter.TransactOpts, amountMinimum)
}

// UnwrapWETH90 is a paid mutator transaction binding the contract method 0x49616997.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorSession) UnwrapWETH90(amountMinimum *big.Int) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.UnwrapWETH90(&_AbiPulsexSwapRouter.TransactOpts, amountMinimum)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactor) UnwrapWETH9WithFee(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.contract.Transact(opts, "unwrapWETH9WithFee", amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) UnwrapWETH9WithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.UnwrapWETH9WithFee(&_AbiPulsexSwapRouter.TransactOpts, amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorSession) UnwrapWETH9WithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.UnwrapWETH9WithFee(&_AbiPulsexSwapRouter.TransactOpts, amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee0 is a paid mutator transaction binding the contract method 0xd4ef38de.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactor) UnwrapWETH9WithFee0(opts *bind.TransactOpts, amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.contract.Transact(opts, "unwrapWETH9WithFee0", amountMinimum, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee0 is a paid mutator transaction binding the contract method 0xd4ef38de.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) UnwrapWETH9WithFee0(amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.UnwrapWETH9WithFee0(&_AbiPulsexSwapRouter.TransactOpts, amountMinimum, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee0 is a paid mutator transaction binding the contract method 0xd4ef38de.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorSession) UnwrapWETH9WithFee0(amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.UnwrapWETH9WithFee0(&_AbiPulsexSwapRouter.TransactOpts, amountMinimum, feeBips, feeRecipient)
}

// WrapETH is a paid mutator transaction binding the contract method 0x1c58db4f.
//
// Solidity: function wrapETH(uint256 value) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactor) WrapETH(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.contract.Transact(opts, "wrapETH", value)
}

// WrapETH is a paid mutator transaction binding the contract method 0x1c58db4f.
//
// Solidity: function wrapETH(uint256 value) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) WrapETH(value *big.Int) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.WrapETH(&_AbiPulsexSwapRouter.TransactOpts, value)
}

// WrapETH is a paid mutator transaction binding the contract method 0x1c58db4f.
//
// Solidity: function wrapETH(uint256 value) payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorSession) WrapETH(value *big.Int) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.WrapETH(&_AbiPulsexSwapRouter.TransactOpts, value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterSession) Receive() (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.Receive(&_AbiPulsexSwapRouter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterTransactorSession) Receive() (*types.Transaction, error) {
	return _AbiPulsexSwapRouter.Contract.Receive(&_AbiPulsexSwapRouter.TransactOpts)
}

// AbiPulsexSwapRouterSetStableSwapIterator is returned from FilterSetStableSwap and is used to iterate over the raw logs and unpacked data for SetStableSwap events raised by the AbiPulsexSwapRouter contract.
type AbiPulsexSwapRouterSetStableSwapIterator struct {
	Event *AbiPulsexSwapRouterSetStableSwap // Event containing the contract specifics and raw log

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
func (it *AbiPulsexSwapRouterSetStableSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiPulsexSwapRouterSetStableSwap)
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
		it.Event = new(AbiPulsexSwapRouterSetStableSwap)
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
func (it *AbiPulsexSwapRouterSetStableSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiPulsexSwapRouterSetStableSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiPulsexSwapRouterSetStableSwap represents a SetStableSwap event raised by the AbiPulsexSwapRouter contract.
type AbiPulsexSwapRouterSetStableSwap struct {
	Factory common.Address
	Info    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetStableSwap is a free log retrieval operation binding the contract event 0x26e41379222b54b0470031bc11852ad23058ffb8983f7cc0e18257d6f7afca9d.
//
// Solidity: event SetStableSwap(address indexed factory, address indexed info)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterFilterer) FilterSetStableSwap(opts *bind.FilterOpts, factory []common.Address, info []common.Address) (*AbiPulsexSwapRouterSetStableSwapIterator, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}
	var infoRule []interface{}
	for _, infoItem := range info {
		infoRule = append(infoRule, infoItem)
	}

	logs, sub, err := _AbiPulsexSwapRouter.contract.FilterLogs(opts, "SetStableSwap", factoryRule, infoRule)
	if err != nil {
		return nil, err
	}
	return &AbiPulsexSwapRouterSetStableSwapIterator{contract: _AbiPulsexSwapRouter.contract, event: "SetStableSwap", logs: logs, sub: sub}, nil
}

// WatchSetStableSwap is a free log subscription operation binding the contract event 0x26e41379222b54b0470031bc11852ad23058ffb8983f7cc0e18257d6f7afca9d.
//
// Solidity: event SetStableSwap(address indexed factory, address indexed info)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterFilterer) WatchSetStableSwap(opts *bind.WatchOpts, sink chan<- *AbiPulsexSwapRouterSetStableSwap, factory []common.Address, info []common.Address) (event.Subscription, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}
	var infoRule []interface{}
	for _, infoItem := range info {
		infoRule = append(infoRule, infoItem)
	}

	logs, sub, err := _AbiPulsexSwapRouter.contract.WatchLogs(opts, "SetStableSwap", factoryRule, infoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiPulsexSwapRouterSetStableSwap)
				if err := _AbiPulsexSwapRouter.contract.UnpackLog(event, "SetStableSwap", log); err != nil {
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

// ParseSetStableSwap is a log parse operation binding the contract event 0x26e41379222b54b0470031bc11852ad23058ffb8983f7cc0e18257d6f7afca9d.
//
// Solidity: event SetStableSwap(address indexed factory, address indexed info)
func (_AbiPulsexSwapRouter *AbiPulsexSwapRouterFilterer) ParseSetStableSwap(log types.Log) (*AbiPulsexSwapRouterSetStableSwap, error) {
	event := new(AbiPulsexSwapRouterSetStableSwap)
	if err := _AbiPulsexSwapRouter.contract.UnpackLog(event, "SetStableSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
