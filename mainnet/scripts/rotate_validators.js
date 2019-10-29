/* Rotate validator script
 Usage Examples
 SIGNER_FILE=`pwd`/mainnet/test/signers.json \
 NEW_VALIDATOR_FILE=`pwd`/mainnet/test/newValidators.json \
 truffle exec scripts/rotate_validators.js --network <network> \
 <gateway_address> 

 newValidators.json example 
 {
    "addresses" : [
        "0xB4F1d19Db467f4990d2b654AeC174325B2ec3012",
        "0xe46C3785F1D0853481773646d7baD281E60BaDbE"
    ],
    "powers" : [100, 100]
 }

 signers.json
 {
    "addresses" : [
        "0x70592BAE00804966295150168467284711E8c2f0",
        "0xe46C3785F1D0853481773646d7baD281E60BaDbE",
        "0x884f0EC39Bb20E5f140A73cc4E1824E911B557Ef",
        "0x2a1CB8c11dA6774CDC7387fb748444Ab826336A7"
    ],
    "privateKeys" : [
        "0xdeadbeef...01",
        "0xdeadbeef...02",
        "0xdeadbeef...03",
        "0xdeadbeef...04"
    ],
    "powers" : [10, 10, 10, 10]
 }

*/
const fs = require('fs')

const ethers = require('ethers')
const Gateway = artifacts.require('Gateway')
const { createSigners, createSignsWithValidators } = require('../test/utils.js')

const ValidatorManagerContract = artifacts.require('ValidatorManagerContract')

module.exports = async function (callback) {
  try {
    const accounts = await web3.eth.getAccounts()
    const alice = accounts[0]

    if (!process.argv[6]) {
      throw new Error('Expected the Ethereum address of the Gateway')
    }
    const signerFile = process.env.SIGNER_FILE
    if (!signerFile) {
      throw new Error("No signer file found.")
    }

    const newValidatorFilePath = process.env.NEW_VALIDATOR_FILE
    if (!newValidatorFilePath) {
      throw new Error("No validator file found.")
    }

    const gatewayAddress = process.argv[6]
    const gateway = await Gateway.at(gatewayAddress)
    const vmcAddr = await gateway.vmc()
    const vmc = await ValidatorManagerContract.at(vmcAddr)
    const validators = await vmc.getValidators()
    console.log("validators", validators)

    let newValidators = await createNewValidators(newValidatorFilePath)
    let nonce = await vmc.nonce.call()
    let hashedData = await createRotateValidatorHash(vmc, nonce, newValidators.addresses, newValidators.powers)
    let signers = await createSigners(signerFile)
    let sigs = await createSignsWithValidators(signers, hashedData, signers.totalPower, 1, validators)

    await vmc.rotateValidators(newValidators.addresses, newValidators.powers, sigs.signers, sigs.v, sigs.r, sigs.s, { from: alice })

    callback()
  } catch (error) {
    callback(error)
  }
}

async function createRotateValidatorHash(vmc, nonce, newValidators, newPowers) {
  let validatorPowerHashed = ethers.utils.solidityKeccak256(['address[]', 'uint64[]'], [newValidators, newPowers])
  let hashedData = ethers.utils.solidityKeccak256(['address', 'uint256', 'bytes32'], [vmc.address, nonce.toString(), validatorPowerHashed])
  return hashedData
}

async function createNewValidators(validatorFilePath) {
  let validatorsFile = validatorFilePath

  if (!validatorsFile) {
    throw new Error("No validators file found.")
  }

  let validators = {
    addresses: [],
    powers: [],
  };

  const validatorsObj = JSON.parse(fs.readFileSync(validatorsFile, "utf8"))
  validators.addresses = validatorsObj.addresses
  validators.powers = validatorsObj.powers
  return validators
}
