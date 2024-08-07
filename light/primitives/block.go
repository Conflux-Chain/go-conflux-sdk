package primitives

import (
	"io"
	"math/big"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

const cip112Epoch = uint64(79050000)

type BlockHeader struct {
	raw              *types.BlockSummary
	evmBaseFeePerGas *big.Int
}

func MustRLPEncodeBlock(coreBlock *types.BlockSummary, evmBaseFeePerGas *big.Int) []byte {
	if coreBlock.BaseFeePerGas != nil && evmBaseFeePerGas == nil {
		panic("EVM base fee per gas empty")
	}

	encoded, err := rlp.EncodeToBytes(BlockHeader{coreBlock, evmBaseFeePerGas})
	if err != nil {
		panic(err)
	}

	return encoded
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
		list = append(list, rlpEncodeOptionSome(*header.raw.PosReference.ToCommonHash()))
	}

	if header.raw.BaseFeePerGas != nil {
		list = append(list, rlpEncodeOptionSome([]interface{}{
			header.raw.BaseFeePerGas.ToInt(),
			header.evmBaseFeePerGas,
		}))
	}

	for _, v := range header.raw.Custom {
		if header.raw.EpochNumber.ToInt().Uint64() >= cip112Epoch {
			list = append(list, v.ToBytes())
		} else {
			list = append(list, rlp.RawValue(v.ToBytes()))
		}
	}

	return rlp.Encode(w, list)
}

// simulate RLP encoding for rust Option type
func rlpEncodeOptionSome(v interface{}) interface{} {
	return []interface{}{v}
}
