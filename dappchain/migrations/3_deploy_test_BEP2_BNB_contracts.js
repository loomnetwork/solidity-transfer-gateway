var testBNBToken = artifacts.require("BNBToken");
const fs = require('fs')

addresses = fs.readFileSync('../bnb_account_addresses_test_migration.txt', 'utf8').split("\n")

let gatewayAddress = ""

const oldBNBAddr = ""

module.exports = (deployer, network) => {
  if (network === 'plasma') return
  deployer.deploy(testBNBToken, gatewayAddress).then(
    async () => {
      let newBNBinstance = await testBNBToken.deployed()
      if (process.env.MIGRATE === 'true'){
        let oldBNBinstance = await testBNBToken.at(oldBNBAddr)
        for(let i = 0; i < addresses.length; i++){
          oldBalance = await oldBNBinstance.balanceOf(addresses[i])
          console.log("address", addresses[i], "balance", oldBalance)
          if (oldBalance > 0) {
            const mintAmount = oldBalance.div('1e+10') // 1e+10 is 10 to power of 10
            await newBNBinstance.mint(addresses[i], mintAmount)
            let newBalance = await newBNBinstance.balanceOf(addresses[i])
            console.log("Address %s newBalance %s ",addresses[i], newBalance.toString())
          }
        }
      }
    }
  )
};
