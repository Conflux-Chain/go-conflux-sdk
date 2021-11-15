package types

import "github.com/ethereum/go-ethereum/common/hexutil"

type TxWithPoolInfo struct {
	Exist              bool         `json:"exist"`
	Packed             bool         `json:"packed"`
	LocalNonce         *hexutil.Big `json:"localNonce"`
	LocalBalance       *hexutil.Big `json:"localBalance"`
	StateNonce         *hexutil.Big `json:"stateNonce"`
	StateBalance       *hexutil.Big `json:"stateBalance"`
	LocalBalanceEnough bool         `json:"localBalanceEnough"`
	StateBalanceEnough bool         `json:"stateBalanceEnough"`
}
type TxPoolPendingNonceRange struct {
	MinNonce *hexutil.Big `json:"minNonce"`
	MaxNonce *hexutil.Big `json:"maxNonce"`
}

type TxPoolStatus struct {
	Deferred   hexutil.Uint64 `json:"deferred"`
	Ready      hexutil.Uint64 `json:"ready"`
	Received   hexutil.Uint64 `json:"received"`
	Unexecuted hexutil.Uint64 `json:"unexecuted"`
}
