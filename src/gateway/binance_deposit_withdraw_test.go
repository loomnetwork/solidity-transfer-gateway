package gateway

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	bnbclient "github.com/binance-chain/go-sdk/client"
	"github.com/binance-chain/go-sdk/client/transaction"
	bnbtypes "github.com/binance-chain/go-sdk/common/types"
	"github.com/binance-chain/go-sdk/keys"
	"github.com/binance-chain/go-sdk/types/msg"
	"github.com/ethereum/go-ethereum/common"
	loom_client "github.com/loomnetwork/go-loom/client"
	"github.com/loomnetwork/go-loom/client/erc20"
	gwclient "github.com/loomnetwork/go-loom/client/gateway_v2"
	"github.com/loomnetwork/go-loom/client/native_coin"
	"github.com/stretchr/testify/suite"
)

const (
	// Loom token asset name on binance dex
	BinanceLoomToken = "LOOM-172"

	// BEP2 token asset name on binance dex
	MoolToken = "MOOL-CBC"

	// BNB native token
	BNBNativeToken = "BNB"
)

type BinanceTransferGatewayTestSuite struct {
	suite.Suite
	oracleWaitTime time.Duration
	loomClient     *loom_client.DAppChainRPCClient

	dappchainGateway *gwclient.DAppChainGateway

	loomCoin                     *native_coin.DAppChainNativeCoin
	bnbToken                     *erc20.DAppChainERC20Contract
	sampleBEP2Token              *erc20.DAppChainERC20Contract
	numMainnetBlockConfirmations int

	// These identities are shared by all the tests
	gatewayCreator *loom_client.Identity
	alice          *loom_client.Identity
	bob            *loom_client.Identity
	tokenOwner     *loom_client.Identity

	tokenOwnerBnbAddress bnbtypes.AccAddress
	aliceBnbAddress      bnbtypes.AccAddress
	bobBnbAddress        bnbtypes.AccAddress

	tokenOwnerDexClient bnbclient.DexClient
	aliceDexClient      bnbclient.DexClient
	bobDexClient        bnbclient.DexClient

	onTestnet bool
	baseURL   string
}

func TestBinanceTransferGatewayTestSuite(t *testing.T) {
	suite.Run(t, new(BinanceTransferGatewayTestSuite))
}

func (s *BinanceTransferGatewayTestSuite) SetupSuite() {
	require := s.Require()
	require.NotEmpty(os.Getenv("LOOM_DIR"), "LOOM_DIR env var should be set to dir containing loom.yml")
	var err error
	s.oracleWaitTime = time.Duration(10) * time.Second
	if os.Getenv("ORACLE_WAIT_TIME") != "" {
		secs, err := strconv.ParseInt(os.Getenv("ORACLE_WAIT_TIME"), 10, 32)
		require.NoError(err)
		s.oracleWaitTime = time.Duration(secs) * time.Second
	}

	bnbNet := os.Getenv("BINANCE_NETWORK")
	if bnbNet == "" || bnbNet == "bnbtestnet" {
		s.onTestnet = true
		bnbtypes.Network = bnbtypes.TestNetwork
		s.baseURL = "testnet-dex.binance.org"
	}

	loomCfg, err := ParseConfig([]string{os.Getenv("LOOM_DIR")})
	require.NoError(err)

	fmt.Println(loomCfg.ChainID, loomCfg.TransferGateway.DAppChainReadURI, loomCfg.TransferGateway.DAppChainWriteURI)

	s.loomClient = loom_client.NewDAppChainRPCClient(
		loomCfg.ChainID,
		loomCfg.TransferGateway.DAppChainWriteURI,
		loomCfg.TransferGateway.DAppChainReadURI,
	)

	s.numMainnetBlockConfirmations = loomCfg.TransferGateway.NumMainnetBlockConfirmations

	s.loomCoin, err = native_coin.ConnectToDAppChainLoomContract(s.loomClient)
	require.NoError(err)
	s.bnbToken, err = NewERC20TokenContract(s.loomClient, "../ethcontract/SampleBEP2Token.abi", "BNBToken")
	require.NoError(err)
	s.sampleBEP2Token, err = NewERC20TokenContract(s.loomClient, "../ethcontract/SampleBEP2Token.abi", "SampleBEP2Token")
	require.NoError(err)

	// Create identities
	var bnbKey string
	var keyManager keys.KeyManager

	// prevent binance API rate limit
	time.Sleep(20 * time.Second)

	bnbKey, dappchainKey := GetBnbKeys("token_owner")
	keyManager, err = keys.NewMnemonicKeyManager(bnbKey)
	require.NoError(err)
	privkey, err := keyManager.ExportAsPrivateKey()
	require.NoError(err)
	s.tokenOwner, err = loom_client.CreateIdentityStr(privkey, dappchainKey, s.loomClient.GetChainID())
	require.NoError(err)
	s.tokenOwnerDexClient, err = bnbclient.NewDexClient(s.baseURL, bnbtypes.Network, keyManager)
	require.NoError(err)
	s.tokenOwnerBnbAddress = keyManager.GetAddr()

	// prevent binance API rate limit
	time.Sleep(20 * time.Second)

	bnbKey, dappchainKey = GetBnbKeys("alice")
	keyManager, err = keys.NewMnemonicKeyManager(bnbKey)
	require.NoError(err)
	privkey, err = keyManager.ExportAsPrivateKey()
	require.NoError(err)
	s.alice, err = loom_client.CreateIdentityStr(privkey, dappchainKey, s.loomClient.GetChainID())
	require.NoError(err)
	s.aliceDexClient, err = bnbclient.NewDexClient(s.baseURL, bnbtypes.Network, keyManager)
	require.NoError(err)
	s.aliceBnbAddress = keyManager.GetAddr()

	// prevent binance API rate limit
	time.Sleep(20 * time.Second)

	fmt.Println(time.Now().UTC())
}

func (s *BinanceTransferGatewayTestSuite) TestLoomDepositAndWithdraw() {
	s.T().Skip("Skip using loomcoin to map with BNB token test. We're not using this feature at the moment.")
	var err error
	require := s.Require()
	require.NoError(err)

	loomCfg, err := ParseConfig([]string{os.Getenv("LOOM_DIR")})
	require.NoError(err)
	// Connect dappchain contracts
	s.dappchainGateway, err = gwclient.ConnectToDAppChainBinanceGateway(s.loomClient, loomCfg.TransferGateway.DAppChainEventsURI)
	require.NoError(err)

	// prevent binance API rate limit
	time.Sleep(60 * time.Second)

	aliceAccount, err := s.aliceDexClient.GetAccount(s.aliceBnbAddress.String())
	require.NoError(err)
	aliceMainnetLoomCoinStartBal := getBEP2TokenBalance(aliceAccount.Balances, BinanceLoomToken)
	fmt.Println("ALICE MAINNET BALANCE", aliceMainnetLoomCoinStartBal)
	aliceLoomCoinStartBal, err := s.loomCoin.BalanceOf(s.alice)
	require.NoError(err)
	fmt.Println("ALICE DAPPCHAIN BALANCE", aliceLoomCoinStartBal)
	curMainnetBalance := aliceMainnetLoomCoinStartBal
	curDppchainBalance := aliceLoomCoinStartBal

	// prevent binance API rate limit
	time.Sleep(10 * time.Second)

	var amount int64 = 10 * 1e8
	var tokenAmount = big.NewInt(amount)

	// Alice deposits to wallet
	payload := []msg.Transfer{
		msg.Transfer{
			ToAddr: s.tokenOwnerBnbAddress,
			Coins: []bnbtypes.Coin{
				bnbtypes.Coin{
					Denom:  BinanceLoomToken,
					Amount: amount,
				},
			},
		},
	}
	// make sure we have put loom address in memo
	aliceDappchainAddr := "loom" + s.alice.LoomAddr.Local.Hex()
	result, err := s.aliceDexClient.SendToken(payload, true, transaction.WithMemo(aliceDappchainAddr))
	require.NoError(err)
	fmt.Println("send token tx hash on Binance Dex: ", result.Hash)

	time.Sleep(15 * time.Second)

	// Alice should now have her loomcoin in the DAppChain Loom contract
	aliceAccount, err = s.aliceDexClient.GetAccount(s.aliceBnbAddress.String())
	require.NoError(err)
	aliceMainnetLoomCoinStartBal = getBEP2TokenBalance(aliceAccount.Balances, BinanceLoomToken)
	fmt.Println("ALICE MAINNET BALANCE AFTER DEPOSIT", aliceMainnetLoomCoinStartBal)
	aliceLoomCoinStartBal, err = s.loomCoin.BalanceOf(s.alice)
	require.NoError(err)
	fmt.Println("ALICE DAPPCHAIN BALANCE AFTER DEPOSIT", aliceLoomCoinStartBal)

	require.EqualValues(
		tokenAmount.String(),
		new(big.Int).Sub(aliceLoomCoinStartBal, curDppchainBalance).String(),
		"Alice should have loom coin deposited the same amount in mainnet transaction")

	require.EqualValues(
		amount,
		curMainnetBalance-aliceMainnetLoomCoinStartBal,
		"Alice should have loom coin different the same amount as deposit in mainnet")

	curMainnetBalance = aliceMainnetLoomCoinStartBal
	curDppchainBalance = aliceLoomCoinStartBal

	// Alice must grant approval to the DAppChain Gateway to take ownership of the tokens when they're withdrawn
	require.NoError(s.loomCoin.Approve(s.alice, s.dappchainGateway.Address, tokenAmount))
	// Now Alice can requests a withdrawal from the DAppChain Gateway...
	// Waiting for oracle to clear a previous receipt
	for i := 0; i < 5; i++ {
		err = s.dappchainGateway.WithdrawLoomToBinanceDex(s.alice, tokenAmount, common.BytesToAddress(s.aliceBnbAddress.Bytes()))
		if err != nil {
			if strings.Contains(err.Error(), "TG003") {
				time.Sleep(5 * time.Second)
			} else {
				require.NoError(err)
			}
		} else {
			break
		}
	}
	require.NoError(err)

	// Let the Oracle fetch pending withdrawals & sign them
	// WARNING: timing issue when tx hash is deleted before this test fetches WithdrawalReceipt
	wr, err := s.dappchainGateway.WithdrawalReceipt(s.alice)
	for i := 0; i < 15; i++ {
		wr, err = s.dappchainGateway.WithdrawalReceipt(s.alice)
		if wr.TxHash != nil {
			break
		}
		fmt.Println("waiting for Oracle to update tx_hash...")
		time.Sleep(10 * time.Second)
	}
	require.NoError(err)
	require.NotNil(wr.TxHash)

	for i := 0; i < 15; i++ {
		wr, err = s.dappchainGateway.WithdrawalReceipt(s.alice)
		if wr == nil {
			break
		}
		fmt.Println("waiting for Oracle to transfer token...")
		time.Sleep(10 * time.Second)
	}
	require.Nil(wr)

	aliceAccount, err = s.aliceDexClient.GetAccount(s.aliceBnbAddress.String())
	require.NoError(err)
	aliceMainnetLoomCoinStartBal = getBEP2TokenBalance(aliceAccount.Balances, BinanceLoomToken)
	fmt.Println("ALICE MAINNET BALANCE AFTER WITHDRAWAL", aliceMainnetLoomCoinStartBal)
	aliceLoomCoinStartBal, err = s.loomCoin.BalanceOf(s.alice)
	require.NoError(err)
	fmt.Println("ALICE DAPPCHAIN BALANCE AFTER WITHDRAWAL", aliceLoomCoinStartBal)

	require.EqualValues(
		tokenAmount.String(),
		new(big.Int).Sub(curDppchainBalance, aliceLoomCoinStartBal).String(),
		"Alice should have loom coin withdrawn the same amount in mainnet transaction")

	require.EqualValues(
		amount,
		aliceMainnetLoomCoinStartBal-curMainnetBalance,
		"Alice should have loom coin different the same amount as withdrawal in mainnet")

}

func (s *BinanceTransferGatewayTestSuite) TestBEP2DepositAndWithdraw() {
	var err error
	require := s.Require()
	require.NoError(err)

	loomCfg, err := ParseConfig([]string{os.Getenv("LOOM_DIR")})
	require.NoError(err)

	// Connect dappchain contracts
	s.dappchainGateway, err = gwclient.ConnectToDAppChainBinanceGateway(s.loomClient, loomCfg.TransferGateway.DAppChainEventsURI)
	require.NoError(err)

	// prevent binance API rate limit
	time.Sleep(60 * time.Second)

	aliceAccount, err := s.aliceDexClient.GetAccount(s.aliceBnbAddress.String())
	require.NoError(err)
	aliceMainnetMoolCoinStartBal := getBEP2TokenBalance(aliceAccount.Balances, MoolToken)
	fmt.Println("ALICE MAINNET BALANCE", aliceMainnetMoolCoinStartBal)
	aliceDappMoolCoinStartBal, err := s.sampleBEP2Token.BalanceOf(s.alice)
	require.NoError(err)
	fmt.Println("ALICE DAPPCHAIN BALANCE", aliceDappMoolCoinStartBal)
	curMainnetBalance := aliceMainnetMoolCoinStartBal
	curDppchainBalance := aliceDappMoolCoinStartBal

	// prevent binance API rate limit
	time.Sleep(5 * time.Second)

	var amount int64 = 10 * 1e8
	var tokenAmount = big.NewInt(amount)

	// Alice deposits to wallet
	payload := []msg.Transfer{
		msg.Transfer{
			ToAddr: s.tokenOwnerBnbAddress,
			Coins: []bnbtypes.Coin{
				bnbtypes.Coin{
					Denom:  MoolToken,
					Amount: amount,
				},
			},
		},
	}
	// make sure we have put loom address in memo
	aliceDappchainAddr := "loom" + s.alice.LoomAddr.Local.Hex()
	result, err := s.aliceDexClient.SendToken(payload, true, transaction.WithMemo(aliceDappchainAddr))
	require.NoError(err)
	fmt.Println("send token tx hash on Binance Dex: ", result.Hash)

	// wait for oracle process
	time.Sleep(30 * time.Second)

	// Alice should now have her MOOL coin in the DAppChain Loom contract
	aliceAccount, err = s.aliceDexClient.GetAccount(s.aliceBnbAddress.String())
	require.NoError(err)
	aliceMainnetMoolCoinStartBal = getBEP2TokenBalance(aliceAccount.Balances, MoolToken)
	fmt.Println("ALICE MAINNET BALANCE AFTER DEPOSIT", aliceMainnetMoolCoinStartBal)
	aliceDappMoolCoinStartBal, err = s.sampleBEP2Token.BalanceOf(s.alice)
	require.NoError(err)
	fmt.Println("ALICE DAPPCHAIN BALANCE AFTER DEPOSIT", aliceDappMoolCoinStartBal)

	require.EqualValues(
		tokenAmount.String(),
		new(big.Int).Sub(aliceDappMoolCoinStartBal, curDppchainBalance).String(),
		"Alice should have MOOL coin deposited the same amount in mainnet transaction")

	require.EqualValues(
		amount,
		curMainnetBalance-aliceMainnetMoolCoinStartBal,
		"Alice should have MOOL coin different the same amount as deposit in mainnet")

	curMainnetBalance = aliceMainnetMoolCoinStartBal
	curDppchainBalance = aliceDappMoolCoinStartBal

	fmt.Printf("Alice is withdrawing %s from DAppChain Gateway...\n", tokenAmount.String())
	// First, Alice need BNB token in the DAppChain so she can use it as transfer fee.
	// Since Alice has no BNB token at the begining, we just transfer some amount from transfer gateway to Alice.
	var feeAmount = big.NewInt(37500)
	toAddr := common.BytesToAddress(s.alice.LoomAddr.Local)
	require.NoError(s.bnbToken.CallEVM("transfer", s.tokenOwner.LoomSigner, toAddr, feeAmount))

	aliceDappBNBCoinStartBal, err := s.bnbToken.BalanceOf(s.alice)
	require.NoError(err)
	fmt.Println("ALICE BNB DAPPCHAIN BALANCE BEFORE WITHDRAWAL", aliceDappBNBCoinStartBal)

	// Alice must grant approval to the DAppChain Gateway to take ownership of the tokens when they're withdrawn
	require.NoError(s.sampleBEP2Token.Approve(s.alice, s.dappchainGateway.Address, tokenAmount))
	// Alice must grant approval to the DAppChain Gateway to take ownership of the bnb token to use as transfer fee
	require.NoError(s.bnbToken.Approve(s.alice, s.dappchainGateway.Address, feeAmount))

	// Now Alice can requests a withdrawal from the DAppChain Gateway...
	// Waiting for oracle to clear a previous receipt
	for i := 0; i < 5; i++ {
		err = s.dappchainGateway.WithdrawBEP2(s.alice, tokenAmount, s.sampleBEP2Token.Address, common.BytesToAddress(s.aliceBnbAddress.Bytes()))
		if err != nil {
			if strings.Contains(err.Error(), "TG003") {
				time.Sleep(5 * time.Second)
			} else {
				require.NoError(err)
			}
		} else {
			break
		}
	}
	require.NoError(err)

	aliceDappBNBCoinStartBal, err = s.bnbToken.BalanceOf(s.alice)
	require.NoError(err)
	fmt.Println("ALICE BNB DAPPCHAIN BALANCE AFTER WITHDRAWAL", aliceDappBNBCoinStartBal)

	require.EqualValues(
		"0",
		aliceDappBNBCoinStartBal.String(),
		"Alice have no BNB token on DappChain after spending it for transfer fee")

	// Let the Oracle fetch pending withdrawals & send trasaction them
	// WARNING: timing issue when tx hash is deleted before this test fetches WithdrawalReceipt
	wr, err := s.dappchainGateway.WithdrawalReceipt(s.alice)
	for i := 0; i < 15; i++ {
		wr, err = s.dappchainGateway.WithdrawalReceipt(s.alice)
		if wr != nil && wr.TxHash != nil {
			break
		}
		fmt.Println("waiting for Oracle to update tx_hash...")
		time.Sleep(10 * time.Second)
	}
	require.NoError(err)
	require.NotNil(wr.TxHash)

	for i := 0; i < 15; i++ {
		wr, err = s.dappchainGateway.WithdrawalReceipt(s.alice)
		if wr == nil {
			break
		}
		fmt.Println("waiting for Oracle to transfer token...")
		time.Sleep(10 * time.Second)
	}
	require.Nil(wr)

	aliceAccount, err = s.aliceDexClient.GetAccount(s.aliceBnbAddress.String())
	require.NoError(err)
	aliceMainnetMoolCoinStartBal = getBEP2TokenBalance(aliceAccount.Balances, MoolToken)
	fmt.Println("ALICE MAINNET BALANCE AFTER WITHDRAWAL", aliceMainnetMoolCoinStartBal)
	aliceDappMoolCoinStartBal, err = s.sampleBEP2Token.BalanceOf(s.alice)
	require.NoError(err)
	fmt.Println("ALICE DAPPCHAIN BALANCE AFTER WITHDRAWAL", aliceDappMoolCoinStartBal)

	require.EqualValues(
		tokenAmount.String(),
		new(big.Int).Sub(curDppchainBalance, aliceDappMoolCoinStartBal).String(),
		"Alice should have loom coin withdrawn the same amount in mainnet transaction")

	require.EqualValues(
		amount,
		aliceMainnetMoolCoinStartBal-curMainnetBalance,
		"Alice should have loom coin different the same amount as withdrawal in mainnet")

}

func (s *BinanceTransferGatewayTestSuite) TestBEP2DepositAndWithdrawOnlyBNB() {
	var err error
	require := s.Require()
	require.NoError(err)

	loomCfg, err := ParseConfig([]string{os.Getenv("LOOM_DIR")})
	require.NoError(err)

	// Connect dappchain contracts
	s.dappchainGateway, err = gwclient.ConnectToDAppChainBinanceGateway(s.loomClient, loomCfg.TransferGateway.DAppChainEventsURI)
	require.NoError(err)

	// prevent binance API rate limit
	time.Sleep(60 * time.Second)

	aliceAccount, err := s.aliceDexClient.GetAccount(s.aliceBnbAddress.String())
	require.NoError(err)
	aliceMainnetBNBStartBal := getBEP2TokenBalance(aliceAccount.Balances, BNBNativeToken)
	fmt.Println("ALICE MAINNET BALANCE", aliceMainnetBNBStartBal)
	aliceDappBNBStartBal, err := s.bnbToken.BalanceOf(s.alice)
	require.NoError(err)
	fmt.Println("ALICE DAPPCHAIN BALANCE", aliceDappBNBStartBal)
	curMainnetBalance := aliceMainnetBNBStartBal
	curDppchainBalance := aliceDappBNBStartBal

	// prevent binance API rate limit
	time.Sleep(10 * time.Second)

	var amount int64 = 0.01 * 1e8
	var tokenAmount = big.NewInt(amount)
	var fee int64 = 37500
	var feeAmount = big.NewInt(fee)

	// Alice deposits to wallet
	payload := []msg.Transfer{
		msg.Transfer{
			ToAddr: s.tokenOwnerBnbAddress,
			Coins: []bnbtypes.Coin{
				bnbtypes.Coin{
					Denom:  BNBNativeToken,
					Amount: amount,
				},
			},
		},
	}
	// make sure we have put loom address in memo
	aliceDappchainAddr := "loom" + s.alice.LoomAddr.Local.Hex()
	result, err := s.aliceDexClient.SendToken(payload, true, transaction.WithMemo(aliceDappchainAddr))
	require.NoError(err)
	fmt.Println("send token tx hash on Binance Dex: ", result.Hash)

	// wait for oracle process
	time.Sleep(60 * time.Second)

	// Alice should now have her BNB coin in the DAppChain Loom contract
	aliceAccount, err = s.aliceDexClient.GetAccount(s.aliceBnbAddress.String())
	require.NoError(err)
	aliceMainnetBNBStartBal = getBEP2TokenBalance(aliceAccount.Balances, BNBNativeToken)
	fmt.Println("ALICE MAINNET BALANCE AFTER DEPOSIT", aliceMainnetBNBStartBal)
	aliceDappBNBStartBal, err = s.bnbToken.BalanceOf(s.alice)
	require.NoError(err)
	fmt.Println("ALICE DAPPCHAIN BALANCE AFTER DEPOSIT", aliceDappBNBStartBal)

	require.EqualValues(
		tokenAmount.String(),
		new(big.Int).Sub(aliceDappBNBStartBal, curDppchainBalance).String(),
		"Alice should have BNB deposited the same amount in mainnet transaction")

	require.EqualValues(
		amount+fee,
		curMainnetBalance-aliceMainnetBNBStartBal,
		"Alice should have BNB different the same amount as deposit in mainnet")

	curMainnetBalance = aliceMainnetBNBStartBal
	curDppchainBalance = aliceDappBNBStartBal

	fmt.Printf("Alice is withdrawing %s from DAppChain Gateway...\n", tokenAmount.String())

	// The expected amount must deduct the transfer fee
	var expectedWithdrawalAmount = big.NewInt(0).Sub(tokenAmount, feeAmount)

	aliceDappBNBStartBal, err = s.bnbToken.BalanceOf(s.alice)
	require.NoError(err)
	fmt.Println("ALICE BNB DAPPCHAIN BALANCE BEFORE WITHDRAWAL", aliceDappBNBStartBal)

	// Alice must grant approval to the DAppChain Gateway to take ownership of the withdrawn tokens
	require.NoError(s.bnbToken.Approve(s.alice, s.dappchainGateway.Address, tokenAmount))

	fmt.Println("Alice is withdrawing from dappchain", expectedWithdrawalAmount.String())
	// Now Alice can requests a withdrawal from the DAppChain Gateway...
	// Waiting for oracle to clear a previous receipt
	for i := 0; i < 5; i++ {
		err = s.dappchainGateway.WithdrawBEP2(s.alice, expectedWithdrawalAmount, s.bnbToken.Address, common.BytesToAddress(s.aliceBnbAddress.Bytes()))
		if err != nil {
			if strings.Contains(err.Error(), "TG003") {
				time.Sleep(5 * time.Second)
			} else {
				require.NoError(err)
			}
		} else {
			break
		}
	}
	require.NoError(err)

	aliceDappBNBStartBal, err = s.bnbToken.BalanceOf(s.alice)
	require.NoError(err)
	fmt.Println("ALICE BNB DAPPCHAIN BALANCE AFTER WITHDRAWAL", aliceDappBNBStartBal)

	require.EqualValues(
		"0",
		aliceDappBNBStartBal.String(),
		"Alice have no BNB token on DappChain after spending it for transfer fee")

	// Let the Oracle fetch pending withdrawals & send trasaction them
	// WARNING: timing issue when tx hash is deleted before this test fetches WithdrawalReceipt
	wr, err := s.dappchainGateway.WithdrawalReceipt(s.alice)
	for i := 0; i < 15; i++ {
		wr, err = s.dappchainGateway.WithdrawalReceipt(s.alice)
		if wr != nil && wr.TxHash != nil {
			break
		}
		fmt.Println("waiting for Oracle to update tx_hash...")
		time.Sleep(10 * time.Second)
	}
	require.NoError(err)
	require.NotNil(wr.TxHash)

	for i := 0; i < 15; i++ {
		wr, err = s.dappchainGateway.WithdrawalReceipt(s.alice)
		if wr == nil {
			break
		}
		fmt.Println("waiting for Oracle to transfer token...")
		time.Sleep(10 * time.Second)
	}
	require.Nil(wr)

	aliceAccount, err = s.aliceDexClient.GetAccount(s.aliceBnbAddress.String())
	require.NoError(err)
	aliceMainnetBNBStartBal = getBEP2TokenBalance(aliceAccount.Balances, BNBNativeToken)
	fmt.Println("ALICE MAINNET BALANCE AFTER WITHDRAWAL", aliceMainnetBNBStartBal)
	aliceDappBNBStartBal, err = s.bnbToken.BalanceOf(s.alice)
	require.NoError(err)
	fmt.Println("ALICE DAPPCHAIN BALANCE AFTER WITHDRAWAL", aliceDappBNBStartBal)

	require.EqualValues(
		tokenAmount.String(),
		new(big.Int).Sub(curDppchainBalance, aliceDappBNBStartBal).String(),
		"Alice should have BNB withdrawn the same amount in mainnet transaction")

	require.EqualValues(
		amount-fee,
		aliceMainnetBNBStartBal-curMainnetBalance,
		"Alice should have BNB different the same amount as withdrawal in mainnet")

}

func getBEP2TokenBalance(balances []bnbtypes.TokenBalance, name string) int64 {
	for _, balance := range balances {
		if balance.Symbol == name {
			return balance.Free.ToInt64()
		}
	}
	return 0
}
