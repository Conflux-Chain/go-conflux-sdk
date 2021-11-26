package types

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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
		// fmt.Printf("k %v l:%#v \n", k, l)
		if reflect.TypeOf(l.Action).Name() != v {
			t.Fatalf("failed to unmarshal action, expected to type %v,actual:%v", v, reflect.TypeOf(l.Action).Name())
		}
	}
}

func TestTraceInTree(t *testing.T) {
	tracesJson := `[{"action":{"callType":"call","from":"NET1037:TYPE.USER:AAP9KTHVCTUNVF030RBKK9K7ZBZYZ12DAJRT1CB5YM","gas":"0x176d6","input":"0x34d97f38","to":"NET1037:TYPE.CONTRACT:ACCAM3USE2128D6NFX3FN9J10V389WE29AV05W9ZP9","value":"0x0"},"blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7","transactionPosition":"0x0","type":"call"},{"action":{"from":"NET1037:TYPE.CONTRACT:ACCAM3USE2128D6NFX3FN9J10V389WE29AV05W9ZP9","gas":"0xf4e9","init":"0x6080604052348015600f57600080fd5b50608a8061001e6000396000f3fe6080604052348015600f57600080fd5b5060043610603c5760003560e01c80636bdebcc9146041578063c2985578146049578063f1fdece4146049575b600080fd5b6047604f565b005b60476052565b33ff5b56fea2646970667358221220b0eee9118dfaab481f80a888713d69aa457ec808435df671f0ab42c85d5f749564736f6c63430008000033","value":"0x0"},"blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7","transactionPosition":"0x0","type":"create"},{"action":{"addr":"NET1037:TYPE.CONTRACT:ACE08JHGG2H2PNWJ7VMNZZUHKEJ557DGE6820MYE39","gasLeft":"0x88c8","outcome":"success","returnData":"0x6080604052348015600f57600080fd5b5060043610603c5760003560e01c80636bdebcc9146041578063c2985578146049578063f1fdece4146049575b600080fd5b6047604f565b005b60476052565b33ff5b56fea2646970667358221220b0eee9118dfaab481f80a888713d69aa457ec808435df671f0ab42c85d5f749564736f6c63430008000033"},"blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7","transactionPosition":"0x0","type":"create_result"},{"action":{"callType":"call","from":"NET1037:TYPE.CONTRACT:ACCAM3USE2128D6NFX3FN9J10V389WE29AV05W9ZP9","gas":"0x8486","input":"0x6bdebcc9","to":"NET1037:TYPE.CONTRACT:ACE08JHGG2H2PNWJ7VMNZZUHKEJ557DGE6820MYE39","value":"0x0"},"blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7","transactionPosition":"0x0","type":"call"},{"action":{"from":"NET1037:TYPE.CONTRACT:ACE08JHGG2H2PNWJ7VMNZZUHKEJ557DGE6820MYE39","to":"NET1037:TYPE.CONTRACT:ACCAM3USE2128D6NFX3FN9J10V389WE29AV05W9ZP9","value":"0x0"},"blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7","transactionPosition":"0x0","type":"internal_transfer_action"},{"action":{"gasLeft":"0x708b","outcome":"success","returnData":"0x"},"blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7","transactionPosition":"0x0","type":"call_result"},{"action":{"gasLeft":"0x727b","outcome":"success","returnData":"0x"},"blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7","transactionPosition":"0x0","type":"call_result"}]`
	flattenTraces := make([]LocalizedTrace, 0)
	err := json.Unmarshal([]byte(tracesJson), &flattenTraces)
	if err != nil {
		t.Fatal(err)
	}
	tree, err := TraceInTree(flattenTraces)
	if err != nil {
		t.Fatal(err)
	}

	marshaled, err := json.Marshal(tree)
	if err != nil {
		t.Fatal(err)
	}
	expectJson := `{"childs":[{"childs":null,"raw":{"action":{"from":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","value":"0x0","gas":"0xf4e9","init":"0x6080604052348015600f57600080fd5b50608a8061001e6000396000f3fe6080604052348015600f57600080fd5b5060043610603c5760003560e01c80636bdebcc9146041578063c2985578146049578063f1fdece4146049575b600080fd5b6047604f565b005b60476052565b33ff5b56fea2646970667358221220b0eee9118dfaab481f80a888713d69aa457ec808435df671f0ab42c85d5f749564736f6c63430008000033"},"type":"create","epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","transactionPosition":"0x0","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7"},"createWithResult":{"Create":{"from":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","value":"0x0","gas":"0xf4e9","init":"0x6080604052348015600f57600080fd5b50608a8061001e6000396000f3fe6080604052348015600f57600080fd5b5060043610603c5760003560e01c80636bdebcc9146041578063c2985578146049578063f1fdece4146049575b600080fd5b6047604f565b005b60476052565b33ff5b56fea2646970667358221220b0eee9118dfaab481f80a888713d69aa457ec808435df671f0ab42c85d5f749564736f6c63430008000033"},"CreateResult":{"outcome":"success","addr":"net1037:ace08jhgg2h2pnwj7vmnzzuhkej557dge6820mye39","gasLeft":"0x88c8","returnData":"0x6080604052348015600f57600080fd5b5060043610603c5760003560e01c80636bdebcc9146041578063c2985578146049578063f1fdece4146049575b600080fd5b6047604f565b005b60476052565b33ff5b56fea2646970667358221220b0eee9118dfaab481f80a888713d69aa457ec808435df671f0ab42c85d5f749564736f6c63430008000033"}}},{"childs":[{"childs":null,"raw":{"action":{"from":"net1037:ace08jhgg2h2pnwj7vmnzzuhkej557dge6820mye39","to":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","value":"0x0"},"type":"internal_transfer_action","epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","transactionPosition":"0x0","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7"},"internalTransferAction":{"from":"net1037:ace08jhgg2h2pnwj7vmnzzuhkej557dge6820mye39","to":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","value":"0x0"}}],"raw":{"action":{"from":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","to":"net1037:ace08jhgg2h2pnwj7vmnzzuhkej557dge6820mye39","value":"0x0","gas":"0x8486","input":"0x6bdebcc9","callType":"call"},"type":"call","epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","transactionPosition":"0x0","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7"},"callWithResult":{"Call":{"from":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","to":"net1037:ace08jhgg2h2pnwj7vmnzzuhkej557dge6820mye39","value":"0x0","gas":"0x8486","input":"0x6bdebcc9","callType":"call"},"CallResult":{"outcome":"success","gasLeft":"0x708b","returnData":"0x"}}}],"raw":{"action":{"from":"net1037:aap9kthvctunvf030rbkk9k7zbzyz12dajrt1cb5ym","to":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","value":"0x0","gas":"0x176d6","input":"0x34d97f38","callType":"call"},"type":"call","epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","transactionPosition":"0x0","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7"},"callWithResult":{"Call":{"from":"net1037:aap9kthvctunvf030rbkk9k7zbzyz12dajrt1cb5ym","to":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","value":"0x0","gas":"0x176d6","input":"0x34d97f38","callType":"call"},"CallResult":{"outcome":"success","gasLeft":"0x727b","returnData":"0x"}}}`
	assert.Equal(t, expectJson, string(marshaled))
	fmt.Printf("tree %v", string(marshaled))
}

func TestFlattenTraces(t *testing.T) {
	treeInJson := `{"childs":[{"childs":null,"raw":{"action":{"from":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","value":"0x0","gas":"0xf4e9","init":"0x6080604052348015600f57600080fd5b50608a8061001e6000396000f3fe6080604052348015600f57600080fd5b5060043610603c5760003560e01c80636bdebcc9146041578063c2985578146049578063f1fdece4146049575b600080fd5b6047604f565b005b60476052565b33ff5b56fea2646970667358221220b0eee9118dfaab481f80a888713d69aa457ec808435df671f0ab42c85d5f749564736f6c63430008000033"},"type":"create","epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","transactionPosition":"0x0","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7"},"createWithResult":{"Create":{"from":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","value":"0x0","gas":"0xf4e9","init":"0x6080604052348015600f57600080fd5b50608a8061001e6000396000f3fe6080604052348015600f57600080fd5b5060043610603c5760003560e01c80636bdebcc9146041578063c2985578146049578063f1fdece4146049575b600080fd5b6047604f565b005b60476052565b33ff5b56fea2646970667358221220b0eee9118dfaab481f80a888713d69aa457ec808435df671f0ab42c85d5f749564736f6c63430008000033"},"CreateResult":{"outcome":"success","addr":"net1037:ace08jhgg2h2pnwj7vmnzzuhkej557dge6820mye39","gasLeft":"0x88c8","returnData":"0x6080604052348015600f57600080fd5b5060043610603c5760003560e01c80636bdebcc9146041578063c2985578146049578063f1fdece4146049575b600080fd5b6047604f565b005b60476052565b33ff5b56fea2646970667358221220b0eee9118dfaab481f80a888713d69aa457ec808435df671f0ab42c85d5f749564736f6c63430008000033"}}},{"childs":[{"childs":null,"raw":{"action":{"from":"net1037:ace08jhgg2h2pnwj7vmnzzuhkej557dge6820mye39","to":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","value":"0x0"},"type":"internal_transfer_action","epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","transactionPosition":"0x0","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7"},"internalTransferAction":{"from":"net1037:ace08jhgg2h2pnwj7vmnzzuhkej557dge6820mye39","to":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","value":"0x0"}}],"raw":{"action":{"from":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","to":"net1037:ace08jhgg2h2pnwj7vmnzzuhkej557dge6820mye39","value":"0x0","gas":"0x8486","input":"0x6bdebcc9","callType":"call"},"type":"call","epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","transactionPosition":"0x0","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7"},"callWithResult":{"Call":{"from":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","to":"net1037:ace08jhgg2h2pnwj7vmnzzuhkej557dge6820mye39","value":"0x0","gas":"0x8486","input":"0x6bdebcc9","callType":"call"},"CallResult":{"outcome":"success","gasLeft":"0x708b","returnData":"0x"}}}],"raw":{"action":{"from":"net1037:aap9kthvctunvf030rbkk9k7zbzyz12dajrt1cb5ym","to":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","value":"0x0","gas":"0x176d6","input":"0x34d97f38","callType":"call"},"type":"call","epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","transactionPosition":"0x0","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7"},"callWithResult":{"Call":{"from":"net1037:aap9kthvctunvf030rbkk9k7zbzyz12dajrt1cb5ym","to":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","value":"0x0","gas":"0x176d6","input":"0x34d97f38","callType":"call"},"CallResult":{"outcome":"success","gasLeft":"0x727b","returnData":"0x"}}}`
	tree := LocalizedTraceNode{}
	err := json.Unmarshal([]byte(treeInJson), &tree)
	if err != nil {
		t.Fatal(err)
	}
	flattened := tree.Flatten()
	marshaled, err := json.Marshal(flattened)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("flattened %v", string(marshaled))
}
