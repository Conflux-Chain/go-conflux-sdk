// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package types

import "github.com/ethereum/go-ethereum/common/hexutil"

// TransactionCore represents a transaction without signature in Conflux
type TransactionCore struct {
	Hash             Hash         `json:"hash"`
	Nonce            *hexutil.Big `json:"nonce"`
	BlockHash        *Hash        `json:"blockHash,omitempty"`
	TransactionIndex *hexutil.Big `json:"transactionIndex,omitempty"`
	From             Address      `json:"from"`
	To               *Address     `json:"to,omitempty"`
	Value            *hexutil.Big `json:"value"`
	GasPrice         *hexutil.Big `json:"gasPrice"`
	Gas              *hexutil.Big `json:"gas"`
	ContractCreated  *Address     `json:"contractCreated,omitempty"`
	Data             string       `json:"data"`
	Status           *hexutil.Big `json:"status,omitempty"`
}

// Transaction represents a transaction with signature in Conflux.
type Transaction struct {
	TransactionCore
	V *hexutil.Big `json:"v"`
	R *hexutil.Big `json:"r"`
	S *hexutil.Big `json:"s"`
}

// TransactionReceipt represents the transaction execution result in Conflux.
type TransactionReceipt struct {
	TransactionHash Hash         `json:"transactionHash"`
	Index           uint         `json:"index"`
	BlockHash       Hash         `json:"blockHash"`
	EpochNumber     *uint64      `json:"epochNumber,omitempty"`
	From            Address      `json:"from"`
	To              *Address     `json:"to,omitempty"`
	GasUsed         *hexutil.Big `json:"gasUsed"`
	ContractCreated *Address     `json:"contractCreated,omitempty"`
	Logs            []LogEntry   `json:"logs"`
	LogsBloom       Bloom        `json:"logsBloom"`
	StateRoot       Hash         `json:"stateRoot"`
	OutcomeStatus   uint8        `json:"outcomeStatus"`
}

// CallRequest represents a request to execute contract.
type CallRequest struct {
	From         *Address     `json:"from,omitempty"`
	To           *Address     `json:"to,omitempty"`
	GasPrice     *hexutil.Big `json:"gasPrice,omitempty"`
	Gas          *hexutil.Big `json:"gas,omitempty"`
	Value        *hexutil.Big `json:"value,omitempty"`
	Data         *string      `json:"data,omitempty"`
	Nonce        *hexutil.Big `json:"nonce,omitempty"`
	StorageLimit *hexutil.Big `json:"storage_limit,omitempty"`
}
