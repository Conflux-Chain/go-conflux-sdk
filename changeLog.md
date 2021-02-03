# Go-conflux-sdk Change Log

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

