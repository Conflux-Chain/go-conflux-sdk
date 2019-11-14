// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package sdk

import (
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// AccountManager manages Conflux accounts.
type AccountManager struct {
	client *Client
}

// NewAccountManager creates an instance of AccountManager.
func NewAccountManager(client *Client) *AccountManager {
	return &AccountManager{client}
}

// Create creates a new account.
func (m *AccountManager) Create(password string) (types.Address, error) {
	result, err := m.client.Debug("new_account", password)
	if err != nil {
		return "", err
	}

	return types.Address(result.(string)), nil
}

// List lists all accounts.
func (m *AccountManager) List() ([]types.Address, error) {
	result, err := m.client.Debug("accounts")
	if err != nil {
		return nil, err
	}

	var accounts []types.Address
	unmarshalRPCResult(result, &accounts)

	return accounts, nil
}

// Unlock unlocks the specified account.
func (m *AccountManager) Unlock(address types.Address, password string, secs ...uint) (bool, error) {
	args := []interface{}{address, password}
	if len(secs) > 0 {
		args = append(args, hexutil.Uint(secs[0]))
	}

	result, err := m.client.Debug("unlock_account", args...)
	if err != nil {
		return false, err
	}

	return result.(bool), nil
}

// Lock locks the specified account.
func (m *AccountManager) Lock(address types.Address) (bool, error) {
	result, err := m.client.Debug("lock_account", address)
	if err != nil {
		return false, err
	}

	return result.(bool), nil
}

// SendTransaction sends transaction and return its hash.
// If password is not provided the account must be unlocked.
func (m *AccountManager) SendTransaction(tx UnsignedTransaction, password ...string) (string, error) {
	tx.applyDefault()

	args := []interface{}{tx}
	if len(password) > 0 {
		args = append(args, password[0])
	}

	result, err := m.client.Debug("send_transaction", args...)
	if err != nil {
		return "", err
	}

	return result.(string), nil
}

// UnsignedTransaction represents a transaction without signature.
type UnsignedTransaction struct {
	From     types.Address  `json:"from"`
	To       *types.Address `json:"to,omitempty"`
	Nonce    *hexutil.Big   `json:"nonce,omitempty"`
	GasPrice *hexutil.Big   `json:"gasPrice"`
	Gas      *hexutil.Big   `json:"gas"`
	Value    *hexutil.Big   `json:"value"`
	Data     string         `json:"data,omitempty"`
}

// DefaultGas is the default gas in a transaction to transfer amount.
const DefaultGas int64 = 21000

// DefaultGasPrice is the default gas price.
var DefaultGasPrice = types.NewBigInt(10000000000) // 10G drip

// ApplyDefault applies default values for empty field.
func (tx *UnsignedTransaction) applyDefault() {
	if tx.GasPrice == nil {
		tx.GasPrice = DefaultGasPrice
	}

	if tx.Gas == nil {
		tx.Gas = types.NewBigInt(DefaultGas)
	}

	if tx.Value == nil {
		tx.Value = types.NewBigInt(0)
	}
}
