package bulk

import (
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	rpc "github.com/openweb3/go-rpc-provider"
)

func get1stEpochIfy(epoch []*types.Epoch) *types.Epoch {
	var realEpoch *types.Epoch
	if len(epoch) > 0 {
		realEpoch = epoch[0]
	}
	return realEpoch
}

func get1stU64Ify(values []uint64) *hexutil.Uint64 {
	if len(values) > 0 {
		_value := hexutil.Uint64(values[0])
		return &_value
	}
	return nil
}

func newBatchElem(result interface{}, method string, args ...interface{}) rpc.BatchElem {
	return rpc.BatchElem{
		Method: method,
		Result: &result,
		Args:   args,
	}
}
