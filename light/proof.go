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

var ErrTransactionExecutionFailed = errors.New("transaction execution failed")

// GetReceiptProofCore returns the receipt proof for specified `txHash` on core space.
//
// If receipt not found, it will return nil and requires client to retry later.
//
// If transaction execution failed, it will return `ErrTransactionExecutionFailed`.
func GetReceiptProofCore(client *sdk.Client, txHash types.Hash) (*contract.TypesReceiptProof, error) {
	receipt, err := client.GetTransactionReceipt(txHash)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get transaction receipt")
	}

	if receipt == nil {
		return nil, nil
	}

	// only return proof for success transaction
	if uint8(receipt.OutcomeStatus) != uint8(enums.NATIVE_SPACE_SUCCESS) {
		return nil, ErrTransactionExecutionFailed
	}

	// rare case that requires to retry
	if receipt.EpochNumber == nil {
		return nil, nil
	}

	return GetReceiptProof(client, uint64(*receipt.EpochNumber), txHash.String())
}

// GetReceiptProofEvm returns the receipt proof for specified `txHash` on eSpace.
//
// If receipt not found, it will return nil and requires client to retry later.
//
// If transaction execution failed, it will return `ErrTransactionExecutionFailed`.
func GetReceiptProofEvm(coreClient *sdk.Client, evmClient *web3go.Client, txHash common.Hash) (*contract.TypesReceiptProof, error) {
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

	return GetReceiptProof(coreClient, receipt.BlockNumber, txHash.Hex())
}

// GetReceiptProof returns the receipt proof for specified `epochNumber` and `txHash` on either core space or eSpace.
//
// If receipt not found, it will return nil and requires client to retry later.
//
// If transaction execution failed, it will return `ErrTransactionExecutionFailed`.
func GetReceiptProof(client *sdk.Client, epochNumber uint64, txHash string) (*contract.TypesReceiptProof, error) {
	epoch := types.NewEpochNumberUint64(epochNumber)
	epochOrHash := types.NewEpochOrBlockHashWithEpoch(epoch)

	epochReceipts, err := client.Debug().GetEpochReceipts(*epochOrHash, true)
	if err != nil {
		return nil, errors.WithMessagef(err, "Failed to get receipts by epoch number %v", epochNumber)
	}

	blockIndex, receipt := matchReceipt(epochReceipts, txHash)
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

	return &contract.TypesReceiptProof{
		EpochNumber:  new(big.Int).SetUint64(epochNumber),
		BlockIndex:   blockIndexKey,
		BlockProof:   mpt.ConvertProofNode(blockProof),
		ReceiptsRoot: receiptsRoot,
		Index:        receiptKey,
		Receipt:      contract.ConvertReceipt(receipt),
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
