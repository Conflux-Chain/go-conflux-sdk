package types

import (
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// ContractDeployOption for setting option when deploying contract
type ContractDeployOption struct {
	UnsignedTransactionBase
	// TimeoutInSecond represents the timeout of deploy contract,
	// default value is 0 which means never timeout
	Timeout time.Duration
}

// ContractMethodCallOption for setting option when call contract method
type ContractMethodCallOption struct {
	From         *Address
	Nonce        *hexutil.Big
	GasPrice     *hexutil.Big
	Gas          *hexutil.Big
	Value        *hexutil.Big
	StorageLimit *hexutil.Uint64
	ChainID      *hexutil.Big
	Epoch        *Epoch
}

// ContractMethodSendOption for setting option when call contract method
type ContractMethodSendOption UnsignedTransactionBase

// CallRequest represents a request to execute contract.
type CallRequest struct {
	From     *Address     `json:"from,omitempty"`
	To       *Address     `json:"to,omitempty"`
	GasPrice *hexutil.Big `json:"gasPrice,omitempty"`
	Gas      *hexutil.Big `json:"gas,omitempty"`
	Value    *hexutil.Big `json:"value,omitempty"`
	// NOTE, cannot use *hexutil.Bytes or hexutil.Bytes here.
	// Otherwise, hexutil.Bytes.UnmarshalJSON will called to
	// unmarshal from nil and cause to errNonString error.
	Data         *string         `json:"data,omitempty"`
	Nonce        *hexutil.Big    `json:"nonce,omitempty"`
	StorageLimit *hexutil.Uint64 `json:"storageLimit,omitempty"`
}

// FillByUnsignedTx fills CallRequest fields by tx
func (request *CallRequest) FillByUnsignedTx(tx *UnsignedTransaction) {
	if tx != nil {
		request.From = tx.From
		request.To = tx.To
		request.GasPrice = tx.GasPrice
		request.Value = tx.Value
		request.StorageLimit = tx.StorageLimit

		if tx.Gas != nil {
			request.Gas = tx.Gas
		}

		hexData := tx.Data.String()
		request.Data = &hexData

		if tx.Nonce != nil {
			request.Nonce = tx.Nonce
		}
	}
}

// FillByCallOption fills CallRequest fields by
func (request *CallRequest) FillByCallOption(option *ContractMethodCallOption) {
	if option != nil {
		request.From = option.From
		request.GasPrice = option.GasPrice
		request.Gas = option.Gas
		request.Value = option.Value
		request.Nonce = option.Nonce
		request.StorageLimit = option.StorageLimit
	}
}
