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
	FromEpoch   *Epoch    `json:"fromEpoch,omitempty"`
	ToEpoch     *Epoch    `json:"toEpoch,omitempty"`
	BlockHashes []Hash    `json:"blockHashes,omitempty"`
	Address     []Address `json:"address,omitempty"`
	Topics      [][]Hash  `json:"topics,omitempty"`
	Limit       *uint8    `json:"limit,omitempty"`
}

// LogEntry represents a summary of event in a smart contract.
type LogEntry struct {
	Address Address `json:"address"`
	Topics  []Hash  `json:"topics"`
	Data    string  `json:"data"`
}

// Log represents the event in a smart contract
type Log struct {
	LogEntry
	BlockHash           *Hash        `json:"blockHash,omitempty"`
	EpochNumber         *hexutil.Big `json:"epochNumber,omitempty"`
	TransactionHash     *Hash        `json:"transactionHash,omitempty"`
	TransactionIndex    *hexutil.Big `json:"transactionIndex,omitempty"`
	LogIndex            *hexutil.Big `json:"logIndex,omitempty"`
	TransactionLogIndex *hexutil.Big `json:"transactionLogIndex,omitempty"`
}

type SubscriptionLog struct {
	Log
	ChainReorg
}

func (s SubscriptionLog) IsRevertLog() bool {
	return !reflect.DeepEqual(s.ChainReorg, ChainReorg{})
}
