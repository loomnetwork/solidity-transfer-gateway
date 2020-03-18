package gateway

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

func GetConfigDir() string {
	ethNet := os.Getenv("ETHEREUM_NETWORK")
	if ethNet == "" {
		ethNet = "ganache"
	}
	dappNet := os.Getenv("DAPPCHAIN_NETWORK")
	if dappNet == "" {
		dappNet = "local"
	}
	// When running "go test" the cwd is set to the package dir, not the root dir
	// where the config is, so gotta do a bit more work to figure out the config dir...
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), "../../e2e_config/"+dappNet+"_"+ethNet)
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
