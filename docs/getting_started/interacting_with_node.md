Start sending requests
----------------------

To send requests:
```golang
client, err := sdk.NewClient("https://test.confluxrpc.com", sdk.ClientOption{
    KeystorePath: "../keystore",
})

epoch, err := client.GetEpochNumber()
```

Send Transaction
----------------------
Create UnsigendTransaction and send transaction by `Client.SendTransaction`, the Client will automatically populate the transaction and sign it before sending.

```golang
client, err := sdk.NewClient("https://test.confluxrpc.com", sdk.ClientOption{
    KeystorePath: "../keystore",
})

var utx types.UnsignedTransaction
utx.From = ... //use default account if not set
utx.To = ...
utx.value = ...
utx.Data = ...
// unlock account 
err = client.AccountManager.UnlockDefault(*utx.From, "hello")
if err != nil {
	log.Fatal(err)
}
txhash, err := client.SendTransaction(utx)
```

