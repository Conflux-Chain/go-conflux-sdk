package signutil

import (
	"crypto/ecdsa"
	"errors"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/utils/addressutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/openweb3/go-sdk-common/privatekeyhelper"
)

func SignTxByPrivateKey(keyString string, tx types.UnsignedTransaction) (*types.SignedTransaction, error) {

	key, err := privatekeyhelper.NewFromKeyString(keyString)
	if err != nil {
		return nil, err
	}

	if tx.From != nil {
		addr := cfxHexAddressByPrivateKey(key)
		if tx.From.GetHexAddress() != addr {
			return nil, errors.New("from of tx conflict with address of private key")
		}
	}

	hash, err := tx.Hash()
	if err != nil {
		return nil, err
	}

	sig, err := crypto.Sign(hash, key)
	if err != nil {
		return nil, err
	}

	v, r, s := sig[64], sig[0:32], sig[32:64]

	signdTx := &types.SignedTransaction{}
	signdTx.UnsignedTransaction = tx
	signdTx.V = v
	signdTx.R = r
	signdTx.S = s

	return signdTx, nil
}

func cfxHexAddressByPrivateKey(key *ecdsa.PrivateKey) string {
	ethAddr := crypto.PubkeyToAddress(key.PublicKey)
	cfxAddr := addressutil.EtherAddressToCfxAddress(ethAddr, false, 1)
	return cfxAddr.GetHexAddress()
}
