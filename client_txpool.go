package sdk

import (
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type RpcTxpoolClient struct {
	core *Client
}

func NewRpcTxpoolClient(core *Client) RpcTxpoolClient {
	return RpcTxpoolClient{core}
}

func (c *RpcTxpoolClient) Status() (val types.TxPoolStatus, err error) {
	err = c.core.CallRPC(&val, "txpool_status")
	return
}

func (c *RpcTxpoolClient) NextNonce(address types.Address) (val *hexutil.Big, err error) {
	err = c.core.CallRPC(&val, "txpool_nextNonce", address)
	return
}

func (c *RpcTxpoolClient) TransactionByAddressAndNonce(address types.Address, nonce *hexutil.Big) (val *types.Transaction, err error) {
	err = c.core.CallRPC(&val, "txpool_transactionByAddressAndNonce", address, nonce)
	return
}

func (c *RpcTxpoolClient) PendingNonceRange(address types.Address) (val types.TxPoolPendingNonceRange, err error) {
	err = c.core.CallRPC(&val, "txpool_pendingNonceRange", address)
	return
}

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
