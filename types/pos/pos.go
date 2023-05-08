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

type VoteParamsInfo struct {
	PowBaseReward *hexutil.Big `json:"powBaseReward"`
	InterestRate  *hexutil.Big `json:"interestRate"`
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

	tmpTx := tmpTransaction{}

	if err := json.Unmarshal(data, &tmpTx); err != nil {
		return errors.WithStack(err)
	}

	*b = Transaction{tmpTx.Hash, tmpTx.From, tmpTx.BlockHash, tmpTx.BlockNumber,
		tmpTx.Timestamp, tmpTx.Number, nil, tmpTx.Status, tmpTx.Type}

	if tmpTx.Payload != nil {
		marshaled, err := json.Marshal(tmpTx.Payload)
		if err != nil {
			return errors.WithStack(err)
		}

		realPayload := TransactionPayload{}
		realPayload.SetTransactionType(tmpTx.Type)
		switch tmpTx.Type {
		case "Election":
			err = json.Unmarshal(marshaled, &realPayload.ElectionPayload)
		case "Retire":
			err = json.Unmarshal(marshaled, &realPayload.RetirePayload)
		case "Register":
			err = json.Unmarshal(marshaled, &realPayload.RegisterPayload)
		case "UpdateVotingPower":
			err = json.Unmarshal(marshaled, &realPayload.UpdateVotingPowerPayload)
		case "PivotDecision":
			err = json.Unmarshal(marshaled, &realPayload.PivotBlockDecision)
		case "Dispute":
			err = json.Unmarshal(marshaled, &realPayload.DisputePayload)
		}
		if err != nil {
			return errors.WithStack(err)
		}

		b.Payload = &realPayload
	}

	return nil
}

// Helper struct to manage validator information for validation
type ValidatorConsensusInfo struct {
	PublicKey string `json:"publicKey"`
	/// None if we do not need VRF.
	VrfPublicKey string         `json:"vrfPublicKey,omitempty"`
	VotingPower  hexutil.Uint64 `json:"votingPower"`
}

// Supports validation of signatures for known authors with individual voting
// powers. This struct can be used for all signature verification operations
// including block and network signature verification, respectively.
type ValidatorVerifier struct {
	// An ordered map of each validator's on-chain account address to its
	// pubkeys and voting power.
	AddressToValidatorInfo map[common.Hash]ValidatorConsensusInfo `json:"addressToValidatorInfo"`
	// The minimum voting power required to achieve a quorum
	QuorumVotingPower hexutil.Uint64 `json:"quorumVotingPower"`
	// Total voting power of all validators (cached from
	// address_to_validator_info)
	TotalVotingPower hexutil.Uint64 `json:"totalVotingPower"`
}

// EpochState represents a trusted validator set to validate messages from the
// specific epoch, it could be updated with EpochChangeProof.
type EpochState struct {
	Epoch    hexutil.Uint64    `json:"epoch"`
	Verifier ValidatorVerifier `json:"verifier"`
	VrfSeed  hexutil.Bytes     `json:"vrfSeed"`
}

/// This structure contains all the information needed for tracking a block
/// without having access to the block or its execution output state. It
/// assumes that the block is the last block executed within the ledger.
type BlockInfo struct {
	/// Epoch number corresponds to the set of validators that are active for
	/// this block.
	Epoch hexutil.Uint64 `json:"epoch"`
	/// The consensus protocol is executed in rounds, which monotonically
	/// increase per epoch.
	Round hexutil.Uint64 `json:"round"`
	/// The identifier (hash) of the block.
	Id common.Hash `json:"id"`
	/// The accumulator root hash after executing this block.
	ExecutedStateId common.Hash `json:"executedStateId"`
	/// The version of the latest transaction after executing this block.
	Version hexutil.Uint64 `json:"version"`
	/// The timestamp this block was proposed by a proposer.
	TimestampUsecs hexutil.Uint64 `json:"timestampUsecs"`
	/// An optional field containing the next epoch info
	NextEpochState *EpochState `json:"nextEpochState"`
	/// TODO(lpl): Remove Option?
	/// The last pivot block selection after executing this block.
	/// None means choosing TreeGraph genesis as the first pivot block.
	Pivot *PivotBlockDecision `json:"pivot"`
}

type LedgerInfo struct {
	CommitInfo BlockInfo `json:"commitInfo"`

	/// Hash of consensus specific data that is opaque to all parts of the
	/// system other than consensus.
	ConsensusDataHash string `json:"consensusDataHash"`
}

type ConsensusSignature string

/// The validator node returns this structure which includes signatures
/// from validators that confirm the state.  The client needs to only pass back
/// the LedgerInfo element since the validator node doesn't need to know the
/// signatures again when the client performs a query, those are only there for
/// the client to be able to verify the state
type LedgerInfoWithSignatures struct {
	LedgerInfo LedgerInfo `json:"ledgerInfo,omitempty"`
	/// The validator is identified by its account address: in order to verify
	/// a signature one needs to retrieve the public key of the validator
	/// for the given epoch.
	Signatures map[common.Hash]ConsensusSignature `json:"signatures"`
}
