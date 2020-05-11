package types

import (
	"fmt"
	"reflect"
	"strings"
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

// Error returns error description
func (e WrappedError) Error() string {
	var innerErrorMsg string
	if e.err != nil {
		t := reflect.TypeOf(e.err).String()
		isJSONErr := strings.Contains(t, "rpc.jsonError")
		if isJSONErr {
			elem := reflect.ValueOf(e.err).Elem()
			data := elem.FieldByName("Data")
			innerErrorMsg = fmt.Sprintf("%v, Data: %v", e.err.Error(), data)
		} else {
			innerErrorMsg = e.err.Error()
		}
	}
	return fmt.Sprintf("%v\n> %v", e.msg, innerErrorMsg)
}

// Unwrap for getting internal error by errors.Unwrap(wrappedError)
func (e WrappedError) Unwrap() error { return e.err }
