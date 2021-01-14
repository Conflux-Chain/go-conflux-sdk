// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package utils

import (
	"encoding/hex"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

// PublicKeyToAddress generate address from public key
//
// Account address in conflux starts with '0x1'
func PublicKeyToAddress(publicKey string) types.Address {

	if Has0xPrefix(publicKey) {
		publicKey = publicKey[2:]
	}

	pubKeyBytes, err := hex.DecodeString(publicKey)
	if err != nil {
		panic(err)
	}

	pubKeyBytes = append([]byte{0x04}, pubKeyBytes...)

	pub, err := crypto.UnmarshalPubkey(pubKeyBytes)
	if err != nil {
		panic(err)
	}

	addr := crypto.PubkeyToAddress(*pub)
	cfxaddr := ToCfxGeneralAddress(addr)
	return cfxaddr
}

// PrivateKeyToPublicKey calculates public key from private key
func PrivateKeyToPublicKey(privateKey string) string {
	if Has0xPrefix(privateKey) {
		privateKey = privateKey[2:]
	}

	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		panic(err)
	}

	pubKeyBytes := crypto.FromECDSAPub(&key.PublicKey)
	pubKey := hexutil.Encode(pubKeyBytes[1:])

	return pubKey
}

// Keccak256 hashes hex string by keccak256 and returns it's hash value
func Keccak256(hexStr string) (string, error) {
	if hexStr[0:2] != "0x" {
		return "", errors.New("input must start with 0x")
	}

	bytes, err := hex.DecodeString(hexStr[2:])
	if err != nil {
		return "", err
	}

	hash := crypto.Keccak256(bytes)
	return "0x" + hex.EncodeToString(hash), nil
}

// ToCfxGeneralAddress converts a normal address to conflux customerd general address
// whose hex string starts with '0x1'
func ToCfxGeneralAddress(address common.Address) types.Address {
	address[0] = (address[0] & 0x0f) | 0x10
	return types.Address(hexutil.Encode(address.Bytes()))
}
