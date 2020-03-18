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

// MainnetCryptoCardsContractABI is the input ABI used to generate the binding from.
const MainnetCryptoCardsContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"InterfaceId_ERC165\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_gateway\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"mintTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"depositToGateway\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MainnetCryptoCardsContractBin is the compiled bytecode used for deploying new contracts.
const MainnetCryptoCardsContractBin = `0x60806040523480156200001157600080fd5b506040516020806200137e83398101604081815291518282018352600b82527f43727970746f43617264730000000000000000000000000000000000000000006020808401919091528351808501909452600384527f43434300000000000000000000000000000000000000000000000000000000009084015291620000c07f01ffc9a700000000000000000000000000000000000000000000000000000000640100000000620001f2810204565b620000f47f80ac58cd00000000000000000000000000000000000000000000000000000000640100000000620001f2810204565b620001287f4f558e7900000000000000000000000000000000000000000000000000000000640100000000620001f2810204565b81516200013d9060059060208501906200025f565b508051620001539060069060208401906200025f565b50620001887f780e9d6300000000000000000000000000000000000000000000000000000000640100000000620001f2810204565b620001bc7f5b5e139f00000000000000000000000000000000000000000000000000000000640100000000620001f2810204565b5050600d805433600160a060020a031991821617909155600e8054909116600160a060020a039290921691909117905562000304565b7fffffffff0000000000000000000000000000000000000000000000000000000080821614156200022257600080fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620002a257805160ff1916838001178555620002d2565b82800160010185558215620002d2579182015b82811115620002d2578251825591602001919060010190620002b5565b50620002e0929150620002e4565b5090565b6200030191905b80821115620002e05760008155600101620002eb565b90565b61106a80620003146000396000f30060806040526004361061011c5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301ffc9a7811461012157806306fdde0314610157578063081812fc146101e1578063095ea7b314610215578063116191b61461023b57806318160ddd1461025057806319fa8f501461027757806323b872dd146102a95780632f745c59146102d357806342842e0e146102f75780634f558e79146103215780634f6ccce7146103395780636352211e1461035157806370a08231146103695780639267daba1461038a57806395d89b41146103a2578063a22cb465146103b7578063b88d4fde146103dd578063bcfaa79d1461044c578063c87b56dd1461046d578063e985e9c514610485575b600080fd5b34801561012d57600080fd5b50610143600160e060020a0319600435166104ac565b604080519115158252519081900360200190f35b34801561016357600080fd5b5061016c6104cb565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101a657818101518382015260200161018e565b50505050905090810190601f1680156101d35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101ed57600080fd5b506101f9600435610562565b60408051600160a060020a039092168252519081900360200190f35b34801561022157600080fd5b50610239600160a060020a036004351660243561057d565b005b34801561024757600080fd5b506101f9610633565b34801561025c57600080fd5b50610265610642565b60408051918252519081900360200190f35b34801561028357600080fd5b5061028c610648565b60408051600160e060020a03199092168252519081900360200190f35b3480156102b557600080fd5b50610239600160a060020a036004358116906024351660443561066c565b3480156102df57600080fd5b50610265600160a060020a036004351660243561070f565b34801561030357600080fd5b50610239600160a060020a036004358116906024351660443561075c565b34801561032d57600080fd5b5061014360043561077d565b34801561034557600080fd5b5061026560043561079a565b34801561035d57600080fd5b506101f96004356107cf565b34801561037557600080fd5b50610265600160a060020a03600435166107f9565b34801561039657600080fd5b5061023960043561082c565b3480156103ae57600080fd5b5061016c610847565b3480156103c357600080fd5b50610239600160a060020a036004351660243515156108a8565b3480156103e957600080fd5b50604080516020601f60643560048181013592830184900484028501840190955281845261023994600160a060020a03813581169560248035909216956044359536956084940191819084018382808284375094975061092c9650505050505050565b34801561045857600080fd5b50610239600160a060020a0360043516610954565b34801561047957600080fd5b5061016c60043561099a565b34801561049157600080fd5b50610143600160a060020a0360043581169060243516610a4f565b600160e060020a03191660009081526020819052604090205460ff1690565b60058054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156105575780601f1061052c57610100808354040283529160200191610557565b820191906000526020600020905b81548152906001019060200180831161053a57829003601f168201915b505050505090505b90565b600090815260026020526040902054600160a060020a031690565b6000610588826107cf565b9050600160a060020a0383811690821614156105a357600080fd5b33600160a060020a03821614806105bf57506105bf8133610a4f565b15156105ca57600080fd5b600082815260026020526040808220805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b600e54600160a060020a031681565b60095490565b7f01ffc9a70000000000000000000000000000000000000000000000000000000081565b6106763382610a7d565b151561068157600080fd5b600160a060020a038316151561069657600080fd5b600160a060020a03821615156106ab57600080fd5b6106b58382610adc565b6106bf8382610b4d565b6106c98282610c54565b8082600160a060020a031684600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b600061071a836107f9565b821061072557600080fd5b600160a060020a038316600090815260076020526040902080548390811061074957fe5b9060005260206000200154905092915050565b610778838383602060405190810160405280600081525061092c565b505050565b600090815260016020526040902054600160a060020a0316151590565b60006107a4610642565b82106107af57600080fd5b60098054839081106107bd57fe5b90600052602060002001549050919050565b600081815260016020526040812054600160a060020a03168015156107f357600080fd5b92915050565b6000600160a060020a038216151561081057600080fd5b50600160a060020a031660009081526003602052604090205490565b600e54610844903390600160a060020a03168361075c565b50565b60068054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156105575780601f1061052c57610100808354040283529160200191610557565b600160a060020a0382163314156108be57600080fd5b336000818152600460209081526040808320600160a060020a03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b61093784848461066c565b61094384848484610c9d565b151561094e57600080fd5b50505050565b600d546000908190600160a060020a0316331461097057600080fd5b600091505b6005821215610778575060095460010161098f8382610e0a565b600190910190610975565b60606109a58261077d565b15156109b057600080fd5b6000828152600b602090815260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529091830182828015610a435780601f10610a1857610100808354040283529160200191610a43565b820191906000526020600020905b815481529060010190602001808311610a2657829003601f168201915b50505050509050919050565b600160a060020a03918216600090815260046020908152604080832093909416825291909152205460ff1690565b600080610a89836107cf565b905080600160a060020a031684600160a060020a03161480610ac4575083600160a060020a0316610ab984610562565b600160a060020a0316145b80610ad45750610ad48185610a4f565b949350505050565b81600160a060020a0316610aef826107cf565b600160a060020a031614610b0257600080fd5b600081815260026020526040902054600160a060020a031615610b49576000818152600260205260409020805473ffffffffffffffffffffffffffffffffffffffff191690555b5050565b6000806000610b5c8585610e59565b600084815260086020908152604080832054600160a060020a0389168452600790925290912054909350610b9790600163ffffffff610eef16565b600160a060020a038616600090815260076020526040902080549193509083908110610bbf57fe5b90600052602060002001549050806007600087600160a060020a0316600160a060020a0316815260200190815260200160002084815481101515610bff57fe5b6000918252602080832090910192909255600160a060020a0387168152600790915260409020805490610c36906000198301611001565b50600093845260086020526040808520859055908452909220555050565b6000610c608383610f01565b50600160a060020a039091166000908152600760209081526040808320805460018101825590845282842081018590559383526008909152902055565b600080610cb285600160a060020a0316610f91565b1515610cc15760019150610e01565b6040517f150b7a020000000000000000000000000000000000000000000000000000000081523360048201818152600160a060020a03898116602485015260448401889052608060648501908152875160848601528751918a169463150b7a0294938c938b938b93909160a490910190602085019080838360005b83811015610d54578181015183820152602001610d3c565b50505050905090810190601f168015610d815780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b158015610da357600080fd5b505af1158015610db7573d6000803e3d6000fd5b505050506040513d6020811015610dcd57600080fd5b5051600160e060020a031981167f150b7a020000000000000000000000000000000000000000000000000000000014925090505b50949350505050565b610e148282610f99565b600980546000838152600a60205260408120829055600182018355919091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af015550565b81600160a060020a0316610e6c826107cf565b600160a060020a031614610e7f57600080fd5b600160a060020a038216600090815260036020526040902054610ea990600163ffffffff610eef16565b600160a060020a03909216600090815260036020908152604080832094909455918152600190915220805473ffffffffffffffffffffffffffffffffffffffff19169055565b600082821115610efb57fe5b50900390565b600081815260016020526040902054600160a060020a031615610f2357600080fd5b6000818152600160208181526040808420805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0388169081179091558452600390915290912054610f7191610ff4565b600160a060020a0390921660009081526003602052604090209190915550565b6000903b1190565b600160a060020a0382161515610fae57600080fd5b610fb88282610c54565b6040518190600160a060020a038416906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b818101828110156107f357fe5b8154818355818111156107785760008381526020902061077891810190830161055f91905b8082111561103a5760008155600101611026565b50905600a165627a7a723058205b9c52212ecf7b84aff1e47eaa8892cdd6608ae21bef25c19c8b9fd30468f5a80029`

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

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCaller) InterfaceIdERC165(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _MainnetCryptoCardsContract.contract.Call(opts, out, "InterfaceId_ERC165")
	return *ret0, err
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) InterfaceIdERC165() ([4]byte, error) {
	return _MainnetCryptoCardsContract.Contract.InterfaceIdERC165(&_MainnetCryptoCardsContract.CallOpts)
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCallerSession) InterfaceIdERC165() ([4]byte, error) {
	return _MainnetCryptoCardsContract.Contract.InterfaceIdERC165(&_MainnetCryptoCardsContract.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MainnetCryptoCardsContract.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MainnetCryptoCardsContract.Contract.BalanceOf(&_MainnetCryptoCardsContract.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MainnetCryptoCardsContract.Contract.BalanceOf(&_MainnetCryptoCardsContract.CallOpts, _owner)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCaller) Exists(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MainnetCryptoCardsContract.contract.Call(opts, out, "exists", _tokenId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) Exists(_tokenId *big.Int) (bool, error) {
	return _MainnetCryptoCardsContract.Contract.Exists(&_MainnetCryptoCardsContract.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCallerSession) Exists(_tokenId *big.Int) (bool, error) {
	return _MainnetCryptoCardsContract.Contract.Exists(&_MainnetCryptoCardsContract.CallOpts, _tokenId)
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
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MainnetCryptoCardsContract.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _MainnetCryptoCardsContract.Contract.GetApproved(&_MainnetCryptoCardsContract.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _MainnetCryptoCardsContract.Contract.GetApproved(&_MainnetCryptoCardsContract.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MainnetCryptoCardsContract.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _MainnetCryptoCardsContract.Contract.IsApprovedForAll(&_MainnetCryptoCardsContract.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _MainnetCryptoCardsContract.Contract.IsApprovedForAll(&_MainnetCryptoCardsContract.CallOpts, _owner, _operator)
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
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MainnetCryptoCardsContract.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _MainnetCryptoCardsContract.Contract.OwnerOf(&_MainnetCryptoCardsContract.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _MainnetCryptoCardsContract.Contract.OwnerOf(&_MainnetCryptoCardsContract.CallOpts, _tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MainnetCryptoCardsContract.contract.Call(opts, out, "supportsInterface", _interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _MainnetCryptoCardsContract.Contract.SupportsInterface(&_MainnetCryptoCardsContract.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _MainnetCryptoCardsContract.Contract.SupportsInterface(&_MainnetCryptoCardsContract.CallOpts, _interfaceId)
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
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCaller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MainnetCryptoCardsContract.contract.Call(opts, out, "tokenByIndex", _index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _MainnetCryptoCardsContract.Contract.TokenByIndex(&_MainnetCryptoCardsContract.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _MainnetCryptoCardsContract.Contract.TokenByIndex(&_MainnetCryptoCardsContract.CallOpts, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MainnetCryptoCardsContract.contract.Call(opts, out, "tokenOfOwnerByIndex", _owner, _index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _MainnetCryptoCardsContract.Contract.TokenOfOwnerByIndex(&_MainnetCryptoCardsContract.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _MainnetCryptoCardsContract.Contract.TokenOfOwnerByIndex(&_MainnetCryptoCardsContract.CallOpts, _owner, _index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MainnetCryptoCardsContract.contract.Call(opts, out, "tokenURI", _tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _MainnetCryptoCardsContract.Contract.TokenURI(&_MainnetCryptoCardsContract.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _MainnetCryptoCardsContract.Contract.TokenURI(&_MainnetCryptoCardsContract.CallOpts, _tokenId)
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
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.Contract.Approve(&_MainnetCryptoCardsContract.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.Contract.Approve(&_MainnetCryptoCardsContract.TransactOpts, _to, _tokenId)
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
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.Contract.SafeTransferFrom(&_MainnetCryptoCardsContract.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.Contract.SafeTransferFrom(&_MainnetCryptoCardsContract.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractTransactor) SetApprovalForAll(opts *bind.TransactOpts, _to common.Address, _approved bool) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.contract.Transact(opts, "setApprovalForAll", _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.Contract.SetApprovalForAll(&_MainnetCryptoCardsContract.TransactOpts, _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractTransactorSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.Contract.SetApprovalForAll(&_MainnetCryptoCardsContract.TransactOpts, _to, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.Contract.TransferFrom(&_MainnetCryptoCardsContract.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _MainnetCryptoCardsContract.Contract.TransferFrom(&_MainnetCryptoCardsContract.TransactOpts, _from, _to, _tokenId)
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
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId indexed uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*MainnetCryptoCardsContractApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _MainnetCryptoCardsContract.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MainnetCryptoCardsContractApprovalIterator{contract: _MainnetCryptoCardsContract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId indexed uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MainnetCryptoCardsContractApproval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _MainnetCryptoCardsContract.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
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
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*MainnetCryptoCardsContractApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _MainnetCryptoCardsContract.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &MainnetCryptoCardsContractApprovalForAllIterator{contract: _MainnetCryptoCardsContract.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MainnetCryptoCardsContractApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _MainnetCryptoCardsContract.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
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
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId indexed uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*MainnetCryptoCardsContractTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _MainnetCryptoCardsContract.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MainnetCryptoCardsContractTransferIterator{contract: _MainnetCryptoCardsContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId indexed uint256)
func (_MainnetCryptoCardsContract *MainnetCryptoCardsContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MainnetCryptoCardsContractTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _MainnetCryptoCardsContract.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
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
