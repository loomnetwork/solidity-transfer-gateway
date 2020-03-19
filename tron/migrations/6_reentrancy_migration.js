var ReentrancyExploit = artifacts.require('./ReentrancyExploit.sol')

module.exports = function (deployer, network, accounts) {
  if (network == 'development' || network == 'development2') {
    deployer.deploy(ReentrancyExploit)
  }
};

