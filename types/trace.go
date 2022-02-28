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
type CallType string
type PocketType string
type CreateType string
type SpaceType string
type ActionType string
type OutcomeType string

const (
	TRACE_CALL                      TraceType = "call"
	TRACE_CALL_RESULT               TraceType = "call_result"
	TRACE_CREATE                    TraceType = "create"
	TRACE_CREATE_RESULT             TraceType = "create_result"
	TRACE_INTERNAL_TRANSFER_ACTIION TraceType = "internal_transfer_action"
)

const (
	CALL_NONE         CallType = "none"
	CALL_CALL         CallType = "call"
	CALL_CALLCODE     CallType = "callcode"
	CALL_DELEGATECALL CallType = "delegatecall"
	CALL_STATICCALL   CallType = "staticcall"
)

const (
	POCKET_BALANCE                     PocketType = "balance"
	POCKET_STAKING_BALANCE             PocketType = "staking_balance"
	POCKET_STORAGE_COLLATERAL          PocketType = "storage_collateral"
	POCKET_SPONSOR_BALANCE_FOR_GAS     PocketType = "sponsor_balance_for_gas"
	POCKET_SPONSOR_BALANCE_FOR_STORAGE PocketType = "sponsor_balance_for_collateral"
	POCKET_MINT_BURN                   PocketType = "mint_or_burn"
	POCKET_GAS_PAYMENT                 PocketType = "gas_payment"
)

const (
	CREATE_NONE    CreateType = "none"
	CREATE_CREATE  CreateType = "create"
	CREATE_CREATE2 CreateType = "create2"
)

const (
	SPACE_EVM    SpaceType = "evm"
	SPACE_NATIVE SpaceType = "native"
	SPACE_NONE   SpaceType = "none"
)

const (
	OUTCOME_SUCCESS  OutcomeType = "success"
	OUTCOME_REVERTED OutcomeType = "reverted"
	OUTCOME_FAIL     OutcomeType = "fail"
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
	Valid               bool            `json:"valid"`
	Type                TraceType       `json:"type"`
	EpochHash           *Hash           `json:"epochHash,omitempty"`
	EpochNumber         *hexutil.Big    `json:"epochNumber,omitempty"`
	BlockHash           *Hash           `json:"blockHash,omitempty"`
	TransactionPosition *hexutil.Uint64 `json:"transactionPosition,omitempty"`
	TransactionHash     *Hash           `json:"transactionHash,omitempty"`
}

// independent action structs
type Call struct {
	Space    SpaceType     `json:"space"`
	From     Address       `json:"from"`
	To       Address       `json:"to"`
	Value    hexutil.Big   `json:"value"`
	Gas      hexutil.Big   `json:"gas"`
	Input    hexutil.Bytes `json:"input"`
	CallType CallType      `json:"callType"`
}

type Create struct {
	Space      SpaceType     `json:"space"`
	From       Address       `json:"from"`
	Value      hexutil.Big   `json:"value"`
	Gas        hexutil.Big   `json:"gas"`
	Init       hexutil.Bytes `json:"init"`
	CreateType CreateType    `json:"createType"`
}

type CallResult struct {
	Outcome    OutcomeType   `json:"outcome"`
	GasLeft    hexutil.Big   `json:"gasLeft"`
	ReturnData hexutil.Bytes `json:"returnData"`
}

type CreateResult struct {
	Outcome    OutcomeType   `json:"outcome"`
	Addr       Address       `json:"addr"`
	GasLeft    hexutil.Big   `json:"gasLeft"`
	ReturnData hexutil.Bytes `json:"returnData"`
}

type InternalTransferAction struct {
	From       Address     `json:"from"`
	FromPocket PocketType  `json:"fromPocket"`
	FromSpace  SpaceType   `json:"fromSpace"`
	To         Address     `json:"to"`
	ToPocket   PocketType  `json:"toPocket"`
	ToSpace    SpaceType   `json:"toSpace"`
	Value      hexutil.Big `json:"value"`
}

type TraceFilter struct {
	FromEpoch   *Epoch    `json:"fromEpoch"`
	ToEpoch     *Epoch    `json:"toEpoch"`
	FromAddress []Address `json:"fromAddress"`
	ToAddress   []Address `json:"toAddress"`
	BlockHashes []Hash    `json:"blockHashes"`
	// action types could be "call","create","callResult","createResult","internalTransferAction"
	ActionTypes []TraceType     `json:"actionTypes"`
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

	type alias LocalizedTrace

	a := alias{}
	err := json.Unmarshal(data, &a)
	if err != nil {
		return err
	}
	*l = LocalizedTrace(a)

	tmp := struct {
		Action map[string]interface{} `json:"action"`
	}{}

	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	var action interface{}
	if action, err = parseAction(tmp.Action, l.Type); err != nil {
		return err
	}
	l.Action = action
	return nil
}

func parseAction(actionInMap map[string]interface{}, actionType TraceType) (interface{}, error) {
	// actionKeys := utils.GetMapSortedKeys(actionInMap)

	// newActionType := actionKeysToTypeMap[strings.Join(actionKeys, "")]
	// if newActionType == "" {
	// 	return nil, fmt.Errorf("uncongonized action type with fields %v", actionKeys)
	// }

	actionJson, err := json.Marshal(actionInMap)
	if err != nil {
		return nil, err
	}

	var result interface{}
	switch actionType {
	case TRACE_CALL:
		action := Call{}
		err = json.Unmarshal(actionJson, &action)
		result = action
	case TRACE_CREATE:
		action := Create{}
		err = json.Unmarshal(actionJson, &action)
		result = action
	case TRACE_CALL_RESULT:
		action := CallResult{}
		err = json.Unmarshal(actionJson, &action)
		result = action
	case TRACE_CREATE_RESULT:
		action := CreateResult{}
		err = json.Unmarshal(actionJson, &action)
		result = action
	case TRACE_INTERNAL_TRANSFER_ACTIION:
		action := InternalTransferAction{}
		err = json.Unmarshal(actionJson, &action)
		result = action
	default:
		return nil, fmt.Errorf("unknown action type %v", actionType)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal %v to %v ", string(actionJson), actionType)
	}

	return result, nil
}

type LocalizedTraceTire []*LocalizedTraceNode

type LocalizedTraceNode struct {
	Type                TraceType       `json:"type"`
	Valid               bool            `json:"valid"`
	EpochHash           *Hash           `json:"epochHash,omitempty"`
	EpochNumber         *hexutil.Big    `json:"epochNumber,omitempty"`
	BlockHash           *Hash           `json:"blockHash,omitempty"`
	TransactionPosition *hexutil.Uint64 `json:"transactionPosition,omitempty"`
	TransactionHash     *Hash           `json:"transactionHash,omitempty"`

	CallWithResult         *TraceCallWithResult    `json:"callWithResult,omitempty"`
	CreateWithResult       *TraceCreateWithResult  `json:"createWithResult,omitempty"`
	InternalTransferAction *InternalTransferAction `json:"internalTransferAction,omitempty"`

	Childs []*LocalizedTraceNode `json:"childs"`
	// Raw    *LocalizedTrace       `json:"raw"`
}

func (l *LocalizedTraceNode) populate(raw LocalizedTrace) {
	l.Type = raw.Type
	l.Valid = raw.Valid
	l.EpochHash = raw.EpochHash
	l.EpochNumber = raw.EpochNumber
	l.BlockHash = raw.BlockHash
	l.TransactionPosition = raw.TransactionPosition
	l.TransactionHash = raw.TransactionHash
}

// TraceInTire convert flattened trace to trie, the convered tiers are stored in the order of the flattened trace like follow.
//
// 		InternalTransfer
// 		call
// 		create
// 		createResult
// 		callResult
// 		InternalTransfer
// 		============>
// 		InternalTransfer
// 		call + callResult
// 			|- create + createResult
// 		InternalTransfer
func TraceInTire(traces []LocalizedTrace) (tier LocalizedTraceTire, err error) {
	cacheStack := new([]*LocalizedTraceNode)

	for _, v := range traces {

		if len(*cacheStack) == 0 {
			n, err := newLocalizedTraceNode(v, cacheStack)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			tier = append(tier, n)
			continue
		}

		lastOpeningNode := (*cacheStack)[len(*cacheStack)-1]
		if v.Type == TRACE_INTERNAL_TRANSFER_ACTIION {
			item, err := newLocalizedTraceNode(v, nil)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			lastOpeningNode.Childs = append(lastOpeningNode.Childs, item)
			continue
		}

		if v.Type == TRACE_CALL || v.Type == TRACE_CREATE {
			item, err := newLocalizedTraceNode(v, cacheStack)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			lastOpeningNode.Childs = append(lastOpeningNode.Childs, item)
			continue
		}

		// call result or create result
		if lastOpeningNode.Type == TRACE_CALL {
			if v.Type != TRACE_CALL_RESULT {
				return nil, fmt.Errorf("expect trace type CallResult, got %v", v.Type)
			}
			cr := v.Action.(CallResult)
			lastOpeningNode.CallWithResult.CallResult = &cr
			*cacheStack = (*cacheStack)[:len(*cacheStack)-1]
		}

		if lastOpeningNode.Type == TRACE_CREATE {
			if v.Type != TRACE_CREATE_RESULT {
				return nil, fmt.Errorf("expect trace type CreateResult, got %v", v.Type)
			}
			cr := v.Action.(CreateResult)
			lastOpeningNode.CreateWithResult.CreateResult = &cr
			*cacheStack = (*cacheStack)[:len(*cacheStack)-1]
		}
	}
	// push call trace, set to child when meet next call trace
	// pop call trace when meet call result trace
	return tier, nil
}

func newLocalizedTraceNode(trace LocalizedTrace, cacheStack *[]*LocalizedTraceNode,
) (*LocalizedTraceNode, error) {
	switch trace.Type {
	case TRACE_CALL:
		action := trace.Action.(Call)
		node := &LocalizedTraceNode{CallWithResult: &TraceCallWithResult{
			&action, nil,
		}}
		node.populate(trace)
		*cacheStack = append(*cacheStack, node)
		return node, nil
	case TRACE_CREATE:
		action := trace.Action.(Create)
		node := &LocalizedTraceNode{CreateWithResult: &TraceCreateWithResult{
			&action, nil,
		}}
		node.populate(trace)
		*cacheStack = append(*cacheStack, node)
		return node, nil
	case TRACE_INTERNAL_TRANSFER_ACTIION:
		action := trace.Action.(InternalTransferAction)
		node := &LocalizedTraceNode{InternalTransferAction: &action}
		node.populate(trace)
		return node, nil
	}
	return nil, fmt.Errorf("could not create new localized trace node by type %v", trace.Type)
}

func (t LocalizedTraceTire) Flatten() (flattened []*LocalizedTraceNode) {
	for _, v := range t {
		flattened = append(flattened, v.Flatten()...)
	}
	return flattened
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
	*Call       `json:"call"`
	*CallResult `json:"callResult"`
}

type TraceCreateWithResult struct {
	*Create       `json:"create"`
	*CreateResult `json:"createResult"`
}
