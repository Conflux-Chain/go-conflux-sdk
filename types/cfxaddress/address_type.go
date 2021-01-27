package cfxaddress

import (
	"encoding/hex"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

/*
[OPTIONAL] Address-type: For the null address (0x0000000000000000000000000000000000000000), address-type must be "type.null". Otherwise,
    match addr[0] & 0xf0
        case b00000000: "type.builtin"
        case b00010000: "type.user"
        case b10000000: "type.contract"
*/

type AddressType string

const (
	AddressTypeBuiltin  AddressType = "builtin"
	AddressTypeUser     AddressType = "user"
	AddressTypeContract AddressType = "contract"
	AddressTypeNull     AddressType = "null"
)

// CalcAddressType calculate address type of hexAddress
func CalcAddressType(hexAddress []byte) (AddressType, error) {
	nullAddr, err := hex.DecodeString("0000000000000000000000000000000000000000")
	if err != nil {
		return "", err
	}
	if reflect.DeepEqual(nullAddr, hexAddress) {
		return AddressTypeNull, nil
	}

	var addressType AddressType
	switch hexAddress[0] & 0xf0 {
	case 0x00:
		addressType = AddressTypeBuiltin
	case 0x10:
		addressType = AddressTypeUser
	case 0x80:
		addressType = AddressTypeContract
	default:
		return "", errors.Errorf("Invalid address %x", hexAddress)
	}
	// fmt.Printf("calc address type of %x : %v\n", hexAddress, addressType)
	return addressType, nil
}

// ToByte returns byte represents of address type according to CIP-37
func (a AddressType) ToByte() (byte, error) {
	switch a {
	case AddressTypeNull:
		return 0x00, nil
	case AddressTypeBuiltin:
		return 0x00, nil
	case AddressTypeUser:
		return 0x10, nil
	case AddressTypeContract:
		return 0x80, nil
	}
	return 0, errors.Errorf("Invalid address type %v", a)
}

func (a AddressType) String() string {
	return fmt.Sprintf("type.%v", string(a))
}
