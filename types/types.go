// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Address represents the 20 byte address of an Conflux account in HEX format.
type Address string

// NewAddress creates a address with specified HEX string.
func NewAddress(hexAddress string) *Address {
	addr := Address(hexAddress)
	return &addr
}

// Hash represents the 32 byte Keccak256 hash of arbitrary data in HEX format.
type Hash string

// Bloom is a hash type with 256 bytes.
type Bloom string

// NewBigInt creates a big number with specified int64 value.
func NewBigInt(x int64) *hexutil.Big {
	n1 := big.NewInt(x)
	n2 := hexutil.Big(*n1)
	return &n2
}
