package internalcontract

import (
	"math/big"
	"sync"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/pkg/errors"
)

// ParamsControl contract
type ParamsControl struct {
	sdk.Contract
}

type ParamsControlVote struct {
	TopicIndex uint16
	Votes      [3]*big.Int
}

var paramsControlMap sync.Map
var paramsControlMu sync.Mutex

// NewParamsControl gets the ParamsControl contract object
func NewParamsControl(client sdk.ClientOperator) (s ParamsControl, err error) {
	netId, err := client.GetNetworkID()
	if err != nil {
		return ParamsControl{}, err
	}
	val, ok := paramsControlMap.Load(netId)
	if !ok {
		paramsControlMu.Lock()
		defer paramsControlMu.Unlock()
		abi := getParamsControlAbi()
		address, e := getParamsControlAddress(client)
		if e != nil {
			return s, errors.Wrap(e, "failed to get ParamsControl address")
		}
		contract, e := sdk.NewContract([]byte(abi), client, &address)
		if e != nil {
			return s, errors.Wrap(e, "failed to new ParamsControl contract")
		}
		val = ParamsControl{Contract: *contract}
		paramsControlMap.Store(netId, val)
	}
	return val.(ParamsControl), nil
}

func getParamsControlAbi() string {
	return "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"vote_round\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"topic_index\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256[3]\",\"name\":\"votes\",\"type\":\"uint256[3]\"}],\"name\":\"CastVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"vote_round\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"topic_index\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256[3]\",\"name\":\"votes\",\"type\":\"uint256[3]\"}],\"name\":\"RevokeVote\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"vote_round\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"topic_index\",\"type\":\"uint16\"},{\"internalType\":\"uint256[3]\",\"name\":\"votes\",\"type\":\"uint256[3]\"}],\"internalType\":\"structParamsControl.Vote[]\",\"name\":\"vote_data\",\"type\":\"tuple[]\"}],\"name\":\"castVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"readVote\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"topic_index\",\"type\":\"uint16\"},{\"internalType\":\"uint256[3]\",\"name\":\"votes\",\"type\":\"uint256[3]\"}],\"internalType\":\"structParamsControl.Vote[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"vote_round\",\"type\":\"uint64\"}],\"name\":\"totalVotes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"topic_index\",\"type\":\"uint16\"},{\"internalType\":\"uint256[3]\",\"name\":\"votes\",\"type\":\"uint256[3]\"}],\"internalType\":\"structParamsControl.Vote[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteRound\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"
}

func getParamsControlAddress(client sdk.ClientOperator) (types.Address, error) {
	addr := cfxaddress.MustNewFromHex("0888000000000000000000000000000000000007")
	err := addr.CompleteByClient(client)
	return addr, err
}

// =================== calls ==================

func (p *ParamsControl) ReadVote(opts *types.ContractMethodCallOption, addr types.Address) ([]ParamsControlVote, error) {
	out := []ParamsControlVote{}
	err := p.Call(opts, &out, "readVote", addr.MustGetCommonAddress())
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (p *ParamsControl) TotalVotes(opts *types.ContractMethodCallOption, vote_round uint64) ([]ParamsControlVote, error) {
	out := []ParamsControlVote{}
	err := p.Call(opts, &out, "totalVotes", vote_round)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (p *ParamsControl) CurrentRound(opts *types.ContractMethodCallOption) (uint64, error) {
	var out uint64
	err := p.Call(opts, &out, "currentRound")
	if err != nil {
		return 0, err
	}
	return out, nil
}

func (p *ParamsControl) PosStakeForVotes(opts *types.ContractMethodCallOption, arg0 uint64) (*big.Int, error) {
	var out = new(big.Int)
	err := p.Call(opts, &out, "posStakeForVotes")
	if err != nil {
		return nil, err
	}
	return out, nil
}

// =================== sends ==================

func (p *ParamsControl) CastVote(opts *types.ContractMethodSendOption, vote_round uint64, vote_data []ParamsControlVote) (types.Hash, error) {
	return p.SendTransaction(opts, "castVote", vote_round, vote_data)
}
