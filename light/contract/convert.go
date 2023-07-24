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
			Account:               v,
			CompressedPublicKey:   info.PublicKey,
			UncompressedPublicKey: ledger.NextEpochValidators[v],
			VotingPower:           uint64(info.VotingPower),
		}

		if len(validator.UncompressedPublicKey) == 0 {
			return LedgerInfoLibEpochState{}, false
		}

		if info.VrfPublicKey != nil {
			validator.VrfPublicKey = *info.VrfPublicKey
		}

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

	for _, v := range ledger.ValidatorsSorted() {
		result.Signatures = append(result.Signatures, LedgerInfoLibAccountSignature{
			Account:            v,
			ConsensusSignature: ledger.Signatures[v],
		})
	}

	return result
}
