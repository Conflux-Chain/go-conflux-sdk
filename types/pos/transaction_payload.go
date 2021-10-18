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
	BlockHash common.Hash    `json:"blockHash"`
}

type DisputePayload struct {
	Address          Address           `json:"address"`
	BlsPubKey        string            `json:"blsPubKey"`
	VrfPubKey        string            `json:"vrfPubKey"`
	ConflictingVotes ConflictSignature `json:"conflictingVotes"`
}

type ConflictSignature struct {
	Proposal [2][]byte
	Vote     [2][]byte
}
