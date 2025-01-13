package types

import (
	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalJsonHash(t *testing.T) {
	table := []struct {
		input       string
		expect      Hash
		expectError bool
	}{
		{
			input:  "\"0x0b50338ca31401a6ccb81a799b4c774ce86af93c1b740c52d12958e35461d999\"",
			expect: "0x0b50338ca31401a6ccb81a799b4c774ce86af93c1b740c52d12958e35461d999",
		},
		{
			input:       "\"0x0b50338ca31401a6ccb81a799b4c774ce86af93c1b740c52d12958e35461d99\"",
			expectError: true,
		},
		{
			input:       "\"\"",
			expectError: true,
		},
		{
			input:       "",
			expectError: true,
		},
	}

	for _, v := range table {
		var actual Hash

		bytes := []byte(v.input)

		// fmt.Printf("bytes: %v\n", string(bytes))
		err := utils.JsonUnmarshal(bytes, &actual)
		// fmt.Printf("err %v\n", err)

		if v.expectError && err == nil {
			t.Fatalf("expect error with input %v, but got %v ", v.input, actual)
		}

		if !v.expectError {
			if err != nil {
				t.Fatalf("unexpect error %v", err)
			}
			if v.expect != actual {
				t.Fatalf("expect %v,got %v", v.expect, actual)
			}
		}

	}
}

func TestHexOrDecimalUint64(t *testing.T) {
	t.Run("Json Marshal", func(t *testing.T) {
		u := HexOrDecimalUint64(10)
		b, err := utils.JsonMarshal(u)
		assert.NoError(t, err)
		assert.Equal(t, string(b), "\"0xa\"")
	})

	t.Run("Json Unmarshal", func(t *testing.T) {
		table := []string{
			"10",
			"\"0xa\"",
		}

		for _, b := range table {
			var u HexOrDecimalUint64
			err := utils.JsonUnmarshal([]byte(b), &u)
			assert.NoError(t, err)

			assert.Equal(t, HexOrDecimalUint64(10), u)
		}
	})
}
