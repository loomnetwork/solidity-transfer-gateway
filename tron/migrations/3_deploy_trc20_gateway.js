const fs = require('fs')
const path = require('path')

const TRC20Gateway = artifacts.require('TRC20Gateway')
const Loom = artifacts.require('LoomToken')

module.exports = async function (deployer, network, accounts) {
  const gatewayCreator = accounts
  let validators = accounts
  loomAddress = Loom.address

  await deployer.deploy(TRC20Gateway, loomAddress, validators, 3, 4, { from: gatewayCreator })
  const loomGatewayInstance = await TRC20Gateway.deployed()
  let logs = []
  
  logs.push(
    `mainnet_loomGateway_validator: "${validators}"`,
    `mainnet_loomGateway_creator_addr: "${gatewayCreator}"`,
    `mainnet_loomGateway_addr: "${loomGatewayInstance.address}"`,
  )

  try {
    const outputDir = path.join(__dirname, `../../e2e_config/${network}`)
    if (!fs.existsSync(outputDir)) {
      fs.mkdirSync(outputDir)
    }
    fs.appendFileSync(path.join(outputDir, 'contracts.yml'), logs.join('\n'))
    fs.appendFileSync(path.join(outputDir, 'contracts.yml'), '\n')
    fs.writeFileSync(path.join(outputDir, 'loomGateway_trc20_addr'), loomGatewayInstance.address)
  } catch (err) {
    console.error(err)
  }

  console.log("\n*************************************************************************")
  console.log(logs.join('\n'))
  console.log("\n*************************************************************************\n")
}
