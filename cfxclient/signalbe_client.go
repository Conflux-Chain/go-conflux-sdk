package cfxclient

import (
	"math/big"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/accounts"
	"github.com/Conflux-Chain/go-conflux-sdk/constants"
	"github.com/Conflux-Chain/go-conflux-sdk/interfaces"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	cfxerrors "github.com/Conflux-Chain/go-conflux-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
)

type SignableClient struct {
	*Client
	wallet interfaces.Wallet
}

func NewSignableClient(c *Client, wallet interfaces.Wallet) SignableClient {
	return SignableClient{c, wallet}
}

//FixMe: Is there better name
func NewSignalbeClientByPath(nodeURL string, keyStorePath string) (SignableClient, error) {

	_client, err := NewClient(nodeURL)
	if err != nil {
		return SignableClient{}, err
	}

	networkId, err := _client.GetNetworkID()
	if err != nil {
		return SignableClient{}, err
	}

	wallet := accounts.NewKeystoreWallet(keyStorePath, uint32(networkId))

	return NewSignableClient(&_client, wallet), nil
}

func (c *SignableClient) GetWallet() interfaces.Wallet {
	return c.wallet
}

func (c *SignableClient) SetRetry(retryCount int, retryInterval time.Duration) *SignableClient {
	c.Client.SetRetry(retryCount, retryInterval)
	return c
}

func (c *SignableClient) SetRequestTimeout(timeout time.Duration) *SignableClient {
	c.Client.SetRequestTimeout(timeout)
	return c
}

// NewTransaction creates an unsigned transaction by parameters,
// and the other fields will be set to values fetched from conflux node.
func (c *SignableClient) NewTransaction(from types.Address, to types.Address, amount *hexutil.Big, data []byte) (types.UnsignedTransaction, error) {
	tx := new(types.UnsignedTransaction)
	tx.From = &from
	tx.To = &to
	tx.Value = amount
	tx.Data = data

	err := c.PopulateTransaction(tx)
	if err != nil {
		return types.UnsignedTransaction{}, errors.Wrap(err, cfxerrors.ErrMsgApplyTxValues)
	}

	return *tx, nil
}

// PopulateTransaction set empty fields to value fetched from conflux node.
func (c *SignableClient) PopulateTransaction(tx *types.UnsignedTransaction) error {

	networkId, err := c.GetNetworkID()
	if err != nil {
		return err
	}

	if c != nil {
		if tx.From == nil {
			// if c.AccountManager != nil {
			defaultAccount, err := c.GetWallet().GetDefault()
			if err != nil {
				return errors.Wrap(err, "failed to get default account")
			}

			if defaultAccount == nil {
				return errors.New("no account found")
			}
			tx.From = defaultAccount
			// }
		}
		tx.From.CompleteByNetworkID(networkId)
		tx.To.CompleteByNetworkID(networkId)

		if tx.Nonce == nil {
			nonce, err := c.Cfx().GetNextNonce(*tx.From, nil)
			if err != nil {
				return errors.Wrap(err, "failed to get nonce")
			}
			tmp := hexutil.Big(*nonce)
			tx.Nonce = &tmp
		}

		if tx.ChainID == nil {
			status, err := c.Cfx().GetStatus()
			if err != nil {
				tx.ChainID = types.NewUint(0)
			} else {
				tx.ChainID = &status.ChainID
			}
		}

		if tx.GasPrice == nil {
			gasPrice, err := c.Cfx().GetGasPrice()
			if err != nil {
				return errors.Wrap(err, "failed to get gas price")
			}

			// conflux responsed gasprice offen be 0, but the min gasprice is 1 when sending transaction, so do this
			if gasPrice.ToInt().Cmp(big.NewInt(constants.MinGasprice)) < 1 {
				gasPrice = types.NewBigInt(constants.MinGasprice)
			}
			tmp := hexutil.Big(*gasPrice)
			tx.GasPrice = &tmp
		}

		if tx.EpochHeight == nil {
			epoch, err := c.Cfx().GetEpochNumber(types.EpochLatestState)
			if err != nil {
				return errors.Wrap(err, "failed to get the latest state epoch number")
			}
			// tx.EpochHeight = (*hexutil.Big)(epoch).toi
			tx.EpochHeight = types.NewUint64(epoch.ToInt().Uint64())
		}

		// The gas and storage limit may be influnced by all fileds of transaction ,so set them at last step.
		if tx.StorageLimit == nil || tx.Gas == nil {
			callReq := new(types.CallRequest)
			callReq.FillByUnsignedTx(tx)

			sm, err := c.Cfx().EstimateGasAndCollateral(*callReq)
			if err != nil {
				return errors.Wrapf(err, "failed to estimate gas and collateral, request = %+v", *callReq)
			}

			// fmt.Printf("callreq, %+v,sm:%+v\n", *callReq, sm)

			if tx.Gas == nil {
				tx.Gas = sm.GasLimit
			}

			if tx.StorageLimit == nil {
				tx.StorageLimit = types.NewUint64(sm.StorageCollateralized.ToInt().Uint64() * 10 / 9)
			}
		}

		tx.ApplyDefault()
	}

	return nil
}

// SendTransaction signs and sends transaction to conflux node and returns the transaction hash.
func (c *SignableClient) SignTransactionAndSend(tx types.UnsignedTransaction) (types.Hash, error) {

	err := c.PopulateTransaction(&tx)
	if err != nil {
		return "", errors.Wrap(err, cfxerrors.ErrMsgApplyTxValues)
	}

	rawData, err := c.GetWallet().SignTransactionAndEncode(tx)
	if err != nil {
		return "", errors.Wrap(err, "failed to sign and encode transaction")
	}

	//send raw tx
	txhash, err := c.Cfx().SendRawTransaction(rawData)
	if err != nil {
		return "", errors.Wrapf(err, "failed to send transaction, raw data = 0x%+x", rawData)
	}
	return txhash, nil
}
