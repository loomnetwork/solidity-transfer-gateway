/**
 * Usage Examples
 * 
 * truffle exec scripts/transfer_gateway_fund.js <gateway-eth-addr> <input-path> <output-path>
 * 
 * Show token holding on mainnet gateway:
 * ./node_modules/.bin/truffle exec scripts/transfer_gateway_funds.js \
 *   0xB2D92ea5592b1F16c7715451aE3A91195a4D5A05 \
 *  ./asia1.json \
 *  ./asia1.json \
 *   --network rinkeby
 * 
 * JSON input example : 
 * {
  "data": [
    {
      "Local": "0x0246A8a81C6BbD3a02b084E1AB89D04E7cFaB6Bb",
      "Foreign": "0x5D48922cC1069323A83bF0b3D170320eFd2f9Ec4"
    }
  ]
  }

 */
const fs = require('fs')
const SampleERC721MintableToken = artifacts.require("SampleERC721MintableToken")
const SampleERC20MintableToken = artifacts.require("SampleERC20MintableToken")
const SampleERC721x = artifacts.require("ERC721XTokenNFT")

const TOKEN_KIND = Object.freeze({
  ETH: "ETH",
  ERC20: "ERC20",
  ERC721: "ERC721",
  ERC721x: "ERC721x",
  UNDEFINED: "Undefined"
})

const ERR_MSG = Object.freeze({
  INVALID_RETURN_VALUE: "Returned values aren't valid, did it run Out of Gas"
})


module.exports = async function (callback) {
  try {
    if (!process.argv[4]) {
      throw new Error('Expected the Ethereum address of the Gateway')
    }

    if (!process.argv[5]) {
      throw new Error('Expected the mapping input path')
    }

    if (!process.argv[6]) {
      throw new Error('Expected the output file name')
    }

    const ContractMapping = require(process.argv[5])
    const gatewayAddr = process.argv[4]
    const filename = process.argv[6]

    let result = await ContractMapping.data.map(async (token) => {
      let tokenAddr = token.Foreign.toLowerCase()
      try {
        // Try if token is erc721x compatible
        let token721xInstance = await SampleERC721x.at(tokenAddr)
        let ownedToken = await token721xInstance.tokensOwned(gatewayAddr)
        return {
          token: tokenAddr,
          balance: ownedToken.indexes.length,
          tokenKind: TOKEN_KIND.ERC721x,
          tokenIndex: ownedToken.indexes,
          balances: ownedToken.balances
        }
      } catch (error) {
        if (!error.message.includes(ERR_MSG.INVALID_RETURN_VALUE)) {
          console.log(error.message);
        }
        
        try {
          let token721Instance = await SampleERC721MintableToken.at(tokenAddr)
          let balance = await token721Instance.balanceOf(gatewayAddr)
          let tokenIndex = await getTokenByIndex(gatewayAddr, tokenAddr)
          return {
            token: tokenAddr,
            balance: balance,
            tokenKind: TOKEN_KIND.ERC721,
            tokenIndex: tokenIndex
          }
        } catch (error) {
          if (error.message.includes(ERR_MSG.INVALID_RETURN_VALUE)) {
            let token20instance = await SampleERC20MintableToken.at(tokenAddr)
            let balance = await token20instance.balanceOf(gatewayAddr)
            return {
              token: tokenAddr,
              tokenKind: TOKEN_KIND.ERC20,
              balance: balance.toString()
            }
          } else { 
            return {
              token: tokenAddr,
              tokenKind: TOKEN_KIND.UNDEFINED
            }
          }
        }
      }
    })

    const resolved = await Promise.all(result)
    fs.writeFileSync(filename, JSON.stringify({ result: resolved }))
    console.log("Completed");
  } catch (err) {
    callback(err)
    throw err
  }
  callback()
}

async function getTokenByIndex(gatewayAddr, tokenAddr) {
  let token721Instance = await SampleERC721MintableToken.at(tokenAddr)
  let ownedToken = []
  try {
    let balance = await token721Instance.balanceOf(gatewayAddr)
    console.log(tokenAddr, "balance", balance.toString());

    for (let index = 0; index < balance; index++) {
      let tokenId = await token721Instance.tokenOfOwnerByIndex(gatewayAddr, index)
      ownedToken.push(tokenId.toString())
      console.log(tokenAddr, "index=", index, "tokenId=", tokenId.toString());
    }

    console.log(tokenAddr, "ownedToken", ownedToken.toString());
    return ownedToken
  } catch (error) {
    throw error
  }
}