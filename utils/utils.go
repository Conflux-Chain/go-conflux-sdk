package utils

import (
	"math/big"
	"reflect"

	"github.com/Conflux-Chain/go-conflux-sdk/constants"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// CalcBlockConfirmationRisk calculates block revert rate
func CalcBlockConfirmationRisk(rawConfirmationRisk *big.Int) *big.Float {
	riskFloat := new(big.Float).SetInt(rawConfirmationRisk)
	maxUint256Float := new(big.Float).SetInt(constants.MaxUint256)
	riskRate := new(big.Float).Quo(riskFloat, maxUint256Float)
	return riskRate
}

// IsNil sepecialy checks if interface object is nil
func IsNil(i interface{}) bool {

	if i == nil {
		return true
	}

	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

// HexStringToBytes converts hex string to bytes
func HexStringToBytes(hexStr string) (hexutil.Bytes, error) {
	if !Has0xPrefix(hexStr) {
		hexStr = "0x" + hexStr
	}
	return hexutil.Decode(hexStr)
}

// Has0xPrefix returns true if input starts with '0x' or '0X'
func Has0xPrefix(input string) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
}
