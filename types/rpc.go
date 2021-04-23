package types

import "github.com/ethereum/go-ethereum/common/hexutil"

type AccountInfo struct {
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
	BestHash         Hash           `json:"bestHash"`
	ChainID          hexutil.Uint   `json:"chainId"`
	NetworkID        hexutil.Uint   `json:"networkId"`
	EpochNumber      hexutil.Uint64 `json:"epochNumber"`
	BlockNumber      hexutil.Uint64 `json:"blockNumber"`
	PendingTxNumber  hexutil.Uint64 `json:"pendingTxNumber"`
	LatestCheckpoint hexutil.Uint64 `json:"latestCheckpoint"`
	LatestConfirmed  hexutil.Uint64 `json:"latestConfirmed"`
	LatestState      hexutil.Uint64 `json:"latestState"`
}

type StorageRoot struct {
	Delta        Hash `json:"delta"`        //delta: H256,
	Intermediate Hash `json:"intermediate"` //intermediate: H256,
	Snapshot     Hash `json:"snapshot"`     //snapshot: H256,
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
	TotalCirculating *hexutil.Big `json:"totalCirculating"`
	TotalIssued      *hexutil.Big `json:"totalIssued"`
	TotalStaking     *hexutil.Big `json:"totalStaking"`
	TotalCollateral  *hexutil.Big `json:"totalCollateral"`
}

type ChainReorg struct {
	RevertTo *hexutil.Big `json:"revertTo"`
}

type AccountPendingInfo struct {
	LocalNonce    *hexutil.Big `json:"localNonce"`
	PendingCount  *hexutil.Big `json:"pendingCount"`
	PendingNonce  *hexutil.Big `json:"pendingNonce"`
	NextPendingTx Hash         `json:"nextPendingTx"`
}
