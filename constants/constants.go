package constants

import "math/big"

const (
	CFXName    = "CFX"
	CFXSymbol  = "CFX"
	CFXDecimal = 18
)

const (
	// MinGasprice represents the mininum gasprice required by conflux chain when sending transactions
	// the value of main net is 1G drip
	MinGasprice = 1000000000
)

var (
	MaxUint256, _ = new(big.Int).SetString("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 0)
)
