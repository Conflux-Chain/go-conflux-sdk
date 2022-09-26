package internalcontract

import (
	"math/big"
	"sync"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/bind"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// PoSRegister contract
type PoSRegister struct {
	sdk.Contract
}

var poSRegisterMap sync.Map
var poSRegisterMu sync.Mutex

// NewPoSRegister gets the PoSRegister contract object
func NewPoSRegister(client sdk.ClientOperator) (s PoSRegister, err error) {
	netId, err := client.GetNetworkID()
	if err != nil {
		return PoSRegister{}, err
	}
	val, ok := poSRegisterMap.Load(netId)
	if !ok {
		poSRegisterMu.Lock()
		defer poSRegisterMu.Unlock()
		abi := getPoSRegisterAbi()
		address, e := getPoSRegisterAddress(client)
		if e != nil {
			return s, errors.Wrap(e, "failed to get PoSRegister address")
		}
		contract, e := sdk.NewContract([]byte(abi), client, &address)
		if e != nil {
			return s, errors.Wrap(e, "failed to new PoSRegister contract")
		}
		val = PoSRegister{Contract: *contract}
		poSRegisterMap.Store(netId, val)
	}
	return val.(PoSRegister), nil
}

func getPoSRegisterAbi() string {
	return "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"votePower\",\"type\":\"uint64\"}],\"name\":\"IncreaseStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"blsPubKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"vrfPubKey\",\"type\":\"bytes\"}],\"name\":\"Register\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"votePower\",\"type\":\"uint64\"}],\"name\":\"Retire\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addressToIdentifier\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"}],\"name\":\"identifierToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"votePower\",\"type\":\"uint64\"}],\"name\":\"increaseStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"indentifier\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"votePower\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"blsPubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"vrfPubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes[2]\",\"name\":\"blsPubKeyProof\",\"type\":\"bytes[2]\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"votePower\",\"type\":\"uint64\"}],\"name\":\"retire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
}

func getPoSRegisterAddress(client sdk.ClientOperator) (types.Address, error) {
	addr := cfxaddress.MustNewFromHex("0888000000000000000000000000000000000005")
	err := addr.CompleteByClient(client)
	return addr, err
}

// =================== calls ==================

func (p *PoSRegister) ReadVote(opts *types.ContractMethodCallOption, addr types.Address) ([]ParamsControlVote, error) {
	out := []ParamsControlVote{}
	err := p.Call(opts, &out, "readVote", addr.MustGetCommonAddress())
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetVotes is a free data retrieval call binding the contract method 0x4c051100.
//
// Solidity: function getVotes(bytes32 identifier) view returns(uint256, uint256)
func (p *PoSRegister) GetVotes(opts *types.ContractMethodCallOption, identifier [32]byte) (*big.Int, *big.Int, error) {
	var out [2]*big.Int
	err := p.Call(opts, &out, "getVotes", identifier)
	if err != nil {
		return nil, nil, err
	}
	return out[0], out[1], nil
}

func (p *PoSRegister) AddressToIdentifier(opts *types.ContractMethodCallOption, addr common.Address) ([32]byte, error) {
	var out [32]byte
	err := p.Call(opts, &out, "addressToIdentifier", addr)
	if err != nil {
		return out, err
	}
	return out, nil
}

func (p *PoSRegister) IdentifierToAddress(opts *types.ContractMethodCallOption, identifier [32]byte) (common.Address, error) {
	var out common.Address
	err := p.Call(opts, &out, "identifierToAddress", identifier)
	if err != nil {
		return out, err
	}
	return out, nil
}

// =================== sends ==================

func (p *PoSRegister) IncreaseStake(opts *bind.TransactOpts, votePower uint64) (types.Hash, error) {
	return p.SendTransaction(opts, "increaseStake", votePower)
}

func (p *PoSRegister) Register(opts *bind.TransactOpts, identifier [32]byte, votePower uint64, blsPubKey []byte, vrfPubKey []byte, blsPubKeyProof [2][]byte) (types.Hash, error) {
	return p.SendTransaction(opts, "register", identifier, votePower, blsPubKey, vrfPubKey, blsPubKeyProof)
}

func (p *PoSRegister) Retire(opts *bind.TransactOpts, votePower uint64) (types.Hash, error) {
	return p.SendTransaction(opts, "retire", votePower)
}
