# Light node contract on Conflux Network

This package aims to deploy light node contract on any other blockchain to verify transaction or receipt from Conflux network.

## Contracts

Please refer to the [repository](https://github.com/Conflux-Chain/conflux-light-contracts) for more details.

## Relay Blocks

Off-chain service is required to relay blocks to light node contract for MPT root verification. Basically, there are two kinds of blocks to relay:

1. **PoS blocks**: indicates the last finalized PoW block.
2. **PoW blocks**: any PoW block could be optional relayed on chain for proof verification. In this way, less block headers in receipt proof is required to be verified on chain and reduce gas cost.

Note, any one could relay the blocks, since all blocks are cryptographically verified.

There is an available component `EvmRelayer` to relay PoS blocks on eSpace:

```go
relayer := light.NewEvmRelayer(coreClient, evmClient, config)
go relayer.Relay()
```

If necessary, `EvmRelayer` could be used to relay partial PoW blocks as well:

```go
relayer.RelayPoWBlocks(headers)
```

## Verify Receipt with Proof
Given a transaction hash, there is available API to generate receipt proof for eSpace.

```go
generator := light.NewProofGenerator(coreClient, evmClient, lightNodeContract)
proof, err := generator.CreateReceiptProofEvm(txHash)
// Handle error
abiEncodedProof := proof.ABIEncode()
```

If there're too many PoW blocks in proof, e.g. 30, client could relay partial PoW blocks on chain at first, so as to avoid `OutOfGas` issue.

```go
proof, err := generator.CreateReceiptProofEvm(txHash)
// Handle error
for len(proof.Headers) > 30 {
    index := len(proof.Headers) - 30
    headers := proof.Headers[index:]
    if err = relayer.RelayPoWBlocks(headers); err != nil {
        // Handle error
    }
    proof.Headers = proof.Headers[:index]
}
abiEncodedProof := proof.ABIEncode()
```

With ABI encoded receipt proof, user could verify against light node contract and get RLP encoded event logs.

```solidity
function verifyProofData(bytes memory receiptProof) external view returns (bool success, string memory message, bytes memory rlpLogs);
```