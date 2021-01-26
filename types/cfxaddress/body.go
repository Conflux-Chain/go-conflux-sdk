package cfxaddress

import (
	"github.com/pkg/errors"
)

/*
Base32-encode address:
To create the payload, first, concatenate the version-byte with addr to get a 21-byte array. Then, encode it left-to-right, mapping each 5-bit sequence to the corresponding ASCII character (see alphabet below). Pad to the right with zero bits to complete any unfinished chunk at the end. In our case, 21-byte payload + 2 bit 0-padding will result in a 34-byte base32-encoded string.

We use the following alphabet: abcdefghjkmnprstuvwxyz0123456789 (i, l, o, q removed).

0x00 => a    0x08 => j    0x10 => u    0x18 => 2
0x01 => b    0x09 => k    0x11 => v    0x19 => 3
0x02 => c    0x0a => m    0x12 => w    0x1a => 4
0x03 => d    0x0b => n    0x13 => x    0x1b => 5
0x04 => e    0x0c => p    0x14 => y    0x1c => 6
0x05 => f    0x0d => r    0x15 => z    0x1d => 7
0x06 => g    0x0e => s    0x16 => 0    0x1e => 8
0x07 => h    0x0f => t    0x17 => 1    0x1f => 9
*/

// Body reperents 5bits byte array of concating version byte with hex address
type Body []byte

// NewBodyByString creates body by base32 string which contains version byte and hex address
func NewBodyByString(base32Str string) (body Body, err error) {
	for _, v := range base32Str {
		index, ok := alphabetToIndexMap[v]
		if !ok {
			err = errors.New("invalid base32 string for body")
		}
		body = append(body, index)
	}
	return
}

// NewBodyByHexAddress convert concat of version type and hex address to 5 bits slice
func NewBodyByHexAddress(vrsByte VersionByte, hexAddress []byte) (b Body, err error) {
	vb, err := vrsByte.ToByte()
	if err != nil {
		err = errors.Wrapf(err, "failed to encode version type %#v", vrsByte)
		return
	}
	concatenate := append([]byte{vb}, hexAddress[:]...)
	bits5, err := convert(concatenate, 8, 5)
	if err != nil {
		err = errors.Wrapf(err, "failed to convert %x from 8 to 5 bits array", concatenate)
		return
	}
	b = bits5
	return
}

// ToHexAddress decode bits5 array to version byte and hex address
func (b Body) ToHexAddress() (vrsType VersionByte, hexAddress []byte, err error) {
	if len(b) == 0 {
		err = errors.New("invalid base32 body")
	}

	val, err := convert(b, 5, 8)
	vrsType = NewVersionByte(val[0])
	hexAddress = val[1:]
	return
}

// String return base32 string
func (b Body) String() string {
	return bits5sToString(b)
}
