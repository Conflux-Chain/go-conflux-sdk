## RPC CallContext/BatchCallContext Hook Middlewire

Client composite MiddlewarableMiddle for support setting middleware by `HookCallContext` for hooking `provider.CallContext` method which is the core of all single RPC-related methods. And `HookBatchCallContext` to set middleware for hooking `BatchCallContext`.

For example, we can custom a logger middleware to log for rpc requests.
```golang
client.HookCallContext(callContextConsoleMiddleware)
```
and the `callContextConsoleMiddleware` implementation is like
```golang
func callContextConsoleMiddleware(f providers.CallFunc) providers.CallFunc {
	return func(ctx context.Context, resultPtr interface{}, method string, args ...interface{}) error {
		fmt.Printf("request %v %v\n", method, args)
		err := f(ctx, resultPtr, method, args...)
		j, _ := json.Marshal(resultPtr)
		fmt.Printf("response %s\n", j)
		return err
	}
}
```

Also, you could 
- customize middleware
- use multiple middlewares

Notice that the middleware chain execution order is like onion, for example, use middleware A first and then middleware B
```go
client.HookCallContext(A)
client.HookCallContext(B)
```
the middleware execution order is
```
A --> B --> client.callRpc --> B --> A
```
