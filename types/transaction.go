// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package types

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Transaction represents a transaction with signature in Conflux.
// it is the response from conflux node when sending rpc request, such as cfx_getTransactionByHash
type Transaction struct {
	Hash             Hash            `json:"hash"`
	Nonce            *hexutil.Big    `json:"nonce"`
	BlockHash        *Hash           `json:"blockHash,omitempty"`
	TransactionIndex *hexutil.Uint64 `json:"transactionIndex,omitempty"`
	From             Address         `json:"from"`
	To               *Address        `json:"to,omitempty"`
	Value            *hexutil.Big    `json:"value"`
	GasPrice         *hexutil.Big    `json:"gasPrice"`
	Gas              *hexutil.Big    `json:"gas"`
	ContractCreated  *Address        `json:"contractCreated,omitempty"`
	Data             string          `json:"data"`
	StorageLimit     *hexutil.Big    `json:"storageLimit,omitempty"`
	EpochHeight      *hexutil.Big    `json:"epochHeight,omitempty"`
	ChainID          *hexutil.Big    `json:"chainId,omitempty"`
	Status           *hexutil.Uint64 `json:"status,omitempty"`

	//signature
	V *hexutil.Big `json:"v"`
	R *hexutil.Big `json:"r"`
	S *hexutil.Big `json:"s"`
}

// TransactionReceipt represents the transaction execution result in Conflux.
// it is the response from conflux node when sending rpc request, such as cfx_getTransactionReceipt
type TransactionReceipt struct {
	TransactionHash Hash            `json:"transactionHash"`
	Index           hexutil.Uint64  `json:"index"`
	BlockHash       Hash            `json:"blockHash"`
	EpochNumber     *hexutil.Uint64 `json:"epochNumber,omitempty"`
	From            Address         `json:"from"`
	To              *Address        `json:"to,omitempty"`
	GasUsed         *hexutil.Big    `json:"gasUsed"`
	GasFee          *hexutil.Big    `json:"gasFee"`
	ContractCreated *Address        `json:"contractCreated,omitempty"`
	Logs            []LogEntry      `json:"logs"`
	LogsBloom       Bloom           `json:"logsBloom"`
	StateRoot       Hash            `json:"stateRoot"`
	OutcomeStatus   hexutil.Uint64  `json:"outcomeStatus"`
}
