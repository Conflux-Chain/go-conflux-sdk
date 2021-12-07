package abiutil

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/pkg/errors"
)

var errorABIJson = `[{
	"inputs": [
	  {
		"internalType": "string",
		"name": "err",
		"type": "string"
	  }
	],
	"name": "Error",
	"outputs": [],
	"stateMutability": "pure",
	"type": "function"
  }]`

var errorABI *abi.ABI

func DecodeErrData(data []byte) (string, error) {
	if errorABI == nil {
		var _tmp abi.ABI
		if err := _tmp.UnmarshalJSON([]byte(errorABIJson)); err != nil {
			return "", errors.WithStack(err)
		}
		errorABI = &_tmp
	}

	if len(data) < 4 {
		return "", errors.New("data must be at least 4 bytes")
	}

	method, err := errorABI.MethodById(data[:4])
	if err != nil {
		return "", errors.WithStack(err)
	}

	if method.Name != "Error" {
		return "", errors.New("not data of contract method Error(string)")
	}

	arguments, err := DecodeParameters(errorABI, "Error", data[4:])
	if err != nil {
		return "", errors.WithStack(err)
	}

	return arguments[0].(string), nil
}

func DecodeParameters(abi *abi.ABI, methodName string, data []byte) ([]interface{}, error) {
	// fmt.Printf("DecodeParameters data %v", data)
	if methodName == "" {
		// constructor
		arguments, err := abi.Constructor.Inputs.Unpack(data)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return arguments, nil
	}

	method, exist := errorABI.Methods["Error"]
	if !exist {
		return nil, fmt.Errorf("method '%s' not found", "Error")
	}

	arguments, err := method.Inputs.Unpack(data)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return arguments, nil
}
