// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package richtypes

// ContractType represents contract type
type ContractType uint

// Contract describe response contract information of scan rest api request
type Contract struct {
	// "address": "exercitation qui anim",
	TypeCode     uint   `json:"typeCode"` // "typeCode": -87108642.2575568,
	ContractName string `json:"name"`     //"name": "fugiat Lorem esse",
	// "webside": "sunt",
	ABI string `json:"abi"`
	// "sourceCode": "Ut cillum exercitation tempor",
	// "icon": "dolor eiusmod in",
	TokenSymbol   string `json:"tokenSymbol"`   // "tokenSymbol": "magna deserunt cillum ullamco",
	TokenDecimals uint64 `json:"tokenDecimals"` // "tokenDecimals": 54105223.43683612,
	TokenIcon     string `json:"tokenIcon"`     // "tokenIcon": "mollit nulla enim",
	TokenName     string `json:"tokenName"`     // "tokenName": "non dolore"
}

// GetContractType return contract type
func (c *Contract) GetContractType() ContractType {
	if c.TypeCode == 0 {
		return GENERAL
	}
	if c.TypeCode >= 100 && c.TypeCode < 200 {
		return ERC20
	}
	if c.TypeCode >= 200 && c.TypeCode < 300 {
		return ERC777
	}
	if c.TypeCode == 201 {
		return FANSCOIN
	}
	if c.TypeCode >= 500 && c.TypeCode < 600 {
		return ERC721
	}
	if c.TypeCode >= 1000 {
		return DEX
	}
	return UNKNOWN
}

const (
	// UNKNOWN contract
	UNKNOWN ContractType = iota
	// GENERAL contract
	GENERAL
	// ERC20 contract
	ERC20
	// ERC777 contract
	ERC777
	// FANSCOIN contract
	FANSCOIN
	// ERC721 contract
	ERC721
	// DEX contract
	DEX
)

// String implements the fmt.Stringer interface
func (c ContractType) String() string {
	dic := make(map[ContractType]string)
	dic[UNKNOWN] = "unknown"
	dic[GENERAL] = "general"
	dic[ERC20] = "erc20"
	dic[ERC777] = "erc777"
	dic[FANSCOIN] = "fanscoin"
	dic[ERC721] = "erc721"
	dic[DEX] = "dex"
	return dic[c]
}
