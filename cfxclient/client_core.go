// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package client

import (
	"context"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/interfaces"
	"github.com/Conflux-Chain/go-conflux-sdk/middleware"
	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/pkg/errors"
)

// Client represents a client to interact with Conflux blockchain.
type ClientCore struct {
	// AccountManager      sdk.AccountManagerOperator
	nodeURL     string
	rpcProvider interfaces.RpcProvider
	networkID   uint32
	// option              config
	callRpcHandler      middleware.CallRpcHandler
	batchCallRpcHandler middleware.BatchCallRpcHandler

	retryCount     int
	retryInterval  time.Duration
	requestTimeout time.Duration
}

// NewClientWithRetry creates a retryable new instance of Client with specified conflux node url and retry options.
//
// the clientOption.RetryInterval will be set to 1 second if pass 0
func NewClientCore(nodeURL string) (*ClientCore, error) {

	var client ClientCore
	client.nodeURL = nodeURL
	client.callRpcHandler = middleware.CallRpcHandlerFunc(client.callRpc)
	client.batchCallRpcHandler = middleware.BatchCallRpcHandlerFunc(client.batchCallRPC)
	client.SetRetry(0, 0)
	client.SetRequestTimeout(0)

	rpcClient, err := rpc.Dial(nodeURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial to fullnode")
	}

	client.rpcProvider = &rpcClientWithRetry{
		inner:      rpcClient,
		retryCount: client.retryCount,
		interval:   client.retryInterval,
	}

	// set networkId
	if client.networkID, err = client.getNetworkId(); err != nil {
		return nil, errors.Wrap(err, "failed to get networkID")
	}

	return &client, nil
}

func (c *ClientCore) SetRetry(retryCount int, retryInterval time.Duration) *ClientCore {
	if retryCount <= 0 {
		retryCount = 0
	}

	if retryInterval <= 0 {
		retryInterval = time.Second
	}

	c.retryCount = retryCount
	c.retryInterval = retryInterval
	return c
}

func (c *ClientCore) SetRequestTimeout(timeout time.Duration) *ClientCore {

	if timeout <= 0 {
		timeout = time.Second * 30
	}

	c.requestTimeout = timeout
	return c
}

// GetNodeURL returns node url
func (c *ClientCore) GetNodeURL() string {
	return c.nodeURL
}

func (c *ClientCore) getNetworkId() (uint32, error) {
	if c.networkID != 0 {
		return c.networkID, nil
	}

	status, err := c.getStatus()
	if err != nil {
		return 0, errors.Wrap(err, "failed to get status")
	}

	return uint32(status.NetworkID), nil
}

// GetStatus returns status of connecting conflux node
func (c *ClientCore) getStatus() (status types.Status, err error) {
	err = c.CallRPC(&status, "cfx_getStatus")
	return
}

// CallRPC performs a JSON-RPC call with the given arguments and unmarshals into
// result if no error occurred.
//
// The result must be a pointer so that package json can unmarshal into it. You
// can also pass nil, in which case the result is ignored.
//
// You could use UseCallRpcMiddleware to add middleware for hooking CallRPC
func (c *ClientCore) CallRPC(result interface{}, method string, args ...interface{}) error {
	return c.callRpcHandler.Handle(result, method, args...)
}

func (c *ClientCore) callRpc(result interface{}, method string, args ...interface{}) error {
	ctx, cancelFunc := c.genContext()
	if cancelFunc != nil {
		defer cancelFunc()
	}
	return c.rpcProvider.CallContext(ctx, result, method, args...)
}

// UseCallRpcMiddleware set middleware to hook CallRpc, for example use middleware.CallRpcLogMiddleware for logging request info.
// You can customize your CallRpcMiddleware and use multi CallRpcMiddleware.
func (c *ClientCore) UseCallRpcMiddleware(middleware middleware.CallRpcMiddleware) {
	c.callRpcHandler = middleware(c.callRpcHandler)
}

// BatchCallRPC sends all given requests as a single batch and waits for the server
// to return a response for all of them.
//
// In contrast to Call, BatchCall only returns I/O errors. Any error specific to
// a request is reported through the Error field of the corresponding BatchElem.
//
// Note that batch calls may not be executed atomically on the server side.
//
// You could use UseBatchCallRpcMiddleware to add middleware for hooking BatchCallRPC
func (c *ClientCore) BatchCallRPC(b []rpc.BatchElem) error {
	return c.batchCallRpcHandler.Handle(b)
}

func (c *ClientCore) batchCallRPC(b []rpc.BatchElem) error {
	ctx, cancelFunc := c.genContext()
	if cancelFunc != nil {
		defer cancelFunc()
	}

	return c.rpcProvider.BatchCallContext(ctx, b)
}

// UseBatchCallRpcMiddleware set middleware to hook BatchCallRpc, for example use middleware.BatchCallRpcLogMiddleware for logging batch request info.
// You can customize your BatchCallRpcMiddleware and use multi BatchCallRpcMiddleware.
func (c *ClientCore) UseBatchCallRpcMiddleware(middleware middleware.BatchCallRpcMiddleware) {
	c.batchCallRpcHandler = middleware(c.batchCallRpcHandler)
}

// Close closes the client, aborting any in-flight requests.
func (c *ClientCore) Close() {
	c.rpcProvider.Close()
}

func (c *ClientCore) wrappedCallRPC(result interface{}, method string, args ...interface{}) error {
	fmtedArgs := c.genRPCParams(args...)
	return c.CallRPC(result, method, fmtedArgs...)
}

func (c *ClientCore) genRPCParams(args ...interface{}) []interface{} {
	// fmt.Println("gen rpc params")
	params := []interface{}{}
	for i := range args {
		// fmt.Printf("args %v:%v\n", i, args[i])
		if !utils.IsNil(args[i]) {
			// fmt.Printf("args %v:%v is not nil\n", i, args[i])

			if tmp, ok := args[i].(cfxaddress.Address); ok {
				tmp.CompleteByNetworkID(c.networkID)
				args[i] = tmp
				// fmt.Printf("complete by networkID,%v; after %v\n", client.networkID, args[i])
			}

			if tmp, ok := args[i].(*cfxaddress.Address); ok {
				tmp.CompleteByNetworkID(c.networkID)
				// fmt.Printf("complete by networkID,%v; after %v\n", client.networkID, args[i])
			}

			if tmp, ok := args[i].(types.CallRequest); ok {
				tmp.From.CompleteByNetworkID(c.networkID)
				tmp.To.CompleteByNetworkID(c.networkID)
				args[i] = tmp
			}

			if tmp, ok := args[i].(*types.CallRequest); ok {
				tmp.From.CompleteByNetworkID(c.networkID)
				tmp.To.CompleteByNetworkID(c.networkID)
			}

			params = append(params, args[i])
		}
	}
	return params
}

func (c *ClientCore) genContext() (context.Context, context.CancelFunc) {
	if c.requestTimeout > 0 {
		return context.WithTimeout(context.Background(), c.requestTimeout)
	}
	return context.Background(), nil
}

func get1stEpochIfy(epoch []*types.Epoch) *types.Epoch {
	var realEpoch *types.Epoch
	if len(epoch) > 0 {
		realEpoch = epoch[0]
	}
	return realEpoch
}
