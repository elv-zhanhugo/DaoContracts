// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package box

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

// BoxMetaData contains all meta data concerning the Box contract.
var BoxMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"newValue\",\"type\":\"string[]\"}],\"name\":\"ValueChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"newValue\",\"type\":\"string[]\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610033610025640100000000610038810204565b64010000000061003c810204565b61008c565b3390565b60008054600160a060020a03838116600160a060020a0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b61085e8061009b6000396000f3fe608060405234801561001057600080fd5b5060043610610073577c010000000000000000000000000000000000000000000000000000000060003504632e64cec18114610078578063715018a6146100965780638da5cb5b146100a0578063ae00106a146100bb578063f2fde38b146100ce575b600080fd5b6100806100e1565b60405161008d9190610449565b60405180910390f35b61009e6101ba565b005b600054604051600160a060020a03909116815260200161008d565b61009e6100c9366004610543565b6101fc565b61009e6100dc36600461065b565b610277565b60606001805480602002602001604051908101604052809291908181526020016000905b828210156101b15783829060005260206000200180546101249061068b565b80601f01602080910402602001604051908101604052809291908181526020018280546101509061068b565b801561019d5780601f106101725761010080835404028352916020019161019d565b820191906000526020600020905b81548152906001019060200180831161018057829003601f168201915b505050505081526020019060010190610105565b50505050905090565b600054600160a060020a031633146101f05760405160e560020a62461bcd0281526004016101e7906106de565b60405180910390fd5b6101fa600061032f565b565b600054600160a060020a031633146102295760405160e560020a62461bcd0281526004016101e7906106de565b805161023c90600190602084019061038c565b507fd265f101d7aa4a64ddf63396a75da5c42673f921d2d659866733f353fe40a6688160405161026c9190610449565b60405180910390a150565b600054600160a060020a031633146102a45760405160e560020a62461bcd0281526004016101e7906106de565b600160a060020a0381166103235760405160e560020a62461bcd02815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016101e7565b61032c8161032f565b50565b60008054600160a060020a0383811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b8280548282559060005260206000209081019282156103d2579160200282015b828111156103d257825182906103c29082610762565b50916020019190600101906103ac565b506103de9291506103e2565b5090565b808211156103de5760006103f682826103ff565b506001016103e2565b50805461040b9061068b565b6000825580601f1061041b575050565b601f01602090049060005260206000209081019061032c91905b808211156103de5760008155600101610435565b60006020808301818452808551808352604086019150604084820287010192508387016000805b838110156104d557888603603f1901855282518051808852835b818110156104a5578281018a01518982018b0152890161048a565b818111156104b557848a838b0101525b50601f01601f191696909601870195509386019391860191600101610470565b509398975050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561053b5761053b6104e3565b604052919050565b6000602080838503121561055657600080fd5b823567ffffffffffffffff8082111561056e57600080fd5b8185019150601f868184011261058357600080fd5b823582811115610595576105956104e3565b8481026105a3868201610512565b918252848101860191868101908a8411156105bd57600080fd5b87870192505b8383101561064d578235868111156105db5760008081fd5b8701603f81018c136105ed5760008081fd5b88810135604088821115610603576106036104e3565b610614828901601f19168c01610512565b8281528e828486010111156106295760008081fd5b828285018d83013760009281018c01929092525083525091870191908701906105c3565b9a9950505050505050505050565b60006020828403121561066d57600080fd5b8135600160a060020a038116811461068457600080fd5b9392505050565b60028104600182168061069f57607f821691505b6020821081036106d8577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b601f82111561075d576000818152602081206020601f8601048101602086101561073a5750805b6020601f860104820191505b8181101561075957828155600101610746565b5050505b505050565b815167ffffffffffffffff81111561077c5761077c6104e3565b6107908161078a845461068b565b84610713565b602080601f8311600181146107c957600084156107ad5750858301515b60028086026008870290910a6000190419821617865550610759565b600085815260208120601f198616915b828110156107f8578886015182559484019460019091019084016107d9565b508582101561081857878501516008601f88160260020a60001904191681555b505050505060020260010190555056fea264697066735822122073e5c6887f2822d32ae1ff9d1fd064bccfd0519a5a6f00653fae88674c37e8ea64736f6c634300080f0033",
}

// BoxABI is the input ABI used to generate the binding from.
// Deprecated: Use BoxMetaData.ABI instead.
var BoxABI = BoxMetaData.ABI

// BoxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BoxMetaData.Bin instead.
var BoxBin = BoxMetaData.Bin

// DeployBox deploys a new Ethereum contract, binding an instance of Box to it.
func DeployBox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Box, error) {
	parsed, err := BoxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BoxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Box{BoxCaller: BoxCaller{contract: contract}, BoxTransactor: BoxTransactor{contract: contract}, BoxFilterer: BoxFilterer{contract: contract}}, nil
}

// Box is an auto generated Go binding around an Ethereum contract.
type Box struct {
	BoxCaller     // Read-only binding to the contract
	BoxTransactor // Write-only binding to the contract
	BoxFilterer   // Log filterer for contract events
}

// BoxCaller is an auto generated read-only Go binding around an Ethereum contract.
type BoxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BoxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BoxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BoxSession struct {
	Contract     *Box              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BoxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BoxCallerSession struct {
	Contract *BoxCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BoxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BoxTransactorSession struct {
	Contract     *BoxTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BoxRaw is an auto generated low-level Go binding around an Ethereum contract.
type BoxRaw struct {
	Contract *Box // Generic contract binding to access the raw methods on
}

// BoxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BoxCallerRaw struct {
	Contract *BoxCaller // Generic read-only contract binding to access the raw methods on
}

// BoxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BoxTransactorRaw struct {
	Contract *BoxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBox creates a new instance of Box, bound to a specific deployed contract.
func NewBox(address common.Address, backend bind.ContractBackend) (*Box, error) {
	contract, err := bindBox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Box{BoxCaller: BoxCaller{contract: contract}, BoxTransactor: BoxTransactor{contract: contract}, BoxFilterer: BoxFilterer{contract: contract}}, nil
}

// NewBoxCaller creates a new read-only instance of Box, bound to a specific deployed contract.
func NewBoxCaller(address common.Address, caller bind.ContractCaller) (*BoxCaller, error) {
	contract, err := bindBox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BoxCaller{contract: contract}, nil
}

// NewBoxTransactor creates a new write-only instance of Box, bound to a specific deployed contract.
func NewBoxTransactor(address common.Address, transactor bind.ContractTransactor) (*BoxTransactor, error) {
	contract, err := bindBox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BoxTransactor{contract: contract}, nil
}

// NewBoxFilterer creates a new log filterer instance of Box, bound to a specific deployed contract.
func NewBoxFilterer(address common.Address, filterer bind.ContractFilterer) (*BoxFilterer, error) {
	contract, err := bindBox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BoxFilterer{contract: contract}, nil
}

// bindBox binds a generic wrapper to an already deployed contract.
func bindBox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BoxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Box *BoxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Box.Contract.BoxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Box *BoxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Box.Contract.BoxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Box *BoxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Box.Contract.BoxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Box *BoxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Box.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Box *BoxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Box.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Box *BoxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Box.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Box *BoxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Box.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Box *BoxSession) Owner() (common.Address, error) {
	return _Box.Contract.Owner(&_Box.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Box *BoxCallerSession) Owner() (common.Address, error) {
	return _Box.Contract.Owner(&_Box.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(string[])
func (_Box *BoxCaller) Retrieve(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Box.contract.Call(opts, &out, "retrieve")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(string[])
func (_Box *BoxSession) Retrieve() ([]string, error) {
	return _Box.Contract.Retrieve(&_Box.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(string[])
func (_Box *BoxCallerSession) Retrieve() ([]string, error) {
	return _Box.Contract.Retrieve(&_Box.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Box *BoxTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Box.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Box *BoxSession) RenounceOwnership() (*types.Transaction, error) {
	return _Box.Contract.RenounceOwnership(&_Box.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Box *BoxTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Box.Contract.RenounceOwnership(&_Box.TransactOpts)
}

// Store is a paid mutator transaction binding the contract method 0xae00106a.
//
// Solidity: function store(string[] newValue) returns()
func (_Box *BoxTransactor) Store(opts *bind.TransactOpts, newValue []string) (*types.Transaction, error) {
	return _Box.contract.Transact(opts, "store", newValue)
}

// Store is a paid mutator transaction binding the contract method 0xae00106a.
//
// Solidity: function store(string[] newValue) returns()
func (_Box *BoxSession) Store(newValue []string) (*types.Transaction, error) {
	return _Box.Contract.Store(&_Box.TransactOpts, newValue)
}

// Store is a paid mutator transaction binding the contract method 0xae00106a.
//
// Solidity: function store(string[] newValue) returns()
func (_Box *BoxTransactorSession) Store(newValue []string) (*types.Transaction, error) {
	return _Box.Contract.Store(&_Box.TransactOpts, newValue)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Box *BoxTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Box.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Box *BoxSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Box.Contract.TransferOwnership(&_Box.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Box *BoxTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Box.Contract.TransferOwnership(&_Box.TransactOpts, newOwner)
}

// BoxOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Box contract.
type BoxOwnershipTransferredIterator struct {
	Event *BoxOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BoxOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoxOwnershipTransferred)
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
		it.Event = new(BoxOwnershipTransferred)
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
func (it *BoxOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoxOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoxOwnershipTransferred represents a OwnershipTransferred event raised by the Box contract.
type BoxOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Box *BoxFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BoxOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Box.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BoxOwnershipTransferredIterator{contract: _Box.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Box *BoxFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BoxOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Box.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoxOwnershipTransferred)
				if err := _Box.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Box *BoxFilterer) ParseOwnershipTransferred(log types.Log) (*BoxOwnershipTransferred, error) {
	event := new(BoxOwnershipTransferred)
	if err := _Box.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BoxValueChangedIterator is returned from FilterValueChanged and is used to iterate over the raw logs and unpacked data for ValueChanged events raised by the Box contract.
type BoxValueChangedIterator struct {
	Event *BoxValueChanged // Event containing the contract specifics and raw log

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
func (it *BoxValueChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoxValueChanged)
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
		it.Event = new(BoxValueChanged)
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
func (it *BoxValueChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoxValueChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoxValueChanged represents a ValueChanged event raised by the Box contract.
type BoxValueChanged struct {
	NewValue []string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterValueChanged is a free log retrieval operation binding the contract event 0xd265f101d7aa4a64ddf63396a75da5c42673f921d2d659866733f353fe40a668.
//
// Solidity: event ValueChanged(string[] newValue)
func (_Box *BoxFilterer) FilterValueChanged(opts *bind.FilterOpts) (*BoxValueChangedIterator, error) {

	logs, sub, err := _Box.contract.FilterLogs(opts, "ValueChanged")
	if err != nil {
		return nil, err
	}
	return &BoxValueChangedIterator{contract: _Box.contract, event: "ValueChanged", logs: logs, sub: sub}, nil
}

// WatchValueChanged is a free log subscription operation binding the contract event 0xd265f101d7aa4a64ddf63396a75da5c42673f921d2d659866733f353fe40a668.
//
// Solidity: event ValueChanged(string[] newValue)
func (_Box *BoxFilterer) WatchValueChanged(opts *bind.WatchOpts, sink chan<- *BoxValueChanged) (event.Subscription, error) {

	logs, sub, err := _Box.contract.WatchLogs(opts, "ValueChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoxValueChanged)
				if err := _Box.contract.UnpackLog(event, "ValueChanged", log); err != nil {
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

// ParseValueChanged is a log parse operation binding the contract event 0xd265f101d7aa4a64ddf63396a75da5c42673f921d2d659866733f353fe40a668.
//
// Solidity: event ValueChanged(string[] newValue)
func (_Box *BoxFilterer) ParseValueChanged(log types.Log) (*BoxValueChanged, error) {
	event := new(BoxValueChanged)
	if err := _Box.contract.UnpackLog(event, "ValueChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
