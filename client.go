// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package sdk

import (
	"context"
	"sync/atomic"

	"math/big"
	"reflect"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/constants"
	"github.com/Conflux-Chain/go-conflux-sdk/middleware"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	sdkerrors "github.com/Conflux-Chain/go-conflux-sdk/types/errors"
	postypes "github.com/Conflux-Chain/go-conflux-sdk/types/pos"
	rpc "github.com/openweb3/go-rpc-provider"

	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
)

const errMsgApplyTxValues = "failed to apply default transaction values"

// Client represents a client to interact with Conflux blockchain.
type Client struct {
	AccountManager      AccountManagerOperator
	nodeURL             string
	rpcRequester        RpcRequester
	networkID           *uint32
	chainID             *uint32
	option              ClientOption
	callRpcHandler      middleware.CallRpcHandler
	batchCallRpcHandler middleware.BatchCallRpcHandler

	rpcPosClient    RpcPosClient
	rpcTxpoolClient RpcTxpoolClient
	rpcDebugClient  RpcDebugClient
}

// ClientOption for set keystore path and flags for retry and timeout
//
// The simplest way to set logger is to use the types.DefaultCallRpcLog and types.DefaultBatchCallRPCLog
type ClientOption struct {
	KeystorePath string
	// retry
	RetryCount    int
	RetryInterval time.Duration
	// timeout of request
	RequestTimeout time.Duration
}

// NewClient creates an instance of Client with specified conflux node url, it will creat account manager if option.KeystorePath not empty.
// 	client, err := sdk.NewClient("https://test.confluxrpc.com", sdk.ClientOption{
// 	    KeystorePath: "your keystore folder path",
// 		RetryCount	: 3,
// 	})
//	// query rpc
// 	epoch, err := client.GetEpochNumber()
// 	if err != nil {
// 		panic(err)
// 	}
//	// send transaction
// 	chainID, err := client.GetNetworkID()
// 	if err!=nil {
// 	    panic(err)
// 	}
// 	from, err :=client.AccountManger().GetDefault()
// 	if err!=nil {
// 	    panic(err)
// 	}
// 	utx, err := client.CreateUnsignedTransaction(*from, cfxaddress.MustNewFromHex("0x1cad0b19bb29d4674531d6f115237e16afce377d", chainID), types.NewBigInt(1), nil)
// 	if err!=nil {
// 	    panic(err)
// 	}
// 	txhash, err := client.SendTransaction(utx)
func NewClient(nodeURL string, option ...ClientOption) (*Client, error) {
	realOption := ClientOption{}
	if len(option) > 0 {
		realOption = option[0]
	}

	client, err := newClientWithRetry(nodeURL, realOption)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new client with retry")
	}

	return client, nil
}

// MustNewClient same to NewClient but panic if failed
func MustNewClient(nodeURL string, option ...ClientOption) *Client {
	client, err := NewClient(nodeURL, option...)
	if err != nil {
		panic(err)
	}
	return client
}

// NewClientWithRPCRequester creates client with specified rpcRequester
func NewClientWithRPCRequester(rpcRequester RpcRequester) (*Client, error) {
	return &Client{
		rpcRequester: rpcRequester,
	}, nil
}

// NewClientWithRetry creates a retryable new instance of Client with specified conflux node url and retry options.
//
// the clientOption.RetryInterval will be set to 1 second if pass 0
func newClientWithRetry(nodeURL string, clientOption ClientOption) (*Client, error) {

	var client Client
	client.nodeURL = nodeURL
	client.option = clientOption
	client.callRpcHandler = middleware.CallRpcHandlerFunc(client.callRpc)
	client.batchCallRpcHandler = middleware.BatchCallRpcHandlerFunc(client.batchCallRPC)
	client.rpcPosClient = RpcPosClient{&client}
	client.rpcTxpoolClient = RpcTxpoolClient{&client}
	client.rpcDebugClient = RpcDebugClient{&client}
	client.option.setDefault()

	rpcClient, err := rpc.Dial(nodeURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial to fullnode")
	}

	if client.option.RetryCount == 0 {
		client.rpcRequester = rpcClient
	} else {
		client.rpcRequester = &rpcClientWithRetry{
			inner:      rpcClient,
			retryCount: client.option.RetryCount,
			interval:   client.option.RetryInterval,
		}
	}

	_, err = client.GetNetworkID()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get networkID")
	}

	if client.option.KeystorePath != "" {
		am := NewAccountManager(client.option.KeystorePath, *client.networkID)
		client.SetAccountManager(am)
	}

	_, err = client.GetChainID()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chainID")
	}

	return &client, nil
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

// Pos returns RpcPosClient for invoke rpc with pos namespace
func (client *Client) Pos() RpcPos {
	return &client.rpcPosClient
}

// TxPool returns RpcTxPoolClient for invoke rpc with txpool namespace
func (client *Client) TxPool() RpcTxpool {
	return &client.rpcTxpoolClient
}

// Debug returns RpcDebugClient for invoke rpc with debug namespace
func (client *Client) Debug() RpcDebug {
	return &client.rpcDebugClient
}

// GetNodeURL returns node url
func (client *Client) GetNodeURL() string {
	return client.nodeURL
}

// GetAccountManager returns account manager of client
func (client *Client) GetAccountManager() AccountManagerOperator {
	return client.AccountManager
}

// NewAddress create conflux address by base32 string or hex40 string, if base32OrHex is base32 and networkID is passed it will create cfx Address use networkID of current client.
func (client *Client) NewAddress(base32OrHex string) (types.Address, error) {
	networkID, err := client.GetNetworkID()
	if err != nil {
		return types.Address{}, err
	}
	return cfxaddress.New(base32OrHex, networkID)
}

// MustNewAddress create conflux address by base32 string or hex40 string, if base32OrHex is base32 and networkID is passed it will create cfx Address use networkID of current client.
// it will painc if error occured.
func (client *Client) MustNewAddress(base32OrHex string) types.Address {
	address, err := client.NewAddress(base32OrHex)
	if err != nil {
		panic(err)
	}
	return address
}

// CallRPC performs a JSON-RPC call with the given arguments and unmarshals into
// result if no error occurred.
//
// The result must be a pointer so that package json can unmarshal into it. You
// can also pass nil, in which case the result is ignored.
//
// You could use UseCallRpcMiddleware to add middleware for hooking CallRPC
func (client *Client) CallRPC(result interface{}, method string, args ...interface{}) error {
	return client.callRpcHandler.Handle(result, method, args...)
}

func (client *Client) callRpc(result interface{}, method string, args ...interface{}) error {
	ctx, cancelFunc := client.genContext()
	if cancelFunc != nil {
		defer cancelFunc()
	}
	return client.rpcRequester.CallContext(ctx, result, method, args...)
}

// UseCallRpcMiddleware set middleware to hook CallRpc, for example use middleware.CallRpcLogMiddleware for logging request info.
// You can customize your CallRpcMiddleware and use multi CallRpcMiddleware.
func (client *Client) UseCallRpcMiddleware(middleware middleware.CallRpcMiddleware) {
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
func (client *Client) BatchCallRPC(b []rpc.BatchElem) error {
	return client.batchCallRpcHandler.Handle(b)
}

func (client *Client) batchCallRPC(b []rpc.BatchElem) error {
	ctx, cancelFunc := client.genContext()
	if cancelFunc != nil {
		defer cancelFunc()
	}

	err := client.rpcRequester.BatchCallContext(ctx, b)
	if err != nil {
		return err
	}

	for i := range b {
		if rpcErr, err2 := utils.ToRpcError(b[i].Error); err2 == nil {
			b[i].Error = rpcErr
		}
	}
	return nil
}

// UseBatchCallRpcMiddleware set middleware to hook BatchCallRpc, for example use middleware.BatchCallRpcLogMiddleware for logging batch request info.
// You can customize your BatchCallRpcMiddleware and use multi BatchCallRpcMiddleware.
func (client *Client) UseBatchCallRpcMiddleware(middleware middleware.BatchCallRpcMiddleware) {
	client.batchCallRpcHandler = middleware(client.batchCallRpcHandler)
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

// GetNetworkID returns networkID of connecting conflux node
func (client *Client) GetNetworkID() (uint32, error) {
	if client.networkID != nil {
		return *client.networkID, nil
	}

	status, err := client.GetStatus()
	if err != nil {
		return 0, errors.Wrap(err, "failed to get status")
	}

	client.networkID = new(uint32)
	atomic.StoreUint32(client.networkID, uint32(status.NetworkID))
	return *client.networkID, nil
}

// GetNetworkIDCached returns chached networkID created when new client
func (client *Client) GetNetworkIDCached() uint32 {
	return *client.networkID
}

// GetNetworkID returns networkID of connecting conflux node
func (client *Client) GetChainID() (uint32, error) {
	if client.chainID != nil {
		return *client.chainID, nil
	}

	status, err := client.GetStatus()
	if err != nil {
		return 0, errors.Wrap(err, "failed to get status")
	}

	client.chainID = new(uint32)
	atomic.StoreUint32(client.chainID, uint32(status.ChainID))
	return *client.chainID, nil
}

// GetChainIDCached returns chached networkID created when new client
func (client *Client) GetChainIDCached() uint32 {
	return *client.chainID
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
func (client *Client) GetCode(address types.Address, epoch ...*types.Epoch) (code hexutil.Bytes, err error) {
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

// GetBlockByHash returns the block of specified block number
func (client *Client) GetBlockByBlockNumber(blockNumer hexutil.Uint64) (block *types.Block, err error) {
	err = client.wrappedCallRPC(&block, "cfx_getBlockByBlockNumber", blockNumer, true)
	return
}

// GetBlockSummaryByBlockNumber returns the block summary of specified block number.
func (client *Client) GetBlockSummaryByBlockNumber(blockNumer hexutil.Uint64) (block *types.BlockSummary, err error) {
	err = client.wrappedCallRPC(&block, "cfx_getBlockByBlockNumber", blockNumer, false)
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
func (client *Client) SendTransaction(tx types.UnsignedTransaction) (types.Hash, error) {

	err := client.ApplyUnsignedTransactionDefault(&tx)
	if err != nil {
		return "", errors.Wrap(err, errMsgApplyTxValues)
	}

	//sign
	if client.AccountManager == nil {
		return "", errors.New("account manager not specified, see SetAccountManager")
	}

	rawData, err := client.AccountManager.SignTransaction(tx)
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
func (client *Client) SendRawTransaction(rawData []byte) (hash types.Hash, err error) {
	tx := types.SignedTransaction{}
	if e := tx.Decode(rawData, client.GetChainIDCached()); e != nil {
		return "", errors.Wrap(e, "invalid raw transaction")
	}
	if tx.UnsignedTransaction.To != nil && tx.UnsignedTransaction.To.GetAddressType() == cfxaddress.AddressTypeUnknown {
		return "", errors.New("to address with unknown type is not allowed ")
	}

	err = client.wrappedCallRPC(&hash, "cfx_sendRawTransaction", hexutil.Encode(rawData))
	return
}

// SignEncodedTransactionAndSend signs RLP encoded transaction "encodedTx" by signature "r,s,v" and sends it to node,
// and returns responsed transaction.
func (client *Client) SignEncodedTransactionAndSend(encodedTx []byte, v byte, r, s []byte) (*types.Transaction, error) {
	tx := new(types.UnsignedTransaction)
	netwrokID, err := client.GetNetworkID()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get networkID")
	}

	err = tx.Decode(encodedTx, netwrokID)
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
	if err == nil {
		return
	}

	if rpcErr, err2 := utils.ToRpcError(err); err2 == nil {
		return result, rpcErr
	}

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
func (client *Client) GetStorageAt(address types.Address, position *hexutil.Big, epoch ...*types.Epoch) (storageEntries hexutil.Bytes, err error) {
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
func (client *Client) GetBlockTraces(blockHash types.Hash) (traces *types.LocalizedBlockTrace, err error) {
	err = client.wrappedCallRPC(&traces, "trace_block", blockHash)
	return
}

// GetFilterTraces returns all traces matching the provided filter.
func (client *Client) FilterTraces(traceFilter types.TraceFilter) (traces []types.LocalizedTrace, err error) {
	err = client.wrappedCallRPC(&traces, "trace_filter", traceFilter)
	return
}

// GetTransactionTraces returns all traces produced at the given transaction.
func (client *Client) GetTransactionTraces(txHash types.Hash) (traces []types.LocalizedTrace, err error) {
	err = client.wrappedCallRPC(&traces, "trace_transaction", txHash)
	return
}

// GetPosRewardByEpoch returns pos rewarded in this epoch
func (client *Client) GetPoSRewardByEpoch(epoch types.Epoch) (reward *postypes.EpochReward, err error) {
	err = client.CallRPC(&reward, "cfx_getPoSRewardByEpoch", epoch)
	return
}

// CreateUnsignedTransaction creates an unsigned transaction by parameters,
// and the other fields will be set to values fetched from conflux node.
func (client *Client) CreateUnsignedTransaction(from types.Address, to types.Address, amount *hexutil.Big, data []byte) (types.UnsignedTransaction, error) {
	tx := new(types.UnsignedTransaction)
	tx.From = &from
	tx.To = &to
	tx.Value = amount
	tx.Data = data

	err := client.ApplyUnsignedTransactionDefault(tx)
	if err != nil {
		return types.UnsignedTransaction{}, errors.Wrap(err, errMsgApplyTxValues)
	}

	return *tx, nil
}

// ApplyUnsignedTransactionDefault set empty fields to value fetched from conflux node.
func (client *Client) ApplyUnsignedTransactionDefault(tx *types.UnsignedTransaction) error {

	networkID, err := client.GetNetworkID()
	if err != nil {
		return errors.Wrap(err, "failed to get networkID")
	}

	chainID, err := client.GetChainID()
	if err != nil {
		return errors.Wrap(err, "failed to get chainID")
	}

	if client != nil {
		if tx.From == nil {
			//TODO: return error if client.AccountManager is nil?
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
		tx.From.CompleteByNetworkID(networkID)
		tx.To.CompleteByNetworkID(networkID)

		if tx.Nonce == nil {
			nonce, err := client.GetNextUsableNonce(*tx.From)
			if err != nil {
				return errors.Wrap(err, "failed to get nonce")
			}
			tmp := hexutil.Big(*nonce)
			tx.Nonce = &tmp
		}

		if tx.ChainID == nil {
			chainID := hexutil.Uint(chainID)
			tx.ChainID = &chainID
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
		txhash, err := client.SendTransaction(*tx)
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
				txReceipt, err := client.GetTransactionReceipt(txhash)
				if err != nil {
					result.Error = errors.Wrapf(err, "failed to get transaction receipt by hash %v", txhash)
					return
				}

				if txReceipt == nil {
					continue
				}

				if txReceipt.OutcomeStatus == 1 {
					result.Error = errors.Errorf("transaction execution failed, reason %v, hash = %v", txReceipt.TxExecErrorMsg, txhash)
					return
				}

				result.DeployedContract = &Contract{abi, client, txReceipt.ContractCreated}
				return
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

// GetAccountPendingInfo gets transaction pending info by account address
func (client *Client) GetAccountPendingInfo(address types.Address) (pendignInfo *types.AccountPendingInfo, err error) {
	err = client.wrappedCallRPC(&pendignInfo, "cfx_getAccountPendingInfo", address)
	return
}

// GetAccountPendingTransactions get transaction pending info by account address
func (client *Client) GetAccountPendingTransactions(address types.Address, startNonce *hexutil.Big, limit *hexutil.Uint64) (pendingTxs types.AccountPendingTransactions, err error) {
	err = client.wrappedCallRPC(&pendingTxs, "cfx_getAccountPendingTransactions", address, startNonce, limit)
	return
}

// GetPoSEconomics returns accumulate interest rate of the given epoch
func (client *Client) GetPoSEconomics(epoch ...*types.Epoch) (posEconomics types.PoSEconomics, err error) {
	err = client.wrappedCallRPC(&posEconomics, "cfx_getPoSEconomics", get1stEpochIfy(epoch))
	return
}

// GetOpenedMethodGroups returns openning method groups
func (client *Client) GetOpenedMethodGroups() (openedGroups []string, err error) {
	err = client.wrappedCallRPC(&openedGroups, "cfx_openedMethodGroups")
	return
}

// =====Debug RPC=====

func (client *Client) GetEpochReceipts(epoch types.Epoch) (receipts [][]types.TransactionReceipt, err error) {
	return client.Debug().GetEpochReceipts(epoch)
}

func (client *Client) GetEpochReceiptsByPivotBlockHash(hash types.Hash) (receipts [][]types.TransactionReceipt, err error) {
	return client.Debug().GetEpochReceiptsByPivotBlockHash(hash)
}

// =======Batch=======

// BatchGetTxByHashes requests transaction informations in bulk by txhashes
func (client *Client) BatchGetTxByHashes(txhashes []types.Hash) (map[types.Hash]*types.Transaction, error) {
	if len(txhashes) == 0 {
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

	if len(blockhashes) == 0 {
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
		if reflect.DeepEqual(be.Result, &types.BlockSummary{}) {
			delete(hashToBlocksummaryMap, bh)
			continue
		}
		hashToBlocksummaryMap[bh] = be.Result.(*types.BlockSummary)
	}
	return hashToBlocksummaryMap, nil
}

// BatchGetBlockSummarysByNumber requests block summary informations in bulk by blocknumbers
func (client *Client) BatchGetBlockSummarysByNumber(blocknumbers []hexutil.Uint64) (map[hexutil.Uint64]*types.BlockSummary, error) {

	if len(blocknumbers) == 0 {
		return make(map[hexutil.Uint64]*types.BlockSummary), nil
	}

	cache := make(map[hexutil.Uint64]*rpc.BatchElem)

	for i := range blocknumbers {
		if cache[blocknumbers[i]] == nil {
			be := rpc.BatchElem{
				Method: "cfx_getBlockByBlockNumber",
				Args:   []interface{}{blocknumbers[i], false},
				Result: &types.BlockSummary{},
			}
			cache[blocknumbers[i]] = &be
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

	numberToBlocksummaryMap := make(map[hexutil.Uint64]*types.BlockSummary)

	for _, bn := range blocknumbers {
		be := cache[bn]
		if reflect.DeepEqual(be.Result, &types.BlockSummary{}) {
			delete(numberToBlocksummaryMap, bn)
			continue
		}
		numberToBlocksummaryMap[bn] = be.Result.(*types.BlockSummary)
	}
	return numberToBlocksummaryMap, nil
}

// BatchGetRawBlockConfirmationRisk requests raw confirmation risk informations in bulk by blockhashes
func (client *Client) BatchGetRawBlockConfirmationRisk(blockhashes []types.Hash) (map[types.Hash]*big.Int, error) {

	if len(blockhashes) == 0 {
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

// SubscribeEpochs subscribes consensus results: the total order of blocks, as expressed by a sequence of epochs. Currently subscriptionEpochType only support "latest_mined" and "latest_state"
func (client *Client) SubscribeEpochs(channel chan types.WebsocketEpochResponse, subscriptionEpochType ...types.Epoch) (*rpc.ClientSubscription, error) {
	if len(subscriptionEpochType) > 0 {
		return client.rpcRequester.Subscribe(context.Background(), "cfx", channel, "epochs", subscriptionEpochType[0].String())
	}
	return client.rpcRequester.Subscribe(context.Background(), "cfx", channel, "epochs")
}

// SubscribeLogs subscribes all logs matching a certain filter, in order.
func (client *Client) SubscribeLogs(channel chan types.SubscriptionLog, filter types.LogFilter) (*rpc.ClientSubscription, error) {
	return client.rpcRequester.Subscribe(context.Background(), "cfx", channel, "logs", filter)
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
	timeout := time.Duration(24 * time.Hour)
	pass := time.Duration(0)

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

		pass += duration
		if pass > timeout {
			return nil, sdkerrors.ErrTimeout
		}
	}
	return txReceipt, nil
}

func (client *Client) GetNextUsableNonce(user types.Address) (nonce *hexutil.Big, err error) {
	hexNonce, err := client.TxPool().NextNonce(user)
	if err != nil {
		hexNonce, err = client.GetNextNonce(user)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return hexNonce, nil
}

// ======== private methods=============

func (client *Client) wrappedCallRPC(result interface{}, method string, args ...interface{}) error {
	fmtedArgs, err := client.genRPCParams(args...)
	if err != nil {
		return errors.WithStack(err)
	}
	return client.CallRPC(result, method, fmtedArgs...)
}

func (client *Client) genRPCParams(args ...interface{}) ([]interface{}, error) {
	// fmt.Println("gen rpc params")
	params := []interface{}{}
	for i := range args {
		// fmt.Printf("args %v:%v\n", i, args[i])
		if !utils.IsNil(args[i]) {
			// fmt.Printf("args %v:%v is not nil\n", i, args[i])

			networkID, err := client.GetNetworkID()
			if err != nil {
				return nil, errors.Wrap(err, "failed to get networkID")
			}

			if tmp, ok := args[i].(cfxaddress.Address); ok {
				tmp.CompleteByNetworkID(networkID)
				args[i] = tmp
				// fmt.Printf("complete by networkID,%v; after %v\n", client.networkID, args[i])
			}

			if tmp, ok := args[i].(*cfxaddress.Address); ok {
				tmp.CompleteByNetworkID(networkID)
				// fmt.Printf("complete by networkID,%v; after %v\n", client.networkID, args[i])
			}

			if tmp, ok := args[i].(types.CallRequest); ok {
				tmp.From.CompleteByNetworkID(networkID)
				tmp.To.CompleteByNetworkID(networkID)
				args[i] = tmp
			}

			if tmp, ok := args[i].(*types.CallRequest); ok {
				tmp.From.CompleteByNetworkID(networkID)
				tmp.To.CompleteByNetworkID(networkID)
			}

			params = append(params, args[i])
		}
	}
	return params, nil
}

func get1stEpochIfy(epoch []*types.Epoch) *types.Epoch {
	var realEpoch *types.Epoch
	if len(epoch) > 0 {
		realEpoch = epoch[0]
	}
	return realEpoch
}

func get1stBoolIfy(values []bool) bool {
	value := false
	if len(values) > 0 {
		value = values[0]
	}
	return value
}

func get1stU64Ify(values []hexutil.Uint64) *hexutil.Uint64 {
	if len(values) > 0 {
		_value := values[0]
		return &_value
	}
	return nil
}

func (client *Client) genContext() (context.Context, context.CancelFunc) {
	if client.option.RequestTimeout > 0 {
		return context.WithTimeout(context.Background(), client.option.RequestTimeout)
	}
	return context.Background(), nil
}
