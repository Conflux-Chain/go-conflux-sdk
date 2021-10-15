package main

import (
	"fmt"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/accounts"
	client "github.com/Conflux-Chain/go-conflux-sdk/cfxclient"
	"github.com/Conflux-Chain/go-conflux-sdk/example/context"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	address "github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
)

func main() {

	//init client
	url := "http://39.97.232.99:12537"
	// url := "http://testnet-jsonrpc.conflux-chain.org:12537"
	_client, err := client.NewClient(url)
	context.PanicIfErrf(err, "failed to new client")

	networkId, err := _client.GetNetworkID()
	context.PanicIfErrf(err, "failed to get status")

	//unlock account
	wallet := accounts.NewKeystoreWallet("../keystore", uint32(networkId))
	err = wallet.TimedUnlockDefault("hello", 300*time.Second)
	if err != nil {
		panic(err)
	}

	client := client.NewSignableClient(&_client, wallet)

	//send transaction
	//send 0.01 cfx
	value := types.NewBigInt(1000000000000000)
	utx, err := client.NewTransaction(address.MustNewFromHex("0x19f4bcf113e0b896d9b34294fd3da86b4adf0302"), cfxaddress.MustNewFromHex("0x1cad0b19bb29d4674531d6f115237e16afce377d"), value, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("creat a new unsigned transaction %+v\n\n", utx)

	txhash, err := client.SignTransactionAndSend(utx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("send transaction hash: %v\n\n", txhash)

	fmt.Println("wait for transaction be packed")
	context.WaitPacked(client, txhash)
}
