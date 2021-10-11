package exampletypes

import (
	"io/ioutil"
	"path"
	"runtime"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/cfxclient"
	"github.com/Conflux-Chain/go-conflux-sdk/contracts"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
)

type Config struct {
	NodeURL                string
	BlockHash              types.Hash
	TransactionHash        types.Hash
	BlockHashOfNewContract types.Hash
	ERC20Address           types.Address
	client                 *cfxclient.SignableClient
	retryClient            *cfxclient.SignableClient
	accountManager         sdk.Wallet
}

func (c *Config) SetWallet(am sdk.Wallet) {
	c.accountManager = am
}

func (c *Config) GetAccountManager() sdk.Wallet {
	return c.accountManager
}

func (c *Config) SetClient(client *cfxclient.SignableClient) {
	c.client = client
}

func (c *Config) GetClient() *cfxclient.SignableClient {
	return c.client
}

func (c *Config) SetRetryClient(client *cfxclient.SignableClient) {
	c.retryClient = client
}

func (c *Config) GetRetryClient() *cfxclient.SignableClient {
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
