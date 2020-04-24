package constants

import "math/big"

const (
	CFXName    = "CFX"
	CFXSymbol  = "CFX"
	CFXDecimal = 18
)

const (
	MinGasprice = 1
)

const (
	RpcConcurrence = 10
)

var (
	MaxUint256, _ = new(big.Int).SetString("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 0)
)
