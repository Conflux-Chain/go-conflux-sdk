package main

import (
	"fmt"
	"math/big"
	"time"

	internalcontract "github.com/Conflux-Chain/go-conflux-sdk/contract_meta/internal_contract"
	"github.com/Conflux-Chain/go-conflux-sdk/example/context"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
)

func testStaking() {

	staking, err := internalcontract.NewStaking(client)
	if err != nil {
		context.PanicIfErrf(err, "failed to get staking contract")
	}

	balance, _ := client.GetBalance(*defaultAccount)
	fmt.Printf("%v has balance %v\n", defaultAccount, balance)

	// getStakingBalance
	stakeBalance, err := staking.GetStakingBalance(nil, *defaultAccount)
	context.PanicIfErrf(err, "getStakingBalance panic")
	fmt.Printf("staking balance is:%v\n", stakeBalance)

	if stakeBalance.Cmp(big.NewInt(0)) == 0 {
		// deposit
		txhash, err := staking.Deposit(&types.ContractMethodSendOption{Nonce: context.GetNextNonceAndIncrease()}, big.NewInt(2e18))
		context.PanicIfErrf(err, "deposit panic")
		receipt, err := client.WaitForTransationReceipt(txhash, time.Second)
		context.PanicIfErrf(err, "wait for pack panic")
		fmt.Printf("depoite done:%v\n", context.JSONFmt(receipt))

		// getStakingBalance
		stakeBalance, err := staking.GetStakingBalance(nil, *defaultAccount)
		context.PanicIfErrf(err, "getStakingBalance panic")
		fmt.Printf("staking balance is:%v\n", stakeBalance)
	}

	// withdraw
	txhash, err := staking.Withdraw(&types.ContractMethodSendOption{Nonce: context.GetNextNonceAndIncrease()}, big.NewInt(100))
	context.PanicIfErrf(err, "withdraw panic")
	receipt, err := client.WaitForTransationReceipt(txhash, time.Second)
	context.PanicIfErrf(err, "wait for pack panic")
	fmt.Printf("withdraw done:%v\n", context.JSONFmt(receipt))

	// getStakingBalance
	stakeBalance, err = staking.GetStakingBalance(nil, *defaultAccount)
	context.PanicIfErrf(err, "getStakingBalance panic")
	fmt.Printf("staking balance is:%v\n", stakeBalance)

	// getLockedStakingBalance
	blockNumber, _ := client.GetEpochNumber()
	lockedStakingBlance, err := staking.GetLockedStakingBalance(nil, *defaultAccount, blockNumber.ToInt())
	context.PanicIfErrf(err, "getLockedStakingBalance panic")
	fmt.Printf("currently locked staking balance is %v \n", lockedStakingBlance)

	if lockedStakingBlance.Cmp(big.NewInt(0)) == 0 {
		// voteLock
		quarterBlockNumber := int64(2 * 3600 * 24 * 31 * 3)
		txhash, err = staking.VoteLock(&types.ContractMethodSendOption{Nonce: context.GetNextNonceAndIncrease()}, big.NewInt(1e18), big.NewInt(quarterBlockNumber))
		context.PanicIfErrf(err, "voteLock panic")
		receipt, err = client.WaitForTransationReceipt(txhash, time.Second)
		context.PanicIfErrf(err, "wait for pack panic")
		fmt.Printf("vote lock %v for %v blocknumber\n", big.NewInt(1e18), quarterBlockNumber)
	}

	// getVotePower
	power, err := staking.GetVotePower(nil, *defaultAccount, blockNumber.ToInt())
	context.PanicIfErrf(err, "getVotePower panic")
	fmt.Printf("currently vote power is %v \n", power)

}
