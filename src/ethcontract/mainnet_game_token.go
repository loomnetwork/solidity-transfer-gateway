// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethcontract

import (
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"math/big"
	"strings"
)

// MainnetGameTokenContractABI is the input ABI used to generate the binding from.
const MainnetGameTokenContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INITIAL_SUPPLY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// MainnetGameTokenContractBin is the compiled bytecode used for deploying new contracts.
const MainnetGameTokenContractBin = `0x60c0604052600960808190527f47616d65546f6b656e000000000000000000000000000000000000000000000060a090815261003e916003919061018f565b506040805180820190915260038082527f47544b000000000000000000000000000000000000000000000000000000000060209092019182526100839160049161018f565b506005805460ff1916601217905534801561009d57600080fd5b506005546100bf90339060ff16600a0a633b9aca00026100c4602090811b901c565b61022a565b6001600160a01b0382166100d757600080fd5b6100f08160025461017660201b61066c1790919060201c565b6002556001600160a01b0382166000908152602081815260409091205461012091839061066c610176821b17901c565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b60008282018381101561018857600080fd5b9392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106101d057805160ff19168380011785556101fd565b828001600101855582156101fd579182015b828111156101fd5782518255916020019190600101906101e2565b5061020992915061020d565b5090565b61022791905b808211156102095760008155600101610213565b90565b6106b1806102396000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063395093511161007157806339509351146101ec57806370a082311461021857806395d89b411461023e578063a457c2d714610246578063a9059cbb14610272578063dd62ed3e1461029e576100b4565b806306fdde03146100b9578063095ea7b31461013657806318160ddd1461017657806323b872dd146101905780632ff2e9dc146101c6578063313ce567146101ce575b600080fd5b6100c16102cc565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100fb5781810151838201526020016100e3565b50505050905090810190601f1680156101285780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101626004803603604081101561014c57600080fd5b506001600160a01b03813516906020013561035a565b604080519115158252519081900360200190f35b61017e610370565b60408051918252519081900360200190f35b610162600480360360608110156101a657600080fd5b506001600160a01b03813581169160208101359091169060400135610376565b61017e6103cd565b6101d66103d5565b6040805160ff9092168252519081900360200190f35b6101626004803603604081101561020257600080fd5b506001600160a01b0381351690602001356103de565b61017e6004803603602081101561022e57600080fd5b50356001600160a01b031661041a565b6100c1610435565b6101626004803603604081101561025c57600080fd5b506001600160a01b038135169060200135610490565b6101626004803603604081101561028857600080fd5b506001600160a01b0381351690602001356104cc565b61017e600480360360408110156102b457600080fd5b506001600160a01b03813581169160200135166104d9565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103525780601f1061032757610100808354040283529160200191610352565b820191906000526020600020905b81548152906001019060200180831161033557829003601f168201915b505050505081565b6000610367338484610504565b50600192915050565b60025490565b600061038384848461058c565b6001600160a01b0384166000908152600160209081526040808320338085529252909120546103c39186916103be908663ffffffff61065716565b610504565b5060019392505050565b633b9aca0081565b60055460ff1681565b3360008181526001602090815260408083206001600160a01b038716845290915281205490916103679185906103be908663ffffffff61066c16565b6001600160a01b031660009081526020819052604090205490565b6004805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103525780601f1061032757610100808354040283529160200191610352565b3360008181526001602090815260408083206001600160a01b038716845290915281205490916103679185906103be908663ffffffff61065716565b600061036733848461058c565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b03821661051757600080fd5b6001600160a01b03831661052a57600080fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b03821661059f57600080fd5b6001600160a01b0383166000908152602081905260409020546105c8908263ffffffff61065716565b6001600160a01b0380851660009081526020819052604080822093909355908416815220546105fd908263ffffffff61066c16565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b60008282111561066657600080fd5b50900390565b60008282018381101561067e57600080fd5b939250505056fea165627a7a723058203c69309bb5cf116a1e4592bbd99d3eb54a5c2d8f0381f77dce0eb8db261d56440029`

// DeployMainnetGameTokenContract deploys a new Ethereum contract, binding an instance of MainnetGameTokenContract to it.
func DeployMainnetGameTokenContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MainnetGameTokenContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainnetGameTokenContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MainnetGameTokenContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MainnetGameTokenContract{MainnetGameTokenContractCaller: MainnetGameTokenContractCaller{contract: contract}, MainnetGameTokenContractTransactor: MainnetGameTokenContractTransactor{contract: contract}, MainnetGameTokenContractFilterer: MainnetGameTokenContractFilterer{contract: contract}}, nil
}

// MainnetGameTokenContract is an auto generated Go binding around an Ethereum contract.
type MainnetGameTokenContract struct {
	MainnetGameTokenContractCaller     // Read-only binding to the contract
	MainnetGameTokenContractTransactor // Write-only binding to the contract
	MainnetGameTokenContractFilterer   // Log filterer for contract events
}

// MainnetGameTokenContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainnetGameTokenContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainnetGameTokenContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainnetGameTokenContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainnetGameTokenContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainnetGameTokenContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainnetGameTokenContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainnetGameTokenContractSession struct {
	Contract     *MainnetGameTokenContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MainnetGameTokenContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainnetGameTokenContractCallerSession struct {
	Contract *MainnetGameTokenContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// MainnetGameTokenContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainnetGameTokenContractTransactorSession struct {
	Contract     *MainnetGameTokenContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// MainnetGameTokenContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainnetGameTokenContractRaw struct {
	Contract *MainnetGameTokenContract // Generic contract binding to access the raw methods on
}

// MainnetGameTokenContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainnetGameTokenContractCallerRaw struct {
	Contract *MainnetGameTokenContractCaller // Generic read-only contract binding to access the raw methods on
}

// MainnetGameTokenContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainnetGameTokenContractTransactorRaw struct {
	Contract *MainnetGameTokenContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMainnetGameTokenContract creates a new instance of MainnetGameTokenContract, bound to a specific deployed contract.
func NewMainnetGameTokenContract(address common.Address, backend bind.ContractBackend) (*MainnetGameTokenContract, error) {
	contract, err := bindMainnetGameTokenContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MainnetGameTokenContract{MainnetGameTokenContractCaller: MainnetGameTokenContractCaller{contract: contract}, MainnetGameTokenContractTransactor: MainnetGameTokenContractTransactor{contract: contract}, MainnetGameTokenContractFilterer: MainnetGameTokenContractFilterer{contract: contract}}, nil
}

// NewMainnetGameTokenContractCaller creates a new read-only instance of MainnetGameTokenContract, bound to a specific deployed contract.
func NewMainnetGameTokenContractCaller(address common.Address, caller bind.ContractCaller) (*MainnetGameTokenContractCaller, error) {
	contract, err := bindMainnetGameTokenContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainnetGameTokenContractCaller{contract: contract}, nil
}

// NewMainnetGameTokenContractTransactor creates a new write-only instance of MainnetGameTokenContract, bound to a specific deployed contract.
func NewMainnetGameTokenContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MainnetGameTokenContractTransactor, error) {
	contract, err := bindMainnetGameTokenContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainnetGameTokenContractTransactor{contract: contract}, nil
}

// NewMainnetGameTokenContractFilterer creates a new log filterer instance of MainnetGameTokenContract, bound to a specific deployed contract.
func NewMainnetGameTokenContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MainnetGameTokenContractFilterer, error) {
	contract, err := bindMainnetGameTokenContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainnetGameTokenContractFilterer{contract: contract}, nil
}

// bindMainnetGameTokenContract binds a generic wrapper to an already deployed contract.
func bindMainnetGameTokenContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainnetGameTokenContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MainnetGameTokenContract *MainnetGameTokenContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MainnetGameTokenContract.Contract.MainnetGameTokenContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MainnetGameTokenContract *MainnetGameTokenContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainnetGameTokenContract.Contract.MainnetGameTokenContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MainnetGameTokenContract *MainnetGameTokenContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MainnetGameTokenContract.Contract.MainnetGameTokenContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MainnetGameTokenContract *MainnetGameTokenContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MainnetGameTokenContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MainnetGameTokenContract *MainnetGameTokenContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainnetGameTokenContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MainnetGameTokenContract *MainnetGameTokenContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MainnetGameTokenContract.Contract.contract.Transact(opts, method, params...)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_MainnetGameTokenContract *MainnetGameTokenContractCaller) INITIALSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MainnetGameTokenContract.contract.Call(opts, out, "INITIAL_SUPPLY")
	return *ret0, err
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_MainnetGameTokenContract *MainnetGameTokenContractSession) INITIALSUPPLY() (*big.Int, error) {
	return _MainnetGameTokenContract.Contract.INITIALSUPPLY(&_MainnetGameTokenContract.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_MainnetGameTokenContract *MainnetGameTokenContractCallerSession) INITIALSUPPLY() (*big.Int, error) {
	return _MainnetGameTokenContract.Contract.INITIALSUPPLY(&_MainnetGameTokenContract.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_MainnetGameTokenContract *MainnetGameTokenContractCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MainnetGameTokenContract.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_MainnetGameTokenContract *MainnetGameTokenContractSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MainnetGameTokenContract.Contract.Allowance(&_MainnetGameTokenContract.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_MainnetGameTokenContract *MainnetGameTokenContractCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MainnetGameTokenContract.Contract.Allowance(&_MainnetGameTokenContract.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_MainnetGameTokenContract *MainnetGameTokenContractCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MainnetGameTokenContract.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_MainnetGameTokenContract *MainnetGameTokenContractSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MainnetGameTokenContract.Contract.BalanceOf(&_MainnetGameTokenContract.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_MainnetGameTokenContract *MainnetGameTokenContractCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MainnetGameTokenContract.Contract.BalanceOf(&_MainnetGameTokenContract.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MainnetGameTokenContract *MainnetGameTokenContractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MainnetGameTokenContract.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MainnetGameTokenContract *MainnetGameTokenContractSession) Decimals() (uint8, error) {
	return _MainnetGameTokenContract.Contract.Decimals(&_MainnetGameTokenContract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MainnetGameTokenContract *MainnetGameTokenContractCallerSession) Decimals() (uint8, error) {
	return _MainnetGameTokenContract.Contract.Decimals(&_MainnetGameTokenContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MainnetGameTokenContract *MainnetGameTokenContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MainnetGameTokenContract.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MainnetGameTokenContract *MainnetGameTokenContractSession) Name() (string, error) {
	return _MainnetGameTokenContract.Contract.Name(&_MainnetGameTokenContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MainnetGameTokenContract *MainnetGameTokenContractCallerSession) Name() (string, error) {
	return _MainnetGameTokenContract.Contract.Name(&_MainnetGameTokenContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MainnetGameTokenContract *MainnetGameTokenContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MainnetGameTokenContract.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MainnetGameTokenContract *MainnetGameTokenContractSession) Symbol() (string, error) {
	return _MainnetGameTokenContract.Contract.Symbol(&_MainnetGameTokenContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MainnetGameTokenContract *MainnetGameTokenContractCallerSession) Symbol() (string, error) {
	return _MainnetGameTokenContract.Contract.Symbol(&_MainnetGameTokenContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MainnetGameTokenContract *MainnetGameTokenContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MainnetGameTokenContract.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MainnetGameTokenContract *MainnetGameTokenContractSession) TotalSupply() (*big.Int, error) {
	return _MainnetGameTokenContract.Contract.TotalSupply(&_MainnetGameTokenContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MainnetGameTokenContract *MainnetGameTokenContractCallerSession) TotalSupply() (*big.Int, error) {
	return _MainnetGameTokenContract.Contract.TotalSupply(&_MainnetGameTokenContract.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.Contract.Approve(&_MainnetGameTokenContract.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.Contract.Approve(&_MainnetGameTokenContract.TransactOpts, spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(spender address, subtractedValue uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(spender address, subtractedValue uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.Contract.DecreaseAllowance(&_MainnetGameTokenContract.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(spender address, subtractedValue uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.Contract.DecreaseAllowance(&_MainnetGameTokenContract.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(spender address, addedValue uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(spender address, addedValue uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.Contract.IncreaseAllowance(&_MainnetGameTokenContract.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(spender address, addedValue uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.Contract.IncreaseAllowance(&_MainnetGameTokenContract.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.Contract.Transfer(&_MainnetGameTokenContract.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.Contract.Transfer(&_MainnetGameTokenContract.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.Contract.TransferFrom(&_MainnetGameTokenContract.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.Contract.TransferFrom(&_MainnetGameTokenContract.TransactOpts, from, to, value)
}

// MainnetGameTokenContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MainnetGameTokenContract contract.
type MainnetGameTokenContractApprovalIterator struct {
	Event *MainnetGameTokenContractApproval // Event containing the contract specifics and raw log

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
func (it *MainnetGameTokenContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainnetGameTokenContractApproval)
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
		it.Event = new(MainnetGameTokenContractApproval)
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
func (it *MainnetGameTokenContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainnetGameTokenContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainnetGameTokenContractApproval represents a Approval event raised by the MainnetGameTokenContract contract.
type MainnetGameTokenContractApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_MainnetGameTokenContract *MainnetGameTokenContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MainnetGameTokenContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MainnetGameTokenContract.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MainnetGameTokenContractApprovalIterator{contract: _MainnetGameTokenContract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_MainnetGameTokenContract *MainnetGameTokenContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MainnetGameTokenContractApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MainnetGameTokenContract.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainnetGameTokenContractApproval)
				if err := _MainnetGameTokenContract.contract.UnpackLog(event, "Approval", log); err != nil {
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

// MainnetGameTokenContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MainnetGameTokenContract contract.
type MainnetGameTokenContractTransferIterator struct {
	Event *MainnetGameTokenContractTransfer // Event containing the contract specifics and raw log

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
func (it *MainnetGameTokenContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainnetGameTokenContractTransfer)
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
		it.Event = new(MainnetGameTokenContractTransfer)
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
func (it *MainnetGameTokenContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainnetGameTokenContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainnetGameTokenContractTransfer represents a Transfer event raised by the MainnetGameTokenContract contract.
type MainnetGameTokenContractTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_MainnetGameTokenContract *MainnetGameTokenContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MainnetGameTokenContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MainnetGameTokenContract.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MainnetGameTokenContractTransferIterator{contract: _MainnetGameTokenContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_MainnetGameTokenContract *MainnetGameTokenContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MainnetGameTokenContractTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MainnetGameTokenContract.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainnetGameTokenContractTransfer)
				if err := _MainnetGameTokenContract.contract.UnpackLog(event, "Transfer", log); err != nil {
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
