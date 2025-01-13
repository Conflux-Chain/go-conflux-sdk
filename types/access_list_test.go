package types

import (
	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonMarshalAccessList(t *testing.T) {
	table := []struct {
		input  AccessList
		expect string
	}{
		{
			input:  nil,
			expect: "null",
		},
		{
			input:  AccessList{},
			expect: "[]",
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		for _, item := range table {
			j, err := utils.JsonMarshal(item.input)
			assert.NoError(t, err)
			assert.Equal(t, item.expect, string(j))
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		for _, item := range table {
			var al *AccessList
			err := utils.JsonUnmarshal([]byte(item.expect), &al)
			assert.NoError(t, err)

			if item.input == nil {
				assert.True(t, al == nil)
			} else {
				assert.Equal(t, len(item.input), len(*al))
				assert.Equal(t, cap(item.input), cap(*al))
			}
		}
	})
}
