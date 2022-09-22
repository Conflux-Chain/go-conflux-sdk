package internalcontract

import (
	"math/big"
	"sync"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// CrossSpaceCall contract
type CrossSpaceCall struct {
	sdk.Contract
}

var corssSpaceCallMap sync.Map
var crossSpaceCallMu sync.Mutex

// NewCrossSpaceCall gets the CrossSpaceCall contract object
func NewCrossSpaceCall(client sdk.ClientOperator) (s CrossSpaceCall, err error) {
	netId, err := client.GetNetworkID()
	if err != nil {
		return CrossSpaceCall{}, err
	}
	val, ok := corssSpaceCallMap.Load(netId)
	if !ok {
		crossSpaceCallMu.Lock()
		defer crossSpaceCallMu.Unlock()
		abi := getCrossSpaceCallAbi()
		address, e := getCrossSpaceCallAddress(client)
		if e != nil {
			return s, errors.Wrap(e, "failed to get CrossSpaceCall address")
		}
		contract, e := sdk.NewContract([]byte(abi), client, &address)
		if e != nil {
			return s, errors.Wrap(e, "failed to new CrossSpaceCall contract")
		}

		val = CrossSpaceCall{Contract: *contract}
		corssSpaceCallMap.Store(netId, val)
	}
	return val.(CrossSpaceCall), nil
}

func getCrossSpaceCallAbi() string {
	return "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes20\",\"name\":\"sender\",\"type\":\"bytes20\"},{\"indexed\":true,\"internalType\":\"bytes20\",\"name\":\"receiver\",\"type\":\"bytes20\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Call\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes20\",\"name\":\"sender\",\"type\":\"bytes20\"},{\"indexed\":true,\"internalType\":\"bytes20\",\"name\":\"contract_address\",\"type\":\"bytes20\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"init\",\"type\":\"bytes\"}],\"name\":\"Create\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"Outcome\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes20\",\"name\":\"sender\",\"type\":\"bytes20\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"to\",\"type\":\"bytes20\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"callEVM\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"init\",\"type\":\"bytes\"}],\"name\":\"createEVM\",\"outputs\":[{\"internalType\":\"bytes20\",\"name\":\"\",\"type\":\"bytes20\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"mappedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"mappedNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"to\",\"type\":\"bytes20\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"staticCallEVM\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"to\",\"type\":\"bytes20\"}],\"name\":\"transferEVM\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdrawFromMapped\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
}

func getCrossSpaceCallAddress(client sdk.ClientOperator) (types.Address, error) {
	addr := cfxaddress.MustNewFromHex("0888000000000000000000000000000000000006")
	err := addr.CompleteByClient(client)
	return addr, err
}

// =================== calls ==================

func (ac *CrossSpaceCall) MappedBalance(opts *types.ContractMethodCallOption, addr types.Address) (*big.Int, error) {
	var tmp *big.Int = new(big.Int)
	err := ac.Call(opts, &tmp, "mappedBalance", addr.MustGetCommonAddress())
	if err != nil {
		return nil, err
	}
	return tmp, nil
}

func (ac *CrossSpaceCall) MappedNonce(opts *types.ContractMethodCallOption, addr types.Address) (*big.Int, error) {
	var tmp *big.Int = new(big.Int)
	err := ac.Call(opts, &tmp, "mappedNonce", addr.MustGetCommonAddress())
	if err != nil {
		return nil, err
	}
	return tmp, nil
}

func (ac *CrossSpaceCall) StaticCallEVM(opts *types.ContractMethodCallOption, to common.Address, data []byte) ([]byte, error) {
	var tmp []byte
	err := ac.Call(opts, &tmp, "staticCallEVM", to, data)
	if err != nil {
		return nil, err
	}
	return tmp, nil
}

// =================== sends ==================

func (s *CrossSpaceCall) CallEVM(opts *types.ContractMethodSendOption, to common.Address, data []byte) (types.Hash, error) {
	return s.SendTransaction(opts, "callEVM", to, data)
}

func (s *CrossSpaceCall) CreateEVM(opts *types.ContractMethodSendOption, init []byte) (types.Hash, error) {
	return s.SendTransaction(opts, "createEVM", init)
}

func (s *CrossSpaceCall) TransferEVM(opts *types.ContractMethodSendOption, to common.Address) (types.Hash, error) {
	return s.SendTransaction(opts, "transferEVM", to)
}

func (s *CrossSpaceCall) WithdrawFromMapped(opts *types.ContractMethodSendOption, value *big.Int) (types.Hash, error) {
	return s.SendTransaction(opts, "withdrawFromMapped", value)
}

// =================== events ==================

type CrossSpaceCallCall struct {
	Sender   common.Address
	Receiver common.Address
	Value    *big.Int
	Nonce    *big.Int
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

func (s *CrossSpaceCall) ParseCall(log types.Log) (*CrossSpaceCallCall, error) {
	event := new(CrossSpaceCallCall)
	if err := s.DecodeEvent(event, "Call", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossSpaceCallCreate struct {
	Sender          [20]byte
	ContractAddress [20]byte
	Value           *big.Int
	Nonce           *big.Int
	Init            []byte
	Raw             types.Log // Blockchain specific contextual infos
}

func (s *CrossSpaceCall) ParseCreate(log types.Log) (*CrossSpaceCallCreate, error) {
	event := new(CrossSpaceCallCreate)
	if err := s.DecodeEvent(event, "Create", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossSpaceCallOutcome represents a Outcome event raised by the CrossSpaceCall contract.
type CrossSpaceCallOutcome struct {
	Success bool
	Raw     types.Log // Blockchain specific contextual infos
}

// ParseOutcome is a log parse operation binding the contract event 0xbc11eabb6efd378a0a489b58a574c6e0d0403060e8a8c7b8eab45db47900edfe.
//
// Solidity: event Outcome(bool success)
func (s *CrossSpaceCall) ParseOutcome(log types.Log) (*CrossSpaceCallOutcome, error) {
	event := new(CrossSpaceCallOutcome)
	if err := s.DecodeEvent(event, "Outcome", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossSpaceCallWithdraw represents a Withdraw event raised by the CrossSpaceCall contract.
type CrossSpaceCallWithdraw struct {
	Sender   [20]byte
	Receiver common.Address
	Value    *big.Int
	Nonce    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// ParseWithdraw is a log parse operation binding the contract event 0x31328e08abcc622b23d8be96d45b371b10e42989dafc8ac56c85b33bb3584b92.
//
// Solidity: event Withdraw(bytes20 indexed sender, address indexed receiver, uint256 value, uint256 nonce)
func (s *CrossSpaceCall) ParseWithdraw(log types.Log) (*CrossSpaceCallWithdraw, error) {
	event := new(CrossSpaceCallWithdraw)
	if err := s.DecodeEvent(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
