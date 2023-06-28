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

// ILightNodeClientState is an auto generated low-level Go binding around an user-defined struct.
type ILightNodeClientState struct {
	Epoch                *big.Int
	Round                *big.Int
	EarliestBlockNumber  *big.Int
	FinalizedBlockNumber *big.Int
	Blocks               *big.Int
	MaxBlocks            *big.Int
}

// LedgerInfoLibAccountSignature is an auto generated low-level Go binding around an user-defined struct.
type LedgerInfoLibAccountSignature struct {
	Account            [32]byte
	ConsensusSignature []byte
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
	Epoch             uint64
	Round             uint64
	Id                [32]byte
	ExecutedStateId   [32]byte
	Version           uint64
	TimestampUsecs    uint64
	NextEpochState    LedgerInfoLibEpochState
	Pivot             LedgerInfoLibDecision
	ConsensusDataHash [32]byte
	Signatures        []LedgerInfoLibAccountSignature
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

// TypesBlockHeader is an auto generated low-level Go binding around an user-defined struct.
type TypesBlockHeader struct {
	ParentHash            [32]byte
	Height                *big.Int
	Timestamp             *big.Int
	Author                common.Address
	TransactionsRoot      [32]byte
	DeferredStateRoot     [32]byte
	DeferredReceiptsRoot  [32]byte
	DeferredLogsBloomHash [32]byte
	Blame                 *big.Int
	Difficulty            *big.Int
	Adaptive              bool
	GasLimit              *big.Int
	RefereeHashes         [][32]byte
	Custom                [][]byte
	Nonce                 *big.Int
	PosReference          [32]byte
}

// TypesReceiptProof is an auto generated low-level Go binding around an user-defined struct.
type TypesReceiptProof struct {
	Headers      []TypesBlockHeader
	BlockIndex   []byte
	BlockProof   []ProofLibProofNode
	ReceiptsRoot [32]byte
	Index        []byte
	Receipt      TypesTxReceipt
	ReceiptProof []ProofLibProofNode
}

// TypesStorageChange is an auto generated low-level Go binding around an user-defined struct.
type TypesStorageChange struct {
	Account     common.Address
	Collaterals uint64
}

// TypesTxLog is an auto generated low-level Go binding around an user-defined struct.
type TypesTxLog struct {
	Addr   common.Address
	Topics [][32]byte
	Data   []byte
	Space  uint8
}

// TypesTxReceipt is an auto generated low-level Go binding around an user-defined struct.
type TypesTxReceipt struct {
	AccumulatedGasUsed    *big.Int
	GasFee                *big.Int
	GasSponsorPaid        bool
	LogBloom              []byte
	Logs                  []TypesTxLog
	OutcomeStatus         uint8
	StorageSponsorPaid    bool
	StorageCollateralized []TypesStorageChange
	StorageReleased       []TypesStorageChange
}

// LightNodeMetaData contains all meta data concerning the LightNode contract.
var LightNodeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"UpdateBlockHeader\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"name\":\"UpdateLightClient\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"clientState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earliestBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalizedBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxBlocks\",\"type\":\"uint256\"}],\"internalType\":\"structILightNode.ClientState\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ledgerInfoUtil\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_mptVerify\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"executedStateId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestampUsecs\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"account\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"compressedPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"uncompressedPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"vrfPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"votingPower\",\"type\":\"uint64\"}],\"internalType\":\"structLedgerInfoLib.ValidatorInfo[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"quorumVotingPower\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"totalVotingPower\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"vrfSeed\",\"type\":\"bytes\"}],\"internalType\":\"structLedgerInfoLib.EpochState\",\"name\":\"nextEpochState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"internalType\":\"structLedgerInfoLib.Decision\",\"name\":\"pivot\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"consensusDataHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"account\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"consensusSignature\",\"type\":\"bytes\"}],\"internalType\":\"structLedgerInfoLib.AccountSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structLedgerInfoLib.LedgerInfoWithSignatures\",\"name\":\"ledgerInfo\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"name\":\"nearestPivot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"removeBlockHeader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"deferredStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"deferredReceiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"deferredLogsBloomHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blame\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"adaptive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"refereeHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes[]\",\"name\":\"custom\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"posReference\",\"type\":\"bytes32\"}],\"internalType\":\"structTypes.BlockHeader[]\",\"name\":\"headers\",\"type\":\"tuple[]\"}],\"name\":\"updateBlockHeader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"executedStateId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestampUsecs\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"account\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"compressedPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"uncompressedPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"vrfPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"votingPower\",\"type\":\"uint64\"}],\"internalType\":\"structLedgerInfoLib.ValidatorInfo[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"quorumVotingPower\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"totalVotingPower\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"vrfSeed\",\"type\":\"bytes\"}],\"internalType\":\"structLedgerInfoLib.EpochState\",\"name\":\"nextEpochState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"internalType\":\"structLedgerInfoLib.Decision\",\"name\":\"pivot\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"consensusDataHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"account\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"consensusSignature\",\"type\":\"bytes\"}],\"internalType\":\"structLedgerInfoLib.AccountSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structLedgerInfoLib.LedgerInfoWithSignatures\",\"name\":\"ledgerInfo\",\"type\":\"tuple\"}],\"name\":\"updateLightClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifiableHeaderRange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"receiptProof\",\"type\":\"bytes\"}],\"name\":\"verifyProofData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"rlpLogs\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"deferredStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"deferredReceiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"deferredLogsBloomHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blame\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"adaptive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"refereeHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes[]\",\"name\":\"custom\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"posReference\",\"type\":\"bytes32\"}],\"internalType\":\"structTypes.BlockHeader[]\",\"name\":\"headers\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"blockIndex\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"nibbles\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"internalType\":\"structProofLib.NibblePath\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[16]\",\"name\":\"children\",\"type\":\"bytes32[16]\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structProofLib.ProofNode[]\",\"name\":\"blockProof\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"receiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"index\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"accumulatedGasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"gasSponsorPaid\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"logBloom\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"topics\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"space\",\"type\":\"uint8\"}],\"internalType\":\"structTypes.TxLog[]\",\"name\":\"logs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8\",\"name\":\"outcomeStatus\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"storageSponsorPaid\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"collaterals\",\"type\":\"uint64\"}],\"internalType\":\"structTypes.StorageChange[]\",\"name\":\"storageCollateralized\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"collaterals\",\"type\":\"uint64\"}],\"internalType\":\"structTypes.StorageChange[]\",\"name\":\"storageReleased\",\"type\":\"tuple[]\"}],\"internalType\":\"structTypes.TxReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"nibbles\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"internalType\":\"structProofLib.NibblePath\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[16]\",\"name\":\"children\",\"type\":\"bytes32[16]\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structProofLib.ProofNode[]\",\"name\":\"receiptProof\",\"type\":\"tuple[]\"}],\"internalType\":\"structTypes.ReceiptProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"verifyReceiptProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"topics\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"space\",\"type\":\"uint8\"}],\"internalType\":\"structTypes.TxLog[]\",\"name\":\"logs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
// Solidity: function clientState() view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_LightNode *LightNodeCaller) ClientState(opts *bind.CallOpts) (ILightNodeClientState, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "clientState")

	if err != nil {
		return *new(ILightNodeClientState), err
	}

	out0 := *abi.ConvertType(out[0], new(ILightNodeClientState)).(*ILightNodeClientState)

	return out0, err

}

// ClientState is a free data retrieval call binding the contract method 0xbd3ce6b0.
//
// Solidity: function clientState() view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_LightNode *LightNodeSession) ClientState() (ILightNodeClientState, error) {
	return _LightNode.Contract.ClientState(&_LightNode.CallOpts)
}

// ClientState is a free data retrieval call binding the contract method 0xbd3ce6b0.
//
// Solidity: function clientState() view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_LightNode *LightNodeCallerSession) ClientState() (ILightNodeClientState, error) {
	return _LightNode.Contract.ClientState(&_LightNode.CallOpts)
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
// Solidity: function verifyProofData(bytes receiptProof) view returns(bool success, string message, bytes rlpLogs)
func (_LightNode *LightNodeCaller) VerifyProofData(opts *bind.CallOpts, receiptProof []byte) (struct {
	Success bool
	Message string
	RlpLogs []byte
}, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "verifyProofData", receiptProof)

	outstruct := new(struct {
		Success bool
		Message string
		RlpLogs []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Success = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Message = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.RlpLogs = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// VerifyProofData is a free data retrieval call binding the contract method 0x16dc5270.
//
// Solidity: function verifyProofData(bytes receiptProof) view returns(bool success, string message, bytes rlpLogs)
func (_LightNode *LightNodeSession) VerifyProofData(receiptProof []byte) (struct {
	Success bool
	Message string
	RlpLogs []byte
}, error) {
	return _LightNode.Contract.VerifyProofData(&_LightNode.CallOpts, receiptProof)
}

// VerifyProofData is a free data retrieval call binding the contract method 0x16dc5270.
//
// Solidity: function verifyProofData(bytes receiptProof) view returns(bool success, string message, bytes rlpLogs)
func (_LightNode *LightNodeCallerSession) VerifyProofData(receiptProof []byte) (struct {
	Success bool
	Message string
	RlpLogs []byte
}, error) {
	return _LightNode.Contract.VerifyProofData(&_LightNode.CallOpts, receiptProof)
}

// VerifyReceiptProof is a free data retrieval call binding the contract method 0x25f0b353.
//
// Solidity: function verifyReceiptProof(((bytes32,uint256,uint256,address,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bool,uint256,bytes32[],bytes[],uint256,bytes32)[],bytes,((bytes32,uint256,uint256),bytes32[16],bytes)[],bytes32,bytes,(uint256,uint256,bool,bytes,(address,bytes32[],bytes,uint8)[],uint8,bool,(address,uint64)[],(address,uint64)[]),((bytes32,uint256,uint256),bytes32[16],bytes)[]) proof) view returns(bool success, (address,bytes32[],bytes,uint8)[] logs)
func (_LightNode *LightNodeCaller) VerifyReceiptProof(opts *bind.CallOpts, proof TypesReceiptProof) (struct {
	Success bool
	Logs    []TypesTxLog
}, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "verifyReceiptProof", proof)

	outstruct := new(struct {
		Success bool
		Logs    []TypesTxLog
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Success = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Logs = *abi.ConvertType(out[1], new([]TypesTxLog)).(*[]TypesTxLog)

	return *outstruct, err

}

// VerifyReceiptProof is a free data retrieval call binding the contract method 0x25f0b353.
//
// Solidity: function verifyReceiptProof(((bytes32,uint256,uint256,address,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bool,uint256,bytes32[],bytes[],uint256,bytes32)[],bytes,((bytes32,uint256,uint256),bytes32[16],bytes)[],bytes32,bytes,(uint256,uint256,bool,bytes,(address,bytes32[],bytes,uint8)[],uint8,bool,(address,uint64)[],(address,uint64)[]),((bytes32,uint256,uint256),bytes32[16],bytes)[]) proof) view returns(bool success, (address,bytes32[],bytes,uint8)[] logs)
func (_LightNode *LightNodeSession) VerifyReceiptProof(proof TypesReceiptProof) (struct {
	Success bool
	Logs    []TypesTxLog
}, error) {
	return _LightNode.Contract.VerifyReceiptProof(&_LightNode.CallOpts, proof)
}

// VerifyReceiptProof is a free data retrieval call binding the contract method 0x25f0b353.
//
// Solidity: function verifyReceiptProof(((bytes32,uint256,uint256,address,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bool,uint256,bytes32[],bytes[],uint256,bytes32)[],bytes,((bytes32,uint256,uint256),bytes32[16],bytes)[],bytes32,bytes,(uint256,uint256,bool,bytes,(address,bytes32[],bytes,uint8)[],uint8,bool,(address,uint64)[],(address,uint64)[]),((bytes32,uint256,uint256),bytes32[16],bytes)[]) proof) view returns(bool success, (address,bytes32[],bytes,uint8)[] logs)
func (_LightNode *LightNodeCallerSession) VerifyReceiptProof(proof TypesReceiptProof) (struct {
	Success bool
	Logs    []TypesTxLog
}, error) {
	return _LightNode.Contract.VerifyReceiptProof(&_LightNode.CallOpts, proof)
}

// Initialize is a paid mutator transaction binding the contract method 0xf3129c45.
//
// Solidity: function initialize(address _controller, address _ledgerInfoUtil, address _mptVerify, (uint64,uint64,bytes32,bytes32,uint64,uint64,(uint64,(bytes32,bytes,bytes,bytes,uint64)[],uint64,uint64,bytes),(bytes32,uint64),bytes32,(bytes32,bytes)[]) ledgerInfo) returns()
func (_LightNode *LightNodeTransactor) Initialize(opts *bind.TransactOpts, _controller common.Address, _ledgerInfoUtil common.Address, _mptVerify common.Address, ledgerInfo LedgerInfoLibLedgerInfoWithSignatures) (*types.Transaction, error) {
	return _LightNode.contract.Transact(opts, "initialize", _controller, _ledgerInfoUtil, _mptVerify, ledgerInfo)
}

// Initialize is a paid mutator transaction binding the contract method 0xf3129c45.
//
// Solidity: function initialize(address _controller, address _ledgerInfoUtil, address _mptVerify, (uint64,uint64,bytes32,bytes32,uint64,uint64,(uint64,(bytes32,bytes,bytes,bytes,uint64)[],uint64,uint64,bytes),(bytes32,uint64),bytes32,(bytes32,bytes)[]) ledgerInfo) returns()
func (_LightNode *LightNodeSession) Initialize(_controller common.Address, _ledgerInfoUtil common.Address, _mptVerify common.Address, ledgerInfo LedgerInfoLibLedgerInfoWithSignatures) (*types.Transaction, error) {
	return _LightNode.Contract.Initialize(&_LightNode.TransactOpts, _controller, _ledgerInfoUtil, _mptVerify, ledgerInfo)
}

// Initialize is a paid mutator transaction binding the contract method 0xf3129c45.
//
// Solidity: function initialize(address _controller, address _ledgerInfoUtil, address _mptVerify, (uint64,uint64,bytes32,bytes32,uint64,uint64,(uint64,(bytes32,bytes,bytes,bytes,uint64)[],uint64,uint64,bytes),(bytes32,uint64),bytes32,(bytes32,bytes)[]) ledgerInfo) returns()
func (_LightNode *LightNodeTransactorSession) Initialize(_controller common.Address, _ledgerInfoUtil common.Address, _mptVerify common.Address, ledgerInfo LedgerInfoLibLedgerInfoWithSignatures) (*types.Transaction, error) {
	return _LightNode.Contract.Initialize(&_LightNode.TransactOpts, _controller, _ledgerInfoUtil, _mptVerify, ledgerInfo)
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

// UpdateBlockHeader is a paid mutator transaction binding the contract method 0xbb9548b9.
//
// Solidity: function updateBlockHeader((bytes32,uint256,uint256,address,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bool,uint256,bytes32[],bytes[],uint256,bytes32)[] headers) returns()
func (_LightNode *LightNodeTransactor) UpdateBlockHeader(opts *bind.TransactOpts, headers []TypesBlockHeader) (*types.Transaction, error) {
	return _LightNode.contract.Transact(opts, "updateBlockHeader", headers)
}

// UpdateBlockHeader is a paid mutator transaction binding the contract method 0xbb9548b9.
//
// Solidity: function updateBlockHeader((bytes32,uint256,uint256,address,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bool,uint256,bytes32[],bytes[],uint256,bytes32)[] headers) returns()
func (_LightNode *LightNodeSession) UpdateBlockHeader(headers []TypesBlockHeader) (*types.Transaction, error) {
	return _LightNode.Contract.UpdateBlockHeader(&_LightNode.TransactOpts, headers)
}

// UpdateBlockHeader is a paid mutator transaction binding the contract method 0xbb9548b9.
//
// Solidity: function updateBlockHeader((bytes32,uint256,uint256,address,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bool,uint256,bytes32[],bytes[],uint256,bytes32)[] headers) returns()
func (_LightNode *LightNodeTransactorSession) UpdateBlockHeader(headers []TypesBlockHeader) (*types.Transaction, error) {
	return _LightNode.Contract.UpdateBlockHeader(&_LightNode.TransactOpts, headers)
}

// UpdateLightClient is a paid mutator transaction binding the contract method 0xf548a7c4.
//
// Solidity: function updateLightClient((uint64,uint64,bytes32,bytes32,uint64,uint64,(uint64,(bytes32,bytes,bytes,bytes,uint64)[],uint64,uint64,bytes),(bytes32,uint64),bytes32,(bytes32,bytes)[]) ledgerInfo) returns()
func (_LightNode *LightNodeTransactor) UpdateLightClient(opts *bind.TransactOpts, ledgerInfo LedgerInfoLibLedgerInfoWithSignatures) (*types.Transaction, error) {
	return _LightNode.contract.Transact(opts, "updateLightClient", ledgerInfo)
}

// UpdateLightClient is a paid mutator transaction binding the contract method 0xf548a7c4.
//
// Solidity: function updateLightClient((uint64,uint64,bytes32,bytes32,uint64,uint64,(uint64,(bytes32,bytes,bytes,bytes,uint64)[],uint64,uint64,bytes),(bytes32,uint64),bytes32,(bytes32,bytes)[]) ledgerInfo) returns()
func (_LightNode *LightNodeSession) UpdateLightClient(ledgerInfo LedgerInfoLibLedgerInfoWithSignatures) (*types.Transaction, error) {
	return _LightNode.Contract.UpdateLightClient(&_LightNode.TransactOpts, ledgerInfo)
}

// UpdateLightClient is a paid mutator transaction binding the contract method 0xf548a7c4.
//
// Solidity: function updateLightClient((uint64,uint64,bytes32,bytes32,uint64,uint64,(uint64,(bytes32,bytes,bytes,bytes,uint64)[],uint64,uint64,bytes),(bytes32,uint64),bytes32,(bytes32,bytes)[]) ledgerInfo) returns()
func (_LightNode *LightNodeTransactorSession) UpdateLightClient(ledgerInfo LedgerInfoLibLedgerInfoWithSignatures) (*types.Transaction, error) {
	return _LightNode.Contract.UpdateLightClient(&_LightNode.TransactOpts, ledgerInfo)
}

// LightNodeUpdateBlockHeaderIterator is returned from FilterUpdateBlockHeader and is used to iterate over the raw logs and unpacked data for UpdateBlockHeader events raised by the LightNode contract.
type LightNodeUpdateBlockHeaderIterator struct {
	Event *LightNodeUpdateBlockHeader // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightNodeUpdateBlockHeaderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightNodeUpdateBlockHeader)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightNodeUpdateBlockHeader)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightNodeUpdateBlockHeaderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightNodeUpdateBlockHeaderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightNodeUpdateBlockHeader represents a UpdateBlockHeader event raised by the LightNode contract.
type LightNodeUpdateBlockHeader struct {
	Account common.Address
	Start   *big.Int
	End     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateBlockHeader is a free log retrieval operation binding the contract event 0x0aa20f5b41aa94d6cf4d4335aa85c09fb6e11a1f67d45d65b629a8f814d1e019.
//
// Solidity: event UpdateBlockHeader(address indexed account, uint256 start, uint256 end)
func (_LightNode *LightNodeFilterer) FilterUpdateBlockHeader(opts *bind.FilterOpts, account []common.Address) (*LightNodeUpdateBlockHeaderIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LightNode.contract.FilterLogs(opts, "UpdateBlockHeader", accountRule)
	if err != nil {
		return nil, err
	}
	return &LightNodeUpdateBlockHeaderIterator{contract: _LightNode.contract, event: "UpdateBlockHeader", logs: logs, sub: sub}, nil
}

// WatchUpdateBlockHeader is a free log subscription operation binding the contract event 0x0aa20f5b41aa94d6cf4d4335aa85c09fb6e11a1f67d45d65b629a8f814d1e019.
//
// Solidity: event UpdateBlockHeader(address indexed account, uint256 start, uint256 end)
func (_LightNode *LightNodeFilterer) WatchUpdateBlockHeader(opts *bind.WatchOpts, sink chan<- *LightNodeUpdateBlockHeader, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LightNode.contract.WatchLogs(opts, "UpdateBlockHeader", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightNodeUpdateBlockHeader)
				if err := _LightNode.contract.UnpackLog(event, "UpdateBlockHeader", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateBlockHeader is a log parse operation binding the contract event 0x0aa20f5b41aa94d6cf4d4335aa85c09fb6e11a1f67d45d65b629a8f814d1e019.
//
// Solidity: event UpdateBlockHeader(address indexed account, uint256 start, uint256 end)
func (_LightNode *LightNodeFilterer) ParseUpdateBlockHeader(log types.Log) (*LightNodeUpdateBlockHeader, error) {
	event := new(LightNodeUpdateBlockHeader)
	if err := _LightNode.contract.UnpackLog(event, "UpdateBlockHeader", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightNodeUpdateLightClientIterator is returned from FilterUpdateLightClient and is used to iterate over the raw logs and unpacked data for UpdateLightClient events raised by the LightNode contract.
type LightNodeUpdateLightClientIterator struct {
	Event *LightNodeUpdateLightClient // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightNodeUpdateLightClientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightNodeUpdateLightClient)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightNodeUpdateLightClient)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightNodeUpdateLightClientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightNodeUpdateLightClientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightNodeUpdateLightClient represents a UpdateLightClient event raised by the LightNode contract.
type LightNodeUpdateLightClient struct {
	Account common.Address
	Epoch   *big.Int
	Round   *big.Int
	Height  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateLightClient is a free log retrieval operation binding the contract event 0x60ffca2d519da71d55982e19f5b347dacc637ae2ad5cbcfe7556cdbb851822e1.
//
// Solidity: event UpdateLightClient(address indexed account, uint256 epoch, uint256 round, uint256 height)
func (_LightNode *LightNodeFilterer) FilterUpdateLightClient(opts *bind.FilterOpts, account []common.Address) (*LightNodeUpdateLightClientIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LightNode.contract.FilterLogs(opts, "UpdateLightClient", accountRule)
	if err != nil {
		return nil, err
	}
	return &LightNodeUpdateLightClientIterator{contract: _LightNode.contract, event: "UpdateLightClient", logs: logs, sub: sub}, nil
}

// WatchUpdateLightClient is a free log subscription operation binding the contract event 0x60ffca2d519da71d55982e19f5b347dacc637ae2ad5cbcfe7556cdbb851822e1.
//
// Solidity: event UpdateLightClient(address indexed account, uint256 epoch, uint256 round, uint256 height)
func (_LightNode *LightNodeFilterer) WatchUpdateLightClient(opts *bind.WatchOpts, sink chan<- *LightNodeUpdateLightClient, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LightNode.contract.WatchLogs(opts, "UpdateLightClient", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightNodeUpdateLightClient)
				if err := _LightNode.contract.UnpackLog(event, "UpdateLightClient", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateLightClient is a log parse operation binding the contract event 0x60ffca2d519da71d55982e19f5b347dacc637ae2ad5cbcfe7556cdbb851822e1.
//
// Solidity: event UpdateLightClient(address indexed account, uint256 epoch, uint256 round, uint256 height)
func (_LightNode *LightNodeFilterer) ParseUpdateLightClient(log types.Log) (*LightNodeUpdateLightClient, error) {
	event := new(LightNodeUpdateLightClient)
	if err := _LightNode.contract.UnpackLog(event, "UpdateLightClient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
