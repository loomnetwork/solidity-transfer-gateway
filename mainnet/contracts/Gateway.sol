pragma solidity ^0.5.7;

import "openzeppelin-solidity/contracts/token/ERC721/IERC721.sol";
import "openzeppelin-solidity/contracts/token/ERC721/IERC721Receiver.sol";

import "erc721x/contracts/Interfaces/ERC721X.sol";
import "erc721x/contracts/Interfaces/ERC721XReceiver.sol";

import "./ERC20Gateway.sol";

import "./ValidatorManagerContract.sol";
import "./IERC721GatewayMintable.sol";

contract Gateway is ERC20Gateway, IERC721Receiver, ERC721XReceiver {

  /// @notice Event to log the deposit of ETH to the Gateway.
  /// @param  from Address of the entity that made the withdrawal.
  /// @param  amount The ETH amount that was deposited
  event ETHReceived(address from, uint256 amount);

  /// @notice Event to log the deposit of an ERC721 to the Gateway.
  /// @param  operator Address of the operator accordingto the erc721 standard
  /// @param  from Address of the entity that made the withdrawal.
  /// @param  tokenId The ERC721 token ID that was deposited
  /// @param  contractAddress Address of the ERC721 token
  /// @param  data Any extra data that was provided during the deposit
  event ERC721Received(address operator, address from, uint256 tokenId, address contractAddress, bytes data);

  /// @notice Event to log the deposit of an ERC721x to the Gateway.
  /// @param  operator Address of the operator accordingto the erc721 standard
  /// @param  from Address of the entity that made the withdrawal.
  /// @param  tokenId The ERC721x token ID that was deposited
  /// @param  amount The ERC721x amount that was deposited
  /// @param  contractAddress Address of the ERC721 token
  /// @param  data Any extra data that was provided during the deposit
  event ERC721XReceived(address operator, address from, uint256 tokenId, uint256 amount, address contractAddress, bytes data);

  /// @notice Event to log the batch deposit of multiple ERC721x to the Gateway.
  /// @param  operator Address of the operator accordingto the erc721 standard
  /// @param  to Address of the entity that made the withdrawal.
  /// @param  tokenTypes The ERC721x token IDs that were deposited
  /// @param  amounts The ERC721x token amounts that wereas deposited
  /// @param  contractAddress Address of the ERC721 token
  /// @param  data Any extra data that was provided during the deposit
  event ERC721XBatchReceived(address operator, address to, uint256[] tokenTypes, uint256[] amounts, address contractAddress, bytes data);

  /// @notice Initialize the contract with the VMC
  /// @param _vmc the validator manager contrct address
  constructor (ValidatorManagerContract _vmc)
    public ERC20Gateway(_vmc) {
  }


  /// @notice Function to withdraw ERC721X tokens from the Gateway.
  /// @param  tokenId The tokenId being withdrawn
  /// @param  amount The amount being withdrawn
  /// @param  contractAddress The address of the token being withdrawn
  /// @param  _signersIndexes Array of indexes of the validator's signatures based on
  ///         the currently elected validators
  /// @param  _v Array of `v` values from the validator signatures
  /// @param  _r Array of `r` values from the validator signatures
  /// @param  _s Array of `s` values from the validator signatures
  function withdrawERC721X(
      uint256 tokenId,
      uint256 amount,
      address contractAddress,
      uint256[] calldata _signersIndexes,
      uint8[] calldata _v,
      bytes32[] calldata _r,
      bytes32[] calldata _s
  )
    external
    gatewayEnabled
  {
    bytes32 message = createMessageWithdraw(
            "\x12Withdraw ERC721X:\n",
            keccak256(abi.encodePacked(tokenId, amount, contractAddress))
    );

    // Ensure enough power has signed the withdrawal
    vmc.checkThreshold(message, _signersIndexes, _v, _r, _s);

    // Replay protection
    nonces[msg.sender]++;

    ERC721X(contractAddress).safeTransferFrom(address(this), msg.sender, tokenId, amount);
    emit TokenWithdrawn(msg.sender, TokenKind.ERC721X, contractAddress, amount);

  }


  /// @notice Function to withdraw ERC721 tokens from the Gateway.
  /// If the given uid doesn't exist the Gateway will mint the uid directly to withdrawer.
  /// @param  uid The uid of the token being withdrawn
  /// @param  contractAddress The address of the token being withdrawn
  /// @param  _signersIndexes Array of indexes of the validator's signatures based on
  ///         the currently elected validators
  /// @param  _v Array of `v` values from the validator signatures
  /// @param  _r Array of `r` values from the validator signatures
  /// @param  _s Array of `s` values from the validator signatures
  function withdrawERC721(
      uint256 uid,
      address contractAddress,
      uint256[] calldata _signersIndexes,
      uint8[] calldata _v,
      bytes32[] calldata _r,
      bytes32[] calldata _s
  )
    external
    gatewayEnabled
  {
    bytes32 message = createMessageWithdraw(
            "\x11Withdraw ERC721:\n",
            keccak256(abi.encodePacked(uid, contractAddress))
    );

    // Ensure enough power has signed the withdrawal
    vmc.checkThreshold(message, _signersIndexes, _v, _r, _s);

    // Replay protection
    nonces[msg.sender]++;

    bool ok = tryERC721TransferFrom(address(this), msg.sender, uid, contractAddress);
    if (!ok) {
        IERC721GatewayMintable(contractAddress).mintTo(msg.sender, uid);
    }
    emit TokenWithdrawn(msg.sender, TokenKind.ERC721, contractAddress, uid);

  }

  /// @notice Function to withdraw ETH from the Gateway.
  /// @param  amount The amount being withdrawn
  /// @param  _signersIndexes Array of indexes of the validator's signatures based on
  ///         the currently elected validators
  /// @param  _v Array of `v` values from the validator signatures
  /// @param  _r Array of `r` values from the validator signatures
  /// @param  _s Array of `s` values from the validator signatures
  function withdrawETH(
      uint256 amount,
      uint256[] calldata _signersIndexes,
      uint8[] calldata _v,
      bytes32[] calldata _r,
      bytes32[] calldata _s
  )
    external
    gatewayEnabled
  {
    bytes32 message = createMessageWithdraw(
            "\x0eWithdraw ETH:\n",
            keccak256(abi.encodePacked(amount))
    );

    // Ensure enough power has signed the withdrawal
    vmc.checkThreshold(message, _signersIndexes, _v, _r, _s);

    // Replay protection
    nonces[msg.sender]++;

    msg.sender.transfer(amount); // not reentrant

    emit TokenWithdrawn(msg.sender, TokenKind.ETH, address(0), amount);
  }

  /// @dev Receiver function for ER721X 1-step deposits to the gateway
  /// @param _from Address of the token owner
  /// @param _tokenId Id of the token or tokenId
  /// @param _amount Amount of tokens received
  /// @return ERC721X_RECEIVED bytes function signature
  function onERC721XReceived(
      address _operator,
      address _from,
      uint256 _tokenId,
      uint256 _amount,
      bytes memory _data
  )
    public
    gatewayEnabled
    returns (bytes4)
  {
    require(isTokenAllowed(msg.sender), "Not a valid token");
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

  /// @dev Receiver function for ER721X 1-step batch-deposits to the gateway
  /// @param _from Address of the token owner
  /// @param _types Ids of the tokens
  /// @param _amounts Amounts of tokens received
  /// @return ERC721X_RECEIVED bytes function signature
  function onERC721XBatchReceived(
          address _operator,
          address _from,
          uint256[] memory _types,
          uint256[] memory _amounts,
          bytes memory _data
  )
    public
    gatewayEnabled
    returns(bytes4)
  {
    require(isTokenAllowed(msg.sender), "Not a valid token");
    uint256 length = _types.length;
    require(length == _amounts.length, "Array lengths do not match");
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

  /// @dev Receiver function for ER721 1-step deposits to the gateway
  /// @param _from Address of the token owner
  /// @param _uid Id of the token or tokenId
  /// @return ERC721_RECEIVED bytes function signature
  function onERC721Received(address _operator, address _from, uint256 _uid, bytes memory _data)
    public
    gatewayEnabled
    returns (bytes4)
  {
    require(isTokenAllowed(msg.sender), "Not a valid token");
    emit ERC721Received(_operator, _from, _uid, msg.sender, _data);
    return this.onERC721Received.selector;
  }

  /// @notice Fallback function just emits an event
  function () external gatewayEnabled payable {
    emit ETHReceived(msg.sender, msg.value);
  }

  function getETH() external view returns (uint256) {
      return address(this).balance;
  }

  // Check if contract owns specific erc721 token
  function getERC721(uint256 uid, address contractAddress) external view returns (bool) {
    return IERC721(contractAddress).ownerOf(uid) == address(this);
  }

  // Returns ERC721 token by uid
  function getERC721X(uint256 tokenId, address contractAddress) external view returns (uint256) {
    return ERC721X(contractAddress).balanceOf(address(this), tokenId);
  }

  // This part is obtained from https://blog.polymath.network/try-catch-in-solidity-handling-the-revert-exception-f53718f76047
  // to handle transaction revert in a try-catch like mechanism until try-catch is standardize.
  function tryERC721TransferFrom(address _from, address _to, uint256 _uid, address _contractAddress) internal returns(bool) {
    (bool success, ) = address(_contractAddress).call(
                            abi.encodeWithSelector(
                                IERC721(_contractAddress).transferFrom.selector, _from, _to, _uid)
                            );
    return success;
  }

}