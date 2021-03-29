package types

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestUnmarshalJSONForAction(t *testing.T) {
	jsonToExpectTypeMap := make(map[string]string)
	call := `{
		"action": {
		"from": "cfx:aatd0wzv4f7f6j33kh5e182z4nscsp59vye32h4yz6",
		"to": "cfx:aatd0wzv4f7f6j33kh5e182z4nscsp59vye32h4yz6",
		"value": "0x1234",
		"gas": "0x3456",
		"input": "0x1234",
		"callType": "call"
	}}`

	create := `{
		"action": {
		"from": "cfx:aatd0wzv4f7f6j33kh5e182z4nscsp59vye32h4yz6",
		"value": "0x1234",
		"gas": "0x3456",
		"init":"0x1234"
	}}`

	callResult := `{
		"action": {
		"outcome":"hello",
		"gasLeft":"0x1234",
		"returnData":"0x1234"
	}}`

	createResult := `{
		"action": {
		"outcome":"hello",
		"addr":"cfx:aatd0wzv4f7f6j33kh5e182z4nscsp59vye32h4yz6",
		"gasLeft":"0x1234",
		"returnData":"0x1234"
	}}`

	InternalTransferAction := `{
		"action": {
		"from": "cfx:aatd0wzv4f7f6j33kh5e182z4nscsp59vye32h4yz6",
		"to": "cfx:aatd0wzv4f7f6j33kh5e182z4nscsp59vye32h4yz6",
		"value": "0x1234"
	}}`

	jsonToExpectTypeMap[call] = "Call"
	jsonToExpectTypeMap[create] = "Create"
	jsonToExpectTypeMap[callResult] = "CallResult"
	jsonToExpectTypeMap[createResult] = "CreateResult"
	jsonToExpectTypeMap[InternalTransferAction] = "InternalTransferAction"

	for k, v := range jsonToExpectTypeMap {
		l := LocalizedTrace{}
		err := json.Unmarshal([]byte(k), &l)
		if err != nil {
			t.Fatalf("failed to unmarshl %v,err:%v", k, err.Error())
		}
		fmt.Printf("k %v l:%#v \n", k, l)
		if reflect.TypeOf(l.Action).Name() != v {
			t.Fatalf("failed to unmarshal action, expected to type %v,actual:%v", v, reflect.TypeOf(l.Action).Name())
		}
	}

}
