// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package contracts

import (
	"time"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	cfxerrors "github.com/Conflux-Chain/go-conflux-sdk/types/errors"
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
	client   sdk.SignableRpcCaller
	abi      abi.ABI
	bytecode []byte
	address  *types.Address
}

// NewContract creates contract by abi and deployed address
func NewContract(clientHelper sdk.SignableRpcCaller, abiJSON []byte, address *types.Address) (*Contract, error) {
	if clientHelper == nil {
		return nil, errors.New("client is nessary")
	}

	var abi abi.ABI
	err := abi.UnmarshalJSON([]byte(abiJSON))
	if err != nil {
		return nil, errors.Wrap(err, "failed unmarshal ABI")
	}

	contract := &Contract{clientHelper, abi, nil, address}
	return contract, nil
}

// DeployContract deploys a contract by abiJSON, bytecode and consturctor params.
// It returns a ContractDeployState instance which contains 3 channels for notifying when state changed.
func DeployContract(client sdk.SignableRpcCaller, option *types.ContractDeployOption, abiJSON []byte,
	bytecode []byte, constroctorParams ...interface{}) *sdk.ContractDeployResult {

	doneChan := make(chan struct{})
	result := sdk.ContractDeployResult{DoneChannel: doneChan}

	go func() {

		defer func() {
			doneChan <- struct{}{}
			close(doneChan)
		}()

		//generate ABI
		var abi abi.ABI
		err := abi.UnmarshalJSON([]byte(abiJSON))
		if err != nil {
			result.Error = errors.Errorf("failed to unmarshal ABI: %+v", abiJSON)
			return
		}

		tx := new(types.UnsignedTransaction)
		if option != nil {
			tx.UnsignedTransactionBase = types.UnsignedTransactionBase(option.UnsignedTransactionBase)
		}

		//recreate contract bytecode with consturctor params
		if len(constroctorParams) > 0 {
			input, err := abi.Pack("", constroctorParams...)
			if err != nil {
				result.Error = errors.Wrapf(err, "failed to encode constructor with args %+v", constroctorParams)
				return
			}

			bytecode = append(bytecode, input...)
		}
		tx.Data = bytecode

		//deploy contract
		txhash, err := client.SignTransactionAndSend(*tx)
		if err != nil {
			result.Error = errors.Wrapf(err, "failed to send transaction, tx = %+v", tx)
			return
		}
		result.TransactionHash = &txhash

		// timeout := time.After(time.Duration(_timeoutIns) * time.Second)
		timeout := time.After(3600 * time.Second)
		if option != nil && option.Timeout != 0 {
			timeout = time.After(option.Timeout)
		}

		ticker := time.Tick(2000 * time.Millisecond)
		// Keep trying until we're time out or get a result or get an error
		for {
			select {
			// Got a timeout! fail with a timeout error
			case t := <-timeout:
				result.Error = errors.Errorf("deploy contract timeout, time = %v, txhash = %v", t, txhash)
				return
			// Got a tick
			case <-ticker:
				txReceipt, err := client.GetTransactionReceipt(txhash)
				if err != nil {
					result.Error = errors.Wrapf(err, "failed to get transaction receipt by hash %v", txhash)
					return
				}

				if txReceipt == nil {
					continue
				}

				if txReceipt.OutcomeStatus == 1 {
					result.Error = errors.Errorf("transaction execution failed, reason %v, hash = %v", txReceipt.TxExecErrorMsg, txhash)
					return
				}

				result.DeployedContract = &Contract{client, abi, bytecode, txReceipt.ContractCreated}
				return
			}
		}
	}()
	return &result
}

func (contract *Contract) ABI() abi.ABI {
	return contract.abi
}

func (contract *Contract) Address() *types.Address {
	return contract.address
}

func (contract *Contract) Bytecode() []byte {
	return contract.bytecode
}

func (contract *Contract) GetRpcCaller() sdk.SignableRpcCaller {
	return contract.client
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
	packed, err := contract.abi.Pack(method, args...)
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
	callRequest.To = contract.address
	hexData := hexutil.Bytes(data).String()
	callRequest.Data = &hexData
	callRequest.FillByCallOption(option)

	var epoch *types.Epoch = nil
	if option != nil && option.Epoch != nil {
		epoch = option.Epoch
	}
	// fmt.Printf("data: %x,hexdata:%v,callRequest.Data:%v\n", data, hexData, *callRequest.Data)
	resultHexStr, err := contract.client.Call(*callRequest, epoch)
	if err != nil {
		return errors.Wrapf(err, "failed to call %+v at epoch %v", *callRequest, epoch)
	}

	if len(resultHexStr) < 2 {
		return errors.Errorf("call response string %v length smaller than 2", resultHexStr)
	}

	bytes := []byte(resultHexStr)

	err = contract.abi.UnpackIntoInterface(resultPtr, method, bytes)
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
	tx.To = contract.address
	tx.Data = data

	err = contract.client.PopulateTransaction(tx)
	if err != nil {
		return "", errors.Wrap(err, cfxerrors.ErrMsgApplyTxValues)
	}

	txhash, err := contract.client.SignTransactionAndSend(*tx)
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
	if contract.address != nil {
		var err error
		*addressPtr, _, err = contract.address.ToCommon()
		if err != nil {
			return errors.Wrap(err, "failed to parse contract address to hex address")
		}
	}

	boundContract := bind.NewBoundContract(*addressPtr, contract.abi, nil, nil, nil)
	err := boundContract.UnpackLog(out, event, eLog)
	if err != nil {
		return errors.Wrapf(err, "failed to unpack log")
	}

	return nil
}
