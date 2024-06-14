// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package types

import (
	"fmt"
	"io"
	"math/big"

	"github.com/Conflux-Chain/go-conflux-sdk/types/cmptutil"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
)

// BlockHeader represents a block header in Conflux.
type BlockHeader struct {
	Hash                  Hash             `json:"hash"`
	ParentHash            Hash             `json:"parentHash"`
	Height                *hexutil.Big     `json:"height"`
	Miner                 Address          `json:"miner"`
	DeferredStateRoot     Hash             `json:"deferredStateRoot"`
	DeferredReceiptsRoot  Hash             `json:"deferredReceiptsRoot"`
	DeferredLogsBloomHash Hash             `json:"deferredLogsBloomHash"`
	Blame                 hexutil.Uint64   `json:"blame"`
	TransactionsRoot      Hash             `json:"transactionsRoot"`
	EpochNumber           *hexutil.Big     `json:"epochNumber"`
	BlockNumber           *hexutil.Big     `json:"blockNumber"`
	GasLimit              *hexutil.Big     `json:"gasLimit"`
	GasUsed               *hexutil.Big     `json:"gasUsed"`
	BaseFeePerGas         *hexutil.Big     `json:"baseFeePerGas"`
	Timestamp             *hexutil.Big     `json:"timestamp"`
	Difficulty            *hexutil.Big     `json:"difficulty"`
	PowQuality            *hexutil.Big     `json:"powQuality"`
	RefereeHashes         []Hash           `json:"refereeHashes"`
	Adaptive              bool             `json:"adaptive"`
	Nonce                 *hexutil.Big     `json:"nonce"`
	Size                  *hexutil.Big     `json:"size"`
	Custom                []cmptutil.Bytes `json:"custom"`
	PosReference          *Hash            `json:"posReference"`
}

func (bh *BlockHeader) toStructForRlp() any {
	if bh.BaseFeePerGas == nil {
		rbh := rlpEncodableBlockHeaderNormal{
			Hash:                  bh.Hash,
			ParentHash:            bh.ParentHash,
			Height:                bh.Height.ToInt(),
			Miner:                 bh.Miner,
			DeferredStateRoot:     bh.DeferredStateRoot,
			DeferredReceiptsRoot:  bh.DeferredReceiptsRoot,
			DeferredLogsBloomHash: bh.DeferredLogsBloomHash,
			Blame:                 bh.Blame,
			TransactionsRoot:      bh.TransactionsRoot,
			EpochNumber:           bh.EpochNumber.ToInt(),
			GasLimit:              bh.GasLimit.ToInt(),
			Timestamp:             bh.Timestamp.ToInt(),
			Difficulty:            bh.Difficulty.ToInt(),
			PowQuality:            bh.PowQuality.ToInt(),
			RefereeHashes:         bh.RefereeHashes,
			Adaptive:              bh.Adaptive,
			Nonce:                 bh.Nonce.ToInt(),
			Size:                  bh.Size.ToInt(),
			Custom:                bh.Custom,
			PosReference:          bh.PosReference,
		}

		if bh.BlockNumber != nil {
			rbh.BlockNumber = &rlpNilableBigInt{bh.BlockNumber.ToInt()}
		}

		if bh.GasUsed != nil {
			rbh.GasUsed = &rlpNilableBigInt{bh.GasUsed.ToInt()}
		}
		return rbh
	} else {
		rbh := rlpEncodeableBlockHeader1559{
			Hash:                  bh.Hash,
			ParentHash:            bh.ParentHash,
			Height:                bh.Height.ToInt(),
			Miner:                 bh.Miner,
			DeferredStateRoot:     bh.DeferredStateRoot,
			DeferredReceiptsRoot:  bh.DeferredReceiptsRoot,
			DeferredLogsBloomHash: bh.DeferredLogsBloomHash,
			Blame:                 bh.Blame,
			TransactionsRoot:      bh.TransactionsRoot,
			EpochNumber:           bh.EpochNumber.ToInt(),
			GasLimit:              bh.GasLimit.ToInt(),
			BaseFeePerGas:         bh.BaseFeePerGas.ToInt(),
			Timestamp:             bh.Timestamp.ToInt(),
			Difficulty:            bh.Difficulty.ToInt(),
			PowQuality:            bh.PowQuality.ToInt(),
			RefereeHashes:         bh.RefereeHashes,
			Adaptive:              bh.Adaptive,
			Nonce:                 bh.Nonce.ToInt(),
			Size:                  bh.Size.ToInt(),
			Custom:                bh.Custom,
			PosReference:          bh.PosReference,
		}

		if bh.BlockNumber != nil {
			rbh.BlockNumber = &rlpNilableBigInt{bh.BlockNumber.ToInt()}
		}

		if bh.GasUsed != nil {
			rbh.GasUsed = &rlpNilableBigInt{bh.GasUsed.ToInt()}
		}
		return rbh
	}
}

// rlpEncodableBlockHeaderNormal block header struct used for rlp encoding
type rlpEncodableBlockHeaderNormal struct {
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
	BlockNumber           *rlpNilableBigInt `rlp:"nil"`
	GasLimit              *big.Int
	GasUsed               *rlpNilableBigInt `rlp:"nil"`
	Timestamp             *big.Int
	Difficulty            *big.Int
	PowQuality            *big.Int
	RefereeHashes         []Hash
	Adaptive              bool
	Nonce                 *big.Int
	Size                  *big.Int
	Custom                []cmptutil.Bytes
	PosReference          *Hash `rlp:"nil"`
}

func (r *rlpEncodableBlockHeaderNormal) toRaw() *BlockHeader {
	bh := &BlockHeader{
		Hash:                  r.Hash,
		ParentHash:            r.ParentHash,
		Height:                (*hexutil.Big)(r.Height),
		Miner:                 r.Miner,
		DeferredStateRoot:     r.DeferredStateRoot,
		DeferredReceiptsRoot:  r.DeferredReceiptsRoot,
		DeferredLogsBloomHash: r.DeferredLogsBloomHash,
		Blame:                 r.Blame,
		TransactionsRoot:      r.TransactionsRoot,
		EpochNumber:           (*hexutil.Big)(r.EpochNumber),
		GasLimit:              (*hexutil.Big)(r.GasLimit),
		BaseFeePerGas:         nil,
		Timestamp:             (*hexutil.Big)(r.Timestamp),
		Difficulty:            (*hexutil.Big)(r.Difficulty),
		PowQuality:            (*hexutil.Big)(r.PowQuality),
		RefereeHashes:         r.RefereeHashes,
		Adaptive:              r.Adaptive,
		Nonce:                 (*hexutil.Big)(r.Nonce),
		Size:                  (*hexutil.Big)(r.Size),
		Custom:                r.Custom,
		PosReference:          r.PosReference,
	}

	if r.BlockNumber != nil {
		bh.BlockNumber = (*hexutil.Big)(r.BlockNumber.Val)
	}

	if r.GasUsed != nil {
		bh.GasUsed = (*hexutil.Big)(r.GasUsed.Val)
	}
	return bh
}

type rlpEncodeableBlockHeader1559 struct {
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
	BlockNumber           *rlpNilableBigInt `rlp:"nil"`
	GasLimit              *big.Int
	GasUsed               *rlpNilableBigInt `rlp:"nil"`
	BaseFeePerGas         *big.Int
	Timestamp             *big.Int
	Difficulty            *big.Int
	PowQuality            *big.Int
	RefereeHashes         []Hash
	Adaptive              bool
	Nonce                 *big.Int
	Size                  *big.Int
	Custom                []cmptutil.Bytes
	PosReference          *Hash `rlp:"nil"`
}

func (r *rlpEncodeableBlockHeader1559) toRaw() *BlockHeader {
	bh := &BlockHeader{
		Hash:                  r.Hash,
		ParentHash:            r.ParentHash,
		Height:                (*hexutil.Big)(r.Height),
		Miner:                 r.Miner,
		DeferredStateRoot:     r.DeferredStateRoot,
		DeferredReceiptsRoot:  r.DeferredReceiptsRoot,
		DeferredLogsBloomHash: r.DeferredLogsBloomHash,
		Blame:                 r.Blame,
		TransactionsRoot:      r.TransactionsRoot,
		EpochNumber:           (*hexutil.Big)(r.EpochNumber),
		GasLimit:              (*hexutil.Big)(r.GasLimit),
		BaseFeePerGas:         (*hexutil.Big)(r.BaseFeePerGas),
		Timestamp:             (*hexutil.Big)(r.Timestamp),
		Difficulty:            (*hexutil.Big)(r.Difficulty),
		PowQuality:            (*hexutil.Big)(r.PowQuality),
		RefereeHashes:         r.RefereeHashes,
		Adaptive:              r.Adaptive,
		Nonce:                 (*hexutil.Big)(r.Nonce),
		Size:                  (*hexutil.Big)(r.Size),
		Custom:                r.Custom,
		PosReference:          r.PosReference,
	}

	if r.BlockNumber != nil {
		bh.BlockNumber = (*hexutil.Big)(r.BlockNumber.Val)
	}

	if r.GasUsed != nil {
		bh.GasUsed = (*hexutil.Big)(r.GasUsed.Val)
	}
	return bh
}

// EncodeRLP implements the rlp.Encoder interface.
func (bh BlockHeader) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, bh.toStructForRlp())
}

// DecodeRLP implements the rlp.Decoder interface.
func (bh *BlockHeader) DecodeRLP(r *rlp.Stream) error {
	bytes, err := r.Raw()
	if err != nil {
		return err
	}

	var list []any
	err = rlp.DecodeBytes(bytes, &list)
	if err != nil {
		return err
	}

	switch len(list) {
	case 22:
		var rbh rlpEncodableBlockHeaderNormal
		if err := rlp.DecodeBytes(bytes, &rbh); err != nil {
			return err
		}
		*bh = *rbh.toRaw()

	case 23:
		var rbh rlpEncodeableBlockHeader1559
		if err := rlp.DecodeBytes(bytes, &rbh); err != nil {
			return err
		}
		*bh = *rbh.toRaw()
	default:
		return fmt.Errorf("invalid block header RLP code length: %d", len(list))
	}

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
