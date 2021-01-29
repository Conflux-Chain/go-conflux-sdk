package types

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/Conflux-Chain/go-conflux-sdk/utils"
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
