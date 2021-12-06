package address

import (
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/pkg/errors"
)

func PubkeyToAddress(publicKey string, networkId uint32) (cfxaddress.Address, error) {
	commAddress, err := utils.PublicKeyToCommonAddress(publicKey)
	if err != nil {
		return cfxaddress.Address{}, errors.WithStack(err)
	}
	return cfxaddress.NewFromCommon(commAddress, networkId)
}
