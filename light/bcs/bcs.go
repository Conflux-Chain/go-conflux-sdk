package bcs

import (
	"errors"
	"io"
	"reflect"
)

func Encode(w io.Writer, val interface{}) (int, error) {
	if val == nil {
		return 0, errors.New("val is nil")
	}

	return write(w, reflect.ValueOf(val))
}

func EncodeToBytes(val interface{}) ([]byte, error) {
	if val == nil {
		return nil, errors.New("val is nil")
	}

	return encodeToBytes(reflect.ValueOf(val))
}

func MustEncodeToBytes(val interface{}) []byte {
	encoded, err := EncodeToBytes(val)
	if err != nil {
		panic(err)
	}
	return encoded
}
