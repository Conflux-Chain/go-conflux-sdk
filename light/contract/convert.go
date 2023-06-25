package contract

import (
	"bytes"
	"math/big"
	"sort"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	postypes "github.com/Conflux-Chain/go-conflux-sdk/types/pos"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func ConvertLogs(logs []types.Log) []TypesTxLog {
	var result []TypesTxLog

	for _, v := range logs {
		var topics [][32]byte
		for _, t := range v.Topics {
			topics = append(topics, *t.ToCommonHash())
		}

		var space uint8
		switch *v.Space {
		case types.SPACE_NATIVE:
			space = LogSpaceNative
		case types.SPACE_EVM:
			space = LogSpaceEthereum
		default:
			panic("invalid space in log entry")
		}

		result = append(result, TypesTxLog{
			Addr:   v.Address.MustGetCommonAddress(),
			Topics: topics,
			Data:   v.Data,
			Space:  space,
		})
	}

	return result
}

func ConstructStorageChanges(receipt *types.TransactionReceipt) (collateralized, released []TypesStorageChange) {
	for _, v := range receipt.StorageReleased {
		released = append(released, TypesStorageChange{
			Account:     v.Address.MustGetCommonAddress(),
			Collaterals: uint64(v.Collaterals),
		})
	}

	if receipt.StorageCollateralized == 0 {
		return
	}

	var account cfxaddress.Address
	if receipt.StorageCoveredBySponsor {
		account = *receipt.To
	} else {
		account = receipt.From
	}

	collateralized = append(collateralized, TypesStorageChange{
		Account:     account.MustGetCommonAddress(),
		Collaterals: uint64(receipt.StorageCollateralized),
	})

	return
}

func ConvertReceipt(receipt *types.TransactionReceipt) TypesTxReceipt {
	storageCollateralized, storageReleased := ConstructStorageChanges(receipt)

	return TypesTxReceipt{
		AccumulatedGasUsed:    receipt.AccumulatedGasUsed.ToInt(),
		GasFee:                receipt.GasFee.ToInt(),
		GasSponsorPaid:        receipt.GasCoveredBySponsor,
		LogBloom:              hexutil.MustDecode(string(receipt.LogsBloom)),
		Logs:                  ConvertLogs(receipt.Logs),
		OutcomeStatus:         uint8(receipt.MustGetOutcomeType()),
		StorageSponsorPaid:    receipt.StorageCoveredBySponsor,
		StorageCollateralized: storageCollateralized,
		StorageReleased:       storageReleased,
	}
}

func ConvertLedger(ledger *postypes.LedgerInfoWithSignatures) TypesLedgerInfoWithSignatures {
	result := TypesLedgerInfoWithSignatures{
		Epoch:             uint64(ledger.LedgerInfo.CommitInfo.Epoch),
		Round:             uint64(ledger.LedgerInfo.CommitInfo.Round),
		Id:                common.BytesToHash(ledger.LedgerInfo.CommitInfo.Id),
		ExecutedStateId:   common.BytesToHash(ledger.LedgerInfo.CommitInfo.ExecutedStateId),
		Version:           uint64(ledger.LedgerInfo.CommitInfo.Version),
		TimestampUsecs:    uint64(ledger.LedgerInfo.CommitInfo.TimestampUsecs),
		ConsensusDataHash: common.BytesToHash(ledger.LedgerInfo.ConsensusDataHash),
	}

	if ledger.LedgerInfo.CommitInfo.NextEpochState != nil {
		result.NextEpochState.Epoch = uint64(ledger.LedgerInfo.CommitInfo.NextEpochState.Epoch)
		var validators sortableValidators
		for k, v := range ledger.LedgerInfo.CommitInfo.NextEpochState.Verifier.AddressToValidatorInfo {
			validator := TypesValidatorInfo{
				Account:     k,
				PublicKey:   v.PublicKey,
				VotingPower: uint64(v.VotingPower),
			}

			if v.VrfPublicKey != nil {
				validator.VrfPublicKey = *v.VrfPublicKey
			}

			validators = append(validators, validator)
		}
		sort.Sort(validators)
		result.NextEpochState.Validators = validators
		result.NextEpochState.QuorumVotingPower = uint64(ledger.LedgerInfo.CommitInfo.NextEpochState.Verifier.QuorumVotingPower)
		result.NextEpochState.TotalVotingPower = uint64(ledger.LedgerInfo.CommitInfo.NextEpochState.Verifier.TotalVotingPower)
		result.NextEpochState.VrfSeed = ledger.LedgerInfo.CommitInfo.NextEpochState.VrfSeed
	}

	if ledger.LedgerInfo.CommitInfo.Pivot != nil {
		result.Pivot.Height = uint64(ledger.LedgerInfo.CommitInfo.Pivot.Height)
		result.Pivot.BlockHash = ledger.LedgerInfo.CommitInfo.Pivot.BlockHash.ToHash()
	}

	var signatures sortableAccountSignatures
	for k, v := range ledger.Signatures {
		signatures = append(signatures, TypesAccountSignature{
			Account:            k,
			ConsensusSignature: v,
		})
	}
	sort.Sort(signatures)
	result.Signatures = signatures

	return result
}

func ConvertBlockHeader(block *types.BlockSummary) TypesBlockHeader {
	var referees [][32]byte
	for _, v := range block.RefereeHashes {
		referees = append(referees, *v.ToCommonHash())
	}

	var custom [][]byte
	for _, v := range block.Custom {
		custom = append(custom, v)
	}

	var posRef [32]byte
	if block.PosReference != nil {
		posRef = *block.PosReference.ToCommonHash()
	}

	return TypesBlockHeader{
		ParentHash:            *block.ParentHash.ToCommonHash(),
		Height:                block.Height.ToInt(),
		Timestamp:             block.Timestamp.ToInt(),
		Author:                block.Miner.MustGetCommonAddress(),
		TransactionsRoot:      *block.TransactionsRoot.ToCommonHash(),
		DeferredStateRoot:     *block.DeferredStateRoot.ToCommonHash(),
		DeferredReceiptsRoot:  *block.DeferredReceiptsRoot.ToCommonHash(),
		DeferredLogsBloomHash: *block.DeferredLogsBloomHash.ToCommonHash(),
		Blame:                 big.NewInt(int64(block.Blame)),
		Difficulty:            block.Difficulty.ToInt(),
		Adaptive:              block.Adaptive,
		GasLimit:              block.GasLimit.ToInt(),
		RefereeHashes:         referees,
		Custom:                custom,
		Nonce:                 block.Nonce.ToInt(),
		PosReference:          posRef,
	}
}

type sortableAccountSignatures []TypesAccountSignature

func (s sortableAccountSignatures) Len() int { return len(s) }
func (s sortableAccountSignatures) Less(i, j int) bool {
	return bytes.Compare(s[i].Account[:], s[j].Account[:]) < 0
}
func (s sortableAccountSignatures) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type sortableValidators []TypesValidatorInfo

func (s sortableValidators) Len() int { return len(s) }
func (s sortableValidators) Less(i, j int) bool {
	return bytes.Compare(s[i].Account[:], s[j].Account[:]) < 0
}
func (s sortableValidators) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
