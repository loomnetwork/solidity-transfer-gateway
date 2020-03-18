const ethutil = require('ethereumjs-util')

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

async function signHash(from, hash) {
  let sig = (await web3.eth.sign(hash, from)).slice(2)
  let r = ethutil.toBuffer('0x' + sig.substring(0, 64))
  let s = ethutil.toBuffer('0x' + sig.substring(64, 128))
  let v = ethutil.toBuffer(parseInt(sig.substring(128, 130), 16) + 27)
  let mode = ethutil.toBuffer(1) // mode = geth
  let signature = '0x' + Buffer.concat([mode, r, s, v]).toString('hex')
  return signature
}

module.exports = {
  signHash,
  assertEventVar,
  Promisify,
}
