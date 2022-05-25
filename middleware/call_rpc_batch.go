package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/fatih/color"
	rpc "github.com/openweb3/go-rpc-provider"
	providers "github.com/openweb3/go-rpc-provider/provider_wrapper"
)

// BatchCallRpcHandler represents interface of batch call rpc handler
type BatchCallRpcHandler interface {
	Handle(ctx context.Context, b []rpc.BatchElem) error
}

type BatchCallRpcHandlerFunc providers.BatchCallContextFunc

// BatchCallRpcMiddleware represents the middleware for batch call rpc
type BatchCallRpcMiddleware func(BatchCallRpcHandler) BatchCallRpcHandler

func (brh BatchCallRpcHandlerFunc) Handle(ctx context.Context, b []rpc.BatchElem) error {
	return brh(ctx, b)
}

// BatchCallRpcConsoleMiddleware is the middleware for console request and response when batch call rpc
func BatchCallRpcConsoleMiddleware(ctx context.Context, handler BatchCallRpcHandler) BatchCallRpcHandler {
	logFn := func(ctx context.Context, b []rpc.BatchElem) error {
		start := time.Now()

		err := handler.Handle(ctx, b)

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
