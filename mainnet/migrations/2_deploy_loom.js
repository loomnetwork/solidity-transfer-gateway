const fs = require('fs')
const path = require('path')
const mkdirp = require('mkdirp')

const Loom = artifacts.require('LoomToken')

module.exports = async function (deployer, network, accounts) {
    // only deploy to testnets
    let loomAddress, loomBlockNumber
    if (network !== 'mainnet') {
        const loomCreator = accounts[7] // Trudy deploys the loomtoken (we use that in the e2e tests)
        await deployer.deploy(Loom, { from: loomCreator })
        const loom = await Loom.deployed()

        let netId = await web3.eth.net.getId()
        let txHash = Loom['networks'][netId].transactionHash
        let tx = await web3.eth.getTransaction(txHash)
        loomAddress = loom.address
        loomBlockNumber = tx.blockNumber
    } else {
         // https://etherscan.io/address/0xa4e8c3ec456107ea67d3075bf9e3df3a75823db0
        loomAddress = "0xa4e8c3ec456107ea67d3075bf9e3df3a75823db0"
        loomBlockNumber = 6925676
    }

    const logs = []
    logs.push(
        `loomtoken_addr: "${loomAddress}"`,
        `loomtoken_blk: "${loomBlockNumber}"`,
    )

    try {
      const outputDir = path.join(__dirname, `../../e2e_config/${network}`)
      if (!fs.existsSync(outputDir)) {
        mkdirp.sync(outputDir)
      }
      fs.writeFileSync(path.join(outputDir, 'contracts.yml'), logs.join('\n'))
      fs.appendFileSync(path.join(outputDir, 'contracts.yml'), '\n')
      fs.writeFileSync(path.join(outputDir, 'loomcoin_eth_addr'), loomAddress)
    } catch (err) {
      console.error(err)
    }
    console.log("\n*************************************************************************")
    console.log(logs.join('\n'))
    console.log("\n*************************************************************************\n")
}
