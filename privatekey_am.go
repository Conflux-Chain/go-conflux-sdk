package sdk

import (
	"crypto/ecdsa"
	"encoding/hex"
	"sync"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	sdkErrors "github.com/Conflux-Chain/go-conflux-sdk/types/errors"
	"github.com/Conflux-Chain/go-conflux-sdk/utils/addressutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/openweb3/go-sdk-common/privatekeyhelper"
	"github.com/pkg/errors"
)

type PrivatekeyAccountManager struct {
	accountsMap map[string]*ecdsa.PrivateKey
	addresses   []types.Address
	mutex       sync.Mutex
	networkID   uint32
}

func NewPrivatekeyAccountManager(privateKeys []string, networkID uint32) *PrivatekeyAccountManager {
	p := &PrivatekeyAccountManager{
		networkID:   networkID,
		accountsMap: make(map[string]*ecdsa.PrivateKey),
	}

	for _, k := range privateKeys {
		p.ImportKey(k, "")
	}
	return p
}

func (p *PrivatekeyAccountManager) Create(passphrase string) (types.Address, error) {
	key, err := privatekeyhelper.NewRandom()
	if err != nil {
		return types.Address{}, err
	}
	addr := p.pushAccount(key)
	return addr, nil
}

func (p *PrivatekeyAccountManager) Import(keyFile string, passphrase string, newPassphrase string) (types.Address, error) {
	key, err := privatekeyhelper.NewFromKeystoreFile(keyFile, passphrase)
	if err != nil {
		return types.Address{}, err
	}
	addr := p.pushAccount(key)
	return addr, nil
}

func (p *PrivatekeyAccountManager) ImportKey(keyString string, passphrase string) (types.Address, error) {
	key, err := privatekeyhelper.NewFromKeyString(keyString)
	if err != nil {
		return types.Address{}, err
	}
	addr := p.pushAccount(key)
	return addr, nil
}

func (p *PrivatekeyAccountManager) Export(address types.Address, passphrase string) (string, error) {
	if !p.Contains(address) {
		return "", sdkErrors.NewAccountNotFoundError(address)
	}

	keystr := hex.EncodeToString(crypto.FromECDSA(p.accountsMap[address.String()]))
	return "0x" + keystr, nil
}

func (p *PrivatekeyAccountManager) Delete(address types.Address, passphrase string) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if !p.Contains(address) {
		return sdkErrors.NewAccountNotFoundError(address)
	}

	delete(p.accountsMap, address.GetHexAddress())
	for i, addr := range p.addresses {
		if addr.Equals(&address) {
			p.addresses = append(p.addresses[:i], p.addresses[i+1:]...)
			break
		}
	}

	return nil
}

func (p *PrivatekeyAccountManager) Update(address types.Address, passphrase string, newPassphrase string) error {
	return nil
}

func (p *PrivatekeyAccountManager) List() []types.Address {
	return p.addresses
}

func (p *PrivatekeyAccountManager) Contains(address types.Address) bool {
	_, ok := p.accountsMap[address.GetHexAddress()]
	return ok
}

func (p *PrivatekeyAccountManager) GetDefault() (*types.Address, error) {
	if len(p.List()) == 0 {
		return nil, sdkErrors.ErrEmptyAddresses
	}
	return &p.List()[0], nil
}

func (p *PrivatekeyAccountManager) Unlock(address types.Address, passphrase string) error {
	return nil
}

func (p *PrivatekeyAccountManager) UnlockDefault(passphrase string) error {
	return nil
}

func (p *PrivatekeyAccountManager) TimedUnlock(address types.Address, passphrase string, timeout time.Duration) error {
	return nil
}

func (p *PrivatekeyAccountManager) TimedUnlockDefault(passphrase string, timeout time.Duration) error {
	return nil
}

func (p *PrivatekeyAccountManager) Lock(address types.Address) error {
	return nil
}

func (p *PrivatekeyAccountManager) SignTransaction(tx types.UnsignedTransaction) ([]byte, error) {
	return p.SignAndEcodeTransactionWithPassphrase(tx, "")
}

func (p *PrivatekeyAccountManager) SignAndEcodeTransactionWithPassphrase(tx types.UnsignedTransaction, passphrase string) ([]byte, error) {
	v, r, s, err := p.Sign(tx, passphrase)
	if err != nil {
		return nil, err
	}

	encoded, err := tx.EncodeWithSignature(v, r, s)
	if err != nil {
		return nil, errors.Wrap(err, errMsgEncodeSignature)
	}

	return encoded, nil
}

func (p *PrivatekeyAccountManager) SignTransactionWithPassphrase(tx types.UnsignedTransaction, passphrase string) (types.SignedTransaction, error) {
	v, r, s, err := p.Sign(tx, passphrase)
	if err != nil {
		return types.SignedTransaction{}, err
	}
	signdTx := types.SignedTransaction{}
	signdTx.UnsignedTransaction = tx
	signdTx.V = v
	signdTx.R = r
	signdTx.S = s
	return signdTx, nil
}

func (p *PrivatekeyAccountManager) Sign(tx types.UnsignedTransaction, passphrase string) (v byte, r []byte, s []byte, err error) {

	if tx.From == nil {
		return 0, nil, nil, errors.New(errMsgFromAddressEmpty)
	}

	hash, err := tx.Hash()
	if err != nil {
		return 0, nil, nil, errors.Wrap(err, errMsgCalculateTxHash)
	}

	if !p.Contains(*tx.From) {
		return 0, nil, nil, sdkErrors.NewAccountNotFoundError(*tx.From)
	}

	sig, err := crypto.Sign(hash, p.accountsMap[tx.From.GetHexAddress()])
	if err != nil {
		return 0, nil, nil, errors.Wrap(err, errMsgSignTx)
	}

	v = sig[64]
	r = sig[0:32]
	s = sig[32:64]
	return v, r, s, nil
}

func (p *PrivatekeyAccountManager) pushAccount(key *ecdsa.PrivateKey) types.Address {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	cfxAddr := CfxAddressOfPrivateKey(key, p.networkID)
	p.addresses = append(p.addresses, cfxAddr)
	p.accountsMap[cfxAddr.GetHexAddress()] = key
	return cfxAddr
}

func CfxAddressOfPrivateKey(key *ecdsa.PrivateKey, networkID uint32) types.Address {
	ethAddr := crypto.PubkeyToAddress(key.PublicKey)
	cfxAddr := addressutil.EtherAddressToCfxAddress(ethAddr, false, networkID)
	return cfxAddr
}
