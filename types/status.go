package types

import "github.com/ethereum/go-ethereum/common/hexutil"

// Status represents current blockchain status
type Status struct {
	BestHash        *Hash        `json:"bestHash"`
	BlockNumber     *hexutil.Big `json:"blockNumber"`
	ChainID         *hexutil.Big `json:"chainId"`
	EpochNumber     *hexutil.Big `json:"epochNumber"`
	PendingTxNumber int          `json:"pendingTxNumber"`
}
