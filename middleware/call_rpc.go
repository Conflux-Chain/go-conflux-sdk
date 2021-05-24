package middleware

import (
	"encoding/json"
	"fmt"

	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/fatih/color"
)

type CallRpcHandler interface {
	Handle(result interface{}, method string, args ...interface{}) error
}

type CallRpcHandlerFunc func(result interface{}, method string, args ...interface{}) error

type CallRpcMiddleWare func(CallRpcHandler) CallRpcHandler

func (c CallRpcHandlerFunc) Handle(result interface{}, method string, args ...interface{}) error {
	return c(result, method, args...)
}

func CallRpcLogMiddleWare(handler CallRpcHandler) CallRpcHandler {
	logFn := func(result interface{}, method string, args ...interface{}) error {
		start := time.Now()

		err := handler.Handle(result, method, args...)

		duration := time.Since(start)
		argsJson, _ := json.Marshal(args)

		if err == nil {
			fmt.Printf("%v Method %v, Params %v, Result %v, Use %v\n",
				color.GreenString("[Call RPC Done]"),
				color.YellowString(method),
				color.CyanString(string(argsJson)),
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
