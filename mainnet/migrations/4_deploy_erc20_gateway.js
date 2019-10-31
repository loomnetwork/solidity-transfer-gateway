const fs = require('fs')
const path = require('path')

const ERC20Gateway = artifacts.require('ERC20Gateway')
const ValidatorManagerContract = artifacts.require('ValidatorManagerContract')

module.exports = async function (deployer, network, accounts) {
  if (network === 'test') { return }
  const gatewayCreator = accounts[0]
  console.log(`Deploying Loomgateway from account: ${gatewayCreator}`)
  let vmcAddress

  if (network == 'local_ganache' || network == 'develop' || network == 'test' || network == 'rinkeby') {
    vmcAddress = ValidatorManagerContract.address
  } else {
    // Insert mainnet VMC contract address
    vmcAddress = "0xa4e8c3ec456107ea67d3075bf9e3df3a75823db0"
  }

  await deployer.deploy(ERC20Gateway, vmcAddress, { from: gatewayCreator })
  const loomGatewayInstance = await ERC20Gateway.deployed()

  let netId = await web3.eth.net.getId()
  let txHash = ERC20Gateway['networks'][netId].transactionHash
  let tx = await web3.eth.getTransaction(txHash)

  let logs = []
  logs.push(
    `mainnet_loomGateway_creator_addr: "${gatewayCreator}"`,
    `mainnet_loomGateway_addr: "${loomGatewayInstance.address}"`,
    `mainnet_loomGateway_blk: "${tx.blockNumber}"`,
  )

  try {
    const outputDir = path.join(__dirname, `../../e2e_config/${network}`)
    if (!fs.existsSync(outputDir)) {
      fs.mkdirSync(outputDir)
    }
    fs.appendFileSync(path.join(outputDir, 'contracts.yml'), logs.join('\n'))
    fs.appendFileSync(path.join(outputDir, 'contracts.yml'), '\n')
    fs.writeFileSync(path.join(outputDir, 'loomGateway_eth_addr'), loomGatewayInstance.address)
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
      'gateway_name': 'loomcoin-gateway',
      'mainnet_address': {
        'chain_id': 'eth',
        'local': Buffer.from(loomGatewayInstance.address.replace(/^0x/, ''), 'hex').toString('base64')
      }
    }
    const migration3 = JSON.stringify(data)
    fs.writeFileSync(path.join(outputDir, 'migration-3.json'), migration3)
  } catch (err) {
    console.error(err)
  }

  console.log("\n*************************************************************************")
  console.log(logs.join('\n'))
  console.log("\n*************************************************************************\n")
}
