// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package types

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Const epoch definitions
var (
	EpochEarliest         *Epoch = &Epoch{"earliest", nil, false}
	EpochLatestCheckpoint *Epoch = &Epoch{"latest_checkpoint", nil, false}
	EpochLatestConfirmed  *Epoch = &Epoch{"latest_confirmed", nil, false}
	EpochLatestState      *Epoch = &Epoch{"latest_state", nil, false}
	EpochLatestMined      *Epoch = &Epoch{"latest_mined", nil, false}
	EpochLatestFinalized  *Epoch = &Epoch{"latest_finalized", nil, false}
)

// Epoch represents an epoch in Conflux.
type Epoch struct {
	name         string
	number       *hexutil.Big
	requirePivot bool
}

// WebsocketEpochResponse represents result of epoch websocket subscription
type WebsocketEpochResponse struct {
	EpochHashesOrdered []Hash       `json:"epochHashesOrdered"`
	EpochNumber        *hexutil.Big `json:"epochNumber"`
}

// NewEpochNumber creates an instance of Epoch with specified number.
func NewEpochNumber(number *hexutil.Big) *Epoch {
	return &Epoch{"", number, false}
}

// NewEpochNumberBig creates an instance of Epoch with specified big number.
func NewEpochNumberBig(number *big.Int) *Epoch {
	return &Epoch{"", NewBigIntByRaw(number), false}
}

// NewEpochNumberUint64 creates an instance of Epoch with specified uint64 number.
func NewEpochNumberUint64(number uint64) *Epoch {
	return &Epoch{"", NewBigInt(number), false}
}

// NewEpochWithBlockHash creates an instance of Epoch with specified block hash.
func NewEpochWithBlockHash(blockHash Hash, requirePivot ...bool) *Epoch {
	if len(requirePivot) == 0 {
		requirePivot = append(requirePivot, false)
	}
	return &Epoch{string(blockHash), nil, requirePivot[0]}
}

// String implements the fmt.Stringer interface
func (e *Epoch) String() string {
	if e.number != nil {
		return e.number.String()
	}

	return e.name
}

// ToInt returns epoch number in type big.Int
func (e *Epoch) ToInt() (result *big.Int, isSuccess bool) {
	if e.number != nil {
		return e.number.ToInt(), true
	}

	if e.name == EpochEarliest.name {
		return common.Big0, true
	}

	return nil, false
}

// Equals checks if e equals target
func (e *Epoch) Equals(target *Epoch) bool {
	if e == nil {
		panic("input could not be nil")
	}

	if target == nil {
		return false
	}

	if e == target {
		return true
	}

	if len(e.name) > 0 || len(target.name) > 0 {
		return e.name == target.name
	}

	if e.number == nil || target.number == nil {
		return e.number == target.number
	}

	return e.number.ToInt().Cmp(target.number.ToInt()) == 0
}

// TODO: support requirePivot
// MarshalText implements the encoding.TextMarshaler interface.
func (e Epoch) MarshalText() ([]byte, error) {
	return []byte(e.String()), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (e *Epoch) UnmarshalJSON(data []byte) error {
	var input string
	if err := json.Unmarshal(data, &input); err != nil {
		return err
	}

	switch input {
	case EpochEarliest.name,
		EpochLatestCheckpoint.name,
		EpochLatestConfirmed.name,
		EpochLatestState.name,
		EpochLatestMined.name,
		EpochLatestFinalized.name:
		e.name = input
		return nil
	default:
		if len(input) == 66 {
			e.name = input
			return nil
		}

		epochNumber, err := hexutil.DecodeBig(input)
		if err != nil {
			return err
		}

		e.number = NewBigIntByRaw(epochNumber)
		return nil
	}
}
