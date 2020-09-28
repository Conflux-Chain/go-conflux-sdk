// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package types

import (
	"math/big"

	"github.com/Conflux-Chain/go-conflux-sdk/constants"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Address represents the 20 byte address of an Conflux account in HEX format.
type Address string

// NewAddress creates a address with specified HEX string.
func NewAddress(hexAddress string) *Address {
	addr := Address(hexAddress)
	return &addr
}

// String implements the interface stringer
func (address *Address) String() string {
	return string(*address)
}

// ToCommonAddress converts address to common.Address
func (address *Address) ToCommonAddress() *common.Address {
	newAddress := common.HexToAddress(string(*address))
	return &newAddress
}

// IsZero returns true if the address is "0x0000000000000000000000000000000000000000"
// or "0X0000000000000000000000000000000000000000", otherwise returns false
func (address *Address) IsZero() bool {
	tmp := address.ToCommonAddress()
	return *tmp == constants.ZeroAddress
}

type AddressType string

const (
	NormalAddress           AddressType = "Normal"
	CustomContractAddress   AddressType = "CustomContract"
	InternalContractAddress AddressType = "InternalContract"
	InvalidAddress          AddressType = "Invalid"
)

// GetAddressType returuns the address type,
// address with prefix "0x1" is normal address and "0x8" is contract address
func (address *Address) GetAddressType() AddressType {
	if address == nil {
		return InvalidAddress
	}

	addrBytes := address.ToCommonAddress().Bytes()

	flag := addrBytes[0] >> 4
	if flag == 0x1 {
		return NormalAddress
	}
	if flag == 0x8 {
		return CustomContractAddress
	}

	if addrBytes[0] == 0x08 && addrBytes[1] == 0x88 {
		return InternalContractAddress
	}

	return InvalidAddress
}

// Hash represents the 32 byte Keccak256 hash of arbitrary data in HEX format.
type Hash string

// ToCommonHash converts hash to common.Hash
func (hash *Hash) ToCommonHash() *common.Hash {
	newHash := common.HexToHash(string(*hash))
	return &newHash
}

// String implements the interface stringer
func (hash *Hash) String() string {
	return string(*hash)
}

// Bloom is a hash type with 256 bytes.
type Bloom string

// NewBigInt creates a big number with specified int64 value.
func NewBigInt(x int64) *hexutil.Big {
	n1 := big.NewInt(x)
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
