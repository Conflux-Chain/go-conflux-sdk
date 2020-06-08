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

## Installation
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
url:= "http://testnet-jsonrpc.conflux-chain.org:12537"
client, err := sdk.NewClient(url)
if err != nil {
	fmt.Println("new client error:", err)
	return
}
am := sdk.NewAccountManager("./keystore")
client.SetAccountManager(am)
```
## package sdk
```
import "github.com/Conflux-Chain/go-conflux-sdk"
```


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

#### func  NewClientWithRPCRequester

```go
func NewClientWithRPCRequester(rpcRequester rpcRequester) (*Client, error)
```
NewClientWithRPCRequester creates client with specified rpcRequester

#### func  NewClientWithRetry

```go
func NewClientWithRetry(nodeURL string, retryCount int, retryInterval time.Duration) (*Client, error)
```
NewClientWithRetry creates a retryable new instance of Client with specified
conflux node url and retry options.

the retryInterval will be set to 1 second if pass 0

#### func (*Client) ApplyUnsignedTransactionDefault

```go
func (client *Client) ApplyUnsignedTransactionDefault(tx *types.UnsignedTransaction) error
```
ApplyUnsignedTransactionDefault set empty fields to value fetched from conflux
node.

#### func (*Client) BatchCallRPC

```go
func (client *Client) BatchCallRPC(b []rpc.BatchElem) error
```
BatchCallRPC sends all given requests as a single batch and waits for the server
to return a response for all of them.

In contrast to Call, BatchCall only returns I/O errors. Any error specific to a
request is reported through the Error field of the corresponding BatchElem.

Note that batch calls may not be executed atomically on the server side.

#### func (*Client) BatchGetBlockConfirmationRisk

```go
func (client *Client) BatchGetBlockConfirmationRisk(blockhashes []types.Hash) (map[types.Hash]*big.Float, error)
```
BatchGetBlockConfirmationRisk acquires confirmation risk informations in bulk by
blockhashes

#### func (*Client) BatchGetBlockSummarys

```go
func (client *Client) BatchGetBlockSummarys(blockhashes []types.Hash) (map[types.Hash]*types.BlockSummary, error)
```
BatchGetBlockSummarys requests block summary informations in bulk by blockhashes

#### func (*Client) BatchGetRawBlockConfirmationRisk

```go
func (client *Client) BatchGetRawBlockConfirmationRisk(blockhashes []types.Hash) (map[types.Hash]*big.Int, error)
```
BatchGetRawBlockConfirmationRisk requests raw confirmation risk informations in
bulk by blockhashes

#### func (*Client) BatchGetTxByHashes

```go
func (client *Client) BatchGetTxByHashes(txhashes []types.Hash) (map[types.Hash]*types.Transaction, error)
```
BatchGetTxByHashes requests transaction informations in bulk by txhashes

#### func (*Client) Call

```go
func (client *Client) Call(request types.CallRequest, epoch *types.Epoch) (*string, error)
```
Call executes a message call transaction "request" at specified epoch, which is
directly executed in the VM of the node, but never mined into the block chain
and returns the contract execution result.

#### func (*Client) CallRPC

```go
func (client *Client) CallRPC(result interface{}, method string, args ...interface{}) error
```
CallRPC performs a JSON-RPC call with the given arguments and unmarshals into
result if no error occurred.

The result must be a pointer so that package json can unmarshal into it. You can
also pass nil, in which case the result is ignored.

#### func (*Client) Close

```go
func (client *Client) Close()
```
Close closes the client, aborting any in-flight requests.

#### func (*Client) CreateUnsignedTransaction

```go
func (client *Client) CreateUnsignedTransaction(from types.Address, to types.Address, amount *hexutil.Big, data []byte) (*types.UnsignedTransaction, error)
```
CreateUnsignedTransaction creates an unsigned transaction by parameters, and the
other fields will be set to values fetched from conflux node.

#### func (*Client) Debug

```go
func (client *Client) Debug(method string, args ...interface{}) (interface{}, error)
```
Debug calls the Conflux debug API.

#### func (*Client) DeployContract

```go
func (client *Client) DeployContract(option *types.ContractDeployOption, abiJSON []byte,
	bytecode []byte, constroctorParams ...interface{}) *ContractDeployResult
```
DeployContract deploys a contract by abiJSON, bytecode and consturctor params.
It returns a ContractDeployState instance which contains 3 channels for
notifying when state changed.

#### func (*Client) EstimateGasAndCollateral

```go
func (client *Client) EstimateGasAndCollateral(request types.CallRequest) (*types.Estimate, error)
```
EstimateGasAndCollateral excutes a message call "request" and returns the amount
of the gas used and storage for collateral

#### func (*Client) GetBalance

```go
func (client *Client) GetBalance(address types.Address, epoch ...*types.Epoch) (*big.Int, error)
```
GetBalance returns the balance of specified address at epoch.

#### func (*Client) GetBestBlockHash

```go
func (client *Client) GetBestBlockHash() (types.Hash, error)
```
GetBestBlockHash returns the current best block hash.

#### func (*Client) GetBlockByEpoch

```go
func (client *Client) GetBlockByEpoch(epoch *types.Epoch) (*types.Block, error)
```
GetBlockByEpoch returns the block of specified epoch. If the epoch is invalid,
return the concrete error.

#### func (*Client) GetBlockByHash

```go
func (client *Client) GetBlockByHash(blockHash types.Hash) (*types.Block, error)
```
GetBlockByHash returns the block of specified blockHash If the block is not
found, return nil.

#### func (*Client) GetBlockConfirmationRisk

```go
func (client *Client) GetBlockConfirmationRisk(blockHash types.Hash) (*big.Float, error)
```
GetBlockConfirmationRisk indicates the probability that the pivot block of the
epoch where the block is located becomes a normal block.

it's (raw confirmation risk coefficient/ (2^256-1))

#### func (*Client) GetBlockSummaryByEpoch

```go
func (client *Client) GetBlockSummaryByEpoch(epoch *types.Epoch) (*types.BlockSummary, error)
```
GetBlockSummaryByEpoch returns the block summary of specified epoch. If the
epoch is invalid, return the concrete error.

#### func (*Client) GetBlockSummaryByHash

```go
func (client *Client) GetBlockSummaryByHash(blockHash types.Hash) (*types.BlockSummary, error)
```
GetBlockSummaryByHash returns the block summary of specified blockHash If the
block is not found, return nil.

#### func (*Client) GetBlocksByEpoch

```go
func (client *Client) GetBlocksByEpoch(epoch *types.Epoch) ([]types.Hash, error)
```
GetBlocksByEpoch returns the blocks hash in the specified epoch.

#### func (*Client) GetCode

```go
func (client *Client) GetCode(address types.Address, epoch ...*types.Epoch) (string, error)
```
GetCode returns the bytecode in HEX format of specified address at epoch.

#### func (*Client) GetContract

```go
func (client *Client) GetContract(abiJSON []byte, deployedAt *types.Address) (*Contract, error)
```
GetContract creates a contract instance according to abi json and it's deployed
address

#### func (*Client) GetEpochNumber

```go
func (client *Client) GetEpochNumber(epoch ...*types.Epoch) (*big.Int, error)
```
GetEpochNumber returns the highest or specified epoch number.

#### func (*Client) GetGasPrice

```go
func (client *Client) GetGasPrice() (*big.Int, error)
```
GetGasPrice returns the recent mean gas price.

#### func (*Client) GetLogs

```go
func (client *Client) GetLogs(filter types.LogFilter) ([]types.Log, error)
```
GetLogs returns logs that matching the specified filter.

#### func (*Client) GetNextNonce

```go
func (client *Client) GetNextNonce(address types.Address, epoch *types.Epoch) (*big.Int, error)
```
GetNextNonce returns the next transaction nonce of address

#### func (*Client) GetNodeURL

```go
func (client *Client) GetNodeURL() string
```
GetNodeURL returns node url

#### func (*Client) GetRawBlockConfirmationRisk

```go
func (client *Client) GetRawBlockConfirmationRisk(blockhash types.Hash) (*big.Int, error)
```
GetRawBlockConfirmationRisk indicates the risk coefficient that the pivot block
of the epoch where the block is located becomes a normal block.

#### func (*Client) GetTransactionByHash

```go
func (client *Client) GetTransactionByHash(txHash types.Hash) (*types.Transaction, error)
```
GetTransactionByHash returns transaction for the specified txHash. If the
transaction is not found, return nil.

#### func (*Client) GetTransactionReceipt

```go
func (client *Client) GetTransactionReceipt(txHash types.Hash) (*types.TransactionReceipt, error)
```
GetTransactionReceipt returns the receipt of specified transaction hash. If no
receipt is found, return nil.

#### func (*Client) SendRawTransaction

```go
func (client *Client) SendRawTransaction(rawData []byte) (types.Hash, error)
```
SendRawTransaction sends signed transaction and returns its hash.

#### func (*Client) SendTransaction

```go
func (client *Client) SendTransaction(tx *types.UnsignedTransaction) (types.Hash, error)
```
SendTransaction signs and sends transaction to conflux node and returns the
transaction hash.

#### func (*Client) SetAccountManager

```go
func (client *Client) SetAccountManager(accountManager AccountManagerOperator)
```
SetAccountManager sets account manager for sign transaction

#### func (*Client) SignEncodedTransactionAndSend

```go
func (client *Client) SignEncodedTransactionAndSend(encodedTx []byte, v byte, r, s []byte) (*types.Transaction, error)
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

Contract represents a smart contract. You can conveniently create contract by
Client.GetContract or Client.DeployContract.

#### func  NewContract

```go
func NewContract(abiJSON []byte, client ClientOperator, address *types.Address) (*Contract, error)
```
NewContract creates contract by abi and deployed address

#### func (*Contract) Call

```go
func (contract *Contract) Call(option *types.ContractMethodCallOption, resultPtr interface{}, method string, args ...interface{}) error
```
Call calls to the contract method with args and fills the excuted result to the
"resultPtr".

the resultPtr should be a pointer of the method output struct type.

please refer
https://github.com/Conflux-Chain/go-conflux-sdk/blob/master/README.md to get the
mappings of solidity types to go types

#### func (*Contract) DecodeEvent

```go
func (contract *Contract) DecodeEvent(out interface{}, event string, log types.LogEntry) error
```
DecodeEvent unpacks a retrieved log into the provided output structure.

please refer
https://github.com/Conflux-Chain/go-conflux-sdk/blob/master/README.md to get the
mappings of solidity types to go types

#### func (*Contract) GetData

```go
func (contract *Contract) GetData(method string, args ...interface{}) ([]byte, error)
```
GetData packs the given method name to conform the ABI of the contract. Method
call's data will consist of method_id, args0, arg1, ... argN. Method id consists
of 4 bytes and arguments are all 32 bytes. Method ids are created from the first
4 bytes of the hash of the methods string signature. (signature =
baz(uint32,string32))

please refer
https://github.com/Conflux-Chain/go-conflux-sdk/blob/master/README.md to get the
mappings of solidity types to go types

#### func (*Contract) SendTransaction

```go
func (contract *Contract) SendTransaction(option *types.ContractMethodSendOption, method string, args ...interface{}) (*types.Hash, error)
```
SendTransaction sends a transaction to the contract method with args and returns
its transaction hash

please refer
https://github.com/Conflux-Chain/go-conflux-sdk/blob/master/README.md to get the
mappings of solidity types to go types

### type ContractDeployResult

```go
type ContractDeployResult struct {
	//DoneChannel channel for notifying when contract deployed done
	DoneChannel      <-chan struct{}
	TransactionHash  *types.Hash
	Error            error
	DeployedContract *Contract
}
```

ContractDeployResult for state change notification when deploying contract
## package utils
```
import "github.com/Conflux-Chain/go-conflux-sdk/utils"
```


#### func  CalcBlockConfirmationRisk

```go
func CalcBlockConfirmationRisk(rawConfirmationRisk *big.Int) *big.Float
```
CalcBlockConfirmationRisk calculates block revert rate

#### func  Keccak256

```go
func Keccak256(hexStr string) (string, error)
```
Keccak256 hashes hex string by keccak256 and returns it's hash value

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
