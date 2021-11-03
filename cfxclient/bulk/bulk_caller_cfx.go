package bulk

import (
	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type BulkCfxCaller struct {
	caller     sdk.ClientOperator
	batchElems *[]rpc.BatchElem
}

func NewBulkCfxCaller(caller sdk.ClientOperator, batchElems *[]rpc.BatchElem) *BulkCfxCaller {
	return &BulkCfxCaller{caller, batchElems}
}

func (b *BulkCfxCaller) Excute() ([]error, error) {
	return batchCall(b.caller, b.batchElems, nil)
}

//ignore

//ignore

func (client *BulkCfxCaller) GetGasPrice() *hexutil.Big {
	result := &hexutil.Big{}
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_gasPrice"))
	return result
}

// GetNextNonce returns the next transaction nonce of address
func (client *BulkCfxCaller) GetNextNonce(address types.Address, epoch ...*types.Epoch) *hexutil.Big {
	result := &hexutil.Big{}
	realEpoch := get1stEpochIfy(epoch)
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getNextNonce", address, realEpoch))
	return result
}

// GetStatus returns status of connecting conflux node
//ignore

func (client *BulkCfxCaller) GetEpochNumber(epoch ...*types.Epoch) *hexutil.Big {
	result := &hexutil.Big{}
	realEpoch := get1stEpochIfy(epoch)
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_epochNumber", realEpoch))

	return result
}

// GetBalance returns the balance of specified address at epoch.
func (client *BulkCfxCaller) GetBalance(address types.Address, epoch ...*types.Epoch) *hexutil.Big {
	result := &hexutil.Big{}
	realEpoch := get1stEpochIfy(epoch)
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getBalance", address, realEpoch))

	return result
}

// GetCode returns the bytecode in HEX format of specified address at epoch.
func (client *BulkCfxCaller) GetCode(address types.Address, epoch ...*types.Epoch) *hexutil.Bytes {
	result := &hexutil.Bytes{}
	realEpoch := get1stEpochIfy(epoch)
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getCode", address, realEpoch))
	return result
}

// GetBlockSummaryByHash returns the block summary of specified blockHash
// If the block is not found, return nil.
func (client *BulkCfxCaller) GetBlockSummaryByHash(blockHash types.Hash) *types.BlockSummary {
	result := &types.BlockSummary{}
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getBlockByHash", blockHash, false))
	return result
}

// GetBlockByHash returns the block of specified blockHash
// If the block is not found, return nil.
func (client *BulkCfxCaller) GetBlockByHash(blockHash types.Hash) *types.Block {
	result := &types.Block{}
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getBlockByHash", blockHash, true))
	return result
}

// GetBlockSummaryByEpoch returns the block summary of specified epoch.
// If the epoch is invalid, return the concrete error.
func (client *BulkCfxCaller) GetBlockSummaryByEpoch(epoch *types.Epoch) *types.BlockSummary {
	result := &types.BlockSummary{}
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getBlockByEpochNumber", epoch, false))
	return result
}

func (client *BulkCfxCaller) GetBlockByBlockNumber(blockNumer hexutil.Uint64) *types.Block {
	result := &types.Block{}
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getBlockByBlockNumber", blockNumer, true))
	return result
}

func (client *BulkCfxCaller) GetBlockSummaryByBlockNumber(blockNumer hexutil.Uint64) *types.BlockSummary {
	result := &types.BlockSummary{}
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getBlockByBlockNumber", blockNumer, false))
	return result
}

// GetBlockByEpoch returns the block of specified epoch.
// If the epoch is invalid, return the concrete error.
func (client *BulkCfxCaller) GetBlockByEpoch(epoch *types.Epoch) *types.Block {
	result := &types.Block{}
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getBlockByEpochNumber", epoch, true))
	return result
}

// GetBestBlockHash returns the current best block hash.
func (client *BulkCfxCaller) GetBestBlockHash() *types.Hash {
	var hash types.Hash
	result := &hash
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getBestBlockHash"))
	return result
}

// GetRawBlockConfirmationRisk indicates the risk coefficient that
// the pivot block of the epoch where the block is located becomes a normal block.
// It will return nil if block not exist
func (client *BulkCfxCaller) GetRawBlockConfirmationRisk(blockhash types.Hash) *hexutil.Big {
	result := &hexutil.Big{}
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getConfirmationRiskByHash", blockhash))
	return result
}

// GetBlockConfirmationRisk indicates the probability that
// the pivot block of the epoch where the block is located becomes a normal block.
//
// it's (raw confirmation risk coefficient/ (2^256-1))
//ignore

//ignore

func (client *BulkCfxCaller) SendRawTransaction(rawData []byte) *types.Hash {
	var hash types.Hash
	result := &hash
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_sendRawTransaction", hexutil.Encode(rawData)))
	return result
}

// Call executes a message call transaction "request" at specified epoch,
// which is directly executed in the VM of the node, but never mined into the block chain
// and returns the contract execution result.
func (client *BulkCfxCaller) Call(request types.CallRequest, epoch *types.Epoch) *hexutil.Bytes {
	var bytes hexutil.Bytes
	result := &bytes
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_call", request, epoch))
	return result
}

// GetLogs returns logs that matching the specified filter.
func (client *BulkCfxCaller) GetLogs(filter types.LogFilter) []types.Log {
	result := make([]types.Log, 0)
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getLogs", filter))
	return result
}

// GetTransactionByHash returns transaction for the specified txHash.
// If the transaction is not found, return nil.
func (client *BulkCfxCaller) GetTransactionByHash(txHash types.Hash) *types.Transaction {
	result := &types.Transaction{}
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getTransactionByHash", txHash))
	return result
}

// EstimateGasAndCollateral excutes a message call "request"
// and returns the amount of the gas used and storage for collateral
func (client *BulkCfxCaller) EstimateGasAndCollateral(request types.CallRequest, epoch ...*types.Epoch) *types.Estimate {
	result := &types.Estimate{}
	realEpoch := get1stEpochIfy(epoch)
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_estimateGasAndCollateral", request, realEpoch))
	return result
}

// GetBlocksByEpoch returns the blocks hash in the specified epoch.
func (client *BulkCfxCaller) GetBlocksByEpoch(epoch *types.Epoch) []types.Hash {
	result := make([]types.Hash, 0)
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getBlocksByEpoch", epoch))
	return result
}

// GetTransactionReceipt returns the receipt of specified transaction hash.
// If no receipt is found, return nil.
func (client *BulkCfxCaller) GetTransactionReceipt(txHash types.Hash) *types.TransactionReceipt {
	result := &types.TransactionReceipt{}
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getTransactionReceipt", txHash))
	return result
}

// ===new rpc===

// GetAdmin returns admin of the given contract, it will return nil if contract not exist
func (client *BulkCfxCaller) GetAdmin(contractAddress types.Address, epoch ...*types.Epoch) *types.Address {
	result := &types.Address{}
	realEpoch := get1stEpochIfy(epoch)
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getAdmin", contractAddress, realEpoch))
	return result
}

// GetSponsorInfo returns sponsor information of the given contract
func (client *BulkCfxCaller) GetSponsorInfo(contractAddress types.Address, epoch ...*types.Epoch) *types.SponsorInfo {
	result := &types.SponsorInfo{}
	realEpoch := get1stEpochIfy(epoch)
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getSponsorInfo", contractAddress, realEpoch))
	return result
}

// GetStakingBalance returns balance of the given account.
func (client *BulkCfxCaller) GetStakingBalance(account types.Address, epoch ...*types.Epoch) *hexutil.Big {
	result := &hexutil.Big{}
	realEpoch := get1stEpochIfy(epoch)
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getStakingBalance", account, realEpoch))
	return result
}

// GetCollateralForStorage returns balance of the given account.
func (client *BulkCfxCaller) GetCollateralForStorage(account types.Address, epoch ...*types.Epoch) *hexutil.Big {
	result := &hexutil.Big{}
	realEpoch := get1stEpochIfy(epoch)
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getCollateralForStorage", account, realEpoch))
	return result
}

// GetStorageAt returns storage entries from a given contract.
func (client *BulkCfxCaller) GetStorageAt(address types.Address, position types.Hash, epoch ...*types.Epoch) *hexutil.Bytes {
	result := &hexutil.Bytes{}
	realEpoch := get1stEpochIfy(epoch)
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getStorageAt", address, position, realEpoch))
	return result
}

// GetStorageRoot returns storage root of given address
func (client *BulkCfxCaller) GetStorageRoot(address types.Address, epoch ...*types.Epoch) *types.StorageRoot {
	result := &types.StorageRoot{}
	realEpoch := get1stEpochIfy(epoch)
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getStorageRoot", address, realEpoch))
	return result
}

// GetBlockByHashWithPivotAssumption returns block with given hash and pivot chain assumption.
func (client *BulkCfxCaller) GetBlockByHashWithPivotAssumption(blockHash types.Hash, pivotHash types.Hash, epoch hexutil.Uint64) *types.Block {
	result := &types.Block{}
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getBlockByHashWithPivotAssumption", blockHash, pivotHash, epoch))
	return result
}

// CheckBalanceAgainstTransaction checks if user balance is enough for the transaction.
func (client *BulkCfxCaller) CheckBalanceAgainstTransaction(accountAddress types.Address,
	contractAddress types.Address,
	gasLimit *hexutil.Big,
	gasPrice *hexutil.Big,
	storageLimit *hexutil.Big,
	epoch ...*types.Epoch) *types.CheckBalanceAgainstTransactionResponse {
	result := &types.CheckBalanceAgainstTransactionResponse{}
	realEpoch := get1stEpochIfy(epoch)
	*client.batchElems = append(*client.batchElems, newBatchElem(result,
		"cfx_checkBalanceAgainstTransaction", accountAddress, contractAddress,
		gasLimit, gasPrice, storageLimit, realEpoch))
	return result
}

// GetSkippedBlocksByEpoch returns skipped block hashes of given epoch
func (client *BulkCfxCaller) GetSkippedBlocksByEpoch(epoch *types.Epoch) []types.Hash {
	result := make([]types.Hash, 0)
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getSkippedBlocksByEpoch", epoch))
	return result
}

// GetAccountInfo returns account related states of the given account
func (client *BulkCfxCaller) GetAccountInfo(account types.Address, epoch ...*types.Epoch) *types.AccountInfo {
	result := &types.AccountInfo{}
	realEpoch := get1stEpochIfy(epoch)
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getAccount", account, realEpoch))
	return result
}

// GetInterestRate returns interest rate of the given epoch
func (client *BulkCfxCaller) GetInterestRate(epoch ...*types.Epoch) *hexutil.Big {
	result := &hexutil.Big{}
	realEpoch := get1stEpochIfy(epoch)
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getInterestRate", realEpoch))

	return result
}

// GetAccumulateInterestRate returns accumulate interest rate of the given epoch
func (client *BulkCfxCaller) GetAccumulateInterestRate(epoch ...*types.Epoch) *hexutil.Big {
	result := &hexutil.Big{}
	realEpoch := get1stEpochIfy(epoch)
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getAccumulateInterestRate", realEpoch))

	return result
}

// GetBlockRewardInfo returns block reward information in an epoch
func (client *BulkCfxCaller) GetBlockRewardInfo(epoch types.Epoch) []types.RewardInfo {
	result := make([]types.RewardInfo, 0)
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getBlockRewardInfo", epoch))
	return result
}

// GetClientVersion returns the client version as a string
func (client *BulkCfxCaller) GetClientVersion() *string {
	var str string
	result := &str
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_clientVersion"))
	return result
}

// GetDepositList returns deposit list of the given account.
func (client *BulkCfxCaller) GetDepositList(address types.Address, epoch ...*types.Epoch) []types.DepositInfo {
	result := make([]types.DepositInfo, 0)
	realEpoch := get1stEpochIfy(epoch)
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getDepositList", address, realEpoch))
	return result
}

// GetVoteList returns vote list of the given account.
func (client *BulkCfxCaller) GetVoteList(address types.Address, epoch ...*types.Epoch) []types.VoteStakeInfo {
	result := make([]types.VoteStakeInfo, 0)
	realEpoch := get1stEpochIfy(epoch)
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getVoteList", address, realEpoch))
	return result
}

// GetSupplyInfo Return information about total token supply.
func (client *BulkCfxCaller) GetSupplyInfo(epoch ...*types.Epoch) *types.TokenSupplyInfo {
	result := &types.TokenSupplyInfo{}
	realEpoch := get1stEpochIfy(epoch)
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getSupplyInfo", realEpoch))
	return result
}

// GetAccountPendingInfo gets transaction pending info by account address
func (client *BulkCfxCaller) GetAccountPendingInfo(address types.Address) *types.AccountPendingInfo {
	result := &types.AccountPendingInfo{}
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getAccountPendingInfo", address))
	return result
}

func (client *BulkCfxCaller) GetAccountPendingTransactions(address types.Address, startNonce *hexutil.Big, limit *hexutil.Uint64) *types.AccountPendingTransactions {
	result := &types.AccountPendingTransactions{}
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getAccountPendingTransactions", address, startNonce, limit))
	return result
}

// GetStatus returns status of connecting conflux node
func (client *BulkCfxCaller) GetStatus() *hexutil.Big {
	result := &hexutil.Big{}
	*client.batchElems = append(*client.batchElems, newBatchElem(result, "cfx_getStatus"))
	return result
}
