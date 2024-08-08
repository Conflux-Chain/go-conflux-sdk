package primitives

import (
	"math/big"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

const cip112Epoch = uint64(79050000)

func MustRLPEncodeBlock(coreBlock *types.BlockSummary, evmBaseFeePerGas *big.Int) []byte {
	var adaptive uint64
	if coreBlock.Adaptive {
		adaptive = 1
	}

	var referees []common.Hash
	for _, v := range coreBlock.RefereeHashes {
		referees = append(referees, *v.ToCommonHash())
	}

	list := []interface{}{
		coreBlock.ParentHash.ToCommonHash(),
		coreBlock.Height.ToInt(),
		coreBlock.Timestamp.ToInt(),
		coreBlock.Miner.MustGetCommonAddress(),
		coreBlock.TransactionsRoot.ToCommonHash(),
		coreBlock.DeferredStateRoot.ToCommonHash(),
		coreBlock.DeferredReceiptsRoot.ToCommonHash(),
		coreBlock.DeferredLogsBloomHash.ToCommonHash(),
		coreBlock.Blame,
		coreBlock.Difficulty.ToInt(),
		adaptive,
		coreBlock.GasLimit.ToInt(),
		referees,
		coreBlock.Nonce.ToInt(),
	}

	if coreBlock.PosReference != nil {
		list = append(list, rlpEncodeOptionSome(*coreBlock.PosReference.ToCommonHash()))
	}

	if coreBlock.BaseFeePerGas != nil {
		if evmBaseFeePerGas == nil {
			panic("EVM base fee is empty")
		}

		list = append(list, rlpEncodeOptionSome([]interface{}{
			coreBlock.BaseFeePerGas.ToInt(),
			evmBaseFeePerGas,
		}))
	}

	for _, v := range coreBlock.Custom {
		if coreBlock.EpochNumber.ToInt().Uint64() >= cip112Epoch {
			list = append(list, v.ToBytes())
		} else {
			list = append(list, rlp.RawValue(v.ToBytes()))
		}
	}

	encoded, err := rlp.EncodeToBytes(list)
	if err != nil {
		panic(err)
	}

	return encoded
}

// simulate RLP encoding for rust Option type
func rlpEncodeOptionSome(v interface{}) interface{} {
	return []interface{}{v}
}
