// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package accounts

import (
	"crypto/ecdsa"
	crand "crypto/rand"
	"encoding/hex"
	"io/ioutil"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	sdkErrors "github.com/Conflux-Chain/go-conflux-sdk/types/errors"
	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

const (
	errMsgFromAddressEmpty = "from address not specified"
	errMsgCalculateTxHash  = "failed to calculate tx hash"
	errMsgSignTx           = "failed to sign tx hash"
	errMsgEncodeSignature  = "failed to encode tx signature"
)

var (
	emptyAccount = accounts.Account{}
)

// KeystoreWallet manages Conflux accounts.
type KeystoreWallet struct {
	ks            *keystore.KeyStore
	cfxAddressDic map[string]accounts.Account
	networkID     uint32
}

// NewKeystoreWallet creates an instance of AccountManager
// based on the keystore directory "keydir".
func NewKeystoreWallet(keydir string, networkID uint32) *KeystoreWallet {
	am := new(KeystoreWallet)
	am.networkID = networkID

	am.ks = keystore.NewKeyStore(keydir, keystore.StandardScryptN, keystore.StandardScryptP)
	am.cfxAddressDic = make(map[string]accounts.Account)

	for _, account := range am.ks.Accounts() {
		addr := getCfxUserAddress(account, networkID)
		am.cfxAddressDic[addr.GetHexAddress()] = account
	}

	return am
}

// Create creates a new account and puts the keystore file into keystore directory
func (m *KeystoreWallet) Create(passphrase string) (address types.Address, err error) {
	account, err := m.ks.NewAccount(passphrase)
	if err != nil {
		return address, err
	}

	addr := getCfxUserAddress(account, m.networkID)
	m.cfxAddressDic[addr.GetHexAddress()] = account
	return addr, nil
}

// CreateEthCompatible creates a new account compatible with eth and puts the keystore file into keystore directory
func (m *KeystoreWallet) CreateEthCompatible(passphrase string) (address types.Address, err error) {
	for {
		privateKeyECDSA, err := ecdsa.GenerateKey(crypto.S256(), crand.Reader)
		if err != nil {
			return address, err
		}

		addr := crypto.PubkeyToAddress(privateKeyECDSA.PublicKey)

		if addr.Bytes()[0]&0xf0 == 0x10 {
			account, err := m.ks.ImportECDSA(privateKeyECDSA, passphrase)
			if err != nil {
				return address, err
			}
			return cfxaddress.NewFromCommon(account.Address, m.networkID)
		}
	}
}

// Import imports account from external key file to keystore directory.
// Returns error if the account already exists.
func (m *KeystoreWallet) Import(keyFile, passphrase, newPassphrase string) (address types.Address, err error) {
	keyJSON, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return address, errors.Wrapf(err, "failed to read key file %v", keyFile)
	}

	key, err := keystore.DecryptKey(keyJSON, passphrase)
	if err != nil {
		return address, errors.Wrap(err, "failed to decrypt key with passphrase")
	}

	if m.ks.HasAddress(key.Address) {
		return address, errors.Errorf("account already exists: %v", keyFile)
	}

	account, err := m.ks.Import(keyJSON, passphrase, newPassphrase)
	if err != nil {
		return address, errors.Wrap(err, "failed to import account into keystore")
	}

	address = getCfxUserAddress(account, m.networkID)

	m.cfxAddressDic[address.GetHexAddress()] = account
	return
}

// ImportKey import account from private key hex string and save to keystore directory
func (m *KeystoreWallet) ImportKey(keyString string, passphrase string) (address types.Address, err error) {
	if utils.Has0xPrefix(keyString) {
		keyString = keyString[2:]
	}

	privateKey, err := crypto.HexToECDSA(keyString)
	if err != nil {
		return address, errors.Wrap(err, "invalid HEX format of private key")
	}

	account, err := m.ks.ImportECDSA(privateKey, passphrase)
	if err != nil {
		return address, errors.Wrap(err, "failed to import private key into keystore")
	}

	address = getCfxUserAddress(account, m.networkID)
	m.cfxAddressDic[address.GetHexAddress()] = account
	return
}

// Export exports private key string of address
func (m *KeystoreWallet) Export(address types.Address, passphrase string) (string, error) {

	a, err := m.account(address)
	if err != nil {
		return "", err
	}

	keyjson, err := m.ks.Export(a, passphrase, passphrase)
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
func (m *KeystoreWallet) Delete(address types.Address, passphrase string) error {
	account, err := m.account(address)
	if err != nil {
		return err
	}
	return m.ks.Delete(account, passphrase)
}

// Update updates the passphrase of specified account.
func (m *KeystoreWallet) Update(address types.Address, passphrase, newPassphrase string) error {
	account, err := m.account(address)
	if err != nil {
		return err
	}
	return m.ks.Update(account, passphrase, newPassphrase)
}

// List lists all accounts in keystore directory.
func (m *KeystoreWallet) List() []types.Address {
	result := make([]types.Address, 0)

	for _, account := range m.ks.Accounts() {

		cfxAddress := getCfxUserAddress(account, m.networkID)
		// fmt.Printf("list %v %v\n", m.networkID, cfxAddress)
		result = append(result, cfxAddress)
	}

	return result
}

// GetDefault returns first account in keystore directory
func (m *KeystoreWallet) GetDefault() (*types.Address, error) {
	list := m.List()
	if len(list) > 0 {
		return &list[0], nil
	}

	return nil, errors.New("no account found")
}

func (m *KeystoreWallet) account(address types.Address) (accounts.Account, error) {
	realAccount := m.cfxAddressDic[address.GetHexAddress()]
	if realAccount == emptyAccount {
		return emptyAccount, sdkErrors.NewAccountNotFoundError(address)
	}

	return realAccount, nil
}

// Unlock unlocks the specified account indefinitely.
func (m *KeystoreWallet) Unlock(address types.Address, passphrase string) error {
	account, err := m.account(address)
	if err != nil {
		return err
	}
	return m.ks.Unlock(account, passphrase)
}

// UnlockDefault unlocks the default account indefinitely.
func (m *KeystoreWallet) UnlockDefault(passphrase string) error {
	defaultAccount, err := m.GetDefault()
	if err != nil {
		return err
	}
	return m.Unlock(*defaultAccount, passphrase)
}

// TimedUnlock unlocks the specified account for a period of time.
func (m *KeystoreWallet) TimedUnlock(address types.Address, passphrase string, timeout time.Duration) error {
	account, err := m.account(address)
	if err != nil {
		return err
	}
	return m.ks.TimedUnlock(account, passphrase, timeout)
}

// TimedUnlockDefault unlocks the specified account for a period of time.
func (m *KeystoreWallet) TimedUnlockDefault(passphrase string, timeout time.Duration) error {
	defaultAccount, err := m.GetDefault()
	if err != nil {
		return err
	}
	return m.TimedUnlock(*defaultAccount, passphrase, timeout)
}

// Lock locks the specified account.
func (m *KeystoreWallet) Lock(address types.Address) error {
	common, _, err := address.ToCommon()
	if err != nil {
		return err
	}
	return m.ks.Lock(common)
}

// SignTransaction signs the hash of unsigned transaction and returns the signed transaction
func (m *KeystoreWallet) SignTransaction(tx types.UnsignedTransaction) (types.SignedTransaction, error) {
	// tx.ApplyDefault()
	empty := types.SignedTransaction{}
	if tx.From == nil {
		return empty, errors.New(errMsgFromAddressEmpty)
	}

	account, err := m.account(*tx.From)
	if err != nil {
		return empty, err
	}

	hash, err := tx.Hash()
	if err != nil {
		return empty, errors.Wrap(err, errMsgCalculateTxHash)
	}

	sig, err := m.ks.SignHash(account, hash)
	if err != nil {
		return empty, errors.Wrap(err, errMsgSignTx)
	}

	signdTx := types.SignedTransaction{}
	signdTx.UnsignedTransaction = tx
	signdTx.V = sig[64]
	signdTx.R = sig[0:32]
	signdTx.S = sig[32:64]

	return signdTx, nil
}

// SignTransactionWithPassphrase signs tx with given passphrase and returns a transction with signature
func (m *KeystoreWallet) SignTransactionWithPassphrase(tx types.UnsignedTransaction, passphrase string) (types.SignedTransaction, error) {
	// tx.ApplyDefault()
	empty := types.SignedTransaction{}
	if tx.From == nil {
		return empty, errors.New(errMsgFromAddressEmpty)
	}

	account, err := m.account(*tx.From)
	if err != nil {
		return empty, err
	}

	hash, err := tx.Hash()
	if err != nil {
		return empty, errors.Wrap(err, errMsgCalculateTxHash)
	}

	sig, err := m.ks.SignHashWithPassphrase(account, passphrase, hash)
	if err != nil {
		return empty, errors.Wrap(err, errMsgSignTx)
	}

	signdTx := types.SignedTransaction{}
	signdTx.UnsignedTransaction = tx
	signdTx.V = sig[64]
	signdTx.R = sig[0:32]
	signdTx.S = sig[32:64]

	return signdTx, nil
}

func getCfxUserAddress(account accounts.Account, networkID uint32) cfxaddress.Address {
	account.Address[0] = account.Address[0]&0x1f | 0x10
	cfxAddress := cfxaddress.MustNewFromCommon(account.Address, networkID)
	return cfxAddress
}
