const ethutil = require('ethereumjs-util')
const ethers = require('ethers')
const keythereum = require("keythereum");
const fs = require('fs')
const Web3 = require('web3')

if (typeof web3 !== 'undefined') {
  web3 = new Web3(web3.currentProvider);
} 
else{
  // for migration as truffle migrate doesn't provide web3
  web3 = new Web3(new Web3.providers.HttpProvider("http://localhost:8545"));
}

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

async function createSigners(signerFilePath) {
  let signersFile = signerFilePath

  if (!signersFile) {
    throw new Error("No signer file found.")
  } 

  var newSigners = {
    addresses: [],
    privateKeys: [],
    powers: [],
    totalPower: 0,
  };

  const signers = JSON.parse(fs.readFileSync(signersFile, "utf8")) 
  let privateKeys = signers.privateKeys;

  for (var i = 0; i < privateKeys.length; i++){
    let privateKey = privateKeys[i]

    let address = signers.addresses[i]
    let power = signers.powers[i]
    newSigners.addresses.push(address);
    newSigners.privateKeys.push(privateKey);
    newSigners.powers.push(power);
    newSigners.totalPower += power
  }
  console.log("signer", newSigners)
  return newSigners
}

async function createAllSigns(newSigners, data){
  var vArray = [], rArray = [], sArray = [], signers = [];
  for (var i = 0; i < newSigners.addresses.length; i++) {
    let signature = await web3.eth.accounts.sign(data, newSigners.privateKeys[i]);

    vArray.push(signature.v);
    rArray.push(ethutil.bufferToHex(signature.r));
    sArray.push(ethutil.bufferToHex(signature.s));
    signers.push(i);

  }

  return {
    signers: signers,
    v: vArray,
    r: rArray,
    s: sArray
  }
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

async function createSignsWithValidators(signers, data, totalPower, percentSign, validators) {
  var vArray = [], rArray = [], sArray = [], signerArray = [];
  var signedPower = 0;
  if (!percentSign) {
      percentSign = 0.95
  }

  for (var i = 0; i < signers.addresses.length; i++) {
      if (signedPower >= percentSign * totalPower) {
          break
      }
      let signature = await web3.eth.accounts.sign(data, signers.privateKeys[i]);
      let validatorIndex = validators.indexOf(signers.addresses[i])
      if (validatorIndex == -1) {
        continue
      }

      vArray[validatorIndex] = signature.v
      rArray[validatorIndex] = ethutil.bufferToHex(signature.r)
      sArray[validatorIndex] = ethutil.bufferToHex(signature.s)

      signerArray[validatorIndex] = validatorIndex
      signedPower += signers.powers[validatorIndex];
  }

  return {
      signers: signerArray,
      v: vArray,
      r: rArray,
      s: sArray,
      signedPower: signedPower,
  }
}


// THIS IS FOR SIGNLE SIG ONLY
async function signHash(from, hash) {
  let sig = (await web3.eth.sign(hash, from)).slice(2)
  let r = ethutil.toBuffer('0x' + sig.substring(0, 64))
  let s = ethutil.toBuffer('0x' + sig.substring(64, 128))
  let v = ethutil.toBuffer(parseInt(sig.substring(128, 130), 16) + 27)
  let mode = ethutil.toBuffer(1) // mode = geth
  let signature = '0x' + Buffer.concat([mode, r, s, v]).toString('hex')
  return signature
}


const expectThrow = async (promise) => {
  try {
    await promise;
  } catch (error) {
    const invalidOpcode = error.message.search('invalid opcode') >= 0;
    const invalidJump = error.message.search('invalid JUMP') >= 0;
    const outOfGas = error.message.search('out of gas') >= 0;
    const revert = error.message.search('revert') >= 0;

    assert(
      invalidOpcode || invalidJump || outOfGas || revert,
      "Expected throw, got '" + error + "' instead",
    );
    return;
  }

  assert.fail('Expected throw not received');
};


module.exports = {
  assertEventVar,
  Promisify,
  createValidators,
  createSigns,
  expectThrow,
  createSigners,
  createAllSigns,
  signHash,
  createSignsWithValidators,
}
