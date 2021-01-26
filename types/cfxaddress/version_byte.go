package cfxaddress

import "github.com/pkg/errors"

/*
Version-byte:
The version byte's most significant bit is reserved and must be 0. The 4 next bits indicate the type of address and the 3 least significant bits indicate the size of the hash.

Type bits	Meaning	Version byte value
0			Conflux			0
Further types might be added as new features are added.

Size bits	Hash size in bits
0			160
1			192
2			224
3			256
4			320
5			384
6			448
7			512
By encoding the size of the hash in the version field, we ensure that it is possible to check that the length of the address is correct.
*/

// VersionByte conmposites by type bits, address type and size bits according above description from CIP-37
type VersionByte struct {
	TypeBits uint8
	// current is constant 0, it's different with AddressType defined in address_type.go
	AddressType uint8
	SizeBits    uint8
}

var (
	hashSizeToBits map[uint]uint8 = make(map[uint]uint8)
)

func init() {
	hashSizeToBits[160] = 0
	hashSizeToBits[192] = 1
	hashSizeToBits[224] = 2
	hashSizeToBits[256] = 3
	hashSizeToBits[320] = 4
	hashSizeToBits[384] = 5
	hashSizeToBits[448] = 6
	hashSizeToBits[512] = 7
}

// ToByte returns byte
func (v VersionByte) ToByte() (byte, error) {
	ret := v.TypeBits & 0x80
	ret = ret | v.AddressType<<3
	ret = ret | v.SizeBits
	return ret, nil
}

// NewVersionByte creates version byte by byte
func NewVersionByte(b byte) (vt VersionByte) {
	vt.TypeBits = b >> 7
	vt.AddressType = (b & 0x7f) >> 3
	vt.SizeBits = b & 0x0f
	return
}

// CalcVersionByte calculates version byte of hex address
func CalcVersionByte(hexAddress []byte) (versionByte VersionByte, err error) {
	versionByte.TypeBits = 0
	versionByte.AddressType = 0
	addrBitsLen := uint(len(hexAddress) * 8)
	versionByte.SizeBits = hashSizeToBits[addrBitsLen]
	if versionByte.SizeBits == 0 && addrBitsLen != 160 {
		return versionByte, errors.Errorf("Invalid hash size %v", addrBitsLen)
	}
	return
}
