package mpt

import (
	"fmt"
	"math/big"

	"github.com/Conflux-Chain/go-conflux-sdk/light/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type ProofNode struct {
	path     NibblePath
	children [CHILDREN_COUNT]common.Hash // for branch node
	value    []byte                      // for leaf node
}

func newProofNode(n *Node) *ProofNode {
	return &ProofNode{
		path:     n.path.Trim(),
		children: n.computeChildrenHashes(),
		value:    n.value,
	}
}

func (n *ProofNode) ComputeMerkle() common.Hash {
	nodeMerkle := computeNodeMerkle(n.children, n.value)
	return n.path.ComputeMerkle(nodeMerkle)
}

func (n ProofNode) String() string {
	var children [16]string
	for i, v := range n.children {
		if v == KECCAKE_EMPTY {
			children[i] = "NULL"
		} else {
			children[i] = v.Hex()
		}
	}

	return fmt.Sprintf("{ path = %v, children = %v, value = %v }", n.path, children, n.value)
}

func (n *Node) Proof(key []byte) ([]*ProofNode, bool) {
	if len(key) == 0 {
		panic("key is empty")
	}

	path := NewNibblePath(key)

	switch n.nType {
	case nodeTypeBranch:
		// empty
		if n.children[0] == nil {
			return nil, false
		}

		// first branch has 2 children at least
		if n.children[1] != nil {
			return n.proof(path)
		}

		// trim the first branch node that has only one child
		childIndex, childPath, ok := path.ToChild()
		if !ok || childIndex != 0 {
			return nil, false
		}

		return n.children[0].proof(childPath)
	case nodeTypeLeaf:
		panic("root should be a branch node")
	default:
		panic("invalid node type")
	}
}

func (n *Node) proof(path NibblePath) ([]*ProofNode, bool) {
	_, path1, path2 := n.path.CommonPrefix(&path)

	// prefix mismatch
	if path1.Length() != 0 {
		return nil, false
	}

	proofNodes := []*ProofNode{newProofNode(n)}

	switch n.nType {
	case nodeTypeLeaf:
		if path2.Length() != 0 {
			return nil, false
		}
	case nodeTypeBranch:
		childIndex, childPath, ok := path2.ToChild()
		if !ok || n.children[childIndex] == nil {
			return nil, false
		}

		childrenProofs, ok := n.children[childIndex].proof(childPath)
		if !ok {
			return nil, false
		}

		proofNodes = append(proofNodes, childrenProofs...)
	default:
		panic("invalid node type")
	}

	return proofNodes, true
}

func Prove(root common.Hash, key, value []byte, proofNodes []*ProofNode) bool {
	if len(key) == 0 || len(proofNodes) == 0 {
		return false
	}

	path := NewNibblePath(key)
	expectedHash := root

	var nibblesLen int
	for _, v := range proofNodes {
		nibblesLen += v.path.Length()
		if v.children[0] != KECCAKE_EMPTY {
			nibblesLen++
		}
	}

	if nibblesLen%2 == 1 {
		if path.nibbles[0] != 0 {
			return false
		}

		path.start++
	}

	for _, v := range proofNodes {
		if nodeHash := v.ComputeMerkle(); nodeHash != expectedHash {
			return false
		}

		_, path1, path2 := v.path.CommonPrefix(&path)
		if path1.Length() != 0 {
			return false
		}

		switch v.children[0] {
		case KECCAKE_EMPTY: // leaf node
			if path2.Length() != 0 {
				return false
			}

			return crypto.Keccak256Hash(value) == crypto.Keccak256Hash(v.value)
		default:
			childIndex, childPath, ok := path2.ToChild()
			if !ok {
				return false
			}

			path = childPath
			expectedHash = v.children[childIndex]
		}
	}

	return false
}

func ConvertProofNode(nodes []*ProofNode) []contract.ProofLibProofNode {
	var result []contract.ProofLibProofNode

	for _, v := range nodes {
		var children [16][32]byte
		for i, child := range v.children {
			children[i] = child
		}

		var nibbles [32]byte
		copy(nibbles[:], v.path.nibbles)

		result = append(result, contract.ProofLibProofNode{
			Path: contract.ProofLibNibblePath{
				Nibbles: nibbles,
				Start:   big.NewInt(int64(v.path.start)),
				End:     big.NewInt(int64(v.path.end)),
			},
			Children: children,
			Value:    v.value,
		})
	}

	return result
}
