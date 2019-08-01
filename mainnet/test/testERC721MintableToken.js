const SampleERC721Token = artifacts.require("SampleERC721MintableToken")

const { expectThrow } = require('./utils.js')

const TOKEN_NAME = "SAMPLETOKEN"

const SYMBOL = "SMPL"

contract('SampleERC721Token', accounts => {
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
        anAddress,
    ] = accounts

    let sampleERC721Token
    
    beforeEach(async () => {
        sampleERC721Token = await SampleERC721Token.new(gateway1, TOKEN_NAME, SYMBOL, {from: creator})
        await sampleERC721Token.addValidator(validator1, {from: creator})
    })

    it(`only validator can add or remove validators`, async () => {
        await sampleERC721Token.addValidator(validator2, {from: validator1})
        await sampleERC721Token.removeValidator(validator2, {from: validator1})
        await expectThrow(sampleERC721Token.addValidator(validator3, {from:notValidator}))
        await expectThrow(sampleERC721Token.removeValidator(validator2, {from:notValidator}))
    })
    
    it(`only validator can add or remove gateways`, async () => {
        await sampleERC721Token.addGateway(gateway2, {from: validator1})
        await expectThrow(sampleERC721Token.addGateway(gateway3, {from:notValidator}))
        await sampleERC721Token.removeGateway(gateway2, {from: validator1})
        await expectThrow(sampleERC721Token.removeGateway(gateway3, {from:notValidator}))
    })

    it(`only gateway can use mintTo()`, async () => {
        let uid  = 1
        let another_uid = 2
        let other_uid = 99
 
        anAddressBalanceBefore = await sampleERC721Token.balanceOf(anAddress)
        await sampleERC721Token.mintTo(anAddress, uid, {from: gateway1})
        anAddressBalanceAfter = await sampleERC721Token.balanceOf(anAddress)
        assert.equal(await sampleERC721Token.ownerOf.call(uid), anAddress)

        await expectThrow(sampleERC721Token.mintTo(anAddress, another_uid, {from: gateway2}))
        await sampleERC721Token.addGateway(gateway3, {from: validator1})

        anAddressBalanceBefore = await sampleERC721Token.balanceOf(anAddress)
        await sampleERC721Token.mintTo(anAddress, another_uid, {from: gateway3})

        anAddressBalanceAfter = await sampleERC721Token.balanceOf(anAddress)
        assert.equal(await sampleERC721Token.ownerOf.call(another_uid), anAddress)

        await expectThrow(sampleERC721Token.mintTo(anAddress, other_uid, {from: notGateway}))
    })

    it(`only validator can use mint()`, async () => {
        let uuid = 9992
        let uuid2 = 1000
        await sampleERC721Token.mint(validator1, uuid, {from: validator1})
        assert.equal(await sampleERC721Token.ownerOf.call(uuid), validator1)

        await expectThrow(sampleERC721Token.mint(validator2, uuid2, {from: notValidator}))
    })

});