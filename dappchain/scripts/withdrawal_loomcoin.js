// Example command
// `truffle exec scripts/withdrawal_loomcoin.js --network asia1 --key ../test.key`
//
// There will be a warning about unsupported key. Just ignore it.
// `Warning: possible unsupported (undocumented in help) command line option: --key`

const yargs = require('yargs');
const { sciNot, loadDappAccount, fetchLoomCoinGatewayInfo } = require('./utils')
const { Address, Contracts } = require('loom-js')
const { LoomCoinTransferGateway, Coin, AddressMapper } = Contracts

// option to get private key
const argv = yargs
  .option('key', {
    alias: 'k',
    description: 'dappchain key file',
    type: 'string'
  })
  .argv

module.exports = async function (callback) {
  let client
  try {
    // account injected by truffle which originated from the given key file
    const accounts = await web3.eth.getAccounts()
    const aliceDapp = accounts[0]
    console.log('Alice dapp address', aliceDapp)

    let info = await fetchLoomCoinGatewayInfo(argv.network)
    const { DAppChainGatewayAddress, MainnetGatewayAddress } = info
    const gatewayDappAddress = Address.fromString(DAppChainGatewayAddress)
    const ethGatewayAddress = Address.fromString(`eth:${MainnetGatewayAddress}`)

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

    // create coin contract
    const coin = await Coin.createAsync(client, aliceDappAddress)
    let aliceCoinBalance = await coin.getBalanceOfAsync(aliceDappAddress)
    console.log('Alice dapp Loom Coin balance', aliceCoinBalance.toString())
    let gatewayCoinBalance = await coin.getBalanceOfAsync(gatewayDappAddress)
    console.log('Gateway dapp Loom Coin balance', gatewayCoinBalance.toString())

    // create gateway contract
    const gatewayContract = await LoomCoinTransferGateway.createAsync(client, aliceDappAddress)

    const withdrawalReceipt = await gatewayContract.withdrawalReceiptAsync(aliceDappAddress)
    if (withdrawalReceipt) {
      throw new Error('pending withdrawal receipt exists')
    }

    // withdraw LoomCoin
    let amount = sciNot(1, 18)
    console.log('Alice approving Loom Coin', amount.toString(), "to Gateway")
    let err = await coin.approveAsync(gatewayContract.address, amount)
    if (err) {
      throw new Error(err)
    }
    let tx = await gatewayContract.withdrawLoomCoinAsync(amount, ethGatewayAddress)
    console.log("Alice withdrawal Loom tx", tx)
    console.log(`run the command to get withdrawl receipt\n$LOOM_BIN gateway withdrawal-receipt ${aliceMainnetAddr.toString()} loomcoin-gateway`)

    callback()
  } catch (error) {
    callback(error)
  } finally {
    if (client) {
      client.disconnect()
    }
  }
}
