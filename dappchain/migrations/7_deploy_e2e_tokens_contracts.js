const fs = require('fs')
const path = require('path')
const SampleERC20Token = artifacts.require("SampleERC20Token");
const SampleERC721xToken = artifacts.require("SampleERC721XToken");
const SampleERC721Token = artifacts.require("SampleERC721Token")

const gatewayAddr = process.env.GATEWAY_ADDR
console.log("gatewayaddr", gatewayAddr)
const MINT_AMOUNT = "1000000000000000000000000"

module.exports = async function (deployer, network, accounts) {
  if (network !== 'e2e') { return }

  let sampleERC20Token, sampleERC20Token2, erc721XToken, cryptoCardToken, sampleERC721Mintable, sampleBEP2Token, sampleBnbToken
  console.log("accounts", accounts)
  try {
    await deployer.then(async () => {
      console.log("trudy dapp is ", accounts[4])

      await deployer.deploy(SampleERC20Token, gatewayAddr, "SampleERC20Token", "SMPL", {from: accounts[4]}); // trudy deploys erc20
      sampleERC20Token = await SampleERC20Token.deployed()
      
      await deployer.deploy(SampleERC20Token, gatewayAddr, "SampleERC20Token2", "SMPL2", {from: accounts[4]}); // trudy deploys erc20
      sampleERC20Token2 = await SampleERC20Token.deployed()

      console.log("dan dapp is ", accounts[3])
      await deployer.deploy(SampleERC721xToken, gatewayAddr, {from: accounts[3]}); // dan deploys erc721
      erc721XToken = await SampleERC721xToken.deployed()

      await deployer.deploy(SampleERC721Token, gatewayAddr, {from: accounts[3]}); // dan deploys erc721
      cryptoCardToken = await SampleERC721Token.deployed()

      await deployer.deploy(SampleERC721Token, gatewayAddr, {from: accounts[3]}); // dan deploys erc721
      sampleERC721Mintable = await SampleERC721Token.deployed()      
    })

    
    let logs = []
    logs.push(`loomchain_SampleERC20Token_1: "${sampleERC20Token.address}"`, 
      `loomchain_SampleERC20Token_2: "${sampleERC20Token2.address}"`,
      `loomchain_SampleERC721XToken_1: "${erc721XToken.address}"`,
      `loomchain_crypto_cards_addr: "${cryptoCardToken.address}"`,
      `loomchain_erc721_mintable_token_addr: "${sampleERC721Mintable.address}"`,
    )


    try {
      const outputDir = path.join(__dirname, `../../e2e_config/local_ganache`)
      if (!fs.existsSync(outputDir)) {
        fs.mkdirSync(outputDir)
      }
      fs.appendFileSync(path.join(outputDir, 'contracts.yml'), logs.join('\n'))
      fs.appendFileSync(path.join(outputDir, 'contracts.yml'), '\n')
      fs.writeFileSync(path.join(outputDir, 'loomchain_SampleERC20Token_1'), sampleERC20Token.address)
      fs.writeFileSync(path.join(outputDir, 'loomchain_SampleERC20Token_2'), sampleERC20Token2.address)
      fs.writeFileSync(path.join(outputDir, 'loomchain_SampleERC721XToken_1'), erc721XToken.address)
      fs.writeFileSync(path.join(outputDir, 'loomchain_crypto_cards_addr'), cryptoCardToken.address)
      fs.writeFileSync(path.join(outputDir, 'loomchain_erc721_mintable_token_addr'), sampleERC721Mintable.address)
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
