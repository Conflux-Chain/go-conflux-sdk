package bulk

import (
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	postypes "github.com/Conflux-Chain/go-conflux-sdk/types/pos"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// BulkPosCaller used for bulk call rpc in one request to improve efficiency
type BulkPosCaller BulkCallerCore

// NewBulkPosCaller creates new BulkPosCaller instance
func NewBulkPosCaller(core BulkCallerCore) *BulkPosCaller {
	return (*BulkPosCaller)(&core)
}

// Execute sends all rpc requests in queue by rpc call "batch" on one request
func (b *BulkPosCaller) Execute() ([]error, error) {
	return batchCall(b.caller, b.batchElems, nil)
}

// GetStatus returns pos chain status
func (client *BulkPosCaller) GetStatus() (*postypes.Status, *error) {
	result := new(postypes.Status)
	err := new(error)

	elem := newBatchElem(result, "pos_getStatus")
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetAccount returns account info at block
func (client *BulkPosCaller) GetAccount(address postypes.Address, blockNumber ...uint64) (*postypes.Account, *error) {
	result := new(postypes.Account)
	err := new(error)
	_view := get1stU64Ify(blockNumber)

	elem := newBatchElem(result, "pos_getAccount", address, _view)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetCommittee returns committee info at block
func (client *BulkPosCaller) GetCommittee(blockNumber ...uint64) (*postypes.CommitteeState, *error) {
	result := new(postypes.CommitteeState)
	err := new(error)
	_view := get1stU64Ify(blockNumber)

	elem := newBatchElem(result, "pos_getCommittee", _view)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetBlockByHash returns block info of block hash
func (client *BulkPosCaller) GetBlockByHash(hash types.Hash) (*postypes.Block, *error) {
	result := new(postypes.Block)
	err := new(error)

	elem := newBatchElem(result, "pos_getBlockByHash", hash)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetBlockByHash returns block at block number
func (client *BulkPosCaller) GetBlockByNumber(blockNumber postypes.BlockNumber) (*postypes.Block, *error) {
	result := new(postypes.Block)
	err := new(error)

	elem := newBatchElem(result, "pos_getBlockByNumber", blockNumber)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetTransactionByNumber returns transaction info of transaction number
func (client *BulkPosCaller) GetTransactionByNumber(txNumber uint64) (*postypes.Transaction, *error) {
	result := new(postypes.Transaction)
	err := new(error)

	elem := newBatchElem(result, "pos_getTransactionByNumber", hexutil.Uint64(txNumber))
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetRewardsByEpoch returns rewards of epoch
func (client *BulkPosCaller) GetRewardsByEpoch(epochNumber uint64) (*postypes.EpochReward, *error) {
	result := new(postypes.EpochReward)
	err := new(error)

	elem := newBatchElem(result, "pos_getRewardsByEpoch", hexutil.Uint64(epochNumber))
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}
