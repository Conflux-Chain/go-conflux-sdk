// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package types

import (
	"math/big"

	address "github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Address = address.Address

// Hash represents the 32 byte Keccak256 hash of arbitrary data in HEX format.
type Hash string

// ToCommonHash converts hash to common.Hash
func (hash Hash) ToCommonHash() *common.Hash {
	newHash := common.HexToHash(string(hash))
	return &newHash
}

// String implements the interface stringer
func (hash Hash) String() string {
	return string(hash)
}

// Bloom is a hash type with 256 bytes.
type Bloom string

// NewBigInt creates a big number with specified uint64 value.
func NewBigInt(x uint64) *hexutil.Big {
	n1 := new(big.Int).SetUint64(x)
	n2 := hexutil.Big(*n1)
	return &n2
}

// NewBigIntByRaw creates a hexutil.big with specified big.int value.
func NewBigIntByRaw(x *big.Int) *hexutil.Big {
	if x == nil {
		return nil
	}
	v := hexutil.Big(*x)
	return &v
}

// NewUint64 creates a hexutil.Uint64 with specified uint64 value.
func NewUint64(x uint64) *hexutil.Uint64 {
	n1 := hexutil.Uint64(x)
	return &n1
}

// NewUint creates a hexutil.Uint with specified uint value.
func NewUint(x uint) *hexutil.Uint {
	n1 := hexutil.Uint(x)
	return &n1
}

// NewBytes creates a hexutil.Bytes with specified input value.
func NewBytes(input []byte) hexutil.Bytes {
	return hexutil.Bytes(input)
}
