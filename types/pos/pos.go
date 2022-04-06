package postypes

import (
	"encoding/json"

	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
)

type Address = common.Hash

type Status struct {
	LatestCommitted hexutil.Uint64  `json:"latestCommitted"`
	Epoch           hexutil.Uint64  `json:"epoch"`
	PivotDecision   Decision        `json:"pivotDecision"`
	LatestVoted     *hexutil.Uint64 `json:"latestVoted"`
	LatestTxNumber  hexutil.Uint64  `json:"latestTxNumber"`
}

type Decision struct {
	BlockHash common.Hash    `json:"blockHash"`
	Height    hexutil.Uint64 `json:"height"`
}

type Account struct {
	Address     Address        `json:"address"`
	BlockNumber hexutil.Uint64 `json:"blockNumber"`
	Status      NodeLockStatus `json:"status"`
}

type NodeLockStatus struct {
	InQueue  []VotePowerState `json:"inQueue"`
	Locked   hexutil.Uint64   `json:"locked"`
	OutQueue []VotePowerState `json:"outQueue"`
	Unlocked hexutil.Uint64   `json:"unlocked"`

	// Equals to the summation of in_queue + locked
	AvailableVotes hexutil.Uint64  `json:"availableVotes"`
	ForceRetired   *hexutil.Uint64 `json:"forceRetired"`

	// If the staking is forfeited, the unlocked votes before forfeiting is
	// exempted.
	Forfeited hexutil.Uint64 `json:"forfeited"`
}

type VotePowerState struct {
	EndBlockNumber hexutil.Uint64 `json:"endBlockNumber"`
	Power          hexutil.Uint64 `json:"power"`
}

type CommitteeState struct {
	CurrentCommittee RpcCommittee  `json:"currentCommittee"`
	Elections        []RpcTermData `json:"elections"`
}

type RpcCommittee struct {
	EpochNumber       hexutil.Uint64    `json:"epochNumber"`
	QuorumVotingPower hexutil.Uint64    `json:"quorumVotingPower"`
	TotalVotingPower  hexutil.Uint64    `json:"totalVotingPower"`
	Nodes             []NodeVotingPower `json:"nodes"`
}

type NodeVotingPower struct {
	Address     common.Hash    `json:"address"`
	VotingPower hexutil.Uint64 `json:"votingPower"`
}

type RpcTermData struct {
	StartBlockNumber hexutil.Uint64    `json:"startBlockNumber"`
	IsFinalized      bool              `json:"isFinalized"`
	TopElectingNodes []NodeVotingPower `json:"topElectingNodes"`
}

type Block struct {
	Hash          common.Hash    `json:"hash"`
	Height        hexutil.Uint64 `json:"height"`
	Epoch         hexutil.Uint64 `json:"epoch"`
	Round         hexutil.Uint64 `json:"round"`
	LastTxNumber  hexutil.Uint64 `json:"lastTxNumber"`
	Miner         *Address       `json:"miner"`
	ParentHash    common.Hash    `json:"parentHash"`
	Timestamp     hexutil.Uint64 `json:"timestamp"`
	PivotDecision *Decision      `json:"pivotDecision"`
	// Transactions  BlockTransactions `json:"transactions"`
	Signatures []Signature `json:"signatures"`
}

type Signature struct {
	Account Address        `json:"account"`
	Votes   hexutil.Uint64 `json:"votes"`
}

type Transaction struct {
	Hash        common.Hash         `json:"hash"`
	From        Address             `json:"from"`
	BlockHash   *common.Hash        `json:"blockHash"`
	BlockNumber *hexutil.Uint64     `json:"blockNumber"`
	Timestamp   *hexutil.Uint64     `json:"timestamp"`
	Number      hexutil.Uint64      `json:"number"`
	Payload     *TransactionPayload `json:"payload"`
	Status      *string             `json:"status"`
	Type        string              `json:"type"`
}

type EpochReward struct {
	PowEpochHash   common.Hash `json:"powEpochHash"`
	AccountRewards []Reward    `json:"accountRewards"`
}

type Reward struct {
	PosAddress Address            `json:"posAddress"`
	PowAddress cfxaddress.Address `json:"powAddress"`
	Reward     hexutil.Big        `json:"reward"`
}

func (b *Transaction) UnmarshalJSON(data []byte) error {

	type tmpTransaction struct {
		Hash        common.Hash     `json:"hash"`
		From        Address         `json:"from"`
		BlockHash   *common.Hash    `json:"blockHash"`
		BlockNumber *hexutil.Uint64 `json:"blockNumber"`
		Timestamp   *hexutil.Uint64 `json:"timestamp"`
		Number      hexutil.Uint64  `json:"number"`
		Payload     interface{}     `json:"payload"`
		Status      *string         `json:"status"`
		Type        string          `json:"type"`
	}

	type tmpPayload struct {
		ElectionPayload
		RetirePayload
		// RegisterPayload
		UpdateVotingPowerPayload
		PivotBlockDecision
		DisputePayload
		ConflictSignature
	}

	tmpTx := tmpTransaction{}

	if err := json.Unmarshal(data, &tmpTx); err != nil {
		return errors.WithStack(err)
	}

	*b = Transaction{tmpTx.Hash, tmpTx.From, tmpTx.BlockHash, tmpTx.BlockNumber,
		tmpTx.Timestamp, tmpTx.Number, nil, tmpTx.Status, tmpTx.Type}

	if tmpTx.Payload != nil {
		marshaed, err := json.Marshal(tmpTx.Payload)
		if err != nil {
			return errors.WithStack(err)
		}

		payload := tmpPayload{}
		err = json.Unmarshal(marshaed, &payload)
		if err != nil {
			return errors.WithStack(err)
		}

		realPayload := TransactionPayload{}
		switch tmpTx.Type {
		case "Election":
			realPayload.ElectionPayload = payload.ElectionPayload
		case "Retire":
			realPayload.RetirePayload = payload.RetirePayload
		case "Register":
			realPayload.RegisterPayload = RegisterPayload{
				PublicKey:    payload.ElectionPayload.PublicKey,
				VrfPublicKey: payload.ElectionPayload.VrfPublicKey,
			}
			// fmt.Printf("realPayload.RegisterPayload: %#v\n\n", realPayload.RegisterPayload)
		case "UpdateVotingPower":
			realPayload.UpdateVotingPowerPayload = payload.UpdateVotingPowerPayload
		case "PivotDecision":
			realPayload.PivotBlockDecision = payload.PivotBlockDecision
		case "Dispute":
			realPayload.DisputePayload = payload.DisputePayload
		}

		realPayload.SetTransactionType(tmpTx.Type)
		// fmt.Printf("tmpTxType:%v,payload:%#v, realPayload %#v", tmpTx.Type, payload, realPayload)
		b.Payload = &realPayload
	}

	return nil
}
