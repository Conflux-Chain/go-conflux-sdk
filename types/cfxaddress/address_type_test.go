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
