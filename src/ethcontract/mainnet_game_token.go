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

// MainnetGameTokenContractABI is the input ABI used to generate the binding from.
const MainnetGameTokenContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x06fdde03\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x095ea7b3\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x18160ddd\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x23b872dd\"},{\"constant\":true,\"inputs\":[],\"name\":\"INITIAL_SUPPLY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x2ff2e9dc\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x313ce567\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x66188463\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x70a08231\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x95d89b41\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xa9059cbb\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xd73dd623\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xdd62ed3e\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"signature\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\",\"signature\":\"0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\",\"signature\":\"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\"}]"

// MainnetGameTokenContractBin is the compiled bytecode used for deploying new contracts.
const MainnetGameTokenContractBin = `0x60c0604052600960808190527f47616d65546f6b656e000000000000000000000000000000000000000000000060a090815261003e91600391906100c7565b506040805180820190915260038082527f47544b00000000000000000000000000000000000000000000000000000000006020909201918252610083916004916100c7565b506005805460ff1916601217905534801561009d57600080fd5b5060055460ff16600a0a633b9aca0002600181905533600090815260208190526040902055610162565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061010857805160ff1916838001178555610135565b82800160010185558215610135579182015b8281111561013557825182559160200191906001019061011a565b50610141929150610145565b5090565b61015f91905b80821115610141576000815560010161014b565b90565b6108b3806101716000396000f3006080604052600436106100b95763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100be578063095ea7b31461014857806318160ddd1461018057806323b872dd146101a75780632ff2e9dc146101d1578063313ce567146101e6578063661884631461021157806370a082311461023557806395d89b4114610256578063a9059cbb1461026b578063d73dd6231461028f578063dd62ed3e146102b3575b600080fd5b3480156100ca57600080fd5b506100d36102da565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561010d5781810151838201526020016100f5565b50505050905090810190601f16801561013a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561015457600080fd5b5061016c600160a060020a0360043516602435610368565b604080519115158252519081900360200190f35b34801561018c57600080fd5b506101956103ce565b60408051918252519081900360200190f35b3480156101b357600080fd5b5061016c600160a060020a03600435811690602435166044356103d4565b3480156101dd57600080fd5b50610195610549565b3480156101f257600080fd5b506101fb610551565b6040805160ff9092168252519081900360200190f35b34801561021d57600080fd5b5061016c600160a060020a036004351660243561055a565b34801561024157600080fd5b50610195600160a060020a0360043516610649565b34801561026257600080fd5b506100d3610664565b34801561027757600080fd5b5061016c600160a060020a03600435166024356106bf565b34801561029b57600080fd5b5061016c600160a060020a036004351660243561079e565b3480156102bf57600080fd5b50610195600160a060020a0360043581169060243516610837565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103605780601f1061033557610100808354040283529160200191610360565b820191906000526020600020905b81548152906001019060200180831161034357829003601f168201915b505050505081565b336000818152600260209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60015490565b600160a060020a0383166000908152602081905260408120548211156103f957600080fd5b600160a060020a038416600090815260026020908152604080832033845290915290205482111561042957600080fd5b600160a060020a038316151561043e57600080fd5b600160a060020a038416600090815260208190526040902054610467908363ffffffff61086216565b600160a060020a03808616600090815260208190526040808220939093559085168152205461049c908363ffffffff61087416565b600160a060020a038085166000908152602081815260408083209490945591871681526002825282812033825290915220546104de908363ffffffff61086216565b600160a060020a03808616600081815260026020908152604080832033845282529182902094909455805186815290519287169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a35060019392505050565b633b9aca0081565b60055460ff1681565b336000908152600260209081526040808320600160a060020a03861684529091528120548083106105ae57336000908152600260209081526040808320600160a060020a03881684529091528120556105e3565b6105be818463ffffffff61086216565b336000908152600260209081526040808320600160a060020a03891684529091529020555b336000818152600260209081526040808320600160a060020a0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b600160a060020a031660009081526020819052604090205490565b6004805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103605780601f1061033557610100808354040283529160200191610360565b336000908152602081905260408120548211156106db57600080fd5b600160a060020a03831615156106f057600080fd5b33600090815260208190526040902054610710908363ffffffff61086216565b3360009081526020819052604080822092909255600160a060020a03851681522054610742908363ffffffff61087416565b600160a060020a038416600081815260208181526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b336000908152600260209081526040808320600160a060020a03861684529091528120546107d2908363ffffffff61087416565b336000818152600260209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60008282111561086e57fe5b50900390565b8181018281101561088157fe5b929150505600a165627a7a7230582085272da7062b31ba32a32cc3da71972fca5b60ac06cc9db479732914fe0b15d90029`

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
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_MainnetGameTokenContract *MainnetGameTokenContractCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MainnetGameTokenContract.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_MainnetGameTokenContract *MainnetGameTokenContractSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _MainnetGameTokenContract.Contract.Allowance(&_MainnetGameTokenContract.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_MainnetGameTokenContract *MainnetGameTokenContractCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _MainnetGameTokenContract.Contract.Allowance(&_MainnetGameTokenContract.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_MainnetGameTokenContract *MainnetGameTokenContractCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MainnetGameTokenContract.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_MainnetGameTokenContract *MainnetGameTokenContractSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MainnetGameTokenContract.Contract.BalanceOf(&_MainnetGameTokenContract.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_MainnetGameTokenContract *MainnetGameTokenContractCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MainnetGameTokenContract.Contract.BalanceOf(&_MainnetGameTokenContract.CallOpts, _owner)
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
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.Contract.Approve(&_MainnetGameTokenContract.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.Contract.Approve(&_MainnetGameTokenContract.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.Contract.DecreaseApproval(&_MainnetGameTokenContract.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.Contract.DecreaseApproval(&_MainnetGameTokenContract.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.Contract.IncreaseApproval(&_MainnetGameTokenContract.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.Contract.IncreaseApproval(&_MainnetGameTokenContract.TransactOpts, _spender, _addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.Contract.Transfer(&_MainnetGameTokenContract.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.Contract.Transfer(&_MainnetGameTokenContract.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.Contract.TransferFrom(&_MainnetGameTokenContract.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_MainnetGameTokenContract *MainnetGameTokenContractTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MainnetGameTokenContract.Contract.TransferFrom(&_MainnetGameTokenContract.TransactOpts, _from, _to, _value)
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
