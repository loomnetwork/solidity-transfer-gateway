pragma solidity <0.6.0;

import "./../ValidatorManagerContract.sol";

contract VMCMock is ValidatorManagerContract {

    constructor (address[] memory _validators, uint64[] memory _powers, uint8 _threshold_num, uint8 _threshold_denom, address _loomAddress)  
    ValidatorManagerContract(_validators, _powers, _threshold_num, _threshold_denom, _loomAddress)
public {   }

}
