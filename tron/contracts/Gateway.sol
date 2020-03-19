pragma solidity ^0.4.23;

import "./TRC20Gateway.sol";

contract Gateway is TRC20Gateway {
    address validatorAddress;
    event TRXReceived(address from, uint256 amount);
    event TRXWithdrawn(address from, TokenKind kind, address contractAddr, uint256 value);

    constructor (address loomToken, address _validators, uint8 _threshold_num, uint8 _threshold_denom)
      public TRC20Gateway(loomToken, _validators, _threshold_num, _threshold_denom) {
        validatorAddress = _validators;
    }

    function withdrawTRX(uint256 amount, bytes32 r, bytes32 s, uint8 v)
      external
      signedByValidator(
          createMessageWithdraw(keccak256(abi.encodePacked(amount))),
          r,s,v
      )
    {
        msg.sender.transfer(amount); // ensure it's not reentrant
        emit TokenWithdrawn(msg.sender, TokenKind.TRX, this, amount);
    }

    function sendToken () external payable {
        emit TRXReceived(msg.sender, msg.value);
    }

    function getTRX() external view returns (uint256) {
        return address(this).balance;
    }
}
