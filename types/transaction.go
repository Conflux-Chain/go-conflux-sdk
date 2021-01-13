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
	BlockHash        *Hash           `json:"blockHash"`
	TransactionIndex *hexutil.Uint64 `json:"transactionIndex"`
	From             Address         `json:"from"`
	To               *Address        `json:"to"`
	Value            *hexutil.Big    `json:"value"`
	GasPrice         *hexutil.Big    `json:"gasPrice"`
	Gas              *hexutil.Big    `json:"gas"`
	ContractCreated  *Address        `json:"contractCreated"`
	Data             string          `json:"data"`
	StorageLimit     *hexutil.Big    `json:"storageLimit"`
	EpochHeight      *hexutil.Big    `json:"epochHeight"`
	ChainID          *hexutil.Big    `json:"chainId"`
	Status           *hexutil.Uint64 `json:"status"`

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
	EpochNumber     *hexutil.Uint64 `json:"epochNumber"`
	From            Address         `json:"from"`
	To              *Address        `json:"to"`
	GasUsed         *hexutil.Big    `json:"gasUsed"`
	GasFee          *hexutil.Big    `json:"gasFee"`
	ContractCreated *Address        `json:"contractCreated"`
	Logs            []Log           `json:"logs"`
	LogsBloom       Bloom           `json:"logsBloom"`
	StateRoot       Hash            `json:"stateRoot"`
	OutcomeStatus   hexutil.Uint64  `json:"outcomeStatus"`
	TxExecErrorMsg  *string         `json:"txExecErrorMsg"`
	// Whether gas costs were covered by the sponsor.
	GasCoveredBySponsor bool `json:"gasCoveredBySponsor"`
	// Whether storage costs were covered by the sponsor.
	StorageCoveredBySponsor bool `json:"storageCoveredBySponsor"`
	// The amount of storage collateralized by the sender.
	StorageCollateralized hexutil.Uint64 `json:"storageCollateralized"`
	// Storage collaterals released during the execution of the transaction.
	StorageReleased []StorageChange `json:"storageReleased"`
}

// StorageChange represents storage change information of the address
type StorageChange struct {
	Address Address `json:"address"`
	/// Number of storage collateral units to deposit / refund (absolute value).
	Collaterals hexutil.Uint64 `json:"collaterals"`
}
