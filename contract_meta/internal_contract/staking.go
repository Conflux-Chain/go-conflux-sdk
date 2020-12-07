package internalcontract

import (
	"math/big"
	"sync"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/example/context"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
)

// Staking contract
type Staking struct {
	sdk.Contract
}

var staking *Staking
var stakingOnce sync.Once

// NewStaking gets the Staking contract object
func NewStaking(client sdk.ClientOperator) *Staking {
	stakingOnce.Do(func() {
		abi := getStakingAbi()
		address := getStakingAddress()
		contract, err := sdk.NewContract([]byte(abi), client, &address)
		context.PanicIfErr(err, "new staking panic")
		staking = &Staking{Contract: *contract}
	})
	return staking
}

// GetStakingBalance returns user's staking balance
func (ac *Staking) GetStakingBalance(option *types.ContractMethodCallOption, user types.Address) (*big.Int, error) {
	var tmp *big.Int = new(big.Int)
	err := ac.Call(option, &tmp, "getStakingBalance", user.ToCommonAddress())
	if err != nil {
		return nil, err
	}
	return tmp, nil
}

// GetLockedStakingBalance returns user's locked staking balance at given blockNumber
// Note: if the blockNumber is less than the current block number, function
// will return current locked staking balance.
func (ac *Staking) GetLockedStakingBalance(option *types.ContractMethodCallOption, user types.Address, blockNumber *big.Int) (*big.Int, error) {
	var tmp *big.Int = new(big.Int)
	err := ac.Call(option, &tmp, "getLockedStakingBalance", user.ToCommonAddress(), blockNumber)
	if err != nil {
		return nil, err
	}
	return tmp, nil
}

// GetVotePower returns user's vote power staking balance at given blockNumber
func (ac *Staking) GetVotePower(option *types.ContractMethodCallOption, user types.Address, blockNumber *big.Int) (*big.Int, error) {
	var tmp *big.Int = new(big.Int)
	err := ac.Call(option, &tmp, "getVotePower", user.ToCommonAddress(), blockNumber)
	if err != nil {
		return nil, err
	}
	return tmp, nil
}

// Deposit `amount` cfx in this contract
func (s *Staking) Deposit(option *types.ContractMethodSendOption, amount *big.Int) (*types.Hash, error) {
	return s.SendTransaction(option, "deposit", amount)
}

// Withdraw `amount` cfx from this contract
func (s *Staking) Withdraw(option *types.ContractMethodSendOption, amount *big.Int) (*types.Hash, error) {
	return s.SendTransaction(option, "withdraw", amount)
}

// VoteLock will locks `amount` cfx from current to next `unlockBlockNumber` blocks for obtain vote power
func (s *Staking) VoteLock(option *types.ContractMethodSendOption, amount *big.Int, unlockBlockNumber *big.Int) (*types.Hash, error) {
	return s.SendTransaction(option, "voteLock", amount, unlockBlockNumber)
}

func getStakingAbi() string {
	return `[
			{
				"inputs": [
					{
						"internalType": "uint256",
						"name": "amount",
						"type": "uint256"
					}
				],
				"name": "deposit",
				"outputs": [],
				"stateMutability": "nonpayable",
				"type": "function"
			},
			{
				"inputs": [
					{
						"internalType": "address",
						"name": "user",
						"type": "address"
					},
					{
						"internalType": "uint256",
						"name": "blockNumber",
						"type": "uint256"
					}
				],
				"name": "getLockedStakingBalance",
				"outputs": [
					{
						"internalType": "uint256",
						"name": "",
						"type": "uint256"
					}
				],
				"stateMutability": "view",
				"type": "function"
			},
			{
				"inputs": [
					{
						"internalType": "address",
						"name": "user",
						"type": "address"
					}
				],
				"name": "getStakingBalance",
				"outputs": [
					{
						"internalType": "uint256",
						"name": "",
						"type": "uint256"
					}
				],
				"stateMutability": "view",
				"type": "function"
			},
			{
				"inputs": [
					{
						"internalType": "address",
						"name": "user",
						"type": "address"
					},
					{
						"internalType": "uint256",
						"name": "blockNumber",
						"type": "uint256"
					}
				],
				"name": "getVotePower",
				"outputs": [
					{
						"internalType": "uint256",
						"name": "",
						"type": "uint256"
					}
				],
				"stateMutability": "view",
				"type": "function"
			},
			{
				"inputs": [
					{
						"internalType": "uint256",
						"name": "amount",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "unlockBlockNumber",
						"type": "uint256"
					}
				],
				"name": "voteLock",
				"outputs": [],
				"stateMutability": "nonpayable",
				"type": "function"
			},
			{
				"inputs": [
					{
						"internalType": "uint256",
						"name": "amount",
						"type": "uint256"
					}
				],
				"name": "withdraw",
				"outputs": [],
				"stateMutability": "nonpayable",
				"type": "function"
			}
		]`
}

func getStakingAddress() types.Address {
	return types.Address("0x0888000000000000000000000000000000000002")
}
