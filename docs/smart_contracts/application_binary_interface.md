Application Binary Interface
============================

The Application Binary Interface (ABI) is a data encoding scheme used in Conflux for working with smart contracts. The types defined in the ABI are the same as those you encounter when writing [Smart Contracts](smart_contracts_overview.md) with Solidity - i.e. *uint8, ..., uint256, int8, ..., int256, bool, string,* etc.

Type mappings
-------------

Use [conflux-abigen]() to create contract binding for invoking with contract convinently.

The ABI type to native golng mappings are as follows:

| solidity types                               | go types                                                                           |
| -------------------------------------------- | ---------------------------------------------------------------------------------- |
| address                                      | common.Address                                                                     |
| uint8,uint16,uint32,uint64                   | uint8,uint16,uint32,uint64                                                         |
| uint24,uint40,uint48,uint56,uint72...uint256 | *big.Int                                                                           |
| int8,int16,int32,int64                       | int8,int16,int32,int64                                                             |
| int24,int40,int48,int56,int72...int256       | *big.Int                                                                           |
| fixed bytes (bytes1,bytes2...bytes32)        | [length]byte                                                                       |
| fixed type T array (T[length])               | [length]TG (TG is go type matched with solidty type T)                             |
| bytes                                        | []byte                                                                             |
| dynamic type T array T[]                     | []TG ((TG is go type matched with solidty type T))                                 |
| string                                       | string                                                                             |
| bool                                         | bool                                                                               |
| tuple                                        | struct eg:`[{"name": "balance","type": "uint256"}]` => `struct {Balance *big.Int}` |


*big.Int types have to be used for numeric types, as numeric types in Ethereum are 256 bit integer values.

Solidity structs will have a corresponding struct generated for them.  The names of the corresponding classes will be the same as the name of the struct in the Solidity contract i.e. `struct Foo` in your smart contract will be called `Foo` in the generated contract binding.