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
	privateKeyStr = "0xd28edbdb7bbe75787b84c5f525f47666a3274bb06561581f00839645f3c26f66"
	publicKeyStr  = "0xc42b53ae2ef95fee489948d33df391c4a9da31b7a3e29cf772c24eb42f74e94ab3bfe00bf29a239c17786a5b921853b7c5344d36694db43aa849e401f91566a5"
	addressStr    = "0x1cecb4a2922b7007e236daf0c797de6e55496e84"
)

func TestPublicKeyToAddress(t *testing.T) {
	expect := addressStr[2:]

	actual := PublicKeyToCommonAddress(publicKeyStr)
	if expect != hex.EncodeToString(actual[:]) {
		t.Errorf("Test PublicKeyToAddress failed, except %v,  actual %x", expect, actual)
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
