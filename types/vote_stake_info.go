package types

import "github.com/ethereum/go-ethereum/common/hexutil"

// VoteStakeInfo represents user vote history
type VoteStakeInfo struct {
	/// This is the number of tokens should be locked before
	/// `unlock_block_number`.
	Amount *hexutil.Big `json:"amount"`
	/// This is the timestamp when the vote right will be invalid, measured in
	/// the number of past blocks.
	UnlockBlockNumber hexutil.Uint64 `json:"unlockBlockNumber"`
}
