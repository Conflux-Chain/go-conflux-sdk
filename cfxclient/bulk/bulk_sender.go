package bulk

import (
	"math/big"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/constants"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
)

// BulkSender used for bulk send unsigned tranactions in one request to improve efficiency,
// it will auto populate missing fields and nonce of unsigned transactions in queue before send.
type BulkSender struct {
	signableCaller     sdk.ClientOperator
	unsignedTxs        []*types.UnsignedTransaction
	bulkEstimateErrors *ErrBulkEstimate
	isPopulated        bool
}

// NewBulkSender creates new bulk sender instance
func NewBulkSender(signableClient sdk.Client) *BulkSender {
	return &BulkSender{
		signableCaller: &signableClient,
	}
}

// AppendTransaction append unsigned transaction to queue
func (b *BulkSender) AppendTransaction(tx *types.UnsignedTransaction) *BulkSender {
	b.unsignedTxs = append(b.unsignedTxs, tx)
	return b
}

// PopulateTransactions fill missing fields and nonce for unsigned transactions in queue.
// nonceSouce means use pending nonce or nonce be the nonce of first tx not setted nonce.
// if set NONCE_TYPE_AUTO, it will use nonce when exist pending txs because of notEnoughCash/notEnoughCash/outDatedStatus/outOfEpochHeight/noncefuture
// and use pending nonce when no pending txs.
func (b *BulkSender) PopulateTransactions(nonceSource types.NonceType) ([]*types.UnsignedTransaction, error) {
	defaultAccount, chainID, networkId, gasPrice, epochHeight, err := b.getChainInfos()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if defaultAccount == nil {
		return nil, errors.Wrap(err, "failed to pupulate, no account found")
	}

	for _, utx := range b.unsignedTxs {
		if utx.From == nil {
			utx.From = defaultAccount
		}
	}

	estimateErrs, err := b.populateGasAndStorage()
	if err != nil {
		return nil, err
	}

	// set nonce
	userUsedNoncesMap := b.gatherUsedNonces()
	userNextNonceCache, err := b.gatherInitNextNonces(nonceSource)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for i, utx := range b.unsignedTxs {
		if estimateErrs != nil && (*estimateErrs)[i] != nil {
			continue
		}

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
			utx.Nonce = (*hexutil.Big)(userNextNonceCache[from])
			// avoid to reuse user used nonce, increase it if transactions used the nonce in cache
			for {
				userNextNonceCache[from] = big.NewInt(0).Add(userNextNonceCache[from], big.NewInt(1))
				if !b.checkIsNonceUsed(userUsedNoncesMap, utx.From, (*hexutil.Big)(userNextNonceCache[from])) {
					break
				}
			}
		}
	}

	// return results, estimatErrs
	b.isPopulated = true
	if estimateErrs != nil {
		b.bulkEstimateErrors = estimateErrs
		return b.unsignedTxs, b.bulkEstimateErrors
	}
	return b.unsignedTxs, nil
}

func (b *BulkSender) populateGasAndStorage() (*ErrBulkEstimate, error) {
	estimatPtrs, errPtrs := make([]*types.Estimate, len(b.unsignedTxs)), make([]*error, len(b.unsignedTxs))
	bulkCaller := NewBulkCaller(b.signableCaller)
	for i, utx := range b.unsignedTxs {
		if utx.StorageLimit != nil && utx.Gas != nil {
			continue
		}
		callReq := new(types.CallRequest)
		callReq.FillByUnsignedTx(utx)

		estimatPtrs[i], errPtrs[i] = bulkCaller.EstimateGasAndCollateral(*callReq)
	}

	err := bulkCaller.Execute()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	estimateErrors := ErrBulkEstimate{}
	for i, e := range errPtrs {
		// not estimate because of both StorageLimit and Gas have values
		if e == nil || *e == nil {
			continue
		}
		estimateErrors[i] = &ErrEstimate{*e}
	}

	for i, utx := range b.unsignedTxs {

		if _, ok := estimateErrors[i]; ok {
			continue
		}

		if utx.StorageLimit != nil && utx.Gas != nil {
			continue
		}

		if utx.Gas == nil {
			utx.Gas = estimatPtrs[i].GasLimit
		}

		// if utx.StorageLimit == nil {
		// 	utx.StorageLimit = types.NewUint64(estimatPtrs[i].StorageCollateralized.ToInt().Uint64())
		// }
	}

	if len(estimateErrors) > 0 {
		return &estimateErrors, nil
	}
	return nil, nil
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

func (b *BulkSender) gatherInitNextNonces(nonceSource types.NonceType) (map[string]*big.Int, error) {
	result := make(map[string]*big.Int)

	bulkCaller := NewBulkCaller(b.signableCaller)
	isUserCached := make(map[string]bool)
	poolNextNonces, poolNextNonceErrs := make(map[string]*hexutil.Big), make(map[string]*error)
	nextNonces, nextNonceErrs := make(map[string]*hexutil.Big), make(map[string]*error)

	for _, utx := range b.unsignedTxs {
		if isUserCached[utx.From.String()] {
			continue
		}
		poolNextNonces[utx.From.String()], poolNextNonceErrs[utx.From.String()] = bulkCaller.Txpool().NextNonce(*utx.From)
		nextNonces[utx.From.String()], nextNonceErrs[utx.From.String()] = bulkCaller.GetNextNonce(*utx.From)
	}

	err := bulkCaller.Execute()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	switch nonceSource {
	case types.NONCE_TYPE_PENDING_NONCE:
		for _, utx := range b.unsignedTxs {
			user := utx.From.String()
			if utx.Nonce != nil || result[user] != nil {
				continue
			}
			if *poolNextNonceErrs[user] != nil {
				return nil, errors.WithStack(*poolNextNonceErrs[user])
			}
			result[utx.From.String()] = poolNextNonces[user].ToInt()
		}
	case types.NONCE_TYPE_NONCE:
		for _, utx := range b.unsignedTxs {
			user := utx.From.String()
			if utx.Nonce != nil || result[user] != nil {
				continue
			}
			if *nextNonceErrs[user] != nil {
				return nil, errors.WithStack(*nextNonceErrs[user])
			}
			result[utx.From.String()] = nextNonces[user].ToInt()
		}
	case types.NONCE_TYPE_AUTO:
		pendingStatus, err := b.getSenderPendingStatus()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		for _, utx := range b.unsignedTxs {
			user := utx.From.String()
			if utx.Nonce != nil || result[user] != nil {
				continue
			}
			if pendingStatus[user] {
				if *nextNonceErrs[user] != nil {
					return nil, errors.WithStack(*nextNonceErrs[user])
				}
				result[utx.From.String()] = nextNonces[user].ToInt()
			} else {
				if *poolNextNonceErrs[user] != nil {
					return nil, errors.WithStack(*poolNextNonceErrs[user])
				}
				result[utx.From.String()] = poolNextNonces[user].ToInt()
			}
		}
	}
	return result, nil
}

func (b *BulkSender) getSenderPendingStatus() (map[string]bool, error) {
	type pendingTxRes struct {
		res *types.AccountPendingTransactions
		err *error
	}

	senderPendingTxRes := make(map[string]*pendingTxRes)
	for _, v := range b.unsignedTxs {
		senderPendingTxRes[v.From.String()] = &pendingTxRes{}
	}

	bulkCaller := NewBulkCaller(b.signableCaller)
	for user, pendingTxRes := range senderPendingTxRes {
		// logrus.WithField("user", user).Info("ready to check pending result")
		res, err := bulkCaller.GetAccountPendingTransactions(cfxaddress.MustNew(user), nil, nil)
		pendingTxRes.res = res
		pendingTxRes.err = err
	}

	// err means timeout
	if err := bulkCaller.Execute(); err != nil {
		return nil, err
	}

	result := make(map[string]bool)
	for user, v := range senderPendingTxRes {
		if v.res.PendingCount > 0 && v.res.FirstTxStatus != nil {
			isPending, _ := v.res.FirstTxStatus.IsPending()
			result[user] = isPending
		}
	}
	return result, nil
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
	_client := b.signableCaller

	_defaultAccount, err := _client.GetAccountManager().GetDefault()
	if err != nil {
		return nil, nil, 0, nil, nil, errors.Wrap(err, "failed to get default account")
	}

	bulkCaller := NewBulkCaller(_client)
	_status, statusErr := bulkCaller.GetStatus()
	_gasPrice, gasPriceErr := bulkCaller.GetGasPrice()
	_epoch, epochErr := bulkCaller.GetEpochNumber(types.EpochLatestState)

	err = bulkCaller.Execute()
	if *statusErr != nil {
		return nil, nil, 0, nil, nil, errors.Wrap(*statusErr, "failed to bulk fetch chain infos")
	}
	if *gasPriceErr != nil {
		return nil, nil, 0, nil, nil, errors.Wrap(*gasPriceErr, "failed to bulk fetch chain infos")
	}
	if *epochErr != nil {
		return nil, nil, 0, nil, nil, errors.Wrap(*epochErr, "failed to bulk fetch chain infos")
	}
	if err != nil {
		return nil, nil, 0, nil, nil, errors.Wrap(err, "failed to bulk fetch chain infos")
	}

	_chainID, _networkId := &_status.ChainID, uint32(_status.NetworkID)
	_epochHeight := types.NewUint64(_epoch.ToInt().Uint64())

	// conflux responsed gasprice offen be 0, but the min gasprice is 1 when sending transaction, so do this
	if _gasPrice.ToInt().Cmp(big.NewInt(constants.MinGasprice)) < 1 {
		_gasPrice = types.NewBigInt(constants.MinGasprice)
	}

	chainIDInUint := (hexutil.Uint)(*_chainID)
	return _defaultAccount, &chainIDInUint, _networkId, _gasPrice, _epochHeight, nil
}

// Clear clear batch elems and errors in queue for new bulk call action
func (b *BulkSender) Clear() {
	b.unsignedTxs = b.unsignedTxs[:0]
	b.isPopulated = false
}

func (b *BulkSender) IsPopulated() bool {
	return b.isPopulated
}

// SignAndSend signs and sends all unsigned transactions in queue by rpc call "batch" on one request
// and returns the result of sending transactions.
// If there is any error on rpc "batch", it will be returned with err not nil.
// If there is no error on rpc "batch", it will return the txHashes or txErrors of sending transactions.
func (b *BulkSender) SignAndSend() (txHashes []*types.Hash, txErrors []error, err error) {
	if !b.IsPopulated() {
		_, err := b.PopulateTransactions(types.NONCE_TYPE_AUTO)
		if err != nil {
			return nil, nil, err
		}
	}

	rawTxs := make([][]byte, len(b.unsignedTxs))

	for i, utx := range b.unsignedTxs {
		var err error
		rawTxs[i], err = b.signableCaller.GetAccountManager().SignTransaction(*utx)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "failed to encode the %vth transaction: %+v", i, utx)
		}
	}

	// send
	bulkCaller := NewBulkCaller(b.signableCaller)
	hashes := make([]*types.Hash, len(rawTxs))
	txErrs := make([]*error, len(rawTxs))
	for i, rawTx := range rawTxs {
		hashes[i], txErrs[i] = bulkCaller.Cfx().SendRawTransaction(rawTx)
	}

	err = bulkCaller.Execute()
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to batch send transactions")
	}

	errorVals := make([]error, len(txErrs))
	for i, err := range txErrs {
		errorVals[i] = *err
	}

	return hashes, errorVals, err
}
