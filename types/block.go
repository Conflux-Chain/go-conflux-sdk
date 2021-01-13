// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package types

import "github.com/ethereum/go-ethereum/common/hexutil"

// BlockHeader represents a block header in Conflux.
type BlockHeader struct {
	Hash                  Hash           `json:"hash"`
	ParentHash            Hash           `json:"parentHash"`
	Height                *hexutil.Big   `json:"height"`
	Miner                 Address        `json:"miner"`
	DeferredStateRoot     Hash           `json:"deferredStateRoot"`
	DeferredReceiptsRoot  Hash           `json:"deferredReceiptsRoot"`
	DeferredLogsBloomHash Hash           `json:"deferredLogsBloomHash"`
	Blame                 hexutil.Uint64 `json:"blame"`
	TransactionsRoot      Hash           `json:"transactionsRoot"`
	EpochNumber           *hexutil.Big   `json:"epochNumber"`
	GasLimit              *hexutil.Big   `json:"gasLimit"`
	GasUsed               *hexutil.Big   `json:"gasUsed"`
	Timestamp             *hexutil.Big   `json:"timestamp"`
	Difficulty            *hexutil.Big   `json:"difficulty"`
	PowQuality            *hexutil.Big   `json:"powQuality"`
	RefereeHashes         []Hash         `json:"refereeHashes"`
	Adaptive              bool           `json:"adaptive"`
	Nonce                 *hexutil.Big   `json:"nonce"`
	Size                  *hexutil.Big   `json:"size"`
}

// BlockSummary includes block header and a list of transaction hashes
type BlockSummary struct {
	BlockHeader
	Transactions []Hash `json:"transactions"`
}

// Block represents a block in Conflux, including block header
// and a list of detailed transactions.
type Block struct {
	BlockHeader
	Transactions []Transaction `json:"transactions"`
}
