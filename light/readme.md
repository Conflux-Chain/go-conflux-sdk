# Light node contract on Conflux Network

This package aims to deploy light node contract on any other blockchain to verify any transaction or receipt from Conflux network.

## Contracts

Please refer to the [repository](https://github.com/Conflux-Chain/conflux-light-contracts) for more details.

## Relay Blocks

Off-chain service is required to relay blocks to light node contract for MPT root verification. Basically, there are 2 steps to relay blocks:

1. **PoS blocks**: indicates the last finalized PoW block.
2. **PoW blocks**: Once a PoS block relayed, all PoW blocks since the last PoS block will be relayed.

Note, any one could relay the blocks, since all blocks are cryptographically verified.

There is an available component `EvmRelayer` to relay blocks on eSpace:

```go
relayer, err := light.NewEvmRelayer(coreClient, evmClient, config)
panicIfErr(err, "Failed to create relayer on eSpace")
go relayer.Relay()
```

## Verify Receipt with Proof
Given a transaction hash, there is available API to generate receipt proof for both core space and eSpace:

```go
proof, err := light.GetReceiptProofEvm(coreClient, evmClient, txHash)
abiEncodedProof := proof.ABIEncode()
```

With ABI encoded receipt proof, user could verify against light node contract and get RLP encoded event logs.

```solidity
function verifyProofData(bytes memory receiptProof) external view returns (bool success, string memory message, bytes memory rlpLogs);
```