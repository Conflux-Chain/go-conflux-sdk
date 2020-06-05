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

// UnsignedTransactionBase represents a transaction without To, Data and signature
type UnsignedTransactionBase struct {
	From         *Address
	Nonce        *hexutil.Big
	GasPrice     *hexutil.Big
	Gas          *hexutil.Big
	Value        *hexutil.Big
	StorageLimit *hexutil.Big
	EpochHeight  *hexutil.Big
	ChainID      *hexutil.Big
}

// UnsignedTransaction represents a transaction without signature,
// it is the transaction information for sending transaction.
type UnsignedTransaction struct {
	UnsignedTransactionBase
	To   *Address
	Data hexutil.Bytes
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
// const defaultGas uint64 = 21000
var defaultGas *hexutil.Big = NewBigInt(21000)

// DefaultGasPrice is the default gas price.
var defaultGasPrice *hexutil.Big = NewBigInt(10000000000) // 10G drip

// ApplyDefault applys default value for these fields if they are empty
func (tx *UnsignedTransaction) ApplyDefault() {
	if tx.GasPrice == nil {
		tx.GasPrice = defaultGasPrice
	}

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
		msg := fmt.Sprintf("encode tx {%+v} error", tx)
		return nil, WrapError(err, msg)
	}

	return crypto.Keccak256(encoded), nil
}

//Encode encodes tx and returns its RLP encoded data
func (tx *UnsignedTransaction) Encode() ([]byte, error) {
	data := *tx.toStructForRlp()
	encoded, err := rlp.EncodeToBytes(data)
	if err != nil {
		msg := fmt.Sprintf("encode data {%+v} to bytes error", data)
		return nil, WrapError(err, msg)
	}
	return encoded, nil
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
func (tx *UnsignedTransaction) Decode(data []byte) error {
	utxForRlp := new(unsignedTransactionForRlp)
	err := rlp.DecodeBytes(data, utxForRlp)
	if err != nil {
		msg := fmt.Sprintf("decode data {%+x} to rlp error", data)
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
		Nonce:        tx.Nonce.ToInt(),
		GasPrice:     tx.GasPrice.ToInt(),
		Gas:          tx.Gas.ToInt(),
		To:           to,
		Value:        tx.Value.ToInt(),
		StorageLimit: tx.StorageLimit.ToInt(),
		EpochHeight:  tx.EpochHeight.ToInt(),
		ChainID:      tx.ChainID.ToInt(),
		Data:         tx.Data,
	}
}

func (tx *unsignedTransactionForRlp) toUnsignedTransaction() *UnsignedTransaction {
	to := Address(strings.ToLower(tx.To.Hex()))
	gasPrice := hexutil.Big(*tx.GasPrice)
	value := hexutil.Big(*tx.Value)
	storageLimit := hexutil.Big(*tx.StorageLimit)
	epochHeight := hexutil.Big(*tx.EpochHeight)

	nonce := hexutil.Big(*tx.Nonce)
	gas := hexutil.Big(*tx.Gas)
	chainid := hexutil.Big(*tx.ChainID)
	return &UnsignedTransaction{
		UnsignedTransactionBase: UnsignedTransactionBase{
			From:         nil,
			Nonce:        &nonce,
			GasPrice:     &gasPrice,
			Gas:          &gas,
			Value:        &value,
			StorageLimit: &storageLimit,
			EpochHeight:  &epochHeight,
			ChainID:      &chainid,
		},
		To:   &to,
		Data: tx.Data,
	}
}
