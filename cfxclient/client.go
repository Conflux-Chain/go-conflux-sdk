// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package cfxclient

import (
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/interfaces"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/pkg/errors"
)

// Client represents a client to interact with Conflux blockchain.
type Client struct {
	*ClientCore
	cfx    *RpcCfxClient
	debug  *RpcDebugClient
	pubsub *RpcPubsubClient
	trace  *RpcTraceClient
	// pos    *RpcPosClient
}

// NewClient creates an instance of Client with specified conflux node url, it will creat account manager if option.KeystorePath not empty.
func NewClient(nodeURL string) (Client, error) {
	core, err := newClientCore(nodeURL)
	if err != nil {
		return Client{}, err
	}
	client := Client{}
	client.ClientCore = &core
	client.cfx = &RpcCfxClient{&core}
	client.debug = &RpcDebugClient{&core}
	client.pubsub = &RpcPubsubClient{&core}
	client.trace = &RpcTraceClient{&core}
	// client.pos = &RpcPosClient{&core}
	return client, nil
}

func (c *Client) Cfx() interfaces.RpcCfxCaller {
	return c.cfx
}
func (c *Client) Debug() interfaces.RpcDebugCaller {
	return c.debug
}
func (c *Client) Pos() interfaces.RpcPosCaller {
	// return c.pos
	return nil
}
func (c *Client) Pubsub() interfaces.RpcPubsubCaller {
	return c.pubsub
}
func (c *Client) Trace() interfaces.RpcTraceCaller {
	return c.trace
}

// =========Helper==========
// NewAddress create conflux address by base32 string or hex40 string, if base32OrHex is base32 and networkID is passed it will create cfx Address use networkID of current c.rpcCaller.
func (c *Client) NewAddress(base32OrHex string) (types.Address, error) {
	networkID, err := c.GetNetworkID()
	if err != nil {
		return types.Address{}, err
	}
	return cfxaddress.New(base32OrHex, networkID)
}

// MustNewAddress create conflux address by base32 string or hex40 string, if base32OrHex is base32 and networkID is passed it will create cfx Address use networkID of current c.rpcCaller.
// it will painc if error occured.
func (c *Client) MustNewAddress(base32OrHex string) types.Address {
	address, err := c.NewAddress(base32OrHex)
	if err != nil {
		panic(err)
	}
	return address
}

// GetNetworkID returns networkID of connecting conflux node
func (client *Client) GetNetworkID() (uint32, error) {
	return client.ClientCore.getNetworkId()
}

// === helper methods ===

// SignEncodedTransactionAndSend signs RLP encoded transaction "encodedTx" by signature "r,s,v" and sends it to node,
// and returns responsed transaction.
func (client *Client) EncodedTransactionAndSend(encodedTx []byte, v byte, r, s []byte) (*types.Transaction, error) {
	tx := new(types.UnsignedTransaction)
	netwrokID, err := client.GetNetworkID()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get networkID")
	}

	err = tx.Decode(encodedTx, netwrokID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode transaction")
	}
	// tx.From = from

	respondTx, err := client.encodeTransactionAndSend(tx, v, r, s)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to sign and send transaction %+v", tx)
	}

	return respondTx, nil
}

func (client *Client) encodeTransactionAndSend(tx *types.UnsignedTransaction, v byte, r, s []byte) (*types.Transaction, error) {
	rlp, err := tx.EncodeWithSignature(v, r, s)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode transaction with signature")
	}

	hash, err := client.Cfx().SendRawTransaction(rlp)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to send transaction, raw data = 0x%+x", rlp)
	}

	respondTx, err := client.Cfx().GetTransactionByHash(hash)
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
