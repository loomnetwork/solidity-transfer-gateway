const ValidatorManagerContact = artifacts.require('ValidatorManagerContract')
const { createSigns, createValidators } = require('./utils.js')
const { shouldFail } = require('openzeppelin-test-helpers');
const { soliditySha3 } = require('web3-utils')

contract('Validator Manager Contract', async (accounts) => {
    let vmc;
    let validators;
    let validatorsTotalPower;

    const threshold = 0.67


    beforeEach('Setup contract', async function() {
        validators = createValidators(21);
        vmc = await ValidatorManagerContact.new(validators.addresses, validators.powers, 2, 3, "0x0000000000000000000000000000000000000000");
        validatorsTotalPower = await vmc.totalPower.call();
    });

    it("Deploys with 1 validator, upgrades to 4-of-4 with equal power and upgrades to 2-3rds weighted by power multisig", async function() {
        validators = createValidators(1);

        // Deploy with 1/1 quorum and 1 validators
        vmc = await ValidatorManagerContact.new(validators.addresses, validators.powers, 1, 1, "0x0000000000000000000000000000000000000000");
        validatorsTotalPower = await vmc.totalPower.call();

        let _validators = await vmc.getValidators.call();
        assert.equal(_validators.length, 1)
        let _powers = await vmc.getPowers.call();
        let _totalPower = await vmc.totalPower.call();

        let multisig = createValidators(4);
        // set all powers for 4 validators to 100
        for (let i in multisig.powers) {
            multisig.powers[i] = 100
        }

        let nonce = await vmc.nonce.call()
        let newVSetHash = soliditySha3(
            vmc.address,
            nonce,
            soliditySha3(
                {t: 'address[]', v: multisig.addresses},
                {t: 'uint256[]', v: multisig.powers}
            )
        )
        let sigs = await createSigns(validators, newVSetHash, validatorsTotalPower, 1)
        await vmc.rotateValidators(multisig.addresses, multisig.powers, sigs.signers, sigs.v, sigs.r, sigs.s, { from: accounts[0] })
        _validators = await vmc.getValidators.call();
        assert.equal(_validators.length, 4)



        // Change quorum to 2/3rds
        nonce = await vmc.nonce.call()
        let newQuorumHash = soliditySha3(
            vmc.address,
            nonce,
            soliditySha3(
                {t: 'uint8', v: 2}, 
                {t: 'uint8', v: 3}
            )
        )
        sigs = await createSigns(multisig, newQuorumHash, 400, 1)
        await vmc.setQuorum(2, 3, sigs.signers, sigs.v, sigs.r, sigs.s, { from: accounts[0] }) 
        assert.equal(await vmc.threshold_num(), 2)
        assert.equal(await vmc.threshold_denom(), 3)


        let newValidators = createValidators(11);
        nonce = await vmc.nonce.call()
        newVSetHash = soliditySha3(
            vmc.address,
            nonce,
            soliditySha3(
                {t: 'address[]', v: newValidators.addresses},
                {t: 'uint256[]', v: newValidators.powers}
            )
        )
        sigs = await createSigns(multisig, newVSetHash, 400, 1)
        await vmc.rotateValidators(newValidators.addresses, newValidators.powers, sigs.signers, sigs.v, sigs.r, sigs.s, { from: accounts[0] })
        _validators = await vmc.getValidators.call();
        assert.equal(_validators.length, 11)

    })

    it('Initialization of the contract', async () => {
        let _validators = await vmc.getValidators.call();
        let _powers = await vmc.getPowers.call();
        let _totalPower = await vmc.totalPower.call();

        assert.strictEqual(_validators.length, _powers.length, "Both initial arrays must have the same length");
        assert.isAtMost(_validators.length, 100, "Validator set should not be larger than 100");

        for (var i = 0; i < validators.addresses.length; i++) {
            assert.equal(_validators[i].toLowerCase(), validators.addresses[i], "Validators not set correctly")
            assert.equal(_powers[i].toNumber(), validators.powers[i], "Powers not set correctly");
        }

        assert.equal(_totalPower.toNumber(), validators.totalPower, "Total power not set correctly")
    });

    it('Rejects duplicate signatures', async () => {
        // even if the quorum required is 10%, you cannot present duplicate sigs to try to augment your voting power
        vmc = await ValidatorManagerContact.new(validators.addresses, validators.powers, 1, 10, "0x0000000000000000000000000000000000000000");
        let nonce = await vmc.nonce.call()
        let hash = soliditySha3("0xdeadbeef")
        let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)

        // Overwrite all signatures with validator 0 sigs.
        let duplicateSigs = sigs
        for (let i in sigs.signers) {
            duplicateSigs.signers[i] = sigs.signers[0]
            duplicateSigs.v[i] = sigs.v[0]
            duplicateSigs.r[i] = sigs.r[0]
            duplicateSigs.s[i] = sigs.s[0]
        }

        assert.equal(await vmc.nonce(), 0, "nonce should start as 0")

        await shouldFail.reverting(vmc.checkThreshold(hash, duplicateSigs.signers, duplicateSigs.v, duplicateSigs.r, duplicateSigs.s, { from: accounts[0] }))

        assert.equal(await vmc.nonce(), 0, "nonce shouldnt increment")
    })

    it('Signs on new validators', async () => {
        let newValidators = createValidators(21);
        let nonce = await vmc.nonce.call()
        let hash = soliditySha3(
            vmc.address,
            nonce,
            soliditySha3(
                {t: 'address[]', v: newValidators.addresses}, 
                {t: 'uint256[]', v: newValidators.powers}
            )
        )
        let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
        
        await vmc.rotateValidators(newValidators.addresses, newValidators.powers, sigs.signers, sigs.v, sigs.r, sigs.s, { from: accounts[0] }) 

        let _validators = await vmc.getValidators.call();
        let _powers = await vmc.getPowers.call();
        let _totalPower = await vmc.totalPower.call();

        for (var i = 0; i < newValidators.addresses.length; i++) {
            assert.equal(_validators[i].toLowerCase(), newValidators.addresses[i], "Validators not set correctly")
            assert.equal(_powers[i].toNumber(), newValidators.powers[i], "Powers not set correctly");
        }

        assert.equal(_totalPower.toNumber(), newValidators.totalPower, "Total power not set correctly")
    })

    it('Fails to sign without enough votes', async () => {
        let newValidators = createValidators(21);
        let nonce = await vmc.nonce.call()
        let hash = soliditySha3(
            vmc.address,
            nonce,
            soliditySha3(
                {t: 'address[]', v: newValidators.addresses}, 
                {t: 'uint256[]', v: newValidators.powers}
            )
        )
        let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold - 0.1)
        
        await shouldFail.reverting(vmc.rotateValidators(newValidators.addresses, newValidators.powers, sigs.signers, sigs.v, sigs.r, sigs.s, { from: accounts[0] }))
    })

    it('Cannot replay signature of an earlier election', async () => {
        let e1_Validators = createValidators(21);
        let nonce = await vmc.nonce.call()
        let hash = soliditySha3(
            vmc.address,
            nonce,
            soliditySha3(
                {t: 'address[]', v: e1_Validators.addresses}, 
                {t: 'uint256[]', v: e1_Validators.powers}
            )
        )
        let sigs = await createSigns(validators, hash, validatorsTotalPower, threshold)
        
        await vmc.rotateValidators(e1_Validators.addresses, e1_Validators.powers, sigs.signers, sigs.v, sigs.r, sigs.s, { from: accounts[0] }) 

        let e2_Validators = createValidators(21);

        // Elections happened once. 
        // now they happen again and e1 need to sign

        nonce = await vmc.nonce.call()
        hash = soliditySha3(
            vmc.address,
            nonce,
            soliditySha3(
                {t: 'address[]', v: e2_Validators.addresses}, 
                {t: 'uint256[]', v: e2_Validators.powers}
            )
        )
        let e2_sigs = await createSigns(e1_Validators, hash, validatorsTotalPower, threshold)

        await vmc.rotateValidators(e2_Validators.addresses, e2_Validators.powers, e2_sigs.signers, e2_sigs.v, e2_sigs.r, e2_sigs.s, { from: accounts[0] })

        // now E1 wants to try to regain control maliciously
        await shouldFail.reverting(vmc.rotateValidators(e1_Validators.addresses, e1_Validators.powers, sigs.signers, sigs.v, sigs.r, sigs.s, { from: accounts[0] }))

    })

})
