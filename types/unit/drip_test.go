package unit

import (
	"math/big"
	"testing"

	"gotest.tools/assert"
)

func TestNewDripFromString(t *testing.T) {
	source := []struct {
		in       string
		out      *Drip
		isOutErr bool
	}{
		{
			in:  "1",
			out: NewDrip(big.NewInt(1)),
		},
		{
			in:  "1 Drip",
			out: NewDrip(big.NewInt(1)),
		},
		{
			in:       "1.1 Drip",
			isOutErr: true,
		},
		{
			in:  "1.1 uCFX",
			out: NewDrip(big.NewInt(1100000000000)),
		},
		{
			in:       "1 uCFX .",
			isOutErr: true,
		},
	}

	for _, v := range source {
		actual, e := NewDripFromString(v.in)

		assert.Equal(t, v.isOutErr, e != nil)

		if !v.isOutErr && actual.Cmp(v.out) != 0 {
			t.Fatal("in:", v.in, "out:", v.out, "actual:", actual)
		}
	}

}
