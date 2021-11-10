package interfaces

import (
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

// Contractor is interface of contract operator
type Contractor interface {
	ABI() abi.ABI
	Address() *types.Address
	Bytecode() []byte

	GetRpcCaller() SignableRpcCaller

	GetData(method string, args ...interface{}) ([]byte, error)
	Call(option *types.ContractMethodCallOption, resultPtr interface{}, method string, args ...interface{}) error
	SendTransaction(option *types.ContractMethodSendOption, method string, args ...interface{}) (types.Hash, error)
	DecodeEvent(out interface{}, event string, log types.Log) error
}
