// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package types

import (
	"math/big"
	"slices"

	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/pkg/errors"
)

type TransactionType uint

func (t TransactionType) Ptr() *TransactionType {
	return &t
}

const (
	TRANSACTION_TYPE_LEGACY TransactionType = iota
	TRANSACTION_TYPE_2930
	TRANSACTION_TYPE_1559
)

var (
	TRANSACTION_TYPE_PREFIX = []byte("cfx")
)

// UnsignedTransactionBase represents a transaction without To, Data and signature
type UnsignedTransactionBase struct {
	From         *Address
	Nonce        *hexutil.Big
	GasPrice     *hexutil.Big
	Gas          *hexutil.Big
	Value        *hexutil.Big
	StorageLimit *hexutil.Uint64
	EpochHeight  *hexutil.Uint64
	ChainID      *hexutil.Uint

	AccessList           AccessList
	MaxPriorityFeePerGas *hexutil.Big
	MaxFeePerGas         *hexutil.Big

	Type *TransactionType
}

// // get the txtype according to feeData
// // - if if has maxFeePerGas or maxPriorityFeePerGas, then set txtype to 2
// // - else if contains accesslist, set txtype to 1
// // - else set txtype to 0
// func (utx UnsignedTransactionBase) GetType() TransactionType {
// 	if utx.Type != nil {
// 		return *utx.Type
// 	}

// 	if utx.MaxPriorityFeePerGas != nil || utx.MaxFeePerGas != nil {
// 		return TRANSACTION_TYPE_LEGACY
// 	}

// 	if utx.AccessList != nil {
// 		return TRANSACTION_TYPE_2930
// 	}

// 	return TRANSACTION_TYPE_LEGACY
// }

// UnsignedTransaction represents a transaction without signature,
// it is the transaction information for sending transaction.
type UnsignedTransaction struct {
	UnsignedTransactionBase
	To   *Address
	Data hexutil.Bytes
}

// unsignedLegacyTransactionForRlp is a intermediate struct for encoding rlp data
type unsignedLegacyTransactionForRlp struct {
	Nonce        *big.Int        `rlp:"nil"`
	GasPrice     *big.Int        `rlp:"nil"`
	Gas          *big.Int        `rlp:"nil"`
	To           *common.Address `rlp:"nil"`
	Value        *big.Int        `rlp:"nil"`
	StorageLimit *hexutil.Uint64 `rlp:"nil"`
	EpochHeight  *hexutil.Uint64 `rlp:"nil"`
	ChainID      *hexutil.Uint   `rlp:"nil"`
	Data         []byte
}

type unsigned2930TransactionForRlp struct {
	Nonce        *big.Int        `rlp:"nil"`
	GasPrice     *big.Int        `rlp:"nil"`
	Gas          *big.Int        `rlp:"nil"`
	To           *common.Address `rlp:"nil"`
	Value        *big.Int        `rlp:"nil"`
	StorageLimit *hexutil.Uint64 `rlp:"nil"`
	EpochHeight  *hexutil.Uint64 `rlp:"nil"`
	ChainID      *hexutil.Uint   `rlp:"nil"`
	Data         []byte
	AccessList   etypes.AccessList
}

type unsigned1559TransactionForRlp struct {
	Nonce                *big.Int        `rlp:"nil"`
	MaxPriorityFeePerGas *big.Int        `rlp:"nil"`
	MaxFeePerGas         *big.Int        `rlp:"nil"`
	Gas                  *big.Int        `rlp:"nil"`
	To                   *common.Address `rlp:"nil"`
	Value                *big.Int        `rlp:"nil"`
	StorageLimit         *hexutil.Uint64 `rlp:"nil"`
	EpochHeight          *hexutil.Uint64 `rlp:"nil"`
	ChainID              *hexutil.Uint   `rlp:"nil"`
	Data                 []byte
	AccessList           etypes.AccessList
}

// DefaultGas is the default gas in a transaction to transfer amount without any data.
// const defaultGas uint64 = 21000
var defaultGas *hexutil.Big = NewBigInt(21000)

// DefaultGasPrice is the default gas price.
// var defaultGasPrice *hexutil.Big = NewBigInt(10000000000) // 10G drip

// ApplyDefault applys default value for these fields if they are empty
func (tx *UnsignedTransaction) ApplyDefault() {
	// if tx.GasPrice == nil {
	// 	tx.GasPrice = defaultGasPrice
	// }

	if tx.Gas == nil {
		tx.Gas = defaultGas
	}

	if tx.Value == nil {
		tx.Value = NewBigInt(0)
	}
}

// Hash hashes the tx by keccak256 and returns the result
func (tx *UnsignedTransaction) Hash() ([]byte, error) {
	encoded, err := tx.Encode()
	if err != nil {
		return nil, err
	}

	return crypto.Keccak256(encoded), nil
}

// Encode encodes tx and returns its RLP encoded data
func (tx *UnsignedTransaction) Encode() ([]byte, error) {
	_tx := *tx
	if tx.Type == nil {
		_tx.Type = TRANSACTION_TYPE_LEGACY.Ptr()
	}

	data, err := tx.toStructForRlp()
	if err != nil {
		return nil, err
	}
	bodyRlp, err := rlp.EncodeToBytes(data)
	if err != nil {
		return nil, err
	}

	if *_tx.Type == TRANSACTION_TYPE_LEGACY {
		return bodyRlp, nil
	} else {
		data := append([]byte{}, TRANSACTION_TYPE_PREFIX...)
		data = append(data, byte(*_tx.Type))
		data = append(data, bodyRlp...)
		return data, nil
	}
}

// EncodeWithSignature encodes tx with signature and return its RLP encoded data
func (tx *UnsignedTransaction) EncodeWithSignature(v byte, r, s []byte) ([]byte, error) {
	signedTx := new(SignedTransaction)
	signedTx.UnsignedTransaction = *tx
	signedTx.V = v
	signedTx.R = r
	signedTx.S = s
	return signedTx.Encode()
}

// Decode decodes RLP encoded data to tx
func (tx *UnsignedTransaction) Decode(data []byte, networkID uint32) error {
	if !slices.Equal(data[:3], TRANSACTION_TYPE_PREFIX) {
		utxForRlp := new(unsignedLegacyTransactionForRlp)
		err := rlp.DecodeBytes(data, utxForRlp)
		if err != nil {
			return err
		}

		*tx = *utxForRlp.toUnsignedTransaction(networkID)
		return nil
	}
	txType := TransactionType(data[3])
	switch txType {
	case TRANSACTION_TYPE_2930:
		txForRlp := new(unsigned2930TransactionForRlp)
		err := rlp.DecodeBytes(data[4:], txForRlp)
		if err != nil {
			return err
		}

		_tx := txForRlp.toUnsignedTransaction(networkID)
		if err != nil {
			return err
		}
		*tx = *_tx
		return nil

	case TRANSACTION_TYPE_1559:
		txForRlp := new(unsigned1559TransactionForRlp)
		err := rlp.DecodeBytes(data[4:], txForRlp)
		if err != nil {
			return err
		}

		_tx := txForRlp.toUnsignedTransaction(networkID)
		if err != nil {
			return err
		}
		*tx = *_tx
		return nil

	default:
		return errors.Errorf("unknown transaction type %d", txType)
	}
}

func (tx *UnsignedTransaction) toStructForRlp() (interface{}, error) {
	txType := tx.Type
	if txType == nil {
		txType = TRANSACTION_TYPE_LEGACY.Ptr()
	}

	var to *common.Address
	if tx.To != nil {
		addr := tx.To.MustGetCommonAddress()
		to = &addr
	}

	switch *txType {
	case TRANSACTION_TYPE_LEGACY:
		return unsignedLegacyTransactionForRlp{
			Nonce:        tx.Nonce.ToInt(),
			GasPrice:     tx.GasPrice.ToInt(),
			Gas:          tx.Gas.ToInt(),
			To:           to,
			Value:        tx.Value.ToInt(),
			StorageLimit: tx.StorageLimit,
			EpochHeight:  tx.EpochHeight,
			ChainID:      tx.ChainID,
			Data:         tx.Data,
		}, nil
	case TRANSACTION_TYPE_2930:
		return unsigned2930TransactionForRlp{
			Nonce:        tx.Nonce.ToInt(),
			GasPrice:     tx.GasPrice.ToInt(),
			Gas:          tx.Gas.ToInt(),
			To:           to,
			Value:        tx.Value.ToInt(),
			StorageLimit: tx.StorageLimit,
			EpochHeight:  tx.EpochHeight,
			ChainID:      tx.ChainID,
			Data:         tx.Data,
			AccessList:   tx.AccessList.ToEthType(),
		}, nil
	case TRANSACTION_TYPE_1559:
		return unsigned1559TransactionForRlp{
			Nonce:                tx.Nonce.ToInt(),
			MaxPriorityFeePerGas: tx.MaxFeePerGas.ToInt(),
			MaxFeePerGas:         tx.MaxFeePerGas.ToInt(),
			Gas:                  tx.Gas.ToInt(),
			To:                   to,
			Value:                tx.Value.ToInt(),
			StorageLimit:         tx.StorageLimit,
			EpochHeight:          tx.EpochHeight,
			ChainID:              tx.ChainID,
			Data:                 tx.Data,
			AccessList:           tx.AccessList.ToEthType(),
		}, nil
	default:
		return nil, errors.New("unkown transaction type")
	}
}

func (tx *unsignedLegacyTransactionForRlp) toUnsignedTransaction(networkID uint32) *UnsignedTransaction {
	var to *cfxaddress.Address
	if tx.To != nil {
		_to := cfxaddress.MustNewFromCommon(*tx.To, networkID)
		to = &_to
	}

	return &UnsignedTransaction{
		UnsignedTransactionBase: UnsignedTransactionBase{
			From:         nil,
			Nonce:        (*hexutil.Big)(tx.Nonce),
			GasPrice:     (*hexutil.Big)(tx.GasPrice),
			Gas:          (*hexutil.Big)(tx.Gas),
			Value:        (*hexutil.Big)(tx.Value),
			StorageLimit: tx.StorageLimit,
			EpochHeight:  tx.EpochHeight,
			ChainID:      tx.ChainID,
		},
		To:   to,
		Data: tx.Data,
	}
}

func (tx *unsigned2930TransactionForRlp) toUnsignedTransaction(networkID uint32) *UnsignedTransaction {
	var to *cfxaddress.Address
	if tx.To != nil {
		_to := cfxaddress.MustNewFromCommon(*tx.To, networkID)
		to = &_to
	}

	accessList := ConvertEthAccessListToCfx(tx.AccessList, networkID)

	return &UnsignedTransaction{
		UnsignedTransactionBase: UnsignedTransactionBase{
			From:         nil,
			Nonce:        (*hexutil.Big)(tx.Nonce),
			GasPrice:     (*hexutil.Big)(tx.GasPrice),
			Gas:          (*hexutil.Big)(tx.Gas),
			Value:        (*hexutil.Big)(tx.Value),
			StorageLimit: tx.StorageLimit,
			EpochHeight:  tx.EpochHeight,
			ChainID:      tx.ChainID,
			AccessList:   accessList,
			Type:         TRANSACTION_TYPE_2930.Ptr(),
		},
		To:   to,
		Data: tx.Data,
	}
}

func (tx *unsigned1559TransactionForRlp) toUnsignedTransaction(networkID uint32) *UnsignedTransaction {
	var to *cfxaddress.Address
	if tx.To != nil {
		_to := cfxaddress.MustNewFromCommon(*tx.To, networkID)
		to = &_to
	}

	accessList := ConvertEthAccessListToCfx(tx.AccessList, networkID)

	return &UnsignedTransaction{
		UnsignedTransactionBase: UnsignedTransactionBase{
			From:                 nil,
			Nonce:                (*hexutil.Big)(tx.Nonce),
			MaxPriorityFeePerGas: (*hexutil.Big)(tx.MaxPriorityFeePerGas),
			MaxFeePerGas:         (*hexutil.Big)(tx.MaxFeePerGas),
			Gas:                  (*hexutil.Big)(tx.Gas),
			Value:                (*hexutil.Big)(tx.Value),
			StorageLimit:         tx.StorageLimit,
			EpochHeight:          tx.EpochHeight,
			ChainID:              tx.ChainID,
			AccessList:           accessList,
			Type:                 TRANSACTION_TYPE_1559.Ptr(),
		},
		To:   to,
		Data: tx.Data,
	}
}

func toUnsignedTransaction(unsignedTxForRlp interface{}, networkID uint32) (*UnsignedTransaction, error) {
	switch val := unsignedTxForRlp.(type) {
	case unsignedLegacyTransactionForRlp:
		return val.toUnsignedTransaction(networkID), nil
	case *unsignedLegacyTransactionForRlp:
		return val.toUnsignedTransaction(networkID), nil
	case unsigned2930TransactionForRlp:
		return val.toUnsignedTransaction(networkID), nil
	case *unsigned2930TransactionForRlp:
		return val.toUnsignedTransaction(networkID), nil
	case unsigned1559TransactionForRlp:
		return val.toUnsignedTransaction(networkID), nil
	case *unsigned1559TransactionForRlp:
		return val.toUnsignedTransaction(networkID), nil
	}
	return nil, errors.New("unkown transaction type")
}
