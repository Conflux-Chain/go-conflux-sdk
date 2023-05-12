package sdk

import "github.com/Conflux-Chain/go-conflux-sdk/types"

type RpcTraceClient struct {
	core *Client
}

// NewRpcPosClient creates a new RpcPosClient instance.
func NewRpcTraceClient(core *Client) RpcPosClient {
	return RpcPosClient{core}
}

// GetBlockTrace returns all traces produced at given block.
func (c *RpcTraceClient) GetBlockTraces(blockHash types.Hash) (traces *types.LocalizedBlockTrace, err error) {
	err = c.core.wrappedCallRPC(&traces, "trace_block", blockHash)
	return
}

// GetFilterTraces returns all traces matching the provided filter.
func (c *RpcTraceClient) FilterTraces(traceFilter types.TraceFilter) (traces []types.LocalizedTrace, err error) {
	err = c.core.wrappedCallRPC(&traces, "trace_filter", traceFilter)
	return
}

// GetTransactionTraces returns all traces produced at the given transaction.
func (c *RpcTraceClient) GetTransactionTraces(txHash types.Hash) (traces []types.LocalizedTrace, err error) {
	err = c.core.wrappedCallRPC(&traces, "trace_transaction", txHash)
	return
}

func (c *RpcTraceClient) GetEpochTraces(epoch types.Epoch) (traces []types.LocalizedTrace, err error) {
	err = c.core.wrappedCallRPC(&traces, "trace_epoch", epoch)
	return
}

// /// Return all traces of both spaces in an epoch.
// #[rpc(name = "trace_epoch")]
// fn epoch_traces(&self, epoch: EpochNumber) -> JsonRpcResult<EpochTrace>;
