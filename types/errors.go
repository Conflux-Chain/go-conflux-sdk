package types

import "fmt"

// AccountNotFoundError represents error of account not found.
type AccountNotFoundError struct {
	Account Address
}

// NewAccountNotFoundError creates a new AccountNotFoundError instance
func NewAccountNotFoundError(address Address) *AccountNotFoundError {
	return &AccountNotFoundError{
		Account: address,
	}
}

// Error implements error interface
func (e *AccountNotFoundError) Error() string {
	return fmt.Sprintf("Not found account %v", e.Account)
}
