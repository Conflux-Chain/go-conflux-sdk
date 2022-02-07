package types

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalJSONForAction(t *testing.T) {
	jsonToExpectTypeMap := make(map[string]string)
	call := `{
		"type": "call",
		"action": {
		"from": "cfx:aatd0wzv4f7f6j33kh5e182z4nscsp59vye32h4yz6",
		"to": "cfx:aatd0wzv4f7f6j33kh5e182z4nscsp59vye32h4yz6",
		"value": "0x1234",
		"gas": "0x3456",
		"input": "0x1234",
		"callType": "call"
	}}`

	create := `{
		"type": "create",
		"action": {
		"from": "cfx:aatd0wzv4f7f6j33kh5e182z4nscsp59vye32h4yz6",
		"value": "0x1234",
		"gas": "0x3456",
		"init":"0x1234",
		"createType":"create"
	}}`

	callResult := `{
		"type": "call_result",
		"action": {
		"outcome":"hello",
		"gasLeft":"0x1234",
		"returnData":"0x1234"
	}}`

	createResult := `{
		"type": "create_result",
		"action": {
		"outcome":"hello",
		"addr":"cfx:aatd0wzv4f7f6j33kh5e182z4nscsp59vye32h4yz6",
		"gasLeft":"0x1234",
		"returnData":"0x1234"
	}}`

	InternalTransferAction := `{
		"type": "internal_transfer_action",
		"action": {
		"from": "cfx:aatd0wzv4f7f6j33kh5e182z4nscsp59vye32h4yz6",
		"fromPocket":"balance",
		"to": "cfx:aatd0wzv4f7f6j33kh5e182z4nscsp59vye32h4yz6",
		"toPocket":"balance",
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
		// fmt.Printf("k %v l:%+v \n", k, l)
		if reflect.TypeOf(l.Action).Name() != v {
			t.Fatalf("failed to unmarshal action, expected to type %v,actual:%v", v, reflect.TypeOf(l.Action).Name())
		}
	}
}

func TestTraceInTree(t *testing.T) {

	cases := []struct {
		In        string
		ExpectOut string
	}{
		{
			In:        `[{"action":{"callType":"call","from":"NET1037:TYPE.USER:AAP9KTHVCTUNVF030RBKK9K7ZBZYZ12DAJRT1CB5YM","gas":"0x176d6","input":"0x34d97f38","to":"NET1037:TYPE.CONTRACT:ACCAM3USE2128D6NFX3FN9J10V389WE29AV05W9ZP9","value":"0x0"},"blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7","transactionPosition":"0x0","type":"call"},{"action":{"from":"NET1037:TYPE.CONTRACT:ACCAM3USE2128D6NFX3FN9J10V389WE29AV05W9ZP9","gas":"0xf4e9","init":"0x6080604052348015600f57600080fd5b50608a8061001e6000396000f3fe6080604052348015600f57600080fd5b5060043610603c5760003560e01c80636bdebcc9146041578063c2985578146049578063f1fdece4146049575b600080fd5b6047604f565b005b60476052565b33ff5b56fea2646970667358221220b0eee9118dfaab481f80a888713d69aa457ec808435df671f0ab42c85d5f749564736f6c63430008000033","value":"0x0","createType":"create"},"blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7","transactionPosition":"0x0","type":"create"},{"action":{"addr":"NET1037:TYPE.CONTRACT:ACE08JHGG2H2PNWJ7VMNZZUHKEJ557DGE6820MYE39","gasLeft":"0x88c8","outcome":"success","returnData":"0x6080604052348015600f57600080fd5b5060043610603c5760003560e01c80636bdebcc9146041578063c2985578146049578063f1fdece4146049575b600080fd5b6047604f565b005b60476052565b33ff5b56fea2646970667358221220b0eee9118dfaab481f80a888713d69aa457ec808435df671f0ab42c85d5f749564736f6c63430008000033"},"blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7","transactionPosition":"0x0","type":"create_result"},{"action":{"callType":"call","from":"NET1037:TYPE.CONTRACT:ACCAM3USE2128D6NFX3FN9J10V389WE29AV05W9ZP9","gas":"0x8486","input":"0x6bdebcc9","to":"NET1037:TYPE.CONTRACT:ACE08JHGG2H2PNWJ7VMNZZUHKEJ557DGE6820MYE39","value":"0x0"},"blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7","transactionPosition":"0x0","type":"call"},{"action":{"from":"NET1037:TYPE.CONTRACT:ACE08JHGG2H2PNWJ7VMNZZUHKEJ557DGE6820MYE39","fromPocket":"balance","to":"NET1037:TYPE.CONTRACT:ACCAM3USE2128D6NFX3FN9J10V389WE29AV05W9ZP9","toPocket":"storage_collateral","value":"0x0"},"blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7","transactionPosition":"0x0","type":"internal_transfer_action"},{"action":{"gasLeft":"0x708b","outcome":"success","returnData":"0x"},"blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7","transactionPosition":"0x0","type":"call_result"},{"action":{"gasLeft":"0x727b","outcome":"success","returnData":"0x"},"blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7","transactionPosition":"0x0","type":"call_result"}]`,
			ExpectOut: `[{"type":"call","valid":false,"epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","transactionPosition":"0x0","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7","callWithResult":{"call":{"space":"","from":"net1037:aap9kthvctunvf030rbkk9k7zbzyz12dajrt1cb5ym","to":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","value":"0x0","gas":"0x176d6","input":"0x34d97f38","callType":"call"},"callResult":{"outcome":"success","gasLeft":"0x727b","returnData":"0x"}},"childs":[{"type":"create","valid":false,"epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","transactionPosition":"0x0","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7","createWithResult":{"create":{"space":"","from":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","value":"0x0","gas":"0xf4e9","init":"0x6080604052348015600f57600080fd5b50608a8061001e6000396000f3fe6080604052348015600f57600080fd5b5060043610603c5760003560e01c80636bdebcc9146041578063c2985578146049578063f1fdece4146049575b600080fd5b6047604f565b005b60476052565b33ff5b56fea2646970667358221220b0eee9118dfaab481f80a888713d69aa457ec808435df671f0ab42c85d5f749564736f6c63430008000033","createType":"create"},"createResult":{"outcome":"success","addr":"net1037:ace08jhgg2h2pnwj7vmnzzuhkej557dge6820mye39","gasLeft":"0x88c8","returnData":"0x6080604052348015600f57600080fd5b5060043610603c5760003560e01c80636bdebcc9146041578063c2985578146049578063f1fdece4146049575b600080fd5b6047604f565b005b60476052565b33ff5b56fea2646970667358221220b0eee9118dfaab481f80a888713d69aa457ec808435df671f0ab42c85d5f749564736f6c63430008000033"}},"childs":null},{"type":"call","valid":false,"epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","transactionPosition":"0x0","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7","callWithResult":{"call":{"space":"","from":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","to":"net1037:ace08jhgg2h2pnwj7vmnzzuhkej557dge6820mye39","value":"0x0","gas":"0x8486","input":"0x6bdebcc9","callType":"call"},"callResult":{"outcome":"success","gasLeft":"0x708b","returnData":"0x"}},"childs":[{"type":"internal_transfer_action","valid":false,"epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","transactionPosition":"0x0","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7","internalTransferAction":{"from":"net1037:ace08jhgg2h2pnwj7vmnzzuhkej557dge6820mye39","fromPocket":"balance","fromSpace":"","to":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","toPocket":"storage_collateral","toSpace":"","value":"0x0"},"childs":null}]}]}]`,
		},
		{
			In:        `[{"action":{"from":"NET1037:TYPE.USER:AAKUN8HGEC6H3WVX1KGZ5M5W1P2ZDTZE0UB98NCN90","fromPocket":"balance","to":"NET1037:TYPE.NULL:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA77VMCZA8","toPocket":"gas_payment","value":"0x2f3bd3d0"},"blockHash":"0x2947e8a6aea15a81993f91267d470cc9d18b839365c8f1e6b20329e20b0cc856","epochHash":"0x2947e8a6aea15a81993f91267d470cc9d18b839365c8f1e6b20329e20b0cc856","epochNumber":"0x87","transactionHash":"0x467f7e2505f3b5c7e4a01ddc361c1977706e545d40d2f46b509ad6d8d404c56c","transactionPosition":"0x0","type":"internal_transfer_action","valid":true},{"action":{"callType":"call","from":"NET1037:TYPE.USER:AAKUN8HGEC6H3WVX1KGZ5M5W1P2ZDTZE0UB98NCN90","gas":"0xdaf5","input":"0xa5050b950000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000300000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000005000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000070000000000000000000000000000000000000000000000000000000000000008","to":"NET1037:TYPE.CONTRACT:ACBDRU3EUSC9RGYVXPT3PVNEFFZZ5FG30PA19KX0H8","value":"0x0"},"blockHash":"0x2947e8a6aea15a81993f91267d470cc9d18b839365c8f1e6b20329e20b0cc856","epochHash":"0x2947e8a6aea15a81993f91267d470cc9d18b839365c8f1e6b20329e20b0cc856","epochNumber":"0x87","transactionHash":"0x467f7e2505f3b5c7e4a01ddc361c1977706e545d40d2f46b509ad6d8d404c56c","transactionPosition":"0x0","type":"call","valid":true},{"action":{"gasLeft":"0x1c24","outcome":"success","returnData":"0x"},"blockHash":"0x2947e8a6aea15a81993f91267d470cc9d18b839365c8f1e6b20329e20b0cc856","epochHash":"0x2947e8a6aea15a81993f91267d470cc9d18b839365c8f1e6b20329e20b0cc856","epochNumber":"0x87","transactionHash":"0x467f7e2505f3b5c7e4a01ddc361c1977706e545d40d2f46b509ad6d8d404c56c","transactionPosition":"0x0","type":"call_result","valid":true},{"action":{"from":"NET1037:TYPE.USER:AAKUN8HGEC6H3WVX1KGZ5M5W1P2ZDTZE0UB98NCN90","fromPocket":"balance","to":"NET1037:TYPE.USER:AAKUN8HGEC6H3WVX1KGZ5M5W1P2ZDTZE0UB98NCN90","toPocket":"storage_collateral","value":"0x4563918244f4000"},"blockHash":"0x2947e8a6aea15a81993f91267d470cc9d18b839365c8f1e6b20329e20b0cc856","epochHash":"0x2947e8a6aea15a81993f91267d470cc9d18b839365c8f1e6b20329e20b0cc856","epochNumber":"0x87","transactionHash":"0x467f7e2505f3b5c7e4a01ddc361c1977706e545d40d2f46b509ad6d8d404c56c","transactionPosition":"0x0","type":"internal_transfer_action","valid":true},{"action":{"from":"NET1037:TYPE.CONTRACT:ACBDRU3EUSC9RGYVXPT3PVNEFFZZ5FG30PA19KX0H8","fromPocket":"storage_collateral","to":"NET1037:TYPE.CONTRACT:ACBDRU3EUSC9RGYVXPT3PVNEFFZZ5FG30PA19KX0H8","toPocket":"sponsor_balance_for_collateral","value":"0xde0b6b3a764000"},"blockHash":"0x2947e8a6aea15a81993f91267d470cc9d18b839365c8f1e6b20329e20b0cc856","epochHash":"0x2947e8a6aea15a81993f91267d470cc9d18b839365c8f1e6b20329e20b0cc856","epochNumber":"0x87","transactionHash":"0x467f7e2505f3b5c7e4a01ddc361c1977706e545d40d2f46b509ad6d8d404c56c","transactionPosition":"0x0","type":"internal_transfer_action","valid":true},{"action":{"from":"NET1037:TYPE.NULL:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA77VMCZA8","fromPocket":"gas_payment","to":"NET1037:TYPE.USER:AAKUN8HGEC6H3WVX1KGZ5M5W1P2ZDTZE0UB98NCN90","toPocket":"balance","value":"0x44b3e40"},"blockHash":"0x2947e8a6aea15a81993f91267d470cc9d18b839365c8f1e6b20329e20b0cc856","epochHash":"0x2947e8a6aea15a81993f91267d470cc9d18b839365c8f1e6b20329e20b0cc856","epochNumber":"0x87","transactionHash":"0x467f7e2505f3b5c7e4a01ddc361c1977706e545d40d2f46b509ad6d8d404c56c","transactionPosition":"0x0","type":"internal_transfer_action","valid":true}]`,
			ExpectOut: `[{"type":"internal_transfer_action","valid":true,"epochHash":"0x2947e8a6aea15a81993f91267d470cc9d18b839365c8f1e6b20329e20b0cc856","epochNumber":"0x87","blockHash":"0x2947e8a6aea15a81993f91267d470cc9d18b839365c8f1e6b20329e20b0cc856","transactionPosition":"0x0","transactionHash":"0x467f7e2505f3b5c7e4a01ddc361c1977706e545d40d2f46b509ad6d8d404c56c","internalTransferAction":{"from":"net1037:aakun8hgec6h3wvx1kgz5m5w1p2zdtze0ub98ncn90","fromPocket":"balance","fromSpace":"","to":"net1037:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa77vmcza8","toPocket":"gas_payment","toSpace":"","value":"0x2f3bd3d0"},"childs":null},{"type":"call","valid":true,"epochHash":"0x2947e8a6aea15a81993f91267d470cc9d18b839365c8f1e6b20329e20b0cc856","epochNumber":"0x87","blockHash":"0x2947e8a6aea15a81993f91267d470cc9d18b839365c8f1e6b20329e20b0cc856","transactionPosition":"0x0","transactionHash":"0x467f7e2505f3b5c7e4a01ddc361c1977706e545d40d2f46b509ad6d8d404c56c","callWithResult":{"call":{"space":"","from":"net1037:aakun8hgec6h3wvx1kgz5m5w1p2zdtze0ub98ncn90","to":"net1037:acbdru3eusc9rgyvxpt3pvneffzz5fg30pa19kx0h8","value":"0x0","gas":"0xdaf5","input":"0xa5050b950000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000300000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000005000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000070000000000000000000000000000000000000000000000000000000000000008","callType":"call"},"callResult":{"outcome":"success","gasLeft":"0x1c24","returnData":"0x"}},"childs":null},{"type":"internal_transfer_action","valid":true,"epochHash":"0x2947e8a6aea15a81993f91267d470cc9d18b839365c8f1e6b20329e20b0cc856","epochNumber":"0x87","blockHash":"0x2947e8a6aea15a81993f91267d470cc9d18b839365c8f1e6b20329e20b0cc856","transactionPosition":"0x0","transactionHash":"0x467f7e2505f3b5c7e4a01ddc361c1977706e545d40d2f46b509ad6d8d404c56c","internalTransferAction":{"from":"net1037:aakun8hgec6h3wvx1kgz5m5w1p2zdtze0ub98ncn90","fromPocket":"balance","fromSpace":"","to":"net1037:aakun8hgec6h3wvx1kgz5m5w1p2zdtze0ub98ncn90","toPocket":"storage_collateral","toSpace":"","value":"0x4563918244f4000"},"childs":null},{"type":"internal_transfer_action","valid":true,"epochHash":"0x2947e8a6aea15a81993f91267d470cc9d18b839365c8f1e6b20329e20b0cc856","epochNumber":"0x87","blockHash":"0x2947e8a6aea15a81993f91267d470cc9d18b839365c8f1e6b20329e20b0cc856","transactionPosition":"0x0","transactionHash":"0x467f7e2505f3b5c7e4a01ddc361c1977706e545d40d2f46b509ad6d8d404c56c","internalTransferAction":{"from":"net1037:acbdru3eusc9rgyvxpt3pvneffzz5fg30pa19kx0h8","fromPocket":"storage_collateral","fromSpace":"","to":"net1037:acbdru3eusc9rgyvxpt3pvneffzz5fg30pa19kx0h8","toPocket":"sponsor_balance_for_collateral","toSpace":"","value":"0xde0b6b3a764000"},"childs":null},{"type":"internal_transfer_action","valid":true,"epochHash":"0x2947e8a6aea15a81993f91267d470cc9d18b839365c8f1e6b20329e20b0cc856","epochNumber":"0x87","blockHash":"0x2947e8a6aea15a81993f91267d470cc9d18b839365c8f1e6b20329e20b0cc856","transactionPosition":"0x0","transactionHash":"0x467f7e2505f3b5c7e4a01ddc361c1977706e545d40d2f46b509ad6d8d404c56c","internalTransferAction":{"from":"net1037:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa77vmcza8","fromPocket":"gas_payment","fromSpace":"","to":"net1037:aakun8hgec6h3wvx1kgz5m5w1p2zdtze0ub98ncn90","toPocket":"balance","toSpace":"","value":"0x44b3e40"},"childs":null}]`,
		},
	}

	for _, v := range cases {
		flattenTraces := make([]LocalizedTrace, 0)
		err := json.Unmarshal([]byte(v.In), &flattenTraces)
		if err != nil {
			t.Fatal(err)
		}
		tree, err := TraceInTire(flattenTraces)
		if err != nil {
			t.Fatal(err)
		}

		marshaled, err := json.Marshal(tree)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, v.ExpectOut, string(marshaled))
	}

}

func TestFlattenTraces(t *testing.T) {
	treeInJson := `{"type":"call","valid":false,"epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","transactionPosition":"0x0","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7","callWithResult":{"call":{"from":"net1037:aap9kthvctunvf030rbkk9k7zbzyz12dajrt1cb5ym","to":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","value":"0x0","gas":"0x176d6","input":"0x34d97f38","callType":"call"},"callResult":{"outcome":"success","gasLeft":"0x727b","returnData":"0x"}},"childs":[{"type":"create","valid":false,"epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","transactionPosition":"0x0","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7","createWithResult":{"create":{"createType":"create","from":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","value":"0x0","gas":"0xf4e9","init":"0x6080604052348015600f57600080fd5b50608a8061001e6000396000f3fe6080604052348015600f57600080fd5b5060043610603c5760003560e01c80636bdebcc9146041578063c2985578146049578063f1fdece4146049575b600080fd5b6047604f565b005b60476052565b33ff5b56fea2646970667358221220b0eee9118dfaab481f80a888713d69aa457ec808435df671f0ab42c85d5f749564736f6c63430008000033"},"createResult":{"outcome":"success","addr":"net1037:ace08jhgg2h2pnwj7vmnzzuhkej557dge6820mye39","gasLeft":"0x88c8","returnData":"0x6080604052348015600f57600080fd5b5060043610603c5760003560e01c80636bdebcc9146041578063c2985578146049578063f1fdece4146049575b600080fd5b6047604f565b005b60476052565b33ff5b56fea2646970667358221220b0eee9118dfaab481f80a888713d69aa457ec808435df671f0ab42c85d5f749564736f6c63430008000033"}},"childs":null},{"type":"call","valid":false,"epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","transactionPosition":"0x0","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7","callWithResult":{"call":{"from":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","to":"net1037:ace08jhgg2h2pnwj7vmnzzuhkej557dge6820mye39","value":"0x0","gas":"0x8486","input":"0x6bdebcc9","callType":"call"},"callResult":{"outcome":"success","gasLeft":"0x708b","returnData":"0x"}},"childs":[{"type":"internal_transfer_action","valid":false,"epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","transactionPosition":"0x0","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7","internalTransferAction":{"from":"net1037:ace08jhgg2h2pnwj7vmnzzuhkej557dge6820mye39","fromPocket":"balance","to":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","toPocket":"storage_collateral","value":"0x0"},"childs":null}]}]}`
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
	expectJson := `[{"type":"call","valid":false,"epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","transactionPosition":"0x0","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7","callWithResult":{"call":{"space":"","from":"net1037:aap9kthvctunvf030rbkk9k7zbzyz12dajrt1cb5ym","to":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","value":"0x0","gas":"0x176d6","input":"0x34d97f38","callType":"call"},"callResult":{"outcome":"success","gasLeft":"0x727b","returnData":"0x"}},"childs":null},{"type":"create","valid":false,"epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","transactionPosition":"0x0","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7","createWithResult":{"create":{"space":"","from":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","value":"0x0","gas":"0xf4e9","init":"0x6080604052348015600f57600080fd5b50608a8061001e6000396000f3fe6080604052348015600f57600080fd5b5060043610603c5760003560e01c80636bdebcc9146041578063c2985578146049578063f1fdece4146049575b600080fd5b6047604f565b005b60476052565b33ff5b56fea2646970667358221220b0eee9118dfaab481f80a888713d69aa457ec808435df671f0ab42c85d5f749564736f6c63430008000033","createType":"create"},"createResult":{"outcome":"success","addr":"net1037:ace08jhgg2h2pnwj7vmnzzuhkej557dge6820mye39","gasLeft":"0x88c8","returnData":"0x6080604052348015600f57600080fd5b5060043610603c5760003560e01c80636bdebcc9146041578063c2985578146049578063f1fdece4146049575b600080fd5b6047604f565b005b60476052565b33ff5b56fea2646970667358221220b0eee9118dfaab481f80a888713d69aa457ec808435df671f0ab42c85d5f749564736f6c63430008000033"}},"childs":null},{"type":"call","valid":false,"epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","transactionPosition":"0x0","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7","callWithResult":{"call":{"space":"","from":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","to":"net1037:ace08jhgg2h2pnwj7vmnzzuhkej557dge6820mye39","value":"0x0","gas":"0x8486","input":"0x6bdebcc9","callType":"call"},"callResult":{"outcome":"success","gasLeft":"0x708b","returnData":"0x"}},"childs":null},{"type":"internal_transfer_action","valid":false,"epochHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","epochNumber":"0x361f","blockHash":"0x4d12913ac843622159b26b830de53194e090de93d2e82ec19e703c5b142258e3","transactionPosition":"0x0","transactionHash":"0x3d63a2a74c0274f3af7fc2ae9215b5d07ae98c4eebdf12a8632b54df028c9fa7","internalTransferAction":{"from":"net1037:ace08jhgg2h2pnwj7vmnzzuhkej557dge6820mye39","fromPocket":"balance","fromSpace":"","to":"net1037:accam3use2128d6nfx3fn9j10v389we29av05w9zp9","toPocket":"storage_collateral","toSpace":"","value":"0x0"},"childs":null}]`
	assert.Equal(t, expectJson, string(marshaled))
	// fmt.Printf("flattened %v", string(marshaled))
}
