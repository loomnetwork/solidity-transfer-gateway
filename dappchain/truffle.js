const { readFileSync } = require('fs')
const LoomTruffleProvider = require('loom-truffle-provider')

// change private key here
const privateKey = readFileSync('test_priv.key', 'utf-8')

function createProvider(provider, privateKey) {

  let providerConfig = {
    chainId: undefined,
    writeUrl: undefined,
    readUrl: undefined
  }

  if (provider === 'local') {
    providerConfig.chainId = 'default'
    providerConfig.writeUrl = 'http://127.0.0.1:46658/rpc'
    providerConfig.readUrl = 'http://127.0.0.1:46658/query'
  }

  const loomTruffleProvider = new LoomTruffleProvider(providerConfig.chainId, providerConfig.writeUrl, providerConfig.readUrl, privateKey)
  loomTruffleProvider.createExtraAccounts(10)
  return loomTruffleProvider
}

const testProvider = createProvider('local', privateKey)

module.exports = {
  // See <http://truffleframework.com/docs/advanced/configuration>
  // for more about customizing your Truffle configuration!
  networks: {
    local: {
      provider: testProvider,
      network_id: '*',
    },
    ganache:{
      host: '127.0.0.1',
      port: 7545,
      network_id: '*', // Match any network id
    }
  }
};
