package utils

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/Conflux-Chain/go-conflux-sdk/utils/abiutil"
	"github.com/pkg/errors"

	rpcutils "github.com/openweb3/go-rpc-provider/utils"
)

type RpcError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// InvalidParamsDetail returns the detail of an Invalid params RPC error.
// ok is true only when the error code is -32602, even if detail is empty.
// It prefers the detail in Data used by newer Core versions and falls back to
// Message for compatibility with older Core versions.
func (e *RpcError) InvalidParamsDetail() (detail string, ok bool) {
	if e == nil || e.Code != -32602 {
		return "", false
	}

	if data, isString := e.Data.(string); isString && data != "" {
		detail = data
		if len(data) >= 2 && data[0] == '"' && data[len(data)-1] == '"' {
			var decoded string
			if err := JsonUnmarshal([]byte(data), &decoded); err == nil {
				detail = decoded
			}
		}
		if detail != "" {
			return detail, true
		}
	}

	for _, prefix := range []string{"Invalid params:", "Invalid parameters:"} {
		if len(e.Message) >= len(prefix) && strings.EqualFold(e.Message[:len(prefix)], prefix) {
			return e.Message[len(prefix):], true
		}
	}

	if strings.EqualFold(e.Message, "Invalid params") || strings.EqualFold(e.Message, "Invalid parameters") {
		return "", true
	}

	return e.Message, true
}

func (e *RpcError) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("%v, data: %v", e.Message, e.Data)
}

// ToRpcError converts a error to JsonError
func ToRpcError(origin error) (*RpcError, error) {

	if !rpcutils.IsRPCJSONError(origin) {
		return nil, errors.New("not a rpc error")
	}

	j, err := JsonMarshal(errors.Cause(origin))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	rpcErr := &RpcError{}
	if err = JsonUnmarshal(j, rpcErr); err != nil {
		return nil, errors.WithStack(err)
	}

	hexStr, ok := rpcErr.Data.(string)
	if !ok {
		return rpcErr, nil
	}

	if !Has0xPrefix(hexStr) {
		return rpcErr, nil
	}

	hexBytes, err := hex.DecodeString(hexStr[2:])
	if err != nil {
		return rpcErr, nil
	}

	data, err := abiutil.DecodeErrData(hexBytes)
	if err != nil {
		return rpcErr, nil
	}

	rpcErr.Message += ": " + data
	return rpcErr, nil
}
