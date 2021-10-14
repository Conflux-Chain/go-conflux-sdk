package postypes

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestUnmarshalTransaction(t *testing.T) {

	type tableTemplate struct {
		data   string
		expect *Transaction
		actual *Transaction
	}

	table := []tableTemplate{
		{
			data:   genTxJsonWithPayload(`"payload": null`, "BlockMetadata"),
			expect: genTxWithPayload(nil, "BlockMetadata"),
			actual: nil,
		},

		{
			data: genTxJsonWithPayload(`"payload": {
				"blockHash": "0xe18e8f0ff9dbb12d0d27451fea177f277445e244e2524e9b5ea370282fc99565",
				"height": "0xb4bb8"
			}`, "PivotDecision"),
			expect: genTxWithPayload(&TransactionPayload{
				transactionType: "PivotDecision",
				PivotBlockDecision: PivotBlockDecision{
					Height:    740280,
					BlockHash: common.HexToHash("0xe18e8f0ff9dbb12d0d27451fea177f277445e244e2524e9b5ea370282fc99565"),
				},
			}, "PivotDecision"),
			actual: nil,
		},

		{
			data: genTxJsonWithPayload(`"payload": {
				"publicKey": "0xb68eb1077fd469f7a9de1c279e74a1bee1fe95ffca2bf88f093e527b0867b7f315b40611a8a10c7da04e9513156ea787",
				"targetTerm": "0x6",
				"vrfProof": "0x0375426a57bfce9c3907f3bb5f621909dc9f7376b9028e72e39bee9d2b51d48d5b0532594580541df6bd829f694b1d7ed6f37dbf7584ec146d90a03ac10312ce17dd3886d50c9e65dc6f27f69fdad59c01",
				"vrfPublicKey": "0x036f20cbf457477fc876202f4fda595ff3bc3d373fb66c23022f1edb67ede1ae34"
			}`, "Election"),
			expect: genTxWithPayload(&TransactionPayload{
				transactionType: "Election",
				ElectionPayload: ElectionPayload{
					PublicKey:    "0xb68eb1077fd469f7a9de1c279e74a1bee1fe95ffca2bf88f093e527b0867b7f315b40611a8a10c7da04e9513156ea787",
					TargetTerm:   6,
					VrfProof:     "0x0375426a57bfce9c3907f3bb5f621909dc9f7376b9028e72e39bee9d2b51d48d5b0532594580541df6bd829f694b1d7ed6f37dbf7584ec146d90a03ac10312ce17dd3886d50c9e65dc6f27f69fdad59c01",
					VrfPublicKey: "0x036f20cbf457477fc876202f4fda595ff3bc3d373fb66c23022f1edb67ede1ae34",
				},
			}, "Election"),
			actual: nil,
		},

		{
			data: genTxJsonWithPayload(`"payload": {
				"nodeId": "0xe3532f3e329b75d738d46c3356a2cbd5cd75c98b13416c5530bb05db3e1a1d89",
				"votes": "0x1"
			}`, "Retire"),
			expect: genTxWithPayload(&TransactionPayload{
				transactionType: "Retire",
				RetirePayload: RetirePayload{
					NodeId: common.HexToHash("0xe3532f3e329b75d738d46c3356a2cbd5cd75c98b13416c5530bb05db3e1a1d89"),
					Votes:  1,
				},
			}, "Retire"),
			actual: nil,
		},

		{
			data: genTxJsonWithPayload(`"payload": {
				"publicKey": "0xb68eb1077fd469f7a9de1c279e74a1bee1fe95ffca2bf88f093e527b0867b7f315b40611a8a10c7da04e9513156ea787",
				"vrfPublicKey": "0x036f20cbf457477fc876202f4fda595ff3bc3d373fb66c23022f1edb67ede1ae34"
			}`, "Register"),
			expect: genTxWithPayload(&TransactionPayload{
				transactionType: "Register",
				RegisterPayload: RegisterPayload{
					PublicKey:    "0xb68eb1077fd469f7a9de1c279e74a1bee1fe95ffca2bf88f093e527b0867b7f315b40611a8a10c7da04e9513156ea787",
					VrfPublicKey: "0x036f20cbf457477fc876202f4fda595ff3bc3d373fb66c23022f1edb67ede1ae34",
				},
			}, "Register"),
			actual: nil,
		},

		{
			data: genTxJsonWithPayload(`"payload": {
				"nodeAddress": "0x941489f181519b7e171f2b485f40909a09fb14716a2fbad5a2996735835257fe",
				"votingPower": "0x3"
			}`, "UpdateVotingPower"),
			expect: genTxWithPayload(&TransactionPayload{
				transactionType: "UpdateVotingPower",
				UpdateVotingPowerPayload: UpdateVotingPowerPayload{
					NodeAddress: common.HexToHash("0x941489f181519b7e171f2b485f40909a09fb14716a2fbad5a2996735835257fe"),
					VotingPower: 3,
				},
			}, "UpdateVotingPower"),
			actual: nil,
		},

		{
			data: genTxJsonWithPayload(`"payload": {
				"blockHash": "0xc50f035607c59ab5fb3f6fce4fed567691fa4c047d14661a5351808e80e9e0a1",
				"height": "0xc498c"
			}`, "PivotDecision"),
			expect: genTxWithPayload(&TransactionPayload{
				transactionType: "PivotDecision",
				PivotBlockDecision: PivotBlockDecision{
					BlockHash: common.HexToHash("0xc50f035607c59ab5fb3f6fce4fed567691fa4c047d14661a5351808e80e9e0a1"),
					Height:    805260,
				},
			}, "PivotDecision"),
			actual: nil,
		},

		{
			data: genTxJsonWithPayload(`"payload": {
			}`, "Dispute"),
			expect: genTxWithPayload(&TransactionPayload{
				transactionType: "Dispute",
				DisputePayload:  DisputePayload{},
			}, "Dispute"),
			actual: nil,
		},
	}

	for _, item := range table {
		// fmt.Printf("data:%+v", item.data)
		if err := json.Unmarshal([]byte(item.data), &item.actual); err != nil {
			t.Fatal(err)
		}

		if item.expect == nil || item.actual == nil {
			if item.expect != item.actual {
				t.Fatalf("expect %+v, actual %+v", item.expect, item.actual)
			}
		}

		if !reflect.DeepEqual(&item.expect.Payload, &item.actual.Payload) {
			t.Fatalf("expect %+v, actual %+v", item.expect.Payload, item.actual.Payload)
		}
	}

}

func getHashPtr(hash common.Hash) *common.Hash {
	return &hash
}

func getStrPtr(str string) *string {
	return &str
}

func genTxWithPayload(tp *TransactionPayload, txType string) *Transaction {
	baseTx := Transaction{
		BlockHash: getHashPtr(common.HexToHash("0xf747537431108da44bc6e72a379c6bfb79472729d6f9b82e312aae0ef1c0c70e")),
		From:      common.HexToHash("0x046ca462890f25ed9394ca9f92c979ff48e1738a81822ecab96d83813c1a433c"),
		Hash:      common.HexToHash("0x088dcb111055951fb4a357f0afd93a2fe7492ec4181ad638451421348cc98d8d"),
		Number:    183,
		Payload:   nil,
		Status:    getStrPtr("Executed"),
		Type:      txType,
	}
	baseTx.Payload = tp
	return &baseTx
}

func genTxJsonWithPayload(payloadJson string, txType string) string {
	return `{
		"blockHash": "0xf747537431108da44bc6e72a379c6bfb79472729d6f9b82e312aae0ef1c0c70e",
		"from": "0x046ca462890f25ed9394ca9f92c979ff48e1738a81822ecab96d83813c1a433c",
		"hash": "0x088dcb111055951fb4a357f0afd93a2fe7492ec4181ad638451421348cc98d8d",
		"number": "0xb7",` +
		payloadJson + `,
		"status": "Executed",
		"type": "` + txType + `"
	}`
}
