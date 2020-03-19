const fs = require('fs')
const path = require('path')

const Loom = artifacts.require('LoomToken')

module.exports = async function (deployer, network, account) {
    // only deploy to testnets
   
    const loomCreator = account
    await deployer.deploy(Loom, { from: loomCreator })
    const loom = await Loom.deployed()
    let loomAddress = loom.address

    const logs = []
    logs.push(
        `loomtoken_addr: "${loomAddress}"`,
    )

    try {
      const outputDir = path.join(__dirname, `../../e2e_config/${network}`)
      if (!fs.existsSync(outputDir)) {
        fs.mkdirSync(outputDir)
      }
      fs.writeFileSync(path.join(outputDir, 'contracts.yml'), logs.join('\n'))
      fs.appendFileSync(path.join(outputDir, 'contracts.yml'), '\n')
      fs.writeFileSync(path.join(outputDir, 'loomcoin_trx_addr'), loomAddress)
    } catch (err) {
      console.error(err)
    }
    console.log("\n*************************************************************************")
    console.log(logs.join('\n'))
    console.log("\n*************************************************************************\n")
}
