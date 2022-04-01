package integrationtest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"gotest.tools/assert"
)

func genCfxTestConfig() rpcTestConfig {

	var rpc2Func map[string]string = map[string]string{
		"cfx_getStatus":                         "GetStatus",
		"cfx_gasPrice":                          "GetGasPrice",
		"cfx_getNextNonce":                      "GetNextNonce",
		"cfx_epochNumber":                       "GetEpochNumber",
		"cfx_getBalance":                        "GetBalance",
		"cfx_getCode":                           "GetCode",
		"cfx_getBestBlockHash":                  "GetBestBlockHash",
		"cfx_getConfirmationRiskByHash":         "GetRawBlockConfirmationRisk",
		"cfx_sendRawTransaction":                "SendRawTransaction",
		"cfx_call":                              "Call",
		"cfx_getLogs":                           "GetLogs",
		"cfx_getTransactionByHash":              "GetTransactionByHash",
		"cfx_estimateGasAndCollateral":          "EstimateGasAndCollateral",
		"cfx_getBlocksByEpoch":                  "GetBlocksByEpoch",
		"cfx_getTransactionReceipt":             "GetTransactionReceipt",
		"cfx_getAdmin":                          "GetAdmin",
		"cfx_getSponsorInfo":                    "GetSponsorInfo",
		"cfx_getStakingBalance":                 "GetStakingBalance",
		"cfx_getCollateralForStorage":           "GetCollateralForStorage",
		"cfx_getStorageAt":                      "GetStorageAt",
		"cfx_getStorageRoot":                    "GetStorageRoot",
		"cfx_getBlockByHashWithPivotAssumption": "GetBlockByHashWithPivotAssumption",
		"cfx_checkBalanceAgainstTransaction":    "CheckBalanceAgainstTransaction",
		"cfx_getSkippedBlocksByEpoch":           "GetSkippedBlocksByEpoch",
		"cfx_getAccount":                        "GetAccountInfo",
		"cfx_getInterestRate":                   "GetInterestRate",
		"cfx_getAccumulateInterestRate":         "GetAccumulateInterestRate",
		"cfx_getBlockRewardInfo":                "GetBlockRewardInfo",
		"cfx_clientVersion":                     "GetClientVersion",
		"cfx_getDepositList":                    "GetDepositList",
		"cfx_getVoteList":                       "GetVoteList",
		"cfx_getSupplyInfo":                     "GetSupplyInfo",
		"trace_block":                           "GetBlockTraces",
		"trace_filter":                          "FilterTraces",
		"trace_transaction":                     "GetTransactionTraces",
		"cfx_getPoSRewardByEpoch":               "GetPosRewardByEpoch",
		"cfx_getAccountPendingInfo":             "GetAccountPendingInfo",
		"cfx_getAccountPendingTransactions":     "GetAccountPendingTransactions",
		"cfx_getPoSEconomics":                   "GetPoSEconomics",
		"cfx_openedMethodGroups":                "GetOpenedMethodGroups",
	}

	var rpc2FuncSelector map[string]func(params []interface{}) (string, []interface{}) = map[string]func(params []interface{}) (string, []interface{}){
		"cfx_getBlockByEpochNumber": func(params []interface{}) (string, []interface{}) {
			if params[1] == false {
				return "GetBlockSummaryByEpoch", []interface{}{params[0]}
			}
			return "GetBlockByEpoch", []interface{}{params[0]}
		},

		"cfx_getBlockByHash": func(params []interface{}) (string, []interface{}) {
			if params[1] == false {
				return "GetBlockSummaryByHash", []interface{}{params[0]}
			}
			return "GetBlockByHash", []interface{}{params[0]}
		},

		"cfx_getBlockByBlockNumber": func(params []interface{}) (string, []interface{}) {
			if params[1] == false {
				return "GetBlockSummaryByBlockNumber", []interface{}{params[0]}
			}
			return "GetBlockByBlockNumber", []interface{}{params[0]}
		},

		"cfx_getAccountPendingTransactions": func(params []interface{}) (string, []interface{}) {
			params = append(params, nil, nil)
			return "GetAccountPendingTransactions", params[:3]
		},
	}

	// ignoreRpc priority is higher than onlyTestRpc
	var ignoreRpc map[string]bool = map[string]bool{
		// "cfx_getBlockByEpochNumber": true,
		"cfx_getLogs": true,
	}

	// onlyTestRpc priority is lower than ignoreRpc
	var onlyTestRpc map[string]bool = map[string]bool{
		// "cfx_getLogs": true,
	}

	return rpcTestConfig{
		examplesUrl: "https://raw.githubusercontent.com/Conflux-Chain/jsonrpc-spec/main/src/cfx/examples.json",
		client:      sdk.MustNewClient("http://47.93.101.243"),

		rpc2Func:         rpc2Func,
		rpc2FuncSelector: rpc2FuncSelector,
		ignoreRpc:        ignoreRpc,
		onlyTestRpc:      onlyTestRpc,
	}

}

func TestClientCFX(t *testing.T) {
	os.Setenv("TESTRPC", "1")
	defer os.Unsetenv("TESTRPC")

	config := genCfxTestConfig()
	doClinetTest(t, config)
}

// request rpc
// compare result
//   order both config result and response result by their fields
//   json marshal then amd compare
func doClinetTest(t *testing.T, config rpcTestConfig) {

	rpc2Func, rpc2FuncSelector, ignoreRpc, onlyTestRpc := config.rpc2Func, config.rpc2FuncSelector, config.ignoreRpc, config.onlyTestRpc

	// read json config
	httpClient := &http.Client{}
	resp, err := httpClient.Get(config.examplesUrl)
	if err != nil {
		t.Fatal(err)
	}
	source := resp.Body
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(source)
	m := &MockRPC{}
	err = json.Unmarshal(b, m)
	if err != nil {
		t.Fatal(err)
	}

	for rpcName, subExamps := range m.Examples {
		if ignoreRpc[rpcName] {
			continue
		}

		if len(onlyTestRpc) > 0 && !onlyTestRpc[rpcName] {
			continue
		}

		subExamp := subExamps[0]

		var sdkFunc string
		var params []interface{}

		if _sdkFunc, ok := rpc2Func[rpcName]; ok {
			sdkFunc, params = _sdkFunc, subExamp.Params
		}

		if sdkFuncSelector, ok := rpc2FuncSelector[rpcName]; ok {
			sdkFunc, params = sdkFuncSelector(subExamp.Params)
		}

		if sdkFunc == "" {
			t.Fatalf("no sdk func for rpc:%s", rpcName)
			continue
		}

		fmt.Printf("\n========== %s %v ==========\n", rpcName, JsonMarshalAndOrdered(params))
		// reflect call sdkFunc
		rpcReuslt, rpcError, err := reflectCall(config.client, sdkFunc, params)
		if err != nil {
			t.Fatal(err)
			continue
		}
		assert.Equal(t, JsonMarshalAndOrdered(subExamp.Result), JsonMarshalAndOrdered(rpcReuslt))
		assert.Equal(t, JsonMarshalAndOrdered(subExamp.Error), JsonMarshalAndOrdered(rpcError))
	}
}
