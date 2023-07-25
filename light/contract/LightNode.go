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

// ILightNodeState is an auto generated low-level Go binding around an user-defined struct.
type ILightNodeState struct {
	Epoch                *big.Int
	Round                *big.Int
	EarliestBlockNumber  *big.Int
	FinalizedBlockNumber *big.Int
	Blocks               *big.Int
	MaxBlocks            *big.Int
}

// LedgerInfoLibDecision is an auto generated low-level Go binding around an user-defined struct.
type LedgerInfoLibDecision struct {
	BlockHash [32]byte
	Height    uint64
}

// LedgerInfoLibEpochState is an auto generated low-level Go binding around an user-defined struct.
type LedgerInfoLibEpochState struct {
	Epoch             uint64
	Validators        []LedgerInfoLibValidatorInfo
	QuorumVotingPower uint64
	TotalVotingPower  uint64
	VrfSeed           []byte
}

// LedgerInfoLibLedgerInfoWithSignatures is an auto generated low-level Go binding around an user-defined struct.
type LedgerInfoLibLedgerInfoWithSignatures struct {
	Epoch               uint64
	Round               uint64
	Id                  [32]byte
	ExecutedStateId     [32]byte
	Version             uint64
	TimestampUsecs      uint64
	NextEpochState      LedgerInfoLibEpochState
	Pivot               LedgerInfoLibDecision
	ConsensusDataHash   [32]byte
	Accounts            [][32]byte
	AggregatedSignature []byte
}

// LedgerInfoLibValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type LedgerInfoLibValidatorInfo struct {
	Account               [32]byte
	CompressedPublicKey   []byte
	UncompressedPublicKey []byte
	VrfPublicKey          []byte
	VotingPower           uint64
}

// ProofLibNibblePath is an auto generated low-level Go binding around an user-defined struct.
type ProofLibNibblePath struct {
	Nibbles [32]byte
	Start   *big.Int
	End     *big.Int
}

// ProofLibProofNode is an auto generated low-level Go binding around an user-defined struct.
type ProofLibProofNode struct {
	Path     ProofLibNibblePath
	Children [16][32]byte
	Value    []byte
}

// TypesReceiptProof is an auto generated low-level Go binding around an user-defined struct.
type TypesReceiptProof struct {
	Headers      [][]byte
	BlockIndex   []byte
	BlockProof   []ProofLibProofNode
	ReceiptsRoot [32]byte
	Index        []byte
	Receipt      []byte
	ReceiptProof []ProofLibProofNode
}

// LightNodeMetaData contains all meta data concerning the LightNode contract.
var LightNodeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"clientState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"headerHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ledgerInfoUtil\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mptVerify\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"account\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"compressedPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"uncompressedPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"vrfPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"votingPower\",\"type\":\"uint64\"}],\"internalType\":\"structLedgerInfoLib.ValidatorInfo[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"quorumVotingPower\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"totalVotingPower\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"vrfSeed\",\"type\":\"bytes\"}],\"internalType\":\"structLedgerInfoLib.EpochState\",\"name\":\"committee\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"executedStateId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestampUsecs\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"account\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"compressedPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"uncompressedPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"vrfPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"votingPower\",\"type\":\"uint64\"}],\"internalType\":\"structLedgerInfoLib.ValidatorInfo[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"quorumVotingPower\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"totalVotingPower\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"vrfSeed\",\"type\":\"bytes\"}],\"internalType\":\"structLedgerInfoLib.EpochState\",\"name\":\"nextEpochState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"internalType\":\"structLedgerInfoLib.Decision\",\"name\":\"pivot\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"consensusDataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"accounts\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"aggregatedSignature\",\"type\":\"bytes\"}],\"internalType\":\"structLedgerInfoLib.LedgerInfoWithSignatures\",\"name\":\"ledgerInfo\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"name\":\"nearestPivot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"executedStateId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestampUsecs\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"account\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"compressedPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"uncompressedPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"vrfPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"votingPower\",\"type\":\"uint64\"}],\"internalType\":\"structLedgerInfoLib.ValidatorInfo[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"quorumVotingPower\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"totalVotingPower\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"vrfSeed\",\"type\":\"bytes\"}],\"internalType\":\"structLedgerInfoLib.EpochState\",\"name\":\"nextEpochState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"internalType\":\"structLedgerInfoLib.Decision\",\"name\":\"pivot\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"consensusDataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"accounts\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"aggregatedSignature\",\"type\":\"bytes\"}],\"internalType\":\"structLedgerInfoLib.LedgerInfoWithSignatures\",\"name\":\"ledgerInfo\",\"type\":\"tuple\"}],\"name\":\"relayPOS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"headers\",\"type\":\"bytes[]\"}],\"name\":\"relayPOW\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"removeBlockHeader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earliestBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalizedBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxBlocks\",\"type\":\"uint256\"}],\"internalType\":\"structILightNode.State\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_blockHeader\",\"type\":\"bytes\"}],\"name\":\"updateBlockHeader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"updateLightClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifiableHeaderRange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_receiptProof\",\"type\":\"bytes\"}],\"name\":\"verifyProofData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"logs\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"headers\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"blockIndex\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"nibbles\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"internalType\":\"structProofLib.NibblePath\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[16]\",\"name\":\"children\",\"type\":\"bytes32[16]\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structProofLib.ProofNode[]\",\"name\":\"blockProof\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"receiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"index\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receipt\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"nibbles\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"internalType\":\"structProofLib.NibblePath\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[16]\",\"name\":\"children\",\"type\":\"bytes32[16]\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structProofLib.ProofNode[]\",\"name\":\"receiptProof\",\"type\":\"tuple[]\"}],\"internalType\":\"structTypes.ReceiptProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"verifyReceiptProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LightNodeABI is the input ABI used to generate the binding from.
// Deprecated: Use LightNodeMetaData.ABI instead.
var LightNodeABI = LightNodeMetaData.ABI

// LightNode is an auto generated Go binding around an Ethereum contract.
type LightNode struct {
	LightNodeCaller     // Read-only binding to the contract
	LightNodeTransactor // Write-only binding to the contract
	LightNodeFilterer   // Log filterer for contract events
}

// LightNodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type LightNodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightNodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LightNodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightNodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LightNodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightNodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LightNodeSession struct {
	Contract     *LightNode        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LightNodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LightNodeCallerSession struct {
	Contract *LightNodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// LightNodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LightNodeTransactorSession struct {
	Contract     *LightNodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// LightNodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type LightNodeRaw struct {
	Contract *LightNode // Generic contract binding to access the raw methods on
}

// LightNodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LightNodeCallerRaw struct {
	Contract *LightNodeCaller // Generic read-only contract binding to access the raw methods on
}

// LightNodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LightNodeTransactorRaw struct {
	Contract *LightNodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLightNode creates a new instance of LightNode, bound to a specific deployed contract.
func NewLightNode(address common.Address, backend bind.ContractBackend) (*LightNode, error) {
	contract, err := bindLightNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LightNode{LightNodeCaller: LightNodeCaller{contract: contract}, LightNodeTransactor: LightNodeTransactor{contract: contract}, LightNodeFilterer: LightNodeFilterer{contract: contract}}, nil
}

// NewLightNodeCaller creates a new read-only instance of LightNode, bound to a specific deployed contract.
func NewLightNodeCaller(address common.Address, caller bind.ContractCaller) (*LightNodeCaller, error) {
	contract, err := bindLightNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LightNodeCaller{contract: contract}, nil
}

// NewLightNodeTransactor creates a new write-only instance of LightNode, bound to a specific deployed contract.
func NewLightNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*LightNodeTransactor, error) {
	contract, err := bindLightNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LightNodeTransactor{contract: contract}, nil
}

// NewLightNodeFilterer creates a new log filterer instance of LightNode, bound to a specific deployed contract.
func NewLightNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*LightNodeFilterer, error) {
	contract, err := bindLightNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LightNodeFilterer{contract: contract}, nil
}

// bindLightNode binds a generic wrapper to an already deployed contract.
func bindLightNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LightNodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LightNode *LightNodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LightNode.Contract.LightNodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LightNode *LightNodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightNode.Contract.LightNodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LightNode *LightNodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LightNode.Contract.LightNodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LightNode *LightNodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LightNode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LightNode *LightNodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightNode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LightNode *LightNodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LightNode.Contract.contract.Transact(opts, method, params...)
}

// ClientState is a free data retrieval call binding the contract method 0xbd3ce6b0.
//
// Solidity: function clientState() view returns(bytes)
func (_LightNode *LightNodeCaller) ClientState(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "clientState")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ClientState is a free data retrieval call binding the contract method 0xbd3ce6b0.
//
// Solidity: function clientState() view returns(bytes)
func (_LightNode *LightNodeSession) ClientState() ([]byte, error) {
	return _LightNode.Contract.ClientState(&_LightNode.CallOpts)
}

// ClientState is a free data retrieval call binding the contract method 0xbd3ce6b0.
//
// Solidity: function clientState() view returns(bytes)
func (_LightNode *LightNodeCallerSession) ClientState() ([]byte, error) {
	return _LightNode.Contract.ClientState(&_LightNode.CallOpts)
}

// HeaderHeight is a free data retrieval call binding the contract method 0xbdb6dead.
//
// Solidity: function headerHeight() view returns(uint256 height)
func (_LightNode *LightNodeCaller) HeaderHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "headerHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HeaderHeight is a free data retrieval call binding the contract method 0xbdb6dead.
//
// Solidity: function headerHeight() view returns(uint256 height)
func (_LightNode *LightNodeSession) HeaderHeight() (*big.Int, error) {
	return _LightNode.Contract.HeaderHeight(&_LightNode.CallOpts)
}

// HeaderHeight is a free data retrieval call binding the contract method 0xbdb6dead.
//
// Solidity: function headerHeight() view returns(uint256 height)
func (_LightNode *LightNodeCallerSession) HeaderHeight() (*big.Int, error) {
	return _LightNode.Contract.HeaderHeight(&_LightNode.CallOpts)
}

// NearestPivot is a free data retrieval call binding the contract method 0x823002e6.
//
// Solidity: function nearestPivot(uint256 height) view returns(uint256)
func (_LightNode *LightNodeCaller) NearestPivot(opts *bind.CallOpts, height *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "nearestPivot", height)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NearestPivot is a free data retrieval call binding the contract method 0x823002e6.
//
// Solidity: function nearestPivot(uint256 height) view returns(uint256)
func (_LightNode *LightNodeSession) NearestPivot(height *big.Int) (*big.Int, error) {
	return _LightNode.Contract.NearestPivot(&_LightNode.CallOpts, height)
}

// NearestPivot is a free data retrieval call binding the contract method 0x823002e6.
//
// Solidity: function nearestPivot(uint256 height) view returns(uint256)
func (_LightNode *LightNodeCallerSession) NearestPivot(height *big.Int) (*big.Int, error) {
	return _LightNode.Contract.NearestPivot(&_LightNode.CallOpts, height)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_LightNode *LightNodeCaller) State(opts *bind.CallOpts) (ILightNodeState, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "state")

	if err != nil {
		return *new(ILightNodeState), err
	}

	out0 := *abi.ConvertType(out[0], new(ILightNodeState)).(*ILightNodeState)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_LightNode *LightNodeSession) State() (ILightNodeState, error) {
	return _LightNode.Contract.State(&_LightNode.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_LightNode *LightNodeCallerSession) State() (ILightNodeState, error) {
	return _LightNode.Contract.State(&_LightNode.CallOpts)
}

// VerifiableHeaderRange is a free data retrieval call binding the contract method 0xe492ebbc.
//
// Solidity: function verifiableHeaderRange() view returns(uint256, uint256)
func (_LightNode *LightNodeCaller) VerifiableHeaderRange(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "verifiableHeaderRange")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// VerifiableHeaderRange is a free data retrieval call binding the contract method 0xe492ebbc.
//
// Solidity: function verifiableHeaderRange() view returns(uint256, uint256)
func (_LightNode *LightNodeSession) VerifiableHeaderRange() (*big.Int, *big.Int, error) {
	return _LightNode.Contract.VerifiableHeaderRange(&_LightNode.CallOpts)
}

// VerifiableHeaderRange is a free data retrieval call binding the contract method 0xe492ebbc.
//
// Solidity: function verifiableHeaderRange() view returns(uint256, uint256)
func (_LightNode *LightNodeCallerSession) VerifiableHeaderRange() (*big.Int, *big.Int, error) {
	return _LightNode.Contract.VerifiableHeaderRange(&_LightNode.CallOpts)
}

// VerifyProofData is a free data retrieval call binding the contract method 0x16dc5270.
//
// Solidity: function verifyProofData(bytes _receiptProof) view returns(bool success, string message, bytes logs)
func (_LightNode *LightNodeCaller) VerifyProofData(opts *bind.CallOpts, _receiptProof []byte) (struct {
	Success bool
	Message string
	Logs    []byte
}, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "verifyProofData", _receiptProof)

	outstruct := new(struct {
		Success bool
		Message string
		Logs    []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Success = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Message = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Logs = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// VerifyProofData is a free data retrieval call binding the contract method 0x16dc5270.
//
// Solidity: function verifyProofData(bytes _receiptProof) view returns(bool success, string message, bytes logs)
func (_LightNode *LightNodeSession) VerifyProofData(_receiptProof []byte) (struct {
	Success bool
	Message string
	Logs    []byte
}, error) {
	return _LightNode.Contract.VerifyProofData(&_LightNode.CallOpts, _receiptProof)
}

// VerifyProofData is a free data retrieval call binding the contract method 0x16dc5270.
//
// Solidity: function verifyProofData(bytes _receiptProof) view returns(bool success, string message, bytes logs)
func (_LightNode *LightNodeCallerSession) VerifyProofData(_receiptProof []byte) (struct {
	Success bool
	Message string
	Logs    []byte
}, error) {
	return _LightNode.Contract.VerifyProofData(&_LightNode.CallOpts, _receiptProof)
}

// VerifyReceiptProof is a free data retrieval call binding the contract method 0xc97777a8.
//
// Solidity: function verifyReceiptProof((bytes[],bytes,((bytes32,uint256,uint256),bytes32[16],bytes)[],bytes32,bytes,bytes,((bytes32,uint256,uint256),bytes32[16],bytes)[]) proof) view returns(bool)
func (_LightNode *LightNodeCaller) VerifyReceiptProof(opts *bind.CallOpts, proof TypesReceiptProof) (bool, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "verifyReceiptProof", proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyReceiptProof is a free data retrieval call binding the contract method 0xc97777a8.
//
// Solidity: function verifyReceiptProof((bytes[],bytes,((bytes32,uint256,uint256),bytes32[16],bytes)[],bytes32,bytes,bytes,((bytes32,uint256,uint256),bytes32[16],bytes)[]) proof) view returns(bool)
func (_LightNode *LightNodeSession) VerifyReceiptProof(proof TypesReceiptProof) (bool, error) {
	return _LightNode.Contract.VerifyReceiptProof(&_LightNode.CallOpts, proof)
}

// VerifyReceiptProof is a free data retrieval call binding the contract method 0xc97777a8.
//
// Solidity: function verifyReceiptProof((bytes[],bytes,((bytes32,uint256,uint256),bytes32[16],bytes)[],bytes32,bytes,bytes,((bytes32,uint256,uint256),bytes32[16],bytes)[]) proof) view returns(bool)
func (_LightNode *LightNodeCallerSession) VerifyReceiptProof(proof TypesReceiptProof) (bool, error) {
	return _LightNode.Contract.VerifyReceiptProof(&_LightNode.CallOpts, proof)
}

// Initialize is a paid mutator transaction binding the contract method 0x435d298a.
//
// Solidity: function initialize(address controller, address ledgerInfoUtil, address mptVerify, (uint64,(bytes32,bytes,bytes,bytes,uint64)[],uint64,uint64,bytes) committee, (uint64,uint64,bytes32,bytes32,uint64,uint64,(uint64,(bytes32,bytes,bytes,bytes,uint64)[],uint64,uint64,bytes),(bytes32,uint64),bytes32,bytes32[],bytes) ledgerInfo) returns()
func (_LightNode *LightNodeTransactor) Initialize(opts *bind.TransactOpts, controller common.Address, ledgerInfoUtil common.Address, mptVerify common.Address, committee LedgerInfoLibEpochState, ledgerInfo LedgerInfoLibLedgerInfoWithSignatures) (*types.Transaction, error) {
	return _LightNode.contract.Transact(opts, "initialize", controller, ledgerInfoUtil, mptVerify, committee, ledgerInfo)
}

// Initialize is a paid mutator transaction binding the contract method 0x435d298a.
//
// Solidity: function initialize(address controller, address ledgerInfoUtil, address mptVerify, (uint64,(bytes32,bytes,bytes,bytes,uint64)[],uint64,uint64,bytes) committee, (uint64,uint64,bytes32,bytes32,uint64,uint64,(uint64,(bytes32,bytes,bytes,bytes,uint64)[],uint64,uint64,bytes),(bytes32,uint64),bytes32,bytes32[],bytes) ledgerInfo) returns()
func (_LightNode *LightNodeSession) Initialize(controller common.Address, ledgerInfoUtil common.Address, mptVerify common.Address, committee LedgerInfoLibEpochState, ledgerInfo LedgerInfoLibLedgerInfoWithSignatures) (*types.Transaction, error) {
	return _LightNode.Contract.Initialize(&_LightNode.TransactOpts, controller, ledgerInfoUtil, mptVerify, committee, ledgerInfo)
}

// Initialize is a paid mutator transaction binding the contract method 0x435d298a.
//
// Solidity: function initialize(address controller, address ledgerInfoUtil, address mptVerify, (uint64,(bytes32,bytes,bytes,bytes,uint64)[],uint64,uint64,bytes) committee, (uint64,uint64,bytes32,bytes32,uint64,uint64,(uint64,(bytes32,bytes,bytes,bytes,uint64)[],uint64,uint64,bytes),(bytes32,uint64),bytes32,bytes32[],bytes) ledgerInfo) returns()
func (_LightNode *LightNodeTransactorSession) Initialize(controller common.Address, ledgerInfoUtil common.Address, mptVerify common.Address, committee LedgerInfoLibEpochState, ledgerInfo LedgerInfoLibLedgerInfoWithSignatures) (*types.Transaction, error) {
	return _LightNode.Contract.Initialize(&_LightNode.TransactOpts, controller, ledgerInfoUtil, mptVerify, committee, ledgerInfo)
}

// RelayPOS is a paid mutator transaction binding the contract method 0x76450598.
//
// Solidity: function relayPOS((uint64,uint64,bytes32,bytes32,uint64,uint64,(uint64,(bytes32,bytes,bytes,bytes,uint64)[],uint64,uint64,bytes),(bytes32,uint64),bytes32,bytes32[],bytes) ledgerInfo) returns()
func (_LightNode *LightNodeTransactor) RelayPOS(opts *bind.TransactOpts, ledgerInfo LedgerInfoLibLedgerInfoWithSignatures) (*types.Transaction, error) {
	return _LightNode.contract.Transact(opts, "relayPOS", ledgerInfo)
}

// RelayPOS is a paid mutator transaction binding the contract method 0x76450598.
//
// Solidity: function relayPOS((uint64,uint64,bytes32,bytes32,uint64,uint64,(uint64,(bytes32,bytes,bytes,bytes,uint64)[],uint64,uint64,bytes),(bytes32,uint64),bytes32,bytes32[],bytes) ledgerInfo) returns()
func (_LightNode *LightNodeSession) RelayPOS(ledgerInfo LedgerInfoLibLedgerInfoWithSignatures) (*types.Transaction, error) {
	return _LightNode.Contract.RelayPOS(&_LightNode.TransactOpts, ledgerInfo)
}

// RelayPOS is a paid mutator transaction binding the contract method 0x76450598.
//
// Solidity: function relayPOS((uint64,uint64,bytes32,bytes32,uint64,uint64,(uint64,(bytes32,bytes,bytes,bytes,uint64)[],uint64,uint64,bytes),(bytes32,uint64),bytes32,bytes32[],bytes) ledgerInfo) returns()
func (_LightNode *LightNodeTransactorSession) RelayPOS(ledgerInfo LedgerInfoLibLedgerInfoWithSignatures) (*types.Transaction, error) {
	return _LightNode.Contract.RelayPOS(&_LightNode.TransactOpts, ledgerInfo)
}

// RelayPOW is a paid mutator transaction binding the contract method 0x16b3b684.
//
// Solidity: function relayPOW(bytes[] headers) returns()
func (_LightNode *LightNodeTransactor) RelayPOW(opts *bind.TransactOpts, headers [][]byte) (*types.Transaction, error) {
	return _LightNode.contract.Transact(opts, "relayPOW", headers)
}

// RelayPOW is a paid mutator transaction binding the contract method 0x16b3b684.
//
// Solidity: function relayPOW(bytes[] headers) returns()
func (_LightNode *LightNodeSession) RelayPOW(headers [][]byte) (*types.Transaction, error) {
	return _LightNode.Contract.RelayPOW(&_LightNode.TransactOpts, headers)
}

// RelayPOW is a paid mutator transaction binding the contract method 0x16b3b684.
//
// Solidity: function relayPOW(bytes[] headers) returns()
func (_LightNode *LightNodeTransactorSession) RelayPOW(headers [][]byte) (*types.Transaction, error) {
	return _LightNode.Contract.RelayPOW(&_LightNode.TransactOpts, headers)
}

// RemoveBlockHeader is a paid mutator transaction binding the contract method 0x8113e308.
//
// Solidity: function removeBlockHeader(uint256 limit) returns()
func (_LightNode *LightNodeTransactor) RemoveBlockHeader(opts *bind.TransactOpts, limit *big.Int) (*types.Transaction, error) {
	return _LightNode.contract.Transact(opts, "removeBlockHeader", limit)
}

// RemoveBlockHeader is a paid mutator transaction binding the contract method 0x8113e308.
//
// Solidity: function removeBlockHeader(uint256 limit) returns()
func (_LightNode *LightNodeSession) RemoveBlockHeader(limit *big.Int) (*types.Transaction, error) {
	return _LightNode.Contract.RemoveBlockHeader(&_LightNode.TransactOpts, limit)
}

// RemoveBlockHeader is a paid mutator transaction binding the contract method 0x8113e308.
//
// Solidity: function removeBlockHeader(uint256 limit) returns()
func (_LightNode *LightNodeTransactorSession) RemoveBlockHeader(limit *big.Int) (*types.Transaction, error) {
	return _LightNode.Contract.RemoveBlockHeader(&_LightNode.TransactOpts, limit)
}

// UpdateBlockHeader is a paid mutator transaction binding the contract method 0xd240f3cf.
//
// Solidity: function updateBlockHeader(bytes _blockHeader) returns()
func (_LightNode *LightNodeTransactor) UpdateBlockHeader(opts *bind.TransactOpts, _blockHeader []byte) (*types.Transaction, error) {
	return _LightNode.contract.Transact(opts, "updateBlockHeader", _blockHeader)
}

// UpdateBlockHeader is a paid mutator transaction binding the contract method 0xd240f3cf.
//
// Solidity: function updateBlockHeader(bytes _blockHeader) returns()
func (_LightNode *LightNodeSession) UpdateBlockHeader(_blockHeader []byte) (*types.Transaction, error) {
	return _LightNode.Contract.UpdateBlockHeader(&_LightNode.TransactOpts, _blockHeader)
}

// UpdateBlockHeader is a paid mutator transaction binding the contract method 0xd240f3cf.
//
// Solidity: function updateBlockHeader(bytes _blockHeader) returns()
func (_LightNode *LightNodeTransactorSession) UpdateBlockHeader(_blockHeader []byte) (*types.Transaction, error) {
	return _LightNode.Contract.UpdateBlockHeader(&_LightNode.TransactOpts, _blockHeader)
}

// UpdateLightClient is a paid mutator transaction binding the contract method 0x330115fc.
//
// Solidity: function updateLightClient(bytes _data) returns()
func (_LightNode *LightNodeTransactor) UpdateLightClient(opts *bind.TransactOpts, _data []byte) (*types.Transaction, error) {
	return _LightNode.contract.Transact(opts, "updateLightClient", _data)
}

// UpdateLightClient is a paid mutator transaction binding the contract method 0x330115fc.
//
// Solidity: function updateLightClient(bytes _data) returns()
func (_LightNode *LightNodeSession) UpdateLightClient(_data []byte) (*types.Transaction, error) {
	return _LightNode.Contract.UpdateLightClient(&_LightNode.TransactOpts, _data)
}

// UpdateLightClient is a paid mutator transaction binding the contract method 0x330115fc.
//
// Solidity: function updateLightClient(bytes _data) returns()
func (_LightNode *LightNodeTransactorSession) UpdateLightClient(_data []byte) (*types.Transaction, error) {
	return _LightNode.Contract.UpdateLightClient(&_LightNode.TransactOpts, _data)
}
