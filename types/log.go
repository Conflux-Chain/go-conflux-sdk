// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package types

import (
	"encoding/json"
	"io"
	"math/big"
	"reflect"

	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/pkg/errors"
)

// LogFilter represents the filter of event in a smart contract.
type LogFilter struct {
	FromEpoch   *Epoch          `json:"fromEpoch,omitempty"`
	ToEpoch     *Epoch          `json:"toEpoch,omitempty"`
	FromBlock   *hexutil.Big    `json:"fromBlock,omitempty"`
	ToBlock     *hexutil.Big    `json:"toBlock,omitempty"`
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
	BlockHash           *Hash         `json:"blockHash,omitempty"`
	EpochNumber         *hexutil.Big  `json:"epochNumber,omitempty"`
	TransactionHash     *Hash         `json:"transactionHash,omitempty"`
	TransactionIndex    *hexutil.Big  `json:"transactionIndex,omitempty"`
	LogIndex            *hexutil.Big  `json:"logIndex,omitempty"`
	TransactionLogIndex *hexutil.Big  `json:"transactionLogIndex,omitempty"`
}

// rlpNilableBigInt nilable pointer to big int used for rlp encoding
type rlpNilableBigInt struct {
	Val *big.Int
}

// rlpEncodableLog log struct used for rlp encoding
type rlpEncodableLog struct {
	Address             Address
	Topics              []Hash
	Data                hexutil.Bytes
	BlockHash           *Hash             `rlp:"nil"`
	EpochNumber         *rlpNilableBigInt `rlp:"nil"`
	TransactionHash     *Hash             `rlp:"nil"`
	TransactionIndex    *rlpNilableBigInt `rlp:"nil"`
	LogIndex            *rlpNilableBigInt `rlp:"nil"`
	TransactionLogIndex *rlpNilableBigInt `rlp:"nil"`
}

func (l Log) toRlpEncodable() rlpEncodableLog {
	rlog := rlpEncodableLog{
		Address: l.Address, Topics: l.Topics, Data: l.Data,
		BlockHash: l.BlockHash, TransactionHash: l.TransactionHash,
	}

	if l.EpochNumber != nil {
		rlog.EpochNumber = &rlpNilableBigInt{l.EpochNumber.ToInt()}
	}

	if l.TransactionIndex != nil {
		rlog.TransactionIndex = &rlpNilableBigInt{l.TransactionIndex.ToInt()}
	}

	if l.LogIndex != nil {
		rlog.LogIndex = &rlpNilableBigInt{l.LogIndex.ToInt()}
	}

	if l.TransactionLogIndex != nil {
		rlog.TransactionLogIndex = &rlpNilableBigInt{l.TransactionLogIndex.ToInt()}
	}
	return rlog
}

func (r rlpEncodableLog) toNormal() Log {
	log := Log{}

	log.Address, log.Topics, log.Data = r.Address, r.Topics, r.Data
	log.BlockHash, log.TransactionHash = r.BlockHash, r.TransactionHash

	if r.EpochNumber != nil {
		log.EpochNumber = (*hexutil.Big)(r.EpochNumber.Val)
	}

	if r.TransactionIndex != nil {
		log.TransactionIndex = (*hexutil.Big)(r.TransactionIndex.Val)
	}

	if r.LogIndex != nil {
		log.LogIndex = (*hexutil.Big)(r.LogIndex.Val)
	}

	if r.TransactionLogIndex != nil {
		log.TransactionLogIndex = (*hexutil.Big)(r.TransactionLogIndex.Val)
	}
	return log
}

// EncodeRLP implements the rlp.Encoder interface.
func (log Log) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, log.toRlpEncodable())
}

// DecodeRLP implements the rlp.Decoder interface.
func (log *Log) DecodeRLP(r *rlp.Stream) error {
	var rlog rlpEncodableLog
	if err := r.Decode(&rlog); err != nil {
		return err
	}

	*log = rlog.toNormal()
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (l *LogFilter) UnmarshalJSON(data []byte) error {
	type tmpLogFilter struct {
		FromEpoch   *Epoch          `json:"fromEpoch,omitempty"`
		ToEpoch     *Epoch          `json:"toEpoch,omitempty"`
		FromBlock   *hexutil.Big    `json:"fromBlock,omitempty"`
		ToBlock     *hexutil.Big    `json:"toBlock,omitempty"`
		BlockHashes []Hash          `json:"blockHashes,omitempty"`
		Address     interface{}     `json:"address,omitempty"`
		Topics      []interface{}   `json:"topics,omitempty"`
		Offset      *hexutil.Uint64 `json:"offset,omitempty"`
		Limit       *hexutil.Uint64 `json:"limit,omitempty"`
	}

	t := tmpLogFilter{}
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}

	var err error
	l.FromEpoch = t.FromEpoch
	l.ToEpoch = t.ToEpoch
	l.FromBlock = t.FromBlock
	l.ToBlock = t.ToBlock
	l.BlockHashes = t.BlockHashes
	l.Offset = t.Offset
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

type SubscriptionLog struct {
	*Log
	*ChainReorg
}

type rlpEncodableSubscriptionLog struct {
	Log        *rlpEncodableLog        `rlp:"nil"`
	ChainReorg *rlpEncodableChainReorg `rlp:"nil"`
}

func (s SubscriptionLog) IsRevertLog() bool {
	return s.ChainReorg != nil
	// return !reflect.DeepEqual(s.ChainReorg, ChainReorg{})
}

func (s SubscriptionLog) MarshalJSON() ([]byte, error) {
	if s.IsRevertLog() {
		return json.Marshal(s.ChainReorg)
	}
	return json.Marshal(s.Log)
}

func (s SubscriptionLog) toRlpEncodable() rlpEncodableSubscriptionLog {
	var rlpLog *rlpEncodableLog
	if s.Log != nil {
		_rlpLog := s.Log.toRlpEncodable()
		rlpLog = &_rlpLog
	}

	var rlpReorg *rlpEncodableChainReorg
	if s.ChainReorg != nil {
		_rlpReorg := s.ChainReorg.toRlpEncodable()
		rlpReorg = &_rlpReorg
	}

	r := rlpEncodableSubscriptionLog{rlpLog, rlpReorg}
	return r
}

func (r rlpEncodableSubscriptionLog) toNormal() SubscriptionLog {
	slog := SubscriptionLog{}
	if r.Log != nil {
		_log := r.Log.toNormal()
		slog.Log = &_log
	}

	if r.ChainReorg != nil {
		_reorg := r.ChainReorg.toNormal()
		slog.ChainReorg = &_reorg
	}
	return slog
}

func (s SubscriptionLog) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, s.toRlpEncodable())
}

func (s *SubscriptionLog) DecodeRLP(r *rlp.Stream) error {
	rlpSubLog := rlpEncodableSubscriptionLog{}
	if err := r.Decode(&rlpSubLog); err != nil {
		return err
	}
	*s = rlpSubLog.toNormal()
	return nil
}
