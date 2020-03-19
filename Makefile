.PHONY: all clean deps abigen deployer

PLUGIN_DIR = $(GOPATH)/src/github.com/loomnetwork/go-loom
GOLANG_PROTOBUF_DIR = $(GOPATH)/src/github.com/golang/protobuf
GENPROTO_DIR = $(GOPATH)/src/google.golang.org/genproto
GOGO_PROTOBUF_DIR = $(GOPATH)/src/github.com/gogo/protobuf
GRPC_DIR = $(GOPATH)/src/google.golang.org/grpc
GO_ETHEREUM_DIR = $(GOPATH)/src/github.com/ethereum/go-ethereum
SSHA3_DIR = $(GOPATH)/src/github.com/miguelmota/go-solidity-sha3
HASHICORP_DIR = $(GOPATH)/src/github.com/hashicorp/go-plugin

# use a modified ethereum git rev
ETHEREUM_GIT_REV = 1fb6138d017a4309105d91f187c126cf979c93f9
# use go-plugin we get 'timeout waiting for connection info' error
HASHICORP_GIT_REV = f4c3476bd38585f9ec669d10ed1686abd52b9961
# pin protobuf related
GENPROTO_GIT_REV = b515fa19cec88c32f305a962f34ae60068947aea

deployer:
	GOPATH=$(GOPATH):`pwd` \
	go build -tags "evm" -o deployer src/deployer/main.go

abigen:
	go build github.com/ethereum/go-ethereum/cmd/abigen
	# Need to run npm compile in the ./mainnet directory for build/contracts to be created.
	cat ./mainnet/build/contracts/ValidatorManagerContract.json | jq '.abi' > ./mainnet/build/ValidatorManagerContract.abi
	./abigen --abi ./mainnet/build/ValidatorManagerContract.abi --pkg ethcontract --type ValidatorManagerContract --out src/ethcontract/mainnet_validator_manager_contract.go 
	cat ./mainnet/build/contracts/ERC20Gateway.json | jq '.abi' > ./mainnet/build/ERC20Gateway.abi
	./abigen --abi ./mainnet/build/ERC20Gateway.abi --pkg ethcontract --type ERC20Gateway --out src/ethcontract/mainnet_erc20_gateway.go 
	cat ./mainnet/build/contracts/Gateway.json | jq '.abi' > ./mainnet/build/MainnetGatewayContract.abi
	./abigen --abi ./mainnet/build/MainnetGatewayContract.abi --pkg ethcontract --type MainnetGatewayContract --out src/ethcontract/mainnet_gateway.go 
	cat ./mainnet/build/contracts/CryptoCards.json | jq '.abi' > ./mainnet/build/MainnetCryptoCardsContract.abi
	cat ./mainnet/build/contracts/CryptoCards.json | jq '.bytecode' -j > ./mainnet/build/MainnetCryptoCardsContract.bin
	./abigen --abi ./mainnet/build/MainnetCryptoCardsContract.abi \
		--bin ./mainnet/build/MainnetCryptoCardsContract.bin \
		--pkg ethcontract --type MainnetCryptoCardsContract \
		--out src/ethcontract/mainnet_crypto_cards.go
	cat ./mainnet/build/contracts/GameToken.json | jq '.abi' > ./mainnet/build/MainnetGameTokenContract.abi
	cat ./mainnet/build/contracts/GameToken.json | jq '.bytecode' -j > ./mainnet/build/MainnetGameTokenContract.bin
	./abigen --abi ./mainnet/build/MainnetGameTokenContract.abi \
		--bin ./mainnet/build/MainnetGameTokenContract.bin \
		--pkg ethcontract --type MainnetGameTokenContract \
		--out src/ethcontract/mainnet_game_token.go
	cat ./mainnet/build/contracts/ERC721XCards.json | jq '.abi' > ./mainnet/build/MainnetERC721XCardsContract.abi
	cat ./mainnet/build/contracts/ERC721XCards.json | jq '.bytecode' -j > ./mainnet/build/MainnetERC721XCardsContract.bin
	./abigen --abi ./mainnet/build/MainnetERC721XCardsContract.abi \
		--bin ./mainnet/build/MainnetERC721XCardsContract.bin \
		--pkg ethcontract --type MainnetERC721XCardsContract \
		--out src/ethcontract/mainnet_erc721x_cards.go
	# Need to run yarn compile in the ./dappchain directory for build/contracts to be created.
	cat ./dappchain/build/contracts/ERC721DAppToken.json | jq '.abi' > ./src/ethcontract/ERC721DAppToken.abi
	cat ./dappchain/build/contracts/ERC721XDAppToken.json | jq '.abi' > ./src/ethcontract/ERC721XDAppToken.abi
	cat ./dappchain/build/contracts/ERC20DAppToken.json | jq '.abi' > ./src/ethcontract/ERC20DAppToken.abi
	cat ./dappchain/build/contracts/SampleERC721Token.json | jq '.abi' > ./src/ethcontract/SampleERC721Token.abi
	cat ./dappchain/build/contracts/SampleERC721Token.json | jq '.bytecode' -j > ./src/ethcontract/SampleERC721Token.bin
	cat ./dappchain/build/contracts/SampleERC20Token.json | jq '.abi' > ./src/ethcontract/SampleERC20Token.abi
	cat ./dappchain/build/contracts/SampleERC20Token.json | jq '.bytecode' -j > ./src/ethcontract/SampleERC20Token.bin
	cat ./dappchain/build/contracts/SampleERC721XToken.json | jq '.abi' > ./src/ethcontract/SampleERC721XToken.abi
	cat ./dappchain/build/contracts/SampleERC721XToken.json | jq '.bytecode' -j > ./src/ethcontract/SampleERC721XToken.bin
	cat ./dappchain/build/contracts/EthCoinIntegrationTest.json | jq '.abi' > ./src/ethcontract/EthCoinIntegrationTest.abi
	cat ./dappchain/build/contracts/EthCoinIntegrationTest.json | jq '.bytecode' -j > ./src/ethcontract/EthCoinIntegrationTest.bin
	cat ./dappchain/build/contracts/TRXToken.json | jq '.abi' > ./src/ethcontract/TRXToken.abi
	cat ./dappchain/build/contracts/TRXToken.json | jq '.bytecode' -j > ./src/ethcontract/TRXToken.bin

$(GO_ETHEREUM_DIR):
	git clone -q https://github.com/loomnetwork/go-ethereum.git $@

$(SSHA3_DIR):
	git clone -q https://github.com/loomnetwork/go-solidity-sha3.git $@

deps: $(GO_ETHEREUM_DIR) $(SSHA3_DIR)
	go get \
		github.com/gogo/protobuf/jsonpb \
		github.com/gogo/protobuf/proto \
		github.com/loomnetwork/go-loom \
		github.com/pkg/errors \
		github.com/go-kit/kit/log \
		github.com/inconshreveable/mousetrap \
		github.com/stretchr/testify \
		github.com/spf13/cobra \
		github.com/spf13/viper \
		github.com/grpc-ecosystem/go-grpc-prometheus \
		github.com/hashicorp/go-plugin \
		github.com/prometheus/client_golang/prometheus \
		github.com/phonkee/go-pubsub \
		github.com/gorilla/websocket \
		github.com/loomnetwork/yubihsm-go \
		github.com/btcsuite/btcd
	# cd $(GOPATH)/src/github.com/loomnetwork/go-loom && git checkout fix-split-sigs # only use when testing new go-loom features
	cd $(GO_ETHEREUM_DIR) && git checkout master && git pull && git checkout $(ETHEREUM_GIT_REV)
	cd $(GOLANG_PROTOBUF_DIR) && git checkout v1.1.0
	cd $(HASHICORP_DIR) && git checkout $(HASHICORP_GIT_REV)
	cd $(GRPC_DIR) && git checkout v1.20.1
	cd $(GENPROTO_DIR) && git checkout master && git pull && git checkout $(GENPROTO_GIT_REV)

clean:
	go clean
