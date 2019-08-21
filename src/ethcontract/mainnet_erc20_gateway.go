// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethcontract

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

// ERC20GatewayABI is the input ABI used to generate the binding from.
const ERC20GatewayABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"vmc\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loomAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_vmc\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TokenWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"loomCoinAddress\",\"type\":\"address\"}],\"name\":\"LoomCoinReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"ERC20Received\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"contractAddress\",\"type\":\"address\"},{\"name\":\"_signersIndexes\",\"type\":\"uint256[]\"},{\"name\":\"_v\",\"type\":\"uint8[]\"},{\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"name\":\"_s\",\"type\":\"bytes32[]\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"depositERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"getERC20\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ERC20Gateway is an auto generated Go binding around an Ethereum contract.
type ERC20Gateway struct {
	ERC20GatewayCaller     // Read-only binding to the contract
	ERC20GatewayTransactor // Write-only binding to the contract
	ERC20GatewayFilterer   // Log filterer for contract events
}

// ERC20GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20GatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20GatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20GatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20GatewaySession struct {
	Contract     *ERC20Gateway     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20GatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20GatewayCallerSession struct {
	Contract *ERC20GatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ERC20GatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20GatewayTransactorSession struct {
	Contract     *ERC20GatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ERC20GatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20GatewayRaw struct {
	Contract *ERC20Gateway // Generic contract binding to access the raw methods on
}

// ERC20GatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20GatewayCallerRaw struct {
	Contract *ERC20GatewayCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20GatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20GatewayTransactorRaw struct {
	Contract *ERC20GatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Gateway creates a new instance of ERC20Gateway, bound to a specific deployed contract.
func NewERC20Gateway(address common.Address, backend bind.ContractBackend) (*ERC20Gateway, error) {
	contract, err := bindERC20Gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Gateway{ERC20GatewayCaller: ERC20GatewayCaller{contract: contract}, ERC20GatewayTransactor: ERC20GatewayTransactor{contract: contract}, ERC20GatewayFilterer: ERC20GatewayFilterer{contract: contract}}, nil
}

// NewERC20GatewayCaller creates a new read-only instance of ERC20Gateway, bound to a specific deployed contract.
func NewERC20GatewayCaller(address common.Address, caller bind.ContractCaller) (*ERC20GatewayCaller, error) {
	contract, err := bindERC20Gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20GatewayCaller{contract: contract}, nil
}

// NewERC20GatewayTransactor creates a new write-only instance of ERC20Gateway, bound to a specific deployed contract.
func NewERC20GatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20GatewayTransactor, error) {
	contract, err := bindERC20Gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20GatewayTransactor{contract: contract}, nil
}

// NewERC20GatewayFilterer creates a new log filterer instance of ERC20Gateway, bound to a specific deployed contract.
func NewERC20GatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20GatewayFilterer, error) {
	contract, err := bindERC20Gateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20GatewayFilterer{contract: contract}, nil
}

// bindERC20Gateway binds a generic wrapper to an already deployed contract.
func bindERC20Gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20GatewayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Gateway *ERC20GatewayRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Gateway.Contract.ERC20GatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Gateway *ERC20GatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Gateway.Contract.ERC20GatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Gateway *ERC20GatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Gateway.Contract.ERC20GatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Gateway *ERC20GatewayCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Gateway *ERC20GatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Gateway *ERC20GatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Gateway.Contract.contract.Transact(opts, method, params...)
}

// GetERC20 is a free data retrieval call binding the contract method 0x4e0dc557.
//
// Solidity: function getERC20(contractAddress address) constant returns(uint256)
func (_ERC20Gateway *ERC20GatewayCaller) GetERC20(opts *bind.CallOpts, contractAddress common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Gateway.contract.Call(opts, out, "getERC20", contractAddress)
	return *ret0, err
}

// GetERC20 is a free data retrieval call binding the contract method 0x4e0dc557.
//
// Solidity: function getERC20(contractAddress address) constant returns(uint256)
func (_ERC20Gateway *ERC20GatewaySession) GetERC20(contractAddress common.Address) (*big.Int, error) {
	return _ERC20Gateway.Contract.GetERC20(&_ERC20Gateway.CallOpts, contractAddress)
}

// GetERC20 is a free data retrieval call binding the contract method 0x4e0dc557.
//
// Solidity: function getERC20(contractAddress address) constant returns(uint256)
func (_ERC20Gateway *ERC20GatewayCallerSession) GetERC20(contractAddress common.Address) (*big.Int, error) {
	return _ERC20Gateway.Contract.GetERC20(&_ERC20Gateway.CallOpts, contractAddress)
}

// LoomAddress is a free data retrieval call binding the contract method 0x37179db8.
//
// Solidity: function loomAddress() constant returns(address)
func (_ERC20Gateway *ERC20GatewayCaller) LoomAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Gateway.contract.Call(opts, out, "loomAddress")
	return *ret0, err
}

// LoomAddress is a free data retrieval call binding the contract method 0x37179db8.
//
// Solidity: function loomAddress() constant returns(address)
func (_ERC20Gateway *ERC20GatewaySession) LoomAddress() (common.Address, error) {
	return _ERC20Gateway.Contract.LoomAddress(&_ERC20Gateway.CallOpts)
}

// LoomAddress is a free data retrieval call binding the contract method 0x37179db8.
//
// Solidity: function loomAddress() constant returns(address)
func (_ERC20Gateway *ERC20GatewayCallerSession) LoomAddress() (common.Address, error) {
	return _ERC20Gateway.Contract.LoomAddress(&_ERC20Gateway.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces( address) constant returns(uint256)
func (_ERC20Gateway *ERC20GatewayCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Gateway.contract.Call(opts, out, "nonces", arg0)
	return *ret0, err
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces( address) constant returns(uint256)
func (_ERC20Gateway *ERC20GatewaySession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _ERC20Gateway.Contract.Nonces(&_ERC20Gateway.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces( address) constant returns(uint256)
func (_ERC20Gateway *ERC20GatewayCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _ERC20Gateway.Contract.Nonces(&_ERC20Gateway.CallOpts, arg0)
}

// Vmc is a free data retrieval call binding the contract method 0x20cc8e51.
//
// Solidity: function vmc() constant returns(address)
func (_ERC20Gateway *ERC20GatewayCaller) Vmc(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Gateway.contract.Call(opts, out, "vmc")
	return *ret0, err
}

// Vmc is a free data retrieval call binding the contract method 0x20cc8e51.
//
// Solidity: function vmc() constant returns(address)
func (_ERC20Gateway *ERC20GatewaySession) Vmc() (common.Address, error) {
	return _ERC20Gateway.Contract.Vmc(&_ERC20Gateway.CallOpts)
}

// Vmc is a free data retrieval call binding the contract method 0x20cc8e51.
//
// Solidity: function vmc() constant returns(address)
func (_ERC20Gateway *ERC20GatewayCallerSession) Vmc() (common.Address, error) {
	return _ERC20Gateway.Contract.Vmc(&_ERC20Gateway.CallOpts)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x392d661c.
//
// Solidity: function depositERC20(amount uint256, contractAddress address) returns()
func (_ERC20Gateway *ERC20GatewayTransactor) DepositERC20(opts *bind.TransactOpts, amount *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Gateway.contract.Transact(opts, "depositERC20", amount, contractAddress)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x392d661c.
//
// Solidity: function depositERC20(amount uint256, contractAddress address) returns()
func (_ERC20Gateway *ERC20GatewaySession) DepositERC20(amount *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Gateway.Contract.DepositERC20(&_ERC20Gateway.TransactOpts, amount, contractAddress)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x392d661c.
//
// Solidity: function depositERC20(amount uint256, contractAddress address) returns()
func (_ERC20Gateway *ERC20GatewayTransactorSession) DepositERC20(amount *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Gateway.Contract.DepositERC20(&_ERC20Gateway.TransactOpts, amount, contractAddress)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xb0116dc7.
//
// Solidity: function withdrawERC20(amount uint256, contractAddress address, _signersIndexes uint256[], _v uint8[], _r bytes32[], _s bytes32[]) returns()
func (_ERC20Gateway *ERC20GatewayTransactor) WithdrawERC20(opts *bind.TransactOpts, amount *big.Int, contractAddress common.Address, _signersIndexes []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _ERC20Gateway.contract.Transact(opts, "withdrawERC20", amount, contractAddress, _signersIndexes, _v, _r, _s)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xb0116dc7.
//
// Solidity: function withdrawERC20(amount uint256, contractAddress address, _signersIndexes uint256[], _v uint8[], _r bytes32[], _s bytes32[]) returns()
func (_ERC20Gateway *ERC20GatewaySession) WithdrawERC20(amount *big.Int, contractAddress common.Address, _signersIndexes []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _ERC20Gateway.Contract.WithdrawERC20(&_ERC20Gateway.TransactOpts, amount, contractAddress, _signersIndexes, _v, _r, _s)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xb0116dc7.
//
// Solidity: function withdrawERC20(amount uint256, contractAddress address, _signersIndexes uint256[], _v uint8[], _r bytes32[], _s bytes32[]) returns()
func (_ERC20Gateway *ERC20GatewayTransactorSession) WithdrawERC20(amount *big.Int, contractAddress common.Address, _signersIndexes []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _ERC20Gateway.Contract.WithdrawERC20(&_ERC20Gateway.TransactOpts, amount, contractAddress, _signersIndexes, _v, _r, _s)
}

// ERC20GatewayERC20ReceivedIterator is returned from FilterERC20Received and is used to iterate over the raw logs and unpacked data for ERC20Received events raised by the ERC20Gateway contract.
type ERC20GatewayERC20ReceivedIterator struct {
	Event *ERC20GatewayERC20Received // Event containing the contract specifics and raw log

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
func (it *ERC20GatewayERC20ReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20GatewayERC20Received)
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
		it.Event = new(ERC20GatewayERC20Received)
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
func (it *ERC20GatewayERC20ReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20GatewayERC20ReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20GatewayERC20Received represents a ERC20Received event raised by the ERC20Gateway contract.
type ERC20GatewayERC20Received struct {
	From            common.Address
	Amount          *big.Int
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterERC20Received is a free log retrieval operation binding the contract event 0xa13cf347fb36122550e414f6fd1a0c2e490cff76331c4dcc20f760891ecca12a.
//
// Solidity: e ERC20Received(from address, amount uint256, contractAddress address)
func (_ERC20Gateway *ERC20GatewayFilterer) FilterERC20Received(opts *bind.FilterOpts) (*ERC20GatewayERC20ReceivedIterator, error) {

	logs, sub, err := _ERC20Gateway.contract.FilterLogs(opts, "ERC20Received")
	if err != nil {
		return nil, err
	}
	return &ERC20GatewayERC20ReceivedIterator{contract: _ERC20Gateway.contract, event: "ERC20Received", logs: logs, sub: sub}, nil
}

// WatchERC20Received is a free log subscription operation binding the contract event 0xa13cf347fb36122550e414f6fd1a0c2e490cff76331c4dcc20f760891ecca12a.
//
// Solidity: e ERC20Received(from address, amount uint256, contractAddress address)
func (_ERC20Gateway *ERC20GatewayFilterer) WatchERC20Received(opts *bind.WatchOpts, sink chan<- *ERC20GatewayERC20Received) (event.Subscription, error) {

	logs, sub, err := _ERC20Gateway.contract.WatchLogs(opts, "ERC20Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20GatewayERC20Received)
				if err := _ERC20Gateway.contract.UnpackLog(event, "ERC20Received", log); err != nil {
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

// ERC20GatewayLoomCoinReceivedIterator is returned from FilterLoomCoinReceived and is used to iterate over the raw logs and unpacked data for LoomCoinReceived events raised by the ERC20Gateway contract.
type ERC20GatewayLoomCoinReceivedIterator struct {
	Event *ERC20GatewayLoomCoinReceived // Event containing the contract specifics and raw log

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
func (it *ERC20GatewayLoomCoinReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20GatewayLoomCoinReceived)
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
		it.Event = new(ERC20GatewayLoomCoinReceived)
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
func (it *ERC20GatewayLoomCoinReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20GatewayLoomCoinReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20GatewayLoomCoinReceived represents a LoomCoinReceived event raised by the ERC20Gateway contract.
type ERC20GatewayLoomCoinReceived struct {
	From            common.Address
	Amount          *big.Int
	LoomCoinAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLoomCoinReceived is a free log retrieval operation binding the contract event 0x91557346f7592c9279b67cc52709a00442f0597658ec38a5fe84568c016331d7.
//
// Solidity: e LoomCoinReceived(from indexed address, amount uint256, loomCoinAddress address)
func (_ERC20Gateway *ERC20GatewayFilterer) FilterLoomCoinReceived(opts *bind.FilterOpts, from []common.Address) (*ERC20GatewayLoomCoinReceivedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ERC20Gateway.contract.FilterLogs(opts, "LoomCoinReceived", fromRule)
	if err != nil {
		return nil, err
	}
	return &ERC20GatewayLoomCoinReceivedIterator{contract: _ERC20Gateway.contract, event: "LoomCoinReceived", logs: logs, sub: sub}, nil
}

// WatchLoomCoinReceived is a free log subscription operation binding the contract event 0x91557346f7592c9279b67cc52709a00442f0597658ec38a5fe84568c016331d7.
//
// Solidity: e LoomCoinReceived(from indexed address, amount uint256, loomCoinAddress address)
func (_ERC20Gateway *ERC20GatewayFilterer) WatchLoomCoinReceived(opts *bind.WatchOpts, sink chan<- *ERC20GatewayLoomCoinReceived, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ERC20Gateway.contract.WatchLogs(opts, "LoomCoinReceived", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20GatewayLoomCoinReceived)
				if err := _ERC20Gateway.contract.UnpackLog(event, "LoomCoinReceived", log); err != nil {
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

// ERC20GatewayTokenWithdrawnIterator is returned from FilterTokenWithdrawn and is used to iterate over the raw logs and unpacked data for TokenWithdrawn events raised by the ERC20Gateway contract.
type ERC20GatewayTokenWithdrawnIterator struct {
	Event *ERC20GatewayTokenWithdrawn // Event containing the contract specifics and raw log

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
func (it *ERC20GatewayTokenWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20GatewayTokenWithdrawn)
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
		it.Event = new(ERC20GatewayTokenWithdrawn)
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
func (it *ERC20GatewayTokenWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20GatewayTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20GatewayTokenWithdrawn represents a TokenWithdrawn event raised by the ERC20Gateway contract.
type ERC20GatewayTokenWithdrawn struct {
	Owner           common.Address
	Kind            uint8
	ContractAddress common.Address
	Value           *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdrawn is a free log retrieval operation binding the contract event 0x591f2d33d85291e32c4067b5a497caf3ddb5b1830eba9909e66006ec3a0051b4.
//
// Solidity: e TokenWithdrawn(owner indexed address, kind uint8, contractAddress address, value uint256)
func (_ERC20Gateway *ERC20GatewayFilterer) FilterTokenWithdrawn(opts *bind.FilterOpts, owner []common.Address) (*ERC20GatewayTokenWithdrawnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ERC20Gateway.contract.FilterLogs(opts, "TokenWithdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20GatewayTokenWithdrawnIterator{contract: _ERC20Gateway.contract, event: "TokenWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokenWithdrawn is a free log subscription operation binding the contract event 0x591f2d33d85291e32c4067b5a497caf3ddb5b1830eba9909e66006ec3a0051b4.
//
// Solidity: e TokenWithdrawn(owner indexed address, kind uint8, contractAddress address, value uint256)
func (_ERC20Gateway *ERC20GatewayFilterer) WatchTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *ERC20GatewayTokenWithdrawn, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ERC20Gateway.contract.WatchLogs(opts, "TokenWithdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20GatewayTokenWithdrawn)
				if err := _ERC20Gateway.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
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
