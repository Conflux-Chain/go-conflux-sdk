package types

import "github.com/ethereum/go-ethereum/common/hexutil"

// Status represents current blockchain status
type Status struct {
	BestHash        *Hash           `json:"bestHash"`
	BlockNumber     *hexutil.Uint64 `json:"blockNumber"`
	ChainID         *hexutil.Uint   `json:"chainId"`
	EpochNumber     *hexutil.Uint64 `json:"epochNumber"`
	PendingTxNumber *hexutil.Uint64 `json:"pendingTxNumber"`
}
