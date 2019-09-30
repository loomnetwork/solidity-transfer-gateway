// Example command
// `truffle exec scripts/withdrawal_token_v1.js --network rinkeby --dapp-network asia1`
//
// There will be a warning about unsupported key. Just ignore it.
// `Warning: possible unsupported (undocumented in help) command line option: --dapp-network`


const yargs = require('yargs')
const { soliditySha3 } = require('web3-utils')
const { fetchETHGatewayInfo } = require('../../dappchain/scripts/utils')
const BN = web3.utils.BN

const Gateway = artifacts.require('ERC20Gateway')

// option to get dapp network key
const argv = yargs
  .option('dapp-network', {
    description: 'dappchain network e.g. local, asia1, us1, extdev',
    type: 'string'
  })
  .argv;

module.exports = async function (callback) {
  try {
    const accounts = await web3.eth.getAccounts()
    const alice = accounts[3]
    console.log("Alice mainnet address", alice)

    const dappNetwork = argv.dappNetwork || 'local'
    console.log("Using dapp network", dappNetwork)
    const info = await fetchETHGatewayInfo(dappNetwork)
    const { MainnetGatewayAddress } = info
    const gatewayAddress = MainnetGatewayAddress
    console.log("Gateway mainnet adress", gatewayAddress)
    const gateway = await Gateway.at(gatewayAddress)

     // sig, amount, and token address can be retreived via `$LOOM_BIN gateway withdrawal-receipt` command
    // base64 signature
    const sig = ''
    // amount in big int
    const amount = new BN('1000000000000000000', 10)
    // token address in hex format 
    const tokenAddress = ''

    if (!sig || !amount || !tokenAddress) {
      throw new Error("sig, amount, and token tokenAddress are required")
    }

    let sigHex = Buffer.from(sig, 'base64').toString('hex')
    sigHex = '0x' + sigHex
    console.log('sigHex:', sigHex)
    
    let nonce = await await gateway.nonces(alice)
    let hash = soliditySha3(
      alice,
      nonce,
      gatewayAddress,
      soliditySha3(amount, tokenAddress)
    )
    console.log('hash:', hash)
    let tx = await gateway.withdrawERC20(amount, sigHex, tokenAddress, { from: alice });
    console.log('withdrawal tx', tx)

    callback()
  } catch (error) {
    callback(error)
  }
}