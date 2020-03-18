#!/bin/bash

# To run this script locally:
# - Append `pwd` to your GOPATH.
# - Set LOOM_BIN env var to point at the loom binary you wish to run the tests with.
# - Set LOOM_VALIDATORS_TOOL env var to point at the validators tool binary if you wish to run the
#   tests on a local DAppChain cluster instead of a single node.
#
# To run the tests with an out-of-process Oracle:
# - Set LOOM_ORACLE env var to point at the oracle binary you wish to run the tests with.
# - Set LOOMCOIN_TGORACLE env var to point at the oracle binary you wish to run the tests with.
# - Disable the in-process Oracles in loom_test_config.yml
#
# To debug the DAppChain node (or in-process Oracle) while running a test:
#
# ./loom_e2e_tests.sh --init
# ... Start DAppChain in debugger ...
# ./loom_e2e_tests.sh --persist
#
# Similarly it's possible to debug the (out-of-process) Oracle by launching it manually via the
# debugger and then executing:
#
# ./loom_e2e_tests.sh


set -exo pipefail

# Loom build to use for tests when running on Jenkins, this build will be automatically downloaded.
# BUILD_NUMBER=build-908
BUILD_NUMBER=master

# These can be toggled via the options below, only useful when running the script locally.
DOWNLOAD_LOOM=false
INIT_DAPPCHAIN=false
REMOVE_LOOM_DIR=true
LAUNCH_ORACLE=false
LAUNCH_DAPPCHAIN=false
LAUNCH_GANACHE=false
SKIP_TESTS=false
WAIT_ON_EXIT=false
ETHEREUM_NETWORK="ganache"
DAPPCHAIN_NETWORK="local"
DEPLOY_TO_DAPPCHAIN=false
DEPLOY_TO_ETHEREUM=false
ORACLE_WAIT_TIME=10
TEST_TO_RUN="ALL"
MAP_CONTRACTS=false
DAPPCHAIN_NODE_COUNT=1

# Scripts options:
# -i / --init    - Reinitializes the DAppChain for a fresh test run.
# --launch-dappchain - Reinitializes and starts the DAppChain node, useful when you want to
#                      launch the DAppChain node manually via the debugger, or use the PlasmaChain
#                      Testnet.
# -p / --persist - Prevents the DAppChain working directory from being wiped on exit, to allow
#                  post test examination of the DAppChain logs etc.
while [[ "$#" > 0 ]]; do case $1 in
  -i|--init) INIT_DAPPCHAIN=true; shift;;
  -p|--persist) REMOVE_LOOM_DIR=false; shift;;
  --download-loom) DOWNLOAD_LOOM=true; shift;;
  --launch-dappchain) LAUNCH_DAPPCHAIN=true; shift;;
  --launch-ganache) LAUNCH_GANACHE=true; shift;;
  --launch-oracle) LAUNCH_ORACLE=true; shift;;
  --nodes) DAPPCHAIN_NODE_COUNT=$2; shift; shift;;
  --skip-tests) SKIP_TESTS=true; shift;;
  --wait-on-exit) WAIT_ON_EXIT=true; shift;;
  --ethereum-network) ETHEREUM_NETWORK="$2"; shift; shift;;
  --dappchain-network) DAPPCHAIN_NETWORK="$2"; shift; shift;;
  --deploy-dappchain-contracts) DEPLOY_TO_DAPPCHAIN=true; shift;;
  --deploy-ethereum-contracts) DEPLOY_TO_ETHEREUM=true; shift;;
  --map-contracts) MAP_CONTRACTS=true; shift;;
  --enable-hsm) ENABLE_HSM=true; shift;;
  --hsmkey-address) HSM_ADDRESS="$2"; shift; shift;;
  --run-test) TEST_TO_RUN="$2"; shift; shift;;
  *) echo "Unknown parameter: $1"; shift; shift;;
esac; done

if [[ "$ETHEREUM_NETWORK" != "ganache" ]]; then
    ORACLE_WAIT_TIME=120
fi

if [[ -z "$TEST_TO_RUN" ]]; then
    TEST_TO_RUN="ALL"
fi

echo "Reinitializing DAppChain? $INIT_DAPPCHAIN"
echo "Launching DAppChain? $LAUNCH_DAPPCHAIN"
echo "Removing LOOM_DIR on exit? $REMOVE_LOOM_DIR"

# Spins up a Ganache node & a DAppChain node
function start_chains {
    if [[ "$LAUNCH_GANACHE" == true ]]; then
        cd $REPO_ROOT/mainnet
        yarn run migrate:dev
        sleep 1
        ganache_pid=`cat ganache.pid`
        echo 'Launched ganache' $ganache_pid
    fi

    if [[ "$LAUNCH_DAPPCHAIN" == true ]]; then
        cd $LOOM_DIR
        if (( DAPPCHAIN_NODE_COUNT > 1 )); then
            $LOOM_VALIDATORS_TOOL run --conf cluster/runner.toml > cluster.log 2>&1 &
            loom_pid=$!
        else
            $LOOM_BIN run > loom.log 2>&1 &
            loom_pid=$!
        fi
        echo "Launched Loom - Log(loom.log) Pid(${loom_pid})"
    fi

    if [[ "$LAUNCH_GANACHE" == true ]] || [[ "$LAUNCH_DAPPCHAIN" == true ]]; then
        # Wait for Ganache & Loom to spin up
        sleep 10
    fi

    if [[ "$LAUNCH_ORACLE" == true ]]; then
        cd $LOOM_DIR
        $LOOM_ORACLE &
        oracle_pid=$!
        echo "Launched Transfer Gateway Oracle - Pid(${oracle_pid})"

        $LOOMCOIN_TGORACLE &
        loomcoin_oracle_pid=$!
        echo "Launched Transfer Gateway Loom Oracle - Pid(${loomcoin_oracle_pid})"
        sleep 5
    fi
}

# Stops the Ganache node & the DAppChain node
function stop_chains {
    if [[ "$LAUNCH_ORACLE" == true ]]; then
        echo "exiting oracle-pid(${oracle_pid})"
        echo "exiting loomcoin-oracle-pid(${loomcoin_oracle_pid})"
        kill -9 "${oracle_pid}" &> /dev/null
        kill -9 "${loomcoin_oracle_pid}" &> /dev/null
    fi
    
    if [[ "$LAUNCH_GANACHE" == true ]]; then
        echo "exiting ganache-pid(${ganache_pid})"
        kill -9 "${ganache_pid}" &> /dev/null
    fi
    
    if [[ "$LAUNCH_DAPPCHAIN" == true ]]; then
        echo "exiting loom-pid(${loom_pid})"
        kill -15 "${loom_pid}" &> /dev/null
    fi
}

# Reset all persisted DAppChain state
function init_dappchain {
    cd $LOOM_DIR

    cp $E2E_CONFIG_DIR/oracle_priv.key ./oracle_priv.key
    cp $E2E_CONFIG_DIR/oracle_eth_priv.key ./oracle_eth_priv.key

    GENESIS_JSON="$E2E_CONFIG_DIR/genesis.json"
    if [[ "$ENABLE_HSM" == true ]]; then
        cp $E2E_CONFIG_DIR/loom.hsm.yml ./loom.yml
        cp $E2E_CONFIG_DIR/oracle_eth_priv_hsm.key ./oracle_eth_priv_hsm.key
        cp $E2E_CONFIG_DIR/oracle_priv_hsm.key ./oracle_priv_hsm.key

        GENESIS_JSON="$E2E_CONFIG_DIR/genesis.hsm.json"

        export ENABLE_HSM="1"
        export HSM_ADDRESS="$HSM_ADDRESS"
    else
        export ENABLE_HSM=""
        cp $E2E_CONFIG_DIR/loom.yml ./loom.yml
    fi
    
    rm -rf app.db
    rm -rf chaindata

    if (( DAPPCHAIN_NODE_COUNT > 1 )); then
        # Disable the in-process TG Oracle since we only want one to be running.
        sed -i "s/OracleEnabled\s*:.*$/OracleEnabled: false/m" $LOOM_DIR/loom.yml

        $LOOM_VALIDATORS_TOOL new \
        -g $GENESIS_JSON \
        -c loom.yml \
        --base-dir `pwd` \
        --contract-dir "" \
        --name cluster \
        --loom-path $LOOM_BIN \
        --log-app-db

        # Override the loom.yaml used by the TG Oracle/tests to connect to the first node.
        NODE_RPC_ADDR=`cat cluster/0/node_rpc_addr`
        sed -i "s/DAppChainReadURI\s*:.*$/DAppChainReadURI: http:\/\/${NODE_RPC_ADDR}\/query/m;\
        s/DAppChainWriteURI\s*:.*$/DAppChainWriteURI: http:\/\/${NODE_RPC_ADDR}\/rpc/m;\
        s/DAppChainEventsURI\s*:.*$/DAppChainEventsURI: ws:\/\/${NODE_RPC_ADDR}\/queryws/m" $LOOM_DIR/loom.yml
    else
        $LOOM_BIN init -f
        cp $GENESIS_JSON ./genesis.json
    fi
    echo 'Loom DAppChain initialized in ' $LOOM_DIR
}

function cleanup {
    stop_chains
}

function download_dappchain {
    cd $REPO_ROOT
    if [[ "`uname`" == 'Darwin' ]]; then
        BUILD_PLATFORM=osx
    else
        BUILD_PLATFORM=linux
    fi
    
    rm -f ./loom; true
    wget https://private.delegatecall.com/loom/${BUILD_PLATFORM}/${BUILD_NUMBER}/loom
    chmod +x loom
    export LOOM_BIN=`pwd`/loom
    
    if (( DAPPCHAIN_NODE_COUNT > 1 )); then
        rm -f ./tgoracle; true
        rm -f ./loomcoin_tgoracle; true
        rm -f ./validators-tool; true

        wget https://private.delegatecall.com/loom/${BUILD_PLATFORM}/${BUILD_NUMBER}/validators-tool
        wget https://private.delegatecall.com/loom/${BUILD_PLATFORM}/${BUILD_NUMBER}/loomcoin_tgoracle
        wget https://private.delegatecall.com/loom/${BUILD_PLATFORM}/${BUILD_NUMBER}/tgoracle

        chmod +x tgoracle
        chmod +x loomcoin_tgoracle
        chmod +x validators-tool

        export LOOMCOIN_TGORACLE=`pwd`/loomcoin_tgoracle
        export LOOM_ORACLE=`pwd`/tgoracle
        export LOOM_VALIDATORS_TOOL=`pwd`/validators-tool
    fi
}

function deploy_test_contracts {
    ETHEREUM_CONTRACTS=""
    DAPPCHAIN_CONTRACTS=""

    if [[ "$DEPLOY_TO_ETHEREUM" == true ]]; then
        ETHEREUM_CONTRACTS="CryptoCards,GameToken,ERC721XCards"
    fi

    if [[ "$DEPLOY_TO_DAPPCHAIN" == true ]]; then
        DAPPCHAIN_CONTRACTS="SampleERC721Token,SampleERC20Token,SampleERC721XToken"
    fi

    if [[ "$DEPLOY_TO_ETHEREUM" == true ]] || [[ "$DEPLOY_TO_DAPPCHAIN" == true ]]; then
        cd $LOOM_DIR
        DAPPCHAIN_NETWORK=$DAPPCHAIN_NETWORK \
        ETHEREUM_NETWORK=$ETHEREUM_NETWORK \
        $REPO_ROOT/deployer deploy --loom-dir "$LOOM_DIR" \
                            --ethereum-contracts "$ETHEREUM_CONTRACTS" \
                            --dappchain-contracts "$DAPPCHAIN_CONTRACTS" \
                            --deployment-file "$E2E_CONFIG_DIR/contracts.yml"
    fi
}

function map_test_contracts {
    DAPPCHAIN_CONTRACTS="SampleERC721Token,SampleERC20Token,SampleERC721XToken"
    
    DAPPCHAIN_NETWORK=$DAPPCHAIN_NETWORK \
    ETHEREUM_NETWORK=$ETHEREUM_NETWORK \
    $REPO_ROOT/deployer map-contracts --timeout "$ORACLE_WAIT_TIME" \
                        --loom-dir "$LOOM_DIR" \
                        --dappchain-contracts "$DAPPCHAIN_CONTRACTS" \
                        --deployment-file "$E2E_CONFIG_DIR/contracts.yml"
}

# BUILD_TAG is usually only set by Jenkins, so when running locally just hardcode some value
if [[ -z "$BUILD_TAG" ]]; then
    BUILD_TAG=123
fi

# REPO_ROOT is set in jenkins.sh, if the script is executed directly just use cwd
if [[ -z "$REPO_ROOT" ]]; then
    REPO_ROOT=`pwd`
fi

LOOM_DIR=`pwd`/tmp/loom-$BUILD_TAG
E2E_CONFIG_DIR=$REPO_ROOT/e2e_config/${DAPPCHAIN_NETWORK}_${ETHEREUM_NETWORK}

if [[ "$INIT_DAPPCHAIN" == true ]]; then
    rm -rf $LOOM_DIR; true
fi

mkdir -p $LOOM_DIR

if [[ "$DOWNLOAD_LOOM" == true ]]; then
    download_dappchain
fi

echo "REPO_ROOT=(${REPO_ROOT})"
echo "GOPATH=(${GOPATH})"

if [[ "$INIT_DAPPCHAIN" == true ]]; then
    init_dappchain
else
    cp $E2E_CONFIG_DIR/loom.yml $LOOM_DIR/loom.yml
fi

trap cleanup EXIT

start_chains
deploy_test_contracts

if [[ "$MAP_CONTRACTS" == true ]]; then
    map_test_contracts
fi

if [[ "$SKIP_TESTS" == false ]]; then
    export GOPATH=$GOPATH:$REPO_ROOT
    cd $REPO_ROOT/src/gateway
    if [[ "$ETHEREUM_NETWORK" == "ganache" ]] && [[ "$TEST_TO_RUN" == "ALL" ]]; then
        LOOM_DIR=$LOOM_DIR \
        DAPPCHAIN_NETWORK=$DAPPCHAIN_NETWORK \
        ETHEREUM_NETWORK=$ETHEREUM_NETWORK \
        ORACLE_WAIT_TIME=$ORACLE_WAIT_TIME \
        go test gateway -tags "evm" -run TestTransferGatewayTestSuite
    else
        # each test takes about 6 mins to complete on Rinkeby, so run them individually to get
        # quicker feedback when something fails
        if [[ "$TEST_TO_RUN" == "ALL" ]] || [[ "$TEST_TO_RUN" == "ERC721DepositAndWithdraw" ]]; then
            LOOM_DIR=$LOOM_DIR \
            DAPPCHAIN_NETWORK=$DAPPCHAIN_NETWORK \
            ETHEREUM_NETWORK=$ETHEREUM_NETWORK \
            ORACLE_WAIT_TIME=$ORACLE_WAIT_TIME \
            go test gateway -tags "evm" -run TestTransferGatewayTestSuite -testify.m ^TestERC721DepositAndWithdraw$
        fi
        if [[ "$TEST_TO_RUN" == "ALL" ]] || [[ "$TEST_TO_RUN" == "ERC721DepositTransferWithdraw" ]]; then
            LOOM_DIR=$LOOM_DIR \
            DAPPCHAIN_NETWORK=$DAPPCHAIN_NETWORK \
            ETHEREUM_NETWORK=$ETHEREUM_NETWORK \
            ORACLE_WAIT_TIME=$ORACLE_WAIT_TIME \
            go test gateway -tags "evm" -run TestTransferGatewayTestSuite -testify.m ^TestERC721DepositTransferWithdraw$
        fi
        if [[ "$TEST_TO_RUN" == "ALL" ]] || [[ "$TEST_TO_RUN" == "ERC721XDepositTransferWithdraw" ]]; then
            LOOM_DIR=$LOOM_DIR \
            DAPPCHAIN_NETWORK=$DAPPCHAIN_NETWORK \
            ETHEREUM_NETWORK=$ETHEREUM_NETWORK \
            ORACLE_WAIT_TIME=$ORACLE_WAIT_TIME \
            go test gateway -tags "evm" -run TestTransferGatewayTestSuite -testify.m ^TestERC721XDepositTransferWithdraw$
        fi
        if [[ "$TEST_TO_RUN" == "ALL" ]] || [[ "$TEST_TO_RUN" == "ERC20DepositAndWithdraw" ]]; then
            LOOM_DIR=$LOOM_DIR \
            DAPPCHAIN_NETWORK=$DAPPCHAIN_NETWORK \
            ETHEREUM_NETWORK=$ETHEREUM_NETWORK \
            ORACLE_WAIT_TIME=$ORACLE_WAIT_TIME \
            go test gateway -tags "evm" -run TestTransferGatewayTestSuite -testify.m ^TestERC20DepositAndWithdraw$
        fi
        if [[ "$TEST_TO_RUN" == "ALL" ]] || [[ "$TEST_TO_RUN" == "ETHDepositAndWithdraw" ]]; then
            LOOM_DIR=$LOOM_DIR \
            DAPPCHAIN_NETWORK=$DAPPCHAIN_NETWORK \
            ETHEREUM_NETWORK=$ETHEREUM_NETWORK \
            ORACLE_WAIT_TIME=$ORACLE_WAIT_TIME \
            go test gateway -tags "evm" -run TestTransferGatewayTestSuite -testify.m ^TestETHDepositAndWithdraw$
        fi
        if [[ "$TEST_TO_RUN" == "ALL" ]] || [[ "$TEST_TO_RUN" == "ETHDepositAndWithdrawWithEVM" ]]; then
            # This test may not work on anything other than Ganache yet...
            LOOM_DIR=$LOOM_DIR \
            DAPPCHAIN_NETWORK=$DAPPCHAIN_NETWORK \
            ETHEREUM_NETWORK=$ETHEREUM_NETWORK \
            ORACLE_WAIT_TIME=$ORACLE_WAIT_TIME \
            go test gateway -tags "evm" -run TestTransferGatewayTestSuite -testify.m ^TestETHDepositAndWithdrawWithEVM$
        fi
        if [[ "$TEST_TO_RUN" == "ALL" ]] || [[ "$TEST_TO_RUN" == "LoomDepositAndWithdraw" ]]; then
            LOOM_DIR=$LOOM_DIR \
            DAPPCHAIN_NETWORK=$DAPPCHAIN_NETWORK \
            ETHEREUM_NETWORK=$ETHEREUM_NETWORK \
            ORACLE_WAIT_TIME=$ORACLE_WAIT_TIME \
            go test gateway -tags "evm" -run TestTransferGatewayTestSuite -testify.m ^TestLoomDepositAndWithdraw$
        fi
    fi
fi

if [[ "$WAIT_ON_EXIT" == true ]]; then
    read -n 1 -s -r -p "Press any key to exit"
fi

if [[ $LOOM_DIR ]] && [[ "$REMOVE_LOOM_DIR" == true ]]; then
    rm -rf $LOOM_DIR
fi