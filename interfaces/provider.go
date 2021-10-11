package interfaces

import (
	"context"

	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
)

type RpcProvider interface {
	Call(resultPtr interface{}, method string, args ...interface{}) error
	CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error
	BatchCall(b []rpc.BatchElem) error
	BatchCallContext(ctx context.Context, b []rpc.BatchElem) error
	Subscribe(ctx context.Context, namespace string, channel interface{}, args ...interface{}) (*rpc.ClientSubscription, error)
	Close()
}
