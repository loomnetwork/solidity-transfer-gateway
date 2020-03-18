package gateway

import (
	"path/filepath"

	"github.com/spf13/viper"
)

type LoomConfig struct {
	ChainID         string
	TransferGateway *TransferGatewayConfig
}

// This is a subset of the settings from loom.yml that are used by the tests.
type TransferGatewayConfig struct {
	// URI of Ethereum node the Oracle should connect to
	EthereumURI       string
	DAppChainReadURI  string
	DAppChainWriteURI string
	// Websocket URI that should be used to subscribe to DAppChain events
	DAppChainEventsURI string
	// Number of Ethereum block confirmations the Oracle should wait for before forwarding events
	// from the Ethereum Gateway contract to the DAppChain Gateway contract.
	NumMainnetBlockConfirmations int
}

func defaultConfig() *LoomConfig {
	return &LoomConfig{
		ChainID: "default",
		TransferGateway: &TransferGatewayConfig{
			EthereumURI:                  "ws://127.0.0.1:8545",
			DAppChainReadURI:             "http://localhost:46658/query",
			DAppChainWriteURI:            "http://localhost:46658/rpc",
			DAppChainEventsURI:           "ws://localhost:46658/queryws",
			NumMainnetBlockConfirmations: 1,
		},
	}
}

// Loads loom.yml or equivalent from one of the usual location, or if overrideCfgDirs is provided
// from one of those config directories.
func ParseConfig(overrideCfgDirs []string) (*LoomConfig, error) {
	v := viper.New()
	v.SetConfigName("loom")
	if len(overrideCfgDirs) == 0 {
		// look for the loom config file in all the places loom itself does
		v.AddConfigPath(".")
		v.AddConfigPath(filepath.Join(".", "config"))
	} else {
		for _, dir := range overrideCfgDirs {
			v.AddConfigPath(dir)
		}
	}
	v.ReadInConfig()
	conf := defaultConfig()
	err := v.Unmarshal(conf)
	if err != nil {
		return nil, err
	}
	return conf, err
}
