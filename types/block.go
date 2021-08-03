// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package types

import (
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
)

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
	BlockNumber           *hexutil.Big   `json:"blockNumber"`
	GasLimit              *hexutil.Big   `json:"gasLimit"`
	GasUsed               *hexutil.Big   `json:"gasUsed"`
	Timestamp             *hexutil.Big   `json:"timestamp"`
	Difficulty            *hexutil.Big   `json:"difficulty"`
	PowQuality            *hexutil.Big   `json:"powQuality"`
	RefereeHashes         []Hash         `json:"refereeHashes"`
	Adaptive              bool           `json:"adaptive"`
	Nonce                 *hexutil.Big   `json:"nonce"`
	Size                  *hexutil.Big   `json:"size"`
	Custom                [][]byte       `json:"custom"`
}

// rlpEncodableBlockHeader block header struct used for rlp encoding
type rlpEncodableBlockHeader struct {
	Hash                  Hash
	ParentHash            Hash
	Height                *big.Int
	Miner                 Address
	DeferredStateRoot     Hash
	DeferredReceiptsRoot  Hash
	DeferredLogsBloomHash Hash
	Blame                 hexutil.Uint64
	TransactionsRoot      Hash
	EpochNumber           *big.Int
	GasLimit              *big.Int
	GasUsed               *big.Int
	Timestamp             *big.Int
	Difficulty            *big.Int
	PowQuality            *big.Int
	RefereeHashes         []Hash
	Adaptive              bool
	Nonce                 *big.Int
	Size                  *big.Int
	Custom                [][]byte
}

// EncodeRLP implements the rlp.Encoder interface.
func (bh BlockHeader) EncodeRLP(w io.Writer) error {
	rbh := rlpEncodableBlockHeader{
		bh.Hash, bh.ParentHash, bh.Height.ToInt(), bh.Miner, bh.DeferredStateRoot,
		bh.DeferredReceiptsRoot, bh.DeferredLogsBloomHash, bh.Blame, bh.TransactionsRoot,
		bh.EpochNumber.ToInt(), bh.GasLimit.ToInt(), bh.GasUsed.ToInt(), bh.Timestamp.ToInt(),
		bh.Difficulty.ToInt(), bh.PowQuality.ToInt(), bh.RefereeHashes, bh.Adaptive,
		bh.Nonce.ToInt(), bh.Size.ToInt(), bh.Custom,
	}

	return rlp.Encode(w, rbh)
}

// DecodeRLP implements the rlp.Decoder interface.
func (bh *BlockHeader) DecodeRLP(r *rlp.Stream) error {
	var rbh rlpEncodableBlockHeader
	if err := r.Decode(&rbh); err != nil {
		return err
	}

	bh.Hash, bh.ParentHash, bh.Height = rbh.Hash, rbh.ParentHash, (*hexutil.Big)(rbh.Height)
	bh.Miner, bh.DeferredStateRoot = rbh.Miner, rbh.DeferredStateRoot
	bh.DeferredReceiptsRoot, bh.DeferredLogsBloomHash = rbh.DeferredReceiptsRoot, rbh.DeferredLogsBloomHash
	bh.Blame, bh.TransactionsRoot = rbh.Blame, rbh.TransactionsRoot
	bh.EpochNumber, bh.GasLimit = (*hexutil.Big)(rbh.EpochNumber), (*hexutil.Big)(rbh.GasLimit)
	bh.GasUsed, bh.Timestamp = (*hexutil.Big)(rbh.GasUsed), (*hexutil.Big)(rbh.Timestamp)
	bh.Difficulty, bh.PowQuality = (*hexutil.Big)(rbh.Difficulty), (*hexutil.Big)(rbh.PowQuality)
	bh.RefereeHashes, bh.Adaptive = rbh.RefereeHashes, rbh.Adaptive
	bh.Nonce, bh.Size, bh.Custom = (*hexutil.Big)(rbh.Nonce), (*hexutil.Big)(rbh.Size), rbh.Custom

	return nil
}

// BlockSummary includes block header and a list of transaction hashes
type BlockSummary struct {
	BlockHeader
	Transactions []Hash `json:"transactions"`
}

// rlpEncodableBlockSummary block summary struct used for rlp encoding
type rlpEncodableBlockSummary struct {
	BlockHeader  BlockHeader
	Transactions []Hash
}

// EncodeRLP implements the rlp.Encoder interface.
func (bs BlockSummary) EncodeRLP(w io.Writer) error {
	rbs := rlpEncodableBlockSummary{
		bs.BlockHeader, bs.Transactions,
	}

	return rlp.Encode(w, rbs)
}

// DecodeRLP implements the rlp.Decoder interface.
func (bs *BlockSummary) DecodeRLP(r *rlp.Stream) error {
	var rbs rlpEncodableBlockSummary
	if err := r.Decode(&rbs); err != nil {
		return err
	}

	bs.BlockHeader = rbs.BlockHeader
	bs.Transactions = rbs.Transactions

	return nil
}

// Block represents a block in Conflux, including block header
// and a list of detailed transactions.
type Block struct {
	BlockHeader
	Transactions []Transaction `json:"transactions"`
}

// rlpEncodableBlock block struct used for rlp encoding
type rlpEncodableBlock struct {
	BlockHeader  BlockHeader
	Transactions []Transaction
}

// EncodeRLP implements the rlp.Encoder interface.
func (block Block) EncodeRLP(w io.Writer) error {
	rblock := rlpEncodableBlock{
		block.BlockHeader, block.Transactions,
	}

	return rlp.Encode(w, rblock)
}

// DecodeRLP implements the rlp.Decoder interface.
func (block *Block) DecodeRLP(r *rlp.Stream) error {
	var rblock rlpEncodableBlock
	if err := r.Decode(&rblock); err != nil {
		return err
	}

	block.BlockHeader = rblock.BlockHeader
	block.Transactions = rblock.Transactions

	return nil
}
