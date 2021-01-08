package types

import (
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

func TestString(t *testing.T) {
	var a *Epoch = &Epoch{}
	if a.String() != "" {
		t.Errorf("expect return empty when a is nil")
	}
}
