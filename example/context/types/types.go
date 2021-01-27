package exampletypes

import (
	"io/ioutil"
	"path"
	"runtime"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
)

type Config struct {
	NodeURL                string
	BlockHash              types.Hash
	TransactionHash        types.Hash
	BlockHashOfNewContract types.Hash
	ERC20Address           types.Address
	client                 *sdk.Client
	retryClient            *sdk.Client
	accountManager         sdk.AccountManagerOperator
}

func (c *Config) SetAccountManager(am sdk.AccountManagerOperator) {
	c.accountManager = am
}

func (c *Config) GetAccountManager() sdk.AccountManagerOperator {
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

func (c *Config) GetErc20Contract() (*sdk.Contract, error) {
	currentDir := getCurrentDir()

	abiPath := path.Join(currentDir, "../contract/erc20.abi")

	abi, err := ioutil.ReadFile(abiPath)
	if err != nil {
		panic(err)
	}

	contract, err := c.GetClient().GetContract([]byte(abi), &c.ERC20Address)
	return contract, err
}

func getCurrentDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("get current file path error")
	}
	currentDir := path.Join(filename, "../")
	return currentDir
}
