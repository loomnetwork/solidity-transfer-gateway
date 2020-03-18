const ERC721XCards = artifacts.require('ERC721XCards')
const Gateway = artifacts.require('Gateway')
const GameToken = artifacts.require('GameToken')
const Loom = artifacts.require('LoomToken')
const BadERC20Token = artifacts.require('BadERC20Token')

const Reentrancy = artifacts.require('./ReentrancyExploit.sol')

const { soliditySha3 } = require('web3-utils')
const { signHash, assertEventVar, Promisify } = require('./utils.js')

import assertRevert from './helpers/assertRevert.js'
import expectThrow from './helpers/expectThrow'


contract('Transfer Gateway V2', async (accounts) => {
  let gateway, erc20, erc721x, loom, badToken
  let [validator, alice, bob] = accounts

  const ERC20_AMOUNT = 10 * 10 ** 18
  const ERC721_UID = 2
  const ETHER_AMOUNT = 3 * 10 ** 18;

  beforeEach(async () => {
    loom = await Loom.new({from : validator })
    gateway = await Gateway.new(loom.address, [validator], 3, 4, [], [], { from: validator })
    erc721x = await ERC721XCards.new(gateway.address, { from: validator })
    erc20 = await GameToken.new({ from: validator })
    badToken = await BadERC20Token.new({ from: validator })
    gateway.toggleToken(erc721x.address, { from: validator })

    // Give Alice some coins
    let tokenIds = [...Array(5).keys()]
    let amounts = [ 100, 100, 1, 300, 1] // token 2 and 4 are NFTs
    let receivers = [ accounts[0], accounts[1], accounts[0], accounts[2], accounts[1]]
    await erc721x.airdrop(tokenIds, amounts, receivers, { from: validator })
    await erc20.transfer(alice, ERC20_AMOUNT, { from: validator })
    await loom.transfer(alice, ERC20_AMOUNT, { from: validator })
    await badToken.transfer(alice, ERC20_AMOUNT, { from: validator })

  })

  describe('Initialization', async() => {
    it('Number of accounts must match numbers of nonces', async () => {
      const loom = await Loom.new({from : validator })
      const accounts = ["0xf1459f9efdc06c7c278a1e9d6187050efe517ab5"]
      const nonces = []
      assertRevert(Gateway.new(loom.address, [validator], 3, 4, accounts, nonces, { from: validator }))
    })
    it('Account nonces should be prepopulated', async () => {
      let accounts = ["0xf1459f9efdc06c7c278a1e9d6187050efe517ab5", "0x9e404e86127931ac193da80c26e49922090d97a0"]
      let nonces = [2, 4]
      const gateway = await Gateway.new(loom.address, [validator], 3, 4, accounts, nonces, { from: validator })
      assert.equal(await gateway.nonces.call(accounts[0]), nonces[0])
      assert.equal(await gateway.nonces.call(accounts[1]), nonces[1])
    })
  })

  describe('Ether deposit / withdrawals', async () => {

    let amount = ETHER_AMOUNT

    it('Alice sends some ether to the Gateway. Validator signs a message that allows her to withdraw the ether.', async () => {
      await depositEther(alice, amount)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(alice, nonce, gateway.address, soliditySha3(amount))
      let sig = await signHash(validator, hash)
      await gateway.withdrawETH(amount, sig, { from: alice }) // if successful alice gets her eth back
    })

    it('Is not reentrant', async () => {
      await depositEther(alice, amount)

      // Bob's contract `reentrancy` is authorized 0.1 * amount, but he'll try to steal more.
      const reentrancy = await Reentrancy.new({from : bob})
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(reentrancy.address, nonce, gateway.address, soliditySha3(amount * 0.1))
      let sig = await signHash(validator, hash)
      await reentrancy.setup(gateway.address, amount * 0.1, sig)
      assertRevert(reentrancy.attack())
    })

    it('Validator signs an invalid amount. Tx fails.', async () => {
      await depositEther(alice, amount)
      let invalid_amount = 10 * amount
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(alice, nonce, gateway.address, soliditySha3(10 * invalid_amount))
      let sig = await signHash(validator, hash)
      assertRevert(gateway.withdrawETH(invalid_amount, sig, { from: alice }))
    })

    it('Validator signs an invalid nonce. Tx fails', async () => {
      await depositEther(alice, amount)
      let invalid_nonce = 42
      let hash = soliditySha3(alice, invalid_nonce, gateway.address, soliditySha3(amount))
      let sig = await signHash(validator, hash)
      assertRevert(gateway.withdrawETH(amount, sig, { from: alice }))
    })


    async function depositEther(from, amount) {
      let tx = await gateway.sendTransaction({ 'from': from, 'value': amount })

      // Check the user's balance
      assertEventVar(tx, 'ETHReceived', 'from', from);
      assertEventVar(tx, 'ETHReceived', 'amount', amount);
      let balance = await gateway.getETH.call()
      assert.equal(balance, amount);
      return tx;
    }
  })

  describe('ERC20 deposit / withdrawals', async () => {

    let amount = ERC20_AMOUNT

    it('Alice sends some ERC20 to the Gateway. Validator signs a message that allows her to withdraw the erc20.', async () => {
      await depositERC20(alice, amount)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(alice, nonce, gateway.address, soliditySha3(amount, erc20.address))
      let sig = await signHash(validator, hash)
      await gateway.withdrawERC20(amount, sig, erc20.address, { from: alice }) // if successful alice gets her erc20 back
      assert.equal(await erc20.balanceOf.call(alice), amount)
    })

    it('Validator signs an invalid amount. Tx fails.', async () => {
      await depositERC20(alice, amount)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(alice, nonce, gateway.address, soliditySha3(10 * amount, erc20.address))
      let sig = await signHash(validator, hash)
      assertRevert(gateway.withdrawERC20(amount, sig, erc20.address, { from: alice }))
      assert.equal(await erc20.balanceOf.call(alice), 0)
    })

    it('Validator signs an invalid nonce. Tx fails', async () => {
      await depositERC20(alice, amount)
      let nonce = 42
      let hash = soliditySha3(alice, nonce, gateway.address, soliditySha3(amount, erc20.address))
      let sig = await signHash(validator, hash)
      assertRevert(gateway.withdrawERC20(amount, sig, erc20.address, { from: alice }))
      assert.equal(await erc20.balanceOf.call(alice), 0)
    })

    it('Loom Withdraw event', async () => {
      await depositLoom(alice, amount)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(alice, nonce, gateway.address, soliditySha3(amount, loom.address))
      let sig = await signHash(validator, hash)
      await gateway.withdrawERC20(amount, sig, loom.address, { from: alice }) // if successful alice gets her erc20 back
      let withdrawn = await gateway.getPastEvents('TokenWithdrawn', {fromBlock: 0, toBlock: "latest"});
      assert.equal(withdrawn[0].args.kind, 4)
      assert.equal(await erc20.balanceOf.call(alice), amount)
    })

    it('Withdraw bad erc20 token', async () => {
      await depositBadERC20(alice, amount)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(alice, nonce, gateway.address, soliditySha3(amount, badToken.address))
      let sig = await signHash(validator, hash)
      await gateway.withdrawERC20(amount, sig, badToken.address, { from: alice }) // if successful alice gets her badToken back
      assert.equal(await badToken.balanceOf.call(alice), amount)
    })

    async function depositLoom(from, amount) {
      await loom.approve(gateway.address, amount, { 'from': from })
      await gateway.depositERC20(amount, loom.address, { 'from': from})
      let received = await gateway.getPastEvents('LoomCoinReceived', {fromBlock: 0, toBlock: "latest"});
      received = received[0].args;
      assert.equal(received.from, from);
      assert.equal(received.amount, amount);
      assert.equal(received.loomCoinAddress, loom.address);
      let balance = await gateway.getERC20.call(loom.address)
      assert.equal(balance, amount)
    }

    async function depositERC20(from, amount) {
      await erc20.approve(gateway.address, amount, { 'from': from })
      await gateway.depositERC20(amount, erc20.address, { 'from': from})
      let received = await gateway.getPastEvents('ERC20Received', {fromBlock: 0, toBlock: "latest"});
      received = received[0].args;
      assert.equal(received.from, from);
      assert.equal(received.amount, amount);
      assert.equal(received.contractAddress, erc20.address);
      let balance = await gateway.getERC20.call(erc20.address)
      assert.equal(balance, amount)
    }

    async function depositBadERC20(from, amount) {
      await badToken.approve(gateway.address, amount, { 'from': from })
      await gateway.depositERC20(amount, badToken.address, { 'from': from})
      let received = await gateway.getPastEvents('ERC20Received', {fromBlock: 0, toBlock: "latest"});
      received = received[0].args;
      assert.equal(received.from, from);
      assert.equal(received.amount, amount);
      assert.equal(received.contractAddress, badToken.address);
      let balance = await gateway.getERC20.call(badToken.address)
      assert.equal(balance, amount)
    }
  })

  describe('ERC721x deposit / withdrawals', async () => {

    it('Alice sends an ERC721 to the Gateway. Validator signs a message that allows her to withdraw the NFT.', async () => {
      let current_account = accounts[0];
      let uid = 2;
      await depositERC721(current_account, uid)
      let nonce = await gateway.nonces.call(current_account)
      let hash = soliditySha3(current_account, nonce, gateway.address, soliditySha3(uid, erc721x.address))
      let sig = await signHash(validator, hash)
      await gateway.withdrawERC721(uid, sig, erc721x.address, { from: current_account }) // if successful current_account gets her erc721x back
      assert.equal(await erc721x.ownerOf.call(uid), current_account)
    })

    it('Alice sends an ERC721x to the Gateway. Validator signs a message that allows her to withdraw the MFT.', async () => {
      let current_account = accounts[2];
      let uid = 3;
      let amount = 100
      await depositERC721X(current_account, uid, amount)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(current_account, nonce, gateway.address, soliditySha3(uid, amount, erc721x.address))
      let sig = await signHash(validator, hash)
      await gateway.withdrawERC721X(uid, amount, sig, erc721x.address, { from: current_account }) 
      assert.equal((await erc721x.balanceOf(current_account, uid)).toNumber(), 300)
    })

    it('Validator signs an invalid uid. Tx fails.', async () => {
      let current_account = accounts[0];
      let uid = 2;
      await depositERC721(current_account, uid)
      let nonce = await gateway.nonces.call(current_account)
      let hash = soliditySha3(current_account, gateway.address, nonce, soliditySha3(3, erc721x.address))
      let sig = await signHash(validator, hash)
      assertRevert(gateway.withdrawERC721(uid, sig, erc721x.address, { from: current_account }))
      assert.equal(await erc721x.ownerOf.call(uid), gateway.address)
    })

    it('Validator signs an invalid nonce. Tx fails', async () => {
      let current_account = accounts[0];
      let uid = 2;
      await depositERC721(current_account, uid)
      let nonce = 42
      let hash = soliditySha3(current_account, gateway.address, nonce, soliditySha3(uid, erc721x.address))
      let sig = await signHash(validator, hash)
      assertRevert(gateway.withdrawERC721(uid, sig, erc721x.address, { from: current_account }))
      assert.equal(await erc721x.ownerOf.call(uid), gateway.address)
    })

    it('Validator signs an invalid balance', async () => {
      let current_account = accounts[2];
      let uid = 3;
      let amount = 100
      await depositERC721X(current_account, uid, amount)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(current_account, nonce, gateway.address, soliditySha3(uid, amount + 10 , erc721x.address))
      let sig = await signHash(validator, hash)
      assertRevert(gateway.withdrawERC721X(uid, amount, sig, erc721x.address, { from: current_account }))
      assert.equal((await erc721x.balanceOf(current_account, uid)).toNumber(), 200)
    })

    async function depositERC721(from, uid) {
      await erc721x.depositToGatewayNFT(uid, { 'from': from })
      let received = await gateway.getPastEvents('ERC721Received', {fromBlock: 0, toBlock: "latest"});
      received = received[0].args;
      assert.equal(received.from, from);
      assert.equal(received.tokenId, uid);
      assert.equal(received.contractAddress, erc721x.address);
      // Check the user's balance
      let balance = await gateway.getERC721.call(uid, erc721x.address)
      assert.ok(balance)
    }

    async function depositERC721X(from, uid, amount) {
      await erc721x.depositToGateway(uid, amount, { 'from': from })
      let received = await gateway.getPastEvents('ERC721XReceived', {fromBlock: 0, toBlock: "latest"});
      received = received[0].args;
      assert.equal(received.from, from);
      assert.equal(received.tokenId, uid);
      assert.equal(received.amount, amount);
      assert.equal(received.contractAddress, erc721x.address);

      // Check the user's balance
      let balance = await gateway.getERC721X.call(uid, erc721x.address)
      assert.equal(balance, amount)
    }
  })
})
