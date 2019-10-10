pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/token/ERC20/StandardToken.sol";
import "./ERC20DAppToken.sol";

contract SampleERC20Token is ERC20DAppToken, StandardToken {
    // Transfer Gateway contract address
    address public gateway;
    address public deployer;
    string public name;
    string public symbol;
    uint8 public decimals = 18;

    /**
     * @dev Constructor function
     */
    constructor(address _gateway, string _name, string _symbol) public {
        gateway = _gateway;
        totalSupply_ = 1000000000 * (10 ** uint256(decimals));
        balances[_gateway] = 1000000000 * (10 ** uint256(decimals));
        deployer = msg.sender;
        name = _name;
        symbol = _symbol;
    }

    function mintToGateway(uint256 _amount) public {
        require(msg.sender == gateway);
        totalSupply_ = totalSupply_.add(_amount);
        balances[gateway] = balances[gateway].add(_amount);
    }

    function mintTo(address _address, uint256 _amount) onlyDeployer public {
        totalSupply_ = totalSupply_.add(_amount);
        balances[_address] = balances[_address].add(_amount);
        emit Transfer(address(0), _address, _amount);
    }

    modifier onlyDeployer() {
        require(deployer == msg.sender, "not authorized to perform this action");
        _;
    }
}
