package interfaces

import (
	"math/big"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/middleware"
	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// RpcCaller is interface of Client
type RpcCaller interface {
	RpcCallerCore
	Cfx() RpcCfxCaller
	Debug() RpcDebugCaller
	Pos() RpcPosCaller
	Pubsub() RpcPubsubCaller
	Trace() RpcTraceCaller
	GetNetworkID() (uint32, error)
}

type SignableRpcCaller interface {
	RpcCaller

	NewTransaction(from types.Address, to types.Address, amount *hexutil.Big, data []byte) (types.UnsignedTransaction, error)
	PopulateTransaction(tx *types.UnsignedTransaction) error
	SignTransactionAndSend(tx types.UnsignedTransaction) (types.Hash, error)

	GetWallet() Wallet
}

type RpcCallerHelper interface {
	NewAddress(base32OrHex string) (types.Address, error)
	MustNewAddress(base32OrHex string) types.Address

	EncodedTransactionAndSend(encodedTx []byte, v byte, r, s []byte) (*types.Transaction, error)

	WaitForTransationBePacked(txhash types.Hash, duration time.Duration) (*types.Transaction, error)
	WaitForTransationReceipt(txhash types.Hash, duration time.Duration) (*types.TransactionReceipt, error)
}

type RpcCallerCore interface {
	GetNodeURL() string

	CallRPC(result interface{}, method string, args ...interface{}) error
	BatchCallRPC(b []rpc.BatchElem) error

	UseCallRpcMiddleware(middleware middleware.CallRpcMiddleware)
	UseBatchCallRpcMiddleware(middleware middleware.BatchCallRpcMiddleware)

	Close()
}

// missing
// cfx_getPoSEconomics
//
type RpcCfxCaller interface {
	GetGasPrice() (*hexutil.Big, error)
	GetEpochNumber(epoch ...*types.Epoch) (*hexutil.Big, error)
	GetBalance(address types.Address, epoch ...*types.Epoch) (*hexutil.Big, error)
	GetAdmin(contractAddress types.Address, epoch ...*types.Epoch) (admin *types.Address, err error)
	GetSponsorInfo(contractAddress types.Address, epoch ...*types.Epoch) (sponsor types.SponsorInfo, err error)
	GetStakingBalance(account types.Address, epoch ...*types.Epoch) (balance *hexutil.Big, err error)
	GetDepositList(address types.Address, epoch ...*types.Epoch) ([]types.DepositInfo, error)
	GetVoteList(address types.Address, epoch ...*types.Epoch) ([]types.VoteStakeInfo, error)
	GetCollateralForStorage(account types.Address, epoch ...*types.Epoch) (storage *hexutil.Big, err error)
	GetCode(address types.Address, epoch ...*types.Epoch) (hexutil.Bytes, error)
	GetStorageAt(address types.Address, position types.Hash, epoch ...*types.Epoch) (storageEntries hexutil.Bytes, err error)
	GetStorageRoot(address types.Address, epoch ...*types.Epoch) (storageRoot *types.StorageRoot, err error)
	GetBlockByHash(blockHash types.Hash) (*types.Block, error)
	GetBlockByHashWithPivotAssumption(blockHash types.Hash, pivotHash types.Hash, epoch hexutil.Uint64) (block types.Block, err error)
	GetBlockByEpoch(epoch *types.Epoch) (*types.Block, error)
	GetBlockByBlockNumber(blockNumer hexutil.Uint64) (block *types.Block, err error)
	GetBestBlockHash() (types.Hash, error)
	GetNextNonce(address types.Address, epoch ...*types.Epoch) (*hexutil.Big, error)

	Call(request types.CallRequest, epoch *types.Epoch) (hexutil.Bytes, error)
	GetLogs(filter types.LogFilter) ([]types.Log, error)
	GetTransactionByHash(txHash types.Hash) (*types.Transaction, error)
	GetAccountPendingInfo(address types.Address) (pendignInfo *types.AccountPendingInfo, err error)
	GetAccountPendingTransactions(address types.Address, startNonce *hexutil.Big, limit *hexutil.Uint64) (pendingTxs types.AccountPendingTransactions, err error)
	EstimateGasAndCollateral(request types.CallRequest, epoch ...*types.Epoch) (types.Estimate, error)
	CheckBalanceAgainstTransaction(accountAddress types.Address,
		contractAddress types.Address,
		gasLimit *hexutil.Big,
		gasPrice *hexutil.Big,
		storageLimit *hexutil.Big,
		epoch ...*types.Epoch) (response types.CheckBalanceAgainstTransactionResponse, err error)
	GetBlocksByEpoch(epoch *types.Epoch) ([]types.Hash, error)
	GetSkippedBlocksByEpoch(epoch *types.Epoch) (blockHashs []types.Hash, err error)
	GetTransactionReceipt(txHash types.Hash) (*types.TransactionReceipt, error)
	GetAccountInfo(account types.Address, epoch ...*types.Epoch) (accountInfo types.AccountInfo, err error)
	GetInterestRate(epoch ...*types.Epoch) (intersetRate *hexutil.Big, err error)
	GetAccumulateInterestRate(epoch ...*types.Epoch) (intersetRate *hexutil.Big, err error)
	GetRawBlockConfirmationRisk(blockhash types.Hash) (*hexutil.Big, error)
	GetBlockConfirmationRisk(blockHash types.Hash) (*big.Float, error)
	GetStatus() (types.Status, error)
	GetBlockRewardInfo(epoch types.Epoch) (rewardInfo []types.RewardInfo, err error)
	GetClientVersion() (clientVersion string, err error)
	GetSupplyInfo(epoch ...*types.Epoch) (info types.TokenSupplyInfo, err error)

	GetBlockSummaryByHash(blockHash types.Hash) (*types.BlockSummary, error)
	GetBlockSummaryByEpoch(epoch *types.Epoch) (*types.BlockSummary, error)
	GetBlockSummaryByBlockNumber(blockNumer hexutil.Uint64) (block *types.BlockSummary, err error)

	SendTransaction(tx types.SignedTransaction) (types.Hash, error)
	SendRawTransaction(rawData []byte) (types.Hash, error)
}

// missing
// txpool_status
// tx_inspect_pending
// tx_inspect
// txpool_inspect
// txpool_content
type RpcDebugCaller interface {
	GetEpochReceipts(epoch types.Epoch) (receipts [][]types.TransactionReceipt, err error)
	GetEpochReceiptsByPivotBlockHash(hash types.Hash) (receipts [][]types.TransactionReceipt, err error)
}

type RpcPosCaller interface {
	// pos_getStatus Status
	// pos_getAccount Account
	// pos_getCommittee CommitteeState
	// pos_getBlockByHash Block
}

type RpcPubsubCaller interface {
	SubscribeNewHeads(channel chan types.BlockHeader) (*rpc.ClientSubscription, error)
	SubscribeEpochs(channel chan types.WebsocketEpochResponse, subscriptionEpochType ...types.Epoch) (*rpc.ClientSubscription, error)
	SubscribeLogs(channel chan types.SubscriptionLog, filter types.LogFilter) (*rpc.ClientSubscription, error)
}

type RpcTestCaller interface {
}

type RpcTraceCaller interface {
	GetBlockTraces(blockHash types.Hash) (*types.LocalizedBlockTrace, error)
	FilterTraces(traceFilter types.TraceFilter) (traces []types.LocalizedTrace, err error)
	GetTransactionTraces(txHash types.Hash) (traces []types.LocalizedTrace, err error)
}
