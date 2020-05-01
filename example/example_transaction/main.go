package main

import (
	"fmt"
	"math/big"
	"time"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

	//send transaction
	//send 0.01 cfx
	tmp, _ := new(big.Int).SetString("1000000000000000", 10)
	value := hexutil.Big(*tmp)
	utx, err := client.CreateUnsignedTransaction(types.Address("0x19f4bcf113e0b896d9b34294fd3da86b4adf0302"), types.Address("0x1cad0b19bb29d4674531d6f115237e16afce377d"), &value, &[]byte{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("creat a new unsigned transaction %+v\n\n", utx)
	txhash, err := client.SendTransaction(utx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("send transaction hash: %v\n\n", txhash)

	fmt.Println("wait for transaction be packed")
	for {
		time.Sleep(time.Duration(1) * time.Second)
		tx, err := client.GetTransactionByHash(txhash)
		if err != nil {
			panic(err)
		}
		if tx.Status != nil {
			fmt.Printf("transaction is packed:%+v\n\n", tx)
			break
		}
	}
}
