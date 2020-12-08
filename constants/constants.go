package constants

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	CFXName    = "CFX"
	CFXSymbol  = "CFX"
	CFXDecimal = 18
)

const (
	// MinGasprice represents the mininum gasprice required by conflux chain when sending transactions
	// the value of main net is 1 Gdrip
	MinGasprice = 1
)

var (
	MaxUint256, _ = new(big.Int).SetString("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 0)
)

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

var (
	ZeroAddress common.Address
)
