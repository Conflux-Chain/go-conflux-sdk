package addressutil

import (
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
)

// EtherAddressToCfxAddress converts an ethereum address to a cfx address, it will change top 4 bit of eth address to 0b0001
func EtherAddressToCfxAddress(ethAddr common.Address, isContract bool, networkID uint32) cfxaddress.Address {
	if isContract {
		ethAddr[0] = ethAddr[0]&0x8f | 0x80
	} else {
		ethAddr[0] = ethAddr[0]&0x1f | 0x10
	}
	return cfxaddress.MustNewFromCommon(ethAddr, networkID)
}
