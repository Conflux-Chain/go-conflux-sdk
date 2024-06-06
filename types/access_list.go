package types

import (
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
)

type AccessList []AccessTuple

type AccessTuple struct {
	Address     cfxaddress.Address `json:"address"`
	StorageKeys []common.Hash      `json:"storageKeys"`
}

func (a *AccessList) ToEthType() *etypes.AccessList {
	var eValue etypes.AccessList
	for _, tuple := range *a {
		eValue = append(eValue, etypes.AccessTuple{
			Address:     tuple.Address.MustGetCommonAddress(),
			StorageKeys: tuple.StorageKeys,
		})
	}
	return &eValue
}

func (a *AccessList) FromEthType(raw *etypes.AccessList, networkID uint32) {
	*a = AccessList{}
	for _, tuple := range *raw {
		*a = append(*a, AccessTuple{
			Address:     cfxaddress.MustNewFromCommon(tuple.Address, networkID),
			StorageKeys: tuple.StorageKeys,
		})
	}
}
