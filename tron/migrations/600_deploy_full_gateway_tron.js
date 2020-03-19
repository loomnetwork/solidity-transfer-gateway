const fs = require('fs')
const path = require('path')

const Gateway = artifacts.require('Gateway')

const tronWeb = require('tronweb')

module.exports = async function (deployer, network, account) {
  const gatewayCreator = account 
  let validator = '0x423ba1feede3e69b0a01ed99a160a1d38ecc3cd1' // equvalent to TG1R35YRPQUwoMnWyvi6CPwD1YZrKkVGmE
  let hexAddr = tronWeb.address.toHex(validator)
  console.log(`Deploying Gateway from account: ${gatewayCreator}`)

  let loomAddr = '0x0000000000000000000000000000000000000000' // we don't have any trc20 address on prod yet
  await deployer.deploy(Gateway, loomAddr, validator, 3, 4, { from: gatewayCreator })
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
