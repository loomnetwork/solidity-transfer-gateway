const ethutil = require('ethereumjs-util')
const keythereum = require("keythereum");

const assertEventVar = (transaction, eventName, eventVar, equalVar) => {
      const event = transaction.logs.find(log => log.event === eventName);
      assert.equal(event.args[eventVar], equalVar, `Event ${event.args[eventVar]} didn't happen`);
};

const Promisify = (inner) =>
  new Promise((resolve, reject) =>
    inner((err, res) => {
      if (err) {
        reject(err)
      } else {
        resolve(res)
      }
    })
  )


/********** UTILS ********/

// Helper functions for the VMC -- taken from
// https://github.com/cosmos/peggy/blob/master/ethereum-contracts/test/utils.js

function createValidators(size) {
    var newValidators = {
        addresses: [],
        pubKeys: [],
        privateKeys: [],
        powers: [],
        totalPower: 0
    };

    let privateKey,hexPrivate, pubKey, address, power;

    for (var i = 0; i < size; i++) {
        privateKey = keythereum.create().privateKey;
        hexPrivate = ethutil.bufferToHex(privateKey);
        address = ethutil.bufferToHex(ethutil.privateToAddress(privateKey));
        pubKey = ethutil.bufferToHex(ethutil.privateToPublic(privateKey));
        power = seededRandomInt(1, 50, i);

        newValidators.addresses.push(address);
        newValidators.privateKeys.push(hexPrivate);
        newValidators.pubKeys.push(pubKey);
        newValidators.powers.push(power);
        newValidators.totalPower += power;
    }

    return newValidators;
}

function seededRandomInt(min, max, seed) {
    seed = (seed * 9301 + 49297) % 233280;
    var rnd = seed / 233280;

    return Math.floor(min + rnd * (max - min));
}

function assignPowersToAccounts(accounts) {
    var newValidators = {
        addresses: accounts,
        powers: []
    };

    for (var i = 0; i < accounts.length; i++) {
        newValidators.powers.push(seededRandomInt(1, 50, i));
    }
    return newValidators;
}

async function createSigns(validators, data, totalPower, percentSign) {
    var vArray = [], rArray = [], sArray = [], signers = [];
    var signedPower = 0;
    if (!percentSign) {
        percentSign = 0.95
    }
    for (var i = 0; i < validators.addresses.length; i++) {
        if (signedPower >= percentSign * totalPower) {
            break
        }
        let signature = await web3.eth.accounts.sign(data, validators.privateKeys[i]);

        vArray.push(signature.v);
        rArray.push(ethutil.bufferToHex(signature.r));
        sArray.push(ethutil.bufferToHex(signature.s));
        signers.push(i);
        signedPower += validators.powers[i];
    }

    return {
        signers: signers,
        v: vArray,
        r: rArray,
        s: sArray,
        signedPower: signedPower
    }
}

module.exports = {
  assertEventVar,
  Promisify,
  createValidators,
  createSigns,
}
