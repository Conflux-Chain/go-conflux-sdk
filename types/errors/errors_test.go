package errors

import (
	"testing"
)

func TestIsBusinessError(t *testing.T) {
	err := BusinessError{}
	actual := IsBusinessError(err)
	if !actual {
		t.Fatalf("expect %v actural %v", true, actual)
	}
}
