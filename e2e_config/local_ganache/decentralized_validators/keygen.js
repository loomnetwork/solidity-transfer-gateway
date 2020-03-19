const fs = require('fs')
const e = require('ethers')

data = []
for (let i = 0 ; i < 6; i++) {
    w = e.Wallet.createRandom()
    data.push({
        "key": w.privateKey,
        "address": w.address
    });
    fs.writeFileSync('validator_' + i, w.privateKey)
}

console.log(JSON.stringify(data, null, 2))
