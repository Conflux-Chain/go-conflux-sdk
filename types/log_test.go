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
			"0x0b50338ca31401a6ccb81a799b4c774ce86af93c1b740c52d12958e35461d999",
			"0x0b50338ca31401a6ccb81a799b4c774ce86af93c1b740c52d12958e35461d99a"
		],
		"address": [],
		"topics": [],
		"limit": "0x1"
	}
	`

	expectTemplate = LogFilter{
		FromEpoch:   NewEpochNumberUint64(0x1),
		ToEpoch:     NewEpochNumberUint64(0x2),
		BlockHashes: []Hash{"0x0b50338ca31401a6ccb81a799b4c774ce86af93c1b740c52d12958e35461d999", "0x0b50338ca31401a6ccb81a799b4c774ce86af93c1b740c52d12958e35461d99a"},
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
	logJson1 := `{"address":"cfx:acg158kvr8zanb1bs048ryb6rtrhr283ma70vz70tx","topics":["0x2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d","0x00000000000000000000000080ae6a88ce3351e9f729e8199f2871ba786ad7c5","0x0000000000000000000000008d545118d91c027c805c552f63a5c00a20ae6aca"],"data":"0x00000000000000000000000000000000000000000000003b16c9e8eeb7c800000000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","blockHash":null,"epochNumber":null,"transactionHash":null,"transactionIndex":null,"logIndex":null,"transactionLogIndex":null}`
	logJson2 := `{"address":"cfx:acc8ya1f2a2bfphxg5ax7a8h29k47d5xsebxfj24nd","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x0000000000000000000000008225b98f888c632241ec3c60b1a528053ef37d4f","0x0000000000000000000000001a7fabd17788269d52ef0850f2e0dcbf444a9403"],"data":"0x00000000000000000000000000000000000000000000000016b276f7e4180350","blockHash":"0x25fd857127dbef9e9a4f851ade9b11ac1cb22967266d915fe00f61ec1e356f54","epochNumber":"0xfce4df","transactionHash":"0xa45f5101687f1cc11ea8ea56eab670cff670c55c84c7feb06ed33b7840346c60","transactionIndex":"0x1","logIndex":"0x0","transactionLogIndex":"0x0"}`

	for _, logJson := range []string{logJson2, logJson1} {
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
}

func TestRLPMarshalSubscriptionLog(t *testing.T) {
	logJson1 := `{"address":"cfx:acg158kvr8zanb1bs048ryb6rtrhr283ma70vz70tx","topics":["0x2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d","0x00000000000000000000000080ae6a88ce3351e9f729e8199f2871ba786ad7c5","0x0000000000000000000000008d545118d91c027c805c552f63a5c00a20ae6aca"],"data":"0x00000000000000000000000000000000000000000000003b16c9e8eeb7c800000000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","blockHash":null,"epochNumber":null,"transactionHash":null,"transactionIndex":null,"logIndex":null,"transactionLogIndex":null}`
	logJson2 := `{"address":"cfx:acc8ya1f2a2bfphxg5ax7a8h29k47d5xsebxfj24nd","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x0000000000000000000000008225b98f888c632241ec3c60b1a528053ef37d4f","0x0000000000000000000000001a7fabd17788269d52ef0850f2e0dcbf444a9403"],"data":"0x00000000000000000000000000000000000000000000000016b276f7e4180350","blockHash":"0x25fd857127dbef9e9a4f851ade9b11ac1cb22967266d915fe00f61ec1e356f54","epochNumber":"0xfce4df","transactionHash":"0xa45f5101687f1cc11ea8ea56eab670cff670c55c84c7feb06ed33b7840346c60","transactionIndex":"0x1","logIndex":"0x0","transactionLogIndex":"0x0"}`

	for _, logJson := range []string{logJson2, logJson1} {
		var log SubscriptionLog
		err := json.Unmarshal([]byte(logJson), &log)
		fatalIfErr(t, err)

		// fmt.Printf("log:%v\n", utils.PrettyJSON(log))
		// RLP marshal log to bytes
		dBytes, err := rlp.EncodeToBytes(log)
		fatalIfErr(t, err)

		// fmt.Printf("log dBytes:%+x\n", dBytes)
		// RLP unmarshal bytes back to log
		var log2 SubscriptionLog
		err = rlp.DecodeBytes(dBytes, &log2)
		fatalIfErr(t, err)
		// fmt.Printf("dlog :%+v\n", log2)
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
}

func TestUnmarshalJSONToLogFilter(t *testing.T) {
	verifyLogFilter(t,
		`"address": null`,
		`"topics": null`,
		nil,
		nil,
	)

	verifyLogFilter(t,
		`"address": ["cfx:acc7uawf5ubtnmezvhu9dhc6sghea0403y2dgpyfjp","cfx:aaejuaaaaaaaaaaaaaaaaaaaaaaaaaaaajrwuc9jnb"]`,
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

	verifyLogFilter(t,
		`"address": "cfx:acc7uawf5ubtnmezvhu9dhc6sghea0403y2dgpyfjp"`,
		`"topics": ["0xA",null,["0xB","0xC"],null]`,
		[]Address{cfxaddress.MustNewFromBase32("cfx:acc7uawf5ubtnmezvhu9dhc6sghea0403y2dgpyfjp")},
		[][]Hash{
			{"0xA"},
			nil,
			{"0xB", "0xC"},
			nil,
		},
	)

	verifyLogFilter(t,
		`"address": "cfx:acc7uawf5ubtnmezvhu9dhc6sghea0403y2dgpyfjp"`,
		`"topics": ["0xA",null,"0xB",null]`,
		[]Address{cfxaddress.MustNewFromBase32("cfx:acc7uawf5ubtnmezvhu9dhc6sghea0403y2dgpyfjp")},
		[][]Hash{
			{"0xA"},
			nil,
			{"0xB"},
			nil,
		},
	)

	verifyLogFilter(t,
		`"address": "cfx:acc7uawf5ubtnmezvhu9dhc6sghea0403y2dgpyfjp"`,
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

func TestUnMarhsalSubscribeLog(t *testing.T) {
	slogJson := `{"address":"cfxtest:acfpjyyu1wp1h79h589dcxhms3z40zgj1y41cj8u3k","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x00000000000000000000000019f4bcf113e0b896d9b34294fd3da86b4adf0302","0x000000000000000000000000160ebef20c1f739957bf9eecd040bce699cc42c6"],"data":"0x000000000000000000000000000000000000000000000000000000000000000a","blockHash":"0x0c4ee0e8487b7556e9fbc07c8b2f8b5d8ea3841cb104ef6ca0ccc373b44446e6","epochNumber":"0x233e0d4","transactionHash":"0x0631cd13ff875fd991c340d4dac5c85fbc36302af866c2a7c69371064250260f","transactionIndex":"0x0","logIndex":"0x0","transactionLogIndex":"0x0"}`

	actual := SubscriptionLog{}
	err := json.Unmarshal([]byte(slogJson), &actual)
	fatalIfErr(t, err)

	l := Log{}
	err = json.Unmarshal([]byte(slogJson), &l)
	fatalIfErr(t, err)

	expect := SubscriptionLog{&l, nil}
	if !reflect.DeepEqual(expect, actual) {
		t.Fatalf("expect %+v, actual %+v", expect, actual)
	}

	slogJson = `{"revertTo":"0x40f"}`
	actual = SubscriptionLog{}
	err = json.Unmarshal([]byte(slogJson), &actual)
	fatalIfErr(t, err)

	r := ChainReorg{}
	err = json.Unmarshal([]byte(slogJson), &r)
	fatalIfErr(t, err)

	expect = SubscriptionLog{nil, &r}
	if !reflect.DeepEqual(expect, actual) {
		t.Fatalf("expect %+v, actual %+v", expect, actual)
	}
}

func TestMarshalSubscriptionLog(t *testing.T) {
	logJson := `{"address":"cfxtest:acfpjyyu1wp1h79h589dcxhms3z40zgj1y41cj8u3k","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x00000000000000000000000019f4bcf113e0b896d9b34294fd3da86b4adf0302","0x000000000000000000000000160ebef20c1f739957bf9eecd040bce699cc42c6"],"data":"0x000000000000000000000000000000000000000000000000000000000000000a","blockHash":"0x0c4ee0e8487b7556e9fbc07c8b2f8b5d8ea3841cb104ef6ca0ccc373b44446e6","epochNumber":"0x233e0d4","transactionHash":"0x0631cd13ff875fd991c340d4dac5c85fbc36302af866c2a7c69371064250260f","transactionIndex":"0x0","logIndex":"0x0","transactionLogIndex":"0x0"}`

	l := Log{}
	err := json.Unmarshal([]byte(logJson), &l)
	fatalIfErr(t, err)

	sLog := SubscriptionLog{&l, nil}
	actual, err := json.Marshal(sLog)
	fatalIfErr(t, err)

	expect, err := json.Marshal(l)
	fatalIfErr(t, err)

	if !reflect.DeepEqual(actual, expect) {
		t.Fatalf("expect %v, actual %v", string(expect), string(actual))
	}

	reorgJson := `{"revertTo":"0x40f"}`
	r := ChainReorg{}
	err = json.Unmarshal([]byte(reorgJson), &r)
	fatalIfErr(t, err)

	sLog = SubscriptionLog{nil, &r}
	actual, err = json.Marshal(sLog)
	fatalIfErr(t, err)

	expect, err = json.Marshal(r)
	fatalIfErr(t, err)

	if !reflect.DeepEqual(actual, expect) {
		t.Fatalf("expect %v, actual %v", string(expect), string(actual))
	}

}
