package main

import (
	"fmt"
	"math/big"
	"time"

	internalcontract "github.com/Conflux-Chain/go-conflux-sdk/contract_meta/internal_contract"
	"github.com/Conflux-Chain/go-conflux-sdk/example/context"
)

func testStaking() {

	staking := internalcontract.NewStaking(client)

	balance, _ := client.GetBalance(*defaultAccount)
	fmt.Printf("%v has balance %v\n", defaultAccount, balance)

	// getStakingBalance
	stakeBalance, err := staking.GetStakingBalance(nil, *defaultAccount)
	context.PanicIfErr(err, "getStakingBalance panic")
	fmt.Printf("staking balance is:%v\n", stakeBalance)

	if stakeBalance.Cmp(big.NewInt(0)) == 0 {
		// deposit
		txhash, err := staking.Deposit(nil, big.NewInt(2e18))
		context.PanicIfErr(err, "deposit panic")
		receipt, err := client.WaitForTransationReceipt(*txhash, time.Second)
		context.PanicIfErr(err, "wait for pack panic")
		fmt.Printf("depoite done:%v\n", context.JsonFmt(receipt))

		// getStakingBalance
		stakeBalance, err := staking.GetStakingBalance(nil, *defaultAccount)
		context.PanicIfErr(err, "getStakingBalance panic")
		fmt.Printf("staking balance is:%v\n", stakeBalance)
	}

	// withdraw
	txhash, err := staking.Withdraw(nil, big.NewInt(100))
	context.PanicIfErr(err, "withdraw panic")
	receipt, err := client.WaitForTransationReceipt(*txhash, time.Second)
	context.PanicIfErr(err, "wait for pack panic")
	fmt.Printf("withdraw done:%v\n", context.JsonFmt(receipt))

	// getStakingBalance
	stakeBalance, err = staking.GetStakingBalance(nil, *defaultAccount)
	context.PanicIfErr(err, "getStakingBalance panic")
	fmt.Printf("staking balance is:%v\n", stakeBalance)

	// getLockedStakingBalance
	blockNumber, _ := client.GetEpochNumber()
	lockedStakingBlance, err := staking.GetLockedStakingBalance(nil, *defaultAccount, blockNumber)
	context.PanicIfErr(err, "getLockedStakingBalance panic")
	fmt.Printf("currently locked staking balance is %v \n", lockedStakingBlance)

	if lockedStakingBlance.Cmp(big.NewInt(0)) == 0 {
		// voteLock
		quarterBlockNumber := int64(2 * 3600 * 24 * 31 * 3)
		txhash, err = staking.VoteLock(nil, big.NewInt(1e18), big.NewInt(quarterBlockNumber))
		context.PanicIfErr(err, "voteLock panic")
		receipt, err = client.WaitForTransationReceipt(*txhash, time.Second)
		context.PanicIfErr(err, "wait for pack panic")
		fmt.Printf("vote lock %v for %v blocknumber\n", big.NewInt(1e18), quarterBlockNumber)
	}

	// getVotePower
	power, err := staking.GetVotePower(nil, *defaultAccount, blockNumber)
	context.PanicIfErr(err, "getVotePower panic")
	fmt.Printf("currently vote power is %v \n", power)

}
