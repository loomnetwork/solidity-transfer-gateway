require('babel-register');
require('babel-polyfill');

const fs = require('fs');
const HDWalletProvider = require('truffle-hdwallet-provider');

function loadSecrets() {
  let secretsFile = process.env.SECRET_FILE
  if (!secretsFile) {
    secretsFile = 'secrets.json'
  }
  if (!fs.existsSync(secretsFile)) {
    throw new Error("No secrets file found.")
  } else {
    return JSON.parse(fs.readFileSync(secretsFile, "utf8"))
  }
}

let ropstenProvider, kovanProvider = {}

const mochaGasSettings = {
  reporter: 'eth-gas-reporter',
  reporterOptions: {
    currency: 'USD',
    gasPrice: 3
  }
}

function provider(network) {
  const secrets = loadSecrets()
  if (!secrets.mnemonic) {
    throw new Error('Mnemonic not specified!')
  }
  if (!secrets.infuraAPIKey) {
    throw new Error('Infura API key not specified!')
  }
  return new HDWalletProvider(secrets.mnemonic, `https://${network}.infura.io/` + secrets.infuraAPIKey, 0, 10)
}

const mocha = process.env.GAS_REPORTER ? mochaGasSettings : {}

module.exports = {
    networks: {
        local_ganache: {
            network_id: '*',
            host: '127.0.0.1',
            port: 8545,
        },
        ropsten: {
            network_id: 3,
            provider: ropstenProvider,
            gas: 4700036,
        },
        kovan: {
            network_id: 42,
            provider: kovanProvider,
            gas: 6.9e6,
        },
        rinkeby: {
            provider: provider('rinkeby'),
            network_id: 4,
            gas: 6.9e6,
            gasPrice: 15000000001,
            skipDryRun: true
        },
        coverage: {
            host: "127.0.0.1",
            network_id: "*",
            port: 8555,
            gas: 0xffffffffff,
            gasPrice: 0x01
        },
        mainnet: {
            provider: provider('mainnet'),
            network_id: 1,
            skipDryRun: true
        }
  },
  build: {},
  mocha,
  solc: {
    optimizer: {
      enabled: true,
      runs: 200
    }
  },
}
