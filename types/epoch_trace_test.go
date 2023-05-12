package types

import (
	"encoding/json"
	"testing"

	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/stretchr/testify/assert"
)

func TestEthLocalizedTraceJson(t *testing.T) {
	// hash := Hash("0x1234567890abcdef")
	// e := EthLocalizedTrace{
	// 	Type: TRACE_CALL,
	// 	Action: EthTraceAction{
	// 		Call: &EthTraceCall{
	// 			From:     common.HexToAddress("cfx:aamgvyzht7h1zxdghb9ee9w26wrz8rd3gj837392dp"),
	// 			To:       common.HexToAddress("cfx:aanht94umnv60ek18w3eskj1np693hp6c68yxze4df"),
	// 			Value:    *NewBigInt(1),
	// 			Gas:      *NewBigInt(1),
	// 			Input:    common.Hex2Bytes("1234"),
	// 			CallType: CALL_CALL,
	// 		},
	// 	},
	// 	// Result: EthTraceRes{
	// 	// 	Call: &EthTraceCallResult{
	// 	// 		GasUsed: *NewBigInt(1),
	// 	// 		Output:  common.Hex2Bytes("1234"),
	// 	// 	},
	// 	// },
	// 	Error:               "error message",
	// 	TraceAddress:        []hexutil.Uint64{*NewUint64(1), *NewUint64(2)},
	// 	Subtraces:           *NewUint64(1),
	// 	TransactionPosition: NewUint64(1),
	// 	TransactionHash:     &hash,
	// 	BlockNumber:         *NewUint64(1),
	// 	BlockHash:           common.HexToHash("0x1d314a81d2ace0503494e52c8acee8ebb7a401cf02560451b1a216334c988ca8"),
	// 	Valid:               true,
	// }

	// j, err := json.Marshal(e)
	// assert.NoError(t, err)
	// fmt.Printf("json\n%s", j)

	input := []string{
		`{"type":"call","action":{"from":"0x00000000000000000000000000000000000000cf","to":"0x00000000000000000000000000000000000000cf","value":"0x1","gas":"0x1","input":"0x1234","callType":"call"},"result":{"gasUsed":"0x1","output":"0x1234"},"traceAddress":["0x1","0x2"],"subtraces":"0x1","transactionPosition":"0x1","transactionHash":"0x81c0a9d47a69de4898f3f63234e8503c630021c7c308265d2f845b34d0e3c5bd","blockNumber":"0x1","blockHash":"0x1d314a81d2ace0503494e52c8acee8ebb7a401cf02560451b1a216334c988ca8","valid":true}`,
		`{"type":"call","action":{"from":"0x00000000000000000000000000000000000000cf","to":"0x00000000000000000000000000000000000000cf","value":"0x1","gas":"0x1","input":"0x1234","callType":"call"},"error":"error message","traceAddress":["0x1","0x2"],"subtraces":"0x1","transactionPosition":"0x1","transactionHash":"0x81c0a9d47a69de4898f3f63234e8503c630021c7c308265d2f845b34d0e3c5bd","blockNumber":"0x1","blockHash":"0x1d314a81d2ace0503494e52c8acee8ebb7a401cf02560451b1a216334c988ca8","valid":true}`,
	}

	for _, v := range input {
		var e EthLocalizedTrace
		err := json.Unmarshal([]byte(v), &e)
		assert.NoError(t, err, v)

		output, err := json.Marshal(e)
		assert.NoError(t, err, v)

		assert.Equal(t, utils.FormatJson(v), utils.FormatJson(string(output)))
	}
}
