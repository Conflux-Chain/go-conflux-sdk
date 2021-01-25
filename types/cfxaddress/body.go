package cfxaddress

import (
	"github.com/pkg/errors"
)

/*
Base32-encode address:
To create the payload, first, concatenate the version-byte with addr to get a 21-byte array. Then, encode it left-to-right, mapping each 5-bit sequence to the corresponding ASCII character (see alphabet below). Pad to the right with zero bits to complete any unfinished chunk at the end. In our case, 21-byte payload + 2 bit 0-padding will result in a 34-byte base32-encoded string.

We use the following alphabet: 0123456789abcdefghjkmnprstuvwxyz (o, i, l, q removed).

0x00 => 0    0x08 => 8    0x10 => g    0x18 => s
0x01 => 1    0x09 => 9    0x11 => h    0x19 => t
0x02 => 2    0x0a => a    0x12 => j    0x1a => u
0x03 => 3    0x0b => b    0x13 => k    0x1b => v
0x04 => 4    0x0c => c    0x14 => m    0x1c => w
0x05 => 5    0x0d => d    0x15 => n    0x1d => x
0x06 => 6    0x0e => e    0x16 => p    0x1e => y
0x07 => 7    0x0f => f    0x17 => r    0x1f => z
*/

// Body reperents by 5bits byte array
type Body []byte

// NewBodyByString ...
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

// ToHexAddress ...
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
