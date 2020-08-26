package utils

import (
	"math/big"
	"reflect"

	"github.com/Conflux-Chain/go-conflux-sdk/constants"
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
