// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package sdk

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/constants"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

// Client represents a client to interact with Conflux blockchain.
type Client struct {
	rpcClient      *rpc.Client
	accountManager *AccountManager
}

// NewClient creates a new instance of Client with specified conflux node url.
func NewClient(nodeURL string) (*Client, error) {
	client, err := rpc.Dial(nodeURL)
	if err != nil {
		return nil, types.WrapError(err, "dail failed")
	}

	return &Client{
		rpcClient: client,
	}, nil
}

// SetAccountManager set account manager for sign transaction
func (c *Client) SetAccountManager(accountManager *AccountManager) {
	c.accountManager = accountManager
}

// GetGasPrice returns the recent mean gas price.
func (c *Client) GetGasPrice() (*big.Int, error) {
	var result interface{}

	if err := c.rpcClient.Call(&result, "cfx_gasPrice"); err != nil {
		msg := "rpc request cfx_gasPrice error"
		return nil, types.WrapError(err, msg)
	}

	return hexutil.DecodeBig(result.(string))
}

// GetNextNonce returns the address next transaction nonce.
func (c *Client) GetNextNonce(address types.Address) (uint64, error) {
	var result interface{}
	if err := c.rpcClient.Call(&result, "cfx_getNextNonce", address); err != nil {
		msg := fmt.Sprintf("rpc request cfx_getNextNonce %+v error", address)
		return 0, types.WrapErrorf(err, msg)
	}

	// remove prefix "0x"
	result = string([]byte(result.(string))[2:])
	nonce, err := strconv.ParseUint(result.(string), 16, 64)
	if err != nil {
		msg := fmt.Sprintf("parse uint %+v error", result)
		return 0, types.WrapError(err, msg)
	}

	return nonce, nil
}

// GetEpochNumber returns the highest or specified epoch number.
func (c *Client) GetEpochNumber(epoch ...*types.Epoch) (*big.Int, error) {
	var result interface{}

	var args []interface{}
	if len(epoch) > 0 {
		args = append(args, epoch[0])
	}

	if err := c.rpcClient.Call(&result, "cfx_epochNumber", args...); err != nil {
		msg := fmt.Sprintf("rpc cfx_epochNumber %+v error", args)
		return nil, types.WrapError(err, msg)
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
		msg := fmt.Sprintf("rpc cfx_getBalance %+v error", args)
		return nil, types.WrapError(err, msg)
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
		msg := fmt.Sprintf("rpc cfx_getCode %+v error", args)
		return "", types.WrapError(err, msg)
	}

	return result.(string), nil
}

// GetBlockSummaryByHash returns the block summary of specified block hash.
// If block not found, return nil.
func (c *Client) GetBlockSummaryByHash(blockHash types.Hash) (*types.BlockSummary, error) {
	var result interface{}

	if err := c.rpcClient.Call(&result, "cfx_getBlockByHash", blockHash, false); err != nil {
		msg := fmt.Sprintf("rpc cfx_getBlockByHash %+v error", blockHash)
		return nil, types.WrapError(err, msg)
	}

	if result == nil {
		return nil, nil
	}

	var block types.BlockSummary
	if err := utils.UnmarshalRPCResult(result, &block); err != nil {
		msg := fmt.Sprintf("UnmarshalRPCResult %+v error", result)
		return nil, types.WrapError(err, msg)
	}

	return &block, nil
}

// GetBlockByHash returns the block of specified block hash.
// If block not found, return nil.
func (c *Client) GetBlockByHash(blockHash types.Hash) (*types.Block, error) {
	var result interface{}

	if err := c.rpcClient.Call(&result, "cfx_getBlockByHash", blockHash, true); err != nil {
		msg := fmt.Sprintf("rpc cfx_getBlockByHash %+v error", blockHash)
		return nil, types.WrapError(err, msg)
	}

	if result == nil {
		return nil, nil
	}

	var block types.Block
	if err := utils.UnmarshalRPCResult(result, &block); err != nil {
		msg := fmt.Sprintf("UnmarshalRPCResult %+v error", result)
		return nil, types.WrapError(err, msg)
	}

	return &block, nil
}

// GetBlockSummaryByEpoch returns the block summary of specified epoch.
// If the epoch is invalid, return the concrete error.
func (c *Client) GetBlockSummaryByEpoch(epoch *types.Epoch) (*types.BlockSummary, error) {
	var result interface{}

	if err := c.rpcClient.Call(&result, "cfx_getBlockByEpochNumber", epoch, false); err != nil {
		msg := fmt.Sprintf("rpc cfx_getBlockByEpochNumber %+v error", epoch)
		return nil, types.WrapError(err, msg)
	}

	var block types.BlockSummary
	if err := utils.UnmarshalRPCResult(result, &block); err != nil {
		msg := fmt.Sprintf("UnmarshalRPCResult %+v error", result)
		return nil, types.WrapError(err, msg)
	}

	return &block, nil
}

// GetBlockByEpoch returns the block of specified epoch.
// If the epoch is invalid, return the concrete error.
func (c *Client) GetBlockByEpoch(epoch *types.Epoch) (*types.Block, error) {
	var result interface{}

	if err := c.rpcClient.Call(&result, "cfx_getBlockByEpochNumber", epoch, true); err != nil {
		msg := fmt.Sprintf("rpc cfx_getBlockByEpochNumber %+v error", epoch)
		return nil, types.WrapError(err, msg)
	}

	var block types.Block
	if err := utils.UnmarshalRPCResult(result, &block); err != nil {
		msg := fmt.Sprintf("UnmarshalRPCResult %+v error", result)
		return nil, types.WrapError(err, msg)
	}

	return &block, nil
}

// GetBestBlockHash returns the current best block hash.
func (c *Client) GetBestBlockHash() (types.Hash, error) {
	var result interface{}

	if err := c.rpcClient.Call(&result, "cfx_getBestBlockHash"); err != nil {
		msg := "rpc cfx_getBestBlockHash error"
		return "", types.WrapError(err, msg)
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
		msg := fmt.Sprintf("rpc cfx_getTransactionCount %+v error", args)
		return nil, types.WrapError(err, msg)
	}

	return hexutil.DecodeBig(result.(string))
}

// SendTransaction sign and send transaction to conflux blockchain and return the transaction hash.
func (c *Client) SendTransaction(tx *types.UnsignedTransaction) (types.Hash, error) {

	err := c.ApplyUnsignedTransactionDefault(tx)
	if err != nil {
		msg := fmt.Sprintf("apply transaction {%+v} default fields error", *tx)
		return "", types.WrapError(err, msg)
	}

	//check balance, return error if balance not enough
	epoch := types.NewEpochNumber(tx.EpochHeight.ToInt())
	balance, err := c.GetBalance(*tx.From, epoch)
	if err != nil {
		msg := fmt.Sprintf("get balance of %+v at ephoc %+v error", tx.From, epoch)
		return "", types.WrapError(err, msg)
	}
	need := big.NewInt(int64(tx.Gas))
	need = need.Add(tx.StorageLimit.ToInt(), need)
	need = need.Mul(tx.GasPrice.ToInt(), need)
	need = need.Add(tx.Value.ToInt(), need)
	need = need.Add(tx.StorageLimit.ToInt(), need)

	if balance.Cmp(need) < 0 {
		msg := fmt.Sprintf("out of balance, need %+v but your balance is %+v", need, balance)
		return "", types.WrapError(err, msg)
	}

	//sign
	// fmt.Printf("ready to send transaction %+v\n\n", tx)
	rawData, err := c.accountManager.SignTransaction(*tx)
	if err != nil {
		msg := fmt.Sprintf("sign transaction {%+v} error", *tx)
		return "", types.WrapError(err, msg)
	}

	// fmt.Printf("signed raw data: %x", rawData)
	//send raw tx
	txhash, err := c.SendRawTransaction(rawData)
	if err != nil {
		msg := fmt.Sprintf("send raw transaction {%+v} error", rawData)
		return "", types.WrapError(err, msg)
	}
	return txhash, nil
}

// SendRawTransaction sends signed transaction and return its hash.
func (c *Client) SendRawTransaction(rawData []byte) (types.Hash, error) {
	var result interface{}

	if err := c.rpcClient.Call(&result, "cfx_sendRawTransaction", hexutil.Encode(rawData)); err != nil {
		msg := fmt.Sprintf("rpc cfx_sendRawTransaction %+x error", rawData)
		// fmt.Printf("call result when error %+v: %+v", err, result)
		return "", types.WrapError(err, msg)
	}

	return types.Hash(result.(string)), nil
}

// SignEncodedTransactionAndSend sign RLP encoded transaction and send it, return responsed transaction.
func (c *Client) SignEncodedTransactionAndSend(encodedTx []byte, v byte, r, s []byte) (*types.Transaction, error) {
	tx := new(types.UnsignedTransaction)
	err := tx.Decode(encodedTx)
	if err != nil {
		msg := fmt.Sprintf("Decode rlp encoded data {%+v} to unsignTransction error", encodedTx)
		return nil, types.WrapError(err, msg)
	}
	// tx.From = from

	respondTx, err := c.signTransactionAndSend(tx, v, r, s)
	if err != nil {
		msg := fmt.Sprintf("sign transaction and send {tx: %+v, r:%+x, s:%+x, v:%v} error", tx, r, s, v)
		return nil, types.WrapError(err, msg)
	}

	return respondTx, nil
}

func (c *Client) signTransactionAndSend(tx *types.UnsignedTransaction, v byte, r, s []byte) (*types.Transaction, error) {
	rlp, err := tx.EncodeWithSignature(v, r, s)
	if err != nil {
		msg := fmt.Sprintf("encode tx %+v with signature { v:%+x, r:%+x, s:%v} error", tx, v, r, s)
		return nil, types.WrapError(err, msg)
	}

	hash, err := c.SendRawTransaction(rlp)
	if err != nil {
		msg := fmt.Sprintf("send signed tx %+x error", rlp)
		return nil, types.WrapError(err, msg)
	}

	respondTx, err := c.GetTransactionByHash(hash)
	if err != nil {
		msg := fmt.Sprintf("get transaction by hash %+v error", hash)
		return nil, types.WrapError(err, msg)
	}
	return respondTx, nil
}

// Call executes contract but not mined into the blockchain,
// and returns the contract execution result.
func (c *Client) Call(request types.CallRequest, epoch *types.Epoch) (*string, error) {
	var rpcResult interface{}

	args := []interface{}{request}
	// if len(epoch) > 0 {
	if epoch != nil {
		// args = append(args, epoch[0])
		args = append(args, epoch)
	}

	if err := c.rpcClient.Call(&rpcResult, "cfx_call", args...); err != nil {
		msg := fmt.Sprintf("rpc cfx_call {%+v} error", args)
		return nil, types.WrapError(err, msg)
	}

	var resultHexStr string
	if err := utils.UnmarshalRPCResult(rpcResult, &resultHexStr); err != nil {
		msg := fmt.Sprintf("UnmarshalRPCResult %+v error", rpcResult)
		return nil, types.WrapError(err, msg)
	}
	return &resultHexStr, nil
}

// GetLogs returns logs that matching the specified filter.
func (c *Client) GetLogs(filter types.LogFilter) ([]types.Log, error) {
	var result interface{}

	if err := c.rpcClient.Call(&result, "cfx_getLogs", filter); err != nil {
		msg := fmt.Sprintf("rpc cfx_getLogs of {%+v} error", filter)
		return nil, types.WrapError(err, msg)
	}

	var log []types.Log
	if err := utils.UnmarshalRPCResult(result, &log); err != nil {
		msg := fmt.Sprintf("UnmarshalRPCResult %+v error", result)
		return nil, types.WrapError(err, msg)
	}

	return log, nil
}

// GetTransactionByHash returns transaction for the specified hash.
// If transaction not found, return nil.
func (c *Client) GetTransactionByHash(txHash types.Hash) (*types.Transaction, error) {
	var result interface{}

	if err := c.rpcClient.Call(&result, "cfx_getTransactionByHash", txHash); err != nil {
		msg := fmt.Sprintf("rpc cfx_getTransactionByHash {%+v} error", txHash)
		return nil, types.WrapError(err, msg)
	}

	if result == nil {
		return nil, nil
	}

	var tx types.Transaction
	if err := utils.UnmarshalRPCResult(result, &tx); err != nil {
		msg := fmt.Sprintf("UnmarshalRPCResult %+v error", result)
		return nil, types.WrapError(err, msg)
	}

	return &tx, nil
}

// EstimateGas estimates the consumed gas of transaction/contract execution.
//
// deprecated, please use EstimateGasAndCollateral instead.
func (c *Client) EstimateGas(request types.CallRequest, epoch ...*types.Epoch) (*big.Int, error) {
	var result interface{}

	args := []interface{}{request}
	if len(epoch) > 0 {
		args = append(args, epoch[0])
	}

	if err := c.rpcClient.Call(&result, "cfx_estimateGas", args...); err != nil {
		msg := fmt.Sprintf("rpc cfx_estimateGas of {%+v} error", args)
		return nil, types.WrapError(err, msg)
	}

	return hexutil.DecodeBig(result.(string))
}

// EstimateGasAndCollateral estimates the consumed gas and storage for collateral of transaction/contract execution.
func (c *Client) EstimateGasAndCollateral(request types.CallRequest) (*types.Estimate, error) {
	var result interface{}

	args := []interface{}{request}

	if err := c.rpcClient.Call(&result, "cfx_estimateGasAndCollateral", args...); err != nil {
		msg := fmt.Sprintf("rpc cfx_estimateGasAndCollateral of {%+v} error", args)
		return nil, types.WrapError(err, msg)
	}
	var estimate types.Estimate
	if err := utils.UnmarshalRPCResult(result, &estimate); err != nil {
		msg := fmt.Sprintf("UnmarshalRPCResult %+v error", result)
		return nil, types.WrapError(err, msg)
	}

	return &estimate, nil
}

// GetBlocksByEpoch returns the blocks in the specified epoch.
func (c *Client) GetBlocksByEpoch(epoch *types.Epoch) ([]types.Hash, error) {
	var result interface{}

	if err := c.rpcClient.Call(&result, "cfx_getBlocksByEpoch", epoch); err != nil {
		msg := fmt.Sprintf("rpc cfx_getBlocksByEpoch {%+v} error", epoch)
		return nil, types.WrapError(err, msg)
	}

	var blocks []types.Hash
	if err := utils.UnmarshalRPCResult(result, &blocks); err != nil {
		msg := fmt.Sprintf("UnmarshalRPCResult %+v error", result)
		return nil, types.WrapError(err, msg)
	}

	return blocks, nil
}

// GetTransactionReceipt returns the receipt of specified transaction hash.
// If receipt not found, return nil.
func (c *Client) GetTransactionReceipt(txHash types.Hash) (*types.TransactionReceipt, error) {
	var result interface{}

	if err := c.rpcClient.Call(&result, "cfx_getTransactionReceipt", txHash); err != nil {
		msg := fmt.Sprintf("rpc cfx_getTransactionReceipt of {%+v} error", txHash)
		return nil, types.WrapError(err, msg)
	}

	if result == nil {
		return nil, nil
	}

	var receipt types.TransactionReceipt
	if err := utils.UnmarshalRPCResult(result, &receipt); err != nil {
		msg := fmt.Sprintf("UnmarshalRPCResult %+v error", result)
		return nil, types.WrapError(err, msg)
	}

	return &receipt, nil
}

// CreateUnsignedTransaction create an UnsignedTransaction instance
func (c *Client) CreateUnsignedTransaction(from types.Address, to types.Address, amount *hexutil.Big, data *[]byte) (*types.UnsignedTransaction, error) {
	tx := new(types.UnsignedTransaction)
	tx.From = &from
	tx.To = &to
	tx.Value = amount
	tx.Data = *data
	err := c.ApplyUnsignedTransactionDefault(tx)
	if err != nil {
		msg := fmt.Sprintf("apply default field of transaction {%+v} error", tx)
		return nil, types.WrapError(err, msg)
	}

	return tx, nil
}

// ApplyUnsignedTransactionDefault set default value for fields which are empty
func (c *Client) ApplyUnsignedTransactionDefault(tx *types.UnsignedTransaction) error {

	if c != nil {
		if tx.From == nil {
			defaultAccount, err := c.accountManager.GetDefault()
			if err != nil {
				return types.WrapError(err, "get default account error")
			}
			// fmt.Printf("default account is: %+v\n\n", defaultAccount)

			if defaultAccount == nil {
				return errors.New("no account exist in keystore directory")
			}
			tx.From = defaultAccount
		}

		if tx.Nonce == 0 {
			nonce, err := c.GetNextNonce(*tx.From)
			if err != nil {
				msg := fmt.Sprintf("get nonce of {%+v} error", tx.From)
				return types.WrapError(err, msg)
			}
			tx.Nonce = nonce
		}

		if tx.GasPrice == nil {
			gasPrice, err := c.GetGasPrice()
			if err != nil {
				msg := "get gas price error"
				return types.WrapError(err, msg)
			}
			// conflux responsed gasprice offen be 0, but the min gasprice is 1 when sending transaction, so do this
			if gasPrice.Cmp(big.NewInt(constants.MinGasprice)) < 1 {
				gasPrice = big.NewInt(1)
			}
			tmp := hexutil.Big(*gasPrice)
			tx.GasPrice = &tmp
		}

		if tx.EpochHeight == nil {
			epoch, err := c.GetEpochNumber(types.EpochLatestState)
			if err != nil {
				msg := fmt.Sprintf("get epoch number of {%+v} error", types.EpochLatestState)
				return types.WrapError(err, msg)
			}
			tx.EpochHeight = (*hexutil.Big)(epoch)
		}

		// The gas and storage limit may influnce by all fileds of transaction ,so set them in last step.
		if tx.StorageLimit == nil || tx.Gas == 0 {
			callReq := new(types.CallRequest)
			callReq.FillByUnsignedTx(tx)

			sm, err := c.EstimateGasAndCollateral(*callReq)
			if err != nil {
				msg := fmt.Sprintf("get estimate gas and collateral by {%+v} error", *callReq)
				return types.WrapError(err, msg)
			}

			if tx.Gas == 0 {
				tx.Gas = sm.GasUsed.ToInt().Uint64()
			}

			if tx.StorageLimit == nil {
				tx.StorageLimit = sm.StorageCollateralized
			}
		}

		tx.ApplyDefault()
	}

	return nil
}

// Debug calls the Conflux debug API.
func (c *Client) Debug(method string, args ...interface{}) (interface{}, error) {
	var result interface{}

	if err := c.rpcClient.Call(&result, method, args...); err != nil {
		msg := fmt.Sprintf("rpc call method {%+v} with args {%+v} error", method, args)
		return nil, types.WrapError(err, msg)
	}

	return result, nil
}

// DeployContract Synchronize deploy a contract by abi, bytecode and options, returns a bool channel for notifying whether contract deploy is successful. The callback for handle
// deploy result.
func (c *Client) DeployContract(abiJSON string, bytecode []byte, option *types.ContractDeployOption, timeout time.Duration, callback func(deployedContract Contractor, hash *types.Hash, err error)) <-chan bool {
	// func (c *Client) DeployContract(abiJSON string, tx *types.UnsignedTransaction, timeout time.Duration) (<-chan bool, error) {
	doneChan := make(chan bool, 1)

	tx := new(types.UnsignedTransaction)
	if option != nil {
		tx.UnsignedTransactionBase = option.UnsignedTransactionBase
	}
	tx.Data = bytecode

	// if option.Data == nil {
	// 	msg := fmt.Sprintf("the Data field of options {%+v} should be bytecode, but it is empty", option)
	// 	callback(nil, nil, errors.New(msg))
	// 	doneChan <- false
	// 	return doneChan
	// }

	//deploy contract
	txhash, err := c.SendTransaction(tx)
	if err != nil {
		msg := fmt.Sprintf("send transaction {%+v} error", tx)
		callback(nil, nil, types.WrapError(err, msg))
		doneChan <- false
		return doneChan
		// return nil, types.WrapError(err, msg)
	}

	var abi abi.ABI
	err = abi.UnmarshalJSON([]byte(abiJSON))
	if err != nil {
		msg := fmt.Sprintf("unmarshal json {%+v} to ABI error", abiJSON)
		callback(nil, nil, types.WrapError(err, msg))
		doneChan <- false
		return doneChan
		// return nil, types.WrapError(err, msg)
	}

	// wait tx be confirmed and excute callback
	if timeout == 0 {
		timeout = 30 * time.Second
	}

	go func(_txhash types.Hash) {
		for i := 0; i < 10; i++ {
			transaction, err := c.GetTransactionByHash(txhash)
			if err != nil {
				msg := fmt.Sprintf("get transaction receipt of txhash %+v error", txhash)
				callback(nil, &_txhash, types.WrapError(err, msg))
				doneChan <- false
				return
			}

			if transaction.Status != nil {
				if transaction.Status.ToInt().Uint64() == 1 {
					msg := fmt.Sprintf("transaction is packed but it is failed,the txhash is %+v", _txhash)
					callback(nil, &_txhash, errors.New(msg))
					doneChan <- false
					return
				}

				contract := &Contract{abi, c, transaction.ContractCreated}
				callback(contract, &_txhash, nil)
				doneChan <- true
				return
			}
			time.Sleep(3 * time.Second)
		}

		msg := fmt.Sprintf("deploy contract timeout after %+v seconds, txhash is %+v", timeout, _txhash)
		callback(nil, &_txhash, errors.New(msg))
		doneChan <- false

	}(txhash)
	return doneChan
}

// GetContract create a contract instance according to abi json and it's deployed address
func (c *Client) GetContract(abiJSON string, deployedOn *types.Address) (*Contract, error) {
	var abi abi.ABI
	err := abi.UnmarshalJSON([]byte(abiJSON))
	if err != nil {
		msg := fmt.Sprintf("unmarshal json {%+v} to ABI error", abiJSON)
		return nil, types.WrapError(err, msg)
	}

	// var contract IContract
	contract := &Contract{abi, c, deployedOn}
	return contract, nil
}

// Close closes the client, aborting any in-flight requests.
func (c *Client) Close() {
	c.rpcClient.Close()
}
