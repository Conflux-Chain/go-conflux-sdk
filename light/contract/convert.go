package contract

import (
	"bytes"
	"sort"

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

	var validators sortableValidators
	for k, v := range state.Verifier.AddressToValidatorInfo {
		validator := LedgerInfoLibValidatorInfo{
			Account:               k,
			CompressedPublicKey:   v.PublicKey,
			UncompressedPublicKey: ledger.NextEpochValidators[k],
			VotingPower:           uint64(v.VotingPower),
		}

		if len(validator.UncompressedPublicKey) == 0 {
			return LedgerInfoLibEpochState{}, false
		}

		if v.VrfPublicKey != nil {
			validator.VrfPublicKey = *v.VrfPublicKey
		}

		validators = append(validators, validator)
	}
	sort.Sort(validators)

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

	var signatures sortableAccountSignatures
	for k, v := range ledger.Signatures {
		signatures = append(signatures, LedgerInfoLibAccountSignature{
			Account:            k,
			ConsensusSignature: v,
		})
	}
	sort.Sort(signatures)
	result.Signatures = signatures

	return result
}

type sortableAccountSignatures []LedgerInfoLibAccountSignature

func (s sortableAccountSignatures) Len() int { return len(s) }
func (s sortableAccountSignatures) Less(i, j int) bool {
	return bytes.Compare(s[i].Account[:], s[j].Account[:]) < 0
}
func (s sortableAccountSignatures) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type sortableValidators []LedgerInfoLibValidatorInfo

func (s sortableValidators) Len() int { return len(s) }
func (s sortableValidators) Less(i, j int) bool {
	return bytes.Compare(s[i].Account[:], s[j].Account[:]) < 0
}
func (s sortableValidators) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
