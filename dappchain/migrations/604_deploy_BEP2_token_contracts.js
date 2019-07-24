var SampleBEP2Token = artifacts.require("SampleBEP2Token");

// add gateway address here
let gatewayAddress = ""

module.exports = function(deployer, network) {
  if (network !== 'plasma') return

  deployer.deploy(SampleBEP2Token, gatewayAddress);
};