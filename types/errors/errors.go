package errors

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
)

// AccountNotFoundError represents error of account not found.
type AccountNotFoundError struct {
	Account types.Address
}

// NewAccountNotFoundError creates a new AccountNotFoundError instance
func NewAccountNotFoundError(address types.Address) *AccountNotFoundError {
	return &AccountNotFoundError{
		Account: address,
	}
}

// Error implements error interface
func (e *AccountNotFoundError) Error() string {
	return fmt.Sprintf("Not found account %v", e.Account)
}

type BusinessError struct {
	Code    ErrorCode
	Message string
}

func (e BusinessError) Error() string {
	return e.Message
}

func IsBusinessError(err error) bool {
	return reflect.TypeOf(err).Name() == "BusinessError"
}

func IsPivotSwitch(err error) bool {
	if err != nil {
		errStr := strings.ToLower(err.Error())
		// pivot hash assumption failed, must be pivot switched
		return strings.Contains(errStr, "pivot assumption failed") || strings.Contains(errStr, "block not found")
	}
	return false
}
