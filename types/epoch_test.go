package types

import (
	"testing"

	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/stretchr/testify/assert"
)

func TestEpochEquals(t *testing.T) {

	a := &Epoch{}
	b := &Epoch{}
	if !a.Equals(b) {
		t.Errorf("expect %v equals %v", a, b)
	}

	a = &Epoch{number: NewBigInt(1)}
	b = &Epoch{number: NewBigInt(1)}
	if !a.Equals(b) {
		t.Errorf("expect %v equals %v", a, b)
	}

	a = &Epoch{}
	b = &Epoch{number: NewBigInt(1)}
	if a.Equals(b) {
		t.Errorf("expect %v not equals %v", a, b)
	}
}

func TestUnmarshalEpoch(t *testing.T) {
	epochStrs := []string{
		"latest_mined",
		"latest_finalized",
	}

	for _, epochStr := range epochStrs {
		err := utils.JsonUnmarshal([]byte(`"`+epochStr+`"`), &Epoch{})
		if err != nil {
			t.Fatalf("failed unmarshal %v, %v", epochStr, err)
		}
	}

}

func TestString(t *testing.T) {
	var a *Epoch = &Epoch{}
	if a.String() != "" {
		t.Errorf("expect return empty when a is nil")
	}
}

func TestUnmarshalEpochOrBlockHash(t *testing.T) {
	// e := NewEpochOrBlockHashWithEpoch(*EpochLatestConfirmed)
	// e := NewEpochOrBlockHashWithBlockHash("0xdbccf46a86aa259e7693536d433558bf8bbf4c88a5ab176be28e4374d7a7a5bc", false)
	// e := EpochOrBlockHash{epochNumber: NewBigInt(1)}
	// j, _ := utils.JsonMarshal(e)
	// fmt.Println(string(j))

	inputs := []string{
		`"0x1"`,
		`"latest_confirmed"`,
		`{"blockHash":"0xdbccf46a86aa259e7693536d433558bf8bbf4c88a5ab176be28e4374d7a7a5bc","requirePivot":true}`,
		`{"blockHash":"0xdbccf46a86aa259e7693536d433558bf8bbf4c88a5ab176be28e4374d7a7a5bc","requirePivot":false}`,
		`{"epochNumber":"0x1"}`,
	}

	for _, input := range inputs {
		var e EpochOrBlockHash
		err := utils.JsonUnmarshal([]byte(input), &e)
		assert.NoError(t, err, input)

		jActual, _ := utils.JsonMarshal(e)
		expect, actual := utils.FormatJson(input), utils.FormatJson(string(jActual))
		assert.Equal(t, expect, actual, input)
	}
}
