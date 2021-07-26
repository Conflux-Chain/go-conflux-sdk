// This test checks that calls with variadic param(s) work.

--> {"jsonrpc": "2.0", "id": 2, "method": "test_variadicArgs", "params": ["x", 3]}
<-- {"jsonrpc":"2.0","id":2,"result":[{"String":"x","Int":3,"Args":null}]}

--> {"jsonrpc": "2.0", "id": 2, "method": "test_variadicArgs", "params": ["x", 3, {"S": "foo"}]}
<-- {"jsonrpc":"2.0","id":2,"result":[{"String":"x","Int":3,"Args":{"S":"foo"}}]}

--> {"jsonrpc": "2.0", "id": 2, "method": "test_variadicArgs", "params": ["x", 3, {"S": "foo"}, {"S": "bar"}]}
<-- {"jsonrpc":"2.0","id":2,"result":[{"String":"x","Int":3,"Args":{"S":"foo"}},{"String":"x","Int":3,"Args":{"S":"bar"}}]}