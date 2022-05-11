package bulk

import (
	"fmt"
	"testing"

	client "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common/hexutil"
	rpc "github.com/openweb3/go-rpc-provider"
)

// use
func TestBulkCall(t *testing.T) {
	_client, err := client.NewClient("https://test.confluxrpc.com")
	// _client.UseBatchCallRpcMiddleware(middleware.BatchCallRpcConsoleMiddleware)
	if err != nil {
		panic(err)
	}
	bulkCaller := NewBulkCaller(_client)

	gasPrice, gasPriceError := bulkCaller.Cfx().GetGasPrice()
	err = bulkCaller.Execute()
	if err != nil {
		panic(err)
	}

	if *gasPriceError != nil {
		fmt.Printf("get price error %v", *gasPriceError)
		panic(*gasPriceError)
	}

	if gasPrice == nil {
		panic("failed get gasPrice")
	}

	fmt.Printf("get get price %v\n", gasPrice)

	addresses := [2]cfxaddress.Address{
		cfxaddress.MustNew("cfxtest:aamjxdgz4m84hjvf2s9rmw5uzd4dkh8aa6krdsh0ep"),
		cfxaddress.MustNew("cfxtest:aak2rra2njvd77ezwjvx04kkds9fzagfe6d5r8e957"),
	}

	var nonces [len(addresses)]*hexutil.Big
	var nonceErrors [len(addresses)]*error
	for i := 0; i < len(nonces); i++ {
		nonces[i], nonceErrors[i] = bulkCaller.Cfx().GetNextNonce(addresses[i], types.NewEpochNumberUint64(1000000000))
	}

	err = bulkCaller.Execute()
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(nonceErrors); i++ {
		if *nonceErrors[i] == nil {
			t.Fatalf("expect get nonce error")
		}
		fmt.Printf("get nonce error: %v\n", *nonceErrors[i])
	}

	bulkCaller.Clear()
	for i := 0; i < len(nonces); i++ {
		nonces[i], nonceErrors[i] = bulkCaller.Cfx().GetNextNonce(addresses[i])
	}
	err = bulkCaller.Execute()
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(nonceErrors); i++ {
		if *nonceErrors[i] != nil {
			panic(*nonceErrors[i])
		}
		fmt.Printf("get nonce: %v\n", nonces[i])
	}

}

func TestBatchOne(t *testing.T) {
	client, err := client.NewClient("http://test.confluxrpc.com")
	if err != nil {
		panic(err)
	}

	var gasPrice *hexutil.Big
	var gasPricePtr = &gasPrice

	var batchElem rpc.BatchElem = rpc.BatchElem{"cfx_gasPrice", nil, &gasPricePtr, nil}

	if err = client.BatchCallRPC([]rpc.BatchElem{batchElem}); err != nil {
		t.Fatal(err)
	}

	fmt.Printf("gasPrice %v\n", gasPrice)
}
