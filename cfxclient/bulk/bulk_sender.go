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
	signalbeCaller sdk.ClientOperator
	unsignedTxs    []*types.UnsignedTransaction
}

// NewBuckSender creates new bulk sender instance
func NewBuckSender(signableClient sdk.Client) *BulkSender {
	return &BulkSender{
		signalbeCaller: &signableClient,
	}
}

// AppendTransaction append unsigned transaction to queue
func (b *BulkSender) AppendTransaction(tx types.UnsignedTransaction) *BulkSender {
	b.unsignedTxs = append(b.unsignedTxs, &tx)
	return b
}

// PopulateTransactions fill missing fields and nonce for unsigned transactions in queue
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
	userNextNonceCache, err := b.gatherInitNextNonces()
	if err != nil {
		return errors.WithStack(err)
	}

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
	return b.populateGasAndStorage()
}

func (b *BulkSender) populateGasAndStorage() error {
	estimatPtrs, errPtrs := make([]*types.Estimate, len(b.unsignedTxs)), make([]*error, len(b.unsignedTxs))
	bulkCaller := NewBulkCaller(b.signalbeCaller)
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
		return errors.WithStack(err)
	}

	for i, utx := range b.unsignedTxs {
		if utx.StorageLimit != nil && utx.Gas != nil {
			continue
		}

		if *errPtrs[i] != nil {
			return errors.WithMessagef(*errPtrs[i], "failed to estimate %vth transaction %+v", i, utx)
		}

		if utx.Gas == nil {
			utx.Gas = estimatPtrs[i].GasLimit
		}

		if utx.StorageLimit == nil {
			utx.StorageLimit = types.NewUint64(estimatPtrs[i].StorageCollateralized.ToInt().Uint64())
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

func (b *BulkSender) gatherInitNextNonces() (map[string]*big.Int, error) {
	result := make(map[string]*big.Int)

	bulkCaller := NewBulkCaller(b.signalbeCaller)
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

	for _, utx := range b.unsignedTxs {
		user := utx.From.String()
		if utx.Nonce != nil || result[user] != nil {
			continue
		}

		if *poolNextNonceErrs[user] == nil {
			result[utx.From.String()] = poolNextNonces[user].ToInt()
			continue
		}

		if *nextNonceErrs[user] == nil {
			result[utx.From.String()] = nextNonces[user].ToInt()
			continue
		}

		return nil, errors.WithStack(*nextNonceErrs[user])
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
	_client := b.signalbeCaller

	_defaultAccount, err := _client.GetAccountManager().GetDefault()
	if err != nil {
		return nil, nil, 0, nil, nil, errors.Wrap(err, "failed to get default account")
	}

	bulkCaller := NewBulkCaller(_client)
	_status, statusErr := bulkCaller.GetStatus()
	_gasPrice, gasPriceErr := bulkCaller.GetGasPrice()
	_epoch, epochErr := bulkCaller.GetEpochNumber(types.EpochLatestState)
	err = bulkCaller.Execute()

	if err != nil || *statusErr != nil || *gasPriceErr != nil || *epochErr != nil {
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
}

// SignAndSend signs and sends all unsigned transactions in queue by rpc call "batch" on one request
// and returns the result of sending transactions.
// If there is any error on rpc "batch", it will be returned with batchErr not nil.
// If there is no error on rpc "batch", it will return the txHashes or txErrors of sending transactions.
func (b *BulkSender) SignAndSend() (txHashes []*types.Hash, txErrors []error, batchErr error) {
	rawTxs := make([][]byte, len(b.unsignedTxs))

	for i, utx := range b.unsignedTxs {
		var err error
		rawTxs[i], err = b.signalbeCaller.GetAccountManager().SignTransaction(*utx)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "failed to encode the %vth transaction: %+v", i, utx)
		}
	}

	// send
	bulkCaller := NewBulkCaller(b.signalbeCaller)
	hashes := make([]*types.Hash, len(rawTxs))
	errs := make([]*error, len(rawTxs))
	for i, rawTx := range rawTxs {
		hashes[i], errs[i] = bulkCaller.Cfx().SendRawTransaction(rawTx)
	}

	batchErr = bulkCaller.Execute()
	if batchErr != nil {
		return nil, nil, errors.Wrapf(batchErr, "failed to batch send transactions")
	}

	errorVals := make([]error, len(errs))
	for i, err := range errs {
		errorVals[i] = *err
	}

	return hashes, errorVals, batchErr
}
