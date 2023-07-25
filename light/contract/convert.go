package contract

import (
	postypes "github.com/Conflux-Chain/go-conflux-sdk/types/pos"
	"github.com/ethereum/go-ethereum/common"
)

func ConvertCommittee(ledger *postypes.LedgerInfoWithSignatures) (LedgerInfoLibEpochState, bool) {
	if ledger == nil {
		return LedgerInfoLibEpochState{}, false
	}

	state := ledger.LedgerInfo.CommitInfo.NextEpochState
	if state == nil {
		return LedgerInfoLibEpochState{}, false
	}

	var validators []LedgerInfoLibValidatorInfo
	for _, v := range ledger.NextEpochValidatorsSorted() {
		info := state.Verifier.AddressToValidatorInfo[v]

		validator := LedgerInfoLibValidatorInfo{
			Account:             v,
			CompressedPublicKey: info.PublicKey,
			VotingPower:         uint64(info.VotingPower),
		}

		if info.VrfPublicKey != nil {
			validator.VrfPublicKey = *info.VrfPublicKey
		}

		uncompressedPubKey, ok := ledger.NextEpochValidators[v]
		if !ok {
			return LedgerInfoLibEpochState{}, false
		}

		validator.UncompressedPublicKey = make([]byte, 128)
		copy(validator.UncompressedPublicKey[16:64], uncompressedPubKey[:48])
		copy(validator.UncompressedPublicKey[80:128], uncompressedPubKey[48:])

		validators = append(validators, validator)
	}

	return LedgerInfoLibEpochState{
		Epoch:             uint64(state.Epoch),
		Validators:        validators,
		QuorumVotingPower: uint64(state.Verifier.QuorumVotingPower),
		TotalVotingPower:  uint64(state.Verifier.TotalVotingPower),
		VrfSeed:           state.VrfSeed,
	}, true
}

func ConvertLedger(ledger *postypes.LedgerInfoWithSignatures) LedgerInfoLibLedgerInfoWithSignatures {
	committee, _ := ConvertCommittee(ledger)

	result := LedgerInfoLibLedgerInfoWithSignatures{
		Epoch:             uint64(ledger.LedgerInfo.CommitInfo.Epoch),
		Round:             uint64(ledger.LedgerInfo.CommitInfo.Round),
		Id:                common.BytesToHash(ledger.LedgerInfo.CommitInfo.Id),
		ExecutedStateId:   common.BytesToHash(ledger.LedgerInfo.CommitInfo.ExecutedStateId),
		Version:           uint64(ledger.LedgerInfo.CommitInfo.Version),
		TimestampUsecs:    uint64(ledger.LedgerInfo.CommitInfo.TimestampUsecs),
		NextEpochState:    committee,
		ConsensusDataHash: common.BytesToHash(ledger.LedgerInfo.ConsensusDataHash),
	}

	if pivot := ledger.LedgerInfo.CommitInfo.Pivot; pivot != nil {
		result.Pivot.Height = uint64(pivot.Height)
		result.Pivot.BlockHash = pivot.BlockHash.ToHash()
	}

	result.AggregatedSignature = make([]byte, 256)
	copy(result.AggregatedSignature[16:64], ledger.AggregatedSignature[:48])
	copy(result.AggregatedSignature[80:128], ledger.AggregatedSignature[48:96])
	copy(result.AggregatedSignature[144:192], ledger.AggregatedSignature[96:144])
	copy(result.AggregatedSignature[208:256], ledger.AggregatedSignature[144:192])
	for _, v := range ledger.ValidatorsSorted() {
		result.Accounts = append(result.Accounts, v)
	}

	return result
}
