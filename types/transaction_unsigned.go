// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package types

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

// UnsignedTransaction represents a transaction without signature,
// it is the transaction information for sending transaction.
type UnsignedTransaction struct {
	From         *Address
	To           *Address
	Nonce        uint64
	GasPrice     *hexutil.Big
	Gas          uint64
	Value        *hexutil.Big
	Data         []byte
	StorageLimit *hexutil.Big
	EpochHeight  *hexutil.Big
	ChainID      uint64
}

// unsignedTransactionForRlp is a intermediate struct for encoding rlp data
type unsignedTransactionForRlp struct {
	Nonce        *big.Int
	GasPrice     *big.Int
	Gas          *big.Int
	To           *common.Address
	Value        *big.Int
	StorageLimit *big.Int
	EpochHeight  *big.Int
	ChainID      *big.Int
	Data         []byte
}

// DefaultGas is the default gas in a transaction to transfer amount without any data.
const defaultGas uint64 = 21000

// DefaultGasPrice is the default gas price.
var defaultGasPrice *hexutil.Big = NewBigInt(10000000000) // 10G drip

// ApplyDefault apply default fields if these filed are empty
func (tx *UnsignedTransaction) ApplyDefault() {
	if tx.GasPrice == nil {
		tx.GasPrice = defaultGasPrice
	}

	if tx.Gas == 0 {
		tx.Gas = defaultGas
	}

	if tx.Value == nil {
		tx.Value = NewBigInt(0)
	}
}

// Hash return transaction Hash
func (tx *UnsignedTransaction) Hash() ([]byte, error) {
	// data := tx.getRlpNeedFields()
	encoded, err := tx.Encode()
	if err != nil {
		msg := fmt.Sprintf("encode tx {%+v} error", tx)
		return nil, WrapError(err, msg)
	}

	return crypto.Keccak256(encoded), nil
}

//Encode encode unsigned transaction and return its RLP encoded data
func (tx *UnsignedTransaction) Encode() ([]byte, error) {
	// data := tx.getRlpNeedFields()
	data := *tx.toStructForRlp()
	encoded, err := rlp.EncodeToBytes(data)
	if err != nil {
		msg := fmt.Sprintf("encode data {%+v} to bytes error", data)
		return nil, WrapError(err, msg)
	}
	return encoded, nil
}

// EncodeWithSignature encode unsigned transaction with signature and return its RLP encoded data
func (tx *UnsignedTransaction) EncodeWithSignature(v byte, r, s []byte) ([]byte, error) {
	signedTx := signedTransactionForRlp{
		UnsignedData: tx.toStructForRlp(),
		V:            v,
		R:            r,
		S:            s,
	}

	encoded, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		msg := fmt.Sprintf("encode data {%+v} to bytes error", signedTx)
		return nil, WrapError(err, msg)
	}

	return encoded, nil
}

// Decode decode RLP encoded data to tx
func (tx *UnsignedTransaction) Decode(data []byte) error {
	utxForRlp := new(unsignedTransactionForRlp)
	err := rlp.DecodeBytes(data, utxForRlp)
	if err != nil {
		msg := fmt.Sprintf("decode data {%+v} to rlp error", data)
		return WrapError(err, msg)
	}

	*tx = *utxForRlp.toUnsignedTransaction()
	return nil
}

func (tx *UnsignedTransaction) toStructForRlp() *unsignedTransactionForRlp {
	var to *common.Address
	if tx.To != nil {
		addr := common.HexToAddress(string(*tx.To))
		to = &addr
	}

	return &unsignedTransactionForRlp{
		Nonce:        new(big.Int).SetUint64(tx.Nonce),
		GasPrice:     tx.GasPrice.ToInt(),
		Gas:          new(big.Int).SetUint64(tx.Gas),
		To:           to,
		Value:        tx.Value.ToInt(),
		StorageLimit: tx.StorageLimit.ToInt(),
		EpochHeight:  tx.EpochHeight.ToInt(),
		ChainID:      new(big.Int).SetUint64(tx.ChainID),
		Data:         tx.Data,
	}
}

func (tx *unsignedTransactionForRlp) toUnsignedTransaction() *UnsignedTransaction {
	to := Address(strings.ToLower(tx.To.Hex()))
	gasPrice := hexutil.Big(*tx.GasPrice)
	value := hexutil.Big(*tx.Value)
	storageLimit := hexutil.Big(*tx.StorageLimit)
	epochHeight := hexutil.Big(*tx.EpochHeight)

	return &UnsignedTransaction{
		From:         nil,
		To:           &to,
		Nonce:        tx.Nonce.Uint64(),
		GasPrice:     &gasPrice,
		Gas:          tx.Gas.Uint64(),
		Value:        &value,
		Data:         tx.Data,
		StorageLimit: &storageLimit,
		EpochHeight:  &epochHeight,
		ChainID:      tx.ChainID.Uint64(),
	}
}
