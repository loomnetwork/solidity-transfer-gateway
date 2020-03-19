const fs = require('fs')
const path = require('path')

const Gateway = artifacts.require('Gateway')
const Loom = artifacts.require('LoomToken')

const tronWeb = require('tronweb')

module.exports = async function (deployer, network, accounts) {
  console.log("Accounts are", accounts)
  const gatewayCreator = accounts
  let validator = '0x423ba1feede3e69b0a01ed99a160a1d38ecc3cd1' // equvalent to TG1R35YRPQUwoMnWyvi6CPwD1YZrKkVGmE
  let hexAddr = tronWeb.address.toHex(validator)
  console.log(`Deploying Gateway from account: ${gatewayCreator}`)

  loomAddress = Loom.address
  await deployer.deploy(Gateway, loomAddress, validator, 3, 4, { from: gatewayCreator })
  const gatewayInstance = await Gateway.deployed()

  const logs = []
    logs.push(
    `mainnet_gateway_validator: "${validator}"`,
    `mainnet_gateway_hexAddr_validator: "${hexAddr}"`,
    `mainnet_gateway_creator_addr: "${gatewayCreator}"`,
    `mainnet_gateway_addr: "${gatewayInstance.address}"`,
  )
  
  try {
    const outputDir = path.join(__dirname, `../../e2e_config/${network}`)
    if (!fs.existsSync(outputDir)) {
      fs.mkdirSync(outputDir)
    }
    fs.appendFileSync(path.join(outputDir, 'contracts.yml'), logs.join('\n'))
    fs.writeFileSync(path.join(outputDir, 'gateway_trx_addr'), gatewayInstance.address)
  } catch (err) {
    console.error(err)
  }

  console.log("\n*************************************************************************")
  console.log(logs.join('\n'))
  console.log("\n*************************************************************************\n")

}
