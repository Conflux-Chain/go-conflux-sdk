package mpt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNibblePathNew(t *testing.T) {
	assert.Equal(t, []byte{0, 1, 0, 2, 0, 3}, NewNibblePath([]byte{1, 2, 3}).value())
	assert.Equal(t, []byte{1, 1, 0, 2, 0, 3}, NewNibblePath([]byte{17, 2, 3}).value())
}

func TestNibblePathLength(t *testing.T) {
	path := NewNibblePath([]byte{1, 2, 3})
	assert.Equal(t, 6, path.Length())

	path.start++
	path.end--
	assert.Equal(t, 4, path.Length())
}

func TestNibblePathToChild(t *testing.T) {
	path := NewNibblePath([]byte{1})

	index, path, ok := path.ToChild()
	assert.True(t, ok)
	assert.Equal(t, byte(0), index)
	assert.Equal(t, []byte{1}, path.value())

	index, path, ok = path.ToChild()
	assert.True(t, ok)
	assert.Equal(t, byte(1), index)
	assert.Equal(t, []byte{}, path.value())

	_, _, ok = path.ToChild()
	assert.False(t, ok)
}

func TestNibblePathCommonPrefix(t *testing.T) {
	// path1 is empty
	path1 := NewNibblePath([]byte{})
	path2 := NewNibblePath([]byte{1, 2, 3})
	prefix, left1, left2 := path1.CommonPrefix(&path2)
	assert.Equal(t, []byte{}, prefix.value())
	assert.Equal(t, []byte{}, left1.value())
	assert.Equal(t, []byte{0, 1, 0, 2, 0, 3}, left2.value())

	// path2 is empty
	path1 = NewNibblePath([]byte{1, 2, 3})
	path2 = NewNibblePath([]byte{})
	prefix, left1, left2 = path1.CommonPrefix(&path2)
	assert.Equal(t, []byte{}, prefix.value())
	assert.Equal(t, []byte{0, 1, 0, 2, 0, 3}, left1.value())
	assert.Equal(t, []byte{}, left2.value())

	// no prefix
	path1 = NewNibblePath([]byte{1, 2, 3})
	path2 = NewNibblePath([]byte{16, 2, 3})
	prefix, left1, left2 = path1.CommonPrefix(&path2)
	assert.Equal(t, []byte{}, prefix.value())
	assert.Equal(t, []byte{0, 1, 0, 2, 0, 3}, left1.value())
	assert.Equal(t, []byte{1, 0, 0, 2, 0, 3}, left2.value())

	// has prefix
	path1 = NewNibblePath([]byte{1, 2, 1})
	path1.start = 2
	path2 = NewNibblePath([]byte{16, 2, 16})
	path2.start = 2
	prefix, left1, left2 = path1.CommonPrefix(&path2)
	assert.Equal(t, []byte{0, 2}, prefix.value())
	assert.Equal(t, []byte{0, 1}, left1.value())
	assert.Equal(t, []byte{1, 0}, left2.value())
}

func TestNibblePathTrim(t *testing.T) {
	path := NewNibblePath([]byte{1, 2, 3})
	path.start = path.end
	assert.Equal(t, []byte(nil), path.Trim().value())

	path = NewNibblePath([]byte{1, 2, 3})

	path.start++
	path.end--
	trimed := path.Trim()
	assert.Equal(t, []byte{0, 1, 0, 2, 0, 3}, trimed.nibbles)
	assert.Equal(t, 1, trimed.start)
	assert.Equal(t, 5, trimed.end)

	path.start++
	path.end--
	trimed = path.Trim()
	assert.Equal(t, []byte{0, 2}, trimed.nibbles)
	assert.Equal(t, 0, trimed.start)
	assert.Equal(t, 2, trimed.end)
}
