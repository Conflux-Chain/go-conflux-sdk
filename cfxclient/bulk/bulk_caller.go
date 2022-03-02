package bulk

import (
	"reflect"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/ethereum/go-ethereum/common/hexutil"
	rpc "github.com/openweb3/go-rpc-provider"
	"github.com/pkg/errors"
)

type BulkCallerCore struct {
	caller     sdk.ClientOperator
	batchElems *[]rpc.BatchElem
	errors     *[]*error
}

func NewBulkCallerCore(rpcCaller sdk.ClientOperator) BulkCallerCore {
	batchElems := make([]rpc.BatchElem, 0, 10)
	errors := make([]*error, 0, 10)

	return BulkCallerCore{
		caller:     rpcCaller,
		batchElems: &batchElems,
		errors:     &errors,
	}
}

func (b *BulkCallerCore) appendElemsAndError(elem rpc.BatchElem, err *error) {
	*b.batchElems = append(*b.batchElems, elem)
	*b.errors = append(*b.errors, err)
}

// BulkCaller used for bulk call rpc in one request to improve efficiency
type BulkCaller struct {
	BulkCallerCore

	outHandlers map[int]*OutputHandler
	*BulkCfxCaller
	customer *BulkCustomCaller

	debug  *BulkDebugCaller
	trace  *BulkTraceCaller
	pos    *BulkPosCaller
	txpool *BulkTxpoolCaller
}

// NewBulkCaller creates new bulk caller instance
func NewBulkCaller(rpcCaller sdk.ClientOperator) *BulkCaller {
	core := NewBulkCallerCore(rpcCaller)
	cfx := NewBulkCfxCaller(core)

	outHandlers := make(map[int]*OutputHandler)
	customer := NewBulkCustomCaller(core, outHandlers)

	return &BulkCaller{
		BulkCallerCore: core,
		outHandlers:    outHandlers,
		BulkCfxCaller:  cfx,
		customer:       customer,

		debug:  NewBulkDebugCaller(core),
		trace:  NewBulkTraceCaller(core),
		pos:    NewBulkPosCaller(core),
		txpool: NewBulkTxpoolCaller(core),
	}
}

// Cfx returns BulkCfxCaller for genereating "cfx" namespace relating rpc request
func (b *BulkCaller) Cfx() *BulkCfxCaller {
	return b.BulkCfxCaller
}

// Debug returns BulkDebugCaller for genereating "debug" namespace relating rpc request
func (b *BulkCaller) Debug() *BulkDebugCaller {
	return b.debug
}

// Trace returns BulkTraceCaller for genereating "trace" namespace relating rpc request
func (b *BulkCaller) Trace() *BulkTraceCaller {
	return b.trace
}

// Pos returns BulkTraceCaller for genereating "pos" namespace relating rpc request
func (b *BulkCaller) Pos() *BulkPosCaller {
	return b.pos
}

// Customer returns BulkCustomCaller for genereating contract relating rpc request which mainly for decoding contract call result with type *hexutil.Big to ABI defined types
func (b *BulkCaller) Customer() *BulkCustomCaller {
	return b.customer
}

// TxPool returns BulkTxpoolCaller for genereating "txpool" namespace relating rpc request
func (b *BulkCaller) Txpool() *BulkTxpoolCaller {
	return b.txpool
}

// Execute sends all rpc requests in queue by rpc call "batch" on one request
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

// Clear clear requests and errors in queue for new bulk call action
func (b *BulkCaller) Clear() {
	*b.BulkCallerCore.batchElems = (*b.BulkCallerCore.batchElems)[:0]
	*b.BulkCallerCore.errors = (*b.BulkCallerCore.errors)[:0]
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
