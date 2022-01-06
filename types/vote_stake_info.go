package types

import (
	"encoding/json"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// VoteStakeInfo represents user vote history
type VoteStakeInfo struct {
	/// This is the number of tokens should be locked before
	/// `unlock_block_number`.
	Amount *hexutil.Big `json:"amount"`
	/// This is the timestamp when the vote right will be invalid, measured in
	/// the number of past blocks.
	UnlockBlockNumber uint64 `json:"unlockBlockNumber"`
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *VoteStakeInfo) UnmarshalJSON(data []byte) error {
	tmp := struct {
		Amount            *hexutil.Big `json:"amount"`
		UnlockBlockNumber interface{}  `json:"unlockBlockNumber"`
	}{}

	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	switch tmp.UnlockBlockNumber.(type) {
	case float64:
		v.UnlockBlockNumber = uint64(tmp.UnlockBlockNumber.(float64))
	case string:
		val, err := strconv.ParseUint(tmp.UnlockBlockNumber.(string), 0, 64)
		if err != nil {
			return err
		}
		v.UnlockBlockNumber = val
	}
	v.Amount = tmp.Amount
	return nil
}
