package utils

//import "encoding/json"
import "github.com/goccy/go-json"

func JSONUnmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func JSONMarshal(v any) ([]byte, error) {
	return json.Marshal(v)
}
