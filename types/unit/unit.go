package unit

import (
	"fmt"
	"math/big"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type UnitType int32

const (
	UNIT_CFX   UnitType = 18
	UNIT_mCFX  UnitType = 15
	UNIT_uCFX  UnitType = 12
	UNIT_GDrip UnitType = 9
	UNIT_MDrip UnitType = 6
	UNIT_KDrip UnitType = 3
	UNIT_Drip  UnitType = 0
)

func unitTypeVal2NameMap() map[UnitType]string {
	return map[UnitType]string{
		UNIT_CFX:   "CFX",
		UNIT_mCFX:  "mCFX",
		UNIT_uCFX:  "uCFX",
		UNIT_GDrip: "GDrip",
		UNIT_MDrip: "MDrip",
		UNIT_KDrip: "KDrip",
		UNIT_Drip:  "Drip",
	}
}

func unitTypeName2ValMap() map[string]UnitType {
	return map[string]UnitType{
		"CFX":   UNIT_CFX,
		"mCFX":  UNIT_mCFX,
		"uCFX":  UNIT_uCFX,
		"GDrip": UNIT_GDrip,
		"MDrip": UNIT_MDrip,
		"KDrip": UNIT_KDrip,
		"Drip":  UNIT_Drip,
	}
}

func (u UnitType) String() string {
	_map := unitTypeVal2NameMap()
	if _, ok := _map[u]; ok {
		return _map[u]
	}
	return "UNKNOWN"
}

func ParseUnitType(unitName string) (*UnitType, error) {
	_map := unitTypeName2ValMap()
	if _, ok := _map[unitName]; ok {
		n := _map[unitName]
		return &n, nil
	}
	return nil, errors.Errorf("unknown unit type: %s", unitName)
}

// formatUnits returns a Decimal representation of value formatted with unit digits (if it is a number) or to the unit specified (if a constant UNIT_XXX).
func formatUnits(valueInDrip *big.Int, unit UnitType) decimal.Decimal {
	multi := decimal.New(1, -int32(unit))
	return decimal.NewFromBigInt(valueInDrip, 0).Mul(multi)
}

// parseUnits returns a *big.Int representation of value, parsed with unit digits (if it is a number) or from the unit specified (if a constant UNIT_XXX).
func parseUnits(value decimal.Decimal, unit UnitType) (*big.Int, error) {
	multi := decimal.New(1, int32(unit))
	valueInDrip := value.Mul(multi)
	if valueInDrip.IsInteger() {
		return value.Mul(multi).BigInt(), nil
	}
	return nil, errors.New("fractional component exceeds decimals")
}

// displayValueWithUnit returns the display format for given drip value.
func displayValueWithUnit(drip *big.Int) string {

	units := []UnitType{UNIT_CFX, UNIT_mCFX, UNIT_uCFX, UNIT_GDrip, UNIT_MDrip, UNIT_KDrip, UNIT_Drip}

	var mappedUnit UnitType
	for _, unit := range units {
		var i, e = big.NewInt(10), big.NewInt(int64(unit))
		i.Exp(i, e, nil)
		if drip.Cmp(i) >= 0 {
			mappedUnit = unit
			break
		}
	}

	return fmt.Sprintf("%v %v", formatUnits(drip, mappedUnit), mappedUnit)
}
