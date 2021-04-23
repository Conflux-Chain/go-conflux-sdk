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
import "."
```


### type AccountManager

```go
type AccountManager struct {
}
```

AccountManager manages Conflux accounts.

#### func  NewAccountManager

```go
func NewAccountManager(keydir string, networkID uint32) *AccountManager
```
NewAccountManager creates an instance of AccountManager based on the keystore
directory "keydir".

#### func (*AccountManager) Create

```go
func (m *AccountManager) Create(passphrase string) (address types.Address, err error)
```
Create creates a new account and puts the keystore file into keystore directory

#### func (*AccountManager) CreateEthCompatible

```go
func (m *AccountManager) CreateEthCompatible(passphrase string) (address types.Address, err error)
```
CreateEthCompatible creates a new account compatible with eth and puts the
keystore file into keystore directory

#### func (*AccountManager) Delete

```go
func (m *AccountManager) Delete(address types.Address, passphrase string) error
```
Delete deletes the specified account and remove the keystore file from keystore
directory.

#### func (*AccountManager) Export

```go
func (m *AccountManager) Export(address types.Address, passphrase string) (string, error)
```
Export exports private key string of address

#### func (*AccountManager) GetDefault

```go
func (m *AccountManager) GetDefault() (*types.Address, error)
```
GetDefault return first account in keystore directory

#### func (*AccountManager) Import

```go
func (m *AccountManager) Import(keyFile, passphrase, newPassphrase string) (address types.Address, err error)
```
Import imports account from external key file to keystore directory. Returns
error if the account already exists.

#### func (*AccountManager) ImportKey

```go
func (m *AccountManager) ImportKey(keyString string, passphrase string) (address types.Address, err error)
```
ImportKey import account from private key hex string and save to keystore
directory

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
func (m *AccountManager) SignTransactionWithPassphrase(tx types.UnsignedTransaction, passphrase string) (types.SignedTransaction, error)
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
	AccountManager AccountManagerOperator
}
```

Client represents a client to interact with Conflux blockchain.

#### func  NewClient

```go
func NewClient(nodeURL string, option ...ClientOption) (*Client, error)
```
NewClient creates an instance of Client with specified conflux node url, it will
creat account manager if option.KeystorePath not empty.

#### func  NewClientWithRPCRequester

```go
func NewClientWithRPCRequester(rpcRequester RpcRequester) (*Client, error)
```
NewClientWithRPCRequester creates client with specified rpcRequester

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
func (client *Client) Call(request types.CallRequest, epoch *types.Epoch) (result hexutil.Bytes, err error)
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

#### func (*Client) CheckBalanceAgainstTransaction

```go
func (client *Client) CheckBalanceAgainstTransaction(accountAddress types.Address,
	contractAddress types.Address,
	gasLimit *hexutil.Big,
	gasPrice *hexutil.Big,
	storageLimit *hexutil.Big,
	epoch ...*types.Epoch) (response types.CheckBalanceAgainstTransactionResponse, err error)
```
CheckBalanceAgainstTransaction checks if user balance is enough for the
transaction.

#### func (*Client) Close

```go
func (client *Client) Close()
```
Close closes the client, aborting any in-flight requests.

#### func (*Client) CreateUnsignedTransaction

```go
func (client *Client) CreateUnsignedTransaction(from types.Address, to types.Address, amount *hexutil.Big, data []byte) (types.UnsignedTransaction, error)
```
CreateUnsignedTransaction creates an unsigned transaction by parameters, and the
other fields will be set to values fetched from conflux node.

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
func (client *Client) EstimateGasAndCollateral(request types.CallRequest, epoch ...*types.Epoch) (estimat types.Estimate, err error)
```
EstimateGasAndCollateral excutes a message call "request" and returns the amount
of the gas used and storage for collateral

#### func (*Client) FilterTraces

```go
func (client *Client) FilterTraces(traceFilter types.TraceFilter) (traces []types.LocalizedTrace, err error)
```
GetFilterTraces returns all traces matching the provided filter.

#### func (*Client) GetAccountInfo

```go
func (client *Client) GetAccountInfo(account types.Address, epoch ...*types.Epoch) (accountInfo types.AccountInfo, err error)
```
GetAccountInfo returns account related states of the given account

#### func (*Client) GetAccountPendingInfo

```go
func (client *Client) GetAccountPendingInfo(address types.Address) (pendignInfo *types.AccountPendingInfo, err error)
```
GetAccountPendingInfo gets transaction pending info by account address

#### func (*Client) GetAccumulateInterestRate

```go
func (client *Client) GetAccumulateInterestRate(epoch ...*types.Epoch) (intersetRate *hexutil.Big, err error)
```
GetAccumulateInterestRate returns accumulate interest rate of the given epoch

#### func (*Client) GetAdmin

```go
func (client *Client) GetAdmin(contractAddress types.Address, epoch ...*types.Epoch) (admin *types.Address, err error)
```
GetAdmin returns admin of the given contract, it will return nil if contract not
exist

#### func (*Client) GetBalance

```go
func (client *Client) GetBalance(address types.Address, epoch ...*types.Epoch) (balance *hexutil.Big, err error)
```
GetBalance returns the balance of specified address at epoch.

#### func (*Client) GetBestBlockHash

```go
func (client *Client) GetBestBlockHash() (hash types.Hash, err error)
```
GetBestBlockHash returns the current best block hash.

#### func (*Client) GetBlockByEpoch

```go
func (client *Client) GetBlockByEpoch(epoch *types.Epoch) (block *types.Block, err error)
```
GetBlockByEpoch returns the block of specified epoch. If the epoch is invalid,
return the concrete error.

#### func (*Client) GetBlockByHash

```go
func (client *Client) GetBlockByHash(blockHash types.Hash) (block *types.Block, err error)
```
GetBlockByHash returns the block of specified blockHash If the block is not
found, return nil.

#### func (*Client) GetBlockByHashWithPivotAssumption

```go
func (client *Client) GetBlockByHashWithPivotAssumption(blockHash types.Hash, pivotHash types.Hash, epoch hexutil.Uint64) (block types.Block, err error)
```
GetBlockByHashWithPivotAssumption returns block with given hash and pivot chain
assumption.

#### func (*Client) GetBlockConfirmationRisk

```go
func (client *Client) GetBlockConfirmationRisk(blockHash types.Hash) (*big.Float, error)
```
GetBlockConfirmationRisk indicates the probability that the pivot block of the
epoch where the block is located becomes a normal block.

it's (raw confirmation risk coefficient/ (2^256-1))

#### func (*Client) GetBlockRewardInfo

```go
func (client *Client) GetBlockRewardInfo(epoch types.Epoch) (rewardInfo []types.RewardInfo, err error)
```
GetBlockRewardInfo returns block reward information in an epoch

#### func (*Client) GetBlockSummaryByEpoch

```go
func (client *Client) GetBlockSummaryByEpoch(epoch *types.Epoch) (blockSummary *types.BlockSummary, err error)
```
GetBlockSummaryByEpoch returns the block summary of specified epoch. If the
epoch is invalid, return the concrete error.

#### func (*Client) GetBlockSummaryByHash

```go
func (client *Client) GetBlockSummaryByHash(blockHash types.Hash) (blockSummary *types.BlockSummary, err error)
```
GetBlockSummaryByHash returns the block summary of specified blockHash If the
block is not found, return nil.

#### func (*Client) GetBlockTraces

```go
func (client *Client) GetBlockTraces(blockHash types.Hash) (traces *types.LocalizedBlockTrace, err error)
```
GetBlockTrace returns all traces produced at given block.

#### func (*Client) GetBlocksByEpoch

```go
func (client *Client) GetBlocksByEpoch(epoch *types.Epoch) (blockHashes []types.Hash, err error)
```
GetBlocksByEpoch returns the blocks hash in the specified epoch.

#### func (*Client) GetClientVersion

```go
func (client *Client) GetClientVersion() (clientVersion string, err error)
```
GetClientVersion returns the client version as a string

#### func (*Client) GetCode

```go
func (client *Client) GetCode(address types.Address, epoch ...*types.Epoch) (code hexutil.Bytes, err error)
```
GetCode returns the bytecode in HEX format of specified address at epoch.

#### func (*Client) GetCollateralForStorage

```go
func (client *Client) GetCollateralForStorage(account types.Address, epoch ...*types.Epoch) (storage *hexutil.Big, err error)
```
GetCollateralForStorage returns balance of the given account.

#### func (*Client) GetContract

```go
func (client *Client) GetContract(abiJSON []byte, deployedAt *types.Address) (*Contract, error)
```
GetContract creates a contract instance according to abi json and it's deployed
address

#### func (*Client) GetDepositList

```go
func (client *Client) GetDepositList(address types.Address, epoch ...*types.Epoch) (depositInfos []types.DepositInfo, err error)
```
GetDepositList returns deposit list of the given account.

#### func (*Client) GetEpochNumber

```go
func (client *Client) GetEpochNumber(epoch ...*types.Epoch) (epochNumber *hexutil.Big, err error)
```
GetEpochNumber returns the highest or specified epoch number.

#### func (*Client) GetEpochReceipts

```go
func (client *Client) GetEpochReceipts(epoch types.Epoch) (receipts [][]types.TransactionReceipt, err error)
```

#### func (*Client) GetGasPrice

```go
func (client *Client) GetGasPrice() (gasPrice *hexutil.Big, err error)
```
GetGasPrice returns the recent mean gas price.

#### func (*Client) GetInterestRate

```go
func (client *Client) GetInterestRate(epoch ...*types.Epoch) (intersetRate *hexutil.Big, err error)
```
GetInterestRate returns interest rate of the given epoch

#### func (*Client) GetLogs

```go
func (client *Client) GetLogs(filter types.LogFilter) (logs []types.Log, err error)
```
GetLogs returns logs that matching the specified filter.

#### func (*Client) GetNetworkID

```go
func (client *Client) GetNetworkID() (uint32, error)
```
GetNetworkID returns networkID of connecting conflux node

#### func (*Client) GetNextNonce

```go
func (client *Client) GetNextNonce(address types.Address, epoch ...*types.Epoch) (nonce *hexutil.Big, err error)
```
GetNextNonce returns the next transaction nonce of address

#### func (*Client) GetNodeURL

```go
func (client *Client) GetNodeURL() string
```
GetNodeURL returns node url

#### func (*Client) GetRawBlockConfirmationRisk

```go
func (client *Client) GetRawBlockConfirmationRisk(blockhash types.Hash) (risk *hexutil.Big, err error)
```
GetRawBlockConfirmationRisk indicates the risk coefficient that the pivot block
of the epoch where the block is located becomes a normal block. It will return
nil if block not exist

#### func (*Client) GetSkippedBlocksByEpoch

```go
func (client *Client) GetSkippedBlocksByEpoch(epoch *types.Epoch) (blockHashs []types.Hash, err error)
```
GetSkippedBlocksByEpoch returns skipped block hashes of given epoch

#### func (*Client) GetSponsorInfo

```go
func (client *Client) GetSponsorInfo(contractAddress types.Address, epoch ...*types.Epoch) (sponsor types.SponsorInfo, err error)
```
GetSponsorInfo returns sponsor information of the given contract

#### func (*Client) GetStakingBalance

```go
func (client *Client) GetStakingBalance(account types.Address, epoch ...*types.Epoch) (balance *hexutil.Big, err error)
```
GetStakingBalance returns balance of the given account.

#### func (*Client) GetStatus

```go
func (client *Client) GetStatus() (status types.Status, err error)
```
GetStatus returns status of connecting conflux node

#### func (*Client) GetStorageAt

```go
func (client *Client) GetStorageAt(address types.Address, position types.Hash, epoch ...*types.Epoch) (storageEntries hexutil.Bytes, err error)
```
GetStorageAt returns storage entries from a given contract.

#### func (*Client) GetStorageRoot

```go
func (client *Client) GetStorageRoot(address types.Address, epoch ...*types.Epoch) (storageRoot *types.StorageRoot, err error)
```
GetStorageRoot returns storage root of given address

#### func (*Client) GetSupplyInfo

```go
func (client *Client) GetSupplyInfo(epoch ...*types.Epoch) (info types.TokenSupplyInfo, err error)
```
GetSupplyInfo Return information about total token supply.

#### func (*Client) GetTransactionByHash

```go
func (client *Client) GetTransactionByHash(txHash types.Hash) (tx *types.Transaction, err error)
```
GetTransactionByHash returns transaction for the specified txHash. If the
transaction is not found, return nil.

#### func (*Client) GetTransactionReceipt

```go
func (client *Client) GetTransactionReceipt(txHash types.Hash) (receipt *types.TransactionReceipt, err error)
```
GetTransactionReceipt returns the receipt of specified transaction hash. If no
receipt is found, return nil.

#### func (*Client) GetTransactionTraces

```go
func (client *Client) GetTransactionTraces(txHash types.Hash) (traces []types.LocalizedTrace, err error)
```
GetTransactionTraces returns all traces produced at the given transaction.

#### func (*Client) GetVoteList

```go
func (client *Client) GetVoteList(address types.Address, epoch ...*types.Epoch) (voteStakeInfos []types.VoteStakeInfo, err error)
```
GetVoteList returns vote list of the given account.

#### func (*Client) MustNewAddress

```go
func (client *Client) MustNewAddress(base32OrHex string) types.Address
```
MustNewAddress create conflux address by base32 string or hex40 string, if
base32OrHex is base32 and networkID is passed it will create cfx Address use
networkID of current client. it will painc if error occured.

#### func (*Client) NewAddress

```go
func (client *Client) NewAddress(base32OrHex string) (types.Address, error)
```
NewAddress create conflux address by base32 string or hex40 string, if
base32OrHex is base32 and networkID is passed it will create cfx Address use
networkID of current client.

#### func (*Client) SendRawTransaction

```go
func (client *Client) SendRawTransaction(rawData []byte) (types.Hash, error)
```
SendRawTransaction sends signed transaction and returns its hash.

#### func (*Client) SendTransaction

```go
func (client *Client) SendTransaction(tx types.UnsignedTransaction) (types.Hash, error)
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

#### func (*Client) SubscribeEpochs

```go
func (client *Client) SubscribeEpochs(channel chan types.WebsocketEpochResponse, subscriptionEpochType ...types.Epoch) (*rpc.ClientSubscription, error)
```
SubscribeEpochs subscribes consensus results: the total order of blocks, as
expressed by a sequence of epochs. Currently subscriptionEpochType only support
"latest_mined" and "latest_state"

#### func (*Client) SubscribeLogs

```go
func (client *Client) SubscribeLogs(logChannel chan types.Log, chainReorgChannel chan types.ChainReorg, filter types.LogFilter) (*rpc.ClientSubscription, error)
```
SubscribeLogs subscribes all logs matching a certain filter, in order.

#### func (*Client) SubscribeNewHeads

```go
func (client *Client) SubscribeNewHeads(channel chan types.BlockHeader) (*rpc.ClientSubscription, error)
```
SubscribeNewHeads subscribes all new block headers participating in the
consensus.

#### func (*Client) WaitForTransationBePacked

```go
func (client *Client) WaitForTransationBePacked(txhash types.Hash, duration time.Duration) (*types.Transaction, error)
```
WaitForTransationBePacked returns transaction when it is packed

#### func (*Client) WaitForTransationReceipt

```go
func (client *Client) WaitForTransationReceipt(txhash types.Hash, duration time.Duration) (*types.TransactionReceipt, error)
```
WaitForTransationReceipt waits for transaction receipt valid

### type ClientOption

```go
type ClientOption struct {
	KeystorePath    string
	RetryCount      int
	RetryInterval   time.Duration
	CallRpcLog      func(method string, args []interface{}, result interface{}, resultError error, duration time.Duration)
	BatchCallRPCLog func(b []rpc.BatchElem, err error, duration time.Duration)
}
```

ClientOption for set keystore path and flags for retry

The simplest way to set logger is to use the types.DefaultCallRpcLog and
types.DefaultBatchCallRPCLog

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
func (contract *Contract) DecodeEvent(out interface{}, event string, log types.Log) error
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
func (contract *Contract) SendTransaction(option *types.ContractMethodSendOption, method string, args ...interface{}) (types.Hash, error)
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
import "."
```


#### func  CalcBlockConfirmationRisk

```go
func CalcBlockConfirmationRisk(rawConfirmationRisk *big.Int) *big.Float
```
CalcBlockConfirmationRisk calculates block revert rate

#### func  GetMapSortedKeys

```go
func GetMapSortedKeys(m map[string]interface{}) []string
```

#### func  GetObjFileds

```go
func GetObjFileds(obj interface{}) []string
```

#### func  GetObjJsonFieldTags

```go
func GetObjJsonFieldTags(obj interface{}) []string
```

#### func  Has0xPrefix

```go
func Has0xPrefix(input string) bool
```
Has0xPrefix returns true if input starts with '0x' or '0X'

#### func  HexStringToBytes

```go
func HexStringToBytes(hexStr string) (hexutil.Bytes, error)
```
HexStringToBytes converts hex string to bytes

#### func  IsNil

```go
func IsNil(i interface{}) bool
```
IsNil sepecialy checks if interface object is nil

#### func  IsRPCJSONError

```go
func IsRPCJSONError(err error) bool
```
IsRPCJSONError returns true if err is rpc error

#### func  Keccak256

```go
func Keccak256(hexStr string) (string, error)
```
Keccak256 hashes hex string by keccak256 and returns it's hash value

#### func  PanicIfErr

```go
func PanicIfErr(err error, msg string)
```
PanicIfErr panic and reports error message

#### func  PanicIfErrf

```go
func PanicIfErrf(err error, msg string, args ...interface{})
```
PanicIfErrf panic and reports error message with args

#### func  PrettyJSON

```go
func PrettyJSON(value interface{}) string
```
PrettyJSON json marshal value and pretty with indent

#### func  PrivateKeyToPublicKey

```go
func PrivateKeyToPublicKey(privateKey string) string
```
PrivateKeyToPublicKey calculates public key from private key

#### func  PublicKeyToCommonAddress

```go
func PublicKeyToCommonAddress(publicKey string) common.Address
```
PublicKeyToCommonAddress generate address from public key

Account address in conflux starts with '0x1'
## package internalcontract
```
import "."
```


### type AdminControl

```go
type AdminControl struct {
	sdk.Contract
}
```

AdminControl contract

#### func  NewAdminControl

```go
func NewAdminControl(client sdk.ClientOperator) (ac AdminControl, err error)
```
NewAdminControl gets the AdminControl contract object

#### func (*AdminControl) Destroy

```go
func (ac *AdminControl) Destroy(option *types.ContractMethodSendOption, contractAddr types.Address) (types.Hash, error)
```
Destroy destroies contract `contractAddr`.

#### func (*AdminControl) GetAdmin

```go
func (ac *AdminControl) GetAdmin(option *types.ContractMethodCallOption, contractAddr types.Address) (types.Address, error)
```
GetAdmin returns admin of specific contract

#### func (*AdminControl) SetAdmin

```go
func (ac *AdminControl) SetAdmin(option *types.ContractMethodSendOption, contractAddr types.Address, newAdmin types.Address) (types.Hash, error)
```
SetAdmin sets the administrator of contract `contractAddr` to `newAdmin`.

### type Sponsor

```go
type Sponsor struct {
	sdk.Contract
}
```

Sponsor represents SponsorWhitelistControl contract

#### func  NewSponsor

```go
func NewSponsor(client sdk.ClientOperator) (s Sponsor, err error)
```
NewSponsor gets the SponsorWhitelistControl contract object

#### func (*Sponsor) AddPrivilegeByAdmin

```go
func (s *Sponsor) AddPrivilegeByAdmin(option *types.ContractMethodSendOption, contractAddr types.Address, userAddresses []types.Address) (types.Hash, error)
```
AddPrivilegeByAdmin for admin adds user to whitelist

#### func (*Sponsor) GetSponsorForCollateral

```go
func (s *Sponsor) GetSponsorForCollateral(option *types.ContractMethodCallOption, contractAddr types.Address) (address types.Address, err error)
```
GetSponsorForCollateral gets collateral sponsor address

#### func (*Sponsor) GetSponsorForGas

```go
func (s *Sponsor) GetSponsorForGas(option *types.ContractMethodCallOption, contractAddr types.Address) (address types.Address, err error)
```
GetSponsorForGas gets gas sponsor address of specific contract

#### func (*Sponsor) GetSponsoredBalanceForCollateral

```go
func (s *Sponsor) GetSponsoredBalanceForCollateral(option *types.ContractMethodCallOption, contractAddr types.Address) (*big.Int, error)
```
GetSponsoredBalanceForCollateral gets current Sponsored Balance for collateral

#### func (*Sponsor) GetSponsoredBalanceForGas

```go
func (s *Sponsor) GetSponsoredBalanceForGas(option *types.ContractMethodCallOption, contractAddr types.Address) (*big.Int, error)
```
GetSponsoredBalanceForGas gets current Sponsored Balance for gas

#### func (*Sponsor) GetSponsoredGasFeeUpperBound

```go
func (s *Sponsor) GetSponsoredGasFeeUpperBound(option *types.ContractMethodCallOption, contractAddr types.Address) (*big.Int, error)
```
GetSponsoredGasFeeUpperBound gets current Sponsored Gas fee upper bound

#### func (*Sponsor) IsAllWhitelisted

```go
func (s *Sponsor) IsAllWhitelisted(option *types.ContractMethodCallOption, contractAddr types.Address) (bool, error)
```
IsAllWhitelisted checks if all users are in a contract's whitelist

#### func (*Sponsor) IsWhitelisted

```go
func (s *Sponsor) IsWhitelisted(option *types.ContractMethodCallOption, contractAddr types.Address, userAddr types.Address) (bool, error)
```
IsWhitelisted checks if a user is in a contract's whitelist

#### func (*Sponsor) RemovePrivilegeByAdmin

```go
func (s *Sponsor) RemovePrivilegeByAdmin(option *types.ContractMethodSendOption, contractAddr types.Address, userAddresses []types.Address) (types.Hash, error)
```
RemovePrivilegeByAdmin for admin removes user from whitelist

#### func (*Sponsor) SetSponsorForCollateral

```go
func (s *Sponsor) SetSponsorForCollateral(option *types.ContractMethodSendOption, contractAddr types.Address) (types.Hash, error)
```
SetSponsorForCollateral for someone sponsor the storage collateral for contract
`contractAddr`, it is payable

#### func (*Sponsor) SetSponsorForGas

```go
func (s *Sponsor) SetSponsorForGas(option *types.ContractMethodSendOption, contractAddr types.Address, upperBound *big.Int) (types.Hash, error)
```
SetSponsorForGas for someone sponsor the gas cost for contract `contractAddr`
with an `upper_bound` for a single transaction, it is payable

### type Staking

```go
type Staking struct {
	sdk.Contract
}
```

Staking contract

#### func  NewStaking

```go
func NewStaking(client sdk.ClientOperator) (s Staking, err error)
```
NewStaking gets the Staking contract object

#### func (*Staking) Deposit

```go
func (s *Staking) Deposit(option *types.ContractMethodSendOption, amount *big.Int) (types.Hash, error)
```
Deposit `amount` cfx in this contract

#### func (*Staking) GetLockedStakingBalance

```go
func (ac *Staking) GetLockedStakingBalance(option *types.ContractMethodCallOption, user types.Address, blockNumber *big.Int) (*big.Int, error)
```
GetLockedStakingBalance returns user's locked staking balance at given
blockNumber Note: if the blockNumber is less than the current block number,
function will return current locked staking balance.

#### func (*Staking) GetStakingBalance

```go
func (ac *Staking) GetStakingBalance(option *types.ContractMethodCallOption, user types.Address) (*big.Int, error)
```
GetStakingBalance returns user's staking balance

#### func (*Staking) GetVotePower

```go
func (ac *Staking) GetVotePower(option *types.ContractMethodCallOption, user types.Address, blockNumber *big.Int) (*big.Int, error)
```
GetVotePower returns user's vote power staking balance at given blockNumber

#### func (*Staking) VoteLock

```go
func (s *Staking) VoteLock(option *types.ContractMethodSendOption, amount *big.Int, unlockBlockNumber *big.Int) (types.Hash, error)
```
VoteLock will locks `amount` cfx from current to next `unlockBlockNumber` blocks
for obtain vote power

#### func (*Staking) Withdraw

```go
func (s *Staking) Withdraw(option *types.ContractMethodSendOption, amount *big.Int) (types.Hash, error)
```
Withdraw `amount` cfx from this contract
