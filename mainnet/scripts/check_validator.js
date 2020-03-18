/**
 * Usage Examples
 * 
 * truffle exec scripts/check_validator.js <gateway-eth-addr> <validator-addr>
 * 
 * To if an address belongs to a TG validator:
 * ./node_modules/.bin/truffle exec scripts/check_validator.js \
 *   0xf5cAD0DB6415a71a5BC67403c87B56b629b4DdaA 0x7292694902bcaf4e1620629e7198cdcb3f572a24 \
 *   --network rpc
 */
const Gateway = artifacts.require('Gateway')

module.exports = async function(callback) {
  try {
    if (!process.argv[4]) {
      throw new Error('Expected the Ethereum address of the Gateway')
    }
    if (!process.argv[5]) {
      throw new Error('Expected an Ethereum address')
    }
    const gatewayAddr = process.argv[4]
    const validatorAddr = process.argv[5].toLowerCase()
    const instance = await Gateway.at(gatewayAddr)

    const isValidator = await instance.checkValidator(validatorAddr)
    console.log(`${validatorAddr} is${isValidator ? '' : ' not'} a validator of ${gatewayAddr}`)
  } catch (err) {
    callback(err)
  }
  callback()
}
