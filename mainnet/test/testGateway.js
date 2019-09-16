// vmc and gateway
const ValidatorManagerContract = artifacts.require('ValidatorManagerContract')
const Gateway = artifacts.require('Gateway')

// Tokens
const ERC721XCards = artifacts.require('ERC721XCards')
const GameToken = artifacts.require('GameToken')
const Loom = artifacts.require('LoomToken')
const BadERC20Token = artifacts.require('BadERC20Token')
const Reentrancy = artifacts.require('ReentrancyExploit')
const SampleERC20Token = artifacts.require('SampleERC20MintableToken')
const SampleERC721Token = artifacts.require("SampleERC721MintableToken")
const NormalERC721Token = artifacts.require("CryptoCards")

const { soliditySha3 } = require('web3-utils')
const { assertEventVar, Promisify, createSigns, createValidators } = require('./utils.js')
const { shouldFail } = require('openzeppelin-test-helpers');


contract('Transfer Gateway', async (accounts) => {
  let gateway, erc20, erc721, erc721x, loom, vmc, erc20Mintable
  let [validator, alice, bob] = accounts
  let owner = validator
  let notOwner = bob
  let validators
  let valIndex, validatorsTotalPower
  let acc

  const TokenKind = {
    ETH: "\x0eWithdraw ETH:\n",
    ERC20: "\x10Withdraw ERC20:\n",
    ERC721: "\x11Withdraw ERC721:\n",
    ERC721X: "\x12Withdraw ERC721X:\n"
  };

  const ERC20_AMOUNT = "10000000000000000000";
  const ERC20_AMOUNT_2 = "10000000000000000000000";
  const ERC20_AMOUNT_3 = "5000000000000000000";
  const ERC721_UID = "2"
  const ETHER_AMOUNT = "3000000000000000000";
  const threshold = 0.67

  beforeEach(async () => {
    validators = createValidators(21);
    loom = await Loom.new({ from: validator })
    vmc = await ValidatorManagerContract.new(validators.addresses, validators.powers, 2, 3, loom.address);
    gateway = await Gateway.new(vmc.address)

    erc721 = await SampleERC721Token.new(gateway.address, {from: validator})
    normalErc721 = await NormalERC721Token.new(gateway.address, {from: validator})
    erc721x = await ERC721XCards.new(gateway.address, "rinkeby.loom.games/", { from: validator })
    erc20 = await GameToken.new({ from: validator })
    badToken = await BadERC20Token.new({ from: validator })
    erc20Mintable = await SampleERC20Token.new(gateway.address, {from: validator})

    await gateway.toggleAllowToken(erc721x.address, true, { from: validator })

    // Give Alice some coins
    let tokenIds = [...Array(5).keys()]
    let amounts = [100, 100, 1, 300, 1] // token 2 and 4 are NFTs
    let receivers = [alice, alice, alice, alice, alice]
    await erc721x.airdrop(tokenIds, amounts, receivers, { from: validator })
    await erc721x.airdrop([0], [50], [bob], { from: validator })
    await erc20.transfer(alice, ERC20_AMOUNT, { from: validator })
    await erc20.transfer(bob, ERC20_AMOUNT, { from: validator })
    await loom.transfer(alice, ERC20_AMOUNT, { from: validator })
    await badToken.transfer(alice, ERC20_AMOUNT, { from: validator })
    await normalErc721.mintTokens(alice, {from: validator})

    validatorsTotalPower = await vmc.totalPower.call();
  })

  describe('Ether deposit / withdrawals + edge cases', async () => {

    let amount = ETHER_AMOUNT

    it('Cannot frontrun rotateValidators call with a withdrawal transaction to prevent being kicked out of the validators', async () => {
      const tinyAmount = 1
      await depositEther(alice, amount)

      const newValidators = createValidators(11);
      const rotateNonce = await vmc.nonce.call() // this should not be modified 
      const newVSetHash = soliditySha3(
        vmc.address,
        rotateNonce,
        soliditySha3(
          { t: 'address[]', v: newValidators.addresses },
          { t: 'uint256[]', v: newValidators.powers }
        )
      )
      let sigs = await createSigns(validators, newVSetHash, validatorsTotalPower, 1)

      // try to withdraw in an attempt to mutate the vmc state
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ETH, alice, nonce, gateway.address, soliditySha3(tinyAmount))
      let withdraw_sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
      // this should not affect the nonce that was signed for the rotateValidators call
      await gateway.withdrawETH(tinyAmount, withdraw_sigs.signers, withdraw_sigs.v, withdraw_sigs.r, withdraw_sigs.s, { from: alice })

      // This should not fail if the nonce is not incremented by external calls
      await vmc.rotateValidators(newValidators.addresses, newValidators.powers, sigs.signers, sigs.v, sigs.r, sigs.s, { from: accounts[0] })
      const _validators = await vmc.getValidators.call();
      assert.equal(_validators.length, newValidators.addresses.length)
    })

    it('Withdraw ether with stake >= equal to threshold', async () => {
      await depositEther(alice, amount)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ETH, alice, nonce, gateway.address, soliditySha3(amount))
      let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
      await gateway.withdrawETH(amount, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice })
    })

    it('Cannot withdraw ether with stake < threshold', async () => {
      await depositEther(alice, amount)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ETH, alice, nonce, gateway.address, soliditySha3(amount))
      let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold - 0.1)
      await shouldFail.reverting(gateway.withdrawETH(amount, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice }))
    })

    it('Cannot withdraw more than has been deposited', async () => {
      await depositEther(alice, amount)
      let invalid_amount = 10 * amount
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ETH, alice, nonce, gateway.address, soliditySha3(invalid_amount))
      let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold - 0.1)
      await shouldFail.reverting(gateway.withdrawETH(invalid_amount.toString(), sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice }))
    })

    it('Cannot reuse a sig to withdraw (invalid nonce)', async () => {
      await depositEther(alice, amount)
      await depositEther(bob, amount) // bob also deposits

      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ETH, alice, nonce, gateway.address, soliditySha3(amount))
      let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
      await gateway.withdrawETH(amount, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice });

      // alice tries to replay the signature - should fail because her
      // nonce has been used alrd
      await shouldFail.reverting(gateway.withdrawETH(amount, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice }))

      // validators update the nonce and alice can indeed withdraw bob's
      // funds (intended)
      nonce = await gateway.nonces.call(alice)
      hash = soliditySha3(TokenKind.ETH, alice, nonce, gateway.address, soliditySha3(amount))
      sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
      await gateway.withdrawETH(amount, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice });
    })

    it('Is not reentrant', async () => {
      await depositEther(alice, amount)

      // Bob's contract `reentrancy` is authorized 0.1 * amount, but he'll try to steal more.
      const reentrancy = await Reentrancy.new({ from: bob })
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ETH, reentrancy.address, nonce, gateway.address, soliditySha3(TokenKind.ETH, 0.1 * amount))
      let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
      await reentrancy.setup(gateway.address, (amount * 0.1).toString(), sigs.signers, sigs.v, sigs.r, sigs.s)
      await shouldFail.reverting(reentrancy.attack())
    })

  })

  describe('ERC20 deposit / withdrawals', async () => {

    let amount = ERC20_AMOUNT
    let largeAmount = ERC20_AMOUNT_2

    it('Withdraw bad ERC20 token', async () => {
      await depositBadERC20(alice, amount)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ERC20, alice, nonce, gateway.address, soliditySha3(amount, badToken.address))
      let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
      await gateway.withdrawERC20(amount, badToken.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice }) // if successful alice gets her badToken back
      assert.equal(await badToken.balanceOf.call(alice), amount)
    })

    it('Withdraw ERC20 with stake >= equal to threshold', async () => {
      await depositERC20(alice, amount)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ERC20, alice, nonce, gateway.address, soliditySha3(amount, erc20.address))
      let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
      await gateway.withdrawERC20(amount, erc20.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice });
      assert.equal(await erc20.balanceOf.call(alice), amount)
    })

    it('Withdraw ERC20 mintable token with stake >= equal to threshold, without deposit', async () => {
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ERC20, alice, nonce, gateway.address, soliditySha3(largeAmount, erc20Mintable.address))
      let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)

      // Gateway doesn't have any tokens so it will mint the entire withdrawal amount
      await gateway.withdrawERC20(largeAmount, erc20Mintable.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice });
      assert.equal(await erc20Mintable.balanceOf.call(alice), largeAmount)
    })

    it('Withdraw ERC20 mintable token with stake >= equal to threshold, withdraw more than the gateway has', async () => {
      await depositERC20(alice, amount)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ERC20, alice, nonce, gateway.address, soliditySha3(largeAmount, erc20Mintable.address))
      let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)

      // Gateway doesn't have enough tokens so it will mint just enough to honor the withdrawal
      await gateway.withdrawERC20(largeAmount, erc20Mintable.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice });
      assert.equal(await erc20Mintable.balanceOf.call(alice), largeAmount)
    })

    it('Withdraw ERC20 mintable token with stake >= equal to threshold, withdraw less than the gateway has', async () => {
      let less_amount = ERC20_AMOUNT_3
      await depositERC20(alice, amount)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ERC20, alice, nonce, gateway.address, soliditySha3(less_amount, erc20Mintable.address))
      let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
      
      // Gateway has enough tokens to honor the withdrawal and doesn't need to mint any
      await gateway.withdrawERC20(less_amount, erc20Mintable.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice });
      assert.equal(await erc20Mintable.balanceOf.call(alice), less_amount)
    })

    it('Withdraw ERC20 with stake < threshold', async () => {
      await depositERC20(alice, amount)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ERC20, alice, nonce, gateway.address, soliditySha3(amount, erc20.address))
      let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold - 0.1)
      await shouldFail.reverting(gateway.withdrawERC20(amount, erc20.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice }))
    })

    it('Cannot withdraw more than has been deposited', async () => {
      await depositERC20(alice, amount)
      let invalid_amount = 10 * amount
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ERC20, alice, nonce, gateway.address, soliditySha3(invalid_amount, erc20.address))
      let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
      await shouldFail.reverting(gateway.withdrawETH(invalid_amount.toString(), sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice }))
    })

    it('Cannot reuse a sig to withdraw ERC20 (invalid nonce)', async () => {
      await depositERC20(alice, amount)
      await depositERC20(bob, amount) // bob also deposits

      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ERC20, alice, nonce, gateway.address, soliditySha3(amount, erc20.address))
      let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
      await gateway.withdrawERC20(amount, erc20.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice });

      // alice tries to replay the signature - should fail because her
      // nonce has been used alrd
      await shouldFail.reverting(gateway.withdrawERC20(amount, erc20.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice }))

      // validators update the nonce and alice can indeed withdraw bob's
      // funds (intended)
      nonce = await gateway.nonces.call(alice)
      hash = soliditySha3(TokenKind.ERC20, alice, nonce, gateway.address, soliditySha3(amount, erc20.address))
      sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
      await gateway.withdrawERC20(amount, erc20.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice });
    })

    it('Loom Withdraw event', async () => {
      await depositLoom(alice, amount)
      let nonce = await gateway.nonces.call(alice)

      let hash = soliditySha3(TokenKind.ERC20, alice, nonce, gateway.address, soliditySha3(amount, loom.address))
      let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
      await gateway.withdrawERC20(amount, loom.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice });

      let withdrawn = await gateway.getPastEvents('TokenWithdrawn', { fromBlock: 0, toBlock: "latest" });
      assert.equal(withdrawn[0].args.kind, 4)
      assert.equal(await erc20.balanceOf.call(alice), amount)
    })

    async function depositLoom(from, amount) {
      await loom.approve(gateway.address, amount, { 'from': from })
      await gateway.depositERC20(amount, loom.address, { 'from': from })
      let received = await gateway.getPastEvents('LoomCoinReceived', { fromBlock: 0, toBlock: "latest" });
      received = received[0].args;
      assert.equal(received.from, from);
      assert.equal(received.amount, amount);
      assert.equal(received.loomCoinAddress, loom.address);
    }
  
    async function depositBadERC20(from, amount) {
      await badToken.approve(gateway.address, amount, { 'from': from })
      await gateway.depositERC20(amount, badToken.address, { 'from': from })
      let received = await gateway.getPastEvents('ERC20Received', { fromBlock: 0, toBlock: "latest" });
      received = received[0].args;
      assert.equal(received.from, from);
      assert.equal(received.amount, amount);
      assert.equal(received.contractAddress, badToken.address);
      let balance = await gateway.getERC20.call(badToken.address)
      assert.equal(balance, amount)
    }


  })

  describe('ERC721 deposits / withdrawals', async () => {

    let uid = 2;
    let other_uid = 4;
    let new_uid = 9999;
    
    it('Withdraw ERC721 with stake >= equal to threshold, without deposit for mintable tokens', async () => {
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ERC721, alice, nonce, gateway.address, soliditySha3(new_uid, erc721.address))
      let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
      await gateway.withdrawERC721(new_uid, erc721.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice });
      assert.equal(await erc721.ownerOf.call(new_uid), alice)
    })
  
    it('Withdraw ERC721 with stake >= equal to threshold', async () => {
      await depositERC721(alice, uid)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ERC721, alice, nonce, gateway.address, soliditySha3(uid, erc721x.address))
      let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
      await gateway.withdrawERC721(uid, erc721x.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice });

      assert.equal(await erc721x.ownerOf.call(uid), alice)
    })  

    it('Withdraw ERC721 with stake >= equal to threshold, for normal token', async () => {
      await depositNormalERC721(alice, uid)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ERC721, alice, nonce, gateway.address, soliditySha3(uid, normalErc721.address))
      let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
      await gateway.withdrawERC721(uid, normalErc721.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice });
      assert.equal(await normalErc721.ownerOf.call(uid), alice)
    })   

    it('Cannot withdraw ERC721 with stake < threshold', async () => {
      await depositERC721(alice, uid)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ERC721, alice, nonce, gateway.address, soliditySha3(uid, erc721x.address))
      let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold - 0.1)
      await shouldFail.reverting(gateway.withdrawERC721(uid, erc721x.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice }));
    })  

    it('Cannot withdraw another uid', async () => {
      await depositERC721(alice, uid)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ERC721, alice, nonce, gateway.address, soliditySha3(uid, erc721x.address))
      let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold - 0.1)
      await shouldFail.reverting(gateway.withdrawERC721(other_uid, erc721x.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice }))
    })

    it('Cannot reuse a sig to withdraw ERC721 (invalid nonce)', async () => {
      await depositERC721(alice, uid)
      await depositERC721(alice, other_uid)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ERC721, alice, nonce, gateway.address, soliditySha3(uid, erc721x.address))
      let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
      await gateway.withdrawERC721(uid, erc721x.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice });

      // alice tries to replay the signature - should fail because her
      // nonce has been used alrd
      await shouldFail.reverting(gateway.withdrawERC721(other_uid, erc721x.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice }))

      // validators update the nonce and alice can indeed withdraw bob's
      // funds (intended)
      nonce = await gateway.nonces.call(alice)
      hash = soliditySha3(TokenKind.ERC721, alice, nonce, gateway.address, soliditySha3(other_uid, erc721x.address))
      sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
      await gateway.withdrawERC721(other_uid, erc721x.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice });
    })

    async function depositNormalERC721(from, uid) {
      await normalErc721.approve(gateway.address, uid, {'from': from})
      await normalErc721.transferFrom(from, gateway.address, uid, { 'from': from })
    }

  });


  describe('ERC721x deposits / withdrawals', async () => {

    let uid = 0;
    let other_uid = 1;
    let amount = 50

    it('Withdraw ERC721x with stake >= equal to threshold', async () => {
      await depositERC721x(alice, uid, amount)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ERC721X, alice, nonce, gateway.address, soliditySha3(uid, amount, erc721x.address))
      let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
      await gateway.withdrawERC721X(uid, amount, erc721x.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice });
    })

    it('Cannot withdraw ERC721x with stake < threshold', async () => {
      await depositERC721x(alice, uid, amount)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ERC721X, alice, nonce, gateway.address, soliditySha3(uid, amount, erc721x.address))
      let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold - 0.1)
      await shouldFail.reverting(gateway.withdrawERC721X(uid, amount, erc721x.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice }));
    })

    it('Cannot withdraw another uid', async () => {
      await depositERC721x(alice, uid, amount)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ERC721X, alice, nonce, gateway.address, soliditySha3(uid, erc721x.address))
      let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
      await shouldFail.reverting(gateway.withdrawERC721X(other_uid, amount, erc721x.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice }))
    })

    it('Cannot withdraw another amount', async () => {
      await depositERC721x(alice, uid, amount - 1)
      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ERC721X, alice, nonce, gateway.address, soliditySha3(uid, amount, erc721x.address))
      let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
      await shouldFail.reverting(gateway.withdrawERC721X(uid, amount - 1, erc721x.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice }))
    })

    it('Cannot reuse a sig to withdraw ERC721 (invalid nonce)', async () => {
      await depositERC721x(alice, uid, amount)
      await depositERC721x(bob, uid, amount)

      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ERC721X, alice, nonce, gateway.address, soliditySha3(uid, amount, erc721x.address))
      let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
      await gateway.withdrawERC721X(uid, amount, erc721x.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice });

      // alice tries to replay the signature - should fail because her
      // nonce has been used alrd
      await shouldFail.reverting(gateway.withdrawERC721X(uid, amount, erc721x.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice }))

      // validators update the nonce and alice can indeed withdraw bob's
      // funds (intended)
      nonce = await gateway.nonces.call(alice)
      hash = soliditySha3(TokenKind.ERC721X, alice, nonce, gateway.address, soliditySha3(uid, amount, erc721x.address))
      sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
      await gateway.withdrawERC721X(uid, amount, erc721x.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice });
    })

  });

  describe('Disable deposits / withdrawals', async () => {
    let amount = ETHER_AMOUNT
    let amountErc20 = ERC20_AMOUNT
    let amountErc721x = 50
    let uidErc721 = 2
    let uidErc721x = 0

    it('Cannot deposit', async () => {
      await gateway.enableGateway(false, { from: validator })
      await shouldFail.reverting(depositEther(alice, amount))
      await shouldFail.reverting(depositERC20(alice, amountErc20))
      await shouldFail.reverting(depositERC721(alice, uidErc721))
      await shouldFail.reverting(depositERC721x(alice, uidErc721x, amountErc721x))
    })

    it('Cannot withdraw', async () => {
      await depositEther(alice, amount)
      await depositERC20(alice, amountErc20)
      await depositERC721(alice, uidErc721)
      await depositERC721x(alice, uidErc721x, amountErc721x)

      await gateway.enableGateway(false, { from: validator })

      let nonce = await gateway.nonces.call(alice)
      let hash = soliditySha3(TokenKind.ETH, alice, nonce, gateway.address, soliditySha3(amount))
      let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
      await shouldFail.reverting(gateway.withdrawETH(amount, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice }))

      nonce = await gateway.nonces.call(alice)
      hash = soliditySha3(TokenKind.ERC20, alice, nonce, gateway.address, soliditySha3(amountErc20, erc20.address))
      sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
      await shouldFail.reverting(gateway.withdrawERC20(amountErc20, erc20.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice }));

      nonce = await gateway.nonces.call(alice)
      hash = soliditySha3(TokenKind.ERC721, alice, nonce, gateway.address, soliditySha3(uidErc721, normalErc721.address))
      sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
      await shouldFail.reverting(gateway.withdrawERC721(uidErc721, normalErc721.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice }));

      nonce = await gateway.nonces.call(alice)
      hash = soliditySha3(TokenKind.ERC721X, alice, nonce, gateway.address, soliditySha3(uidErc721x, amountErc721x, erc721x.address))
      sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
      await shouldFail.reverting(gateway.withdrawERC721X(uidErc721x, amountErc721x, erc721x.address, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice }));
    })

  });

  describe('Gateway permissions', async () => {

    it('Only owner can disable and enable gateway', async () => {
      let gatewayStatus = await gateway.getGatewayEnabled()
      assert.equal(gatewayStatus, true, "gateway status has to be true")

      await gateway.enableGateway(false, {from: owner})
      gatewayStatus = await gateway.getGatewayEnabled()
      assert.equal(gatewayStatus, false, "gateway status has to be false")

      await shouldFail.reverting(gateway.enableGateway(true, {from: notOwner}))
    })

    it('Only owner can disable and enable toggleAllowAnyToken', async () => {
      let allowAnyTokenStatus = await gateway.getAllowAnyToken()
      await assert.equal(allowAnyTokenStatus, true, "allowAnyTokenStatus has to be true")

      await gateway.toggleAllowAnyToken(false, {from: owner})
      allowAnyTokenStatus = await gateway.getAllowAnyToken()
      await assert.equal(allowAnyTokenStatus, false, "allowAnyTokenStatus has to be false")

      await shouldFail.reverting(gateway.toggleAllowAnyToken(true, {from: notOwner}))
    })

    it('Only owner can disable and enable toggleAllowToken', async () => {
      await gateway.toggleAllowAnyToken(false, {from: owner})

      let isTokenAllowedStatus = await gateway.isTokenAllowed(erc20.address)
       assert.equal(isTokenAllowedStatus, false, "isTokenAllowedStatus has to be false")

      await gateway.toggleAllowToken(erc20.address, true, {from: owner})
      isTokenAllowedStatus = await gateway.isTokenAllowed(erc20.address)
      assert.equal(isTokenAllowedStatus, true, "isTokenAllowedStatus has to be true")

      await shouldFail.reverting(gateway.toggleAllowToken(erc20.address, false, {from: notOwner}))
    })
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
