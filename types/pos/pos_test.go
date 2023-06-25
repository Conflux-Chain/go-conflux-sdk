package postypes

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
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
					BlockHash: "0xe18e8f0ff9dbb12d0d27451fea177f277445e244e2524e9b5ea370282fc99565",
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
					BlockHash: "0xc50f035607c59ab5fb3f6fce4fed567691fa4c047d14661a5351808e80e9e0a1",
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

func TestUnmarshalLedgerInfoWithSignatures(t *testing.T) {
	expect := `{"ledgerInfo":{"commitInfo":{"epoch":"0x1234","executedStateId":"0x0000000000000000000000000000000000000000000000000000000000000000","id":"0x364c28971296b9e487023c17fd4bed577e6fbb621452e4a633926757bc53f14d","nextEpochState":{"epoch":"0x1235","verifier":{"addressToValidatorInfo":{"0x4dbe13565ee9eef1277f72d1e27b688a1087af13c721d3bf715d6e752151382b":{"publicKey":"811e103da17df9f6c0c09bad05b70c31bbbbab28d0d7e3f3a345a91989bc8cc3ec55c962d23f502d67ae6c990b3f7ae5","votingPower":"0x47","vrfPublicKey":"03ee1e0bf295eb52c699719b8977ffd3c932b923947449c8d161bbc6d46478d4ed"},"0x941f97b7f878abc15dcd289e643e639f22484d94c0908a1ff4afee53f142df0c":{"publicKey":"b23b0c6ef4e1a68a0d3a65061674a9013650c5105d1eb647b85b94e3781d034ae7e5e5419b2d71fb01b1f988412dd39d","votingPower":"0x21","vrfPublicKey":"02babc924839a01bf02a3b8de2ec9689323dabc60259a543e1c219607c1caaa537"},"0x9f77c402afa0b04b2e38d39381890387a8fd2c621c0862129146b2de335cd964":{"publicKey":"b157f238403a5b980546fd19ca48f79a2613e3e3a91d14ee69908b8816e4c53665370b2fbd0db62cc4aa0e8caeedc9b5","votingPower":"0x3b","vrfPublicKey":"02b49bbfe4e9e7523dc916c37697f8d71fefc9583691dfc453946bd408173d60a6"},"0xc035f831f4808318d19d3eaf472fe3b42b68f8646a91eed00e1311f8241a865c":{"publicKey":"b51a7f24113e37ea2196af5fcd9a6bd46e0bf861e01c341151860ee9de82092c35260ad0730fe05d41e1bf33bc1ab111","votingPower":"0x3f","vrfPublicKey":"03fe232b62992e3f70118f7c6ecce444adfa156c739e3ed2ac08baf24d873baf0a"},"0xcfc244d9dfa7f3ba4a804a5e0bbb17267a1322492d020b2d310a650f69a318ac":{"publicKey":"951c6d12712c257d7ca1a566b6bb85f6095d92be85d38d3a01af21f4d7c51f5c98b2fff27ffc5e57a81aa8725d7c7040","votingPower":"0x4a","vrfPublicKey":"03f0964ebf270aa24b17b474f7f9113373d71152ecfe4114f0eedd0b2a6b2cb7b5"}},"quorumVotingPower":"0xc9","totalVotingPower":"0x12c"},"vrfSeed":"0x6d18142638c22c9e4d214b605732526aa85b6676745f7353969b0bfb3787675e"},"pivot":{"blockHash":"0x722478fecd9ddc2572cbf5f39bc528a242c26c57af82955f4c2393dbcb1f9f0d","height":"0x5ca7b0"},"round":"0xa","timestampUsecs":"0x5f51aa3fb0255","version":"0x1beaf"},"consensusDataHash":"0x6835b2c19308ea9074cf0c14c495cff97f07ef81d97c4578f34cf685cf241053"},"signatures":{"0x4dbe13565ee9eef1277f72d1e27b688a1087af13c721d3bf715d6e752151382b":"13b122399759d82c1a5ba4da64b2883f9907ae42f79e7b2a39f180977d48fe53b856c226a03e42c1fac59a8c46b0ce1905761d9b7047d09a641d23f0eda310cd8ca71eccbf58c707aa717de246da20847de9ff92201b4630dfc4b45fffe8986a178572b70cbdbbfd9276db5b88c90d17b88f189ebc398e85f471ed839a5d86f7e6c7956b46cd23a0ecf85d595b33592214e274778e2d29d5d71cd5aee5882a3c226ee84197e49d842788489798edd7c7a4f0844cb18058161e6c25dfd3b16832","0x941f97b7f878abc15dcd289e643e639f22484d94c0908a1ff4afee53f142df0c":"12dd5f36d363aeacb7403dbe1935d700b46700b673b89a57c50e07ca05f860f5e2361a65e839c5128516d5692ead813f065fd67e2c7bb4904bdd3f497883c559c4361d6194ac37c65f3e2884a8217d62900ebb853739451de2a14e766cdd92b918be1733f4ff5a631bfaffaab4144659f764311b7d3706221b957a6b0f0a7408df29cc2df47713e137b2eb0a5f0bc4f9129101669b838645e03437633f39a73dffa107416176965fa50f203e1d44cd1b419d0a08bb1b38d1d8029251d6604504","0x9f77c402afa0b04b2e38d39381890387a8fd2c621c0862129146b2de335cd964":"13dddfc3e5b0bd13b4b76a3a487b4b605a9d4b2f5c478a27fbc0f33df0f7bcdbbd98dfa54962937a92f060d0ff47b20d0b1be3a6bdc389cc66988bc28b499405c092fa5be9947765bc53c85b87cd13a4786cda654fa92ae4a46b6b14f3d3856c05f97ba9c1dcb847c27147a6f883ba56921f144995b78ee7a4250242d064bf00dd4ed7ca9abed2c329d2b49714aab647001e5ad41ab58205d1a0c3ccf4c4ce08bf1e18eaf20f58d26d84c806395481c043ddae95c841107fb62cf31e45b054ee","0xc035f831f4808318d19d3eaf472fe3b42b68f8646a91eed00e1311f8241a865c":"01daa63518edbbaadab02d561ee004e5f8ae1d003798441877a455847e9bbcbb85035ef8b2b05e5ebe9fda06769547ca198ed8e6ac98b53fdf85a1bb5f60d935a260d67a21f2c5ff004f867d9ad4b4db1378413bfd3bc82110a8eb190462a97b0909c26a1b31d3e6c45bbc0b527165b8c4b788c8279b33f7c6a03c6f26745e4a982864abfe499cedc2829a507575a3561719cccc31edd4b3ec91339371c2ee5d533b5595967ea62bfc9eebad0d7e44b78bd2094e50d3729373a6fa2df275c169","0xcfc244d9dfa7f3ba4a804a5e0bbb17267a1322492d020b2d310a650f69a318ac":"0982c02fbf4c267d88835cd669a8ad38935aca0d21b782b5c9e8c164591244be19faaeb770691ff8502d23880d4e60d009b21f7cace14eb7af7395887b576bc0c861bde2d1dd4fbc311947c1d9e6a5b8324293f936e687fbd8e21292aa1eb26d08f15f610c4579a240fad3d0ad8d6f45dce06f0d3345e90362cc4fc39f6dbaae1fd65b9380c987542e167d51b2c949740accaa882c29669d613eab6a0837c169597bc1ff68bf34e062582e5c668c547aca0e2ef2ef8dc6ebd64064cbddf58d10"}}`
	var l LedgerInfoWithSignatures
	err := json.Unmarshal([]byte(expect), &l)
	assert.NoError(t, err)

	actual, err := json.Marshal(l)
	assert.NoError(t, err)

	fExpect, fActual := utils.FormatJson(expect), utils.FormatJson(string(actual))
	assert.Equal(t, fExpect, fActual)
}

func TestUnmarshalEpochState(t *testing.T) {
	expect := `{"epoch":"0x1234","verifier":{"addressToValidatorInfo":{"0x4dbe13565ee9eef1277f72d1e27b688a1087af13c721d3bf715d6e752151382b":{"publicKey":"811e103da17df9f6c0c09bad05b70c31bbbbab28d0d7e3f3a345a91989bc8cc3ec55c962d23f502d67ae6c990b3f7ae5","votingPower":"0x45","vrfPublicKey":"03ee1e0bf295eb52c699719b8977ffd3c932b923947449c8d161bbc6d46478d4ed"},"0x941f97b7f878abc15dcd289e643e639f22484d94c0908a1ff4afee53f142df0c":{"publicKey":"b23b0c6ef4e1a68a0d3a65061674a9013650c5105d1eb647b85b94e3781d034ae7e5e5419b2d71fb01b1f988412dd39d","votingPower":"0x22","vrfPublicKey":"02babc924839a01bf02a3b8de2ec9689323dabc60259a543e1c219607c1caaa537"},"0x9f77c402afa0b04b2e38d39381890387a8fd2c621c0862129146b2de335cd964":{"publicKey":"b157f238403a5b980546fd19ca48f79a2613e3e3a91d14ee69908b8816e4c53665370b2fbd0db62cc4aa0e8caeedc9b5","votingPower":"0x3d","vrfPublicKey":"02b49bbfe4e9e7523dc916c37697f8d71fefc9583691dfc453946bd408173d60a6"},"0xc035f831f4808318d19d3eaf472fe3b42b68f8646a91eed00e1311f8241a865c":{"publicKey":"b51a7f24113e37ea2196af5fcd9a6bd46e0bf861e01c341151860ee9de82092c35260ad0730fe05d41e1bf33bc1ab111","votingPower":"0x3e","vrfPublicKey":"03fe232b62992e3f70118f7c6ecce444adfa156c739e3ed2ac08baf24d873baf0a"},"0xcfc244d9dfa7f3ba4a804a5e0bbb17267a1322492d020b2d310a650f69a318ac":{"publicKey":"951c6d12712c257d7ca1a566b6bb85f6095d92be85d38d3a01af21f4d7c51f5c98b2fff27ffc5e57a81aa8725d7c7040","votingPower":"0x4a","vrfPublicKey":"03f0964ebf270aa24b17b474f7f9113373d71152ecfe4114f0eedd0b2a6b2cb7b5"}},"quorumVotingPower":"0xc9","totalVotingPower":"0x12c"},"vrfSeed":"0xc1db63a832a8aa3ee92623359ddcf0acdecc26c4d08c4b50f47cd72cc52cfbe0"}`
	var l EpochState
	err := json.Unmarshal([]byte(expect), &l)
	assert.NoError(t, err)

	actual, err := json.Marshal(l)
	assert.NoError(t, err)

	fExpect, fActual := utils.FormatJson(expect), utils.FormatJson(string(actual))
	assert.Equal(t, fExpect, fActual)
}

// {"epoch":"0x1234","verifier":{"addressToValidatorInfo":{"0x4dbe13565ee9eef1277f72d1e27b688a1087af13c721d3bf715d6e752151382b":{"publicKey":"811e103da17df9f6c0c09bad05b70c31bbbbab28d0d7e3f3a345a91989bc8cc3ec55c962d23f502d67ae6c990b3f7ae5","votingPower":"0x45","vrfPublicKey":"03ee1e0bf295eb52c699719b8977ffd3c932b923947449c8d161bbc6d46478d4ed"},"0x941f97b7f878abc15dcd289e643e639f22484d94c0908a1ff4afee53f142df0c":{"publicKey":"b23b0c6ef4e1a68a0d3a65061674a9013650c5105d1eb647b85b94e3781d034ae7e5e5419b2d71fb01b1f988412dd39d","votingPower":"0x22","vrfPublicKey":"02babc924839a01bf02a3b8de2ec9689323dabc60259a543e1c219607c1caaa537"},"0x9f77c402afa0b04b2e38d39381890387a8fd2c621c0862129146b2de335cd964":{"publicKey":"b157f238403a5b980546fd19ca48f79a2613e3e3a91d14ee69908b8816e4c53665370b2fbd0db62cc4aa0e8caeedc9b5","votingPower":"0x3d","vrfPublicKey":"02b49bbfe4e9e7523dc916c37697f8d71fefc9583691dfc453946bd408173d60a6"},"0xc035f831f4808318d19d3eaf472fe3b42b68f8646a91eed00e1311f8241a865c":{"publicKey":"b51a7f24113e37ea2196af5fcd9a6bd46e0bf861e01c341151860ee9de82092c35260ad0730fe05d41e1bf33bc1ab111","votingPower":"0x3e","vrfPublicKey":"03fe232b62992e3f70118f7c6ecce444adfa156c739e3ed2ac08baf24d873baf0a"},"0xcfc244d9dfa7f3ba4a804a5e0bbb17267a1322492d020b2d310a650f69a318ac":{"publicKey":"951c6d12712c257d7ca1a566b6bb85f6095d92be85d38d3a01af21f4d7c51f5c98b2fff27ffc5e57a81aa8725d7c7040","votingPower":"0x4a","vrfPublicKey":"03f0964ebf270aa24b17b474f7f9113373d71152ecfe4114f0eedd0b2a6b2cb7b5"}},"quorumVotingPower":"0xc9","totalVotingPower":"0x12c"},"vrfSeed":"0xc1db63a832a8aa3ee92623359ddcf0acdecc26c4d08c4b50f47cd72cc52cfbe0"}
