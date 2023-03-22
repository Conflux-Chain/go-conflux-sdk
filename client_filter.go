package sdk

import (
	"github.com/Conflux-Chain/go-conflux-sdk/types"
)

type RpcFilterClient struct {
	core *Client
}

// NewRpcPosClient creates a new RpcPosClient instance.
func NewRpcFilterClient(core *Client) RpcPosClient {
	return RpcPosClient{core}
}

func (c *RpcFilterClient) NewFilter() (filterId types.H128, err error) {
	err = c.core.CallRPC(&filterId, "cfx_newFilter")
	return
}

func (c *RpcFilterClient) NewBlockFilter() (filterId types.H128, err error) {
	err = c.core.CallRPC(&filterId, "cfx_newBlockFilter")
	return
}

func (c *RpcFilterClient) NewPendingTransactionFilter() (filterId types.H128, err error) {
	err = c.core.CallRPC(&filterId, "cfx_newPendingTransactionFilter")
	return
}

func (c *RpcFilterClient) GetFilterChanges(filterId types.H128) (cfxFilterChanges *types.CfxFilterChanges, err error) {
	err = c.core.CallRPC(&cfxFilterChanges, "cfx_getFilterChanges", filterId)
	return

}

func (c *RpcFilterClient) GetFilterLogs(filterID types.H128) (logs []types.Log, err error) {
	err = c.core.CallRPC(&logs, "cfx_getFilterLogs", filterID)
	return
}

func (c *RpcFilterClient) UninstallFilter(filterId types.H128) (isUninstalled bool, err error) {
	err = c.core.CallRPC(&isUninstalled, "cfx_uninstallFilter", filterId)
	return
}
