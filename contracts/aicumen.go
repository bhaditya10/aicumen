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
const AicumenABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getEntity\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setEntity\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AicumenBin is the compiled bytecode used for deploying new contracts.
const AicumenBin = `0x608060405234801561001057600080fd5b506104d2806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063070c412b1461003b578063760c284d14610156575b600080fd5b6100e16004803603602081101561005157600080fd5b81019060208101813564010000000081111561006c57600080fd5b82018360208201111561007e57600080fd5b803590602001918460018302840111640100000000831117156100a057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610297945050505050565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561011b578181015183820152602001610103565b50505050905090810190601f1680156101485780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102836004803603604081101561016c57600080fd5b81019060208101813564010000000081111561018757600080fd5b82018360208201111561019957600080fd5b803590602001918460018302840111640100000000831117156101bb57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561020e57600080fd5b82018360208201111561022057600080fd5b8035906020019184600183028401116401000000008311171561024257600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061038b945050505050565b604080519115158252519081900360200190f35b60606000826040518082805190602001908083835b602083106102cb5780518252601f1990920191602091820191016102ac565b518151600019602094850361010090810a820192831692199390931691909117909252949092019687526040805197889003820188208054601f600260018316159098029095011695909504928301829004820288018201905281875292945092505083018282801561037f5780601f106103545761010080835404028352916020019161037f565b820191906000526020600020905b81548152906001019060200180831161036257829003601f168201915b50505050509050919050565b6000826000836040518082805190602001908083835b602083106103c05780518252601f1990920191602091820191016103a1565b51815160209384036101000a60001901801990921691161790529201948552506040519384900381019093208451610401959194919091019250905061040b565b5060019392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061044c57805160ff1916838001178555610479565b82800160010185558215610479579182015b8281111561047957825182559160200191906001019061045e565b50610485929150610489565b5090565b6104a391905b80821115610485576000815560010161048f565b9056fea165627a7a72305820b35641fac88521aba45f4a324c5dfc42b58c7ef59a38caa9694c7aeb59542b320029`

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

// GetEntity is a free data retrieval call binding the contract method 0x070c412b.
//
// Solidity: function getEntity(string name) constant returns(string)
func (_Aicumen *AicumenCaller) GetEntity(opts *bind.CallOpts, name string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Aicumen.contract.Call(opts, out, "getEntity", name)
	return *ret0, err
}

// GetEntity is a free data retrieval call binding the contract method 0x070c412b.
//
// Solidity: function getEntity(string name) constant returns(string)
func (_Aicumen *AicumenSession) GetEntity(name string) (string, error) {
	return _Aicumen.Contract.GetEntity(&_Aicumen.CallOpts, name)
}

// GetEntity is a free data retrieval call binding the contract method 0x070c412b.
//
// Solidity: function getEntity(string name) constant returns(string)
func (_Aicumen *AicumenCallerSession) GetEntity(name string) (string, error) {
	return _Aicumen.Contract.GetEntity(&_Aicumen.CallOpts, name)
}

// SetEntity is a paid mutator transaction binding the contract method 0x760c284d.
//
// Solidity: function setEntity(string addr, string name) returns(bool)
func (_Aicumen *AicumenTransactor) SetEntity(opts *bind.TransactOpts, addr string, name string) (*types.Transaction, error) {
	return _Aicumen.contract.Transact(opts, "setEntity", addr, name)
}

// SetEntity is a paid mutator transaction binding the contract method 0x760c284d.
//
// Solidity: function setEntity(string addr, string name) returns(bool)
func (_Aicumen *AicumenSession) SetEntity(addr string, name string) (*types.Transaction, error) {
	return _Aicumen.Contract.SetEntity(&_Aicumen.TransactOpts, addr, name)
}

// SetEntity is a paid mutator transaction binding the contract method 0x760c284d.
//
// Solidity: function setEntity(string addr, string name) returns(bool)
func (_Aicumen *AicumenTransactorSession) SetEntity(addr string, name string) (*types.Transaction, error) {
	return _Aicumen.Contract.SetEntity(&_Aicumen.TransactOpts, addr, name)
}
