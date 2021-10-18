package postypes

import "github.com/ethereum/go-ethereum/common/hexutil"

type BlockNumber struct {
	name   string
	number hexutil.Uint64
}

var (
	BlockEarliest        *BlockNumber = &BlockNumber{"earliest", 0}
	BlockLatestCommitted *BlockNumber = &BlockNumber{"latest_committed", 0}
	BlockLatestVoted     *BlockNumber = &BlockNumber{"latest_voted", 0}
)

func NewBlockNumber(number uint64) BlockNumber {
	return BlockNumber{"", hexutil.Uint64(number)}
}

// String implements the fmt.Stringer interface
func (e *BlockNumber) String() string {
	if e.name == "" {
		return e.number.String()
	}

	return e.name
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e BlockNumber) MarshalText() ([]byte, error) {
	// fmt.Println("marshal text for epoch")
	return []byte(e.String()), nil
}
