package unitutil

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/shopspring/decimal"
)

type UnitType int32

const (
	UNIT_CFX   UnitType = 18
	UNIT_PDrip UnitType = 15
	UNIT_TDrip UnitType = 12
	UNIT_GDrip UnitType = 9
	UNIT_MDrip UnitType = 6
	UNIT_KDrip UnitType = 3
	UNIT_Drip  UnitType = 0
)

func (u UnitType) String() string {
	switch u {
	case UNIT_CFX:
		return "CFX"
	case UNIT_PDrip:
		return "PDrip"
	case UNIT_TDrip:
		return "TDrip"
	case UNIT_GDrip:
		return "GDrip"
	case UNIT_MDrip:
		return "MDrip"
	case UNIT_KDrip:
		return "KDrip"
	case UNIT_Drip:
		return "Drip"
	default:
		return "UNKNOWN"
	}
}

// FormatUnits returns a Decimal representation of value formatted with unit digits (if it is a number) or to the unit specified (if a constant UNIT_XXX).
func FormatUnits(valueInDrip *big.Int, unit UnitType) decimal.Decimal {
	multi := decimal.New(1, -int32(unit))
	return decimal.NewFromBigInt(valueInDrip, 0).Mul(multi)
}

// FormatCFX equals to calling FormatUnits(value, UNIT_CFX).
func FormatCFX(valueInDrip *big.Int) decimal.Decimal {
	return FormatUnits(valueInDrip, UNIT_CFX)
}

// ParseUnits returns a *big.Int representation of value, parsed with unit digits (if it is a number) or from the unit specified (if a constant UNIT_XXX).
func ParseUnits(value decimal.Decimal, unit UnitType) (*big.Int, error) {
	multi := decimal.New(1, int32(unit))
	valueInDrip := value.Mul(multi)
	if valueInDrip.IsInteger() {
		return value.Mul(multi).BigInt(), nil
	}
	return nil, errors.New("fractional component exceeds decimals")
}

// ParseCFX equals to calling ParseUnits(value, UNIT_CFX).
func ParseCFX(value decimal.Decimal) (*big.Int, error) {
	return ParseUnits(value, UNIT_Drip)
}

// DisplayValueWithUnit returns the display format for given drip value.
func DisplayValueWithUnit(drip *big.Int) string {

	units := []UnitType{UNIT_CFX, UNIT_PDrip, UNIT_TDrip, UNIT_GDrip, UNIT_MDrip, UNIT_KDrip, UNIT_Drip}

	var mappedUnit UnitType
	for _, unit := range units {
		var i, e = big.NewInt(10), big.NewInt(int64(unit))
		i.Exp(i, e, nil)
		if drip.Cmp(i) >= 0 {
			mappedUnit = unit
			break
		}
	}

	return fmt.Sprintf("%v %v", FormatUnits(drip, mappedUnit), mappedUnit)
}

// MustParseUnits same as ParseUnits but panic if error
func MustParseUnits(value decimal.Decimal, unit UnitType) *big.Int {
	r, err := ParseUnits(value, unit)
	if err != nil {
		panic(err)
	}
	return r
}

// MustParseCFX same as ParseCFX but panic if error
func MustParseCFX(value decimal.Decimal) *big.Int {
	return MustParseUnits(value, UNIT_CFX)
}
