package sdk

import (
	"github.com/Conflux-Chain/go-conflux-sdk/types"
)

type RpcDebugClient struct {
	core *Client
}

func NewRpcDebugClient(core *Client) RpcDebugClient {
	return RpcDebugClient{core}
}

// return account ready + deferred transactions
func (c *RpcDebugClient) TxpoolGetAccountTransactions(address types.Address) (val []types.Transaction, err error) {
	err = c.core.CallRPC(&val, "txpool_accountTransactions", address)
	return
}
