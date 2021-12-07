package abiutil

import (
	"encoding/hex"
	"testing"

	"gotest.tools/assert"
)

func TestDecodeErrData(t *testing.T) {
	data, err := hex.DecodeString("08c379a0000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000294552433732313a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e0000000000000000000000000000000000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	res, err := DecodeErrData(data)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, res, "ERC721: owner query for nonexistent token")
}
