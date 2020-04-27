package utils

import (
	"encoding/json"
	"fmt"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
)

// UnmarshalRPCResult unmarshals rpc result to v struct and fills in it
func UnmarshalRPCResult(result interface{}, v interface{}) error {
	encoded, err := json.Marshal(result)
	if err != nil {
		msg := fmt.Sprintf("json marshal %v error", result)
		return types.WrapError(err, msg)
	}

	if err = json.Unmarshal(encoded, v); err != nil {
		msg := fmt.Sprintf("json unmarshal %v error", encoded)
		return types.WrapError(err, msg)
	}

	return nil
}
