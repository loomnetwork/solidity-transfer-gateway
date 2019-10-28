// vmc and gateway
const ValidatorManagerContract = artifacts.require('ValidatorManagerContract')
const Gateway = artifacts.require('Gateway')

// Tokens
const ERC721XCards = artifacts.require('ERC721XCards')
const GameToken = artifacts.require('GameToken')
const Loom = artifacts.require('LoomToken')
const SampleERC20Token = artifacts.require('SampleERC20MintableToken')
const SampleERC721Token = artifacts.require("SampleERC721MintableToken")
const NormalERC721Token = artifacts.require("CryptoCards")

// utils
const { soliditySha3 } = require('web3-utils')
const { assertEventVar, createSigns, createValidators } = require('./utils.js')


contract('Transfer Gateway - Single Validator', async (accounts) => {
  let gateway, erc20, erc721, erc721x, loom, vmc, erc20Mintable
  let [validator, alice] = accounts

  const TokenKind = {
    ETH: "\x0eWithdraw ETH:\n",
    ERC20: "\x10Withdraw ERC20:\n",
    ERC721: "\x11Withdraw ERC721:\n",
    ERC721X: "\x12Withdraw ERC721X:\n"
  };

  const ERC20_AMOUNT = "10000000000000000000";
  const ERC20_AMOUNT_2 = "10000000000000000000000";
  const ETHER_AMOUNT = "3000000000000000000";
  const threshold = 1.0

  beforeEach(async () => {
    oneValidator = createValidators(1);
    loom = await Loom.new({ from: validator })
    vmc = await ValidatorManagerContract.new(oneValidator.addresses, oneValidator.powers, 2, 3, loom.address);
    gateway = await Gateway.new(vmc.address)

    erc20 = await GameToken.new({ from: validator })
    erc20Mintable = await SampleERC20Token.new(gateway.address, { from: validator })

    oneValidatorTotalPower = await vmc.totalPower.call();

    erc721 = await SampleERC721Token.new(gateway.address, { from: validator })
    erc721_2 = await NormalERC721Token.new(gateway.address, { from: validator })
    erc721x = await ERC721XCards.new(gateway.address, "rinkeby.loom.games/", { from: validator })

    await erc20.transfer(alice, ERC20_AMOUNT, { from: validator })
    await erc721_2.mintTokens(alice, { from: validator })

    let tokenIds = [...Array(5).keys()]
    let amounts = [100, 100, 1, 300, 1] // token 2 and 4 are NFTs
    let receivers = [alice, alice, alice, alice, alice]
    await erc721x.airdrop(tokenIds, amounts, receivers, { from: validator })


  })

  describe('ERC20, ERC721, and ETH deposit / withdrawals with only one validator', async () => {

    it('Withdraw ether with stake >= equal to threshold', async () => {
      let amount = ETHER_AMOUNT
      await depositEther(alice, amount)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ETH, alice, nonce, gateway.address, soliditySha3(amount))
      let sigs = await createSigns(oneValidator, hash, oneValidatorTotalPower, threshold)
      await gateway.withdrawETH(amount, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice })
    })

    it('Withdraw ERC20 with stake >= equal to threshold', async () => {
      let amount = ERC20_AMOUNT
      await depositERC20(alice, amount)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ERC20, alice, nonce, gateway.address, soliditySha3(amount, erc20.address))
      let sigs = await createSigns(oneValidator, hash, oneValidatorTotalPower, threshold)
      await gateway.withdrawERC20(amount, erc20.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice });
      assert.equal(await erc20.balanceOf.call(alice), amount)
    })

    it('Withdraw ERC20 mintable token with stake >= equal to threshold, without deposit', async () => {
      let amount = ERC20_AMOUNT_2
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ERC20, alice, nonce, gateway.address, soliditySha3(amount, erc20Mintable.address))
      let sigs = await createSigns(oneValidator, hash, oneValidatorTotalPower, threshold)

      // Gateway doesn't have any tokens so it will mint the entire withdrawal amount
      await gateway.withdrawERC20(amount, erc20Mintable.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice });
      assert.equal(await erc20Mintable.balanceOf.call(alice), amount)
    })

    let uid = 2;
    let new_uid = 9999;

    it('Withdraw ERC721 with stake >= equal to threshold, without deposit for mintable tokens', async () => {
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ERC721, alice, nonce, gateway.address, soliditySha3(new_uid, erc721.address))
      let sigs = await createSigns(oneValidator, hash, oneValidatorTotalPower, threshold)
      await gateway.withdrawERC721(new_uid, erc721.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice });
      assert.equal(await erc721.ownerOf.call(new_uid), alice)
    })

    it('Withdraw ERC721 with stake >= equal to threshold', async () => {
      uid = 3
      await depositNormalERC721ToGateway(alice, uid)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ERC721, alice, nonce, gateway.address, soliditySha3(uid, erc721_2.address))
      let sigs = await createSigns(oneValidator, hash, oneValidatorTotalPower, threshold)
      await gateway.withdrawERC721(uid, erc721_2.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice });

      assert.equal(await erc721_2.ownerOf.call(uid), alice)
    })

    it('Withdraw ERC721x with stake >= equal to threshold', async () => {
      uid = 0;
      amount = 50
      await depositERC721x(alice, uid, amount)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ERC721X, alice, nonce, gateway.address, soliditySha3(uid, amount, erc721x.address))
      let sigs = await createSigns(oneValidator, hash, oneValidatorTotalPower, threshold)
      await gateway.withdrawERC721X(uid, amount, erc721x.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice });
    })

    async function depositNormalERC721ToGateway(from, uid) {
      await erc721_2.approve(gateway.address, uid, { 'from': from })
      await erc721_2.transferFrom(from, gateway.address, uid, { 'from': from })
    }

  });

  async function depositEther(from, amount) {
    let tx = await gateway.sendTransaction({ 'from': from, 'value': amount })

    // Check the user's balance
    assertEventVar(tx, 'ETHReceived', 'from', from);
    assertEventVar(tx, 'ETHReceived', 'amount', amount);
    return tx;
  }

  async function depositERC20(from, amount) {
    await erc20.approve(gateway.address, amount, { 'from': from })
    await gateway.depositERC20(amount, erc20.address, { 'from': from })
  }

  async function depositERC721(from, uid) {
    await erc721x.depositToGatewayNFT(uid, { 'from': from })
  }

  async function depositERC721x(from, uid, amount) {
    await erc721x.depositToGateway(uid, amount, { 'from': from })
  }

})
