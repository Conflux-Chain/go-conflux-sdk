package postypes

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type TransactionPayload struct {
	RetirePayload
	RegisterPayload
	UpdateVotingPowerPayload
	PivotBlockDecision
	DisputePayload
	ConflictSignature
}

type RetirePayload struct {
	NodeId Address        `json:"nodeId"`
	Votes  hexutil.Uint64 `json:"votes"`
}

type RegisterPayload struct {
	PublicKey    common.Hash `json:"publicKey"`
	VrfPublicKey common.Hash `json:"vrfPublicKey"`
}

type UpdateVotingPowerPayload struct {
	NodeAddress Address        `json:"nodeAddress"`
	VotingPower hexutil.Uint64 `json:"votingPower"`
}

type PivotBlockDecision struct {
	Height    uint64      `json:"height"`
	BlockHash common.Hash `json:"blockHash"`
}

type DisputePayload struct {
	Address          Address           `json:"address"`
	BlsPubKey        common.Hash       `json:"blsPubKey"`
	VrfPubKey        common.Hash       `json:"vrfPubKey"`
	ConflictingVotes ConflictSignature `json:"conflictingVotes"`
}

type ConflictSignature struct {
	Proposal [2][]byte
	Vote     [2][]byte
}
