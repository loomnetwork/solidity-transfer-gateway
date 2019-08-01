const SampleERC20MintableToken = artifacts.require('SampleERC20MintableToken.sol')

const { expectThrow } = require('./utils')

const MINT_BALANCE = 10000000

const delay = ms => new Promise(resolve => setTimeout(resolve, ms));

const TOKEN_NAME = "SampleERC20MintableToken"
const SYMBOL = "SMPL"

contract('SampleERC20MintableToken', accounts => {
    const [
        creator,
        validator1,
        validator2,
        validator3,
        gateway1,
        gateway2,
        gateway3,
        notValidator,
        notGateway,
        withdrawer1
    ] = accounts

    let sampleERC20MintableToken
    
    beforeEach(async () => {
        sampleERC20MintableToken = await SampleERC20MintableToken.new(withdrawer1, TOKEN_NAME, SYMBOL)
        await sampleERC20MintableToken.addValidator(validator1, {from: creator})  
        await sampleERC20MintableToken.addGateway(gateway1, {from: validator1})    
    })

    it(`only validator can add or remove validators`, async () => {
        await sampleERC20MintableToken.addValidator(validator2, {from: validator1})
        await sampleERC20MintableToken.removeValidator(validator2, {from: validator1})
        await expectThrow(sampleERC20MintableToken.addValidator(validator3, {from:notValidator}))
        await expectThrow(sampleERC20MintableToken.removeValidator(validator2, {from:notValidator}))
    })
    
    it(`only validator can add or remove gateways`, async () => {
        await sampleERC20MintableToken.addGateway(gateway2, {from: validator1})
        await sampleERC20MintableToken.removeGateway(gateway2, {from: validator1})
        await expectThrow(sampleERC20MintableToken.addGateway(gateway3, {from:notValidator}))
        await expectThrow(sampleERC20MintableToken.removeGateway(gateway3, {from:notValidator}))
    })

    

    it(`only gateway can use mintTo()`, async () => {
        let withdrawer1BalanceBefore = await sampleERC20MintableToken.balanceOf(withdrawer1)
        await sampleERC20MintableToken.mintTo(withdrawer1, MINT_BALANCE, {from: gateway1})
        let withdrawer1BalanceAfter = await sampleERC20MintableToken.balanceOf(withdrawer1)
        assert.strictEqual(withdrawer1BalanceAfter.toNumber() - withdrawer1BalanceBefore.toNumber(), MINT_BALANCE, `gateway is expected to mint {$MINT_BALANCE} to {$withdrawer1}`)

        await expectThrow(sampleERC20MintableToken.mintTo(withdrawer1, MINT_BALANCE, {from: gateway2}))
        await sampleERC20MintableToken.addGateway(gateway3, {from: validator1})

        withdrawer1BalanceBefore = await sampleERC20MintableToken.balanceOf(withdrawer1)
        await sampleERC20MintableToken.mintTo(withdrawer1, MINT_BALANCE, {from: gateway3})

        withdrawer1BalanceAfter = await sampleERC20MintableToken.balanceOf(withdrawer1)
        assert.strictEqual(withdrawer1BalanceAfter.toNumber() - withdrawer1BalanceBefore.toNumber(), MINT_BALANCE, `gateway is expected to mint {$MINT_BALANCE} to {$withdrawer1}`)

        await expectThrow(sampleERC20MintableToken.mintTo(withdrawer1, MINT_BALANCE, {from: notGateway}))

    })

    it(`only validator can use mint()`, async () => {
        let balacneBefore =  await sampleERC20MintableToken.balanceOf(withdrawer1)
        await sampleERC20MintableToken.mint(withdrawer1, MINT_BALANCE, {from: validator1})
        let balanceAfter = await sampleERC20MintableToken.balanceOf(withdrawer1)
        assert.strictEqual(balanceAfter - balacneBefore, MINT_BALANCE,`validator1 is expected to mint`)

        await expectThrow(sampleERC20MintableToken.mint(withdrawer1, MINT_BALANCE, {from: notValidator}))
    })

});