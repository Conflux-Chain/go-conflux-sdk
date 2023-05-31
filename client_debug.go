package sdk

import (
	"fmt"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	sdkErrors "github.com/Conflux-Chain/go-conflux-sdk/types/errors"
)

// RpcDebugClient used to access debug namespace RPC of Conflux blockchain.
type RpcDebugClient struct {
	core *Client
}

// NewRpcDebugClient creates a new RpcDebugClient instance.
func NewRpcDebugClient(core *Client) RpcDebugClient {
	return RpcDebugClient{core}
}

// TxpoolGetAccountTransactions returns account ready + deferred transactions
func (c *RpcDebugClient) TxpoolGetAccountTransactions(address types.Address) (val []types.Transaction, err error) {
	err = c.core.CallRPC(&val, "txpool_accountTransactions", address)
	return
}

// GetEpochReceiptsByEpochNumber returns epoch receipts by epoch number
func (c *RpcDebugClient) GetEpochReceipts(epoch types.EpochOrBlockHash, include_eth_recepits ...bool) (receipts [][]types.TransactionReceipt, err error) {
	includeEth := get1stBoolIfy(include_eth_recepits)
	err = c.core.CallRPC(&receipts, "cfx_getEpochReceipts", epoch, includeEth)
	if ok, code := sdkErrors.DetectErrorCode(err); ok {
		err = sdkErrors.BusinessError{Code: code, Inner: err}
	}
	return
}

// GetEpochReceiptsByPivotBlockHash returns epoch receipts by pivot block hash
func (c *RpcDebugClient) GetEpochReceiptsByPivotBlockHash(hash types.Hash) (receipts [][]types.TransactionReceipt, err error) {
	err = c.core.CallRPC(&receipts, "cfx_getEpochReceipts", fmt.Sprintf("hash:%v", hash))
	if ok, code := sdkErrors.DetectErrorCode(err); ok {
		err = sdkErrors.BusinessError{Code: code, Inner: err}
	}
	return
}

func (c *RpcDebugClient) GetEpochReceiptProofByTransaction(hash types.Hash) (proof *types.EpochReceiptProof, err error) {
	err = c.core.CallRPC(&proof, "debug_getEpochReceiptProofByTransaction", hash)
	if ok, code := sdkErrors.DetectErrorCode(err); ok {
		err = sdkErrors.BusinessError{Code: code, Inner: err}
	}
	return
}

func (c *RpcDebugClient) GetTransactionsByEpoch(epoch types.Epoch) (wrapTransactions []types.WrapTransaction, err error) {
	err = c.core.CallRPC(&wrapTransactions, "debug_getTransactionsByEpoch", epoch)
	if ok, code := sdkErrors.DetectErrorCode(err); ok {
		err = sdkErrors.BusinessError{Code: code, Inner: err}
	}
	return
}

func (c *RpcDebugClient) GetTransactionsByBlock(hash types.Hash) (wrapTransactions []types.WrapTransaction, err error) {
	err = c.core.CallRPC(&wrapTransactions, "debug_getTransactionsByBlock", hash)
	if ok, code := sdkErrors.DetectErrorCode(err); ok {
		err = sdkErrors.BusinessError{Code: code, Inner: err}
	}
	return
}
