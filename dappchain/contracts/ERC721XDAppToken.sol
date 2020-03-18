pragma solidity ^0.4.24;

import "./ERC721X/Interfaces/ERC721X.sol";

/**
 * @title ERC721X interface for token contracts deployed to Loom DAppChains.
 */
contract ERC721XDAppToken is ERC721X {
    // Called by the DAppChain Gateway contract to mint ERC721X tokens that have been deposited to
    // the Mainnet gateway.
    //
    // NOTE: This function will only be called by the DAppChain Gateway contract if it doesn't own
    // the tokens it needs to transfer, so it's possible to omit this function if you wish to
    // manually allocate tokens for the Gateway.
    function mintToGateway(uint256 tokenId, uint256 amount) public;
}
