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

// SampleERC20MintableTokenABI is the input ABI used to generate the binding from.
const SampleERC20MintableTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_gateway\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"gateway\",\"type\":\"address\"}],\"name\":\"GatewayAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"gateway\",\"type\":\"address\"}],\"name\":\"GatewayRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mintTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newValidator\",\"type\":\"address\"}],\"name\":\"addValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_gateway\",\"type\":\"address\"}],\"name\":\"addGateway\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_gateway\",\"type\":\"address\"}],\"name\":\"removeGateway\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"gateway\",\"type\":\"address\"}],\"name\":\"isGateway\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SampleERC20MintableTokenBin is the compiled bytecode used for deploying new contracts.
const SampleERC20MintableTokenBin = `0x60806040523480156200001157600080fd5b50604051602080620011af833981018060405260208110156200003357600080fd5b5051620000473362000117602090811b901c565b6001600160a01b03811660009081526004602090815260408083208054600160ff19918216811790925533855260078452938290208054909416179092558151808301909252600d8083527f65726332306d696e7461626c650000000000000000000000000000000000000092909101918252620000c891600591620001f5565b506040805180820190915260058082527f4d4e54323000000000000000000000000000000000000000000000000000000060209092019182526200010f91600691620001f5565b50506200029a565b620001328160036200016960201b62000e141790919060201c565b6040516001600160a01b038216907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a250565b6001600160a01b0381166200017d57600080fd5b6200018f8282620001bf60201b60201c565b156200019a57600080fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b60006001600160a01b038216620001d557600080fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200023857805160ff191683800117855562000268565b8280016001018555821562000268579182015b82811115620002685782518255916020019190600101906200024b565b50620002769291506200027a565b5090565b6200029791905b8082111562000276576000815560010162000281565b90565b610f0580620002aa6000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c806370a08231116100b8578063a457c2d71161007c578063a457c2d7146103ec578063a9059cbb14610418578063aa271e1a14610444578063bfe69d9a1461046a578063dd62ed3e14610490578063facd743b146104be57610142565b806370a082311461036a5780638a885e351461039057806395d89b41146103b6578063983b2d56146103be57806398650275146103e457610142565b8063395093511161010a578063395093511461027257806340a141ff1461029e57806340c10f19146102c6578063449a52f8146102f25780634d238c8e1461031e57806368bb37951461034457610142565b806306fdde0314610147578063095ea7b3146101c457806318160ddd1461020457806323b872dd1461021e578063313ce56714610254575b600080fd5b61014f6104e4565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610189578181015183820152602001610171565b50505050905090810190601f1680156101b65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101f0600480360360408110156101da57600080fd5b506001600160a01b038135169060200135610572565b604080519115158252519081900360200190f35b61020c610588565b60408051918252519081900360200190f35b6101f06004803603606081101561023457600080fd5b506001600160a01b0381358116916020810135909116906040013561058e565b61025c6105e5565b6040805160ff9092168252519081900360200190f35b6101f06004803603604081101561028857600080fd5b506001600160a01b0381351690602001356105ea565b6102c4600480360360208110156102b457600080fd5b50356001600160a01b0316610626565b005b6101f0600480360360408110156102dc57600080fd5b506001600160a01b0381351690602001356106d4565b6102c46004803603604081101561030857600080fd5b506001600160a01b038135169060200135610734565b6102c46004803603602081101561033457600080fd5b50356001600160a01b03166107ae565b6102c46004803603602081101561035a57600080fd5b50356001600160a01b031661085f565b61020c6004803603602081101561038057600080fd5b50356001600160a01b0316610910565b6102c4600480360360208110156103a657600080fd5b50356001600160a01b031661092b565b61014f6109d9565b6102c4600480360360208110156103d457600080fd5b50356001600160a01b0316610a34565b6102c4610a52565b6101f06004803603604081101561040257600080fd5b506001600160a01b038135169060200135610a5d565b6101f06004803603604081101561042e57600080fd5b506001600160a01b038135169060200135610a99565b6101f06004803603602081101561045a57600080fd5b50356001600160a01b0316610aa6565b6101f06004803603602081101561048057600080fd5b50356001600160a01b0316610abf565b61020c600480360360408110156104a657600080fd5b506001600160a01b0381358116916020013516610add565b6101f0600480360360208110156104d457600080fd5b50356001600160a01b0316610b08565b6005805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561056a5780601f1061053f5761010080835404028352916020019161056a565b820191906000526020600020905b81548152906001019060200180831161054d57829003601f168201915b505050505081565b600061057f338484610b26565b50600192915050565b60025490565b600061059b848484610bae565b6001600160a01b0384166000908152600160209081526040808320338085529252909120546105db9186916105d6908663ffffffff610c7916565b610b26565b5060019392505050565b601281565b3360008181526001602090815260408083206001600160a01b0387168452909152812054909161057f9185906105d6908663ffffffff610c8e16565b3360009081526007602052604090205460ff16151560011461067c57604051600160e51b62461bcd028152600401808060200182810382526031815260200180610ea96031913960400191505060405180910390fd5b6001600160a01b038116600081815260076020908152604091829020805460ff19169055815192835290517fe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f19281900390910190a150565b3360009081526007602052604081205460ff16151560011461072a57604051600160e51b62461bcd028152600401808060200182810382526031815260200180610ea96031913960400191505060405180910390fd5b61057f8383610ca7565b3360009081526004602052604090205460ff1615156001146107a05760408051600160e51b62461bcd02815260206004820152601e60248201527f6f6e6c792067617465776179732061726520616c6c6f776564206d696e740000604482015290519081900360640190fd5b6107aa8282610ca7565b5050565b3360009081526007602052604090205460ff16151560011461080457604051600160e51b62461bcd028152600401808060200182810382526031815260200180610ea96031913960400191505060405180910390fd5b6001600160a01b038116600081815260076020908152604091829020805460ff19166001179055815192835290517fe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec38849879281900390910190a150565b3360009081526007602052604090205460ff1615156001146108b557604051600160e51b62461bcd028152600401808060200182810382526031815260200180610ea96031913960400191505060405180910390fd5b6001600160a01b038116600081815260046020908152604091829020805460ff19166001179055815192835290517f7137528d21fb7b0b9462886348954009edac49570e27ef9dba8bb3a676fc11e29281900390910190a150565b6001600160a01b031660009081526020819052604090205490565b3360009081526007602052604090205460ff16151560011461098157604051600160e51b62461bcd028152600401808060200182810382526031815260200180610ea96031913960400191505060405180910390fd5b6001600160a01b038116600081815260046020908152604091829020805460ff19169055815192835290517f6ad07cb5676ab639fa3821550f0ec4379900d542e2d8bd8f0f8158d8bd352da29281900390910190a150565b6006805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561056a5780601f1061053f5761010080835404028352916020019161056a565b610a3d33610aa6565b610a4657600080fd5b610a4f81610d4f565b50565b610a5b33610d97565b565b3360008181526001602090815260408083206001600160a01b0387168452909152812054909161057f9185906105d6908663ffffffff610c7916565b600061057f338484610bae565b6000610ab960038363ffffffff610ddf16565b92915050565b6001600160a01b031660009081526004602052604090205460ff1690565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b031660009081526007602052604090205460ff1690565b6001600160a01b038216610b3957600080fd5b6001600160a01b038316610b4c57600080fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b038216610bc157600080fd5b6001600160a01b038316600090815260208190526040902054610bea908263ffffffff610c7916565b6001600160a01b038085166000908152602081905260408082209390935590841681522054610c1f908263ffffffff610c8e16565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600082821115610c8857600080fd5b50900390565b600082820183811015610ca057600080fd5b9392505050565b6001600160a01b038216610cba57600080fd5b600254610ccd908263ffffffff610c8e16565b6002556001600160a01b038216600090815260208190526040902054610cf9908263ffffffff610c8e16565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b610d6060038263ffffffff610e1416565b6040516001600160a01b038216907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a250565b610da860038263ffffffff610e6016565b6040516001600160a01b038216907fe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb6669290600090a250565b60006001600160a01b038216610df457600080fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b6001600160a01b038116610e2757600080fd5b610e318282610ddf565b15610e3b57600080fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b6001600160a01b038116610e7357600080fd5b610e7d8282610ddf565b610e8657600080fd5b6001600160a01b0316600090815260209190915260409020805460ff1916905556fe6f6e6c792076616c696461746f727320617574686f72697a656420746f20706572666f726d207468697320616374696f6ea165627a7a7230582096e72567e3919b0a88fb4e1cce93964e445020a8624c4f16e26307d101843fcc0029`

// DeploySampleERC20MintableToken deploys a new Ethereum contract, binding an instance of SampleERC20MintableToken to it.
func DeploySampleERC20MintableToken(auth *bind.TransactOpts, backend bind.ContractBackend, _gateway common.Address) (common.Address, *types.Transaction, *SampleERC20MintableToken, error) {
	parsed, err := abi.JSON(strings.NewReader(SampleERC20MintableTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SampleERC20MintableTokenBin), backend, _gateway)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SampleERC20MintableToken{SampleERC20MintableTokenCaller: SampleERC20MintableTokenCaller{contract: contract}, SampleERC20MintableTokenTransactor: SampleERC20MintableTokenTransactor{contract: contract}, SampleERC20MintableTokenFilterer: SampleERC20MintableTokenFilterer{contract: contract}}, nil
}

// SampleERC20MintableToken is an auto generated Go binding around an Ethereum contract.
type SampleERC20MintableToken struct {
	SampleERC20MintableTokenCaller     // Read-only binding to the contract
	SampleERC20MintableTokenTransactor // Write-only binding to the contract
	SampleERC20MintableTokenFilterer   // Log filterer for contract events
}

// SampleERC20MintableTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SampleERC20MintableTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleERC20MintableTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SampleERC20MintableTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleERC20MintableTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SampleERC20MintableTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleERC20MintableTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SampleERC20MintableTokenSession struct {
	Contract     *SampleERC20MintableToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SampleERC20MintableTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SampleERC20MintableTokenCallerSession struct {
	Contract *SampleERC20MintableTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// SampleERC20MintableTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SampleERC20MintableTokenTransactorSession struct {
	Contract     *SampleERC20MintableTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// SampleERC20MintableTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type SampleERC20MintableTokenRaw struct {
	Contract *SampleERC20MintableToken // Generic contract binding to access the raw methods on
}

// SampleERC20MintableTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SampleERC20MintableTokenCallerRaw struct {
	Contract *SampleERC20MintableTokenCaller // Generic read-only contract binding to access the raw methods on
}

// SampleERC20MintableTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SampleERC20MintableTokenTransactorRaw struct {
	Contract *SampleERC20MintableTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSampleERC20MintableToken creates a new instance of SampleERC20MintableToken, bound to a specific deployed contract.
func NewSampleERC20MintableToken(address common.Address, backend bind.ContractBackend) (*SampleERC20MintableToken, error) {
	contract, err := bindSampleERC20MintableToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SampleERC20MintableToken{SampleERC20MintableTokenCaller: SampleERC20MintableTokenCaller{contract: contract}, SampleERC20MintableTokenTransactor: SampleERC20MintableTokenTransactor{contract: contract}, SampleERC20MintableTokenFilterer: SampleERC20MintableTokenFilterer{contract: contract}}, nil
}

// NewSampleERC20MintableTokenCaller creates a new read-only instance of SampleERC20MintableToken, bound to a specific deployed contract.
func NewSampleERC20MintableTokenCaller(address common.Address, caller bind.ContractCaller) (*SampleERC20MintableTokenCaller, error) {
	contract, err := bindSampleERC20MintableToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SampleERC20MintableTokenCaller{contract: contract}, nil
}

// NewSampleERC20MintableTokenTransactor creates a new write-only instance of SampleERC20MintableToken, bound to a specific deployed contract.
func NewSampleERC20MintableTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*SampleERC20MintableTokenTransactor, error) {
	contract, err := bindSampleERC20MintableToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SampleERC20MintableTokenTransactor{contract: contract}, nil
}

// NewSampleERC20MintableTokenFilterer creates a new log filterer instance of SampleERC20MintableToken, bound to a specific deployed contract.
func NewSampleERC20MintableTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*SampleERC20MintableTokenFilterer, error) {
	contract, err := bindSampleERC20MintableToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SampleERC20MintableTokenFilterer{contract: contract}, nil
}

// bindSampleERC20MintableToken binds a generic wrapper to an already deployed contract.
func bindSampleERC20MintableToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SampleERC20MintableTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SampleERC20MintableToken *SampleERC20MintableTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SampleERC20MintableToken.Contract.SampleERC20MintableTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SampleERC20MintableToken *SampleERC20MintableTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.SampleERC20MintableTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SampleERC20MintableToken *SampleERC20MintableTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.SampleERC20MintableTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SampleERC20MintableToken *SampleERC20MintableTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SampleERC20MintableToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_SampleERC20MintableToken *SampleERC20MintableTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleERC20MintableToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_SampleERC20MintableToken *SampleERC20MintableTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SampleERC20MintableToken.Contract.Allowance(&_SampleERC20MintableToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_SampleERC20MintableToken *SampleERC20MintableTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SampleERC20MintableToken.Contract.Allowance(&_SampleERC20MintableToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_SampleERC20MintableToken *SampleERC20MintableTokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleERC20MintableToken.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_SampleERC20MintableToken *SampleERC20MintableTokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SampleERC20MintableToken.Contract.BalanceOf(&_SampleERC20MintableToken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_SampleERC20MintableToken *SampleERC20MintableTokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SampleERC20MintableToken.Contract.BalanceOf(&_SampleERC20MintableToken.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SampleERC20MintableToken *SampleERC20MintableTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _SampleERC20MintableToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SampleERC20MintableToken *SampleERC20MintableTokenSession) Decimals() (uint8, error) {
	return _SampleERC20MintableToken.Contract.Decimals(&_SampleERC20MintableToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SampleERC20MintableToken *SampleERC20MintableTokenCallerSession) Decimals() (uint8, error) {
	return _SampleERC20MintableToken.Contract.Decimals(&_SampleERC20MintableToken.CallOpts)
}

// IsGateway is a free data retrieval call binding the contract method 0xbfe69d9a.
//
// Solidity: function isGateway(gateway address) constant returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenCaller) IsGateway(opts *bind.CallOpts, gateway common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SampleERC20MintableToken.contract.Call(opts, out, "isGateway", gateway)
	return *ret0, err
}

// IsGateway is a free data retrieval call binding the contract method 0xbfe69d9a.
//
// Solidity: function isGateway(gateway address) constant returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenSession) IsGateway(gateway common.Address) (bool, error) {
	return _SampleERC20MintableToken.Contract.IsGateway(&_SampleERC20MintableToken.CallOpts, gateway)
}

// IsGateway is a free data retrieval call binding the contract method 0xbfe69d9a.
//
// Solidity: function isGateway(gateway address) constant returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenCallerSession) IsGateway(gateway common.Address) (bool, error) {
	return _SampleERC20MintableToken.Contract.IsGateway(&_SampleERC20MintableToken.CallOpts, gateway)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(account address) constant returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SampleERC20MintableToken.contract.Call(opts, out, "isMinter", account)
	return *ret0, err
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(account address) constant returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenSession) IsMinter(account common.Address) (bool, error) {
	return _SampleERC20MintableToken.Contract.IsMinter(&_SampleERC20MintableToken.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(account address) constant returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenCallerSession) IsMinter(account common.Address) (bool, error) {
	return _SampleERC20MintableToken.Contract.IsMinter(&_SampleERC20MintableToken.CallOpts, account)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(validator address) constant returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenCaller) IsValidator(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SampleERC20MintableToken.contract.Call(opts, out, "isValidator", validator)
	return *ret0, err
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(validator address) constant returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenSession) IsValidator(validator common.Address) (bool, error) {
	return _SampleERC20MintableToken.Contract.IsValidator(&_SampleERC20MintableToken.CallOpts, validator)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(validator address) constant returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenCallerSession) IsValidator(validator common.Address) (bool, error) {
	return _SampleERC20MintableToken.Contract.IsValidator(&_SampleERC20MintableToken.CallOpts, validator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SampleERC20MintableToken *SampleERC20MintableTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SampleERC20MintableToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SampleERC20MintableToken *SampleERC20MintableTokenSession) Name() (string, error) {
	return _SampleERC20MintableToken.Contract.Name(&_SampleERC20MintableToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SampleERC20MintableToken *SampleERC20MintableTokenCallerSession) Name() (string, error) {
	return _SampleERC20MintableToken.Contract.Name(&_SampleERC20MintableToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SampleERC20MintableToken *SampleERC20MintableTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SampleERC20MintableToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SampleERC20MintableToken *SampleERC20MintableTokenSession) Symbol() (string, error) {
	return _SampleERC20MintableToken.Contract.Symbol(&_SampleERC20MintableToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SampleERC20MintableToken *SampleERC20MintableTokenCallerSession) Symbol() (string, error) {
	return _SampleERC20MintableToken.Contract.Symbol(&_SampleERC20MintableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SampleERC20MintableToken *SampleERC20MintableTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleERC20MintableToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SampleERC20MintableToken *SampleERC20MintableTokenSession) TotalSupply() (*big.Int, error) {
	return _SampleERC20MintableToken.Contract.TotalSupply(&_SampleERC20MintableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SampleERC20MintableToken *SampleERC20MintableTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _SampleERC20MintableToken.Contract.TotalSupply(&_SampleERC20MintableToken.CallOpts)
}

// AddGateway is a paid mutator transaction binding the contract method 0x68bb3795.
//
// Solidity: function addGateway(_gateway address) returns()
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactor) AddGateway(opts *bind.TransactOpts, _gateway common.Address) (*types.Transaction, error) {
	return _SampleERC20MintableToken.contract.Transact(opts, "addGateway", _gateway)
}

// AddGateway is a paid mutator transaction binding the contract method 0x68bb3795.
//
// Solidity: function addGateway(_gateway address) returns()
func (_SampleERC20MintableToken *SampleERC20MintableTokenSession) AddGateway(_gateway common.Address) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.AddGateway(&_SampleERC20MintableToken.TransactOpts, _gateway)
}

// AddGateway is a paid mutator transaction binding the contract method 0x68bb3795.
//
// Solidity: function addGateway(_gateway address) returns()
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactorSession) AddGateway(_gateway common.Address) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.AddGateway(&_SampleERC20MintableToken.TransactOpts, _gateway)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(account address) returns()
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _SampleERC20MintableToken.contract.Transact(opts, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(account address) returns()
func (_SampleERC20MintableToken *SampleERC20MintableTokenSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.AddMinter(&_SampleERC20MintableToken.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(account address) returns()
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactorSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.AddMinter(&_SampleERC20MintableToken.TransactOpts, account)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(newValidator address) returns()
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactor) AddValidator(opts *bind.TransactOpts, newValidator common.Address) (*types.Transaction, error) {
	return _SampleERC20MintableToken.contract.Transact(opts, "addValidator", newValidator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(newValidator address) returns()
func (_SampleERC20MintableToken *SampleERC20MintableTokenSession) AddValidator(newValidator common.Address) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.AddValidator(&_SampleERC20MintableToken.TransactOpts, newValidator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(newValidator address) returns()
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactorSession) AddValidator(newValidator common.Address) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.AddValidator(&_SampleERC20MintableToken.TransactOpts, newValidator)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _SampleERC20MintableToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.Approve(&_SampleERC20MintableToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.Approve(&_SampleERC20MintableToken.TransactOpts, spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(spender address, subtractedValue uint256) returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SampleERC20MintableToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(spender address, subtractedValue uint256) returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.DecreaseAllowance(&_SampleERC20MintableToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(spender address, subtractedValue uint256) returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.DecreaseAllowance(&_SampleERC20MintableToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(spender address, addedValue uint256) returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SampleERC20MintableToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(spender address, addedValue uint256) returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.IncreaseAllowance(&_SampleERC20MintableToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(spender address, addedValue uint256) returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.IncreaseAllowance(&_SampleERC20MintableToken.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SampleERC20MintableToken.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.Mint(&_SampleERC20MintableToken.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.Mint(&_SampleERC20MintableToken.TransactOpts, _to, _amount)
}

// MintTo is a paid mutator transaction binding the contract method 0x449a52f8.
//
// Solidity: function mintTo(_to address, _amount uint256) returns()
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactor) MintTo(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SampleERC20MintableToken.contract.Transact(opts, "mintTo", _to, _amount)
}

// MintTo is a paid mutator transaction binding the contract method 0x449a52f8.
//
// Solidity: function mintTo(_to address, _amount uint256) returns()
func (_SampleERC20MintableToken *SampleERC20MintableTokenSession) MintTo(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.MintTo(&_SampleERC20MintableToken.TransactOpts, _to, _amount)
}

// MintTo is a paid mutator transaction binding the contract method 0x449a52f8.
//
// Solidity: function mintTo(_to address, _amount uint256) returns()
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactorSession) MintTo(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.MintTo(&_SampleERC20MintableToken.TransactOpts, _to, _amount)
}

// RemoveGateway is a paid mutator transaction binding the contract method 0x8a885e35.
//
// Solidity: function removeGateway(_gateway address) returns()
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactor) RemoveGateway(opts *bind.TransactOpts, _gateway common.Address) (*types.Transaction, error) {
	return _SampleERC20MintableToken.contract.Transact(opts, "removeGateway", _gateway)
}

// RemoveGateway is a paid mutator transaction binding the contract method 0x8a885e35.
//
// Solidity: function removeGateway(_gateway address) returns()
func (_SampleERC20MintableToken *SampleERC20MintableTokenSession) RemoveGateway(_gateway common.Address) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.RemoveGateway(&_SampleERC20MintableToken.TransactOpts, _gateway)
}

// RemoveGateway is a paid mutator transaction binding the contract method 0x8a885e35.
//
// Solidity: function removeGateway(_gateway address) returns()
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactorSession) RemoveGateway(_gateway common.Address) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.RemoveGateway(&_SampleERC20MintableToken.TransactOpts, _gateway)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(validator address) returns()
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactor) RemoveValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _SampleERC20MintableToken.contract.Transact(opts, "removeValidator", validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(validator address) returns()
func (_SampleERC20MintableToken *SampleERC20MintableTokenSession) RemoveValidator(validator common.Address) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.RemoveValidator(&_SampleERC20MintableToken.TransactOpts, validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(validator address) returns()
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactorSession) RemoveValidator(validator common.Address) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.RemoveValidator(&_SampleERC20MintableToken.TransactOpts, validator)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactor) RenounceMinter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleERC20MintableToken.contract.Transact(opts, "renounceMinter")
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_SampleERC20MintableToken *SampleERC20MintableTokenSession) RenounceMinter() (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.RenounceMinter(&_SampleERC20MintableToken.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactorSession) RenounceMinter() (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.RenounceMinter(&_SampleERC20MintableToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SampleERC20MintableToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.Transfer(&_SampleERC20MintableToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.Transfer(&_SampleERC20MintableToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SampleERC20MintableToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.TransferFrom(&_SampleERC20MintableToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_SampleERC20MintableToken *SampleERC20MintableTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SampleERC20MintableToken.Contract.TransferFrom(&_SampleERC20MintableToken.TransactOpts, from, to, value)
}

// SampleERC20MintableTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SampleERC20MintableToken contract.
type SampleERC20MintableTokenApprovalIterator struct {
	Event *SampleERC20MintableTokenApproval // Event containing the contract specifics and raw log

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
func (it *SampleERC20MintableTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleERC20MintableTokenApproval)
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
		it.Event = new(SampleERC20MintableTokenApproval)
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
func (it *SampleERC20MintableTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleERC20MintableTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleERC20MintableTokenApproval represents a Approval event raised by the SampleERC20MintableToken contract.
type SampleERC20MintableTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_SampleERC20MintableToken *SampleERC20MintableTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SampleERC20MintableTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SampleERC20MintableToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SampleERC20MintableTokenApprovalIterator{contract: _SampleERC20MintableToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_SampleERC20MintableToken *SampleERC20MintableTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SampleERC20MintableTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SampleERC20MintableToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleERC20MintableTokenApproval)
				if err := _SampleERC20MintableToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// SampleERC20MintableTokenGatewayAddedIterator is returned from FilterGatewayAdded and is used to iterate over the raw logs and unpacked data for GatewayAdded events raised by the SampleERC20MintableToken contract.
type SampleERC20MintableTokenGatewayAddedIterator struct {
	Event *SampleERC20MintableTokenGatewayAdded // Event containing the contract specifics and raw log

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
func (it *SampleERC20MintableTokenGatewayAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleERC20MintableTokenGatewayAdded)
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
		it.Event = new(SampleERC20MintableTokenGatewayAdded)
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
func (it *SampleERC20MintableTokenGatewayAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleERC20MintableTokenGatewayAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleERC20MintableTokenGatewayAdded represents a GatewayAdded event raised by the SampleERC20MintableToken contract.
type SampleERC20MintableTokenGatewayAdded struct {
	Gateway common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGatewayAdded is a free log retrieval operation binding the contract event 0x7137528d21fb7b0b9462886348954009edac49570e27ef9dba8bb3a676fc11e2.
//
// Solidity: e GatewayAdded(gateway address)
func (_SampleERC20MintableToken *SampleERC20MintableTokenFilterer) FilterGatewayAdded(opts *bind.FilterOpts) (*SampleERC20MintableTokenGatewayAddedIterator, error) {

	logs, sub, err := _SampleERC20MintableToken.contract.FilterLogs(opts, "GatewayAdded")
	if err != nil {
		return nil, err
	}
	return &SampleERC20MintableTokenGatewayAddedIterator{contract: _SampleERC20MintableToken.contract, event: "GatewayAdded", logs: logs, sub: sub}, nil
}

// WatchGatewayAdded is a free log subscription operation binding the contract event 0x7137528d21fb7b0b9462886348954009edac49570e27ef9dba8bb3a676fc11e2.
//
// Solidity: e GatewayAdded(gateway address)
func (_SampleERC20MintableToken *SampleERC20MintableTokenFilterer) WatchGatewayAdded(opts *bind.WatchOpts, sink chan<- *SampleERC20MintableTokenGatewayAdded) (event.Subscription, error) {

	logs, sub, err := _SampleERC20MintableToken.contract.WatchLogs(opts, "GatewayAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleERC20MintableTokenGatewayAdded)
				if err := _SampleERC20MintableToken.contract.UnpackLog(event, "GatewayAdded", log); err != nil {
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

// SampleERC20MintableTokenGatewayRemovedIterator is returned from FilterGatewayRemoved and is used to iterate over the raw logs and unpacked data for GatewayRemoved events raised by the SampleERC20MintableToken contract.
type SampleERC20MintableTokenGatewayRemovedIterator struct {
	Event *SampleERC20MintableTokenGatewayRemoved // Event containing the contract specifics and raw log

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
func (it *SampleERC20MintableTokenGatewayRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleERC20MintableTokenGatewayRemoved)
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
		it.Event = new(SampleERC20MintableTokenGatewayRemoved)
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
func (it *SampleERC20MintableTokenGatewayRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleERC20MintableTokenGatewayRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleERC20MintableTokenGatewayRemoved represents a GatewayRemoved event raised by the SampleERC20MintableToken contract.
type SampleERC20MintableTokenGatewayRemoved struct {
	Gateway common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGatewayRemoved is a free log retrieval operation binding the contract event 0x6ad07cb5676ab639fa3821550f0ec4379900d542e2d8bd8f0f8158d8bd352da2.
//
// Solidity: e GatewayRemoved(gateway address)
func (_SampleERC20MintableToken *SampleERC20MintableTokenFilterer) FilterGatewayRemoved(opts *bind.FilterOpts) (*SampleERC20MintableTokenGatewayRemovedIterator, error) {

	logs, sub, err := _SampleERC20MintableToken.contract.FilterLogs(opts, "GatewayRemoved")
	if err != nil {
		return nil, err
	}
	return &SampleERC20MintableTokenGatewayRemovedIterator{contract: _SampleERC20MintableToken.contract, event: "GatewayRemoved", logs: logs, sub: sub}, nil
}

// WatchGatewayRemoved is a free log subscription operation binding the contract event 0x6ad07cb5676ab639fa3821550f0ec4379900d542e2d8bd8f0f8158d8bd352da2.
//
// Solidity: e GatewayRemoved(gateway address)
func (_SampleERC20MintableToken *SampleERC20MintableTokenFilterer) WatchGatewayRemoved(opts *bind.WatchOpts, sink chan<- *SampleERC20MintableTokenGatewayRemoved) (event.Subscription, error) {

	logs, sub, err := _SampleERC20MintableToken.contract.WatchLogs(opts, "GatewayRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleERC20MintableTokenGatewayRemoved)
				if err := _SampleERC20MintableToken.contract.UnpackLog(event, "GatewayRemoved", log); err != nil {
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

// SampleERC20MintableTokenMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the SampleERC20MintableToken contract.
type SampleERC20MintableTokenMinterAddedIterator struct {
	Event *SampleERC20MintableTokenMinterAdded // Event containing the contract specifics and raw log

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
func (it *SampleERC20MintableTokenMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleERC20MintableTokenMinterAdded)
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
		it.Event = new(SampleERC20MintableTokenMinterAdded)
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
func (it *SampleERC20MintableTokenMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleERC20MintableTokenMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleERC20MintableTokenMinterAdded represents a MinterAdded event raised by the SampleERC20MintableToken contract.
type SampleERC20MintableTokenMinterAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: e MinterAdded(account indexed address)
func (_SampleERC20MintableToken *SampleERC20MintableTokenFilterer) FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*SampleERC20MintableTokenMinterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SampleERC20MintableToken.contract.FilterLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &SampleERC20MintableTokenMinterAddedIterator{contract: _SampleERC20MintableToken.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: e MinterAdded(account indexed address)
func (_SampleERC20MintableToken *SampleERC20MintableTokenFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *SampleERC20MintableTokenMinterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SampleERC20MintableToken.contract.WatchLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleERC20MintableTokenMinterAdded)
				if err := _SampleERC20MintableToken.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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

// SampleERC20MintableTokenMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the SampleERC20MintableToken contract.
type SampleERC20MintableTokenMinterRemovedIterator struct {
	Event *SampleERC20MintableTokenMinterRemoved // Event containing the contract specifics and raw log

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
func (it *SampleERC20MintableTokenMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleERC20MintableTokenMinterRemoved)
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
		it.Event = new(SampleERC20MintableTokenMinterRemoved)
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
func (it *SampleERC20MintableTokenMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleERC20MintableTokenMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleERC20MintableTokenMinterRemoved represents a MinterRemoved event raised by the SampleERC20MintableToken contract.
type SampleERC20MintableTokenMinterRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: e MinterRemoved(account indexed address)
func (_SampleERC20MintableToken *SampleERC20MintableTokenFilterer) FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*SampleERC20MintableTokenMinterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SampleERC20MintableToken.contract.FilterLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &SampleERC20MintableTokenMinterRemovedIterator{contract: _SampleERC20MintableToken.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: e MinterRemoved(account indexed address)
func (_SampleERC20MintableToken *SampleERC20MintableTokenFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *SampleERC20MintableTokenMinterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SampleERC20MintableToken.contract.WatchLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleERC20MintableTokenMinterRemoved)
				if err := _SampleERC20MintableToken.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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

// SampleERC20MintableTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SampleERC20MintableToken contract.
type SampleERC20MintableTokenTransferIterator struct {
	Event *SampleERC20MintableTokenTransfer // Event containing the contract specifics and raw log

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
func (it *SampleERC20MintableTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleERC20MintableTokenTransfer)
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
		it.Event = new(SampleERC20MintableTokenTransfer)
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
func (it *SampleERC20MintableTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleERC20MintableTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleERC20MintableTokenTransfer represents a Transfer event raised by the SampleERC20MintableToken contract.
type SampleERC20MintableTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_SampleERC20MintableToken *SampleERC20MintableTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SampleERC20MintableTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SampleERC20MintableToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SampleERC20MintableTokenTransferIterator{contract: _SampleERC20MintableToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_SampleERC20MintableToken *SampleERC20MintableTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SampleERC20MintableTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SampleERC20MintableToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleERC20MintableTokenTransfer)
				if err := _SampleERC20MintableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// SampleERC20MintableTokenValidatorAddedIterator is returned from FilterValidatorAdded and is used to iterate over the raw logs and unpacked data for ValidatorAdded events raised by the SampleERC20MintableToken contract.
type SampleERC20MintableTokenValidatorAddedIterator struct {
	Event *SampleERC20MintableTokenValidatorAdded // Event containing the contract specifics and raw log

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
func (it *SampleERC20MintableTokenValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleERC20MintableTokenValidatorAdded)
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
		it.Event = new(SampleERC20MintableTokenValidatorAdded)
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
func (it *SampleERC20MintableTokenValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleERC20MintableTokenValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleERC20MintableTokenValidatorAdded represents a ValidatorAdded event raised by the SampleERC20MintableToken contract.
type SampleERC20MintableTokenValidatorAdded struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorAdded is a free log retrieval operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: e ValidatorAdded(validator address)
func (_SampleERC20MintableToken *SampleERC20MintableTokenFilterer) FilterValidatorAdded(opts *bind.FilterOpts) (*SampleERC20MintableTokenValidatorAddedIterator, error) {

	logs, sub, err := _SampleERC20MintableToken.contract.FilterLogs(opts, "ValidatorAdded")
	if err != nil {
		return nil, err
	}
	return &SampleERC20MintableTokenValidatorAddedIterator{contract: _SampleERC20MintableToken.contract, event: "ValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorAdded is a free log subscription operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: e ValidatorAdded(validator address)
func (_SampleERC20MintableToken *SampleERC20MintableTokenFilterer) WatchValidatorAdded(opts *bind.WatchOpts, sink chan<- *SampleERC20MintableTokenValidatorAdded) (event.Subscription, error) {

	logs, sub, err := _SampleERC20MintableToken.contract.WatchLogs(opts, "ValidatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleERC20MintableTokenValidatorAdded)
				if err := _SampleERC20MintableToken.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
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

// SampleERC20MintableTokenValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the SampleERC20MintableToken contract.
type SampleERC20MintableTokenValidatorRemovedIterator struct {
	Event *SampleERC20MintableTokenValidatorRemoved // Event containing the contract specifics and raw log

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
func (it *SampleERC20MintableTokenValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleERC20MintableTokenValidatorRemoved)
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
		it.Event = new(SampleERC20MintableTokenValidatorRemoved)
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
func (it *SampleERC20MintableTokenValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleERC20MintableTokenValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleERC20MintableTokenValidatorRemoved represents a ValidatorRemoved event raised by the SampleERC20MintableToken contract.
type SampleERC20MintableTokenValidatorRemoved struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: e ValidatorRemoved(validator address)
func (_SampleERC20MintableToken *SampleERC20MintableTokenFilterer) FilterValidatorRemoved(opts *bind.FilterOpts) (*SampleERC20MintableTokenValidatorRemovedIterator, error) {

	logs, sub, err := _SampleERC20MintableToken.contract.FilterLogs(opts, "ValidatorRemoved")
	if err != nil {
		return nil, err
	}
	return &SampleERC20MintableTokenValidatorRemovedIterator{contract: _SampleERC20MintableToken.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: e ValidatorRemoved(validator address)
func (_SampleERC20MintableToken *SampleERC20MintableTokenFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *SampleERC20MintableTokenValidatorRemoved) (event.Subscription, error) {

	logs, sub, err := _SampleERC20MintableToken.contract.WatchLogs(opts, "ValidatorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleERC20MintableTokenValidatorRemoved)
				if err := _SampleERC20MintableToken.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
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
