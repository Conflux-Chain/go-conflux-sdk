package bulk

import (
	"github.com/Conflux-Chain/go-conflux-sdk/types"
)

// BulkTraceCaller used for bulk call rpc in one request to improve efficiency
type BulkTraceCaller BulkCallerCore

// NewBulkTraceCaller creates new BulkTraceCaller instance
func NewBulkTraceCaller(core BulkCallerCore) *BulkTraceCaller {
	return (*BulkTraceCaller)(&core)
}

// Execute sends all rpc requests in queue by rpc call "batch" on one request
func (b *BulkTraceCaller) Execute() ([]error, error) {
	return batchCall(b.caller, b.batchElems, nil)
}

func (client *BulkTraceCaller) GetBlockTraces(blockHash types.Hash) (*types.LocalizedBlockTrace, *error) {
	result := new(types.LocalizedBlockTrace)
	err := new(error)

	elem := newBatchElem(result, "trace_block", blockHash)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetFilterTraces returns all traces matching the provided filter.
func (client *BulkTraceCaller) FilterTraces(traceFilter types.TraceFilter) (*[]types.LocalizedTrace, *error) {
	result := new([]types.LocalizedTrace)
	err := new(error)

	elem := newBatchElem(&result, "trace_filter", traceFilter)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetTransactionTraces returns all traces produced at the given transaction.
func (client *BulkTraceCaller) GetTransactionTraces(txHash types.Hash) (*[]types.LocalizedTrace, *error) {
	result := new([]types.LocalizedTrace)
	err := new(error)

	elem := newBatchElem(&result, "trace_transaction", txHash)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}
