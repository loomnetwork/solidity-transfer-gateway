// Example command
// `truffle exec scripts/withdrawal_token_v2.js --network rinkeby --dapp-network asia1`
//
// There will be a warning about unsupported key. Just ignore it.
// `Warning: possible unsupported (undocumented in help) command line option: --dapp-network`


const yargs = require('yargs')
const { MessagePrefix, createWithdrawalHash, parseSigs } = require('./utils')
const { fetchETHGatewayInfo } = require('../../dappchain/scripts/utils')
const BN = web3.utils.BN

// option to get dapp network key
const argv = yargs
  .option('dapp-network', {
    description: 'dappchain network e.g. local, asia1, us1, extdev',
    type: 'string'
  })
  .argv;

const Gateway = artifacts.require('ERC20Gateway')
const ValidatorManagerContract = artifacts.require('ValidatorManagerContract')

module.exports = async function (callback) {
  let client
  try {
    // mainnet account
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
    const vmcAddr = await gateway.vmc()
    const vmc = await ValidatorManagerContract.at(vmcAddr)
    const vmValidators = await vmc.getValidators()
    const validators = vmValidators.map(val => val.toLowerCase())

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

    const prefix = MessagePrefix.ERC20Prefix
    const params = {
      prefix,
      gatewayAddress,
      amount,
      tokenAddress,
      nonce: await gateway.nonces(alice),
      userEthAddress: alice,
    }
    console.log('params:', params)
    let hash = createWithdrawalHash(params)
    console.log('hash:', hash)
    let sigHex = Buffer.from(sig, 'base64').toString('hex')
    console.log('sigHex:', sigHex)
    let payloads = parseSigs(sigHex, hash, validators)
    console.log('payload:', payloads)
    const { vs, rs, ss, valIndexes } = payloads

    let tx = await gateway.withdrawERC20(
      amount,
      tokenAddress,
      valIndexes,
      vs,
      rs,
      ss,
      { from: alice }
    )
    console.log('withdrawal tx', tx)

    callback()
  } catch (error) {
    callback(error)
  } finally {
    if (client) {
      client.disconnect()
    }
  }
}