package types

import "github.com/ethereum/go-ethereum/common/hexutil"

// DepositInfo represents user deposit history
type DepositInfo struct {
	AccumulatedInterestRate *hexutil.Big `json:"accumulatedInterestRate"`
	Amount                  *hexutil.Big `json:"amount"`
	DepositTime             uint64       `json:"depositTime"`
}
