package types

import (
	"reflect"

	"github.com/pkg/errors"
)

// IsRpcJsonError returns true if err is rpc error
func IsRpcJsonError(err error) bool {
	// get the underlying error in case of error wrapped
	if err = errors.Cause(err); err == nil {
		return false
	}

	t := reflect.TypeOf(err).String()

	return t == "*rpc.jsonError"
}
