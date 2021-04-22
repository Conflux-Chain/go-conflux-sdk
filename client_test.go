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
	"errors"
	"testing"

	. "bou.ke/monkey"
	// "github.com/ethereum/go-ethereum/rpc"

	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
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

				client, err := newClientWithRetry("", ClientOption{})
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

				client, err := newClientWithRetry("", ClientOption{})
				// fmt.Printf("client:%+v,err:%+v", client, err)

				Convey("Return client instance", func() {
					So(err, ShouldEqual, nil)
					So(client, ShouldNotEqual, nil)
				})

			})

		})

	})

}
