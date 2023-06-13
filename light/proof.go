package light

import (
	"math/big"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/light/contract"
	"github.com/Conflux-Chain/go-conflux-sdk/light/mpt"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/enums"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go"
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

	return g.getReceiptProof(receipt.BlockNumber, txHash.Hex())
}

func (g *ProofGenerator) getReceiptProof(epochNumber uint64, txHash string) (*contract.TypesReceiptProof, error) {
	epoch := types.NewEpochNumberUint64(epochNumber)
	epochOrHash := types.NewEpochOrBlockHashWithEpoch(epoch)

	epochReceipts, err := g.coreClient.Debug().GetEpochReceipts(*epochOrHash, true)
	if err != nil {
		return nil, errors.WithMessagef(err, "Failed to get receipts by epoch number %v", epochNumber)
	}

	blockIndex, receipt := g.matchReceipt(epochReceipts, txHash)
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

	headers, err := g.getHeaderChain(epochNumber + deferredExecutionEpochs)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get header chain")
	}

	return &contract.TypesReceiptProof{
		Headers:      headers,
		BlockIndex:   blockIndexKey,
		BlockProof:   mpt.ConvertProofNode(blockProof),
		ReceiptsRoot: receiptsRoot,
		Index:        receiptKey,
		Receipt:      contract.ConvertReceipt(receipt),
		ReceiptProof: mpt.ConvertProofNode(receiptProof),
	}, nil
}

func (ProofGenerator) matchReceipt(epochReceipts [][]types.TransactionReceipt, txHash string) (blockIndex int, receipt *types.TransactionReceipt) {
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

func (g *ProofGenerator) getHeaderChain(epochNumber uint64) ([]contract.TypesBlockHeader, error) {
	var chain []contract.TypesBlockHeader

	pivot, err := g.lightNode.NearestPivot(nil, new(big.Int).SetUint64(epochNumber))
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get nearest pivot on chain")
	}

	// pivot should be >= epoch number

	for i := epochNumber; i <= pivot.Uint64(); i++ {
		block, err := g.coreClient.GetBlockSummaryByEpoch(types.NewEpochNumberUint64(i))
		if err != nil {
			return nil, errors.WithMessage(err, "Failed to get block summary by epoch")
		}

		chain = append(chain, contract.ConvertBlockHeader(block))
	}

	return chain, nil
}
