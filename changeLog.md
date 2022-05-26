Go-conflux-sdk Change Log
============
v1.3.0
------------
-   Use MiddlewareableProvider in client for convinenet to HookCallContext/HookBatchCallContext
-   Remove client.UseCallRpcMiddleware and client.UseBatchCallRpcMiddleware, and use client.HookCallContext and client.HookBatchCallContext to replace them

v1.2.0
------------
-   add context in CallRpcHandlerFunc/BatchCallRpcHandlerFunc for hook context when Call/Batchcall
-   AccountManagerOperator support ImportKey and Export interface
> Note: It's not compatiable with previous version

V1.1.5
------------
-   support setting max connection number for client with http(s) protocal

V1.1.4
------------
-   Fix RPC GetPoSRewardByEpoch

V1.1.0
------------
-   Complete docs

v1.0.29
------------
-   Support unit convert utils
-   Upgrade go-ethereum dependency version to v1.10.15
-   Update RPC methods and structs

v1.0.27
------------
-   Set Block.Custom type to cmptutil.bytes which could json unmarshal both hex string or number array
-   Update PoS rpc type `Block`
-   Update PoW rpc type `Trace`

v1.0.24
------------
-   Support convert trace to tire
-   Forbid send transaction with to address in unkonwn type
-   Fix rlp decode transaction
-   Support rpc GetPosRewardByEpoch and GetPoSEconomics
-   Return unknown type for address start without '0x0,0x8,0x1'

v1.0.19
------------
-   Decode rpc error in Client.BatchCall

v1.0.18
------------
-   Decode rpc error in Client.Call

v1.0.17
------------
- Use tx pool pending nonce in `Client.ApplyUnsignedTransactionDefault` to ensure nonce is correct when continuous sending transactions
- Use bulk caller to populate transactions when bulk send transactions

v1.0.16
------------
- Support tx pool and debug RPC methods

v1.0.15
------------
- Add bulk caller and bulk sender for sending batch RPC requests by one request, see the example from [example_bulk](https://github.com/conflux-fans/go-conflux-sdk-examples/tree/main/example_bulk)
- Move example to independent repo [go-conflux-sdk-example](https://github.com/conflux-fans/go-conflux-sdk-examples)

v1.0.14
------------
- Add POS RPC

v1.0.13
------------
- Add API GetBlockSummaryByBlockNumber

v1.0.12
------------
- Fix test for Marshal/UnMarshal Block

v1.0.11
------------
- Add `blockNumber` to block related methods `cfx_getBlockByHash`, `cfx_getBlockByEpochNumber`, `cfx_getBlockByHashWithPivotAssumption` which need `Conflux-rust v1.1.5` or above.
- Add new RPC method `cfx_getBlockByBlockNumber`
- Refactor SubscribeLogs for avoiding losing the timing sequence of Chain-Reorg and Log
- Add variadic arguments support for rpc service

v1.0.10
------------
- Set default RPC request timeout to 30s
- Remove addition error msg in wrappedCallRPC
- Add method GetAccountPendingTransactions in client

v1.0.9
------------
- Apply middleware for hooking call RPC and batch call RPC
- Support set request RPC timeout in Client

v1.0.0
------------
Note: v1.0.0 is not incompatible with v0.x, the changes are below
- Change address format follow [CIP-37](https://github.com/Conflux-Chain/CIPs/blob/master/CIPs/cip-37.md)
- Unmarshal logfilter according to full node struct
- RPC functions follow the rule: input and output to be value or pointer according to whether it could be nil
- 
v0.4.11
------------
- Fix bug

v0.4.10
------------
- Support RPC block_trace
- Fix the amount of TIME_WAIT when concurrency request

v0.4.9
------------
- Support sdk.Client.GetSupplyInfo
- Support internal contract

