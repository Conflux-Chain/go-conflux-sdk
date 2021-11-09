package bulk

import (
	"reflect"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
)

type BulkCallerCore struct {
	caller     sdk.ClientOperator
	batchElems *[]rpc.BatchElem
	errors     *[]*error
}

func (b *BulkCallerCore) appendElemsAndError(elem rpc.BatchElem, err *error) {
	*b.batchElems = append(*b.batchElems, elem)
	*b.errors = append(*b.errors, err)
}

type BulkCaller struct {
	BulkCallerCore

	outHandlers map[int]*OutputHandler
	*BulkCfxCaller
	customer *BulkCustomCaller
}

func NewBulkerCaller(rpcCaller sdk.ClientOperator) *BulkCaller {
	batchElems := make([]rpc.BatchElem, 0, 10)
	errors := make([]*error, 0, 10)
	outHandlers := make(map[int]*OutputHandler)

	core := BulkCallerCore{
		caller:     rpcCaller,
		batchElems: &batchElems,
		errors:     &errors,
	}

	cfx := NewBulkCfxCaller(core)
	customer := NewBulkCustomCaller(core, outHandlers)

	return &BulkCaller{
		BulkCallerCore: core,
		outHandlers:    outHandlers,
		BulkCfxCaller:  cfx,
		customer:       customer,
	}
}

func (b *BulkCaller) Cfx() *BulkCfxCaller {
	return b.BulkCfxCaller
}

func (b *BulkCaller) Customer() *BulkCustomCaller {
	return b.customer
}

func (b *BulkCaller) Execute() error {
	_errors, _err := batchCall(b.BulkCallerCore.caller, b.BulkCallerCore.batchElems, b.outHandlers)
	if _err != nil {
		return _err
	}
	for i, v := range _errors {
		errPtr := (*b.BulkCallerCore.errors)[i]
		*errPtr = v
	}
	return nil
}

func (b *BulkCaller) Clear() {
	*b.BulkCallerCore.batchElems = (*b.BulkCallerCore.batchElems)[:0]
}

func batchCall(caller sdk.ClientOperator,
	batchElems *[]rpc.BatchElem,
	outHandlers map[int]*OutputHandler,
) ([]error, error) {
	if len(*batchElems) == 0 {
		return nil, nil
	}

	err := caller.BatchCallRPC(*batchElems)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	_errors := make([]error, len(*batchElems))
	for i, v := range *batchElems {
		_errors[i] = v.Error
	}

	if outHandlers == nil {
		return _errors, nil
	}

	for i, v := range *batchElems {
		if v.Error != nil {
			continue
		}

		handler := outHandlers[i]
		if handler != nil {

			var rawOut interface{} = *v.Result.(*interface{})
			val, ok := rawOut.(*hexutil.Bytes)
			if !ok {
				_errors[i] = errors.Errorf("response result type must be *[]byte or *hexutil.Bytes, got %v", reflect.TypeOf(rawOut))
				continue
			}

			err := (*handler)(*val)
			if err != nil {
				_errors[i] = errors.WithStack(err)
			}
		}
	}
	return _errors, nil
}

func newBatchElem(result interface{}, method string, args ...interface{}) rpc.BatchElem {
	return rpc.BatchElem{
		Method: method,
		Result: &result,
		Args:   args,
	}
}

// func appendBatchElem()
