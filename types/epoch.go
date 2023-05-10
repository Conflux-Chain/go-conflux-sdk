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

// // NewEpochWithBlockHash creates an instance of Epoch with specified block hash.
// func NewEpochWithBlockHash(blockHash Hash, requirePivot ...bool) *Epoch {
// 	if len(requirePivot) == 0 {
// 		requirePivot = append(requirePivot, false)
// 	}
// 	return &Epoch{"", nil, blockHash.ToCommonHash(), requirePivot[0]}
// }

// String implements the fmt.Stringer interface
func (e *Epoch) String() string {
	if e.number != nil {
		return e.number.String()
	}

	return e.name
}

// @depercated
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
	BlockHash    *common.Hash
	RequirePivot bool
}

// String implements the fmt.Stringer interface
func (e *EpochOrBlockHash) String() string {
	if e.epoch != nil {
		return e.epoch.String()
	}

	if e.epochNumber != nil {
		return e.epochNumber.String()
	}

	if e.BlockHash != nil {
		return e.BlockHash.String()
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

	if e.BlockHash != nil {
		return json.Marshal(struct {
			BlockHash    common.Hash `json:"blockHash"`
			RequirePivot bool        `json:"requirePivot"`
		}{
			BlockHash:    *e.BlockHash,
			RequirePivot: e.RequirePivot,
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
	if err == nil {
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
		e.BlockHash = val.BlockHash
		e.RequirePivot = val.RequirePivot
		return nil
	}
	return err

	// var input string
	// err = json.Unmarshal(data, &input)
	// if err != nil {
	// 	return err
	// }

	// switch input {
	// case EpochEarliest.name:
	// 	fallthrough
	// case EpochLatestCheckpoint.name:
	// 	fallthrough
	// case EpochLatestConfirmed.name:
	// 	fallthrough
	// case EpochLatestState.name:
	// 	fallthrough
	// case EpochLatestMined.name:
	// 	fallthrough
	// case EpochLatestFinalized.name:
	// 	e.name = input
	// 	return nil
	// default:
	// 	if len(input) == 66 {
	// 		hash := common.Hash{}
	// 		err := hash.UnmarshalText([]byte(input))
	// 		if err != nil {
	// 			return err
	// 		}
	// 		e.BlockHash = &hash
	// 		return nil
	// 	} else {
	// 		blckNum, err := hexutil.DecodeBig(input)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		e.number = NewBigIntByRaw(blckNum)
	// 		return nil
	// 	}
	// }

}

// NewEpochWithBlockHash creates an instance of Epoch with specified epoch.
func NewEpochOrBlockHashWithEpoch(epoch Epoch) *EpochOrBlockHash {
	return &EpochOrBlockHash{epoch: &epoch}
}

// NewEpochWithBlockHash creates an instance of Epoch with specified block hash.
func NewEpochOrBlockHashWithBlockHash(blockHash Hash, requirePivot ...bool) *EpochOrBlockHash {
	if len(requirePivot) == 0 {
		requirePivot = append(requirePivot, false)
	}
	return &EpochOrBlockHash{nil, nil, blockHash.ToCommonHash(), requirePivot[0]}
}
