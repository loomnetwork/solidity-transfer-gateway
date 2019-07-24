const assert = require('assert')
const {CryptoUtils, Client, LocalAddress, LoomProvider, Contracts, Address, SpeculativeNonceTxMiddleware, SignedTxMiddleware} = require('loom-js')
const Web3 = require('web3')
const BN = require('bn.js')
const bech32 = require('bech32')
const axios = require('axios');
const erc20abi = require('../dappchain/build/contracts/BNBToken.json')
const BnbApiClient = require('@binance-chain/javascript-sdk');

// This script expects the following env vars to be set:
// DAPP_ENV <- DAppChain to run against (local)
// BNB_COIN_ADDRESS_HEX = '0x...' <- the address of the BNB-ERC20 on the DAppChain
// BINANCE_WALLET_KEY <- hex-encoded private key of Binance containing funds for the test
// ALICE_DAPP_KEY <- base64 encoded ed25519 private key of DAppChain account that will be used in the test
// BINANCE_ADDRESS <- address of the Binance account which will be used to deposit funds to the DAppChain, e.g. tbnb14sa7gnlalxd0e336clc0ltgke6e6hdanyl6pqq

const AMOUNT_INT = 1300000 
const amount = 0.013; // amount float
const FEE = 37500

const message = 'loom8731f082aa764cf7c559e749892c40a4623b9ee4'; // supposed to be loom+addressOf(ALICE_DAPP_KEY)
let api
if (process.env.DAPP_ENV != 'plasma') {
  api = 'https://testnet-dex.binance.org/'; /// api string
} else {
  api = 'https://dex.binance.org/'; /// api string
}

const delay = ms => new Promise(resolve => setTimeout(resolve, ms));

const bnbClient = new BnbApiClient(api);

function decodeAddress(value) {
  const decodeAddress = bech32.decode(value)
  return Buffer.from(bech32.fromWords(decodeAddress.words))
}

function depositToken(symbol="BNB"){
  httpClient
    .get(sequenceURL)
    .then((res) => {
        const sequence = res.data.sequence || 0
        return bnbClient.transfer(addressFrom, process.env.BINANCE_ADDRESS, amount, symbol, message, sequence)
    })
    .then((result) => {
        if (result.status === 200) {
            console.log('success', result.result[0].hash);
        } else {
            console.error('error', result);
        }
    })
    .catch((error) => {
        console.error('error', error);
    });
}

let httpClient, sequenceURL, addressFrom


describe("BinanceGateway", () => {

  beforeEach(async function(){
  this.timeout(40000)

  if (process.env.DAPP_ENV != 'plasma'){
    bnbClient.chooseNetwork("testnet"); // or this can be "mainnet"
  } else {
    bnbClient.chooseNetwork("mainnet")
  }
  bnbClient.setPrivateKey(process.env.BINANCE_WALLET_KEY);
  
  await bnbClient.initChain();
  
  addressFrom = bnbClient.getClientKeyAddress(); // sender address string (e.g. bnb1...)

  httpClient = axios.create({ baseURL: api });
  sequenceURL = `${api}api/v1/account/${addressFrom}/sequence`;

  const aliceDappchainPrivateKey = CryptoUtils.B64ToUint8Array(process.env.ALICE_DAPP_KEY)
  const aliceDappchainPublicKey = CryptoUtils.publicKeyFromPrivateKey(aliceDappchainPrivateKey)
  const aliceDappchainAddress = LocalAddress.fromPublicKey(aliceDappchainPublicKey)
  aliceDappchainHex = aliceDappchainAddress.toString()
  console.log("aliceDappchainHex", aliceDappchainHex)

  if (process.env.DAPP_ENV == undefined || process.env.DAPP_ENV == 'local') {
    dappclient = new Client(
      'default',
      'ws://127.0.0.1:46658/websocket',
      'ws://127.0.0.1:46658/queryws'
      )
  }

  dappclient.txMiddleware = [new SpeculativeNonceTxMiddleware(aliceDappchainPublicKey, dappclient), new SignedTxMiddleware(aliceDappchainPrivateKey)]
    
  loomWeb3 = new Web3(new LoomProvider(dappclient, aliceDappchainPrivateKey))
  const aliceLocalAddress = new Address(dappclient.chainId, aliceDappchainAddress)

  try {
    dappchainGateway = await Contracts.BinanceTransferGateway.createAsync(
      dappclient,
      aliceLocalAddress
    )
  } catch (err) {
    console.log(err)
  }

  try {
    bnbDappTokenInstance = new loomWeb3.eth.Contract(erc20abi.abi, process.env.BNB_COIN_ADDRESS_HEX)
  } catch (err) {
    console.log(err)
  }

  });

  after( () => {
    dappclient.disconnect()
    }
  );

  describe("Deposit and Withdraw flow", function(){
    it('should deposit BNB to gateway contract and withdraw back in the same BNB amount - fee', async function() {
      this.timeout(120000)
      let aliceDappBalanceBefore = await bnbDappTokenInstance.methods.balanceOf(aliceDappchainHex).call({from:aliceDappchainHex})
      let aliceDappBalanceAfter = await bnbDappTokenInstance.methods.balanceOf(aliceDappchainHex).call({from:aliceDappchainHex})
      console.log("aliceDappBalanceBefore", aliceDappBalanceBefore)
      if (process.env.DEPOSIT == "true") {
        console.log("sending bnbtoken to hot wallet")
        await depositToken()
        while (aliceDappBalanceAfter == aliceDappBalanceBefore) {
          aliceDappBalanceAfter = await bnbDappTokenInstance.methods.balanceOf(aliceDappchainHex).call({from:aliceDappchainHex})
          await delay(5000)
        }
        console.log("alicebalanceAfter", aliceDappBalanceAfter)
      }

      if (process.env.DEPOSIT_MOOL == "true") {
        depositToken("MOOL")
      }

      if (process.env.DEPOSIT == "true") {
        assert.equal(aliceDappBalanceAfter - aliceDappBalanceBefore, AMOUNT_INT, `Alice balance should be increased by ${AMOUNT_INT}`)
      } else {
        assert.equal(aliceDappBalanceAfter , aliceDappBalanceBefore, `balance should be the same`)
      }

      // call dappchain gateway to withdraw token
      // first we approve dappchain gateway to take money
      console.log("approving bnbtoken to gateway")
      let dappchainGatewayHex = dappchainGateway.address.toString()
      console.log("dappchainGatewayHex", dappchainGatewayHex)
      if (process.env.DAPP_ENV == 'extdev'){
        dappchainGatewayHex = dappchainGatewayHex.slice(18)
      } else if (process.env.DAPP_ENV == 'asia1') {
        dappchainGatewayHex = dappchainGatewayHex.slice(6)
      } else {
        dappchainGatewayHex = dappchainGatewayHex.slice(8)
      }

      await bnbDappTokenInstance.methods.approve(dappchainGatewayHex, AMOUNT_INT+FEE).send({from:aliceDappchainHex})
      
      let approveBalance = 0

      while(approveBalance == 0){
        approveBalance = await bnbDappTokenInstance.methods.allowance(aliceDappchainHex, dappchainGatewayHex).call({from:aliceDappchainHex})
        await delay(5000)
      }
      
      const bnbTokenAddress = Address.fromString("default:"+bnbDappTokenInstance._address.toLowerCase())
      const tmp = decodeAddress(addressFrom)
      recepient = new Address('binance', new LocalAddress(tmp))
      let aliceDappBalanceBeforeWithdraw = await bnbDappTokenInstance.methods.balanceOf(aliceDappchainHex).call({from:aliceDappchainHex})
      console.log("aliceDappBalanceBeforeWithdraw", aliceDappBalanceBeforeWithdraw)
      try {
        let receipt = await dappchainGateway.withdrawTokenAsync(new BN(AMOUNT_INT  - FEE, 10), bnbTokenAddress, recepient)
      } catch (err){
        console.log("withdrawTokenAsync", err)
      }
      let aliceDappBalanceAfterWithdraw = await bnbDappTokenInstance.methods.balanceOf(aliceDappchainHex).call({from:aliceDappchainHex})
      console.log("aliceDappBalanceAfterWithdraw", aliceDappBalanceAfterWithdraw)
      assert.equal(aliceDappBalanceBeforeWithdraw - aliceDappBalanceAfterWithdraw, AMOUNT_INT, `Alice balance should be decreased by ${AMOUNT_INT}`)

    });
  });
});
