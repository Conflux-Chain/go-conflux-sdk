---
id: go_sdk
title: Golang SDK
custom_edit_url: https://github.com/Conflux-Chain/go-conflux-sdk/edit/master/api.md
keywords:
  - conflux
  - go
  - sdk
---
# API Reference
## Getting Started
The go-conflux-sdk module is a collection of packages which contain specific functionality for the conflux ecosystem.

- The package `sdk` is for interacting with conflux chain, account manager and operating smart contracts
- The package `utils` contains useful helper functions for Dapp developers.

## Install
You can get Conflux Golang API directly or use go module as below
```
go get github.com/Conflux-Chain/go-conflux-sdk
```
You can also add the Conflux Golang API into vendor folder.
```
govendor fetch github.com/Conflux-Chain/go-conflux-sdk
```

After that you need to create a client instance with node url and an account manager instance.
```go
//create account manager and unlock account
am := sdk.NewAccountManager("./keystore")
err := am.TimedUnlockDefault("password", 30 * time.Second)
if err != nil {
	panic(err)
}

//init client
client, err := sdk.NewClient("http://testnet-jsonrpc.conflux-chain.org:12537")
if err != nil {
	panic(err)
}
client.SetAccountManager(am)
```
## package sdk

 import "github.com/Conflux-Chain/go-conflux-sdk"

### type AccountManager

```go
type AccountManager struct {
}
```

AccountManager manages Conflux accounts.

#### func  NewAccountManager

```go
func NewAccountManager(keydir string) *AccountManager
```
NewAccountManager creates an instance of AccountManager based on the keystore
directory "keydir".

#### func (*AccountManager) Create

```go
func (m *AccountManager) Create(passphrase string) (types.Address, error)
```
Create creates a new account and puts the keystore file into keystore directory

#### func (*AccountManager) Delete

```go
func (m *AccountManager) Delete(address types.Address, passphrase string) error
```
Delete deletes the specified account and remove the keystore file from keystore
directory.

#### func (*AccountManager) GetDefault

```go
func (m *AccountManager) GetDefault() (*types.Address, error)
```
GetDefault return first account in keystore directory

#### func (*AccountManager) Import

```go
func (m *AccountManager) Import(keyFile, passphrase, newPassphrase string) (types.Address, error)
```
Import imports account from external key file to keystore directory. Returns
error if the account already exists.

#### func (*AccountManager) List

```go
func (m *AccountManager) List() []types.Address
```
List lists all accounts in keystore directory.

#### func (*AccountManager) Lock

```go
func (m *AccountManager) Lock(address types.Address) error
```
Lock locks the specified account.

#### func (*AccountManager) Sign

```go
func (m *AccountManager) Sign(tx types.UnsignedTransaction, passphrase string) (v byte, r, s []byte, err error)
```
Sign signs tx by passphrase and returns the signature

#### func (*AccountManager) SignAndEcodeTransactionWithPassphrase

```go
func (m *AccountManager) SignAndEcodeTransactionWithPassphrase(tx types.UnsignedTransaction, passphrase string) ([]byte, error)
```
SignAndEcodeTransactionWithPassphrase signs tx with given passphrase and return
its RLP encoded data.

#### func (*AccountManager) SignTransaction

```go
func (m *AccountManager) SignTransaction(tx types.UnsignedTransaction) ([]byte, error)
```
SignTransaction signs tx and returns its RLP encoded data.

#### func (*AccountManager) SignTransactionWithPassphrase

```go
func (m *AccountManager) SignTransactionWithPassphrase(tx types.UnsignedTransaction, passphrase string) (*types.SignedTransaction, error)
```
SignTransactionWithPassphrase signs tx with given passphrase and returns a
transction with signature

#### func (*AccountManager) TimedUnlock

```go
func (m *AccountManager) TimedUnlock(address types.Address, passphrase string, timeout time.Duration) error
```
TimedUnlock unlocks the specified account for a period of time.

#### func (*AccountManager) TimedUnlockDefault

```go
func (m *AccountManager) TimedUnlockDefault(passphrase string, timeout time.Duration) error
```
TimedUnlockDefault unlocks the specified account for a period of time.

#### func (*AccountManager) Unlock

```go
func (m *AccountManager) Unlock(address types.Address, passphrase string) error
```
Unlock unlocks the specified account indefinitely.

#### func (*AccountManager) UnlockDefault

```go
func (m *AccountManager) UnlockDefault(passphrase string) error
```
UnlockDefault unlocks the default account indefinitely.

#### func (*AccountManager) Update

```go
func (m *AccountManager) Update(address types.Address, passphrase, newPassphrase string) error
```
Update updates the passphrase of specified account.

### type Client

```go
type Client struct {
}
```

Client represents a client to interact with Conflux blockchain.

#### func  NewClient

```go
func NewClient(nodeURL string) (*Client, error)
```
NewClient creates a new instance of Client with specified conflux node url.

#### func (*Client) ApplyUnsignedTransactionDefault

```go
func (c *Client) ApplyUnsignedTransactionDefault(tx *types.UnsignedTransaction) error
```
ApplyUnsignedTransactionDefault set empty fields to value fetched from conflux
node.

#### func (*Client) Call

```go
func (c *Client) Call(request types.CallRequest, epoch *types.Epoch) (*string, error)
```
Call executes a message call transaction "request" at specified epoch, which is
directly executed in the VM of the node, but never mined into the block chain
and returns the contract execution result.

#### func (*Client) CallRPC

```go
func (c *Client) CallRPC(result interface{}, method string, args ...interface{}) error
```
CallRPC performs a JSON-RPC call with the given arguments and unmarshals into
result if no error occurred.

The result must be a pointer so that package json can unmarshal into it. You can
also pass nil, in which case the result is ignored.

#### func (*Client) Close

```go
func (c *Client) Close()
```
Close closes the client, aborting any in-flight requests.

#### func (*Client) CreateUnsignedTransaction

```go
func (c *Client) CreateUnsignedTransaction(from types.Address, to types.Address, amount *hexutil.Big, data *[]byte) (*types.UnsignedTransaction, error)
```
CreateUnsignedTransaction creates an unsigned transaction by parameters, and the
other fields will be set to values fetched from conflux node.

#### func (*Client) Debug

```go
func (c *Client) Debug(method string, args ...interface{}) (interface{}, error)
```
Debug calls the Conflux debug API.

#### func (*Client) DeployContract

```go
func (c *Client) DeployContract(abiJSON string, bytecode []byte, option *types.ContractDeployOption, timeout time.Duration, callback func(deployedContract Contractor, hash *types.Hash, err error)) <-chan struct{}
```
DeployContract deploys a contract Function A deploys a contract synchronously by
abiJSON, bytecode and option. It returns a channel for notifying when deploy
completed. And the callback for handling the deploy result.

#### func (*Client) EstimateGasAndCollateral

```go
func (c *Client) EstimateGasAndCollateral(request types.CallRequest) (*types.Estimate, error)
```
EstimateGasAndCollateral excutes a message call "request" and returns the amount
of the gas used and storage for collateral

#### func (*Client) GetBalance

```go
func (c *Client) GetBalance(address types.Address, epoch ...*types.Epoch) (*big.Int, error)
```
GetBalance returns the balance of specified address at epoch.

#### func (*Client) GetBestBlockHash

```go
func (c *Client) GetBestBlockHash() (types.Hash, error)
```
GetBestBlockHash returns the current best block hash.

#### func (*Client) GetBlockByEpoch

```go
func (c *Client) GetBlockByEpoch(epoch *types.Epoch) (*types.Block, error)
```
GetBlockByEpoch returns the block of specified epoch. If the epoch is invalid,
return the concrete error.

#### func (*Client) GetBlockByHash

```go
func (c *Client) GetBlockByHash(blockHash types.Hash) (*types.Block, error)
```
GetBlockByHash returns the block of specified blockHash If the block is not
found, return nil.

#### func (*Client) GetBlockConfirmRiskByHash

```go
func (c *Client) GetBlockConfirmRiskByHash(blockHash types.Hash) (*big.Int, error)
```
GetBlockConfirmRiskByHash indicates the risk coefficient that the pivot block of
the epoch where the block is located becomes an normal block.

#### func (*Client) GetBlockRevertRateByHash

```go
func (c *Client) GetBlockRevertRateByHash(blockHash types.Hash) (*big.Float, error)
```
GetBlockRevertRateByHash indicates the probability that the pivot block of the
epoch where the block is located becomes an ordinary block.

it's (confirm risk coefficient/ (2^256-1))

#### func (*Client) GetBlockSummaryByEpoch

```go
func (c *Client) GetBlockSummaryByEpoch(epoch *types.Epoch) (*types.BlockSummary, error)
```
GetBlockSummaryByEpoch returns the block summary of specified epoch. If the
epoch is invalid, return the concrete error.

#### func (*Client) GetBlockSummaryByHash

```go
func (c *Client) GetBlockSummaryByHash(blockHash types.Hash) (*types.BlockSummary, error)
```
GetBlockSummaryByHash returns the block summary of specified blockHash If the
block is not found, return nil.

#### func (*Client) GetBlocksByEpoch

```go
func (c *Client) GetBlocksByEpoch(epoch *types.Epoch) ([]types.Hash, error)
```
GetBlocksByEpoch returns the blocks hash in the specified epoch.

#### func (*Client) GetCode

```go
func (c *Client) GetCode(address types.Address, epoch ...*types.Epoch) (string, error)
```
GetCode returns the bytecode in HEX format of specified address at epoch.

#### func (*Client) GetContract

```go
func (c *Client) GetContract(abiJSON string, deployedAt *types.Address) (*Contract, error)
```
GetContract creates a contract instance according to abi json and it's deployed
address

#### func (*Client) GetEpochNumber

```go
func (c *Client) GetEpochNumber(epoch ...*types.Epoch) (*big.Int, error)
```
GetEpochNumber returns the highest or specified epoch number.

#### func (*Client) GetGasPrice

```go
func (c *Client) GetGasPrice() (*big.Int, error)
```
GetGasPrice returns the recent mean gas price.

#### func (*Client) GetLogs

```go
func (c *Client) GetLogs(filter types.LogFilter) ([]types.Log, error)
```
GetLogs returns logs that matching the specified filter.

#### func (*Client) GetNextNonce

```go
func (c *Client) GetNextNonce(address types.Address) (uint64, error)
```
GetNextNonce returns the next transaction nonce of address

#### func (*Client) GetTransactionByHash

```go
func (c *Client) GetTransactionByHash(txHash types.Hash) (*types.Transaction, error)
```
GetTransactionByHash returns transaction for the specified txHash. If the
transaction is not found, return nil.

#### func (*Client) GetTransactionReceipt

```go
func (c *Client) GetTransactionReceipt(txHash types.Hash) (*types.TransactionReceipt, error)
```
GetTransactionReceipt returns the receipt of specified transaction hash. If no
receipt is found, return nil.

#### func (*Client) SendRawTransaction

```go
func (c *Client) SendRawTransaction(rawData []byte) (types.Hash, error)
```
SendRawTransaction sends signed transaction and returns its hash.

#### func (*Client) SendTransaction

```go
func (c *Client) SendTransaction(tx *types.UnsignedTransaction) (types.Hash, error)
```
SendTransaction signs and sends transaction to conflux node and returns the
transaction hash.

#### func (*Client) SetAccountManager

```go
func (c *Client) SetAccountManager(accountManager *AccountManager)
```
SetAccountManager sets account manager for sign transaction

#### func (*Client) SignEncodedTransactionAndSend

```go
func (c *Client) SignEncodedTransactionAndSend(encodedTx []byte, v byte, r, s []byte) (*types.Transaction, error)
```
SignEncodedTransactionAndSend signs RLP encoded transaction "encodedTx" by
signature "r,s,v" and sends it to node, and returns responsed transaction.

### type Contract

```go
type Contract struct {
        ABI     abi.ABI
        Client  ClientOperator
        Address *types.Address
}
```

Contract represents a smart contract

#### func (*Contract) Call

```go
func (c *Contract) Call(option *types.ContractMethodCallOption, resultPtr interface{}, method string, args ...interface{}) error
```
Call calls to the contract method with args and fills the excuted result to the
"resultPtr".

the resultPtr should be a pointer of the method output struct type.

#### func (*Contract) GetData

```go
func (c *Contract) GetData(method string, args ...interface{}) (*[]byte, error)
```
GetData packs the given method name to conform the ABI of the contract "c".
Method call's data will consist of method_id, args0, arg1, ... argN. Method id
consists of 4 bytes and arguments are all 32 bytes. Method ids are created from
the first 4 bytes of the hash of the methods string signature. (signature =
baz(uint32,string32))

#### func (*Contract) SendTransaction

```go
func (c *Contract) SendTransaction(option *types.ContractMethodSendOption, method string, args ...interface{}) (*types.Hash, error)
```
SendTransaction sends a transaction to the contract method with args and returns
its transaction hash

## package utils

    import "github.com/Conflux-Chain/go-conflux-sdk/utils"


#### func  Keccak256

```go
func Keccak256(hexStr string) (string, error)
```
Keccak256 hashs hex string by keccak256 and returns it's hash value

#### func  PrivateKeyToPublicKey

```go
func PrivateKeyToPublicKey(privateKey string) string
```
PrivateKeyToPublicKey calculates public key from private key

#### func  PublicKeyToAddress

```go
func PublicKeyToAddress(publicKey string) types.Address
```
PublicKeyToAddress generate address from public key

Account address in conflux starts with '0x1'

#### func  ToCfxGeneralAddress

```go
func ToCfxGeneralAddress(address common.Address) types.Address
```
ToCfxGeneralAddress converts a normal address to conflux customerd general
address whose hex string starts with '0x1'