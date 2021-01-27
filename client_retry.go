package sdk

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

	remain := r.retryCount
	for {

		err := r.inner.Call(resultPtr, method, args...)
		if err == nil {
			return nil
		}

		if utils.IsRPCJSONError(err) {
			return err
		}

		remain--
		// fmt.Printf("remain retry count: %v\n", remain)
		if remain == 0 {
			return errors.Wrap(err, "rpc call timeout")
		}

		if r.interval > 0 {
			time.Sleep(r.interval)
		}
	}
}

func (r *rpcClientWithRetry) BatchCall(b []rpc.BatchElem) error {
	err := r.inner.BatchCall(b)
	if err == nil {
		return nil
	}

	if r.retryCount <= 0 {
		return err
	}

	remain := r.retryCount
	for {
		if err = r.inner.BatchCall(b); err == nil {
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
