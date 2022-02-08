package cmptutil

import (
	"bytes"
	"encoding/json"
	"testing"

	"gotest.tools/assert"
)

func TestMarshalBytes(t *testing.T) {
	source := []struct {
		in  Bytes
		out string
	}{
		{
			in:  Bytes([]byte{0x1}),
			out: `"0x01"`,
		},
		{
			in:  nil,
			out: `"0x"`,
		},
	}

	for _, v := range source {
		j, e := json.Marshal(v.in)
		if e != nil {
			t.Fatal(e)
		}
		// fmt.Printf("%v %s\n", j, j)
		assert.Equal(t, v.out, string(j))
	}
}

func TestUnmarshalBytes(t *testing.T) {
	source := []struct {
		in  string
		out Bytes
	}{
		{
			in:  `"0x0102"`,
			out: Bytes([]byte{0x1, 0x2}),
		},
		{
			in:  `[1,2]`,
			out: Bytes([]byte{0x1, 0x2}),
		},
		{
			in:  `"0x"`,
			out: nil,
		},
		{
			in:  `[]`,
			out: nil,
		},
	}

	for _, v := range source {
		var b Bytes
		e := json.Unmarshal([]byte(v.in), &b)
		if e != nil {
			t.Fatal(e)
		}
		// fmt.Printf("%v %v\n", v.out, b)
		isequal := bytes.Compare(v.out.ToBytes(), b.ToBytes())
		assert.Equal(t, isequal, 0)
	}
}
