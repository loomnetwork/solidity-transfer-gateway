
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

const assertEventVar = (transaction, eventName, eventVar, equalVar) => {
  const event = transaction.logs.find(log => log.event === eventName);
  assert.equal(event.args[eventVar], equalVar, `Event ${eventVar} didn't happen`);
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




module.exports = {
  expectThrow,
  assertEventVar,
  Promisify,
}
