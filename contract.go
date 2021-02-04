// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package sdk

import (
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

// Contract represents a smart contract.
// You can conveniently create contract by Client.GetContract or Client.DeployContract.
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
		return nil, err
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
	// fmt.Printf("get data of method %v with args %v: %x \n", method, args, data)
	if err != nil {
		return errors.Wrap(err, "failed to encode call data")
	}

	callRequest := new(types.CallRequest)
	callRequest.To = contract.Address
	hexData := hexutil.Bytes(data).String()
	callRequest.Data = &hexData
	callRequest.FillByCallOption(option)

	var epoch *types.Epoch = nil
	if option != nil && option.Epoch != nil {
		epoch = option.Epoch
	}
	// fmt.Printf("data: %x,hexdata:%v,callRequest.Data:%v\n", data, hexData, *callRequest.Data)
	resultHexStr, err := contract.Client.Call(*callRequest, epoch)
	if err != nil {
		return errors.Wrapf(err, "failed to call %+v at epoch %v", *callRequest, epoch)
	}

	if len(resultHexStr) < 2 {
		return errors.Errorf("call response string %v length smaller than 2", resultHexStr)
	}

	bytes := []byte(resultHexStr)

	err = contract.ABI.UnpackIntoInterface(resultPtr, method, bytes)
	if err != nil {
		return errors.Wrapf(err, "failed to decode call result, encoded data = %v", resultHexStr)
	}

	return nil
}

// SendTransaction sends a transaction to the contract method with args and returns its transaction hash
//
// please refer https://github.com/Conflux-Chain/go-conflux-sdk/blob/master/README.md to
// get the mappings of solidity types to go types
func (contract *Contract) SendTransaction(option *types.ContractMethodSendOption, method string, args ...interface{}) (types.Hash, error) {

	data, err := contract.GetData(method, args...)
	if err != nil {
		return "", errors.Wrap(err, "failed to encode call data")
	}

	tx := new(types.UnsignedTransaction)
	if option != nil {
		tx.UnsignedTransactionBase = types.UnsignedTransactionBase(*option)
	}
	tx.To = contract.Address
	tx.Data = data

	err = contract.Client.ApplyUnsignedTransactionDefault(tx)
	if err != nil {
		return "", errors.Wrap(err, errMsgApplyTxValues)
	}

	txhash, err := contract.Client.SendTransaction(*tx)
	if err != nil {
		return "", errors.Wrapf(err, "failed to send transaction %+v", tx)
	}
	return txhash, nil
}

// DecodeEvent unpacks a retrieved log into the provided output structure.
//
// please refer https://github.com/Conflux-Chain/go-conflux-sdk/blob/master/README.md to
// get the mappings of solidity types to go types
func (contract *Contract) DecodeEvent(out interface{}, event string, log types.Log) error {

	topics := make([]common.Hash, len(log.Topics))
	for i, v := range log.Topics {
		topics[i] = *v.ToCommonHash()
	}
	eLog := etypes.Log{}
	eLog.Topics = topics
	eLog.Data = []byte(log.Data)
	// fmt.Printf("elog: %+v\n", eLog)

	addressPtr := new(common.Address)
	if contract.Address != nil {
		var err error
		*addressPtr, _, err = contract.Address.ToCommon()
		if err != nil {
			return errors.Wrap(err, "failed to parse contract address to hex address")
		}
	}

	boundContract := bind.NewBoundContract(*addressPtr, contract.ABI, nil, nil, nil)
	err := boundContract.UnpackLog(out, event, eLog)
	if err != nil {
		return errors.Wrapf(err, "failed to unpack log")
	}

	return nil
}
