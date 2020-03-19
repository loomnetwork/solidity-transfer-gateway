pragma solidity ^0.4.23;

import "tron-contracts/contracts/tokens/TRC20/TRC20.sol";
import "./ValidatorManagerContract.sol";

contract TRC20Gateway is ValidatorManagerContract {

    event TRC20Received(address from, uint256 amount, address contractAddress);

    /**
    * Event to log the withdrawal of a token from the Gateway.
    * @param from Address of the entity that made the withdrawal.
    * @param kind The type of token withdrawn (TRC20/ERC721/TRX).
    * @param contractAddress Address of token contract the token belong to.
    * @param value For ERC721 this is the uid of the token, for TRX/TRC20 this is the amount.
    */
    event TokenWithdrawn(address indexed from, TokenKind kind, address contractAddress, uint256 value);
    event LoomCoinReceived(address indexed from, uint256 amount, address loomCoinAddress);

    bool public allowAnyToken = false;
    address public loomAddress;
    enum TokenKind {
        ETH,
        ERC20,
        ERC721,
        ERC721X,
        LoomCoin,
        TRX,
        TRC20
    }

    constructor (address _loomAddress, address _validators, uint8 _threshold_num, uint8 _threshold_denom)
      public ValidatorManagerContract(_validators, _threshold_num, _threshold_denom) {
        loomAddress = _loomAddress;
    }

    // Withdrawal functions
    function withdrawTRC20(uint256 amount, bytes32 r, bytes32 s, uint8 v, address contractAddress)
      external
      signedByValidator(
          createMessageWithdraw(
              keccak256(abi.encodePacked(amount, contractAddress))),
              r, s, v
      )
    {
        require(TRC20(contractAddress).transfer(msg.sender, amount), "Transfer failed");
        
        emit TokenWithdrawn(msg.sender, TokenKind.TRC20, contractAddress, amount);
    }

    // Approve and Deposit function for 2-step deposits
    // Requires first to have called `approve` on the specified TRC20 contract
    function depositTRC20(uint256 amount, address contractAddress) external {
        require(TRC20(contractAddress).transferFrom(msg.sender, address(this), amount), "Transfer failed");
        emit TRC20Received(msg.sender, amount, contractAddress);
        if (contractAddress == loomAddress) {
            emit LoomCoinReceived(msg.sender, amount, contractAddress);
        }
    }

    function getTRC20(address contractAddress) external view returns (uint256) {
        return TRC20(contractAddress).balanceOf(address(this));
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
