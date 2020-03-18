const { readFileSync } = require('fs')
const LoomTruffleProvider = require('loom-truffle-provider')

const chainId    = 'default'
const writeUrl   = 'http://127.0.0.1:46658/rpc'
const readUrl    = 'http://127.0.0.1:46658/query'
const privateKey = readFileSync('test_priv.key', 'utf-8')

const loomTruffleProvider = new LoomTruffleProvider(chainId, writeUrl, readUrl, privateKey)

module.exports = {
  // See <http://truffleframework.com/docs/advanced/configuration>
  // for more about customizing your Truffle configuration!
  networks: {
    loom_dapp_chain: {
      provider: loomTruffleProvider,
      network_id: '*'
    },
    development: {
      host: "127.0.0.1",
      port: 8545,
      network_id: "*" // Match any network id
    }
  }
};
