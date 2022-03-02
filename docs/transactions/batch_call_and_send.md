
## Batch Query Information and Send Transaction
When we need to query many pieces of information or send many transactions, we may need to send many requests to the RPC server, and it may cause request limitation and low efficiency. So we provided batch methods for you to send a batch request at one time to avoid this case and improve efficiency.

Please see example from [example_bulk](https://github.com/conflux-fans/go-conflux-sdk-examples/tree/main/example_bulk)
### Batch query information
1. New `BulkCaller`
2. `BulkCaller.Cfx().XXX` *(XXX means RPC methods)* to append request, and the returned result and error are pointers for saving results after requests are sent.
   > Besides `Cfx`, there are also `Debug`, `Trace`, `Pos` methods for acquiring RPC methods for the corresponding namespace
3. `BulkCaller.Execute` to send requests.
4. The result and error pointer of step 2 are filled by request results
5. `BulkCaller.Clear` to clear request cache for new bulk call action.

```golang
gasPrice, gasPriceErr := bulkCaller.Cfx().GetGasPrice()
nonce, nonceErr := bulkCaller.Cfx().GetNextNonce(<addresses>)
err = bulkCaller.Execute()
if err != nil {
	panic(fmt.Sprintf("%+v", err))
}
// *gasPriceErr means query `GetGasPrice` failed
if *gasPriceErr != nil{
    // handle error
}
// *nonceErr means query `GetNextNonce` failed
if *nonceErr != nil{
    // handle error
}
// gasPrice and nonce are queried result
fmt.Printf("gasPrice %v", gasPrice)
fmt.Printf("nonce %v", nonce)
```
[Example Location](https://github.com/conflux-fans/go-conflux-sdk-examples/blob/9074ff226371a3610e5f98cfb4bd32c4ae3d126e/example_bulk/main.go#L68)

### Batch call contract
1. New `BulkCaller`
2. Use [`abigen`](../cfxabigen.md) to generate contract binding
3. There is a struct called `XXXBulkCaller` *(XXX means your contract name)* for bulk call contract methods
4. `XXXBulkCaller.YourContractMethod` to append request to its first parameter which is `BulkCaller` instance created in step 1, and the returned result and error arepointersr for saving results after requests be sent.
5. `BulkCaller.Execute` to send requests.
6. The result and error pointer of step 4 are filled by request results
7. `BulkCaller.Clear` to clear request cache for new bulk call contract action.

It's ok to batch call normal RPC methods and contract calls by BulkCaller.

```golang
balance0, balance0Err := mTokenBulkCaller.BalanceOf(*bulkCaller, nil, <address0>)
balance1, balance1Err := mTokenBulkCaller.BalanceOf(*bulkCaller, nil, <address1>)
err = bulkCaller.Execute()
if err != nil {
	panic(fmt.Sprintf("%+v", err))
}
// *balance0Err means query `BalanceOf` failed
if *balance0Err != nil{
    // handle error
}
...
// balance0 and balance1 are queried result
fmt.Printf("balance0 %v", balance0)
```

[Example Location](https://github.com/conflux-fans/go-conflux-sdk-examples/blob/9074ff226371a3610e5f98cfb4bd32c4ae3d126e/example_bulk/main.go#L70)

### Batch send transaction
1. New `BulkSender`
2. `BulkSender.AppendTransaction` to append an unsigned transaction
3. `BulkSender.SignAndSend` to send requests. The transaction hashes and errors will be returned. All of them are slice with the same length of appended transactions.
4. `BulkSender.Clear` to clear request cache for new bulk send action.

```golang
bulkSender.
AppendTransaction(newTx(&froms[0], &tos[0], types.NewBigInt(100))).
AppendTransaction(newTx(&froms[1], &tos[1], types.NewBigInt(200), types.NewBigInt(3)))
if err := bulkSender.PopulateTransactions(); err != nil {
    // handle error
}

hashes, errors, err := bulkSender.SignAndSend()
if err != nil {
    // handle error
}

for i := 0; i < len(hashes); i++ {
    if errors[i] != nil {
        fmt.Printf("sign and send the %vth tx error %v\n", i, errors[i])
    } else {
        fmt.Printf("the %vth tx hash %v\n", i, hashes[i])
    }
}
```
[Example Location](https://github.com/conflux-fans/go-conflux-sdk-examples/blob/9074ff226371a3610e5f98cfb4bd32c4ae3d126e/example_bulk/main.go#L127)

### Batch send contract transaction
1. Use [`abigen`](../cfxabigen.md) to generate contract binding
2. There is a struct called `XXXBulkTransactor` *(XXX means your contract name)* for bulk send contract transactions
3. `BulkSender.SignAndSend` to send requests. The transaction hashes and errors will be returned. All of them are slice with the same length of appended transactions.
4. `BulkSender.Clear` to clear request cache for new bulk send action.

```golang
bulkSender.
AppendTransaction(mTokenBulkSender.Transfer(nil, tos[1].MustGetCommonAddress(), big.NewIn(1))).
AppendTransaction(mTokenBulkSender.Transfer(nil, tos[2].MustGetCommonAddress(), big.NewIn(2))).

if err := bulkSender.PopulateTransactions(); err != nil {
    // handle error
}

hashes, errors, err := bulkSender.SignAndSend()
if err != nil {
    // handle error
}

for i := 0; i < len(hashes); i++ {
    if errors[i] != nil {
        fmt.Printf("sign and send the %vth tx error %v\n", i, errors[i])
    } else {
        fmt.Printf("the %vth tx hash %v\n", i, hashes[i])
    }
}
```
[Example Location](https://github.com/conflux-fans/go-conflux-sdk-examples/blob/9074ff226371a3610e5f98cfb4bd32c4ae3d126e/example_bulk/main.go#L142)