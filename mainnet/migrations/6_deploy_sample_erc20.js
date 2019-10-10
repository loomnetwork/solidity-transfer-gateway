const fs = require('fs')
const path = require('path')

const SampleERC20MintableToken = artifacts.require('SampleERC20MintableToken')
const GameToken = artifacts.require('GameToken')
const Gateway = artifacts.require('Gateway')


module.exports = async function (deployer, network, accounts) {
    if (network !== 'local_ganache') { return }
    console.log("gateway network address is ", Gateway.network.address)
    let gatewayAddr = Gateway.network.address
    let netId = await web3.eth.net.getId()
    let gameToken, sampleERC20Token
    let gameTokenTxHash, sampleERC20Token2TxHash

    console.log("trudy is ", accounts[7])
    try {
        await deployer.then(async () => {
          await deployer.deploy(GameToken, {from: accounts[7]}); // trudy deploys erc20
          gameToken = await GameToken.deployed()
          gameTokenTxHash = GameToken['networks'][netId].transactionHash

          await deployer.deploy(SampleERC20MintableToken, gatewayAddr, {from: accounts[7]}); // trudy deploys erc20
          sampleERC20Token = await SampleERC20MintableToken.deployed()
          sampleERC20Token2TxHash = SampleERC20MintableToken['networks'][netId].transactionHash

        })


    
        
        let logs = []
        logs.push(`mainnet_game_token_addr: "${gameToken.address}"`,
                `mainnet_game_token_tx: "${gameTokenTxHash}"`,
                `mainnet_erc20_mintable_token_addr: "${sampleERC20Token.address}"`,
                `mainnet_erc20_mintable_token_tx: "${sampleERC20Token2TxHash}"`)

    
        try {
          const outputDir = path.join(__dirname, `../../e2e_config/local_ganache`)
          if (!fs.existsSync(outputDir)) {
            fs.mkdirSync(outputDir)
          }
          fs.appendFileSync(path.join(outputDir, 'contracts.yml'), logs.join('\n'))
          fs.appendFileSync(path.join(outputDir, 'contracts.yml'), '\n')
          fs.writeFileSync(path.join(outputDir, 'mainnet_game_token_addr'), gameToken.address)
          fs.writeFileSync(path.join(outputDir, 'mainnet_game_token_tx'), gameTokenTxHash)
          fs.writeFileSync(path.join(outputDir, 'mainnet_erc20_mintable_token_addr'), sampleERC20Token.address)
          fs.writeFileSync(path.join(outputDir, 'mainnet_erc20_mintable_token_tx'), sampleERC20Token2TxHash)
        } catch (err) {
          console.error(err)
        }
    
        console.log("\n*************************************************************************")
        console.log(logs.join('\n'))
        console.log("\n*************************************************************************\n")
    
      } catch (err){
        console.log(err)
      }
    


}