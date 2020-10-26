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
	"github.com/ethereum/go-ethereum/crypto"
)

// AccountManager manages Conflux accounts.
type AccountManager struct {
	ks            *keystore.KeyStore
	cfxAddressDic map[string]*accounts.Account
}

// NewAccountManager creates an instance of AccountManager
// based on the keystore directory "keydir".
func NewAccountManager(keydir string) *AccountManager {
	am := new(AccountManager)

	am.ks = keystore.NewKeyStore(keydir, keystore.StandardScryptN, keystore.StandardScryptP)
	am.cfxAddressDic = make(map[string]*accounts.Account)

	for _, account := range am.ks.Accounts() {
		cfxAddress := utils.ToCfxGeneralAddress(account.Address)
		tmp := account
		am.cfxAddressDic[string(cfxAddress)] = &tmp
	}

	return am
}

// Create creates a new account and puts the keystore file into keystore directory
func (m *AccountManager) Create(passphrase string) (types.Address, error) {
	account, err := m.ks.NewAccount(passphrase)
	if err != nil {
		msg := fmt.Sprintf("create account with passphrase %+v error", passphrase)
		return "", types.WrapError(err, msg)
	}

	cfxAddress := utils.ToCfxGeneralAddress(account.Address)
	m.cfxAddressDic[string(cfxAddress)] = &account
	return cfxAddress, nil
}

// Import imports account from external key file to keystore directory.
// Returns error if the account already exists.
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

	cfxAddress := utils.ToCfxGeneralAddress(account.Address)
	m.cfxAddressDic[string(cfxAddress)] = &account
	return cfxAddress, nil
}

// ImportKey import account from private key hex string and save to keystore directory
func (m *AccountManager) ImportKey(keyString string, passphrase string) (types.Address, error) {
	if utils.Has0xPrefix(keyString) {
		keyString = keyString[2:]
	}

	privateKey, err := crypto.HexToECDSA(keyString)
	if err != nil {
		msg := fmt.Sprintf("convert hexstring %v to private key error", keyString)
		return "", types.WrapError(err, msg)
	}

	account, err := m.ks.ImportECDSA(privateKey, passphrase)
	if err != nil {
		msg := fmt.Sprintf("import account by privatkey {%+v}, passphrase %+v error", keyString, passphrase)
		return "", types.WrapError(err, msg)
	}

	cfxAddress := utils.ToCfxGeneralAddress(account.Address)
	m.cfxAddressDic[string(cfxAddress)] = &account
	return cfxAddress, nil
}

// Delete deletes the specified account and remove the keystore file from keystore directory.
func (m *AccountManager) Delete(address types.Address, passphrase string) error {
	account := m.account(address)
	if account == nil {
		return nil
	}
	return m.ks.Delete(*account, passphrase)
}

// Update updates the passphrase of specified account.
func (m *AccountManager) Update(address types.Address, passphrase, newPassphrase string) error {
	account := m.account(address)
	if account == nil {
		return types.NewAccountNotFoundError(address)
	}
	return m.ks.Update(*account, passphrase, newPassphrase)
}

// List lists all accounts in keystore directory.
func (m *AccountManager) List() []types.Address {
	result := make([]types.Address, 0)

	for _, account := range m.ks.Accounts() {
		cfxAddress := utils.ToCfxGeneralAddress(account.Address)
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
	if account == nil {
		return types.NewAccountNotFoundError(address)
	}
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
	if account == nil {
		return types.NewAccountNotFoundError(address)
	}
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

// SignTransaction signs tx and returns its RLP encoded data.
func (m *AccountManager) SignTransaction(tx types.UnsignedTransaction) ([]byte, error) {
	// tx.ApplyDefault()
	if tx.From == nil {
		return nil, errors.New("From is empty, it is necessary for sign")
	}

	account := m.account(*tx.From)
	// fmt.Printf("get account of address %+v is %+v\n\n", tx.From, account)
	if account == nil {
		return nil, types.NewAccountNotFoundError(*tx.From)
	}

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

// SignAndEcodeTransactionWithPassphrase signs tx with given passphrase and return its RLP encoded data.
func (m *AccountManager) SignAndEcodeTransactionWithPassphrase(tx types.UnsignedTransaction, passphrase string) ([]byte, error) {
	// tx.ApplyDefault()
	if tx.From == nil {
		return nil, errors.New("From is empty, it is necessary for sign")
	}

	account := m.account(*tx.From)
	if account == nil {
		return nil, types.NewAccountNotFoundError(*tx.From)
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

// SignTransactionWithPassphrase signs tx with given passphrase and returns a transction with signature
func (m *AccountManager) SignTransactionWithPassphrase(tx types.UnsignedTransaction, passphrase string) (*types.SignedTransaction, error) {
	// tx.ApplyDefault()
	if tx.From == nil {
		return nil, errors.New("From is empty, it is necessary for sign")
	}

	account := m.account(*tx.From)
	if account == nil {
		return nil, types.NewAccountNotFoundError(*tx.From)
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

	signdTx := new(types.SignedTransaction)
	signdTx.UnsignedTransaction = tx
	signdTx.V = sig[64]
	signdTx.R = sig[0:32]
	signdTx.S = sig[32:64]

	return signdTx, nil
}

// Sign signs tx by passphrase and returns the signature
func (m *AccountManager) Sign(tx types.UnsignedTransaction, passphrase string) (v byte, r, s []byte, err error) {
	// tx.ApplyDefault()
	if tx.From == nil {
		return 0, nil, nil, errors.New("From is empty, it is necessary for sign")
	}

	account := m.account(*tx.From)
	if account == nil {
		// msg := fmt.Sprintf("no account of address %+v is found in keystore directory ", *tx.From)
		return 0, nil, nil, types.NewAccountNotFoundError(*tx.From)
	}

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
