package middleware

import (
	"fmt"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/rpc"
	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/fatih/color"
)

type BatchCallRpcHandler interface {
	Handle(b []rpc.BatchElem) error
}

type BatchCallRpcHandlerFunc func(b []rpc.BatchElem) error

type BatchCallRpcMiddleWare func(BatchCallRpcHandler) BatchCallRpcHandler

func (brh BatchCallRpcHandlerFunc) Handle(b []rpc.BatchElem) error {
	return brh(b)
}

func BatchCallRpcLogMiddleWare(handler BatchCallRpcHandler) BatchCallRpcHandler {
	logFn := func(b []rpc.BatchElem) error {
		start := time.Now()

		err := handler.Handle(b)

		duration := time.Since(start)
		if err == nil {
			fmt.Printf("%v BatchElems %v, Use %v\n",
				color.GreenString("[Batch Call RPC Done]"),
				color.CyanString(utils.PrettyJSON(b)),
				color.CyanString(duration.String()))
			return nil
		}
		fmt.Printf("%v BatchElems %v, Error: %v, Use %v\n",
			color.RedString("[Batch Call RPC Fail]"),
			color.CyanString(utils.PrettyJSON(b)),
			color.RedString(fmt.Sprintf("%+v", err)),
			duration)
		return err
	}
	return BatchCallRpcHandlerFunc(logFn)
}
