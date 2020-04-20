// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package sdk

import (
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// AccountManager manages Conflux accounts.
type AccountManager struct {
	ks            *keystore.KeyStore
	cfxAddressDic map[string]*accounts.Account
}

// NewAccountManager creates an instance of AccountManager.
func NewAccountManager(keydir string) *AccountManager {
	am := new(AccountManager)

	am.ks = keystore.NewKeyStore(keydir, keystore.StandardScryptN, keystore.StandardScryptP)
	am.cfxAddressDic = make(map[string]*accounts.Account)

	for _, account := range am.ks.Accounts() {
		cfxAddress := utils.ToGeneralAddress(account.Address)
		tmp := account
		am.cfxAddressDic[string(cfxAddress)] = &tmp
	}

	return am
}

// Create creates a new account.
func (m *AccountManager) Create(passphrase string) (types.Address, error) {
	account, err := m.ks.NewAccount(passphrase)
	if err != nil {
		msg := fmt.Sprintf("create account with passphrase %+v error", passphrase)
		return "", types.WrapError(err, msg)
	}

	cfxAddress := utils.ToGeneralAddress(account.Address)
	m.cfxAddressDic[string(cfxAddress)] = &account
	return cfxAddress, nil
}

// Import imports account from external key file.
// Return error if the account already exists.
func (m *AccountManager) Import(keyFile, passphrase, newPassphrase string) (types.Address, error) {
	keyJSON, err := ioutil.ReadFile(keyFile)
	if err != nil {
		msg := fmt.Sprintf("read file %+v error", keyFile)
		return "", types.WrapError(err, msg)
	}

	key, err := keystore.DecryptKey(keyJSON, passphrase)
	if err != nil {
		msg := fmt.Sprintf("decrypt key %+v with passphrase %+v error", keyJSON, passphrase)
		return "", types.WrapError(err, msg)
	}

	if m.ks.HasAddress(key.Address) {
		return "", fmt.Errorf("account already exists: %s", hexutil.Encode(key.Address.Bytes()))
	}

	account, err := m.ks.Import(keyJSON, passphrase, newPassphrase)
	if err != nil {
		msg := fmt.Sprintf("import account by keystore {%+v}, passphrase %+v, new passphrase %+v error", keyJSON, passphrase, newPassphrase)
		return "", types.WrapError(err, msg)
	}

	cfxAddress := utils.ToGeneralAddress(account.Address)
	m.cfxAddressDic[string(cfxAddress)] = &account
	return cfxAddress, nil
}

// Delete deletes the specified account.
func (m *AccountManager) Delete(address types.Address, passphrase string) error {
	account := m.account(address)
	return m.ks.Delete(*account, passphrase)
}

// Update updates the passphrase of specified account.
func (m *AccountManager) Update(address types.Address, passphrase, newPassphrase string) error {
	account := m.account(address)
	return m.ks.Update(*account, passphrase, newPassphrase)
}

// List lists all accounts.
func (m *AccountManager) List() []types.Address {
	result := make([]types.Address, 0)

	for _, account := range m.ks.Accounts() {
		cfxAddress := utils.ToGeneralAddress(account.Address)
		result = append(result, cfxAddress)
	}

	return result
}

// GetDefault return first account in keystore directory
func (m *AccountManager) GetDefault() (*types.Address, error) {
	list := m.List()
	if len(list) > 0 {
		return &list[0], nil
	}
	msg := fmt.Sprintf("no account exist in keystore directory")
	return nil, errors.New(msg)
}

func (m *AccountManager) account(address types.Address) *accounts.Account {
	realAccount := m.cfxAddressDic[string(address)]
	return realAccount

}

// Unlock unlocks the specified account indefinitely.
func (m *AccountManager) Unlock(address types.Address, passphrase string) error {
	account := m.account(address)
	return m.ks.Unlock(*account, passphrase)
}

// UnlockDefault unlocks the default account indefinitely.
func (m *AccountManager) UnlockDefault(passphrase string) error {
	defaultAccount, err := m.GetDefault()
	if err != nil {
		return types.WrapError(err, "get default account error")
	}
	return m.Unlock(*defaultAccount, passphrase)
}

// TimedUnlock unlocks the specified account for a period of time.
func (m *AccountManager) TimedUnlock(address types.Address, passphrase string, timeout time.Duration) error {
	account := m.account(address)
	return m.ks.TimedUnlock(*account, passphrase, timeout)
}

// TimedUnlockDefault unlocks the specified account for a period of time.
func (m *AccountManager) TimedUnlockDefault(passphrase string, timeout time.Duration) error {
	defaultAccount, err := m.GetDefault()
	if err != nil {
		return types.WrapError(err, "get default account error")
	}
	return m.TimedUnlock(*defaultAccount, passphrase, timeout)
}

// Lock locks the specified account.
func (m *AccountManager) Lock(address types.Address) error {
	return m.ks.Lock(common.HexToAddress(string(address)))
}

// SignTransaction signs a transaction and return its RLP encoded data.
func (m *AccountManager) SignTransaction(tx types.UnsignedTransaction) ([]byte, error) {
	// tx.ApplyDefault()

	account := m.account(*tx.From)
	// fmt.Printf("get account of address %+v is %+v\n\n", tx.From, account)

	hash, err := tx.Hash()
	if err != nil {
		msg := fmt.Sprintf("calculate tx hash of %+v error", tx)
		return nil, types.WrapError(err, msg)
	}

	sig, err := m.ks.SignHash(*account, hash)
	if err != nil {
		msg := fmt.Sprintf("sign tx hash {%+x} by account %+v error", hash, account)
		return nil, types.WrapError(err, msg)
	}

	encoded, err := tx.EncodeWithSignature(sig[64], sig[0:32], sig[32:64])
	if err != nil {
		msg := fmt.Sprintf("encode tx %+v with signature %+v error", tx, sig)
		return nil, types.WrapError(err, msg)
	}

	return encoded, nil
}

// SignTransactionWithPassphrase signs a transaction with given passphrase and return its RLP encoded data.
func (m *AccountManager) SignTransactionWithPassphrase(tx types.UnsignedTransaction, passphrase string) ([]byte, error) {
	tx.ApplyDefault()

	account := m.account(*tx.From)
	if account == nil {
		msg := fmt.Sprintf("no account of address %+v is found in keystore directory ", *tx.From)
		return nil, errors.New(msg)
	}

	hash, err := tx.Hash()
	if err != nil {
		msg := fmt.Sprintf("calculate tx hash of %+v error", tx)
		return nil, types.WrapError(err, msg)
	}

	sig, err := m.ks.SignHashWithPassphrase(*account, passphrase, hash)
	if err != nil {
		msg := fmt.Sprintf("sign tx hash {%+x} by account %+v with passphrase %+v error", hash, account, passphrase)
		return nil, types.WrapError(err, msg)
	}

	encoded, err := tx.EncodeWithSignature(sig[64], sig[0:32], sig[32:64])
	if err != nil {
		msg := fmt.Sprintf("encode tx %+v with signature %+v error", tx, sig)
		return nil, types.WrapError(err, msg)
	}

	return encoded, nil
}

// Sign return signature of transaction
func (m *AccountManager) Sign(tx types.UnsignedTransaction, passphrase string) (v byte, r, s []byte, err error) {
	tx.ApplyDefault()
	account := m.account(*tx.From)
	// fmt.Printf("the address %+v 's account is %+v\n\n", tx.From, account)
	hash, err := tx.Hash()
	if err != nil {
		msg := fmt.Sprintf("calculate tx hash of %+v error", tx)
		return 0, nil, nil, types.WrapError(err, msg)
	}

	sig, err := m.ks.SignHashWithPassphrase(*account, passphrase, hash)
	if err != nil {
		msg := fmt.Sprintf("sign tx hash {%+x} by account %+v with passphrase %+v error", hash, account, passphrase)
		return 0, nil, nil, types.WrapError(err, msg)
	}
	v = sig[64]
	r = sig[0:32]
	s = sig[32:64]
	return v, r, s, nil
}
