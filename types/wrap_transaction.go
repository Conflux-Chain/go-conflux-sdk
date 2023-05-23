package types

import "github.com/openweb3/web3go/types"

type WrapTransaction struct {
	NativeTransaction *Transaction       `json:"nativeTransaction"`
	EthTransaction    *types.Transaction `json:"ethTransaction"`
}
