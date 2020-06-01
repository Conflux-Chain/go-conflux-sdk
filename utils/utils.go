package utils

import (
	"math/big"

	"github.com/Conflux-Chain/go-conflux-sdk/constants"
)

// CalcBlockRevertRate calculates block revert rate
func CalcBlockRevertRate(confirmRisk *big.Int) *big.Float {
	riskFloat := new(big.Float).SetInt(confirmRisk)
	maxUint256Float := new(big.Float).SetInt(constants.MaxUint256)
	riskRate := new(big.Float).Quo(riskFloat, maxUint256Float)
	return riskRate
}
