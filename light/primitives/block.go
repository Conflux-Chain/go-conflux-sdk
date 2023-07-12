package primitives

import (
	"io"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type BlockHeader struct {
	raw *types.BlockSummary
}

func MustRLPEncodeBlock(block *types.BlockSummary) []byte {
	val := ConvertBlock(block)
	encoded, err := rlp.EncodeToBytes(val)
	if err != nil {
		panic(err)
	}
	return encoded
}

func ConvertBlock(block *types.BlockSummary) BlockHeader {
	return BlockHeader{block}
}

// EncodeRLP implements the rlp.Encoder interface.
func (header BlockHeader) EncodeRLP(w io.Writer) error {
	var adaptive uint64
	if header.raw.Adaptive {
		adaptive = 1
	}

	var referees []common.Hash
	for _, v := range header.raw.RefereeHashes {
		referees = append(referees, *v.ToCommonHash())
	}

	list := []interface{}{
		header.raw.ParentHash.ToCommonHash(),
		header.raw.Height.ToInt(),
		header.raw.Timestamp.ToInt(),
		header.raw.Miner.MustGetCommonAddress(),
		header.raw.TransactionsRoot.ToCommonHash(),
		header.raw.DeferredStateRoot.ToCommonHash(),
		header.raw.DeferredReceiptsRoot.ToCommonHash(),
		header.raw.DeferredLogsBloomHash.ToCommonHash(),
		header.raw.Blame,
		header.raw.Difficulty.ToInt(),
		adaptive,
		header.raw.GasLimit.ToInt(),
		referees,
		header.raw.Nonce.ToInt(),
	}

	if header.raw.PosReference != nil {
		// simulate RLP encoding for rust Option type
		item := []common.Hash{*header.raw.PosReference.ToCommonHash()}
		list = append(list, item)
	}

	for _, v := range header.raw.Custom {
		list = append(list, rlp.RawValue(v))
	}

	return rlp.Encode(w, list)
}
