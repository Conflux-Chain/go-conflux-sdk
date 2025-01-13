// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package types

import (
	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"io"
	"math/big"

	"github.com/Conflux-Chain/go-conflux-sdk/types/enums"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/pkg/errors"
)

// Transaction represents a transaction with signature in Conflux.
// it is the response from conflux node when sending rpc request, such as cfx_getTransactionByHash
type Transaction struct {
	// Space            *string         `json:"space,omitempty"` //currently it is always "nil", so comment it now and uncomment it if need later
	TransactionType  *hexutil.Uint64 `json:"type,omitempty"`
	Hash             Hash            `json:"hash"`
	Nonce            *hexutil.Big    `json:"nonce"`
	BlockHash        *Hash           `json:"blockHash"`
	TransactionIndex *hexutil.Uint64 `json:"transactionIndex"`
	From             Address         `json:"from"`
	To               *Address        `json:"to"`
	Value            *hexutil.Big    `json:"value"`
	GasPrice         *hexutil.Big    `json:"gasPrice,omitempty"`
	Gas              *hexutil.Big    `json:"gas"`
	ContractCreated  *Address        `json:"contractCreated"`
	Data             string          `json:"data"`
	StorageLimit     *hexutil.Big    `json:"storageLimit"`
	EpochHeight      *hexutil.Big    `json:"epochHeight"`
	ChainID          *hexutil.Big    `json:"chainId"`
	Status           *hexutil.Uint64 `json:"status"`

	AccessList           AccessList   `json:"accessList,omitempty"`
	MaxPriorityFeePerGas *hexutil.Big `json:"maxPriorityFeePerGas,omitempty"`
	MaxFeePerGas         *hexutil.Big `json:"maxFeePerGas,omitempty"`

	//signature
	V       *hexutil.Big    `json:"v"`
	R       *hexutil.Big    `json:"r"`
	S       *hexutil.Big    `json:"s"`
	YParity *hexutil.Uint64 `json:"yParity,omitempty"`
}

// The custom serialization method is designed so that the access list
// is not ignored when it is an array with zero elements.
func (t Transaction) MarshalJSON() ([]byte, error) {
	type Alias Transaction
	if t.AccessList == nil {
		return utils.JsonMarshal((Alias)(t))
	}
	return utils.JsonMarshal(&struct {
		AccessList AccessList `json:"accessList"`
		Alias
	}{
		AccessList: t.AccessList,
		Alias:      (Alias)(t),
	})
}

type legacyTransactionForRlp struct {
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

func (tx *legacyTransactionForRlp) toRaw() *Transaction {
	return &Transaction{
		TransactionType:  NewUint64(uint64(TRANSACTION_TYPE_LEGACY)),
		Hash:             tx.Hash,
		Nonce:            (*hexutil.Big)(tx.Nonce),
		BlockHash:        tx.BlockHash,
		TransactionIndex: tx.TransactionIndex,
		From:             tx.From,
		To:               tx.To,
		Value:            (*hexutil.Big)(tx.Value),
		GasPrice:         (*hexutil.Big)(tx.GasPrice),
		Gas:              (*hexutil.Big)(tx.Gas),
		ContractCreated:  tx.ContractCreated,
		Data:             tx.Data,
		StorageLimit:     (*hexutil.Big)(tx.StorageLimit),
		EpochHeight:      (*hexutil.Big)(tx.EpochHeight),
		ChainID:          (*hexutil.Big)(tx.ChainID),
		Status:           tx.Status,

		V: (*hexutil.Big)(tx.V),
		R: (*hexutil.Big)(tx.R),
		S: (*hexutil.Big)(tx.S),
	}
}

type accessListTransactionForRlp struct {
	TransactionType  *hexutil.Uint64
	Hash             Hash
	Nonce            *big.Int
	BlockHash        *Hash
	TransactionIndex *hexutil.Uint64
	From             Address
	To               *Address `rlp:"nil"`
	Value            *big.Int
	Gas              *big.Int
	GasPrice         *big.Int
	ContractCreated  *Address `rlp:"nil"` // nil means contract creation
	Data             string
	StorageLimit     *big.Int
	EpochHeight      *big.Int
	ChainID          *big.Int
	Status           *hexutil.Uint64

	AccessList AccessList

	//signature
	V       *big.Int
	R       *big.Int
	S       *big.Int
	YParity *hexutil.Uint64
}

func (tx *accessListTransactionForRlp) toRaw() *Transaction {
	return &Transaction{
		TransactionType:  NewUint64(uint64(*tx.TransactionType)),
		Hash:             tx.Hash,
		Nonce:            (*hexutil.Big)(tx.Nonce),
		BlockHash:        tx.BlockHash,
		TransactionIndex: tx.TransactionIndex,
		From:             tx.From,
		To:               tx.To,
		Value:            (*hexutil.Big)(tx.Value),
		GasPrice:         (*hexutil.Big)(tx.GasPrice),
		Gas:              (*hexutil.Big)(tx.Gas),
		ContractCreated:  tx.ContractCreated,
		Data:             tx.Data,
		StorageLimit:     (*hexutil.Big)(tx.StorageLimit),
		EpochHeight:      (*hexutil.Big)(tx.EpochHeight),
		ChainID:          (*hexutil.Big)(tx.ChainID),
		Status:           tx.Status,

		AccessList: tx.AccessList,

		V: (*hexutil.Big)(tx.V),
		R: (*hexutil.Big)(tx.R),
		S: (*hexutil.Big)(tx.S),
	}
}

// dynamicFeeTransactionForRlp transaction struct used for rlp encoding
type dynamicFeeTransactionForRlp struct {
	TransactionType  *hexutil.Uint64
	Hash             Hash
	Nonce            *big.Int
	BlockHash        *Hash
	TransactionIndex *hexutil.Uint64
	From             Address
	To               *Address `rlp:"nil"`
	Value            *big.Int
	Gas              *big.Int
	GasPrice         *big.Int
	ContractCreated  *Address `rlp:"nil"` // nil means contract creation
	Data             string
	StorageLimit     *big.Int
	EpochHeight      *big.Int
	ChainID          *big.Int
	Status           *hexutil.Uint64

	AccessList           AccessList
	MaxPriorityFeePerGas *big.Int
	MaxFeePerGas         *big.Int

	//signature
	V       *big.Int
	R       *big.Int
	S       *big.Int
	YParity *hexutil.Uint64
}

func (tx *dynamicFeeTransactionForRlp) toRaw() *Transaction {
	return &Transaction{
		TransactionType:  NewUint64(uint64(*tx.TransactionType)),
		Hash:             tx.Hash,
		Nonce:            (*hexutil.Big)(tx.Nonce),
		BlockHash:        tx.BlockHash,
		TransactionIndex: tx.TransactionIndex,
		From:             tx.From,
		To:               tx.To,
		Value:            (*hexutil.Big)(tx.Value),
		Gas:              (*hexutil.Big)(tx.Gas),
		GasPrice:         (*hexutil.Big)(tx.GasPrice),
		ContractCreated:  tx.ContractCreated,
		Data:             tx.Data,
		StorageLimit:     (*hexutil.Big)(tx.StorageLimit),
		EpochHeight:      (*hexutil.Big)(tx.EpochHeight),
		ChainID:          (*hexutil.Big)(tx.ChainID),
		Status:           tx.Status,

		AccessList:           tx.AccessList,
		MaxFeePerGas:         (*hexutil.Big)(tx.MaxFeePerGas),
		MaxPriorityFeePerGas: (*hexutil.Big)(tx.MaxPriorityFeePerGas),

		V: (*hexutil.Big)(tx.V),
		R: (*hexutil.Big)(tx.R),
		S: (*hexutil.Big)(tx.S),
	}
}

func (tx *Transaction) toStructForRlp() (interface{}, error) {

	txType := getTransactionType(tx.TransactionType)

	switch txType {
	case TRANSACTION_TYPE_LEGACY:
		return legacyTransactionForRlp{
			Hash:             tx.Hash,
			Nonce:            tx.Nonce.ToInt(),
			BlockHash:        tx.BlockHash,
			TransactionIndex: tx.TransactionIndex,
			From:             tx.From,
			To:               tx.To,
			Value:            tx.Value.ToInt(),
			GasPrice:         tx.GasPrice.ToInt(),
			Gas:              tx.Gas.ToInt(),
			ContractCreated:  tx.ContractCreated,
			Data:             tx.Data,
			StorageLimit:     tx.StorageLimit.ToInt(),
			EpochHeight:      tx.EpochHeight.ToInt(),
			ChainID:          tx.ChainID.ToInt(),
			Status:           tx.Status,
			V:                tx.V.ToInt(),
			R:                tx.R.ToInt(),
			S:                tx.S.ToInt(),
		}, nil
	case TRANSACTION_TYPE_2930:
		inner := accessListTransactionForRlp{
			TransactionType:  tx.TransactionType,
			Hash:             tx.Hash,
			Nonce:            tx.Nonce.ToInt(),
			BlockHash:        tx.BlockHash,
			TransactionIndex: tx.TransactionIndex,
			From:             tx.From,
			To:               tx.To,
			Value:            tx.Value.ToInt(),
			Gas:              tx.Gas.ToInt(),
			GasPrice:         tx.GasPrice.ToInt(),
			ContractCreated:  tx.ContractCreated,
			Data:             tx.Data,
			StorageLimit:     tx.StorageLimit.ToInt(),
			EpochHeight:      tx.EpochHeight.ToInt(),
			ChainID:          tx.ChainID.ToInt(),
			Status:           tx.Status,
			AccessList:       tx.AccessList,
			V:                tx.V.ToInt(),
			R:                tx.R.ToInt(),
			S:                tx.S.ToInt(),
			YParity:          tx.YParity,
		}
		return inner, nil
	case TRANSACTION_TYPE_1559:
		inner := dynamicFeeTransactionForRlp{
			TransactionType:      tx.TransactionType,
			Hash:                 tx.Hash,
			Nonce:                tx.Nonce.ToInt(),
			BlockHash:            tx.BlockHash,
			TransactionIndex:     tx.TransactionIndex,
			From:                 tx.From,
			To:                   tx.To,
			Value:                tx.Value.ToInt(),
			Gas:                  tx.Gas.ToInt(),
			GasPrice:             tx.GasPrice.ToInt(),
			ContractCreated:      tx.ContractCreated,
			Data:                 tx.Data,
			StorageLimit:         tx.StorageLimit.ToInt(),
			EpochHeight:          tx.EpochHeight.ToInt(),
			ChainID:              tx.ChainID.ToInt(),
			Status:               tx.Status,
			AccessList:           tx.AccessList,
			MaxPriorityFeePerGas: tx.MaxFeePerGas.ToInt(),
			MaxFeePerGas:         tx.MaxFeePerGas.ToInt(),
			V:                    tx.V.ToInt(),
			R:                    tx.R.ToInt(),
			S:                    tx.S.ToInt(),
			YParity:              tx.YParity,
		}
		return inner, nil
	default:
		return nil, errors.New("unkown transaction type")
	}
}

// EncodeRLP implements the rlp.Encoder interface.
func (tx Transaction) EncodeRLP(w io.Writer) error {
	txType := getTransactionType(tx.TransactionType)

	txForRlp, err := tx.toStructForRlp()
	if err != nil {
		return err
	}
	bodyRlp, err := rlp.EncodeToBytes(txForRlp)
	if err != nil {
		return err
	}

	if txType == TRANSACTION_TYPE_LEGACY {
		_, err := w.Write(bodyRlp)
		return err
	} else {
		prefixRlp, err := rlp.EncodeToBytes(TRANSACTION_TYPE_PREFIX)
		if err != nil {
			return err
		}

		typeRlp, err := rlp.EncodeToBytes(txType)
		if err != nil {
			return err
		}

		data := append(prefixRlp, typeRlp...)
		data = append(data, bodyRlp...)
		_, err = w.Write(data)
		return err
	}
}

// DecodeRLP implements the rlp.Decoder interface.
func (tx *Transaction) DecodeRLP(r *rlp.Stream) error {

	// legacy
	txForRlp := new(legacyTransactionForRlp)
	err := r.Decode(txForRlp)
	if err == nil {
		*tx = *txForRlp.toRaw()
		return nil
	} else {
		prefix, err := r.Bytes()
		if err != nil {
			return err
		}

		if string(prefix) != string(TRANSACTION_TYPE_PREFIX) {
			return errors.New("invalid transaction")
		}

		t, err := r.Bytes()
		if err != nil {
			return err
		}

		switch TransactionType(t[0]) {
		case TRANSACTION_TYPE_2930:
			txForRlp := new(accessListTransactionForRlp)
			err := r.Decode(&txForRlp)
			if err != nil {
				return err
			}
			*tx = *txForRlp.toRaw()
			return nil

		case TRANSACTION_TYPE_1559:
			txForRlp := new(dynamicFeeTransactionForRlp)
			err := r.Decode(&txForRlp)
			if err != nil {
				return err
			}
			*tx = *txForRlp.toRaw()
			return nil

		default:
			return errors.New("unknown transaction type")
		}
	}
}

// TransactionReceipt represents the transaction execution result in Conflux.
// it is the response from conflux node when sending rpc request, such as cfx_getTransactionReceipt
type TransactionReceipt struct {
	Type               *hexutil.Uint64 `json:"type,omitempty"`
	TransactionHash    Hash            `json:"transactionHash"`
	Index              hexutil.Uint64  `json:"index"`
	BlockHash          Hash            `json:"blockHash"`
	EpochNumber        *hexutil.Uint64 `json:"epochNumber"`
	From               Address         `json:"from"`
	To                 *Address        `json:"to"`
	GasUsed            *hexutil.Big    `json:"gasUsed"`
	AccumulatedGasUsed *hexutil.Big    `json:"accumulatedGasUsed,omitempty"`
	GasFee             *hexutil.Big    `json:"gasFee"`
	EffectiveGasPrice  *hexutil.Big    `json:"effectiveGasPrice"`
	ContractCreated    *Address        `json:"contractCreated"`
	Logs               []Log           `json:"logs"`
	LogsBloom          Bloom           `json:"logsBloom"`
	StateRoot          Hash            `json:"stateRoot"`
	OutcomeStatus      hexutil.Uint64  `json:"outcomeStatus"`
	TxExecErrorMsg     *string         `json:"txExecErrorMsg"`
	// Whether gas costs were covered by the sponsor.
	GasCoveredBySponsor bool `json:"gasCoveredBySponsor"`
	// Whether storage costs were covered by the sponsor.
	StorageCoveredBySponsor bool `json:"storageCoveredBySponsor"`
	// The amount of storage collateralized by the sender.
	StorageCollateralized hexutil.Uint64 `json:"storageCollateralized"`
	// Storage collaterals released during the execution of the transaction.
	StorageReleased []StorageChange `json:"storageReleased"`
	Space           *SpaceType      `json:"space,omitempty"`
	BurntGasFee     *hexutil.Big    `json:"burntGasFee,omitempty"`
}

func (r *TransactionReceipt) GetOutcomeType() (enums.TransactionOutcome, error) {
	switch *r.Space {
	case SPACE_NATIVE:
		outcome := enums.NativeSpaceOutcome(r.OutcomeStatus)
		switch outcome {
		case enums.NATIVE_SPACE_SUCCESS:
			return enums.TRANSACTION_OUTCOME_SUCCESS, nil
		case enums.NATIVE_SPACE_EXCEPTION_WITH_NONCE_BUMPING:
			return enums.TRANSACTION_OUTCOME_FAILURE, nil
		case enums.NATIVE_SPACE_EXCEPTION_WITHOUT_NONCE_BUMPING:
			return enums.TRANSACTION_OUTCOME_SKIPPED, nil
		}
	case SPACE_EVM:
		outcome := enums.EvmSpaceOutcome(r.OutcomeStatus)
		switch outcome {
		case enums.EVM_SPACE_SUCCESS:
			return enums.TRANSACTION_OUTCOME_SUCCESS, nil
		case enums.EVM_SPACE_FAIL:
			return enums.TRANSACTION_OUTCOME_FAILURE, nil
		case enums.EVM_SPACE_SKIPPED:
			return enums.TRANSACTION_OUTCOME_SKIPPED, nil
		}
	default:
		return enums.TransactionOutcome(0xff), errors.New("unknown space")
	}
	return enums.TransactionOutcome(0xff), errors.New("unknown outcome status")
}

func (r *TransactionReceipt) MustGetOutcomeType() enums.TransactionOutcome {
	result, err := r.GetOutcomeType()
	if err != nil {
		panic(err)
	}
	return result
}

func (tx *TransactionReceipt) toStructForRlp() (interface{}, error) {
	_tx := tx
	if tx.Type == nil {
		_tx.Type = NewUint64(uint64(*TRANSACTION_TYPE_LEGACY.Ptr()))
	}

	txType := TransactionType(*_tx.Type)
	switch txType {
	case TRANSACTION_TYPE_LEGACY:
		return legacyTxReceiptForRlp{
			TransactionHash:         tx.TransactionHash,
			Index:                   tx.Index,
			BlockHash:               tx.BlockHash,
			EpochNumber:             tx.EpochNumber,
			From:                    tx.From,
			To:                      tx.To,
			GasUsed:                 tx.GasUsed.ToInt(),
			GasFee:                  tx.GasFee.ToInt(),
			ContractCreated:         tx.ContractCreated,
			Logs:                    tx.Logs,
			LogsBloom:               tx.LogsBloom,
			StateRoot:               tx.StateRoot,
			OutcomeStatus:           tx.OutcomeStatus,
			TxExecErrorMsg:          tx.TxExecErrorMsg,
			GasCoveredBySponsor:     tx.GasCoveredBySponsor,
			StorageCoveredBySponsor: tx.StorageCoveredBySponsor,
			StorageCollateralized:   tx.StorageCollateralized,
			StorageReleased:         tx.StorageReleased,
		}, nil
	case TRANSACTION_TYPE_2930:
		fallthrough
	case TRANSACTION_TYPE_1559:
		inner := txReceiptWithTypeForRlp{
			Type:                    *tx.Type,
			TransactionHash:         tx.TransactionHash,
			Index:                   tx.Index,
			BlockHash:               tx.BlockHash,
			EpochNumber:             tx.EpochNumber,
			From:                    tx.From,
			To:                      tx.To,
			GasUsed:                 tx.GasUsed.ToInt(),
			AccumulatedGasUsed:      tx.AccumulatedGasUsed.ToInt(),
			GasFee:                  tx.GasFee.ToInt(),
			EffectiveGasPrice:       tx.EffectiveGasPrice.ToInt(),
			ContractCreated:         tx.ContractCreated,
			Logs:                    tx.Logs,
			LogsBloom:               tx.LogsBloom,
			StateRoot:               tx.StateRoot,
			OutcomeStatus:           tx.OutcomeStatus,
			TxExecErrorMsg:          tx.TxExecErrorMsg,
			GasCoveredBySponsor:     tx.GasCoveredBySponsor,
			StorageCoveredBySponsor: tx.StorageCoveredBySponsor,
			StorageCollateralized:   tx.StorageCollateralized,
			StorageReleased:         tx.StorageReleased,
			Space:                   tx.Space,
			BurntGasFee:             tx.BurntGasFee.ToInt(),
		}
		return inner, nil
	default:
		return nil, errors.New("unkown transaction type")
	}
}

// StorageChange represents storage change information of the address
type StorageChange struct {
	Address Address `json:"address"`
	/// Number of storage collateral units to deposit / refund (absolute value).
	Collaterals hexutil.Uint64 `json:"collaterals"`
}

// legacyTxReceiptForRlp transaction receipt struct used for rlp encoding
type legacyTxReceiptForRlp struct {
	TransactionHash         Hash
	Index                   hexutil.Uint64
	BlockHash               Hash
	EpochNumber             *hexutil.Uint64
	From                    Address
	To                      *Address `rlp:"nil"`
	GasUsed                 *big.Int
	GasFee                  *big.Int
	ContractCreated         *Address `rlp:"nil"` // nil means contract creation
	Logs                    []Log
	LogsBloom               Bloom
	StateRoot               Hash
	OutcomeStatus           hexutil.Uint64
	TxExecErrorMsg          *string `rlp:"nil"`
	GasCoveredBySponsor     bool
	StorageCoveredBySponsor bool
	StorageCollateralized   hexutil.Uint64
	StorageReleased         []StorageChange
}

func (l *legacyTxReceiptForRlp) toRaw() *TransactionReceipt {
	return &TransactionReceipt{
		Type:                    NewUint64(uint64(TRANSACTION_TYPE_LEGACY)),
		TransactionHash:         l.TransactionHash,
		Index:                   l.Index,
		BlockHash:               l.BlockHash,
		EpochNumber:             l.EpochNumber,
		From:                    l.From,
		To:                      l.To,
		GasUsed:                 (*hexutil.Big)(l.GasUsed),
		GasFee:                  (*hexutil.Big)(l.GasFee),
		ContractCreated:         l.ContractCreated,
		Logs:                    l.Logs,
		LogsBloom:               l.LogsBloom,
		StateRoot:               l.StateRoot,
		OutcomeStatus:           l.OutcomeStatus,
		TxExecErrorMsg:          l.TxExecErrorMsg,
		GasCoveredBySponsor:     l.GasCoveredBySponsor,
		StorageCoveredBySponsor: l.StorageCoveredBySponsor,
		StorageCollateralized:   l.StorageCollateralized,
		StorageReleased:         l.StorageReleased,
	}
}

type txReceiptWithTypeForRlp struct {
	Type                    hexutil.Uint64
	TransactionHash         Hash
	Index                   hexutil.Uint64
	BlockHash               Hash
	EpochNumber             *hexutil.Uint64
	From                    Address
	To                      *Address `rlp:"nil"`
	GasUsed                 *big.Int
	AccumulatedGasUsed      *big.Int
	GasFee                  *big.Int
	EffectiveGasPrice       *big.Int
	ContractCreated         *Address `rlp:"nil"` // nil means contract creation
	Logs                    []Log
	LogsBloom               Bloom
	StateRoot               Hash
	OutcomeStatus           hexutil.Uint64
	TxExecErrorMsg          *string `rlp:"nil"`
	GasCoveredBySponsor     bool
	StorageCoveredBySponsor bool
	StorageCollateralized   hexutil.Uint64
	StorageReleased         []StorageChange
	Space                   *SpaceType
	BurntGasFee             *big.Int
}

func (t *txReceiptWithTypeForRlp) toRaw() *TransactionReceipt {
	return &TransactionReceipt{
		Type:                    &t.Type,
		TransactionHash:         t.TransactionHash,
		Index:                   t.Index,
		BlockHash:               t.BlockHash,
		EpochNumber:             t.EpochNumber,
		From:                    t.From,
		To:                      t.To,
		GasUsed:                 (*hexutil.Big)(t.GasUsed),
		AccumulatedGasUsed:      (*hexutil.Big)(t.AccumulatedGasUsed),
		GasFee:                  (*hexutil.Big)(t.GasFee),
		EffectiveGasPrice:       (*hexutil.Big)(t.EffectiveGasPrice),
		ContractCreated:         t.ContractCreated,
		Logs:                    t.Logs,
		LogsBloom:               t.LogsBloom,
		StateRoot:               t.StateRoot,
		OutcomeStatus:           t.OutcomeStatus,
		TxExecErrorMsg:          t.TxExecErrorMsg,
		GasCoveredBySponsor:     t.GasCoveredBySponsor,
		StorageCoveredBySponsor: t.StorageCoveredBySponsor,
		StorageCollateralized:   t.StorageCollateralized,
		StorageReleased:         t.StorageReleased,
		Space:                   t.Space,
		BurntGasFee:             (*hexutil.Big)(t.BurntGasFee),
	}
}

// EncodeRLP implements the rlp.Encoder interface.
func (tr TransactionReceipt) EncodeRLP(w io.Writer) error {
	txType := getTransactionType(tr.Type)

	txForRlp, err := tr.toStructForRlp()
	if err != nil {
		return err
	}
	bodyRlp, err := rlp.EncodeToBytes(txForRlp)
	if err != nil {
		return err
	}

	if txType == TRANSACTION_TYPE_LEGACY {
		_, err := w.Write(bodyRlp)
		return err
	} else {
		prefixRlp, err := rlp.EncodeToBytes(TRANSACTION_TYPE_PREFIX)
		if err != nil {
			return err
		}

		typeRlp, err := rlp.EncodeToBytes(txType)
		if err != nil {
			return err
		}

		data := append(prefixRlp, typeRlp...)
		data = append(data, bodyRlp...)
		_, err = w.Write(data)
		return err
	}
}

// DecodeRLP implements the rlp.Decoder interface.
func (tr *TransactionReceipt) DecodeRLP(r *rlp.Stream) error {

	// legacy
	trForRlp := new(legacyTxReceiptForRlp)
	err := r.Decode(trForRlp)
	if err == nil {
		*tr = *trForRlp.toRaw()
		return nil
	} else {
		prefix, err := r.Bytes()
		if err != nil {
			return err
		}

		if string(prefix) != string(TRANSACTION_TYPE_PREFIX) {
			return errors.New("invalid rlp data of transaction receipt")
		}

		t, err := r.Bytes()
		if err != nil {
			return err
		}

		switch TransactionType(t[0]) {
		case TRANSACTION_TYPE_2930:
			fallthrough
		case TRANSACTION_TYPE_1559:
			trForRlp := new(txReceiptWithTypeForRlp)
			err := r.Decode(&trForRlp)
			if err != nil {
				return err
			}
			*tr = *trForRlp.toRaw()
			return nil

		default:
			return errors.New("unknown transaction type")
		}
	}
}

type AccountPendingTransactions struct {
	PendingTransactions []Transaction `json:"pendingTransactions"`
	// type maybe string/Pending
	FirstTxStatus *TransactionStatus `json:"firstTxStatus"`
	PendingCount  hexutil.Uint64     `json:"pendingCount"`
}

type PendingReason string

const (
	PENDING_REASON_FUTURE_NONCE     PendingReason = "futureNonce"
	PENDING_REASON_NOT_ENOUGH_CASH  PendingReason = "notEnoughCash"
	PENDING_REASON_OLD_EPOCH_HEIGHT PendingReason = "oldEpochHeight"
	PENDING_REASON_OUTDATED_STATUS  PendingReason = "outdatedStatus"
)

type pending struct {
	PendingReason PendingReason `json:"pending"`
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
		return string(ts.pending.PendingReason)
	}
	return ""
}

func (ts TransactionStatus) MarshalJSON() ([]byte, error) {
	if ts.packedOrReady != "" {
		return utils.JsonMarshal(ts.packedOrReady)
	}
	if (ts.pending != pending{}) {
		return utils.JsonMarshal(ts.pending)
	}
	return []byte{}, nil
}

func (ts *TransactionStatus) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var pendingreason pending
	if err := utils.JsonUnmarshal(data, &pendingreason); err == nil {
		ts.pending = pendingreason
		return nil
	}

	var tmp string
	if err := utils.JsonUnmarshal(data, &tmp); err == nil {
		ts.packedOrReady = tmp
		return nil
	}

	return errors.Errorf("failed to json unmarshal %v to TransactionStatus", string(data))
}

func (ts *TransactionStatus) IsPending() (bool, PendingReason) {
	return ts.pending != pending{}, ts.pending.PendingReason
}

func getTransactionType(raw *hexutil.Uint64) TransactionType {
	if raw == nil {
		return TRANSACTION_TYPE_LEGACY
	}

	return TransactionType(*raw)
}
