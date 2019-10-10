var SampleERc20Token = artifacts.require("SampleERC20Token");

let gatewayAddress = ""

module.exports = async function (deployer, network, accounts) {
  await deployer.deploy(SampleERc20Token, gatewayAddress);
};
