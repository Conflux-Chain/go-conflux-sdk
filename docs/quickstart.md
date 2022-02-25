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
The most simple way to interact with contract is generator contract binding by `conflux-abigen`, see details from [here](https://github.com/Conflux-Chain/conflux-abigen)
