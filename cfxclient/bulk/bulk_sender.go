package bulk

import (
	"fmt"
	"math/big"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/constants"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
)

type BulkSender struct {
	// bulkCaller     *BulkCaller
	signalbeCaller sdk.ClientOperator
	unsignedTxs    []*types.UnsignedTransaction
	// wallet         interfaces.Wallet
}

func NewBuckSender(signableClient sdk.Client) *BulkSender {
	return &BulkSender{
		// bulkCaller:     bulkCaller,
		signalbeCaller: &signableClient,
	}
}

func (b *BulkSender) AppendTransaction(tx types.UnsignedTransaction) *BulkSender {
	b.unsignedTxs = append(b.unsignedTxs, &tx)
	return b
}

func (b *BulkSender) PopulateTransactions() error {
	defaultAccount, chainID, networkId, gasPrice, epochHeight, err := b.getChainInfos()
	if err != nil {
		return errors.WithStack(err)
	}
	if defaultAccount == nil {
		return errors.Wrap(err, "failed to pupulate, no account found")
	}

	for _, utx := range b.unsignedTxs {
		if utx.From == nil {
			utx.From = defaultAccount
		}
	}

	userUsedNoncesMap := b.gatherUsedNonces()
	// fill nonce
	userNextNonceCache := make(map[string]*big.Int, len(b.unsignedTxs))
	for _, utx := range b.unsignedTxs {
		utx.From.CompleteByNetworkID(networkId)
		utx.To.CompleteByNetworkID(networkId)

		if utx.ChainID == nil {
			utx.ChainID = chainID
		}

		if utx.GasPrice == nil {
			utx.GasPrice = gasPrice
		}

		if utx.EpochHeight == nil {
			utx.EpochHeight = epochHeight
		}

		if utx.Value == nil {
			utx.Value = types.NewBigInt(0)
		}

		if utx.Nonce == nil {
			from := utx.From.String()
			if userNextNonceCache[from] == nil {
				// TODO: get pending nonce
				hexNonce, err := b.signalbeCaller.GetNextNonce(*utx.From)
				if err != nil {
					return errors.WithStack(err)
				}
				userNextNonceCache[from] = hexNonce.ToInt()
			}

			utx.Nonce = (*hexutil.Big)(userNextNonceCache[from])
			// avoid to reuse user used nonce, increase it if transactions used the nonce in cache
			for {
				fmt.Printf("check nonce use")
				userNextNonceCache[from] = big.NewInt(0).Add(userNextNonceCache[from], big.NewInt(1))
				if !b.checkIsNonceUsed(userUsedNoncesMap, utx.From, (*hexutil.Big)(userNextNonceCache[from])) {
					break
				}
			}

			fmt.Printf("utx.Nonce %v\n", utx.Nonce)
		}
	}

	for _, utx := range b.unsignedTxs {
		// The gas and storage limit may be influnced by all fileds of transaction ,so set them at last step.
		if utx.StorageLimit == nil || utx.Gas == nil {
			callReq := new(types.CallRequest)
			callReq.FillByUnsignedTx(utx)

			estimat, err := b.signalbeCaller.EstimateGasAndCollateral(*callReq)
			if err != nil {
				return errors.Wrapf(err, "failed to estimate gas and collateral, request = %+v", *callReq)
			}

			if utx.Gas == nil {
				utx.Gas = estimat.GasLimit
			}

			if utx.StorageLimit == nil {
				utx.StorageLimit = types.NewUint64(estimat.StorageCollateralized.ToInt().Uint64())
			}
		}
	}
	return nil
}

func (b *BulkSender) gatherUsedNonces() map[string]map[string]bool {
	result := make(map[string]map[string]bool)
	for _, utx := range b.unsignedTxs {
		if utx.Nonce != nil && utx.From != nil {
			from, nonce := utx.From.String(), utx.Nonce.String()
			if result[from] == nil {
				result[from] = make(map[string]bool)
			}
			result[from][nonce] = true
		}
	}
	return result
}

func (b *BulkSender) checkIsNonceUsed(usedCaches map[string]map[string]bool, user *cfxaddress.Address, nonce *hexutil.Big) bool {
	hasCache, ok := usedCaches[user.String()]
	if ok {
		return hasCache[nonce.String()]
	}
	return false
}

func (b *BulkSender) getChainInfos() (
	defaultAccount *cfxaddress.Address,
	chainID *hexutil.Uint,
	networkId uint32,
	gasPrice *hexutil.Big,
	epochHeight *hexutil.Uint64,
	err error,
) {
	_client := b.signalbeCaller
	defaultAccount, err = _client.GetAccountManager().GetDefault()
	if err != nil {
		return nil, nil, 0, nil, nil, errors.Wrap(err, "failed to get default account")
	}

	status, err := _client.GetStatus()
	if err != nil {
		return nil, nil, 0, nil, nil, errors.WithStack(err)
	}
	chainID = &status.ChainID

	networkId, err = _client.GetNetworkID()
	if err != nil {
		return nil, nil, 0, nil, nil, errors.WithStack(err)
	}

	gasPrice, err = _client.GetGasPrice()
	if err != nil {
		return nil, nil, 0, nil, nil, errors.Wrap(err, "failed to get gas price")
	}

	// conflux responsed gasprice offen be 0, but the min gasprice is 1 when sending transaction, so do this
	if gasPrice.ToInt().Cmp(big.NewInt(constants.MinGasprice)) < 1 {
		gasPrice = types.NewBigInt(constants.MinGasprice)
	}

	epoch, err := _client.GetEpochNumber(types.EpochLatestState)
	if err != nil {
		return nil, nil, 0, nil, nil, errors.Wrap(err, "failed to get the latest state epoch number")
	}
	epochHeight = types.NewUint64(epoch.ToInt().Uint64())

	return defaultAccount, chainID, networkId, gasPrice, epochHeight, nil
}

func (b *BulkSender) SignAndSend() ([]*types.Hash, []error, error) {
	rawTxs := make([][]byte, len(b.unsignedTxs))

	for i, utx := range b.unsignedTxs {
		var err error
		rawTxs[i], err = b.signalbeCaller.GetAccountManager().SignTransaction(*utx)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "failed to encode the %vth transaction: %+v", i, utx)
		}
	}

	// send
	bulkCaller := NewBulkerCaller(b.signalbeCaller)
	hashes := make([]*types.Hash, len(rawTxs))
	for i, rawTx := range rawTxs {
		fmt.Printf("raw Tx: %x\n", rawTx)
		hashes[i] = bulkCaller.Cfx().SendRawTransaction(rawTx)
	}

	_errors, err := bulkCaller.Excute()
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to batch send transactions")
	}
	return hashes, _errors, err
}

// // ? nonce 是否可以由用户设置
// // 每个用户数量不能大于2000
// func (b *BulkSender) Execute() ([]types.Hash, []error, error) {
// 	// get pending nonce
// 	// check from, to, gasprice>1, epochheight<current+100000, chainID match, gas>=21000
// 	// estimate all
// 	// check balance
// 	// 		1. gas = sponsored  ? 0 : gas; storage = sponsored ? 0 : sotrage/1024*1e18
// 	// 		2. balance > value + gas * gasprice + storage
// 	// sign txs
// 	signedTxs := make([]types.SignedTransaction, len(b.unsignedTxs))
// 	for i, v := range b.unsignedTxs {
// 		signedTxs[i] = b.signalbeCaller.GetWallet().SignTransaction()

// 	}
// 	// batch call

// }
