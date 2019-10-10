const fs = require('fs')
const path = require('path')

const ERC721XCards = artifacts.require('ERC721XCards')
const Gateway = artifacts.require('Gateway')


module.exports = async function (deployer, network, accounts) {
    if (network !== 'local_ganache') { return }
    console.log("gateway network address is ", Gateway.network.address)
    await deployer.deploy(ERC721XCards,  Gateway.network.address, "URI", {from: accounts[4]}) // dan deploys erc721x
    const erc721x = await ERC721XCards.deployed()
    erc721x.airdrop([1], [100], [Gateway.network.address], {from: accounts[4]})
    erc721x.airdrop([3], [50], [Gateway.network.address], {from: accounts[4]})

    let netId = await web3.eth.net.getId()
    let txHash = ERC721XCards['networks'][netId].transactionHash

    
    let logs = []
    logs.push(`mainnet_erc721x_cards_addr: "${erc721x.address}"`)
    logs.push(`mainnet_erc721x_cards_tx: "${txHash}"`)

    try {
        const outputDir = path.join(__dirname, `../../e2e_config/local_ganache`)
        if (!fs.existsSync(outputDir)) {
            fs.mkdirSync(outputDir)
        }
        fs.appendFileSync(path.join(outputDir, 'contracts.yml'), logs.join('\n'))
        fs.appendFileSync(path.join(outputDir, 'contracts.yml'), '\n')
        fs.writeFileSync(path.join(outputDir, 'mainnet_erc721x_cards_addr'), erc721x.address)
        fs.writeFileSync(path.join(outputDir, 'mainnet_erc721x_cards_tx'), txHash)

    } catch (err) {
        console.error(err)
    }

    console.log("\n*************************************************************************")
    console.log(logs.join('\n'))
    console.log("\n*************************************************************************\n")

}