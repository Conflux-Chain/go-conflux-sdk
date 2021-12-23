package cfxaddress

import (
	"encoding/hex"
	"testing"
)

func TestCalcAddressType(t *testing.T) {
	verifyAddressType(t, "006d49f8505410eb4e671d51f7d96d2c87807b09", AddressTypeBuiltin)
	verifyAddressType(t, "106d49f8505410eb4e671d51f7d96d2c87807b09", AddressTypeUser)
	verifyAddressType(t, "806d49f8505410eb4e671d51f7d96d2c87807b09", AddressTypeContract)
	verifyAddressType(t, "0000000000000000000000000000000000000000", AddressTypeNull)
	verifyAddressType(t, "2000000000000000000000000000000000000000", AddressTypeUnknown)
	expectCalcError(t, []byte{1, 2})
	expectCalcError(t, []byte{})
	expectCalcError(t, nil)
}

func verifyAddressType(t *testing.T, hexAddress string, expect AddressType) {
	addr, err := hex.DecodeString(hexAddress)
	fatalIfErr(t, err)
	actual, err := CalcAddressType(addr)
	fatalIfErr(t, err)
	if actual != expect {
		t.Fatalf("expect %v actual %v", expect, actual)
	}
}

func expectCalcError(t *testing.T, bytes []byte) {
	if _, err := CalcAddressType(bytes); err == nil {
		t.Fatalf("should return error")
	}
}
