## Transfer of CFX from one party to another

The sending of CFX between two parties requires a minimal number of details of the transaction object:

`to`

 the destination wallet address

`value`

 the amount of CFX you wish to send to the destination address

```golang
utx, _ := client.CreateUnsignedTransaction(<from>,<to>,<value>,nil)
txhash, err := client.SendTransaction(utx)
```

wait transaction be executed
```golang
receipt, err := client.WaitForTransationReceipt(txhash)
```

OutcomeStatus of transaction receipt is 0 means transaction successfully