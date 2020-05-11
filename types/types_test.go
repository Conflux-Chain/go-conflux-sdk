package types

import "testing"

func TestAddressIsZero(t *testing.T) {
	zeroAddrs := []Address{Address("0x0000000000000000000000000000000000000000"), Address("0X0000000000000000000000000000000000000000")}
	for _, addr := range zeroAddrs {
		if !addr.IsZero() {
			t.Errorf("expect %+v be zero address", addr)
		}
		if !(&addr).IsZero() {
			t.Errorf("expect %+v be zero address", &addr)
		}
	}

	normalAddr := Address("0x0000000000fa000d00e000000000000000000000")
	if normalAddr.IsZero() {
		t.Errorf("expect %+v be zero address", normalAddr)
	}
	if (&normalAddr).IsZero() {
		t.Errorf("expect %+v be zero address", &normalAddr)
	}
}
