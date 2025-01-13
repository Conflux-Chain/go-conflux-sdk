package utils

import (
	sysJson "encoding/json"
	libJson "github.com/goccy/go-json"
	"os"
)

var unmarshal = sysJson.Unmarshal
var marshal = sysJson.Marshal

func JsonUnmarshal(data []byte, v any) error {
	return unmarshal(data, v)
}

func JsonMarshal(v any) ([]byte, error) {
	return marshal(v)
}

func init() {
	if os.Getenv("UseGoCcyJson") != "" {
		unmarshal = libJson.Unmarshal
		marshal = libJson.Marshal
	}
}
