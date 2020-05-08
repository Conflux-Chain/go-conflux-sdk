// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package sdk

import (
	"encoding/hex"
	"fmt"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

// Contract represents a smart contract
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

// GetData packs the given method name to conform the ABI of the contract "c". Method call's data
// will consist of method_id, args0, arg1, ... argN. Method id consists
// of 4 bytes and arguments are all 32 bytes.
// Method ids are created from the first 4 bytes of the hash of the
// methods string signature. (signature = baz(uint32,string32))
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
	// j, err := json.Marshal(callRequest)
	// fmt.Printf("callrequest of call: %s, err:%+v\n\n", j, err)
	var epoch *types.Epoch = nil
	if option != nil && option.Epoch != nil {
		epoch = option.Epoch
	}
	resultHexStr, err := contract.Client.Call(*callRequest, epoch)

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
	// fmt.Printf("outptr:%+v", resultPtr)

	return nil
}

// SendTransaction sends a transaction to the contract method with args and returns its transaction hash
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
