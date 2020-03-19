# Transfer Gateway V2

A transfer gateway which supports TRX, TRC20 tokens.

# Setup locally

```bash
cd tron
make deps
yarn install
yarn compile

cd ../dappchain
yarn install
yarn compile

cd ../
make deps
make deployer
```

In the case that you edit the smart contracts, make sure to always regenerate the go bindings for them, via `make abigen` :)

# Setup Tronbox

see more details on `tron/README.md`

# Test with Shasta

Use the flag `--gateway-type tron-gateway` when running test to specifiy gateway type.

It might be cumbersome to to setup Tron's event server locally, testing with Shasta testnet would be comparatively faster. In this case, we just need another flag `--tron-network shasta` to run with the test command.

For example:

Run test with a new dappchain node with already deployed Tron transfer gateway:

```bash
export $LOOM_BIN=`your path to loom binary`
./loom_e2e_tests.sh --init --gateway-type tron-gateway \
        --launch-dappchain \
        --deploy-dappchain-contracts \
        --tron-network shasta \
        --map-contracts
```

In case we want to also deploy a new transfer gateway contract on Tron shasta test net, run:

```bash
export $LOOM_BIN=`your path to loom binary`
./loom_e2e_tests.sh --init --gateway-type tron-gateway \
        --launch-dappchain \
        --deploy-tron-contracts \
        --tron-network shasta \
        --deploy-dappchain-contracts \
        --map-contracts
```