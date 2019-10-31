const fs = require('fs')
const path = require('path')

const Gateway = artifacts.require('Gateway')
const ValidatorManagerContract = artifacts.require('ValidatorManagerContract')

module.exports = async function (deployer, network, accounts) {
  if (network === 'test') { return }
  const gatewayCreator = accounts[0]
  console.log(`Deploying Gateway from account: ${gatewayCreator}`)
  let vmcAddress

  if (network == 'local_ganache' || network == 'develop' || network == 'test' || network == 'rinkeby') {
    vmcAddress = ValidatorManagerContract.address
  } else {
    // Insert mainnet VMC contract address
    vmcAddress = "0xa4e8c3ec456107ea67d3075bf9e3df3a75823db0"
  }

  await deployer.deploy(Gateway, vmcAddress, { from: gatewayCreator })
  const gatewayInstance = await Gateway.deployed()

  const netId = await web3.eth.net.getId()
  const txHash = Gateway['networks'][netId].transactionHash
  const tx = await web3.eth.getTransaction(txHash)

  const logs = []
  logs.push(
    `mainnet_gateway_creator_addr: "${gatewayCreator}"`,
    `mainnet_gateway_addr: "${gatewayInstance.address}"`,
    `mainnet_gateway_blk: "${tx.blockNumber}"`,
  )

  try {
    const outputDir = path.join(__dirname, `../../e2e_config/${network}`)
    if (!fs.existsSync(outputDir)) {
      fs.mkdirSync(outputDir)
    }
    fs.appendFileSync(path.join(outputDir, 'contracts.yml'), logs.join('\n'))
    fs.appendFileSync(path.join(outputDir, 'contracts.yml'), '\n')
    fs.writeFileSync(path.join(outputDir, 'gateway_eth_addr'), gatewayInstance.address)
  } catch (err) {
    console.error(err)
  }

  // Create migration file from the deployed Gateway contract
  try {
    const outputDir = path.join(__dirname, `../../e2e_config/${network}`)
    if (!fs.existsSync(outputDir)) {
      fs.mkdirSync(outputDir)
    }
    const data = {
      'gateway_name': 'gateway',
      'mainnet_address': {
        'chain_id': 'eth',
        'local': Buffer.from(gatewayInstance.address.replace(/^0x/,''), 'hex').toString('base64')
      }
    }
    const migration2 = JSON.stringify(data)
    fs.writeFileSync(path.join(outputDir, 'migration-2.json'), migration2)
  } catch (err) {
    console.error(err)
  }

  console.log("\n*************************************************************************")
  console.log(logs.join('\n'))
  console.log("\n*************************************************************************\n")
}
