// +build evm

package gateway

import (
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/auth"
	loomclient "github.com/loomnetwork/go-loom/client"
	"github.com/loomnetwork/go-loom/client/erc20"
	"github.com/spf13/viper"
)

func GetConfigDir() string {
	gwType := os.Getenv("GATEWAY_TYPE")
	if gwType == "" {
		gwType = "eth"
	}

	var mainNet string
	switch gwType {
	case "eth":
		mainNet = os.Getenv("ETHEREUM_NETWORK")
		if mainNet == "" {
			mainNet = "ganache"
		}
	case "tron-gateway":
		mainNet = os.Getenv("TRON_NETWORK")
		if mainNet == "" {
			mainNet = "shasta"
		}
	case "binance-gateway":
		mainNet = os.Getenv("BINANCE_NETWORK")
		if mainNet == "" {
			mainNet = "bnbtestnet"
		}
	}

	dappNet := os.Getenv("DAPPCHAIN_NETWORK")
	if dappNet == "" {
		dappNet = "local"
	}
	// When running "go test" the cwd is set to the package dir, not the root dir
	// where the config is, so gotta do a bit more work to figure out the config dir...
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), "../../e2e_config/"+dappNet+"_"+mainNet)
}

// Loads gateway_test_config.yml or equivalent from project root
func parseConfig(configName string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(configName)
	v.AddConfigPath(GetConfigDir())
	return v, v.ReadInConfig()
}

// TODO: make this less stupid
func GetTestAccountKey(name string) string {
	cfg, err := parseConfig("test_keys")
	if err != nil {
		log.Fatalf("failed to load config file: %v", err)
	}
	return cfg.GetString(name)
}

func GetMainnetContractCfgString(name string) string {
	cfg, err := parseConfig("contracts")
	if err != nil {
		log.Fatalf("failed to load config file: %v", err)
	}
	return cfg.GetString(name)
}

func LoadDAppChainContractABI(contractName string) (*abi.ABI, error) {
	_, filename, _, _ := runtime.Caller(0)
	abiPath := filepath.Join(filepath.Dir(filename), "../../src/ethcontract", contractName+".abi")
	abiBytes, err := ioutil.ReadFile(abiPath)
	if err != nil {
		return nil, err
	}
	contractABI, err := abi.JSON(strings.NewReader(string(abiBytes)))
	if err != nil {
		return nil, err
	}
	return &contractABI, nil
}

func LoadDAppChainContractCode(contractName string) ([]byte, error) {
	_, filename, _, _ := runtime.Caller(0)
	binPath := filepath.Join(filepath.Dir(filename), "../../src/ethcontract", contractName+".bin")
	hexByteCode, err := ioutil.ReadFile(binPath)
	if err != nil {
		return nil, err
	}
	return common.FromHex(string(hexByteCode)), nil
}

func GetKeys(name string) (string, string) {
	ethKey := GetTestAccountKey(name + "_eth")
	dappchainKey := GetTestAccountKey(name + "_dapp")
	return ethKey, dappchainKey
}

func GetTronKeys(name string) (string, string) {
	tronKey := GetTestAccountKey(name + "_tron")
	dappchainKey := GetTestAccountKey(name + "_dapp")
	return tronKey, dappchainKey
}

func GetBnbKeys(name string) (string, string) {
	bnbKey := GetTestAccountKey(name + "_bnb")
	dappchainKey := GetTestAccountKey(name + "_dapp")
	return bnbKey, dappchainKey
}

func ConnectToTokenContract(
	loomClient *loomclient.DAppChainRPCClient, contractABIPath string, contractName string,
) (*loomclient.MirroredTokenContract, error) {
	abiBytes, err := ioutil.ReadFile(contractABIPath)
	if err != nil {
		return nil, err
	}
	contractABI, err := abi.JSON(strings.NewReader(string(abiBytes)))
	if err != nil {
		return nil, err
	}

	contractAddr, err := loomClient.Resolve(contractName)
	if err != nil {
		return nil, err
	}

	return &loomclient.MirroredTokenContract{
		Contract:    loomclient.NewEvmContract(loomClient, contractAddr.Local),
		ContractABI: &contractABI,
		ChainID:     loomClient.GetChainID(),
		Address:     contractAddr,
	}, nil
}

func ConnectToTokenContractByAddress(
	loomClient *loomclient.DAppChainRPCClient, contractABIPath string, contractName string,
	contractAddr loom.Address,
) (*loomclient.MirroredTokenContract, error) {
	abiBytes, err := ioutil.ReadFile(contractABIPath)
	if err != nil {
		return nil, err
	}
	contractABI, err := abi.JSON(strings.NewReader(string(abiBytes)))
	if err != nil {
		return nil, err
	}

	return &loomclient.MirroredTokenContract{
		Contract:    loomclient.NewEvmContract(loomClient, contractAddr.Local),
		ContractABI: &contractABI,
		ChainID:     loomClient.GetChainID(),
		Address:     contractAddr,
	}, nil
}

func DeployTokenToDAppChain(loomClient *loomclient.DAppChainRPCClient, contractABIPath string,
	contractBinPath string, contractName string, gatewayAddr loom.Address, creator auth.Signer,
) (*loomclient.MirroredTokenContract, error) {
	abiBytes, err := ioutil.ReadFile(contractABIPath)
	if err != nil {
		return nil, err
	}
	contractABI, err := abi.JSON(strings.NewReader(string(abiBytes)))
	if err != nil {
		return nil, err
	}

	hexByteCode, err := ioutil.ReadFile(contractBinPath)
	if err != nil {
		return nil, err
	}
	byteCode := common.FromHex(string(hexByteCode))
	// append constructor args to bytecode
	input, err := contractABI.Pack("", common.BytesToAddress(gatewayAddr.Local))
	if err != nil {
		return nil, err
	}
	byteCode = append(byteCode, input...)
	contract, _, err := loomclient.DeployContract(loomClient, byteCode, creator, contractName)
	if err != nil {
		return nil, err
	}
	return &loomclient.MirroredTokenContract{
		Contract:    contract,
		ContractABI: &contractABI,
		ChainID:     loomClient.GetChainID(),
		Address:     contract.Address,
	}, nil
}

func NewERC20TokenContract(
	loomClient *loomclient.DAppChainRPCClient, contractABIPath string, contractName string,
) (*erc20.DAppChainERC20Contract, error) {
	mirroredTokenContract, err := ConnectToTokenContract(loomClient, contractABIPath, contractName)
	if err != nil {
		return nil, err
	}
	return &erc20.DAppChainERC20Contract{MirroredTokenContract: mirroredTokenContract}, nil
}

func sciNot(m int64) *big.Int {
	n := int64(18)
	ret := big.NewInt(10)
	ret.Exp(ret, big.NewInt(n), nil)
	ret.Mul(ret, big.NewInt(m))
	return ret
}

func amountAsString(m *big.Int) string {
	n := int64(18)
	return new(big.Rat).SetFrac(m, new(big.Int).Exp(big.NewInt(10), big.NewInt(n), nil)).FloatString(4)
}
