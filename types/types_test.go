package types

import (
	"testing"
)

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

func TestGetAddressType(t *testing.T) {
	normalAddr := Address("0x1200000000fa000d00e000000000000000000000")
	customContractAddr := Address("0x8300000000fa000d00e000000000000000000000")
	internalContractAddr := Address("0x0888000000fa000d00e000000000000000000002")
	invalidAddr := Address("0x3300000000fa000d00e000000000000000000000")

	addressType := normalAddr.GetAddressType()
	if addressType != NormalAddress {
		t.Errorf("expect %+v be normal address, actual is %v", normalAddr, addressType)
	}

	addressType = customContractAddr.GetAddressType()
	if addressType != CustomContractAddress {
		t.Errorf("expect %+v be contract address, actual is %v", customContractAddr, addressType)
	}

	addressType = internalContractAddr.GetAddressType()
	if addressType != InternalContractAddress {
		t.Errorf("expect %+v be contract address, actual is %v", internalContractAddr, addressType)
	}

	addressType = invalidAddr.GetAddressType()
	if addressType != InvalidAddress {
		t.Errorf("expect %+v be unknown address,actual is %v", invalidAddr, addressType)
	}
}
