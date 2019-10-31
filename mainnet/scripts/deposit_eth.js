// This script deposits ETH to the mainnet Gateway contract
// 
// Example command
// `truffle exec scripts/deposit_eth.js --network rinkeby --dapp-network asia1
//
// There will be a warning about unsupported key. Just ignore it.
// `Warning: possible unsupported (undocumented in help) command line option: --dapp-network`

const yargs = require('yargs')
const Gateway = artifacts.require('Gateway')
const { sciNot } = require('./utils')
const { fetchETHGatewayInfo } = require('../../dappchain/scripts/utils')

const argv = yargs
  .option('dapp-network', {
    description: 'dappchain network e.g. local, asia1, us1, extdev',
    type: 'string'
  })
  .argv;

module.exports = async function (callback) {
  try {
    const accounts = await web3.eth.getAccounts()
    const alice = accounts[4]

    const dappNetwork = argv.dappNetwork || 'local'
    console.log("Using dapp network", dappNetwork)
    const info = await fetchETHGatewayInfo(dappNetwork)
    const { MainnetGatewayAddress } = info
    const gatewayAddress = MainnetGatewayAddress
    console.log("Gateway mainnet adress", gatewayAddress)

    const gateway = await Gateway.at(gatewayAddress)
    console.log("Alice address", alice)
    let balance = await web3.eth.getBalance(alice)
    console.log('Alice token balance', balance.toString(10))

    // CHANGE THE AMOUNT TO WHATEVER YOU WANT HERE
    let amount = sciNot(1, 18) // default to 1 x 10^18

    console.log('Alice depositing', amount.toString())
    tx = await gateway.send(amount, { from: alice })
    console.log('Alice deposited tx', tx)

    callback()
  } catch (error) {
    callback(error)
  }
}
