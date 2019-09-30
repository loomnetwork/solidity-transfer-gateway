// Example command
// `truffle exec scripts/withdrawal_token.js --network asia1 --key ../test.key --token 0x73dF46Ded0D43CF5eA2A32A6cea9F9830326867B`
//
// There will be a warning about unsupported key. Just ignore it.
// `Warning: possible unsupported (undocumented in help) command line option: --key`

const yargs = require('yargs');
const { sciNot, loadDappAccount, fetchETHGatewayInfo } = require('./utils')
const { Address, Contracts } = require('loom-js')
const { TransferGateway, AddressMapper } = Contracts
const SampleERC20Token = artifacts.require('SampleERC20Token')

// option to get private key
const argv = yargs
  .option('key', {
    alias: 'k',
    description: 'dappchain key file',
    type: 'string'
  })
  .option('token', {
    description: 'token address in dappchain side e.g. 0x73dF46Ded0D43CF5eA2A32A6cea9F9830326867B',
    type: 'string'
  })
  .argv;

module.exports = async function (callback) {
  let client
  try {
     // account injected by truffle which originated from the given key file
     const accounts = await web3.eth.getAccounts()
     const aliceDapp = accounts[0]
     console.log('Alice dapp address', aliceDapp)
 
     let info = await fetchETHGatewayInfo(argv.network)
     const { DAppChainGatewayAddress, MainnetGatewayAddress } = info
     const gatewayDappAddress = Address.fromString(DAppChainGatewayAddress)
     const ethGatewayAddress = Address.fromString(`eth:${MainnetGatewayAddress}`)

     if (!argv.token) {
       throw new Error('token address on dappchain is required')
     }
     const tokenAddress = argv.token.toLowerCase()
     console.log("Token address", tokenAddress)
 
     // create Dappchain client
     if (!argv.key) {
       throw new Error('key file is required')
     }
     const dappClient = loadDappAccount(argv.network, argv.key)
     client = dappClient.client
     const aliceDappAddress = Address.fromString(`${client.chainId}:${aliceDapp}`)
 
     // create address mapper
     const addressMapper = await AddressMapper.createAsync(client, aliceDappAddress)
     const hasMapping = await addressMapper.hasMappingAsync(aliceDappAddress)
     if (!hasMapping) {
       throw new Error(`no mapping from ${aliceDapp} to mainnet address`)
     }
     const addrMapping = await addressMapper.getMappingAsync(aliceDappAddress)
     const aliceMainnetAddr = addrMapping.to
     console.log('Alice mainnet address', aliceMainnetAddr.local.toString())

    const token = await SampleERC20Token.at(tokenAddress)
    let aliceBalance = await token.balanceOf(aliceDapp)
    console.log('Alice dapp token balance', aliceBalance.toString())
    let gatewayBalance = await token.balanceOf(gatewayDappAddress.local.toString())
    console.log('Gateway dapp token balance', gatewayBalance.toString())

    // create gateway contract
    const gatewayContract = await TransferGateway.createAsync(client, aliceDappAddress)

    const withdrawalReceipt = await gatewayContract.withdrawalReceiptAsync(aliceDappAddress)
    if (withdrawalReceipt) {
      throw new Error("pending withdrawal receipt exists")
    }

    const tokenContractAddress = Address.fromString(`${client.chainId}:${tokenAddress}`)
    let amount = sciNot(1, 18)
    console.log('Alice approving token', amount.toString(), "to Gateway")
    let tx = await token.approve(gatewayContract.address.local.toString(), amount)
    console.log("Alice approval s tx", tx)
    tx = await gatewayContract.withdrawERC20Async(amount, tokenContractAddress)
    console.log("Alice withdrawal token tx", tx)
    console.log(`run the command to get withdrawl receipt\n$LOOM_BIN gateway withdrawal-receipt ${aliceMainnetAddr.toString()} gateway`)

    callback()
  } catch (error) {
    callback(error)
  } finally {
    if (client) {
      client.disconnect()
    }
  }
}
