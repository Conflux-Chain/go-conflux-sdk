package sdk

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
	providers "github.com/openweb3/go-rpc-provider/provider_wrapper"
	"github.com/stretchr/testify/assert"
)

func TestInterfaceImplementation(t *testing.T) {
	var _ ClientOperator = &Client{}
}

func TestNewClientNotCrash(t *testing.T) {
	NewClient("https://test.confluxrpc.com", ClientOption{
		KeystorePath: "./keystore",
	})
}

func TestClose(t *testing.T) {
	c := MustNewClient("https://test.confluxrpc.com")
	c.Close()
}

func TestClientHookCallContext(t *testing.T) {
	c := MustNewClient("https://test.confluxrpc.com")
	mp := c.Provider()
	mp.HookCallContext(callContextMid1)
	mp.HookCallContext(callContextMid2)
	c.GetStatus()
}

func TestEpochWhenGetNextnonce(t *testing.T) {
	c := MustNewClient("https://test.confluxrpc.com", ClientOption{Logger: os.Stdout})
	c.GetNextNonce(cfxaddress.MustNew("cfxtest:aaskvgxcfej371g4ecepx9an78ngrke5ay9f8jtbgg"), types.NewEpochOrBlockHashWithEpoch(types.EpochLatestMined))
}

func callContextMid1(f providers.CallContextFunc) providers.CallContextFunc {
	return func(ctx context.Context, result interface{}, method string, args ...interface{}) error {
		ctx = context.WithValue(ctx, "foo", "bar")
		return f(ctx, result, method, args...)
	}
}

func callContextMid2(f providers.CallContextFunc) providers.CallContextFunc {
	return func(ctx context.Context, result interface{}, method string, args ...interface{}) error {
		fmt.Printf("ctx value of foo: %+v\n", ctx.Value("foo"))
		return f(ctx, result, method, args...)
	}
}

func TestSendTransaction(t *testing.T) {
	client := MustNewClient("http://net8888cfx.confluxrpc.com", ClientOption{
		KeystorePath: "./keystore",
		Logger:       os.Stdout,
	})
	client.AccountManager.ImportKey("0x0ccb34a57c54b3e61effebbc3cf3baaf5a03cd07a90836f38d97e08bcd1662f3", "hello")
	// assert.NoError(t, err)
	client.AccountManager.UnlockDefault("hello")

	addr, _ := client.AccountManager.GetDefault()

	fmt.Printf("addr %v\n", addr)

	txHash, err := client.SendTransaction(types.UnsignedTransaction{
		UnsignedTransactionBase: types.UnsignedTransactionBase{
			AccessList: types.AccessList{types.AccessTuple{
				Address: *addr,
				StorageKeys: []common.Hash{
					common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
				},
			}},
			// Gas: types.NewBigInt(30000),
		},
		To: addr,
	})
	assert.NoError(t, err)

	fmt.Println(txHash)

}
