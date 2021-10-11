// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package cfxclient

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
	nodeURL             string
	rpcProvider         interfaces.RpcProvider
	networkID           uint32
	option              ClientOption
	callRpcHandler      middleware.CallRpcHandler
	batchCallRpcHandler middleware.BatchCallRpcHandler
}

// ClientOption for set keystore path and flags for retry
//
// The simplest way to set logger is to use the types.DefaultCallRpcLog and types.DefaultBatchCallRPCLog
type ClientOption struct {
	// KeystorePath string
	// retry
	RetryCount    int
	RetryInterval time.Duration
	// timeout of request
	RequestTimeout time.Duration
}

// NewClient creates an instance of Client with specified conflux node url, it will creat account manager if option.KeystorePath not empty.
func newClientCore(nodeURL string) (ClientCore, error) {
	realOption := ClientOption{}
	// if len(option) > 0 {
	// 	realOption = option[0]
	// }

	client, err := newClientCoreWithRetry(nodeURL, realOption)
	if err != nil {
		return ClientCore{}, errors.Wrap(err, "failed to new client with retry")
	}

	return client, nil
}

// NewClientWithRetry creates a retryable new instance of Client with specified conflux node url and retry options.
//
// the clientOption.RetryInterval will be set to 1 second if pass 0
func newClientCoreWithRetry(nodeURL string, clientOption ClientOption) (ClientCore, error) {

	var client ClientCore
	client.nodeURL = nodeURL
	client.option = clientOption
	client.callRpcHandler = middleware.CallRpcHandlerFunc(client.callRpc)
	client.batchCallRpcHandler = middleware.BatchCallRpcHandlerFunc(client.batchCallRPC)
	client.option.setDefault()

	rpcClient, err := rpc.Dial(nodeURL)
	if err != nil {
		return ClientCore{}, errors.Wrap(err, "failed to dial to fullnode")
	}

	if client.option.RetryCount == 0 {
		client.rpcProvider = rpcClient
	} else {
		client.rpcProvider = &rpcClientWithRetry{
			inner:      rpcClient,
			retryCount: client.option.RetryCount,
			interval:   client.option.RetryInterval,
		}
	}

	// set networkId
	if client.networkID, err = client.getNetworkId(); err != nil {
		return ClientCore{}, errors.Wrap(err, "failed to get networkID")
	}

	return client, nil
}

func (co *ClientOption) setDefault() {
	if co.RequestTimeout == 0 {
		co.RequestTimeout = time.Second * 30
	}
	// Interval 0 is meaningless and may lead full node busy, so default sets it to 1 second
	if co.RetryInterval == 0 {
		co.RetryInterval = time.Second
	}
}

func (client *ClientCore) SetRetry(retryCount int, retryInterval time.Duration) *ClientCore {
	client.option.RetryCount = retryCount
	client.option.RetryInterval = retryInterval
	return client
}

func (client *ClientCore) SetRequestTimeout(timeout time.Duration) *ClientCore {
	client.option.RequestTimeout = timeout
	return client
}

// GetNodeURL returns node url
func (client *ClientCore) GetNodeURL() string {
	return client.nodeURL
}

func (client *ClientCore) getNetworkId() (uint32, error) {
	if client.networkID != 0 {
		return client.networkID, nil
	}

	status, err := client.getStatus()
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
func (client *ClientCore) CallRPC(result interface{}, method string, args ...interface{}) error {
	return client.callRpcHandler.Handle(result, method, args...)
}

func (client *ClientCore) callRpc(result interface{}, method string, args ...interface{}) error {
	ctx, cancelFunc := client.genContext()
	if cancelFunc != nil {
		defer cancelFunc()
	}
	return client.rpcProvider.CallContext(ctx, result, method, args...)
}

// UseCallRpcMiddleware set middleware to hook CallRpc, for example use middleware.CallRpcLogMiddleware for logging request info.
// You can customize your CallRpcMiddleware and use multi CallRpcMiddleware.
func (client *ClientCore) UseCallRpcMiddleware(middleware middleware.CallRpcMiddleware) {
	client.callRpcHandler = middleware(client.callRpcHandler)
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
func (client *ClientCore) BatchCallRPC(b []rpc.BatchElem) error {
	return client.batchCallRpcHandler.Handle(b)
}

func (client *ClientCore) batchCallRPC(b []rpc.BatchElem) error {
	ctx, cancelFunc := client.genContext()
	if cancelFunc != nil {
		defer cancelFunc()
	}

	return client.rpcProvider.BatchCallContext(ctx, b)
}

// UseBatchCallRpcMiddleware set middleware to hook BatchCallRpc, for example use middleware.BatchCallRpcLogMiddleware for logging batch request info.
// You can customize your BatchCallRpcMiddleware and use multi BatchCallRpcMiddleware.
func (client *ClientCore) UseBatchCallRpcMiddleware(middleware middleware.BatchCallRpcMiddleware) {
	client.batchCallRpcHandler = middleware(client.batchCallRpcHandler)
}

// Close closes the client, aborting any in-flight requests.
func (client *ClientCore) Close() {
	client.rpcProvider.Close()
}

func (client *ClientCore) wrappedCallRPC(result interface{}, method string, args ...interface{}) error {
	fmtedArgs := client.genRPCParams(args...)
	return client.CallRPC(result, method, fmtedArgs...)
}

func (client *ClientCore) genRPCParams(args ...interface{}) []interface{} {
	// fmt.Println("gen rpc params")
	params := []interface{}{}
	for i := range args {
		// fmt.Printf("args %v:%v\n", i, args[i])
		if !utils.IsNil(args[i]) {
			// fmt.Printf("args %v:%v is not nil\n", i, args[i])

			if tmp, ok := args[i].(cfxaddress.Address); ok {
				tmp.CompleteByNetworkID(client.networkID)
				args[i] = tmp
				// fmt.Printf("complete by networkID,%v; after %v\n", client.networkID, args[i])
			}

			if tmp, ok := args[i].(*cfxaddress.Address); ok {
				tmp.CompleteByNetworkID(client.networkID)
				// fmt.Printf("complete by networkID,%v; after %v\n", client.networkID, args[i])
			}

			if tmp, ok := args[i].(types.CallRequest); ok {
				tmp.From.CompleteByNetworkID(client.networkID)
				tmp.To.CompleteByNetworkID(client.networkID)
				args[i] = tmp
			}

			if tmp, ok := args[i].(*types.CallRequest); ok {
				tmp.From.CompleteByNetworkID(client.networkID)
				tmp.To.CompleteByNetworkID(client.networkID)
			}

			params = append(params, args[i])
		}
	}
	return params
}

func get1stEpochIfy(epoch []*types.Epoch) *types.Epoch {
	var realEpoch *types.Epoch
	if len(epoch) > 0 {
		realEpoch = epoch[0]
	}
	return realEpoch
}

func (client *ClientCore) genContext() (context.Context, context.CancelFunc) {
	if client.option.RequestTimeout > 0 {
		return context.WithTimeout(context.Background(), client.option.RequestTimeout)
	}
	return context.Background(), nil
}
