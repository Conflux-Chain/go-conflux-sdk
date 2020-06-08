package main

import (
	"fmt"
	"math/big"
	"time"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
)

var am *sdk.AccountManager

func main() {
	//unlock account
	am = sdk.NewAccountManager("../keystore")
	err := am.TimedUnlockDefault("hello", 300*time.Second)
	if err != nil {
		panic(err)
	}

	//init client without retry and excute it
	url := "http://123.57.45.90:12537"
	// url := "http://testnet-jsonrpc.conflux-chain.org:12537"
	client, err := sdk.NewClient(url)
	if err != nil {
		panic(err)
	}
	client.SetAccountManager(am)
	fmt.Println("\n======================================")
	fmt.Println("start excute client methods without retry")
	run(client)

	//init client with retry and excute it
	client, err = sdk.NewClientWithRetry(url, 10, time.Second)
	if err != nil {
		panic(err)
	}
	client.SetAccountManager(am)
	fmt.Println("\n======================================")
	fmt.Println("start excute client methods with retry")
	run(client)
}

func run(client *sdk.Client) {
	gasPrice, err := client.GetGasPrice()
	if err != nil {
		fmt.Printf("gasprice error: %#v\n\n", err)
	} else {
		fmt.Printf("gasprice: %#v\n\n", gasPrice)
	}

	epochs := []*types.Epoch{nil, types.EpochLatestState, types.NewEpochNumber(big.NewInt(1061848))}
	for _, e := range epochs {
		nonce, err := client.GetNextNonce(types.Address("0x19f4bcf113e0b896d9b34294fd3da86b4adf0302"), e)
		if err != nil {
			fmt.Printf("nonce of epoch %v error: %v\n\n", e, err)
		} else {
			fmt.Printf("nonce of epoch %v : %v\n\n", e, nonce)
		}
	}

	status, err := client.GetStatus()
	if err != nil {
		fmt.Printf("get status error: %v\n\n", err)
	} else {
		fmt.Printf("get status result: %+v\n\n", status)
	}

	addr := types.Address("0x19f4bcf113e0b896d9b34294fd3da86b4adf0302")
	balance, err := client.GetBalance(addr)
	if err != nil {
		fmt.Printf("get balance of %v: %v\n\n", addr, err)
	} else {
		fmt.Printf("balance of %v: %#v\n\n", addr, balance)
	}

	bestBlockHash, err := client.GetBestBlockHash()
	if err != nil {
		fmt.Printf("get best block hash error: %v\n\n", err)
	} else {
		fmt.Printf("best block hash: %#v\n\n", bestBlockHash)
	}

	epochNumber, err := client.GetEpochNumber()
	if err != nil {
		fmt.Printf("get epochNumber error: %#v\n\n", err)
	} else {
		fmt.Printf("epochNumber: %#v\n\n", epochNumber)
	}

	epochNumber = big.NewInt(10)
	block, err := client.GetBlockByEpoch(types.NewEpochNumber(epochNumber))
	if err != nil {
		fmt.Printf("get block of epoch %v error:%#v\n\n", epochNumber, err)
	} else {
		fmt.Printf("block of epoch %v:%#v\n\n", epochNumber, block)
	}

	blocks, err := client.GetBlocksByEpoch(types.NewEpochNumber(epochNumber))
	if err != nil {
		fmt.Printf("get blocks of epoch %v error:%#v\n\n", epochNumber, err)
	} else {
		fmt.Printf("blocks of epoch %v:%#v\n\n", epochNumber, blocks)
	}

	blockHash := types.Hash("0xc3d83152acdc1296845be415a0b590627ce406dc355f0479398b99ca6dbc6a79")
	block, err = client.GetBlockByHash(blockHash)
	if err != nil {
		fmt.Printf("get block of hash %v error:%#v\n\n", blockHash, err)
	} else {
		fmt.Printf("block of hash %v:%#v\n\n", blockHash, block)
	}

	blockSummary, err := client.GetBlockSummaryByEpoch(types.NewEpochNumber(epochNumber))
	if err != nil {
		fmt.Printf("get block summary of epoch %v error:%#v\n\n", epochNumber, err)
	} else {
		fmt.Printf("block summary of epoch %v:%#v\n\n", epochNumber, blockSummary)
	}

	blockSummary, err = client.GetBlockSummaryByHash(blockHash)
	if err != nil {
		fmt.Printf("get block summary of block hash %v error:%#v\n\n", blockHash, err)
	} else {
		fmt.Printf("block summary of block hash %v:%#v\n\n", blockHash, blockSummary)
	}

	contractAddr := *types.NewAddress("0xa70ddf9b9750c575db453eea6a041f4c8536785a")
	code, err := client.GetCode(contractAddr)
	if err != nil {
		fmt.Printf("get code of address %v err: %v\n\n", contractAddr, err)
	} else {
		fmt.Printf("get code of address %v:%#v\n\n", contractAddr, code)
	}

	txhash := types.Hash("0x2234bb87229cf36481fdd58f632d2b5f573a62a968eb1fd341e98905e50c81e8")
	tx, err := client.GetTransactionByHash(txhash)
	if err != nil {
		fmt.Printf("get Transaction By Hash %v error:%v\n\n", txhash, err)
	} else {
		fmt.Printf("get Transaction By Hash %v:%#v\n\n", txhash, tx)
	}

	receipt, err := client.GetTransactionReceipt(txhash)
	if err != nil {
		fmt.Printf("transaction receipt of txhash %v error:%v\n\n", txhash, err)
	} else {
		fmt.Printf("transaction receipt of txhash %v:%#v\n\n", txhash, receipt)
	}

	rawtx := signTx(client)
	txhash, err = client.SendRawTransaction(rawtx)
	if err != nil {
		fmt.Printf("send Signed Transaction result error :%v\n\n", err)
	} else {
		fmt.Printf("send Signed Transaction result :%#v\n\n", txhash)
	}
	if err == nil {
		tx = waitTxBePacked(client, txhash)
	}
	time.Sleep(10 * time.Second)

	//send transaction
	utx, err := client.CreateUnsignedTransaction(types.Address("0x19f4bcf113e0b896d9b34294fd3da86b4adf0302"), types.Address("0x1cad0b19bb29d4674531d6f115237e16afce377d"), types.NewBigInt(1000000), nil)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("creat a new unsigned transaction %+v\n\n", utx)
	}
	txhash, err = client.SendTransaction(utx)
	if err != nil {
		fmt.Printf("send transaction error: %v\n\n", err)
	} else {
		fmt.Printf("send transaction done, tx hash is: %v\n\n", txhash)
	}
	if err == nil {
		tx = waitTxBePacked(client, txhash)
	}
	time.Sleep(10 * time.Second)

	if tx != nil {
		blockHash = *tx.BlockHash
		risk, err := client.GetRawBlockConfirmationRisk(blockHash)
		if err != nil {
			fmt.Printf("get risk of block %v error: %v\n", blockHash, err)
		} else {
			fmt.Printf("get risk of block %v : %v\n", blockHash, risk)
		}

		rate, err := client.GetBlockConfirmationRisk(blockHash)
		if err != nil {
			fmt.Printf("get revert rate of block %v error: %v\n", blockHash, err)
		} else {
			fmt.Printf("get revert rate of block %v : %v\n", blockHash, rate)
		}
	}

	b := new(types.Block)
	err = client.CallRPC(b, "cfx_getBlockByHash", "0x6d78977bf3882baf55ade9ffd4682754c66228bd42a6da4d2b5666544fe179bc", false)
	if err != nil {
		fmt.Printf("use CallRPC get block by hash error:%+v\n\n", err)
	} else {
		fmt.Printf("use CallRPC get block by hash result:%+v\n\n", b)
	}
}

func signTx(client *sdk.Client) []byte {
	unSignedTx := types.UnsignedTransaction{
		UnsignedTransactionBase: types.UnsignedTransactionBase{
			From:     types.NewAddress("0x19f4bcf113e0b896d9b34294fd3da86b4adf0302"),
			Value:    types.NewBigInt(100),
			GasPrice: types.NewBigInt(10000000000),
			// ChainID:  types.NewBigInt(1),
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
	fmt.Printf("signed tx %+v result:\n0x%x\n\n", unSignedTx, signedTx)
	return signedTx
}

func waitTxBePacked(client *sdk.Client, txhash types.Hash) *types.Transaction {
	fmt.Printf("wait be packed")
	var tx *types.Transaction
	var err error
	for {
		time.Sleep(time.Duration(1) * time.Second)
		tx, err = client.GetTransactionByHash(txhash)
		if err != nil {
			panic(err)
		}
		fmt.Print(".")
		if tx.Status != nil {
			fmt.Printf("\ntransaction is packed:%+v\n\n", tx)
			break
		}
	}
	return tx
}
