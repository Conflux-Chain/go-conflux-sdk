package unit

import (
	"math/big"
	"strings"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type Drip struct {
	value *big.Int
}

func NewDrip(value *big.Int) *Drip {
	return &Drip{value}
}

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

func (d *Drip) ParseFrom(value decimal.Decimal, unit UnitType) error {
	r, err := parseUnits(value, unit)
	if err != nil {
		return nil
	}
	*d = Drip{r}
	return nil
}

func (d *Drip) ParseFromCFX(value decimal.Decimal) error {
	return d.ParseFrom(value, UNIT_CFX)
}

func (d *Drip) Format(unit UnitType) decimal.Decimal {
	return formatUnits(d.BigInt(), unit)
}

func (d *Drip) FormatCFX() decimal.Decimal {
	return d.Format(UNIT_CFX)
}

func (d *Drip) BigInt() *big.Int {
	return d.value
}

func (d *Drip) Cmp(y *Drip) int {
	return d.value.Cmp(y.value)
}

func (d Drip) String() string {
	return displayValueWithUnit(d.value)
}
