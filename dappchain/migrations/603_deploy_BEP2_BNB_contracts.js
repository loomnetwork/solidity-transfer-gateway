var BNBToken = artifacts.require("BNBToken");
const fs = require('fs')

addresses = fs.readFileSync('../bnb_account_addresses_prod.txt', 'utf8').split("\n")

let gatewayAddress = "0x7E0DF5C9fF8898F0e1B4Af4D133Ef557A0641AA8" // dappchain bnb gateway address

const oldBNBAddr = "0xcf2851b1ad63d093238ea296524be8d7cd920e0b" // this is bnb coin in plasma chain

module.exports = function(deployer, network) {
  if (network !== 'plasma') return

  deployer.deploy(BNBToken, gatewayAddress).then(
    async () => {
      let oldBNBinstance = await BNBToken.at(oldBNBAddr)
      let newBNBinstance = await BNBToken.deployed()
      for(let i = 0; i < addresses.length; i++){
        oldBalance = await oldBNBinstance.balanceOf(addresses[i])
        console.log("address", addresses[i], "balance", oldBalance)
        if (oldBalance > 0) {
          newBNBinstance.mint(addresses[i], oldBalance)
          newBalance = await newBNBinstance.balanceOf(addresses[i])
          console.log("address", addresses[i], "balance", newBalance)
        }
      }
    }
  )
};