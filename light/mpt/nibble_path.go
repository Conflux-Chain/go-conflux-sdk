package mpt

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type NibblePath struct {
	nibbles []byte
	start   int // inclusive
	end     int // exclusive
}

func NewNibblePath(key []byte) NibblePath {
	nibbles := make([]byte, 2*len(key))

	for i := 0; i < len(key); i++ {
		nibbles[2*i] = key[i] >> 4
		nibbles[2*i+1] = key[i] & 0x0F
	}

	return NibblePath{nibbles, 0, len(nibbles)}
}

func (path *NibblePath) Length() int {
	return path.end - path.start
}

func (path NibblePath) value() []byte {
	return path.nibbles[path.start:path.end]
}

func (path *NibblePath) ToChild() (childIndex byte, childPath NibblePath, ok bool) {
	if path.start == path.end {
		return
	}

	childIndex = path.nibbles[path.start]
	childPath = NibblePath{path.nibbles, path.start + 1, path.end}
	ok = true

	return
}

func (path *NibblePath) CommonPrefix(other *NibblePath) (prefix, path1, path2 NibblePath) {
	var offset int

	for {
		pathPos := path.start + offset
		if pathPos >= path.end {
			break
		}

		otherPos := other.start + offset
		if otherPos >= other.end {
			break
		}

		if path.nibbles[pathPos] != other.nibbles[otherPos] {
			break
		}

		offset++
	}

	prefix = NibblePath{
		nibbles: path.nibbles,
		start:   path.start,
		end:     path.start + offset,
	}

	path1 = NibblePath{
		nibbles: path.nibbles,
		start:   path.start + offset,
		end:     path.end,
	}

	path2 = NibblePath{
		nibbles: other.nibbles,
		start:   other.start + offset,
		end:     other.end,
	}

	return
}

func (path *NibblePath) ComputeMerkle(nodeMerkle common.Hash) common.Hash {
	if path.Length() == 0 {
		return nodeMerkle
	}

	var buffer []byte
	var start = path.start

	pathInfo := 128
	if start%2 == 1 {
		pathInfo += 64
	}
	pathInfo += path.Length() % 63
	buffer = append(buffer, byte(pathInfo))
	buffer = append(buffer, path.fullBytes()...)
	buffer = append(buffer, nodeMerkle.Bytes()...)

	return crypto.Keccak256Hash(buffer)
}

func (path *NibblePath) fullBytes() []byte {
	var result []byte

	start := path.start
	if start%2 == 1 {
		result = append(result, path.nibbles[start])
		start++
	}

	end := path.end
	if end%2 == 1 {
		end++
	}

	for i := start; i < end; i += 2 {
		key := path.nibbles[i]<<4 + path.nibbles[i+1]
		result = append(result, key)
	}

	return result
}

func (path *NibblePath) Trim() NibblePath {
	if path.start == path.end {
		return NibblePath{}
	}

	result := *path

	start := result.start
	if start%2 == 1 {
		start--
	}

	if start > 0 {
		result.nibbles = result.nibbles[start:]
		result.start -= start
		result.end -= start
	}

	end := result.end
	if end%2 == 1 {
		end++
	}

	if end < len(result.nibbles) {
		result.nibbles = result.nibbles[:end]
	}

	return result
}
