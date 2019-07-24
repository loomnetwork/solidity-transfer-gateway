pragma solidity 0.4.24;

import 'openzeppelin-solidity/contracts/token/ERC20/MintableToken.sol';

contract BNBToken is MintableToken {
    address public gateway;
    string public name = "BNBToken";
    string public symbol = "BNB";
    uint8 public constant decimals = 8; // Need to have exact 8 decimals because of BEP2 specifications
    uint256 public INITIAL_SUPPLY = 0 * (10 ** uint256(decimals));
    mapping (address => bool) validator;

    constructor(address _gateway) public {
        gateway = _gateway;
        validator[msg.sender] = true;
        totalSupply_ = INITIAL_SUPPLY;
        balances[msg.sender] = INITIAL_SUPPLY;
    }

    function mintToGateway(uint256 _amount) public {
        require(msg.sender == gateway, "only the gateway is allowed to mint");
        totalSupply_ = totalSupply_.add(_amount);
        balances[gateway] = balances[gateway].add(_amount);
    }

    // Overloaded `mint` function of Mintable token for onlyValidator
    function mint(address _to, uint256 _amount) onlyValidator canMint public returns (bool) {
        totalSupply_ = totalSupply_.add(_amount);
        balances[_to] = balances[_to].add(_amount);
        emit Mint(_to, _amount);
        emit Transfer(address(0), _to, _amount);
        return true;
    }

    function addValidator(address newValidator) onlyValidator public {
        validator[newValidator] = true;
    }

    function removeValidator(address _validator) onlyValidator public {
        validator[_validator] = false;
    }

    modifier onlyValidator() {
        require(validator[msg.sender] == true, "You don\'t have permission to action with this token");
        _;
    }

}
