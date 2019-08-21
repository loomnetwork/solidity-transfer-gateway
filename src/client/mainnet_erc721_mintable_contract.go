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

type MainnetERC721MintableContract struct {
	contract  *ethcontract.SampleERC721MintableToken
	ethClient *ethclient.Client

	TxTimeout time.Duration
	Address   common.Address
	TxHash    string
}

func (c *MainnetERC721MintableContract) Mint(from *client.Identity, to common.Address, tokenID *big.Int) error {
	tx, err := c.contract.Mint(client.DefaultTransactOptsForIdentity(from), to, tokenID)
	if err != nil {
		return err
	}
	return client.WaitForTxConfirmation(context.TODO(), c.ethClient, tx, c.TxTimeout)
}

func (c *MainnetERC721MintableContract) MintTo(from *client.Identity, to common.Address, tokenID *big.Int) error {
	tx, err := c.contract.MintTo(client.DefaultTransactOptsForIdentity(from), to, tokenID)
	if err != nil {
		return err
	}
	return client.WaitForTxConfirmation(context.TODO(), c.ethClient, tx, c.TxTimeout)
}

func (c *MainnetERC721MintableContract) BalanceOf(caller *client.Identity) (*big.Int, error) {
	bal, err := c.contract.BalanceOf(nil, caller.MainnetAddr)
	if err != nil {
		return nil, err
	}
	return bal, nil
}

func (c *MainnetERC721MintableContract) OwnerOf(tokenID *big.Int) (common.Address, error) {
	return c.contract.OwnerOf(nil, tokenID)
}

func ConnectToMainnetERC721MintableContract(ethClient *ethclient.Client, address string) (*MainnetERC721MintableContract, error) {
	contractAddr := common.HexToAddress(address)
	contract, err := ethcontract.NewSampleERC721MintableToken(contractAddr, ethClient)
	if err != nil {
		return nil, err
	}
	return &MainnetERC721MintableContract{
		contract:  contract,
		ethClient: ethClient,
		Address:   contractAddr,
	}, nil
}

func DeployMainnetERC721MintableContract(
	ethClient *ethclient.Client, creator *client.Identity, gatewayAddr common.Address,
) (*MainnetERC721MintableContract, error) {
	addr, tx, contract, err := ethcontract.DeploySampleERC721MintableToken(
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
	return &MainnetERC721MintableContract{
		contract:  contract,
		ethClient: ethClient,
		Address:   addr,
		TxHash:    tx.Hash().Hex(),
	}, nil
}
