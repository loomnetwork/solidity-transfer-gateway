const BEP2TokenTemplate = artifacts.require('BEP2TokenTemplate.sol')

const { expectThrow } = require('./helpers')

const MINT_BALANCE = 10000000

const delay = ms => new Promise(resolve => setTimeout(resolve, ms));

const TOKEN_NAME = "template"
const SYMBOL = "TEMP"

contract('BEP2TokenTemplate', accounts => {
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

    let bep2TokenTemplate
    
    beforeEach(async () => {
        bep2TokenTemplate = await BEP2TokenTemplate.new(gateway1, TOKEN_NAME, SYMBOL)
        await bep2TokenTemplate.addValidator(validator1, {from: creator})
    })

    it(`only validator can add validators`, async () => {
        await bep2TokenTemplate.addValidator(validator2, {from: validator1})
        await expectThrow(bep2TokenTemplate.addValidator(validator3, {from:notValidator}))
    })
    
    it(`only validator can add or remove gateways`, async () => {
        await bep2TokenTemplate.addGateway(gateway2, {from: validator1})
        await expectThrow(bep2TokenTemplate.addGateway(gateway3, {from:notValidator}))
        await bep2TokenTemplate.removeGateway(gateway2, {from: validator1})
        await expectThrow(bep2TokenTemplate.removeGateway(gateway3, {from:notValidator}))
    })

    it(`only gateway can use mintToGateway()`, async () => {
        gateway1BalanceBefore = await bep2TokenTemplate.balanceOf(gateway1)
        await bep2TokenTemplate.mintToGateway(MINT_BALANCE, {from: gateway1})
        gateway1BalanceAfter = await bep2TokenTemplate.balanceOf(gateway1)
        assert.strictEqual(gateway1BalanceAfter.toNumber() - gateway1BalanceBefore.toNumber(), MINT_BALANCE, `gateway is expected to mint {$MINT_BALANCE}`)

        await expectThrow(bep2TokenTemplate.mintToGateway(MINT_BALANCE, {from: gateway2}))
        await bep2TokenTemplate.addGateway(gateway3, {from: validator1})

        gateway3BalanceBefore = await bep2TokenTemplate.balanceOf(gateway3)
        await bep2TokenTemplate.mintToGateway(MINT_BALANCE, {from: gateway3})

        gateway3BalanceAfter = await bep2TokenTemplate.balanceOf(gateway3)
        assert.strictEqual(gateway3BalanceAfter.toNumber() - gateway3BalanceBefore.toNumber(), MINT_BALANCE, `gateway is expected to mint {$MINT_BALANCE}`)


        await expectThrow(bep2TokenTemplate.mintToGateway(MINT_BALANCE, {from: notGateway}))

    })

});