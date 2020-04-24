package types

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// ContractDeployOption for setting option when deploying contract
type ContractDeployOption UnsignedTransactionBase

// ContractMethodCallOption for setting option when call contract method
type ContractMethodCallOption UnsignedTransactionBase

// ContractMethodSendOption for setting option when call contract method
type ContractMethodSendOption UnsignedTransactionBase

// CallRequest represents a request to execute contract.
type CallRequest struct {
	From         *Address     `json:"from,omitempty"`
	To           *Address     `json:"to,omitempty"`
	GasPrice     *hexutil.Big `json:"gasPrice,omitempty"`
	Gas          *hexutil.Big `json:"gas,omitempty"`
	Value        *hexutil.Big `json:"value,omitempty"`
	Data         string       `json:"data,omitempty"`
	Nonce        *hexutil.Big `json:"nonce,omitempty"`
	StorageLimit *hexutil.Big `json:"storage_limit,omitempty"`
}

// FillByUnsignedTx fills CallRequest fields by tx
func (cq *CallRequest) FillByUnsignedTx(tx *UnsignedTransaction) {
	cq.From = tx.From
	cq.To = tx.To
	cq.GasPrice = tx.GasPrice
	cq.Value = tx.Value
	cq.StorageLimit = tx.StorageLimit

	if tx.Gas != 0 {
		cq.Gas = NewBigInt(int64(tx.Gas))
	}

	_data := "0x" + hex.EncodeToString(tx.Data)
	cq.Data = _data

	if tx.Nonce != 0 {
		cq.Nonce = NewBigInt(int64(tx.Nonce))
	}

}
