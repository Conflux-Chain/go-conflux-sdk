package unit

import (
	"math/big"
	"testing"

	"github.com/shopspring/decimal"
	"gotest.tools/assert"
)

func TestFormatUnits(t *testing.T) {
	source := []struct {
		in   *big.Int
		unit UnitType
		out  decimal.Decimal
	}{
		{
			in:   big.NewInt(1),
			unit: UNIT_CFX,
			out:  MustNewDecimalFromString("0.000000000000000001"),
		},
		{
			in:   big.NewInt(1),
			unit: UNIT_GDrip,
			out:  MustNewDecimalFromString("0.000000001"),
		},
		{
			in:   big.NewInt(1),
			unit: UNIT_Drip,
			out:  MustNewDecimalFromString("1"),
		},
		{
			in:   big.NewInt(1234567890),
			unit: UNIT_GDrip,
			out:  MustNewDecimalFromString("1.23456789"),
		},
	}

	for _, v := range source {
		r := formatUnits(v.in, v.unit)
		// fmt.Printf("%v %v %v\n", v.in, r, v.out)
		isEqual := r.Cmp(v.out)
		assert.Equal(t, 0, isEqual)
	}
}

func TestParseUnits(t *testing.T) {
	source := []struct {
		in         decimal.Decimal
		unit       UnitType
		out        *big.Int
		isOutError bool
	}{
		{
			in:   MustNewDecimalFromString("0.000000000000000001"),
			unit: UNIT_CFX,
			out:  big.NewInt(1),
		},
		{
			in:   MustNewDecimalFromString("0.000000001"),
			unit: UNIT_GDrip,
			out:  big.NewInt(1),
		},
		{
			in:   MustNewDecimalFromString("1"),
			unit: UNIT_Drip,
			out:  big.NewInt(1),
		},
		{
			in:   MustNewDecimalFromString("1.23456789"),
			unit: UNIT_GDrip,
			out:  big.NewInt(1234567890),
		},
		{
			in:         MustNewDecimalFromString("1.2345678901234567"),
			unit:       UNIT_GDrip,
			isOutError: true,
		},
	}

	for _, v := range source {
		r, err := parseUnits(v.in, v.unit)
		// fmt.Printf("%v %v %v %v\n", v.in, r, v.out, err)
		if !v.isOutError {

			isEqual := r.Cmp(v.out)
			assert.Equal(t, 0, isEqual)
			continue
		}
		assert.Assert(t, err != nil)
	}
}

func TestDisplay(t *testing.T) {
	source := []struct {
		in  *big.Int
		out string
	}{
		{
			in:  big.NewInt(1),
			out: "1 Drip",
		},
		{
			in:  big.NewInt(1000),
			out: "1 KDrip",
		},
		{
			in:  big.NewInt(1000000),
			out: "1 MDrip",
		},
		{
			in:  big.NewInt(1000000000),
			out: "1 GDrip",
		},
		{
			in:  big.NewInt(1000000000000),
			out: "1 uCFX",
		},
		{
			in:  big.NewInt(1000000000000000),
			out: "1 mCFX",
		},
		{
			in:  big.NewInt(1000000000000000000),
			out: "1 CFX",
		},
		{
			in:  new(big.Int).Mul(big.NewInt(1000000000000000000), big.NewInt(1000)),
			out: "1000 CFX",
		},
	}

	for _, v := range source {
		r := displayValueWithUnit(v.in)
		// fmt.Printf("%v %v %v\n", v.in, r, v.out)
		isEqual := r == v.out
		assert.Equal(t, true, isEqual)
	}
}

func MustNewDecimalFromString(value string) decimal.Decimal {
	v, err := decimal.NewFromString(value)
	if err != nil {
		panic(err)
	}
	return v
}
