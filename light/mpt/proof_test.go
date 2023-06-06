package mpt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProof(t *testing.T) {
	valueFunc := func(i int) []byte {
		return []byte(fmt.Sprintf("leaf node value - %v", i))
	}

	for leafNodes := 1; leafNodes <= 300; leafNodes++ {
		keyLen := MinReprBytes(leafNodes)

		var root Node
		for i := 0; i < leafNodes; i++ {
			key := ToIndexBytes(i, keyLen)
			root.Insert(key, valueFunc(i))
		}

		txsRoot := root.Hash()

		for i := 0; i < leafNodes; i++ {
			key := ToIndexBytes(i, keyLen)
			proofNodes, ok := root.Proof(key)
			assert.True(t, ok, "Failed to generate proof, leaf nodes = %v, index = %v", leafNodes, i)

			proved := Prove(txsRoot, key, valueFunc(i), proofNodes)
			assert.True(t, proved, "Failed to prove, leaf nodes = %v, index = %v", leafNodes, i)
		}
	}
}
