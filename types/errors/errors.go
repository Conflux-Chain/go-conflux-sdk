package errors

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
)

const ErrMsgApplyTxValues = "failed to apply default transaction values"

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
	Code  ErrorCode
	Inner error
}

func (e BusinessError) Error() string {
	return e.Inner.Error()
}

func IsBusinessError(err error) bool {
	return reflect.TypeOf(err).Name() == "BusinessError"
}

// DetectErrorCode detect error code according to string of err.Error(), ok indicate accroding errorcode is found
func DetectErrorCode(err error) (ok bool, code ErrorCode) {
	if err != nil {
		errStr := strings.ToLower(err.Error())

		if strings.Contains(errStr, "pivot assumption failed") {
			return true, CodePivotAssumption
		}

		if strings.Contains(errStr, "block not found") {
			return true, CodeBlockNotFound
		}
	}
	return false, 0
}
