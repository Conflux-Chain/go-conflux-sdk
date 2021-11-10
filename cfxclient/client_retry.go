package client

import (
	"context"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/pkg/errors"
)

type rpcClientWithRetry struct {
	inner      *rpc.Client
	retryCount int
	interval   time.Duration
}

func (r *rpcClientWithRetry) Call(resultPtr interface{}, method string, args ...interface{}) error {
	return r.CallContext(context.Background(), resultPtr, method, args...)
}

func (r *rpcClientWithRetry) CallContext(ctx context.Context, resultPtr interface{}, method string, args ...interface{}) error {
	remain := r.retryCount
	for {

		err := r.inner.CallContext(ctx, resultPtr, method, args...)
		if err == nil {
			return nil
		}

		if utils.IsRPCJSONError(err) {
			return err
		}

		// fmt.Printf("remain retry count: %v\n", remain)
		if remain == 0 {
			return errors.Wrap(err, "rpc call timeout")
		}

		remain--

		if r.interval > 0 {
			time.Sleep(r.interval)
		}
	}
}

func (r *rpcClientWithRetry) BatchCall(b []rpc.BatchElem) error {
	return r.BatchCallContext(context.Background(), b)
}

func (r *rpcClientWithRetry) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	err := r.inner.BatchCallContext(ctx, b)
	if err == nil {
		return nil
	}

	if r.retryCount <= 0 {
		return err
	}

	remain := r.retryCount
	for {
		if err = r.inner.BatchCallContext(ctx, b); err == nil {
			return nil
		}

		remain--
		if remain == 0 {
			return errors.Wrap(err, "batch rpc call timeout")
		}

		if r.interval > 0 {
			time.Sleep(r.interval)
		}
	}
}

func (r *rpcClientWithRetry) Subscribe(ctx context.Context, namespace string, channel interface{}, args ...interface{}) (*rpc.ClientSubscription, error) {
	return r.inner.Subscribe(ctx, namespace, channel, args...)
}

func (r *rpcClientWithRetry) Close() {
	if r != nil && r.inner != nil {
		r.inner.Close()
	}
}
