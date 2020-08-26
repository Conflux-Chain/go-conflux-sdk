package types

import (
	"fmt"
	"reflect"
)

// WrappedError for wrapping old error and new error message
type WrappedError struct {
	msg string
	err error
}

// WrapError creates a WrappedError
func WrapError(err error, msg string) error {
	return WrappedError{msg, err}
}

// WrapErrorf creates a WrappedError with formated message
func WrapErrorf(err error, msgPattern string, values ...interface{}) error {
	msg := fmt.Sprintf(msgPattern, values...)
	return WrappedError{msg, err}
}

// IsRpcJsonError returns true if err is rpc error
func IsRpcJsonError(err error) bool {
	if err == nil {
		return false
	}

	t := reflect.TypeOf(err).String()
	isJSONErr := t == "*rpc.jsonError"
	if isJSONErr {
		return true
	}

	isWrappedError := t == "types.WrappedError"
	if isWrappedError {
		err = err.(WrappedError).Unwrap()
		return IsRpcJsonError(err)
	}
	return false
}

// Error returns error description
func (e WrappedError) Error() string {
	var innerErrorMsg string
	if e.err != nil {
		innerErrorMsg = e.err.Error()

		t := reflect.TypeOf(e.err).String()
		isJSONErr := t == "*rpc.jsonError"
		if isJSONErr {
			elem := reflect.ValueOf(e.err).Elem()
			data := elem.FieldByName("Data")
			if !data.IsNil() {
				innerErrorMsg = fmt.Sprintf("%v, Data: %v", e.err.Error(), data)
			}
		}
	}
	return fmt.Sprintf("%v\n> %v", e.msg, innerErrorMsg)
}

// Unwrap for getting internal error by errors.Unwrap(wrappedError)
func (e WrappedError) Unwrap() error { return e.err }
