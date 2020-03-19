pragma solidity ^0.5.7;

import "openzeppelin-solidity/contracts/token/ERC20/IERC20.sol";
import "openzeppelin-solidity/contracts/token/ERC20/SafeERC20.sol";
import "./ValidatorManagerContract.sol";

contract ERC20Gateway {
    using SafeERC20 for IERC20;

  /// @notice Event to log the withdrawal of a token from the Gateway.
  /// @param  owner Address of the entity that made the withdrawal.
  /// @param  kind The type of token withdrawn (ERC20/ERC721/ETH).
  /// @param  contractAddress Address of token contract the token belong to.
  /// @param  value For ERC721 this is the uid of the token, for ETH/ERC20 this is the amount.
  event TokenWithdrawn(address indexed owner, TokenKind kind, address contractAddress, uint256 value);

  /// @notice Event to log the deposit of a LOOM to the Gateway.
  /// @param  from Address of the entity that made the withdrawal.
  /// @param  amount The LOOM token amount that was deposited
  /// @param  loomCoinAddress Address of the LOOM token
  event LoomCoinReceived(address indexed from, uint256 amount, address loomCoinAddress);

  /// @notice Event to log the deposit of a ERC20 to the Gateway.
  /// @param  from Address of the entity that made the withdrawal.
  /// @param  amount The ERC20 token amount that was deposited
  /// @param  contractAddress Address of the ERC20 token
  event ERC20Received(address from, uint256 amount, address contractAddress);

  /// The LOOM token address
  address public loomAddress;

  /// The nonces per withdrawer to prevent replays
  mapping (address => uint256) public nonces;

  /// The Validator Manager Contract
  ValidatorManagerContract public vmc;

  /// Enum for the various types of each token to notify clients during
  /// deposits and withdrawals
  enum TokenKind {
    ETH,
    ERC20,
    ERC721,
    ERC721X,
    LoomCoin
  }

  /// @notice Initialize the contract with the VMC
  /// @param _vmc the validator manager contrct address
  constructor(ValidatorManagerContract _vmc) public {
    vmc = _vmc;
    loomAddress = vmc.loomAddress();
  }

  /// @notice Function to withdraw ERC20 tokens from the Gateway. Emits a
  /// ERC20Withdrawn event, or a LoomCoinWithdrawn event if the coin is LOOM
  /// token, according to the ValidatorManagerContract.
  /// @param  amount The amount being withdrawn
  /// @param  contractAddress The address of the token being withdrawn
  /// @param  _signersIndexes Array of indexes of the validator's signatures based on
  ///         the currently elected validators
  /// @param  _v Array of `v` values from the validator signatures
  /// @param  _r Array of `r` values from the validator signatures
  /// @param  _s Array of `s` values from the validator signatures
  function withdrawERC20(
      uint256 amount, 
      address contractAddress,
      uint256[] calldata _signersIndexes,
      uint8[] calldata _v,
      bytes32[] calldata _r,
      bytes32[] calldata _s
  )
    external
  {
    bytes32 message = createMessageWithdraw(
            "\x10Withdraw ERC20:\n",
            keccak256(abi.encodePacked(amount, contractAddress))
    );

    // Ensure enough power has signed the withdrawal
    vmc.checkThreshold(message, _signersIndexes, _v, _r, _s);

    // Replay protection
    nonces[msg.sender]++;

    IERC20(contractAddress).safeTransfer(msg.sender, amount);
    
    TokenKind kind = contractAddress == loomAddress ? TokenKind.LoomCoin : TokenKind.ERC20;
    emit TokenWithdrawn(msg.sender, kind, contractAddress, amount);

  }

  // Approve and Deposit function for 2-step deposits
  // Requires first to have called `approve` on the specified ERC20 contract
  function depositERC20(uint256 amount, address contractAddress) external {
    IERC20(contractAddress).safeTransferFrom(msg.sender, address(this), amount);

    emit ERC20Received(msg.sender, amount, contractAddress);
    if (contractAddress == loomAddress) {
        emit LoomCoinReceived(msg.sender, amount, contractAddress);
    }
  }

  function getERC20(address contractAddress) external view returns (uint256) {
      return IERC20(contractAddress).balanceOf(address(this));
  }

    /// @notice Creates the message hash that includes replay protection and
    ///         binds the hash to this contract only.
    /// @param  hash The hash of the message being signed
    /// @return A hash on the hash of the message
  function createMessageWithdraw(string memory prefix, bytes32 hash)
    internal
    view
    returns (bytes32)
  {
    return keccak256(
      abi.encodePacked(
        prefix,
        msg.sender,
        nonces[msg.sender],
        address(this),
        hash
      )
    );
  }
}
