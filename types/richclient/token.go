// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package richtypes

import (
	"math/big"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
)

// Token describes token detail messages, such as erc20, erc777, fanscoin and so on.
type Token struct {
	TokenName    string         `json:"tokenName"`
	TokenSymbol  string         `json:"tokenSymbol"`
	TokenDecimal uint64         `json:"tokenDecimal"`
	Address      *types.Address `json:"address"`
	TypeCode     uint           `json:"typeCode"`
}

// TokenWithBalance describes token with balace information
type TokenWithBalance struct {
	Token
	Balance string `json:"balance"`
}

// TokenWithBlanceList describes list of token with balance
type TokenWithBlanceList struct {
	Total uint64  `json:"total"`
	List  []Token `json:"list"`
}

// TokenTransferEvent describes token transfer event information
type TokenTransferEvent struct {
	Token
	TransactionHash types.Hash    `json:"transactionHash"`
	Status          uint64        `json:"status"`
	From            types.Address `json:"from"`
	To              types.Address `json:"to"`
	Value           string        `json:"value"`
	Timestamp       uint64        `json:"timestamp"`
	BlockHash       types.Hash
	RevertRate      *big.Float
}

// TokenTransferEventList describes list of token tranfer event information
type TokenTransferEventList struct {
	Total     uint64               `json:"total"`
	ListLimit uint64               `json:"listLimit"`
	List      []TokenTransferEvent `json:"list"`
}
