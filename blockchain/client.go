package blockchain

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"tx-bot/config"
)

type Client struct {
	client     *ethclient.Client
	config     *config.Config
	privateKey *ecdsa.PrivateKey
	auth       *bind.TransactOpts
}

func NewClient(cfg *config.Config) (*Client, error) {
	client, err := ethclient.Dial(cfg.Web3ProviderURL)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(cfg.ChainID))
	if err != nil {
		return nil, err
	}

	auth.GasLimit = cfg.GasLimit
	auth.GasPrice = big.NewInt(cfg.GasPriceGwei * 1e9) // Convert Gwei to Wei

	return &Client{
		client:     client,
		config:     cfg,
		privateKey: privateKey,
		auth:       auth,
	}, nil
}

func (c *Client) SendTransaction(ctx context.Context) error {
	// Generate random amount between min and max transfer amount
	amount := c.config.MinTransferAmount + rand.Float64()*(c.config.MaxTransferAmount-c.config.MinTransferAmount)
	
	// Convert amount to Wei using big.Float for precise calculation
	amountFloat := new(big.Float).SetFloat64(amount)
	weiMultiplier := new(big.Float).SetFloat64(1e18)
	amountFloat.Mul(amountFloat, weiMultiplier)
	
	// Convert to big.Int
	amountWei := new(big.Int)
	amountFloat.Int(amountWei)

	// Get the nonce
	nonce, err := c.client.PendingNonceAt(ctx, c.auth.From)
	if err != nil {
		return err
	}

	// Create transaction
	tx := types.NewTransaction(
		nonce,
		c.auth.From, // Send to self for now
		amountWei,
		c.config.GasLimit,
		big.NewInt(c.config.GasPriceGwei * 1e9),
		nil,
	)

	// Sign the transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(c.config.ChainID)), c.privateKey)
	if err != nil {
		return err
	}

	// Send the transaction
	return c.client.SendTransaction(ctx, signedTx)
}

func (c *Client) GetRandomSleepTime() time.Duration {
	// Random sleep time between min and max sleep time
	sleepTime := c.config.MinSleepTime + rand.Intn(c.config.MaxSleepTime-c.config.MinSleepTime)
	return time.Duration(sleepTime) * time.Second
}

func (c *Client) GetRandomLongSleepTime() time.Duration {
	// Random long sleep time between min and max long sleep time
	sleepTime := c.config.MinLongSleepTime + rand.Intn(c.config.MaxLongSleepTime-c.config.MinLongSleepTime)
	return time.Duration(sleepTime) * time.Second
} 