// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package sdk

import (
	"math/big"
	"net/http"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// HTTPRequester represents a http requester
type HTTPRequester interface {
	Get(url string) (resp *http.Response, err error)
}

// Contractor ...
type Contractor interface {
	GetData(method string, args ...interface{}) (*[]byte, error)
	Call(callRequest types.CallRequest, method string, args ...interface{}) (interface{}, error)
	SendTransaction(callRequest types.CallRequest, method string, args ...interface{}) (*types.Hash, error)
}

// ClientOperator ...
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
	SignEncodedTransactionAndSend(encodedTx []byte, v byte, r, s []byte) (*types.Transaction, error)
	Call(request types.CallRequest, epoch ...*types.Epoch) (string, error)
	GetLogs(filter types.LogFilter) ([]types.Log, error)
	GetTransactionByHash(txHash types.Hash) (*types.Transaction, error)
	EstimateGas(request types.CallRequest, epoch ...*types.Epoch) (*big.Int, error)
	EstimateGasAndCollateral(request types.CallRequest) (*types.Estimate, error)
	GetBlocksByEpoch(epoch *types.Epoch) ([]types.Hash, error)
	GetTransactionReceipt(txHash types.Hash) (*types.TransactionReceipt, error)
	CreateUnsignedTransaction(from types.Address, to types.Address, amount *hexutil.Big, data *[]byte) (*types.UnsignedTransaction, error)
	Debug(method string, args ...interface{}) (interface{}, error)
	Close()
	NewContract(abiJSON string, address types.Address) (*Contract, error)
}
