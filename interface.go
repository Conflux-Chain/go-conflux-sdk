// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package sdk

import (
	"math/big"
	"net/http"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	// rpc "github.com/ethereum/go-ethereum/rpc"
)

// HTTPRequester is interface for emitting a http requester
type HTTPRequester interface {
	Get(url string) (resp *http.Response, err error)
}

// Contractor is interface of contract operator
type Contractor interface {
	GetData(method string, args ...interface{}) ([]byte, error)
	Call(option *types.ContractMethodCallOption, resultPtr interface{}, method string, args ...interface{}) error
	SendTransaction(option *types.ContractMethodSendOption, method string, args ...interface{}) (*types.Hash, error)
	DecodeEvent(out interface{}, event string, log types.LogEntry) error
}

// ClientOperator is interface of operate actions on client
type ClientOperator interface {
	GetGasPrice() (*big.Int, error)
	GetEpochNumber(epoch ...*types.Epoch) (*big.Int, error)
	GetBalance(address types.Address, epoch ...*types.Epoch) (*big.Int, error)
	GetCode(address types.Address, epoch ...*types.Epoch) (string, error)
	GetBlockSummaryByHash(blockHash types.Hash) (*types.BlockSummary, error)
	GetBlockByHash(blockHash types.Hash) (*types.Block, error)
	GetBlockSummaryByEpoch(epoch *types.Epoch) (*types.BlockSummary, error)
	GetBlockByEpoch(epoch *types.Epoch) (*types.Block, error)
	GetBestBlockHash() (types.Hash, error)
	GetRawBlockConfirmRisk(blockhash types.Hash) (*big.Int, error)
	GetBlockConfirmRisk(blockHash types.Hash) (*big.Float, error)
	SendRawTransaction(rawData []byte) (types.Hash, error)
	SendTransaction(tx *types.UnsignedTransaction) (types.Hash, error)
	SetAccountManager(accountManager AccountManagerOperator)
	SignEncodedTransactionAndSend(encodedTx []byte, v byte, r, s []byte) (*types.Transaction, error)
	Call(request types.CallRequest, epoch *types.Epoch) (*string, error)
	CallRPC(result interface{}, method string, args ...interface{}) error
	BatchCallRPC(b []rpc.BatchElem) error
	GetLogs(filter types.LogFilter) ([]types.Log, error)
	GetTransactionByHash(txHash types.Hash) (*types.Transaction, error)
	EstimateGasAndCollateral(request types.CallRequest) (*types.Estimate, error)
	GetBlocksByEpoch(epoch *types.Epoch) ([]types.Hash, error)
	GetTransactionReceipt(txHash types.Hash) (*types.TransactionReceipt, error)
	CreateUnsignedTransaction(from types.Address, to types.Address, amount *hexutil.Big, data []byte) (*types.UnsignedTransaction, error)
	ApplyUnsignedTransactionDefault(tx *types.UnsignedTransaction) error
	Debug(method string, args ...interface{}) (interface{}, error)
	Close()
	GetContract(abiJSON []byte, deployedAt *types.Address) (*Contract, error)
	// DeployContract(abiJSON string, bytecode []byte, option *types.ContractDeployOption, timeout time.Duration, callback func(deployedContract Contractor, hash *types.Hash, err error)) <-chan struct{}
	DeployContract(option *types.ContractDeployOption, abiJSON []byte,
		bytecode []byte, constroctorParams ...interface{}) *ContractDeployResult

	BatchGetTxByHashs(txhashs []types.Hash) (map[types.Hash]*types.Transaction, error)
	BatchGetBlockConfirmationRisk(blockhashs []types.Hash) (map[types.Hash]*big.Float, error)
	BatchGetRawBlockConfirmationRisk(blockhashs []types.Hash) (map[types.Hash]*big.Int, error)
	BatchGetBlockSummarys(blockhashs []types.Hash) (map[types.Hash]*types.BlockSummary, error)
	GetNodeURL() string
}

// AccountManagerOperator is interface of operate actions on account manager
type AccountManagerOperator interface {
	Create(passphrase string) (types.Address, error)
	Import(keyFile, passphrase, newPassphrase string) (types.Address, error)
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
	SignTransactionWithPassphrase(tx types.UnsignedTransaction, passphrase string) (*types.SignedTransaction, error)
	Sign(tx types.UnsignedTransaction, passphrase string) (v byte, r, s []byte, err error)
}

type rpcRequester interface {
	Call(resultPtr interface{}, method string, args ...interface{}) error
	BatchCall(b []rpc.BatchElem) error
	Close()
}
