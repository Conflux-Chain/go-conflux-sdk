Middleware
-----------------------------

Client applies the method `UseCallRpcMiddleware` to set middleware for hooking `callRpc` method which is the core of all single RPC-related methods. And `UseBatchCallRpcMiddleware` to set middleware for hooking `batchCallRPC`.

For example, use `CallRpcConsoleMiddleware` to log for rpc requests.
```golang
client.UseCallRpcMiddleware(middleware.CallRpcConsoleMiddleware)
```

Also, you could 
- customize middleware
- use multiple middlewares

Notice that the middleware chain execution order is like onion, for example, use middleware A first and then middleware B
```go
client.UseCallRpcMiddleware(A)
client.UseCallRpcMiddleware(B)
```
the middleware execution order is
```
B --> A --> client.callRpc --> A --> B
```