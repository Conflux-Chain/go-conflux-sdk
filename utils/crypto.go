package utils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"math/big"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// PublicKeyToAddress generate address from public key
//
// Account address hex starts with '0x1'
func PublicKeyToAddress(publicKey string) types.Address {
	pubKey := new(big.Int)
	_, ok := pubKey.SetString(publicKey, 0)
	if !ok {
		panic("publicKey is invalid")
	}

	// _publicKey := hexutil.MustDecodeBig(publicKey).Bytes()
	val := crypto.Keccak256(pubKey.Bytes())[12:]
	val[0] = (val[0] & 0x0f) | 0x10
	return types.Address(hexutil.Encode(val))
}

// PrivateKeyToPublicKey calculate public key from private key
func PrivateKeyToPublicKey(privateKey string) string {
	prvKey := new(big.Int)
	_, ok := prvKey.SetString(privateKey, 0)
	if !ok {
		panic("privateKey is invalid.")
	}

	c := crypto.S256()
	// _privateKey := hexutil.MustDecodeBig(privateKey).Bytes()
	pubKeyX, pubKeyY := c.ScalarBaseMult(prvKey.Bytes())
	pubKeyBytes := crypto.FromECDSAPub(&ecdsa.PublicKey{
		Curve: c,
		X:     pubKeyX,
		Y:     pubKeyY,
	})

	pubKey := hexutil.Encode(pubKeyBytes[1:])
	return pubKey
}

// Keccak256 hash hex string and return it's hash value
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
