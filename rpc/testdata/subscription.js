// This test checks basic subscription support.

--> {"jsonrpc":"2.0","id":1,"method":"nftest_subscribe","params":["someSubscription",5,1]}
<-- {"jsonrpc":"2.0","id":1,"result":"0x1"}
<-- {"jsonrpc":"2.0","method":"nftest_subscription","params":{"subscription":"0x1","result":1}}
<-- {"jsonrpc":"2.0","method":"nftest_subscription","params":{"subscription":"0x1","result":2}}
<-- {"jsonrpc":"2.0","method":"nftest_subscription","params":{"subscription":"0x1","result":3}}
<-- {"jsonrpc":"2.0","method":"nftest_subscription","params":{"subscription":"0x1","result":4}}
<-- {"jsonrpc":"2.0","method":"nftest_subscription","params":{"subscription":"0x1","result":5}}

--> {"jsonrpc":"2.0","id":1,"method":"nftest_subscribe","params":["echoVariadicArgs"]}
<-- {"jsonrpc":"2.0","id":1,"result":"0x2"}
<-- {"jsonrpc":"2.0","method":"nftest_subscription","params":{"subscription":"0x2","result":"hi world"}}

--> {"jsonrpc":"2.0","id":1,"method":"nftest_subscribe","params":["echoVariadicArgs", "david"]}
<-- {"jsonrpc":"2.0","id":1,"result":"0x3"}
<-- {"jsonrpc":"2.0","method":"nftest_subscription","params":{"subscription":"0x3","result":"hi world"}}
<-- {"jsonrpc":"2.0","method":"nftest_subscription","params":{"subscription":"0x3","result":"hi david"}}

--> {"jsonrpc":"2.0","id":1,"method":"nftest_subscribe","params":["echoVariadicArgs", "jimmy", "lily"]}
<-- {"jsonrpc":"2.0","id":1,"result":"0x4"}
<-- {"jsonrpc":"2.0","method":"nftest_subscription","params":{"subscription":"0x4","result":"hi world"}}
<-- {"jsonrpc":"2.0","method":"nftest_subscription","params":{"subscription":"0x4","result":"hi jimmy"}}
<-- {"jsonrpc":"2.0","method":"nftest_subscription","params":{"subscription":"0x4","result":"hi lily"}}

--> {"jsonrpc":"2.0","id":2,"method":"nftest_echo","params":[11]}
<-- {"jsonrpc":"2.0","id":2,"result":11}
