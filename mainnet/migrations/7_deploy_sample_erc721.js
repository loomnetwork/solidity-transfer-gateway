const fs = require('fs')
const path = require('path')

const SampleERC721MintableToken = artifacts.require('SampleERC721MintableToken')
const Cards = artifacts.require('CryptoCards')

const Gateway = artifacts.require('Gateway')


module.exports = async function (deployer, network, accounts) {
    if (network !== 'local_ganache') { return }
    console.log("gateway network address is ", Gateway.network.address)
    let mainnetCryptoCard, sampleERC721mintableToken
    let mainnetCryptoCardTxHash, sampleERC721mintableTokenTxHash
    let netId = await web3.eth.net.getId()

    console.log("dan is ", accounts[4])

    await deployer.then(async () => {
        await deployer.deploy(Cards, Gateway.network.address, { from: accounts[4] }) // dan deploys erc721
        mainnetCryptoCard = await Cards.deployed()
        mainnetCryptoCardTxHash = Cards['networks'][netId].transactionHash 

        await deployer.deploy(SampleERC721MintableToken, Gateway.network.address, { from: accounts[4] }) // dan deploys erc721
        sampleERC721mintableToken = await SampleERC721MintableToken.deployed()
        sampleERC721mintableTokenTxHash = SampleERC721MintableToken['networks'][netId].transactionHash
    })

    
    let logs = []
    logs.push(`mainnet_crypto_cards_addr: "${mainnetCryptoCard.address}"`,
        `mainnet_crypto_cards_tx: "${mainnetCryptoCardTxHash}"`,
        `mainnet_erc721_mintable_token_addr: "${sampleERC721mintableToken.address}"`,
        `mainnet_erc721_mintable_token_tx: "${sampleERC721mintableTokenTxHash}"`)

    try {
        const outputDir = path.join(__dirname, `../../e2e_config/local_ganache`)
        if (!fs.existsSync(outputDir)) {
            fs.mkdirSync(outputDir)
        }
        fs.appendFileSync(path.join(outputDir, 'contracts.yml'), logs.join('\n'))
        fs.appendFileSync(path.join(outputDir, 'contracts.yml'), '\n')
        fs.writeFileSync(path.join(outputDir, 'mainnet_crypto_cards_addr'), mainnetCryptoCard.address)
        fs.writeFileSync(path.join(outputDir, 'mainnet_crypto_cards_tx'), mainnetCryptoCardTxHash)
        fs.writeFileSync(path.join(outputDir, 'mainnet_erc721_mintable_token_addr'), sampleERC721mintableToken.address)
        fs.writeFileSync(path.join(outputDir, 'mainnet_erc721_mintable_token_tx'), sampleERC721mintableTokenTxHash)

    } catch (err) {
        console.error(err)
    }

    console.log("\n*************************************************************************")
    console.log(logs.join('\n'))
    console.log("\n*************************************************************************\n")

}