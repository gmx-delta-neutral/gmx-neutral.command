// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package glpmanager

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
)

// GlpmanagerMetaData contains all meta data concerning the Glpmanager contract.
var GlpmanagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdg\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_glp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_cooldownDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"aumInUsdg\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"glpSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdgAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"AddLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"glpAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"aumInUsdg\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"glpSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdgAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidity\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_COOLDOWN_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRICE_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDG_DECIMALS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minUsdg\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minGlp\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fundingAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minUsdg\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minGlp\",\"type\":\"uint256\"}],\"name\":\"addLiquidityForAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aumAddition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aumDeduction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cooldownDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"maximise\",\"type\":\"bool\"}],\"name\":\"getAum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"maximise\",\"type\":\"bool\"}],\"name\":\"getAumInUsdg\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAums\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inPrivateMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isHandler\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastAddedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_glpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_glpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"removeLiquidityForAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_aumAddition\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_aumDeduction\",\"type\":\"uint256\"}],\"name\":\"setAumAdjustment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cooldownDuration\",\"type\":\"uint256\"}],\"name\":\"setCooldownDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_handler\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_inPrivateMode\",\"type\":\"bool\"}],\"name\":\"setInPrivateMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdg\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GlpmanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use GlpmanagerMetaData.ABI instead.
var GlpmanagerABI = GlpmanagerMetaData.ABI

// Glpmanager is an auto generated Go binding around an Ethereum contract.
type Glpmanager struct {
	GlpmanagerCaller     // Read-only binding to the contract
	GlpmanagerTransactor // Write-only binding to the contract
	GlpmanagerFilterer   // Log filterer for contract events
}

// GlpmanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type GlpmanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlpmanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GlpmanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlpmanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GlpmanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlpmanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GlpmanagerSession struct {
	Contract     *Glpmanager       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GlpmanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GlpmanagerCallerSession struct {
	Contract *GlpmanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GlpmanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GlpmanagerTransactorSession struct {
	Contract     *GlpmanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GlpmanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type GlpmanagerRaw struct {
	Contract *Glpmanager // Generic contract binding to access the raw methods on
}

// GlpmanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GlpmanagerCallerRaw struct {
	Contract *GlpmanagerCaller // Generic read-only contract binding to access the raw methods on
}

// GlpmanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GlpmanagerTransactorRaw struct {
	Contract *GlpmanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlpmanager creates a new instance of Glpmanager, bound to a specific deployed contract.
func NewGlpmanager(address common.Address, backend bind.ContractBackend) (*Glpmanager, error) {
	contract, err := bindGlpmanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Glpmanager{GlpmanagerCaller: GlpmanagerCaller{contract: contract}, GlpmanagerTransactor: GlpmanagerTransactor{contract: contract}, GlpmanagerFilterer: GlpmanagerFilterer{contract: contract}}, nil
}

// NewGlpmanagerCaller creates a new read-only instance of Glpmanager, bound to a specific deployed contract.
func NewGlpmanagerCaller(address common.Address, caller bind.ContractCaller) (*GlpmanagerCaller, error) {
	contract, err := bindGlpmanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GlpmanagerCaller{contract: contract}, nil
}

// NewGlpmanagerTransactor creates a new write-only instance of Glpmanager, bound to a specific deployed contract.
func NewGlpmanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*GlpmanagerTransactor, error) {
	contract, err := bindGlpmanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GlpmanagerTransactor{contract: contract}, nil
}

// NewGlpmanagerFilterer creates a new log filterer instance of Glpmanager, bound to a specific deployed contract.
func NewGlpmanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*GlpmanagerFilterer, error) {
	contract, err := bindGlpmanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GlpmanagerFilterer{contract: contract}, nil
}

// bindGlpmanager binds a generic wrapper to an already deployed contract.
func bindGlpmanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GlpmanagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Glpmanager *GlpmanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Glpmanager.Contract.GlpmanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Glpmanager *GlpmanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Glpmanager.Contract.GlpmanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Glpmanager *GlpmanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Glpmanager.Contract.GlpmanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Glpmanager *GlpmanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Glpmanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Glpmanager *GlpmanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Glpmanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Glpmanager *GlpmanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Glpmanager.Contract.contract.Transact(opts, method, params...)
}

// MAXCOOLDOWNDURATION is a free data retrieval call binding the contract method 0x1e9049cf.
//
// Solidity: function MAX_COOLDOWN_DURATION() view returns(uint256)
func (_Glpmanager *GlpmanagerCaller) MAXCOOLDOWNDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Glpmanager.contract.Call(opts, &out, "MAX_COOLDOWN_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXCOOLDOWNDURATION is a free data retrieval call binding the contract method 0x1e9049cf.
//
// Solidity: function MAX_COOLDOWN_DURATION() view returns(uint256)
func (_Glpmanager *GlpmanagerSession) MAXCOOLDOWNDURATION() (*big.Int, error) {
	return _Glpmanager.Contract.MAXCOOLDOWNDURATION(&_Glpmanager.CallOpts)
}

// MAXCOOLDOWNDURATION is a free data retrieval call binding the contract method 0x1e9049cf.
//
// Solidity: function MAX_COOLDOWN_DURATION() view returns(uint256)
func (_Glpmanager *GlpmanagerCallerSession) MAXCOOLDOWNDURATION() (*big.Int, error) {
	return _Glpmanager.Contract.MAXCOOLDOWNDURATION(&_Glpmanager.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_Glpmanager *GlpmanagerCaller) PRICEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Glpmanager.contract.Call(opts, &out, "PRICE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_Glpmanager *GlpmanagerSession) PRICEPRECISION() (*big.Int, error) {
	return _Glpmanager.Contract.PRICEPRECISION(&_Glpmanager.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_Glpmanager *GlpmanagerCallerSession) PRICEPRECISION() (*big.Int, error) {
	return _Glpmanager.Contract.PRICEPRECISION(&_Glpmanager.CallOpts)
}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_Glpmanager *GlpmanagerCaller) USDGDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Glpmanager.contract.Call(opts, &out, "USDG_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_Glpmanager *GlpmanagerSession) USDGDECIMALS() (*big.Int, error) {
	return _Glpmanager.Contract.USDGDECIMALS(&_Glpmanager.CallOpts)
}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_Glpmanager *GlpmanagerCallerSession) USDGDECIMALS() (*big.Int, error) {
	return _Glpmanager.Contract.USDGDECIMALS(&_Glpmanager.CallOpts)
}

// AumAddition is a free data retrieval call binding the contract method 0x196b68cb.
//
// Solidity: function aumAddition() view returns(uint256)
func (_Glpmanager *GlpmanagerCaller) AumAddition(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Glpmanager.contract.Call(opts, &out, "aumAddition")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AumAddition is a free data retrieval call binding the contract method 0x196b68cb.
//
// Solidity: function aumAddition() view returns(uint256)
func (_Glpmanager *GlpmanagerSession) AumAddition() (*big.Int, error) {
	return _Glpmanager.Contract.AumAddition(&_Glpmanager.CallOpts)
}

// AumAddition is a free data retrieval call binding the contract method 0x196b68cb.
//
// Solidity: function aumAddition() view returns(uint256)
func (_Glpmanager *GlpmanagerCallerSession) AumAddition() (*big.Int, error) {
	return _Glpmanager.Contract.AumAddition(&_Glpmanager.CallOpts)
}

// AumDeduction is a free data retrieval call binding the contract method 0xb172bb0c.
//
// Solidity: function aumDeduction() view returns(uint256)
func (_Glpmanager *GlpmanagerCaller) AumDeduction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Glpmanager.contract.Call(opts, &out, "aumDeduction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AumDeduction is a free data retrieval call binding the contract method 0xb172bb0c.
//
// Solidity: function aumDeduction() view returns(uint256)
func (_Glpmanager *GlpmanagerSession) AumDeduction() (*big.Int, error) {
	return _Glpmanager.Contract.AumDeduction(&_Glpmanager.CallOpts)
}

// AumDeduction is a free data retrieval call binding the contract method 0xb172bb0c.
//
// Solidity: function aumDeduction() view returns(uint256)
func (_Glpmanager *GlpmanagerCallerSession) AumDeduction() (*big.Int, error) {
	return _Glpmanager.Contract.AumDeduction(&_Glpmanager.CallOpts)
}

// CooldownDuration is a free data retrieval call binding the contract method 0x35269315.
//
// Solidity: function cooldownDuration() view returns(uint256)
func (_Glpmanager *GlpmanagerCaller) CooldownDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Glpmanager.contract.Call(opts, &out, "cooldownDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CooldownDuration is a free data retrieval call binding the contract method 0x35269315.
//
// Solidity: function cooldownDuration() view returns(uint256)
func (_Glpmanager *GlpmanagerSession) CooldownDuration() (*big.Int, error) {
	return _Glpmanager.Contract.CooldownDuration(&_Glpmanager.CallOpts)
}

// CooldownDuration is a free data retrieval call binding the contract method 0x35269315.
//
// Solidity: function cooldownDuration() view returns(uint256)
func (_Glpmanager *GlpmanagerCallerSession) CooldownDuration() (*big.Int, error) {
	return _Glpmanager.Contract.CooldownDuration(&_Glpmanager.CallOpts)
}

// GetAum is a free data retrieval call binding the contract method 0x03391476.
//
// Solidity: function getAum(bool maximise) view returns(uint256)
func (_Glpmanager *GlpmanagerCaller) GetAum(opts *bind.CallOpts, maximise bool) (*big.Int, error) {
	var out []interface{}
	err := _Glpmanager.contract.Call(opts, &out, "getAum", maximise)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAum is a free data retrieval call binding the contract method 0x03391476.
//
// Solidity: function getAum(bool maximise) view returns(uint256)
func (_Glpmanager *GlpmanagerSession) GetAum(maximise bool) (*big.Int, error) {
	return _Glpmanager.Contract.GetAum(&_Glpmanager.CallOpts, maximise)
}

// GetAum is a free data retrieval call binding the contract method 0x03391476.
//
// Solidity: function getAum(bool maximise) view returns(uint256)
func (_Glpmanager *GlpmanagerCallerSession) GetAum(maximise bool) (*big.Int, error) {
	return _Glpmanager.Contract.GetAum(&_Glpmanager.CallOpts, maximise)
}

// GetAumInUsdg is a free data retrieval call binding the contract method 0x68a0a3e0.
//
// Solidity: function getAumInUsdg(bool maximise) view returns(uint256)
func (_Glpmanager *GlpmanagerCaller) GetAumInUsdg(opts *bind.CallOpts, maximise bool) (*big.Int, error) {
	var out []interface{}
	err := _Glpmanager.contract.Call(opts, &out, "getAumInUsdg", maximise)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAumInUsdg is a free data retrieval call binding the contract method 0x68a0a3e0.
//
// Solidity: function getAumInUsdg(bool maximise) view returns(uint256)
func (_Glpmanager *GlpmanagerSession) GetAumInUsdg(maximise bool) (*big.Int, error) {
	return _Glpmanager.Contract.GetAumInUsdg(&_Glpmanager.CallOpts, maximise)
}

// GetAumInUsdg is a free data retrieval call binding the contract method 0x68a0a3e0.
//
// Solidity: function getAumInUsdg(bool maximise) view returns(uint256)
func (_Glpmanager *GlpmanagerCallerSession) GetAumInUsdg(maximise bool) (*big.Int, error) {
	return _Glpmanager.Contract.GetAumInUsdg(&_Glpmanager.CallOpts, maximise)
}

// GetAums is a free data retrieval call binding the contract method 0xed0d1c04.
//
// Solidity: function getAums() view returns(uint256[])
func (_Glpmanager *GlpmanagerCaller) GetAums(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Glpmanager.contract.Call(opts, &out, "getAums")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAums is a free data retrieval call binding the contract method 0xed0d1c04.
//
// Solidity: function getAums() view returns(uint256[])
func (_Glpmanager *GlpmanagerSession) GetAums() ([]*big.Int, error) {
	return _Glpmanager.Contract.GetAums(&_Glpmanager.CallOpts)
}

// GetAums is a free data retrieval call binding the contract method 0xed0d1c04.
//
// Solidity: function getAums() view returns(uint256[])
func (_Glpmanager *GlpmanagerCallerSession) GetAums() ([]*big.Int, error) {
	return _Glpmanager.Contract.GetAums(&_Glpmanager.CallOpts)
}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_Glpmanager *GlpmanagerCaller) Glp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Glpmanager.contract.Call(opts, &out, "glp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_Glpmanager *GlpmanagerSession) Glp() (common.Address, error) {
	return _Glpmanager.Contract.Glp(&_Glpmanager.CallOpts)
}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_Glpmanager *GlpmanagerCallerSession) Glp() (common.Address, error) {
	return _Glpmanager.Contract.Glp(&_Glpmanager.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Glpmanager *GlpmanagerCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Glpmanager.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Glpmanager *GlpmanagerSession) Gov() (common.Address, error) {
	return _Glpmanager.Contract.Gov(&_Glpmanager.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Glpmanager *GlpmanagerCallerSession) Gov() (common.Address, error) {
	return _Glpmanager.Contract.Gov(&_Glpmanager.CallOpts)
}

// InPrivateMode is a free data retrieval call binding the contract method 0x070eacee.
//
// Solidity: function inPrivateMode() view returns(bool)
func (_Glpmanager *GlpmanagerCaller) InPrivateMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Glpmanager.contract.Call(opts, &out, "inPrivateMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InPrivateMode is a free data retrieval call binding the contract method 0x070eacee.
//
// Solidity: function inPrivateMode() view returns(bool)
func (_Glpmanager *GlpmanagerSession) InPrivateMode() (bool, error) {
	return _Glpmanager.Contract.InPrivateMode(&_Glpmanager.CallOpts)
}

// InPrivateMode is a free data retrieval call binding the contract method 0x070eacee.
//
// Solidity: function inPrivateMode() view returns(bool)
func (_Glpmanager *GlpmanagerCallerSession) InPrivateMode() (bool, error) {
	return _Glpmanager.Contract.InPrivateMode(&_Glpmanager.CallOpts)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_Glpmanager *GlpmanagerCaller) IsHandler(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Glpmanager.contract.Call(opts, &out, "isHandler", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_Glpmanager *GlpmanagerSession) IsHandler(arg0 common.Address) (bool, error) {
	return _Glpmanager.Contract.IsHandler(&_Glpmanager.CallOpts, arg0)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_Glpmanager *GlpmanagerCallerSession) IsHandler(arg0 common.Address) (bool, error) {
	return _Glpmanager.Contract.IsHandler(&_Glpmanager.CallOpts, arg0)
}

// LastAddedAt is a free data retrieval call binding the contract method 0x8b770e11.
//
// Solidity: function lastAddedAt(address ) view returns(uint256)
func (_Glpmanager *GlpmanagerCaller) LastAddedAt(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Glpmanager.contract.Call(opts, &out, "lastAddedAt", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastAddedAt is a free data retrieval call binding the contract method 0x8b770e11.
//
// Solidity: function lastAddedAt(address ) view returns(uint256)
func (_Glpmanager *GlpmanagerSession) LastAddedAt(arg0 common.Address) (*big.Int, error) {
	return _Glpmanager.Contract.LastAddedAt(&_Glpmanager.CallOpts, arg0)
}

// LastAddedAt is a free data retrieval call binding the contract method 0x8b770e11.
//
// Solidity: function lastAddedAt(address ) view returns(uint256)
func (_Glpmanager *GlpmanagerCallerSession) LastAddedAt(arg0 common.Address) (*big.Int, error) {
	return _Glpmanager.Contract.LastAddedAt(&_Glpmanager.CallOpts, arg0)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_Glpmanager *GlpmanagerCaller) Usdg(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Glpmanager.contract.Call(opts, &out, "usdg")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_Glpmanager *GlpmanagerSession) Usdg() (common.Address, error) {
	return _Glpmanager.Contract.Usdg(&_Glpmanager.CallOpts)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_Glpmanager *GlpmanagerCallerSession) Usdg() (common.Address, error) {
	return _Glpmanager.Contract.Usdg(&_Glpmanager.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Glpmanager *GlpmanagerCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Glpmanager.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Glpmanager *GlpmanagerSession) Vault() (common.Address, error) {
	return _Glpmanager.Contract.Vault(&_Glpmanager.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Glpmanager *GlpmanagerCallerSession) Vault() (common.Address, error) {
	return _Glpmanager.Contract.Vault(&_Glpmanager.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1ece366a.
//
// Solidity: function addLiquidity(address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_Glpmanager *GlpmanagerTransactor) AddLiquidity(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _Glpmanager.contract.Transact(opts, "addLiquidity", _token, _amount, _minUsdg, _minGlp)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1ece366a.
//
// Solidity: function addLiquidity(address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_Glpmanager *GlpmanagerSession) AddLiquidity(_token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _Glpmanager.Contract.AddLiquidity(&_Glpmanager.TransactOpts, _token, _amount, _minUsdg, _minGlp)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1ece366a.
//
// Solidity: function addLiquidity(address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_Glpmanager *GlpmanagerTransactorSession) AddLiquidity(_token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _Glpmanager.Contract.AddLiquidity(&_Glpmanager.TransactOpts, _token, _amount, _minUsdg, _minGlp)
}

// AddLiquidityForAccount is a paid mutator transaction binding the contract method 0x17eb2a15.
//
// Solidity: function addLiquidityForAccount(address _fundingAccount, address _account, address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_Glpmanager *GlpmanagerTransactor) AddLiquidityForAccount(opts *bind.TransactOpts, _fundingAccount common.Address, _account common.Address, _token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _Glpmanager.contract.Transact(opts, "addLiquidityForAccount", _fundingAccount, _account, _token, _amount, _minUsdg, _minGlp)
}

// AddLiquidityForAccount is a paid mutator transaction binding the contract method 0x17eb2a15.
//
// Solidity: function addLiquidityForAccount(address _fundingAccount, address _account, address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_Glpmanager *GlpmanagerSession) AddLiquidityForAccount(_fundingAccount common.Address, _account common.Address, _token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _Glpmanager.Contract.AddLiquidityForAccount(&_Glpmanager.TransactOpts, _fundingAccount, _account, _token, _amount, _minUsdg, _minGlp)
}

// AddLiquidityForAccount is a paid mutator transaction binding the contract method 0x17eb2a15.
//
// Solidity: function addLiquidityForAccount(address _fundingAccount, address _account, address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_Glpmanager *GlpmanagerTransactorSession) AddLiquidityForAccount(_fundingAccount common.Address, _account common.Address, _token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _Glpmanager.Contract.AddLiquidityForAccount(&_Glpmanager.TransactOpts, _fundingAccount, _account, _token, _amount, _minUsdg, _minGlp)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x8fed0b2c.
//
// Solidity: function removeLiquidity(address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_Glpmanager *GlpmanagerTransactor) RemoveLiquidity(opts *bind.TransactOpts, _tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Glpmanager.contract.Transact(opts, "removeLiquidity", _tokenOut, _glpAmount, _minOut, _receiver)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x8fed0b2c.
//
// Solidity: function removeLiquidity(address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_Glpmanager *GlpmanagerSession) RemoveLiquidity(_tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Glpmanager.Contract.RemoveLiquidity(&_Glpmanager.TransactOpts, _tokenOut, _glpAmount, _minOut, _receiver)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x8fed0b2c.
//
// Solidity: function removeLiquidity(address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_Glpmanager *GlpmanagerTransactorSession) RemoveLiquidity(_tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Glpmanager.Contract.RemoveLiquidity(&_Glpmanager.TransactOpts, _tokenOut, _glpAmount, _minOut, _receiver)
}

// RemoveLiquidityForAccount is a paid mutator transaction binding the contract method 0x71d597ad.
//
// Solidity: function removeLiquidityForAccount(address _account, address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_Glpmanager *GlpmanagerTransactor) RemoveLiquidityForAccount(opts *bind.TransactOpts, _account common.Address, _tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Glpmanager.contract.Transact(opts, "removeLiquidityForAccount", _account, _tokenOut, _glpAmount, _minOut, _receiver)
}

// RemoveLiquidityForAccount is a paid mutator transaction binding the contract method 0x71d597ad.
//
// Solidity: function removeLiquidityForAccount(address _account, address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_Glpmanager *GlpmanagerSession) RemoveLiquidityForAccount(_account common.Address, _tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Glpmanager.Contract.RemoveLiquidityForAccount(&_Glpmanager.TransactOpts, _account, _tokenOut, _glpAmount, _minOut, _receiver)
}

// RemoveLiquidityForAccount is a paid mutator transaction binding the contract method 0x71d597ad.
//
// Solidity: function removeLiquidityForAccount(address _account, address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_Glpmanager *GlpmanagerTransactorSession) RemoveLiquidityForAccount(_account common.Address, _tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Glpmanager.Contract.RemoveLiquidityForAccount(&_Glpmanager.TransactOpts, _account, _tokenOut, _glpAmount, _minOut, _receiver)
}

// SetAumAdjustment is a paid mutator transaction binding the contract method 0x9116c4ae.
//
// Solidity: function setAumAdjustment(uint256 _aumAddition, uint256 _aumDeduction) returns()
func (_Glpmanager *GlpmanagerTransactor) SetAumAdjustment(opts *bind.TransactOpts, _aumAddition *big.Int, _aumDeduction *big.Int) (*types.Transaction, error) {
	return _Glpmanager.contract.Transact(opts, "setAumAdjustment", _aumAddition, _aumDeduction)
}

// SetAumAdjustment is a paid mutator transaction binding the contract method 0x9116c4ae.
//
// Solidity: function setAumAdjustment(uint256 _aumAddition, uint256 _aumDeduction) returns()
func (_Glpmanager *GlpmanagerSession) SetAumAdjustment(_aumAddition *big.Int, _aumDeduction *big.Int) (*types.Transaction, error) {
	return _Glpmanager.Contract.SetAumAdjustment(&_Glpmanager.TransactOpts, _aumAddition, _aumDeduction)
}

// SetAumAdjustment is a paid mutator transaction binding the contract method 0x9116c4ae.
//
// Solidity: function setAumAdjustment(uint256 _aumAddition, uint256 _aumDeduction) returns()
func (_Glpmanager *GlpmanagerTransactorSession) SetAumAdjustment(_aumAddition *big.Int, _aumDeduction *big.Int) (*types.Transaction, error) {
	return _Glpmanager.Contract.SetAumAdjustment(&_Glpmanager.TransactOpts, _aumAddition, _aumDeduction)
}

// SetCooldownDuration is a paid mutator transaction binding the contract method 0x966be075.
//
// Solidity: function setCooldownDuration(uint256 _cooldownDuration) returns()
func (_Glpmanager *GlpmanagerTransactor) SetCooldownDuration(opts *bind.TransactOpts, _cooldownDuration *big.Int) (*types.Transaction, error) {
	return _Glpmanager.contract.Transact(opts, "setCooldownDuration", _cooldownDuration)
}

// SetCooldownDuration is a paid mutator transaction binding the contract method 0x966be075.
//
// Solidity: function setCooldownDuration(uint256 _cooldownDuration) returns()
func (_Glpmanager *GlpmanagerSession) SetCooldownDuration(_cooldownDuration *big.Int) (*types.Transaction, error) {
	return _Glpmanager.Contract.SetCooldownDuration(&_Glpmanager.TransactOpts, _cooldownDuration)
}

// SetCooldownDuration is a paid mutator transaction binding the contract method 0x966be075.
//
// Solidity: function setCooldownDuration(uint256 _cooldownDuration) returns()
func (_Glpmanager *GlpmanagerTransactorSession) SetCooldownDuration(_cooldownDuration *big.Int) (*types.Transaction, error) {
	return _Glpmanager.Contract.SetCooldownDuration(&_Glpmanager.TransactOpts, _cooldownDuration)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Glpmanager *GlpmanagerTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _Glpmanager.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Glpmanager *GlpmanagerSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _Glpmanager.Contract.SetGov(&_Glpmanager.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Glpmanager *GlpmanagerTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _Glpmanager.Contract.SetGov(&_Glpmanager.TransactOpts, _gov)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_Glpmanager *GlpmanagerTransactor) SetHandler(opts *bind.TransactOpts, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _Glpmanager.contract.Transact(opts, "setHandler", _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_Glpmanager *GlpmanagerSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _Glpmanager.Contract.SetHandler(&_Glpmanager.TransactOpts, _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_Glpmanager *GlpmanagerTransactorSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _Glpmanager.Contract.SetHandler(&_Glpmanager.TransactOpts, _handler, _isActive)
}

// SetInPrivateMode is a paid mutator transaction binding the contract method 0x6a86da19.
//
// Solidity: function setInPrivateMode(bool _inPrivateMode) returns()
func (_Glpmanager *GlpmanagerTransactor) SetInPrivateMode(opts *bind.TransactOpts, _inPrivateMode bool) (*types.Transaction, error) {
	return _Glpmanager.contract.Transact(opts, "setInPrivateMode", _inPrivateMode)
}

// SetInPrivateMode is a paid mutator transaction binding the contract method 0x6a86da19.
//
// Solidity: function setInPrivateMode(bool _inPrivateMode) returns()
func (_Glpmanager *GlpmanagerSession) SetInPrivateMode(_inPrivateMode bool) (*types.Transaction, error) {
	return _Glpmanager.Contract.SetInPrivateMode(&_Glpmanager.TransactOpts, _inPrivateMode)
}

// SetInPrivateMode is a paid mutator transaction binding the contract method 0x6a86da19.
//
// Solidity: function setInPrivateMode(bool _inPrivateMode) returns()
func (_Glpmanager *GlpmanagerTransactorSession) SetInPrivateMode(_inPrivateMode bool) (*types.Transaction, error) {
	return _Glpmanager.Contract.SetInPrivateMode(&_Glpmanager.TransactOpts, _inPrivateMode)
}

// GlpmanagerAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the Glpmanager contract.
type GlpmanagerAddLiquidityIterator struct {
	Event *GlpmanagerAddLiquidity // Event containing the contract specifics and raw log

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
func (it *GlpmanagerAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlpmanagerAddLiquidity)
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
		it.Event = new(GlpmanagerAddLiquidity)
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
func (it *GlpmanagerAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlpmanagerAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlpmanagerAddLiquidity represents a AddLiquidity event raised by the Glpmanager contract.
type GlpmanagerAddLiquidity struct {
	Account    common.Address
	Token      common.Address
	Amount     *big.Int
	AumInUsdg  *big.Int
	GlpSupply  *big.Int
	UsdgAmount *big.Int
	MintAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x38dc38b96482be64113daffd8d464ebda93e856b70ccfc605e69ccf892ab981e.
//
// Solidity: event AddLiquidity(address account, address token, uint256 amount, uint256 aumInUsdg, uint256 glpSupply, uint256 usdgAmount, uint256 mintAmount)
func (_Glpmanager *GlpmanagerFilterer) FilterAddLiquidity(opts *bind.FilterOpts) (*GlpmanagerAddLiquidityIterator, error) {

	logs, sub, err := _Glpmanager.contract.FilterLogs(opts, "AddLiquidity")
	if err != nil {
		return nil, err
	}
	return &GlpmanagerAddLiquidityIterator{contract: _Glpmanager.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x38dc38b96482be64113daffd8d464ebda93e856b70ccfc605e69ccf892ab981e.
//
// Solidity: event AddLiquidity(address account, address token, uint256 amount, uint256 aumInUsdg, uint256 glpSupply, uint256 usdgAmount, uint256 mintAmount)
func (_Glpmanager *GlpmanagerFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *GlpmanagerAddLiquidity) (event.Subscription, error) {

	logs, sub, err := _Glpmanager.contract.WatchLogs(opts, "AddLiquidity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlpmanagerAddLiquidity)
				if err := _Glpmanager.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0x38dc38b96482be64113daffd8d464ebda93e856b70ccfc605e69ccf892ab981e.
//
// Solidity: event AddLiquidity(address account, address token, uint256 amount, uint256 aumInUsdg, uint256 glpSupply, uint256 usdgAmount, uint256 mintAmount)
func (_Glpmanager *GlpmanagerFilterer) ParseAddLiquidity(log types.Log) (*GlpmanagerAddLiquidity, error) {
	event := new(GlpmanagerAddLiquidity)
	if err := _Glpmanager.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlpmanagerRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the Glpmanager contract.
type GlpmanagerRemoveLiquidityIterator struct {
	Event *GlpmanagerRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *GlpmanagerRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlpmanagerRemoveLiquidity)
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
		it.Event = new(GlpmanagerRemoveLiquidity)
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
func (it *GlpmanagerRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlpmanagerRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlpmanagerRemoveLiquidity represents a RemoveLiquidity event raised by the Glpmanager contract.
type GlpmanagerRemoveLiquidity struct {
	Account    common.Address
	Token      common.Address
	GlpAmount  *big.Int
	AumInUsdg  *big.Int
	GlpSupply  *big.Int
	UsdgAmount *big.Int
	AmountOut  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x87b9679bb9a4944bafa98c267e7cd4a00ab29fed48afdefae25f0fca5da27940.
//
// Solidity: event RemoveLiquidity(address account, address token, uint256 glpAmount, uint256 aumInUsdg, uint256 glpSupply, uint256 usdgAmount, uint256 amountOut)
func (_Glpmanager *GlpmanagerFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts) (*GlpmanagerRemoveLiquidityIterator, error) {

	logs, sub, err := _Glpmanager.contract.FilterLogs(opts, "RemoveLiquidity")
	if err != nil {
		return nil, err
	}
	return &GlpmanagerRemoveLiquidityIterator{contract: _Glpmanager.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x87b9679bb9a4944bafa98c267e7cd4a00ab29fed48afdefae25f0fca5da27940.
//
// Solidity: event RemoveLiquidity(address account, address token, uint256 glpAmount, uint256 aumInUsdg, uint256 glpSupply, uint256 usdgAmount, uint256 amountOut)
func (_Glpmanager *GlpmanagerFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *GlpmanagerRemoveLiquidity) (event.Subscription, error) {

	logs, sub, err := _Glpmanager.contract.WatchLogs(opts, "RemoveLiquidity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlpmanagerRemoveLiquidity)
				if err := _Glpmanager.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0x87b9679bb9a4944bafa98c267e7cd4a00ab29fed48afdefae25f0fca5da27940.
//
// Solidity: event RemoveLiquidity(address account, address token, uint256 glpAmount, uint256 aumInUsdg, uint256 glpSupply, uint256 usdgAmount, uint256 amountOut)
func (_Glpmanager *GlpmanagerFilterer) ParseRemoveLiquidity(log types.Log) (*GlpmanagerRemoveLiquidity, error) {
	event := new(GlpmanagerRemoveLiquidity)
	if err := _Glpmanager.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
