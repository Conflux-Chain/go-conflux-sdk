// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package sdk

import (
	"encoding/hex"
	"encoding/json"
	"math/big"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

// Client represents a client to interact with Conflux blockchain.
type Client struct {
	rpcClient *rpc.Client
}

// NewClient creates a new instance of Client with specified url.
func NewClient(url string) (*Client, error) {
	client, err := rpc.Dial(url)
	if err != nil {
		return nil, err
	}

	return &Client{
		rpcClient: client,
	}, nil
}

// GasPrice returns the recent mean gas price.
func (c *Client) GasPrice() (*big.Int, error) {
	var result interface{}

	if err := c.rpcClient.Call(&result, "cfx_gasPrice"); err != nil {
		return nil, err
	}

	return hexutil.DecodeBig(result.(string))
}

// GetEpochNumber returns the highest or specified epoch number.
func (c *Client) GetEpochNumber(epoch ...*types.Epoch) (*big.Int, error) {
	var result interface{}

	var args []interface{}
	if len(epoch) > 0 {
		args = append(args, epoch[0])
	}

	if err := c.rpcClient.Call(&result, "cfx_epochNumber", args...); err != nil {
		return nil, err
	}

	return hexutil.DecodeBig(result.(string))
}

// GetBalance returns the balance of specified account.
func (c *Client) GetBalance(address types.Address, epoch ...*types.Epoch) (*big.Int, error) {
	var result interface{}

	args := []interface{}{address}
	if len(epoch) > 0 {
		args = append(args, epoch[0])
	}

	if err := c.rpcClient.Call(&result, "cfx_getBalance", args...); err != nil {
		return nil, err
	}

	return hexutil.DecodeBig(result.(string))
}

// GetCode returns the bytecodes in HEX format of specified contract.
func (c *Client) GetCode(address types.Address, epoch ...*types.Epoch) (string, error) {
	var result interface{}

	args := []interface{}{address}
	if len(epoch) > 0 {
		args = append(args, epoch[0])
	}

	if err := c.rpcClient.Call(&result, "cfx_getCode", args...); err != nil {
		return "", err
	}

	return result.(string), nil
}

// GetBlockSummaryByHash returns the block summary of specified block hash.
// If block not found, return nil.
func (c *Client) GetBlockSummaryByHash(blockHash types.Hash) (*types.BlockSummary, error) {
	var result interface{}

	if err := c.rpcClient.Call(&result, "cfx_getBlockByHash", blockHash, false); err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	var block types.BlockSummary
	if err := unmarshalRPCResult(result, &block); err != nil {
		return nil, err
	}

	return &block, nil
}

// GetBlockByHash returns the block of specified block hash.
// If block not found, return nil.
func (c *Client) GetBlockByHash(blockHash types.Hash) (*types.Block, error) {
	var result interface{}

	if err := c.rpcClient.Call(&result, "cfx_getBlockByHash", blockHash, true); err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	var block types.Block
	if err := unmarshalRPCResult(result, &block); err != nil {
		return nil, err
	}

	return &block, nil
}

func unmarshalRPCResult(result interface{}, v interface{}) error {
	encoded, err := json.Marshal(result)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(encoded, v); err != nil {
		return err
	}

	return nil
}

// GetBlockSummaryByEpoch returns the block summary of specified epoch.
// If the epoch is invalid, return the concrete error.
func (c *Client) GetBlockSummaryByEpoch(epoch *types.Epoch) (*types.BlockSummary, error) {
	var result interface{}

	if err := c.rpcClient.Call(&result, "cfx_getBlockByEpochNumber", epoch, false); err != nil {
		return nil, err
	}

	var block types.BlockSummary
	if err := unmarshalRPCResult(result, &block); err != nil {
		return nil, err
	}

	return &block, nil
}

// GetBlockByEpoch returns the block of specified epoch.
// If the epoch is invalid, return the concrete error.
func (c *Client) GetBlockByEpoch(epoch *types.Epoch) (*types.Block, error) {
	var result interface{}

	if err := c.rpcClient.Call(&result, "cfx_getBlockByEpochNumber", epoch, true); err != nil {
		return nil, err
	}

	var block types.Block
	if err := unmarshalRPCResult(result, &block); err != nil {
		return nil, err
	}

	return &block, nil
}

// GetBestBlockHash returns the current best block hash.
func (c *Client) GetBestBlockHash() (types.Hash, error) {
	var result interface{}

	if err := c.rpcClient.Call(&result, "cfx_getBestBlockHash"); err != nil {
		return "", err
	}

	return types.Hash(result.(string)), nil
}

// GetTransactionCount returns the number of transactions sent from given address.
// If epoch specified, returns the number of transactions in the state of specified epoch.
// Otherwise, returns the number of transactions in latest state.
func (c *Client) GetTransactionCount(address types.Address, epoch ...*types.Epoch) (*big.Int, error) {
	var result interface{}

	args := []interface{}{address}
	if len(epoch) > 0 {
		args = append(args, epoch[0])
	}

	if err := c.rpcClient.Call(&result, "cfx_getTransactionCount", args...); err != nil {
		return nil, err
	}

	return hexutil.DecodeBig(result.(string))
}

// SendSignedTransaction sends signed transaction and return its hash.
func (c *Client) SendSignedTransaction(rawData []byte) (types.Hash, error) {
	var result interface{}

	if err := c.rpcClient.Call(&result, "cfx_sendRawTransaction", hexutil.Encode(rawData)); err != nil {
		return "", err
	}

	return types.Hash(result.(string)), nil
}

// SignEncodedTransactionAndSend sign RLP encoded transaction and send it, return responsed transaction
func (c *Client) SignEncodedTransactionAndSend(encodedTx []byte, v byte, r, s []byte) (*types.Transaction, error) {
	tx, err := types.DecodeRlpToUnsignTransction(encodedTx)
	if err != nil {
		return nil, err
	}

	respondTx, err := c.signTransactionAndSend(tx, v, r, s)
	if err != nil {
		return nil, err
	}

	return respondTx, nil
}

func (c *Client) signTransactionAndSend(tx *types.UnsignedTransaction, v byte, r, s []byte) (*types.Transaction, error) {
	rlp, err := tx.EncodeWithSignature(v, r, s)
	if err != nil {
		return nil, err
	}
	hash, err := c.SendSignedTransaction(rlp)
	if err != nil {
		return nil, err
	}
	respondTx, err := c.GetTransactionByHash(hash)
	if err != nil {
		return nil, err
	}
	return respondTx, nil
}

// Call executes contract but not mined into the blockchain,
// and returns the contract execution result.
func (c *Client) Call(request types.CallRequest, epoch ...*types.Epoch) (string, error) {
	var result interface{}

	args := []interface{}{request}
	if len(epoch) > 0 {
		args = append(args, epoch[0])
	}

	if err := c.rpcClient.Call(&result, "cfx_call", args...); err != nil {
		return "", err
	}

	return result.(string), nil
}

// GetLogs returns logs that matching the specified filter.
func (c *Client) GetLogs(filter types.LogFilter) ([]types.Log, error) {
	var result interface{}

	if err := c.rpcClient.Call(&result, "cfx_getLogs", filter); err != nil {
		return nil, err
	}

	var log []types.Log
	if err := unmarshalRPCResult(result, &log); err != nil {
		return nil, err
	}

	return log, nil
}

// GetTransactionByHash returns transaction for the specified hash.
// If transaction not found, return nil.
func (c *Client) GetTransactionByHash(txHash types.Hash) (*types.Transaction, error) {
	var result interface{}

	if err := c.rpcClient.Call(&result, "cfx_getTransactionByHash", txHash); err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	var tx types.Transaction
	if err := unmarshalRPCResult(result, &tx); err != nil {
		return nil, err
	}

	return &tx, nil
}

// EstimateGas estimates the consumed gas of transaction/contract execution.
func (c *Client) EstimateGas(request types.CallRequest, epoch ...*types.Epoch) (*big.Int, error) {
	var result interface{}

	args := []interface{}{request}
	if len(epoch) > 0 {
		args = append(args, epoch[0])
	}

	if err := c.rpcClient.Call(&result, "cfx_estimateGas", args...); err != nil {
		return nil, err
	}

	return hexutil.DecodeBig(result.(string))
}

// EstimateGasAndCollateral estimates the consumed gas and storage for collateral of transaction/contract execution.
func (c *Client) EstimateGasAndCollateral(request types.CallRequest) (*types.Estimate, error) {
	var result interface{}

	args := []interface{}{request}

	if err := c.rpcClient.Call(&result, "cfx_estimateGasAndCollateral", args...); err != nil {
		return nil, err
	}
	var estimate types.Estimate
	if err := unmarshalRPCResult(result, &estimate); err != nil {
		return nil, err
	}

	return &estimate, nil
}

// GetBlocksByEpoch returns the blocks in the specified epoch.
func (c *Client) GetBlocksByEpoch(epoch *types.Epoch) ([]types.Hash, error) {
	var result interface{}

	if err := c.rpcClient.Call(&result, "cfx_getBlocksByEpoch", epoch); err != nil {
		return nil, err
	}

	var blocks []types.Hash
	if err := unmarshalRPCResult(result, &blocks); err != nil {
		return nil, err
	}

	return blocks, nil
}

// GetTransactionReceipt returns the receipt of specified transaction hash.
// If receipt not found, return nil.
func (c *Client) GetTransactionReceipt(txHash types.Hash) (*types.Receipt, error) {
	var result interface{}

	if err := c.rpcClient.Call(&result, "cfx_getTransactionReceipt", txHash); err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	var receipt types.Receipt
	if err := unmarshalRPCResult(result, &receipt); err != nil {
		return nil, err
	}

	return &receipt, nil
}

// CreateUnsignedTransaction create an UnsignedTransaction instance
func (c *Client) CreateUnsignedTransaction(from types.Address, to types.Address, amount hexutil.Big) (*types.UnsignedTransaction, error) {
	tx := new(types.UnsignedTransaction)
	tx.From = from
	tx.To = &to
	tx.Value = &amount
	err := c.applyUnsignedTransactionDefault(tx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *Client) applyUnsignedTransactionDefault(tx *types.UnsignedTransaction) error {
	tx.ApplyDefault()
	if c != nil {
		if tx.EpochHeight == nil {
			epoch, err := c.GetEpochNumber(types.EpochLatestState)
			if err != nil {
				return err
			}
			tx.EpochHeight = (*hexutil.Big)(epoch)
		}

		if tx.StorageLimit == nil {
			callReq := new(types.CallRequest)
			callReq.To = tx.To
			dataStr := "0x" + hex.EncodeToString(tx.Data)
			callReq.Data = &dataStr
			sm, err := c.EstimateGasAndCollateral(*callReq)
			if err != nil {
				return err
			}
			tx.StorageLimit = sm.StorageCollateralized
		}
	}

	return nil
}

// Debug calls the Conflux debug API.
func (c *Client) Debug(method string, args ...interface{}) (interface{}, error) {
	var result interface{}

	if err := c.rpcClient.Call(&result, method, args...); err != nil {
		return nil, err
	}

	return result, nil
}

// Close closes the client, aborting any in-flight requests.
func (c *Client) Close() {
	c.rpcClient.Close()
}
