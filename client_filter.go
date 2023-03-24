package sdk

import (
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	rpc "github.com/openweb3/go-rpc-provider"
)

type RpcFilterClient struct {
	core *Client
}

func NewRpcFilterClient(core *Client) RpcPosClient {
	return RpcPosClient{core}
}

func (c *RpcFilterClient) NewFilter() (filterId *rpc.ID, err error) {
	err = c.core.CallRPC(&filterId, "cfx_newFilter")
	return
}

func (c *RpcFilterClient) NewBlockFilter() (filterId *rpc.ID, err error) {
	err = c.core.CallRPC(&filterId, "cfx_newBlockFilter")
	return
}

func (c *RpcFilterClient) NewPendingTransactionFilter() (filterId *rpc.ID, err error) {
	err = c.core.CallRPC(&filterId, "cfx_newPendingTransactionFilter")
	return
}

func (c *RpcFilterClient) GetFilterChanges(filterId rpc.ID) (cfxFilterChanges *types.CfxFilterChanges, err error) {
	err = c.core.CallRPC(&cfxFilterChanges, "cfx_getFilterChanges", filterId)
	return

}

func (c *RpcFilterClient) GetFilterLogs(filterID rpc.ID) (logs []types.Log, err error) {
	err = c.core.CallRPC(&logs, "cfx_getFilterLogs", filterID)
	return
}

func (c *RpcFilterClient) UninstallFilter(filterId rpc.ID) (isUninstalled bool, err error) {
	err = c.core.CallRPC(&isUninstalled, "cfx_uninstallFilter", filterId)
	return
}
