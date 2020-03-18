/**
 * Usage Examples
 * 
 * To enable allowAnyToken on the Ethereum Gateway:
 * ACCOUNT=0x7292694902bcAF4E1620629E7198cDcb3f572A24 \
 * GATEWAY_ETH_ADDR=0xf5cAD0DB6415a71a5BC67403c87B56b629b4DdaA \
 * ./node_modules/.bin/truffle exec scripts/toggle_allow_any_token.js true --network rpc
 * 
 * To disable allowAnyToken on the Ethereum Gateway:
 * ACCOUNT=0x7292694902bcAF4E1620629E7198cDcb3f572A24 \
 * GATEWAY_ETH_ADDR=0xf5cAD0DB6415a71a5BC67403c87B56b629b4DdaA \
 * ./node_modules/.bin/truffle exec scripts/toggle_allow_any_token.js false --network rpc
 */
const Gateway = artifacts.require('Gateway')

module.exports = async function(callback) {
  const account = process.env.ACCOUNT
  const gatewayAddr = process.env.GATEWAY_ETH_ADDR

  try {
    let allowAnyToken
    if (process.argv[4] === 'true') {
      allowAnyToken = true
    } else if (process.argv[4] === 'false') {
      allowAnyToken = false
    } else {
      throw new Error('Expected true or false')
    }
    const instance = await Gateway.at(gatewayAddr)

    const enabled = await instance.allowAnyToken()
    console.log(`allowAnyToken is currently ${enabled}`)

    if (enabled !== allowAnyToken) {
      console.log(`setting allowAnyToken to ${allowAnyToken}`)
      await instance.toggleAllowAnyToken(allowAnyToken, { from: account })
    }
  } catch (err) {
    callback(err)
  }
  callback()
}
