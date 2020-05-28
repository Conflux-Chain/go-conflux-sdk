// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package sdk

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
)

// Contract represents a smart contract.
// You can conveniently create contract by Client.GetContrat or Client.DeployContract.
type Contract struct {
	ABI     abi.ABI
	Client  ClientOperator
	Address *types.Address
}

// ContractDeployResult for state change notification when deploying contract
type ContractDeployResult struct {
	//DoneChannel channel for notifying when contract deployed done
	DoneChannel      <-chan struct{}
	TransactionHash  *types.Hash
	Error            error
	DeployedContract *Contract
}

// NewContract creates contract by abi and deployed address
func NewContract(abiJSON []byte, client ClientOperator, address *types.Address) (*Contract, error) {
	if client == nil {
		client = (*Client)(nil)
	}
	return client.GetContract(abiJSON, address)
}

// GetData packs the given method name to conform the ABI of the contract. Method call's data
// will consist of method_id, args0, arg1, ... argN. Method id consists
// of 4 bytes and arguments are all 32 bytes.
// Method ids are created from the first 4 bytes of the hash of the
// methods string signature. (signature = baz(uint32,string32))
//
// please refer https://github.com/Conflux-Chain/go-conflux-sdk/blob/master/README.md to
// get the mappings of solidity types to go types
func (contract *Contract) GetData(method string, args ...interface{}) ([]byte, error) {
	packed, err := contract.ABI.Pack(method, args...)
	if err != nil {
		msg := fmt.Sprintf("encode method %+v with args %+v error", method, args)
		return nil, types.WrapError(err, msg)
	}

	return packed, nil
}

// Call calls to the contract method with args and fills the excuted result to the "resultPtr".
//
// the resultPtr should be a pointer of the method output struct type.
//
// please refer https://github.com/Conflux-Chain/go-conflux-sdk/blob/master/README.md to
// get the mappings of solidity types to go types
func (contract *Contract) Call(option *types.ContractMethodCallOption, resultPtr interface{}, method string, args ...interface{}) error {

	data, err := contract.GetData(method, args...)
	if err != nil {
		msg := fmt.Sprintf("get data of method %+v with args %+v error", method, args)
		return types.WrapError(err, msg)
	}

	callRequest := new(types.CallRequest)
	callRequest.To = contract.Address
	callRequest.Data = "0x" + hex.EncodeToString(data)
	callRequest.FillByCallOption(option)

	var epoch *types.Epoch = nil
	if option != nil && option.Epoch != nil {
		epoch = option.Epoch
	}
	resultHexStr, err := contract.Client.Call(*callRequest, epoch)
	if err != nil {
		msg := fmt.Sprintf("call {%+v} at epoch %+v error", *callRequest, epoch)
		return types.WrapError(err, msg)
	}

	if len(*resultHexStr) < 2 {
		return fmt.Errorf("call response string %v length smaller than 2", resultHexStr)
	}

	bytes, err := hex.DecodeString((*resultHexStr)[2:])
	if err != nil {
		msg := fmt.Sprintf("decode hex string %s to bytes error", (*resultHexStr)[2:])
		return types.WrapError(err, msg)
	}

	err = contract.ABI.Unpack(resultPtr, method, bytes)
	if err != nil {
		msg := fmt.Sprintf("unpack bytes {%x} to method %v output on abi %+v error", bytes, method, contract.ABI)
		return types.WrapError(err, msg)
	}

	return nil
}

// SendTransaction sends a transaction to the contract method with args and returns its transaction hash
//
// please refer https://github.com/Conflux-Chain/go-conflux-sdk/blob/master/README.md to
// get the mappings of solidity types to go types
func (contract *Contract) SendTransaction(option *types.ContractMethodSendOption, method string, args ...interface{}) (*types.Hash, error) {

	data, err := contract.GetData(method, args...)
	if err != nil {
		msg := fmt.Sprintf("get data of method %+v with args %+v error", method, args)
		return nil, types.WrapError(err, msg)
	}

	tx := new(types.UnsignedTransaction)
	if option != nil {
		tx.UnsignedTransactionBase = types.UnsignedTransactionBase(*option)
	}
	tx.To = contract.Address
	tx.Data = data

	err = contract.Client.ApplyUnsignedTransactionDefault(tx)
	if err != nil {
		msg := fmt.Sprintf("apply default for tx {%+v} error", tx)
		return nil, types.WrapError(err, msg)
	}

	txhash, err := contract.Client.SendTransaction(tx)
	if err != nil {
		msg := fmt.Sprintf("send transaction {%+v} error", tx)
		return nil, types.WrapError(err, msg)
	}
	return &txhash, nil
}

// DecodeEvent unpacks a retrieved log into the provided output structure.
//
// please refer https://github.com/Conflux-Chain/go-conflux-sdk/blob/master/README.md to
// get the mappings of solidity types to go types
func (contract *Contract) DecodeEvent(out interface{}, event string, log types.LogEntry) error {

	topics := make([]common.Hash, len(log.Topics))
	for i, v := range log.Topics {
		topics[i] = *v.ToCommonHash()
	}
	eLog := etypes.Log{}
	eLog.Topics = topics
	eLog.Data, _ = hex.DecodeString(strings.Replace(log.Data, "0x", "", -1))
	// fmt.Printf("elog: %+v\n", eLog)

	addressPtr := new(common.Address)
	if contract.Address != nil {
		addressPtr = contract.Address.ToCommonAddress()
	}

	boundContract := bind.NewBoundContract(*addressPtr, contract.ABI, nil, nil, nil)
	err := boundContract.UnpackLog(out, event, eLog)
	if err != nil {
		return err
	}

	return nil
}
