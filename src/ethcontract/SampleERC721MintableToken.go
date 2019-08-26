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

// SampleERC721MintableTokenABI is the input ABI used to generate the binding from.
const SampleERC721MintableTokenABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_gateway\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"gateway\",\"type\":\"address\"}],\"name\":\"GatewayAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"gateway\",\"type\":\"address\"}],\"name\":\"GatewayRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"mintTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newValidator\",\"type\":\"address\"}],\"name\":\"addValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_gateway\",\"type\":\"address\"}],\"name\":\"addGateway\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_gateway\",\"type\":\"address\"}],\"name\":\"removeGateway\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SampleERC721MintableTokenBin is the compiled bytecode used for deploying new contracts.
const SampleERC721MintableTokenBin = `0x60806040523480156200001157600080fd5b5060405160208062001ac4833981018060405260208110156200003357600080fd5b5051600e805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015620000bf5780601f106200009357610100808354040283529160200191620000bf565b820191906000526020600020905b815481529060010190602001808311620000a157829003601f168201915b5050600f8054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815295509193509150830182828015620001515780601f10620001255761010080835404028352916020019162000151565b820191906000526020600020905b8154815290600101906020018083116200013357829003601f168201915b50505050508181620001706301ffc9a760e01b620002c960201b60201c565b620001886380ac58cd60e01b620002c960201b60201c565b62000199336200033660201b60201c565b620001b163780e9d6360e01b620002c960201b60201c565b8151620001c690600a90602085019062000414565b508051620001dc90600b90602084019062000414565b50620001f5635b5e139f60e01b620002c960201b60201c565b505050506001600160a01b0381166000908152600d602090815260408083208054600160ff19918216811790925533855260108452938290208054909416179092558151808301909252600a8083527f4552433732314d696e7400000000000000000000000000000000000000000000929091019182526200027a91600e9162000414565b506040805180820190915260068082527f4d4e5437323100000000000000000000000000000000000000000000000000006020909201918252620002c191600f9162000414565b5050620004b9565b7fffffffff000000000000000000000000000000000000000000000000000000008082161415620002f957600080fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b620003518160056200038860201b620014571790919060201c565b6040516001600160a01b038216907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a250565b6001600160a01b0381166200039c57600080fd5b620003ae8282620003de60201b60201c565b15620003b957600080fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b60006001600160a01b038216620003f457600080fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200045757805160ff191683800117855562000487565b8280016001018555821562000487579182015b82811115620004875782518255916020019190600101906200046a565b506200049592915062000499565b5090565b620004b691905b80821115620004955760008155600101620004a0565b90565b6115fb80620004c96000396000f3fe608060405234801561001057600080fd5b50600436106101735760003560e01c80634f6ccce7116100de578063983b2d5611610097578063aa271e1a11610071578063aa271e1a146104fd578063b88d4fde14610523578063c87b56dd146105e9578063e985e9c51461060657610173565b8063983b2d56146104a157806398650275146104c7578063a22cb465146104cf57610173565b80634f6ccce7146103ed5780636352211e1461040a57806368bb37951461042757806370a082311461044d5780638a885e351461047357806395d89b411461049957610173565b80632f745c59116101305780632f745c59146102e757806340a141ff1461031357806340c10f191461033957806342842e0e14610365578063449a52f81461039b5780634d238c8e146103c757610173565b806301ffc9a71461017857806306fdde03146101b3578063081812fc14610230578063095ea7b31461026957806318160ddd1461029757806323b872dd146102b1575b600080fd5b61019f6004803603602081101561018e57600080fd5b50356001600160e01b031916610634565b604080519115158252519081900360200190f35b6101bb610653565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101f55781810151838201526020016101dd565b50505050905090810190601f1680156102225780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61024d6004803603602081101561024657600080fd5b50356106e1565b604080516001600160a01b039092168252519081900360200190f35b6102956004803603604081101561027f57600080fd5b506001600160a01b038135169060200135610711565b005b61029f6107be565b60408051918252519081900360200190f35b610295600480360360608110156102c757600080fd5b506001600160a01b038135811691602081013590911690604001356107c5565b61029f600480360360408110156102fd57600080fd5b506001600160a01b0381351690602001356107e8565b6102956004803603602081101561032957600080fd5b50356001600160a01b0316610835565b61019f6004803603604081101561034f57600080fd5b506001600160a01b0381351690602001356108e3565b6102956004803603606081101561037b57600080fd5b506001600160a01b0381358116916020810135909116906040013561094c565b610295600480360360408110156103b157600080fd5b506001600160a01b038135169060200135610967565b610295600480360360208110156103dd57600080fd5b50356001600160a01b03166109e1565b61029f6004803603602081101561040357600080fd5b5035610a92565b61024d6004803603602081101561042057600080fd5b5035610ac6565b6102956004803603602081101561043d57600080fd5b50356001600160a01b0316610aee565b61029f6004803603602081101561046357600080fd5b50356001600160a01b0316610b9f565b6102956004803603602081101561048957600080fd5b50356001600160a01b0316610bd5565b6101bb610c83565b610295600480360360208110156104b757600080fd5b50356001600160a01b0316610cde565b610295610cfc565b610295600480360360408110156104e557600080fd5b506001600160a01b0381351690602001351515610d07565b61019f6004803603602081101561051357600080fd5b50356001600160a01b0316610d8b565b6102956004803603608081101561053957600080fd5b6001600160a01b0382358116926020810135909116916040820135919081019060808101606082013564010000000081111561057457600080fd5b82018360208201111561058657600080fd5b803590602001918460018302840111640100000000831117156105a857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610d9e945050505050565b6101bb600480360360208110156105ff57600080fd5b5035610dc4565b61019f6004803603604081101561061c57600080fd5b506001600160a01b0381358116916020013516610e77565b6001600160e01b03191660009081526020819052604090205460ff1690565b600e805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106d95780601f106106ae576101008083540402835291602001916106d9565b820191906000526020600020905b8154815290600101906020018083116106bc57829003601f168201915b505050505081565b60006106ec82610ea5565b6106f557600080fd5b506000908152600260205260409020546001600160a01b031690565b600061071c82610ac6565b9050806001600160a01b0316836001600160a01b0316141561073d57600080fd5b336001600160a01b038216148061075957506107598133610e77565b61076257600080fd5b60008281526002602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b6008545b90565b6107cf3382610ec2565b6107d857600080fd5b6107e3838383610f21565b505050565b60006107f383610b9f565b82106107fe57600080fd5b6001600160a01b038316600090815260066020526040902080548390811061082257fe5b9060005260206000200154905092915050565b3360009081526010602052604090205460ff16151560011461088b57604051600160e51b62461bcd02815260040180806020018281038252603181526020018061159f6031913960400191505060405180910390fd5b6001600160a01b038116600081815260106020908152604091829020805460ff19169055815192835290517fe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f19281900390910190a150565b3360009081526010602052604081205460ff16151560011461093957604051600160e51b62461bcd02815260040180806020018281038252603181526020018061159f6031913960400191505060405180910390fd5b6109438383610f40565b50600192915050565b6107e383838360405180602001604052806000815250610d9e565b336000908152600d602052604090205460ff1615156001146109d35760408051600160e51b62461bcd02815260206004820152601e60248201527f6f6e6c792067617465776179732061726520616c6c6f776564206d696e740000604482015290519081900360640190fd5b6109dd8282610f40565b5050565b3360009081526010602052604090205460ff161515600114610a3757604051600160e51b62461bcd02815260040180806020018281038252603181526020018061159f6031913960400191505060405180910390fd5b6001600160a01b038116600081815260106020908152604091829020805460ff19166001179055815192835290517fe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec38849879281900390910190a150565b6000610a9c6107be565b8210610aa757600080fd5b60088281548110610ab457fe5b90600052602060002001549050919050565b6000818152600160205260408120546001600160a01b031680610ae857600080fd5b92915050565b3360009081526010602052604090205460ff161515600114610b4457604051600160e51b62461bcd02815260040180806020018281038252603181526020018061159f6031913960400191505060405180910390fd5b6001600160a01b0381166000818152600d6020908152604091829020805460ff19166001179055815192835290517f7137528d21fb7b0b9462886348954009edac49570e27ef9dba8bb3a676fc11e29281900390910190a150565b60006001600160a01b038216610bb457600080fd5b6001600160a01b0382166000908152600360205260409020610ae890610f5d565b3360009081526010602052604090205460ff161515600114610c2b57604051600160e51b62461bcd02815260040180806020018281038252603181526020018061159f6031913960400191505060405180910390fd5b6001600160a01b0381166000818152600d6020908152604091829020805460ff19169055815192835290517f6ad07cb5676ab639fa3821550f0ec4379900d542e2d8bd8f0f8158d8bd352da29281900390910190a150565b600f805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106d95780601f106106ae576101008083540402835291602001916106d9565b610ce733610d8b565b610cf057600080fd5b610cf981610f61565b50565b610d0533610fa9565b565b6001600160a01b038216331415610d1d57600080fd5b3360008181526004602090815260408083206001600160a01b03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b6000610ae860058363ffffffff610ff116565b610da98484846107c5565b610db584848484611026565b610dbe57600080fd5b50505050565b6060610dcf82610ea5565b610dd857600080fd5b6000828152600c602090815260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529091830182828015610e6b5780601f10610e4057610100808354040283529160200191610e6b565b820191906000526020600020905b815481529060010190602001808311610e4e57829003601f168201915b50505050509050919050565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b6000908152600160205260409020546001600160a01b0316151590565b600080610ece83610ac6565b9050806001600160a01b0316846001600160a01b03161480610f095750836001600160a01b0316610efe846106e1565b6001600160a01b0316145b80610f195750610f198185610e77565b949350505050565b610f2c83838361115f565b610f36838261123f565b6107e38282611334565b610f4a8282611372565b610f548282611334565b6109dd81611413565b5490565b610f7260058263ffffffff61145716565b6040516001600160a01b038216907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a250565b610fba60058263ffffffff6114a316565b6040516001600160a01b038216907fe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb6669290600090a250565b60006001600160a01b03821661100657600080fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b600061103a846001600160a01b03166114eb565b61104657506001610f19565b604051600160e11b630a85bd0102815233600482018181526001600160a01b03888116602485015260448401879052608060648501908152865160848601528651600095928a169463150b7a029490938c938b938b939260a4019060208501908083838e5b838110156110c35781810151838201526020016110ab565b50505050905090810190601f1680156110f05780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b15801561111257600080fd5b505af1158015611126573d6000803e3d6000fd5b505050506040513d602081101561113c57600080fd5b50516001600160e01b031916600160e11b630a85bd010214915050949350505050565b826001600160a01b031661117282610ac6565b6001600160a01b03161461118557600080fd5b6001600160a01b03821661119857600080fd5b6111a1816114f1565b6001600160a01b03831660009081526003602052604090206111c29061152c565b6001600160a01b03821660009081526003602052604090206111e390611543565b60008181526001602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6001600160a01b03821660009081526006602052604081205461126990600163ffffffff61154c16565b600083815260076020526040902054909150808214611304576001600160a01b03841660009081526006602052604081208054849081106112a657fe5b906000526020600020015490508060066000876001600160a01b03166001600160a01b0316815260200190815260200160002083815481106112e457fe5b600091825260208083209091019290925591825260079052604090208190555b6001600160a01b038416600090815260066020526040902080549061132d906000198301611561565b5050505050565b6001600160a01b0390911660009081526006602081815260408084208054868652600784529185208290559282526001810183559183529091200155565b6001600160a01b03821661138557600080fd5b61138e81610ea5565b1561139857600080fd5b600081815260016020908152604080832080546001600160a01b0319166001600160a01b0387169081179091558352600390915290206113d790611543565b60405181906001600160a01b038416906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b600880546000838152600960205260408120829055600182018355919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b6001600160a01b03811661146a57600080fd5b6114748282610ff1565b1561147e57600080fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b6001600160a01b0381166114b657600080fd5b6114c08282610ff1565b6114c957600080fd5b6001600160a01b0316600090815260209190915260409020805460ff19169055565b3b151590565b6000818152600260205260409020546001600160a01b031615610cf957600090815260026020526040902080546001600160a01b0319169055565b805461153f90600163ffffffff61154c16565b9055565b80546001019055565b60008282111561155b57600080fd5b50900390565b8154818355818111156107e3576000838152602090206107e39181019083016107c291905b8082111561159a5760008155600101611586565b509056fe6f6e6c792076616c696461746f727320617574686f72697a656420746f20706572666f726d207468697320616374696f6ea165627a7a72305820373e7ddda8966dfb8e4e687f092144a947b04074aec64d2101a18eff96febe560029`

// DeploySampleERC721MintableToken deploys a new Ethereum contract, binding an instance of SampleERC721MintableToken to it.
func DeploySampleERC721MintableToken(auth *bind.TransactOpts, backend bind.ContractBackend, _gateway common.Address) (common.Address, *types.Transaction, *SampleERC721MintableToken, error) {
	parsed, err := abi.JSON(strings.NewReader(SampleERC721MintableTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SampleERC721MintableTokenBin), backend, _gateway)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SampleERC721MintableToken{SampleERC721MintableTokenCaller: SampleERC721MintableTokenCaller{contract: contract}, SampleERC721MintableTokenTransactor: SampleERC721MintableTokenTransactor{contract: contract}, SampleERC721MintableTokenFilterer: SampleERC721MintableTokenFilterer{contract: contract}}, nil
}

// SampleERC721MintableToken is an auto generated Go binding around an Ethereum contract.
type SampleERC721MintableToken struct {
	SampleERC721MintableTokenCaller     // Read-only binding to the contract
	SampleERC721MintableTokenTransactor // Write-only binding to the contract
	SampleERC721MintableTokenFilterer   // Log filterer for contract events
}

// SampleERC721MintableTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SampleERC721MintableTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleERC721MintableTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SampleERC721MintableTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleERC721MintableTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SampleERC721MintableTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleERC721MintableTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SampleERC721MintableTokenSession struct {
	Contract     *SampleERC721MintableToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SampleERC721MintableTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SampleERC721MintableTokenCallerSession struct {
	Contract *SampleERC721MintableTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// SampleERC721MintableTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SampleERC721MintableTokenTransactorSession struct {
	Contract     *SampleERC721MintableTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// SampleERC721MintableTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type SampleERC721MintableTokenRaw struct {
	Contract *SampleERC721MintableToken // Generic contract binding to access the raw methods on
}

// SampleERC721MintableTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SampleERC721MintableTokenCallerRaw struct {
	Contract *SampleERC721MintableTokenCaller // Generic read-only contract binding to access the raw methods on
}

// SampleERC721MintableTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SampleERC721MintableTokenTransactorRaw struct {
	Contract *SampleERC721MintableTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSampleERC721MintableToken creates a new instance of SampleERC721MintableToken, bound to a specific deployed contract.
func NewSampleERC721MintableToken(address common.Address, backend bind.ContractBackend) (*SampleERC721MintableToken, error) {
	contract, err := bindSampleERC721MintableToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SampleERC721MintableToken{SampleERC721MintableTokenCaller: SampleERC721MintableTokenCaller{contract: contract}, SampleERC721MintableTokenTransactor: SampleERC721MintableTokenTransactor{contract: contract}, SampleERC721MintableTokenFilterer: SampleERC721MintableTokenFilterer{contract: contract}}, nil
}

// NewSampleERC721MintableTokenCaller creates a new read-only instance of SampleERC721MintableToken, bound to a specific deployed contract.
func NewSampleERC721MintableTokenCaller(address common.Address, caller bind.ContractCaller) (*SampleERC721MintableTokenCaller, error) {
	contract, err := bindSampleERC721MintableToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SampleERC721MintableTokenCaller{contract: contract}, nil
}

// NewSampleERC721MintableTokenTransactor creates a new write-only instance of SampleERC721MintableToken, bound to a specific deployed contract.
func NewSampleERC721MintableTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*SampleERC721MintableTokenTransactor, error) {
	contract, err := bindSampleERC721MintableToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SampleERC721MintableTokenTransactor{contract: contract}, nil
}

// NewSampleERC721MintableTokenFilterer creates a new log filterer instance of SampleERC721MintableToken, bound to a specific deployed contract.
func NewSampleERC721MintableTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*SampleERC721MintableTokenFilterer, error) {
	contract, err := bindSampleERC721MintableToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SampleERC721MintableTokenFilterer{contract: contract}, nil
}

// bindSampleERC721MintableToken binds a generic wrapper to an already deployed contract.
func bindSampleERC721MintableToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SampleERC721MintableTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SampleERC721MintableToken *SampleERC721MintableTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SampleERC721MintableToken.Contract.SampleERC721MintableTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SampleERC721MintableToken *SampleERC721MintableTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.SampleERC721MintableTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SampleERC721MintableToken *SampleERC721MintableTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.SampleERC721MintableTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SampleERC721MintableToken *SampleERC721MintableTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SampleERC721MintableToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SampleERC721MintableToken *SampleERC721MintableTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SampleERC721MintableToken *SampleERC721MintableTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_SampleERC721MintableToken *SampleERC721MintableTokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleERC721MintableToken.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_SampleERC721MintableToken *SampleERC721MintableTokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SampleERC721MintableToken.Contract.BalanceOf(&_SampleERC721MintableToken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_SampleERC721MintableToken *SampleERC721MintableTokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SampleERC721MintableToken.Contract.BalanceOf(&_SampleERC721MintableToken.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(tokenId uint256) constant returns(address)
func (_SampleERC721MintableToken *SampleERC721MintableTokenCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SampleERC721MintableToken.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(tokenId uint256) constant returns(address)
func (_SampleERC721MintableToken *SampleERC721MintableTokenSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SampleERC721MintableToken.Contract.GetApproved(&_SampleERC721MintableToken.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(tokenId uint256) constant returns(address)
func (_SampleERC721MintableToken *SampleERC721MintableTokenCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SampleERC721MintableToken.Contract.GetApproved(&_SampleERC721MintableToken.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(owner address, operator address) constant returns(bool)
func (_SampleERC721MintableToken *SampleERC721MintableTokenCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SampleERC721MintableToken.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(owner address, operator address) constant returns(bool)
func (_SampleERC721MintableToken *SampleERC721MintableTokenSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SampleERC721MintableToken.Contract.IsApprovedForAll(&_SampleERC721MintableToken.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(owner address, operator address) constant returns(bool)
func (_SampleERC721MintableToken *SampleERC721MintableTokenCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SampleERC721MintableToken.Contract.IsApprovedForAll(&_SampleERC721MintableToken.CallOpts, owner, operator)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(account address) constant returns(bool)
func (_SampleERC721MintableToken *SampleERC721MintableTokenCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SampleERC721MintableToken.contract.Call(opts, out, "isMinter", account)
	return *ret0, err
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(account address) constant returns(bool)
func (_SampleERC721MintableToken *SampleERC721MintableTokenSession) IsMinter(account common.Address) (bool, error) {
	return _SampleERC721MintableToken.Contract.IsMinter(&_SampleERC721MintableToken.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(account address) constant returns(bool)
func (_SampleERC721MintableToken *SampleERC721MintableTokenCallerSession) IsMinter(account common.Address) (bool, error) {
	return _SampleERC721MintableToken.Contract.IsMinter(&_SampleERC721MintableToken.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SampleERC721MintableToken *SampleERC721MintableTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SampleERC721MintableToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SampleERC721MintableToken *SampleERC721MintableTokenSession) Name() (string, error) {
	return _SampleERC721MintableToken.Contract.Name(&_SampleERC721MintableToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SampleERC721MintableToken *SampleERC721MintableTokenCallerSession) Name() (string, error) {
	return _SampleERC721MintableToken.Contract.Name(&_SampleERC721MintableToken.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(tokenId uint256) constant returns(address)
func (_SampleERC721MintableToken *SampleERC721MintableTokenCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SampleERC721MintableToken.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(tokenId uint256) constant returns(address)
func (_SampleERC721MintableToken *SampleERC721MintableTokenSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SampleERC721MintableToken.Contract.OwnerOf(&_SampleERC721MintableToken.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(tokenId uint256) constant returns(address)
func (_SampleERC721MintableToken *SampleERC721MintableTokenCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SampleERC721MintableToken.Contract.OwnerOf(&_SampleERC721MintableToken.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(interfaceId bytes4) constant returns(bool)
func (_SampleERC721MintableToken *SampleERC721MintableTokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SampleERC721MintableToken.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(interfaceId bytes4) constant returns(bool)
func (_SampleERC721MintableToken *SampleERC721MintableTokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SampleERC721MintableToken.Contract.SupportsInterface(&_SampleERC721MintableToken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(interfaceId bytes4) constant returns(bool)
func (_SampleERC721MintableToken *SampleERC721MintableTokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SampleERC721MintableToken.Contract.SupportsInterface(&_SampleERC721MintableToken.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SampleERC721MintableToken *SampleERC721MintableTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SampleERC721MintableToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SampleERC721MintableToken *SampleERC721MintableTokenSession) Symbol() (string, error) {
	return _SampleERC721MintableToken.Contract.Symbol(&_SampleERC721MintableToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SampleERC721MintableToken *SampleERC721MintableTokenCallerSession) Symbol() (string, error) {
	return _SampleERC721MintableToken.Contract.Symbol(&_SampleERC721MintableToken.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(index uint256) constant returns(uint256)
func (_SampleERC721MintableToken *SampleERC721MintableTokenCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleERC721MintableToken.contract.Call(opts, out, "tokenByIndex", index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(index uint256) constant returns(uint256)
func (_SampleERC721MintableToken *SampleERC721MintableTokenSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _SampleERC721MintableToken.Contract.TokenByIndex(&_SampleERC721MintableToken.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(index uint256) constant returns(uint256)
func (_SampleERC721MintableToken *SampleERC721MintableTokenCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _SampleERC721MintableToken.Contract.TokenByIndex(&_SampleERC721MintableToken.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(owner address, index uint256) constant returns(uint256)
func (_SampleERC721MintableToken *SampleERC721MintableTokenCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleERC721MintableToken.contract.Call(opts, out, "tokenOfOwnerByIndex", owner, index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(owner address, index uint256) constant returns(uint256)
func (_SampleERC721MintableToken *SampleERC721MintableTokenSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _SampleERC721MintableToken.Contract.TokenOfOwnerByIndex(&_SampleERC721MintableToken.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(owner address, index uint256) constant returns(uint256)
func (_SampleERC721MintableToken *SampleERC721MintableTokenCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _SampleERC721MintableToken.Contract.TokenOfOwnerByIndex(&_SampleERC721MintableToken.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(tokenId uint256) constant returns(string)
func (_SampleERC721MintableToken *SampleERC721MintableTokenCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SampleERC721MintableToken.contract.Call(opts, out, "tokenURI", tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(tokenId uint256) constant returns(string)
func (_SampleERC721MintableToken *SampleERC721MintableTokenSession) TokenURI(tokenId *big.Int) (string, error) {
	return _SampleERC721MintableToken.Contract.TokenURI(&_SampleERC721MintableToken.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(tokenId uint256) constant returns(string)
func (_SampleERC721MintableToken *SampleERC721MintableTokenCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _SampleERC721MintableToken.Contract.TokenURI(&_SampleERC721MintableToken.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SampleERC721MintableToken *SampleERC721MintableTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleERC721MintableToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SampleERC721MintableToken *SampleERC721MintableTokenSession) TotalSupply() (*big.Int, error) {
	return _SampleERC721MintableToken.Contract.TotalSupply(&_SampleERC721MintableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SampleERC721MintableToken *SampleERC721MintableTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _SampleERC721MintableToken.Contract.TotalSupply(&_SampleERC721MintableToken.CallOpts)
}

// AddGateway is a paid mutator transaction binding the contract method 0x68bb3795.
//
// Solidity: function addGateway(_gateway address) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenTransactor) AddGateway(opts *bind.TransactOpts, _gateway common.Address) (*types.Transaction, error) {
	return _SampleERC721MintableToken.contract.Transact(opts, "addGateway", _gateway)
}

// AddGateway is a paid mutator transaction binding the contract method 0x68bb3795.
//
// Solidity: function addGateway(_gateway address) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenSession) AddGateway(_gateway common.Address) (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.AddGateway(&_SampleERC721MintableToken.TransactOpts, _gateway)
}

// AddGateway is a paid mutator transaction binding the contract method 0x68bb3795.
//
// Solidity: function addGateway(_gateway address) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenTransactorSession) AddGateway(_gateway common.Address) (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.AddGateway(&_SampleERC721MintableToken.TransactOpts, _gateway)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(account address) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenTransactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _SampleERC721MintableToken.contract.Transact(opts, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(account address) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.AddMinter(&_SampleERC721MintableToken.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(account address) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenTransactorSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.AddMinter(&_SampleERC721MintableToken.TransactOpts, account)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(newValidator address) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenTransactor) AddValidator(opts *bind.TransactOpts, newValidator common.Address) (*types.Transaction, error) {
	return _SampleERC721MintableToken.contract.Transact(opts, "addValidator", newValidator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(newValidator address) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenSession) AddValidator(newValidator common.Address) (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.AddValidator(&_SampleERC721MintableToken.TransactOpts, newValidator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(newValidator address) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenTransactorSession) AddValidator(newValidator common.Address) (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.AddValidator(&_SampleERC721MintableToken.TransactOpts, newValidator)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(to address, tokenId uint256) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SampleERC721MintableToken.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(to address, tokenId uint256) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.Approve(&_SampleERC721MintableToken.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(to address, tokenId uint256) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.Approve(&_SampleERC721MintableToken.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _tokenId uint256) returns(bool)
func (_SampleERC721MintableToken *SampleERC721MintableTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _SampleERC721MintableToken.contract.Transact(opts, "mint", _to, _tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _tokenId uint256) returns(bool)
func (_SampleERC721MintableToken *SampleERC721MintableTokenSession) Mint(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.Mint(&_SampleERC721MintableToken.TransactOpts, _to, _tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _tokenId uint256) returns(bool)
func (_SampleERC721MintableToken *SampleERC721MintableTokenTransactorSession) Mint(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.Mint(&_SampleERC721MintableToken.TransactOpts, _to, _tokenId)
}

// MintTo is a paid mutator transaction binding the contract method 0x449a52f8.
//
// Solidity: function mintTo(_address address, _tokenId uint256) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenTransactor) MintTo(opts *bind.TransactOpts, _address common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _SampleERC721MintableToken.contract.Transact(opts, "mintTo", _address, _tokenId)
}

// MintTo is a paid mutator transaction binding the contract method 0x449a52f8.
//
// Solidity: function mintTo(_address address, _tokenId uint256) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenSession) MintTo(_address common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.MintTo(&_SampleERC721MintableToken.TransactOpts, _address, _tokenId)
}

// MintTo is a paid mutator transaction binding the contract method 0x449a52f8.
//
// Solidity: function mintTo(_address address, _tokenId uint256) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenTransactorSession) MintTo(_address common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.MintTo(&_SampleERC721MintableToken.TransactOpts, _address, _tokenId)
}

// RemoveGateway is a paid mutator transaction binding the contract method 0x8a885e35.
//
// Solidity: function removeGateway(_gateway address) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenTransactor) RemoveGateway(opts *bind.TransactOpts, _gateway common.Address) (*types.Transaction, error) {
	return _SampleERC721MintableToken.contract.Transact(opts, "removeGateway", _gateway)
}

// RemoveGateway is a paid mutator transaction binding the contract method 0x8a885e35.
//
// Solidity: function removeGateway(_gateway address) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenSession) RemoveGateway(_gateway common.Address) (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.RemoveGateway(&_SampleERC721MintableToken.TransactOpts, _gateway)
}

// RemoveGateway is a paid mutator transaction binding the contract method 0x8a885e35.
//
// Solidity: function removeGateway(_gateway address) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenTransactorSession) RemoveGateway(_gateway common.Address) (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.RemoveGateway(&_SampleERC721MintableToken.TransactOpts, _gateway)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(validator address) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenTransactor) RemoveValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _SampleERC721MintableToken.contract.Transact(opts, "removeValidator", validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(validator address) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenSession) RemoveValidator(validator common.Address) (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.RemoveValidator(&_SampleERC721MintableToken.TransactOpts, validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(validator address) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenTransactorSession) RemoveValidator(validator common.Address) (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.RemoveValidator(&_SampleERC721MintableToken.TransactOpts, validator)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenTransactor) RenounceMinter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleERC721MintableToken.contract.Transact(opts, "renounceMinter")
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenSession) RenounceMinter() (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.RenounceMinter(&_SampleERC721MintableToken.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenTransactorSession) RenounceMinter() (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.RenounceMinter(&_SampleERC721MintableToken.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(from address, to address, tokenId uint256, _data bytes) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _SampleERC721MintableToken.contract.Transact(opts, "safeTransferFrom", from, to, tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(from address, to address, tokenId uint256, _data bytes) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.SafeTransferFrom(&_SampleERC721MintableToken.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(from address, to address, tokenId uint256, _data bytes) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.SafeTransferFrom(&_SampleERC721MintableToken.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(to address, approved bool) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, error) {
	return _SampleERC721MintableToken.contract.Transact(opts, "setApprovalForAll", to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(to address, approved bool) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.SetApprovalForAll(&_SampleERC721MintableToken.TransactOpts, to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(to address, approved bool) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenTransactorSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.SetApprovalForAll(&_SampleERC721MintableToken.TransactOpts, to, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokenId uint256) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SampleERC721MintableToken.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokenId uint256) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.TransferFrom(&_SampleERC721MintableToken.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokenId uint256) returns()
func (_SampleERC721MintableToken *SampleERC721MintableTokenTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SampleERC721MintableToken.Contract.TransferFrom(&_SampleERC721MintableToken.TransactOpts, from, to, tokenId)
}

// SampleERC721MintableTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SampleERC721MintableToken contract.
type SampleERC721MintableTokenApprovalIterator struct {
	Event *SampleERC721MintableTokenApproval // Event containing the contract specifics and raw log

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
func (it *SampleERC721MintableTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleERC721MintableTokenApproval)
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
		it.Event = new(SampleERC721MintableTokenApproval)
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
func (it *SampleERC721MintableTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleERC721MintableTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleERC721MintableTokenApproval represents a Approval event raised by the SampleERC721MintableToken contract.
type SampleERC721MintableTokenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, approved indexed address, tokenId indexed uint256)
func (_SampleERC721MintableToken *SampleERC721MintableTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*SampleERC721MintableTokenApprovalIterator, error) {

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

	logs, sub, err := _SampleERC721MintableToken.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SampleERC721MintableTokenApprovalIterator{contract: _SampleERC721MintableToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, approved indexed address, tokenId indexed uint256)
func (_SampleERC721MintableToken *SampleERC721MintableTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SampleERC721MintableTokenApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _SampleERC721MintableToken.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleERC721MintableTokenApproval)
				if err := _SampleERC721MintableToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// SampleERC721MintableTokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the SampleERC721MintableToken contract.
type SampleERC721MintableTokenApprovalForAllIterator struct {
	Event *SampleERC721MintableTokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *SampleERC721MintableTokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleERC721MintableTokenApprovalForAll)
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
		it.Event = new(SampleERC721MintableTokenApprovalForAll)
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
func (it *SampleERC721MintableTokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleERC721MintableTokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleERC721MintableTokenApprovalForAll represents a ApprovalForAll event raised by the SampleERC721MintableToken contract.
type SampleERC721MintableTokenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(owner indexed address, operator indexed address, approved bool)
func (_SampleERC721MintableToken *SampleERC721MintableTokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*SampleERC721MintableTokenApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SampleERC721MintableToken.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &SampleERC721MintableTokenApprovalForAllIterator{contract: _SampleERC721MintableToken.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(owner indexed address, operator indexed address, approved bool)
func (_SampleERC721MintableToken *SampleERC721MintableTokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *SampleERC721MintableTokenApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SampleERC721MintableToken.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleERC721MintableTokenApprovalForAll)
				if err := _SampleERC721MintableToken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// SampleERC721MintableTokenGatewayAddedIterator is returned from FilterGatewayAdded and is used to iterate over the raw logs and unpacked data for GatewayAdded events raised by the SampleERC721MintableToken contract.
type SampleERC721MintableTokenGatewayAddedIterator struct {
	Event *SampleERC721MintableTokenGatewayAdded // Event containing the contract specifics and raw log

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
func (it *SampleERC721MintableTokenGatewayAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleERC721MintableTokenGatewayAdded)
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
		it.Event = new(SampleERC721MintableTokenGatewayAdded)
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
func (it *SampleERC721MintableTokenGatewayAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleERC721MintableTokenGatewayAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleERC721MintableTokenGatewayAdded represents a GatewayAdded event raised by the SampleERC721MintableToken contract.
type SampleERC721MintableTokenGatewayAdded struct {
	Gateway common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGatewayAdded is a free log retrieval operation binding the contract event 0x7137528d21fb7b0b9462886348954009edac49570e27ef9dba8bb3a676fc11e2.
//
// Solidity: e GatewayAdded(gateway address)
func (_SampleERC721MintableToken *SampleERC721MintableTokenFilterer) FilterGatewayAdded(opts *bind.FilterOpts) (*SampleERC721MintableTokenGatewayAddedIterator, error) {

	logs, sub, err := _SampleERC721MintableToken.contract.FilterLogs(opts, "GatewayAdded")
	if err != nil {
		return nil, err
	}
	return &SampleERC721MintableTokenGatewayAddedIterator{contract: _SampleERC721MintableToken.contract, event: "GatewayAdded", logs: logs, sub: sub}, nil
}

// WatchGatewayAdded is a free log subscription operation binding the contract event 0x7137528d21fb7b0b9462886348954009edac49570e27ef9dba8bb3a676fc11e2.
//
// Solidity: e GatewayAdded(gateway address)
func (_SampleERC721MintableToken *SampleERC721MintableTokenFilterer) WatchGatewayAdded(opts *bind.WatchOpts, sink chan<- *SampleERC721MintableTokenGatewayAdded) (event.Subscription, error) {

	logs, sub, err := _SampleERC721MintableToken.contract.WatchLogs(opts, "GatewayAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleERC721MintableTokenGatewayAdded)
				if err := _SampleERC721MintableToken.contract.UnpackLog(event, "GatewayAdded", log); err != nil {
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

// SampleERC721MintableTokenGatewayRemovedIterator is returned from FilterGatewayRemoved and is used to iterate over the raw logs and unpacked data for GatewayRemoved events raised by the SampleERC721MintableToken contract.
type SampleERC721MintableTokenGatewayRemovedIterator struct {
	Event *SampleERC721MintableTokenGatewayRemoved // Event containing the contract specifics and raw log

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
func (it *SampleERC721MintableTokenGatewayRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleERC721MintableTokenGatewayRemoved)
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
		it.Event = new(SampleERC721MintableTokenGatewayRemoved)
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
func (it *SampleERC721MintableTokenGatewayRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleERC721MintableTokenGatewayRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleERC721MintableTokenGatewayRemoved represents a GatewayRemoved event raised by the SampleERC721MintableToken contract.
type SampleERC721MintableTokenGatewayRemoved struct {
	Gateway common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGatewayRemoved is a free log retrieval operation binding the contract event 0x6ad07cb5676ab639fa3821550f0ec4379900d542e2d8bd8f0f8158d8bd352da2.
//
// Solidity: e GatewayRemoved(gateway address)
func (_SampleERC721MintableToken *SampleERC721MintableTokenFilterer) FilterGatewayRemoved(opts *bind.FilterOpts) (*SampleERC721MintableTokenGatewayRemovedIterator, error) {

	logs, sub, err := _SampleERC721MintableToken.contract.FilterLogs(opts, "GatewayRemoved")
	if err != nil {
		return nil, err
	}
	return &SampleERC721MintableTokenGatewayRemovedIterator{contract: _SampleERC721MintableToken.contract, event: "GatewayRemoved", logs: logs, sub: sub}, nil
}

// WatchGatewayRemoved is a free log subscription operation binding the contract event 0x6ad07cb5676ab639fa3821550f0ec4379900d542e2d8bd8f0f8158d8bd352da2.
//
// Solidity: e GatewayRemoved(gateway address)
func (_SampleERC721MintableToken *SampleERC721MintableTokenFilterer) WatchGatewayRemoved(opts *bind.WatchOpts, sink chan<- *SampleERC721MintableTokenGatewayRemoved) (event.Subscription, error) {

	logs, sub, err := _SampleERC721MintableToken.contract.WatchLogs(opts, "GatewayRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleERC721MintableTokenGatewayRemoved)
				if err := _SampleERC721MintableToken.contract.UnpackLog(event, "GatewayRemoved", log); err != nil {
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

// SampleERC721MintableTokenMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the SampleERC721MintableToken contract.
type SampleERC721MintableTokenMinterAddedIterator struct {
	Event *SampleERC721MintableTokenMinterAdded // Event containing the contract specifics and raw log

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
func (it *SampleERC721MintableTokenMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleERC721MintableTokenMinterAdded)
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
		it.Event = new(SampleERC721MintableTokenMinterAdded)
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
func (it *SampleERC721MintableTokenMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleERC721MintableTokenMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleERC721MintableTokenMinterAdded represents a MinterAdded event raised by the SampleERC721MintableToken contract.
type SampleERC721MintableTokenMinterAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: e MinterAdded(account indexed address)
func (_SampleERC721MintableToken *SampleERC721MintableTokenFilterer) FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*SampleERC721MintableTokenMinterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SampleERC721MintableToken.contract.FilterLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &SampleERC721MintableTokenMinterAddedIterator{contract: _SampleERC721MintableToken.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: e MinterAdded(account indexed address)
func (_SampleERC721MintableToken *SampleERC721MintableTokenFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *SampleERC721MintableTokenMinterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SampleERC721MintableToken.contract.WatchLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleERC721MintableTokenMinterAdded)
				if err := _SampleERC721MintableToken.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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

// SampleERC721MintableTokenMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the SampleERC721MintableToken contract.
type SampleERC721MintableTokenMinterRemovedIterator struct {
	Event *SampleERC721MintableTokenMinterRemoved // Event containing the contract specifics and raw log

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
func (it *SampleERC721MintableTokenMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleERC721MintableTokenMinterRemoved)
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
		it.Event = new(SampleERC721MintableTokenMinterRemoved)
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
func (it *SampleERC721MintableTokenMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleERC721MintableTokenMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleERC721MintableTokenMinterRemoved represents a MinterRemoved event raised by the SampleERC721MintableToken contract.
type SampleERC721MintableTokenMinterRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: e MinterRemoved(account indexed address)
func (_SampleERC721MintableToken *SampleERC721MintableTokenFilterer) FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*SampleERC721MintableTokenMinterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SampleERC721MintableToken.contract.FilterLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &SampleERC721MintableTokenMinterRemovedIterator{contract: _SampleERC721MintableToken.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: e MinterRemoved(account indexed address)
func (_SampleERC721MintableToken *SampleERC721MintableTokenFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *SampleERC721MintableTokenMinterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SampleERC721MintableToken.contract.WatchLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleERC721MintableTokenMinterRemoved)
				if err := _SampleERC721MintableToken.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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

// SampleERC721MintableTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SampleERC721MintableToken contract.
type SampleERC721MintableTokenTransferIterator struct {
	Event *SampleERC721MintableTokenTransfer // Event containing the contract specifics and raw log

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
func (it *SampleERC721MintableTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleERC721MintableTokenTransfer)
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
		it.Event = new(SampleERC721MintableTokenTransfer)
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
func (it *SampleERC721MintableTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleERC721MintableTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleERC721MintableTokenTransfer represents a Transfer event raised by the SampleERC721MintableToken contract.
type SampleERC721MintableTokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, tokenId indexed uint256)
func (_SampleERC721MintableToken *SampleERC721MintableTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*SampleERC721MintableTokenTransferIterator, error) {

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

	logs, sub, err := _SampleERC721MintableToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SampleERC721MintableTokenTransferIterator{contract: _SampleERC721MintableToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, tokenId indexed uint256)
func (_SampleERC721MintableToken *SampleERC721MintableTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SampleERC721MintableTokenTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _SampleERC721MintableToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleERC721MintableTokenTransfer)
				if err := _SampleERC721MintableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// SampleERC721MintableTokenValidatorAddedIterator is returned from FilterValidatorAdded and is used to iterate over the raw logs and unpacked data for ValidatorAdded events raised by the SampleERC721MintableToken contract.
type SampleERC721MintableTokenValidatorAddedIterator struct {
	Event *SampleERC721MintableTokenValidatorAdded // Event containing the contract specifics and raw log

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
func (it *SampleERC721MintableTokenValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleERC721MintableTokenValidatorAdded)
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
		it.Event = new(SampleERC721MintableTokenValidatorAdded)
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
func (it *SampleERC721MintableTokenValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleERC721MintableTokenValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleERC721MintableTokenValidatorAdded represents a ValidatorAdded event raised by the SampleERC721MintableToken contract.
type SampleERC721MintableTokenValidatorAdded struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorAdded is a free log retrieval operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: e ValidatorAdded(validator address)
func (_SampleERC721MintableToken *SampleERC721MintableTokenFilterer) FilterValidatorAdded(opts *bind.FilterOpts) (*SampleERC721MintableTokenValidatorAddedIterator, error) {

	logs, sub, err := _SampleERC721MintableToken.contract.FilterLogs(opts, "ValidatorAdded")
	if err != nil {
		return nil, err
	}
	return &SampleERC721MintableTokenValidatorAddedIterator{contract: _SampleERC721MintableToken.contract, event: "ValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorAdded is a free log subscription operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: e ValidatorAdded(validator address)
func (_SampleERC721MintableToken *SampleERC721MintableTokenFilterer) WatchValidatorAdded(opts *bind.WatchOpts, sink chan<- *SampleERC721MintableTokenValidatorAdded) (event.Subscription, error) {

	logs, sub, err := _SampleERC721MintableToken.contract.WatchLogs(opts, "ValidatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleERC721MintableTokenValidatorAdded)
				if err := _SampleERC721MintableToken.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
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

// SampleERC721MintableTokenValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the SampleERC721MintableToken contract.
type SampleERC721MintableTokenValidatorRemovedIterator struct {
	Event *SampleERC721MintableTokenValidatorRemoved // Event containing the contract specifics and raw log

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
func (it *SampleERC721MintableTokenValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleERC721MintableTokenValidatorRemoved)
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
		it.Event = new(SampleERC721MintableTokenValidatorRemoved)
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
func (it *SampleERC721MintableTokenValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleERC721MintableTokenValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleERC721MintableTokenValidatorRemoved represents a ValidatorRemoved event raised by the SampleERC721MintableToken contract.
type SampleERC721MintableTokenValidatorRemoved struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: e ValidatorRemoved(validator address)
func (_SampleERC721MintableToken *SampleERC721MintableTokenFilterer) FilterValidatorRemoved(opts *bind.FilterOpts) (*SampleERC721MintableTokenValidatorRemovedIterator, error) {

	logs, sub, err := _SampleERC721MintableToken.contract.FilterLogs(opts, "ValidatorRemoved")
	if err != nil {
		return nil, err
	}
	return &SampleERC721MintableTokenValidatorRemovedIterator{contract: _SampleERC721MintableToken.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: e ValidatorRemoved(validator address)
func (_SampleERC721MintableToken *SampleERC721MintableTokenFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *SampleERC721MintableTokenValidatorRemoved) (event.Subscription, error) {

	logs, sub, err := _SampleERC721MintableToken.contract.WatchLogs(opts, "ValidatorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleERC721MintableTokenValidatorRemoved)
				if err := _SampleERC721MintableToken.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
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
