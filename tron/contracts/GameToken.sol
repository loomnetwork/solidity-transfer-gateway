pragma solidity ^0.4.23;


import "tron-contracts/contracts/tokens/TRC20/TRC20.sol";
import "tron-contracts/contracts/tokens/TRC20/TRC20Detailed.sol";

contract GameToken is TRC20, TRC20Detailed {
    uint8 public constant DECIMALS = 18;
    uint256 public constant INITIAL_SUPPLY = 1000000000;

    /**
     * @dev Constructor that gives msg.sender all of existing tokens.
     */
    constructor () public TRC20Detailed("GameToken", "GTK", DECIMALS) {
        _mint(msg.sender, INITIAL_SUPPLY);
    }
}