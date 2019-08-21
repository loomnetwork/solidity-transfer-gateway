pragma solidity ^0.5.2;

import "openzeppelin-solidity/contracts/token/ERC721/ERC721Mintable.sol";
import "openzeppelin-solidity/contracts/token/ERC721/ERC721Full.sol";
import "./IERC721GatewayMintable.sol";

/**
 * @title ERC721 example for token contracts to be deployed to Ethereum.
 */
contract SampleERC721MintableToken is ERC721Mintable, ERC721Full, IERC721GatewayMintable {
    mapping (address => bool) gateways;
    string public name;
    string public symbol;
    mapping (address => bool) validators;

    event ValidatorAdded(address validator);
    event ValidatorRemoved(address validator);

    event GatewayAdded(address gateway);
    event GatewayRemoved(address gateway);

    constructor(address _gateway) public ERC721Full(name, symbol) {
        gateways[_gateway] = true;
        validators[msg.sender] = true;
        name = "ERC721Mint";
        symbol = "MNT721";
    }

    function mintTo(address _address, uint256 _tokenId) public onlyGateway {
        _mint(_address, _tokenId);
    }

    /**
     * @dev Override function to mint tokens
     * @param _to The address that will receive the minted tokens.
     * @param _tokenId The token id to mint.
     * @return A boolean that indicates if the operation was successful.
     */
    function mint(address _to, uint256 _tokenId) public onlyValidator returns (bool) {
        _mint(_to, _tokenId);
        return true;
    }

    function addValidator(address newValidator) onlyValidator public {
        validators[newValidator] = true;
        emit ValidatorAdded(newValidator);
    }

    function removeValidator(address validator) onlyValidator public {
        validators[validator] = false;
        emit ValidatorRemoved(validator);
    }

    modifier onlyValidator() {
        require(validators[msg.sender] == true, "only validators authorized to perform this action");
        _;
    }

    modifier onlyGateway(){
        require(gateways[msg.sender] == true, "only gateways are allowed mint");
        _;
    }

    function addGateway(address _gateway) onlyValidator public {
        gateways[_gateway] = true;
        emit GatewayAdded(_gateway);
    }

    function removeGateway(address _gateway) onlyValidator public {
        gateways[_gateway] = false;
        emit GatewayRemoved(_gateway);
    }
}