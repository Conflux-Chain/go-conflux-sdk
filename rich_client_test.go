package sdk

// import (
// 	"encoding/json"
// 	"reflect"
// 	"testing"

// 	"github.com/Conflux-Chain/go-conflux-sdk/types"
// 	scantypes "github.com/Conflux-Chain/go-conflux-sdk/types/scan"
// 	testutils "github.com/Conflux-Chain/go-conflux-sdk/utils/test_utils"
// )

// func TestGet(t *testing.T) {
// 	type student struct {
// 		Name string `json:"name"`
// 		Age  uint   `json:"age"`
// 	}

// 	expect := scantypes.Response{
// 		Code:    0,
// 		Message: "good",
// 		Result: student{
// 			Name: "xiaohong",
// 			Age:  10,
// 		},
// 	}

// 	var httpRequster testutils.HttpClientMock
// 	rspBody, _ := json.Marshal(expect)
// 	httpRequster.SetHandler("", string(rspBody))

// 	var stu student
// 	s := ScanServer{
// 		Scheme:        "http",
// 		HostName:      "test",
// 		HTTPRequester: &httpRequster,
// 	}

// 	err := s.Get("/test", nil, &stu)
// 	if err != nil {
// 		t.Error(err.Error())
// 	}

// 	if !reflect.DeepEqual(expect.Result, stu) {
// 		t.Errorf("expect:%+v,actual:%v", expect.Result, stu)
// 	}
// }

// func TestGetTokenByIdentifier(t *testing.T) {
// 	expect := scantypes.Response{
// 		Code:    0,
// 		Message: "good",
// 		Result: scantypes.Contract{
// 			TokenName:     "miniERC20",
// 			TokenSymbol:   "ERC20",
// 			TokenDecimals: 10,
// 			TypeCode:      uint(1),
// 		},
// 	}

// 	// mock
// 	var httpRequster testutils.HttpClientMock
// 	rspBody, _ := json.Marshal(expect)
// 	httpRequster.SetHandler("", string(rspBody))
// 	contractManager.HTTPRequester = &httpRequster

// 	contract, err := GetTokenByIdentifier(*types.NewAddress("0x8f08cb92a481c72ba69c869ad7dc5b9e5320fa70"))
// 	if err != nil {
// 		t.Error(err.Error())
// 		return
// 	}
// 	if !reflect.DeepEqual(expect.Result, *contract) {
// 		t.Errorf("expect:%+v,actual:%+v", expect.Result, *contract)
// 	}
// }

// func TestGetAccountTokens(t *testing.T) {
// 	// mock cfxScanBackend+txListPath results for main coin transaction
// 	main_coin_mock_response := ``

// 	main_coin_expect := scantypes.Contract{
// 		TokenName:     "miniERC20",
// 		TokenSymbol:   "ERC20",
// 		TokenDecimals: 10,
// 		TypeCode:      uint(1),
// 	}

// 	// mock cfxScanBackend+accountTokenTxListPath results for token transfer events
// 	token_mock_response := `{
// 		"code": 0,
// 		"message": "good",
// 		"result": {
// 		"total": "100",
// 		"listLimit": 10,
// 		"list": [
// 		{
// 			"transactionHash": "0x2ed4fcc7f6c7317a733123ebf1c931da3e190bfc6bc48c5d6602c45886a4b878",
// 			"status": "0",
// 			"from": "0x160ebef20c1f739957bf9eecd040bce699cc42c6",
// 			"to": "0x170ddf9b9750c575db453eea6a041f4c8536785a",
// 			"value": "10000000",
// 			"timestamp": 1234567890,
// 			"tokenName": "token name",
// 			"tokenSymbol": "tkn",
// 			"tokenDecimal": 10,
// 			"tokenIcon": "",
// 			"address": "0x160ebef20c1f739957bf9eecd040bce699cc42c9",
// 			"typeCode": 100
// 			}
// 		]
// 		}
// 	}`

// 	// mock get transaction result by hash "0x2ed4fcc7f6c7317a733123ebf1c931da3e190bfc6bc48c5d6602c45886a4b878"
// 	// mock get block result by hash ""

// 	token_expect := scantypes.TokenTransferEventList{
// 		Total:     100,
// 		ListLimit: 10,
// 		List: []scantypes.TokenTransferEvent{
// 			scantypes.TokenTransferEvent{
// 				Token: scantypes.Token{
// 					TokenName:    "token name",
// 					TokenSymbol:  "tkn",
// 					TokenDecimal: 10,
// 					Address:      types.NewAddress("0x160ebef20c1f739957bf9eecd040bce699cc42c9"),
// 					TypeCode:     100,
// 				},
// 				TransactionHash: "0x2ed4fcc7f6c7317a733123ebf1c931da3e190bfc6bc48c5d6602c45886a4b878",
// 				Status:          0,
// 				From:            types.Address("0x160ebef20c1f739957bf9eecd040bce699cc42c6"),
// 				To:              types.Address("0x170ddf9b9750c575db453eea6a041f4c8536785a"),
// 				Value:           "10000000",
// 				Timestamp:       1234567890,
// 				BlockHash:       "",
// 				RevertRate:      0,
// 			},
// 		},
// 	}

// 	var httpRequster testutils.HttpClientMock
// 	rspBody := token_mock_response
// 	httpRequster.SetHandler("", string(rspBody))
// 	cfxScanBackend.HTTPRequester = &httpRequster

// 	tks, err := GetAccountTokens(*types.NewAddress("0x1b6487d1db89869bc79879cc4dedee24cc2b9518"))
// 	if err != nil {
// 		t.Error(err.Error())
// 	}
// 	t.Errorf("%+v", tks)
// }
