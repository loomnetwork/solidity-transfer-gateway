const ERC721XCards = artifacts.require('ERC721XCards')
const Gateway = artifacts.require('Gateway')


module.exports = async function (deployer, network, accounts) {
    if (network !== 'local_ganache') { return }
    console.log("gateway network address is ", Gateway.network.address)
    await deployer.deploy(ERC721XCards,  Gateway.network.address, "URI", {from: accounts[0]})
    const erc721x = await ERC721XCards.deployed()
    erc721x.airdrop([1], [100], [Gateway.network.address], {from: accounts[0]})
    erc721x.airdrop([3], [50], [Gateway.network.address], {from: accounts[0]})

}