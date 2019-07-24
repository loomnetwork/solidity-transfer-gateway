var SampleERc20Token = artifacts.require("SampleERC20Token");

let gatewayAddress = ""

module.exports = function(deployer) {
  deployer.deploy(SampleERc20Token, gatewayAddress);
};