package primitives

import (
	"io"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/pkg/errors"
)

func MustRLPEncodeReceipt(receipt *types.TransactionReceipt) []byte {
	storageCollateralized, storageReleased := constructStorageChanges(receipt)

	val := []interface{}{
		receipt.AccumulatedGasUsed.ToInt(),
		receipt.GasFee.ToInt(),
		Bool(receipt.GasCoveredBySponsor),
		hexutil.MustDecode(string(receipt.LogsBloom)),
		convertLogs(receipt.Logs),
		uint8(receipt.MustGetOutcomeType()),
		Bool(receipt.StorageCoveredBySponsor),
		storageCollateralized,
		storageReleased,
	}

	if receipt.BurntGasFee != nil {
		val = append(val, receipt.BurntGasFee)
	}

	encoded, err := rlp.EncodeToBytes(val)
	if err != nil {
		panic(err)
	}
	return encoded
}

type StorageChange struct {
	Account     common.Address
	Collaterals uint64
}

func constructStorageChanges(receipt *types.TransactionReceipt) (collateralized, released []StorageChange) {
	for _, v := range receipt.StorageReleased {
		released = append(released, StorageChange{
			Account:     v.Address.MustGetCommonAddress(),
			Collaterals: uint64(v.Collaterals),
		})
	}

	if receipt.StorageCollateralized == 0 {
		return
	}

	var account cfxaddress.Address
	if receipt.StorageCoveredBySponsor {
		account = *receipt.To
	} else {
		account = receipt.From
	}

	collateralized = append(collateralized, StorageChange{
		Account:     account.MustGetCommonAddress(),
		Collaterals: uint64(receipt.StorageCollateralized),
	})

	return
}

const (
	LogSpaceNative   uint8 = 1
	LogSpaceEthereum uint8 = 2
)

type TxLog struct {
	Addr   common.Address
	Topics []common.Hash
	Data   []byte
	Space  uint8
}

// EncodeRLP implements the rlp.Encoder interface.
func (log TxLog) EncodeRLP(w io.Writer) error {
	switch log.Space {
	case LogSpaceNative:
		return rlp.Encode(w, []interface{}{log.Addr, log.Topics, log.Data})
	case LogSpaceEthereum:
		return rlp.Encode(w, []interface{}{log.Addr, log.Topics, log.Data, log.Space})
	default:
		return errors.Errorf("invalid log space %v", log.Space)
	}
}

func convertLogs(logs []types.Log) []TxLog {
	var result []TxLog

	for _, v := range logs {
		var topics []common.Hash
		for _, t := range v.Topics {
			topics = append(topics, *t.ToCommonHash())
		}

		var space uint8
		switch *v.Space {
		case types.SPACE_NATIVE:
			space = LogSpaceNative
		case types.SPACE_EVM:
			space = LogSpaceEthereum
		default:
			panic("invalid space in log entry")
		}

		result = append(result, TxLog{
			Addr:   v.Address.MustGetCommonAddress(),
			Topics: topics,
			Data:   v.Data,
			Space:  space,
		})
	}

	return result
}
