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
	Bin: "0x608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6107dc8061007e6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80632e64cec11461005c578063715018a61461007a5780638da5cb5b14610084578063ae00106a1461009f578063f2fde38b146100b2575b600080fd5b6100646100c5565b60405161007191906103fd565b60405180910390f35b61008261019e565b005b6000546040516001600160a01b039091168152602001610071565b6100826100ad3660046104df565b6101dd565b6100826100c03660046105f8565b610255565b60606001805480602002602001604051908101604052809291908181526020016000905b8282101561019557838290600052602060002001805461010890610628565b80601f016020809104026020016040519081016040528092919081815260200182805461013490610628565b80156101815780601f1061015657610100808354040283529160200191610181565b820191906000526020600020905b81548152906001019060200180831161016457829003601f168201915b5050505050815260200190600101906100e9565b50505050905090565b6000546001600160a01b031633146101d15760405162461bcd60e51b81526004016101c890610662565b60405180910390fd5b6101db60006102f0565b565b6000546001600160a01b031633146102075760405162461bcd60e51b81526004016101c890610662565b805161021a906001906020840190610340565b507fd265f101d7aa4a64ddf63396a75da5c42673f921d2d659866733f353fe40a6688160405161024a91906103fd565b60405180910390a150565b6000546001600160a01b0316331461027f5760405162461bcd60e51b81526004016101c890610662565b6001600160a01b0381166102e45760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016101c8565b6102ed816102f0565b50565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b828054828255906000526020600020908101928215610386579160200282015b82811115610386578251829061037690826106e6565b5091602001919060010190610360565b50610392929150610396565b5090565b808211156103925760006103aa82826103b3565b50600101610396565b5080546103bf90610628565b6000825580601f106103cf575050565b601f0160209004906000526020600020908101906102ed91905b8082111561039257600081556001016103e9565b6000602080830181845280855180835260408601915060408160051b87010192508387016000805b8381101561048a57888603603f1901855282518051808852835b8181101561045a578281018a01518982018b0152890161043f565b8181111561046a57848a838b0101525b50601f01601f191696909601870195509386019391860191600101610425565b509398975050505050505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156104d7576104d7610498565b604052919050565b600060208083850312156104f257600080fd5b823567ffffffffffffffff8082111561050a57600080fd5b8185019150601f868184011261051f57600080fd5b82358281111561053157610531610498565b8060051b6105408682016104ae565b918252848101860191868101908a84111561055a57600080fd5b87870192505b838310156105ea578235868111156105785760008081fd5b8701603f81018c1361058a5760008081fd5b888101356040888211156105a0576105a0610498565b6105b1828901601f19168c016104ae565b8281528e828486010111156105c65760008081fd5b828285018d83013760009281018c0192909252508352509187019190870190610560565b9a9950505050505050505050565b60006020828403121561060a57600080fd5b81356001600160a01b038116811461062157600080fd5b9392505050565b600181811c9082168061063c57607f821691505b60208210810361065c57634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b601f8211156106e157600081815260208120601f850160051c810160208610156106be5750805b601f850160051c820191505b818110156106dd578281556001016106ca565b5050505b505050565b815167ffffffffffffffff81111561070057610700610498565b6107148161070e8454610628565b84610697565b602080601f83116001811461074957600084156107315750858301515b600019600386901b1c1916600185901b1785556106dd565b600085815260208120601f198616915b8281101561077857888601518255948401946001909101908401610759565b50858210156107965787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea2646970667358221220878c1fda8500df93dd637e18f251bb8e50b22949645a8b09dd3cb2a9496391ca64736f6c634300080f0033",
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
