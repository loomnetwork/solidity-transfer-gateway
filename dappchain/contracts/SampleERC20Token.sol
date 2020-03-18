pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/token/ERC20/StandardToken.sol";
import "./ERC20DAppToken.sol";

contract SampleERC20Token is ERC20DAppToken, StandardToken {
    // Transfer Gateway contract address
    address public gateway;

    string public name = "SampleERC20Token";
    string public symbol = "DCC";
    uint8 public decimals = 18;
    
    /**
     * @dev Constructor function
     */
    constructor(address _gateway) public {
        gateway = _gateway;
        totalSupply_ = 1000000000 * (10 ** uint256(decimals));
        balances[_gateway] = totalSupply_;
    }

    function mintToGateway(uint256 _amount) public {
        require(msg.sender == gateway);
        totalSupply_ = totalSupply_.add(_amount);
        balances[gateway] = balances[gateway].add(_amount);
    }
}
