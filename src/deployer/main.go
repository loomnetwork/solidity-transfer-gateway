package main

import (
	"client"
	"encoding/hex"
	"fmt"
	"gateway"
	"io/ioutil"
	"math/big"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	bnbclient "github.com/binance-chain/go-sdk/client"
	bnbtypes "github.com/binance-chain/go-sdk/common/types"
	"github.com/binance-chain/go-sdk/keys"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	loom "github.com/loomnetwork/go-loom"
	tgtypes "github.com/loomnetwork/go-loom/builtin/types/transfer_gateway"
	loom_client "github.com/loomnetwork/go-loom/client"
	"github.com/loomnetwork/go-loom/client/erc20"
	"github.com/loomnetwork/go-loom/client/erc721"
	"github.com/loomnetwork/go-loom/client/erc721x"
	gw "github.com/loomnetwork/go-loom/client/gateway"
	vmc "github.com/loomnetwork/go-loom/client/validator_manager"
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

func newDeployCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "deploy",
		Short: "Deploys test contracts",
		RunE:  deploy,
	}
}

func newDeployTronCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "deploy-tron",
		Short: "Deploys test contracts using Tron gateway",
		RunE:  deployTron,
	}
}

func newDeployBinanceCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "deploy-binance",
		Short: "Deploys test contracts using Binance gateway",
		RunE:  deployBinance,
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

func deploy(cmd *cobra.Command, args []string) error {
	ethereumContractsToDeploy := map[string]bool{}
	if len(cmdFlags.EthereumContractNames) > 0 {
		for _, contractName := range cmdFlags.EthereumContractNames {
			ethereumContractsToDeploy[contractName] = true
		}
	}

	dAppChainContractsToDeploy := map[string]bool{}
	if len(cmdFlags.DAppChainContractNames) > 0 {
		for _, contractName := range cmdFlags.DAppChainContractNames {
			dAppChainContractsToDeploy[contractName] = true
		}
	}
	fmt.Println("eth contract to deploy", ethereumContractsToDeploy)
	fmt.Println("dapp contract to deploy", dAppChainContractsToDeploy)

	loomCfg, err := gateway.ParseConfig([]string{cmdFlags.LoomDir})
	if err != nil {
		return errors.Wrap(err, "failed to parse loom config")
	}

	ethKey, dappchainKey := gateway.GetKeys("dan")
	fmt.Println("dan", ethKey, dappchainKey)
	erc721Creator, err := loom_client.CreateIdentityStr(ethKey, dappchainKey, loomCfg.ChainID)
	if err != nil {
		return errors.Wrap(err, "failed to create identity1")
	}

	ethKey, dappchainKey = gateway.GetKeys("trudy")
	erc20Creator, err := loom_client.CreateIdentityStr(ethKey, dappchainKey, loomCfg.ChainID)
	if err != nil {
		return errors.Wrap(err, "failed to create identity2")
	}

	deploymentInfo, err := parseEthereumDeploymentInfo(cmdFlags.EthereumDeploymentInfoPath)
	if err != nil {
		return errors.Wrap(err, "failed to load deployment info file")
	}

	// Deploy contracts to Ethereum

	if len(ethereumContractsToDeploy) > 0 {
		ethClient, err := ethclient.Dial(loomCfg.TransferGateway.EthereumURI)
		if err != nil {
			return errors.Wrap(err, "failed to connect to Ethereum network")
		}

		mainnetGatewayAddr := gateway.GetMainnetContractCfgString("mainnet_gateway_addr")
		mainnetGateway, err := gw.ConnectToMainnetGateway(ethClient, mainnetGatewayAddr)
		if err != nil {
			return errors.Wrap(err, "failed to connect to Gateway on Ethereum network")
		}

		// Get Validator key
		ethKey, dappchainKey = gateway.GetKeys("oracle")
		val, err := loom_client.CreateIdentityStr(ethKey, dappchainKey, loomCfg.ChainID)
		if err != nil {
			return errors.Wrap(err, "failed to create identity for validator")
		}

		mainnetVmcAddr := gateway.GetMainnetContractCfgString("mainnet_validatorManagerContract_addr")
		vmc, err := vmc.ConnectToMainnetVMCClient(ethClient, mainnetVmcAddr)
		if err != nil {
			return errors.Wrap(err, "failed to connect to Validator Manager Contract on Ethereum network")
		}

		ind := big.NewInt(0)
		if err := vmc.ToggleAllowAnyToken(val, true, ind); err != nil {
			return errors.Wrap(err, "failed to allow all tokens on VMC")
		}

		if ethereumContractsToDeploy["CryptoCards"] {
			contract, err := client.DeployMainnetCardsContract(ethClient, erc721Creator, mainnetGateway.Address)
			if err != nil {
				return errors.Wrap(err, "failed to deploy CryptoCards contract")
			}

			deploymentInfo.Set("mainnet_crypto_cards_addr", contract.Address)
			deploymentInfo.Set("mainnet_crypto_cards_tx", contract.TxHash)
		}

		if ethereumContractsToDeploy["ERC721XCards"] {
			contract, err := client.DeployMainnetERC721XContract(ethClient, erc721Creator, mainnetGateway.Address)
			if err != nil {
				return errors.Wrap(err, "failed to deploy ZBGCard contract")
			}

			deploymentInfo.Set("mainnet_erc721x_cards_addr", contract.Address)
			deploymentInfo.Set("mainnet_erc721x_cards_tx", contract.TxHash)
		}

		if ethereumContractsToDeploy["GameToken"] {
			contract, err := client.DeployMainnetERC20Contract(ethClient, erc20Creator, mainnetGateway.Address)
			if err != nil {
				return errors.Wrap(err, "failed to deploy GameToken contract")
			}

			deploymentInfo.Set("mainnet_game_token_addr", contract.Address)
			deploymentInfo.Set("mainnet_game_token_tx", contract.TxHash)
		}

		if ethereumContractsToDeploy["SampleERC20MintableToken"] {
			contract, err := client.DeployMainnetERC20MintableContract(ethClient, erc20Creator, mainnetGateway.Address)
			if err != nil {
				return errors.Wrap(err, "failed to deploy SampleERC20MintableToken contract")
			}

			deploymentInfo.Set("mainnet_erc20_mintable_token_addr", contract.Address)
			deploymentInfo.Set("mainnet_erc20_mintable_token_tx", contract.TxHash)
		}

		if ethereumContractsToDeploy["SampleERC721MintableToken"] {
			contract, err := client.DeployMainnetERC721MintableContract(ethClient, erc721Creator, mainnetGateway.Address)
			if err != nil {
				return errors.Wrap(err, "failed to deploy SampleERC721MintableToken contract")
			}

			deploymentInfo.Set("mainnet_erc721_mintable_token_addr", contract.Address)
			deploymentInfo.Set("mainnet_erc721_mintable_token_tx", contract.TxHash)
		}
	}

	// Deploy contracts to DAppChain

	if len(dAppChainContractsToDeploy) > 0 {
		contractDir := cmdFlags.ContractDir
		if contractDir == "" {
			contractDir, _ = os.Getwd()
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

		if dAppChainContractsToDeploy["SampleERC721Token"] {
			c, err := gateway.DeployTokenToDAppChain(
				loomClient, path.Join(contractDir, "SampleERC721Token.abi"),
				path.Join(contractDir, "SampleERC721Token.bin"),
				"SampleERC721Token", loomGateway.Address, erc721Creator.LoomSigner)
			if err != nil {
				return errors.Wrap(err, "failed to deploy SampleERC721Token")
			}
			fmt.Printf("SampleERC721Token at %v\n", c.Address)
		}

		if dAppChainContractsToDeploy["SampleERC721XToken"] {
			c, err := erc721x.DeployERC721XToDAppChain(
				loomClient, "SampleERC721XToken", loomGateway.Address, erc721Creator.LoomSigner)
			if err != nil {
				return errors.Wrap(err, "failed to deploy SampleERC721XToken")
			}
			fmt.Printf("SampleERC721XToken at %v\n", c.Address)
		}

		if dAppChainContractsToDeploy["SampleERC20Token"] {
			c, err := erc20.DeployERC20ToDAppChain(
				loomClient, "SampleERC20Token", loomGateway.Address, erc20Creator.LoomSigner)
			if err != nil {
				return errors.Wrap(err, "failed to deploy SampleERC20Token")
			}
			fmt.Printf("SampleERC20Token at %v\n", c.Address)
		}

		if dAppChainContractsToDeploy["SampleERC20Token2"] {
			c, err := gateway.DeployTokenToDAppChain(
				loomClient, path.Join(contractDir, "SampleERC20Token.abi"),
				path.Join(contractDir, "SampleERC20Token.bin"),
				"SampleERC20Token2", loomGateway.Address, erc20Creator.LoomSigner)
			if err != nil {
				return errors.Wrap(err, "failed to deploy SampleERC20Token2")
			}
			addr := c.Address.Local.String()
			fmt.Printf("SampleERC20Token2 at %v\n", addr)
			deploymentInfo.Set("dapp_token_for_erc20_mintable_token_addr", addr)
		}

		if dAppChainContractsToDeploy["SampleERC721Token2"] {
			c, err := gateway.DeployTokenToDAppChain(
				loomClient, path.Join(contractDir, "SampleERC721Token.abi"),
				path.Join(contractDir, "SampleERC721Token.bin"),
				"SampleERC721Token2", loomGateway.Address, erc721Creator.LoomSigner)
			if err != nil {
				return errors.Wrap(err, "failed to deploy SampleERC721Token2")
			}
			addr := c.Address.Local.String()
			fmt.Printf("SampleERC721Token2 at %v\n", addr)
			deploymentInfo.Set("dapp_token_for_erc721_mintable_token_addr", addr)
		}
	}

	return deploymentInfo.WriteConfig()
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
		loomCards, err := erc721.ConnectERC721ToDAppChain(loomClient, "SampleERC721Token")
		if err != nil {
			return err
		}

		ethContractAddress := deploymentInfo.GetString("mainnet_crypto_cards_addr")
		ethContractTxHash := deploymentInfo.GetString("mainnet_crypto_cards_tx")
		if !common.IsHexAddress(ethContractAddress) || ethContractTxHash == "" {
			return errors.New("missing Ethereum address and/or tx hash for ERC721 contract")
		}

		err = loomGateway.AddContractMapping(
			common.HexToAddress(ethContractAddress), loomCards.Address,
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
		loomERC721X, err := erc721x.ConnectERC721XToDAppChain(loomClient, "SampleERC721XToken")
		if err != nil {
			return err
		}

		ethContractAddress := deploymentInfo.GetString("mainnet_erc721x_cards_addr")
		ethContractTxHash := deploymentInfo.GetString("mainnet_erc721x_cards_tx")
		if !common.IsHexAddress(ethContractAddress) || ethContractTxHash == "" {
			return errors.New("missing Ethereum address and/or tx hash for ERC721X contract")
		}

		err = loomGateway.AddContractMapping(
			common.HexToAddress(ethContractAddress), loomERC721X.Address,
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
		loomCoin, err := erc20.ConnectERC20ToDAppChain(loomClient, "SampleERC20Token")
		if err != nil {
			return err
		}

		ethContractAddress := deploymentInfo.GetString("mainnet_game_token_addr")
		ethContractTxHash := deploymentInfo.GetString("mainnet_game_token_tx")
		if !common.IsHexAddress(ethContractAddress) || ethContractTxHash == "" {
			return errors.New("missing Ethereum address and/or tx hash for ERC20 contract")
		}

		err = loomGateway.AddContractMapping(
			common.HexToAddress(ethContractAddress), loomCoin.Address,
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
		dAppSampleERC20TokenAddr := deploymentInfo.GetString("dapp_token_for_erc20_mintable_token_addr")
		ethContractAddress := deploymentInfo.GetString("mainnet_erc20_mintable_token_addr")
		ethContractTxHash := deploymentInfo.GetString("mainnet_erc20_mintable_token_tx")
		if !common.IsHexAddress(ethContractAddress) || ethContractTxHash == "" {
			return errors.New("missing Ethereum address and/or tx hash for ERC20 contract")
		}
		dapplocalAddr, err := loom.LocalAddressFromHexString(dAppSampleERC20TokenAddr)
		if err != nil {
			return errors.Wrapf(err, "failed to convert %s to LocalAddress", dAppSampleERC20TokenAddr)
		}

		dappAddr := &loom.Address{
			ChainID: loomClient.GetChainID(),
			Local:   dapplocalAddr,
		}

		err = loomGateway.AddContractMapping(
			common.HexToAddress(ethContractAddress), *dappAddr,
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
		dAppSampleERC721TokenAddr := deploymentInfo.GetString("dapp_token_for_erc721_mintable_token_addr")
		ethContractAddress := deploymentInfo.GetString("mainnet_erc721_mintable_token_addr")
		ethContractTxHash := deploymentInfo.GetString("mainnet_erc721_mintable_token_tx")
		if !common.IsHexAddress(ethContractAddress) || ethContractTxHash == "" {
			return errors.New("missing Ethereum address and/or tx hash for ERC721 contract")
		}

		dapplocalAddr, err := loom.LocalAddressFromHexString(dAppSampleERC721TokenAddr)
		if err != nil {
			return errors.Wrapf(err, "failed to convert %s to LocalAddress", dAppSampleERC721TokenAddr)
		}

		dappAddr := &loom.Address{
			ChainID: loomClient.GetChainID(),
			Local:   dapplocalAddr,
		}

		err = loomGateway.AddContractMapping(
			common.HexToAddress(ethContractAddress), *dappAddr,
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

func deployBinance(cmd *cobra.Command, args []string) error {
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

	bnbMnemonic, dappchainKey := gateway.GetBnbKeys("token_owner")
	keyManager, err := keys.NewMnemonicKeyManager(bnbMnemonic)
	if err != nil {
		return err
	}
	bnbKey, err := keyManager.ExportAsPrivateKey()
	if err != nil {
		return err
	}
	tokenOwner, err := loom_client.CreateIdentityStr(bnbKey, dappchainKey, loomCfg.ChainID)
	if err != nil {
		return errors.Wrap(err, "failed to create identity")
	}

	if len(dAppChainContractsToDeploy) > 0 {
		loomClient := loom_client.NewDAppChainRPCClient(
			loomCfg.ChainID,
			loomCfg.TransferGateway.DAppChainWriteURI,
			loomCfg.TransferGateway.DAppChainReadURI,
		)

		loomGateway, err := gw.ConnectToDAppChainBinanceGateway(loomClient, loomCfg.TransferGateway.DAppChainEventsURI)
		if err != nil {
			return errors.Wrap(err, "failed to connect to Binance Gateway on DAppChain")
		}

		contractDir := cmdFlags.ContractDir
		if contractDir == "" {
			contractDir, _ = os.Getwd()
		}

		if dAppChainContractsToDeploy["BNBToken"] {
			c, err := gateway.DeployTokenToDAppChain(
				loomClient, path.Join(contractDir, "BNBToken.abi"),
				path.Join(contractDir, "BNBToken.bin"),
				"BNBToken", loomGateway.Address, tokenOwner.LoomSigner)
			if err != nil {
				return errors.Wrap(err, "failed to deploy BNBToken")
			}
			fmt.Printf("BNBToken at %v\n", c.Address)
			e2eDir := path.Dir(cmdFlags.EthereumDeploymentInfoPath)
			if err := os.MkdirAll(e2eDir, 0744); err != nil {
				return errors.Wrap(err, "failed to create directory")
			}
			filename := path.Join(e2eDir, "dapp_bnb_token_address")
			err = ioutil.WriteFile(filename, []byte(c.Address.String()), 0744)
			if err != nil {
				return errors.Wrap(err, "failed to write file dapp_bnb_token_address")
			}
			fmt.Println("wrote to file...", filename)
		}

		if dAppChainContractsToDeploy["SampleBEP2Token"] {
			c, err := gateway.DeployTokenToDAppChain(
				loomClient, path.Join(contractDir, "SampleBEP2Token.abi"),
				path.Join(contractDir, "SampleBEP2Token.bin"),
				"SampleBEP2Token", loomGateway.Address, tokenOwner.LoomSigner)
			if err != nil {
				return errors.Wrap(err, "failed to deploy SampleBEP2Token")
			}
			fmt.Printf("SampleBEP2Token at %v\n", c.Address)
			e2eDir := path.Dir(cmdFlags.EthereumDeploymentInfoPath)
			if err := os.MkdirAll(e2eDir, 0744); err != nil {
				return errors.Wrap(err, "failed to create directory")
			}
			filename := path.Join(e2eDir, "dapp_sample_bep2_token_address")
			err = ioutil.WriteFile(filename, []byte(c.Address.String()), 0744)
			if err != nil {
				return errors.Wrap(err, "failed to write file dapp_sample_bep2_token_address")
			}
			fmt.Println("wrote to file...", filename)
		}
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

	if dAppChainContracts["BNBToken"] {
		bnbToken, err := gateway.NewERC20TokenContract(loomClient, path.Join(contractDir, "SampleBEP2Token.abi"), "BNBToken")
		if err != nil {
			return err
		}
		// Fake token contract that will be mapped to native BNB token on Binance Dex
		fakeMainnetBNBTokenAddress := loom.MustParseAddress("binance:0x0000000000000000000000000000000000424e42")
		err = loomGateway.AddAuthorizedBinanceContractMapping(
			common.HexToAddress(fakeMainnetBNBTokenAddress.Local.Hex()), bnbToken.Address,
			gatewayOwner,
		)
		if err != nil {
			return errors.Wrap(err, "failed to map BNBToken contracts")
		}
		fmt.Printf("mapped %s <==> %s\n", bnbToken.Address.String(), fakeMainnetBNBTokenAddress.String())
	}

	if dAppChainContracts["SampleBEP2Token"] {
		moolToken, err := gateway.NewERC20TokenContract(loomClient, path.Join(contractDir, "SampleBEP2Token.abi"), "SampleBEP2Token")
		if err != nil {
			return err
		}
		// MOOL-CBC is assumed to have already issued on BinanceChain
		tokenNameHex := hex.EncodeToString([]byte("MOOL-CBC"))
		fakeMainnetMoolTokenAddress := common.HexToAddress(tokenNameHex)
		err = loomGateway.AddBinanceContractMapping(
			fakeMainnetMoolTokenAddress, moolToken.Address,
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

		fmt.Printf("mapped %s <==> %s\n", moolToken.Address.String(), fakeMainnetMoolTokenAddress.String())
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
		newDeployCmd(),
		newMapContractsCmd(),
		newDeployTronCmd(),
		newMapTronContractsCmd(),
		newIssueTokenCmd(),
		newDeployBinanceCmd(),
		newMapBinanceContractsCmd(),
	)

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
