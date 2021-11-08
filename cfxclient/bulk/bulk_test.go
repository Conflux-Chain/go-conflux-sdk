package bulk

import (
	"fmt"
	"math/big"
	"testing"

	client "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// use
func TestBulkCall(t *testing.T) {

	raw, success := big.NewInt(0).SetString("0x12837843846827364abcdef12334fbd45ac123acd", 0)
	val := (*hexutil.Big)(raw)
	fmt.Printf("raw %v, val %v, sucess: %v\n", raw, val, success)
	// return

	_client, err := client.NewClient("https://test.confluxrpc.com")
	if err != nil {
		panic(err)
	}
	bulkCaller := NewBulkerCaller(_client)

	gasPrice := bulkCaller.GetGasPrice()
	_errors, err := bulkCaller.Execute()
	if err != nil {
		panic(err)
	}
	gasPriceError := _errors[0]
	if gasPriceError != nil {
		fmt.Printf("get price error %v", gasPriceError)
		panic(gasPriceError)
	}

	if gasPrice == nil {
		panic("failed get gasPrice")
	}
	// return

	addresses := [2]cfxaddress.Address{
		cfxaddress.MustNew("cfxtest:aamjxdgz4m84hjvf2s9rmw5uzd4dkh8aa6krdsh0ep"),
		cfxaddress.MustNew("cfxtest:aak2rra2njvd77ezwjvx04kkds9fzagfe6d5r8e957"),
	}

	var nonces [len(addresses)]*hexutil.Big
	for i := 0; i < len(nonces); i++ {
		nonces[i] = bulkCaller.GetNextNonce(addresses[i])
	}

	errors, err := bulkCaller.Execute()
	if err != nil {
		panic(err)
	}

	nonceErrors := errors[1 : 1+len(addresses)]
	for i := 0; i < len(nonceErrors); i++ {
		if nonceErrors[i] != nil {
			panic(nonceErrors[i])
		}
		fmt.Printf("get nonce of address %v %v\n", addresses[i], &nonces[i])
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
