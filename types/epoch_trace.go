package types

import (
	"encoding/json"

	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
)

type EthTraceCall struct {
	From     common.Address `json:"from"`
	To       common.Address `json:"to"`
	Value    hexutil.Big    `json:"value"`
	Gas      hexutil.Big    `json:"gas"`
	Input    hexutil.Bytes  `json:"input"`
	CallType CallType       `json:"callType"`
}

type EthTraceCreate struct {
	From       common.Address `json:"from"`
	Value      hexutil.Big    `json:"value"`
	Gas        hexutil.Big    `json:"gas"`
	Init       hexutil.Bytes  `json:"input"`
	CreateType CreateType     `json:"createType"`
}

type EthTraceAction struct {
	traceype TraceType
	Call     *EthTraceCall
	Create   *EthTraceCreate
}

func (e EthTraceAction) MarshalJSON() ([]byte, error) {
	switch e.traceype {
	case TRACE_CALL:
		return json.Marshal(e.Call)
	case TRACE_CREATE:
		return json.Marshal(e.Create)
	}
	return nil, errors.Errorf("unsupport trace type %s", e.traceype)
}

func (e *EthTraceAction) UnmarshalJSON(data []byte) error {
	switch e.traceype {
	case TRACE_CALL:
		return json.Unmarshal(data, &e.Call)
	case TRACE_CREATE:
		return json.Unmarshal(data, &e.Create)
	}
	return errors.Errorf("unsupport trace type %s", e.traceype)
}

type EthTraceCallResult struct {
	GasUsed hexutil.Big   `json:"gasUsed"`
	Output  hexutil.Bytes `json:"output"`
}

type EthTraceCreateResult struct {
	GasUsed hexutil.Big    `json:"gasUsed"`
	Code    hexutil.Bytes  `json:"code"`
	Address common.Address `json:"address"`
}

type EthTraceRes struct {
	traceype TraceType
	Call     *EthTraceCallResult
	Create   *EthTraceCreateResult
}

func (e EthTraceRes) MarshalJSON() ([]byte, error) {
	switch e.traceype {
	case TRACE_CALL:
		return json.Marshal(e.Call)
	case TRACE_CREATE:
		return json.Marshal(e.Create)
	}
	return nil, errors.Errorf("unsupport trace type %s", e.traceype)
}

func (e *EthTraceRes) UnmarshalJSON(data []byte) error {
	switch e.traceype {
	case TRACE_CALL:
		return json.Unmarshal(data, &e.Call)
	case TRACE_CREATE:
		return json.Unmarshal(data, &e.Create)
	}
	return errors.Errorf("unsupport trace type %s", e.traceype)
}

type EthLocalizedTrace struct {
	Type                TraceType        `json:"type"`
	Action              EthTraceAction   `json:"action"`
	Result              *EthTraceRes     `json:"result,omitempty"`
	Error               string           `json:"error,omitempty"`
	TraceAddress        []hexutil.Uint64 `json:"traceAddress"`
	Subtraces           hexutil.Uint64   `json:"subtraces"`
	TransactionPosition *hexutil.Uint64  `json:"transactionPosition"`
	TransactionHash     *Hash            `json:"transactionHash"`
	BlockNumber         hexutil.Uint64   `json:"blockNumber"`
	BlockHash           common.Hash      `json:"blockHash"`
	Valid               bool             `json:"valid"`
}

func (e EthLocalizedTrace) MarshalJSON() ([]byte, error) {
	e.Action.traceype = e.Type
	if e.Result != nil {
		e.Result.traceype = e.Type
	}
	type Alias EthLocalizedTrace
	return json.Marshal(Alias(e))
}

func (e *EthLocalizedTrace) UnmarshalJSON(data []byte) error {
	type Alias EthLocalizedTrace
	type Tmp struct {
		Alias
		Action interface{} `json:"action"`
		Result interface{} `json:"result,omitempty"`
	}

	var tmp Tmp
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	tmp.Alias.Action.traceype = tmp.Type
	jAction, _ := json.Marshal(tmp.Action)
	if err := json.Unmarshal(jAction, &tmp.Alias.Action); err != nil {
		return err
	}

	if tmp.Error == "" {
		tmp.Alias.Result = &EthTraceRes{}
		tmp.Alias.Result.traceype = tmp.Type
		jResult, _ := json.Marshal(tmp.Result)
		if err := json.Unmarshal(jResult, &tmp.Alias.Result); err != nil {
			return err
		}
	}

	*e = EthLocalizedTrace(tmp.Alias)
	return nil
}

type EpochTrace struct {
	CfxTraces        []*LocalizedTrace
	EthTraces        []*EthLocalizedTrace
	MirrorAddressMap map[common.Address]cfxaddress.Address
}

/*
/// Trace
// #[derive(Debug)]
// pub struct LocalizedTrace {
//     /// Action
//     pub action: Action,
//     /// Result
//     pub result: Res,
//     /// Trace address
//     pub trace_address: Vec<U64>,
//     /// Subtraces
//     pub subtraces: U64,
//     /// Transaction position
//     pub transaction_position: Option<U64>,
//     /// Transaction hash
//     pub transaction_hash: Option<H256>,
//     /// Block Number
//     pub block_number: U64,
//     /// Block Hash
//     pub block_hash: H256,
//     /// Valid
//     pub valid: bool,
// }

impl Serialize for LocalizedTrace {
    fn serialize<S>(&self, serializer: S) -> Result<S::Ok, S::Error>
    where S: Serializer {
        let mut struc = serializer.serialize_struct("LocalizedTrace", 9)?;
        match self.action {
            Action::Call(ref call) => {
                struc.serialize_field("type", "call")?;
                struc.serialize_field("action", call)?;
            }
            Action::Create(ref create) => {
                struc.serialize_field("type", "create")?;
                struc.serialize_field("action", create)?;
            }
        }

        match self.result {
            Res::Call(ref call) => struc.serialize_field("result", call)?,
            Res::Create(ref create) => {
                struc.serialize_field("result", create)?
            }
            Res::FailedCall(ref error) => {
                struc.serialize_field("error", &error.to_string())?
            }
            Res::FailedCreate(ref error) => {
                struc.serialize_field("error", &error.to_string())?
            }
            Res::None => {
                struc.serialize_field("result", &None as &Option<u8>)?
            }
        }

        struc.serialize_field("traceAddress", &self.trace_address)?;
        struc.serialize_field("subtraces", &self.subtraces)?;
        struc.serialize_field(
            "transactionPosition",
            &self.transaction_position,
        )?;
        struc.serialize_field("transactionHash", &self.transaction_hash)?;
        struc.serialize_field("blockNumber", &self.block_number)?;
        struc.serialize_field("blockHash", &self.block_hash)?;
        struc.serialize_field("valid", &self.valid)?;

        struc.end()
    }
}
*/
