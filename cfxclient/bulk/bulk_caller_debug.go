package bulk

import (
	"fmt"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/utils"
)

// BulkDebugCaller used for bulk call rpc in one request to improve efficiency
type BulkDebugCaller BulkCallerCore

// NewBulkDebugCaller creates new BulkDebugCaller instance
func NewBulkDebugCaller(core BulkCallerCore) *BulkDebugCaller {
	return (*BulkDebugCaller)(&core)
}

// Execute sends all rpc requests in queue by rpc call "batch" on one request
func (b *BulkDebugCaller) Execute() ([]error, error) {
	return batchCall(b.caller, b.batchElems, nil)
}

// GetEpochReceiptsByEpochNumber returns epoch receipts by epoch number
func (client *BulkDebugCaller) GetEpochReceipts(epoch types.EpochOrBlockHash, include_eth_recepits ...bool) (*[][]types.TransactionReceipt, *error) {
	result := new([][]types.TransactionReceipt)
	err := new(error)

	includeEth := utils.Get1stBoolIfy(include_eth_recepits)
	elem := newBatchElem(result, "cfx_getEpochReceipts", epoch, includeEth)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)

	return result, err
}

// GetEpochReceiptsByPivotBlockHash returns epoch receipts by pivot block hash
func (client *BulkDebugCaller) GetEpochReceiptsByPivotBlockHash(hash types.Hash) (*[][]types.TransactionReceipt, *error) {
	result := new([][]types.TransactionReceipt)
	err := new(error)

	elem := newBatchElem(result, "cfx_getEpochReceipts", fmt.Sprintf("hash:%v", hash))
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)

	return result, err
}
