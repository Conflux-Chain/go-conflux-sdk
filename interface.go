// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package sdk

import (
	"math/big"
	"net/http"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// HTTPRequester represents a http requester
type HTTPRequester interface {
	Get(url string) (resp *http.Response, err error)
}

// Contractor is interface of contract operator
type Contractor interface {
	GetData(method string, args ...interface{}) (*[]byte, error)
	Call(option *types.ContractMethodCallOption, resultPtr interface{}, method string, args ...interface{}) error
	SendTransaction(option *types.ContractMethodSendOption, method string, args ...interface{}) (*types.Hash, error)
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
	GetTransactionCount(address types.Address, epoch ...*types.Epoch) (*big.Int, error)
	SendRawTransaction(rawData []byte) (types.Hash, error)
	SendTransaction(tx *types.UnsignedTransaction) (types.Hash, error)
	SignEncodedTransactionAndSend(encodedTx []byte, v byte, r, s []byte) (*types.Transaction, error)
	Call(request types.CallRequest, epoch *types.Epoch) (*string, error)
	GetLogs(filter types.LogFilter) ([]types.Log, error)
	GetTransactionByHash(txHash types.Hash) (*types.Transaction, error)
	EstimateGasAndCollateral(request types.CallRequest) (*types.Estimate, error)
	GetBlocksByEpoch(epoch *types.Epoch) ([]types.Hash, error)
	GetTransactionReceipt(txHash types.Hash) (*types.TransactionReceipt, error)
	CreateUnsignedTransaction(from types.Address, to types.Address, amount *hexutil.Big, data *[]byte) (*types.UnsignedTransaction, error)
	ApplyUnsignedTransactionDefault(tx *types.UnsignedTransaction) error
	Debug(method string, args ...interface{}) (interface{}, error)
	Close()
	GetContract(abiJSON string, deployedAt *types.Address) (*Contract, error)
	DeployContract(abiJSON string, bytecode []byte, option *types.ContractDeployOption, timeout time.Duration, callback func(deployedContract Contractor, hash *types.Hash, err error)) <-chan struct{}
}
