package bulk

import (
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// BulkTxpoolCaller used for bulk call rpc in one request to improve efficiency
type BulkTxpoolCaller BulkCallerCore

// NewBulkTxpoolCaller creates new BulkTxpoolCaller instance
func NewBulkTxpoolCaller(core BulkCallerCore) *BulkTxpoolCaller {
	return (*BulkTxpoolCaller)(&core)
}

// Execute sends all rpc requests in queue by rpc call "batch" on one request
func (b *BulkTxpoolCaller) Execute() ([]error, error) {
	return batchCall(b.caller, b.batchElems, nil)
}

func (client *BulkTxpoolCaller) Status() (*types.TxPoolStatus, *error) {
	result := new(types.TxPoolStatus)
	err := new(error)

	elem := newBatchElem(result, "txpool_status")
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

func (client *BulkTxpoolCaller) NextNonce(address types.Address) (*hexutil.Big, *error) {
	result := new(hexutil.Big)
	err := new(error)

	elem := newBatchElem(result, "txpool_nextNonce", address)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

func (client *BulkTxpoolCaller) TransactionByAddressAndNonce(address types.Address, nonce *hexutil.Big) (*types.Transaction, *error) {
	result := new(types.Transaction)
	err := new(error)

	elem := newBatchElem(result, "txpool_transactionByAddressAndNonce", address, nonce)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

func (client *BulkTxpoolCaller) PendingNonceRange(address types.Address) (*types.TxPoolPendingNonceRange, *error) {
	result := new(types.TxPoolPendingNonceRange)
	err := new(error)

	elem := newBatchElem(result, "txpool_pendingNonceRange", address)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

func (client *BulkTxpoolCaller) TxWithPoolInfo(hash types.Hash) (*types.TxWithPoolInfo, *error) {
	result := new(types.TxWithPoolInfo)
	err := new(error)

	elem := newBatchElem(result, "txpool_txWithPoolInfo", hash)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

/// Get transaction pending info by account address
func (client *BulkTxpoolCaller) AccountPendingInfo(address types.Address) (*types.AccountPendingInfo, *error) {
	result := new(types.AccountPendingInfo)
	err := new(error)

	elem := newBatchElem(result, "txpool_accountPendingInfo", address)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

/// Get transaction pending info by account address
func (client *BulkTxpoolCaller) AccountPendingTransactions(address types.Address, maybeStartNonce *hexutil.Big, maybeLimit *hexutil.Uint64) (*types.AccountPendingTransactions, *error) {
	result := new(types.AccountPendingTransactions)
	err := new(error)

	elem := newBatchElem(result, "txpool_accountPendingTransactions", address, maybeStartNonce, maybeLimit)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}
