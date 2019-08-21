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

type MainnetERC721XContract struct {
	contract  *ethcontract.MainnetERC721XCardsContract
	ethClient *ethclient.Client

	TxTimeout time.Duration
	Address   common.Address
	TxHash    string
}

func (c *MainnetERC721XContract) MintTokens(contractOwner *client.Identity, tokenID *big.Int, amount *big.Int, recipient *client.Identity) error {
	tx, err := c.contract.MintTokens(
		client.DefaultTransactOptsForIdentity(contractOwner),
		recipient.MainnetAddr, tokenID, amount,
	)
	if err != nil {
		return err
	}
	return client.WaitForTxConfirmation(context.TODO(), c.ethClient, tx, c.TxTimeout)
}

func (c *MainnetERC721XContract) DepositToGateway(caller *client.Identity, tokenID, amount *big.Int) error {
	tx, err := c.contract.DepositToGateway(client.DefaultTransactOptsForIdentity(caller), tokenID, amount)
	if err != nil {
		return err
	}
	return client.WaitForTxConfirmation(context.TODO(), c.ethClient, tx, c.TxTimeout)
}

func (c *MainnetERC721XContract) BalanceOf(caller *client.Identity, tokenID *big.Int) (*big.Int, error) {
	bal, err := c.contract.BalanceOfToken(nil, caller.MainnetAddr, tokenID)
	if err != nil {
		return nil, err
	}
	return bal, nil
}

func (c *MainnetERC721XContract) TokenOfOwnerByIndex(caller *client.Identity, index int) (*big.Int, error) {
	tokenID, err := c.contract.TokenOfOwnerByIndex(nil, caller.MainnetAddr, new(big.Int).SetInt64(int64(index)))
	if err != nil {
		return nil, err
	}
	return tokenID, nil
}

func (c *MainnetERC721XContract) OwnerOf(tokenID *big.Int) (common.Address, error) {
	return c.contract.OwnerOf(nil, tokenID)
}

func ConnectToMainnetERC721XContract(ethClient *ethclient.Client, address string) (*MainnetERC721XContract, error) {
	contractAddr := common.HexToAddress(address)
	contract, err := ethcontract.NewMainnetERC721XCardsContract(contractAddr, ethClient)
	if err != nil {
		return nil, err
	}
	return &MainnetERC721XContract{
		contract:  contract,
		ethClient: ethClient,
		Address:   contractAddr,
	}, nil
}

func DeployMainnetERC721XContract(
	ethClient *ethclient.Client, creator *client.Identity, gatewayAddr common.Address,
) (*MainnetERC721XContract, error) {
	addr, tx, contract, err := ethcontract.DeployMainnetERC721XCardsContract(
		client.DefaultTransactOptsForIdentity(creator),
		ethClient,
		gatewayAddr,
		"baseTokenURI",
	)
	if err != nil {
		return nil, err
	}
	if err := client.WaitForTxConfirmation(context.TODO(), ethClient, tx, 0); err != nil {
		return nil, err
	}
	return &MainnetERC721XContract{
		contract:  contract,
		ethClient: ethClient,
		Address:   addr,
		TxHash:    tx.Hash().Hex(),
	}, nil
}
