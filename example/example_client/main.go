package main

import (
	"fmt"
	"math/big"
	"time"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/example/context"
	exampletypes "github.com/Conflux-Chain/go-conflux-sdk/example/context/types"
	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	address "github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var (
	am     sdk.AccountManagerOperator
	client *sdk.Client
	// retryClient    *sdk.Client
	config         *exampletypes.Config
	epochs         []*types.Epoch
	defaultAccount *types.Address
	nextNonce      *big.Int
)

func init() {
	config = context.PrepareForClientExample()
	am = config.GetAccountManager()
	defaultAccount, _ = am.GetDefault()

	epochs = []*types.Epoch{nil,
		types.EpochEarliest,
		types.EpochLatestCheckpoint,
		types.EpochLatestConfirmed,
		types.EpochLatestMined,
		types.EpochLatestState,
	}
}

func main() {

	fmt.Printf("\n=======start excute client methods without retry=========\n")
	run(config.GetClient())
	fmt.Printf("\n=======excute client methods without retry end!==========\n")

	fmt.Printf("\n=======start excute client methods with retry============\n")
	run(config.GetRetryClient())
	fmt.Printf("\n=======excute client methods with retry end!=============\n")
}

func run(_client *sdk.Client) {
	client = _client

	newAddress()
	mustNewAddress()

	getEpochNumber()
	getGasPrice()
	getNextNonce()
	getStatus()
	getBalance()
	getBestBlockHash()
	getBlockByEpoch()
	getBlocksByEpoch()
	getBlockByHash()
	getBlockSummaryByEpoch()
	getBlockSummaryByHash()
	getCode()
	getTransactionByHash()
	estimateGasAndCollateral()
	getTransactionReceipt()
	sendRawTransaction()
	createAndSendUnsignedTransaction()
	getRawBlockConfirmationRisk()
	getBlockConfirmationRisk()
	getSupplyInfo()
	callRPC()

	getAdmin()
	getSponsorInfo()
	getStakingBalance()
	getCollateralForStorage()
	getStorageAt()
	getStorageRoot()
	getBlockByHashWithPivotAssumption()
	checkBalanceAgainstTransaction()
	getSkippedBlocksByEpoch()
	getAccountInfo()
	getInterestRate()
	getAccumulateInterestRate()
	getBlockRewardInfo()
	getClientVersion()
	getEpochReceipts()
	getAccountPendingInfo()

	traceBlock()
	traceFilter()
	tarceTransaction()

	subscribeNewHeads()
	subscribeEpochs()
	subscribeLogs()

	batchCall()
}

func newAddress() {
	addr, err := client.NewAddress("0x0000000000000000000000000000000000000000")
	printResult("NewAddress", []interface{}{"0x0000000000000000000000000000000000000000"}, addr, err)
	addr, err = client.NewAddress("0x2000000000000000000000000000000000000000")
	printResult("NewAddress", []interface{}{"0x2000000000000000000000000000000000000000"}, addr, err)

	addr, err = client.NewAddress("cfxtest:aap9kthvctunvf030rbkk9k7zbzyz12dajp1u3sp4g")
	printResult("NewAddress", []interface{}{"cfxtest:aap9kthvctunvf030rbkk9k7zbzyz12dajp1u3sp4g"}, addr, err)
	addr, err = client.NewAddress("cfx:aap9kthvctunvf030rbkk9k7zbzyz12dajg6dkjg8p")
	printResult("NewAddress", []interface{}{"cfx:aap9kthvctunvf030rbkk9k7zbzyz12dajg6dkjg8p"}, addr, err)
}

func mustNewAddress() {
	addr := client.MustNewAddress("0x0000000000000000000000000000000000000000")
	printResult("MustNewAddress", []interface{}{"0x0000000000000000000000000000000000000000"}, addr, nil)
	// addr = client.MustNewAddress("0x2000000000000000000000000000000000000000")
	// printResult("MustNewAddress", []interface{}{"0x2000000000000000000000000000000000000000"}, addr, nil)

	addr = client.MustNewAddress("cfxtest:aap9kthvctunvf030rbkk9k7zbzyz12dajp1u3sp4g")
	printResult("MustNewAddress", []interface{}{"cfxtest:aap9kthvctunvf030rbkk9k7zbzyz12dajp1u3sp4g"}, addr, nil)
	// addr = client.MustNewAddress("cfx:aap9kthvctunvf030rbkk9k7zbzyz12dajg6dkjg8p")
	// printResult("MustNewAddress", []interface{}{"cfx:aap9kthvctunvf030rbkk9k7zbzyz12dajg6dkjg8p"}, addr, nil)
}

func getAdmin() {
	result, err := client.GetAdmin(config.ERC20Address, nil)
	printResult("GetAdmin", []interface{}{config.ERC20Address, nil}, result, err)

	result, err = client.GetAdmin(address.MustNewFromHex("0x0000000000000000000000000000000000000000"), nil)
	printResult("GetAdmin", []interface{}{address.MustNewFromHex("0x0000000000000000000000000000000000000000"), nil}, result, err)
}

func getSponsorInfo() {
	// result, err := client.GetSponsorInfo(*defaultAccount, nil)
	result, err := client.GetSponsorInfo(config.ERC20Address, nil)
	printResult("GetSponsorInfo", []interface{}{config.ERC20Address, nil}, result, err)
}

func getStakingBalance() {
	result, err := client.GetStakingBalance(*defaultAccount, nil)
	printResult("GetStakingBalance", []interface{}{*defaultAccount, nil}, result, err)
}

func getCollateralForStorage() {
	result, err := client.GetCollateralForStorage(*defaultAccount, nil)
	printResult("GetCollateralForStorage", []interface{}{*defaultAccount, nil}, result, err)
}

func getStorageAt() {
	result, err := client.GetStorageAt(config.ERC20Address, "0x8549225e0f8e0f4a2ea0d9c0e562e986994ded65da69d91aa3768ac6da0a1635", nil)
	printResult("GetStorageAt", []interface{}{config.ERC20Address, "0x8549225e0f8e0f4a2ea0d9c0e562e986994ded65da69d91aa3768ac6da0a1635", nil}, result, err)
}

func getStorageRoot() {
	result, err := client.GetStorageRoot(config.ERC20Address, nil)
	printResult("GetStorageRoot", []interface{}{config.ERC20Address, nil}, result, err)
}

func getBlockByHashWithPivotAssumption() {
	result, err := client.GetBlockByHashWithPivotAssumption(types.Hash("0x08de0feea8cc989029f86a00ef6aabbf4de16d9bf21207c8ba9f011f10b1456d"), types.Hash("0x8cf781d04606e195f7fc5e03a73d8e2ef5bf7d9bfba11b11e73cd056f190c67a"), hexutil.Uint64(0x176334))
	printResult("GetBlockByHashWithPivotAssumption", []interface{}{types.Hash("0x08de0feea8cc989029f86a00ef6aabbf4de16d9bf21207c8ba9f011f10b1456d"), types.Hash("0x8cf781d04606e195f7fc5e03a73d8e2ef5bf7d9bfba11b11e73cd056f190c67a"), hexutil.Uint64(0x176334)}, result, err)
}

func checkBalanceAgainstTransaction() {
	result, err := client.CheckBalanceAgainstTransaction(*defaultAccount, config.ERC20Address, types.NewBigInt(1000), types.NewBigInt(1000), types.NewBigInt(1000), nil)
	printResult("CheckBalanceAgainstTransaction", []interface{}{*defaultAccount, *defaultAccount, types.NewBigInt(1000), types.NewBigInt(1000), types.NewBigInt(1000), types.EpochLatestState}, result, err)
}

func getSkippedBlocksByEpoch() {
	result, err := client.GetSkippedBlocksByEpoch(types.EpochLatestState)
	printResult("GetSkippedBlocksByEpoch", []interface{}{nil}, result, err)
}

func getAccountInfo() {
	result, err := client.GetAccountInfo(*defaultAccount, nil)
	printResult("GetAccountInfo", []interface{}{*defaultAccount, nil}, result, err)
}

// GetInterestRate()
func getInterestRate() {
	result, err := client.GetInterestRate(nil)
	printResult("GetInterestRate", []interface{}{nil}, result, err)
}

// GetAccumulateInterestRate()
func getAccumulateInterestRate() {
	result, err := client.GetAccumulateInterestRate(nil)
	printResult("GetAccumulateInterestRate", []interface{}{nil}, result, err)
}

// GetBlockRewardInfo()
func getBlockRewardInfo() {
	result, err := client.GetBlockRewardInfo(*types.EpochLatestState)
	printResult("GetBlockRewardInfo", []interface{}{*types.EpochLatestState}, result, err)
}

// ClientVersion()
func getClientVersion() {
	result, err := client.GetClientVersion()
	printResult("ClientVersion", []interface{}{}, result, err)
}

func getEpochNumber() {
	fmt.Println("- start get epoch number")
	for _, e := range epochs {
		epochnumber, err := client.GetEpochNumber(e)
		if err != nil {
			fmt.Printf("- get epoch %v error: %v\n\n", e, err.Error())
		} else {
			fmt.Printf("	epoch of %v : %v\n\n", e, epochnumber)
		}
	}
}

func getGasPrice() {

	gasPrice, err := client.GetGasPrice()
	if err != nil {
		fmt.Printf("- gasprice error: %#v\n\n", err.Error())
	} else {
		fmt.Printf("- gasprice: %v\n\n", gasPrice.String())
	}

}

func getNextNonce() {
	fmt.Println("- start get nextNonce")
	for _, e := range epochs {
		nonce, err := client.GetNextNonce(*defaultAccount, e)
		if err != nil {
			fmt.Printf("	nonce of epoch %v error: %v\n\n", e, err.Error())
		} else {
			fmt.Printf("	nonce of epoch %v : %v\n\n", e, nonce)
		}
	}
}

func getStatus() {
	status, err := client.GetStatus()
	if err != nil {
		fmt.Printf("- get status error: %v\n\n", err.Error())
	} else {
		fmt.Printf("- get status result:\n%v\n\n", context.JSONFmt(status))
	}

}

func getBalance() {

	addr := *defaultAccount
	balance, err := client.GetBalance(addr)
	if err != nil {
		fmt.Printf("- get balance of %v: %v\n\n", addr, err.Error())
	} else {
		fmt.Printf("- balance of %v: %#v\n\n", addr, balance)
	}
}

func getBestBlockHash() {
	bestBlockHash, err := client.GetBestBlockHash()
	if err != nil {
		fmt.Printf("- get best block hash error: %v\n\n", err.Error())
	} else {
		fmt.Printf("- best block hash: %#v\n\n", bestBlockHash)
	}
}

func getBlockByEpoch() {
	epochNumber, err := client.GetEpochNumber()
	block, err := client.GetBlockByEpoch(types.NewEpochNumber(epochNumber))
	if err != nil {
		fmt.Printf("- get block of epoch %v error:%#v\n\n", epochNumber, err.Error())
	} else {
		fmt.Printf("- block of epoch %v:\n%v\n\n", epochNumber, context.JSONFmt(block))
	}
}

func getBlocksByEpoch() {
	epochNumber, err := client.GetEpochNumber()
	blocks, err := client.GetBlocksByEpoch(types.NewEpochNumber(epochNumber))
	if err != nil {
		fmt.Printf("- get blocks of epoch %v error:%#v\n\n", epochNumber, err.Error())
	} else {
		fmt.Printf("- blocks of epoch %v:%#v\n\n", epochNumber, blocks)
	}

}

func getBlockByHash() {
	blockHash := types.Hash(config.BlockHash)
	block, err := client.GetBlockByHash(blockHash)
	if err != nil {
		fmt.Printf("- get block of hash %v error:\n%#v\n\n", blockHash, err.Error())
	} else {
		fmt.Printf("- block of hash %v:\n%v\n\n", blockHash, context.JSONFmt(block))
	}
}

func getBlockSummaryByEpoch() {
	epochNumber, err := client.GetEpochNumber()
	blockSummary, err := client.GetBlockSummaryByEpoch(types.NewEpochNumber(epochNumber))
	if err != nil {
		fmt.Printf("- get block summary of epoch %v error:%#v\n\n", epochNumber, err.Error())
	} else {
		fmt.Printf("- block summary of epoch %v:\n%v\n\n", epochNumber, context.JSONFmt(blockSummary))
	}
}

func getBlockSummaryByHash() {

	blockHash := types.Hash(config.BlockHash)
	blockSummary, err := client.GetBlockSummaryByHash(blockHash)
	if err != nil {
		fmt.Printf("- get block summary of block hash %v error:%#v\n\n", blockHash, err.Error())
	} else {
		fmt.Printf("- block summary of block hash %v:\n%v\n\n", blockHash, context.JSONFmt(blockSummary))
	}
}

func getCode() {
	contractAddr := *defaultAccount // config.ERC20Address
	// code, err := client.GetCode(contractAddr)
	code, err := client.GetCode(address.MustNewFromHex("0x19f4bcf113e0b896d9b34294fd3da86b4adf0301"))
	if err != nil {
		fmt.Printf("- get code of address %v err: %v\n\n", contractAddr, err.Error())
	} else {
		fmt.Printf("- get code of address %v:%v\n\n", contractAddr, code)
	}
}

func getTransactionByHash() {
	txhash := types.Hash(config.TransactionHash)
	tx, err := client.GetTransactionByHash(txhash)
	if err != nil {
		fmt.Printf("- get Transaction By Hash %v error:%v\n\n", txhash, err.Error())
	} else {
		fmt.Printf("- get Transaction By Hash %v:\n%v\n\n", txhash, context.JSONFmt(tx))
	}
}

func getTransactionReceipt() {
	txhash := types.Hash(config.TransactionHash)
	receipt, err := client.GetTransactionReceipt(txhash)
	if err != nil {
		fmt.Printf("- transaction receipt of txhash %v error:%v\n\n", txhash, err.Error())
	} else {
		fmt.Printf("- transaction receipt of txhash %v:\n%v\n\n", txhash, context.JSONFmt(receipt))
	}
}

func estimateGasAndCollateral() {
	to := cfxaddress.MustNewFromHex("0x10f4bcf113e0b896d9b34294fd3da86b4adf0302")
	request := types.CallRequest{
		To:    &to,
		Value: types.NewBigInt(1),
	}
	est, err := client.EstimateGasAndCollateral(request)
	if err != nil {
		fmt.Printf("- estimate request %v error: %v\n\n", request, err.Error())
	} else {
		fmt.Printf("- estimate result: %v\n\n", context.JSONFmt(est))
	}
}

func sendRawTransaction() {
	rawtx := context.CreateSignedTx(client)
	txhash, err := client.SendRawTransaction(rawtx)
	if err != nil {
		fmt.Printf("- send Signed Transaction result error :%v\n\n", err.Error())
	} else {
		fmt.Printf("- send Signed Transaction result :%#v\n\n", txhash)
	}
	if err == nil {
		context.WaitPacked(client, txhash)
	}
	// time.Sleep(10 * time.Second)
}

func createAndSendUnsignedTransaction() {
	//send transaction
	chainID, err := client.GetNetworkID()
	context.PanicIfErrf(err, "failed to get chainID")
	utx, err := client.CreateUnsignedTransaction(*defaultAccount, cfxaddress.MustNewFromHex("0x1cad0b19bb29d4674531d6f115237e16afce377d", chainID), types.NewBigInt(1000000), nil)
	if err != nil {
		panic(err)
	}
	utx.Nonce = context.GetNextNonceAndIncrease()
	fmt.Printf("- creat a new unsigned transaction:\n%v\n\n", context.JSONFmt(utx))

	txhash, err := client.SendTransaction(utx)
	if err != nil {
		fmt.Printf("- send transaction error: %v\n\n", err.Error())
	} else {
		fmt.Printf("- send transaction done, tx hash is: %v\n\n", txhash)
	}
	if err == nil {
		context.WaitPacked(client, txhash)
	}
	// time.Sleep(10 * time.Second)
}

func getRawBlockConfirmationRisk() {
	risk, err := client.GetRawBlockConfirmationRisk(config.BlockHash)
	if err != nil {
		fmt.Printf("- get risk of block %v error: %v\n", config.BlockHash, err.Error())
	} else {
		fmt.Printf("- get risk of block %v : %v\n", config.BlockHash, risk)
	}

	risk, err = client.GetRawBlockConfirmationRisk(types.Hash("0x0000000000000000000000000000000000000000000000000000000000000000"))
	if err != nil {
		fmt.Printf("- get risk of block %v error: %v\n", config.BlockHash, err.Error())
	} else {
		fmt.Printf("- get risk of block %v : %v\n", config.BlockHash, risk)
	}

}

func getBlockConfirmationRisk() {
	rate, err := client.GetBlockConfirmationRisk(config.BlockHash)
	if err != nil {
		fmt.Printf("- get revert rate of block %v error: %v\n", config.BlockHash, err.Error())
	} else {
		fmt.Printf("- get revert rate of block %v : %v\n", config.BlockHash, rate)
	}
}

func getSupplyInfo() {
	info, err := client.GetSupplyInfo()
	if err != nil {
		fmt.Printf("- get supply info error: %v\n", err.Error())
	} else {
		fmt.Printf("- get supply info :%v \n", context.JSONFmt(info))
	}
}

func getEpochReceipts() {
	b, _ := client.GetBlockByHash(config.BlockHashOfNewContract)

	receipts, err := client.GetEpochReceipts(*types.NewEpochNumber(b.EpochNumber))
	if err != nil {
		fmt.Printf("- get epoch receipts error: %+v\n", err)
	} else {
		fmt.Printf("- get rpoch receipts info:%v \n", context.JSONFmt(receipts))
	}
}

func getAccountPendingInfo() {
	fmt.Println("default account:", *defaultAccount)
	result, err := client.GetAccountPendingInfo(*defaultAccount)
	printResult("GetAccountPendingInfo", []interface{}{*defaultAccount}, result, err)
}

func traceBlock() {
	traces, err := client.GetBlockTraces(config.BlockHashOfNewContract)
	if err != nil {
		fmt.Printf("- get block trace of create error: %+v\n", err)
	} else {
		fmt.Printf("- get block info of create: %+v\n", context.JSONFmt(traces))
	}

	traces, err = client.GetBlockTraces(config.BlockHash)
	if err != nil {
		fmt.Printf("- get block trace of call error: %+v\n", err)
	} else {
		fmt.Printf("- get block info of call: %+v\n", context.JSONFmt(traces))
	}
}

func traceFilter() {
	traces, err := client.FilterTraces(types.TraceFilter{
		FromEpoch:   types.NewEpochNumberUint64(1),
		ActionTypes: []string{"call", "create"},
	})
	if err != nil {
		fmt.Printf("- filter trace error: %+v\n", err)
	} else {
		fmt.Printf("- filter trace result: %+v", context.JSONFmt(traces))
	}
}

func tarceTransaction() {
	traces, err := client.GetTransactionTraces(config.TransactionHash)
	if err != nil {
		fmt.Printf("- get transcation trace of create error: %+v\n", err)
	} else {
		fmt.Printf("- get transcation trace of create: %+v", context.JSONFmt(traces))
	}
}

func callRPC() {
	b := new(types.Block)
	err := client.CallRPC(b, "cfx_getBlockByHash", config.BlockHash, true)
	if err != nil {
		fmt.Printf("- use CallRPC get block by hash error:%+v\n\n", err.Error())
	} else {
		fmt.Printf("- use CallRPC get block by hash result:\n%v\n\n", context.JSONFmt(b))
	}
}

func subscribeNewHeads() {
	fmt.Printf("- subscribe new heads\n")
	channel := make(chan types.BlockHeader, 100)
	sub, err := client.SubscribeNewHeads(channel)
	if err != nil {
		fmt.Printf("subscribe block header error:%+v\n", err.Error())
		return
	}

	errorchan := sub.Err()
	for i := 0; i < 10; i++ {
		select {
		case err = <-errorchan:
			fmt.Printf("subscription new heads error:%v\n", err.Error())
			// sub.Unsubscribe()
			// return
		case head := <-channel:
			fmt.Printf("received new header:%+v\n", head)
		}
	}
	sub.Unsubscribe()
}

func subscribeEpochs() {
	fmt.Printf("- subscribe epochs\n")
	channel := make(chan types.WebsocketEpochResponse, 100)
	sub, err := client.SubscribeEpochs(channel)
	if err != nil {
		fmt.Printf("subscribe epoch error:%+v\n", err.Error())
		return
	}

	errorchan := sub.Err()
	for i := 0; i < 10; i++ {
		select {
		case err = <-errorchan:
			fmt.Printf("subscription epoch error:%v\n", err.Error())
			// sub.Unsubscribe()
			// return
		case epoch := <-channel:
			fmt.Printf("received new epoch:%+v\n", epoch)
		}
	}
	sub.Unsubscribe()
}

func subscribeLogs() {
	fmt.Printf("- subscribe logs\n")
	logChannel := make(chan types.Log, 100)
	reorgChannel := make(chan types.ChainReorg, 100)
	sub, err := client.SubscribeLogs(logChannel, reorgChannel, types.LogFilter{
		Topics: [][]types.Hash{{"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"}}})
	if err != nil {
		fmt.Printf("subscribe log error:%+v\n", err.Error())
		return
	}

	go func() {
		time.Sleep(time.Second)
		for i := 0; i < 10; i++ {
			// send an erc tx
			contract, err := config.GetErc20Contract()
			if err != nil {
				panic(fmt.Sprintf("get erc20 contract error: %v", err.Error()))
			}
			//send transction for contract method
			to := cfxaddress.MustNewFromHex("0x160ebef20c1f739957bf9eecd040bce699cc42c6")
			nonce := context.GetNextNonceAndIncrease()
			txhash, err := contract.SendTransaction(&types.ContractMethodSendOption{
				Nonce: nonce,
			}, "transfer", to.MustGetCommonAddress(), big.NewInt(10))
			if err != nil {
				panic(err)
			}
			fmt.Printf("transfer %v erc20 token to %v done, tx hash: %v\n", 10, to, txhash)
			receipt, _ := client.WaitForTransationReceipt(txhash, time.Second)
			if receipt != nil {
				fmt.Printf("tx %v is executed.\n", txhash)
			}
		}
	}()

	errorchan := sub.Err()

	for i := 0; i < 10; i++ {
		select {
		case err = <-errorchan:
			fmt.Printf("subscription error:%v\n", err.Error())
			// sub.Unsubscribe()
			// return
		case log := <-logChannel:
			fmt.Printf("received new log:%+v\n\n", log)
		case reorg := <-reorgChannel:
			fmt.Printf("received new reorg:%+v\n\n", reorg)
		}
	}
	sub.Unsubscribe()
}

func batchCall() {
	elems := make([]rpc.BatchElem, 2)
	elems[0] = rpc.BatchElem{Method: "cfx_epochNumber", Result: &hexutil.Big{}, Args: []interface{}{}}
	elems[1] = rpc.BatchElem{Method: "cfx_getBalance", Result: &hexutil.Big{}, Args: []interface{}{client.MustNewAddress("cfxtest:aap9kthvctunvf030rbkk9k7zbzyz12dajp1u3sp4g")}}
	err := client.BatchCallRPC(elems)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("batch call rpc done:%+v\n", elems)
	}
}

func printResult(method string, args []interface{}, result interface{}, err error) {
	if err != nil {
		fmt.Printf("- call method %v with args %+v error: %v\n\n", method, args, err.Error())
	} else {
		fmt.Printf("- call method %v with args %+v result: %+v\n\n", method, args, context.JSONFmt(result))
	}
}
