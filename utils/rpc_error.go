package utils

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/Conflux-Chain/go-conflux-sdk/utils/abiutil"
	"github.com/pkg/errors"

	rpcutils "github.com/openweb3/go-rpc-provider/utils"
)

type RpcError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
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

	j, err := json.Marshal(origin)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	rpcErr := &RpcError{}
	if err = json.Unmarshal(j, rpcErr); err != nil {
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
