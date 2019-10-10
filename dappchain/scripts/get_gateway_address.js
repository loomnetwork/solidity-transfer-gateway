const { Client, NonceTxMiddleware, SignedTxMiddleware,
    CryptoUtils, Contracts } = require("loom-js");
let yaml = require('js-yaml');
let fs = require('fs');

async function loadDappClient() {
    try {
        const loomYamlPath = process.env.CFG
        const cfg = yaml.safeLoad(fs.readFileSync(loomYamlPath, 'utf8'))


        const privateKeyStr = fs.readFileSync('../test.key', 'utf-8')
        const privateKey = CryptoUtils.B64ToUint8Array(privateKeyStr)
        const publicKey = CryptoUtils.publicKeyFromPrivateKey(privateKey)
        let writeURI, readURI, gatewayType

        gatewayType = process.argv[2]
        if (!process.argv[2]) {
            gatewayType = 'eth'
        }

        if (gatewayType === 'eth') {
            writeURI = cfg.TransferGateway.DAppChainWriteURI
            readURI = cfg.TransferGateway.DAppChainReadURI
        } else if (gatewayType === 'binance') {
            writeURI = cfg.BinanceTransferGateway.DAppChainWriteURI
            readURI = cfg.BinanceTransferGateway.DAppChainReadURI
        } else {
            throw new Error("gatewayType not found")
        }
        
        writeURI = writeURI.replace("http", "ws").replace("rpc", "websocket")
        readURI = readURI.replace("http", "ws").replace("query", "queryws")

        loomClient = new Client(cfg.ChainID,
            writeURI,
            readURI
        )

        loomClient.txMiddleware = [
            new NonceTxMiddleware(publicKey, loomClient),
            new SignedTxMiddleware(privateKey)
        ]

        loomClient.on('error', msg => {
            console.error('PlasmaChain connection error', msg)
        })

        let gatewayContractAddr
        if (gatewayType === 'eth') {
            gatewayContractAddr = await loomClient.getContractAddressAsync('gateway')
        } else if (gatewayType === 'binance') {
            gatewayContractAddr = await loomClient.getContractAddressAsync('binance-gateway')
        } else {
            throw new Error("gatewayType not found")
        }

        console.log(gatewayContractAddr.local.toString())

    } catch (err) {
        console.log(err)
    } finally {
        loomClient.disconnect()
    }
}

return loadDappClient()
