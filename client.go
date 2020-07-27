// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/constants"
	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Client represents a client to interact with Conflux blockchain.
type Client struct {
	nodeURL        string
	rpcRequester   rpcRequester
	accountManager AccountManagerOperator
}

// NewClient creates a new instance of Client with specified conflux node url.
func NewClient(nodeURL string) (*Client, error) {
	client, err := NewClientWithRetry(nodeURL, 0, 0)
	return client, err
}

// NewClientWithRetry creates a retryable new instance of Client with specified conflux node url and retry options.
//
// the retryInterval will be set to 1 second if pass 0
func NewClientWithRetry(nodeURL string, retryCount int, retryInterval time.Duration) (*Client, error) {

	var client Client
	client.nodeURL = nodeURL

	rpcClient, err := rpc.Dial(nodeURL)
	if err != nil {
		return nil, types.WrapError(err, "dail failed")
	}

	if retryCount == 0 {
		client.rpcRequester = rpcClient
	} else {
		// Interval 0 is meaningless and may lead full node busy, so default sets it to 1 second
		if retryInterval == 0 {
			retryInterval = time.Second
		}

		client.rpcRequester = &rpcClientWithRetry{
			inner:      rpcClient,
			retryCount: retryCount,
			interval:   retryInterval,
		}
	}

	return &client, nil
}

// NewClientWithRPCRequester creates client with specified rpcRequester
func NewClientWithRPCRequester(rpcRequester rpcRequester) (*Client, error) {
	return &Client{
		rpcRequester: rpcRequester,
	}, nil
}

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

		if types.IsRpcJsonError(err) {
			return err
		}

		remain--
		fmt.Printf("remain retry count: %v\n", remain)
		if remain < 0 {
			msg := fmt.Sprintf("timeout when call %v with args %v", method, args)
			return types.WrapError(err, msg)
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
			msg := fmt.Sprintf("timeout when batch call %+v", b)
			return types.WrapError(err, msg)
		}

		if r.interval > 0 {
			time.Sleep(r.interval)
		}
	}
}

func (r *rpcClientWithRetry) Close() {
	if r != nil && r.inner != nil {
		r.inner.Close()
	}
}

// GetNodeURL returns node url
func (client *Client) GetNodeURL() string {
	return client.nodeURL
}

// CallRPC performs a JSON-RPC call with the given arguments and unmarshals into
// result if no error occurred.
//
// The result must be a pointer so that package json can unmarshal into it. You
// can also pass nil, in which case the result is ignored.
func (client *Client) CallRPC(result interface{}, method string, args ...interface{}) error {
	return client.rpcRequester.Call(result, method, args...)
}

// BatchCallRPC sends all given requests as a single batch and waits for the server
// to return a response for all of them.
//
// In contrast to Call, BatchCall only returns I/O errors. Any error specific to
// a request is reported through the Error field of the corresponding BatchElem.
//
// Note that batch calls may not be executed atomically on the server side.
func (client *Client) BatchCallRPC(b []rpc.BatchElem) error {
	return client.rpcRequester.BatchCall(b)
}

// SetAccountManager sets account manager for sign transaction
func (client *Client) SetAccountManager(accountManager AccountManagerOperator) {
	client.accountManager = accountManager
}

// GetGasPrice returns the recent mean gas price.
func (client *Client) GetGasPrice() (*big.Int, error) {
	var result interface{}

	if err := client.rpcRequester.Call(&result, "cfx_gasPrice"); err != nil {
		msg := "rpc request cfx_gasPrice error"
		return nil, types.WrapError(err, msg)
	}

	return hexutil.DecodeBig(result.(string))
}

// GetNextNonce returns the next transaction nonce of address
func (client *Client) GetNextNonce(address types.Address, epoch *types.Epoch) (*big.Int, error) {
	var result interface{}
	args := []interface{}{address}
	if epoch != nil {
		args = append(args, epoch)
	}

	if err := client.rpcRequester.Call(&result, "cfx_getNextNonce", args...); err != nil {
		msg := fmt.Sprintf("rpc request cfx_getNextNonce %+v error", address)
		return nil, types.WrapErrorf(err, msg)
	}
	return hexutil.DecodeBig(result.(string))
}

// GetStatus returns chainID of connecting conflux node
func (client *Client) GetStatus() (*types.Status, error) {
	var result types.Status

	if err := client.rpcRequester.Call(&result, "cfx_getStatus"); err != nil {
		return nil, types.WrapErrorf(err, "rpc request cfx_getStatus error")
	}
	return &result, nil
}

// GetEpochNumber returns the highest or specified epoch number.
func (client *Client) GetEpochNumber(epoch ...*types.Epoch) (*big.Int, error) {
	var result interface{}

	var args []interface{}
	if len(epoch) > 0 {
		args = append(args, epoch[0])
	}

	if err := client.rpcRequester.Call(&result, "cfx_epochNumber", args...); err != nil {
		msg := fmt.Sprintf("rpc cfx_epochNumber %+v error", args)
		return nil, types.WrapError(err, msg)
	}

	return hexutil.DecodeBig(result.(string))
}

// GetBalance returns the balance of specified address at epoch.
func (client *Client) GetBalance(address types.Address, epoch ...*types.Epoch) (*big.Int, error) {
	var result interface{}

	args := []interface{}{address}
	if len(epoch) > 0 {
		args = append(args, epoch[0])
	}

	if err := client.rpcRequester.Call(&result, "cfx_getBalance", args...); err != nil {
		msg := fmt.Sprintf("rpc cfx_getBalance %+v error", args)
		return nil, types.WrapError(err, msg)
	}

	return hexutil.DecodeBig(result.(string))
}

// GetCode returns the bytecode in HEX format of specified address at epoch.
func (client *Client) GetCode(address types.Address, epoch ...*types.Epoch) (string, error) {
	var result interface{}

	args := []interface{}{address}
	if len(epoch) > 0 {
		args = append(args, epoch[0])
	}

	if err := client.rpcRequester.Call(&result, "cfx_getCode", args...); err != nil {
		msg := fmt.Sprintf("rpc cfx_getCode %+v error", args)
		return "", types.WrapError(err, msg)
	}
	// fmt.Printf("get code of %v res: %+v, result.(string) is %v\n", address, result, result.(string))

	return result.(string), nil
}

// GetBlockSummaryByHash returns the block summary of specified blockHash
// If the block is not found, return nil.
func (client *Client) GetBlockSummaryByHash(blockHash types.Hash) (*types.BlockSummary, error) {
	var result interface{}

	if err := client.rpcRequester.Call(&result, "cfx_getBlockByHash", blockHash, false); err != nil {
		msg := fmt.Sprintf("rpc cfx_getBlockByHash %+v error", blockHash)
		return nil, types.WrapError(err, msg)
	}

	if result == nil {
		return nil, nil
	}

	var block types.BlockSummary
	if err := unmarshalRPCResult(result, &block); err != nil {
		msg := fmt.Sprintf("UnmarshalRPCResult %+v error", result)
		return nil, types.WrapError(err, msg)
	}

	return &block, nil
}

// GetBlockByHash returns the block of specified blockHash
// If the block is not found, return nil.
func (client *Client) GetBlockByHash(blockHash types.Hash) (*types.Block, error) {
	var result interface{}

	if err := client.rpcRequester.Call(&result, "cfx_getBlockByHash", blockHash, true); err != nil {
		msg := fmt.Sprintf("rpc cfx_getBlockByHash %+v error", blockHash)
		return nil, types.WrapError(err, msg)
	}

	if result == nil {
		return nil, nil
	}

	var block types.Block
	if err := unmarshalRPCResult(result, &block); err != nil {
		msg := fmt.Sprintf("UnmarshalRPCResult %+v error", result)
		return nil, types.WrapError(err, msg)
	}

	return &block, nil
}

// GetBlockSummaryByEpoch returns the block summary of specified epoch.
// If the epoch is invalid, return the concrete error.
func (client *Client) GetBlockSummaryByEpoch(epoch *types.Epoch) (*types.BlockSummary, error) {
	var result interface{}

	if err := client.rpcRequester.Call(&result, "cfx_getBlockByEpochNumber", epoch, false); err != nil {
		msg := fmt.Sprintf("rpc cfx_getBlockByEpochNumber %+v error", epoch)
		return nil, types.WrapError(err, msg)
	}

	var block types.BlockSummary
	if err := unmarshalRPCResult(result, &block); err != nil {
		msg := fmt.Sprintf("UnmarshalRPCResult %+v error", result)
		return nil, types.WrapError(err, msg)
	}

	return &block, nil
}

// GetBlockByEpoch returns the block of specified epoch.
// If the epoch is invalid, return the concrete error.
func (client *Client) GetBlockByEpoch(epoch *types.Epoch) (*types.Block, error) {
	var result interface{}

	if err := client.rpcRequester.Call(&result, "cfx_getBlockByEpochNumber", epoch, true); err != nil {
		msg := fmt.Sprintf("rpc cfx_getBlockByEpochNumber %+v error", epoch)
		return nil, types.WrapError(err, msg)
	}

	var block types.Block
	if err := unmarshalRPCResult(result, &block); err != nil {
		msg := fmt.Sprintf("UnmarshalRPCResult %+v error", result)
		return nil, types.WrapError(err, msg)
	}

	return &block, nil
}

// GetBestBlockHash returns the current best block hash.
func (client *Client) GetBestBlockHash() (types.Hash, error) {
	var result interface{}

	if err := client.rpcRequester.Call(&result, "cfx_getBestBlockHash"); err != nil {
		msg := "rpc cfx_getBestBlockHash error"
		return "", types.WrapError(err, msg)
	}

	return types.Hash(result.(string)), nil
}

// GetRawBlockConfirmationRisk indicates the risk coefficient that
// the pivot block of the epoch where the block is located becomes a normal block.
func (client *Client) GetRawBlockConfirmationRisk(blockhash types.Hash) (*big.Int, error) {
	var result interface{}

	args := []interface{}{blockhash}

	if err := client.rpcRequester.Call(&result, "cfx_getConfirmationRiskByHash", args...); err != nil {
		msg := fmt.Sprintf("rpc cfx_getConfirmationRiskByHash %+v error", args)
		return nil, types.WrapError(err, msg)
	}

	if result == nil {

		block, err := client.GetBlockSummaryByHash(blockhash)
		if err != nil {
			msg := fmt.Sprintf("get block by hash %+v error", blockhash)
			return nil, types.WrapError(err, msg)
		}
		if block != nil && block.EpochNumber != nil {
			return big.NewInt(0), nil
		}

		return constants.MaxUint256, nil
	}

	return hexutil.DecodeBig(result.(string))
}

// GetBlockConfirmationRisk indicates the probability that
// the pivot block of the epoch where the block is located becomes a normal block.
//
// it's (raw confirmation risk coefficient/ (2^256-1))
func (client *Client) GetBlockConfirmationRisk(blockHash types.Hash) (*big.Float, error) {
	risk, err := client.GetRawBlockConfirmationRisk(blockHash)
	if err != nil {
		msg := fmt.Sprintf("get block confirmation risk by hash %+v error", blockHash)
		return nil, types.WrapError(err, msg)
	}
	if risk == nil {
		return nil, nil
	}

	riskFloat := new(big.Float).SetInt(risk)
	maxUint256Float := new(big.Float).SetInt(constants.MaxUint256)

	riskRate := new(big.Float).Quo(riskFloat, maxUint256Float)
	return riskRate, nil
}

// SendTransaction signs and sends transaction to conflux node and returns the transaction hash.
func (client *Client) SendTransaction(tx *types.UnsignedTransaction) (types.Hash, error) {

	err := client.ApplyUnsignedTransactionDefault(tx)
	if err != nil {
		msg := fmt.Sprintf("apply transaction {%+v} default fields error", *tx)
		return "", types.WrapError(err, msg)
	}

	// commet it becasue there are some contract need not pay gas.
	//
	// //check balance, return error if balance not enough
	// epoch := types.NewEpochNumber(tx.EpochHeight.ToInt())
	// balance, err := c.GetBalance(*tx.From, epoch)
	// if err != nil {
	// 	msg := fmt.Sprintf("get balance of %+v at ephoc %+v error", tx.From, epoch)
	// 	return "", types.WrapError(err, msg)
	// }
	// need := big.NewInt(int64(tx.Gas))
	// need = need.Add(tx.StorageLimit.ToInt(), need)
	// need = need.Mul(tx.GasPrice.ToInt(), need)
	// need = need.Add(tx.Value.ToInt(), need)
	// need = need.Add(tx.StorageLimit.ToInt(), need)

	// if balance.Cmp(need) < 0 {
	// 	msg := fmt.Sprintf("out of balance, need %+v but your balance is %+v", need, balance)
	// 	return "", types.WrapError(err, msg)
	// }

	//sign
	if client.accountManager == nil {
		msg := fmt.Sprintf("sign transaction need account manager, please call SetAccountManager to set it.")
		return "", errors.New(msg)
	}

	rawData, err := client.accountManager.SignTransaction(*tx)
	if err != nil {
		msg := fmt.Sprintf("sign transaction {%+v} error", *tx)
		return "", types.WrapError(err, msg)
	}

	//send raw tx
	txhash, err := client.SendRawTransaction(rawData)
	if err != nil {
		msg := fmt.Sprintf("send raw transaction 0x%+x error", rawData)
		return "", types.WrapError(err, msg)
	}
	return txhash, nil
}

// SendRawTransaction sends signed transaction and returns its hash.
func (client *Client) SendRawTransaction(rawData []byte) (types.Hash, error) {
	var result interface{}

	if err := client.rpcRequester.Call(&result, "cfx_sendRawTransaction", hexutil.Encode(rawData)); err != nil {
		msg := fmt.Sprintf("rpc cfx_sendRawTransaction 0x%+x error", rawData)
		return "", types.WrapError(err, msg)
	}

	return types.Hash(result.(string)), nil
}

// SignEncodedTransactionAndSend signs RLP encoded transaction "encodedTx" by signature "r,s,v" and sends it to node,
// and returns responsed transaction.
func (client *Client) SignEncodedTransactionAndSend(encodedTx []byte, v byte, r, s []byte) (*types.Transaction, error) {
	tx := new(types.UnsignedTransaction)
	err := tx.Decode(encodedTx)
	if err != nil {
		msg := fmt.Sprintf("Decode rlp encoded data {%+v} to unsignTransction error", encodedTx)
		return nil, types.WrapError(err, msg)
	}
	// tx.From = from

	respondTx, err := client.signTransactionAndSend(tx, v, r, s)
	if err != nil {
		msg := fmt.Sprintf("sign transaction and send {tx: %+v, r:%+x, s:%+x, v:%v} error", tx, r, s, v)
		return nil, types.WrapError(err, msg)
	}

	return respondTx, nil
}

func (client *Client) signTransactionAndSend(tx *types.UnsignedTransaction, v byte, r, s []byte) (*types.Transaction, error) {
	rlp, err := tx.EncodeWithSignature(v, r, s)
	if err != nil {
		msg := fmt.Sprintf("encode tx %+v with signature { v:%+x, r:%+x, s:%v} error", tx, v, r, s)
		return nil, types.WrapError(err, msg)
	}

	hash, err := client.SendRawTransaction(rlp)
	if err != nil {
		msg := fmt.Sprintf("send signed tx %+x error", rlp)
		return nil, types.WrapError(err, msg)
	}

	respondTx, err := client.GetTransactionByHash(hash)
	if err != nil {
		msg := fmt.Sprintf("get transaction by hash %+v error", hash)
		return nil, types.WrapError(err, msg)
	}
	return respondTx, nil
}

// Call executes a message call transaction "request" at specified epoch,
// which is directly executed in the VM of the node, but never mined into the block chain
// and returns the contract execution result.
func (client *Client) Call(request types.CallRequest, epoch *types.Epoch) (*string, error) {
	var rpcResult interface{}

	args := []interface{}{request}
	// if len(epoch) > 0 {
	if epoch != nil {
		// args = append(args, epoch[0])
		args = append(args, epoch)
	}

	if err := client.rpcRequester.Call(&rpcResult, "cfx_call", args...); err != nil {
		msg := fmt.Sprintf("rpc cfx_call {%+v} error", args)
		return nil, types.WrapError(err, msg)
	}

	var resultHexStr string
	if err := unmarshalRPCResult(rpcResult, &resultHexStr); err != nil {
		msg := fmt.Sprintf("UnmarshalRPCResult %+v error", rpcResult)
		return nil, types.WrapError(err, msg)
	}
	return &resultHexStr, nil
}

// GetLogs returns logs that matching the specified filter.
func (client *Client) GetLogs(filter types.LogFilter) ([]types.Log, error) {
	var result interface{}

	if err := client.rpcRequester.Call(&result, "cfx_getLogs", filter); err != nil {
		msg := fmt.Sprintf("rpc cfx_getLogs of {%+v} error", filter)
		return nil, types.WrapError(err, msg)
	}

	var log []types.Log
	if err := unmarshalRPCResult(result, &log); err != nil {
		msg := fmt.Sprintf("UnmarshalRPCResult %+v error", result)
		return nil, types.WrapError(err, msg)
	}

	return log, nil
}

// GetTransactionByHash returns transaction for the specified txHash.
// If the transaction is not found, return nil.
func (client *Client) GetTransactionByHash(txHash types.Hash) (*types.Transaction, error) {
	var result interface{}

	if err := client.rpcRequester.Call(&result, "cfx_getTransactionByHash", txHash); err != nil {
		msg := fmt.Sprintf("rpc cfx_getTransactionByHash {%+v} error", txHash)
		return nil, types.WrapError(err, msg)
	}

	if result == nil {
		return nil, nil
	}

	var tx types.Transaction
	if err := unmarshalRPCResult(result, &tx); err != nil {
		msg := fmt.Sprintf("UnmarshalRPCResult %+v error", result)
		return nil, types.WrapError(err, msg)
	}

	return &tx, nil
}

// EstimateGasAndCollateral excutes a message call "request"
// and returns the amount of the gas used and storage for collateral
func (client *Client) EstimateGasAndCollateral(request types.CallRequest) (*types.Estimate, error) {
	var result interface{}

	args := []interface{}{request}

	if err := client.rpcRequester.Call(&result, "cfx_estimateGasAndCollateral", args...); err != nil {
		msg := fmt.Sprintf("rpc cfx_estimateGasAndCollateral of {%+v} error", args)
		return nil, types.WrapError(err, msg)
	}
	var estimate types.Estimate
	if err := unmarshalRPCResult(result, &estimate); err != nil {
		msg := fmt.Sprintf("UnmarshalRPCResult %+v error", result)
		return nil, types.WrapError(err, msg)
	}

	return &estimate, nil
}

// GetBlocksByEpoch returns the blocks hash in the specified epoch.
func (client *Client) GetBlocksByEpoch(epoch *types.Epoch) ([]types.Hash, error) {
	var result interface{}

	if err := client.rpcRequester.Call(&result, "cfx_getBlocksByEpoch", epoch); err != nil {
		msg := fmt.Sprintf("rpc cfx_getBlocksByEpoch {%+v} error", epoch)
		return nil, types.WrapError(err, msg)
	}

	var blocks []types.Hash
	if err := unmarshalRPCResult(result, &blocks); err != nil {
		msg := fmt.Sprintf("UnmarshalRPCResult %+v error", result)
		return nil, types.WrapError(err, msg)
	}

	return blocks, nil
}

// GetTransactionReceipt returns the receipt of specified transaction hash.
// If no receipt is found, return nil.
func (client *Client) GetTransactionReceipt(txHash types.Hash) (*types.TransactionReceipt, error) {
	var result interface{}

	if err := client.rpcRequester.Call(&result, "cfx_getTransactionReceipt", txHash); err != nil {
		msg := fmt.Sprintf("rpc cfx_getTransactionReceipt of {%+v} error", txHash)
		return nil, types.WrapError(err, msg)
	}

	if result == nil {
		return nil, nil
	}

	var receipt types.TransactionReceipt
	if err := unmarshalRPCResult(result, &receipt); err != nil {
		msg := fmt.Sprintf("UnmarshalRPCResult %+v error", result)
		return nil, types.WrapError(err, msg)
	}

	return &receipt, nil
}

// CreateUnsignedTransaction creates an unsigned transaction by parameters,
// and the other fields will be set to values fetched from conflux node.
func (client *Client) CreateUnsignedTransaction(from types.Address, to types.Address, amount *hexutil.Big, data []byte) (*types.UnsignedTransaction, error) {
	tx := new(types.UnsignedTransaction)
	tx.From = &from
	tx.To = &to
	tx.Value = amount
	tx.Data = data

	err := client.ApplyUnsignedTransactionDefault(tx)
	if err != nil {
		msg := fmt.Sprintf("apply default field of transaction {%+v} error", tx)
		return nil, types.WrapError(err, msg)
	}

	return tx, nil
}

// ApplyUnsignedTransactionDefault set empty fields to value fetched from conflux node.
func (client *Client) ApplyUnsignedTransactionDefault(tx *types.UnsignedTransaction) error {

	if client != nil {
		if tx.From == nil {
			if client.accountManager != nil {
				defaultAccount, err := client.accountManager.GetDefault()
				if err != nil {
					return types.WrapError(err, "get default account error")
				}

				if defaultAccount == nil {
					return errors.New("no account exist in keystore directory")
				}
				tx.From = defaultAccount
			}
		}

		if tx.Nonce == nil {
			nonce, err := client.GetNextNonce(*tx.From, nil)
			if err != nil {
				msg := fmt.Sprintf("get nonce of {%+v} error", tx.From)
				return types.WrapError(err, msg)
			}
			tmp := hexutil.Big(*nonce)
			tx.Nonce = &tmp
		}

		if tx.ChainID == nil {
			status, err := client.GetStatus()
			if err != nil {
				tx.ChainID = types.NewUint(0)
			} else {
				tx.ChainID = status.ChainID
			}
		}

		if tx.GasPrice == nil {
			gasPrice, err := client.GetGasPrice()
			if err != nil {
				msg := "get gas price error"
				return types.WrapError(err, msg)
			}

			// conflux responsed gasprice offen be 0, but the min gasprice is 1 when sending transaction, so do this
			if gasPrice.Cmp(big.NewInt(constants.MinGasprice)) < 1 {
				gasPrice = big.NewInt(constants.MinGasprice)
			}
			tmp := hexutil.Big(*gasPrice)
			tx.GasPrice = &tmp
		}

		if tx.EpochHeight == nil {
			epoch, err := client.GetEpochNumber(types.EpochLatestState)
			if err != nil {
				msg := fmt.Sprintf("get epoch number of {%+v} error", types.EpochLatestState)
				return types.WrapError(err, msg)
			}
			// tx.EpochHeight = (*hexutil.Big)(epoch).toi
			tx.EpochHeight = types.NewUint64(epoch.Uint64())
		}

		// The gas and storage limit may be influnced by all fileds of transaction ,so set them at last step.
		if tx.StorageLimit == nil || tx.Gas == nil {
			callReq := new(types.CallRequest)
			callReq.FillByUnsignedTx(tx)

			sm, err := client.EstimateGasAndCollateral(*callReq)
			if err != nil {
				msg := fmt.Sprintf("get estimate gas and collateral by {%+v} error", *callReq)
				return types.WrapError(err, msg)
			}

			if tx.Gas == nil {
				tx.Gas = sm.GasUsed
			}

			if tx.StorageLimit == nil {
				tx.StorageLimit = types.NewUint64(sm.StorageCollateralized.ToInt().Uint64())
			}
		}

		tx.ApplyDefault()
	}

	return nil
}

// Debug calls the Conflux debug API.
func (client *Client) Debug(method string, args ...interface{}) (interface{}, error) {
	var result interface{}

	if err := client.rpcRequester.Call(&result, method, args...); err != nil {
		msg := fmt.Sprintf("rpc call method {%+v} with args {%+v} error", method, args)
		return nil, types.WrapError(err, msg)
	}

	return result, nil
}

// DeployContract deploys a contract by abiJSON, bytecode and consturctor params.
// It returns a ContractDeployState instance which contains 3 channels for notifying when state changed.
func (client *Client) DeployContract(option *types.ContractDeployOption, abiJSON []byte,
	bytecode []byte, constroctorParams ...interface{}) *ContractDeployResult {

	doneChan := make(chan struct{})
	result := ContractDeployResult{DoneChannel: doneChan}

	go func() {

		defer func() {
			doneChan <- struct{}{}
			close(doneChan)
		}()

		//generate ABI
		var abi abi.ABI
		err := abi.UnmarshalJSON([]byte(abiJSON))
		if err != nil {
			msg := fmt.Sprintf("unmarshal json {%+v} to ABI error", abiJSON)
			result.Error = types.WrapError(err, msg)
			return
		}

		tx := new(types.UnsignedTransaction)
		if option != nil {
			tx.UnsignedTransactionBase = types.UnsignedTransactionBase(option.UnsignedTransactionBase)
		}

		//recreate contract bytecode with consturctor params
		if len(constroctorParams) > 0 {
			input, err := abi.Pack("", constroctorParams...)
			if err != nil {
				msg := fmt.Sprintf("encode constrctor with args %+v error", constroctorParams)
				result.Error = types.WrapError(err, msg)
				return
			}

			bytecode = append(bytecode, input...)
		}
		tx.Data = bytecode

		//deploy contract
		txhash, err := client.SendTransaction(tx)
		if err != nil {
			msg := fmt.Sprintf("send transaction {%+v} error", tx)
			result.Error = types.WrapError(err, msg)
			return
		}
		result.TransactionHash = &txhash

		// timeout := time.After(time.Duration(_timeoutIns) * time.Second)
		timeout := time.After(3600 * time.Second)
		if option != nil && option.Timeout != 0 {
			timeout = time.After(option.Timeout)
		}

		ticker := time.Tick(2000 * time.Millisecond)
		// Keep trying until we're time out or get a result or get an error
		for {
			select {
			// Got a timeout! fail with a timeout error
			case t := <-timeout:
				msg := fmt.Sprintf("deploy contract time out after %v, txhash is %+v", t, txhash)
				result.Error = errors.New(msg)
				return
			// Got a tick
			case <-ticker:
				transaction, err := client.GetTransactionByHash(txhash)
				if err != nil {
					msg := fmt.Sprintf("get transaction receipt of txhash %+v error", txhash)
					result.Error = types.WrapError(err, msg)
					return
				}

				if transaction.Status != nil {
					if *transaction.Status == 1 {
						msg := fmt.Sprintf("transaction is packed but it is failed,the txhash is %+v", txhash)
						result.Error = errors.New(msg)
						return
					}

					result.DeployedContract = &Contract{abi, client, transaction.ContractCreated}
					return
				}
			}
		}
	}()
	return &result
}

// GetContract creates a contract instance according to abi json and it's deployed address
func (client *Client) GetContract(abiJSON []byte, deployedAt *types.Address) (*Contract, error) {
	var abi abi.ABI
	err := abi.UnmarshalJSON([]byte(abiJSON))
	if err != nil {
		msg := fmt.Sprintf("unmarshal json {%+v} to ABI error", abiJSON)
		return nil, types.WrapError(err, msg)
	}

	contract := &Contract{abi, client, deployedAt}
	return contract, nil
}

// BatchGetTxByHashes requests transaction informations in bulk by txhashes
func (client *Client) BatchGetTxByHashes(txhashes []types.Hash) (map[types.Hash]*types.Transaction, error) {
	if txhashes == nil || len(txhashes) == 0 {
		return make(map[types.Hash]*types.Transaction), nil
	}

	cache := make(map[types.Hash]*rpc.BatchElem)
	for i := range txhashes {
		if cache[txhashes[i]] == nil {
			be := rpc.BatchElem{
				Method: "cfx_getTransactionByHash",
				Args:   []interface{}{txhashes[i]},
				Result: &types.Transaction{},
			}
			cache[txhashes[i]] = &be
		}
	}

	bes := make([]rpc.BatchElem, 0, len(cache))
	for _, v := range cache {
		bes = append(bes, *v)
	}
	// fmt.Printf("send BatchItems: %+v \n", bes)
	if err := client.BatchCallRPC(bes); err != nil {
		return nil, err
	}

	hashToTxMap := make(map[types.Hash]*types.Transaction)
	for _, th := range txhashes {
		be := cache[th]
		if (be.Result == types.Transaction{}) {
			hashToTxMap[th] = nil
			continue
		}
		hashToTxMap[th] = be.Result.(*types.Transaction)
	}

	return hashToTxMap, nil
}

// BatchGetBlockSummarys requests block summary informations in bulk by blockhashes
func (client *Client) BatchGetBlockSummarys(blockhashes []types.Hash) (map[types.Hash]*types.BlockSummary, error) {

	if blockhashes == nil || len(blockhashes) == 0 {
		return make(map[types.Hash]*types.BlockSummary), nil
	}

	cache := make(map[types.Hash]*rpc.BatchElem)

	for i := range blockhashes {
		if cache[blockhashes[i]] == nil {
			be := rpc.BatchElem{
				Method: "cfx_getBlockByHash",
				Args:   []interface{}{blockhashes[i], false},
				Result: &types.BlockSummary{},
			}
			cache[blockhashes[i]] = &be
		}
	}

	// generate bes
	bes := make([]rpc.BatchElem, 0, len(cache))
	for _, v := range cache {
		bes = append(bes, *v)
	}

	if err := client.BatchCallRPC(bes); err != nil {
		return nil, err
	}

	hashToBlocksummaryMap := make(map[types.Hash]*types.BlockSummary)

	for _, bh := range blockhashes {
		be := cache[bh]
		if (be.Result == types.Transaction{}) {
			hashToBlocksummaryMap[bh] = nil
			continue
		}
		hashToBlocksummaryMap[bh] = be.Result.(*types.BlockSummary)
	}
	return hashToBlocksummaryMap, nil
}

// BatchGetRawBlockConfirmationRisk requests raw confirmation risk informations in bulk by blockhashes
func (client *Client) BatchGetRawBlockConfirmationRisk(blockhashes []types.Hash) (map[types.Hash]*big.Int, error) {

	if blockhashes == nil || len(blockhashes) == 0 {
		return make(map[types.Hash]*big.Int), nil
	}

	// get risks
	riskCache := make(map[types.Hash]*rpc.BatchElem)
	for i := range blockhashes {
		if riskCache[blockhashes[i]] == nil {
			var riskStr string
			be := rpc.BatchElem{
				Method: "cfx_getConfirmationRiskByHash",
				Args:   []interface{}{blockhashes[i]},
				Result: &riskStr,
			}
			riskCache[blockhashes[i]] = &be
		}
	}

	bes := make([]rpc.BatchElem, 0, len(riskCache))
	for _, v := range riskCache {
		bes = append(bes, *v)
	}

	if err := client.BatchCallRPC(bes); err != nil {
		return nil, err
	}

	// get block summary of blockhashes without risk
	noRiskBlockhashes := make([]types.Hash, 0)
	for _, bh := range blockhashes {
		be := riskCache[bh]
		if len(*be.Result.(*string)) == 0 {
			noRiskBlockhashes = append(noRiskBlockhashes, bh)
		}
	}

	hashToBlocksummaryMap := make(map[types.Hash]*types.BlockSummary)
	if len(noRiskBlockhashes) > 0 {
		var err error
		hashToBlocksummaryMap, err = client.BatchGetBlockSummarys(noRiskBlockhashes)
		if err != nil {
			return nil, err
		}
	}

	hashToRiskMap := make(map[types.Hash]*big.Int)
	for _, bh := range blockhashes {
		be := riskCache[bh]
		riskStr := *be.Result.(*string)
		if len(riskStr) == 0 {
			blkSummary := hashToBlocksummaryMap[bh]
			if blkSummary != nil && blkSummary.EpochNumber != nil {
				hashToRiskMap[bh] = big.NewInt(0)
			} else {
				hashToRiskMap[bh] = constants.MaxUint256
			}
			continue
			// hashToRiskMap[bh] = nil
			// continue
		}
		risk, err := hexutil.DecodeBig(riskStr)
		if err != nil {
			return nil, err
		}
		hashToRiskMap[bh] = risk
	}
	return hashToRiskMap, nil
}

// BatchGetBlockConfirmationRisk acquires confirmation risk informations in bulk by blockhashes
func (client *Client) BatchGetBlockConfirmationRisk(blockhashes []types.Hash) (map[types.Hash]*big.Float, error) {
	hashToRiskMap, err := client.BatchGetRawBlockConfirmationRisk(blockhashes)
	if err != nil {
		return nil, err
	}

	hashToRevertRateMap := make(map[types.Hash]*big.Float)
	for bh, risk := range hashToRiskMap {
		hashToRevertRateMap[bh] = utils.CalcBlockConfirmationRisk(risk)
	}
	return hashToRevertRateMap, nil
}

// Close closes the client, aborting any in-flight requests.
func (client *Client) Close() {
	client.rpcRequester.Close()
}

// ===new rpc===

// GetAdmin returns admin of the given contract
func (client *Client) GetAdmin(account types.Address, epoch *types.Epoch) (admin types.Address, err error) {
	err = client.wrappedCallRPC(&admin, "cfx_getAdmin", account, epoch)
	return
}

// GetSponsorInfo returns sponsor information of the given contract
func (client *Client) GetSponsorInfo(contractAddress types.Address, epoch *types.Epoch) (sponsor types.SponsorInfo, err error) {
	sponsor = types.SponsorInfo{}
	err = client.wrappedCallRPC(&sponsor, "cfx_getSponsorInfo", contractAddress, epoch)
	return
}

// GetStakingBalance returns balance of the given account.
func (client *Client) GetStakingBalance(account types.Address, epoch *types.Epoch) (balance *hexutil.Big, err error) {
	balance = types.NewBigInt(0)
	err = client.wrappedCallRPC(balance, "cfx_getStakingBalance", account, epoch)
	return
}

// GetCollateralForStorage returns balance of the given account.
func (client *Client) GetCollateralForStorage(account types.Address, epoch *types.Epoch) (storage *hexutil.Big, err error) {
	storage = types.NewBigInt(0)
	err = client.wrappedCallRPC(storage, "cfx_getCollateralForStorage", account, epoch)
	return
}

// GetStorageAt returns storage entries from a given contract.
func (client *Client) GetStorageAt(address types.Address, position types.Hash, epoch *types.Epoch) (storageEntries *hexutil.Big, err error) {
	storageEntries = types.NewBigInt(0)
	err = client.wrappedCallRPC(storageEntries, "cfx_getStorageAt", address, position, epoch)
	return
}

// GetStorageRoot returns storage root of given address
func (client *Client) GetStorageRoot(address types.Address, epoch *types.Epoch) (storageRoot types.StorageRoot, err error) {
	err = client.wrappedCallRPC(&storageRoot, "cfx_getStorageRoot", address, epoch)
	return
}

// GetBlockByHashWithPivotAssumption returns block with given hash and pivot chain assumption.
func (client *Client) GetBlockByHashWithPivotAssumption(blockHash types.Hash, pivotHash types.Hash, epoch hexutil.Uint64) (block types.Block, err error) {
	err = client.wrappedCallRPC(&block, "cfx_getBlockByHashWithPivotAssumption", blockHash, pivotHash, epoch)
	return
}

// CheckBalanceAgainstTransaction checks if user balance is enough for the transaction.
func (client *Client) CheckBalanceAgainstTransaction(accountAddress types.Address,
	contractAddress types.Address,
	gasLimit *hexutil.Big,
	gasPrice *hexutil.Big,
	storageLimit *hexutil.Big,
	epoch *types.Epoch) (response types.CheckBalanceAgainstTransactionResponse, err error) {
	err = client.wrappedCallRPC(&response,
		"cfx_checkBalanceAgainstTransaction", accountAddress, contractAddress,
		gasLimit, gasPrice, storageLimit, epoch)
	return
}

// GetSkippedBlocksByEpoch returns skipped block hashes of given epoch
func (client *Client) GetSkippedBlocksByEpoch(epoch *types.Epoch) (blockHashs []types.Hash, err error) {
	blockHashs = make([]types.Hash, 0)
	err = client.wrappedCallRPC(&blockHashs, "cfx_getSkippedBlocksByEpoch", epoch)
	return
}

// GetAccountInfo returns account related states of the given account
func (client *Client) GetAccountInfo(account types.Address, epoch *types.Epoch) (accountInfo types.AccountInfo, err error) {
	err = client.wrappedCallRPC(&accountInfo, "cfx_getAccount", account, epoch)
	return
}

// GetInterestRate returns interest rate of the given epoch
func (client *Client) GetInterestRate(epoch *types.Epoch) (intersetRate *hexutil.Big, err error) {
	intersetRate = types.NewBigInt(0)
	err = client.wrappedCallRPC(intersetRate, "cfx_getInterestRate", epoch)
	return
}

// GetAccumulateInterestRate returns accumulate interest rate of the given epoch
func (client *Client) GetAccumulateInterestRate(epoch *types.Epoch) (intersetRate *hexutil.Big, err error) {
	intersetRate = types.NewBigInt(0)
	err = client.wrappedCallRPC(intersetRate, "cfx_getAccumulateInterestRate", epoch)
	return
}

// GetBlockRewardInfo returns block reward information in an epoch
func (client *Client) GetBlockRewardInfo(epoch *types.Epoch) (rewardInfo []types.RewardInfo, err error) {
	err = client.wrappedCallRPC(&rewardInfo, "cfx_getBlockRewardInfo", epoch)
	return
}

// GetClientVersion returns the client version as a string
func (client *Client) GetClientVersion() (clientVersion string, err error) {
	err = client.wrappedCallRPC(&clientVersion, "cfx_clientVersion")
	return
}

// === helper methods ===

func (client *Client) wrappedCallRPC(result interface{}, method string, args ...interface{}) error {
	fmtedArgs := genRPCParams(args...)
	if err := client.CallRPC(result, method, fmtedArgs...); err != nil {
		msg := fmt.Sprintf("rpc request method: %v with args: %v error\n", method, fmtedArgs)
		return types.WrapError(err, msg)
	}
	// fmt.Printf("rpcRequesterCall method %+v with args %+v: %+v\n", method, fmtedArgs, result)
	return nil
}

func genRPCParams(args ...interface{}) []interface{} {
	params := []interface{}{}
	for i := range args {
		// fmt.Printf("args %v:%v\n", i, args[i])
		if !utils.IsNil(args[i]) {
			// fmt.Printf("args %v:%v is not nil\n", i, args[i])
			params = append(params, args[i])
		}
	}
	return params
}

func unmarshalRPCResult(result interface{}, v interface{}) error {
	encoded, err := json.Marshal(result)
	if err != nil {
		msg := fmt.Sprintf("json marshal %v error", result)
		return types.WrapError(err, msg)
	}

	if err = json.Unmarshal(encoded, v); err != nil {
		msg := fmt.Sprintf("json unmarshal 0x%x error", encoded)
		return types.WrapError(err, msg)
	}

	return nil
}
