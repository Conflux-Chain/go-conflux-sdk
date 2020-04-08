// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package sdk

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// AccountManager manages Conflux accounts.
type AccountManager struct {
	ks *keystore.KeyStore
}

// NewAccountManager creates an instance of AccountManager.
func NewAccountManager(keydir string) *AccountManager {
	return &AccountManager{
		ks: keystore.NewKeyStore(keydir, keystore.StandardScryptN, keystore.StandardScryptP),
	}
}

// Create creates a new account.
func (m *AccountManager) Create(passphrase string) (types.Address, error) {
	account, err := m.ks.NewAccount(passphrase)
	if err != nil {
		return "", err
	}

	return types.Address(hexutil.Encode(account.Address.Bytes())), nil
}

// Import imports account from external key file.
// Return error if the account already exists.
func (m *AccountManager) Import(keyFile, passphrase, newPassphrase string) (types.Address, error) {
	keyJSON, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return "", err
	}

	key, err := keystore.DecryptKey(keyJSON, passphrase)
	if err != nil {
		return "", err
	}

	if m.ks.HasAddress(key.Address) {
		return "", fmt.Errorf("account already exists: %s", hexutil.Encode(key.Address.Bytes()))
	}

	account, err := m.ks.Import(keyJSON, passphrase, newPassphrase)
	if err != nil {
		return "", err
	}

	return types.Address(hexutil.Encode(account.Address.Bytes())), nil
}

// Delete deletes the specified account.
func (m *AccountManager) Delete(address types.Address, passphrase string) error {
	account := m.account(address)
	return m.ks.Delete(account, passphrase)
}

// Update updates the passphrase of specified account.
func (m *AccountManager) Update(address types.Address, passphrase, newPassphrase string) error {
	account := m.account(address)
	return m.ks.Update(account, passphrase, newPassphrase)
}

// List lists all accounts.
func (m *AccountManager) List() []types.Address {
	result := make([]types.Address, 0)

	for _, account := range m.ks.Accounts() {
		address := types.Address(hexutil.Encode(account.Address.Bytes()))
		result = append(result, address)
	}

	return result
}

func (m *AccountManager) account(address types.Address) accounts.Account {
	return accounts.Account{
		Address: common.HexToAddress(string(address)),
		URL: accounts.URL{
			Scheme: keystore.KeyStoreScheme,
			Path:   "",
		},
	}
}

// Unlock unlocks the specified account indefinitely.
func (m *AccountManager) Unlock(address types.Address, passphrase string) error {
	account := m.account(address)
	return m.ks.Unlock(account, passphrase)
}

// TimedUnlock unlocks the specified account for a period of time.
func (m *AccountManager) TimedUnlock(address types.Address, passphrase string, timeout time.Duration) error {
	account := m.account(address)
	return m.ks.TimedUnlock(account, passphrase, timeout)
}

// Lock locks the specified account.
func (m *AccountManager) Lock(address types.Address) error {
	return m.ks.Lock(common.HexToAddress(string(address)))
}

// SignTransaction signs a transaction and return its RLP encoded data.
func (m *AccountManager) SignTransaction(tx types.UnsignedTransaction) ([]byte, error) {
	tx.ApplyDefault()

	account := m.account(tx.From)
	sig, err := m.ks.SignHash(account, tx.Hash())
	if err != nil {
		return nil, err
	}

	encoded, err := tx.EncodeWithSignature(sig[64], sig[0:32], sig[32:64])
	if err != nil {
		return nil, err
	}

	return encoded, nil
}

// SignTransactionWithPassphrase signs a transaction with given passphrase and return its RLP encoded data.
func (m *AccountManager) SignTransactionWithPassphrase(tx types.UnsignedTransaction, passphrase string) ([]byte, error) {
	tx.ApplyDefault()

	account := m.account(tx.From)
	sig, err := m.ks.SignHashWithPassphrase(account, passphrase, tx.Hash())
	if err != nil {
		return nil, err
	}

	encoded, err := tx.EncodeWithSignature(sig[64], sig[0:32], sig[32:64])
	if err != nil {
		return nil, err
	}

	return encoded, nil
}

// Sign return signature of transaction
func (m *AccountManager) Sign(tx types.UnsignedTransaction, passphrase string) (v byte, r, s []byte, err error) {
	tx.ApplyDefault()
	account := m.account(tx.From)
	sig, err := m.ks.SignHashWithPassphrase(account, passphrase, tx.Hash())
	if err != nil {
		return 0, []byte{}, []byte{}, err
	}
	v = sig[64]
	r = sig[0:32]
	s = sig[32:64]
	return v, r, s, nil
}
