// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// EventMetaData contains all meta data concerning the Event contract.
var EventMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ticketHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"clone\",\"type\":\"address\"}],\"name\":\"TicketCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ticketHash\",\"type\":\"string\"}],\"name\":\"createTicket\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ticketAddress\",\"type\":\"address\"}],\"name\":\"doesTicketExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"eventMinters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"eventTickets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factoryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ticketFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factoryAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ticketAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"mintNFT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ticketAddress\",\"type\":\"address\"}],\"name\":\"registerTicket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// EventABI is the input ABI used to generate the binding from.
// Deprecated: Use EventMetaData.ABI instead.
var EventABI = EventMetaData.ABI

// Event is an auto generated Go binding around an Ethereum contract.
type Event struct {
	EventCaller     // Read-only binding to the contract
	EventTransactor // Write-only binding to the contract
	EventFilterer   // Log filterer for contract events
}

// EventCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventSession struct {
	Contract     *Event            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventCallerSession struct {
	Contract *EventCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventTransactorSession struct {
	Contract     *EventTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventRaw struct {
	Contract *Event // Generic contract binding to access the raw methods on
}

// EventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventCallerRaw struct {
	Contract *EventCaller // Generic read-only contract binding to access the raw methods on
}

// EventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventTransactorRaw struct {
	Contract *EventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEvent creates a new instance of Event, bound to a specific deployed contract.
func NewEvent(address common.Address, backend bind.ContractBackend) (*Event, error) {
	contract, err := bindEvent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Event{EventCaller: EventCaller{contract: contract}, EventTransactor: EventTransactor{contract: contract}, EventFilterer: EventFilterer{contract: contract}}, nil
}

// NewEventCaller creates a new read-only instance of Event, bound to a specific deployed contract.
func NewEventCaller(address common.Address, caller bind.ContractCaller) (*EventCaller, error) {
	contract, err := bindEvent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventCaller{contract: contract}, nil
}

// NewEventTransactor creates a new write-only instance of Event, bound to a specific deployed contract.
func NewEventTransactor(address common.Address, transactor bind.ContractTransactor) (*EventTransactor, error) {
	contract, err := bindEvent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventTransactor{contract: contract}, nil
}

// NewEventFilterer creates a new log filterer instance of Event, bound to a specific deployed contract.
func NewEventFilterer(address common.Address, filterer bind.ContractFilterer) (*EventFilterer, error) {
	contract, err := bindEvent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventFilterer{contract: contract}, nil
}

// bindEvent binds a generic wrapper to an already deployed contract.
func bindEvent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EventMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Event *EventRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Event.Contract.EventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Event *EventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Event.Contract.EventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Event *EventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Event.Contract.EventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Event *EventCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Event.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Event *EventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Event.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Event *EventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Event.Contract.contract.Transact(opts, method, params...)
}

// DoesTicketExist is a free data retrieval call binding the contract method 0xd23bf05e.
//
// Solidity: function doesTicketExist(address ticketAddress) view returns(bool)
func (_Event *EventCaller) DoesTicketExist(opts *bind.CallOpts, ticketAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Event.contract.Call(opts, &out, "doesTicketExist", ticketAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DoesTicketExist is a free data retrieval call binding the contract method 0xd23bf05e.
//
// Solidity: function doesTicketExist(address ticketAddress) view returns(bool)
func (_Event *EventSession) DoesTicketExist(ticketAddress common.Address) (bool, error) {
	return _Event.Contract.DoesTicketExist(&_Event.CallOpts, ticketAddress)
}

// DoesTicketExist is a free data retrieval call binding the contract method 0xd23bf05e.
//
// Solidity: function doesTicketExist(address ticketAddress) view returns(bool)
func (_Event *EventCallerSession) DoesTicketExist(ticketAddress common.Address) (bool, error) {
	return _Event.Contract.DoesTicketExist(&_Event.CallOpts, ticketAddress)
}

// EventMinters is a free data retrieval call binding the contract method 0xdcb57a3a.
//
// Solidity: function eventMinters(address ) view returns(address)
func (_Event *EventCaller) EventMinters(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Event.contract.Call(opts, &out, "eventMinters", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EventMinters is a free data retrieval call binding the contract method 0xdcb57a3a.
//
// Solidity: function eventMinters(address ) view returns(address)
func (_Event *EventSession) EventMinters(arg0 common.Address) (common.Address, error) {
	return _Event.Contract.EventMinters(&_Event.CallOpts, arg0)
}

// EventMinters is a free data retrieval call binding the contract method 0xdcb57a3a.
//
// Solidity: function eventMinters(address ) view returns(address)
func (_Event *EventCallerSession) EventMinters(arg0 common.Address) (common.Address, error) {
	return _Event.Contract.EventMinters(&_Event.CallOpts, arg0)
}

// EventTickets is a free data retrieval call binding the contract method 0x89a46c92.
//
// Solidity: function eventTickets(uint256 ) view returns(address)
func (_Event *EventCaller) EventTickets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Event.contract.Call(opts, &out, "eventTickets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EventTickets is a free data retrieval call binding the contract method 0x89a46c92.
//
// Solidity: function eventTickets(uint256 ) view returns(address)
func (_Event *EventSession) EventTickets(arg0 *big.Int) (common.Address, error) {
	return _Event.Contract.EventTickets(&_Event.CallOpts, arg0)
}

// EventTickets is a free data retrieval call binding the contract method 0x89a46c92.
//
// Solidity: function eventTickets(uint256 ) view returns(address)
func (_Event *EventCallerSession) EventTickets(arg0 *big.Int) (common.Address, error) {
	return _Event.Contract.EventTickets(&_Event.CallOpts, arg0)
}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() view returns(address)
func (_Event *EventCaller) FactoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Event.contract.Call(opts, &out, "factoryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() view returns(address)
func (_Event *EventSession) FactoryAddress() (common.Address, error) {
	return _Event.Contract.FactoryAddress(&_Event.CallOpts)
}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() view returns(address)
func (_Event *EventCallerSession) FactoryAddress() (common.Address, error) {
	return _Event.Contract.FactoryAddress(&_Event.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Event *EventCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Event.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Event *EventSession) Owner() (common.Address, error) {
	return _Event.Contract.Owner(&_Event.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Event *EventCallerSession) Owner() (common.Address, error) {
	return _Event.Contract.Owner(&_Event.CallOpts)
}

// CreateTicket is a paid mutator transaction binding the contract method 0x271ba406.
//
// Solidity: function createTicket(string name, uint256 maxSupply, uint256 mintPrice, string ticketHash) returns(address)
func (_Event *EventTransactor) CreateTicket(opts *bind.TransactOpts, name string, maxSupply *big.Int, mintPrice *big.Int, ticketHash string) (*types.Transaction, error) {
	return _Event.contract.Transact(opts, "createTicket", name, maxSupply, mintPrice, ticketHash)
}

// CreateTicket is a paid mutator transaction binding the contract method 0x271ba406.
//
// Solidity: function createTicket(string name, uint256 maxSupply, uint256 mintPrice, string ticketHash) returns(address)
func (_Event *EventSession) CreateTicket(name string, maxSupply *big.Int, mintPrice *big.Int, ticketHash string) (*types.Transaction, error) {
	return _Event.Contract.CreateTicket(&_Event.TransactOpts, name, maxSupply, mintPrice, ticketHash)
}

// CreateTicket is a paid mutator transaction binding the contract method 0x271ba406.
//
// Solidity: function createTicket(string name, uint256 maxSupply, uint256 mintPrice, string ticketHash) returns(address)
func (_Event *EventTransactorSession) CreateTicket(name string, maxSupply *big.Int, mintPrice *big.Int, ticketHash string) (*types.Transaction, error) {
	return _Event.Contract.CreateTicket(&_Event.TransactOpts, name, maxSupply, mintPrice, ticketHash)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _ticketFactory, address owner, address _factoryAddress) returns()
func (_Event *EventTransactor) Initialize(opts *bind.TransactOpts, _ticketFactory common.Address, owner common.Address, _factoryAddress common.Address) (*types.Transaction, error) {
	return _Event.contract.Transact(opts, "initialize", _ticketFactory, owner, _factoryAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _ticketFactory, address owner, address _factoryAddress) returns()
func (_Event *EventSession) Initialize(_ticketFactory common.Address, owner common.Address, _factoryAddress common.Address) (*types.Transaction, error) {
	return _Event.Contract.Initialize(&_Event.TransactOpts, _ticketFactory, owner, _factoryAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _ticketFactory, address owner, address _factoryAddress) returns()
func (_Event *EventTransactorSession) Initialize(_ticketFactory common.Address, owner common.Address, _factoryAddress common.Address) (*types.Transaction, error) {
	return _Event.Contract.Initialize(&_Event.TransactOpts, _ticketFactory, owner, _factoryAddress)
}

// MintNFT is a paid mutator transaction binding the contract method 0xeacabe14.
//
// Solidity: function mintNFT(address ticketAddress, string tokenURI) payable returns(uint256)
func (_Event *EventTransactor) MintNFT(opts *bind.TransactOpts, ticketAddress common.Address, tokenURI string) (*types.Transaction, error) {
	return _Event.contract.Transact(opts, "mintNFT", ticketAddress, tokenURI)
}

// MintNFT is a paid mutator transaction binding the contract method 0xeacabe14.
//
// Solidity: function mintNFT(address ticketAddress, string tokenURI) payable returns(uint256)
func (_Event *EventSession) MintNFT(ticketAddress common.Address, tokenURI string) (*types.Transaction, error) {
	return _Event.Contract.MintNFT(&_Event.TransactOpts, ticketAddress, tokenURI)
}

// MintNFT is a paid mutator transaction binding the contract method 0xeacabe14.
//
// Solidity: function mintNFT(address ticketAddress, string tokenURI) payable returns(uint256)
func (_Event *EventTransactorSession) MintNFT(ticketAddress common.Address, tokenURI string) (*types.Transaction, error) {
	return _Event.Contract.MintNFT(&_Event.TransactOpts, ticketAddress, tokenURI)
}

// RegisterTicket is a paid mutator transaction binding the contract method 0xc1126a36.
//
// Solidity: function registerTicket(address ticketAddress) returns()
func (_Event *EventTransactor) RegisterTicket(opts *bind.TransactOpts, ticketAddress common.Address) (*types.Transaction, error) {
	return _Event.contract.Transact(opts, "registerTicket", ticketAddress)
}

// RegisterTicket is a paid mutator transaction binding the contract method 0xc1126a36.
//
// Solidity: function registerTicket(address ticketAddress) returns()
func (_Event *EventSession) RegisterTicket(ticketAddress common.Address) (*types.Transaction, error) {
	return _Event.Contract.RegisterTicket(&_Event.TransactOpts, ticketAddress)
}

// RegisterTicket is a paid mutator transaction binding the contract method 0xc1126a36.
//
// Solidity: function registerTicket(address ticketAddress) returns()
func (_Event *EventTransactorSession) RegisterTicket(ticketAddress common.Address) (*types.Transaction, error) {
	return _Event.Contract.RegisterTicket(&_Event.TransactOpts, ticketAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Event *EventTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Event.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Event *EventSession) RenounceOwnership() (*types.Transaction, error) {
	return _Event.Contract.RenounceOwnership(&_Event.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Event *EventTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Event.Contract.RenounceOwnership(&_Event.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Event *EventTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Event.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Event *EventSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Event.Contract.TransferOwnership(&_Event.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Event *EventTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Event.Contract.TransferOwnership(&_Event.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Event *EventTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Event.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Event *EventSession) Withdraw() (*types.Transaction, error) {
	return _Event.Contract.Withdraw(&_Event.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Event *EventTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Event.Contract.Withdraw(&_Event.TransactOpts)
}

// EventInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Event contract.
type EventInitializedIterator struct {
	Event *EventInitialized // Event containing the contract specifics and raw log

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
func (it *EventInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventInitialized)
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
		it.Event = new(EventInitialized)
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
func (it *EventInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventInitialized represents a Initialized event raised by the Event contract.
type EventInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Event *EventFilterer) FilterInitialized(opts *bind.FilterOpts) (*EventInitializedIterator, error) {

	logs, sub, err := _Event.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EventInitializedIterator{contract: _Event.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Event *EventFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EventInitialized) (event.Subscription, error) {

	logs, sub, err := _Event.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventInitialized)
				if err := _Event.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Event *EventFilterer) ParseInitialized(log types.Log) (*EventInitialized, error) {
	event := new(EventInitialized)
	if err := _Event.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Event contract.
type EventOwnershipTransferredIterator struct {
	Event *EventOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EventOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventOwnershipTransferred)
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
		it.Event = new(EventOwnershipTransferred)
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
func (it *EventOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventOwnershipTransferred represents a OwnershipTransferred event raised by the Event contract.
type EventOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Event *EventFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EventOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EventOwnershipTransferredIterator{contract: _Event.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Event *EventFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EventOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventOwnershipTransferred)
				if err := _Event.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Event *EventFilterer) ParseOwnershipTransferred(log types.Log) (*EventOwnershipTransferred, error) {
	event := new(EventOwnershipTransferred)
	if err := _Event.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventTicketCreatedIterator is returned from FilterTicketCreated and is used to iterate over the raw logs and unpacked data for TicketCreated events raised by the Event contract.
type EventTicketCreatedIterator struct {
	Event *EventTicketCreated // Event containing the contract specifics and raw log

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
func (it *EventTicketCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventTicketCreated)
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
		it.Event = new(EventTicketCreated)
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
func (it *EventTicketCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventTicketCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventTicketCreated represents a TicketCreated event raised by the Event contract.
type EventTicketCreated struct {
	TicketHash string
	Clone      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTicketCreated is a free log retrieval operation binding the contract event 0xd822f61f6a447402944f4fd46b908accaa48625d1599487cd6ae7b197cb12c63.
//
// Solidity: event TicketCreated(string ticketHash, address clone)
func (_Event *EventFilterer) FilterTicketCreated(opts *bind.FilterOpts) (*EventTicketCreatedIterator, error) {

	logs, sub, err := _Event.contract.FilterLogs(opts, "TicketCreated")
	if err != nil {
		return nil, err
	}
	return &EventTicketCreatedIterator{contract: _Event.contract, event: "TicketCreated", logs: logs, sub: sub}, nil
}

// WatchTicketCreated is a free log subscription operation binding the contract event 0xd822f61f6a447402944f4fd46b908accaa48625d1599487cd6ae7b197cb12c63.
//
// Solidity: event TicketCreated(string ticketHash, address clone)
func (_Event *EventFilterer) WatchTicketCreated(opts *bind.WatchOpts, sink chan<- *EventTicketCreated) (event.Subscription, error) {

	logs, sub, err := _Event.contract.WatchLogs(opts, "TicketCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventTicketCreated)
				if err := _Event.contract.UnpackLog(event, "TicketCreated", log); err != nil {
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

// ParseTicketCreated is a log parse operation binding the contract event 0xd822f61f6a447402944f4fd46b908accaa48625d1599487cd6ae7b197cb12c63.
//
// Solidity: event TicketCreated(string ticketHash, address clone)
func (_Event *EventFilterer) ParseTicketCreated(log types.Log) (*EventTicketCreated, error) {
	event := new(EventTicketCreated)
	if err := _Event.contract.UnpackLog(event, "TicketCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
