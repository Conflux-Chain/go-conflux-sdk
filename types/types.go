// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package types

import (
	"errors"
	"math/big"
	"strings"
	"unicode"

	"github.com/Conflux-Chain/go-conflux-sdk/constants"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// Address represents the 20 byte address of an Conflux account in HEX format.
type Address string

// NewAddress creates an address with specified HEX string.
func NewAddress(hexAddress string) *Address {
	addr := Address(hexAddress)
	return &addr
}

// NewAddressFromCommon creates an address from common.Address
func NewAddressFromCommon(address common.Address) *Address {
	hex := hexutil.Encode(address[:])
	return NewAddress(hex)
}

// String implements the interface stringer
func (address Address) String() string {
	return string(address)
}

// ToCommonAddress converts address to common.Address
func (address Address) ToCommonAddress() *common.Address {
	newAddress := common.HexToAddress(string(address))
	return &newAddress
}

// IsZero returns true if the address is "0x0000000000000000000000000000000000000000"
// or "0X0000000000000000000000000000000000000000", otherwise returns false
func (address Address) IsZero() bool {
	tmp := address.ToCommonAddress()
	return *tmp == constants.ZeroAddress
}

// Checksum returns checksum address
func (address Address) Checksum() Address {
	str := address.String()
	lower := strings.ToLower(str[2:])

	hash := crypto.Keccak256([]byte(lower))
	result := []rune{}

	for i, v := range str[2:] {
		byteIndex := i / 2
		valueOfHashChar := byte(0)

		if i%2 == 0 {
			valueOfHashChar = hash[byteIndex] >> 4
		} else {
			valueOfHashChar = hash[byteIndex] & 0x0f
		}

		if valueOfHashChar >= 8 {
			v = unicode.ToUpper(v)
		}

		result = append(result, v)
	}
	return "0x" + Address(result)
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
func (address Address) GetAddressType() (AddressType, error) {

	if !common.IsHexAddress(string(address)) {
		return InvalidAddress, errors.New("address should be start with 0x prefix and length should be 42")
	}

	// if not lowercase and not equal to checksum address, invalid
	if strings.ToLower(string(address)) != address.String() {
		if address != address.Checksum() {
			return InvalidAddress, errors.New("address checksum fail")
		}
	}

	addrBytes := address.ToCommonAddress()

	flag := addrBytes[0] >> 4
	if flag == 0x1 {
		return NormalAddress, nil
	}
	if flag == 0x8 {
		return CustomContractAddress, nil
	}

	if addrBytes[0] == 0x08 && addrBytes[1] == 0x88 {
		return InternalContractAddress, nil
	}

	return InvalidAddress, errors.New("address should be start with 0x1, 0x8 or 0x088")
}

// IsValid return true if address is valid
func (address Address) IsValid() bool {
	addreeType, _ := address.GetAddressType()
	return addreeType != InvalidAddress
}

// func (address *Address) Checksum() Address {
// 	bytes := address.ToCommonAddress().Bytes()
// 	hash := address.ToCommonAddress().Hash().Bytes()
// 	for i := 0; i < len(bytes); i++ {
// 		hash[i]>=8?"":""
// 	}
// }

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
