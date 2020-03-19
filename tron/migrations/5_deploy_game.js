const GameToken = artifacts.require('GameToken')

module.exports = function (deployer, network, accounts) {
  if (network == 'development' || network == 'development2') {
    deployer.deploy(GameToken)
  }
}
