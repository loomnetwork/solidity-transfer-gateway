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

if [[ -z "$ETHEREUM_NETWORK" ]]; then
    cd $REPO_ROOT/mainnet
    yarn install
    yarn lint
    yarn test

    cd $REPO_ROOT/dappchain
    yarn install
    yarn compile
fi

cd $REPO_ROOT
export GOPATH=/tmp/gopath-$BUILD_TAG
make clean
make deps
make deployer

if [[ -z "$ETHEREUM_NETWORK" ]]; then
    pkill -f ganache || true
    REPO_ROOT=`pwd` \
    bash loom_e2e_tests.sh --download-loom --nodes 4 --skip-tests

    # run the tests on a single node
    REPO_ROOT=`pwd` \
    LOOM_BIN=$REPO_ROOT/loom \
    bash loom_e2e_tests.sh --init \
                           --launch-dappchain --launch-ganache \
                           --deploy-dappchain-contracts --deploy-ethereum-contracts \
                           --map-contracts

    # run the tests on a single node with yubihsm (disabled until we setup new remote signer)
    #pkill -f ganache || true
    #REPO_ROOT=`pwd` \
    #LOOM_BIN=$REPO_ROOT/loom \
    #bash loom_e2e_tests.sh --init \
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

    # run the tests again on a 4-node cluster with yubihsm
    #pkill -f ganache || true
    #REPO_ROOT=`pwd` \
    #LOOM_BIN=$REPO_ROOT/loom \
    #LOOMCOIN_TGORACLE=$REPO_ROOT/loomcoin_tgoracle \
    #LOOM_ORACLE=$REPO_ROOT/tgoracle \
    #LOOM_VALIDATORS_TOOL=$REPO_ROOT/validators-tool \
    #bash loom_e2e_tests.sh --init \
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

    #REPO_ROOT=`pwd` \
    #bash loom_e2e_tests.sh --dappchain-network "$DAPPCHAIN_NETWORK" \
    #                       --ethereum-network "$ETHEREUM_NETWORK" \
    #                       --run-test "$TEST_TO_RUN" \
    #                       --enable-hsm --hsmkey-address 0x2669Ff29f3D3e78DAFd2dB842Cb9d0dDb96D90f2
fi
