const TronWeb = require('tronweb')
const assert = require('assert')
const yaml = require('js-yaml')
const fs = require('fs')
const {CryptoUtils, Client, LocalAddress, LoomProvider, Contracts, TronWebSigner, Address, createDefaultTxMiddleware} = require('loom-js')
const Web3 = require('web3')
const BN = require('bn.js')
const request = require('request');

const erc20abi = require('../dappchain/build/contracts/TRXToken.json')

const AddressMapper = Contracts.AddressMapper

// READ from yaml files
const contracts = yaml.safeLoad(fs.readFileSync('../e2e_config/local_shasta/contracts.yml', 'utf8'))

const fullNode = new TronWeb.providers.HttpProvider('https://api.shasta.trongrid.io');
const solidityNode = new TronWeb.providers.HttpProvider('https://api.shasta.trongrid.io');
const eventServer = 'https://api.shasta.trongrid.io';
let newAccount, aliceDappchainHex, gateway, alice, aliceTRXAddress, aliceTrxAddrObj
let newAccountPromise = TronWeb.createAccount()
const delay = ms => new Promise(resolve => setTimeout(resolve, ms));
let dappchainGateway, trxDappTokenInstance, tronWeb, loomWeb3
const TRX_AMOUNT = 1000000
const TRON_DAPP_ADDRESS = '0x0000000000000000000000000000000000000001'
let dappclient, TRX_COIN_ADDRESS_HEX

describe("TronGateway", () => {

  beforeEach(async function(){
  this.timeout(40000)

  newAccountPromise.then(function() {
      console.log("new account created"); // "Stuff worked!"
    }, function(err) {
      console.log(err); // Error: "It broke"
    });

  newAccount = await newAccountPromise
  tronWeb = new TronWeb(fullNode, solidityNode, eventServer, newAccount.privateKey);
  alice = tronWeb.defaultAddress.base58
  aliceTRXAddress = tronWeb.address.toHex(alice)
  aliceTRXAddress = '0x'+ aliceTRXAddress.substring(2,100)
  aliceTrxAddrObj =  Address.fromString(`tron:${aliceTRXAddress}`)


  const aliceDappchainPrivateKey = CryptoUtils.generatePrivateKey()
  const aliceDappchainPublicKey = CryptoUtils.publicKeyFromPrivateKey(aliceDappchainPrivateKey)
  const aliceDappchainAddress = LocalAddress.fromPublicKey(aliceDappchainPublicKey)
  aliceDappchainHex = aliceDappchainAddress.toString()

  if (process.env.DAPP_ENV == undefined || process.env.DAPP_ENV == 'local') {
    dappclient = new Client(
      'default',
      'ws://127.0.0.1:46658/websocket',
      'ws://127.0.0.1:46658/queryws'
      )
    TRX_COIN_ADDRESS_HEX = '0xb4a4f7bf5c71074e6c56ff27609aeba09011dda8'
  } else if (process.env.DAPP_ENV == 'asia1') {
    dappclient = new Client(
      'asia1',
      'wss://test-z-asia1.dappchains.com/websocket',
      'wss://test-z-asia1.dappchains.com/queryws'
    )
    TRX_COIN_ADDRESS_HEX = '0xb4a4f7bf5c71074e6c56ff27609aeba09011dda8'
  }

  dappclient.txMiddleware = createDefaultTxMiddleware(dappclient, aliceDappchainPrivateKey)
    
  loomWeb3 = new Web3(new LoomProvider(dappclient, aliceDappchainPrivateKey))
  const aliceLocalAddress = new Address(dappclient.chainId, aliceDappchainAddress)
  const signer = new TronWebSigner(tronWeb, aliceTRXAddress)

  async function getContract(address) {
    const res = await tronWeb.contract().at(address);
    return res
  }

  // request TRX Token from shasta faucet
  console.log("request TRX Token from shasta faucet")

  var headers = {
      'accept-encoding': 'gzip, deflate, br',
      'accept-language': 'en-US,en;q=0.9,th;q=0.8',
      'content-type': 'application/x-www-form-urlencoded; charset=UTF-8',
      'accept': '*/*',
      'referer': 'https://www.trongrid.io/shasta/',
      'authority': 'www.trongrid.io',
      'x-requested-with': 'XMLHttpRequest'
  };

  var dataString = 'value='+newAccount.address.base58;

  var options = {
      url: 'https://www.trongrid.io/shasta/submit',
      method: 'POST',
      headers: headers,
      body: dataString
  };

  function callback(error, response, body) {
      if (!error && response.statusCode == 200) {
          console.log(body);
      }
  }

  await request(options, callback);
  gateway = await getContract(contracts.mainnet_gateway_addr);

  const addressMapper = await AddressMapper.createAsync(
    dappclient,
    aliceLocalAddress
  )

  console.log(`Mapping user DAppChain address ${aliceLocalAddress} to ${aliceTrxAddrObj.toString()}`)  
  console.log("Check account mapping")
  if (!(await addressMapper.hasMappingAsync(aliceLocalAddress))) {
    console.log("Account has not been mapped with any account. Mapping...")
    await addressMapper.addIdentityMappingAsync(
      aliceLocalAddress,
      aliceTrxAddrObj,
      signer
    )
  }

  try {
    dappchainGateway = await Contracts.TronTransferGateway.createAsync(
      dappclient,
      aliceLocalAddress
    )
  } catch (err) {
    console.log(err)
  }

  try {
    trxDappTokenInstance = new loomWeb3.eth.Contract(erc20abi.abi, TRX_COIN_ADDRESS_HEX)
  } catch (err) {
    console.log(err)
  }

  });

  after( () => {
    dappclient.disconnect()
    }
  );

  describe("Deposit and Withdraw flow", function(){
    it('should deposit TRX to gateway contract and withdraw back in the same TRX amount', async function() {
      this.timeout(60000)
      const balanceBefore = await gateway.getTRX().call();
      console.log("Gateway balance before", balanceBefore.toString())
      let aliceDappBalanceBefore = await trxDappTokenInstance.methods.balanceOf(aliceDappchainHex).call({from:aliceDappchainHex})
      console.log("aliceDappBalanceBefore ", aliceDappBalanceBefore)
      const txHash = await gateway.sendToken().send({'from': alice, 'callValue': TRX_AMOUNT})
      await delay(10000)
      const balanceAfter = await gateway.getTRX().call();
      const depositEvent =  await tronWeb.getEventByTransactionID(txHash);
      assert.equal(balanceAfter - balanceBefore, TRX_AMOUNT, `gateway balance should be increased by ${TRX_AMOUNT}`)
      assert.equal(depositEvent[0].result.from, aliceTRXAddress, "deposit address should be Alice address")
      assert.equal(depositEvent[0].result.amount, TRX_AMOUNT.toString(), `deposit amount should be ${TRX_AMOUNT}`)
      console.log("Gateway balance before", balanceAfter.toString())

      let aliceDappBalanceAfter = 0
      
      while (aliceDappBalanceAfter == 0) {
        aliceDappBalanceAfter = await trxDappTokenInstance.methods.balanceOf(aliceDappchainHex).call({from:aliceDappchainHex})
        console.log("aliceDappBalanceAfter is ", aliceDappBalanceAfter)
        await delay(5000)
      }

      // call dappchain gateway to withdraw token
      // first we approve dappchain gateway to take money
      console.log("approving trxtoken to gateway")
      let dappchainGatewayHex = dappchainGateway.address.toString()
      dappchainGatewayHex = dappchainGatewayHex.slice(8)
      await trxDappTokenInstance.methods.approve(dappchainGatewayHex, TRX_AMOUNT).send({from:aliceDappchainHex}) // approve stuck if await. No idea why
      
      let approveBalance = 0

      while(approveBalance == 0){
        approveBalance = await trxDappTokenInstance.methods.allowance(aliceDappchainHex, dappchainGatewayHex).call({from:aliceDappchainHex})
        delay(5000)
      }
      
      const timeout = 60 * 1000
      const gatewayContract = dappchainGateway
      const receiveSignedWithdrawalEvent = new Promise((resolve, reject) => {
        let timer = setTimeout(
          () => reject(new Error('Timeout while waiting for withdrawal to be signed')),
          timeout
        )
        const listener = event => {
          console.log("receiveSignedWithdrawalEvent resolved")

          if (
            event.tokenContract.local.toString() === TRON_DAPP_ADDRESS &&
            event.tokenOwner.toString() === aliceTrxAddrObj.toString()
          ) {
            clearTimeout(timer)
            timer = null
            gatewayContract.removeAllListeners(Contracts.TransferGateway.EVENT_TOKEN_WITHDRAWAL)
            console.log('Oracle signed tx ', CryptoUtils.bytesToHexAddr(event.sig))
            resolve(event)
          }
        }
        gatewayContract.on(Contracts.TransferGateway.EVENT_TOKEN_WITHDRAWAL, listener)
      })

      let receipt
      let sig
      try 
      {
        let trxTokenAddress = Address.fromString("default:"+trxDappTokenInstance._address.toLowerCase())
        receipt = await dappchainGateway.withdrawTRXAsync(new BN(TRX_AMOUNT, 10), trxTokenAddress, aliceTrxAddrObj)
      } catch (err) {
        console.log("error withdraw TRX token, reason: ", err)
        try {
          receipt = await dappchainGateway.withdrawalReceiptAsync(Address.fromString("default:"+aliceDappchainHex))
          sig = receipt.oracleSignature
        } catch (err2) {
          console.log(err2)
        }
      }

      receipt = await receiveSignedWithdrawalEvent
      sig = receipt.sig
      
      sig = CryptoUtils.bytesToHexAddr(sig)

      if ( sig.length > 132) {
        byteToOmit = sig.length - 132 + 2 // +2 from `0x`
        sig = sig.slice(byteToOmit)
        sig = '0x' + sig
      }

      const r = sig.slice(0, 66)
      const s = '0x' + sig.slice(66, 130)
      let v = '0x' + sig.slice(130, 132)
      v = loomWeb3.utils.toDecimal(v)
      

      let withdrawTxHash = await gateway.withdrawTRX(TRX_AMOUNT, r, s, v).send({ from: alice }) // if successful alice gets her trx back
      await delay(7000)
      const withdrawEvent =  await tronWeb.getEventByTransactionID(withdrawTxHash);
      assert.equal(withdrawEvent[0].result.from, aliceTRXAddress, "Withdrawal address should be the same as alice TRX address")
      assert.equal(withdrawEvent[0].result.value, TRX_AMOUNT.toString(), "Withdrawal amount should be the same as TRX_AMOUNT")

    });
  });
});
