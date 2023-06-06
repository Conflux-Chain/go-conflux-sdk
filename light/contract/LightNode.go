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
	Epoch                 *big.Int
	Round                 *big.Int
	EarliestBlockNumber   *big.Int
	FinalizedBlockNumber  *big.Int
	RelayBlockStartNumber *big.Int
	RelayBlockEndNumber   *big.Int
	RelayBlockEndHash     [32]byte
	Blocks                *big.Int
	MaxBlocks             *big.Int
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

// TypesAccountSignature is an auto generated low-level Go binding around an user-defined struct.
type TypesAccountSignature struct {
	Account            [32]byte
	ConsensusSignature []byte
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

// TypesDecision is an auto generated low-level Go binding around an user-defined struct.
type TypesDecision struct {
	BlockHash [32]byte
	Height    uint64
}

// TypesEpochState is an auto generated low-level Go binding around an user-defined struct.
type TypesEpochState struct {
	Epoch             uint64
	Validators        []TypesValidatorInfo
	QuorumVotingPower uint64
	TotalVotingPower  uint64
	VrfSeed           []byte
}

// TypesLedgerInfoWithSignatures is an auto generated low-level Go binding around an user-defined struct.
type TypesLedgerInfoWithSignatures struct {
	Epoch             uint64
	Round             uint64
	Id                [32]byte
	ExecutedStateId   [32]byte
	Version           uint64
	TimestampUsecs    uint64
	NextEpochState    TypesEpochState
	Pivot             TypesDecision
	ConsensusDataHash [32]byte
	Signatures        []TypesAccountSignature
}

// TypesReceiptProof is an auto generated low-level Go binding around an user-defined struct.
type TypesReceiptProof struct {
	EpochNumber  *big.Int
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

// TypesValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type TypesValidatorInfo struct {
	Account      [32]byte
	PublicKey    []byte
	VrfPublicKey []byte
	VotingPower  uint64
}

// LightNodeMetaData contains all meta data concerning the LightNode contract.
var LightNodeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"UpdateBlockHeader\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"UpdateLightClient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFER_EXECUTION_BLOCKS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clientState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earliestBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalizedBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"relayBlockStartNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"relayBlockEndNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"relayBlockEndHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxBlocks\",\"type\":\"uint256\"}],\"internalType\":\"structILightNode.ClientState\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deferredReceiptsRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_mptVerify\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"executedStateId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestampUsecs\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"account\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"vrfPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"votingPower\",\"type\":\"uint64\"}],\"internalType\":\"structTypes.ValidatorInfo[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"quorumVotingPower\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"totalVotingPower\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"vrfSeed\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.EpochState\",\"name\":\"nextEpochState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"internalType\":\"structTypes.Decision\",\"name\":\"pivot\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"consensusDataHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"account\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"consensusSignature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.AccountSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structTypes.LedgerInfoWithSignatures\",\"name\":\"ledgerInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"deferredStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"deferredReceiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"deferredLogsBloomHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blame\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"adaptive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"refereeHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes[]\",\"name\":\"custom\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"posReference\",\"type\":\"bytes32\"}],\"internalType\":\"structTypes.BlockHeader\",\"name\":\"header\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mptVerify\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"removeBlockHeader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setMaxBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"togglePause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"deferredStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"deferredReceiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"deferredLogsBloomHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blame\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"adaptive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"refereeHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes[]\",\"name\":\"custom\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"posReference\",\"type\":\"bytes32\"}],\"internalType\":\"structTypes.BlockHeader[]\",\"name\":\"headers\",\"type\":\"tuple[]\"}],\"name\":\"updateBlockHeader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"executedStateId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestampUsecs\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"account\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"vrfPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"votingPower\",\"type\":\"uint64\"}],\"internalType\":\"structTypes.ValidatorInfo[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"quorumVotingPower\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"totalVotingPower\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"vrfSeed\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.EpochState\",\"name\":\"nextEpochState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"internalType\":\"structTypes.Decision\",\"name\":\"pivot\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"consensusDataHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"account\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"consensusSignature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.AccountSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structTypes.LedgerInfoWithSignatures\",\"name\":\"ledgerInfo\",\"type\":\"tuple\"}],\"name\":\"updateLightClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifiableHeaderRange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"receiptProof\",\"type\":\"bytes\"}],\"name\":\"verifyProofData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"rlpLogs\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"blockIndex\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"nibbles\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"internalType\":\"structProofLib.NibblePath\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[16]\",\"name\":\"children\",\"type\":\"bytes32[16]\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structProofLib.ProofNode[]\",\"name\":\"blockProof\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"receiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"index\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"accumulatedGasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"gasSponsorPaid\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"logBloom\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"topics\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"space\",\"type\":\"uint8\"}],\"internalType\":\"structTypes.TxLog[]\",\"name\":\"logs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8\",\"name\":\"outcomeStatus\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"storageSponsorPaid\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"collaterals\",\"type\":\"uint64\"}],\"internalType\":\"structTypes.StorageChange[]\",\"name\":\"storageCollateralized\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"collaterals\",\"type\":\"uint64\"}],\"internalType\":\"structTypes.StorageChange[]\",\"name\":\"storageReleased\",\"type\":\"tuple[]\"}],\"internalType\":\"structTypes.TxReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"nibbles\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"internalType\":\"structProofLib.NibblePath\",\"name\":\"path\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[16]\",\"name\":\"children\",\"type\":\"bytes32[16]\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structProofLib.ProofNode[]\",\"name\":\"receiptProof\",\"type\":\"tuple[]\"}],\"internalType\":\"structTypes.ReceiptProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"verifyReceiptProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"topics\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"space\",\"type\":\"uint8\"}],\"internalType\":\"structTypes.TxLog[]\",\"name\":\"logs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// DEFEREXECUTIONBLOCKS is a free data retrieval call binding the contract method 0x2d46a4e5.
//
// Solidity: function DEFER_EXECUTION_BLOCKS() view returns(uint256)
func (_LightNode *LightNodeCaller) DEFEREXECUTIONBLOCKS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "DEFER_EXECUTION_BLOCKS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFEREXECUTIONBLOCKS is a free data retrieval call binding the contract method 0x2d46a4e5.
//
// Solidity: function DEFER_EXECUTION_BLOCKS() view returns(uint256)
func (_LightNode *LightNodeSession) DEFEREXECUTIONBLOCKS() (*big.Int, error) {
	return _LightNode.Contract.DEFEREXECUTIONBLOCKS(&_LightNode.CallOpts)
}

// DEFEREXECUTIONBLOCKS is a free data retrieval call binding the contract method 0x2d46a4e5.
//
// Solidity: function DEFER_EXECUTION_BLOCKS() view returns(uint256)
func (_LightNode *LightNodeCallerSession) DEFEREXECUTIONBLOCKS() (*big.Int, error) {
	return _LightNode.Contract.DEFEREXECUTIONBLOCKS(&_LightNode.CallOpts)
}

// ClientState is a free data retrieval call binding the contract method 0xbd3ce6b0.
//
// Solidity: function clientState() view returns((uint256,uint256,uint256,uint256,uint256,uint256,bytes32,uint256,uint256))
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
// Solidity: function clientState() view returns((uint256,uint256,uint256,uint256,uint256,uint256,bytes32,uint256,uint256))
func (_LightNode *LightNodeSession) ClientState() (ILightNodeClientState, error) {
	return _LightNode.Contract.ClientState(&_LightNode.CallOpts)
}

// ClientState is a free data retrieval call binding the contract method 0xbd3ce6b0.
//
// Solidity: function clientState() view returns((uint256,uint256,uint256,uint256,uint256,uint256,bytes32,uint256,uint256))
func (_LightNode *LightNodeCallerSession) ClientState() (ILightNodeClientState, error) {
	return _LightNode.Contract.ClientState(&_LightNode.CallOpts)
}

// DeferredReceiptsRoots is a free data retrieval call binding the contract method 0xb659adfb.
//
// Solidity: function deferredReceiptsRoots(uint256 ) view returns(bytes32)
func (_LightNode *LightNodeCaller) DeferredReceiptsRoots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "deferredReceiptsRoots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DeferredReceiptsRoots is a free data retrieval call binding the contract method 0xb659adfb.
//
// Solidity: function deferredReceiptsRoots(uint256 ) view returns(bytes32)
func (_LightNode *LightNodeSession) DeferredReceiptsRoots(arg0 *big.Int) ([32]byte, error) {
	return _LightNode.Contract.DeferredReceiptsRoots(&_LightNode.CallOpts, arg0)
}

// DeferredReceiptsRoots is a free data retrieval call binding the contract method 0xb659adfb.
//
// Solidity: function deferredReceiptsRoots(uint256 ) view returns(bytes32)
func (_LightNode *LightNodeCallerSession) DeferredReceiptsRoots(arg0 *big.Int) ([32]byte, error) {
	return _LightNode.Contract.DeferredReceiptsRoots(&_LightNode.CallOpts, arg0)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_LightNode *LightNodeCaller) GetAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "getAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_LightNode *LightNodeSession) GetAdmin() (common.Address, error) {
	return _LightNode.Contract.GetAdmin(&_LightNode.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_LightNode *LightNodeCallerSession) GetAdmin() (common.Address, error) {
	return _LightNode.Contract.GetAdmin(&_LightNode.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_LightNode *LightNodeCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_LightNode *LightNodeSession) GetImplementation() (common.Address, error) {
	return _LightNode.Contract.GetImplementation(&_LightNode.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_LightNode *LightNodeCallerSession) GetImplementation() (common.Address, error) {
	return _LightNode.Contract.GetImplementation(&_LightNode.CallOpts)
}

// MptVerify is a free data retrieval call binding the contract method 0x6f30b5d4.
//
// Solidity: function mptVerify() view returns(address)
func (_LightNode *LightNodeCaller) MptVerify(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "mptVerify")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MptVerify is a free data retrieval call binding the contract method 0x6f30b5d4.
//
// Solidity: function mptVerify() view returns(address)
func (_LightNode *LightNodeSession) MptVerify() (common.Address, error) {
	return _LightNode.Contract.MptVerify(&_LightNode.CallOpts)
}

// MptVerify is a free data retrieval call binding the contract method 0x6f30b5d4.
//
// Solidity: function mptVerify() view returns(address)
func (_LightNode *LightNodeCallerSession) MptVerify() (common.Address, error) {
	return _LightNode.Contract.MptVerify(&_LightNode.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LightNode *LightNodeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LightNode *LightNodeSession) Paused() (bool, error) {
	return _LightNode.Contract.Paused(&_LightNode.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LightNode *LightNodeCallerSession) Paused() (bool, error) {
	return _LightNode.Contract.Paused(&_LightNode.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LightNode *LightNodeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LightNode.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LightNode *LightNodeSession) ProxiableUUID() ([32]byte, error) {
	return _LightNode.Contract.ProxiableUUID(&_LightNode.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LightNode *LightNodeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _LightNode.Contract.ProxiableUUID(&_LightNode.CallOpts)
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

// VerifyReceiptProof is a free data retrieval call binding the contract method 0x9445e4ef.
//
// Solidity: function verifyReceiptProof((uint256,bytes,((bytes32,uint256,uint256),bytes32[16],bytes)[],bytes32,bytes,(uint256,uint256,bool,bytes,(address,bytes32[],bytes,uint8)[],uint8,bool,(address,uint64)[],(address,uint64)[]),((bytes32,uint256,uint256),bytes32[16],bytes)[]) proof) view returns(bool success, (address,bytes32[],bytes,uint8)[] logs)
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

// VerifyReceiptProof is a free data retrieval call binding the contract method 0x9445e4ef.
//
// Solidity: function verifyReceiptProof((uint256,bytes,((bytes32,uint256,uint256),bytes32[16],bytes)[],bytes32,bytes,(uint256,uint256,bool,bytes,(address,bytes32[],bytes,uint8)[],uint8,bool,(address,uint64)[],(address,uint64)[]),((bytes32,uint256,uint256),bytes32[16],bytes)[]) proof) view returns(bool success, (address,bytes32[],bytes,uint8)[] logs)
func (_LightNode *LightNodeSession) VerifyReceiptProof(proof TypesReceiptProof) (struct {
	Success bool
	Logs    []TypesTxLog
}, error) {
	return _LightNode.Contract.VerifyReceiptProof(&_LightNode.CallOpts, proof)
}

// VerifyReceiptProof is a free data retrieval call binding the contract method 0x9445e4ef.
//
// Solidity: function verifyReceiptProof((uint256,bytes,((bytes32,uint256,uint256),bytes32[16],bytes)[],bytes32,bytes,(uint256,uint256,bool,bytes,(address,bytes32[],bytes,uint8)[],uint8,bool,(address,uint64)[],(address,uint64)[]),((bytes32,uint256,uint256),bytes32[16],bytes)[]) proof) view returns(bool success, (address,bytes32[],bytes,uint8)[] logs)
func (_LightNode *LightNodeCallerSession) VerifyReceiptProof(proof TypesReceiptProof) (struct {
	Success bool
	Logs    []TypesTxLog
}, error) {
	return _LightNode.Contract.VerifyReceiptProof(&_LightNode.CallOpts, proof)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _admin) returns()
func (_LightNode *LightNodeTransactor) ChangeAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _LightNode.contract.Transact(opts, "changeAdmin", _admin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _admin) returns()
func (_LightNode *LightNodeSession) ChangeAdmin(_admin common.Address) (*types.Transaction, error) {
	return _LightNode.Contract.ChangeAdmin(&_LightNode.TransactOpts, _admin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _admin) returns()
func (_LightNode *LightNodeTransactorSession) ChangeAdmin(_admin common.Address) (*types.Transaction, error) {
	return _LightNode.Contract.ChangeAdmin(&_LightNode.TransactOpts, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0x80171a9d.
//
// Solidity: function initialize(address _controller, address _mptVerify, (uint64,uint64,bytes32,bytes32,uint64,uint64,(uint64,(bytes32,bytes,bytes,uint64)[],uint64,uint64,bytes),(bytes32,uint64),bytes32,(bytes32,bytes)[]) ledgerInfo, (bytes32,uint256,uint256,address,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bool,uint256,bytes32[],bytes[],uint256,bytes32) header) returns()
func (_LightNode *LightNodeTransactor) Initialize(opts *bind.TransactOpts, _controller common.Address, _mptVerify common.Address, ledgerInfo TypesLedgerInfoWithSignatures, header TypesBlockHeader) (*types.Transaction, error) {
	return _LightNode.contract.Transact(opts, "initialize", _controller, _mptVerify, ledgerInfo, header)
}

// Initialize is a paid mutator transaction binding the contract method 0x80171a9d.
//
// Solidity: function initialize(address _controller, address _mptVerify, (uint64,uint64,bytes32,bytes32,uint64,uint64,(uint64,(bytes32,bytes,bytes,uint64)[],uint64,uint64,bytes),(bytes32,uint64),bytes32,(bytes32,bytes)[]) ledgerInfo, (bytes32,uint256,uint256,address,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bool,uint256,bytes32[],bytes[],uint256,bytes32) header) returns()
func (_LightNode *LightNodeSession) Initialize(_controller common.Address, _mptVerify common.Address, ledgerInfo TypesLedgerInfoWithSignatures, header TypesBlockHeader) (*types.Transaction, error) {
	return _LightNode.Contract.Initialize(&_LightNode.TransactOpts, _controller, _mptVerify, ledgerInfo, header)
}

// Initialize is a paid mutator transaction binding the contract method 0x80171a9d.
//
// Solidity: function initialize(address _controller, address _mptVerify, (uint64,uint64,bytes32,bytes32,uint64,uint64,(uint64,(bytes32,bytes,bytes,uint64)[],uint64,uint64,bytes),(bytes32,uint64),bytes32,(bytes32,bytes)[]) ledgerInfo, (bytes32,uint256,uint256,address,bytes32,bytes32,bytes32,bytes32,uint256,uint256,bool,uint256,bytes32[],bytes[],uint256,bytes32) header) returns()
func (_LightNode *LightNodeTransactorSession) Initialize(_controller common.Address, _mptVerify common.Address, ledgerInfo TypesLedgerInfoWithSignatures, header TypesBlockHeader) (*types.Transaction, error) {
	return _LightNode.Contract.Initialize(&_LightNode.TransactOpts, _controller, _mptVerify, ledgerInfo, header)
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

// SetMaxBlocks is a paid mutator transaction binding the contract method 0x5092515e.
//
// Solidity: function setMaxBlocks(uint256 val) returns()
func (_LightNode *LightNodeTransactor) SetMaxBlocks(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _LightNode.contract.Transact(opts, "setMaxBlocks", val)
}

// SetMaxBlocks is a paid mutator transaction binding the contract method 0x5092515e.
//
// Solidity: function setMaxBlocks(uint256 val) returns()
func (_LightNode *LightNodeSession) SetMaxBlocks(val *big.Int) (*types.Transaction, error) {
	return _LightNode.Contract.SetMaxBlocks(&_LightNode.TransactOpts, val)
}

// SetMaxBlocks is a paid mutator transaction binding the contract method 0x5092515e.
//
// Solidity: function setMaxBlocks(uint256 val) returns()
func (_LightNode *LightNodeTransactorSession) SetMaxBlocks(val *big.Int) (*types.Transaction, error) {
	return _LightNode.Contract.SetMaxBlocks(&_LightNode.TransactOpts, val)
}

// TogglePause is a paid mutator transaction binding the contract method 0x57d159c6.
//
// Solidity: function togglePause(bool flag) returns(bool)
func (_LightNode *LightNodeTransactor) TogglePause(opts *bind.TransactOpts, flag bool) (*types.Transaction, error) {
	return _LightNode.contract.Transact(opts, "togglePause", flag)
}

// TogglePause is a paid mutator transaction binding the contract method 0x57d159c6.
//
// Solidity: function togglePause(bool flag) returns(bool)
func (_LightNode *LightNodeSession) TogglePause(flag bool) (*types.Transaction, error) {
	return _LightNode.Contract.TogglePause(&_LightNode.TransactOpts, flag)
}

// TogglePause is a paid mutator transaction binding the contract method 0x57d159c6.
//
// Solidity: function togglePause(bool flag) returns(bool)
func (_LightNode *LightNodeTransactorSession) TogglePause(flag bool) (*types.Transaction, error) {
	return _LightNode.Contract.TogglePause(&_LightNode.TransactOpts, flag)
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

// UpdateLightClient is a paid mutator transaction binding the contract method 0xda027b4e.
//
// Solidity: function updateLightClient((uint64,uint64,bytes32,bytes32,uint64,uint64,(uint64,(bytes32,bytes,bytes,uint64)[],uint64,uint64,bytes),(bytes32,uint64),bytes32,(bytes32,bytes)[]) ledgerInfo) returns()
func (_LightNode *LightNodeTransactor) UpdateLightClient(opts *bind.TransactOpts, ledgerInfo TypesLedgerInfoWithSignatures) (*types.Transaction, error) {
	return _LightNode.contract.Transact(opts, "updateLightClient", ledgerInfo)
}

// UpdateLightClient is a paid mutator transaction binding the contract method 0xda027b4e.
//
// Solidity: function updateLightClient((uint64,uint64,bytes32,bytes32,uint64,uint64,(uint64,(bytes32,bytes,bytes,uint64)[],uint64,uint64,bytes),(bytes32,uint64),bytes32,(bytes32,bytes)[]) ledgerInfo) returns()
func (_LightNode *LightNodeSession) UpdateLightClient(ledgerInfo TypesLedgerInfoWithSignatures) (*types.Transaction, error) {
	return _LightNode.Contract.UpdateLightClient(&_LightNode.TransactOpts, ledgerInfo)
}

// UpdateLightClient is a paid mutator transaction binding the contract method 0xda027b4e.
//
// Solidity: function updateLightClient((uint64,uint64,bytes32,bytes32,uint64,uint64,(uint64,(bytes32,bytes,bytes,uint64)[],uint64,uint64,bytes),(bytes32,uint64),bytes32,(bytes32,bytes)[]) ledgerInfo) returns()
func (_LightNode *LightNodeTransactorSession) UpdateLightClient(ledgerInfo TypesLedgerInfoWithSignatures) (*types.Transaction, error) {
	return _LightNode.Contract.UpdateLightClient(&_LightNode.TransactOpts, ledgerInfo)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LightNode *LightNodeTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _LightNode.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LightNode *LightNodeSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _LightNode.Contract.UpgradeTo(&_LightNode.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LightNode *LightNodeTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _LightNode.Contract.UpgradeTo(&_LightNode.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LightNode *LightNodeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LightNode.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LightNode *LightNodeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LightNode.Contract.UpgradeToAndCall(&_LightNode.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LightNode *LightNodeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LightNode.Contract.UpgradeToAndCall(&_LightNode.TransactOpts, newImplementation, data)
}

// LightNodeAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the LightNode contract.
type LightNodeAdminChangedIterator struct {
	Event *LightNodeAdminChanged // Event containing the contract specifics and raw log

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
func (it *LightNodeAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightNodeAdminChanged)
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
		it.Event = new(LightNodeAdminChanged)
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
func (it *LightNodeAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightNodeAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightNodeAdminChanged represents a AdminChanged event raised by the LightNode contract.
type LightNodeAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_LightNode *LightNodeFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*LightNodeAdminChangedIterator, error) {

	logs, sub, err := _LightNode.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &LightNodeAdminChangedIterator{contract: _LightNode.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_LightNode *LightNodeFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *LightNodeAdminChanged) (event.Subscription, error) {

	logs, sub, err := _LightNode.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightNodeAdminChanged)
				if err := _LightNode.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_LightNode *LightNodeFilterer) ParseAdminChanged(log types.Log) (*LightNodeAdminChanged, error) {
	event := new(LightNodeAdminChanged)
	if err := _LightNode.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightNodeBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the LightNode contract.
type LightNodeBeaconUpgradedIterator struct {
	Event *LightNodeBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *LightNodeBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightNodeBeaconUpgraded)
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
		it.Event = new(LightNodeBeaconUpgraded)
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
func (it *LightNodeBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightNodeBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightNodeBeaconUpgraded represents a BeaconUpgraded event raised by the LightNode contract.
type LightNodeBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_LightNode *LightNodeFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*LightNodeBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _LightNode.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &LightNodeBeaconUpgradedIterator{contract: _LightNode.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_LightNode *LightNodeFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *LightNodeBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _LightNode.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightNodeBeaconUpgraded)
				if err := _LightNode.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_LightNode *LightNodeFilterer) ParseBeaconUpgraded(log types.Log) (*LightNodeBeaconUpgraded, error) {
	event := new(LightNodeBeaconUpgraded)
	if err := _LightNode.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightNodeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LightNode contract.
type LightNodeInitializedIterator struct {
	Event *LightNodeInitialized // Event containing the contract specifics and raw log

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
func (it *LightNodeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightNodeInitialized)
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
		it.Event = new(LightNodeInitialized)
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
func (it *LightNodeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightNodeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightNodeInitialized represents a Initialized event raised by the LightNode contract.
type LightNodeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LightNode *LightNodeFilterer) FilterInitialized(opts *bind.FilterOpts) (*LightNodeInitializedIterator, error) {

	logs, sub, err := _LightNode.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LightNodeInitializedIterator{contract: _LightNode.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LightNode *LightNodeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LightNodeInitialized) (event.Subscription, error) {

	logs, sub, err := _LightNode.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightNodeInitialized)
				if err := _LightNode.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LightNode *LightNodeFilterer) ParseInitialized(log types.Log) (*LightNodeInitialized, error) {
	event := new(LightNodeInitialized)
	if err := _LightNode.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightNodePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the LightNode contract.
type LightNodePausedIterator struct {
	Event *LightNodePaused // Event containing the contract specifics and raw log

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
func (it *LightNodePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightNodePaused)
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
		it.Event = new(LightNodePaused)
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
func (it *LightNodePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightNodePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightNodePaused represents a Paused event raised by the LightNode contract.
type LightNodePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LightNode *LightNodeFilterer) FilterPaused(opts *bind.FilterOpts) (*LightNodePausedIterator, error) {

	logs, sub, err := _LightNode.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LightNodePausedIterator{contract: _LightNode.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LightNode *LightNodeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LightNodePaused) (event.Subscription, error) {

	logs, sub, err := _LightNode.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightNodePaused)
				if err := _LightNode.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LightNode *LightNodeFilterer) ParsePaused(log types.Log) (*LightNodePaused, error) {
	event := new(LightNodePaused)
	if err := _LightNode.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightNodeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the LightNode contract.
type LightNodeUnpausedIterator struct {
	Event *LightNodeUnpaused // Event containing the contract specifics and raw log

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
func (it *LightNodeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightNodeUnpaused)
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
		it.Event = new(LightNodeUnpaused)
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
func (it *LightNodeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightNodeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightNodeUnpaused represents a Unpaused event raised by the LightNode contract.
type LightNodeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LightNode *LightNodeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LightNodeUnpausedIterator, error) {

	logs, sub, err := _LightNode.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LightNodeUnpausedIterator{contract: _LightNode.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LightNode *LightNodeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LightNodeUnpaused) (event.Subscription, error) {

	logs, sub, err := _LightNode.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightNodeUnpaused)
				if err := _LightNode.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LightNode *LightNodeFilterer) ParseUnpaused(log types.Log) (*LightNodeUnpaused, error) {
	event := new(LightNodeUnpaused)
	if err := _LightNode.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateLightClient is a free log retrieval operation binding the contract event 0x027143ef91911a3444f6b749bf632b3946978f1e2d22c6679246489a8d591794.
//
// Solidity: event UpdateLightClient(address indexed account, uint256 epoch, uint256 round)
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

// WatchUpdateLightClient is a free log subscription operation binding the contract event 0x027143ef91911a3444f6b749bf632b3946978f1e2d22c6679246489a8d591794.
//
// Solidity: event UpdateLightClient(address indexed account, uint256 epoch, uint256 round)
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

// ParseUpdateLightClient is a log parse operation binding the contract event 0x027143ef91911a3444f6b749bf632b3946978f1e2d22c6679246489a8d591794.
//
// Solidity: event UpdateLightClient(address indexed account, uint256 epoch, uint256 round)
func (_LightNode *LightNodeFilterer) ParseUpdateLightClient(log types.Log) (*LightNodeUpdateLightClient, error) {
	event := new(LightNodeUpdateLightClient)
	if err := _LightNode.contract.UnpackLog(event, "UpdateLightClient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightNodeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the LightNode contract.
type LightNodeUpgradedIterator struct {
	Event *LightNodeUpgraded // Event containing the contract specifics and raw log

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
func (it *LightNodeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightNodeUpgraded)
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
		it.Event = new(LightNodeUpgraded)
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
func (it *LightNodeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightNodeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightNodeUpgraded represents a Upgraded event raised by the LightNode contract.
type LightNodeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LightNode *LightNodeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*LightNodeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LightNode.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &LightNodeUpgradedIterator{contract: _LightNode.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LightNode *LightNodeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *LightNodeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LightNode.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightNodeUpgraded)
				if err := _LightNode.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LightNode *LightNodeFilterer) ParseUpgraded(log types.Log) (*LightNodeUpgraded, error) {
	event := new(LightNodeUpgraded)
	if err := _LightNode.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
