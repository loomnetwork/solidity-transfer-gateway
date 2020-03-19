const ethutil = require('ethereumjs-util')
const assertEventVar = (response, eventName, eventVar, equalVar) => {

  assert.notEqual(response.length, 0, `Your transaction was reverted.`)
  const event = response.find(log => log.name === eventName);
  if (tronWeb.isAddress(equalVar)) {
    // in case assert address: event emitted address in '0x' fomat so need to convert first
    let addrHex = tronWeb.address.toHex(equalVar)
    assert.equal(event.result[eventVar], '0x' + addrHex.slice(2), `Event ${event.result[eventVar]} didn't happen..`);
  } else {
    assert.equal(event.result[eventVar], equalVar, `Event ${event.result[eventVar]} didn't happen...`);
  }
};

const assertRevert = (response) => {
  assert.equal(response.length, 0, `Assert fail. Your transaction should revert.`)
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

async function signHash(hash, privateKeyHex) {
  let message = ethutil.toBuffer(hash)
  let privateKey = new Buffer.from(privateKeyHex, "hex") // this is the gateway address
  let sig = ethutil.ecsign(message, privateKey);
  return sig
}

module.exports = {
  signHash,
  assertEventVar,
  assertRevert,
  Promisify,
}
