package internalcontract

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/stretchr/testify/assert"
)

func TestUnpackToTuple(t *testing.T) {
	data, err := hex.DecodeString("00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000")
	assert.NoError(t, err)

	abiStr := "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"votePower\",\"type\":\"uint64\"}],\"name\":\"IncreaseStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"blsPubKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"vrfPubKey\",\"type\":\"bytes\"}],\"name\":\"Register\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"votePower\",\"type\":\"uint64\"}],\"name\":\"Retire\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addressToIdentifier\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"}],\"name\":\"identifierToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"votePower\",\"type\":\"uint64\"}],\"name\":\"increaseStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"indentifier\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"votePower\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"blsPubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"vrfPubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes[2]\",\"name\":\"blsPubKeyProof\",\"type\":\"bytes[2]\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"votePower\",\"type\":\"uint64\"}],\"name\":\"retire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	_abi, err := abi.JSON(bytes.NewReader([]byte(abiStr)))
	assert.NoError(t, err)

	var out [2]*big.Int
	err = _abi.UnpackIntoInterface(&out, "getVotes", data)
	assert.NoError(t, err)

	assert.Equal(t, out[0].String(), "1")
	assert.Equal(t, out[1].String(), "0")
}
