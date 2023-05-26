// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// WebsocketEpochResponse represents result of epoch websocket subscription
type WebsocketEpochResponse struct {
	EpochHashesOrdered []Hash       `json:"epochHashesOrdered"`
	EpochNumber        *hexutil.Big `json:"epochNumber"`
}

// Epoch represents an epoch in Conflux.
type Epoch struct {
	name   string
	number *hexutil.Big
}

// Const epoch definitions
var (
	EpochEarliest         *Epoch = &Epoch{"earliest", nil}
	EpochLatestCheckpoint *Epoch = &Epoch{"latest_checkpoint", nil}
	EpochLatestConfirmed  *Epoch = &Epoch{"latest_confirmed", nil}
	EpochLatestState      *Epoch = &Epoch{"latest_state", nil}
	EpochLatestMined      *Epoch = &Epoch{"latest_mined", nil}
	EpochLatestFinalized  *Epoch = &Epoch{"latest_finalized", nil}
)

// NewEpochNumber creates an instance of Epoch with specified number.
func NewEpochNumber(number *hexutil.Big) *Epoch {
	return &Epoch{"", number}
}

// NewEpochNumberBig creates an instance of Epoch with specified big number.
func NewEpochNumberBig(number *big.Int) *Epoch {
	return &Epoch{"", NewBigIntByRaw(number)}
}

// NewEpochNumberUint64 creates an instance of Epoch with specified uint64 number.
func NewEpochNumberUint64(number uint64) *Epoch {
	return &Epoch{"", NewBigInt(number)}
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

type EpochOrBlockHash struct {
	epoch        *Epoch
	epochNumber  *hexutil.Big
	blockHash    *common.Hash
	requirePivot bool
}

// IsEpoch returns epoch if it is epoch, and the 2rd return value represents if is epoch.
func (e *EpochOrBlockHash) IsEpoch() (*Epoch, bool) {
	if e.epoch != nil {
		return e.epoch, true
	}
	if e.epochNumber != nil {
		return NewEpochNumber(e.epochNumber), true
	}
	return nil, false
}

// IsBlockHash returns "block hash" and "require pivot" if it is blockhash, and the 3rd return value represents if is block hash.
func (e *EpochOrBlockHash) IsBlockHash() (*common.Hash, bool, bool) {
	if e.blockHash != nil {
		return e.blockHash, e.requirePivot, true
	}
	return nil, false, false
}

// String implements the fmt.Stringer interface
func (e *EpochOrBlockHash) String() string {
	if e.epoch != nil {
		return e.epoch.String()
	}

	if e.epochNumber != nil {
		return e.epochNumber.String()
	}

	if e.blockHash != nil {
		return e.blockHash.String()
	}

	return "nil"
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e EpochOrBlockHash) MarshalText() ([]byte, error) {
	return []byte(e.String()), nil
}

func (e EpochOrBlockHash) MarshalJSON() ([]byte, error) {
	if e.epoch != nil {
		return json.Marshal(e.epoch)
	}

	if e.epochNumber != nil {
		return json.Marshal(struct {
			EpochNumber *hexutil.Big `json:"epochNumber"`
		}{
			EpochNumber: e.epochNumber,
		})
	}

	if e.blockHash != nil {
		return json.Marshal(struct {
			BlockHash    common.Hash `json:"blockHash"`
			RequirePivot bool        `json:"requirePivot"`
		}{
			BlockHash:    *e.blockHash,
			RequirePivot: e.requirePivot,
		})
	}
	return nil, errors.New("unkown EpochOrBlockHash")
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (e *EpochOrBlockHash) UnmarshalJSON(data []byte) error {

	var epoch Epoch
	err := json.Unmarshal(data, &epoch)
	if err == nil {
		e.epoch = &epoch
		return nil
	}

	type tmpEpoch struct {
		EpochNumber  *hexutil.Big `json:"epochNumber"`
		BlockHash    *common.Hash `json:"blockHash"`
		RequirePivot bool         `json:"requirePivot"`
	}

	var val tmpEpoch
	err = json.Unmarshal(data, &val)
	if err != nil {
		return err
	}

	if val.EpochNumber != nil && val.BlockHash != nil {
		return fmt.Errorf("cannot specify both BlockHash and EpochNumber, choose one or the other")
	}
	if val.EpochNumber != nil && val.RequirePivot {
		return fmt.Errorf("cannot specify both EpochNumber and RequirePivot, choose one or the other")
	}
	if val.EpochNumber != nil {
		e.epochNumber = val.EpochNumber
		return nil
	}
	e.blockHash = val.BlockHash
	e.requirePivot = val.RequirePivot
	return nil
}

// NewEpochOrBlockHashWithEpoch creates an instance of Epoch with specified epoch.
func NewEpochOrBlockHashWithEpoch(epoch *Epoch) *EpochOrBlockHash {
	if epoch == nil {
		epoch = EpochLatestCheckpoint
	}
	return &EpochOrBlockHash{epoch: epoch}
}

// NewEpochOrBlockHashWithBlockHash creates an instance of Epoch with specified block hash.
func NewEpochOrBlockHashWithBlockHash(blockHash Hash, requirePivot ...bool) *EpochOrBlockHash {
	if len(requirePivot) == 0 {
		requirePivot = append(requirePivot, false)
	}
	return &EpochOrBlockHash{nil, nil, blockHash.ToCommonHash(), requirePivot[0]}
}
