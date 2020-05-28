package utils

import (
	"math/big"

	"github.com/Conflux-Chain/go-conflux-sdk/constants"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// CalcBlockRevertRate calculates block revert rate
func CalcBlockRevertRate(confirmRisk *big.Int) *big.Float {
	riskFloat := new(big.Float).SetInt(confirmRisk)
	maxUint256Float := new(big.Float).SetInt(constants.MaxUint256)
	riskRate := new(big.Float).Quo(riskFloat, maxUint256Float)
	return riskRate
}

func ConvertBigIntToHexutilBig(input *big.Int) *hexutil.Big {
	_tmp := hexutil.Big(*input)
	return &_tmp
}
