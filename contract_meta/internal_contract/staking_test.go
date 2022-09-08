package internalcontract

import (
	"testing"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/stretchr/testify/assert"
)

func TestNewStaking(t *testing.T) {
	mainClient := sdk.MustNewClient("http://main.confluxrpc.com")
	stake_cfx, err := NewStaking(mainClient)
	assert.NoError(t, err)

	netId, _ := stake_cfx.Contract.Client.GetNetworkID()
	assert.Equal(t, uint32(1029), netId)

	testClient := sdk.MustNewClient("http://test.confluxrpc.com")
	stake_test, err := NewStaking(testClient)
	assert.NoError(t, err)

	netId, _ = stake_test.Contract.Client.GetNetworkID()
	assert.Equal(t, uint32(1), netId)
}
