[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://github.com/Conflux-Chain/go-conflux-sdk/blob/master/LICENSE)
[![Documentation](https://img.shields.io/badge/Documentation-GoDoc-green.svg)](https://godoc.org/github.com/Conflux-Chain/go-conflux-sdk)

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

- Send a signed transaction

    `Client.SendSignedTransaction(rawData []byte)`

To send multiple transactions at a time, you can unlock the account at first, then send multiple transactions without passphrase. To send a single transaction, you can just only send the transaction with passphrase.

## Deploy/Call Smart Contract
To deploy or call a smart contract, you can sign a transaction with `Data` field set in `UnsignedTransaction` struct. When deploy a smart contract, you can use ***solc*** to compile the smart contract to get the contract bytecodes, which is set to the `Data` field. To all a contract, you can import the [ABI](https://github.com/ethereum/go-ethereum/tree/master/accounts/abi) library from [go-etherem](https://github.com/ethereum/go-ethereum) to get the encoded method call, which is set to the `Data` field.

### ABI Example
```go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func main() {
	abiJSON, err := ioutil.ReadFile(`E:\solidity\SimpleStorage\SimpleStorage.abi`)
	if err != nil {
		fmt.Println("failed to read ABI file:", err)
		return
	}

	var abi abi.ABI
	if err = json.Unmarshal(abiJSON, &abi); err != nil {
		fmt.Println("failed to unmarshal ABI JSON:", err)
		return
	}

	var val *big.Int = big.NewInt(6)
	encoded, err := abi.Pack("set", val)
	if err != nil {
		fmt.Println("failed to pack ABI:", err)
		return
	}

	fmt.Println(hexutil.Encode(encoded))
}
```