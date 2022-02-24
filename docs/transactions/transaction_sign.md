## Transaction mechanisms

When you have a valid account created with some CFX, you should sign the transaction first


## Offline transaction signing

Offline transaction signing allows you to sign a transaction using your [Account Manager]() within go-conflux-sdk, allowing you to have complete control over your private credentials. A transaction created offline can then be sent to any Ethereum client on the network, which will propagate the transaction out to other nodes, provided it is a valid transaction.

```golang
am := new AccountManager(...)

utx := types.UnsignedTransaction{...}

// you should unlock account before sign
err := am.Unlock(*utx.From, <passphrase>)
if err != nil{
        // handle error
}

am.SignTransaction(utx)
```

The simpliest way to sign and send a transaction is `Client.SendTransaction`, it will automatically sign transaction.

```golang
utx := types.UnsignedTransaction{...}
client.SendTransaction(utx)
```