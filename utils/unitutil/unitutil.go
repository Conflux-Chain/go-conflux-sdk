package unitutil

import (
	"math/big"

	"github.com/shopspring/decimal"
)

type UnitType int32

const (
	UNIT_CFX   UnitType = 18
	UNIT_GDrip UnitType = 9
	UNIT_Drip  UnitType = 1
)

func FormatUnits(valueInDrip *big.Int, unit UnitType) decimal.Decimal {
	return decimal.NewFromBigInt(valueInDrip, -int32(unit))
}

func FormatCFX(valueInDrip *big.Int) decimal.Decimal {
	return FormatUnits(valueInDrip, UNIT_CFX)
}

func ParseUnits(valueInCFX decimal.Decimal, unit UnitType) *big.Int {
	multi := decimal.New(1, int32(unit))
	return valueInCFX.Mul(multi).BigInt()
}

func ParseDrip(valueInCFX decimal.Decimal, unit UnitType) *big.Int {
	return ParseUnits(valueInCFX, UNIT_Drip)
}
