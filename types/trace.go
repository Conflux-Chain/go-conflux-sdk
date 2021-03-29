package types

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
)

type LocalizedBlockTrace struct {
	TransactionTraces []LocalizedTransactionTrace `json:"transactionTraces"`
	EpochHash         Hash                        `json:"epochHash"`
	EpochNumber       hexutil.Big                 `json:"epochNumber"`
	BlockHash         Hash                        `json:"blockHash"`
}

type LocalizedTransactionTrace struct {
	Traces              []LocalizedTrace `json:"traces"`
	TransactionPosition hexutil.Uint64   `json:"transactionPosition"`
	TransactionHash     Hash             `json:"transactionHash"`
}

type LocalizedTrace struct {
	Action              interface{}     `json:"action"`
	Type                string          `json:"type"`
	EpochHash           *Hash           `json:"epochHash,omitempty"`
	EpochNumber         *hexutil.Big    `json:"epochNumber,omitempty"`
	BlockHash           *Hash           `json:"blockHash,omitempty"`
	TransactionPosition *hexutil.Uint64 `json:"transactionPosition,omitempty"`
	TransactionHash     *Hash           `json:"transactionHash,omitempty"`
}

// independent action structs
type Call struct {
	From     Address       `json:"from"`
	To       Address       `json:"to"`
	Value    hexutil.Big   `json:"value"`
	Gas      hexutil.Big   `json:"gas"`
	Input    hexutil.Bytes `json:"input"`
	CallType string        `json:"callType"`
}

type Create struct {
	From  Address       `json:"from"`
	Value hexutil.Big   `json:"value"`
	Gas   hexutil.Big   `json:"gas"`
	Init  hexutil.Bytes `json:"init"`
}

type CallResult struct {
	Outcome    string        `json:"outcome"`
	GasLeft    hexutil.Big   `json:"gasLeft"`
	ReturnData hexutil.Bytes `json:"returnData"`
}

type CreateResult struct {
	Outcome    string        `json:"outcome"`
	Addr       Address       `json:"addr"`
	GasLeft    hexutil.Big   `json:"gasLeft"`
	ReturnData hexutil.Bytes `json:"returnData"`
}

type InternalTransferAction struct {
	From  Address     `json:"from"`
	To    Address     `json:"to"`
	Value hexutil.Big `json:"value"`
}

type TraceFilter struct {
	FromEpoch   *Epoch `json:"fromEpoch"`
	ToEpoch     *Epoch `json:"toEpoch"`
	BlockHashes []Hash `json:"blockHashes"`
	// action types could be "call","create","callResult","createResult","internalTransferAction"
	ActionTypes []string        `json:"actionTypes"`
	After       *hexutil.Uint64 `json:"after"`
	Count       *hexutil.Uint64 `json:"count"`
}

var actionKeysToTypeMap = make(map[string]string)

func init() {
	// get all action object keys
	actionObjs := []interface{}{
		Call{}, Create{}, CallResult{}, CreateResult{}, InternalTransferAction{},
	}

	for _, v := range actionObjs {
		tags := utils.GetObjJsonFieldTags(v)
		actionKeysToTypeMap[strings.Join(tags, "")] = reflect.TypeOf(v).Name()
	}
}

// UnmarshalJSON unmarshals Input and Init type from []byte to hexutil.Bytes
func (l *LocalizedTrace) UnmarshalJSON(data []byte) error {

	tmp := struct {
		Action              map[string]interface{} `json:"action"`
		Type                string                 `json:"type"`
		EpochHash           *Hash                  `json:"epochHash"`
		EpochNumber         *hexutil.Big           `json:"epochNumber"`
		BlockHash           *Hash                  `json:"blockHash"`
		TransactionPosition *hexutil.Uint64        `json:"transactionPosition"`
		TransactionHash     *Hash                  `json:"transactionHash"`
	}{}

	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	l.Type = tmp.Type
	l.EpochHash = tmp.EpochHash
	l.EpochNumber = tmp.EpochNumber
	l.BlockHash = tmp.BlockHash
	l.TransactionPosition = tmp.TransactionPosition
	l.TransactionHash = tmp.TransactionHash

	var action interface{}
	if action, err = parseAction(tmp.Action); err != nil {
		return err
	}
	l.Action = action

	return nil
}

func parseAction(actionInMap map[string]interface{}) (interface{}, error) {
	actionKeys := utils.GetMapSortedKeys(actionInMap)

	newActionType := actionKeysToTypeMap[strings.Join(actionKeys, "")]
	if newActionType == "" {
		return nil, fmt.Errorf("uncongonized action type with fields %v", actionKeys)
	}

	actionJson, err := json.Marshal(actionInMap)
	if err != nil {
		return nil, err
	}

	var result interface{}
	switch newActionType {
	case "Call":
		action := Call{}
		err = json.Unmarshal(actionJson, &action)
		result = action
	case "Create":
		action := Create{}
		err = json.Unmarshal(actionJson, &action)
		result = action
	case "CallResult":
		action := CallResult{}
		err = json.Unmarshal(actionJson, &action)
		result = action
	case "CreateResult":
		action := CreateResult{}
		err = json.Unmarshal(actionJson, &action)
		result = action
	case "InternalTransferAction":
		action := InternalTransferAction{}
		err = json.Unmarshal(actionJson, &action)
		result = action
	}
	if err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal %v to %v ", string(actionJson), newActionType)
	}

	return result, nil
}
