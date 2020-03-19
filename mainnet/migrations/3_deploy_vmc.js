const fs = require('fs')
const path = require('path')

const Loom = artifacts.require('LoomToken')
const ValidatorManagerContract = artifacts.require('ValidatorManagerContract')
const MockVMC = artifacts.require('VMCMock')

module.exports = async function (deployer, network, accounts) {
  if (network === 'test') { return }
  const vmcCreator = accounts[0]
  let validators = []
  let powers =  []
  let threshold_num = 2
  let threshold_denom = 3
  
  console.log(`Deploying Validator Manager Contract from account: ${vmcCreator}`)

  console.log(accounts)

  const secretsFile = process.env.SECRET_FILE
  let secrets = null

  if ((secretsFile == '') || !fs.existsSync(secretsFile)) {
    console.log("No secrets file found. Can't deploy VMC without validators!")
  } else {
    secrets = JSON.parse(fs.readFileSync(secretsFile, "utf8"))
    validators = secrets.validators.map(v => v.address)
    powers = secrets.validators.map(v => v.power)
    threshold_num = secrets.threshold_num
    threshold_denom = secrets.threshold_denom
  }

  console.log(validators)

    // Deploy the VMC
    let loomAddress
    if (network !== 'mainnet') {
        loomAddress = Loom.address
    } else {
        loomAddress = '0xa4e8c3ec456107ea67d3075bf9e3df3a75823db0' // TODO: Hardcode mainnet address for VMC - this MUST be always the same
    }

    deployedVMC = network === 'local_ganache' ? MockVMC : ValidatorManagerContract
    await deployer.deploy(deployedVMC, validators, powers, threshold_num, threshold_denom, loomAddress, { from: vmcCreator })
    const validatorManagerContractInstance = await deployedVMC.deployed()

    let netId = await web3.eth.net.getId()
    let txHash = deployedVMC['networks'][netId].transactionHash
    let tx = await web3.eth.getTransaction(txHash)

    let logs = []
    for (let i = 0; i < validators.length; i++) {
      logs.push(`mainnet_validatorManagerContract_validator_${i}: "Address: ${validators[i]} / Power: ${powers[i]}"`)
    }
    logs.push(
      `mainnet_validatorManagerContract_creator_addr: "${vmcCreator}"`,
      `mainnet_validatorManagerContract_addr: "${validatorManagerContractInstance.address}"`,
      `mainnet_validatorManagerContract_blk: "${tx.blockNumber}"`,
    )

    try {
      const outputDir = path.join(__dirname, `../../e2e_config/${network}`)
      if (!fs.existsSync(outputDir)) {
        fs.mkdirSync(outputDir)
      }
      fs.appendFileSync(path.join(outputDir, 'contracts.yml'), logs.join('\n'))
      fs.appendFileSync(path.join(outputDir, 'contracts.yml'), '\n')
      fs.writeFileSync(path.join(outputDir, 'validatorManagerContract_eth_addr'), validatorManagerContractInstance.address)
    } catch (err) {
      console.error(err)
    }

    console.log("\n*************************************************************************")
    console.log(logs.join('\n'))
    console.log("\n*************************************************************************\n")
}
