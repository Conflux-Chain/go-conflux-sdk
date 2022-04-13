[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://github.com/Conflux-Chain/go-conflux-sdk/blob/master/LICENSE)
[![Documentation](https://img.shields.io/badge/Documentation-GoDoc-green.svg)](https://godoc.org/github.com/Conflux-Chain/go-conflux-sdk)
[![Build Status](https://travis-ci.org/Conflux-Chain/go-conflux-sdk.svg?branch=master)](https://travis-ci.org/Conflux-Chain/go-conflux-sdk)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/conflux-chain/go-conflux-sdk)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/conflux-chain/go-conflux-sdk)

# Conflux Golang API

The Conflux Golang API allows any Golang client to interact with a local or remote Conflux node based on JSON-RPC 2.0 protocol. With Conflux Golang API, users can easily manage accounts, send transactions, deploy smart contracts, and query blockchain information.

Please read the [documentation](https://docs.confluxnetwork.org/go-conflux-sdk) for more.

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
The most simple way to interact with contract is generator contract binding by `conflux-abigen`, see details from [here](./cfxabigen.md)

