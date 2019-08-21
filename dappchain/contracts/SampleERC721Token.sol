pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/token/ERC721/ERC721Token.sol";
import "./ERC721DAppToken.sol";

contract SampleERC721Token is ERC721DAppToken, ERC721Token {
    // Transfer Gateway contract address
    address public gateway;
    address public deployer;


    /**
     * @dev Constructor function
     */
    constructor(address _gateway) ERC721Token("SampleERC721Token", "SDT") public {
        gateway = _gateway;
        deployer = msg.sender;
    }

    function mintToGateway(uint256 _uid) public
    {
        require(msg.sender == gateway);
        _mint(gateway, _uid);
    }

    function mintTo(address _address, uint256 _uid) onlyDeployer public {
        _mint(_address, _uid);
    }

    modifier onlyDeployer() {
        require(deployer == msg.sender, "not authorized to perform this action");
        _;
    }

}
