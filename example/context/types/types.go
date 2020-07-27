package exampletypes

import (
	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
)

type Config struct {
	NodeURL         string
	BlockHash       types.Hash
	TransactionHash types.Hash
	ERC20Address    types.Address
	client          *sdk.Client
	retryClient     *sdk.Client
	accountManager  *sdk.AccountManager
}

func (c *Config) SetAccountManager(am *sdk.AccountManager) {
	c.accountManager = am
}

func (c *Config) GetAccountManager() *sdk.AccountManager {
	return c.accountManager
}

func (c *Config) SetClient(client *sdk.Client) {
	c.client = client
}

func (c *Config) GetClient() *sdk.Client {
	return c.client
}

func (c *Config) SetRetryClient(client *sdk.Client) {
	c.retryClient = client
}

func (c *Config) GetRetryClient() *sdk.Client {
	return c.retryClient
}
