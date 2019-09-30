const BN = require('bn.js')
const axios = require('axios')
const Web3 = require('web3')
const fs = require('fs')
const {
  Client, NonceTxMiddleware, SignedTxMiddleware,
  LocalAddress, CryptoUtils, LoomProvider
} = require('loom-js')

// sciNot returns a big number whose value is m x 10^n
module.exports.sciNot = (m, n) => {
  let b = new BN('10', 10)
  b = b.pow(new BN(n))
  b = b.mul(new BN(m))
  return b
}

// delay delays in milliseconds
module.exports.delay = ms => new Promise(resolve => setTimeout(resolve, ms))

// loadDappAccount returns an object containing dapp provider info
module.exports.loadDappAccount = (provider, keyfile) => {
  const privateKeyStr = fs.readFileSync(keyfile, 'utf-8')
  const privateKey = CryptoUtils.B64ToUint8Array(privateKeyStr)
  const publicKey = CryptoUtils.publicKeyFromPrivateKey(privateKey)

  let providerConfig = {
    chainId: undefined,
    writeUrl: undefined,
    readUrl: undefined
  }

  if (provider === 'local') {
    providerConfig.chainId = 'default'
    providerConfig.writeUrl = 'http://127.0.0.1:46658/websocket'
    providerConfig.readUrl = 'http://127.0.0.1:46658/queryws'
  } else {
    throw new Error(`not support provider${provider}`)
  }

  const client = new Client(
    providerConfig.chainId,
    providerConfig.writeUrl,
    providerConfig.readUrl
  )

  client.txMiddleware = [
    new NonceTxMiddleware(publicKey, client),
    new SignedTxMiddleware(privateKey)
  ]
  client.on('error', msg => {
    console.error('PlasmaChain connection error', msg)
  })

  return {
    account: LocalAddress.fromPublicKey(publicKey).toString(),
    web3js: new Web3(new LoomProvider(client, privateKey)),
    client
  }
}

module.exports.fetchETHGatewayInfo = async (network) => {
  let url
  switch (network) {
    case 'local':
      url = "http://127.0.0.1:9998/status"
      break
    default:
      throw new Error(`unknown network ${network}`)
  }

  const { data } = await axios.get(url)
  return data
}


module.exports.fetchLoomCoinGatewayInfo = async (network) => {
  let url
  switch (network) {
    case 'local':
      url = "http://127.0.0.1:9997/status"
      break
    default:
      throw new Error(`unknown network ${network}`)
  }

  const { data } = await axios.get(url)
  return data
}