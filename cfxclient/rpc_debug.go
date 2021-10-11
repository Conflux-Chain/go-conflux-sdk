package cfxclient

import (
	"fmt"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	sdkErrors "github.com/Conflux-Chain/go-conflux-sdk/types/errors"
)

type RpcDebugClient struct {
	core *ClientCore
}

func NewRpcDebugClient(core *ClientCore) RpcDebugClient {
	return RpcDebugClient{core}
}

// =====Debug RPC=====

func (client *RpcDebugClient) GetEpochReceipts(epoch types.Epoch) (receipts [][]types.TransactionReceipt, err error) {
	err = client.core.wrappedCallRPC(&receipts, "cfx_getEpochReceipts", epoch)
	if ok, code := sdkErrors.DetectErrorCode(err); ok {
		err = sdkErrors.BusinessError{Code: code, Inner: err}
	}
	return
}

func (client *RpcDebugClient) GetEpochReceiptsByPivotBlockHash(hash types.Hash) (receipts [][]types.TransactionReceipt, err error) {
	err = client.core.wrappedCallRPC(&receipts, "cfx_getEpochReceipts", fmt.Sprintf("hash:%v", hash))
	if ok, code := sdkErrors.DetectErrorCode(err); ok {
		err = sdkErrors.BusinessError{Code: code, Inner: err}
	}
	return
}
