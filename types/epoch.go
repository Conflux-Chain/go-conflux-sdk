// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Const epoch definitions
var (
	EpochEarliest         *Epoch = &Epoch{"earliest", nil}
	EpochLatestCheckpoint *Epoch = &Epoch{"latest_checkpoint", nil}
	EpochLatestConfirmed  *Epoch = &Epoch{"latest_confirmed", nil}
	EpochLatestState      *Epoch = &Epoch{"latest_state", nil}
	EpochLatestMined      *Epoch = &Epoch{"latest_mined", nil}
)

// Epoch represents an epoch in Conflux.
type Epoch struct {
	name   string
	number *big.Int
}

// WebsocketEpochResponse represents result of epoch websocket subscription
type WebsocketEpochResponse struct {
	EpochHashesOrdered []Hash       `json:"epochHashesOrdered"`
	EpochNumber        *hexutil.Big `json:"epochNumber"`
}

// NewEpochNumber creates an instance of Epoch with specified number.
func NewEpochNumber(number *big.Int) *Epoch {
	return &Epoch{"", number}
}

// NewEpochWithBlockHash creates an instance of Epoch with specified block hash.
func NewEpochWithBlockHash(blockHash Hash) *Epoch {
	return &Epoch{string(blockHash), nil}
}

// String implements the fmt.Stringer interface
func (e *Epoch) String() string {
	if len(e.name) > 0 {
		return e.name
	}

	return hexutil.EncodeBig(e.number)
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e *Epoch) MarshalText() ([]byte, error) {
	return []byte(e.String()), nil
}
