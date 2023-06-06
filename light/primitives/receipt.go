package primitives

import (
	"math/big"

	"github.com/Conflux-Chain/go-conflux-sdk/light/contract"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Receipt struct {
	AccumulatedGasUsed    *big.Int
	GasFee                *big.Int
	GasSponsorPaid        Bool
	LogBloom              []byte
	Logs                  []contract.TypesTxLog
	OutcomeStatus         uint8
	StorageSponsorPaid    Bool
	StorageCollateralized []contract.TypesStorageChange
	StorageReleased       []contract.TypesStorageChange
}

func ConvertReceipt(receipt *types.TransactionReceipt) Receipt {
	storageCollateralized, storageReleased := contract.ConstructStorageChanges(receipt)

	return Receipt{
		AccumulatedGasUsed:    receipt.AccumulatedGasUsed.ToInt(),
		GasFee:                receipt.GasFee.ToInt(),
		GasSponsorPaid:        Bool(receipt.GasCoveredBySponsor),
		LogBloom:              hexutil.MustDecode(string(receipt.LogsBloom)),
		Logs:                  contract.ConvertLogs(receipt.Logs),
		OutcomeStatus:         uint8(receipt.MustGetOutcomeType()),
		StorageSponsorPaid:    Bool(receipt.StorageCoveredBySponsor),
		StorageCollateralized: storageCollateralized,
		StorageReleased:       storageReleased,
	}
}
