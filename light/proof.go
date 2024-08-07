package light

import (
	"math/big"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/light/contract"
	"github.com/Conflux-Chain/go-conflux-sdk/light/mpt"
	"github.com/Conflux-Chain/go-conflux-sdk/light/primitives"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/enums"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go"
	evmTypes "github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
)

const deferredExecutionEpochs uint64 = 5

var ErrTransactionExecutionFailed = errors.New("transaction execution failed")

type ProofGenerator struct {
	coreClient *sdk.Client
	lightNode  *contract.LightNodeCaller
}

func NewProofGenerator(coreClient *sdk.Client, relayerClient *web3go.Client, lightNodeContract common.Address) *ProofGenerator {
	caller, _ := relayerClient.ToClientForContract()
	lightNode, err := contract.NewLightNodeCaller(lightNodeContract, caller)
	if err != nil {
		panic(err.Error())
	}

	return &ProofGenerator{
		coreClient: coreClient,
		lightNode:  lightNode,
	}
}

// CreateReceiptProofEvm returns the receipt proof for specified `txHash` on eSpace.
//
// If receipt not found, it will return nil and requires client to retry later.
//
// If transaction execution failed, it will return `ErrTransactionExecutionFailed`.
func (g *ProofGenerator) CreateReceiptProofEvm(evmClient *web3go.Client, txHash common.Hash) (*contract.TypesReceiptProof, error) {
	receipt, err := evmClient.Eth.TransactionReceipt(txHash)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get transaction receipt")
	}

	if receipt == nil {
		return nil, nil
	}

	if receipt.Status == nil {
		return nil, errors.New("receipt status is nil")
	}

	// only return proof for success transaction
	if uint8(*receipt.Status) != uint8(enums.EVM_SPACE_SUCCESS) {
		return nil, ErrTransactionExecutionFailed
	}

	pivot, err := g.lightNode.NearestPivot(nil, new(big.Int).SetUint64(receipt.BlockNumber+deferredExecutionEpochs))
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get nearest pivot on chain")
	}

	return CreateReceiptProofEvm(g.coreClient, evmClient, txHash, receipt.BlockNumber, pivot.Uint64())
}

func CreateReceiptProofEvm(coreClient *sdk.Client, evmClient *web3go.Client, txHash common.Hash, epochNumber uint64, pivot uint64) (*contract.TypesReceiptProof, error) {
	if epochNumber+deferredExecutionEpochs > pivot {
		return nil, errors.New("invalid pivot")
	}

	epoch := types.NewEpochNumberUint64(epochNumber)
	epochOrHash := types.NewEpochOrBlockHashWithEpoch(epoch)

	epochReceipts, err := coreClient.Debug().GetEpochReceipts(*epochOrHash, true)
	if err != nil {
		return nil, errors.WithMessagef(err, "Failed to get receipts by epoch number %v", epochNumber)
	}

	blockIndex, receipt := matchReceipt(epochReceipts, txHash.Hex())
	if receipt == nil {
		return nil, nil
	}

	if receipt.MustGetOutcomeType() != enums.TRANSACTION_OUTCOME_SUCCESS {
		return nil, ErrTransactionExecutionFailed
	}

	subtrees, root := CreateReceiptsMPT(epochReceipts)

	blockIndexKey := mpt.IndexToKey(blockIndex, len(subtrees))
	blockProof, ok := root.Proof(blockIndexKey)
	if !ok {
		return nil, errors.New("Failed to generate block proof")
	}

	receiptsRoot := subtrees[blockIndex].Hash()
	receiptKey := mpt.IndexToKey(int(receipt.Index), len(epochReceipts[blockIndex]))
	receiptProof, ok := subtrees[blockIndex].Proof(receiptKey)
	if !ok {
		return nil, errors.New("Failed to generate receipt proof")
	}

	var headers [][]byte
	for i := epochNumber + deferredExecutionEpochs; i <= pivot; i++ {
		coreBlock, err := coreClient.GetBlockSummaryByEpoch(types.NewEpochNumberUint64(i))
		if err != nil {
			return nil, errors.WithMessagef(err, "Failed to get core block summary by epoch %v", i)
		}

		if coreBlock == nil {
			return nil, errors.Errorf("Core block not found by epoch %v", i)
		}

		var evmBlock *evmTypes.Block
		if coreBlock.BaseFeePerGas != nil {
			evmBlock, err = evmClient.Eth.BlockByNumber(evmTypes.NewBlockNumber(int64(i)), false)
			if err != nil {
				return nil, errors.WithMessagef(err, "Failed to get evm block by block number %v", i)
			}

			if evmBlock == nil {
				return nil, errors.Errorf("Evm block not found by block number %v", i)
			}

			if evmBlock.BaseFeePerGas == nil {
				return nil, errors.Errorf("There is no base fee in evm block by number %v", i)
			}
		}

		headers = append(headers, primitives.MustRLPEncodeBlock(coreBlock, evmBlock.BaseFeePerGas))
	}

	return &contract.TypesReceiptProof{
		Headers:      headers,
		BlockIndex:   blockIndexKey,
		BlockProof:   mpt.ConvertProofNode(blockProof),
		ReceiptsRoot: receiptsRoot,
		Index:        receiptKey,
		Receipt:      primitives.MustRLPEncodeReceipt(receipt),
		ReceiptProof: mpt.ConvertProofNode(receiptProof),
	}, nil
}

func matchReceipt(epochReceipts [][]types.TransactionReceipt, txHash string) (blockIndex int, receipt *types.TransactionReceipt) {
	for i, blockReceipts := range epochReceipts {
		for _, v := range blockReceipts {
			if v.MustGetOutcomeType() == enums.TRANSACTION_OUTCOME_SKIPPED {
				continue
			}

			if v.TransactionHash.String() != txHash {
				continue
			}

			return i, &v
		}
	}

	return 0, nil
}
