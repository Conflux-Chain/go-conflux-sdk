package postypes

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Address common.Hash

type Status struct {
	LatestCommitted hexutil.Uint64  `json:"latestCommitted"`
	Epoch           hexutil.Uint64  `json:"epoch"`
	PivotDecision   hexutil.Uint64  `json:"pivotDecision"`
	LatestVoted     *hexutil.Uint64 `json:"latestVoted"`
}

type Account struct {
	Address     Address        `json:"address"`
	BlockNumber hexutil.Uint64 `json:"blockNumber"`
	Status      NodeLockStatus `json:"status"`
}

type NodeLockStatus struct {
	InQueue  []VotePowerState `json:"inQueue"`
	Locked   *hexutil.Uint64  `json:"locked"`
	OutQueue []VotePowerState `json:"outQueue"`
	Unlocked hexutil.Uint64   `json:"unlocked"`

	// Equals to the summation of in_queue + locked
	AvailableVotes hexutil.Uint64 `json:"availableVotes"`
	ForceRetired   bool           `json:"forceRetired"`
	// If the staking is forfeited, the unlocked votes before forfeiting is
	// exempted.
	ExemptFromForfeit *hexutil.Uint64 `json:"exemptFromForfeit"`
}

type VotePowerState struct {
	StartBlockNumber hexutil.Uint64 `json:"startBlockNumber"`
	Power            hexutil.Uint64 `json:"power"`
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
	Hash          common.Hash       `json:"hash"`
	Height        hexutil.Uint64    `json:"height"`
	Epoch         hexutil.Uint64    `json:"epoch"`
	Round         hexutil.Uint64    `json:"round"`
	NextTxNumber  hexutil.Uint64    `json:"nextTxNumber"`
	Miner         Address           `json:"miner"`
	ParentHash    common.Hash       `json:"parentHash"`
	Timestamp     hexutil.Uint64    `json:"timestamp"`
	PivotDecision *hexutil.Uint64   `json:"pivotDecision"`
	Transactions  BlockTransactions `json:"transactions"`
	Signatures    []Signature       `json:"signatures"`
}

type BlockTransactions struct {
	Hashes []common.Hash
	Full   []Transaction
}

func (b *BlockTransactions) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &b.Hashes); err != nil {
		err = json.Unmarshal(data, &b.Full)
		return err
	}
	return nil
}

func (b BlockTransactions) MarshalJSON() ([]byte, error) {
	if len(b.Hashes) > 0 {
		return json.Marshal(b.Hashes)
	}
	return json.Marshal(b.Full)
}

type Signature struct {
	Account Address
	Votes   hexutil.Uint64
}

type Transaction struct {
	Hash      common.Hash           `json:"Hash"`
	From      Address               `json:"From"`
	BlockHash *common.Hash          `json:"BlockHash"`
	Number    hexutil.Uint64        `json:"Number"`
	Payload   *TransactionPayload   `json:"Payload"`
	Status    *RpcTransactionStatus `json:"Status"`
}

type RpcTransactionStatus struct {
}
