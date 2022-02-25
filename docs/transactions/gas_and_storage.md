## Gas

When a transaction takes place in Conflux, a transaction cost must be paid to the client that executes the transaction on your behalf, committing the output of this transaction to the Conflux blockchain.

This cost is measured in gas, where gas is the number of instructions used to execute a transaction in the Conflux Virtual Machine. Please refer to the [Homestead
documentation](http://ethdocs.org/en/latest/contracts-and-transactions/account-types-gas-and-transactions.html?highlight=gas#what-is-gas) for detail.

What this means for you when working with Conflux clients is that there are two parameters that are used to dictate how much CFX you wish to spend in order for a transaction to complete:

_Gas price_

> This is the amount you are prepared in conflux per unit of gas. 

_Gas limit_

> This is the total amount of gas you are happy to spend on the
> transaction execution. There is an upper limit of how large a single
> transaction can be in an Conflux block which restricts this value
> typically to less then 30,000,000.

These parameters taken together dictate the maximum amount of CFX you are willing to spend on transaction costs. i.e. you can spend no more then gas price \* gas limit. The gas price can also affect how quickly a transaction takes place depending on what other transactions are available with a more profitable gas price for miners.

You may need to adjust these parameters to ensure that transactions take place in a timely manner.

## Storage Collateral

In addition to paying gas, you also need to spend storage collateral according to the storage occupancy in the smart contract, see details from [developer site](https://developer.confluxnetwork.org/introduction/en/conflux_storage)