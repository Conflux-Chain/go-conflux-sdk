package primitives

import (
	"io"

	"github.com/ethereum/go-ethereum/rlp"
)

// Bool represents a bool value that could be RLP encoded in Rust version.
type Bool bool

func (b Bool) EncodeRLP(w io.Writer) error {
	if b {
		// go-ethereum encoded as 0x80
		return rlp.Encode(w, "\x01")
	}

	return rlp.Encode(w, "\x00")
}
