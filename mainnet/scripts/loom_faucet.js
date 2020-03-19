const Loom = artifacts.require('LoomToken')
const Gateway = artifacts.require('ERC20Gateway')

module.exports = async function(callback) {
    try {
        const account = process.env.ACCOUNT
        const gatewayAddr = process.env.GATEWAY_ETH_ADDR
        const accounts = await web3.eth.getAccounts()



        const gateway = await Gateway.at(gatewayAddr)
        const loomAddr = await gateway.loomAddress()
        const instance = await Loom.at(loomAddr)
        const amount = 10 * 10e18
        console.log("Balance of sender", (await instance.balanceOf(accounts[7])).toString())
        console.log("Balance of receiver before", (await instance.balanceOf(account)).toString())

        await instance.transfer(account, amount.toString(), {from: accounts[7]})
        console.log("Balance of receiver after", (await instance.balanceOf(account)).toString())


    } catch (err) {
        callback(err)
    }
    callback()
}
