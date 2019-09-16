const fs = require('fs')
const path = require('path')

const SampleERC20MintableToken = artifacts.require('SampleERC20MintableToken')
const Gateway = artifacts.require('Gateway')


module.exports = async function (deployer, network, accounts) {
    if (network !== 'local_ganache') { return }
    console.log("gateway network address is ", Gateway.network.address)
    await deployer.deploy(SampleERC20MintableToken, Gateway.network.address, {from: accounts[0]})
}