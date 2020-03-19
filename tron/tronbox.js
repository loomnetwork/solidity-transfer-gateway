const port = process.env.HOST_PORT || 9090

module.exports = {
  networks: {
    development: {
      // For trontools/quickstart docker image
      privateKey: '11b5b4c7ff9e69e89e66ef614d8d63a81f2799c3b2bad938a6ae3b0a10a12772',
      userFeePercentage: 30,
      // feeLimit: 1e8,
      fullHost: 'http://127.0.0.1:' + port,
      network_id: port
    },
    development2: {
      // For trontools/quickstart docker image
      privateKey: 'da146374a75310b9666e834ee4ad0866d6f4035967bfc76217c5a495fff9f0d0',
      userFeePercentage: 30,
      feeLimit: 1e8,
      fullHost: "http://127.0.0.1:9091",
      network_id: "9091"
    },
    shasta: {
      privateKey: process.env.PRIVATE_KEY_SHASTA,
      userFeePercentage: 50,
      feeLimit: 1e8,
      fullHost: "https://api.shasta.trongrid.io",
      network_id: "2"
    },
    mainnet: {
      privateKey: process.env.PRIVATE_KEY_MAINNET,
      userFeePercentage: 100,
      feeLimit: 1e8,
      fullHost: "https://api.trongrid.io",
      network_id: "1"
    }
  }
}