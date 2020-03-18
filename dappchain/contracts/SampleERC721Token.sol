pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/token/ERC721/ERC721Token.sol";
import "./ERC721DAppToken.sol";

contract SampleERC721Token is ERC721DAppToken, ERC721Token {
    // Transfer Gateway contract address
    address public gateway;

    /**
     * @dev Constructor function
     */
    constructor(address _gateway) ERC721Token("SampleERC721Token", "SDT") public {
        gateway = _gateway;
    }

    function mintToGateway(uint256 _uid) public
    {
        require(msg.sender == gateway);
        _mint(gateway, _uid);
    }
}
