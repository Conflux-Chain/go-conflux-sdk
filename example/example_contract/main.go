package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"path"
	"runtime"
	"time"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/example/context"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	config := context.PrepareForContractExample()
	client := config.GetClient()
	currentDir := getCurrentDir()

	//deploy contract
	fmt.Println("start deploy contract...")
	abiPath := path.Join(currentDir, "./contract/erc20.abi")
	bytecodePath := path.Join(currentDir, "./contract/erc20.bytecode")
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

	result := client.DeployContract(nil, abi, bytecode, big.NewInt(100000), "biu", uint8(10), "BIU")
	_ = <-result.DoneChannel
	if result.Error != nil {
		panic(result.Error)
	}
	contract = result.DeployedContract
	fmt.Printf("deploy contract by client.DeployContract done\ncontract address: %+v\ntxhash:%v\n\n", contract.Address, result.TransactionHash)

	time.Sleep(10 * time.Second)

	// or get contract by deployed address
	// deployedAt := types.Address("0x8d1089f00c40dcc290968b366889e85e67024662")
	// contract, err := client.GetContract(string(abi), &deployedAt)
	// if err != nil {
	// 	panic(err)
	// }

	//get data for send/call contract method
	chainID, err := contract.Client.GetNetworkID()
	context.PanicIfErrf(err, "failed to get chainID")
	user := cfxaddress.MustNewFromHex("0x19f4bcf113e0b896d9b34294fd3da86b4adf0302", chainID)
	data, err := contract.GetData("balanceOf", user.MustGetCommonAddress())
	if err != nil {
		panic(err)
	}
	fmt.Printf("get data of method balanceOf result: 0x%x\n\n", data)

	//call contract method
	//Note: the output struct type need match method output type of ABI, go type "*big.Int" match abi type "uint256", go type "struct{Balance *big.Int}" match abi tuple type "(balance uint256)"
	balance := &struct{ Balance *big.Int }{}
	err = contract.Call(nil, balance, "balanceOf", user.MustGetCommonAddress())
	if err != nil {
		panic(err)
	}
	fmt.Printf("balance of address %v in contract is: %+v\n\n", user, balance)

	// name symbol decimals
	name := ""
	err = contract.Call(nil, &name, "name")
	if err != nil {
		panic(err)
	}
	fmt.Printf("name of contract is: %+v\n\n", name)

	symbol := ""
	err = contract.Call(nil, &symbol, "symbol")
	if err != nil {
		panic(err)
	}
	fmt.Printf("symbol of contract is: %+v\n\n", name)

	decimals := uint8(0)
	err = contract.Call(nil, &decimals, "decimals")
	if err != nil {
		panic(err)
	}
	fmt.Printf("decimals of contract is: %+v\n\n", decimals)

	//send transction for contract method
	to := cfxaddress.MustNewFromHex("0x160ebef20c1f739957bf9eecd040bce699cc42c6")
	txhash, err := contract.SendTransaction(nil, "transfer", to.MustGetCommonAddress(), big.NewInt(10))
	if err != nil {
		panic(err)
	}

	fmt.Printf("transfer %v erc20 token to %v done, tx hash: %v\n\n", 10, to, txhash)

	context.WaitPacked(client, txhash)

	fmt.Println("wait be excuted")
	time.Sleep(10 * time.Second)

	//get event log and decode it
	receipt, err := client.GetTransactionReceipt(txhash)
	if err != nil {
		panic(err)
	}

	// decode Transfer Event
	var Transfer struct {
		From  common.Address
		To    common.Address
		Value *big.Int
	}

	if receipt == nil || len(receipt.Logs) == 0 {
		panic(fmt.Errorf("receipt is nil or no logs in receipt, receipt: %v", receipt))
	}

	err = contract.DecodeEvent(&Transfer, "Transfer", receipt.Logs[0])
	if err != nil {
		panic(err)
	}
	fmt.Printf("decoded transfer event: {From: 0x%x, To: 0x%x, Value: %v} \n", Transfer.From, Transfer.To, Transfer.Value)
}

func getCurrentDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("get current file path error")
	}
	currentDir := path.Join(filename, "../")
	return currentDir
}
