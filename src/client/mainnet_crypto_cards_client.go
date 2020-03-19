package client

import (
	"context"
	"ethcontract"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/loomnetwork/go-loom/client"
)

type MainnetCryptoCardsClient struct {
	contract  *ethcontract.MainnetCryptoCardsContract
	ethClient *ethclient.Client

	TxTimeout time.Duration
	Address   common.Address
	TxHash    string
}

func (c *MainnetCryptoCardsClient) MintTokens(contractOwner *client.Identity, recipient *client.Identity) error {
	tx, err := c.contract.MintTokens(client.DefaultTransactOptsForIdentity(contractOwner), recipient.MainnetAddr)
	if err != nil {
		return err
	}
	return client.WaitForTxConfirmation(context.TODO(), c.ethClient, tx, c.TxTimeout)
}

func (c *MainnetCryptoCardsClient) DepositToGateway(caller *client.Identity, tokenID *big.Int) error {
	tx, err := c.contract.DepositToGateway(client.DefaultTransactOptsForIdentity(caller), tokenID)
	if err != nil {
		return err
	}
	return client.WaitForTxConfirmation(context.TODO(), c.ethClient, tx, c.TxTimeout)
}

func (c *MainnetCryptoCardsClient) BalanceOf(caller *client.Identity) (uint64, error) {
	bal, err := c.contract.BalanceOf(nil, caller.MainnetAddr)
	if err != nil {
		return 0, err
	}
	return bal.Uint64(), nil
}

func (c *MainnetCryptoCardsClient) TokenOfOwnerByIndex(caller *client.Identity, index int) (*big.Int, error) {
	tokenID, err := c.contract.TokenOfOwnerByIndex(nil, caller.MainnetAddr, new(big.Int).SetInt64(int64(index)))
	if err != nil {
		return nil, err
	}
	return tokenID, nil
}

func (c *MainnetCryptoCardsClient) OwnerOf(tokenID *big.Int) (common.Address, error) {
	return c.contract.OwnerOf(nil, tokenID)
}

func ConnectToMainnetCards(ethClient *ethclient.Client, address string) (*MainnetCryptoCardsClient, error) {
	contractAddr := common.HexToAddress(address)
	contract, err := ethcontract.NewMainnetCryptoCardsContract(contractAddr, ethClient)
	if err != nil {
		return nil, err
	}
	return &MainnetCryptoCardsClient{
		contract:  contract,
		ethClient: ethClient,
		Address:   contractAddr,
	}, nil
}

func DeployMainnetCardsContract(
	ethClient *ethclient.Client, creator *client.Identity, gatewayAddr common.Address,
) (*MainnetCryptoCardsClient, error) {
	addr, tx, contract, err := ethcontract.DeployMainnetCryptoCardsContract(
		client.DefaultTransactOptsForIdentity(creator),
		ethClient,
		gatewayAddr,
	)
	if err != nil {
		return nil, err
	}
	if err := client.WaitForTxConfirmation(context.TODO(), ethClient, tx, 0); err != nil {
		return nil, err
	}
	return &MainnetCryptoCardsClient{
		contract:  contract,
		ethClient: ethClient,
		Address:   addr,
		TxHash:    tx.Hash().Hex(),
	}, nil
}
