#!/bin/bash

# If the ETHEREUM_NETWORK env var isn't set this script spin up a DAppChain node & Ganache to run
# the e2e tests on. Otherwise it will run the tests against the Ethereum and DAppChain networks
# specified by the ETHEREUM_NETWORK and DAPPCHAIN_NETWORK env vars.
#
# Currently supported values for ETHEREUM_NETWORK (case-sensitive, if set):
# - rinkeby
#
# Currently supported values for DAPPCHAIN_NETWORK (case-sensitive, if set):
# - pc_testnet (PlasmaChain Testnet)
#
# The following env vars may also be set to tweak the script flow:
# TEST_TO_RUN - Can be used to specify a single test to run.

set -exo pipefail

REPO_ROOT=`pwd`
pkill -f loom || true
pkill -f loom-gateway || true
pkill -f binance_tgoracle || true

#load specific number of LOOM_BIN
rm -rf loom loom-gateway binance_tgoracle tgoracle loomcoin_tgoracle
export BUILD_ID=build-1225



# Check available platforms
PLATFORM='unknown'
unamestr=`uname`
if [[ "$unamestr" == 'Linux' ]]; then
  PLATFORM='linux'
elif [[ "$unamestr" == 'Darwin' ]]; then
  PLATFORM='osx'
else
  echo "Platform not supported on this script yet"
  exit -1
fi



wget https://private.delegatecall.com/loom/$PLATFORM/$BUILD_ID/loom-gateway
chmod +x loom-gateway
mv loom-gateway loom
export LOOM_BIN=`pwd`/loom

wget https://private.delegatecall.com/loom/$PLATFORM/$BUILD_ID/loomcoin_tgoracle
chmod +x loomcoin_tgoracle
wget https://private.delegatecall.com/loom/$PLATFORM/$BUILD_ID/tgoracle
chmod +x tgoracle

if [[ -z "$ETHEREUM_NETWORK" ]]; then
    cd $REPO_ROOT/mainnet
    yarn install
    yarn lint
    yarn compile
    yarn test

    cd $REPO_ROOT/dappchain
    yarn install
    yarn compile
fi

cd $REPO_ROOT
export GOPATH=/tmp/gopath-$BUILD_TAG
mkdir -p $GOPATH/bin
export PATH=$PATH:$GOPATH/bin

make clean
make deps
make vendor-deps
make deployer

if [[ -z "$ETHEREUM_NETWORK" ]]; then
    pkill -f ganache || true
    REPO_ROOT=`pwd` \
    bash loom_e2e_tests.sh --download-loom --nodes 4 --skip-tests

    # run the tests on a single node
    REPO_ROOT=`pwd` \
    LOOM_BIN=$REPO_ROOT/loom \
    LOOMCOIN_TGORACLE=$REPO_ROOT/loomcoin_tgoracle \
    LOOM_ORACLE=$REPO_ROOT/tgoracle \
    bash loom_e2e_tests.sh --init \
                           --launch-dappchain --launch-ganache  --launch-oracle \
                           --deploy-dappchain-contracts --deploy-ethereum-contracts \
                           --map-contracts

    # # run the tests on a single node with yubihsm (disabled until we setup new remote signer)
    # pkill -f ganache || true
    # REPO_ROOT=`pwd` \
    # LOOM_BIN=$REPO_ROOT/loom \
    # bash loom_e2e_tests.sh --init \
    #                       --launch-dappchain --launch-ganache \
    #                       --deploy-dappchain-contracts --deploy-ethereum-contracts \
    #                       --map-contracts \
    #                       --run-test ERC721DepositAndWithdraw \
    #                       --enable-hsm --hsmkey-address 0x2669Ff29f3D3e78DAFd2dB842Cb9d0dDb96D90f2
    
    # run the tests again on a 4-node cluster...
    pkill -f ganache || true
    REPO_ROOT=`pwd` \
    LOOM_BIN=$REPO_ROOT/loom \
    LOOMCOIN_TGORACLE=$REPO_ROOT/loomcoin_tgoracle \
    LOOM_ORACLE=$REPO_ROOT/tgoracle \
    LOOM_VALIDATORS_TOOL=$REPO_ROOT/validators-tool \
    bash loom_e2e_tests.sh --init \
                           --launch-dappchain --launch-ganache --launch-oracle \
                           --deploy-dappchain-contracts --deploy-ethereum-contracts \
                           --map-contracts \
                           --nodes 4

    # # run the tests again on a 4-node cluster with yubihsm
    # pkill -f ganache || true
    # REPO_ROOT=`pwd` \
    # LOOM_BIN=$REPO_ROOT/loom \
    # LOOMCOIN_TGORACLE=$REPO_ROOT/loomcoin_tgoracle \
    # LOOM_ORACLE=$REPO_ROOT/tgoracle \
    # LOOM_VALIDATORS_TOOL=$REPO_ROOT/validators-tool \
    # bash loom_e2e_tests.sh --init \
    #                       --launch-dappchain --launch-ganache --launch-oracle \
    #                       --deploy-dappchain-contracts --deploy-ethereum-contracts \
    #                       --map-contracts \
    #                       --nodes 4 \
    #                       --run-test ERC721DepositAndWithdraw \
    #                       --enable-hsm --hsmkey-address 0x2669Ff29f3D3e78DAFd2dB842Cb9d0dDb96D90f2
else
    REPO_ROOT=`pwd` \
    bash loom_e2e_tests.sh --dappchain-network "$DAPPCHAIN_NETWORK" \
                           --ethereum-network "$ETHEREUM_NETWORK" \
                           --run-test "$TEST_TO_RUN"

    REPO_ROOT=`pwd` \
    bash loom_e2e_tests.sh --dappchain-network "$DAPPCHAIN_NETWORK" \
                          --ethereum-network "$ETHEREUM_NETWORK" \
                          --run-test "$TEST_TO_RUN" \
                          --enable-hsm --hsmkey-address 0x2669Ff29f3D3e78DAFd2dB842Cb9d0dDb96D90f2
fi

# cd tron

# make deps

# cd ../

## Tron gateway get wiped, and for some reasons we can't deploy a new gateway at the moment. So, we disable the end to end test script
# ## Run Tron test on Shasta
# REPO_ROOT=`pwd` \
# bash loom_e2e_tests.sh --init \
#     --gateway-type tron-gateway \
#     --tron-network shasta \
#     --launch-dappchain \
#     --deploy-dappchain-contracts \
#     --map-contracts

# Run Binance Gateway e2e test
cd $REPO_ROOT
export ORACLE_BUILD_NUMBER=build-27
wget https://private.delegatecall.com/binance_tgoracle/linux/$ORACLE_BUILD_NUMBER/binance_tgoracle
chmod +x binance_tgoracle

LOOM_ORACLE=$REPO_ROOT/binance_tgoracle \
REPO_ROOT=`pwd` \
bash loom_e2e_tests.sh --init \
    --gateway-type binance-gateway \
    --binance-network bnbtestnet \
    --launch-dappchain \
    --launch-oracle \
    --map-contracts \
    --deploy-dappchain-contracts \
    --run-test ALL \
    --reset-latest-block-num \
    --set-transfer-fee
