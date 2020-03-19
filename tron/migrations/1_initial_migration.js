var Migrations = artifacts.require('./Migrations.sol')

module.exports = function (deployer, network, accounts) {
  console.log(`Deploying initial migration from account: ${accounts}`)
  deployer.deploy(Migrations)
};
