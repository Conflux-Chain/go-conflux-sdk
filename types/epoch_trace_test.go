package types

import (
	"encoding/json"
	"testing"

	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/stretchr/testify/assert"
)

func TestEthLocalizedTraceJson(t *testing.T) {
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
