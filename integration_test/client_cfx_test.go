package integrationtest

import (
	"os"
	"testing"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/openweb3/go-sdk-common/rpctest"
)

func genCfxTestConfig() rpctest.RpcTestConfig {

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
		"cfx_getPoSRewardByEpoch":               "GetPoSRewardByEpoch",
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
	var ignoreRpc map[string]bool = map[string]bool{}

	var ignoreExamples map[string]bool = map[string]bool{
		"cfx_getBlockByEpochNumber-1649303708800": true, // TODO: Epoch number type is U256 and conflux-rust is U64
	}

	// onlyTestRpc priority is lower than ignoreRpc
	var onlyTestRpc map[string]bool = map[string]bool{}

	return rpctest.RpcTestConfig{
		ExamplesUrl: "https://raw.githubusercontent.com/Conflux-Chain/jsonrpc-spec/main/src/cfx/examples.json",
		Client:      sdk.MustNewClient("http://47.93.101.243", sdk.ClientOption{Logger: os.Stdout}),

		Rpc2Func:         rpc2Func,
		Rpc2FuncSelector: rpc2FuncSelector,
		IgnoreRpcs:       ignoreRpc,
		IgnoreExamples:   ignoreExamples,
		OnlyTestRpcs:     onlyTestRpc,
	}

}

func TestClientCFX(t *testing.T) {
	cfxaddress.SetConfig(cfxaddress.Config{
		AddressStringVerbose: true,
	})

	config := genCfxTestConfig()
	rpctest.DoClientTest(t, config)
}
