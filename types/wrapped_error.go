package types

import "fmt"

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
	return fmt.Sprintf("%v\n> %v", e.msg, e.err.Error())
}

// Unwrap for getting internal error by errors.Unwrap(wrappedError)
func (e WrappedError) Unwrap() error { return e.err }
