// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dex

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

// DexMetaData contains all meta data concerning the Dex contract.
var DexMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractERC20Votes\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Bought\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractERC20Votes\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161034138038061034183398101604081905261002f91610054565b600080546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b6102ae806100936000396000f3fe6080604052600436106100295760003560e01c8063a6f2ae3a1461002e578063fc0c546a14610038575b600080fd5b610036610074565b005b34801561004457600080fd5b50600054610058906001600160a01b031681565b6040516001600160a01b03909116815260200160405180910390f35b600080546040516370a0823160e01b81523060048201523492916001600160a01b0316906370a0823190602401602060405180830381865afa1580156100be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100e29190610236565b9050600082116101395760405162461bcd60e51b815260206004820152601b60248201527f596f75206e65656420746f2073656e6420736f6d65206574686572000000000060448201526064015b60405180910390fd5b808211156101895760405162461bcd60e51b815260206004820181905260248201527f4e6f7420656e6f75676820746f6b656e7320696e2074686520726573657276656044820152606401610130565b60005460405163a9059cbb60e01b8152336004820152602481018490526001600160a01b039091169063a9059cbb906044016020604051808303816000875af11580156101da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101fe919061024f565b506040518281527f4e08ba899977cf7d4c2964bce71c6b9a7ef76ee5166a4c1249a1e08016e33ef19060200160405180910390a15050565b60006020828403121561024857600080fd5b5051919050565b60006020828403121561026157600080fd5b8151801515811461027157600080fd5b939250505056fea2646970667358221220d8b507a3e8dba35e955f798f8c312fcb8e09e51d7d7f077397729655a364279564736f6c63430008100033",
}

// DexABI is the input ABI used to generate the binding from.
// Deprecated: Use DexMetaData.ABI instead.
var DexABI = DexMetaData.ABI

// DexBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DexMetaData.Bin instead.
var DexBin = DexMetaData.Bin

// DeployDex deploys a new Ethereum contract, binding an instance of Dex to it.
func DeployDex(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *Dex, error) {
	parsed, err := DexMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DexBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Dex{DexCaller: DexCaller{contract: contract}, DexTransactor: DexTransactor{contract: contract}, DexFilterer: DexFilterer{contract: contract}}, nil
}

// Dex is an auto generated Go binding around an Ethereum contract.
type Dex struct {
	DexCaller     // Read-only binding to the contract
	DexTransactor // Write-only binding to the contract
	DexFilterer   // Log filterer for contract events
}

// DexCaller is an auto generated read-only Go binding around an Ethereum contract.
type DexCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DexTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DexFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DexSession struct {
	Contract     *Dex              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DexCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DexCallerSession struct {
	Contract *DexCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DexTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DexTransactorSession struct {
	Contract     *DexTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DexRaw is an auto generated low-level Go binding around an Ethereum contract.
type DexRaw struct {
	Contract *Dex // Generic contract binding to access the raw methods on
}

// DexCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DexCallerRaw struct {
	Contract *DexCaller // Generic read-only contract binding to access the raw methods on
}

// DexTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DexTransactorRaw struct {
	Contract *DexTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDex creates a new instance of Dex, bound to a specific deployed contract.
func NewDex(address common.Address, backend bind.ContractBackend) (*Dex, error) {
	contract, err := bindDex(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dex{DexCaller: DexCaller{contract: contract}, DexTransactor: DexTransactor{contract: contract}, DexFilterer: DexFilterer{contract: contract}}, nil
}

// NewDexCaller creates a new read-only instance of Dex, bound to a specific deployed contract.
func NewDexCaller(address common.Address, caller bind.ContractCaller) (*DexCaller, error) {
	contract, err := bindDex(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DexCaller{contract: contract}, nil
}

// NewDexTransactor creates a new write-only instance of Dex, bound to a specific deployed contract.
func NewDexTransactor(address common.Address, transactor bind.ContractTransactor) (*DexTransactor, error) {
	contract, err := bindDex(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DexTransactor{contract: contract}, nil
}

// NewDexFilterer creates a new log filterer instance of Dex, bound to a specific deployed contract.
func NewDexFilterer(address common.Address, filterer bind.ContractFilterer) (*DexFilterer, error) {
	contract, err := bindDex(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DexFilterer{contract: contract}, nil
}

// bindDex binds a generic wrapper to an already deployed contract.
func bindDex(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DexABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dex *DexRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dex.Contract.DexCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dex *DexRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dex.Contract.DexTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dex *DexRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dex.Contract.DexTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dex *DexCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dex.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dex *DexTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dex.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dex *DexTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dex.Contract.contract.Transact(opts, method, params...)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Dex *DexCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dex.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Dex *DexSession) Token() (common.Address, error) {
	return _Dex.Contract.Token(&_Dex.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Dex *DexCallerSession) Token() (common.Address, error) {
	return _Dex.Contract.Token(&_Dex.CallOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() payable returns()
func (_Dex *DexTransactor) Buy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dex.contract.Transact(opts, "buy")
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() payable returns()
func (_Dex *DexSession) Buy() (*types.Transaction, error) {
	return _Dex.Contract.Buy(&_Dex.TransactOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() payable returns()
func (_Dex *DexTransactorSession) Buy() (*types.Transaction, error) {
	return _Dex.Contract.Buy(&_Dex.TransactOpts)
}

// DexBoughtIterator is returned from FilterBought and is used to iterate over the raw logs and unpacked data for Bought events raised by the Dex contract.
type DexBoughtIterator struct {
	Event *DexBought // Event containing the contract specifics and raw log

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
func (it *DexBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexBought)
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
		it.Event = new(DexBought)
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
func (it *DexBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexBought represents a Bought event raised by the Dex contract.
type DexBought struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBought is a free log retrieval operation binding the contract event 0x4e08ba899977cf7d4c2964bce71c6b9a7ef76ee5166a4c1249a1e08016e33ef1.
//
// Solidity: event Bought(uint256 amount)
func (_Dex *DexFilterer) FilterBought(opts *bind.FilterOpts) (*DexBoughtIterator, error) {

	logs, sub, err := _Dex.contract.FilterLogs(opts, "Bought")
	if err != nil {
		return nil, err
	}
	return &DexBoughtIterator{contract: _Dex.contract, event: "Bought", logs: logs, sub: sub}, nil
}

// WatchBought is a free log subscription operation binding the contract event 0x4e08ba899977cf7d4c2964bce71c6b9a7ef76ee5166a4c1249a1e08016e33ef1.
//
// Solidity: event Bought(uint256 amount)
func (_Dex *DexFilterer) WatchBought(opts *bind.WatchOpts, sink chan<- *DexBought) (event.Subscription, error) {

	logs, sub, err := _Dex.contract.WatchLogs(opts, "Bought")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexBought)
				if err := _Dex.contract.UnpackLog(event, "Bought", log); err != nil {
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

// ParseBought is a log parse operation binding the contract event 0x4e08ba899977cf7d4c2964bce71c6b9a7ef76ee5166a4c1249a1e08016e33ef1.
//
// Solidity: event Bought(uint256 amount)
func (_Dex *DexFilterer) ParseBought(log types.Log) (*DexBought, error) {
	event := new(DexBought)
	if err := _Dex.contract.UnpackLog(event, "Bought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
