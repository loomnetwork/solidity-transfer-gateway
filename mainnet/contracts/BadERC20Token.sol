pragma solidity ^0.4.24;
import "openzeppelin-solidity/contracts/math/SafeMath.sol";

contract BadERC20Token {
  
  event Approval(address indexed owner, address indexed spender, uint256 value);
  event Transfer(address indexed from, address indexed to, uint256 value);
  string public name = "BadToken";
  string public symbol = "BTK";
  uint8 public decimals = 18;
  using SafeMath for uint256;

  uint256 public totalSupply_;
  mapping (address => uint256) public balances;
  mapping (address => mapping (address => uint256)) internal allowed;

  // one billion in initial supply
  uint256 public constant INITIAL_SUPPLY = 1000000000;

  constructor () public {
    totalSupply_ = INITIAL_SUPPLY * (10 ** uint256(decimals));
    balances[msg.sender] = totalSupply_;
  }

  function transfer(address _to, uint256 _value) public {
    require(_value <= balances[msg.sender]);
    require(_to != address(0));

    balances[msg.sender] = balances[msg.sender].sub(_value);
    balances[_to] = balances[_to].add(_value);
    emit Transfer(msg.sender, _to, _value);
  }

  function transferFrom(
    address _from,
    address _to,
    uint256 _value
  )
    public returns (bool)
  {
    require(_value <= balances[_from]);
    require(_value <= allowed[_from][msg.sender]);
    require(_to != address(0));

    balances[_from] = balances[_from].sub(_value);
    balances[_to] = balances[_to].add(_value);
    allowed[_from][msg.sender] = allowed[_from][msg.sender].sub(_value);
    emit Transfer(_from, _to, _value);
    return true;
  }

  function approve(address _spender, uint256 _value) public returns (bool) {
    allowed[msg.sender][_spender] = _value;
    emit Approval(msg.sender, _spender, _value);
    return true;
  }

  function allowance(
    address _owner,
    address _spender
   )
    public
    view
    returns (uint256)
  {
    return allowed[_owner][_spender];
  }

  function balanceOf(address _owner) public view returns (uint256) {
    return balances[_owner];
  }
}
