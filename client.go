// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package sdk

import (
	"context"
	"encoding/json"
	"math/big"
	"reflect"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/constants"
	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
)

const errMsgApplyTxValues = "failed to apply default transaction values"

// Client represents a client to interact with Conflux blockchain.
type Client struct {
	AccountManager AccountManagerOperator
	nodeURL        string
	rpcRequester   rpcRequester
	chainID        uint32
}

type ClientOption struct {
	KeystorePath  string
	RetryCount    int
	RetryInterval time.Duration
}

// NewClient creates a new instance of Client with specified conflux node url.
func NewClient(nodeURL string, option ...ClientOption) (*Client, error) {
	realOption := ClientOption{}
	if len(option) > 0 {
		realOption = option[0]
	}

	client, err := newClientWithRetry(nodeURL, realOption.KeystorePath, realOption.RetryCount, realOption.RetryInterval)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new client with retry")
	}

	return client, nil
}

// NewClientWithRPCRequester creates client with specified rpcRequester
func NewClientWithRPCRequester(rpcRequester rpcRequester) (*Client, error) {
	return &Client{
		rpcRequester: rpcRequester,
	}, nil
}

// NewClientWithRetry creates a retryable new instance of Client with specified conflux node url and retry options.
//
// the retryInterval will be set to 1 second if pass 0
func newClientWithRetry(nodeURL string, keystorePath string, retryCount int, retryInterval time.Duration) (*Client, error) {

	var client Client
	client.nodeURL = nodeURL

	rpcClient, err := rpc.Dial(nodeURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial to fullnode")
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

	client.chainID, err = client.GetChainID()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chainID")
	}

	if keystorePath != "" {
		am := NewAccountManager(keystorePath, client.chainID)
		client.SetAccountManager(am)
	}

	return &client, nil
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
	client.AccountManager = accountManager
}

// GetGasPrice returns the recent mean gas price.
func (client *Client) GetGasPrice() (gasPrice *hexutil.Big, err error) {
	err = client.wrappedCallRPC(&gasPrice, "cfx_gasPrice")
	return
}

// GetNextNonce returns the next transaction nonce of address
func (client *Client) GetNextNonce(address types.Address, epoch ...*types.Epoch) (nonce *hexutil.Big, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.wrappedCallRPC(&nonce, "cfx_getNextNonce", address, realEpoch)
	return
}

// GetStatus returns status of connecting conflux node
func (client *Client) GetStatus() (status types.Status, err error) {
	err = client.wrappedCallRPC(&status, "cfx_getStatus")
	return
}

// GetChainID returns chainID of connecting conflux node
func (client *Client) GetChainID() (uint32, error) {
	status, err := client.GetStatus()
	if err != nil {
		return 0, errors.Wrap(err, "failed to get status")
	}
	return uint32(*status.ChainID), nil
}

// GetEpochNumber returns the highest or specified epoch number.
func (client *Client) GetEpochNumber(epoch ...*types.Epoch) (epochNumber *hexutil.Big, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.wrappedCallRPC(&epochNumber, "cfx_epochNumber", realEpoch)
	if err != nil {
		epochNumber = nil
	}
	return
}

// GetBalance returns the balance of specified address at epoch.
func (client *Client) GetBalance(address types.Address, epoch ...*types.Epoch) (balance *hexutil.Big, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.wrappedCallRPC(&balance, "cfx_getBalance", address, realEpoch)
	if err != nil {
		balance = nil
	}
	return
}

// GetCode returns the bytecode in HEX format of specified address at epoch.
func (client *Client) GetCode(address types.Address, epoch ...*types.Epoch) (code string, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.wrappedCallRPC(&code, "cfx_getCode", address, realEpoch)
	return
}

// GetBlockSummaryByHash returns the block summary of specified blockHash
// If the block is not found, return nil.
func (client *Client) GetBlockSummaryByHash(blockHash types.Hash) (blockSummary *types.BlockSummary, err error) {
	err = client.wrappedCallRPC(&blockSummary, "cfx_getBlockByHash", blockHash, false)
	return
}

// GetBlockByHash returns the block of specified blockHash
// If the block is not found, return nil.
func (client *Client) GetBlockByHash(blockHash types.Hash) (block *types.Block, err error) {
	err = client.wrappedCallRPC(&block, "cfx_getBlockByHash", blockHash, true)
	return
}

// GetBlockSummaryByEpoch returns the block summary of specified epoch.
// If the epoch is invalid, return the concrete error.
func (client *Client) GetBlockSummaryByEpoch(epoch *types.Epoch) (blockSummary *types.BlockSummary, err error) {
	err = client.wrappedCallRPC(&blockSummary, "cfx_getBlockByEpochNumber", epoch, false)
	return
}

// GetBlockByEpoch returns the block of specified epoch.
// If the epoch is invalid, return the concrete error.
func (client *Client) GetBlockByEpoch(epoch *types.Epoch) (block *types.Block, err error) {
	err = client.wrappedCallRPC(&block, "cfx_getBlockByEpochNumber", epoch, true)
	return
}

// GetBestBlockHash returns the current best block hash.
func (client *Client) GetBestBlockHash() (hash types.Hash, err error) {
	err = client.wrappedCallRPC(&hash, "cfx_getBestBlockHash")
	return
}

// GetRawBlockConfirmationRisk indicates the risk coefficient that
// the pivot block of the epoch where the block is located becomes a normal block.
// It will return nil if block not exist
func (client *Client) GetRawBlockConfirmationRisk(blockhash types.Hash) (risk *hexutil.Big, err error) {
	err = client.wrappedCallRPC(&risk, "cfx_getConfirmationRiskByHash", blockhash)
	return
}

// GetBlockConfirmationRisk indicates the probability that
// the pivot block of the epoch where the block is located becomes a normal block.
//
// it's (raw confirmation risk coefficient/ (2^256-1))
func (client *Client) GetBlockConfirmationRisk(blockHash types.Hash) (*big.Float, error) {
	risk, err := client.GetRawBlockConfirmationRisk(blockHash)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to cfx_getConfirmationRiskByHash %v", blockHash)
	}
	if risk == nil {
		return nil, nil
	}

	riskFloat := new(big.Float).SetInt(risk.ToInt())
	maxUint256Float := new(big.Float).SetInt(constants.MaxUint256)

	riskRate := new(big.Float).Quo(riskFloat, maxUint256Float)
	return riskRate, nil
}

// SendTransaction signs and sends transaction to conflux node and returns the transaction hash.
func (client *Client) SendTransaction(tx *types.UnsignedTransaction) (types.Hash, error) {

	err := client.ApplyUnsignedTransactionDefault(tx)
	if err != nil {
		return "", errors.Wrap(err, errMsgApplyTxValues)
	}
	// fmt.Printf("tx info: %+v", tx)
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
	if client.AccountManager == nil {
		return "", errors.New("account manager not specified, see SetAccountManager")
	}

	rawData, err := client.AccountManager.SignTransaction(*tx)
	if err != nil {
		return "", errors.Wrap(err, "failed to sign transaction")
	}

	//send raw tx
	txhash, err := client.SendRawTransaction(rawData)
	if err != nil {
		return "", errors.Wrapf(err, "failed to send transaction, raw data = 0x%+x", rawData)
	}
	return txhash, nil
}

// SendRawTransaction sends signed transaction and returns its hash.
func (client *Client) SendRawTransaction(rawData []byte) (types.Hash, error) {
	var result interface{}

	if err := client.rpcRequester.Call(&result, "cfx_sendRawTransaction", hexutil.Encode(rawData)); err != nil {
		return "", err
	}

	return types.Hash(result.(string)), nil
}

// SignEncodedTransactionAndSend signs RLP encoded transaction "encodedTx" by signature "r,s,v" and sends it to node,
// and returns responsed transaction.
func (client *Client) SignEncodedTransactionAndSend(encodedTx []byte, v byte, r, s []byte) (*types.Transaction, error) {
	tx := new(types.UnsignedTransaction)
	err := tx.Decode(encodedTx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode transaction")
	}
	// tx.From = from

	respondTx, err := client.signTransactionAndSend(tx, v, r, s)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to sign and send transaction %+v", tx)
	}

	return respondTx, nil
}

func (client *Client) signTransactionAndSend(tx *types.UnsignedTransaction, v byte, r, s []byte) (*types.Transaction, error) {
	rlp, err := tx.EncodeWithSignature(v, r, s)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode transaction with signature")
	}

	hash, err := client.SendRawTransaction(rlp)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to send transaction, raw data = 0x%+x", rlp)
	}

	respondTx, err := client.GetTransactionByHash(hash)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get transaction by hash %v", hash)
	}
	return respondTx, nil
}

// Call executes a message call transaction "request" at specified epoch,
// which is directly executed in the VM of the node, but never mined into the block chain
// and returns the contract execution result.
func (client *Client) Call(request types.CallRequest, epoch *types.Epoch) (result hexutil.Bytes, err error) {
	err = client.wrappedCallRPC(&result, "cfx_call", request, epoch)
	return
}

// GetLogs returns logs that matching the specified filter.
func (client *Client) GetLogs(filter types.LogFilter) (logs []types.Log, err error) {
	err = client.wrappedCallRPC(&logs, "cfx_getLogs", filter)
	return
}

// GetTransactionByHash returns transaction for the specified txHash.
// If the transaction is not found, return nil.
func (client *Client) GetTransactionByHash(txHash types.Hash) (tx *types.Transaction, err error) {
	err = client.wrappedCallRPC(&tx, "cfx_getTransactionByHash", txHash)
	return
}

// EstimateGasAndCollateral excutes a message call "request"
// and returns the amount of the gas used and storage for collateral
func (client *Client) EstimateGasAndCollateral(request types.CallRequest, epoch ...*types.Epoch) (estimat types.Estimate, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.wrappedCallRPC(&estimat, "cfx_estimateGasAndCollateral", request, realEpoch)
	return
}

// GetBlocksByEpoch returns the blocks hash in the specified epoch.
func (client *Client) GetBlocksByEpoch(epoch *types.Epoch) (blockHashes []types.Hash, err error) {
	err = client.wrappedCallRPC(&blockHashes, "cfx_getBlocksByEpoch", epoch)
	return
}

// GetTransactionReceipt returns the receipt of specified transaction hash.
// If no receipt is found, return nil.
func (client *Client) GetTransactionReceipt(txHash types.Hash) (receipt *types.TransactionReceipt, err error) {
	err = client.wrappedCallRPC(&receipt, "cfx_getTransactionReceipt", txHash)
	return
}

// ===new rpc===

// GetAdmin returns admin of the given contract, it will return nil if contract not exist
func (client *Client) GetAdmin(contractAddress types.Address, epoch ...*types.Epoch) (admin *types.Address, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.wrappedCallRPC(&admin, "cfx_getAdmin", contractAddress, realEpoch)
	return
}

// GetSponsorInfo returns sponsor information of the given contract
func (client *Client) GetSponsorInfo(contractAddress types.Address, epoch ...*types.Epoch) (sponsor types.SponsorInfo, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.wrappedCallRPC(&sponsor, "cfx_getSponsorInfo", contractAddress, realEpoch)
	return
}

// GetStakingBalance returns balance of the given account.
func (client *Client) GetStakingBalance(account types.Address, epoch ...*types.Epoch) (balance *hexutil.Big, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.wrappedCallRPC(&balance, "cfx_getStakingBalance", account, realEpoch)
	return
}

// GetCollateralForStorage returns balance of the given account.
func (client *Client) GetCollateralForStorage(account types.Address, epoch ...*types.Epoch) (storage *hexutil.Big, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.wrappedCallRPC(&storage, "cfx_getCollateralForStorage", account, realEpoch)
	return
}

// GetStorageAt returns storage entries from a given contract.
func (client *Client) GetStorageAt(address types.Address, position types.Hash, epoch ...*types.Epoch) (storageEntries hexutil.Bytes, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.wrappedCallRPC(&storageEntries, "cfx_getStorageAt", address, position, realEpoch)
	return
}

// GetStorageRoot returns storage root of given address
func (client *Client) GetStorageRoot(address types.Address, epoch ...*types.Epoch) (storageRoot *types.StorageRoot, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.wrappedCallRPC(&storageRoot, "cfx_getStorageRoot", address, realEpoch)
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
	epoch ...*types.Epoch) (response types.CheckBalanceAgainstTransactionResponse, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.wrappedCallRPC(&response,
		"cfx_checkBalanceAgainstTransaction", accountAddress, contractAddress,
		gasLimit, gasPrice, storageLimit, realEpoch)
	return
}

// GetSkippedBlocksByEpoch returns skipped block hashes of given epoch
func (client *Client) GetSkippedBlocksByEpoch(epoch *types.Epoch) (blockHashs []types.Hash, err error) {
	err = client.wrappedCallRPC(&blockHashs, "cfx_getSkippedBlocksByEpoch", epoch)
	return
}

// GetAccountInfo returns account related states of the given account
func (client *Client) GetAccountInfo(account types.Address, epoch ...*types.Epoch) (accountInfo types.AccountInfo, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.wrappedCallRPC(&accountInfo, "cfx_getAccount", account, realEpoch)
	return
}

// GetInterestRate returns interest rate of the given epoch
func (client *Client) GetInterestRate(epoch ...*types.Epoch) (intersetRate *hexutil.Big, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.wrappedCallRPC(&intersetRate, "cfx_getInterestRate", realEpoch)
	if err != nil {
		intersetRate = nil
	}
	return
}

// GetAccumulateInterestRate returns accumulate interest rate of the given epoch
func (client *Client) GetAccumulateInterestRate(epoch ...*types.Epoch) (intersetRate *hexutil.Big, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.wrappedCallRPC(&intersetRate, "cfx_getAccumulateInterestRate", realEpoch)
	if err != nil {
		intersetRate = nil
	}
	return
}

// GetBlockRewardInfo returns block reward information in an epoch
func (client *Client) GetBlockRewardInfo(epoch types.Epoch) (rewardInfo []types.RewardInfo, err error) {
	err = client.wrappedCallRPC(&rewardInfo, "cfx_getBlockRewardInfo", epoch)
	return
}

// GetClientVersion returns the client version as a string
func (client *Client) GetClientVersion() (clientVersion string, err error) {
	err = client.wrappedCallRPC(&clientVersion, "cfx_clientVersion")
	return
}

// GetDepositList returns deposit list of the given account.
func (client *Client) GetDepositList(address types.Address, epoch ...*types.Epoch) (depositInfos []types.DepositInfo, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.wrappedCallRPC(&depositInfos, "cfx_getDepositList", address, realEpoch)
	return
}

// GetVoteList returns vote list of the given account.
func (client *Client) GetVoteList(address types.Address, epoch ...*types.Epoch) (voteStakeInfos []types.VoteStakeInfo, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.wrappedCallRPC(&voteStakeInfos, "cfx_getVoteList", address, realEpoch)
	return
}

// GetSupplyInfo Return information about total token supply.
func (client *Client) GetSupplyInfo(epoch ...*types.Epoch) (info types.TokenSupplyInfo, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.wrappedCallRPC(&info, "cfx_getSupplyInfo", realEpoch)
	return
}

// GetBlockTrace returns all traces produced at given block.
func (client *Client) GetBlockTrace(blockHash types.Hash) (trace *types.LocalizedBlockTrace, err error) {
	err = client.wrappedCallRPC(&trace, "trace_block", blockHash)
	return
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
		return nil, errors.Wrap(err, errMsgApplyTxValues)
	}

	return tx, nil
}

// ApplyUnsignedTransactionDefault set empty fields to value fetched from conflux node.
func (client *Client) ApplyUnsignedTransactionDefault(tx *types.UnsignedTransaction) error {

	if client != nil {
		if tx.From == nil {
			if client.AccountManager != nil {
				defaultAccount, err := client.AccountManager.GetDefault()
				if err != nil {
					return errors.Wrap(err, "failed to get default account")
				}

				if defaultAccount == nil {
					return errors.New("no account found")
				}
				tx.From = defaultAccount
			}
		}
		tx.From.CompleteAddressByChainID(client.chainID)
		tx.To.CompleteAddressByChainID(client.chainID)

		if tx.Nonce == nil {
			nonce, err := client.GetNextNonce(*tx.From, nil)
			if err != nil {
				return errors.Wrap(err, "failed to get nonce")
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
				return errors.Wrap(err, "failed to get gas price")
			}

			// conflux responsed gasprice offen be 0, but the min gasprice is 1 when sending transaction, so do this
			if gasPrice.ToInt().Cmp(big.NewInt(constants.MinGasprice)) < 1 {
				gasPrice = types.NewBigInt(constants.MinGasprice)
			}
			tmp := hexutil.Big(*gasPrice)
			tx.GasPrice = &tmp
		}

		if tx.EpochHeight == nil {
			epoch, err := client.GetEpochNumber(types.EpochLatestState)
			if err != nil {
				return errors.Wrap(err, "failed to get the latest state epoch number")
			}
			// tx.EpochHeight = (*hexutil.Big)(epoch).toi
			tx.EpochHeight = types.NewUint64(epoch.ToInt().Uint64())
		}

		// The gas and storage limit may be influnced by all fileds of transaction ,so set them at last step.
		if tx.StorageLimit == nil || tx.Gas == nil {
			callReq := new(types.CallRequest)
			callReq.FillByUnsignedTx(tx)

			sm, err := client.EstimateGasAndCollateral(*callReq)
			if err != nil {
				return errors.Wrapf(err, "failed to estimate gas and collateral, request = %+v", *callReq)
			}

			// fmt.Printf("callreq, %+v,sm:%+v\n", *callReq, sm)

			if tx.Gas == nil {
				tx.Gas = sm.GasLimit
			}

			if tx.StorageLimit == nil {
				tx.StorageLimit = types.NewUint64(sm.StorageCollateralized.ToInt().Uint64() * 10 / 9)
			}
		}

		tx.ApplyDefault()
	}

	return nil
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
			result.Error = errors.Errorf("failed to unmarshal ABI: %+v", abiJSON)
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
				result.Error = errors.Wrapf(err, "failed to encode constructor with args %+v", constroctorParams)
				return
			}

			bytecode = append(bytecode, input...)
		}
		tx.Data = bytecode

		//deploy contract
		txhash, err := client.SendTransaction(tx)
		if err != nil {
			result.Error = errors.Wrapf(err, "failed to send transaction, tx = %+v", tx)
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
				result.Error = errors.Errorf("deploy contract timeout, time = %v, txhash = %v", t, txhash)
				return
			// Got a tick
			case <-ticker:
				transaction, err := client.GetTransactionByHash(txhash)
				if err != nil {
					result.Error = errors.Wrapf(err, "failed to get transaction by hash %v", txhash)
					return
				}

				if transaction.Status != nil {
					if *transaction.Status == 1 {
						result.Error = errors.Errorf("transaction execution failed, hash = %v", txhash)
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
		return nil, errors.Wrap(err, "failed unmarshal ABI")
	}

	contract := &Contract{abi, client, deployedAt}
	return contract, nil
}

// =======Batch=======

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
		if reflect.DeepEqual(be.Result, types.Transaction{}) {
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
		if reflect.DeepEqual(be.Result, types.Transaction{}) {
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

// === pub/sub ===

// SubscribeNewHeads subscribes all new block headers participating in the consensus.
func (client *Client) SubscribeNewHeads(channel chan types.BlockHeader) (*rpc.ClientSubscription, error) {
	return client.rpcRequester.Subscribe(context.Background(), "cfx", channel, "newHeads")
}

// SubscribeEpochs subscribes consensus results: the total order of blocks, as expressed by a sequence of epochs.
func (client *Client) SubscribeEpochs(channel chan types.WebsocketEpochResponse) (*rpc.ClientSubscription, error) {
	return client.rpcRequester.Subscribe(context.Background(), "cfx", channel, "epochs")
}

// SubscribeLogs subscribes all logs matching a certain filter, in order.
func (client *Client) SubscribeLogs(logChannel chan types.Log, chainReorgChannel chan types.ChainReorg, filter types.LogFilter) (*rpc.ClientSubscription, error) {
	channel := make(chan types.SubscriptionLog, 100)
	clientSubscrip, err := client.rpcRequester.Subscribe(context.Background(), "cfx", channel, "logs", filter)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			subscriptionLog, isOpen := <-channel
			if !isOpen {
				return
			}
			if subscriptionLog.IsRevertLog() {
				chainReorgChannel <- subscriptionLog.ChainReorg
			} else {
				logChannel <- subscriptionLog.Log
			}
		}
	}()
	return clientSubscrip, nil
}

// === helper methods ===

// WaitForTransationBePacked returns transaction when it is packed
func (client *Client) WaitForTransationBePacked(txhash types.Hash, duration time.Duration) (*types.Transaction, error) {
	// fmt.Printf("wait for transaction %v be packed\n", txhash)
	if duration == 0 {
		duration = time.Second
	}

	var tx *types.Transaction
	for {
		time.Sleep(duration)
		var err error
		tx, err = client.GetTransactionByHash(txhash)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get transaction by hash %v", txhash)
		}

		if tx.Status != nil {
			// fmt.Printf("transaction is packed:%+v\n\n", JsonFmt(tx))
			break
		}
	}
	return tx, nil
}

// WaitForTransationReceipt waits for transaction receipt valid
func (client *Client) WaitForTransationReceipt(txhash types.Hash, duration time.Duration) (*types.TransactionReceipt, error) {
	// fmt.Printf("wait for transaction %v be packed\n", txhash)
	if duration == 0 {
		duration = time.Second
	}

	var txReceipt *types.TransactionReceipt
	for {
		time.Sleep(duration)
		var err error
		txReceipt, err = client.GetTransactionReceipt(txhash)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get transaction receipt")
		}

		if txReceipt != nil {
			break
		}
	}
	return txReceipt, nil
}

func (client *Client) wrappedCallRPC(result interface{}, method string, args ...interface{}) error {
	fmtedArgs := client.genRPCParams(args...)
	return client.CallRPC(result, method, fmtedArgs...)
}

func (client *Client) genRPCParams(args ...interface{}) []interface{} {
	params := []interface{}{}
	for i := range args {
		// fmt.Printf("args %v:%v\n", i, args[i])
		if !utils.IsNil(args[i]) {
			// fmt.Printf("args %v:%v is not nil\n", i, args[i])
			t := reflect.TypeOf(args[i])
			// fmt.Printf("typeof %v is %v\n", args[i], t)
			if t == reflect.TypeOf(cfxaddress.Address{}) {
				tmp := args[i].(cfxaddress.Address)
				tmp.CompleteAddressByChainID(client.chainID)
				// fmt.Printf("complete by chainID,%v", client.chainID)
			}

			if t == reflect.TypeOf(&cfxaddress.Address{}) {
				tmp := args[i].(*cfxaddress.Address)
				tmp.CompleteAddressByChainID(client.chainID)
				// fmt.Printf("complete by chainID,%v", client.chainID)
			}

			if t == reflect.TypeOf(types.CallRequest{}) {
				tmp := args[i].(types.CallRequest)
				tmp.From.CompleteAddressByChainID(client.chainID)
				tmp.To.CompleteAddressByChainID(client.chainID)
			}

			if t == reflect.TypeOf(&types.CallRequest{}) {
				tmp := args[i].(*types.CallRequest)
				tmp.From.CompleteAddressByChainID(client.chainID)
				tmp.To.CompleteAddressByChainID(client.chainID)
			}

			params = append(params, args[i])
		}
	}
	return params
}

func unmarshalRPCResult(result interface{}, v interface{}) error {
	encoded, err := json.Marshal(result)
	if err != nil {
		return errors.Wrap(err, "failed to marshal RPC result to JSON")
	}

	if err = json.Unmarshal(encoded, v); err != nil {
		return errors.Wrap(err, "failed to unmarshal JSON to RPC result")
	}

	return nil
}

func get1stEpochIfy(epoch []*types.Epoch) *types.Epoch {
	var realEpoch *types.Epoch
	if len(epoch) > 0 {
		realEpoch = epoch[0]
	}
	return realEpoch
}
