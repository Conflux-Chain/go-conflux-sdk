package sdk

// GoConvey COMPOSER
// Test NewClient
// 	Subject: New client use rpc client
// 		Given a node url and retry params
// 			When rpc dail error
// 				Return error
// 			When rpc dail success
// 				Return client instance
import (
	"context"
	"fmt"
	"testing"

	. "bou.ke/monkey"
	// "github.com/ethereum/go-ethereum/rpc"

	rpc "github.com/openweb3/go-rpc-provider"
	providers "github.com/openweb3/go-rpc-provider/provider_wrapper"
	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
)

func _TestNewClient(t *testing.T) {

	Convey("Subject: New client use rpc client", t, func() {

		Convey("Given a node url and retry params", func() {

			Convey("When rpc dail error", func() {
				//stub for rpc.Dail
				guard := Patch(rpc.Dial, func(_ string) (*rpc.Client, error) {
					return nil, errors.New("rpc dail fail")
				})
				defer guard.Unpatch()

				client, err := newClientWithOption("", ClientOption{})
				Convey("Return error", func() {
					So(err, ShouldNotEqual, nil)
					So(client, ShouldEqual, nil)
				})
			})

			Convey("When rpc dail success", func() {
				//stub for rpc.Dail
				guard := Patch(rpc.Dial, func(_ string) (*rpc.Client, error) {
					return &rpc.Client{}, nil
				})
				defer guard.Unpatch()

				client, err := newClientWithOption("", ClientOption{})
				// fmt.Printf("client:%+v,err:%+v", client, err)

				Convey("Return client instance", func() {
					So(err, ShouldEqual, nil)
					So(client, ShouldNotEqual, nil)
				})
			})
		})
	})

}

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
