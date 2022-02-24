Compiling Solidity source code 
------------------------------

Compilation to bytecode is performed by the Solidity compiler, `solc`. You can install the compiler, locally following the instructions as per [the project documentation](http://solidity.readthedocs.io/en/develop/installing-solidity.html).

To compile the Solidity code run:

``` bash
$ solc <contract>.sol --bin --abi --optimize -o <output-dir>/
```

The `--bin` and `--abi` compiler arguments are both required to take full advantage of working with smart contracts from web3j.

`--bin`

 Outputs a Solidity binary file containing the hex-encoded binary to provide with the transaction request. This is required only for `deploy` and `isValid` [Solidity smart contract wrappers](construction_and_deployment.md#solidity-smart-contract-wrappers) methods.

`--abi`

 Outputs a Solidity [Application Binary Interface](application_binary_interface.md) file which details all of the publicly accessible contract methods and their associated parameters. These details along with the contract address are crucial for interacting with smart contracts. The ABI file is also used for the generation of [Solidity smart contract wrappers](construction_and_deployment.md#solidity-smart-contract-wrappers)

There is also a `--gas` argument for providing estimates of the [Gas](../transactions/gas.md) required to create a contract and transact with its methods.

Alternatively, you can write and compile Solidity code in your browser via the [browser-solidity](https://remix.ethereum.org/#optimize=false&evmVersion=null&version=soljson-v0.5.1+commit.c8a2cb62.js) project. browser-solidity is great for smaller smart contracts, but you may run into issues working with larger contracts.

You can also compile Solidity code via Ethereum clients such as Geth and OpenEthereum, using the JSON-RPC method [eth_compileSolidity](https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_compileSolidity) which is also supported in web3j. However, the Solidity compiler must be
installed on the client for this to work.

There are further options available, please refer to the [relevant section](https://ethereum-homestead.readthedocs.io/en/latest/contracts-and-transactions/contracts.html#compiling-a-contract) in the Homestead documentation.
