# Transactions

Broadly speaking there are three types transactions supported on Conflux:

1.  [Transfer of CFX from one party to another](transfer_eth.md#transfer-of-ether-from-one-party-to-another)
2.  [Creation of a smart contract](../cfxabigen.md#deploy-contract)
3.  [Transacting with a smart contract](../cfxabigen.md#transacting-with-an-conflux-contract)

To undertake any of these transactions, it is necessary to have CFX (the fuel of the Conflux blockchain) residing in the Conflux account which the transactions are taking place from. This is to pay for the [Gas Costs and Storage Collateral](gas_and_storage.md), they are the transaction execution cost for the Conflux client that performs the transaction on your behalf, committing the result to the Conflux blockchain. The storage collateral will be refund when the storage used in contract released. Instructions for obtaining CFX are described below in [Obtaining CFX](obtaining_cfx.md)

Additionally, it is possible to query the state of a smart contract, this is described in [Querying the state of a smart contract](../cfxabigen.md#accessing-an-conflux-contract)








