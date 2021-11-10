package client

import (
	"math/big"

	"github.com/Conflux-Chain/go-conflux-sdk/constants"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
)

type RpcCfxClient struct {
	core *ClientCore
}

func NewRpcCfxClient(core *ClientCore) RpcCfxClient {
	return RpcCfxClient{core}
}

// GetGasPrice returns the recent mean gas price.
func (client *RpcCfxClient) GetGasPrice() (gasPrice *hexutil.Big, err error) {
	err = client.core.wrappedCallRPC(&gasPrice, "cfx_gasPrice")
	return
}

// GetNextNonce returns the next transaction nonce of address
func (client *RpcCfxClient) GetNextNonce(address types.Address, epoch ...*types.Epoch) (nonce *hexutil.Big, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.core.wrappedCallRPC(&nonce, "cfx_getNextNonce", address, realEpoch)
	return
}

// GetStatus returns status of connecting conflux node
func (client *RpcCfxClient) GetStatus() (status types.Status, err error) {
	return client.core.getStatus()
}

// GetEpochNumber returns the highest or specified epoch number.
func (client *RpcCfxClient) GetEpochNumber(epoch ...*types.Epoch) (epochNumber *hexutil.Big, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.core.wrappedCallRPC(&epochNumber, "cfx_epochNumber", realEpoch)
	if err != nil {
		epochNumber = nil
	}
	return
}

// GetBalance returns the balance of specified address at epoch.
func (client *RpcCfxClient) GetBalance(address types.Address, epoch ...*types.Epoch) (balance *hexutil.Big, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.core.wrappedCallRPC(&balance, "cfx_getBalance", address, realEpoch)
	if err != nil {
		balance = nil
	}
	return
}

// GetCode returns the bytecode in HEX format of specified address at epoch.
func (client *RpcCfxClient) GetCode(address types.Address, epoch ...*types.Epoch) (code hexutil.Bytes, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.core.wrappedCallRPC(&code, "cfx_getCode", address, realEpoch)
	return
}

// GetBlockSummaryByHash returns the block summary of specified blockHash
// If the block is not found, return nil.
func (client *RpcCfxClient) GetBlockSummaryByHash(blockHash types.Hash) (blockSummary *types.BlockSummary, err error) {
	err = client.core.wrappedCallRPC(&blockSummary, "cfx_getBlockByHash", blockHash, false)
	return
}

// GetBlockByHash returns the block of specified blockHash
// If the block is not found, return nil.
func (client *RpcCfxClient) GetBlockByHash(blockHash types.Hash) (block *types.Block, err error) {
	err = client.core.wrappedCallRPC(&block, "cfx_getBlockByHash", blockHash, true)
	return
}

// GetBlockSummaryByEpoch returns the block summary of specified epoch.
// If the epoch is invalid, return the concrete error.
func (client *RpcCfxClient) GetBlockSummaryByEpoch(epoch *types.Epoch) (blockSummary *types.BlockSummary, err error) {
	err = client.core.wrappedCallRPC(&blockSummary, "cfx_getBlockByEpochNumber", epoch, false)
	return
}

func (client *RpcCfxClient) GetBlockByBlockNumber(blockNumer hexutil.Uint64) (block *types.Block, err error) {
	err = client.core.wrappedCallRPC(&block, "cfx_getBlockByBlockNumber", blockNumer, true)
	return
}

func (client *RpcCfxClient) GetBlockSummaryByBlockNumber(blockNumer hexutil.Uint64) (block *types.BlockSummary, err error) {
	err = client.core.wrappedCallRPC(&block, "cfx_getBlockByBlockNumber", blockNumer, false)
	return
}

// GetBlockByEpoch returns the block of specified epoch.
// If the epoch is invalid, return the concrete error.
func (client *RpcCfxClient) GetBlockByEpoch(epoch *types.Epoch) (block *types.Block, err error) {
	err = client.core.wrappedCallRPC(&block, "cfx_getBlockByEpochNumber", epoch, true)
	return
}

// GetBestBlockHash returns the current best block hash.
func (client *RpcCfxClient) GetBestBlockHash() (hash types.Hash, err error) {
	err = client.core.wrappedCallRPC(&hash, "cfx_getBestBlockHash")
	return
}

// GetRawBlockConfirmationRisk indicates the risk coefficient that
// the pivot block of the epoch where the block is located becomes a normal block.
// It will return nil if block not exist
func (client *RpcCfxClient) GetRawBlockConfirmationRisk(blockhash types.Hash) (risk *hexutil.Big, err error) {
	err = client.core.wrappedCallRPC(&risk, "cfx_getConfirmationRiskByHash", blockhash)
	return
}

// GetBlockConfirmationRisk indicates the probability that
// the pivot block of the epoch where the block is located becomes a normal block.
//
// it's (raw confirmation risk coefficient/ (2^256-1))
func (client *RpcCfxClient) GetBlockConfirmationRisk(blockHash types.Hash) (*big.Float, error) {
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
func (client *RpcCfxClient) SendTransaction(tx types.SignedTransaction) (types.Hash, error) {

	// err := client.core.PopulateTransaction(&tx)
	// if err != nil {
	// 	return "", errors.Wrap(err, cfxerrors.ErrMsgApplyTxValues)
	// }

	// //sign
	// if client.core.AccountManager == nil {
	// 	return "", errors.New("account manager not specified, see SetAccountManager")
	// }

	// rawData, err := client.core.AccountManager.SignTransaction(tx)
	// if err != nil {
	// 	return "", errors.Wrap(err, "failed to sign transaction")
	// }

	rawData, err := tx.Encode()
	if err != nil {
		return "", errors.Wrap(err, "failed to encode transaction")
	}
	//send raw tx
	txhash, err := client.SendRawTransaction(rawData)
	if err != nil {
		return "", errors.Wrapf(err, "failed to send transaction, raw data = 0x%+x", rawData)
	}
	return txhash, nil
}

// SendRawTransaction sends signed transaction and returns its hash.
func (client *RpcCfxClient) SendRawTransaction(rawData []byte) (hash types.Hash, err error) {
	err = client.core.wrappedCallRPC(&hash, "cfx_sendRawTransaction", hexutil.Encode(rawData))
	return
}

// Call executes a message call transaction "request" at specified epoch,
// which is directly executed in the VM of the node, but never mined into the block chain
// and returns the contract execution result.
func (client *RpcCfxClient) Call(request types.CallRequest, epoch *types.Epoch) (result hexutil.Bytes, err error) {
	err = client.core.wrappedCallRPC(&result, "cfx_call", request, epoch)
	return
}

// GetLogs returns logs that matching the specified filter.
func (client *RpcCfxClient) GetLogs(filter types.LogFilter) (logs []types.Log, err error) {
	err = client.core.wrappedCallRPC(&logs, "cfx_getLogs", filter)
	return
}

// GetTransactionByHash returns transaction for the specified txHash.
// If the transaction is not found, return nil.
func (client *RpcCfxClient) GetTransactionByHash(txHash types.Hash) (tx *types.Transaction, err error) {
	err = client.core.wrappedCallRPC(&tx, "cfx_getTransactionByHash", txHash)
	return
}

// EstimateGasAndCollateral excutes a message call "request"
// and returns the amount of the gas used and storage for collateral
func (client *RpcCfxClient) EstimateGasAndCollateral(request types.CallRequest, epoch ...*types.Epoch) (estimat types.Estimate, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.core.wrappedCallRPC(&estimat, "cfx_estimateGasAndCollateral", request, realEpoch)
	return
}

// GetBlocksByEpoch returns the blocks hash in the specified epoch.
func (client *RpcCfxClient) GetBlocksByEpoch(epoch *types.Epoch) (blockHashes []types.Hash, err error) {
	err = client.core.wrappedCallRPC(&blockHashes, "cfx_getBlocksByEpoch", epoch)
	return
}

// GetTransactionReceipt returns the receipt of specified transaction hash.
// If no receipt is found, return nil.
func (client *RpcCfxClient) GetTransactionReceipt(txHash types.Hash) (receipt *types.TransactionReceipt, err error) {
	err = client.core.wrappedCallRPC(&receipt, "cfx_getTransactionReceipt", txHash)
	return
}

// ===new rpc===

// GetAdmin returns admin of the given contract, it will return nil if contract not exist
func (client *RpcCfxClient) GetAdmin(contractAddress types.Address, epoch ...*types.Epoch) (admin *types.Address, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.core.wrappedCallRPC(&admin, "cfx_getAdmin", contractAddress, realEpoch)
	return
}

// GetSponsorInfo returns sponsor information of the given contract
func (client *RpcCfxClient) GetSponsorInfo(contractAddress types.Address, epoch ...*types.Epoch) (sponsor types.SponsorInfo, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.core.wrappedCallRPC(&sponsor, "cfx_getSponsorInfo", contractAddress, realEpoch)
	return
}

// GetStakingBalance returns balance of the given account.
func (client *RpcCfxClient) GetStakingBalance(account types.Address, epoch ...*types.Epoch) (balance *hexutil.Big, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.core.wrappedCallRPC(&balance, "cfx_getStakingBalance", account, realEpoch)
	return
}

// GetCollateralForStorage returns balance of the given account.
func (client *RpcCfxClient) GetCollateralForStorage(account types.Address, epoch ...*types.Epoch) (storage *hexutil.Big, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.core.wrappedCallRPC(&storage, "cfx_getCollateralForStorage", account, realEpoch)
	return
}

// GetStorageAt returns storage entries from a given contract.
func (client *RpcCfxClient) GetStorageAt(address types.Address, position types.Hash, epoch ...*types.Epoch) (storageEntries hexutil.Bytes, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.core.wrappedCallRPC(&storageEntries, "cfx_getStorageAt", address, position, realEpoch)
	return
}

// GetStorageRoot returns storage root of given address
func (client *RpcCfxClient) GetStorageRoot(address types.Address, epoch ...*types.Epoch) (storageRoot *types.StorageRoot, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.core.wrappedCallRPC(&storageRoot, "cfx_getStorageRoot", address, realEpoch)
	return
}

// GetBlockByHashWithPivotAssumption returns block with given hash and pivot chain assumption.
func (client *RpcCfxClient) GetBlockByHashWithPivotAssumption(blockHash types.Hash, pivotHash types.Hash, epoch hexutil.Uint64) (block types.Block, err error) {
	err = client.core.wrappedCallRPC(&block, "cfx_getBlockByHashWithPivotAssumption", blockHash, pivotHash, epoch)
	return
}

// CheckBalanceAgainstTransaction checks if user balance is enough for the transaction.
func (client *RpcCfxClient) CheckBalanceAgainstTransaction(accountAddress types.Address,
	contractAddress types.Address,
	gasLimit *hexutil.Big,
	gasPrice *hexutil.Big,
	storageLimit *hexutil.Big,
	epoch ...*types.Epoch) (response types.CheckBalanceAgainstTransactionResponse, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.core.wrappedCallRPC(&response,
		"cfx_checkBalanceAgainstTransaction", accountAddress, contractAddress,
		gasLimit, gasPrice, storageLimit, realEpoch)
	return
}

// GetSkippedBlocksByEpoch returns skipped block hashes of given epoch
func (client *RpcCfxClient) GetSkippedBlocksByEpoch(epoch *types.Epoch) (blockHashs []types.Hash, err error) {
	err = client.core.wrappedCallRPC(&blockHashs, "cfx_getSkippedBlocksByEpoch", epoch)
	return
}

// GetAccountInfo returns account related states of the given account
func (client *RpcCfxClient) GetAccountInfo(account types.Address, epoch ...*types.Epoch) (accountInfo types.AccountInfo, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.core.wrappedCallRPC(&accountInfo, "cfx_getAccount", account, realEpoch)
	return
}

// GetInterestRate returns interest rate of the given epoch
func (client *RpcCfxClient) GetInterestRate(epoch ...*types.Epoch) (intersetRate *hexutil.Big, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.core.wrappedCallRPC(&intersetRate, "cfx_getInterestRate", realEpoch)
	if err != nil {
		intersetRate = nil
	}
	return
}

// GetAccumulateInterestRate returns accumulate interest rate of the given epoch
func (client *RpcCfxClient) GetAccumulateInterestRate(epoch ...*types.Epoch) (intersetRate *hexutil.Big, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.core.wrappedCallRPC(&intersetRate, "cfx_getAccumulateInterestRate", realEpoch)
	if err != nil {
		intersetRate = nil
	}
	return
}

// GetBlockRewardInfo returns block reward information in an epoch
func (client *RpcCfxClient) GetBlockRewardInfo(epoch types.Epoch) (rewardInfo []types.RewardInfo, err error) {
	err = client.core.wrappedCallRPC(&rewardInfo, "cfx_getBlockRewardInfo", epoch)
	return
}

// GetClientVersion returns the client version as a string
func (client *RpcCfxClient) GetClientVersion() (clientVersion string, err error) {
	err = client.core.wrappedCallRPC(&clientVersion, "cfx_clientVersion")
	return
}

// GetDepositList returns deposit list of the given account.
func (client *RpcCfxClient) GetDepositList(address types.Address, epoch ...*types.Epoch) (depositInfos []types.DepositInfo, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.core.wrappedCallRPC(&depositInfos, "cfx_getDepositList", address, realEpoch)
	return
}

// GetVoteList returns vote list of the given account.
func (client *RpcCfxClient) GetVoteList(address types.Address, epoch ...*types.Epoch) (voteStakeInfos []types.VoteStakeInfo, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.core.wrappedCallRPC(&voteStakeInfos, "cfx_getVoteList", address, realEpoch)
	return
}

// GetSupplyInfo Return information about total token supply.
func (client *RpcCfxClient) GetSupplyInfo(epoch ...*types.Epoch) (info types.TokenSupplyInfo, err error) {
	realEpoch := get1stEpochIfy(epoch)
	err = client.core.wrappedCallRPC(&info, "cfx_getSupplyInfo", realEpoch)
	return
}

// GetAccountPendingInfo gets transaction pending info by account address
func (client *RpcCfxClient) GetAccountPendingInfo(address types.Address) (pendignInfo *types.AccountPendingInfo, err error) {
	err = client.core.wrappedCallRPC(&pendignInfo, "cfx_getAccountPendingInfo", address)
	return
}

func (client *RpcCfxClient) GetAccountPendingTransactions(address types.Address, startNonce *hexutil.Big, limit *hexutil.Uint64) (pendingTxs types.AccountPendingTransactions, err error) {
	err = client.core.wrappedCallRPC(&pendingTxs, "cfx_getAccountPendingTransactions", address, startNonce, limit)
	return
}
