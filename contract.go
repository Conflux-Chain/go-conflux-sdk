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

// GetData packs the given method name to conform the ABI of the contract "c". Method call's data
// will consist of method_id, args0, arg1, ... argN. Method id consists
// of 4 bytes and arguments are all 32 bytes.
// Method ids are created from the first 4 bytes of the hash of the
// methods string signature. (signature = baz(uint32,string32))
func (c *Contract) GetData(method string, args ...interface{}) (*[]byte, error) {
	packed, err := c.ABI.Pack(method, args...)
	if err != nil {
		msg := fmt.Sprintf("encode method %+v with args %+v error", method, args)
		return nil, types.WrapError(err, msg)
	}

	return &packed, nil
}

// Call calls to the contract method with args and fills the excuted result to the "resultPtr".
//
// the resultPtr should be a pointer of the method output struct type.
func (c *Contract) Call(option *types.ContractMethodCallOption, resultPtr interface{}, method string, args ...interface{}) error {

	data, err := c.GetData(method, args...)
	if err != nil {
		msg := fmt.Sprintf("get data of method %+v with args %+v error", method, args)
		return types.WrapError(err, msg)
	}

	tx := new(types.UnsignedTransaction)
	if option != nil {
		tx.UnsignedTransactionBase = types.UnsignedTransactionBase(*option)
	}
	tx.To = c.Address
	tx.Data = *data
	callRequest := new(types.CallRequest)
	callRequest.FillByUnsignedTx(tx)

	// j, err := json.Marshal(callRequest)
	// fmt.Printf("callrequest of call: %s, err:%+v\n\n", j, err)

	resultHexStr, err := c.Client.Call(*callRequest, types.EpochLatestState)

	if len(*resultHexStr) < 2 {
		return fmt.Errorf("call response string %v length smaller than 2", resultHexStr)
	}

	bytes, err := hex.DecodeString((*resultHexStr)[2:])
	if err != nil {
		msg := fmt.Sprintf("decode hex string %s to bytes error", (*resultHexStr)[2:])
		return types.WrapError(err, msg)
	}

	err = c.ABI.Unpack(resultPtr, method, bytes)
	if err != nil {
		msg := fmt.Sprintf("unpack bytes {%x} to method %v output on abi %+v error", bytes, method, c.ABI)
		return types.WrapError(err, msg)
	}
	// fmt.Printf("outptr:%+v", resultPtr)

	return nil
}

// SendTransaction sends a transaction to the contract method with args and returns its transaction hash
func (c *Contract) SendTransaction(option *types.ContractMethodSendOption, method string, args ...interface{}) (*types.Hash, error) {

	data, err := c.GetData(method, args...)
	if err != nil {
		msg := fmt.Sprintf("get data of method %+v with args %+v error", method, args)
		return nil, types.WrapError(err, msg)
	}

	tx := new(types.UnsignedTransaction)
	if option != nil {
		tx.UnsignedTransactionBase = types.UnsignedTransactionBase(*option)
	}
	tx.To = c.Address
	tx.Data = *data

	err = c.Client.ApplyUnsignedTransactionDefault(tx)
	if err != nil {
		msg := fmt.Sprintf("apply default for tx {%+v} error", tx)
		return nil, types.WrapError(err, msg)
	}

	txhash, err := c.Client.SendTransaction(tx)
	if err != nil {
		msg := fmt.Sprintf("send transaction {%+v} error", tx)
		return nil, types.WrapError(err, msg)
	}
	return &txhash, nil
}
