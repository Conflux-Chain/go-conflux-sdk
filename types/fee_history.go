package types

import "github.com/ethereum/go-ethereum/common/hexutil"

type FeeHistory struct {
	OldestBlock      *hexutil.Big     `json:"oldest_block"`
	Reward           [][]*hexutil.Big `json:"reward,omitempty"`
	BaseFee          []*hexutil.Big   `json:"base_fee_per_gas,omitempty"`
	GasUsedRatio     []float64        `json:"gas_used_ratio"`
	BlobBaseFee      []*hexutil.Big   `json:"base_fee_per_blob_gas,omitempty"`
	BlobGasUsedRatio []float64        `json:"blob_gas_used_ratio,omitempty"`
}
