pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/token/ERC721/ERC721.sol";
import "openzeppelin-solidity/contracts/token/ERC721/ERC721Receiver.sol";

import "erc721x/contracts/Interfaces/ERC721X.sol";
import "erc721x/contracts/Interfaces/ERC721XReceiver.sol";

import "./ERC20Gateway.sol";

contract Gateway is ERC20Gateway, ERC721Receiver, ERC721XReceiver {

  event ETHReceived(address from, uint256 amount);
  event ERC721Received(address operator, address from, uint256 tokenId, address contractAddress, bytes data);
  event ERC721XReceived(address operator, address from, uint256 tokenId, uint256 amount, address contractAddress, bytes data);
  event ERC721XBatchReceived(address operator, address to, uint256[] tokenTypes, uint256[] amounts, address contractAddress, bytes data);

  constructor (
    address loomToken, address[] _validators, uint8 _threshold_num, uint8 _threshold_denom,
    address[] _accounts, uint256[] _nonces
  ) public ERC20Gateway(loomToken, _validators, _threshold_num, _threshold_denom, _accounts, _nonces) {
  }

  function withdrawERC721X(uint256 tokenId, uint256 amount, bytes sig, address contractAddress)
    external
    signedByValidator(
        createMessageWithdraw(
            keccak256(abi.encodePacked(tokenId, amount, contractAddress))),
            sig
    )
  {
    ERC721X(contractAddress).safeTransferFrom(address(this), msg.sender, tokenId, amount);
    emit TokenWithdrawn(msg.sender, TokenKind.ERC721X, contractAddress, amount);
  }


  function withdrawERC721(uint256 uid, bytes sig, address contractAddress)
    external
    signedByValidator(
        createMessageWithdraw(
            keccak256(abi.encodePacked(uid, contractAddress))),
            sig
    )
  {
    ERC721(contractAddress).safeTransferFrom(address(this),  msg.sender, uid);
    emit TokenWithdrawn(msg.sender, TokenKind.ERC721, contractAddress, uid);
  }

  function withdrawETH(uint256 amount, bytes sig)
    external
    signedByValidator(
        createMessageWithdraw(
            keccak256(abi.encodePacked(amount))),
            sig
    )
  {
    msg.sender.transfer(amount); // ensure it's not reentrant
    emit TokenWithdrawn(msg.sender, TokenKind.ETH, address(0), amount);
  }

  /// @dev Receiver functions for 1-step deposits to the gateway
  /// @param _from Address of the token owner
  /// @param _tokenId Id of the token or tokenId
  /// @param _amount Amount of tokens received
  /// @return ERC721X_RECEIVED bytes function signature
  function onERC721XReceived(
      address _operator,
      address _from,
      uint256 _tokenId,
      uint256 _amount,
      bytes _data
  )
    public
    returns (bytes4)
  {
    require(allowAnyToken || allowedTokens[msg.sender], "Not a valid token");
    emit ERC721XReceived(
        _operator,
        _from,
        _tokenId,
        _amount,
        msg.sender,
        _data
    );
    return ERC721X_RECEIVED;
  }

  function onERC721XBatchReceived(
          address _operator,
          address _from,
          uint256[] _types,
          uint256[] _amounts,
          bytes _data
  )
    public
    returns(bytes4) 
  {
    require(allowAnyToken || allowedTokens[msg.sender], "Not a valid token");
    uint256 length = _types.length;
    require(length == _amounts.length);
    emit ERC721XBatchReceived(
        _operator,
        _from,
        _types,
        _amounts,
        msg.sender,
        _data
    );
    return ERC721X_BATCH_RECEIVE_SIG;
  }

  function onERC721Received(address _operator, address _from, uint256 _uid, bytes _data)
    public
    returns (bytes4) 
  {
    require(allowAnyToken || allowedTokens[msg.sender], "Not a valid token");
    emit ERC721Received(_operator, _from, _uid, msg.sender, _data);
    return ERC721_RECEIVED;
  }

  function () external payable {
    emit ETHReceived(msg.sender, msg.value);
  }

  function getETH() external view returns (uint256) {
      return address(this).balance;
  }

  // Check if contract owns specific erc721 token
  function getERC721(uint256 uid, address contractAddress) external view returns (bool) {
    return ERC721(contractAddress).ownerOf(uid) == address(this);
  }

  // Returns ERC721 token by uid
  function getERC721X(uint256 tokenId, address contractAddress) external view returns (uint256) {
    return ERC721X(contractAddress).balanceOf(address(this), tokenId);
  }
}
