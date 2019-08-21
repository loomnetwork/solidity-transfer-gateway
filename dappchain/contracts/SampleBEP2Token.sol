pragma solidity 0.4.24;

import './BEP2TokenTemplate.sol';

contract SampleBEP2Token is BEP2TokenTemplate {
    string public name;
    string public symbol;
    constructor(address _gateway) BEP2TokenTemplate(_gateway) public {
        name = "sampleBEP2Token";
        symbol = "SBEP2";
    }

}
