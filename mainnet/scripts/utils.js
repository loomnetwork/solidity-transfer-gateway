// Utility functions
const BN = require('bn.js')
const ethers = require('ethers')

const ORACLE_SIG_SIZE_WITH_MODE = 134 // '0x'.length + (65 + 1) * 2
const ORACLE_SIG_SIZE = 132 // '0x'.length + 65 * 2
const MessagePrefix = Object.freeze({
  ETHPrefix: '\x0eWithdraw ETH:\n',
  ERC20Prefix: '\x10Withdraw ERC20:\n',
  ERC721Prefix: '\x11Withdraw ERC721:\n',
  ERC721XPrefix: '\x12Withdraw ERC721X:\n'
})
module.exports.MessagePrefix = MessagePrefix

// sciNot returns a big number whose value is m x 10^n
module.exports.sciNot = (m, n) => {
  let b = new BN('10', 10)
  b = b.pow(new BN(n))
  b = b.mul(new BN(m))
  return b
}

// delay delays in milliseconds
module.exports.delay = ms => new Promise(resolve => setTimeout(resolve, ms))

// parseSigs parses signature and return an object of vs, rs, ss
module.exports.parseSigs = (sig, hash, validators) => {
  let vs = []
  let rs = []
  let ss = []
  let valIndexes = []

  let sigs
  if (sig.length === ORACLE_SIG_SIZE_WITH_MODE) {
    // using old oracle but new mainnet contract requires removing the 'mode' bit from the signature
    sigs = ['0x' + sig.slice(4)]
  } else if (sig.length === ORACLE_SIG_SIZE) {
    // if the oracle signs without a mode
    sigs = [sig]
  } else {
    // else split sig string into 65 byte array of sigs
    sigs = sig
      // .slice(2)
      .match(/.{1,130}/g)
      .map(s => '0x' + s)
  }

  console.log(sigs)

  // Split signature in v,r,s arrays
  // Store the ordering of the validators' signatures in `valIndexes`
  for (let i in sigs) {
    const _hash = ethers.utils.arrayify(ethers.utils.hashMessage(ethers.utils.arrayify(hash)))
    const recAddress = ethers.utils.recoverAddress(_hash, sigs[i]).toLowerCase()
    const ind = validators.indexOf(recAddress)
    if (ind == -1) {
      // skip if invalid signature
      continue
    }

    valIndexes.push(validators.indexOf(recAddress))

    const s = ethers.utils.splitSignature(sigs[i])
    vs.push(s.v)
    rs.push(s.r)
    ss.push(s.s)
  }

  vs = mapOrder(vs, valIndexes)
  rs = mapOrder(rs, valIndexes)
  ss = mapOrder(ss, valIndexes)
  valIndexes.sort()
  return { vs, rs, ss, valIndexes }
}


/**
 * @param gatewayAddress Mainnet gateway address
 * @param nonce nonce value from gateway
 * @param prefix
 * @param tokenAddress
 * @param userEthAddress
 * ETH, ERC20
 * @param amount amount to withdraw
 * ERC721, ERC721x
 * @param uid
 */
module.exports.createWithdrawalHash = (params) => {
  let amountHashed

  if (params.prefix === MessagePrefix.ETHPrefix) {
    amountHashed = ethers.utils.solidityKeccak256(
      ['uint256'],
      [params.amount.toString()]
    )
  } else if (params.prefix === MessagePrefix.ERC20Prefix) {
    amountHashed = ethers.utils.solidityKeccak256(
      ['uint256', 'address'],
      [params.amount.toString(), params.tokenAddress]
    )
  } else if (params.prefix === MessagePrefix.ERC721Prefix) {
    amountHashed = ethers.utils.solidityKeccak256(
      ['uint256', 'address'],
      [params.uid.toString(), params.tokenAddress]
    )
  } else if (params.prefix === MessagePrefix.ERC721XPrefix) {
    amountHashed = ethers.utils.solidityKeccak256(
      ['uint256', 'uint256', 'address'],
      [params.uid.toString(), params.amount.toString(), params.tokenAddress]
    )
  }

  const msg = ethers.utils.solidityKeccak256(
    ['string', 'address', 'uint256', 'address', 'bytes32'],
    [params.prefix, params.userEthAddress, params.nonce.toString(), params.gatewayAddress, amountHashed]
  )

  return msg
}

function mapOrder(array, order) {
  if (array.length === 1) {
    return array
  }

  const sortedArray = []

  for (const i in order) {
    const ind = order.indexOf(parseInt(i, 10))
    sortedArray.push(array[ind])
  }

  return sortedArray
}
