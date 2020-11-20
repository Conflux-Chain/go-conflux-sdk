package context

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"time"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
)

func JsonFmt(v interface{}) string {
	j, e := json.Marshal(v)
	if e != nil {
		panic(e)
	}
	var str bytes.Buffer
	_ = json.Indent(&str, j, "", "    ")
	return str.String()
}

func WaitPacked(client *sdk.Client, txhash types.Hash) *types.TransactionReceipt {
	fmt.Printf("wait for transaction %v be packed\n", txhash)
	var tx *types.TransactionReceipt
	for {
		time.Sleep(time.Duration(1) * time.Second)
		var err error
		tx, err = client.GetTransactionReceipt(txhash)
		if err != nil {
			panic(err)
		}
		fmt.Print(".")
		if tx != nil {
			fmt.Printf("transaction is packed:%+v\n\n", JsonFmt(tx))
			break
		}
	}
	return tx
}

func CreateSignedTx(client *sdk.Client) []byte {
	unSignedTx := types.UnsignedTransaction{
		UnsignedTransactionBase: types.UnsignedTransactionBase{
			// From:  types.NewAddress("0x19f4bcf113e0b896d9b34294fd3da86b4adf0302"),
			From:  defaultAccount,
			Value: types.NewBigInt(100),
			// GasPrice: types.NewBigInt(10000000000),
			// ChainID:  types.NewBigInt(1),
			Nonce: getNextNonceAndIncrease(),
		},
		To: types.NewAddress("0x10f4bcf113e0b896d9b34294fd3da86b4adf0302"),
	}
	err := client.ApplyUnsignedTransactionDefault(&unSignedTx)
	if err != nil {
		panic(err)
	}

	signedTx, err := am.SignAndEcodeTransactionWithPassphrase(unSignedTx, "hello")
	if err != nil {
		panic(err)
	}
	fmt.Printf("signed tx %v result:\n0x%x\n\n", JsonFmt(unSignedTx), signedTx)
	return signedTx
}

func DeployIfNotExist(contractAddress types.Address, abiFilePath string, bytecodeFilePath string) *sdk.Contract {
	isAddress := len(contractAddress) == 42 && (contractAddress)[0:2] == "0x"
	isCodeExist := false

	if isAddress {
		code, err := client.GetCode(contractAddress)
		// fmt.Printf("err: %v,code:%v\n", err, len(code))
		if err == nil && len(code) > 0 && code != "0x" {
			isCodeExist = true
		}
	}

	fmt.Printf("%v isAddress:%v, isCodeExist:%v\n", contractAddress, isAddress, isCodeExist)
	if isAddress && isCodeExist {
		abi, err := ioutil.ReadFile(abiFilePath)
		if err != nil {
			panic(err)
		}
		contract, err := client.GetContract(abi, &contractAddress)
		if err != nil {
			panic(err)
		}
		return contract
	}

	contract := deployContractWithConstroctor(abiFilePath, bytecodeFilePath, big.NewInt(100000), "biu", uint8(10), "BIU")
	return contract
}

func deployContractWithConstroctor(abiFile string, bytecodeFile string, params ...interface{}) *sdk.Contract {
	fmt.Println("start deploy contract with construcotr")
	abi, err := ioutil.ReadFile(abiFile)
	if err != nil {
		panic(err)
	}

	bytecodeHexStr, err := ioutil.ReadFile(bytecodeFile)
	if err != nil {
		panic(err)
	}

	bytecode, err := hex.DecodeString(string(bytecodeHexStr))
	if err != nil {
		panic(err)
	}

	option := types.ContractDeployOption{}
	option.Nonce = getNextNonceAndIncrease()
	result := client.DeployContract(&option, abi, bytecode, params...)

	_ = <-result.DoneChannel
	if result.Error != nil {
		panic(result.Error)
	}
	contract := result.DeployedContract
	fmt.Printf("deploy contract with abi: %v, bytecode: %v done\ncontract address: %+v\ntxhash:%v\n\n", abiFile, bytecodeFile, contract.Address, result.TransactionHash)

	return contract
}
