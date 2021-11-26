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

type TraceType string

const (
	CALL_TYPE                      TraceType = "call"
	CALL_RESULT_TYPE               TraceType = "call_result"
	CREATE_TYPE                    TraceType = "create"
	CREATE_RESULT_TYPE             TraceType = "create_result"
	INTERNAL_TRANSFER_ACTIION_TYPE TraceType = "internal_transfer_action"
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
	Type                TraceType       `json:"type"`
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
		Type                TraceType              `json:"type"`
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

// TODO: should move to go sdk=========================
type LocalizedTraceNode struct {
	Childs []*LocalizedTraceNode `json:"childs"`
	Raw    *LocalizedTrace       `json:"raw"`

	CallWithResult         *TraceCallWithResult    `json:"callWithResult,omitempty"`
	CreateWithResult       *TraceCreateWithResult  `json:"createWithResult,omitempty"`
	InternalTransferAction *InternalTransferAction `json:"internalTransferAction,omitempty"`
}

func TraceInTree(traces []LocalizedTrace) (node *LocalizedTraceNode, err error) {
	cacheStack := new([]*LocalizedTraceNode)

	for _, v := range traces {
		if node == nil {
			node, err = newLocalizedTraceNode(v, cacheStack)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			continue
		}

		lastOpendingNode := (*cacheStack)[len(*cacheStack)-1]
		if v.Type == INTERNAL_TRANSFER_ACTIION_TYPE {
			item, err := newLocalizedTraceNode(v, nil)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			lastOpendingNode.Childs = append(lastOpendingNode.Childs, item)
			continue
		}

		if v.Type == CALL_TYPE || v.Type == CREATE_TYPE {
			item, err := newLocalizedTraceNode(v, cacheStack)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			lastOpendingNode.Childs = append(lastOpendingNode.Childs, item)
			continue
		}

		// call result or create result
		if lastOpendingNode.Raw.Type == CALL_TYPE {
			if v.Type != CALL_RESULT_TYPE {
				return nil, fmt.Errorf("expect trace type CallResult, got %v", v.Type)
			}
			cr := v.Action.(CallResult)
			lastOpendingNode.CallWithResult.CallResult = &cr
			*cacheStack = (*cacheStack)[:len(*cacheStack)-1]
		}

		if lastOpendingNode.Raw.Type == CREATE_TYPE {
			if v.Type != CREATE_RESULT_TYPE {
				return nil, fmt.Errorf("expect trace type CreateResult, got %v", v.Type)
			}
			cr := v.Action.(CreateResult)
			lastOpendingNode.CreateWithResult.CreateResult = &cr
			*cacheStack = (*cacheStack)[:len(*cacheStack)-1]
		}
	}
	// push call trace, set to child when meet next call trace
	// pop call trace when meet call result trace
	return node, nil
}

func newLocalizedTraceNode(trace LocalizedTrace, cacheStack *[]*LocalizedTraceNode,
) (*LocalizedTraceNode, error) {
	switch trace.Type {
	case CALL_TYPE:
		action := trace.Action.(Call)
		node := &LocalizedTraceNode{Raw: &trace, CallWithResult: &TraceCallWithResult{
			&action, nil,
		}}
		*cacheStack = append(*cacheStack, node)
		return node, nil
	case CREATE_TYPE:
		action := trace.Action.(Create)
		node := &LocalizedTraceNode{Raw: &trace, CreateWithResult: &TraceCreateWithResult{
			&action, nil,
		}}
		*cacheStack = append(*cacheStack, node)
		return node, nil
	case INTERNAL_TRANSFER_ACTIION_TYPE:
		action := trace.Action.(InternalTransferAction)
		return &LocalizedTraceNode{Raw: &trace, InternalTransferAction: &action}, nil
	}
	return nil, fmt.Errorf("could not create new localized trace node by type %v", trace.Type)
}

func (l LocalizedTraceNode) Flatten() (flattened []*LocalizedTraceNode) {
	flattened = append(flattened, &l)
	for _, v := range l.Childs {
		flattened = append(flattened, v.Flatten()...)
	}
	// clear childs
	l.Childs = nil
	return
}

type TraceCallWithResult struct {
	*Call
	*CallResult
}

type TraceCreateWithResult struct {
	*Create
	*CreateResult
}
