package interfaces

import (
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
)

// Wallet is interface of operate actions on account manager
type Wallet interface {
	Create(passphrase string) (types.Address, error)
	Import(keyFile, passphrase, newPassphrase string) (types.Address, error)
	Delete(address types.Address, passphrase string) error
	Update(address types.Address, passphrase, newPassphrase string) error
	List() []types.Address
	GetDefault() (*types.Address, error)
	Unlock(address types.Address, passphrase string) error
	UnlockDefault(passphrase string) error
	TimedUnlock(address types.Address, passphrase string, timeout time.Duration) error
	TimedUnlockDefault(passphrase string, timeout time.Duration) error
	Lock(address types.Address) error

	SignTransaction(tx types.UnsignedTransaction) (types.SignedTransaction, error)
	SignTransactionAndEncode(tx types.UnsignedTransaction) ([]byte, error)
	SignTransactionWithPassphrase(tx types.UnsignedTransaction, passphrase string) (types.SignedTransaction, error)
	SignTransactionWithPassphraseAndEcode(tx types.UnsignedTransaction, passphrase string) ([]byte, error)
	CalcSignature(tx types.UnsignedTransaction, passphrase string) (v byte, r, s []byte, err error)
}
