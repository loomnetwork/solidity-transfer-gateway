const Gateway = artifacts.require('Gateway')
const { soliditySha3 } = require('web3-utils')
const { signHash, Promisify } = require('./utils.js')

import assertRevert from './helpers/assertRevert.js'

contract('Transfer Gateway V2', async (accounts) => {
  let gateway;

  let createSigs = async function(signers, validator, operation) {

      let nonce = await gateway.nonce();
      let hash = soliditySha3(gateway.address, nonce, soliditySha3(operation, validator))

      let sigV = []
      let sigR = []
      let sigS = []

      for (var i=0; i<signers.length; i++) {
          let sig =  (await web3.eth.sign(hash, signers[i])).slice(2);
          let r = '0x' + sig.substring(0, 64)
          let s = '0x' + sig.substring(64, 128)
          let v = parseInt(sig.substring(128, 130), 16) + 27
          sigV.push(v)
          sigR.push(r)
          sigS.push(s)
      }

      return {sigV: sigV, sigR: sigR, sigS: sigS}
  }

  describe('Governance of Validators', async () => {

    let getThreshold = function(n, num, denom) {
      return parseInt(n * num / denom);
    }

    it('Adds a new validator with 3/4 majority out of 5 validators (3 needed)', async () => {
      let n = 5;
      let threshold = getThreshold(n, 3, 4);
      let validators = accounts.slice(0, n);
      let new_validator = accounts[9];
      gateway = await Gateway.new('0x0000000000000000000000000000000000000000', validators, 3, 4, [], []);
      assert.equal(await gateway.numValidators(), n);

      let signers = accounts.slice(0, threshold);
      signers.sort();
      let sigs = await createSigs(signers, new_validator, 'add');
      await gateway.addValidator(new_validator, sigs.sigV, sigs.sigR, sigs.sigS)
      assert.equal(await gateway.numValidators(), n+1);
    });

    it('Removes a validator with 3/4 majority out of 5 validators (3 needed)', async () => {
      let n = 5;
      let threshold = getThreshold(n, 3, 4);
      let validators = accounts.slice(0, n);
      let deleted_validator = accounts[n-1]; // will remove the last validator
      gateway = await Gateway.new('0x0000000000000000000000000000000000000000', validators, 3, 4, [], []);

      let signers = accounts.slice(0, threshold);
      signers.sort();
      let sigs = await createSigs(signers, deleted_validator, 'remove');
      await gateway.removeValidator(deleted_validator, sigs.sigV, sigs.sigR, sigs.sigS);
      assert.equal(await gateway.numValidators(), n-1);
    });


    it('Fails to add a new validator without majority', async () => {
      let n = 5;
      let threshold = getThreshold(n, 3, 4);
      let validators = accounts.slice(0, n);
      let new_validator = accounts[9];
      gateway = await Gateway.new('0x0000000000000000000000000000000000000000', validators, 3, 4, [], []);
      assert.equal(await gateway.numValidators(), n);

      let signers = accounts.slice(0, threshold-1);
      signers.sort();
      let sigs = await createSigs(signers, new_validator, 'add');
      assertRevert(gateway.addValidator(new_validator, sigs.sigV, sigs.sigR, sigs.sigS))
      assert.equal(await gateway.numValidators(), n);
    });

    it('Adds a new validator with >3/4 majority out of 5 validators', async () => {
      let n = 5;
      let threshold = getThreshold(n, 3, 4);
      let validators = accounts.slice(0, n);
      let new_validator = accounts[9];
      gateway = await Gateway.new('0x0000000000000000000000000000000000000000', validators, 3, 4, [], []);

      let signers = accounts.slice(0, threshold+1);
      signers.sort();
      let sigs = await createSigs(signers, new_validator, 'add');
      await gateway.addValidator(new_validator, sigs.sigV, sigs.sigR, sigs.sigS)
    });

    it('Adds a new validator, removes new validator, and re-adds same validator', async () => {
      let n = 5;
      let threshold = getThreshold(n, 3, 4);
      let validators = accounts.slice(0, n);
      let new_validator = accounts[9];
      gateway = await Gateway.new('0x0000000000000000000000000000000000000000', validators, 3, 4, [], []);

      let signers = accounts.slice(0, threshold+1);
      signers.sort();
      let sigs = await createSigs(signers, new_validator, 'add');
      await gateway.addValidator(new_validator, sigs.sigV, sigs.sigR, sigs.sigS);
      sigs = await createSigs(signers, new_validator, 'remove');
      await gateway.removeValidator(new_validator, sigs.sigV, sigs.sigR, sigs.sigS);
      sigs = await createSigs(signers, new_validator, 'add');
      await gateway.addValidator(new_validator, sigs.sigV, sigs.sigR, sigs.sigS);
    });

    it('Validators: A -> Validators: A,B -> Validators: B', async () => {
      let n = 1;
      let threshold = getThreshold(n, 3, 4);
      let validators = accounts.slice(0, n);
      let new_validator = accounts[9];
      gateway = await Gateway.new('0x0000000000000000000000000000000000000000', validators, 3, 4, [], []);

      let signers = accounts.slice(0, threshold+1);
      signers.sort();
      let sigs = await createSigs(signers, new_validator, 'add');
      await gateway.addValidator(new_validator, sigs.sigV, sigs.sigR, sigs.sigS);
      sigs = await createSigs(signers, validators[0], 'remove');
      await gateway.removeValidator(validators[0], sigs.sigV, sigs.sigR, sigs.sigS);
      assert.equal(await gateway.numValidators(), 1)
      let addedValidator = (await gateway.getPastEvents('AddedValidator', { fromBlock: 0, toBlock: "latest"}))
      let removedValidator = (await gateway.getPastEvents('RemovedValidator', { fromBlock: 0, toBlock: "latest"}))
      assert.equal(addedValidator[0].args.validator, removedValidator[0].args.validator)
      assert.equal(addedValidator[1].args.validator, new_validator)
       
    });

    it('Rejects replayed commands', async () => {
      let n = 5;
      let threshold = getThreshold(n, 3, 4);
      let validators = accounts.slice(0, n);
      let new_validator = accounts[9];
      gateway = await Gateway.new('0x0000000000000000000000000000000000000000', validators, 3, 4, [], []);

      let signers = accounts.slice(0, threshold+1);
      signers.sort();
      let addSigs = await createSigs(signers, new_validator, 'add');
      await gateway.addValidator(new_validator, addSigs.sigV, addSigs.sigR, addSigs.sigS);
      let removeSigs = await createSigs(signers, new_validator, 'remove');
      await gateway.removeValidator(new_validator, removeSigs.sigV, removeSigs.sigR, removeSigs.sigS);
      // A replay attempt using old signatures for `addValidator`.
      assertRevert(gateway.addValidator(new_validator, addSigs.sigV, addSigs.sigR, addSigs.sigS));
      assert.equal(await gateway.numValidators(), n);
    });

    it('Fails to add new validator with 0 sigs, 3/4 majority out of 1 total validators (1 needed)', async () => {
      let n = 1;
      let validators = accounts.slice(0, n);
      let new_validator = accounts[9];
      gateway = await Gateway.new('0x0000000000000000000000000000000000000000', validators, 3, 4, [], []);
      assert.equal(await gateway.numValidators(), n);

      let signers = [];
      signers.sort();
      let sigs = await createSigs(signers, new_validator, 'add');
      assertRevert(gateway.addValidator(new_validator, sigs.sigV, sigs.sigR, sigs.sigS));
      assert.equal(await gateway.numValidators(), n);
    });

  });

})
