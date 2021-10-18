package postypes

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
)

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

// UnmarshalJSON implements the json.Unmarshaler interface.
func (e *BlockNumber) UnmarshalJSON(data []byte) error {
	var input string

	if err := json.Unmarshal(data, &input); err != nil {
		return err
	}

	hexU64Pattern := `(?i)^0x[a-f0-9]*$`
	if ok, _ := regexp.Match(hexU64Pattern, []byte(input)); ok {
		numU64, err := hexutil.DecodeUint64(input)
		if err != nil {
			return errors.WithStack(err)
		}
		e.number = hexutil.Uint64(numU64)
		return nil
	}

	switch input {
	case BlockEarliest.name, BlockLatestCommitted.name, BlockLatestVoted.name:
		e.name = input
		return nil
	}

	return fmt.Errorf(`unsupported pos block number tag %v`, input)
}
