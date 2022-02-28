package sdk

import (
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	postypes "github.com/Conflux-Chain/go-conflux-sdk/types/pos"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// RpcPosClient used to access pos namespace RPC of Conflux blockchain.
type RpcPosClient struct {
	core *Client
}

// NewRpcPosClient creates a new RpcPosClient instance.
func NewRpcPosClient(core *Client) RpcPosClient {
	return RpcPosClient{core}
}

// GetStatus returns pos chain status
func (c *RpcPosClient) GetStatus() (status postypes.Status, err error) {
	err = c.core.CallRPC(&status, "pos_getStatus")
	return
}

// GetAccount returns account info at block
func (c *RpcPosClient) GetAccount(address postypes.Address, blockNumber ...hexutil.Uint64) (account postypes.Account, err error) {
	_view := get1stU64Ify(blockNumber)
	err = c.core.CallRPC(&account, "pos_getAccount", address, _view)
	return
}

// GetCommittee returns committee info at block
func (c *RpcPosClient) GetCommittee(blockNumber ...hexutil.Uint64) (committee postypes.CommitteeState, err error) {
	_view := get1stU64Ify(blockNumber)
	err = c.core.CallRPC(&committee, "pos_getCommittee", _view)
	return
}

// GetBlockByHash returns block info of block hash
func (c *RpcPosClient) GetBlockByHash(hash types.Hash) (block *postypes.Block, err error) {
	err = c.core.CallRPC(&block, "pos_getBlockByHash", hash)
	return
}

// GetBlockByHash returns block at block number
func (c *RpcPosClient) GetBlockByNumber(blockNumber postypes.BlockNumber) (block *postypes.Block, err error) {
	err = c.core.CallRPC(&block, "pos_getBlockByNumber", blockNumber)
	return
}

// GetTransactionByNumber returns transaction info of transaction number
func (c *RpcPosClient) GetTransactionByNumber(txNumber hexutil.Uint64) (transaction *postypes.Transaction, err error) {
	err = c.core.CallRPC(&transaction, "pos_getTransactionByNumber", txNumber)
	return
}

// GetRewardsByEpoch returns rewards of epoch
func (c *RpcPosClient) GetRewardsByEpoch(epochNumber hexutil.Uint64) (reward postypes.EpochReward, err error) {
	err = c.core.CallRPC(&reward, "pos_getRewardsByEpoch", epochNumber)
	return
}
