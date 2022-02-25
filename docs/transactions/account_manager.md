## Manage Accounts
Use `AccountManager` struct to manage accounts at the local machine. Account will be saved as keystore file in specified folder when creates account.

Specify folder for save keystore files to create Account Manager, we need pass networkID be parameter because of conflux address display dependent on it.
```golang
am := new AccountManager(<Your_Keystore_Folder>, <NetworkID>)
```

Feature
----------
- Create an account.
- Import an account by keystore file.
- Import an account by private key.
- Export an account to be private key.
- Update an account.
- Delete an account.
- List all accounts.
- Unlock/Lock an account.
- Sign a transaction.

See API introduction from [API Doc]()