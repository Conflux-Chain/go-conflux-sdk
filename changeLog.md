# Go-conflux-sdk Change Log
## v1.0.13
- Add API GetBlockSummaryByBlockNumber
## v1.0.12
- Fix test for Marshal/UnMarshal Block
## v1.0.11
- Add `blockNumber` to block related methods `cfx_getBlockByHash`, `cfx_getBlockByEpochNumber`, `cfx_getBlockByHashWithPivotAssumption` which need `Conflux-rust v1.1.5` or above.
- Add new RPC method `cfx_getBlockByBlockNumber`
- Refactor SubscribeLogs for avoiding lossing timing sequence of Chain-Reorg and Log
- Add variadic arguments support for rpc service
## v1.0.10
- Set default rpc request timeout to 30s
- Remove addition error msg in wrappedCallRPC
- Add method GetAccountPendingTransactions in client
## v1.0.9
- Apply middleware for hooking call rpc and batch call rpc
- Support set request rpc timeout in Client
## v1.0.0
Note: v1.0.0 is not impatable with v0.x, the changes are
- Change address format follow [CIP-37](https://github.com/Conflux-Chain/CIPs/blob/master/CIPs/cip-37.md)
- Unmarshal logfilter according to full node struct
- RPC functions follow rule: input and output to be value or pointer according to whether it could be nil
## v0.4.11
- Fix bug

## v0.4.10
- Support rpc block_trace
- Fix amount of TIME_WAIT when concurrency request

## v0.4.9
- Support sdk.Client.GetSupplyInfo
- Support internal contract

