package sdk

import (
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// RpcTxpoolClient used to access txpool namespace RPC of Conflux blockchain.
type RpcTxpoolClient struct {
	core *Client
}

// NewRpcTxpoolClient creates a new RpcTxpoolClient instance.
func NewRpcTxpoolClient(core *Client) RpcTxpoolClient {
	return RpcTxpoolClient{core}
}

// Status returns txpool status
func (c *RpcTxpoolClient) Status() (val types.TxPoolStatus, err error) {
	err = c.core.CallRPC(&val, "txpool_status")
	return
}

// NextNonce returns next nonce of account, including pending transactions
func (c *RpcTxpoolClient) NextNonce(address types.Address) (val *hexutil.Big, err error) {
	err = c.core.CallRPC(&val, "txpool_nextNonce", address)
	return
}

// TransactionByAddressAndNonce returns transaction info in txpool by account address and nonce
func (c *RpcTxpoolClient) TransactionByAddressAndNonce(address types.Address, nonce *hexutil.Big) (val *types.Transaction, err error) {
	err = c.core.CallRPC(&val, "txpool_transactionByAddressAndNonce", address, nonce)
	return
}

// PendingNonceRange returns pending nonce range in txpool of account
func (c *RpcTxpoolClient) PendingNonceRange(address types.Address) (val types.TxPoolPendingNonceRange, err error) {
	err = c.core.CallRPC(&val, "txpool_pendingNonceRange", address)
	return
}

// TxWithPoolInfo returns transaction with txpool info by transaction hash
func (c *RpcTxpoolClient) TxWithPoolInfo(hash types.Hash) (val types.TxWithPoolInfo, err error) {
	err = c.core.CallRPC(&val, "txpool_txWithPoolInfo", hash)
	return
}

/// Get transaction pending info by account address
func (c *RpcTxpoolClient) AccountPendingInfo(address types.Address) (val *types.AccountPendingInfo, err error) {
	err = c.core.CallRPC(&val, "txpool_accountPendingInfo", address)
	return
}

/// Get transaction pending info by account address
func (c *RpcTxpoolClient) AccountPendingTransactions(address types.Address, maybeStartNonce *hexutil.Big, maybeLimit *hexutil.Uint64) (val types.AccountPendingTransactions, err error) {
	err = c.core.CallRPC(&val, "txpool_accountPendingTransactions", address, maybeStartNonce, maybeLimit)
	return
}
