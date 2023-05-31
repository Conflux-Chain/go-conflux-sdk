package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type AccountInfo struct {
	Address Address `json:"address"`

	Balance *hexutil.Big `json:"balance"`
	// : U256,
	Nonce *hexutil.Big `json:"nonce"`
	// : U256,
	CodeHash Hash `json:"codeHash"`
	//  : H256,
	StakingBalance *hexutil.Big `json:"stakingBalance"`
	// : U256,
	CollateralForStorage *hexutil.Big `json:"collateralForStorage"`
	// : U256,
	AccumulatedInterestReturn *hexutil.Big `json:"accumulatedInterestReturn"`
	// : U256,
	Admin Address `json:"admin"`
	// : H160,
}

//Estimate represents estimated gas will be used and storage will be collateralized when transaction excutes
type Estimate struct {
	GasLimit              *hexutil.Big `json:"gasLimit"`
	GasUsed               *hexutil.Big `json:"gasUsed"`
	StorageCollateralized *hexutil.Big `json:"storageCollateralized"`
}

type RewardInfo struct {
	BlockHash Hash `json:"blockHash"`
	// H256,
	Author Address `json:"author"`
	// H160,
	TotalReward *hexutil.Big `json:"totalReward"`
	// U256,
	BaseReward *hexutil.Big `json:"baseReward"`
	// U256,
	TxFee *hexutil.Big `json:"txFee"`
	// U256,
}

type SponsorInfo struct {
	SponsorForGas               Address      `json:"sponsorForGas"`
	SponsorForCollateral        Address      `json:"sponsorForCollateral"`
	SponsorGasBound             *hexutil.Big `json:"sponsorGasBound"`
	SponsorBalanceForGas        *hexutil.Big `json:"sponsorBalanceForGas"`
	SponsorBalanceForCollateral *hexutil.Big `json:"sponsorBalanceForCollateral"`
}

// Status represents current blockchain status
type Status struct {
	BestHash             Hash           `json:"bestHash"`
	ChainID              hexutil.Uint64 `json:"chainId"`
	EthereumSpaceChainId hexutil.Uint64 `json:"ethereumSpaceChainId"`
	NetworkID            hexutil.Uint64 `json:"networkId"`
	EpochNumber          hexutil.Uint64 `json:"epochNumber"`
	BlockNumber          hexutil.Uint64 `json:"blockNumber"`
	PendingTxNumber      hexutil.Uint64 `json:"pendingTxNumber"`
	LatestCheckpoint     hexutil.Uint64 `json:"latestCheckpoint"`
	LatestConfirmed      hexutil.Uint64 `json:"latestConfirmed"`
	LatestState          hexutil.Uint64 `json:"latestState"`
	LatestFinalized      hexutil.Uint64 `json:"latestFinalized"`
}

type StorageRoot struct {
	Delta        *Hash `json:"delta"`        //delta: H256,
	Intermediate *Hash `json:"intermediate"` //intermediate: H256,
	Snapshot     *Hash `json:"snapshot"`     //snapshot: H256,
}

type CheckBalanceAgainstTransactionResponse struct {
	/// Whether the account should pay transaction fee by self.
	WillPayTxFee bool `json:"willPayTxFee"`
	/// Whether the account should pay collateral by self.
	WillPayCollateral bool `json:"willPayCollateral"`
	/// Whether the account balance is enough for this transaction.
	IsBalanceEnough bool `json:"isBalanceEnough"`
}

type TokenSupplyInfo struct {
	TotalCirculating  *hexutil.Big `json:"totalCirculating"`
	TotalIssued       *hexutil.Big `json:"totalIssued"`
	TotalStaking      *hexutil.Big `json:"totalStaking"`
	TotalCollateral   *hexutil.Big `json:"totalCollateral"`
	TotalEspaceTokens *hexutil.Big `json:"totalEspaceTokens"`
}

type ChainReorg struct {
	RevertTo *hexutil.Big `json:"revertTo"`
}

type rlpEncodableChainReorg struct {
	RevertTo *big.Int
}

func (c ChainReorg) toRlpEncodable() rlpEncodableChainReorg {
	if c.RevertTo != nil {
		return rlpEncodableChainReorg{c.RevertTo.ToInt()}
	}
	return rlpEncodableChainReorg{}
}

func (r rlpEncodableChainReorg) toNormal() ChainReorg {
	if r.RevertTo != nil {
		return ChainReorg{NewBigIntByRaw(r.RevertTo)}
	}
	return ChainReorg{}
}

type AccountPendingInfo struct {
	LocalNonce    *hexutil.Big `json:"localNonce"`
	PendingCount  *hexutil.Big `json:"pendingCount"`
	PendingNonce  *hexutil.Big `json:"pendingNonce"`
	NextPendingTx Hash         `json:"nextPendingTx"`
}

type StorageCollateralInfo struct {
	TotalStorageTokens     *hexutil.Big `json:"totalStorageTokens"`
	ConvertedStoragePoints *hexutil.Big `json:"convertedStoragePoints"`
	UsedStoragePoints      *hexutil.Big `json:"usedStoragePoints"`
}

type EpochReceiptProof struct {
	BlockIndexProof   TrieProof `json:"block_index_proof"`
	BlockReceiptProof TrieProof `json:"block_receipt_proof"`
}

type TrieProof []VanillaTrieNode

type VanillaTrieNode struct {
	ChildrenTable  []common.Hash `json:"childrenTable"`
	CompressedPath struct {
		PathMask  hexutil.Uint  `json:"pathMask"`
		PathSlice hexutil.Bytes `json:"pathSlice"`
	} `json:"compressedPath"`
	MerkleHash common.Hash    `json:"merkleHash"`
	MptValue   *hexutil.Bytes `json:"mptValue"`
}

// pub struct EpochReceiptProof {
//     pub block_index_proof: TrieProof,
//     pub block_receipt_proof: TrieProof,
// }

// #[derive(Clone, Debug, Default, PartialEq)]
// pub struct TrieProof {
//     /// The first node must be the root node. Child node must come later than
//     /// one of its parent node.
//     nodes: Vec<TrieProofNode>,
//     merkle_to_node_index: HashMap<MerkleHash, usize>,
//     number_leaf_nodes: u32,
// }

// #[derive(Clone, Debug, Default, PartialEq, Serialize, Deserialize)]
// pub struct TrieProofNode(VanillaTrieNode<MerkleHash>);
