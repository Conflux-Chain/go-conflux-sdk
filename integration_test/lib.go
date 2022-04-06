package integrationtest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/pkg/errors"
	"gotest.tools/assert"
)

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

		fmt.Printf("\n========== %s %v ==========\n", rpcName, jsonMarshalAndOrdered(params))
		// reflect call sdkFunc
		rpcReuslt, rpcError, err := reflectCall(config.client, sdkFunc, params)
		if err != nil {
			t.Fatal(err)
			continue
		}
		assert.Equal(t, jsonMarshalAndOrdered(subExamp.Result), jsonMarshalAndOrdered(rpcReuslt))
		assert.Equal(t, jsonMarshalAndOrdered(subExamp.Error), jsonMarshalAndOrdered(rpcError))
	}
}

func reflectCall(c interface{}, sdkFunc string, params []interface{}) (resp interface{}, respError interface{}, err error) {
	typeOfClient := reflect.TypeOf(c)
	if method, ok := typeOfClient.MethodByName(sdkFunc); ok {
		in := make([]reflect.Value, len(params)+1)
		in[0] = reflect.ValueOf(c)
		// params marshal/unmarshal -> func params type
		for i, param := range params {
			// unmarshal params
			pType := method.Type.In(i + 1)

			// get element type if is variadic function for last param
			if method.Type.IsVariadic() && i == method.Type.NumIn()-2 {
				pType = pType.Elem()
			}

			vPtr := reflect.New(pType).Interface()
			vPtr = convertType(param, vPtr)
			v := reflect.ValueOf(vPtr).Elem().Interface()
			in[i+1] = reflect.ValueOf(v)
		}
		out := method.Func.Call(in)
		fmt.Printf("func %v, params %v, resp type %T, respError type %T, out %v\n", sdkFunc, jsonMarshalAndOrdered(getReflectValuesInterfaces(in[1:])), out[0].Interface(), out[1].Interface(), jsonMarshalAndOrdered(getReflectValuesInterfaces(out)))
		return out[0].Interface(), out[1].Interface(), nil
	}
	return nil, nil, errors.Errorf("not found method %v", sdkFunc)
}

func getReflectValuesInterfaces(values []reflect.Value) []interface{} {
	var result []interface{}
	for _, v := range values {
		result = append(result, v.Interface())
	}
	return result
}

// cfx_getBlockByEpochNumber  GetBlockSummaryByEpoch 0x0, false
// rpc_name => func(params) sdkFuncName sdkFuncParams
func convertType(from interface{}, to interface{}) interface{} {
	jp, err := json.Marshal(from)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jp, &to)
	if err != nil {
		panic(err)
	}
	return to
}

func orderJson(j []byte, indent ...bool) []byte {
	var r interface{}
	err := json.Unmarshal(j, &r)
	if err != nil {
		panic(err)
	}

	isIndent := false
	if len(indent) > 0 {
		isIndent = indent[0]
	}

	if isIndent {
		j, err = json.MarshalIndent(r, "", "  ")
	} else {
		j, err = json.Marshal(r)
	}

	if err != nil {
		panic(err)
	}

	return j
}

func jsonMarshalAndOrdered(v interface{}, indent ...bool) string {

	j, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(orderJson(j, indent...))
}
