package bcs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	val := map[byte]byte{
		'e': 'f',
		'a': 'b',
		'c': 'd',
	}

	assert.Equal(t, []byte{3, 'a', 'b', 'c', 'd', 'e', 'f'}, MustEncodeToBytes(val))
}
