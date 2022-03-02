
## Interaca with Smart Contract

The simplest and recommended way is to use [conflux-abigen](../cfxabigen.md) to generate contract binding to deploy and invoke with contract

### [Deploy Contract](../cfxabigen.md#deploy-contract)

### [Accessing an Conflux contract](../cfxabigen.md#accessing-an-conflux-contract)

### [Transacting with an Conflux contract](../cfxabigen.md#transacting-with-an-conflux-contract)

### [Batch Accessing an Conflux contract](../cfxabigen.md#batch-accessing-an-conflux-contract)

### [Batch Transacting with an Conflux contract](../cfxabigen.md#batch-transacting-with-an-conflux-contract)

## Direct interact with go-conflux-sdk ***[Depreated]***

However, but not recommended, you also can use `Client.DeployContract` to deploy a contract or use `Client.GetContract` to get a contract by deployed address. Then you can use the contract instance to operate the contract, there are GetData/Call/SendTransaction. Please see [api document](https://github.com/Conflux-Chain/go-conflux-sdk/blob/master/api.md) and [example](https://github.com/conflux-fans/go-conflux-sdk-examples/tree/main/example_contract) for detail.