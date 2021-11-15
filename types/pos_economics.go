package types

import "github.com/ethereum/go-ethereum/common/hexutil"

type PoSEconomics struct {
	TotalPosStakingTokens    *hexutil.Big   `json:"totalPosStakingTokens"`
	DistributablePosInterest *hexutil.Big   `json:"distributablePosInterest"`
	LastDistributeBlock      hexutil.Uint64 `json:"lastDistributeBlock"`
}
