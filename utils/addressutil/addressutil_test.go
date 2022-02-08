package addressutil

import (
	"testing"

	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/stretchr/testify/assert"
)

func TestCfxMappedEVMSpaceAddress(t *testing.T) {
	cfxAddr := cfxaddress.MustNewFromBase32("cfx:aak2rra2njvd77ezwjvx04kkds9fzagfe6ku8scz91")
	evmAddr, _ := CfxMappedEVMSpaceAddress(cfxAddr)
	assert.Equal(t, "0x12Bf6283CcF8Ad6ffA63f7Da63EDc217228d839A", evmAddr.String())
}
