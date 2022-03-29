package integrationtest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/pkg/errors"
	"gotest.tools/assert"
)

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
		fmt.Printf("func %v, params %v, resp type %T, respError type %T, out %v\n", sdkFunc, JsonMarshalAndOrdered(getReflectValuesInterfaces(in[1:])), out[0].Interface(), out[1].Interface(), JsonMarshalAndOrdered(getReflectValuesInterfaces(out)))
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

func orderJson(j []byte) []byte {
	var r interface{}
	err := json.Unmarshal(j, &r)
	if err != nil {
		panic(err)
	}

	j, err = json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}

	return j
}

func JsonMarshalAndOrdered(v interface{}) string {

	j, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(orderJson(j))
}

func TestJsonMarshalMapping(t *testing.T) {
	m := map[string]interface{}{
		"c": struct {
			C3 string
			C2 string
		}{"C3", "C2"},
		"b": "A",
		"a": "A",
	}
	j, _ := json.Marshal(m)

	var m2 interface{}
	json.Unmarshal(j, &m2)
	j2, _ := json.Marshal(m2)

	assert.Equal(t, string(j), string(j2))

	fmt.Println(string(j))
}

func TestUnmarshalMockRPC(t *testing.T) {
	content, err := ioutil.ReadFile("./tmp/rpc_example.json")
	if err != nil {
		panic(err)
	}
	m := &MockRPC{}
	err = json.Unmarshal(content, m)
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
}

func TestReflectCall(t *testing.T) {
	c := sdk.MustNewClient("http://47.93.101.243")
	typeOfClient := reflect.TypeOf(c)
	if method, ok := typeOfClient.MethodByName("GetStatus"); ok {
		in := make([]reflect.Value, 1)
		in[0] = reflect.ValueOf(c)
		// for i, param := range params {
		// 	in[i] = reflect.ValueOf(param)
		// }
		// fmt.Printf("func %v, params %+v\n", sdkFunc, in)
		out := method.Func.Call(in)
		//  out[0].Interface(), out[1].Interface(), nil
		fmt.Printf("%v", out)
	}

}

func TestReflect(t *testing.T) {
	a := 1
	typea := reflect.TypeOf(a)
	vptrInReflect := reflect.New(typea)
	vptr := vptrInReflect.Interface()
	e := json.Unmarshal([]byte("1"), vptr)
	if e != nil {
		t.Fatal(e)
	}

	v := reflect.ValueOf(vptr).Elem().Interface().(int)
	fmt.Printf("%T %v\n", v, v)

	typea = reflect.TypeOf([]int{}).Elem()
	fmt.Printf("%T %v\n", typea, typea)
}
