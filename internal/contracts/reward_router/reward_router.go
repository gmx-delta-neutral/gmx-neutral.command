// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rewardrouter

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

// RewardrouterMetaData contains all meta data concerning the Rewardrouter contract.
var RewardrouterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeGlp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeGmx\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnstakeGlp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnstakeGmx\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"acceptTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"}],\"name\":\"batchCompoundForAccounts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchStakeGmxForAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bnGmx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimEsGmx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"compound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"compoundForAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"esGmx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGlpTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glpManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glpVester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gmx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gmxVester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_shouldClaimGmx\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_shouldStakeGmx\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_shouldClaimEsGmx\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_shouldStakeEsGmx\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_shouldStakeMultiplierPoints\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_shouldClaimWeth\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_shouldConvertWethToEth\",\"type\":\"bool\"}],\"name\":\"handleRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gmx\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_esGmx\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bnGmx\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_glp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakedGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bonusGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeGlpTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakedGlpTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_glpManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gmxVester\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_glpVester\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minUsdg\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minGlp\",\"type\":\"uint256\"}],\"name\":\"mintAndStakeGlp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minUsdg\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minGlp\",\"type\":\"uint256\"}],\"name\":\"mintAndStakeGlpETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingReceivers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"signalTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stakeEsGmx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stakeGmx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stakeGmxForAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedGlpTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_glpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"unstakeAndRedeemGlp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_glpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minOut\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"unstakeAndRedeemGlpETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unstakeEsGmx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unstakeGmx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RewardrouterABI is the input ABI used to generate the binding from.
// Deprecated: Use RewardrouterMetaData.ABI instead.
var RewardrouterABI = RewardrouterMetaData.ABI

// Rewardrouter is an auto generated Go binding around an Ethereum contract.
type Rewardrouter struct {
	RewardrouterCaller     // Read-only binding to the contract
	RewardrouterTransactor // Write-only binding to the contract
	RewardrouterFilterer   // Log filterer for contract events
}

// RewardrouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardrouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardrouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardrouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardrouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardrouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardrouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardrouterSession struct {
	Contract     *Rewardrouter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewardrouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardrouterCallerSession struct {
	Contract *RewardrouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RewardrouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardrouterTransactorSession struct {
	Contract     *RewardrouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RewardrouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardrouterRaw struct {
	Contract *Rewardrouter // Generic contract binding to access the raw methods on
}

// RewardrouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardrouterCallerRaw struct {
	Contract *RewardrouterCaller // Generic read-only contract binding to access the raw methods on
}

// RewardrouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardrouterTransactorRaw struct {
	Contract *RewardrouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewardrouter creates a new instance of Rewardrouter, bound to a specific deployed contract.
func NewRewardrouter(address common.Address, backend bind.ContractBackend) (*Rewardrouter, error) {
	contract, err := bindRewardrouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rewardrouter{RewardrouterCaller: RewardrouterCaller{contract: contract}, RewardrouterTransactor: RewardrouterTransactor{contract: contract}, RewardrouterFilterer: RewardrouterFilterer{contract: contract}}, nil
}

// NewRewardrouterCaller creates a new read-only instance of Rewardrouter, bound to a specific deployed contract.
func NewRewardrouterCaller(address common.Address, caller bind.ContractCaller) (*RewardrouterCaller, error) {
	contract, err := bindRewardrouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardrouterCaller{contract: contract}, nil
}

// NewRewardrouterTransactor creates a new write-only instance of Rewardrouter, bound to a specific deployed contract.
func NewRewardrouterTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardrouterTransactor, error) {
	contract, err := bindRewardrouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardrouterTransactor{contract: contract}, nil
}

// NewRewardrouterFilterer creates a new log filterer instance of Rewardrouter, bound to a specific deployed contract.
func NewRewardrouterFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardrouterFilterer, error) {
	contract, err := bindRewardrouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardrouterFilterer{contract: contract}, nil
}

// bindRewardrouter binds a generic wrapper to an already deployed contract.
func bindRewardrouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RewardrouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rewardrouter *RewardrouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rewardrouter.Contract.RewardrouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rewardrouter *RewardrouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewardrouter.Contract.RewardrouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rewardrouter *RewardrouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rewardrouter.Contract.RewardrouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rewardrouter *RewardrouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rewardrouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rewardrouter *RewardrouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewardrouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rewardrouter *RewardrouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rewardrouter.Contract.contract.Transact(opts, method, params...)
}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_Rewardrouter *RewardrouterCaller) BnGmx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewardrouter.contract.Call(opts, &out, "bnGmx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_Rewardrouter *RewardrouterSession) BnGmx() (common.Address, error) {
	return _Rewardrouter.Contract.BnGmx(&_Rewardrouter.CallOpts)
}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_Rewardrouter *RewardrouterCallerSession) BnGmx() (common.Address, error) {
	return _Rewardrouter.Contract.BnGmx(&_Rewardrouter.CallOpts)
}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_Rewardrouter *RewardrouterCaller) BonusGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewardrouter.contract.Call(opts, &out, "bonusGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_Rewardrouter *RewardrouterSession) BonusGmxTracker() (common.Address, error) {
	return _Rewardrouter.Contract.BonusGmxTracker(&_Rewardrouter.CallOpts)
}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_Rewardrouter *RewardrouterCallerSession) BonusGmxTracker() (common.Address, error) {
	return _Rewardrouter.Contract.BonusGmxTracker(&_Rewardrouter.CallOpts)
}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_Rewardrouter *RewardrouterCaller) EsGmx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewardrouter.contract.Call(opts, &out, "esGmx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_Rewardrouter *RewardrouterSession) EsGmx() (common.Address, error) {
	return _Rewardrouter.Contract.EsGmx(&_Rewardrouter.CallOpts)
}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_Rewardrouter *RewardrouterCallerSession) EsGmx() (common.Address, error) {
	return _Rewardrouter.Contract.EsGmx(&_Rewardrouter.CallOpts)
}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_Rewardrouter *RewardrouterCaller) FeeGlpTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewardrouter.contract.Call(opts, &out, "feeGlpTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_Rewardrouter *RewardrouterSession) FeeGlpTracker() (common.Address, error) {
	return _Rewardrouter.Contract.FeeGlpTracker(&_Rewardrouter.CallOpts)
}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_Rewardrouter *RewardrouterCallerSession) FeeGlpTracker() (common.Address, error) {
	return _Rewardrouter.Contract.FeeGlpTracker(&_Rewardrouter.CallOpts)
}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_Rewardrouter *RewardrouterCaller) FeeGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewardrouter.contract.Call(opts, &out, "feeGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_Rewardrouter *RewardrouterSession) FeeGmxTracker() (common.Address, error) {
	return _Rewardrouter.Contract.FeeGmxTracker(&_Rewardrouter.CallOpts)
}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_Rewardrouter *RewardrouterCallerSession) FeeGmxTracker() (common.Address, error) {
	return _Rewardrouter.Contract.FeeGmxTracker(&_Rewardrouter.CallOpts)
}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_Rewardrouter *RewardrouterCaller) Glp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewardrouter.contract.Call(opts, &out, "glp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_Rewardrouter *RewardrouterSession) Glp() (common.Address, error) {
	return _Rewardrouter.Contract.Glp(&_Rewardrouter.CallOpts)
}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_Rewardrouter *RewardrouterCallerSession) Glp() (common.Address, error) {
	return _Rewardrouter.Contract.Glp(&_Rewardrouter.CallOpts)
}

// GlpManager is a free data retrieval call binding the contract method 0xfa6db1bc.
//
// Solidity: function glpManager() view returns(address)
func (_Rewardrouter *RewardrouterCaller) GlpManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewardrouter.contract.Call(opts, &out, "glpManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlpManager is a free data retrieval call binding the contract method 0xfa6db1bc.
//
// Solidity: function glpManager() view returns(address)
func (_Rewardrouter *RewardrouterSession) GlpManager() (common.Address, error) {
	return _Rewardrouter.Contract.GlpManager(&_Rewardrouter.CallOpts)
}

// GlpManager is a free data retrieval call binding the contract method 0xfa6db1bc.
//
// Solidity: function glpManager() view returns(address)
func (_Rewardrouter *RewardrouterCallerSession) GlpManager() (common.Address, error) {
	return _Rewardrouter.Contract.GlpManager(&_Rewardrouter.CallOpts)
}

// GlpVester is a free data retrieval call binding the contract method 0x3671df25.
//
// Solidity: function glpVester() view returns(address)
func (_Rewardrouter *RewardrouterCaller) GlpVester(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewardrouter.contract.Call(opts, &out, "glpVester")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlpVester is a free data retrieval call binding the contract method 0x3671df25.
//
// Solidity: function glpVester() view returns(address)
func (_Rewardrouter *RewardrouterSession) GlpVester() (common.Address, error) {
	return _Rewardrouter.Contract.GlpVester(&_Rewardrouter.CallOpts)
}

// GlpVester is a free data retrieval call binding the contract method 0x3671df25.
//
// Solidity: function glpVester() view returns(address)
func (_Rewardrouter *RewardrouterCallerSession) GlpVester() (common.Address, error) {
	return _Rewardrouter.Contract.GlpVester(&_Rewardrouter.CallOpts)
}

// Gmx is a free data retrieval call binding the contract method 0x31e67c71.
//
// Solidity: function gmx() view returns(address)
func (_Rewardrouter *RewardrouterCaller) Gmx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewardrouter.contract.Call(opts, &out, "gmx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gmx is a free data retrieval call binding the contract method 0x31e67c71.
//
// Solidity: function gmx() view returns(address)
func (_Rewardrouter *RewardrouterSession) Gmx() (common.Address, error) {
	return _Rewardrouter.Contract.Gmx(&_Rewardrouter.CallOpts)
}

// Gmx is a free data retrieval call binding the contract method 0x31e67c71.
//
// Solidity: function gmx() view returns(address)
func (_Rewardrouter *RewardrouterCallerSession) Gmx() (common.Address, error) {
	return _Rewardrouter.Contract.Gmx(&_Rewardrouter.CallOpts)
}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_Rewardrouter *RewardrouterCaller) GmxVester(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewardrouter.contract.Call(opts, &out, "gmxVester")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_Rewardrouter *RewardrouterSession) GmxVester() (common.Address, error) {
	return _Rewardrouter.Contract.GmxVester(&_Rewardrouter.CallOpts)
}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_Rewardrouter *RewardrouterCallerSession) GmxVester() (common.Address, error) {
	return _Rewardrouter.Contract.GmxVester(&_Rewardrouter.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Rewardrouter *RewardrouterCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewardrouter.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Rewardrouter *RewardrouterSession) Gov() (common.Address, error) {
	return _Rewardrouter.Contract.Gov(&_Rewardrouter.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Rewardrouter *RewardrouterCallerSession) Gov() (common.Address, error) {
	return _Rewardrouter.Contract.Gov(&_Rewardrouter.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Rewardrouter *RewardrouterCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Rewardrouter.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Rewardrouter *RewardrouterSession) IsInitialized() (bool, error) {
	return _Rewardrouter.Contract.IsInitialized(&_Rewardrouter.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Rewardrouter *RewardrouterCallerSession) IsInitialized() (bool, error) {
	return _Rewardrouter.Contract.IsInitialized(&_Rewardrouter.CallOpts)
}

// PendingReceivers is a free data retrieval call binding the contract method 0xe1b9db89.
//
// Solidity: function pendingReceivers(address ) view returns(address)
func (_Rewardrouter *RewardrouterCaller) PendingReceivers(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Rewardrouter.contract.Call(opts, &out, "pendingReceivers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingReceivers is a free data retrieval call binding the contract method 0xe1b9db89.
//
// Solidity: function pendingReceivers(address ) view returns(address)
func (_Rewardrouter *RewardrouterSession) PendingReceivers(arg0 common.Address) (common.Address, error) {
	return _Rewardrouter.Contract.PendingReceivers(&_Rewardrouter.CallOpts, arg0)
}

// PendingReceivers is a free data retrieval call binding the contract method 0xe1b9db89.
//
// Solidity: function pendingReceivers(address ) view returns(address)
func (_Rewardrouter *RewardrouterCallerSession) PendingReceivers(arg0 common.Address) (common.Address, error) {
	return _Rewardrouter.Contract.PendingReceivers(&_Rewardrouter.CallOpts, arg0)
}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_Rewardrouter *RewardrouterCaller) StakedGlpTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewardrouter.contract.Call(opts, &out, "stakedGlpTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_Rewardrouter *RewardrouterSession) StakedGlpTracker() (common.Address, error) {
	return _Rewardrouter.Contract.StakedGlpTracker(&_Rewardrouter.CallOpts)
}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_Rewardrouter *RewardrouterCallerSession) StakedGlpTracker() (common.Address, error) {
	return _Rewardrouter.Contract.StakedGlpTracker(&_Rewardrouter.CallOpts)
}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_Rewardrouter *RewardrouterCaller) StakedGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewardrouter.contract.Call(opts, &out, "stakedGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_Rewardrouter *RewardrouterSession) StakedGmxTracker() (common.Address, error) {
	return _Rewardrouter.Contract.StakedGmxTracker(&_Rewardrouter.CallOpts)
}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_Rewardrouter *RewardrouterCallerSession) StakedGmxTracker() (common.Address, error) {
	return _Rewardrouter.Contract.StakedGmxTracker(&_Rewardrouter.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Rewardrouter *RewardrouterCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewardrouter.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Rewardrouter *RewardrouterSession) Weth() (common.Address, error) {
	return _Rewardrouter.Contract.Weth(&_Rewardrouter.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Rewardrouter *RewardrouterCallerSession) Weth() (common.Address, error) {
	return _Rewardrouter.Contract.Weth(&_Rewardrouter.CallOpts)
}

// AcceptTransfer is a paid mutator transaction binding the contract method 0x655603a4.
//
// Solidity: function acceptTransfer(address _sender) returns()
func (_Rewardrouter *RewardrouterTransactor) AcceptTransfer(opts *bind.TransactOpts, _sender common.Address) (*types.Transaction, error) {
	return _Rewardrouter.contract.Transact(opts, "acceptTransfer", _sender)
}

// AcceptTransfer is a paid mutator transaction binding the contract method 0x655603a4.
//
// Solidity: function acceptTransfer(address _sender) returns()
func (_Rewardrouter *RewardrouterSession) AcceptTransfer(_sender common.Address) (*types.Transaction, error) {
	return _Rewardrouter.Contract.AcceptTransfer(&_Rewardrouter.TransactOpts, _sender)
}

// AcceptTransfer is a paid mutator transaction binding the contract method 0x655603a4.
//
// Solidity: function acceptTransfer(address _sender) returns()
func (_Rewardrouter *RewardrouterTransactorSession) AcceptTransfer(_sender common.Address) (*types.Transaction, error) {
	return _Rewardrouter.Contract.AcceptTransfer(&_Rewardrouter.TransactOpts, _sender)
}

// BatchCompoundForAccounts is a paid mutator transaction binding the contract method 0x1af276a6.
//
// Solidity: function batchCompoundForAccounts(address[] _accounts) returns()
func (_Rewardrouter *RewardrouterTransactor) BatchCompoundForAccounts(opts *bind.TransactOpts, _accounts []common.Address) (*types.Transaction, error) {
	return _Rewardrouter.contract.Transact(opts, "batchCompoundForAccounts", _accounts)
}

// BatchCompoundForAccounts is a paid mutator transaction binding the contract method 0x1af276a6.
//
// Solidity: function batchCompoundForAccounts(address[] _accounts) returns()
func (_Rewardrouter *RewardrouterSession) BatchCompoundForAccounts(_accounts []common.Address) (*types.Transaction, error) {
	return _Rewardrouter.Contract.BatchCompoundForAccounts(&_Rewardrouter.TransactOpts, _accounts)
}

// BatchCompoundForAccounts is a paid mutator transaction binding the contract method 0x1af276a6.
//
// Solidity: function batchCompoundForAccounts(address[] _accounts) returns()
func (_Rewardrouter *RewardrouterTransactorSession) BatchCompoundForAccounts(_accounts []common.Address) (*types.Transaction, error) {
	return _Rewardrouter.Contract.BatchCompoundForAccounts(&_Rewardrouter.TransactOpts, _accounts)
}

// BatchStakeGmxForAccount is a paid mutator transaction binding the contract method 0x0db79e52.
//
// Solidity: function batchStakeGmxForAccount(address[] _accounts, uint256[] _amounts) returns()
func (_Rewardrouter *RewardrouterTransactor) BatchStakeGmxForAccount(opts *bind.TransactOpts, _accounts []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Rewardrouter.contract.Transact(opts, "batchStakeGmxForAccount", _accounts, _amounts)
}

// BatchStakeGmxForAccount is a paid mutator transaction binding the contract method 0x0db79e52.
//
// Solidity: function batchStakeGmxForAccount(address[] _accounts, uint256[] _amounts) returns()
func (_Rewardrouter *RewardrouterSession) BatchStakeGmxForAccount(_accounts []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Rewardrouter.Contract.BatchStakeGmxForAccount(&_Rewardrouter.TransactOpts, _accounts, _amounts)
}

// BatchStakeGmxForAccount is a paid mutator transaction binding the contract method 0x0db79e52.
//
// Solidity: function batchStakeGmxForAccount(address[] _accounts, uint256[] _amounts) returns()
func (_Rewardrouter *RewardrouterTransactorSession) BatchStakeGmxForAccount(_accounts []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Rewardrouter.Contract.BatchStakeGmxForAccount(&_Rewardrouter.TransactOpts, _accounts, _amounts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Rewardrouter *RewardrouterTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewardrouter.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Rewardrouter *RewardrouterSession) Claim() (*types.Transaction, error) {
	return _Rewardrouter.Contract.Claim(&_Rewardrouter.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Rewardrouter *RewardrouterTransactorSession) Claim() (*types.Transaction, error) {
	return _Rewardrouter.Contract.Claim(&_Rewardrouter.TransactOpts)
}

// ClaimEsGmx is a paid mutator transaction binding the contract method 0x5fe3945f.
//
// Solidity: function claimEsGmx() returns()
func (_Rewardrouter *RewardrouterTransactor) ClaimEsGmx(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewardrouter.contract.Transact(opts, "claimEsGmx")
}

// ClaimEsGmx is a paid mutator transaction binding the contract method 0x5fe3945f.
//
// Solidity: function claimEsGmx() returns()
func (_Rewardrouter *RewardrouterSession) ClaimEsGmx() (*types.Transaction, error) {
	return _Rewardrouter.Contract.ClaimEsGmx(&_Rewardrouter.TransactOpts)
}

// ClaimEsGmx is a paid mutator transaction binding the contract method 0x5fe3945f.
//
// Solidity: function claimEsGmx() returns()
func (_Rewardrouter *RewardrouterTransactorSession) ClaimEsGmx() (*types.Transaction, error) {
	return _Rewardrouter.Contract.ClaimEsGmx(&_Rewardrouter.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_Rewardrouter *RewardrouterTransactor) ClaimFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewardrouter.contract.Transact(opts, "claimFees")
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_Rewardrouter *RewardrouterSession) ClaimFees() (*types.Transaction, error) {
	return _Rewardrouter.Contract.ClaimFees(&_Rewardrouter.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_Rewardrouter *RewardrouterTransactorSession) ClaimFees() (*types.Transaction, error) {
	return _Rewardrouter.Contract.ClaimFees(&_Rewardrouter.TransactOpts)
}

// Compound is a paid mutator transaction binding the contract method 0xf69e2046.
//
// Solidity: function compound() returns()
func (_Rewardrouter *RewardrouterTransactor) Compound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewardrouter.contract.Transact(opts, "compound")
}

// Compound is a paid mutator transaction binding the contract method 0xf69e2046.
//
// Solidity: function compound() returns()
func (_Rewardrouter *RewardrouterSession) Compound() (*types.Transaction, error) {
	return _Rewardrouter.Contract.Compound(&_Rewardrouter.TransactOpts)
}

// Compound is a paid mutator transaction binding the contract method 0xf69e2046.
//
// Solidity: function compound() returns()
func (_Rewardrouter *RewardrouterTransactorSession) Compound() (*types.Transaction, error) {
	return _Rewardrouter.Contract.Compound(&_Rewardrouter.TransactOpts)
}

// CompoundForAccount is a paid mutator transaction binding the contract method 0x2a9f4083.
//
// Solidity: function compoundForAccount(address _account) returns()
func (_Rewardrouter *RewardrouterTransactor) CompoundForAccount(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Rewardrouter.contract.Transact(opts, "compoundForAccount", _account)
}

// CompoundForAccount is a paid mutator transaction binding the contract method 0x2a9f4083.
//
// Solidity: function compoundForAccount(address _account) returns()
func (_Rewardrouter *RewardrouterSession) CompoundForAccount(_account common.Address) (*types.Transaction, error) {
	return _Rewardrouter.Contract.CompoundForAccount(&_Rewardrouter.TransactOpts, _account)
}

// CompoundForAccount is a paid mutator transaction binding the contract method 0x2a9f4083.
//
// Solidity: function compoundForAccount(address _account) returns()
func (_Rewardrouter *RewardrouterTransactorSession) CompoundForAccount(_account common.Address) (*types.Transaction, error) {
	return _Rewardrouter.Contract.CompoundForAccount(&_Rewardrouter.TransactOpts, _account)
}

// HandleRewards is a paid mutator transaction binding the contract method 0x30b70002.
//
// Solidity: function handleRewards(bool _shouldClaimGmx, bool _shouldStakeGmx, bool _shouldClaimEsGmx, bool _shouldStakeEsGmx, bool _shouldStakeMultiplierPoints, bool _shouldClaimWeth, bool _shouldConvertWethToEth) returns()
func (_Rewardrouter *RewardrouterTransactor) HandleRewards(opts *bind.TransactOpts, _shouldClaimGmx bool, _shouldStakeGmx bool, _shouldClaimEsGmx bool, _shouldStakeEsGmx bool, _shouldStakeMultiplierPoints bool, _shouldClaimWeth bool, _shouldConvertWethToEth bool) (*types.Transaction, error) {
	return _Rewardrouter.contract.Transact(opts, "handleRewards", _shouldClaimGmx, _shouldStakeGmx, _shouldClaimEsGmx, _shouldStakeEsGmx, _shouldStakeMultiplierPoints, _shouldClaimWeth, _shouldConvertWethToEth)
}

// HandleRewards is a paid mutator transaction binding the contract method 0x30b70002.
//
// Solidity: function handleRewards(bool _shouldClaimGmx, bool _shouldStakeGmx, bool _shouldClaimEsGmx, bool _shouldStakeEsGmx, bool _shouldStakeMultiplierPoints, bool _shouldClaimWeth, bool _shouldConvertWethToEth) returns()
func (_Rewardrouter *RewardrouterSession) HandleRewards(_shouldClaimGmx bool, _shouldStakeGmx bool, _shouldClaimEsGmx bool, _shouldStakeEsGmx bool, _shouldStakeMultiplierPoints bool, _shouldClaimWeth bool, _shouldConvertWethToEth bool) (*types.Transaction, error) {
	return _Rewardrouter.Contract.HandleRewards(&_Rewardrouter.TransactOpts, _shouldClaimGmx, _shouldStakeGmx, _shouldClaimEsGmx, _shouldStakeEsGmx, _shouldStakeMultiplierPoints, _shouldClaimWeth, _shouldConvertWethToEth)
}

// HandleRewards is a paid mutator transaction binding the contract method 0x30b70002.
//
// Solidity: function handleRewards(bool _shouldClaimGmx, bool _shouldStakeGmx, bool _shouldClaimEsGmx, bool _shouldStakeEsGmx, bool _shouldStakeMultiplierPoints, bool _shouldClaimWeth, bool _shouldConvertWethToEth) returns()
func (_Rewardrouter *RewardrouterTransactorSession) HandleRewards(_shouldClaimGmx bool, _shouldStakeGmx bool, _shouldClaimEsGmx bool, _shouldStakeEsGmx bool, _shouldStakeMultiplierPoints bool, _shouldClaimWeth bool, _shouldConvertWethToEth bool) (*types.Transaction, error) {
	return _Rewardrouter.Contract.HandleRewards(&_Rewardrouter.TransactOpts, _shouldClaimGmx, _shouldStakeGmx, _shouldClaimEsGmx, _shouldStakeEsGmx, _shouldStakeMultiplierPoints, _shouldClaimWeth, _shouldConvertWethToEth)
}

// Initialize is a paid mutator transaction binding the contract method 0x2fdd983d.
//
// Solidity: function initialize(address _weth, address _gmx, address _esGmx, address _bnGmx, address _glp, address _stakedGmxTracker, address _bonusGmxTracker, address _feeGmxTracker, address _feeGlpTracker, address _stakedGlpTracker, address _glpManager, address _gmxVester, address _glpVester) returns()
func (_Rewardrouter *RewardrouterTransactor) Initialize(opts *bind.TransactOpts, _weth common.Address, _gmx common.Address, _esGmx common.Address, _bnGmx common.Address, _glp common.Address, _stakedGmxTracker common.Address, _bonusGmxTracker common.Address, _feeGmxTracker common.Address, _feeGlpTracker common.Address, _stakedGlpTracker common.Address, _glpManager common.Address, _gmxVester common.Address, _glpVester common.Address) (*types.Transaction, error) {
	return _Rewardrouter.contract.Transact(opts, "initialize", _weth, _gmx, _esGmx, _bnGmx, _glp, _stakedGmxTracker, _bonusGmxTracker, _feeGmxTracker, _feeGlpTracker, _stakedGlpTracker, _glpManager, _gmxVester, _glpVester)
}

// Initialize is a paid mutator transaction binding the contract method 0x2fdd983d.
//
// Solidity: function initialize(address _weth, address _gmx, address _esGmx, address _bnGmx, address _glp, address _stakedGmxTracker, address _bonusGmxTracker, address _feeGmxTracker, address _feeGlpTracker, address _stakedGlpTracker, address _glpManager, address _gmxVester, address _glpVester) returns()
func (_Rewardrouter *RewardrouterSession) Initialize(_weth common.Address, _gmx common.Address, _esGmx common.Address, _bnGmx common.Address, _glp common.Address, _stakedGmxTracker common.Address, _bonusGmxTracker common.Address, _feeGmxTracker common.Address, _feeGlpTracker common.Address, _stakedGlpTracker common.Address, _glpManager common.Address, _gmxVester common.Address, _glpVester common.Address) (*types.Transaction, error) {
	return _Rewardrouter.Contract.Initialize(&_Rewardrouter.TransactOpts, _weth, _gmx, _esGmx, _bnGmx, _glp, _stakedGmxTracker, _bonusGmxTracker, _feeGmxTracker, _feeGlpTracker, _stakedGlpTracker, _glpManager, _gmxVester, _glpVester)
}

// Initialize is a paid mutator transaction binding the contract method 0x2fdd983d.
//
// Solidity: function initialize(address _weth, address _gmx, address _esGmx, address _bnGmx, address _glp, address _stakedGmxTracker, address _bonusGmxTracker, address _feeGmxTracker, address _feeGlpTracker, address _stakedGlpTracker, address _glpManager, address _gmxVester, address _glpVester) returns()
func (_Rewardrouter *RewardrouterTransactorSession) Initialize(_weth common.Address, _gmx common.Address, _esGmx common.Address, _bnGmx common.Address, _glp common.Address, _stakedGmxTracker common.Address, _bonusGmxTracker common.Address, _feeGmxTracker common.Address, _feeGlpTracker common.Address, _stakedGlpTracker common.Address, _glpManager common.Address, _gmxVester common.Address, _glpVester common.Address) (*types.Transaction, error) {
	return _Rewardrouter.Contract.Initialize(&_Rewardrouter.TransactOpts, _weth, _gmx, _esGmx, _bnGmx, _glp, _stakedGmxTracker, _bonusGmxTracker, _feeGmxTracker, _feeGlpTracker, _stakedGlpTracker, _glpManager, _gmxVester, _glpVester)
}

// MintAndStakeGlp is a paid mutator transaction binding the contract method 0x364e2311.
//
// Solidity: function mintAndStakeGlp(address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_Rewardrouter *RewardrouterTransactor) MintAndStakeGlp(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _Rewardrouter.contract.Transact(opts, "mintAndStakeGlp", _token, _amount, _minUsdg, _minGlp)
}

// MintAndStakeGlp is a paid mutator transaction binding the contract method 0x364e2311.
//
// Solidity: function mintAndStakeGlp(address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_Rewardrouter *RewardrouterSession) MintAndStakeGlp(_token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _Rewardrouter.Contract.MintAndStakeGlp(&_Rewardrouter.TransactOpts, _token, _amount, _minUsdg, _minGlp)
}

// MintAndStakeGlp is a paid mutator transaction binding the contract method 0x364e2311.
//
// Solidity: function mintAndStakeGlp(address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_Rewardrouter *RewardrouterTransactorSession) MintAndStakeGlp(_token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _Rewardrouter.Contract.MintAndStakeGlp(&_Rewardrouter.TransactOpts, _token, _amount, _minUsdg, _minGlp)
}

// MintAndStakeGlpETH is a paid mutator transaction binding the contract method 0x53a8aa03.
//
// Solidity: function mintAndStakeGlpETH(uint256 _minUsdg, uint256 _minGlp) payable returns(uint256)
func (_Rewardrouter *RewardrouterTransactor) MintAndStakeGlpETH(opts *bind.TransactOpts, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _Rewardrouter.contract.Transact(opts, "mintAndStakeGlpETH", _minUsdg, _minGlp)
}

// MintAndStakeGlpETH is a paid mutator transaction binding the contract method 0x53a8aa03.
//
// Solidity: function mintAndStakeGlpETH(uint256 _minUsdg, uint256 _minGlp) payable returns(uint256)
func (_Rewardrouter *RewardrouterSession) MintAndStakeGlpETH(_minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _Rewardrouter.Contract.MintAndStakeGlpETH(&_Rewardrouter.TransactOpts, _minUsdg, _minGlp)
}

// MintAndStakeGlpETH is a paid mutator transaction binding the contract method 0x53a8aa03.
//
// Solidity: function mintAndStakeGlpETH(uint256 _minUsdg, uint256 _minGlp) payable returns(uint256)
func (_Rewardrouter *RewardrouterTransactorSession) MintAndStakeGlpETH(_minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _Rewardrouter.Contract.MintAndStakeGlpETH(&_Rewardrouter.TransactOpts, _minUsdg, _minGlp)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Rewardrouter *RewardrouterTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _Rewardrouter.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Rewardrouter *RewardrouterSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _Rewardrouter.Contract.SetGov(&_Rewardrouter.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Rewardrouter *RewardrouterTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _Rewardrouter.Contract.SetGov(&_Rewardrouter.TransactOpts, _gov)
}

// SignalTransfer is a paid mutator transaction binding the contract method 0xef9aacfd.
//
// Solidity: function signalTransfer(address _receiver) returns()
func (_Rewardrouter *RewardrouterTransactor) SignalTransfer(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _Rewardrouter.contract.Transact(opts, "signalTransfer", _receiver)
}

// SignalTransfer is a paid mutator transaction binding the contract method 0xef9aacfd.
//
// Solidity: function signalTransfer(address _receiver) returns()
func (_Rewardrouter *RewardrouterSession) SignalTransfer(_receiver common.Address) (*types.Transaction, error) {
	return _Rewardrouter.Contract.SignalTransfer(&_Rewardrouter.TransactOpts, _receiver)
}

// SignalTransfer is a paid mutator transaction binding the contract method 0xef9aacfd.
//
// Solidity: function signalTransfer(address _receiver) returns()
func (_Rewardrouter *RewardrouterTransactorSession) SignalTransfer(_receiver common.Address) (*types.Transaction, error) {
	return _Rewardrouter.Contract.SignalTransfer(&_Rewardrouter.TransactOpts, _receiver)
}

// StakeEsGmx is a paid mutator transaction binding the contract method 0xef8c5994.
//
// Solidity: function stakeEsGmx(uint256 _amount) returns()
func (_Rewardrouter *RewardrouterTransactor) StakeEsGmx(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Rewardrouter.contract.Transact(opts, "stakeEsGmx", _amount)
}

// StakeEsGmx is a paid mutator transaction binding the contract method 0xef8c5994.
//
// Solidity: function stakeEsGmx(uint256 _amount) returns()
func (_Rewardrouter *RewardrouterSession) StakeEsGmx(_amount *big.Int) (*types.Transaction, error) {
	return _Rewardrouter.Contract.StakeEsGmx(&_Rewardrouter.TransactOpts, _amount)
}

// StakeEsGmx is a paid mutator transaction binding the contract method 0xef8c5994.
//
// Solidity: function stakeEsGmx(uint256 _amount) returns()
func (_Rewardrouter *RewardrouterTransactorSession) StakeEsGmx(_amount *big.Int) (*types.Transaction, error) {
	return _Rewardrouter.Contract.StakeEsGmx(&_Rewardrouter.TransactOpts, _amount)
}

// StakeGmx is a paid mutator transaction binding the contract method 0xf3daeacc.
//
// Solidity: function stakeGmx(uint256 _amount) returns()
func (_Rewardrouter *RewardrouterTransactor) StakeGmx(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Rewardrouter.contract.Transact(opts, "stakeGmx", _amount)
}

// StakeGmx is a paid mutator transaction binding the contract method 0xf3daeacc.
//
// Solidity: function stakeGmx(uint256 _amount) returns()
func (_Rewardrouter *RewardrouterSession) StakeGmx(_amount *big.Int) (*types.Transaction, error) {
	return _Rewardrouter.Contract.StakeGmx(&_Rewardrouter.TransactOpts, _amount)
}

// StakeGmx is a paid mutator transaction binding the contract method 0xf3daeacc.
//
// Solidity: function stakeGmx(uint256 _amount) returns()
func (_Rewardrouter *RewardrouterTransactorSession) StakeGmx(_amount *big.Int) (*types.Transaction, error) {
	return _Rewardrouter.Contract.StakeGmx(&_Rewardrouter.TransactOpts, _amount)
}

// StakeGmxForAccount is a paid mutator transaction binding the contract method 0x5da4b8dd.
//
// Solidity: function stakeGmxForAccount(address _account, uint256 _amount) returns()
func (_Rewardrouter *RewardrouterTransactor) StakeGmxForAccount(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Rewardrouter.contract.Transact(opts, "stakeGmxForAccount", _account, _amount)
}

// StakeGmxForAccount is a paid mutator transaction binding the contract method 0x5da4b8dd.
//
// Solidity: function stakeGmxForAccount(address _account, uint256 _amount) returns()
func (_Rewardrouter *RewardrouterSession) StakeGmxForAccount(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Rewardrouter.Contract.StakeGmxForAccount(&_Rewardrouter.TransactOpts, _account, _amount)
}

// StakeGmxForAccount is a paid mutator transaction binding the contract method 0x5da4b8dd.
//
// Solidity: function stakeGmxForAccount(address _account, uint256 _amount) returns()
func (_Rewardrouter *RewardrouterTransactorSession) StakeGmxForAccount(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Rewardrouter.Contract.StakeGmxForAccount(&_Rewardrouter.TransactOpts, _account, _amount)
}

// UnstakeAndRedeemGlp is a paid mutator transaction binding the contract method 0x0f3aa554.
//
// Solidity: function unstakeAndRedeemGlp(address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_Rewardrouter *RewardrouterTransactor) UnstakeAndRedeemGlp(opts *bind.TransactOpts, _tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Rewardrouter.contract.Transact(opts, "unstakeAndRedeemGlp", _tokenOut, _glpAmount, _minOut, _receiver)
}

// UnstakeAndRedeemGlp is a paid mutator transaction binding the contract method 0x0f3aa554.
//
// Solidity: function unstakeAndRedeemGlp(address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_Rewardrouter *RewardrouterSession) UnstakeAndRedeemGlp(_tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Rewardrouter.Contract.UnstakeAndRedeemGlp(&_Rewardrouter.TransactOpts, _tokenOut, _glpAmount, _minOut, _receiver)
}

// UnstakeAndRedeemGlp is a paid mutator transaction binding the contract method 0x0f3aa554.
//
// Solidity: function unstakeAndRedeemGlp(address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_Rewardrouter *RewardrouterTransactorSession) UnstakeAndRedeemGlp(_tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Rewardrouter.Contract.UnstakeAndRedeemGlp(&_Rewardrouter.TransactOpts, _tokenOut, _glpAmount, _minOut, _receiver)
}

// UnstakeAndRedeemGlpETH is a paid mutator transaction binding the contract method 0xabb5e5e2.
//
// Solidity: function unstakeAndRedeemGlpETH(uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_Rewardrouter *RewardrouterTransactor) UnstakeAndRedeemGlpETH(opts *bind.TransactOpts, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Rewardrouter.contract.Transact(opts, "unstakeAndRedeemGlpETH", _glpAmount, _minOut, _receiver)
}

// UnstakeAndRedeemGlpETH is a paid mutator transaction binding the contract method 0xabb5e5e2.
//
// Solidity: function unstakeAndRedeemGlpETH(uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_Rewardrouter *RewardrouterSession) UnstakeAndRedeemGlpETH(_glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Rewardrouter.Contract.UnstakeAndRedeemGlpETH(&_Rewardrouter.TransactOpts, _glpAmount, _minOut, _receiver)
}

// UnstakeAndRedeemGlpETH is a paid mutator transaction binding the contract method 0xabb5e5e2.
//
// Solidity: function unstakeAndRedeemGlpETH(uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_Rewardrouter *RewardrouterTransactorSession) UnstakeAndRedeemGlpETH(_glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Rewardrouter.Contract.UnstakeAndRedeemGlpETH(&_Rewardrouter.TransactOpts, _glpAmount, _minOut, _receiver)
}

// UnstakeEsGmx is a paid mutator transaction binding the contract method 0x64f64467.
//
// Solidity: function unstakeEsGmx(uint256 _amount) returns()
func (_Rewardrouter *RewardrouterTransactor) UnstakeEsGmx(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Rewardrouter.contract.Transact(opts, "unstakeEsGmx", _amount)
}

// UnstakeEsGmx is a paid mutator transaction binding the contract method 0x64f64467.
//
// Solidity: function unstakeEsGmx(uint256 _amount) returns()
func (_Rewardrouter *RewardrouterSession) UnstakeEsGmx(_amount *big.Int) (*types.Transaction, error) {
	return _Rewardrouter.Contract.UnstakeEsGmx(&_Rewardrouter.TransactOpts, _amount)
}

// UnstakeEsGmx is a paid mutator transaction binding the contract method 0x64f64467.
//
// Solidity: function unstakeEsGmx(uint256 _amount) returns()
func (_Rewardrouter *RewardrouterTransactorSession) UnstakeEsGmx(_amount *big.Int) (*types.Transaction, error) {
	return _Rewardrouter.Contract.UnstakeEsGmx(&_Rewardrouter.TransactOpts, _amount)
}

// UnstakeGmx is a paid mutator transaction binding the contract method 0x078580d2.
//
// Solidity: function unstakeGmx(uint256 _amount) returns()
func (_Rewardrouter *RewardrouterTransactor) UnstakeGmx(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Rewardrouter.contract.Transact(opts, "unstakeGmx", _amount)
}

// UnstakeGmx is a paid mutator transaction binding the contract method 0x078580d2.
//
// Solidity: function unstakeGmx(uint256 _amount) returns()
func (_Rewardrouter *RewardrouterSession) UnstakeGmx(_amount *big.Int) (*types.Transaction, error) {
	return _Rewardrouter.Contract.UnstakeGmx(&_Rewardrouter.TransactOpts, _amount)
}

// UnstakeGmx is a paid mutator transaction binding the contract method 0x078580d2.
//
// Solidity: function unstakeGmx(uint256 _amount) returns()
func (_Rewardrouter *RewardrouterTransactorSession) UnstakeGmx(_amount *big.Int) (*types.Transaction, error) {
	return _Rewardrouter.Contract.UnstakeGmx(&_Rewardrouter.TransactOpts, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_Rewardrouter *RewardrouterTransactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Rewardrouter.contract.Transact(opts, "withdrawToken", _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_Rewardrouter *RewardrouterSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Rewardrouter.Contract.WithdrawToken(&_Rewardrouter.TransactOpts, _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_Rewardrouter *RewardrouterTransactorSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Rewardrouter.Contract.WithdrawToken(&_Rewardrouter.TransactOpts, _token, _account, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Rewardrouter *RewardrouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewardrouter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Rewardrouter *RewardrouterSession) Receive() (*types.Transaction, error) {
	return _Rewardrouter.Contract.Receive(&_Rewardrouter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Rewardrouter *RewardrouterTransactorSession) Receive() (*types.Transaction, error) {
	return _Rewardrouter.Contract.Receive(&_Rewardrouter.TransactOpts)
}

// RewardrouterStakeGlpIterator is returned from FilterStakeGlp and is used to iterate over the raw logs and unpacked data for StakeGlp events raised by the Rewardrouter contract.
type RewardrouterStakeGlpIterator struct {
	Event *RewardrouterStakeGlp // Event containing the contract specifics and raw log

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
func (it *RewardrouterStakeGlpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardrouterStakeGlp)
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
		it.Event = new(RewardrouterStakeGlp)
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
func (it *RewardrouterStakeGlpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardrouterStakeGlpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardrouterStakeGlp represents a StakeGlp event raised by the Rewardrouter contract.
type RewardrouterStakeGlp struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeGlp is a free log retrieval operation binding the contract event 0xa4725d47fa458d9222498e4d63f34527cf7318c1506f89d9092b35fdbcb64f3a.
//
// Solidity: event StakeGlp(address account, uint256 amount)
func (_Rewardrouter *RewardrouterFilterer) FilterStakeGlp(opts *bind.FilterOpts) (*RewardrouterStakeGlpIterator, error) {

	logs, sub, err := _Rewardrouter.contract.FilterLogs(opts, "StakeGlp")
	if err != nil {
		return nil, err
	}
	return &RewardrouterStakeGlpIterator{contract: _Rewardrouter.contract, event: "StakeGlp", logs: logs, sub: sub}, nil
}

// WatchStakeGlp is a free log subscription operation binding the contract event 0xa4725d47fa458d9222498e4d63f34527cf7318c1506f89d9092b35fdbcb64f3a.
//
// Solidity: event StakeGlp(address account, uint256 amount)
func (_Rewardrouter *RewardrouterFilterer) WatchStakeGlp(opts *bind.WatchOpts, sink chan<- *RewardrouterStakeGlp) (event.Subscription, error) {

	logs, sub, err := _Rewardrouter.contract.WatchLogs(opts, "StakeGlp")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardrouterStakeGlp)
				if err := _Rewardrouter.contract.UnpackLog(event, "StakeGlp", log); err != nil {
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

// ParseStakeGlp is a log parse operation binding the contract event 0xa4725d47fa458d9222498e4d63f34527cf7318c1506f89d9092b35fdbcb64f3a.
//
// Solidity: event StakeGlp(address account, uint256 amount)
func (_Rewardrouter *RewardrouterFilterer) ParseStakeGlp(log types.Log) (*RewardrouterStakeGlp, error) {
	event := new(RewardrouterStakeGlp)
	if err := _Rewardrouter.contract.UnpackLog(event, "StakeGlp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardrouterStakeGmxIterator is returned from FilterStakeGmx and is used to iterate over the raw logs and unpacked data for StakeGmx events raised by the Rewardrouter contract.
type RewardrouterStakeGmxIterator struct {
	Event *RewardrouterStakeGmx // Event containing the contract specifics and raw log

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
func (it *RewardrouterStakeGmxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardrouterStakeGmx)
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
		it.Event = new(RewardrouterStakeGmx)
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
func (it *RewardrouterStakeGmxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardrouterStakeGmxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardrouterStakeGmx represents a StakeGmx event raised by the Rewardrouter contract.
type RewardrouterStakeGmx struct {
	Account common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeGmx is a free log retrieval operation binding the contract event 0xad0723806aa1e5a8fb826fc9f0c5b589e585a6b60dc768a1b20691c95062d2d6.
//
// Solidity: event StakeGmx(address account, address token, uint256 amount)
func (_Rewardrouter *RewardrouterFilterer) FilterStakeGmx(opts *bind.FilterOpts) (*RewardrouterStakeGmxIterator, error) {

	logs, sub, err := _Rewardrouter.contract.FilterLogs(opts, "StakeGmx")
	if err != nil {
		return nil, err
	}
	return &RewardrouterStakeGmxIterator{contract: _Rewardrouter.contract, event: "StakeGmx", logs: logs, sub: sub}, nil
}

// WatchStakeGmx is a free log subscription operation binding the contract event 0xad0723806aa1e5a8fb826fc9f0c5b589e585a6b60dc768a1b20691c95062d2d6.
//
// Solidity: event StakeGmx(address account, address token, uint256 amount)
func (_Rewardrouter *RewardrouterFilterer) WatchStakeGmx(opts *bind.WatchOpts, sink chan<- *RewardrouterStakeGmx) (event.Subscription, error) {

	logs, sub, err := _Rewardrouter.contract.WatchLogs(opts, "StakeGmx")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardrouterStakeGmx)
				if err := _Rewardrouter.contract.UnpackLog(event, "StakeGmx", log); err != nil {
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

// ParseStakeGmx is a log parse operation binding the contract event 0xad0723806aa1e5a8fb826fc9f0c5b589e585a6b60dc768a1b20691c95062d2d6.
//
// Solidity: event StakeGmx(address account, address token, uint256 amount)
func (_Rewardrouter *RewardrouterFilterer) ParseStakeGmx(log types.Log) (*RewardrouterStakeGmx, error) {
	event := new(RewardrouterStakeGmx)
	if err := _Rewardrouter.contract.UnpackLog(event, "StakeGmx", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardrouterUnstakeGlpIterator is returned from FilterUnstakeGlp and is used to iterate over the raw logs and unpacked data for UnstakeGlp events raised by the Rewardrouter contract.
type RewardrouterUnstakeGlpIterator struct {
	Event *RewardrouterUnstakeGlp // Event containing the contract specifics and raw log

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
func (it *RewardrouterUnstakeGlpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardrouterUnstakeGlp)
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
		it.Event = new(RewardrouterUnstakeGlp)
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
func (it *RewardrouterUnstakeGlpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardrouterUnstakeGlpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardrouterUnstakeGlp represents a UnstakeGlp event raised by the Rewardrouter contract.
type RewardrouterUnstakeGlp struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnstakeGlp is a free log retrieval operation binding the contract event 0x1cb6202519b6b6c72ba5ed11e2c3f53af3cea010f96bfc705584e53e75cf034c.
//
// Solidity: event UnstakeGlp(address account, uint256 amount)
func (_Rewardrouter *RewardrouterFilterer) FilterUnstakeGlp(opts *bind.FilterOpts) (*RewardrouterUnstakeGlpIterator, error) {

	logs, sub, err := _Rewardrouter.contract.FilterLogs(opts, "UnstakeGlp")
	if err != nil {
		return nil, err
	}
	return &RewardrouterUnstakeGlpIterator{contract: _Rewardrouter.contract, event: "UnstakeGlp", logs: logs, sub: sub}, nil
}

// WatchUnstakeGlp is a free log subscription operation binding the contract event 0x1cb6202519b6b6c72ba5ed11e2c3f53af3cea010f96bfc705584e53e75cf034c.
//
// Solidity: event UnstakeGlp(address account, uint256 amount)
func (_Rewardrouter *RewardrouterFilterer) WatchUnstakeGlp(opts *bind.WatchOpts, sink chan<- *RewardrouterUnstakeGlp) (event.Subscription, error) {

	logs, sub, err := _Rewardrouter.contract.WatchLogs(opts, "UnstakeGlp")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardrouterUnstakeGlp)
				if err := _Rewardrouter.contract.UnpackLog(event, "UnstakeGlp", log); err != nil {
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

// ParseUnstakeGlp is a log parse operation binding the contract event 0x1cb6202519b6b6c72ba5ed11e2c3f53af3cea010f96bfc705584e53e75cf034c.
//
// Solidity: event UnstakeGlp(address account, uint256 amount)
func (_Rewardrouter *RewardrouterFilterer) ParseUnstakeGlp(log types.Log) (*RewardrouterUnstakeGlp, error) {
	event := new(RewardrouterUnstakeGlp)
	if err := _Rewardrouter.contract.UnpackLog(event, "UnstakeGlp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardrouterUnstakeGmxIterator is returned from FilterUnstakeGmx and is used to iterate over the raw logs and unpacked data for UnstakeGmx events raised by the Rewardrouter contract.
type RewardrouterUnstakeGmxIterator struct {
	Event *RewardrouterUnstakeGmx // Event containing the contract specifics and raw log

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
func (it *RewardrouterUnstakeGmxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardrouterUnstakeGmx)
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
		it.Event = new(RewardrouterUnstakeGmx)
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
func (it *RewardrouterUnstakeGmxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardrouterUnstakeGmxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardrouterUnstakeGmx represents a UnstakeGmx event raised by the Rewardrouter contract.
type RewardrouterUnstakeGmx struct {
	Account common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnstakeGmx is a free log retrieval operation binding the contract event 0xce8eb393006add0768cc6cefb3ca0fc4787015ce1ac86bd800e72a7999310345.
//
// Solidity: event UnstakeGmx(address account, address token, uint256 amount)
func (_Rewardrouter *RewardrouterFilterer) FilterUnstakeGmx(opts *bind.FilterOpts) (*RewardrouterUnstakeGmxIterator, error) {

	logs, sub, err := _Rewardrouter.contract.FilterLogs(opts, "UnstakeGmx")
	if err != nil {
		return nil, err
	}
	return &RewardrouterUnstakeGmxIterator{contract: _Rewardrouter.contract, event: "UnstakeGmx", logs: logs, sub: sub}, nil
}

// WatchUnstakeGmx is a free log subscription operation binding the contract event 0xce8eb393006add0768cc6cefb3ca0fc4787015ce1ac86bd800e72a7999310345.
//
// Solidity: event UnstakeGmx(address account, address token, uint256 amount)
func (_Rewardrouter *RewardrouterFilterer) WatchUnstakeGmx(opts *bind.WatchOpts, sink chan<- *RewardrouterUnstakeGmx) (event.Subscription, error) {

	logs, sub, err := _Rewardrouter.contract.WatchLogs(opts, "UnstakeGmx")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardrouterUnstakeGmx)
				if err := _Rewardrouter.contract.UnpackLog(event, "UnstakeGmx", log); err != nil {
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

// ParseUnstakeGmx is a log parse operation binding the contract event 0xce8eb393006add0768cc6cefb3ca0fc4787015ce1ac86bd800e72a7999310345.
//
// Solidity: event UnstakeGmx(address account, address token, uint256 amount)
func (_Rewardrouter *RewardrouterFilterer) ParseUnstakeGmx(log types.Log) (*RewardrouterUnstakeGmx, error) {
	event := new(RewardrouterUnstakeGmx)
	if err := _Rewardrouter.contract.UnpackLog(event, "UnstakeGmx", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
