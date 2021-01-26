package cfxaddress

import (
	"github.com/pkg/errors"
)

/*
Calculate checksum: We use the checksum algorithm of Bitcoin Cash, defined here.
uint64_t PolyMod(const data &v) {
    uint64_t c = 1;
    for (uint8_t d : v) {
        uint8_t c0 = c >> 35;
        c = ((c & 0x07ffffffff) << 5) ^ d;

        if (c0 & 0x01) c ^= 0x98f2bc8e61;
        if (c0 & 0x02) c ^= 0x79b76d99e2;
        if (c0 & 0x04) c ^= 0xf33e5fb3c4;
        if (c0 & 0x08) c ^= 0xae2eabe2a8;
        if (c0 & 0x10) c ^= 0x1e4f43e470;
    }

    return c ^ 1;
}
The checksum is calculated over the following data:

The lower 5 bits of each character of the prefix. - e.g. "bit..." becomes 2,9,20,...
A zero for the separator (5 zero bits).
The payload by chunks of 5 bits. If necessary, the payload is padded to the right with zero bits to complete any unfinished chunk at the end.
Eight zeros as a "template" for the checksum.
Importantly, optional fields (like address-type) are NOT part of the checksum computation.

The 40-bit number returned by PolyMod is split into eight 5-bit numbers (msb first). The payload and the checksum are then encoded according to the base32 character table.
*/

// Checksum represents by 5bits byte array
type Checksum [8]byte

// CalcChecksum calculates checksum by network type and body
func CalcChecksum(nt NetworkType, body Body) (c Checksum, err error) {
	var lower5bitsNettype []byte
	for _, v := range nt.String() {
		lower5bitsNettype = append(lower5bitsNettype, byte(v)&0x1f)
	}
	separator := byte(0)
	payload5Bits := body
	template := [8]byte{}

	checksumInput := append(lower5bitsNettype, separator)
	checksumInput = append(checksumInput, payload5Bits...)
	checksumInput = append(checksumInput, template[:]...)

	// fmt.Printf("checksumInput:%x\n", checksumInput)

	uint64Chc := polymod(checksumInput)
	// fmt.Printf("uint64Chc:%x\n", uint64Chc)

	low40BitsChc := uint64ToBytes(uint64Chc)[3:]
	// fmt.Printf("low40BitsChc of %x:%x\n", uint64ToBytes(uint64Chc), low40BitsChc)

	checksumIn5Bits, err := convert(low40BitsChc, 8, 5)
	// fmt.Printf("low40BitsChcIn5Bits:%x\n", checksumIn5Bits)

	if err != nil {
		err = errors.Wrapf(err, "failed to convert %v from 8 to 5 bits", low40BitsChc)
		return
	}
	copy(c[:], checksumIn5Bits)
	return
}

// String returns base32 string of checksum according to CIP-37
func (c Checksum) String() string {
	return bits5sToString(c[:])
}
