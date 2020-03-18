pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/token/ERC20/ERC20.sol";
import "./ValidatorManagerContract.sol";

contract ERC20Gateway is ValidatorManagerContract {

  event ERC20Received(address from, uint256 amount, address contractAddress);

  /**
   * Event to log the withdrawal of a token from the Gateway.
   * @param owner Address of the entity that made the withdrawal.
   * @param kind The type of token withdrawn (ERC20/ERC721/ETH).
   * @param contractAddress Address of token contract the token belong to.
   * @param value For ERC721 this is the uid of the token, for ETH/ERC20 this is the amount.
   */
  event TokenWithdrawn(address indexed owner, TokenKind kind, address contractAddress, uint256 value);
  event LoomCoinReceived(address indexed from, uint256 amount, address loomCoinAddress);

  bool public allowAnyToken = false;
  address public loomAddress;
  enum TokenKind {
    ETH,
    ERC20,
    ERC721,
    ERC721X,
    LoomCoin
  }

  constructor (
    address _loomAddress, address[] _validators, uint8 _threshold_num, uint8 _threshold_denom,
    address[] _accounts, uint256[] _nonces
  ) public ValidatorManagerContract(_validators, _threshold_num, _threshold_denom, _accounts, _nonces) {
        loomAddress = _loomAddress;
  }

  // Withdrawal functions
  function withdrawERC20(uint256 amount, bytes sig, address contractAddress)
    external
    signedByValidator(
        createMessageWithdraw(
            keccak256(abi.encodePacked(amount, contractAddress))),
            sig
    )
  {
    require(safeTransfer(contractAddress, msg.sender, amount), "Transfer failed"); // revert() if result is false
    TokenKind kind = contractAddress == loomAddress ? TokenKind.LoomCoin : TokenKind.ERC20;
    emit TokenWithdrawn(msg.sender, kind, contractAddress, amount);
  }

  // Workaround for buggy ERC20 contracts like BNB.
  // See https://medium.com/coinmonks/missing-return-value-bug-at-least-130-tokens-affected-d67bf08521ca
  function safeTransfer(address _tokenAddress, address _to, uint256 _value) internal returns (bool result) {
      bytes memory msg = abi.encodeWithSignature("transfer(address,uint256)", _to, _value);
      uint msgSize = msg.length;

      assembly {
          // pre-set scratch space to all bits set
          mstore(0x00, 0xff)

          // note: this requires tangerine whistle compatible EVM
          if iszero(call(gas(), _tokenAddress, 0, add(msg, 0x20), msgSize, 0x00, 0x20)) { revert(0, 0) }
          
          switch mload(0x00)
          case 0xff {
              result := 1
          }
          case 0x01 {
              result := 1
          }
          case 0x00 {
              result := 0
          }
          default {
              revert(0, 0)
          }
        }
    }

  // Approve and Deposit function for 2-step deposits
  // Requires first to have called `approve` on the specified ERC20 contract
  function depositERC20(uint256 amount, address contractAddress) external {
    require(ERC20(contractAddress).transferFrom(msg.sender, address(this), amount), "Transfer failed");
    emit ERC20Received(msg.sender, amount, contractAddress);
    if (contractAddress == loomAddress) {
        emit LoomCoinReceived(msg.sender, amount, contractAddress);
    }
  }

  function getERC20(address contractAddress) external view returns (uint256) {
    return ERC20(contractAddress).balanceOf(address(this));
  }

  function toggleAllowAnyToken(bool _allow) public onlyValidator {
    allowAnyToken = _allow;
  }

  /// @dev Creates message to check the validator sign
  function createMessageWithdraw(bytes32 hash)
    internal
    view
    returns (bytes32)
  {
    return keccak256(
      abi.encodePacked(
        msg.sender,
        nonces[msg.sender],
        address(this),
        hash
      )
    );
  }
}
