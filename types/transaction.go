// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package types

import (
	"encoding/json"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/pkg/errors"
)

// Transaction represents a transaction with signature in Conflux.
// it is the response from conflux node when sending rpc request, such as cfx_getTransactionByHash
type Transaction struct {
	// Space            *string          `json:"space,omitempty"` //currently it is always "nil", so comment it now and uncomment it if need later
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

// rlpEncodableTransaction transaction struct used for rlp encoding
type rlpEncodableTransaction struct {
	Hash             Hash
	Nonce            *big.Int
	BlockHash        *Hash
	TransactionIndex *hexutil.Uint64
	From             Address
	To               *Address `rlp:"nil"`
	Value            *big.Int
	GasPrice         *big.Int
	Gas              *big.Int
	ContractCreated  *Address `rlp:"nil"` // nil means contract creation
	Data             string
	StorageLimit     *big.Int
	EpochHeight      *big.Int
	ChainID          *big.Int
	Status           *hexutil.Uint64

	//signature
	V *big.Int
	R *big.Int
	S *big.Int
}

// EncodeRLP implements the rlp.Encoder interface.
func (tx Transaction) EncodeRLP(w io.Writer) error {
	rtx := rlpEncodableTransaction{
		tx.Hash, tx.Nonce.ToInt(), tx.BlockHash, tx.TransactionIndex, tx.From, tx.To,
		tx.Value.ToInt(), tx.GasPrice.ToInt(), tx.Gas.ToInt(), tx.ContractCreated, tx.Data,
		tx.StorageLimit.ToInt(), tx.EpochHeight.ToInt(), tx.ChainID.ToInt(), tx.Status,
		tx.V.ToInt(), tx.R.ToInt(), tx.S.ToInt(),
	}

	return rlp.Encode(w, rtx)
}

// DecodeRLP implements the rlp.Decoder interface.
func (tx *Transaction) DecodeRLP(r *rlp.Stream) error {
	var rtx rlpEncodableTransaction
	if err := r.Decode(&rtx); err != nil {
		return err
	}

	tx.Hash, tx.Nonce, tx.BlockHash = rtx.Hash, (*hexutil.Big)(rtx.Nonce), rtx.BlockHash
	tx.TransactionIndex, tx.From, tx.To = rtx.TransactionIndex, rtx.From, rtx.To
	tx.Value, tx.GasPrice = (*hexutil.Big)(rtx.Value), (*hexutil.Big)(rtx.GasPrice)
	tx.Gas, tx.ContractCreated, tx.Data = (*hexutil.Big)(rtx.Gas), rtx.ContractCreated, rtx.Data
	tx.StorageLimit, tx.EpochHeight = (*hexutil.Big)(rtx.StorageLimit), (*hexutil.Big)(rtx.EpochHeight)
	tx.ChainID, tx.Status, tx.V = (*hexutil.Big)(rtx.ChainID), rtx.Status, (*hexutil.Big)(rtx.V)
	tx.R, tx.S = (*hexutil.Big)(rtx.R), (*hexutil.Big)(rtx.S)

	return nil
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

// rlpEncodableTransactionReceipt transaction receipt struct used for rlp encoding
type rlpEncodableTransactionReceipt struct {
	TransactionHash Hash
	Index           hexutil.Uint64
	BlockHash       Hash
	EpochNumber     *hexutil.Uint64
	From            Address
	To              *Address `rlp:"nil"`
	GasUsed         *big.Int
	GasFee          *big.Int
	ContractCreated *Address `rlp:"nil"` // nil means contract creation
	Logs            []Log
	LogsBloom       Bloom
	StateRoot       Hash
	OutcomeStatus   hexutil.Uint64
	TxExecErrorMsg  *string `rlp:"nil"`
	// Whether gas costs were covered by the sponsor.
	GasCoveredBySponsor bool
	// Whether storage costs were covered by the sponsor.
	StorageCoveredBySponsor bool
	// The amount of storage collateralized by the sender.
	StorageCollateralized hexutil.Uint64
	// Storage collaterals released during the execution of the transaction.
	StorageReleased []StorageChange
}

// EncodeRLP implements the rlp.Encoder interface.
func (tr TransactionReceipt) EncodeRLP(w io.Writer) error {
	rtx := rlpEncodableTransactionReceipt{
		tr.TransactionHash, tr.Index, tr.BlockHash, tr.EpochNumber, tr.From, tr.To,
		tr.GasUsed.ToInt(), tr.GasFee.ToInt(), tr.ContractCreated, tr.Logs, tr.LogsBloom,
		tr.StateRoot, tr.OutcomeStatus, tr.TxExecErrorMsg, tr.GasCoveredBySponsor,
		tr.StorageCoveredBySponsor, tr.StorageCollateralized, tr.StorageReleased,
	}

	return rlp.Encode(w, rtx)
}

// DecodeRLP implements the rlp.Decoder interface.
func (tr *TransactionReceipt) DecodeRLP(r *rlp.Stream) error {
	var rtr rlpEncodableTransactionReceipt
	if err := r.Decode(&rtr); err != nil {
		return err
	}

	tr.TransactionHash, tr.Index, tr.BlockHash = rtr.TransactionHash, rtr.Index, rtr.BlockHash
	tr.EpochNumber, tr.From, tr.To = rtr.EpochNumber, rtr.From, rtr.To
	tr.GasUsed, tr.GasFee = (*hexutil.Big)(rtr.GasUsed), (*hexutil.Big)(rtr.GasFee)
	tr.ContractCreated, tr.Logs, tr.LogsBloom = rtr.ContractCreated, rtr.Logs, rtr.LogsBloom
	tr.StateRoot, tr.OutcomeStatus, tr.TxExecErrorMsg = rtr.StateRoot, rtr.OutcomeStatus, rtr.TxExecErrorMsg
	tr.GasCoveredBySponsor, tr.StorageCoveredBySponsor = rtr.GasCoveredBySponsor, rtr.StorageCoveredBySponsor
	tr.StorageCollateralized, tr.StorageReleased = rtr.StorageCollateralized, rtr.StorageReleased

	return nil
}

type AccountPendingTransactions struct {
	PendingTransactions []Transaction `json:"pendingTransactions"`
	// type maybe string/Pending
	FirstTxStatus *TransactionStatus `json:"firstTxStatus"`
	PendingCount  hexutil.Uint64     `json:"pendingCount"`
}

type pending struct {
	PendingReason string `json:"pending"`
}

type TransactionStatus struct {
	packedOrReady string
	pending       pending
}

func (ts TransactionStatus) String() string {
	if ts.packedOrReady != "" {
		return ts.packedOrReady
	}
	if (ts.pending != pending{}) {
		return ts.pending.PendingReason
	}
	return ""
}

func (ts TransactionStatus) MarshalJSON() ([]byte, error) {
	if ts.packedOrReady != "" {
		return json.Marshal(ts.packedOrReady)
	}
	if (ts.pending != pending{}) {
		return json.Marshal(ts.pending)
	}
	return []byte{}, nil
}

func (ts *TransactionStatus) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var pendingreason pending
	if err := json.Unmarshal(data, &pendingreason); err == nil {
		ts.pending = pendingreason
		return nil
	}

	var tmp string
	if err := json.Unmarshal(data, &tmp); err == nil {
		ts.packedOrReady = tmp
		return nil
	}

	return errors.Errorf("failed to json unmarshal %v to TransactionStatus", string(data))
}

func (ts *TransactionStatus) IsPending() (bool, string) {
	return ts.pending != pending{}, ts.pending.PendingReason
}
