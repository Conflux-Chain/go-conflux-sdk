// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package types

type signedTransactionForRlp struct {
	UnsignedData *unsignedTransactionForRlp
	V            byte
	R            []byte
	S            []byte
}

// SignedTransaction represents a transaction with signature,
// it is the transaction information for sending transaction.
type SignedTransaction struct {
	UnsignedTransaction
	V byte
	R []byte
	S []byte
}
