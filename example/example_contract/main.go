package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"time"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
)

func main() {

	//unlock account
	am := sdk.NewAccountManager("../keystore")
	err := am.TimedUnlockDefault("hello", 300*time.Second)
	if err != nil {
		panic(err)
	}

	//init client
	client, err := sdk.NewClient("http://testnet-jsonrpc.conflux-chain.org:12537")
	if err != nil {
		panic(err)
	}
	client.SetAccountManager(am)

	//deploy contract
	fmt.Println("start deploy contract...")
	abiPath := "./contract/erc20.abi"
	bytecodePath := "./contract/erc20.bytecode"
	var contract *sdk.Contract

	abi, err := ioutil.ReadFile(abiPath)
	if err != nil {
		panic(err)
	}

	bytecodeHexStr, err := ioutil.ReadFile(bytecodePath)
	if err != nil {
		panic(err)
	}

	bytecode, err := hex.DecodeString(string(bytecodeHexStr))
	if err != nil {
		panic(err)
	}

	doneChan := client.DeployContract(string(abi), bytecode, nil, time.Duration(time.Second*30), func(c sdk.Contractor, txhash *types.Hash, err error) {
		if err != nil {
			panic(err)
		}
		contract = c.(*sdk.Contract)
		fmt.Printf("deploy contract by client.DeployContract done\ncontract address: %+v\ntxhash:%v\n\n", *contract.Address, txhash)
	})

	_ = <-doneChan
	fmt.Println("wait for epoch excution for 15 seconds...")
	time.Sleep(15 * time.Second)

	// or get contract by deployed address
	// deployedAt := types.Address("0x8d1089f00c40dcc290968b366889e85e67024662")
	// contract, err := client.GetContract(string(abi), &deployedAt)
	// if err != nil {
	// 	panic(err)
	// }

	//get data for send/call contract method
	user := types.Address("0x19f4bcf113e0b896d9b34294fd3da86b4adf0302")
	data, err := contract.GetData("balanceOf", user.ToCommonAddress())
	if err != nil {
		panic(err)
	}
	fmt.Printf("get data of method balanceOf is: 0x%x\n\n", data)

	//call contract method
	//Note: the output struct type need match method output type of ABI, go type "*big.Int" match abi type "uint256", go type "struct{Balance *big.Int}" match abi tuple type "(balance uint256)"
	balance := &struct{ Balance *big.Int }{}
	err = contract.Call(nil, balance, "balanceOf", user.ToCommonAddress())
	if err != nil {
		panic(err)
	}
	fmt.Printf("address %v balance in contract is: %+v\n\n", user, balance)

	//send transction for contract method
	to := types.Address("0x160ebef20c1f739957bf9eecd040bce699cc42c6")
	txhash, err := contract.SendTransaction(nil, "transfer", to.ToCommonAddress(), big.NewInt(10))
	if err != nil {
		panic(err)
	}

	fmt.Printf("transfer %v erc20 token to %v done, tx hash: %v\n\n", 10, to, txhash)

	fmt.Println("wait for transaction be packed...")
	for {
		time.Sleep(time.Duration(1) * time.Second)
		tx, err := client.GetTransactionByHash(*txhash)
		if err != nil {
			panic(err)
		}
		if tx.Status != nil {
			fmt.Printf("transaction is packed.")
			break
		}
	}
}
