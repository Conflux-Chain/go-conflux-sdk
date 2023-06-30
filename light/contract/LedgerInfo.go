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

// LedgerInfoMetaData contains all meta data concerning the LedgerInfo contract.
var LedgerInfoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"}],\"name\":\"batchVerifyBLS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"executedStateId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestampUsecs\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"account\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"compressedPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"uncompressedPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"vrfPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"votingPower\",\"type\":\"uint64\"}],\"internalType\":\"structLedgerInfoLib.ValidatorInfo[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"quorumVotingPower\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"totalVotingPower\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"vrfSeed\",\"type\":\"bytes\"}],\"internalType\":\"structLedgerInfoLib.EpochState\",\"name\":\"nextEpochState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"internalType\":\"structLedgerInfoLib.Decision\",\"name\":\"pivot\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"consensusDataHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"account\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"consensusSignature\",\"type\":\"bytes\"}],\"internalType\":\"structLedgerInfoLib.AccountSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structLedgerInfoLib.LedgerInfoWithSignatures\",\"name\":\"ledgerInfo\",\"type\":\"tuple\"}],\"name\":\"bcsEncode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"precompile\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"outputLen\",\"type\":\"uint256\"}],\"name\":\"callPrecompile\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"expandMessageXmd\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"hashToCurve\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"hashToField\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"verifyBLS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"g2Message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"verifyBLSHashed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LedgerInfoABI is the input ABI used to generate the binding from.
// Deprecated: Use LedgerInfoMetaData.ABI instead.
var LedgerInfoABI = LedgerInfoMetaData.ABI

// LedgerInfo is an auto generated Go binding around an Ethereum contract.
type LedgerInfo struct {
	LedgerInfoCaller     // Read-only binding to the contract
	LedgerInfoTransactor // Write-only binding to the contract
	LedgerInfoFilterer   // Log filterer for contract events
}

// LedgerInfoCaller is an auto generated read-only Go binding around an Ethereum contract.
type LedgerInfoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LedgerInfoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LedgerInfoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LedgerInfoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LedgerInfoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LedgerInfoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LedgerInfoSession struct {
	Contract     *LedgerInfo       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LedgerInfoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LedgerInfoCallerSession struct {
	Contract *LedgerInfoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// LedgerInfoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LedgerInfoTransactorSession struct {
	Contract     *LedgerInfoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LedgerInfoRaw is an auto generated low-level Go binding around an Ethereum contract.
type LedgerInfoRaw struct {
	Contract *LedgerInfo // Generic contract binding to access the raw methods on
}

// LedgerInfoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LedgerInfoCallerRaw struct {
	Contract *LedgerInfoCaller // Generic read-only contract binding to access the raw methods on
}

// LedgerInfoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LedgerInfoTransactorRaw struct {
	Contract *LedgerInfoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLedgerInfo creates a new instance of LedgerInfo, bound to a specific deployed contract.
func NewLedgerInfo(address common.Address, backend bind.ContractBackend) (*LedgerInfo, error) {
	contract, err := bindLedgerInfo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LedgerInfo{LedgerInfoCaller: LedgerInfoCaller{contract: contract}, LedgerInfoTransactor: LedgerInfoTransactor{contract: contract}, LedgerInfoFilterer: LedgerInfoFilterer{contract: contract}}, nil
}

// NewLedgerInfoCaller creates a new read-only instance of LedgerInfo, bound to a specific deployed contract.
func NewLedgerInfoCaller(address common.Address, caller bind.ContractCaller) (*LedgerInfoCaller, error) {
	contract, err := bindLedgerInfo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LedgerInfoCaller{contract: contract}, nil
}

// NewLedgerInfoTransactor creates a new write-only instance of LedgerInfo, bound to a specific deployed contract.
func NewLedgerInfoTransactor(address common.Address, transactor bind.ContractTransactor) (*LedgerInfoTransactor, error) {
	contract, err := bindLedgerInfo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LedgerInfoTransactor{contract: contract}, nil
}

// NewLedgerInfoFilterer creates a new log filterer instance of LedgerInfo, bound to a specific deployed contract.
func NewLedgerInfoFilterer(address common.Address, filterer bind.ContractFilterer) (*LedgerInfoFilterer, error) {
	contract, err := bindLedgerInfo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LedgerInfoFilterer{contract: contract}, nil
}

// bindLedgerInfo binds a generic wrapper to an already deployed contract.
func bindLedgerInfo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LedgerInfoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LedgerInfo *LedgerInfoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LedgerInfo.Contract.LedgerInfoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LedgerInfo *LedgerInfoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LedgerInfo.Contract.LedgerInfoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LedgerInfo *LedgerInfoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LedgerInfo.Contract.LedgerInfoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LedgerInfo *LedgerInfoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LedgerInfo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LedgerInfo *LedgerInfoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LedgerInfo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LedgerInfo *LedgerInfoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LedgerInfo.Contract.contract.Transact(opts, method, params...)
}

// BatchVerifyBLS is a free data retrieval call binding the contract method 0xdd003b5f.
//
// Solidity: function batchVerifyBLS(bytes[] signatures, bytes message, bytes[] publicKeys) view returns(bool)
func (_LedgerInfo *LedgerInfoCaller) BatchVerifyBLS(opts *bind.CallOpts, signatures [][]byte, message []byte, publicKeys [][]byte) (bool, error) {
	var out []interface{}
	err := _LedgerInfo.contract.Call(opts, &out, "batchVerifyBLS", signatures, message, publicKeys)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BatchVerifyBLS is a free data retrieval call binding the contract method 0xdd003b5f.
//
// Solidity: function batchVerifyBLS(bytes[] signatures, bytes message, bytes[] publicKeys) view returns(bool)
func (_LedgerInfo *LedgerInfoSession) BatchVerifyBLS(signatures [][]byte, message []byte, publicKeys [][]byte) (bool, error) {
	return _LedgerInfo.Contract.BatchVerifyBLS(&_LedgerInfo.CallOpts, signatures, message, publicKeys)
}

// BatchVerifyBLS is a free data retrieval call binding the contract method 0xdd003b5f.
//
// Solidity: function batchVerifyBLS(bytes[] signatures, bytes message, bytes[] publicKeys) view returns(bool)
func (_LedgerInfo *LedgerInfoCallerSession) BatchVerifyBLS(signatures [][]byte, message []byte, publicKeys [][]byte) (bool, error) {
	return _LedgerInfo.Contract.BatchVerifyBLS(&_LedgerInfo.CallOpts, signatures, message, publicKeys)
}

// BcsEncode is a free data retrieval call binding the contract method 0x838eeceb.
//
// Solidity: function bcsEncode((uint64,uint64,bytes32,bytes32,uint64,uint64,(uint64,(bytes32,bytes,bytes,bytes,uint64)[],uint64,uint64,bytes),(bytes32,uint64),bytes32,(bytes32,bytes)[]) ledgerInfo) pure returns(bytes)
func (_LedgerInfo *LedgerInfoCaller) BcsEncode(opts *bind.CallOpts, ledgerInfo LedgerInfoLibLedgerInfoWithSignatures) ([]byte, error) {
	var out []interface{}
	err := _LedgerInfo.contract.Call(opts, &out, "bcsEncode", ledgerInfo)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// BcsEncode is a free data retrieval call binding the contract method 0x838eeceb.
//
// Solidity: function bcsEncode((uint64,uint64,bytes32,bytes32,uint64,uint64,(uint64,(bytes32,bytes,bytes,bytes,uint64)[],uint64,uint64,bytes),(bytes32,uint64),bytes32,(bytes32,bytes)[]) ledgerInfo) pure returns(bytes)
func (_LedgerInfo *LedgerInfoSession) BcsEncode(ledgerInfo LedgerInfoLibLedgerInfoWithSignatures) ([]byte, error) {
	return _LedgerInfo.Contract.BcsEncode(&_LedgerInfo.CallOpts, ledgerInfo)
}

// BcsEncode is a free data retrieval call binding the contract method 0x838eeceb.
//
// Solidity: function bcsEncode((uint64,uint64,bytes32,bytes32,uint64,uint64,(uint64,(bytes32,bytes,bytes,bytes,uint64)[],uint64,uint64,bytes),(bytes32,uint64),bytes32,(bytes32,bytes)[]) ledgerInfo) pure returns(bytes)
func (_LedgerInfo *LedgerInfoCallerSession) BcsEncode(ledgerInfo LedgerInfoLibLedgerInfoWithSignatures) ([]byte, error) {
	return _LedgerInfo.Contract.BcsEncode(&_LedgerInfo.CallOpts, ledgerInfo)
}

// CallPrecompile is a free data retrieval call binding the contract method 0x2c283b15.
//
// Solidity: function callPrecompile(address precompile, bytes input, uint256 outputLen) view returns(bytes)
func (_LedgerInfo *LedgerInfoCaller) CallPrecompile(opts *bind.CallOpts, precompile common.Address, input []byte, outputLen *big.Int) ([]byte, error) {
	var out []interface{}
	err := _LedgerInfo.contract.Call(opts, &out, "callPrecompile", precompile, input, outputLen)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CallPrecompile is a free data retrieval call binding the contract method 0x2c283b15.
//
// Solidity: function callPrecompile(address precompile, bytes input, uint256 outputLen) view returns(bytes)
func (_LedgerInfo *LedgerInfoSession) CallPrecompile(precompile common.Address, input []byte, outputLen *big.Int) ([]byte, error) {
	return _LedgerInfo.Contract.CallPrecompile(&_LedgerInfo.CallOpts, precompile, input, outputLen)
}

// CallPrecompile is a free data retrieval call binding the contract method 0x2c283b15.
//
// Solidity: function callPrecompile(address precompile, bytes input, uint256 outputLen) view returns(bytes)
func (_LedgerInfo *LedgerInfoCallerSession) CallPrecompile(precompile common.Address, input []byte, outputLen *big.Int) ([]byte, error) {
	return _LedgerInfo.Contract.CallPrecompile(&_LedgerInfo.CallOpts, precompile, input, outputLen)
}

// ExpandMessageXmd is a free data retrieval call binding the contract method 0xcb6415ad.
//
// Solidity: function expandMessageXmd(bytes message) pure returns(bytes)
func (_LedgerInfo *LedgerInfoCaller) ExpandMessageXmd(opts *bind.CallOpts, message []byte) ([]byte, error) {
	var out []interface{}
	err := _LedgerInfo.contract.Call(opts, &out, "expandMessageXmd", message)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ExpandMessageXmd is a free data retrieval call binding the contract method 0xcb6415ad.
//
// Solidity: function expandMessageXmd(bytes message) pure returns(bytes)
func (_LedgerInfo *LedgerInfoSession) ExpandMessageXmd(message []byte) ([]byte, error) {
	return _LedgerInfo.Contract.ExpandMessageXmd(&_LedgerInfo.CallOpts, message)
}

// ExpandMessageXmd is a free data retrieval call binding the contract method 0xcb6415ad.
//
// Solidity: function expandMessageXmd(bytes message) pure returns(bytes)
func (_LedgerInfo *LedgerInfoCallerSession) ExpandMessageXmd(message []byte) ([]byte, error) {
	return _LedgerInfo.Contract.ExpandMessageXmd(&_LedgerInfo.CallOpts, message)
}

// HashToCurve is a free data retrieval call binding the contract method 0x95961df8.
//
// Solidity: function hashToCurve(bytes message) view returns(bytes)
func (_LedgerInfo *LedgerInfoCaller) HashToCurve(opts *bind.CallOpts, message []byte) ([]byte, error) {
	var out []interface{}
	err := _LedgerInfo.contract.Call(opts, &out, "hashToCurve", message)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// HashToCurve is a free data retrieval call binding the contract method 0x95961df8.
//
// Solidity: function hashToCurve(bytes message) view returns(bytes)
func (_LedgerInfo *LedgerInfoSession) HashToCurve(message []byte) ([]byte, error) {
	return _LedgerInfo.Contract.HashToCurve(&_LedgerInfo.CallOpts, message)
}

// HashToCurve is a free data retrieval call binding the contract method 0x95961df8.
//
// Solidity: function hashToCurve(bytes message) view returns(bytes)
func (_LedgerInfo *LedgerInfoCallerSession) HashToCurve(message []byte) ([]byte, error) {
	return _LedgerInfo.Contract.HashToCurve(&_LedgerInfo.CallOpts, message)
}

// HashToField is a free data retrieval call binding the contract method 0x1c5490f2.
//
// Solidity: function hashToField(bytes message) view returns(bytes)
func (_LedgerInfo *LedgerInfoCaller) HashToField(opts *bind.CallOpts, message []byte) ([]byte, error) {
	var out []interface{}
	err := _LedgerInfo.contract.Call(opts, &out, "hashToField", message)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// HashToField is a free data retrieval call binding the contract method 0x1c5490f2.
//
// Solidity: function hashToField(bytes message) view returns(bytes)
func (_LedgerInfo *LedgerInfoSession) HashToField(message []byte) ([]byte, error) {
	return _LedgerInfo.Contract.HashToField(&_LedgerInfo.CallOpts, message)
}

// HashToField is a free data retrieval call binding the contract method 0x1c5490f2.
//
// Solidity: function hashToField(bytes message) view returns(bytes)
func (_LedgerInfo *LedgerInfoCallerSession) HashToField(message []byte) ([]byte, error) {
	return _LedgerInfo.Contract.HashToField(&_LedgerInfo.CallOpts, message)
}

// VerifyBLS is a free data retrieval call binding the contract method 0xeb5f906b.
//
// Solidity: function verifyBLS(bytes signature, bytes message, bytes publicKey) view returns(bool)
func (_LedgerInfo *LedgerInfoCaller) VerifyBLS(opts *bind.CallOpts, signature []byte, message []byte, publicKey []byte) (bool, error) {
	var out []interface{}
	err := _LedgerInfo.contract.Call(opts, &out, "verifyBLS", signature, message, publicKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyBLS is a free data retrieval call binding the contract method 0xeb5f906b.
//
// Solidity: function verifyBLS(bytes signature, bytes message, bytes publicKey) view returns(bool)
func (_LedgerInfo *LedgerInfoSession) VerifyBLS(signature []byte, message []byte, publicKey []byte) (bool, error) {
	return _LedgerInfo.Contract.VerifyBLS(&_LedgerInfo.CallOpts, signature, message, publicKey)
}

// VerifyBLS is a free data retrieval call binding the contract method 0xeb5f906b.
//
// Solidity: function verifyBLS(bytes signature, bytes message, bytes publicKey) view returns(bool)
func (_LedgerInfo *LedgerInfoCallerSession) VerifyBLS(signature []byte, message []byte, publicKey []byte) (bool, error) {
	return _LedgerInfo.Contract.VerifyBLS(&_LedgerInfo.CallOpts, signature, message, publicKey)
}

// VerifyBLSHashed is a free data retrieval call binding the contract method 0xac1cab32.
//
// Solidity: function verifyBLSHashed(bytes signature, bytes g2Message, bytes publicKey) view returns(bool)
func (_LedgerInfo *LedgerInfoCaller) VerifyBLSHashed(opts *bind.CallOpts, signature []byte, g2Message []byte, publicKey []byte) (bool, error) {
	var out []interface{}
	err := _LedgerInfo.contract.Call(opts, &out, "verifyBLSHashed", signature, g2Message, publicKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyBLSHashed is a free data retrieval call binding the contract method 0xac1cab32.
//
// Solidity: function verifyBLSHashed(bytes signature, bytes g2Message, bytes publicKey) view returns(bool)
func (_LedgerInfo *LedgerInfoSession) VerifyBLSHashed(signature []byte, g2Message []byte, publicKey []byte) (bool, error) {
	return _LedgerInfo.Contract.VerifyBLSHashed(&_LedgerInfo.CallOpts, signature, g2Message, publicKey)
}

// VerifyBLSHashed is a free data retrieval call binding the contract method 0xac1cab32.
//
// Solidity: function verifyBLSHashed(bytes signature, bytes g2Message, bytes publicKey) view returns(bool)
func (_LedgerInfo *LedgerInfoCallerSession) VerifyBLSHashed(signature []byte, g2Message []byte, publicKey []byte) (bool, error) {
	return _LedgerInfo.Contract.VerifyBLSHashed(&_LedgerInfo.CallOpts, signature, g2Message, publicKey)
}
