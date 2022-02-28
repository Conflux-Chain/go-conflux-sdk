package bulk

import (
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	postypes "github.com/Conflux-Chain/go-conflux-sdk/types/pos"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// BulkCfxCaller used for bulk call rpc in one request to improve efficiency
type BulkCfxCaller BulkCallerCore

// NewBulkCfxCaller creates new BulkCfxCaller instance
func NewBulkCfxCaller(core BulkCallerCore) *BulkCfxCaller {
	return (*BulkCfxCaller)(&core)
}

// Execute sends all rpc requests in queue by rpc call "batch" on one request
func (b *BulkCfxCaller) Execute() ([]error, error) {
	return batchCall(b.caller, b.batchElems, nil)
}

// GetGasPrice returns the recent mean gas price.
func (client *BulkCfxCaller) GetGasPrice() (*hexutil.Big, *error) {
	result := new(hexutil.Big)
	err := new(error)

	elem := newBatchElem(result, "cfx_gasPrice")
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetNextNonce returns the next transaction nonce of address
func (client *BulkCfxCaller) GetNextNonce(address types.Address, epoch ...*types.Epoch) (*hexutil.Big, *error) {
	result := new(hexutil.Big)
	err := new(error)
	realEpoch := get1stEpochIfy(epoch)

	elem := newBatchElem(result, "cfx_getNextNonce", address, realEpoch)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetStatus returns status of connecting conflux node
func (client *BulkCfxCaller) GetStatus() (*types.Status, *error) {
	result := new(types.Status)
	err := new(error)

	elem := newBatchElem(result, "cfx_getStatus")
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetEpochNumber returns the highest or specified epoch number.
func (client *BulkCfxCaller) GetEpochNumber(epoch ...*types.Epoch) (*hexutil.Big, *error) {
	result := new(hexutil.Big)
	err := new(error)
	realEpoch := get1stEpochIfy(epoch)

	elem := newBatchElem(result, "cfx_epochNumber", realEpoch)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)

	return result, err
}

// GetBalance returns the balance of specified address at epoch.
func (client *BulkCfxCaller) GetBalance(address types.Address, epoch ...*types.Epoch) (*hexutil.Big, *error) {
	result := new(hexutil.Big)
	err := new(error)
	realEpoch := get1stEpochIfy(epoch)

	elem := newBatchElem(result, "cfx_getBalance", address, realEpoch)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)

	return result, err
}

// GetCode returns the bytecode in HEX format of specified address at epoch.
func (client *BulkCfxCaller) GetCode(address types.Address, epoch ...*types.Epoch) (*hexutil.Bytes, *error) {
	result := new(hexutil.Bytes)
	err := new(error)
	realEpoch := get1stEpochIfy(epoch)

	elem := newBatchElem(result, "cfx_getCode", address, realEpoch)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetBlockSummaryByHash returns the block summary of specified blockHash
// If the block is not found, return nil.
func (client *BulkCfxCaller) GetBlockSummaryByHash(blockHash types.Hash) (*types.BlockSummary, *error) {
	result := new(types.BlockSummary)
	err := new(error)

	elem := newBatchElem(result, "cfx_getBlockByHash", blockHash, false)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetBlockByHash returns the block of specified blockHash
// If the block is not found, return nil.
func (client *BulkCfxCaller) GetBlockByHash(blockHash types.Hash) (*types.Block, *error) {
	result := new(types.Block)
	err := new(error)

	elem := newBatchElem(result, "cfx_getBlockByHash", blockHash, true)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetBlockSummaryByEpoch returns the block summary of specified epoch.
// If the epoch is invalid, return the concrete error.
func (client *BulkCfxCaller) GetBlockSummaryByEpoch(epoch *types.Epoch) (*types.BlockSummary, *error) {
	result := new(types.BlockSummary)
	err := new(error)

	elem := newBatchElem(result, "cfx_getBlockByEpochNumber", epoch, false)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetBlockByHash returns the block of specified block number
func (client *BulkCfxCaller) GetBlockByBlockNumber(blockNumer hexutil.Uint64) (*types.Block, *error) {
	result := new(types.Block)
	err := new(error)

	elem := newBatchElem(result, "cfx_getBlockByBlockNumber", blockNumer, true)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetBlockSummaryByBlockNumber returns the block summary of specified block number.
func (client *BulkCfxCaller) GetBlockSummaryByBlockNumber(blockNumer hexutil.Uint64) (*types.BlockSummary, *error) {
	result := new(types.BlockSummary)
	err := new(error)

	elem := newBatchElem(result, "cfx_getBlockByBlockNumber", blockNumer, false)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetBlockByEpoch returns the block of specified epoch.
// If the epoch is invalid, return the concrete error.
func (client *BulkCfxCaller) GetBlockByEpoch(epoch *types.Epoch) (*types.Block, *error) {
	result := new(types.Block)
	err := new(error)

	elem := newBatchElem(result, "cfx_getBlockByEpochNumber", epoch, true)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetBestBlockHash returns the current best block hash.
func (client *BulkCfxCaller) GetBestBlockHash() (*types.Hash, *error) {
	result := new(types.Hash)
	err := new(error)

	elem := newBatchElem(result, "cfx_getBestBlockHash")
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetRawBlockConfirmationRisk indicates the risk coefficient that
// the pivot block of the epoch where the block is located becomes a normal block.
// It will return nil if block not exist
func (client *BulkCfxCaller) GetRawBlockConfirmationRisk(blockhash types.Hash) (*hexutil.Big, *error) {
	result := new(hexutil.Big)
	err := new(error)

	elem := newBatchElem(result, "cfx_getConfirmationRiskByHash", blockhash)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

func (client *BulkCfxCaller) SendRawTransaction(rawData []byte) (*types.Hash, *error) {
	result := new(types.Hash)
	err := new(error)

	elem := newBatchElem(result, "cfx_sendRawTransaction", hexutil.Encode(rawData))
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// Call executes a message call transaction "request" at specified epoch,
// which is directly executed in the VM of the node, but never mined into the block chain
// and returns the contract execution result.
func (client *BulkCfxCaller) Call(request types.CallRequest, epoch *types.Epoch) (*hexutil.Bytes, *error) {
	result := new(hexutil.Bytes)
	err := new(error)

	elem := newBatchElem(result, "cfx_call", request, epoch)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetLogs returns logs that matching the specified filter.
func (client *BulkCfxCaller) GetLogs(filter types.LogFilter) ([]types.Log, *error) {
	result := make([]types.Log, 0)
	err := new(error)

	elem := newBatchElem(result, "cfx_getLogs", filter)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetTransactionByHash returns transaction for the specified txHash.
// If the transaction is not found, return nil.
func (client *BulkCfxCaller) GetTransactionByHash(txHash types.Hash) (*types.Transaction, *error) {
	result := new(types.Transaction)
	err := new(error)

	elem := newBatchElem(result, "cfx_getTransactionByHash", txHash)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// EstimateGasAndCollateral excutes a message call "request"
// and returns the amount of the gas used and storage for collateral
func (client *BulkCfxCaller) EstimateGasAndCollateral(request types.CallRequest, epoch ...*types.Epoch) (*types.Estimate, *error) {
	result := new(types.Estimate)
	err := new(error)
	realEpoch := get1stEpochIfy(epoch)

	elem := newBatchElem(result, "cfx_estimateGasAndCollateral", request, realEpoch)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetBlocksByEpoch returns the blocks hash in the specified epoch.
func (client *BulkCfxCaller) GetBlocksByEpoch(epoch *types.Epoch) ([]types.Hash, *error) {
	result := make([]types.Hash, 0)
	err := new(error)

	elem := newBatchElem(result, "cfx_getBlocksByEpoch", epoch)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetTransactionReceipt returns the receipt of specified transaction hash.
// If no receipt is found, return nil.
func (client *BulkCfxCaller) GetTransactionReceipt(txHash types.Hash) (*types.TransactionReceipt, *error) {
	result := new(types.TransactionReceipt)
	err := new(error)

	elem := newBatchElem(result, "cfx_getTransactionReceipt", txHash)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// ===new rpc===

// GetAdmin returns admin of the given contract, it will return nil if contract not exist
func (client *BulkCfxCaller) GetAdmin(contractAddress types.Address, epoch ...*types.Epoch) (*types.Address, *error) {
	result := new(types.Address)
	err := new(error)
	realEpoch := get1stEpochIfy(epoch)

	elem := newBatchElem(result, "cfx_getAdmin", contractAddress, realEpoch)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetSponsorInfo returns sponsor information of the given contract
func (client *BulkCfxCaller) GetSponsorInfo(contractAddress types.Address, epoch ...*types.Epoch) (*types.SponsorInfo, *error) {
	result := new(types.SponsorInfo)
	err := new(error)
	realEpoch := get1stEpochIfy(epoch)

	elem := newBatchElem(result, "cfx_getSponsorInfo", contractAddress, realEpoch)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetStakingBalance returns balance of the given account.
func (client *BulkCfxCaller) GetStakingBalance(account types.Address, epoch ...*types.Epoch) (*hexutil.Big, *error) {
	result := new(hexutil.Big)
	err := new(error)
	realEpoch := get1stEpochIfy(epoch)

	elem := newBatchElem(result, "cfx_getStakingBalance", account, realEpoch)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetCollateralForStorage returns balance of the given account.
func (client *BulkCfxCaller) GetCollateralForStorage(account types.Address, epoch ...*types.Epoch) (*hexutil.Big, *error) {
	result := new(hexutil.Big)
	err := new(error)
	realEpoch := get1stEpochIfy(epoch)

	elem := newBatchElem(result, "cfx_getCollateralForStorage", account, realEpoch)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetStorageAt returns storage entries from a given contract.
func (client *BulkCfxCaller) GetStorageAt(address types.Address, position types.Hash, epoch ...*types.Epoch) (*hexutil.Bytes, *error) {
	result := new(hexutil.Bytes)
	err := new(error)
	realEpoch := get1stEpochIfy(epoch)

	elem := newBatchElem(result, "cfx_getStorageAt", address, position, realEpoch)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetStorageRoot returns storage root of given address
func (client *BulkCfxCaller) GetStorageRoot(address types.Address, epoch ...*types.Epoch) (*types.StorageRoot, *error) {
	result := new(types.StorageRoot)
	err := new(error)
	realEpoch := get1stEpochIfy(epoch)

	elem := newBatchElem(result, "cfx_getStorageRoot", address, realEpoch)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetBlockByHashWithPivotAssumption returns block with given hash and pivot chain assumption.
func (client *BulkCfxCaller) GetBlockByHashWithPivotAssumption(blockHash types.Hash, pivotHash types.Hash, epoch hexutil.Uint64) (*types.Block, *error) {
	result := new(types.Block)
	err := new(error)

	elem := newBatchElem(result, "cfx_getBlockByHashWithPivotAssumption", blockHash, pivotHash, epoch)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// CheckBalanceAgainstTransaction checks if user balance is enough for the transaction.
func (client *BulkCfxCaller) CheckBalanceAgainstTransaction(accountAddress types.Address,
	contractAddress types.Address,
	gasLimit *hexutil.Big,
	gasPrice *hexutil.Big,
	storageLimit *hexutil.Big,
	epoch ...*types.Epoch) (*types.CheckBalanceAgainstTransactionResponse, *error) {
	result := new(types.CheckBalanceAgainstTransactionResponse)
	err := new(error)
	realEpoch := get1stEpochIfy(epoch)

	elem := newBatchElem(result,
		"cfx_checkBalanceAgainstTransaction", accountAddress, contractAddress,
		gasLimit, gasPrice, storageLimit, realEpoch)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetSkippedBlocksByEpoch returns skipped block hashes of given epoch
func (client *BulkCfxCaller) GetSkippedBlocksByEpoch(epoch *types.Epoch) ([]types.Hash, *error) {
	result := make([]types.Hash, 0)
	err := new(error)

	elem := newBatchElem(result, "cfx_getSkippedBlocksByEpoch", epoch)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetAccountInfo returns account related states of the given account
func (client *BulkCfxCaller) GetAccountInfo(account types.Address, epoch ...*types.Epoch) (*types.AccountInfo, *error) {
	result := new(types.AccountInfo)
	err := new(error)
	realEpoch := get1stEpochIfy(epoch)

	elem := newBatchElem(result, "cfx_getAccount", account, realEpoch)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetInterestRate returns interest rate of the given epoch
func (client *BulkCfxCaller) GetInterestRate(epoch ...*types.Epoch) (*hexutil.Big, *error) {
	result := new(hexutil.Big)
	err := new(error)
	realEpoch := get1stEpochIfy(epoch)

	elem := newBatchElem(result, "cfx_getInterestRate", realEpoch)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)

	return result, err
}

// GetAccumulateInterestRate returns accumulate interest rate of the given epoch
func (client *BulkCfxCaller) GetAccumulateInterestRate(epoch ...*types.Epoch) (*hexutil.Big, *error) {
	result := new(hexutil.Big)
	err := new(error)
	realEpoch := get1stEpochIfy(epoch)

	elem := newBatchElem(result, "cfx_getAccumulateInterestRate", realEpoch)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)

	return result, err
}

// GetBlockRewardInfo returns block reward information in an epoch
func (client *BulkCfxCaller) GetBlockRewardInfo(epoch types.Epoch) ([]types.RewardInfo, *error) {
	result := make([]types.RewardInfo, 0)
	err := new(error)

	elem := newBatchElem(result, "cfx_getBlockRewardInfo", epoch)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetClientVersion returns the client version as a string
func (client *BulkCfxCaller) GetClientVersion() (*string, *error) {
	result := new(string)
	err := new(error)

	elem := newBatchElem(result, "cfx_clientVersion")
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetDepositList returns deposit list of the given account.
func (client *BulkCfxCaller) GetDepositList(address types.Address, epoch ...*types.Epoch) ([]types.DepositInfo, *error) {
	result := make([]types.DepositInfo, 0)
	err := new(error)
	realEpoch := get1stEpochIfy(epoch)

	elem := newBatchElem(result, "cfx_getDepositList", address, realEpoch)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetVoteList returns vote list of the given account.
func (client *BulkCfxCaller) GetVoteList(address types.Address, epoch ...*types.Epoch) ([]types.VoteStakeInfo, *error) {
	result := make([]types.VoteStakeInfo, 0)
	err := new(error)
	realEpoch := get1stEpochIfy(epoch)

	elem := newBatchElem(result, "cfx_getVoteList", address, realEpoch)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetSupplyInfo Return information about total token supply.
func (client *BulkCfxCaller) GetSupplyInfo(epoch ...*types.Epoch) (*types.TokenSupplyInfo, *error) {
	result := new(types.TokenSupplyInfo)
	err := new(error)
	realEpoch := get1stEpochIfy(epoch)

	elem := newBatchElem(result, "cfx_getSupplyInfo", realEpoch)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetAccountPendingInfo gets transaction pending info by account address
func (client *BulkCfxCaller) GetAccountPendingInfo(address types.Address) (*types.AccountPendingInfo, *error) {
	result := new(types.AccountPendingInfo)
	err := new(error)

	elem := newBatchElem(result, "cfx_getAccountPendingInfo", address)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetAccountPendingTransactions get transaction pending info by account address
func (client *BulkCfxCaller) GetAccountPendingTransactions(address types.Address, startNonce *hexutil.Big, limit *hexutil.Uint64) (*types.AccountPendingTransactions, *error) {
	result := new(types.AccountPendingTransactions)
	err := new(error)

	elem := newBatchElem(result, "cfx_getAccountPendingTransactions", address, startNonce, limit)
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetPoSEconomics returns accumulate interest rate of the given epoch
func (client *BulkCfxCaller) GetPoSEconomics(epoch ...*types.Epoch) (*types.PoSEconomics, *error) {
	result := new(types.PoSEconomics)
	err := new(error)

	elem := newBatchElem(result, "cfx_getPoSEconomics", get1stEpochIfy(epoch))
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetOpenedMethodGroups returns openning method groups
func (client *BulkCfxCaller) GetOpenedMethodGroups() (*[]string, *error) {
	result := new([]string)
	err := new(error)

	elem := newBatchElem(result, "cfx_openedMethodGroups")
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}

// GetPoSRewardByEpoch returns PoS reward in the epoch
func (client *BulkCfxCaller) GetPoSRewardByEpoch(epoch types.Epoch) (*postypes.EpochReward, *error) {
	result := new(postypes.EpochReward)
	err := new(error)

	elem := newBatchElem(result, "cfx_getPoSRewardByEpoch")
	(*BulkCallerCore)(client).appendElemsAndError(elem, err)
	return result, err
}
