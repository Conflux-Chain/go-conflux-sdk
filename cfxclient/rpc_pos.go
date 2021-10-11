package cfxclient

import (
	postypes "github.com/Conflux-Chain/go-conflux-sdk/types/pos"
)

type RpcPosClient struct {
	core *ClientCore
}

func NewRpcPosClient(core *ClientCore) RpcPosClient {
	return RpcPosClient{core}
}

func (c *RpcPosClient) GetPosStatus() postypes.Status {
	return postypes.Status{}
}
