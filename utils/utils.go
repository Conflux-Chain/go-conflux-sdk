package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"sort"
	"strings"

	"github.com/Conflux-Chain/go-conflux-sdk/constants"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
)

// CalcBlockConfirmationRisk calculates block revert rate
func CalcBlockConfirmationRisk(rawConfirmationRisk *big.Int) *big.Float {
	riskFloat := new(big.Float).SetInt(rawConfirmationRisk)
	maxUint256Float := new(big.Float).SetInt(constants.MaxUint256)
	riskRate := new(big.Float).Quo(riskFloat, maxUint256Float)
	return riskRate
}

// IsNil sepecialy checks if interface object is nil
func IsNil(i interface{}) bool {

	if i == nil {
		return true
	}

	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

// HexStringToBytes converts hex string to bytes
func HexStringToBytes(hexStr string) (hexutil.Bytes, error) {
	if !Has0xPrefix(hexStr) {
		hexStr = "0x" + hexStr
	}
	return hexutil.Decode(hexStr)
}

// Has0xPrefix returns true if input starts with '0x' or '0X'
func Has0xPrefix(input string) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
}

// IsRPCJSONError returns true if err is rpc error
func IsRPCJSONError(err error) bool {
	t := reflect.TypeOf(errors.Cause(err)).String()
	return t == "*rpc.jsonError" || t == "rpc.jsonError"
}

// PanicIfErrf panic and reports error message with args
func PanicIfErrf(err error, msg string, args ...interface{}) {
	if err != nil {
		fmt.Printf(msg, args...)
		fmt.Println()
		panic(err)
	}
}

// PanicIfErr panic and reports error message
func PanicIfErr(err error, msg string) {
	if err != nil {
		fmt.Printf(msg)
		fmt.Println()
		panic(err)
	}
}

// PrettyJSON json marshal value and pretty with indent
func PrettyJSON(value interface{}) string {
	j, e := json.Marshal(value)
	if e != nil {
		panic(e)
	}
	var str bytes.Buffer
	_ = json.Indent(&str, j, "", "    ")
	return str.String()
}

func GetObjJsonFieldTags(obj interface{}) []string {
	val := reflect.ValueOf(obj)
	var fieldNames []string
	for i := 0; i < val.Type().NumField(); i++ {
		t := val.Type().Field(i)
		fieldName := t.Name

		if jsonTag := t.Tag.Get("json"); jsonTag != "" && jsonTag != "-" {
			fieldName = jsonTag
			if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
				fieldName = jsonTag[:commaIdx]
			}
		}
		// fmt.Printf("json tag:%v\n", t.Tag.Get("json"))

		// fmt.Println(fieldName)
		fieldNames = append(fieldNames, fieldName)
	}
	less := func(i, j int) bool {
		return strings.Compare(fieldNames[i], fieldNames[j]) > 0
	}
	sort.Slice(fieldNames, less)
	return fieldNames
}

func GetObjFileds(obj interface{}) []string {
	val := reflect.ValueOf(obj).Elem()
	var fieldNames []string
	for i := 0; i < val.NumField(); i++ {
		fieldNames = append(fieldNames, val.Type().Field(i).Name)
	}
	return fieldNames
}

func GetMapSortedKeys(m map[string]interface{}) []string {
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
		// fmt.Printf("k:%v\n", k)
	}
	less := func(i, j int) bool {
		return strings.Compare(keys[i], keys[j]) > 0
	}
	sort.Slice(keys, less)

	return keys
}
