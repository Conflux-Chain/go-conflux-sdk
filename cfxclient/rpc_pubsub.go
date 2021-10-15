package client

import (
	"context"

	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
)

type RpcPubsubClient struct {
	core *ClientCore
}

func NewRpcPubsubClient(core *ClientCore) RpcPubsubClient {
	return RpcPubsubClient{core}
}

// SubscribeNewHeads subscribes all new block headers participating in the consensus.
func (client *RpcPubsubClient) SubscribeNewHeads(channel chan types.BlockHeader) (*rpc.ClientSubscription, error) {
	return client.core.rpcProvider.Subscribe(context.Background(), "cfx", channel, "newHeads")
}

// SubscribeEpochs subscribes consensus results: the total order of blocks, as expressed by a sequence of epochs. Currently subscriptionEpochType only support "latest_mined" and "latest_state"
func (client *RpcPubsubClient) SubscribeEpochs(channel chan types.WebsocketEpochResponse, subscriptionEpochType ...types.Epoch) (*rpc.ClientSubscription, error) {
	if len(subscriptionEpochType) > 0 {
		return client.core.rpcProvider.Subscribe(context.Background(), "cfx", channel, "epochs", subscriptionEpochType[0].String())
	}
	return client.core.rpcProvider.Subscribe(context.Background(), "cfx", channel, "epochs")
}

// SubscribeLogs subscribes all logs matching a certain filter, in order.
func (client *RpcPubsubClient) SubscribeLogs(channel chan types.SubscriptionLog, filter types.LogFilter) (*rpc.ClientSubscription, error) {
	return client.core.rpcProvider.Subscribe(context.Background(), "cfx", channel, "logs", filter)
}
