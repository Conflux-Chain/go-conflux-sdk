package types

import (
	"encoding/json"
	"testing"
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
		err := json.Unmarshal([]byte(`"`+epochStr+`"`), &Epoch{})
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
