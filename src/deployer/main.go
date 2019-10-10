package main

import (
	"encoding/hex"
	"fmt"
	"gateway"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	bnbclient "github.com/binance-chain/go-sdk/client"
	bnbtypes "github.com/binance-chain/go-sdk/common/types"
	"github.com/binance-chain/go-sdk/keys"
	"github.com/ethereum/go-ethereum/common"
	loom "github.com/loomnetwork/go-loom"
	tgtypes "github.com/loomnetwork/go-loom/builtin/types/transfer_gateway"
	loom_client "github.com/loomnetwork/go-loom/client"
	"github.com/loomnetwork/go-loom/client/erc20"
	gw "github.com/loomnetwork/go-loom/client/gateway"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type RootCmdFlags struct {
	LoomDir                    string
	EthereumContractNames      []string
	DAppChainContractNames     []string
	EthereumDeploymentInfoPath string
	ContractDir                string
}

var cmdFlags RootCmdFlags
var RootCmd = &cobra.Command{
	Use:   "deployer",
	Short: "e2e test contracts deployer",
}

func newDeployTronCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "deploy-tron",
		Short: "Deploys test contracts using Tron gateway",
		RunE:  deployTron,
	}
}

var mapContractsTimeout int

func newMapContractsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "map-contracts",
		Short: "Adds contract mappings for test contracts",
		RunE:  mapContracts,
	}
	cmd.Flags().IntVar(&mapContractsTimeout, "timeout", 10,
		"Max number of seconds to wait for Oracle to confirm contract mapping.")

	return cmd
}

func newMapTronContractsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "map-tron-contracts",
		Short: "Adds tron contract mappings for test contracts",
		RunE:  mapTronContracts,
	}
	cmd.Flags().IntVar(&mapContractsTimeout, "timeout", 10,
		"Max number of seconds to wait for Oracle to confirm contract mapping.")

	return cmd
}

func newMapBinanceContractsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "map-binance-contracts",
		Short: "Adds binance hex address mappings for test contracts",
		RunE:  mapBinanceContracts,
	}
	cmd.Flags().IntVar(&mapContractsTimeout, "timeout", 10,
		"Max number of seconds to wait for Oracle to confirm contract mapping.")

	return cmd
}

var privateKeyFile string

func newIssueTokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bnb-issue-token",
		Short: "Issue token on BNB network",
		RunE:  bnbIssueToken,
	}
	cmd.Flags().StringVar(&privateKeyFile, "private-key", "", "Private key file")
	return cmd
}

func mapContracts(cmd *cobra.Command, args []string) error {
	dAppChainContracts := map[string]bool{}
	if len(cmdFlags.DAppChainContractNames) > 0 {
		for _, contractName := range cmdFlags.DAppChainContractNames {
			dAppChainContracts[contractName] = true
		}
	}

	if len(dAppChainContracts) == 0 {
		return nil
	}

	loomCfg, err := gateway.ParseConfig([]string{cmdFlags.LoomDir})
	if err != nil {
		return errors.Wrap(err, "failed to parse loom config")
	}

	ethKey, dappchainKey := gateway.GetKeys("dan")
	erc721Creator, err := loom_client.CreateIdentityStr(ethKey, dappchainKey, loomCfg.ChainID)
	if err != nil {
		return errors.Wrap(err, "failed to create identity3")
	}

	ethKey, dappchainKey = gateway.GetKeys("trudy")
	erc20Creator, err := loom_client.CreateIdentityStr(ethKey, dappchainKey, loomCfg.ChainID)
	if err != nil {
		return errors.Wrap(err, "failed to create identity4")
	}

	deploymentInfo, err := parseEthereumDeploymentInfo(cmdFlags.EthereumDeploymentInfoPath)
	if err != nil {
		return errors.Wrap(err, "failed to load deployment info file")
	}

	loomClient := loom_client.NewDAppChainRPCClient(
		loomCfg.ChainID,
		loomCfg.TransferGateway.DAppChainWriteURI,
		loomCfg.TransferGateway.DAppChainReadURI,
	)

	loomGateway, err := gw.ConnectToDAppChainGateway(loomClient, loomCfg.TransferGateway.DAppChainEventsURI)
	if err != nil {
		return errors.Wrap(err, "failed to connect to Gateway on DAppChain")
	}

	oracleWaitTime := time.Duration(mapContractsTimeout) * time.Second

	var contractMappingSub *gw.EventSub
	contractMappingConfirmedCh := make(chan *tgtypes.TransferGatewayContractMappingConfirmed, 1)

	if len(dAppChainContracts) > 0 {
		contractMappingSub, err = loomGateway.WatchContractMappingConfirmed(contractMappingConfirmedCh)
		if err != nil {
			return errors.Wrap(err, "failed to subscribe to DAppChain events")
		}
	}

	if dAppChainContracts["SampleERC721Token"] {
		localContractAddress := deploymentInfo.GetString("loomchain_crypto_cards_addr")
		local, err := loom.LocalAddressFromHexString(localContractAddress)
		if err != nil {
			return errors.Wrap(err, "parsing address fail")
		}
		loomLocalContractAddress := loom.Address{
			ChainID: loomCfg.ChainID,
			Local:   local,
		}

		ethContractAddress := deploymentInfo.GetString("mainnet_crypto_cards_addr")
		ethContractTxHash := deploymentInfo.GetString("mainnet_crypto_cards_tx")
		if !common.IsHexAddress(ethContractAddress) || ethContractTxHash == "" {
			return errors.New("missing Ethereum address and/or tx hash for ERC721 contract")
		}

		err = loomGateway.AddContractMapping(
			common.HexToAddress(ethContractAddress), loomLocalContractAddress,
			erc721Creator, ethContractTxHash,
		)
		if err != nil {
			return errors.Wrap(err, "failed to map ERC721 contracts")
		}

		// Let the Oracle fetch pending contract mappings and confirm them
		select {
		case <-contractMappingConfirmedCh:
		case <-time.After(oracleWaitTime):
			return errors.New("timeout while waiting for ContractMappingConfirmed event for ERC721 contracts")
		}
	}

	if dAppChainContracts["SampleERC721XToken"] {
		localContractAddress := deploymentInfo.GetString("loomchain_SampleERC721XToken_1")
		local, err := loom.LocalAddressFromHexString(localContractAddress)
		if err != nil {
			return errors.Wrap(err, "parsing address fail")
		}
		loomLocalContractAddress := loom.Address{
			ChainID: loomCfg.ChainID,
			Local:   local,
		}

		ethContractAddress := deploymentInfo.GetString("mainnet_erc721x_cards_addr")
		ethContractTxHash := deploymentInfo.GetString("mainnet_erc721x_cards_tx")
		if !common.IsHexAddress(ethContractAddress) || ethContractTxHash == "" {
			return errors.New("missing Ethereum address and/or tx hash for ERC721X contract")
		}

		err = loomGateway.AddContractMapping(
			common.HexToAddress(ethContractAddress), loomLocalContractAddress,
			erc721Creator, ethContractTxHash,
		)
		if err != nil {
			return errors.Wrap(err, "failed to map ERC721X contracts")
		}

		// Let the Oracle fetch pending contract mappings and confirm them
		select {
		case <-contractMappingConfirmedCh:
		case <-time.After(oracleWaitTime):
			return errors.New("timeout while waiting for ContractMappingConfirmed event for ERC721X contracts")
		}
	}

	if dAppChainContracts["SampleERC20Token"] {
		localContractAddress := deploymentInfo.GetString("loomchain_SampleERC20Token_1")
		local, err := loom.LocalAddressFromHexString(localContractAddress)
		if err != nil {
			return errors.Wrap(err, "parsing address fail")
		}
		loomLocalContractAddress := loom.Address{
			ChainID: loomCfg.ChainID,
			Local:   local,
		}

		ethContractAddress := deploymentInfo.GetString("mainnet_game_token_addr")
		ethContractTxHash := deploymentInfo.GetString("mainnet_game_token_tx")
		if !common.IsHexAddress(ethContractAddress) || ethContractTxHash == "" {
			return errors.New("missing Ethereum address and/or tx hash for ERC20 contract")
		}

		err = loomGateway.AddContractMapping(
			common.HexToAddress(ethContractAddress), loomLocalContractAddress,
			erc20Creator, ethContractTxHash,
		)
		if err != nil {
			return errors.Wrap(err, "failed to map ERC20 contracts")
		}

		// Let the Oracle fetch pending contract mappings and confirm them
		select {
		case <-contractMappingConfirmedCh:
		case <-time.After(oracleWaitTime):
			return errors.New("timeout while waiting for ContractMappingConfirmed event for ERC20 contracts")
		}
	}

	if dAppChainContracts["SampleERC20Token2"] {
		localContractAddress := deploymentInfo.GetString("loomchain_SampleERC20Token_2")
		local, err := loom.LocalAddressFromHexString(localContractAddress)
		if err != nil {
			return errors.Wrap(err, "parsing address fail")
		}
		loomLocalContractAddress := loom.Address{
			ChainID: loomCfg.ChainID,
			Local:   local,
		}

		ethContractAddress := deploymentInfo.GetString("mainnet_erc20_mintable_token_addr")
		ethContractTxHash := deploymentInfo.GetString("mainnet_erc20_mintable_token_tx")
		if !common.IsHexAddress(ethContractAddress) || ethContractTxHash == "" {
			return errors.New("missing Ethereum address and/or tx hash for ERC20 contract")
		}

		err = loomGateway.AddContractMapping(
			common.HexToAddress(ethContractAddress), loomLocalContractAddress,
			erc20Creator, ethContractTxHash,
		)
		if err != nil {
			return errors.Wrap(err, "failed to map ERC20 contracts")
		}

		// Let the Oracle fetch pending contract mappings and confirm them
		select {
		case <-contractMappingConfirmedCh:
		case <-time.After(oracleWaitTime):
			return errors.New("timeout while waiting for ContractMappingConfirmed event for ERC20 contracts")
		}
	}

	if dAppChainContracts["SampleERC721Token2"] {
		localContractAddress := deploymentInfo.GetString("loomchain_erc721_mintable_token_addr")
		local, err := loom.LocalAddressFromHexString(localContractAddress)
		if err != nil {
			return errors.Wrap(err, "parsing address fail")
		}
		loomLocalContractAddress := loom.Address{
			ChainID: loomCfg.ChainID,
			Local:   local,
		}

		ethContractAddress := deploymentInfo.GetString("mainnet_erc721_mintable_token_addr")
		ethContractTxHash := deploymentInfo.GetString("mainnet_erc721_mintable_token_tx")
		if !common.IsHexAddress(ethContractAddress) || ethContractTxHash == "" {
			return errors.New("missing Ethereum address and/or tx hash for ERC721 contract")
		}

		err = loomGateway.AddContractMapping(
			common.HexToAddress(ethContractAddress), loomLocalContractAddress,
			erc721Creator, ethContractTxHash,
		)
		if err != nil {
			return errors.Wrap(err, "failed to map ERC721 contracts")
		}

		// Let the Oracle fetch pending contract mappings and confirm them
		select {
		case <-contractMappingConfirmedCh:
		case <-time.After(oracleWaitTime):
			return errors.New("timeout while waiting for ContractMappingConfirmed event for ERC721 contracts")
		}
	}

	if contractMappingSub != nil {
		contractMappingSub.Close()
	}
	return nil
}

func deployTron(cmd *cobra.Command, args []string) error {
	dAppChainContractsToDeploy := map[string]bool{}
	if len(cmdFlags.DAppChainContractNames) > 0 {
		for _, contractName := range cmdFlags.DAppChainContractNames {
			dAppChainContractsToDeploy[contractName] = true
		}
	}

	loomCfg, err := gateway.ParseConfig([]string{cmdFlags.LoomDir})
	if err != nil {
		return errors.Wrap(err, "failed to parse loom config")
	}

	tronKey, dappchainKey := gateway.GetTronKeys("trudy")
	erc20Creator, err := loom_client.CreateIdentityStr(tronKey, dappchainKey, loomCfg.ChainID)
	if err != nil {
		return errors.Wrap(err, "failed to create identity5")
	}

	// Deploy contracts to DAppChain

	if len(dAppChainContractsToDeploy) > 0 {
		loomClient := loom_client.NewDAppChainRPCClient(
			loomCfg.ChainID,
			loomCfg.TransferGateway.DAppChainWriteURI,
			loomCfg.TransferGateway.DAppChainReadURI,
		)

		loomGateway, err := gw.ConnectToDAppChainTronGateway(loomClient, loomCfg.TransferGateway.DAppChainEventsURI)
		if err != nil {
			return errors.Wrap(err, "failed to connect to Gateway on DAppChain")
		}

		if dAppChainContractsToDeploy["TRXToken"] {
			c, err := erc20.DeployERC20ToDAppChain(
				loomClient, "TRXToken", loomGateway.Address, erc20Creator.LoomSigner)
			if err != nil {
				return errors.Wrap(err, "failed to deploy TRXToken")
			}
			fmt.Printf("TRXToken at %v\n", c.Address)
			// write to file for tron test
			e2eDir := path.Dir(cmdFlags.EthereumDeploymentInfoPath)
			if err := os.MkdirAll(e2eDir, 0744); err != nil {
				return errors.Wrap(err, "failed to create directory")
			}
			filename := path.Join(e2eDir, "dapp_trx_token_address")
			err = ioutil.WriteFile(filename, []byte(c.Address.String()), 0744)
			if err != nil {
				return errors.Wrap(err, "failed to write file dapp_trx_token_address")
			}
			fmt.Println("wrote to file...", filename)
		}

		if dAppChainContractsToDeploy["SampleERC20Token"] {
			c, err := erc20.DeployERC20ToDAppChain(
				loomClient, "SampleERC20Token", loomGateway.Address, erc20Creator.LoomSigner)
			if err != nil {
				return errors.Wrap(err, "failed to deploy SampleERC20Token")
			}
			fmt.Printf("SampleERC20Token at %v\n", c.Address)
			// write to file for tron test
			e2eDir := path.Dir(cmdFlags.EthereumDeploymentInfoPath)
			if err := os.MkdirAll(e2eDir, 0744); err != nil {
				return errors.Wrap(err, "failed to create directory")
			}
			filename := path.Join(e2eDir, "dapp_trc20_token_address")
			err = ioutil.WriteFile(filename, []byte(c.Address.String()), 0744)
			if err != nil {
				return errors.Wrap(err, "failed to write file dapp_trc20_token_address")
			}
			fmt.Println("wrote to file...", filename)
		}
	}

	return nil
}

func mapTronContracts(cmd *cobra.Command, args []string) error {
	dAppChainContracts := map[string]bool{}
	if len(cmdFlags.DAppChainContractNames) > 0 {
		for _, contractName := range cmdFlags.DAppChainContractNames {
			dAppChainContracts[contractName] = true
		}
	}

	if len(dAppChainContracts) == 0 {
		return nil
	}

	loomCfg, err := gateway.ParseConfig([]string{cmdFlags.LoomDir})
	if err != nil {
		return errors.Wrap(err, "failed to parse loom config")
	}

	tronKey, dappchainKey := gateway.GetTronKeys("trudy")
	erc20Creator, err := loom_client.CreateIdentityStr(tronKey, dappchainKey, loomCfg.ChainID)
	if err != nil {
		return errors.Wrap(err, "failed to create identity6")
	}

	tronKey, dappchainKey = gateway.GetTronKeys("gateway_owner")
	gatewayOwner, err := loom_client.CreateIdentityStr(tronKey, dappchainKey, loomCfg.ChainID)
	if err != nil {
		return errors.Wrap(err, "failed to create identity7")
	}

	deploymentInfo, err := parseEthereumDeploymentInfo(cmdFlags.EthereumDeploymentInfoPath)
	if err != nil {
		return errors.Wrap(err, "failed to load deployment info file")
	}

	loomClient := loom_client.NewDAppChainRPCClient(
		loomCfg.ChainID,
		loomCfg.TransferGateway.DAppChainWriteURI,
		loomCfg.TransferGateway.DAppChainReadURI,
	)

	loomGateway, err := gw.ConnectToDAppChainTronGateway(loomClient, loomCfg.TransferGateway.DAppChainEventsURI)
	if err != nil {
		return errors.Wrap(err, "failed to connect to Gateway on DAppChain")
	}

	oracleWaitTime := time.Duration(mapContractsTimeout) * time.Second

	var contractMappingSub *gw.EventSub
	contractMappingConfirmedCh := make(chan *tgtypes.TransferGatewayContractMappingConfirmed, 1)

	if len(dAppChainContracts) > 0 {
		contractMappingSub, err = loomGateway.WatchContractMappingConfirmed(contractMappingConfirmedCh)
		if err != nil {
			return errors.Wrap(err, "failed to subscribe to DAppChain events")
		}
	}

	if dAppChainContracts["TRXToken"] {
		TRXToken, err := erc20.ConnectERC20ToDAppChain(loomClient, "TRXToken")
		if err != nil {
			return err
		}
		fakeTRXContractAddress := loom.MustParseAddress("tron:0x0000000000000000000000000000000000000001")
		err = loomGateway.AddAuthorizedTronContractMapping(
			common.HexToAddress(fakeTRXContractAddress.Local.Hex()), TRXToken.Address,
			gatewayOwner,
		)
		if err != nil {
			return errors.Wrap(err, "failed to map ERC20 contracts")
		}
	}

	if dAppChainContracts["SampleERC20Token"] {
		loomCoin, err := erc20.ConnectERC20ToDAppChain(loomClient, "SampleERC20Token")
		if err != nil {
			return err
		}

		tronContractAddress := deploymentInfo.GetString("loomtoken_addr")
		tronContractAddress = strings.TrimPrefix(tronContractAddress, "41")
		if !common.IsHexAddress(tronContractAddress) {
			return errors.New("missing Tron address for ERC20 contract")
		}

		// we are not able txHash when we deploy contract via tronbox.
		// so the hacky way to get gateway checking it to use tronContractAddress
		// as a key for the gateway.
		err = loomGateway.AddTronContractMapping(
			common.HexToAddress(tronContractAddress), loomCoin.Address,
			erc20Creator, tronContractAddress,
		)
		if err != nil {
			return errors.Wrap(err, "failed to map ERC20 contracts")
		}

		// Let the Oracle fetch pending contract mappings and confirm them
		select {
		case <-contractMappingConfirmedCh:
		case <-time.After(oracleWaitTime):
			return errors.New("timeout while waiting for ContractMappingConfirmed event for ERC20 contracts")
		}
	}

	if contractMappingSub != nil {
		contractMappingSub.Close()
	}
	return nil
}

func mapBinanceContracts(cmd *cobra.Command, args []string) error {
	dAppChainContracts := map[string]bool{}
	if len(cmdFlags.DAppChainContractNames) > 0 {
		for _, contractName := range cmdFlags.DAppChainContractNames {
			dAppChainContracts[contractName] = true
		}
	}

	if len(dAppChainContracts) == 0 {
		return nil
	}

	loomCfg, err := gateway.ParseConfig([]string{cmdFlags.LoomDir})
	if err != nil {
		return errors.Wrap(err, "failed to parse loom config")
	}

	bnbKey, dappchainKey := gateway.GetBnbKeys("gateway_owner")
	keyManager, err := keys.NewMnemonicKeyManager(bnbKey)
	if err != nil {
		return err
	}
	privkey, err := keyManager.ExportAsPrivateKey()
	if err != nil {
		return err
	}
	gatewayOwner, err := loom_client.CreateIdentityStr(privkey, dappchainKey, loomCfg.ChainID)
	if err != nil {
		return errors.Wrap(err, "failed to create identity")
	}

	bnbKey, dappchainKey = gateway.GetBnbKeys("token_owner")
	keyManager, err = keys.NewMnemonicKeyManager(bnbKey)
	if err != nil {
		return err
	}
	privkey, err = keyManager.ExportAsPrivateKey()
	if err != nil {
		return err
	}
	tokenOwner, err := loom_client.CreateIdentityStr(privkey, dappchainKey, loomCfg.ChainID)
	if err != nil {
		return err
	}

	loomClient := loom_client.NewDAppChainRPCClient(
		loomCfg.ChainID,
		loomCfg.TransferGateway.DAppChainWriteURI,
		loomCfg.TransferGateway.DAppChainReadURI,
	)

	loomGateway, err := gw.ConnectToDAppChainBinanceGateway(loomClient, loomCfg.TransferGateway.DAppChainEventsURI)
	if err != nil {
		return errors.Wrap(err, "failed to connect to Gateway on DAppChain")
	}

	oracleWaitTime := time.Duration(mapContractsTimeout) * time.Second

	var contractMappingSub *gw.EventSub
	contractMappingConfirmedCh := make(chan *tgtypes.TransferGatewayContractMappingConfirmed, 1)

	if len(dAppChainContracts) > 0 {
		contractMappingSub, err = loomGateway.WatchContractMappingConfirmed(contractMappingConfirmedCh)
		if err != nil {
			return errors.Wrap(err, "failed to subscribe to DAppChain events")
		}
	}

	contractDir := cmdFlags.ContractDir
	if contractDir == "" {
		contractDir, _ = os.Getwd()
	}

	deploymentInfo, err := parseEthereumDeploymentInfo(cmdFlags.EthereumDeploymentInfoPath)
	if err != nil {
		return errors.Wrap(err, "failed to load deployment info file")
	}

	if dAppChainContracts["BNBToken"] {
		localContractAddress := deploymentInfo.GetString("loomchain_bnb_token_addr")
		local, err := loom.LocalAddressFromHexString(localContractAddress)
		if err != nil {
			return errors.Wrap(err, "parsing address fail")
		}
		loomLocalContractAddress := loom.Address{
			ChainID: loomCfg.ChainID,
			Local:   local,
		}

		// Fake token contract that will be mapped to native BNB token on Binance Dex
		fakeMainnetBNBTokenAddress := loom.MustParseAddress("binance:0x0000000000000000000000000000000000424e42")
		err = loomGateway.AddAuthorizedBinanceContractMapping(
			common.HexToAddress(fakeMainnetBNBTokenAddress.Local.Hex()), loomLocalContractAddress,
			gatewayOwner,
		)
		if err != nil {
			return errors.Wrap(err, "failed to map BNBToken contracts")
		}
		fmt.Printf("mapped %s <==> %s\n", localContractAddress, fakeMainnetBNBTokenAddress.String())
	}

	if dAppChainContracts["SampleBEP2Token"] {
		localContractAddress := deploymentInfo.GetString("loomchain_bep2_token_addr")
		local, err := loom.LocalAddressFromHexString(localContractAddress)
		if err != nil {
			return errors.Wrap(err, "parsing address fail")
		}
		loomLocalContractAddress := loom.Address{
			ChainID: loomCfg.ChainID,
			Local:   local,
		}

		// MOOL-CBC is assumed to have already issued on BinanceChain
		tokenNameHex := hex.EncodeToString([]byte("MOOL-CBC"))
		fakeMainnetMoolTokenAddress := common.HexToAddress(tokenNameHex)
		err = loomGateway.AddBinanceContractMapping(
			fakeMainnetMoolTokenAddress, loomLocalContractAddress,
			tokenOwner,
		)
		if err != nil {
			return errors.Wrap(err, "failed to map SampleBEP2Token contracts")
		}

		// Let the Oracle fetch pending contract mappings and confirm them
		select {
		case <-contractMappingConfirmedCh:
		case <-time.After(oracleWaitTime):
			return errors.New("timeout while waiting for ContractMappingConfirmed event for ERC20 contracts")
		}

		fmt.Printf("mapped %s <==> %s\n", localContractAddress, fakeMainnetMoolTokenAddress.String())
	}

	if contractMappingSub != nil {
		contractMappingSub.Close()
	}

	return nil
}

func bnbIssueToken(cmd *cobra.Command, args []string) error {
	keyManager, err := keys.NewPrivateKeyManager("466090730f432eaa3a412ca2431e829999f781adc65a5917a27def68e6928e58")
	if err != nil {
		return err
	}

	client, err := bnbclient.NewDexClient("testnet-dex.binance.org", bnbtypes.TestNetwork, keyManager)
	if err != nil {
		return err
	}
	// supply needs to be multiplied by 10^8
	issue, err := client.IssueToken("MOOL_Token", "MOOL", 100000000000000000, true, true)
	if err != nil {
		return err
	}
	fmt.Printf("result: %+v", issue)
	return nil
}

func parseEthereumDeploymentInfo(filename string) (*viper.Viper, error) {
	v := viper.New()
	name := filepath.Base(filename)
	name = strings.TrimSuffix(name, filepath.Ext(name))
	dir := filepath.Dir(filename)
	v.SetConfigName(name)
	v.AddConfigPath(dir)
	return v, v.ReadInConfig()
}

func main() {
	pflags := RootCmd.PersistentFlags()
	pflags.StringVar(&cmdFlags.LoomDir, "loom-dir", "", "Directory containing loom.yml")
	pflags.StringVar(&cmdFlags.EthereumDeploymentInfoPath, "deployment-file", "",
		"YAML file containing info about contracts deployed to Ethereum")
	pflags.StringSliceVar(&cmdFlags.EthereumContractNames, "ethereum-contracts", nil, "Names of contracts to deploy to Ethereum network")
	pflags.StringSliceVar(&cmdFlags.DAppChainContractNames, "dappchain-contracts", nil, "Names of contracts to deploy to DAppChain")
	pflags.StringVar(&cmdFlags.ContractDir, "contract-dir", "", "Directory containing contract abi and bin. Default to current dir")
	RootCmd.MarkFlagRequired("loom-dir")
	RootCmd.MarkFlagRequired("deployment-file")
	RootCmd.MarkFlagFilename("deployment-file")

	RootCmd.AddCommand(
		newMapContractsCmd(),
		newDeployTronCmd(),
		newMapTronContractsCmd(),
		newIssueTokenCmd(),
		newMapBinanceContractsCmd(),
	)

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
