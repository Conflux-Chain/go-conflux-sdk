package exampletypes

import (
	"io/ioutil"
	"path"
	"runtime"

	client "github.com/Conflux-Chain/go-conflux-sdk/cfxclient"
	"github.com/Conflux-Chain/go-conflux-sdk/contracts"
	"github.com/Conflux-Chain/go-conflux-sdk/interfaces"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
)

type Config struct {
	NodeURL                string
	BlockHash              types.Hash
	TransactionHash        types.Hash
	BlockHashOfNewContract types.Hash
	ERC20Address           types.Address
	client                 *client.SignableClient
	retryClient            *client.SignableClient
	accountManager         interfaces.Wallet
}

func (c *Config) SetWallet(am interfaces.Wallet) {
	c.accountManager = am
}

func (c *Config) GetAccountManager() interfaces.Wallet {
	return c.accountManager
}

func (c *Config) SetClient(_client *client.SignableClient) {
	c.client = _client
}

func (c *Config) GetClient() *client.SignableClient {
	return c.client
}

func (c *Config) SetRetryClient(_client *client.SignableClient) {
	c.retryClient = _client
}

func (c *Config) GetRetryClient() *client.SignableClient {
	return c.retryClient
}

func (c *Config) GetErc20Contract() (*contracts.Contract, error) {
	currentDir := getCurrentDir()

	abiPath := path.Join(currentDir, "../contract/erc20.abi")

	abi, err := ioutil.ReadFile(abiPath)
	if err != nil {
		panic(err)
	}

	contract, err := contracts.NewContract(c.client, []byte(abi), &c.ERC20Address)
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
