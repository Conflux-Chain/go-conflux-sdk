// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package sdk

import (
	"math/big"
	"net/http"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	postypes "github.com/Conflux-Chain/go-conflux-sdk/types/pos"
	"github.com/ethereum/go-ethereum/common/hexutil"
	rpc "github.com/openweb3/go-rpc-provider"
	"github.com/openweb3/go-rpc-provider/interfaces"
)

// HTTPRequester is interface for emitting a http requester
type HTTPRequester interface {
	Get(url string) (resp *http.Response, err error)
}

// Contractor is interface of contract operator
type Contractor interface {
	GetData(method string, args ...interface{}) ([]byte, error)
	Call(option *types.ContractMethodCallOption, resultPtr interface{}, method string, args ...interface{}) error
	SendTransaction(option *types.ContractMethodSendOption, method string, args ...interface{}) (types.Hash, error)
	DecodeEvent(out interface{}, event string, log types.Log) error
}

// ClientOperator is interface of operate actions on client
type ClientOperator interface {
	GetNodeURL() string
	NewAddress(base32OrHex string) (types.Address, error)
	MustNewAddress(base32OrHex string) types.Address

	CallRPC(result interface{}, method string, args ...interface{}) error
	BatchCallRPC(b []rpc.BatchElem) error

	SetAccountManager(accountManager AccountManagerOperator)
	GetAccountManager() AccountManagerOperator

	Pos() RpcPos
	TxPool() RpcTxpool
	Debug() RpcDebug
	Filter() RpcFilter
	Trace() RpcTrace

	GetGasPrice() (*hexutil.Big, error)
	GetNextNonce(address types.Address, epoch ...*types.EpochOrBlockHash) (*hexutil.Big, error)
	GetStatus() (types.Status, error)
	GetNetworkID() (uint32, error)
	GetChainID() (uint32, error)
	GetEpochNumber(epoch ...*types.Epoch) (*hexutil.Big, error)
	GetBalance(address types.Address, epoch ...*types.EpochOrBlockHash) (*hexutil.Big, error)
	GetCode(address types.Address, epoch ...*types.EpochOrBlockHash) (hexutil.Bytes, error)
	GetBlockSummaryByHash(blockHash types.Hash) (*types.BlockSummary, error)
	GetBlockByHash(blockHash types.Hash) (*types.Block, error)
	GetBlockSummaryByEpoch(epoch *types.Epoch) (*types.BlockSummary, error)
	GetBlockByEpoch(epoch *types.Epoch) (*types.Block, error)
	GetBlockByBlockNumber(blockNumer hexutil.Uint64) (block *types.Block, err error)
	GetBlockSummaryByBlockNumber(blockNumer hexutil.Uint64) (block *types.BlockSummary, err error)
	GetBestBlockHash() (types.Hash, error)
	GetRawBlockConfirmationRisk(blockhash types.Hash) (*hexutil.Big, error)
	GetBlockConfirmationRisk(blockHash types.Hash) (*big.Float, error)

	SendTransaction(tx types.UnsignedTransaction) (types.Hash, error)
	SendRawTransaction(rawData []byte) (types.Hash, error)
	SignEncodedTransactionAndSend(encodedTx []byte, v byte, r, s []byte) (*types.Transaction, error)

	Call(request types.CallRequest, epoch *types.EpochOrBlockHash) (hexutil.Bytes, error)

	GetLogs(filter types.LogFilter) ([]types.Log, error)
	GetTransactionByHash(txHash types.Hash) (*types.Transaction, error)
	EstimateGasAndCollateral(request types.CallRequest, epoch ...*types.Epoch) (types.Estimate, error)
	GetBlocksByEpoch(epoch *types.Epoch) ([]types.Hash, error)
	GetTransactionReceipt(txHash types.Hash) (*types.TransactionReceipt, error)
	GetAdmin(contractAddress types.Address, epoch ...*types.Epoch) (admin *types.Address, err error)
	GetSponsorInfo(contractAddress types.Address, epoch ...*types.Epoch) (sponsor types.SponsorInfo, err error)
	GetStakingBalance(account types.Address, epoch ...*types.Epoch) (balance *hexutil.Big, err error)
	GetCollateralForStorage(account types.Address, epoch ...*types.Epoch) (storage *hexutil.Big, err error)
	GetStorageAt(address types.Address, position *hexutil.Big, epoch ...*types.EpochOrBlockHash) (storageEntries hexutil.Bytes, err error)
	GetStorageRoot(address types.Address, epoch ...*types.Epoch) (storageRoot *types.StorageRoot, err error)
	GetBlockByHashWithPivotAssumption(blockHash types.Hash, pivotHash types.Hash, epoch hexutil.Uint64) (block types.Block, err error)
	CheckBalanceAgainstTransaction(accountAddress types.Address,
		contractAddress types.Address,
		gasLimit *hexutil.Big,
		gasPrice *hexutil.Big,
		storageLimit *hexutil.Big,
		epoch ...*types.Epoch) (response types.CheckBalanceAgainstTransactionResponse, err error)
	GetSkippedBlocksByEpoch(epoch *types.Epoch) (blockHashs []types.Hash, err error)
	GetAccountInfo(account types.Address, epoch ...*types.Epoch) (accountInfo types.AccountInfo, err error)
	GetInterestRate(epoch ...*types.Epoch) (intersetRate *hexutil.Big, err error)
	GetAccumulateInterestRate(epoch ...*types.Epoch) (intersetRate *hexutil.Big, err error)
	GetBlockRewardInfo(epoch types.Epoch) (rewardInfo []types.RewardInfo, err error)

	GetClientVersion() (clientVersion string, err error)
	GetDepositList(address types.Address, epoch ...*types.Epoch) ([]types.DepositInfo, error)
	GetVoteList(address types.Address, epoch ...*types.Epoch) ([]types.VoteStakeInfo, error)
	GetSupplyInfo(epoch ...*types.Epoch) (info types.TokenSupplyInfo, err error)

	CreateUnsignedTransaction(from types.Address, to types.Address, amount *hexutil.Big, data []byte) (types.UnsignedTransaction, error)
	ApplyUnsignedTransactionDefault(tx *types.UnsignedTransaction) error

	DeployContract(option *types.ContractDeployOption, abiJSON []byte,
		bytecode []byte, constroctorParams ...interface{}) *ContractDeployResult
	GetContract(abiJSON []byte, deployedAt *types.Address) (*Contract, error)
	GetAccountPendingInfo(address types.Address) (pendignInfo *types.AccountPendingInfo, err error)
	GetAccountPendingTransactions(address types.Address, startNonce *hexutil.Big, limit *hexutil.Uint64) (pendingTxs types.AccountPendingTransactions, err error)
	GetPoSEconomics(epoch ...*types.Epoch) (posEconomics types.PoSEconomics, err error)
	GetOpenedMethodGroups() (openedGroups []string, err error)
	GetPoSRewardByEpoch(epoch types.Epoch) (reward *postypes.EpochReward, err error)

	GetEpochReceipts(epoch types.EpochOrBlockHash) (receipts [][]types.TransactionReceipt, err error)
	GetEpochReceiptsByPivotBlockHash(hash types.Hash) (receipts [][]types.TransactionReceipt, err error)

	GetParamsFromVote(epoch ...*types.Epoch) (info postypes.VoteParamsInfo, err error)

	BatchGetTxByHashes(txhashes []types.Hash) (map[types.Hash]*types.Transaction, error)
	BatchGetBlockSummarys(blockhashes []types.Hash) (map[types.Hash]*types.BlockSummary, error)
	BatchGetBlockSummarysByNumber(blocknumbers []hexutil.Uint64) (map[hexutil.Uint64]*types.BlockSummary, error)
	BatchGetRawBlockConfirmationRisk(blockhashes []types.Hash) (map[types.Hash]*big.Int, error)
	BatchGetBlockConfirmationRisk(blockhashes []types.Hash) (map[types.Hash]*big.Float, error)

	Close()

	SubscribeNewHeads(channel chan types.BlockHeader) (*rpc.ClientSubscription, error)
	SubscribeEpochs(channel chan types.WebsocketEpochResponse, subscriptionEpochType ...types.Epoch) (*rpc.ClientSubscription, error)
	SubscribeLogs(channel chan types.SubscriptionLog, filter types.LogFilter) (*rpc.ClientSubscription, error)

	WaitForTransationBePacked(txhash types.Hash, duration time.Duration) (*types.Transaction, error)
	WaitForTransationReceipt(txhash types.Hash, duration time.Duration) (*types.TransactionReceipt, error)
	GetNextUsableNonce(user types.Address) (nonce *hexutil.Big, err error)

	GetChainIDCached() uint32
	GetNetworkIDCached() uint32
}

type RpcTrace interface {
	GetBlockTraces(blockHash types.Hash) (*types.LocalizedBlockTrace, error)
	FilterTraces(traceFilter types.TraceFilter) (traces []types.LocalizedTrace, err error)
	GetTransactionTraces(txHash types.Hash) (traces []types.LocalizedTrace, err error)
}

type RpcPos interface {
	GetStatus() (postypes.Status, error)
	GetAccount(address postypes.Address, blockNumber ...hexutil.Uint64) (postypes.Account, error)
	GetCommittee(blockNumber ...hexutil.Uint64) (postypes.CommitteeState, error)
	GetBlockByHash(types.Hash) (*postypes.Block, error)
	GetBlockByNumber(blockNumber postypes.BlockNumber) (*postypes.Block, error)
	GetTransactionByNumber(txNumber hexutil.Uint64) (*postypes.Transaction, error)
	GetRewardsByEpoch(epochNumber hexutil.Uint64) (postypes.EpochReward, error)
}

type RpcTxpool interface {
	Status() (val types.TxPoolStatus, err error)
	NextNonce(address types.Address) (val *hexutil.Big, err error)
	TransactionByAddressAndNonce(address types.Address, nonce *hexutil.Big) (val *types.Transaction, err error)
	PendingNonceRange(address types.Address) (val types.TxPoolPendingNonceRange, err error)
	TxWithPoolInfo(hash types.Hash) (val types.TxWithPoolInfo, err error)
	AccountPendingInfo(address types.Address) (val *types.AccountPendingInfo, err error)
	AccountPendingTransactions(address types.Address, maybeStartNonce *hexutil.Big, maybeLimit *hexutil.Uint64) (val types.AccountPendingTransactions, err error)
}

type RpcDebug interface {
	TxpoolGetAccountTransactions(address types.Address) (val []types.Transaction, err error)
	GetEpochReceipts(epoch types.EpochOrBlockHash, include_eth_recepits ...bool) ([][]types.TransactionReceipt, error)
	GetEpochReceiptsByPivotBlockHash(hash types.Hash) (receipts [][]types.TransactionReceipt, err error)
	GetEpochReceiptProofByTransaction(hash types.Hash) (proof *string, err error)
	GetTransactionsByEpoch(epoch types.Epoch) (wrapTransactions []types.WrapTransaction, err error)
	GetTransactionsByBlock(hash types.Hash) (wrapTransactions []types.WrapTransaction, err error)
}

type RpcFilter interface {
	NewFilter(logFilter types.LogFilter) (filterId *rpc.ID, err error)
	NewBlockFilter() (filterId *rpc.ID, err error)
	NewPendingTransactionFilter() (filterId *rpc.ID, err error)
	GetFilterChanges(filterId rpc.ID) (cfxFilterChanges *types.CfxFilterChanges, err error)
	GetFilterLogs(filterID rpc.ID) (logs []types.Log, err error)
	UninstallFilter(filterId rpc.ID) (isUninstalled bool, err error)
}

// AccountManagerOperator is interface of operate actions on account manager
type AccountManagerOperator interface {
	Create(passphrase string) (types.Address, error)
	Import(keyFile, passphrase, newPassphrase string) (types.Address, error)
	ImportKey(keyString string, passphrase string) (types.Address, error)
	Export(address types.Address, passphrase string) (string, error)
	Delete(address types.Address, passphrase string) error
	Update(address types.Address, passphrase, newPassphrase string) error
	List() []types.Address
	GetDefault() (*types.Address, error)
	Unlock(address types.Address, passphrase string) error
	UnlockDefault(passphrase string) error
	TimedUnlock(address types.Address, passphrase string, timeout time.Duration) error
	TimedUnlockDefault(passphrase string, timeout time.Duration) error
	Lock(address types.Address) error
	SignTransaction(tx types.UnsignedTransaction) ([]byte, error)
	SignAndEcodeTransactionWithPassphrase(tx types.UnsignedTransaction, passphrase string) ([]byte, error)
	SignTransactionWithPassphrase(tx types.UnsignedTransaction, passphrase string) (types.SignedTransaction, error)
	Sign(tx types.UnsignedTransaction, passphrase string) (v byte, r, s []byte, err error)
}

// reserve for forward compatbility
type RpcRequester = interfaces.Provider
