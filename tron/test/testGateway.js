const Gateway = artifacts.require('Gateway')
const LoomToken = artifacts.require('LoomToken')
const GameToken = artifacts.require('GameToken')
const Reentrancy = artifacts.require('./ReentrancyExploit.sol')
const { soliditySha3 } = require('web3-utils')
const { signHash, assertEventVar, assertRevert } = require('./utils.js')

contract.only('Transfer Gateway V2', async (accounts) => {
  let gateway, trc20, loom
  let [validator, alice, bob] = accounts
  // This is the same as e2e_config/local_shasta/oracle_tron_priv.key
  let signerPrivkey = 'd16911f08c31d2875039b72d83e2c0acf0dd9854f67bc108472cf3ee16422638'
  const TRC20_AMOUNT = 100
  const TRX_AMOUNT = 300;
  const delay = ms => new Promise(resolve => setTimeout(resolve, ms));
  let aliceHex, gwAddrHex, trc20AddrHex;

  beforeEach(async () => {
    // loom = await LoomToken.new({ from: validator })
    loom = await LoomToken.deployed({from : validator })
    // gateway = await Gateway.new(loom.address, validator, 3, 4, { from: validator })
    gateway = await Gateway.deployed({ from: validator })
    // trc20 = await GameToken.new({ from: validator })
    trc20 = await GameToken.deployed({ from: validator })

    await loom.transfer(alice, TRC20_AMOUNT, { from: validator })
    await trc20.transfer(alice, TRC20_AMOUNT, { from: validator })

    let aliceAddr = tronWeb.address.toHex(alice)
    let gwAddr = tronWeb.address.toHex(gateway.address)
    let trc20Addr = await tronWeb.address.toHex(trc20.address)
    
    aliceHex = "0x"+aliceAddr.substring(2,100)
    gwAddrHex = "0x"+gwAddr.substring(2,100)
    trc20AddrHex = '0x'+ trc20Addr.substring(2,100)

  })

  describe('Tron deposit / withdrawals', async () => {

    async function depositTron(from, amount) {
      const balanceBefore = await gateway.getTRX()
      let tx = await gateway.sendToken({ 'from': from, 'callValue': amount })
      await delay(5000)
      let response = await tronWeb.getEventByTransactionID(tx);
      // Check the user's balance
      assertEventVar(response, 'TRXReceived', 'from', from);
      assertEventVar(response, 'TRXReceived', 'amount', amount);
      let balanceAfter = await gateway.getTRX()
      assert.equal(balanceAfter.toNumber() - balanceBefore.toNumber(), amount);
      return tx;
    }

    let amount = TRX_AMOUNT

    it('Alice sends some Tron to the Gateway. Validator signs a message that allows her to withdraw the Tron.', async () => {
      await depositTron(alice, amount)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(aliceHex, nonce, gwAddrHex, soliditySha3(amount))
      let sig = await signHash(hash, signerPrivkey) // we don't have prefix in oracle, so I made a sign function.
      let r = '0x' + sig['r'].toString('hex')
      let s = '0x' + sig['s'].toString('hex')
      let v = sig['v']
      let txId = await gateway.withdrawTRX(amount, r, s, v, { from: alice }) // if successful alice gets her trx back
      await delay(5000)
      let response = await tronWeb.getEventByTransactionID(txId);
      assertEventVar(response, 'TokenWithdrawn', 'from', alice);
      assertEventVar(response, 'TokenWithdrawn', 'value', amount);
    })

    it('Is not reentrant', async () => {
      await depositTron(alice, amount)
      // Bob's contract `reentrancy` is authorized 0.1 * amount, but he'll try to steal more.
      const reentrancy = await Reentrancy.deployed({ from: bob })
      let reentrancyHex = tronWeb.address.toHex(reentrancy.address)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(aliceHex, nonce, gwAddrHex, soliditySha3(amount))
      let sig = await signHash(hash, signerPrivkey) // we don't have prefix in oracle, so I made a sign function.
      let r = '0x' + sig['r'].toString('hex')
      let s = '0x' + sig['s'].toString('hex')
      let v = sig['v']   
      let txId = await reentrancy.setup(gwAddrHex, amount * 0.1, r,s,v)
      await delay(5000)
      let response = await tronWeb.getEventByTransactionID(txId);
      assertRevert(response)
    })

    it('Validator signs an invalid amount. Tx fails.', async () => {
      await depositTron(alice, amount)
      let invalid_amount = 10 * amount
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(aliceHex, nonce, gwAddrHex, soliditySha3(10 * invalid_amount))
      let sig = await signHash(hash, signerPrivkey) // we don't have prefix in oracle, so I made a sign function.
      let r = '0x' + sig['r'].toString('hex')
      let s = '0x' + sig['s'].toString('hex')
      let v = sig['v']
      let txId = await gateway.withdrawTRX(invalid_amount, r,s,v, { from: alice })
      await delay(5000)
      let response = await tronWeb.getEventByTransactionID(txId);
      assertRevert(response)
    })

    it('Validator signs an invalid nonce. Tx fails', async () => {
      let balanceBefore = await trc20.balanceOf.call(alice)
      await depositTron(alice, amount)
      let invalid_nonce = 42
      let hash = soliditySha3(aliceHex, invalid_nonce, gwAddrHex, soliditySha3(amount))
      let sig = await signHash(hash, signerPrivkey) // we don't have prefix in oracle, so I made a sign function.
      let r = '0x' + sig['r'].toString('hex')
      let s = '0x' + sig['s'].toString('hex')
      let v = sig['v']
      let txId = await gateway.withdrawTRX(amount, r,s,v, { from: alice })
      await delay(5000)
      let response = await tronWeb.getEventByTransactionID(txId);
      assertRevert(response)
      let balanceAfter = await trc20.balanceOf.call(alice)
      assert.equal(balanceBefore.toNumber() - balanceAfter.toNumber(), 0)
    })



  })

  describe('TRC20 deposit / withdrawals', async () => {

    let amount = TRC20_AMOUNT

    async function depositTRC20(from, amount) {
      let balanceBefore = await gateway.getTRC20(trc20.address)
      await trc20.approve(gateway.address, amount, { 'from': from })
      let tx = await gateway.depositTRC20(amount, trc20.address, { 'from': from})

      await delay(5000)
      let response = await tronWeb.getEventByTransactionID(tx);
      assertEventVar(response, 'TRC20Received', 'from', from);
      assertEventVar(response, 'TRC20Received', 'amount', amount);
      assertEventVar(response, 'TRC20Received', 'contractAddress', trc20.address);
      let balanceAfter = await gateway.getTRC20(trc20.address)
      assert.equal(balanceAfter.toNumber() - balanceBefore.toNumber(), amount);
    }

    it('Alice sends some TRC20 to the Gateway. Validator signs a message that allows her to withdraw the TRC20.', async () => {
      await depositTRC20(alice, amount)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(aliceHex, nonce, gwAddrHex, soliditySha3(amount, trc20AddrHex))
      let sig = await signHash(hash, signerPrivkey) // we don't have prefix in oracle, so I made a sign function.
      let r = '0x' + sig['r'].toString('hex')
      let s = '0x' + sig['s'].toString('hex')
      let v = sig['v']
      let txId = await gateway.withdrawTRC20(amount, r, s, v, trc20AddrHex, { from: alice }) // if successful alice gets her trc20 back
      await delay(5000)
      let response = await tronWeb.getEventByTransactionID(txId);
      assertEventVar(response, 'TokenWithdrawn', 'from', aliceHex);
      assertEventVar(response, 'TokenWithdrawn', 'value', amount);

    })

    it('Validator signs an invalid amount. Tx fails.', async () => {
  
      await depositTRC20(alice, amount)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(aliceHex, nonce, gwAddrHex, soliditySha3(10 * amount, trc20AddrHex))
      let sig = await signHash(hash, signerPrivkey) // we don't have prefix in oracle, so I made a sign function.
      let r = '0x' + sig['r'].toString('hex')
      let s = '0x' + sig['s'].toString('hex')
      let v = sig['v']
      let txId = await gateway.withdrawTRC20(amount, r, s, v, trc20AddrHex, { from: alice })
      delay(5000)
      let response = await tronWeb.getEventByTransactionID(txId);
      assertRevert(response)
    })

    it('Validator signs an invalid nonce. Tx fails', async () => {
      let balanceBefore = await trc20.balanceOf.call(alice)

      await depositTRC20(alice, amount)
      let nonce = 42
      let hash = soliditySha3(aliceHex, nonce, gwAddrHex, soliditySha3(amount, trc20AddrHex))
      let sig = await signHash(hash, signerPrivkey) // we don't have prefix in oracle, so I made a sign function.
      let r = '0x' + sig['r'].toString('hex')
      let s = '0x' + sig['s'].toString('hex')
      let v = sig['v']
      let txId = await gateway.withdrawTRC20(amount, r, s, v, trc20AddrHex, { from: alice })
      await delay(5000)
      let response = await tronWeb.getEventByTransactionID(txId);
      assertRevert(response)
      let balanceAfter = await trc20.balanceOf.call(alice)
      assert.equal(balanceBefore.toNumber() - balanceAfter.toNumber(), amount)
    })
  
  })

})
