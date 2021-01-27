package cfxaddress

import (
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	verifyConvert(t, []byte{0xff, 0xff, 0xff}, 8, 5, []byte{0x1f, 0x1f, 0x1f, 0x1f, 0x1e})
	verifyConvert(t, []byte{0x1f, 0x1f, 0x1f, 0x1f, 0x1e}, 5, 8, []byte{0xff, 0xff, 0xff})

	verifyConvert(t, []byte{0xaa, 0xaa, 0xaa}, 8, 5, []byte{0b10101, 0b01010, 0b10101, 0b01010, 0b10100})
	verifyConvert(t, []byte{0b10101, 0b01010, 0b10101, 0b01010, 0b10100}, 5, 8, []byte{0xaa, 0xaa, 0xaa})
}

func verifyConvert(t *testing.T, ihnput []byte, inbits uint, outbits uint, expectOutput []byte) {
	actualOutput, _ := convert(ihnput, inbits, outbits)
	if !reflect.DeepEqual(actualOutput, expectOutput) {
		t.Errorf("expect: %b, actual: %b", expectOutput, actualOutput)
	}
}

func TestUint64ToBytes(t *testing.T) {
	actual := uint64ToBytes(0x1b94eb1f)
	expect := []byte{0x00, 0x00, 0x00, 0x00, 0x1b, 0x94, 0xeb, 0x1f}
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("expect %x actual %x", expect, actual)
	}

}
