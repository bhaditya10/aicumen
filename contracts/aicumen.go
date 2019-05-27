// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AicumenABI is the input ABI used to generate the binding from.
const AicumenABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"entityAddress\",\"type\":\"address\"}],\"name\":\"isEntity\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"entityStructs\",\"outputs\":[{\"name\":\"entityData\",\"type\":\"string\"},{\"name\":\"isEntity\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"entityAddress\",\"type\":\"address\"}],\"name\":\"getEntity\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"entityAddress\",\"type\":\"address\"},{\"name\":\"entityData\",\"type\":\"string\"}],\"name\":\"newEntity\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AicumenBin is the compiled bytecode used for deploying new contracts.
const AicumenBin = `0x608060405234801561001057600080fd5b50610525806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806314887c58146100515780634dd448bf1461008b57806375894e8c1461013457806396c48c2d146101cf575b600080fd5b6100776004803603602081101561006757600080fd5b50356001600160a01b0316610285565b604080519115158252519081900360200190f35b6100b1600480360360208110156100a157600080fd5b50356001600160a01b03166102a6565b604051808060200183151515158152602001828103825284818151815260200191508051906020019080838360005b838110156100f85781810151838201526020016100e0565b50505050905090810190601f1680156101255780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b61015a6004803603602081101561014a57600080fd5b50356001600160a01b031661034d565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561019457818101518382015260200161017c565b50505050905090810190601f1680156101c15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610077600480360360408110156101e557600080fd5b6001600160a01b03823516919081019060408101602082013564010000000081111561021057600080fd5b82018360208201111561022257600080fd5b8035906020019184600183028401116401000000008311171561024457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506103f6945050505050565b6001600160a01b031660009081526020819052604090206001015460ff1690565b6000602081815291815260409081902080548251601f60026000196101006001861615020190931692909204918201859004850281018501909352808352909283919083018282801561033a5780601f1061030f5761010080835404028352916020019161033a565b820191906000526020600020905b81548152906001019060200180831161031d57829003601f168201915b5050506001909301549192505060ff1682565b6001600160a01b0381166000908152602081815260409182902080548351601f60026000196101006001861615020190931692909204918201849004840281018401909452808452606093928301828280156103ea5780601f106103bf576101008083540402835291602001916103ea565b820191906000526020600020905b8154815290600101906020018083116103cd57829003601f168201915b50505050509050919050565b600061040183610285565b1561040b57600080fd5b6001600160a01b03831660009081526020818152604090912083516104329285019061045e565b5050506001600160a01b031660009081526020819052604090206001908101805460ff19168217905590565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061049f57805160ff19168380011785556104cc565b828001600101855582156104cc579182015b828111156104cc5782518255916020019190600101906104b1565b506104d89291506104dc565b5090565b6104f691905b808211156104d857600081556001016104e2565b9056fea165627a7a72305820bda8a03d924394093d8a5917c0ca1548a14ac64b2c1725609fb3867a3c850a6c0029`

// DeployAicumen deploys a new Ethereum contract, binding an instance of Aicumen to it.
func DeployAicumen(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Aicumen, error) {
	parsed, err := abi.JSON(strings.NewReader(AicumenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AicumenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Aicumen{AicumenCaller: AicumenCaller{contract: contract}, AicumenTransactor: AicumenTransactor{contract: contract}, AicumenFilterer: AicumenFilterer{contract: contract}}, nil
}

// Aicumen is an auto generated Go binding around an Ethereum contract.
type Aicumen struct {
	AicumenCaller     // Read-only binding to the contract
	AicumenTransactor // Write-only binding to the contract
	AicumenFilterer   // Log filterer for contract events
}

// AicumenCaller is an auto generated read-only Go binding around an Ethereum contract.
type AicumenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AicumenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AicumenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AicumenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AicumenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AicumenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AicumenSession struct {
	Contract     *Aicumen          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AicumenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AicumenCallerSession struct {
	Contract *AicumenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AicumenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AicumenTransactorSession struct {
	Contract     *AicumenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AicumenRaw is an auto generated low-level Go binding around an Ethereum contract.
type AicumenRaw struct {
	Contract *Aicumen // Generic contract binding to access the raw methods on
}

// AicumenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AicumenCallerRaw struct {
	Contract *AicumenCaller // Generic read-only contract binding to access the raw methods on
}

// AicumenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AicumenTransactorRaw struct {
	Contract *AicumenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAicumen creates a new instance of Aicumen, bound to a specific deployed contract.
func NewAicumen(address common.Address, backend bind.ContractBackend) (*Aicumen, error) {
	contract, err := bindAicumen(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aicumen{AicumenCaller: AicumenCaller{contract: contract}, AicumenTransactor: AicumenTransactor{contract: contract}, AicumenFilterer: AicumenFilterer{contract: contract}}, nil
}

// NewAicumenCaller creates a new read-only instance of Aicumen, bound to a specific deployed contract.
func NewAicumenCaller(address common.Address, caller bind.ContractCaller) (*AicumenCaller, error) {
	contract, err := bindAicumen(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AicumenCaller{contract: contract}, nil
}

// NewAicumenTransactor creates a new write-only instance of Aicumen, bound to a specific deployed contract.
func NewAicumenTransactor(address common.Address, transactor bind.ContractTransactor) (*AicumenTransactor, error) {
	contract, err := bindAicumen(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AicumenTransactor{contract: contract}, nil
}

// NewAicumenFilterer creates a new log filterer instance of Aicumen, bound to a specific deployed contract.
func NewAicumenFilterer(address common.Address, filterer bind.ContractFilterer) (*AicumenFilterer, error) {
	contract, err := bindAicumen(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AicumenFilterer{contract: contract}, nil
}

// bindAicumen binds a generic wrapper to an already deployed contract.
func bindAicumen(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AicumenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aicumen *AicumenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Aicumen.Contract.AicumenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aicumen *AicumenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aicumen.Contract.AicumenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aicumen *AicumenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aicumen.Contract.AicumenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aicumen *AicumenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Aicumen.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aicumen *AicumenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aicumen.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aicumen *AicumenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aicumen.Contract.contract.Transact(opts, method, params...)
}

// EntityStructs is a free data retrieval call binding the contract method 0x4dd448bf.
//
// Solidity: function entityStructs(address ) constant returns(string entityData, bool isEntity)
func (_Aicumen *AicumenCaller) EntityStructs(opts *bind.CallOpts, arg0 common.Address) (struct {
	EntityData string
	IsEntity   bool
}, error) {
	ret := new(struct {
		EntityData string
		IsEntity   bool
	})
	out := ret
	err := _Aicumen.contract.Call(opts, out, "entityStructs", arg0)
	return *ret, err
}

// EntityStructs is a free data retrieval call binding the contract method 0x4dd448bf.
//
// Solidity: function entityStructs(address ) constant returns(string entityData, bool isEntity)
func (_Aicumen *AicumenSession) EntityStructs(arg0 common.Address) (struct {
	EntityData string
	IsEntity   bool
}, error) {
	return _Aicumen.Contract.EntityStructs(&_Aicumen.CallOpts, arg0)
}

// EntityStructs is a free data retrieval call binding the contract method 0x4dd448bf.
//
// Solidity: function entityStructs(address ) constant returns(string entityData, bool isEntity)
func (_Aicumen *AicumenCallerSession) EntityStructs(arg0 common.Address) (struct {
	EntityData string
	IsEntity   bool
}, error) {
	return _Aicumen.Contract.EntityStructs(&_Aicumen.CallOpts, arg0)
}

// GetEntity is a free data retrieval call binding the contract method 0x75894e8c.
//
// Solidity: function getEntity(address entityAddress) constant returns(string)
func (_Aicumen *AicumenCaller) GetEntity(opts *bind.CallOpts, entityAddress common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Aicumen.contract.Call(opts, out, "getEntity", entityAddress)
	return *ret0, err
}

// GetEntity is a free data retrieval call binding the contract method 0x75894e8c.
//
// Solidity: function getEntity(address entityAddress) constant returns(string)
func (_Aicumen *AicumenSession) GetEntity(entityAddress common.Address) (string, error) {
	return _Aicumen.Contract.GetEntity(&_Aicumen.CallOpts, entityAddress)
}

// GetEntity is a free data retrieval call binding the contract method 0x75894e8c.
//
// Solidity: function getEntity(address entityAddress) constant returns(string)
func (_Aicumen *AicumenCallerSession) GetEntity(entityAddress common.Address) (string, error) {
	return _Aicumen.Contract.GetEntity(&_Aicumen.CallOpts, entityAddress)
}

// IsEntity is a free data retrieval call binding the contract method 0x14887c58.
//
// Solidity: function isEntity(address entityAddress) constant returns(bool isIndeed)
func (_Aicumen *AicumenCaller) IsEntity(opts *bind.CallOpts, entityAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Aicumen.contract.Call(opts, out, "isEntity", entityAddress)
	return *ret0, err
}

// IsEntity is a free data retrieval call binding the contract method 0x14887c58.
//
// Solidity: function isEntity(address entityAddress) constant returns(bool isIndeed)
func (_Aicumen *AicumenSession) IsEntity(entityAddress common.Address) (bool, error) {
	return _Aicumen.Contract.IsEntity(&_Aicumen.CallOpts, entityAddress)
}

// IsEntity is a free data retrieval call binding the contract method 0x14887c58.
//
// Solidity: function isEntity(address entityAddress) constant returns(bool isIndeed)
func (_Aicumen *AicumenCallerSession) IsEntity(entityAddress common.Address) (bool, error) {
	return _Aicumen.Contract.IsEntity(&_Aicumen.CallOpts, entityAddress)
}

// NewEntity is a paid mutator transaction binding the contract method 0x96c48c2d.
//
// Solidity: function newEntity(address entityAddress, string entityData) returns(bool success)
func (_Aicumen *AicumenTransactor) NewEntity(opts *bind.TransactOpts, entityAddress common.Address, entityData string) (*types.Transaction, error) {
	return _Aicumen.contract.Transact(opts, "newEntity", entityAddress, entityData)
}

// NewEntity is a paid mutator transaction binding the contract method 0x96c48c2d.
//
// Solidity: function newEntity(address entityAddress, string entityData) returns(bool success)
func (_Aicumen *AicumenSession) NewEntity(entityAddress common.Address, entityData string) (*types.Transaction, error) {
	return _Aicumen.Contract.NewEntity(&_Aicumen.TransactOpts, entityAddress, entityData)
}

// NewEntity is a paid mutator transaction binding the contract method 0x96c48c2d.
//
// Solidity: function newEntity(address entityAddress, string entityData) returns(bool success)
func (_Aicumen *AicumenTransactorSession) NewEntity(entityAddress common.Address, entityData string) (*types.Transaction, error) {
	return _Aicumen.Contract.NewEntity(&_Aicumen.TransactOpts, entityAddress, entityData)
}
