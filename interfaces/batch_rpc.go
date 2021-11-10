package interfaces

import (
	"math/big"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type BatchRpcInvoker interface {
	BatchGetTxByHashes(txhashes []types.Hash) (map[types.Hash]*types.Transaction, error)
	BatchGetBlockSummarys(blockhashes []types.Hash) (map[types.Hash]*types.BlockSummary, error)
	BatchGetBlockSummarysByNumber(blocknumbers []hexutil.Uint64) (map[hexutil.Uint64]*types.BlockSummary, error)
	BatchGetRawBlockConfirmationRisk(blockhashes []types.Hash) (map[types.Hash]*big.Int, error)
	BatchGetBlockConfirmationRisk(blockhashes []types.Hash) (map[types.Hash]*big.Float, error)
}
