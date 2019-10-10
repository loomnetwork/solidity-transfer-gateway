const fs = require('fs')
const path = require('path')

const SampleBEP2Token = artifacts.require("BNBToken")


const gatewayAddr = process.env.GATEWAY_ADDR
console.log("gatewayaddr", gatewayAddr)

module.exports = async function (deployer, network, accounts) {
    if (network !== 'e2e') { return }
    let  sampleBEP2Token, sampleBnbToken
    try {
        await deployer.then(async () => {
          console.log("trudy dapp is ", accounts[4])
          await deployer.deploy(SampleBEP2Token, gatewayAddr, {from: accounts[4]}); // trudy deploys erc20
          sampleBEP2Token = await SampleBEP2Token.deployed()
    
          await deployer.deploy(SampleBEP2Token, gatewayAddr, {from: accounts[4]}); // trudy deploys erc20
          sampleBnbToken = await SampleBEP2Token.deployed()
        })

        
        let logs = []
        logs.push(`loomchain_bep2_token_addr: "${sampleBEP2Token.address}"`,
        `loomchain_bnb_token_addr: "${sampleBnbToken.address}"`,
        )
        
        const outputDir = path.join(__dirname, `../../e2e_config/local_bnbtestnet`)
        if (!fs.existsSync(outputDir)) {
            fs.mkdirSync(outputDir)
        }
        fs.appendFileSync(path.join(outputDir, 'contracts.yml'), logs.join('\n'))
        fs.appendFileSync(path.join(outputDir, 'contracts.yml'), '\n')
        fs.writeFileSync(path.join(outputDir, 'loomchain_bep2_token_addr'), sampleBEP2Token.address)
        fs.writeFileSync(path.join(outputDir, 'loomchain_bnb_token_addr'), sampleBnbToken.address)
    } catch (err){
    console.log(err)
  }
}
