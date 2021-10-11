package main

import (
	"fmt"
	"math/big"
	"time"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/cfxclient"
	"github.com/Conflux-Chain/go-conflux-sdk/example/context"
	exampletypes "github.com/Conflux-Chain/go-conflux-sdk/example/context/types"
	"github.com/Conflux-Chain/go-conflux-sdk/middleware"
	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/utils"

	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var (
	am     sdk.Wallet
	client cfxclient.SignableClient
	// retryClient    *sdk.Client
	config         *exampletypes.Config
	epochs         []*types.Epoch
	defaultAccount *types.Address
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
	config.GetClient().UseCallRpcMiddleware(middleware.CallRpcConsoleMiddleware)
	config.GetClient().UseBatchCallRpcMiddleware(middleware.BatchCallRpcConsoleMiddleware)

	config.GetRetryClient().UseCallRpcMiddleware(middleware.CallRpcConsoleMiddleware)
	config.GetRetryClient().UseBatchCallRpcMiddleware(middleware.BatchCallRpcConsoleMiddleware)

	fmt.Printf("\n=======start excute client methods without retry=========\n")
	run(*config.GetClient())
	fmt.Printf("\n=======excute client methods without retry end!==========\n")

	fmt.Printf("\n=======start excute client methods with retry============\n")
	run(*config.GetRetryClient())
	fmt.Printf("\n=======excute client methods with retry end!=============\n")
}

func run(_client cfxclient.SignableClient) {
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
	GetBlockByBlockNumber()
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
	getEpochReceiptsByPivotBlockHash()
	getAccountPendingInfo()
	getAccountPendingTransactions()

	traceBlock()
	traceFilter()
	tarceTransaction()

	subscribeNewHeads()
	subscribeEpochs()
	subscribeLogs()

	batchCall()
}

func newAddress() {
	fmt.Println("\n- start new address")
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
	fmt.Println("\n- start must new address")
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
	fmt.Println("\n- start get admin")
	client.GetAdmin(config.ERC20Address, nil)
	// printResult("GetAdmin", []interface{}{config.ERC20Address, nil}, result, err)

	client.GetAdmin(cfxaddress.MustNewFromHex("0x0000000000000000000000000000000000000000"), nil)
	// printResult("GetAdmin", []interface{}{address.MustNewFromHex("0x0000000000000000000000000000000000000000"), nil}, result, err)
}

func getSponsorInfo() {
	fmt.Println("\n- start get sponsor info")
	// result, err := client.GetSponsorInfo(*defaultAccount, nil)
	client.GetSponsorInfo(config.ERC20Address, nil)
	// printResult("GetSponsorInfo", []interface{}{config.ERC20Address, nil}, result, err)
}

func getStakingBalance() {
	fmt.Println("\n- start get staking balance")
	client.GetStakingBalance(*defaultAccount, nil)
	// printResult("GetStakingBalance", []interface{}{*defaultAccount, nil}, result, err)
}

func getCollateralForStorage() {
	fmt.Println("\n- start get collateral for storage")
	client.GetCollateralForStorage(*defaultAccount, nil)
	// printResult("GetCollateralForStorage", []interface{}{*defaultAccount, nil}, result, err)
}

func getStorageAt() {
	fmt.Println("\n- start get storage at")
	client.GetStorageAt(config.ERC20Address, "0x8549225e0f8e0f4a2ea0d9c0e562e986994ded65da69d91aa3768ac6da0a1635", nil)
	// printResult("GetStorageAt", []interface{}{config.ERC20Address, "0x8549225e0f8e0f4a2ea0d9c0e562e986994ded65da69d91aa3768ac6da0a1635", nil}, result, err)
}

func getStorageRoot() {
	fmt.Println("\n- start get storage root")
	client.GetStorageRoot(config.ERC20Address, nil)
	// printResult("GetStorageRoot", []interface{}{config.ERC20Address, nil}, result, err)
}

func getBlockByHashWithPivotAssumption() {
	fmt.Println("\n- start get block hash with pivot assumption")
	client.GetBlockByHashWithPivotAssumption(types.Hash("0x08de0feea8cc989029f86a00ef6aabbf4de16d9bf21207c8ba9f011f10b1456d"), types.Hash("0x8cf781d04606e195f7fc5e03a73d8e2ef5bf7d9bfba11b11e73cd056f190c67a"), hexutil.Uint64(0x176334))
	// printResult("GetBlockByHashWithPivotAssumption", []interface{}{types.Hash("0x08de0feea8cc989029f86a00ef6aabbf4de16d9bf21207c8ba9f011f10b1456d"), types.Hash("0x8cf781d04606e195f7fc5e03a73d8e2ef5bf7d9bfba11b11e73cd056f190c67a"), hexutil.Uint64(0x176334)}, result, err)
}

func checkBalanceAgainstTransaction() {
	fmt.Println("\n- start check balance against transaction")
	client.CheckBalanceAgainstTransaction(*defaultAccount, config.ERC20Address, types.NewBigInt(1000), types.NewBigInt(1000), types.NewBigInt(1000), nil)
	// printResult("CheckBalanceAgainstTransaction", []interface{}{*defaultAccount, *defaultAccount, types.NewBigInt(1000), types.NewBigInt(1000), types.NewBigInt(1000), types.EpochLatestState}, result, err)
}

func getSkippedBlocksByEpoch() {
	fmt.Println("\n- start get skipped blocks by epoch")
	client.GetSkippedBlocksByEpoch(types.EpochLatestState)
	// printResult("GetSkippedBlocksByEpoch", []interface{}{nil}, result, err)
}

func getAccountInfo() {
	fmt.Println("\n- start get account info")
	client.GetAccountInfo(*defaultAccount, nil)
	// printResult("GetAccountInfo", []interface{}{*defaultAccount, nil}, result, err)
}

// GetInterestRate()
func getInterestRate() {
	fmt.Println("\n- start get interest rate")
	client.GetInterestRate(nil)
	// printResult("GetInterestRate", []interface{}{nil}, result, err)
}

// GetAccumulateInterestRate()
func getAccumulateInterestRate() {
	fmt.Println("\n- start get accumulate interest rate")
	client.GetAccumulateInterestRate(nil)
	// printResult("GetAccumulateInterestRate", []interface{}{nil}, result, err)
}

// GetBlockRewardInfo()
func getBlockRewardInfo() {
	fmt.Println("\n- start get block reward info")
	client.GetBlockRewardInfo(*types.EpochLatestState)
	// printResult("GetBlockRewardInfo", []interface{}{*types.EpochLatestState}, result, err)
}

// ClientVersion()
func getClientVersion() {
	fmt.Println("\n- start get client version")
	client.GetClientVersion()
	// printResult("ClientVersion", []interface{}{}, result, err)
}

func getEpochNumber() {
	fmt.Println("\n- start get epoch number")
	for _, e := range epochs {
		client.GetEpochNumber(e)
	}
}

func getGasPrice() {
	fmt.Println("\n- start get gas price")
	client.GetGasPrice()
}

func getNextNonce() {
	fmt.Println("\n- start get nextNonce")
	for _, e := range epochs {
		client.GetNextNonce(*defaultAccount, e)
	}
}

func getStatus() {
	fmt.Println("\n- start get status")
	client.GetStatus()
}

func getBalance() {
	fmt.Println("\n- start get balance")
	addr := *defaultAccount
	client.GetBalance(addr)
}

func getBestBlockHash() {
	fmt.Println("\n- start get best block hash")
	client.GetBestBlockHash()
}

func getBlockByEpoch() {
	fmt.Println("\n- start get block by epoch")
	if epochNumber, err := client.GetEpochNumber(); err == nil {
		client.GetBlockByEpoch(types.NewEpochNumber(epochNumber))
	}
}

func GetBlockByBlockNumber() {
	fmt.Println("\n- start get block by block number")

	b, err := client.GetBlockByHash(config.BlockHash)
	utils.PanicIfErr(err)

	_blockNumber := hexutil.Uint64(b.BlockNumber.ToInt().Uint64())
	client.GetBlockByBlockNumber(_blockNumber)
	client.GetBlockSummaryByBlockNumber(_blockNumber)
}

func getBlocksByEpoch() {
	fmt.Println("\n- start get blocks by epoch")
	if epochNumber, err := client.GetEpochNumber(); err == nil {
		client.GetBlocksByEpoch(types.NewEpochNumber(epochNumber))
	}
}

func getBlockByHash() {
	fmt.Println("\n- start get block by hash")
	blockHash := types.Hash(config.BlockHash)
	client.GetBlockByHash(blockHash)
}

func getBlockSummaryByEpoch() {
	fmt.Println("\n- start get block summary by epoch")
	if epochNumber, err := client.GetEpochNumber(); err == nil {
		client.GetBlockSummaryByEpoch(types.NewEpochNumber(epochNumber))
	}
}

func getBlockSummaryByHash() {
	fmt.Println("\n- start get block summary by hash")
	blockHash := types.Hash(config.BlockHash)
	client.GetBlockSummaryByHash(blockHash)
}

func getCode() {
	fmt.Println("\n- start get code")
	contractAddr := *defaultAccount // config.ERC20Address
	client.GetCode(contractAddr)
	client.GetCode(cfxaddress.MustNewFromHex("0x19f4bcf113e0b896d9b34294fd3da86b4adf0301"))
}

func getTransactionByHash() {
	fmt.Println("\n- start get transaction by hash")
	txhash := types.Hash(config.TransactionHash)
	client.GetTransactionByHash(txhash)
}

func getTransactionReceipt() {
	fmt.Println("\n- start get transaction receipt")
	txhash := types.Hash(config.TransactionHash)
	client.GetTransactionReceipt(txhash)
}

func estimateGasAndCollateral() {
	fmt.Println("\n- start estimate gas and collateral")
	to := cfxaddress.MustNewFromHex("0x10f4bcf113e0b896d9b34294fd3da86b4adf0302")
	request := types.CallRequest{
		To:    &to,
		Value: types.NewBigInt(1),
	}
	client.EstimateGasAndCollateral(request)
}

func sendRawTransaction() {
	fmt.Println("\n- start send raw transaction")
	rawtx := context.CreateSignedTx(&client)
	txhash, err := client.SendRawTransaction(rawtx)
	if err == nil {
		context.WaitPacked(client, txhash)
	}
	// time.Sleep(10 * time.Second)
}

func createAndSendUnsignedTransaction() {
	//send transaction
	fmt.Println("\n- start create and send unsigned transaction")
	chainID, err := client.GetNetworkID()
	context.PanicIfErrf(err, "failed to get chainID")

	utx, err := client.NewTransaction(*defaultAccount, cfxaddress.MustNewFromHex("0x1cad0b19bb29d4674531d6f115237e16afce377d", chainID), types.NewBigInt(1000000), nil)
	context.PanicIfErrf(err, "failed to creat unsigned tx")

	utx.Nonce = context.GetNextNonceAndIncrease()
	fmt.Printf("- creat a new unsigned transaction:\n%v\n\n", context.JSONFmt(utx))

	txhash, err := client.SignTransactionAndSend(utx)
	if err == nil {
		context.WaitPacked(client, txhash)
	}
	// time.Sleep(10 * time.Second)
}

func getRawBlockConfirmationRisk() {
	fmt.Println("\n- start get raw block confirmation risk")
	client.GetRawBlockConfirmationRisk(config.BlockHash)
	client.GetRawBlockConfirmationRisk(types.Hash("0x0000000000000000000000000000000000000000000000000000000000000000"))
}

func getBlockConfirmationRisk() {
	fmt.Println("\n- start get block confirmation risk")
	rate, err := client.GetBlockConfirmationRisk(config.BlockHash)
	if err != nil {
		fmt.Printf("- get revert rate of block %v error: %v\n", config.BlockHash, err.Error())
	} else {
		fmt.Printf("- get revert rate of block %v : %v\n", config.BlockHash, rate)
	}
}

func getSupplyInfo() {
	fmt.Println("\n- start get supply info")
	client.GetSupplyInfo()
}

func getEpochReceipts() {
	fmt.Println("\n- start get epoch receipts")
	b, _ := client.GetBlockByHash(config.BlockHashOfNewContract)
	client.GetEpochReceipts(*types.NewEpochNumber(b.EpochNumber))
}

func getEpochReceiptsByPivotBlockHash() {
	fmt.Println("\n- start get epoch receipts by block hash")
	client.GetEpochReceiptsByPivotBlockHash(config.BlockHashOfNewContract)
}

func getAccountPendingInfo() {
	fmt.Println("\n- start get account pending info")
	fmt.Println("default account:", *defaultAccount)
	client.GetAccountPendingInfo(*defaultAccount)
	// printResult("GetAccountPendingInfo", []interface{}{*defaultAccount}, result, err)
}

func getAccountPendingTransactions() {
	fmt.Println("\n- start get account pending transactions")
	client.GetAccountPendingTransactions(*defaultAccount, types.NewBigInt(0), types.NewUint64(100))

	to := client.MustNewAddress("cfxtest:aasm4c231py7j34fghntcfkdt2nm9xv1tyce66w5u3")
	utx, _ := client.NewTransaction(*defaultAccount, to, types.NewBigInt(1000000), nil)
	utx.Nonce = context.GetNextNonceAndIncrease()

	client.SignTransactionAndSend(utx)
	client.GetAccountPendingTransactions(*defaultAccount, types.NewBigInt(0), types.NewUint64(100))
	// for avoiding block following tests
	// client.WaitForTransationReceipt(hash, time.Second*2)
}

func traceBlock() {
	fmt.Println("\n- start get block trace")
	client.GetBlockTraces(config.BlockHashOfNewContract)
	client.GetBlockTraces(config.BlockHash)
}

func traceFilter() {
	fmt.Println("\n- start trace filter")
	client.FilterTraces(types.TraceFilter{
		FromEpoch:   types.NewEpochNumberUint64(1),
		ActionTypes: []string{"call", "create"},
	})
}

func tarceTransaction() {
	fmt.Println("\n- start trace transaction")
	client.GetTransactionTraces(config.TransactionHash)
}

func callRPC() {
	fmt.Println("\n- start call rpc")
	b := new(types.Block)
	client.CallRPC(b, "cfx_getBlockByHash", config.BlockHash, true)

}

func subscribeNewHeads() {
	fmt.Printf("\n- subscribe new heads\n")
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
	fmt.Printf("\n- subscribe epochs\n")
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
	fmt.Printf("\n- subscribe logs\n")
	channel := make(chan types.SubscriptionLog, 100)
	sub, err := client.SubscribeLogs(channel, types.LogFilter{
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
			to := client.MustNewAddress("0x160ebef20c1f739957bf9eecd040bce699cc42c6")
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
		case log := <-channel:
			fmt.Printf("received new log:%v\n\n", utils.PrettyJSON(log))
		}
	}
	sub.Unsubscribe()
}

func batchCall() {
	fmt.Println("\n- start batch call")
	elems := make([]rpc.BatchElem, 2)
	elems[0] = rpc.BatchElem{Method: "cfx_epochNumber", Result: &hexutil.Big{}, Args: []interface{}{}}
	elems[1] = rpc.BatchElem{Method: "cfx_getBalance", Result: &hexutil.Big{}, Args: []interface{}{client.MustNewAddress("cfxtest:aap9kthvctunvf030rbkk9k7zbzyz12dajp1u3sp4g")}}
	client.BatchCallRPC(elems)
}

func printResult(method string, args []interface{}, result interface{}, err error) {
	if err != nil {
		fmt.Printf("- function %v with args %+v error: %v\n\n", method, args, err.Error())
	} else {
		fmt.Printf("- function %v with args %+v result: %+v\n\n", method, args, context.JSONFmt(result))
	}
}
