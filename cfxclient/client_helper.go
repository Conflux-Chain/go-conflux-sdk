package cfxclient

import (
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/pkg/errors"
)

type ClientHelper struct {
	client *Client
}

func NewClientHelper(client *Client) ClientHelper {
	return ClientHelper{client}
}

// =========Helper==========
// NewAddress create conflux address by base32 string or hex40 string, if base32OrHex is base32 and networkID is passed it will create cfx Address use networkID of current c.rpcCaller.
func (c *ClientHelper) NewAddress(base32OrHex string) (types.Address, error) {
	networkID, err := c.client.GetNetworkID()
	if err != nil {
		return types.Address{}, err
	}
	return cfxaddress.New(base32OrHex, networkID)
}

// MustNewAddress create conflux address by base32 string or hex40 string, if base32OrHex is base32 and networkID is passed it will create cfx Address use networkID of current c.rpcCaller.
// it will painc if error occured.
func (c *ClientHelper) MustNewAddress(base32OrHex string) types.Address {
	address, err := c.NewAddress(base32OrHex)
	if err != nil {
		panic(err)
	}
	return address
}

// === helper methods ===

// SignEncodedTransactionAndSend signs RLP encoded transaction "encodedTx" by signature "r,s,v" and sends it to node,
// and returns responsed transaction.
func (c *ClientHelper) EncodedTransactionAndSend(encodedTx []byte, v byte, r, s []byte) (*types.Transaction, error) {
	tx := new(types.UnsignedTransaction)
	netwrokID, err := c.client.GetNetworkID()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get networkID")
	}

	err = tx.Decode(encodedTx, netwrokID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode transaction")
	}
	// tx.From = from

	respondTx, err := c.encodeTransactionAndSend(tx, v, r, s)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to sign and send transaction %+v", tx)
	}

	return respondTx, nil
}

func (c *ClientHelper) encodeTransactionAndSend(tx *types.UnsignedTransaction, v byte, r, s []byte) (*types.Transaction, error) {
	rlp, err := tx.EncodeWithSignature(v, r, s)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode transaction with signature")
	}

	hash, err := c.client.Cfx().SendRawTransaction(rlp)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to send transaction, raw data = 0x%+x", rlp)
	}

	respondTx, err := c.client.Cfx().GetTransactionByHash(hash)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get transaction by hash %v", hash)
	}
	return respondTx, nil
}

// WaitForTransationBePacked returns transaction when it is packed
func (c *Client) WaitForTransationBePacked(txhash types.Hash, duration time.Duration) (*types.Transaction, error) {
	// fmt.Printf("wait for transaction %v be packed\n", txhash)
	if duration == 0 {
		duration = time.Second
	}

	var tx *types.Transaction
	for {
		time.Sleep(duration)
		var err error
		tx, err = c.Cfx().GetTransactionByHash(txhash)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get transaction by hash %v", txhash)
		}

		if tx.Status != nil {
			// fmt.Printf("transaction is packed:%+v\n\n", JsonFmt(tx))
			break
		}
	}
	return tx, nil
}

// WaitForTransationReceipt waits for transaction receipt valid
func (c *Client) WaitForTransationReceipt(txhash types.Hash, duration time.Duration) (*types.TransactionReceipt, error) {
	// fmt.Printf("wait for transaction %v be packed\n", txhash)
	if duration == 0 {
		duration = time.Second
	}

	var txReceipt *types.TransactionReceipt
	for {
		time.Sleep(duration)
		var err error
		txReceipt, err = c.Cfx().GetTransactionReceipt(txhash)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get transaction receipt")
		}

		if txReceipt != nil {
			break
		}
	}
	return txReceipt, nil
}
