// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package types

import (
	"encoding/json"
	"reflect"

	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
)

// LogFilter represents the filter of event in a smart contract.
type LogFilter struct {
	FromEpoch   *Epoch          `json:"fromEpoch,omitempty"`
	ToEpoch     *Epoch          `json:"toEpoch,omitempty"`
	BlockHashes []Hash          `json:"blockHashes,omitempty"`
	Address     []Address       `json:"address,omitempty"`
	Topics      [][]Hash        `json:"topics,omitempty"`
	Offset      *hexutil.Uint64 `json:"offset,omitempty"`
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

// UnmarshalJSON implements the json.Unmarshaler interface.
func (l *LogFilter) UnmarshalJSON(data []byte) error {
	type tmpLogFilter struct {
		FromEpoch   *Epoch          `json:"fromEpoch,omitempty"`
		ToEpoch     *Epoch          `json:"toEpoch,omitempty"`
		BlockHashes []Hash          `json:"blockHashes,omitempty"`
		Address     interface{}     `json:"address,omitempty"`
		Topics      []interface{}   `json:"topics,omitempty"`
		Limit       *hexutil.Uint64 `json:"limit,omitempty"`
	}

	t := tmpLogFilter{}
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}

	var err error
	l.FromEpoch = t.FromEpoch
	l.ToEpoch = t.ToEpoch
	l.BlockHashes = t.BlockHashes
	l.Limit = t.Limit
	if l.Address, err = resolveToAddresses(t.Address); err != nil {
		return err
	}
	if l.Topics, err = resolveToTopicsList(t.Topics); err != nil {
		return err
	}
	return nil
}

func resolveToAddresses(val interface{}) ([]Address, error) {
	// if val is nil, return
	if val == nil {
		return nil, nil
	}

	// if val is string, new address and return
	if addrStr, ok := val.(string); ok {
		addr, err := cfxaddress.NewFromBase32(addrStr)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create address by %v", addrStr)
		}
		return []Address{addr}, nil
	}

	// if val is string slice, new every item to cfxaddress
	if addrStrList, ok := val.([]interface{}); ok {
		addrList := make([]Address, 0)
		for _, v := range addrStrList {
			vStr, ok := v.(string)
			if !ok {
				return nil, errors.Errorf("could not conver type %v to address", reflect.TypeOf(v))
			}

			addr, err := cfxaddress.NewFromBase32(vStr)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to create address by %v", v)
			}
			addrList = append(addrList, addr)
		}
		return addrList, nil
	}

	return nil, errors.Errorf("failed to unmarshal %#v to address or address list", val)
}

func resolveToTopicsList(val []interface{}) ([][]Hash, error) {
	// if val is nil, return
	if val == nil {
		return nil, nil
	}

	// otherwise, convert every item to topics
	topicsList := make([][]Hash, 0)

	for _, v := range val {
		hashes, err := resolveToHashes(v)
		if err != nil {
			return nil, err
		}
		topicsList = append(topicsList, hashes)
	}
	return topicsList, nil
}

func resolveToHashes(val interface{}) ([]Hash, error) {
	// if val is nil, return
	if val == nil {
		return nil, nil
	}

	// if val is string, return
	if hashStr, ok := val.(string); ok {
		return []Hash{Hash(hashStr)}, nil
	}

	// if val is string slice, append every item
	if addrStrList, ok := val.([]interface{}); ok {
		addrList := make([]Hash, 0)
		for _, v := range addrStrList {
			vStr, ok := v.(string)
			if !ok {
				return nil, errors.Errorf("could not conver type %v to hash", reflect.TypeOf(v))
			}

			addrList = append(addrList, Hash(vStr))
		}
		return addrList, nil
	}

	return nil, errors.Errorf("failed to convert %v to hash or hashes", val)
}
