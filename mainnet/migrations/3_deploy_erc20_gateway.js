const fs = require('fs')
const path = require('path')
const mkdirp = require('mkdirp')

const ERC20Gateway = artifacts.require('ERC20Gateway')
const Loom = artifacts.require('LoomToken')

module.exports = async function (deployer, network, accounts) {
  const gatewayCreator = accounts[0]
  let validators = []
  
  console.log(`Deploying Gateway from account: ${gatewayCreator}`)

  if (network == 'local_ganache' || network == 'develop' || network == 'test' ) {
    if (!process.env.ENABLE_HSM) {
      validators.push(accounts[9])
    } else {
      validators.push(accounts[9], process.env.HSM_ADDRESS)
    }
  } else {
    const secretsFile = process.env.SECRET_FILE
    let secrets = null

    if ((secretsFile == '') || !fs.existsSync(secretsFile)) {
      console.log("No secrets file found. Can't deploy Gateway without validators!")
    } else {
      secrets = JSON.parse(fs.readFileSync(secretsFile, "utf8"))
      validators = secrets.validators
    }
  }

    // Deploy the Loom ERC20 gateway
    let loomAddress
    if (network !== 'mainnet') {
        loomAddress = Loom.address
    } else {
        loomAddress = "0xa4e8c3ec456107ea67d3075bf9e3df3a75823db0"
    }
    await deployer.deploy(ERC20Gateway, loomAddress, validators, 3, 4, [], [], { from: gatewayCreator })
    const loomGatewayInstance = await ERC20Gateway.deployed()

    let netId = await web3.eth.net.getId()
    let txHash = ERC20Gateway['networks'][netId].transactionHash
    let tx = await web3.eth.getTransaction(txHash)

    let logs = []
    for (let i = 0; i < validators.length; i++) {
      logs.push(`mainnet_loomGateway_validator_${i}: "${validators[i]}"`)
    }
    logs.push(
      `mainnet_loomGateway_creator_addr: "${gatewayCreator}"`,
      `mainnet_loomGateway_addr: "${loomGatewayInstance.address}"`,
      `mainnet_loomGateway_blk: "${tx.blockNumber}"`,
    )

    try {
      const outputDir = path.join(__dirname, `../../e2e_config/${network}`)
      if (!fs.existsSync(outputDir)) {
        mkdirp.sync(outputDir)
      }
      fs.appendFileSync(path.join(outputDir, 'contracts.yml'), logs.join('\n'))
      fs.appendFileSync(path.join(outputDir, 'contracts.yml'), '\n')
      fs.writeFileSync(path.join(outputDir, 'loomGateway_eth_addr'), loomGatewayInstance.address)
    } catch (err) {
      console.error(err)
    }

    console.log("\n*************************************************************************")
    console.log(logs.join('\n'))
    console.log("\n*************************************************************************\n")
}
