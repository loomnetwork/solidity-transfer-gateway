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

// MainnetCryptoCardsContractABI is the input ABI used to generate the binding from.
const MainnetCryptoCardsContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_gateway\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"mintTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"depositToGateway\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MainnetCryptoCardsContractBin is the compiled bytecode used for deploying new contracts.
const MainnetCryptoCardsContractBin = `0x60806040523480156200001157600080fd5b506040516020806200126f833981018060405260208110156200003357600080fd5b5051604080518082018252600b81527f43727970746f43617264730000000000000000000000000000000000000000006020828101919091528251808401909352600383527f43434300000000000000000000000000000000000000000000000000000000008382015290919082908290620000d5907f01ffc9a7000000000000000000000000000000000000000000000000000000009062000183811b901c565b620000ed6380ac58cd60e01b6200018360201b60201c565b6200010563780e9d6360e01b6200018360201b60201c565b81516200011a906009906020850190620001f0565b5080516200013090600a906020840190620001f0565b5062000149635b5e139f60e01b6200018360201b60201c565b5050600d8054336001600160a01b031991821617909155600e80549091166001600160a01b03949094169390931790925550620002959050565b7fffffffff000000000000000000000000000000000000000000000000000000008082161415620001b357600080fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200023357805160ff191683800117855562000263565b8280016001018555821562000263579182015b828111156200026357825182559160200191906001019062000246565b506200027192915062000275565b5090565b6200029291905b808211156200027157600081556001016200027c565b90565b610fca80620002a56000396000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c80634f6ccce7116100ad578063a22cb46511610071578063a22cb46514610384578063b88d4fde146103b2578063bcfaa79d14610478578063c87b56dd1461049e578063e985e9c5146104bb57610121565b80634f6ccce7146102ff5780636352211e1461031c57806370a08231146103395780639267daba1461035f57806395d89b411461037c57610121565b8063116191b6116100f4578063116191b61461024557806318160ddd1461024d57806323b872dd146102675780632f745c591461029d57806342842e0e146102c957610121565b806301ffc9a71461012657806306fdde0314610161578063081812fc146101de578063095ea7b314610217575b600080fd5b61014d6004803603602081101561013c57600080fd5b50356001600160e01b0319166104e9565b604080519115158252519081900360200190f35b610169610508565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101a357818101518382015260200161018b565b50505050905090810190601f1680156101d05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101fb600480360360208110156101f457600080fd5b503561059f565b604080516001600160a01b039092168252519081900360200190f35b6102436004803603604081101561022d57600080fd5b506001600160a01b0381351690602001356105cf565b005b6101fb61067c565b61025561068b565b60408051918252519081900360200190f35b6102436004803603606081101561027d57600080fd5b506001600160a01b03813581169160208101359091169060400135610691565b610255600480360360408110156102b357600080fd5b506001600160a01b0381351690602001356106b4565b610243600480360360608110156102df57600080fd5b506001600160a01b03813581169160208101359091169060400135610701565b6102556004803603602081101561031557600080fd5b503561071c565b6101fb6004803603602081101561033257600080fd5b5035610750565b6102556004803603602081101561034f57600080fd5b50356001600160a01b0316610778565b6102436004803603602081101561037557600080fd5b50356107ae565b6101696107c9565b6102436004803603604081101561039a57600080fd5b506001600160a01b038135169060200135151561082a565b610243600480360360808110156103c857600080fd5b6001600160a01b0382358116926020810135909116916040820135919081019060808101606082013564010000000081111561040357600080fd5b82018360208201111561041557600080fd5b8035906020019184600183028401116401000000008311171561043757600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506108ae945050505050565b6102436004803603602081101561048e57600080fd5b50356001600160a01b03166108d4565b610169600480360360208110156104b457600080fd5b503561091d565b61014d600480360360408110156104d157600080fd5b506001600160a01b03813581169160200135166109d0565b6001600160e01b03191660009081526020819052604090205460ff1690565b60098054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156105945780601f1061056957610100808354040283529160200191610594565b820191906000526020600020905b81548152906001019060200180831161057757829003601f168201915b505050505090505b90565b60006105aa826109fe565b6105b357600080fd5b506000908152600260205260409020546001600160a01b031690565b60006105da82610750565b9050806001600160a01b0316836001600160a01b031614156105fb57600080fd5b336001600160a01b0382161480610617575061061781336109d0565b61062057600080fd5b60008281526002602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b600e546001600160a01b031681565b60075490565b61069b3382610a1b565b6106a457600080fd5b6106af838383610a7a565b505050565b60006106bf83610778565b82106106ca57600080fd5b6001600160a01b03831660009081526005602052604090208054839081106106ee57fe5b9060005260206000200154905092915050565b6106af838383604051806020016040528060008152506108ae565b600061072661068b565b821061073157600080fd5b6007828154811061073e57fe5b90600052602060002001549050919050565b6000818152600160205260408120546001600160a01b03168061077257600080fd5b92915050565b60006001600160a01b03821661078d57600080fd5b6001600160a01b038216600090815260036020526040902061077290610a99565b600e546107c69033906001600160a01b031683610701565b50565b600a8054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156105945780601f1061056957610100808354040283529160200191610594565b6001600160a01b03821633141561084057600080fd5b3360008181526004602090815260408083206001600160a01b03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b6108b9848484610691565b6108c584848484610a9d565b6108ce57600080fd5b50505050565b600d546001600160a01b031633146108eb57600080fd5b60005b600581121561091957600061090161068b565b60010190506109108382610bd6565b506001016108ee565b5050565b6060610928826109fe565b61093157600080fd5b6000828152600b602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845290918301828280156109c45780601f10610999576101008083540402835291602001916109c4565b820191906000526020600020905b8154815290600101906020018083116109a757829003601f168201915b50505050509050919050565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b6000908152600160205260409020546001600160a01b0316151590565b600080610a2783610750565b9050806001600160a01b0316846001600160a01b03161480610a625750836001600160a01b0316610a578461059f565b6001600160a01b0316145b80610a725750610a7281856109d0565b949350505050565b610a85838383610bf3565b610a8f8382610cd3565b6106af8282610dc8565b5490565b6000610ab1846001600160a01b0316610e06565b610abd57506001610a72565b604051600160e11b630a85bd0102815233600482018181526001600160a01b03888116602485015260448401879052608060648501908152865160848601528651600095928a169463150b7a029490938c938b938b939260a4019060208501908083838e5b83811015610b3a578181015183820152602001610b22565b50505050905090810190601f168015610b675780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b158015610b8957600080fd5b505af1158015610b9d573d6000803e3d6000fd5b505050506040513d6020811015610bb357600080fd5b50516001600160e01b031916600160e11b630a85bd010214915050949350505050565b610be08282610e0c565b610bea8282610dc8565b61091981610ead565b826001600160a01b0316610c0682610750565b6001600160a01b031614610c1957600080fd5b6001600160a01b038216610c2c57600080fd5b610c3581610ef1565b6001600160a01b0383166000908152600360205260409020610c5690610f2c565b6001600160a01b0382166000908152600360205260409020610c7790610f43565b60008181526001602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6001600160a01b038216600090815260056020526040812054610cfd90600163ffffffff610f4c16565b600083815260066020526040902054909150808214610d98576001600160a01b0384166000908152600560205260408120805484908110610d3a57fe5b906000526020600020015490508060056000876001600160a01b03166001600160a01b031681526020019081526020016000208381548110610d7857fe5b600091825260208083209091019290925591825260069052604090208190555b6001600160a01b0384166000908152600560205260409020805490610dc1906000198301610f61565b5050505050565b6001600160a01b0390911660009081526005602081815260408084208054868652600684529185208290559282526001810183559183529091200155565b3b151590565b6001600160a01b038216610e1f57600080fd5b610e28816109fe565b15610e3257600080fd5b600081815260016020908152604080832080546001600160a01b0319166001600160a01b038716908117909155835260039091529020610e7190610f43565b60405181906001600160a01b038416906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b600780546000838152600860205260408120829055600182018355919091527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6880155565b6000818152600260205260409020546001600160a01b0316156107c657600090815260026020526040902080546001600160a01b0319169055565b8054610f3f90600163ffffffff610f4c16565b9055565b80546001019055565b600082821115610f5b57600080fd5b50900390565b8154818355818111156106af576000838152602090206106af91810190830161059c91905b80821115610f9a5760008155600101610f86565b509056fea165627a7a7230582085944150cc80b2d127d4a64a723c758f2fc1b54da91e97dce46a5a58fe09a5310029`

// DeployMainnetCryptoCardsContract deploys a new Ethereum contract, binding an instance of MainnetCryptoCardsContract to it.
func DeployMainnetCryptoCardsContract(auth *bind.TransactOpts, backend bind.ContractBackend, _gateway common.Address) (common.Address, *types.Transaction, *MainnetCryptoCardsContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainnetCryptoCardsContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MainnetCryptoCardsContractBin), backend, _gateway)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MainnetCryptoCardsContract{MainnetCryptoCardsContractCaller: MainnetCryptoCardsContractCaller{contract: contract}, MainnetCryptoCardsContractTransactor: MainnetCryptoCardsContractTransactor{contract: contract}, MainnetCryptoCardsContractFilterer: MainnetCryptoCardsContractFilterer{contract: contract}}, nil
}

// MainnetCryptoCardsContract is an auto generated Go binding around an Ethereum contract.
type MainnetCryptoCardsContract struct {
	MainnetCryptoCardsContractCaller     // Read-only binding to the contract
	MainnetCryptoCardsContractTransactor // Write-only binding to the contract
	MainnetCryptoCardsContractFilterer   // Log filterer for contract events
}

// MainnetCryptoCardsContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainnetCryptoCardsContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainnetCryptoCardsContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainnetCryptoCardsContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainnetCryptoCardsContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainnetCryptoCardsContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainnetCryptoCardsContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainnetCryptoCardsContractSession struct {
	Contract     *MainnetCryptoCardsContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MainnetCryptoCardsContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainnetCryptoCardsContractCallerSession struct {
	Contract *MainnetCryptoCardsContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// MainnetCryptoCardsContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainnetCryptoCardsContractTransactorSession struct {
	Contract     *MainnetCryptoCardsContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// MainnetCryptoCardsContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainnetCryptoCardsContractRaw struct {
	Contract *MainnetCryptoCardsContract // Generic contract binding to access the raw methods on
}

// MainnetCryptoCardsContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainnetCryptoCardsContractCallerRaw struct {
	Contract *MainnetCryptoCardsContractCaller // Generic read-only contract binding to access the raw methods on
}

// MainnetCryptoCardsContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainnetCryptoCardsContractTransactorRaw struct {
	Contract *MainnetCryptoCardsContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMainnetCryptoCardsContract creates a new instance of MainnetCryptoCardsContract, bound to a specific deployed contract.
func NewMainnetCryptoCardsContract(address common.Address, backend bind.ContractBackend) (*MainnetCryptoCardsContract, error) {
	contract, err := bindMainnetCryptoCardsContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MainnetCryptoCardsContract{MainnetCryptoCardsContractCaller: MainnetCryptoCardsContractCaller{contract: contract}, MainnetCryptoCardsContractTransactor: MainnetCryptoCardsContractTransactor{contract: contract}, MainnetCryptoCardsContractFilterer: MainnetCryptoCardsContractFilterer{contract: contract}}, nil
}

// NewMainnetCryptoCardsContractCaller creates a new read-only instance of MainnetCryptoCardsContract, bound to a specific deployed contract.
func NewMainnetCryptoCardsContractCaller(address common.Address, caller bind.ContractCaller) (*MainnetCryptoCardsContractCaller, error) {
	contract, err := bindMainnetCryptoCardsContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainnetCryptoCardsContractCaller{contract: contract}, nil
}

// NewMainnetCryptoCardsContractTransactor creates a new write-only instance of MainnetCryptoCardsContract, bound to a specific deployed contract.
func NewMainnetCryptoCardsContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MainnetCryptoCardsContractTransactor, error) {
	contract, err := bindMainnetCryptoCardsContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainnetCryptoCardsContractTransactor{contract: contract}, nil
}

// NewMainnetCryptoCardsContractFilterer creates a new log filterer instance of MainnetCryptoCardsContract, bound to a specific deployed contract.
func NewMainnetCryptoCardsContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MainnetCryptoCardsContractFilterer, error) {
	contract, err := bindMainnetCryptoCardsContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainnetCryptoCardsContractFilterer{contract: contract}, nil
}

// bindMainnetCryptoCardsContract binds a generic wrapper to an already deployed contract.
func bindMainnetCryptoCardsContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainnetCryptoCardsContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MainnetCryptoCardsContract.Contract.MainnetCryptoCardsContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.Contract.MainnetCryptoCardsContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.Contract.MainnetCryptoCardsContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MainnetCryptoCardsContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MainnetCryptoCardsContract.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MainnetCryptoCardsContract.Contract.BalanceOf(&_MainnetCryptoCardsContract.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MainnetCryptoCardsContract.Contract.BalanceOf(&_MainnetCryptoCardsContract.CallOpts, owner)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() constant returns(address)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MainnetCryptoCardsContract.contract.Call(opts, out, "gateway")
	return *ret0, err
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() constant returns(address)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) Gateway() (common.Address, error) {
	return _MainnetCryptoCardsContract.Contract.Gateway(&_MainnetCryptoCardsContract.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() constant returns(address)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCallerSession) Gateway() (common.Address, error) {
	return _MainnetCryptoCardsContract.Contract.Gateway(&_MainnetCryptoCardsContract.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(tokenId uint256) constant returns(address)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MainnetCryptoCardsContract.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(tokenId uint256) constant returns(address)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _MainnetCryptoCardsContract.Contract.GetApproved(&_MainnetCryptoCardsContract.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(tokenId uint256) constant returns(address)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _MainnetCryptoCardsContract.Contract.GetApproved(&_MainnetCryptoCardsContract.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(owner address, operator address) constant returns(bool)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MainnetCryptoCardsContract.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(owner address, operator address) constant returns(bool)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MainnetCryptoCardsContract.Contract.IsApprovedForAll(&_MainnetCryptoCardsContract.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(owner address, operator address) constant returns(bool)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MainnetCryptoCardsContract.Contract.IsApprovedForAll(&_MainnetCryptoCardsContract.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MainnetCryptoCardsContract.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) Name() (string, error) {
	return _MainnetCryptoCardsContract.Contract.Name(&_MainnetCryptoCardsContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCallerSession) Name() (string, error) {
	return _MainnetCryptoCardsContract.Contract.Name(&_MainnetCryptoCardsContract.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(tokenId uint256) constant returns(address)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MainnetCryptoCardsContract.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(tokenId uint256) constant returns(address)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _MainnetCryptoCardsContract.Contract.OwnerOf(&_MainnetCryptoCardsContract.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(tokenId uint256) constant returns(address)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _MainnetCryptoCardsContract.Contract.OwnerOf(&_MainnetCryptoCardsContract.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(interfaceId bytes4) constant returns(bool)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MainnetCryptoCardsContract.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(interfaceId bytes4) constant returns(bool)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MainnetCryptoCardsContract.Contract.SupportsInterface(&_MainnetCryptoCardsContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(interfaceId bytes4) constant returns(bool)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MainnetCryptoCardsContract.Contract.SupportsInterface(&_MainnetCryptoCardsContract.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MainnetCryptoCardsContract.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) Symbol() (string, error) {
	return _MainnetCryptoCardsContract.Contract.Symbol(&_MainnetCryptoCardsContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCallerSession) Symbol() (string, error) {
	return _MainnetCryptoCardsContract.Contract.Symbol(&_MainnetCryptoCardsContract.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(index uint256) constant returns(uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MainnetCryptoCardsContract.contract.Call(opts, out, "tokenByIndex", index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(index uint256) constant returns(uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _MainnetCryptoCardsContract.Contract.TokenByIndex(&_MainnetCryptoCardsContract.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(index uint256) constant returns(uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _MainnetCryptoCardsContract.Contract.TokenByIndex(&_MainnetCryptoCardsContract.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(owner address, index uint256) constant returns(uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MainnetCryptoCardsContract.contract.Call(opts, out, "tokenOfOwnerByIndex", owner, index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(owner address, index uint256) constant returns(uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _MainnetCryptoCardsContract.Contract.TokenOfOwnerByIndex(&_MainnetCryptoCardsContract.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(owner address, index uint256) constant returns(uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _MainnetCryptoCardsContract.Contract.TokenOfOwnerByIndex(&_MainnetCryptoCardsContract.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(tokenId uint256) constant returns(string)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MainnetCryptoCardsContract.contract.Call(opts, out, "tokenURI", tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(tokenId uint256) constant returns(string)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) TokenURI(tokenId *big.Int) (string, error) {
	return _MainnetCryptoCardsContract.Contract.TokenURI(&_MainnetCryptoCardsContract.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(tokenId uint256) constant returns(string)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _MainnetCryptoCardsContract.Contract.TokenURI(&_MainnetCryptoCardsContract.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MainnetCryptoCardsContract.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) TotalSupply() (*big.Int, error) {
	return _MainnetCryptoCardsContract.Contract.TotalSupply(&_MainnetCryptoCardsContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCallerSession) TotalSupply() (*big.Int, error) {
	return _MainnetCryptoCardsContract.Contract.TotalSupply(&_MainnetCryptoCardsContract.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(to address, tokenId uint256) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(to address, tokenId uint256) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.Contract.Approve(&_MainnetCryptoCardsContract.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(to address, tokenId uint256) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.Contract.Approve(&_MainnetCryptoCardsContract.TransactOpts, to, tokenId)
}

// DepositToGateway is a paid mutator transaction binding the contract method 0x9267daba.
//
// Solidity: function depositToGateway(tokenId uint256) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractTransactor) DepositToGateway(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.contract.Transact(opts, "depositToGateway", tokenId)
}

// DepositToGateway is a paid mutator transaction binding the contract method 0x9267daba.
//
// Solidity: function depositToGateway(tokenId uint256) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) DepositToGateway(tokenId *big.Int) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.Contract.DepositToGateway(&_MainnetCryptoCardsContract.TransactOpts, tokenId)
}

// DepositToGateway is a paid mutator transaction binding the contract method 0x9267daba.
//
// Solidity: function depositToGateway(tokenId uint256) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractTransactorSession) DepositToGateway(tokenId *big.Int) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.Contract.DepositToGateway(&_MainnetCryptoCardsContract.TransactOpts, tokenId)
}

// MintTokens is a paid mutator transaction binding the contract method 0xbcfaa79d.
//
// Solidity: function mintTokens(_to address) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractTransactor) MintTokens(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.contract.Transact(opts, "mintTokens", _to)
}

// MintTokens is a paid mutator transaction binding the contract method 0xbcfaa79d.
//
// Solidity: function mintTokens(_to address) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) MintTokens(_to common.Address) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.Contract.MintTokens(&_MainnetCryptoCardsContract.TransactOpts, _to)
}

// MintTokens is a paid mutator transaction binding the contract method 0xbcfaa79d.
//
// Solidity: function mintTokens(_to address) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractTransactorSession) MintTokens(_to common.Address) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.Contract.MintTokens(&_MainnetCryptoCardsContract.TransactOpts, _to)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(from address, to address, tokenId uint256, _data bytes) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.contract.Transact(opts, "safeTransferFrom", from, to, tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(from address, to address, tokenId uint256, _data bytes) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.Contract.SafeTransferFrom(&_MainnetCryptoCardsContract.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(from address, to address, tokenId uint256, _data bytes) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.Contract.SafeTransferFrom(&_MainnetCryptoCardsContract.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(to address, approved bool) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractTransactor) SetApprovalForAll(opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.contract.Transact(opts, "setApprovalForAll", to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(to address, approved bool) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.Contract.SetApprovalForAll(&_MainnetCryptoCardsContract.TransactOpts, to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(to address, approved bool) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractTransactorSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.Contract.SetApprovalForAll(&_MainnetCryptoCardsContract.TransactOpts, to, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokenId uint256) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokenId uint256) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.Contract.TransferFrom(&_MainnetCryptoCardsContract.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokenId uint256) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.Contract.TransferFrom(&_MainnetCryptoCardsContract.TransactOpts, from, to, tokenId)
}

// MainnetCryptoCardsContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MainnetCryptoCardsContract contract.
type MainnetCryptoCardsContractApprovalIterator struct {
	Event *MainnetCryptoCardsContractApproval // Event containing the contract specifics and raw log

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
func (it *MainnetCryptoCardsContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainnetCryptoCardsContractApproval)
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
		it.Event = new(MainnetCryptoCardsContractApproval)
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
func (it *MainnetCryptoCardsContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainnetCryptoCardsContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainnetCryptoCardsContractApproval represents a Approval event raised by the MainnetCryptoCardsContract contract.
type MainnetCryptoCardsContractApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, approved indexed address, tokenId indexed uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*MainnetCryptoCardsContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MainnetCryptoCardsContract.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MainnetCryptoCardsContractApprovalIterator{contract: _MainnetCryptoCardsContract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, approved indexed address, tokenId indexed uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MainnetCryptoCardsContractApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MainnetCryptoCardsContract.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainnetCryptoCardsContractApproval)
				if err := _MainnetCryptoCardsContract.contract.UnpackLog(event, "Approval", log); err != nil {
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

// MainnetCryptoCardsContractApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the MainnetCryptoCardsContract contract.
type MainnetCryptoCardsContractApprovalForAllIterator struct {
	Event *MainnetCryptoCardsContractApprovalForAll // Event containing the contract specifics and raw log

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
func (it *MainnetCryptoCardsContractApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainnetCryptoCardsContractApprovalForAll)
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
		it.Event = new(MainnetCryptoCardsContractApprovalForAll)
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
func (it *MainnetCryptoCardsContractApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainnetCryptoCardsContractApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainnetCryptoCardsContractApprovalForAll represents a ApprovalForAll event raised by the MainnetCryptoCardsContract contract.
type MainnetCryptoCardsContractApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(owner indexed address, operator indexed address, approved bool)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*MainnetCryptoCardsContractApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MainnetCryptoCardsContract.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MainnetCryptoCardsContractApprovalForAllIterator{contract: _MainnetCryptoCardsContract.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(owner indexed address, operator indexed address, approved bool)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MainnetCryptoCardsContractApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MainnetCryptoCardsContract.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainnetCryptoCardsContractApprovalForAll)
				if err := _MainnetCryptoCardsContract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// MainnetCryptoCardsContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MainnetCryptoCardsContract contract.
type MainnetCryptoCardsContractTransferIterator struct {
	Event *MainnetCryptoCardsContractTransfer // Event containing the contract specifics and raw log

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
func (it *MainnetCryptoCardsContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainnetCryptoCardsContractTransfer)
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
		it.Event = new(MainnetCryptoCardsContractTransfer)
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
func (it *MainnetCryptoCardsContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainnetCryptoCardsContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainnetCryptoCardsContractTransfer represents a Transfer event raised by the MainnetCryptoCardsContract contract.
type MainnetCryptoCardsContractTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, tokenId indexed uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*MainnetCryptoCardsContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MainnetCryptoCardsContract.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MainnetCryptoCardsContractTransferIterator{contract: _MainnetCryptoCardsContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, tokenId indexed uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MainnetCryptoCardsContractTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MainnetCryptoCardsContract.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainnetCryptoCardsContractTransfer)
				if err := _MainnetCryptoCardsContract.contract.UnpackLog(event, "Transfer", log); err != nil {
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
