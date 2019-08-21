const SampleBEP2Token = artifacts.require('SampleBEP2Token.sol')

const { expectThrow } = require('./helpers')

const MINT_BALANCE = 10000000

const delay = ms => new Promise(resolve => setTimeout(resolve, ms));

const TOKEN_NAME = "SampleBEP2Token"
const SYMBOL = "BEP2"

contract('SampleBEP2Token', accounts => {
    const [
        creator,
        validator1,
        validator2,
        validator3,
        gateway1,
        gateway2,
        gateway3,
        notValidator,
        notGateway
    ] = accounts

    let sampleBEP2Token
    
    beforeEach(async () => {
        sampleBEP2Token = await SampleBEP2Token.new(gateway1, TOKEN_NAME, SYMBOL)
        await sampleBEP2Token.addValidator(validator1, {from: creator})      
    })

    it(`only validator can add validators`, async () => {
        await sampleBEP2Token.addValidator(validator2, {from: validator1})
        await expectThrow(sampleBEP2Token.addValidator(validator3, {from:notValidator}))
    })
    
    it(`only validator can add or remove gateways`, async () => {
        await sampleBEP2Token.addGateway(gateway2, {from: validator1})
        await expectThrow(sampleBEP2Token.addGateway(gateway3, {from:notValidator}))
        await sampleBEP2Token.removeGateway(gateway2, {from: validator1})
        await expectThrow(sampleBEP2Token.removeGateway(gateway3, {from:notValidator}))
    })

    it(`only gateway can use mintToGateway()`, async () => {
        gateway1BalanceBefore = await sampleBEP2Token.balanceOf(gateway1)
        await sampleBEP2Token.mintToGateway(MINT_BALANCE, {from: gateway1})
        gateway1BalanceAfter = await sampleBEP2Token.balanceOf(gateway1)
        assert.strictEqual(gateway1BalanceAfter.toNumber() - gateway1BalanceBefore.toNumber(), MINT_BALANCE, `gateway is expected to mint {$MINT_BALANCE}`)

        await expectThrow(sampleBEP2Token.mintToGateway(MINT_BALANCE, {from: gateway2}))
        await sampleBEP2Token.addGateway(gateway3, {from: validator1})

        gateway3BalanceBefore = await sampleBEP2Token.balanceOf(gateway3)
        await sampleBEP2Token.mintToGateway(MINT_BALANCE, {from: gateway3})

        gateway3BalanceAfter = await sampleBEP2Token.balanceOf(gateway3)
        assert.strictEqual(gateway3BalanceAfter.toNumber() - gateway3BalanceBefore.toNumber(), MINT_BALANCE, `gateway is expected to mint {$MINT_BALANCE}`)


        await expectThrow(sampleBEP2Token.mintToGateway(MINT_BALANCE, {from: notGateway}))

    })

});