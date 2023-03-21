package bulk

import (
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
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

// Check if user's 1st tx in tx pool is pending, pending means not on "ready" state, maybe notEnoughCash/outdatedStatus/nonceFuture/epochHeightOutBound.
func CheckIfUser1stTxIsPending(bulkCaller *BulkCaller, user []cfxaddress.Address) (map[string]bool, error) {
	type pendingTxRes struct {
		res *types.AccountPendingTransactions
		err *error
	}

	senderPendingTxRes := make(map[string]*pendingTxRes)
	for _, u := range user {
		senderPendingTxRes[u.String()] = &pendingTxRes{}
	}

	// bulkCaller := NewBulkCaller(b.signableCaller)
	for user, pendingTxRes := range senderPendingTxRes {
		// logrus.WithField("user", user).Info("ready to check pending result")
		res, err := bulkCaller.GetAccountPendingTransactions(cfxaddress.MustNew(user), nil, nil)
		pendingTxRes.res = res
		pendingTxRes.err = err
	}

	// err means timeout
	if err := bulkCaller.Execute(); err != nil {
		return nil, err
	}

	result := make(map[string]bool)
	for user, v := range senderPendingTxRes {
		if v.res.PendingCount > 0 && v.res.FirstTxStatus != nil {
			isPending, _ := v.res.FirstTxStatus.IsPending()
			result[user] = isPending
		}
	}
	return result, nil
}
