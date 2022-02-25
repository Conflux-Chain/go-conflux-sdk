## The transaction nonce 

The nonce is an increasing numeric value which is used to uniquely identify transactions. A nonce can only be used once and until a transaction is mined, it is possible to send multiple versions of a transaction with the same nonce, however, once mined, any subsequent submissions will be rejected.

You can obtain the next available nonce via the [cfx_getNextNonce](https://developer.confluxnetwork.org/conflux-doc/docs/json_rpc#cfx_getnextnonce) method:

```golang
nonce,err := client.GetNextNonce(<address>, <epoch>)
```
If you want send transaction continuously, the better way is to obtain the next available nonce within the pending transactions in txpool via [txpool_nextNonce](https://developer.confluxnetwork.org/conflux-doc/docs/RPCs/txpool_rpc#txpool_nextnonce) method:
```golang
nonce,err := client.TxPool().NextNonce(<address>)
```

However, go-conflux-sdk apply function `Client.GetNextUsableNonce` for convinent, it will get nonce by `txpool_nextNonce` first and if failed then get nonce by `cfx_getNextNonce`
```golang
nonce,err := client.GetNextUsableNonce(<address>)
```

The nonce can then be used to create your transaction object
```golang
var utx types.UnsignedTransaction
utx.Nonce = nonce
```