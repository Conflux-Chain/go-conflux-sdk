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

// signedTransactionForRlp is a intermediate struct for encoding rlp data
type signedTransactionForRlp struct {
	UnsignedData *unsignedTransactionForRlp
	V            byte
	R            *big.Int
	S            *big.Int
}

// SignedTransaction represents a transaction with signature,
// it is the transaction information for sending transaction.
type SignedTransaction struct {
	UnsignedTransaction UnsignedTransaction
	V                   byte
	R                   hexutil.Bytes
	S                   hexutil.Bytes
}

// Decode decodes RLP encoded data to tx
func (tx *SignedTransaction) Decode(data []byte, networkID uint32) error {
	txForRlp := new(signedTransactionForRlp)
	err := rlp.DecodeBytes(data, txForRlp)
	if err != nil {
		return err
	}

	*tx = *txForRlp.toSignedTransaction(networkID)
	return nil
}

//Encode encodes tx and returns its RLP encoded data
func (tx *SignedTransaction) Encode() ([]byte, error) {
	txForRlp := *tx.toStructForRlp()
	encoded, err := rlp.EncodeToBytes(txForRlp)
	if err != nil {
		return nil, err
	}

	return encoded, nil
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

func (tx *SignedTransaction) toStructForRlp() *signedTransactionForRlp {
	txForRlp := signedTransactionForRlp{
		UnsignedData: tx.UnsignedTransaction.toStructForRlp(),
		V:            tx.V,
		R:            big.NewInt(0).SetBytes(tx.R),
		S:            big.NewInt(0).SetBytes(tx.S),
	}
	return &txForRlp
}

func (tx *signedTransactionForRlp) toSignedTransaction(networkID uint32) *SignedTransaction {
	unsigned := tx.UnsignedData.toUnsignedTransaction(networkID)
	return &SignedTransaction{
		UnsignedTransaction: *unsigned,
		V:                   tx.V,
		R:                   tx.R.Bytes(),
		S:                   tx.S.Bytes(),
	}
}
