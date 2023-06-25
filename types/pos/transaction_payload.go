package postypes

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
)

type TransactionPayload struct {
	transactionType string

	ElectionPayload
	RetirePayload
	RegisterPayload
	UpdateVotingPowerPayload
	PivotBlockDecision
	DisputePayload
}

func (t *TransactionPayload) SetTransactionType(txType string) {
	t.transactionType = txType
}

func (b TransactionPayload) MarshalJSON() ([]byte, error) {
	switch b.transactionType {
	case "Election":
		return json.Marshal(b.ElectionPayload)
	case "Retire":
		return json.Marshal(b.RetirePayload)
	case "Register":
		return json.Marshal(b.RegisterPayload)
	case "UpdateVotingPower":
		return json.Marshal(b.UpdateVotingPowerPayload)
	case "PivotDecision":
		return json.Marshal(b.PivotBlockDecision)
	case "Dispute":
		return json.Marshal(b.DisputePayload)
	}
	return nil, nil
}

func (b *TransactionPayload) UnmarshalJSON(data []byte) error {
	return errors.New("not support unmarshal TransactionPayload directly, because need transactionType info")
}

type ElectionPayload struct {
	PublicKey    string         `json:"publicKey"`
	VrfPublicKey string         `json:"vrfPublicKey"`
	TargetTerm   hexutil.Uint64 `json:"targetTerm"`
	VrfProof     string         `json:"vrfProof"`
}

type RetirePayload struct {
	NodeId Address        `json:"nodeId"`
	Votes  hexutil.Uint64 `json:"votes"`
}

type RegisterPayload struct {
	PublicKey    string `json:"publicKey"`
	VrfPublicKey string `json:"vrfPublicKey"`
}

type UpdateVotingPowerPayload struct {
	NodeAddress Address        `json:"nodeAddress"`
	VotingPower hexutil.Uint64 `json:"votingPower"`
}

type PivotBlockDecision struct {
	Height    hexutil.Uint64 `json:"height"`
	BlockHash H256           `json:"blockHash"`
}

// for BCS serialization purpose
type H256 string

func (h H256) ToHash() common.Hash {
	return common.HexToHash(string(h))
}

func (h H256) String() string {
	return string(h)
}

type DisputePayload struct {
	Address          Address          `json:"address"`
	BlsPublicKey     string           `json:"blsPublicKey"`
	VrfPublicKey     string           `json:"vrfPublicKey"`
	ConflictingVotes ConflictingVotes `json:"conflictingVotes"`
}

type ConflictingVotes struct {
	ConflictVoteType string `json:"conflictVoteType"`
	First            string `json:"first"`
	Second           string `json:"second"`
}

type ConflictSignature struct {
	Proposal [2][]byte
	Vote     [2][]byte
}
