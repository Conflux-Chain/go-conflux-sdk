package types

import (
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	H128Length = 16
)

var (
	h128T = reflect.TypeOf(H128{})
)

type H128 [H128Length]byte

func HexToH128(s string) H128 {
	var h H128
	h.SetBytes(common.FromHex(s))
	return h
}

func (h H128) String() string { return h.Hex() }

func (h H128) Hex() string { return hexutil.Encode(h[:]) }

func (h H128) MarshalText() ([]byte, error) {
	return hexutil.Bytes(h[:]).MarshalText()
}

func (h *H128) UnmarshalText(input []byte) error {
	return hexutil.UnmarshalFixedText("H128", input, h[:])
}

func (h *H128) UnmarshalJSON(data []byte) error {
	return hexutil.UnmarshalFixedJSON(h128T, data, h[:])
}

// SetBytes sets the H128 to the value of b.
// If b is larger than len(h), b will be cropped from the left.
func (h *H128) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-H128Length:]
	}
	copy(h[H128Length-len(b):], b)
}
