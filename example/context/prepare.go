package context

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/BurntSushi/toml"
	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	exampletypes "github.com/Conflux-Chain/go-conflux-sdk/example/context/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var (
	config     exampletypes.Config
	client     *sdk.Client
	currentDir string
	configPath string
	// am             *sdk.AccountManager
	defaultAccount *types.Address
	nextNonce      *hexutil.Big
)

// PrepareForClientExample ...
func PrepareForClientExample() *exampletypes.Config {
	fmt.Printf("\n=======start prepare config===========\n")
	getConfig()
	initClient()
	generateBlockHashAndTxHash()
	deployContract(false)
	saveConfig()
	fmt.Printf("\n=======prepare config done!===========\n")
	return &config
}

// PrepareForContractExample ...
func PrepareForContractExample() *exampletypes.Config {
	fmt.Printf("\n=======start prepare config===========\n")
	getConfig()
	initClient()
	saveConfig()
	fmt.Print("\n=======prepare config done!===========\n")
	return &config
}

func getConfig() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("get current file path error")
	}
	currentDir = path.Join(filename, "../")
	configPath = path.Join(currentDir, "config.toml")
	// cp := make(map[string]string)
	config = exampletypes.Config{}
	_, err := toml.DecodeFile(configPath, &config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("- to get config done: %+v\n", JSONFmt(config))
}

func initClient() {
	// url := "http://testnet-jsonrpc.conflux-chain.org:12537"

	// init client
	var err error

	keyStorePath := path.Join(currentDir, "keystore")
	client, err = sdk.NewClient(config.NodeURL, sdk.ClientOption{KeystorePath: keyStorePath})
	if err != nil {
		panic(err)
	}

	config.SetClient(client)

	// init retry client
	option := sdk.ClientOption{
		KeystorePath:  keyStorePath,
		RetryCount:    10,
		RetryInterval: time.Second,
		// CallRpcLog:      types.DefaultCallRPCLog,
		// BatchCallRPCLog: types.DefaultBatchCallRPCLog,
	}
	retryclient, err := sdk.NewClient(config.NodeURL, option)
	if err != nil {
		panic(err)
	}
	config.SetRetryClient(retryclient)

	defaultAccount, err = client.AccountManager.GetDefault()
	if err != nil {
		panic(err)
	}
	client.AccountManager.UnlockDefault("hello")
	config.SetAccountManager(client.AccountManager)

	nextNonce, err = client.GetNextNonce(*defaultAccount, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("- to init client done")
}

func generateBlockHashAndTxHash() {

	block, err1 := client.GetBlockByHash(config.BlockHash)
	tx, err2 := client.GetTransactionByHash(config.TransactionHash)
	if block == nil || err1 != nil || tx == nil || err2 != nil {
		utx, err := client.CreateUnsignedTransaction(*defaultAccount, cfxaddress.MustNewFromHex("0x10f4bcf113e0b896d9b34294fd3da86b4adf0302"), types.NewBigInt(1), nil)
		if err != nil {
			panic(err)
		}
		utx.Nonce = GetNextNonceAndIncrease()
		txhash, err := client.SendTransaction(utx)
		if err != nil {
			panic(err)
		}
		config.TransactionHash = txhash

		WaitPacked(client, txhash)

		tx, err := client.GetTransactionByHash(txhash)
		if err != nil {
			panic(err)
		}
		config.BlockHash = *tx.BlockHash
	}

	fmt.Println("- gen txhash done")
}

func deployContract(force bool) *sdk.Contract {
	// check erc20 and erc777 address, if len !==42 or getcode error, deploy
	erc20Contract, txhash := DeployIfNotExist(config.ERC20Address, path.Join(currentDir, "contract/erc20.abi"), path.Join(currentDir, "contract/erc20.bytecode"), force)
	if erc20Contract != nil {
		config.ERC20Address = *erc20Contract.Address
	}
	if txhash != nil {
		receipt := WaitPacked(client, *txhash)
		config.BlockHashOfNewContract = receipt.BlockHash
	}

	fmt.Println("- to deploy contracts if not exist done")
	return erc20Contract
}

func saveConfig() {
	f, err := os.OpenFile(configPath, os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	encoder := toml.NewEncoder(f)
	err = encoder.Encode(config)
	if err != nil {
		panic(err)
	}
	fmt.Println("- to save config done")
}
