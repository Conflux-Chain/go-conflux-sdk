package cfxaddress

import (
	"errors"
	"strings"
)

const (
	alphabet = "abcdefghjkmnprstuvwxyz0123456789"
)

var (
	alphabetToIndexMap map[rune]byte = make(map[rune]byte)
)

func init() {
	for i, v := range alphabet {
		alphabetToIndexMap[v] = byte(i)
	}
}

// convert converts data from inbits byte array to outbits byte array
func convert(data []byte, inbits uint, outbits uint) ([]byte, error) {
	// fmt.Printf("convert %b from %v bits to %v bits\n", data, inbits, outbits)
	// only support bits length<=8
	if inbits > 8 || outbits > 8 {
		return nil, errors.New("only support bits length<=8")
	}

	accBits := uint(0) //accumulate bit length
	acc := uint16(0)   //accumulate value
	var ret []byte
	for _, d := range data {
		acc = acc<<uint16(inbits) | uint16(d)
		// fmt.Printf("acc1: %b\n", acc)
		accBits += inbits
		for accBits >= outbits {
			val := byte(acc >> uint16(accBits-outbits))
			// fmt.Printf("5bits val:%v\n", val)
			ret = append(ret, val)
			// fmt.Printf("ret: %b\n", ret)
			acc = acc & uint16(1<<(accBits-outbits)-1)
			// fmt.Printf("acc2: %b\n", acc)
			accBits -= outbits
		}
	}
	// if acc > 0 || accBits > 0 {
	if accBits > 0 && (inbits > outbits) {
		ret = append(ret, byte(acc<<uint16(outbits-accBits)))
	}
	// fmt.Printf("ret %b\n", ret)
	return ret, nil
}

func bits5sToString(dataInBits5 []byte) string {
	sb := strings.Builder{}
	for _, v := range dataInBits5 {
		sb.WriteRune(rune(alphabet[v]))
	}
	return sb.String()
}

// Modification based on https://github.com/hlb8122/rust-bitcoincash-addr in MIT License.
// A copy of the original license is included in LICENSE.rust-bitcoincash-addr.

// https://github.com/bitcoincashorg/bitcoincash.org/blob/master/spec/cashaddr.md#checksum
func polymod(v []byte) uint64 {
	c := uint64(1)
	for _, d := range v {
		c0 := byte(c >> 35)
		c = ((c & 0x07ffffffff) << 5) ^ uint64(d)
		if c0&0x01 != 0 {
			c ^= 0x98f2bc8e61
		}
		if c0&0x02 != 0 {
			c ^= 0x79b76d99e2
		}
		if c0&0x04 != 0 {
			c ^= 0xf33e5fb3c4
		}
		if c0&0x08 != 0 {
			c ^= 0xae2eabe2a8
		}
		if c0&0x10 != 0 {
			c ^= 0x1e4f43e470
		}
	}
	return c ^ 1
}

func uint64ToBytes(num uint64) []byte {
	r := make([]byte, 8)
	for i := 0; i < 8; i++ {
		r[7-i] = byte(num >> uint(i*8))
	}
	return r
}
