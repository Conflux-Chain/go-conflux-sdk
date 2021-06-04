package types

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/ethereum/go-ethereum/rlp"
)

var (
	logFilterStrTemplate = `{
		"fromEpoch": "0x1",
		"toEpoch": "0x2",
		"blockHashes": [
			"0x1",
			"0x2"
		],
		"address": [],
		"topics": [],
		"limit": "0x1"
	}
	`

	expectTemplate = LogFilter{
		FromEpoch:   NewEpochNumberUint64(0x1),
		ToEpoch:     NewEpochNumberUint64(0x2),
		BlockHashes: []Hash{"0x1", "0x2"},
		Address:     []Address{},
		Topics: [][]Hash{
			{"0xA"},
			nil,
			{"0xB", "0xC"},
			nil,
		},
		Limit: NewUint64(0x1),
	}
)

func TestRLPMarshalLog(t *testing.T) {
	logJson := `{"address":"cfx:acg158kvr8zanb1bs048ryb6rtrhr283ma70vz70tx","topics":["0x2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d","0x00000000000000000000000080ae6a88ce3351e9f729e8199f2871ba786ad7c5","0x0000000000000000000000008d545118d91c027c805c552f63a5c00a20ae6aca"],"data":"0x00000000000000000000000000000000000000000000003b16c9e8eeb7c800000000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","blockHash":null,"epochNumber":null,"transactionHash":null,"transactionIndex":null,"logIndex":null,"transactionLogIndex":null}`

	var log Log
	err := json.Unmarshal([]byte(logJson), &log)
	fatalIfErr(t, err)
	// RLP marshal log to bytes
	dBytes, err := rlp.EncodeToBytes(log)
	fatalIfErr(t, err)
	// RLP unmarshal bytes back to log
	var log2 Log
	err = rlp.DecodeBytes(dBytes, &log2)
	fatalIfErr(t, err)
	// Json marshal log
	jBytes1, err := json.Marshal(log)
	fatalIfErr(t, err)
	logJsonStr := string(jBytes1)
	// Json marshal log2
	jBytes2, err := json.Marshal(log2)
	fatalIfErr(t, err)
	logJsonStr2 := string(jBytes2)

	if logJsonStr != logJsonStr2 {
		t.Fatalf("expect %#v, actual %#v", logJsonStr, logJsonStr2)
	}
}

func TestUnmarshalJSONToLogFilter(t *testing.T) {
	verifyLogFilter(t, `"address": null`,
		`"topics": null`,
		nil,
		nil,
	)

	verifyLogFilter(t, `"address": ["cfx:acc7uawf5ubtnmezvhu9dhc6sghea0403y2dgpyfjp","cfx:aaejuaaaaaaaaaaaaaaaaaaaaaaaaaaaajrwuc9jnb"]`,
		`"topics": ["0xA",null,["0xB","0xC"],null]`,
		[]Address{cfxaddress.MustNewFromBase32("cfx:acc7uawf5ubtnmezvhu9dhc6sghea0403y2dgpyfjp"),
			cfxaddress.MustNewFromBase32("cfx:aaejuaaaaaaaaaaaaaaaaaaaaaaaaaaaajrwuc9jnb")},
		[][]Hash{
			{"0xA"},
			nil,
			{"0xB", "0xC"},
			nil,
		},
	)

	verifyLogFilter(t, `"address": "cfx:acc7uawf5ubtnmezvhu9dhc6sghea0403y2dgpyfjp"`,
		`"topics": ["0xA",null,["0xB","0xC"],null]`,
		[]Address{cfxaddress.MustNewFromBase32("cfx:acc7uawf5ubtnmezvhu9dhc6sghea0403y2dgpyfjp")},
		[][]Hash{
			{"0xA"},
			nil,
			{"0xB", "0xC"},
			nil,
		},
	)

	verifyLogFilter(t, `"address": "cfx:acc7uawf5ubtnmezvhu9dhc6sghea0403y2dgpyfjp"`,
		`"topics": ["0xA",null,"0xB",null]`,
		[]Address{cfxaddress.MustNewFromBase32("cfx:acc7uawf5ubtnmezvhu9dhc6sghea0403y2dgpyfjp")},
		[][]Hash{
			{"0xA"},
			nil,
			{"0xB"},
			nil,
		},
	)

	verifyLogFilter(t, `"address": "cfx:acc7uawf5ubtnmezvhu9dhc6sghea0403y2dgpyfjp"`,
		`"topics": ["0xA"]`,
		[]Address{cfxaddress.MustNewFromBase32("cfx:acc7uawf5ubtnmezvhu9dhc6sghea0403y2dgpyfjp")},
		[][]Hash{
			{"0xA"},
		},
	)
}

func verifyLogFilter(t *testing.T, address string, topics string, expectAddress []Address, expectTopics [][]Hash) {
	input := strings.Replace(logFilterStrTemplate, `"address": []`, address, -1)
	input = strings.Replace(input, `"topics": []`, topics, -1)

	expect := expectTemplate
	expect.Address = expectAddress
	expect.Topics = expectTopics

	actual := LogFilter{}
	err := json.Unmarshal([]byte(input), &actual)
	if err != nil {
		t.Error(err)
		t.Fatalf("failed to unmarshal %v", input)
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Fatalf("expect %v, actual %v", utils.PrettyJSON(expect), utils.PrettyJSON(actual))
	}
}
