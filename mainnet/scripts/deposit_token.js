// Example command
// `truffle exec scripts/withdrawal_token_multisig.js --network rinkeby --dapp-network asia1 --token 0x8dc9659ef712A4E234E07f2DD537F1e6206fd59f`
//
// There will be a warning about unsupported key. Just ignore it.
// `Warning: possible unsupported (undocumented in help) command line option: --dapp-network`

const yargs = require('yargs')
const Gateway = artifacts.require('Gateway')
const SampleERC20MintableToken = artifacts.require('SampleERC20MintableToken')
const { sciNot, delay } = require('./utils')
const { fetchETHGatewayInfo } = require('../../dappchain/scripts/utils')


const argv = yargs
  .option('dapp-network', {
    description: 'dappchain network e.g. local, asia1, us1, extdev',
    type: 'string'
  })
  .option('token', {
    description: 'mainnet token e.g. 0x8dc9659ef712A4E234E07f2DD537F1e6206fd59f',
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

    const tokenAddress = argv.token 
    if (!tokenAddress) {
      throw new Error("token address is required")
    }

    const gateway = await Gateway.at(gatewayAddress)
    const token = await SampleERC20MintableToken.at(tokenAddress)
    console.log("Alice address", alice)
    let balance = await token.balanceOf(alice)
    console.log('Alice token balance', balance.toString(10))
    let amount = sciNot(1, 18)

    console.log('Alice approving', amount.toString(), 'to gateway', gateway.address)
    let tx = await token.approve(gateway.address, amount, { from: alice })
    console.log('Alice approved tx', tx)
    await delay(5000)
    console.log('Alice depositing', amount.toString(), 'token address', token.address)
    tx = await gateway.depositERC20(amount, token.address, { from: alice })
    console.log('Alice deposited tx', tx)
    
    callback()
  } catch (error) {
    callback(error)
  }
}
