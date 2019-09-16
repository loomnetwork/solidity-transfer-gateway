const SampleERC721MintableToken = artifacts.require('SampleERC721MintableToken')
const Gateway = artifacts.require('Gateway')


module.exports = async function (deployer, network, accounts) {
    if (network !== 'local_ganache') { return }
    console.log("gateway network address is ", Gateway.network.address)
    await deployer.deploy(SampleERC721MintableToken,  Gateway.network.address, {from: accounts[0]})
}