package cfxclient

import (
	"testing"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
)

func TestSignalbeClientInterfaceImpl(test *testing.T) {
	h := NewSignableClient(nil, nil)
	var _ sdk.SignableRpcCaller = &h
}
