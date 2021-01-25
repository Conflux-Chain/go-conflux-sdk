package types

import (
	"reflect"

	"github.com/pkg/errors"
)

// IsRPCJSONError returns true if err is rpc error
func IsRPCJSONError(err error) bool {
	t := reflect.TypeOf(err).String()

	if t == "*rpc.jsonError" {
		return true
	}

	if errors.Cause(err) == err {
		return false
	}

	return IsRPCJSONError(errors.Cause(err))
}
