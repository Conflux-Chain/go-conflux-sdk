package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"time"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
)

func main() {

	//init client
	client, err := sdk.NewClient("https://test.confluxrpc.com", sdk.ClientOption{
		KeystorePath: "../context/keystore",
	})
	utils.PanicIfErr(err, "failed to new client")

	// unlock default account
	client.AccountManager.TimedUnlockDefault("hello", 300*time.Second)

	//deploy contract
	abi := mustGetAbi("../context/contract/erc20.abi")
	bytecode := mustGetBytecode("../context/contract/erc20.bytecode")

	result := client.DeployContract(nil, abi, bytecode, big.NewInt(100000), "biu", uint8(10), "BIU")
	<-result.DoneChannel
	utils.PanicIfErr(result.Error, "failed to deploy contract")

	contract := result.DeployedContract
	fmt.Printf("deploy contract by client.DeployContract done\ncontract address: %+v\ntxhash:%v\n\n", contract.Address, result.TransactionHash)

	// or get contract by deployed address
	// deployedAt := client.MustNewAddress("cfxtest:acgkhpdz61g11parejzbftznnt8gds15mp4wg54j5c")
	// contract, err := client.GetContract(abi, &deployedAt)

	//get data for send/call contract method
	user := client.MustNewAddress("cfxtest:aap9kthvctunvf030rbkk9k7zbzyz12dajp1u3sp4g")
	data, err := contract.GetData("balanceOf", user.MustGetCommonAddress())
	utils.PanicIfErr(err, "failed to get data")
	fmt.Printf("get data of method balanceOf: 0x%x\n\n", data)

	//call contract method
	//Note: the output struct type need match method output type of ABI, go type "*big.Int" match abi type "uint256", go type "struct{Balance *big.Int}" match abi tuple type "(balance uint256)"
	balance := &struct{ Balance *big.Int }{}
	err = contract.Call(nil, balance, "balanceOf", user.MustGetCommonAddress())
	utils.PanicIfErr(err, "failed to get balance")
	fmt.Printf("balance of address %v in contract is: %+v\n\n", user, balance)

	//send transction for contract method
	to := client.MustNewAddress("cfxtest:acgkhpdz61g11parejzbftznnt8gds15mp4wg54j5c")
	txhash, err := contract.SendTransaction(nil, "transfer", to.MustGetCommonAddress(), big.NewInt(10))
	utils.PanicIfErr(err, "failed to transfer")
	fmt.Printf("transfer %v erc20 token to %v done, tx hash: %v\n\n", 10, to, txhash)

	//get event log and decode it
	fmt.Println("wait for transaction be packed...")
	receipt, err := client.WaitForTransationReceipt(txhash, time.Second*2)
	utils.PanicIfErr(err, "failed to get tx receipt")
	fmt.Printf("get receipt: %+v\n\n", receipt)

	// decode Transfer Event
	var Transfer struct {
		From  common.Address
		To    common.Address
		Value *big.Int
	}

	err = contract.DecodeEvent(&Transfer, "Transfer", receipt.Logs[0])
	utils.PanicIfErr(err, "failed to decode event")
	fmt.Printf("decoded transfer event: {From: 0x%x, To: 0x%x, Value: %v} ", Transfer.From, Transfer.To, Transfer.Value)
}

func mustGetAbi(filePath string) []byte {
	abi, err := ioutil.ReadFile(filePath)
	utils.PanicIfErr(err, "failed to get abi")
	return abi
}

func mustGetBytecode(filePath string) []byte {
	bytecodeHexStr, err := ioutil.ReadFile(filePath)
	utils.PanicIfErr(err, "failed to read bytecode")

	bytecode, err := hex.DecodeString(string(bytecodeHexStr))
	utils.PanicIfErr(err, "failed to decode bytecode")
	return bytecode
}
