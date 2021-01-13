// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package types

import (
	"reflect"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// LogFilter represents the filter of event in a smart contract.
type LogFilter struct {
	FromEpoch   *Epoch          `json:"fromEpoch,omitempty"`
	ToEpoch     *Epoch          `json:"toEpoch,omitempty"`
	BlockHashes []Hash          `json:"blockHashes,omitempty"`
	Address     []Address       `json:"address,omitempty"`
	Topics      [][]Hash        `json:"topics,omitempty"`
	Limit       *hexutil.Uint64 `json:"limit,omitempty"`
}

// Log represents the event in a smart contract
type Log struct {
	Address             Address       `json:"address"`
	Topics              []Hash        `json:"topics"`
	Data                hexutil.Bytes `json:"data"`
	BlockHash           *Hash         `json:"blockHash"`
	EpochNumber         *hexutil.Big  `json:"epochNumber"`
	TransactionHash     *Hash         `json:"transactionHash"`
	TransactionIndex    *hexutil.Big  `json:"transactionIndex"`
	LogIndex            *hexutil.Big  `json:"logIndex"`
	TransactionLogIndex *hexutil.Big  `json:"transactionLogIndex"`
}

type SubscriptionLog struct {
	Log
	ChainReorg
}

func (s SubscriptionLog) IsRevertLog() bool {
	return !reflect.DeepEqual(s.ChainReorg, ChainReorg{})
}
