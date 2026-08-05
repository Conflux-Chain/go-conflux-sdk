[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://github.com/Conflux-Chain/go-conflux-sdk/blob/master/LICENSE)
[![Documentation](https://img.shields.io/badge/Documentation-GoDoc-green.svg)](https://godoc.org/github.com/Conflux-Chain/go-conflux-sdk)
[![Build Status](https://travis-ci.org/Conflux-Chain/go-conflux-sdk.svg?branch=master)](https://travis-ci.org/Conflux-Chain/go-conflux-sdk)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/conflux-chain/go-conflux-sdk)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/conflux-chain/go-conflux-sdk)

# Conflux Golang API

The Conflux Golang API allows any Golang client to interact with a local or remote Conflux node based on JSON-RPC 2.0 protocol. With Conflux Golang API, users can easily manage accounts, send transactions, deploy smart contracts, and query blockchain information.

Please read the [documentation](https://confluxnetwork.gitbook.io/go-conflux-sdk) for more.

And read the API documentation from [here](https://pkg.go.dev/github.com/Conflux-Chain/go-conflux-sdk).

## Install go-conflux-sdk
```
go get github.com/Conflux-Chain/go-conflux-sdk
```
You can also add the Conflux Golang API into the vendor folder.
```
govendor fetch github.com/Conflux-Chain/go-conflux-sdk
```
## Use go-conflux-sdk

### Create Client
usd `sdk.NewClient` to creat a client for interact with conflux-rust node, the `sdk.ClientOption` is for setting `Account Manager` keystore folder path and retry options.
```golang
client, err := sdk.NewClient("https://test.confluxrpc.com", sdk.ClientOption{
    KeystorePath: "../context/keystore",
})
```
### Query RPC
```golang
epoch, err := client.GetEpochNumber()
```
### Handle Invalid params Errors
Conflux Core v3.0.3 usually puts the details of an `Invalid params` error in the RPC error `message`. For framework-level parameter errors, Core v3.1.0 and later use the fixed message `Invalid params` and put the details in `data`. `utils.RpcError.InvalidParamsDetail` handles both formats.

Convert the error returned by an SDK call with `utils.ToRpcError` before reading the details, because not every error is an RPC error:
```golang
_, err := client.GetEpochNumber()
if err != nil {
    rpcErr, conversionErr := utils.ToRpcError(err)
    if conversionErr != nil {
        // Handle a non-RPC error.
        return conversionErr
    }

    detail, ok := rpcErr.InvalidParamsDetail()
    if ok {
        // ok means the RPC error code is -32602; detail may still be empty.
        fmt.Printf("Invalid params: %s\n", detail)
    }
}
```
### Send Transaction
```golang
chainID, err := client.GetNetworkID()
if err!=nil {
    panic(err)
}

from, err :=client.AccountManger().GetDefault()
if err!=nil {
    panic(err)
}

utx, err := client.CreateUnsignedTransaction(*from, cfxaddress.MustNewFromHex("0x1cad0b19bb29d4674531d6f115237e16afce377d", chainID), types.NewBigInt(1), nil)
if err!=nil {
    panic(err)
}

txhash, err := client.SendTransaction(utx)
```
### Interact With Smart Contract
The most simple way to interact with contract is generator contract binding by `conflux-abigen`, see details from [here](./docs/cfxabigen.md)

