package client

import (
	"testing"

	"github.com/Conflux-Chain/go-conflux-sdk/interfaces"
)

func TestSignalbeClientInterfaceImpl(test *testing.T) {
	h := NewSignableClient(nil, nil)
	var _ interfaces.SignableRpcCaller = &h
}
