// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ProvableMetaData contains all meta data concerning the Provable contract.
var ProvableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"valueHash\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"nibbles\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"internalType\":\"structProofLib.NibblePath\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[16]\",\"name\":\"children\",\"type\":\"bytes32[16]\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structProofLib.ProofNode[]\",\"name\":\"nodes\",\"type\":\"tuple[]\"}],\"name\":\"prove\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blockIndex\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"nibbles\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"internalType\":\"structProofLib.NibblePath\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[16]\",\"name\":\"children\",\"type\":\"bytes32[16]\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structProofLib.ProofNode[]\",\"name\":\"blockProof\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"receiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"index\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"accumulatedGasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"gasSponsorPaid\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"logBloom\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"topics\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"space\",\"type\":\"uint8\"}],\"internalType\":\"structTypes.TxLog[]\",\"name\":\"logs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8\",\"name\":\"outcomeStatus\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"storageSponsorPaid\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"collaterals\",\"type\":\"uint64\"}],\"internalType\":\"structTypes.StorageChange[]\",\"name\":\"storageCollateralized\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"collaterals\",\"type\":\"uint64\"}],\"internalType\":\"structTypes.StorageChange[]\",\"name\":\"storageReleased\",\"type\":\"tuple[]\"}],\"internalType\":\"structTypes.TxReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"nibbles\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"internalType\":\"structProofLib.NibblePath\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[16]\",\"name\":\"children\",\"type\":\"bytes32[16]\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structProofLib.ProofNode[]\",\"name\":\"receiptProof\",\"type\":\"tuple[]\"}],\"name\":\"proveReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// ProvableABI is the input ABI used to generate the binding from.
// Deprecated: Use ProvableMetaData.ABI instead.
var ProvableABI = ProvableMetaData.ABI

// Provable is an auto generated Go binding around an Ethereum contract.
type Provable struct {
	ProvableCaller     // Read-only binding to the contract
	ProvableTransactor // Write-only binding to the contract
	ProvableFilterer   // Log filterer for contract events
}

// ProvableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProvableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProvableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProvableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProvableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProvableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProvableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProvableSession struct {
	Contract     *Provable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProvableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProvableCallerSession struct {
	Contract *ProvableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ProvableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProvableTransactorSession struct {
	Contract     *ProvableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ProvableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProvableRaw struct {
	Contract *Provable // Generic contract binding to access the raw methods on
}

// ProvableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProvableCallerRaw struct {
	Contract *ProvableCaller // Generic read-only contract binding to access the raw methods on
}

// ProvableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProvableTransactorRaw struct {
	Contract *ProvableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProvable creates a new instance of Provable, bound to a specific deployed contract.
func NewProvable(address common.Address, backend bind.ContractBackend) (*Provable, error) {
	contract, err := bindProvable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Provable{ProvableCaller: ProvableCaller{contract: contract}, ProvableTransactor: ProvableTransactor{contract: contract}, ProvableFilterer: ProvableFilterer{contract: contract}}, nil
}

// NewProvableCaller creates a new read-only instance of Provable, bound to a specific deployed contract.
func NewProvableCaller(address common.Address, caller bind.ContractCaller) (*ProvableCaller, error) {
	contract, err := bindProvable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProvableCaller{contract: contract}, nil
}

// NewProvableTransactor creates a new write-only instance of Provable, bound to a specific deployed contract.
func NewProvableTransactor(address common.Address, transactor bind.ContractTransactor) (*ProvableTransactor, error) {
	contract, err := bindProvable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProvableTransactor{contract: contract}, nil
}

// NewProvableFilterer creates a new log filterer instance of Provable, bound to a specific deployed contract.
func NewProvableFilterer(address common.Address, filterer bind.ContractFilterer) (*ProvableFilterer, error) {
	contract, err := bindProvable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProvableFilterer{contract: contract}, nil
}

// bindProvable binds a generic wrapper to an already deployed contract.
func bindProvable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProvableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Provable *ProvableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Provable.Contract.ProvableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Provable *ProvableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Provable.Contract.ProvableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Provable *ProvableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Provable.Contract.ProvableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Provable *ProvableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Provable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Provable *ProvableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Provable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Provable *ProvableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Provable.Contract.contract.Transact(opts, method, params...)
}

// Prove is a free data retrieval call binding the contract method 0x283bf72c.
//
// Solidity: function prove(bytes32 root, bytes key, bytes32 valueHash, ((bytes32,uint256,uint256),bytes32[16],bytes)[] nodes) pure returns(bool)
func (_Provable *ProvableCaller) Prove(opts *bind.CallOpts, root [32]byte, key []byte, valueHash [32]byte, nodes []ProofLibProofNode) (bool, error) {
	var out []interface{}
	err := _Provable.contract.Call(opts, &out, "prove", root, key, valueHash, nodes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Prove is a free data retrieval call binding the contract method 0x283bf72c.
//
// Solidity: function prove(bytes32 root, bytes key, bytes32 valueHash, ((bytes32,uint256,uint256),bytes32[16],bytes)[] nodes) pure returns(bool)
func (_Provable *ProvableSession) Prove(root [32]byte, key []byte, valueHash [32]byte, nodes []ProofLibProofNode) (bool, error) {
	return _Provable.Contract.Prove(&_Provable.CallOpts, root, key, valueHash, nodes)
}

// Prove is a free data retrieval call binding the contract method 0x283bf72c.
//
// Solidity: function prove(bytes32 root, bytes key, bytes32 valueHash, ((bytes32,uint256,uint256),bytes32[16],bytes)[] nodes) pure returns(bool)
func (_Provable *ProvableCallerSession) Prove(root [32]byte, key []byte, valueHash [32]byte, nodes []ProofLibProofNode) (bool, error) {
	return _Provable.Contract.Prove(&_Provable.CallOpts, root, key, valueHash, nodes)
}

// ProveReceipt is a free data retrieval call binding the contract method 0xcad91694.
//
// Solidity: function proveReceipt(bytes32 blockRoot, bytes blockIndex, ((bytes32,uint256,uint256),bytes32[16],bytes)[] blockProof, bytes32 receiptsRoot, bytes index, (uint256,uint256,bool,bytes,(address,bytes32[],bytes,uint8)[],uint8,bool,(address,uint64)[],(address,uint64)[]) receipt, ((bytes32,uint256,uint256),bytes32[16],bytes)[] receiptProof) pure returns(bool)
func (_Provable *ProvableCaller) ProveReceipt(opts *bind.CallOpts, blockRoot [32]byte, blockIndex []byte, blockProof []ProofLibProofNode, receiptsRoot [32]byte, index []byte, receipt TypesTxReceipt, receiptProof []ProofLibProofNode) (bool, error) {
	var out []interface{}
	err := _Provable.contract.Call(opts, &out, "proveReceipt", blockRoot, blockIndex, blockProof, receiptsRoot, index, receipt, receiptProof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProveReceipt is a free data retrieval call binding the contract method 0xcad91694.
//
// Solidity: function proveReceipt(bytes32 blockRoot, bytes blockIndex, ((bytes32,uint256,uint256),bytes32[16],bytes)[] blockProof, bytes32 receiptsRoot, bytes index, (uint256,uint256,bool,bytes,(address,bytes32[],bytes,uint8)[],uint8,bool,(address,uint64)[],(address,uint64)[]) receipt, ((bytes32,uint256,uint256),bytes32[16],bytes)[] receiptProof) pure returns(bool)
func (_Provable *ProvableSession) ProveReceipt(blockRoot [32]byte, blockIndex []byte, blockProof []ProofLibProofNode, receiptsRoot [32]byte, index []byte, receipt TypesTxReceipt, receiptProof []ProofLibProofNode) (bool, error) {
	return _Provable.Contract.ProveReceipt(&_Provable.CallOpts, blockRoot, blockIndex, blockProof, receiptsRoot, index, receipt, receiptProof)
}

// ProveReceipt is a free data retrieval call binding the contract method 0xcad91694.
//
// Solidity: function proveReceipt(bytes32 blockRoot, bytes blockIndex, ((bytes32,uint256,uint256),bytes32[16],bytes)[] blockProof, bytes32 receiptsRoot, bytes index, (uint256,uint256,bool,bytes,(address,bytes32[],bytes,uint8)[],uint8,bool,(address,uint64)[],(address,uint64)[]) receipt, ((bytes32,uint256,uint256),bytes32[16],bytes)[] receiptProof) pure returns(bool)
func (_Provable *ProvableCallerSession) ProveReceipt(blockRoot [32]byte, blockIndex []byte, blockProof []ProofLibProofNode, receiptsRoot [32]byte, index []byte, receipt TypesTxReceipt, receiptProof []ProofLibProofNode) (bool, error) {
	return _Provable.Contract.ProveReceipt(&_Provable.CallOpts, blockRoot, blockIndex, blockProof, receiptsRoot, index, receipt, receiptProof)
}
