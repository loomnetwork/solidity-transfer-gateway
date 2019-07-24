var SampleBEP2Token = artifacts.require("BNBToken");

let gatewayAddress = ""

module.exports = function(deployer, network) {
  if (network === 'plasma') return

  deployer.deploy(SampleBEP2Token, gatewayAddress);
};