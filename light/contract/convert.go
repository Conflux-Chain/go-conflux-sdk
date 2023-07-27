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

		uncompressedPubKey, ok := ledger.NextEpochValidators[v]
		if !ok {
			return LedgerInfoLibEpochState{}, false
		}

		validator := LedgerInfoLibValidatorInfo{
			Account:               v,
			UncompressedPublicKey: ABIEncodePublicKey(uncompressedPubKey),
			VotingPower:           uint64(info.VotingPower),
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

	result.AggregatedSignature = ABIEncodeSignature(ledger.AggregatedSignature)
	for _, v := range ledger.ValidatorsSorted() {
		result.Accounts = append(result.Accounts, v)
	}

	return result
}

func ABIEncodeSignature(signature []byte) []byte {
	if len(signature) != 192 {
		return signature
	}

	encoded := make([]byte, 256)

	copy(encoded[16:64], signature[:48])
	copy(encoded[80:128], signature[48:96])
	copy(encoded[144:192], signature[96:144])
	copy(encoded[208:256], signature[144:192])

	return encoded
}

func ABIEncodePublicKey(publicKey []byte) []byte {
	if len(publicKey) != 96 {
		return publicKey
	}

	encoded := make([]byte, 128)

	copy(encoded[16:64], publicKey[:48])
	copy(encoded[80:128], publicKey[48:])

	return encoded
}
