// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package types

import (
	"math/big"

	"github.com/Conflux-Chain/go-conflux-sdk/utils/addressutil"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/pkg/errors"
)

// SignedTransaction represents a transaction with signature,
// it is the transaction information for sending transaction.
type SignedTransaction struct {
	UnsignedTransaction UnsignedTransaction
	V                   byte
	R                   hexutil.Bytes
	S                   hexutil.Bytes
}

// signedTransactionForRlp is a intermediate struct for encoding rlp data
type signedTransactionForRlp struct {
	UnsignedData interface{}
	V            byte
	R            *big.Int
	S            *big.Int
}

// geneic format of type signedTransactionForRlp
type gSignedTransactionForRlp[T any] struct {
	UnsignedData T
	V            byte
	R            *big.Int
	S            *big.Int
}

// convert from gSignedTransactionForRlp to signedTransactionForRlp
func (g gSignedTransactionForRlp[T]) toRaw() *signedTransactionForRlp {
	return &signedTransactionForRlp{
		UnsignedData: g.UnsignedData,
		V:            g.V,
		R:            g.R,
		S:            g.S,
	}
}

// Decode decodes RLP encoded data to tx
func (tx *SignedTransaction) Decode(data []byte, networkID uint32) error {
	if len(data) < 4 {
		return errors.New("data should not be less than 4")
	}

	if string(data[:3]) != string(TRANSACTION_TYPE_PREFIX) {
		txForRlp := new(gSignedTransactionForRlp[unsignedLegacyTransactionForRlp])
		err := rlp.DecodeBytes(data, txForRlp)
		if err != nil {
			return err
		}

		_tx, err := txForRlp.toRaw().toSignedTransaction(networkID)
		if err != nil {
			return err
		}
		*tx = *_tx
		return nil
	}

	txType := TransactionType(data[3])
	switch txType {
	case TRANSACTION_TYPE_2930:
		txForRlp := new(gSignedTransactionForRlp[unsigned2930TransactionForRlp])
		err := rlp.DecodeBytes(data[4:], txForRlp)
		if err != nil {
			return err
		}

		_tx, err := txForRlp.toRaw().toSignedTransaction(networkID)
		if err != nil {
			return err
		}
		*tx = *_tx
		return nil

	case TRANSACTION_TYPE_1559:
		txForRlp := new(gSignedTransactionForRlp[unsigned1559TransactionForRlp])
		err := rlp.DecodeBytes(data[4:], txForRlp)
		if err != nil {
			return err
		}

		_tx, err := txForRlp.toRaw().toSignedTransaction(networkID)
		if err != nil {
			return err
		}
		*tx = *_tx
		return nil

	default:
		return errors.Errorf("unknown transaction type %d", txType)
	}
}

// Encode encodes tx and returns its RLP encoded data; Encode format is "cfx||tx_type||body_rlp"
func (tx *SignedTransaction) Encode() ([]byte, error) {
	_tx := *tx
	if tx.UnsignedTransaction.Type == nil {
		_tx.UnsignedTransaction.Type = TRANSACTION_TYPE_LEGACY.Ptr()
	}

	txForRlp, err := _tx.toStructForRlp()
	if err != nil {
		return nil, err
	}

	bodyRlp, err := rlp.EncodeToBytes(txForRlp)
	if err != nil {
		return nil, err
	}

	if *_tx.UnsignedTransaction.Type == TRANSACTION_TYPE_LEGACY {
		return bodyRlp, nil
	} else {
		data := append([]byte{}, TRANSACTION_TYPE_PREFIX...)
		data = append(data, byte(*_tx.UnsignedTransaction.Type))
		data = append(data, bodyRlp...)
		return data, nil
	}
}

// Hash calculates the hash of the transaction rlp encode result
func (tx *SignedTransaction) Hash() ([]byte, error) {
	encoded, err := tx.Encode()
	if err != nil {
		return nil, err
	}

	return crypto.Keccak256(encoded), nil
}

// Sender recovers the sender from a signed transaction
func (tx *SignedTransaction) Sender(networkId uint32) (Address, error) {
	hash, err := tx.UnsignedTransaction.Hash()
	if err != nil {
		return Address{}, errors.WithStack(err)
	}

	pub, err := crypto.Ecrecover(hash, tx.Signature())
	if err != nil {
		return Address{}, errors.WithStack(err)
	}

	pubStr := (hexutil.Bytes(pub)).String()
	return addressutil.PubkeyToAddress(pubStr, networkId)
}

// Signature returns the signature of the transaction
func (tx *SignedTransaction) Signature() []byte {
	sig := []byte(tx.R)
	sig = append(sig, tx.S...)
	sig = append(sig, tx.V)
	return sig
}

func (tx *SignedTransaction) toStructForRlp() (*signedTransactionForRlp, error) {
	untxForRlp, err := tx.UnsignedTransaction.toStructForRlp()
	if err != nil {
		return nil, err
	}

	txForRlp := signedTransactionForRlp{
		UnsignedData: untxForRlp,
		V:            tx.V,
		R:            big.NewInt(0).SetBytes(tx.R),
		S:            big.NewInt(0).SetBytes(tx.S),
	}
	return &txForRlp, nil
}

func (tx *signedTransactionForRlp) toSignedTransaction(networkID uint32) (*SignedTransaction, error) {
	unsigned, err := toUnsignedTransaction(tx.UnsignedData, networkID)
	if err != nil {
		return nil, err
	}
	return &SignedTransaction{
		UnsignedTransaction: *unsigned,
		V:                   tx.V,
		R:                   tx.R.Bytes(),
		S:                   tx.S.Bytes(),
	}, nil
}
