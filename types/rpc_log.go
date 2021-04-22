package types

import (
	"fmt"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
	"github.com/Conflux-Chain/go-conflux-sdk/utils"
)

func DefaultCallRPCLog(method string, args []interface{}, result interface{}, resultError error, duration time.Duration) {
	if resultError == nil {
		fmt.Printf("call rpc %v sucessfully, args %+v, result %+v, use %v\n", method, utils.PrettyJSON(args), utils.PrettyJSON(result), duration)
		return
	}
	fmt.Printf("call rpc %v failed, args %+v, error: %+v, use %v\n", method, args, resultError, duration)
}

func DefaultBatchCallRPCLog(b []rpc.BatchElem, err error, duration time.Duration) {
	if err == nil {
		fmt.Printf("batch call %+v sucessfully, use %v\n", b, duration)
		return
	}
	fmt.Printf("batch call %+v failed,  error: %+v, use %v\n", b, err, duration)
}
