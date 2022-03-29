package integrationtest

type MockRPC struct {
	Version  string `json:"version"`
	Examples map[string][]struct {
		Name        string        `json:"name"`
		Description string        `json:"description"`
		Params      []interface{} `json:"params"`
		Result      interface{}   `json:"result"`
		Error       interface{}   `json:"error"`
	} `json:"examples"`
}

type rpcTestConfig struct {
	examplesUrl string
	client      interface{}

	rpc2Func         map[string]string
	rpc2FuncSelector map[string]func(params []interface{}) (string, []interface{})
	ignoreRpc        map[string]bool
	onlyTestRpc      map[string]bool
}
