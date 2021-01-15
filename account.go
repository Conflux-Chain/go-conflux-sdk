// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package sdk

import (
	"crypto/ecdsa"
	crand "crypto/rand"
	"encoding/hex"
	"io/ioutil"
	"strings"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

const (
	errMsgFromAddressEmpty = "from address not specified"
	errMsgCalculateTxHash  = "failed to calculate tx hash"
	errMsgSignTx           = "failed to sign tx hash"
	errMsgEncodeSignature  = "failed to encode tx signature"
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
		return "", err
	}

	cfxAddress := utils.ToCfxGeneralAddress(account.Address)
	m.cfxAddressDic[string(cfxAddress)] = &account
	return cfxAddress, nil
}

// CreateEthCompatible creates a new account compatible with eth and puts the keystore file into keystore directory
func (m *AccountManager) CreateEthCompatible(passphrase string) (types.Address, error) {
	for {
		privateKeyECDSA, err := ecdsa.GenerateKey(crypto.S256(), crand.Reader)
		if err != nil {
			return "", err
		}

		addr := crypto.PubkeyToAddress(privateKeyECDSA.PublicKey)

		if addr.Bytes()[0]&0xf0 == 0x10 {
			account, err := m.ks.ImportECDSA(privateKeyECDSA, passphrase)
			if err != nil {
				return "", err
			}
			return *types.NewAddressFromCommon(account.Address), nil
		}
	}
}

// Import imports account from external key file to keystore directory.
// Returns error if the account already exists.
func (m *AccountManager) Import(keyFile, passphrase, newPassphrase string) (types.Address, error) {
	keyJSON, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return "", errors.Wrapf(err, "failed to read key file %v", keyFile)
	}

	key, err := keystore.DecryptKey(keyJSON, passphrase)
	if err != nil {
		return "", errors.Wrap(err, "failed to decrypt key with passphrase")
	}

	if m.ks.HasAddress(key.Address) {
		return "", errors.Errorf("account already exists: %v", key.Address.String())
	}

	account, err := m.ks.Import(keyJSON, passphrase, newPassphrase)
	if err != nil {
		return "", errors.Wrap(err, "failed to import account into keystore")
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
		return "", errors.Wrap(err, "invalid HEX format of private key")
	}

	account, err := m.ks.ImportECDSA(privateKey, passphrase)
	if err != nil {
		return "", errors.Wrap(err, "failed to import private key into keystore")
	}

	cfxAddress := utils.ToCfxGeneralAddress(account.Address)
	m.cfxAddressDic[string(cfxAddress)] = &account
	return cfxAddress, nil
}

// Export exports private key string of address
func (m *AccountManager) Export(address types.Address, passphrase string) (string, error) {

	a, err := m.account(address)
	if err != nil {
		return "", err
	}

	keyjson, err := m.ks.Export(*a, passphrase, passphrase)
	if err != nil {
		return "", errors.Wrap(err, "failed to export account")
	}

	key, err := keystore.DecryptKey(keyjson, passphrase)
	if err != nil {
		return "", errors.Wrap(err, "failed to decrypt key file")
	}

	keystr := hex.EncodeToString(crypto.FromECDSA(key.PrivateKey))
	return "0x" + keystr, nil
}

// Delete deletes the specified account and remove the keystore file from keystore directory.
func (m *AccountManager) Delete(address types.Address, passphrase string) error {
	account, err := m.account(address)
	if err != nil {
		return err
	}
	return m.ks.Delete(*account, passphrase)
}

// Update updates the passphrase of specified account.
func (m *AccountManager) Update(address types.Address, passphrase, newPassphrase string) error {
	account, err := m.account(address)
	if err != nil {
		return err
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

	return nil, errors.New("no account found")
}

func (m *AccountManager) account(address types.Address) (*accounts.Account, error) {
	_, err := address.GetAddressType()
	if err != nil {
		return nil, err
	}

	realAccount := m.cfxAddressDic[strings.ToLower(string(address))]
	if realAccount == nil {
		return nil, types.NewAccountNotFoundError(address)
	}

	return realAccount, nil
}

// Unlock unlocks the specified account indefinitely.
func (m *AccountManager) Unlock(address types.Address, passphrase string) error {
	account, err := m.account(address)
	if err != nil {
		return err
	}
	return m.ks.Unlock(*account, passphrase)
}

// UnlockDefault unlocks the default account indefinitely.
func (m *AccountManager) UnlockDefault(passphrase string) error {
	defaultAccount, err := m.GetDefault()
	if err != nil {
		return err
	}
	return m.Unlock(*defaultAccount, passphrase)
}

// TimedUnlock unlocks the specified account for a period of time.
func (m *AccountManager) TimedUnlock(address types.Address, passphrase string, timeout time.Duration) error {
	account, err := m.account(address)
	if err != nil {
		return err
	}
	return m.ks.TimedUnlock(*account, passphrase, timeout)
}

// TimedUnlockDefault unlocks the specified account for a period of time.
func (m *AccountManager) TimedUnlockDefault(passphrase string, timeout time.Duration) error {
	defaultAccount, err := m.GetDefault()
	if err != nil {
		return err
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
		return nil, errors.New(errMsgFromAddressEmpty)
	}

	account, err := m.account(*tx.From)
	if err != nil {
		return nil, err
	}

	hash, err := tx.Hash()
	if err != nil {
		return nil, errors.Wrap(err, errMsgCalculateTxHash)
	}

	sig, err := m.ks.SignHash(*account, hash)
	if err != nil {
		return nil, errors.Wrap(err, errMsgSignTx)
	}

	encoded, err := tx.EncodeWithSignature(sig[64], sig[0:32], sig[32:64])
	if err != nil {
		return nil, errors.Wrap(err, errMsgEncodeSignature)
	}

	return encoded, nil
}

// SignAndEcodeTransactionWithPassphrase signs tx with given passphrase and return its RLP encoded data.
func (m *AccountManager) SignAndEcodeTransactionWithPassphrase(tx types.UnsignedTransaction, passphrase string) ([]byte, error) {
	// tx.ApplyDefault()
	if tx.From == nil {
		return nil, errors.New(errMsgFromAddressEmpty)
	}

	account, err := m.account(*tx.From)
	if err != nil {
		return nil, err
	}

	hash, err := tx.Hash()
	if err != nil {
		return nil, errors.Wrap(err, errMsgCalculateTxHash)
	}

	sig, err := m.ks.SignHashWithPassphrase(*account, passphrase, hash)
	if err != nil {
		return nil, errors.Wrap(err, errMsgSignTx)
	}

	encoded, err := tx.EncodeWithSignature(sig[64], sig[0:32], sig[32:64])
	if err != nil {
		return nil, errors.Wrap(err, errMsgEncodeSignature)
	}

	return encoded, nil
}

// SignTransactionWithPassphrase signs tx with given passphrase and returns a transction with signature
func (m *AccountManager) SignTransactionWithPassphrase(tx types.UnsignedTransaction, passphrase string) (*types.SignedTransaction, error) {
	// tx.ApplyDefault()
	if tx.From == nil {
		return nil, errors.New(errMsgFromAddressEmpty)
	}

	account, err := m.account(*tx.From)
	if err != nil {
		return nil, err
	}

	hash, err := tx.Hash()
	if err != nil {
		return nil, errors.Wrap(err, errMsgCalculateTxHash)
	}

	sig, err := m.ks.SignHashWithPassphrase(*account, passphrase, hash)
	if err != nil {
		return nil, errors.Wrap(err, errMsgSignTx)
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
		return 0, nil, nil, errors.New(errMsgFromAddressEmpty)
	}

	account, err := m.account(*tx.From)
	if account == nil {
		return 0, nil, nil, err
	}

	hash, err := tx.Hash()
	if err != nil {
		return 0, nil, nil, errors.Wrap(err, errMsgCalculateTxHash)
	}

	sig, err := m.ks.SignHashWithPassphrase(*account, passphrase, hash)
	if err != nil {
		return 0, nil, nil, errors.Wrap(err, errMsgSignTx)
	}
	v = sig[64]
	r = sig[0:32]
	s = sig[32:64]
	return v, r, s, nil
}
