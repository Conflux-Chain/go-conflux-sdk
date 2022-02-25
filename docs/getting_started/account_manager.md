# Manage Accounts

Use `AccountManager` struct to manage accounts at the local machine.
- Create/Import/Export/Update/Delete an account.
- List all accounts.
- Unlock/Lock an account.
- Sign a transaction.

## Create an AccountManager instance
Create account manager use folder path storing Keystore files and networkID because of [cfxaddress]() display independent.
```
    am, err := accounts.NewAccountManager("../keystore", 1)
```
## Get AccountManger from Client instance
```golang
    client, err := sdk.NewClient("https://test.confluxrpc.com", sdk.ClientOption{
            KeystorePath: "../keystore",
        })

    am := client.AccountManger
```

## Sign Transaction
The client will automatically sign a transaction when using `Client.SendTransaction` to send a transaction, you also could sign a transaction manually

```golang

    am, err := accounts.NewAccountManager("../keystore", 1)

    unSignedTx := types.UnsignedTransaction{...}
    am.SignAndEcodeTransactionWithPassphrase(unSignedTx, "passport")
```

See all apis from [AccountManager API Doc]()