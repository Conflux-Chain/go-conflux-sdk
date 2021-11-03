package bulk

import (
	"fmt"
	"reflect"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
)

type BulkCaller struct {
	caller      sdk.ClientOperator
	batchElems  *[]rpc.BatchElem
	outHandlers map[int]*OutputHandler

	cfx      *BulkCfxCaller
	customer *BulkCustomCaller
}

func NewBulkerCaller(rpcCaller sdk.ClientOperator) *BulkCaller {
	batchElems := make([]rpc.BatchElem, 0, 10)
	outHandlers := make(map[int]*OutputHandler)

	cfx := NewBulkCfxCaller(rpcCaller, &batchElems)
	customer := NewBulkCustomCaller(rpcCaller, &batchElems, outHandlers)

	return &BulkCaller{
		caller:      rpcCaller,
		batchElems:  &batchElems,
		outHandlers: outHandlers,

		cfx:      cfx,
		customer: customer,
	}
}

func (b *BulkCaller) Cfx() *BulkCfxCaller {
	return b.cfx
}

func (b *BulkCaller) Customer() *BulkCustomCaller {
	return b.customer
}

func (b *BulkCaller) Excute() ([]error, error) {
	// fmt.Printf("b: %#v", b)
	return batchCall(b.caller, b.batchElems, b.outHandlers)
}

func batchCall(caller sdk.ClientOperator,
	batchElems *[]rpc.BatchElem,
	outHandlers map[int]*OutputHandler,
) ([]error, error) {
	fmt.Printf("outHandlers %v\n", outHandlers)
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
			// fmt.Printf("v.Result %v\n", reflect.ValueOf(v.Result).Elem())
			// fmt.Printf("v.Result Type %v\n", reflect.TypeOf(v.Result))
			// fmt.Printf("v.Result To hexutil.Bytes %v\n", (*v.Result.(*interface{})).(*hexutil.Bytes))

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

			// switch rawOut.(type) {
			// case *hexutil.Bytes:
			// 	err := (*handler)(*rawOut.(*hexutil.Bytes))
			// 	if err != nil {
			// 		_errors[i] = errors.WithStack(err)
			// 	}
			// }
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
