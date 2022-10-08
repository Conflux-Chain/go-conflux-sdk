package integrationtest

import (
	"fmt"
	"sync"
	"testing"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
)

func _TestResubHeads(t *testing.T) {
	client := sdk.MustNewClient("wss://test.confluxrpc.com/ws")

	headc := make(chan types.BlockHeader)
	sub := client.SubscribeNewHeadsWitReconn(headc)

	retry := 0

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case err := <-sub.Err():
				fmt.Printf("sub error: %v\n", err)
				retry++
				if retry >= 20 {
					sub.Unsubscribe()
					return
				}
			case <-sub.ResubSuccess():
				fmt.Println("sub success")
				retry = 0
			case h := <-headc:
				fmt.Printf("received head %v\n", h)
			}
		}
	}()
	wg.Wait()
}
