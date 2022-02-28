package unit

import (
	"math/big"
	"strings"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

// Drip is minimal unit of conflux network coin, 1CFX = 10^18 DRIP
type Drip struct {
	value *big.Int
}

func NewDrip(value *big.Int) *Drip {
	return &Drip{value}
}

// NewDripFromString create Drip from string, prettyValue could be one part or two parts,
// if one part like "12345" that equals to "12345 Drip", if two parts like "1.2 CFX", the second part is unit
// 		NewDripFromString("12345") => 12345 Drip
// 		NewDripFromString("1.2 CFX") => 1.2 CFX
func NewDripFromString(prettyValue string) (*Drip, error) {
	parts := strings.Split(prettyValue, " ")

	if len(parts) == 1 {
		val, ok := new(big.Int).SetString(parts[0], 0)
		if !ok {
			return nil, errors.New("invalid value")
		}
		return &Drip{val}, nil
	}

	if len(parts) != 2 {
		return nil, errors.New("invalid format")
	}

	value, err := decimal.NewFromString(parts[0])
	if err != nil {
		return nil, errors.WithStack(err)
	}

	unit, err := ParseUnitType(parts[1])
	if err != nil {
		return nil, errors.WithStack(err)
	}

	core, err := parseUnits(value, *unit)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &Drip{core}, nil
}

// ParseFrom parse drip from value with uint
// 		ParseFrom(10, UNIT_GCFX) => 10_000_000_000 Drip
func (d *Drip) ParseFrom(value decimal.Decimal, unit UnitType) error {
	r, err := parseUnits(value, unit)
	if err != nil {
		return nil
	}
	*d = Drip{r}
	return nil
}

// ParseFrom same to ParseFrom and unit is CFX
// 		ParseFrom(10) => 10_000_000_000 Drip
func (d *Drip) ParseFromCFX(value decimal.Decimal) error {
	return d.ParseFrom(value, UNIT_CFX)
}

// Format format drip to value with unit
// 		d := NewDrip(big.NewInt(1000000000))
// 		d.Format(UNIT_CFX) => 0.0000000001
func (d *Drip) Format(unit UnitType) decimal.Decimal {
	return formatUnits(d.BigInt(), unit)
}

// FormatCFX format drip to value with unit CFX
// 		d := NewDrip(big.NewInt(1000000000))
// 		d.Format() => 0.0000000001
func (d *Drip) FormatCFX() decimal.Decimal {
	return d.Format(UNIT_CFX)
}

// BigInt return drip value as big.Int
func (d *Drip) BigInt() *big.Int {
	return d.value
}

// Cmp compare drip value with another drip value
func (d *Drip) Cmp(y *Drip) int {
	return d.value.Cmp(y.value)
}

// String implements Stringer interface
func (d Drip) String() string {
	return displayValueWithUnit(d.value)
}
