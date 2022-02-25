## Transaction mechanisms

When you have a valid account created with some CFX, you should sign the transaction first


## Transaction signing

Offline transaction signing allows you to sign a transaction using your [Account Manager]() within go-conflux-sdk, allowing you to have complete control over your private credentials. A transaction created offline can then be sent to any Ethereum client on the network, which will propagate the transaction out to other nodes, provided it is a valid transaction.

- Sign a transaction with an unlocked account: `AccountManager.SignTransaction(tx UnsignedTransaction)`

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

- Or sign a transaction with the passphrase for the locked account:

    `AccountManager.SignTransactionWithPassphrase(tx UnsignedTransaction, passphrase string)`

## Send Transaction
To send a transaction, you need to sign the transaction at a local machine offline as [below](./transaction_sign.md#offline-transaction-signing) and send the signed transaction to a local or remote Conflux node.

- Send an unsigned transaction

    `Client.SendTransaction(tx types.UnsignedTransaction)`

- Send an encoded transaction

    `Client.SendRawTransaction(rawData []byte)`

To send multiple transactions at a time, you can unlock the account at first, then send multiple transactions without the passphrase. To send a single transaction, you can just only send the transaction with the passphrase.