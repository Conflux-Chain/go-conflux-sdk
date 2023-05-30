package bind

import (
	"errors"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	rpc "github.com/openweb3/go-rpc-provider"
)

var (
	// ErrNoCode is returned by call and transact operations for which the requested
	// recipient contract to operate on does not exist in the state db or does not
	// have any code associated with it (i.e. suicided).
	ErrNoCode = errors.New("no contract code at given address")

	// This error is raised when attempting to perform a pending state action
	// on a backend that doesn't implement PendingContractCaller.
	ErrNoPendingState = errors.New("backend does not support pending state")
)

// ContractCaller defines the methods needed to allow operating with a contract on a read
// only basis.
type ContractCaller interface {
	// CodeAt returns the code of the given account. This is needed to differentiate
	// between contract internal errors and the local chain being out of sync.
	GetCode(address types.Address, epoch ...*types.Epoch) (code hexutil.Bytes, err error)
	// ContractCall executes an Ethereum contract call with the specified data as the
	// input.
	Call(call types.CallRequest, epoch *types.Epoch) (result hexutil.Bytes, err error)
}

// ContractTransactor defines the methods needed to allow operating with a contract
// on a write only basis. Besides the transacting method, the remainder are helpers
// used when the user does not provide some needed values, but rather leaves it up
// to the transactor to decide.
type ContractTransactor interface {

	// EstimateGasAndCollateral tries to estimate the gas and storage collateral
	// needed to execute a specific transaction based on the current pending state
	// of the backend blockchain. There is no guarantee that this is the true gas
	// limit requirement as other transactions may be added or removed by miners,
	// but it should provide a basis for setting a reasonable default.
	EstimateGasAndCollateral(request types.CallRequest, epoch ...*types.Epoch) (estimat types.Estimate, err error)
	// SendTransaction injects the transaction into the pending pool for execution.
	SendTransaction(tx types.UnsignedTransaction) (types.Hash, error)

	ApplyUnsignedTransactionDefault(tx *types.UnsignedTransaction) error
}

// ContractFilterer defines the methods needed to access log events using one-off
// queries or continuous event subscriptions.
type ContractFilterer interface {
	// FilterLogs executes a log filter operation, blocking during execution and
	// returning all the results in one batch.
	GetLogs(filter types.LogFilter) (logs []types.Log, err error)

	// SubscribeLogs creates a background log filtering operation, returning
	// a subscription immediately, which can be used to stream the found events.
	SubscribeLogs(channel chan types.SubscriptionLog, filter types.LogFilter) (*rpc.ClientSubscription, error)
}

// DeployBackend wraps the operations needed by WaitMined and WaitDeployed.
type DeployBackend interface {
	TransactionReceipt(txHash common.Hash) (*types.TransactionReceipt, error)
	GetCode(address types.Address, epoch ...*types.Epoch) (code hexutil.Bytes, err error)
}

// ContractBackend defines the methods needed to work with contracts on a read-write basis.
type ContractBackend interface {
	ContractCaller
	ContractTransactor
	ContractFilterer
}
