[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://github.com/Conflux-Chain/go-conflux-sdk/blob/master/LICENSE)
[![Documentation](https://img.shields.io/badge/Documentation-GoDoc-green.svg)](https://godoc.org/github.com/Conflux-Chain/go-conflux-sdk)
[![Build Status](https://travis-ci.org/Conflux-Chain/go-conflux-sdk.svg?branch=master)](https://travis-ci.org/Conflux-Chain/go-conflux-sdk)

# Conflux Golang API

The Conflux Golang API allows any Golang client to interact with a local or remote Conflux node based on JSON-RPC 2.0 protocol. With Conflux Golang API, user can easily manage accounts, send transactions, deploy smart contracts and query blockchain information.

## Install
```
go get github.com/Conflux-Chain/go-conflux-sdk
```
You can also add the Conflux Golang API into vendor folder.
```
govendor fetch github.com/Conflux-Chain/go-conflux-sdk
```

## Usage

[api document](https://github.com/Conflux-Chain/go-conflux-sdk/blob/master/api.md)

## Manage Accounts
Use `AccountManager` struct to manage accounts at local machine.
- Create/Import/Update/Delete an account.
- List all accounts.
- Unlock/Lock an account.
- Sign a transaction.

## Query Conflux Information
Use `Client` struct to query Conflux blockchain information, such as block, epoch, transaction, receipt. Following is an example to query the current epoch number:
```go
package main

import (
	"fmt"

	conflux "github.com/Conflux-Chain/go-conflux-sdk"
)

func main() {
	client, err := conflux.NewClient("http://52.175.52.111:12537")
	if err != nil {
		fmt.Println("failed to create client:", err)
		return
	}
	defer client.Close()

	epoch, err := client.GetEpochNumber()
	if err != nil {
		fmt.Println("failed to get epoch number:", err)
		return
	}

	fmt.Println("Current epoch number:", epoch)
}
```

## Send Transaction
To send a transaction, you need to sign the transaction at local machine, and send the signed transaction to local or remote Conflux node.
- Sign a transaction with unlocked account:

    `AccountManager.SignTransaction(tx UnsignedTransaction)`

- Sign a transaction with passphrase for locked account:

	`AccountManager.SignTransactionWithPassphrase(tx UnsignedTransaction, passphrase string)`

- Send a unsigned transaction

    `Client.SendTransaction(tx types.UnsignedTransaction)`

- Send a encoded transaction

    `Client.SendRawTransaction(rawData []byte)`

- Encode a encoded unsigned transaction with signature and send transaction

    `Client.SignEncodedTransactionAndSend(encodedTx []byte, v byte, r, s []byte)`

To send multiple transactions at a time, you can unlock the account at first, then send multiple transactions without passphrase. To send a single transaction, you can just only send the transaction with passphrase.

## Deploy/Invoke Smart Contract

**The simpiest and recommend way is to use [conflux-abigen](https://github.com/Conflux-Chain/conflux-abigen) to generate contract binding to deploy and invoke with contract**

***[Depreated]***
However you also can use `Client.DeployContract` to deploy a contract or use `Client.GetContract` to get a contract by deployed address. Then you can use the contract instance to operate contract, there are GetData/Call/SendTransaction. Please see [api document](https://github.com/Conflux-Chain/go-conflux-sdk/blob/master/api.md) for detail.



### Contract Example ***[Depreated]***
Please reference [contract example]((https://github.com/Conflux-Chain/go-conflux-sdk/blob/master/example/example_contract)) for all source code
```go
package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"time"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
)

func main() {

	//init client
	client, err := sdk.NewClient("https://test.confluxrpc.com", sdk.ClientOption{
		KeystorePath: "../context/keystore",
	})
	utils.PanicIfErr(err, "failed to new client")

	// unlock default account
	client.AccountManager.TimedUnlockDefault("hello", 300*time.Second)

	//deploy contract
	abi := your contract abi
	bytecode := your contract bytecode

	result := client.DeployContract(nil, abi, bytecode, big.NewInt(100000), "biu", uint8(10), "BIU")
	<-result.DoneChannel
	utils.PanicIfErr(result.Error, "failed to deploy contract")

	contract := result.DeployedContract
	fmt.Printf("deploy contract by client.DeployContract done\ncontract address: %+v\ntxhash:%v\n\n", contract.Address, result.TransactionHash)

	// or get contract by deployed address
	// deployedAt := client.MustNewAddress("cfxtest:acgkhpdz61g11parejzbftznnt8gds15mp4wg54j5c")
	// contract, err := client.GetContract(abi, &deployedAt)

	//get data for send/call contract method
	user := client.MustNewAddress("cfxtest:aap9kthvctunvf030rbkk9k7zbzyz12dajp1u3sp4g")
	data, err := contract.GetData("balanceOf", user.MustGetCommonAddress())
	utils.PanicIfErr(err, "failed to get data")
	fmt.Printf("get data of method balanceOf: 0x%x\n\n", data)

	//call contract method
	//Note: the output struct type need match method output type of ABI, go type "*big.Int" match abi type "uint256", go type "struct{Balance *big.Int}" match abi tuple type "(balance uint256)"
	balance := &struct{ Balance *big.Int }{}
	err = contract.Call(nil, balance, "balanceOf", user.MustGetCommonAddress())
	utils.PanicIfErr(err, "failed to get balance")
	fmt.Printf("balance of address %v in contract is: %+v\n\n", user, balance)

	//send transction for contract method
	to := client.MustNewAddress("cfxtest:acgkhpdz61g11parejzbftznnt8gds15mp4wg54j5c")
	txhash, err := contract.SendTransaction(nil, "transfer", to.MustGetCommonAddress(), big.NewInt(10))
	utils.PanicIfErr(err, "failed to transfer")
	fmt.Printf("transfer %v erc20 token to %v done, tx hash: %v\n\n", 10, to, txhash)

	//get event log and decode it
	fmt.Println("wait for transaction be packed...")
	receipt, err := client.WaitForTransationReceipt(txhash, time.Second*2)
	utils.PanicIfErr(err, "failed to get tx receipt")
	fmt.Printf("get receipt: %+v\n\n", receipt)

	// decode Transfer Event
	var Transfer struct {
		From  common.Address
		To    common.Address
		Value *big.Int
	}

	err = contract.DecodeEvent(&Transfer, "Transfer", receipt.Logs[0])
	utils.PanicIfErr(err, "failed to decode event")
	fmt.Printf("decoded transfer event: {From: 0x%x, To: 0x%x, Value: %v} ", Transfer.From, Transfer.To, Transfer.Value)
}
```

## Subscribe Epochs/BlockHeads/Logs

Please find Publish-Subscribe API documentation from https://developer.confluxnetwork.org/conflux-doc/docs/pubsub

It should be noted that when subscribing logs, a `SubscribeLogs` object is received. It has two fields `Log` and `ChainRerog`, one of them must be nil and the other not. When Log is not nil, it means that a Log is received. When field `ChainReorg` is not nil, that means chainreorg is occurred. That represents the log related to epoch greater than or equal to `ChainReog.RevertTo` will become invalid, and the dapp needs to be dealt with at the business level.

## Use middleware to hook rpc request

Client applies method `UseCallRpcMiddleware` to set middleware for hooking `callRpc` method which is the core of all single rpc related methods. And `UseBatchCallRpcMiddleware` to set middleware for hooking `batchCallRPC`.

For example, use `CallRpcConsoleMiddleware` to log for rpc requests.
```golang
client.UseCallRpcMiddleware(middleware.CallRpcConsoleMiddleware)
```

Also you could 
- customize middleware
- use multiple middleware

Notice that the middleware chain exectuion order is like onion, for example, use middleware A first and then middleware B
```go
client.UseCallRpcMiddleware(A)
client.UseCallRpcMiddleware(B)
```
the middleware execution order is
```
B --> A --> client.callRpc --> A --> B
```

## Appendix
### Mapping of solidity types to go types 
This is a mapping table for map solidity types to go types when using contract methods GetData/Call/SendTransaction/DecodeEvent
| solidity types                               | go types                                                                          |
|----------------------------------------------|-----------------------------------------------------------------------------------|
| address                                      | common.Address                                                                    |
| uint8,uint16,uint32,uint64                   | uint8,uint16,uint32,uint64                                                        |
| uint24,uint40,uint48,uint56,uint72...uint256 | *big.Int                                                                          |
| int8,int16,int32,int64                       | int8,int16,int32,int64                                                            |
| int24,int40,int48,int56,int72...int256       | *big.Int                                                                          |
| fixed bytes (bytes1,bytes2...bytes32)        | [length]byte                                                                      |
| fixed type T array (T[length])               | [length]TG (TG is go type matched with solidty type T)                            |
| bytes                                        | []byte                                                                            |
| dynamic type T array T[]                     | []TG ((TG is go type matched with solidty type T))                                |
| function                                     | [24]byte                                                                          |
| string                                       | string                                                                            |
| bool                                         | bool                                                                              |
| tuple                                        | struct  eg:[{"name": "balance","type": "uint256"}] => struct {Balance *big.Int} |
