package unitutil

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
			out:  decimal.NewFromBigInt(big.NewInt(1), -18),
		},
		{
			in:   big.NewInt(1),
			unit: UNIT_GDrip,
			out:  decimal.NewFromBigInt(big.NewInt(1), -9),
		},
		{
			in:   big.NewInt(1),
			unit: UNIT_Drip,
			out:  decimal.NewFromBigInt(big.NewInt(1), -1),
		},
		{
			in:   big.NewInt(1234567890),
			unit: UNIT_GDrip,
			out:  decimal.NewFromBigInt(big.NewInt(1234567890), -9),
		},
	}

	for _, v := range source {
		r := FormatUnits(v.in, v.unit)
		// fmt.Printf("%v %v %v\n", v.in, r, v.out)
		isEqual := r.Cmp(v.out)
		assert.Equal(t, 0, isEqual)
	}
}

func TestParseUnits(t *testing.T) {
	source := []struct {
		out  *big.Int
		unit UnitType
		in   decimal.Decimal
	}{
		{
			out:  big.NewInt(1),
			unit: UNIT_CFX,
			in:   decimal.NewFromBigInt(big.NewInt(1), -18),
		},
		{
			out:  big.NewInt(1),
			unit: UNIT_GDrip,
			in:   decimal.NewFromBigInt(big.NewInt(1), -9),
		},
		{
			out:  big.NewInt(1),
			unit: UNIT_Drip,
			in:   decimal.NewFromBigInt(big.NewInt(1), -1),
		},
		{
			out:  big.NewInt(1234567890),
			unit: UNIT_GDrip,
			in:   decimal.NewFromBigInt(big.NewInt(1234567890), -9),
		},
	}

	for _, v := range source {
		r := ParseUnits(v.in, v.unit)
		// fmt.Printf("%v %v %v\n", v.in, r, v.out)
		isEqual := r.Cmp(v.out)
		assert.Equal(t, 0, isEqual)
	}
}
