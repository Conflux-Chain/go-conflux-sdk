package bulk

import (
	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type BulkCustomCaller struct {
	caller      sdk.ClientOperator
	batchElems  *[]rpc.BatchElem
	outHandlers map[int]*OutputHandler
}

type OutputHandler func(out []byte) error

func NewBulkCustomCaller(caller sdk.ClientOperator,
	batchElems *[]rpc.BatchElem,
	outHandlers map[int]*OutputHandler,
) *BulkCustomCaller {
	return &BulkCustomCaller{caller, batchElems, outHandlers}
}

func (client *BulkCustomCaller) ContractCall(request types.CallRequest,
	epoch *types.Epoch,
	outDecoder OutputHandler,
) {
	v := &hexutil.Bytes{}
	client.outHandlers[len(*client.batchElems)] = &outDecoder
	*client.batchElems = append(*client.batchElems, newBatchElem(v, "cfx_call", request, epoch))
}
