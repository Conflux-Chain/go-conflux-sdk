// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package utils

import (
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

const (
	privateKeyStr = "0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"
	publicKeyStr  = "0x4646ae5047316b4230d0086c8acec687f00b1cd9d1dc634f6cb358ac0a9a8ffffe77b4dd0a4bfb95851f3b7355c781dd60f8418fc8a65d14907aff47c903a559"
	addressStr    = "0x1cad0b19bb29d4674531d6f115237e16afce377c"
)

func TestPublicKeyToAddress(t *testing.T) {
	expect := addressStr

	address := PublicKeyToAddress(publicKeyStr)
	if expect != string(address) {
		t.Errorf("Test PublicKeyToAddress failed, except %v,  actual %v", expect, address)
	}
}

func TestPrivateKeyToPublicKey(t *testing.T) {
	expect := publicKeyStr

	pubKey := PrivateKeyToPublicKey(privateKeyStr)
	if expect != string(pubKey) {
		t.Errorf("Test PrivateKeyToPublicKey failed, except %v,  actual %v", expect, pubKey)
	}
}

func TestKeccak256(t *testing.T) {

	inputBytes := []byte{0x12, 0x34, 0x56, 0x78}
	hash := crypto.Keccak256(inputBytes)
	expect := "0x" + hex.EncodeToString(hash)
	// t.Error(expect)

	input := "0x12345678"
	actual, err := Keccak256(input)
	if err != nil {
		t.Error(err)
	}
	if actual != expect {
		t.Errorf("Test Keccak256 failed, expect %+v, actual %+v", expect, actual)
	}
}
