pragma solidity <0.6.0;

import "openzeppelin-solidity/contracts/token/ERC20/ERC20.sol";

contract LoomToken is ERC20 {
  string public name = "Loom Token";
  string public symbol = "LOOM";
  uint8 public decimals = 18;

  // one billion in initial supply
  uint256 public constant INITIAL_SUPPLY = 1000000000;

  constructor () public {
    _mint(msg.sender, INITIAL_SUPPLY * (10 ** uint256(decimals)));
  }
}
