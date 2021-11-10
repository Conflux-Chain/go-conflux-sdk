package client

import (
	"github.com/Conflux-Chain/go-conflux-sdk/types"
)

type RpcTraceClient struct {
	core *ClientCore
}

func NewRpcTraceClient(core *ClientCore) RpcTraceClient {
	return RpcTraceClient{core}
}

// GetBlockTrace returns all traces produced at given block.
func (client *RpcTraceClient) GetBlockTraces(blockHash types.Hash) (traces *types.LocalizedBlockTrace, err error) {
	err = client.core.wrappedCallRPC(&traces, "trace_block", blockHash)
	return
}

// GetFilterTraces returns all traces matching the provided filter.
func (client *RpcTraceClient) FilterTraces(traceFilter types.TraceFilter) (traces []types.LocalizedTrace, err error) {
	err = client.core.wrappedCallRPC(&traces, "trace_filter", traceFilter)
	return
}

// GetTransactionTraces returns all traces produced at the given transaction.
func (client *RpcTraceClient) GetTransactionTraces(txHash types.Hash) (traces []types.LocalizedTrace, err error) {
	err = client.core.wrappedCallRPC(&traces, "trace_transaction", txHash)
	return
}
