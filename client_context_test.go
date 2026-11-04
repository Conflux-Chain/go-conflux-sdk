package sdk

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	providers "github.com/openweb3/go-rpc-provider/provider_wrapper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newLocalRPCClient(t *testing.T) *Client {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			ID     json.RawMessage `json:"id"`
			Method string          `json:"method"`
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			t.Errorf("decode local RPC request: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		assert.Equal(t, "cfx_getStatus", request.Method)
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      request.ID,
			"result": types.Status{
				BestHash:  "0x0000000000000000000000000000000000000000000000000000000000000000",
				NetworkID: 1,
				ChainID:   1,
			},
		})
		assert.NoError(t, err)
	}))
	t.Cleanup(server.Close)
	client, err := NewClient(server.URL)
	require.NoError(t, err)
	t.Cleanup(client.Close)
	return client
}

func namespaceRPCCalls(client *Client) []struct {
	name   string
	method string
	call   func() error
} {
	trace := NewRpcTraceClient(client)
	filter := NewRpcFilterClient(client)
	return []struct {
		name   string
		method string
		call   func() error
	}{
		{"cfx", "cfx_getStatus", func() error { _, err := client.GetStatus(); return err }},
		{"pos", "pos_getStatus", func() error { _, err := client.Pos().GetStatus(); return err }},
		{"txpool", "txpool_status", func() error { _, err := client.TxPool().Status(); return err }},
		{"debug", "cfx_getEpochReceipts", func() error {
			_, err := client.Debug().GetEpochReceipts(*types.NewEpochOrBlockHashWithEpoch(types.EpochLatestState))
			return err
		}},
		{"filter", "cfx_newBlockFilter", func() error { _, err := client.Filter().NewBlockFilter(); return err }},
		{"trace", "trace_epoch", func() error { _, err := client.Trace().GetEpochTraces(*types.EpochLatestState); return err }},
		{"promoted_trace", "trace_epoch", func() error { _, err := client.GetEpochTraces(*types.EpochLatestState); return err }},
		{"trace_constructor", "trace_epoch", func() error { _, err := trace.GetEpochTraces(*types.EpochLatestState); return err }},
		{"filter_constructor", "cfx_newBlockFilter", func() error { _, err := filter.NewBlockFilter(); return err }},
	}
}

func TestClientWithContextNamespaces(t *testing.T) {
	constructors := []struct {
		name string
		new  func(*testing.T) *Client
	}{
		{"url", newLocalRPCClient},
		{"provider", func(t *testing.T) *Client {
			client, err := NewClientWithProvider(newLocalRPCClient(t).Provider())
			require.NoError(t, err)
			return client
		}},
		{"rpc_requester", func(t *testing.T) *Client {
			client, err := NewClientWithRPCRequester(newLocalRPCClient(t).Provider())
			require.NoError(t, err)
			return client
		}},
	}
	for _, constructor := range constructors {
		t.Run(constructor.name, func(t *testing.T) {
			client := constructor.new(t)
			// Avoid a separate network ID lookup when formatting canceled trace calls.
			client.SetNetworkId(1)
			type contextKey struct{}
			ctx, cancel := context.WithDeadline(context.WithValue(context.Background(), contextKey{}, "derived"), time.Now().Add(time.Hour))
			defer cancel()
			derived := client.WithContext(ctx)
			siblingCtx, cancelSibling := context.WithCancel(context.WithValue(context.Background(), contextKey{}, "sibling"))
			sibling := client.WithContext(siblingCtx)
			cancelSibling()
			chainedCtx := context.WithValue(context.Background(), contextKey{}, "chained")
			expiredCtx, cancelExpired := context.WithDeadline(context.WithValue(context.Background(), contextKey{}, "expired"), time.Unix(1, 0))
			defer cancelExpired()

			var receivedValue interface{}
			var receivedDeadline time.Time
			var receivedHasDeadline bool
			var receivedErr error
			var receivedMethod string
			client.Provider().HookCallContext(func(next providers.CallContextFunc) providers.CallContextFunc {
				return func(ctx context.Context, result interface{}, method string, args ...interface{}) error {
					receivedValue = ctx.Value(contextKey{})
					receivedDeadline, receivedHasDeadline = ctx.Deadline()
					receivedErr = ctx.Err()
					receivedMethod = method
					return ctx.Err()
				}
			})
			variants := []struct {
				name   string
				client *Client
				ctx    context.Context
			}{
				{"original", client, context.Background()},
				{"derived", derived, ctx},
				{"canceled_sibling", sibling, siblingCtx},
				{"chained", derived.WithContext(chainedCtx), chainedCtx},
				{"nil_context", derived.WithContext(nil), context.Background()},
				{"expired", derived.WithContext(expiredCtx), expiredCtx},
				{"derived_again", derived, ctx},
				{"original_again", client, context.Background()},
			}
			for _, variant := range variants {
				t.Run(variant.name, func(t *testing.T) {
					for _, call := range namespaceRPCCalls(variant.client) {
						t.Run(call.name, func(t *testing.T) {
							receivedMethod = ""
							err := call.call()
							if variant.ctx.Err() == nil {
								require.NoError(t, err)
							} else {
								require.ErrorIs(t, err, variant.ctx.Err())
							}
							require.Equal(t, call.method, receivedMethod)
							assert.Equal(t, variant.ctx.Value(contextKey{}), receivedValue, "%s lost the caller's context value", call.method)
							assert.Equal(t, variant.ctx.Err(), receivedErr, "%s lost the caller's cancellation", call.method)
							if deadline, ok := variant.ctx.Deadline(); ok {
								require.True(t, receivedHasDeadline, "%s lost the caller's deadline", call.method)
								assert.Equal(t, deadline, receivedDeadline)
							}
						})
					}
				})
			}
		})
	}
}
