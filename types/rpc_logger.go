package types

import (
	"fmt"
	"reflect"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
)

type DefaultCallRPCLogger struct {
}

type DefaultBatchCallRPCLogger struct {
}

func (d DefaultCallRPCLogger) Info(method string, args []interface{}, result interface{}, duration time.Duration) {
	fmt.Printf("call rpc %v, args %+v, result %+v, use %v\n", method, args, reflect.ValueOf(result).Elem(), duration)
}

func (d DefaultCallRPCLogger) Error(method string, args []interface{}, resultError error, duration time.Duration) {
	fmt.Printf("call rpc %v, args %+v, error %+v, use %v\n", method, args, resultError, duration)
}

func (d DefaultBatchCallRPCLogger) Info(b []rpc.BatchElem, duration time.Duration) {
	fmt.Printf("batch call %+v, use %v\n", b, duration)
}

func (d DefaultBatchCallRPCLogger) Error(b []rpc.BatchElem, err error, duration time.Duration) {
	fmt.Printf("batch call %+v,  error %+v, use %v\n", b, err, duration)
}
