Working with smart contracts and `conflux-abigen`
--------------------------------------------------------------

`conflux-abigen` can auto-generate smart contract binding to deploy and interact with smart contracts.

## Install 
``` bash
$ go install github.com/Conflux-Chain/conflux-abigen/cmd/cfxabigen
```

## Usage

To generate the contract binding, compile your smart contract:

``` bash
$ solc <contract>.sol --bin --abi --optimize -o <output-dir>/
```

Then generate the binding code using the [conflux-abigen]():

``` bash
$ cfxabigen --abi /path/to/<smart-contract>.abi --bin /path/to/<smart-contract>.bin --pkg main  --out <smart-contract>.go
```

The generated code will be like:
``` golang
// DeployMyERC20Token deploys a new Conflux contract, binding an instance of MyERC20Token to it.
func DeployMyERC20Token(auth *bind.TransactOpts, ...) (*types.UnsignedTransaction, *types.Hash, *MyERC20Token, error) {
        ...
}

// NewMyERC20Token creates a new instance of MyERC20Token, bound to a specific deployed contract.
func NewMyERC20Token(address types.Address, backend bind.ContractBackend) (*MyERC20Token, error) {
        ...
}


// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_MyERC20Token *MyERC20TokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyERC20Token.contract.Transact(opts, "transfer", _to, _value)
}


// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_MyERC20Token *MyERC20TokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
        ...
}

```

Now you can create and deploy your smart contract use method `Deploy<YourContractName>`:

```golang
client, err := sdk.NewClient("https://test.confluxrpc.com", sdk.ClientOption{
	KeystorePath: "../keystore",
})
if err != nil {
	log.Fatal(err)
}       
err = client.AccountManager.UnlockDefault("hello")
if err != nil {
	log.Fatal(err)
}       
tx, hash, yourContract, err := DeployYourContract(<*bind.TransactOpts>, client,<PARAM1>,...,<PARAMN>)
if err != nil {
	panic(err)
}
// yourContract is deployed contract instance
```

Or use an existing contract:

```golang
yourContract, err := NewYourContract(<YourContractAddress>, client)
```

To transact with a smart contract:

```golang
unsignedTx, txHash, err := yourContract.someMethod(<*bind.TransactOpts>,<param1>,...)
```

To call a smart contract:

```golang
result,err := yourContract.someMethod(<*bind.CallOpts>, <param1>, ...);
```

For more information refer to  [conflux-abigen](..//cfxabigen.md#readme)

Smart contract examples
-----------------------

`conflux-abigen` applies a very simple smart contract exmaple at [conflux-abigen-example](https://github.com/conflux-fans/conflux-abigen-example)