package middleware

import (
	"context"
	"encoding/json"
	"fmt"

	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/fatih/color"
	providers "github.com/openweb3/go-rpc-provider/provider_wrapper"
)

// CallRpcHandler represents interface of call rpc handler
type CallRpcHandler interface {
	Handle(ctx context.Context, result interface{}, method string, args ...interface{}) error
}

type CallRpcHandlerFunc providers.CallContextFunc

// CallRpcMiddleware represents the middleware for call rpc
type CallRpcMiddleware func(CallRpcHandler) CallRpcHandler

func (c CallRpcHandlerFunc) Handle(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	return c(ctx, result, method, args...)
}

// CallRpcConsoleMiddleware is the middleware for console request and response when call rpc
func CallRpcConsoleMiddleware(handler CallRpcHandler) CallRpcHandler {
	logFn := func(ctx context.Context, result interface{}, method string, args ...interface{}) error {

		argsStr := fmt.Sprintf("%+v", args)
		argsJson, err := json.Marshal(args)
		if err == nil {
			argsStr = string(argsJson)
		}

		start := time.Now()
		err = handler.Handle(ctx, result, method, args...)
		duration := time.Since(start)

		if err == nil {
			fmt.Printf("%v Method %v, Params %v, Result %v, Use %v\n",
				color.GreenString("[Call RPC Done]"),
				color.YellowString(method),
				color.CyanString(argsStr),
				color.CyanString(utils.PrettyJSON(result)),
				color.CyanString(duration.String()))
			return nil
		}

		color.Red("%v Method %v, Params %v, Error %v, Use %v\n",
			color.RedString("[Call RPC Fail]"),
			color.YellowString(method),
			color.CyanString(string(argsJson)),
			color.RedString(fmt.Sprintf("%+v", err)),
			color.CyanString(duration.String()))
		return err
	}
	return CallRpcHandlerFunc(logFn)
}
