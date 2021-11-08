package bulk

import (
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type BulkCustomCaller struct {
	BulkCallerCore
	outHandlers map[int]*OutputHandler
}

type OutputHandler func(out []byte) error

func NewBulkCustomCaller(core BulkCallerCore,
	outHandlers map[int]*OutputHandler,
) *BulkCustomCaller {
	return &BulkCustomCaller{core, outHandlers}
}

func (client *BulkCustomCaller) ContractCall(request types.CallRequest,
	epoch *types.Epoch,
	outDecoder OutputHandler,
	errPtr *error,
) {
	v := &hexutil.Bytes{}
	client.outHandlers[len(*client.batchElems)] = &outDecoder
	elem := newBatchElem(v, "cfx_call", request, epoch)
	client.BulkCallerCore.appendElemsAndError(elem, errPtr)
}
