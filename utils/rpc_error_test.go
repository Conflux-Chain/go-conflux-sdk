package utils

import "testing"

func TestRpcErrorInvalidParamsDetail(t *testing.T) {
	tests := []struct {
		name       string
		rpcErr     *RpcError
		wantDetail string
		wantOK     bool
	}{
		{name: "nil receiver", wantDetail: "", wantOK: false},
		{
			name:       "different code ignores matching data and message",
			rpcErr:     &RpcError{Code: -32601, Data: "data detail", Message: "Invalid params:message detail"},
			wantDetail: "",
			wantOK:     false,
		},
		{
			name:       "plain string data wins",
			rpcErr:     &RpcError{Code: -32602, Data: "data detail", Message: "Invalid params:message detail"},
			wantDetail: "data detail",
			wantOK:     true,
		},
		{
			name:       "quoted JSON string data is decoded once",
			rpcErr:     &RpcError{Code: -32602, Data: `"quoted detail"`},
			wantDetail: "quoted detail",
			wantOK:     true,
		},
		{
			name:       "quoted JSON escaped slash proves JSON decoding",
			rpcErr:     &RpcError{Code: -32602, Data: `"path \/rpc"`},
			wantDetail: "path /rpc",
			wantOK:     true,
		},
		{
			name:       "invalid params prefix is removed case insensitively",
			rpcErr:     &RpcError{Code: -32602, Message: "iNvAlId PaRaMs: keep leading space"},
			wantDetail: " keep leading space",
			wantOK:     true,
		},
		{
			name:       "invalid parameters prefix is removed case insensitively",
			rpcErr:     &RpcError{Code: -32602, Message: "INVALID PARAMETERS:detail"},
			wantDetail: "detail",
			wantOK:     true,
		},
		{
			name:       "generic invalid params without colon has no detail",
			rpcErr:     &RpcError{Code: -32602, Message: "Invalid params"},
			wantDetail: "",
			wantOK:     true,
		},
		{
			name:       "generic invalid params with colon has no detail",
			rpcErr:     &RpcError{Code: -32602, Message: "Invalid params:"},
			wantDetail: "",
			wantOK:     true,
		},
		{
			name:       "generic invalid parameters without colon has no detail",
			rpcErr:     &RpcError{Code: -32602, Message: "INVALID PARAMETERS"},
			wantDetail: "",
			wantOK:     true,
		},
		{
			name:       "generic invalid parameters with colon has no detail",
			rpcErr:     &RpcError{Code: -32602, Message: "INVALID PARAMETERS:"},
			wantDetail: "",
			wantOK:     true,
		},
		{
			name:       "message without generic prefix is preserved",
			rpcErr:     &RpcError{Code: -32602, Message: "execution rejected the argument"},
			wantDetail: "execution rejected the argument",
			wantOK:     true,
		},
		{
			name:       "empty string data falls back to message",
			rpcErr:     &RpcError{Code: -32602, Data: "", Message: "Invalid params:fallback detail"},
			wantDetail: "fallback detail",
			wantOK:     true,
		},
		{
			name:       "non-string data falls back to message",
			rpcErr:     &RpcError{Code: -32602, Data: map[string]string{"detail": "ignored"}, Message: "Invalid parameters:fallback detail"},
			wantDetail: "fallback detail",
			wantOK:     true,
		},
		{
			name:       "invalid quoted data is preserved",
			rpcErr:     &RpcError{Code: -32602, Data: `"invalid \q escape"`},
			wantDetail: `"invalid \q escape"`,
			wantOK:     true,
		},
		{
			name:       "only leading quote is preserved",
			rpcErr:     &RpcError{Code: -32602, Data: `"leading quote only`},
			wantDetail: `"leading quote only`,
			wantOK:     true,
		},
		{
			name:       "only trailing quote is preserved",
			rpcErr:     &RpcError{Code: -32602, Data: `trailing quote only"`},
			wantDetail: `trailing quote only"`,
			wantOK:     true,
		},
		{
			name:       "decoded empty data falls back to message",
			rpcErr:     &RpcError{Code: -32602, Data: `""`, Message: "Invalid params:decoded empty fallback"},
			wantDetail: "decoded empty fallback",
			wantOK:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDetail, gotOK := tt.rpcErr.InvalidParamsDetail()
			if gotDetail != tt.wantDetail || gotOK != tt.wantOK {
				t.Fatalf("InvalidParamsDetail() = (%q, %v), want (%q, %v)", gotDetail, gotOK, tt.wantDetail, tt.wantOK)
			}
		})
	}
}
