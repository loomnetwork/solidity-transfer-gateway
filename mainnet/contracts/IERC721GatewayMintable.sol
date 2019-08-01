pragma solidity ^0.5.2;

import "openzeppelin-solidity/contracts/token/ERC721/IERC721.sol";

/**
 * @title ERC721 interface for token contracts deployed to mainnet that let Ethereum Gateway mint the token.
 */
contract IERC721GatewayMintable is IERC721 {
    // Called by the Ethereum Gateway contract to mint tokens.
    //
    // NOTE: the Ethereum gateway will call this method unconditionally.
    function mintTo(address _to, uint256 _tokenId) public;
}
