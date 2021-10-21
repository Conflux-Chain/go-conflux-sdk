package postypes

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func TestBlockNumberUnmarshalJSON(t *testing.T) {
	table := []struct {
		input       string
		expect      BlockNumber
		expectError bool
	}{
		{
			input:  `"0x12"`,
			expect: BlockNumber{number: hexutil.Uint64(0x12)},
		},
		{
			input:  `"latest_committed"`,
			expect: BlockNumber{name: "latest_committed"},
		},
		{
			input:       `"0x12232323232323232323232"`,
			expectError: true,
		},
		{
			input:       `"LatestCommitted"`,
			expectError: true,
		},
	}

	for _, v := range table {
		var actual BlockNumber
		err := json.Unmarshal([]byte(v.input), &actual)

		// fmt.Printf("actual %+v, err %v", actual, err)

		if v.expectError {
			if err == nil {
				t.Fatalf("expect error, actual %+v", actual)
			}
			continue
		}

		if err != nil {
			t.Fatalf("unexpect error:%v", err)
		}

		if !reflect.DeepEqual(actual, v.expect) {
			t.Fatalf("expect %+v, actual %+v", v.expect, actual)
		}

	}
}
