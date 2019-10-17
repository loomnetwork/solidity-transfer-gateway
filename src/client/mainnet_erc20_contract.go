package client

import (
	"context"
	"ethcontract"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/loomnetwork/go-loom/client"
)

type MainnetERC20Contract struct {
	contract  *ethcontract.MainnetGameTokenContract
	ethClient *ethclient.Client

	TxTimeout time.Duration
	Address   common.Address
	TxHash    string
}

func (c *MainnetERC20Contract) BalanceOf(caller *client.Identity) (*big.Int, error) {
	bal, err := c.contract.BalanceOf(nil, caller.MainnetAddr)
	if err != nil {
		return nil, err
	}
	return bal, nil
}

func (c *MainnetERC20Contract) TransferFrom(to *client.Identity, from *client.Identity, amount *big.Int) error {
	tx, err := c.contract.TransferFrom(client.DefaultTransactOptsForIdentity(from), from.MainnetAddr, to.MainnetAddr, amount)
	if err != nil {
		return err
	}
	return client.WaitForTxConfirmation(context.TODO(), c.ethClient, tx, c.TxTimeout)
}

func (c *MainnetERC20Contract) Approve(from *client.Identity, to common.Address, amount *big.Int) error {
	tx, err := c.contract.Approve(client.DefaultTransactOptsForIdentity(from), to, amount)
	if err != nil {
		return err
	}
	return client.WaitForTxConfirmation(context.TODO(), c.ethClient, tx, c.TxTimeout)
}

func (c *MainnetERC20Contract) Transfer(to *client.Identity, from *client.Identity, amount *big.Int) error {
	tx, err := c.contract.Transfer(client.DefaultTransactOptsForIdentity(from), to.MainnetAddr, amount)
	if err != nil {
		return err
	}
	return client.WaitForTxConfirmation(context.TODO(), c.ethClient, tx, c.TxTimeout)
}

// TransferTx calls  Transfer and waits for it to complete. It returns tx and error.
func (c *MainnetERC20Contract) TransferTx(caller *client.Identity, to common.Address, amount *big.Int) (*ethtypes.Transaction, error) {
	tx, err := c.contract.Transfer(client.DefaultTransactOptsForIdentity(caller), to, amount)
	if err != nil {
		return nil, err
	}
	err = client.WaitForTxConfirmation(context.TODO(), c.ethClient, tx, c.TxTimeout)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func ConnectToMainnetERC20Contract(ethClient *ethclient.Client, address string) (*MainnetERC20Contract, error) {
	contractAddr := common.HexToAddress(address)
	contract, err := ethcontract.NewMainnetGameTokenContract(contractAddr, ethClient)
	if err != nil {
		return nil, err
	}
	return &MainnetERC20Contract{
		contract:  contract,
		ethClient: ethClient,
		Address:   contractAddr,
	}, nil
}

func DeployMainnetERC20Contract(
	ethClient *ethclient.Client, creator *client.Identity, gatewayAddr common.Address,
) (*MainnetERC20Contract, error) {
	addr, tx, contract, err := ethcontract.DeployMainnetGameTokenContract(
		client.DefaultTransactOptsForIdentity(creator),
		ethClient,
	)
	if err != nil {
		return nil, err
	}
	if err := client.WaitForTxConfirmation(context.TODO(), ethClient, tx, 0); err != nil {
		return nil, err
	}
	return &MainnetERC20Contract{
		contract:  contract,
		ethClient: ethClient,
		Address:   addr,
		TxHash:    tx.Hash().Hex(),
	}, nil
}
