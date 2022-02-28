package cmptutil

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Bytes for unmarshalling both hex and byte array
type Bytes []byte

func (b Bytes) MarshalText() ([]byte, error) {
	return hexutil.Bytes(b).MarshalText()
}

func (b *Bytes) UnmarshalJSON(data []byte) error {
	var hex hexutil.Bytes
	if err := json.Unmarshal(data, &hex); err == nil {
		*b = Bytes([]byte(hex))
		return nil
	}

	nums := make([]uint, 0)
	if err := json.Unmarshal(data, &nums); err != nil {
		return err
	}

	for _, v := range nums {
		*b = append(*b, byte(v))
	}
	return nil
}

func (b *Bytes) ToBytes() []byte {
	return []byte(*b)
}

func (b *Bytes) ToHexBytes() hexutil.Bytes {
	return hexutil.Bytes(*b)
}
