package mpt

import (
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const CHILDREN_COUNT = 16

var (
	KECCAKE_EMPTY      = crypto.Keccak256Hash()
	LEAF_NODE_CHILDREN = [CHILDREN_COUNT]common.Hash{
		KECCAKE_EMPTY, KECCAKE_EMPTY, KECCAKE_EMPTY, KECCAKE_EMPTY,
		KECCAKE_EMPTY, KECCAKE_EMPTY, KECCAKE_EMPTY, KECCAKE_EMPTY,
		KECCAKE_EMPTY, KECCAKE_EMPTY, KECCAKE_EMPTY, KECCAKE_EMPTY,
		KECCAKE_EMPTY, KECCAKE_EMPTY, KECCAKE_EMPTY, KECCAKE_EMPTY,
	}
)

type nodeType int

const (
	nodeTypeBranch = iota
	nodeTypeLeaf
)

type Node struct {
	nType    nodeType
	path     NibblePath
	value    []byte                // for leaf node only
	children [CHILDREN_COUNT]*Node // unlike Ethereum, no extension node in Conflux

	merkle atomic.Value
}

func mustNewLeafNode(path NibblePath, value []byte) (childIndex byte, leaf *Node) {
	index, childPath, ok := path.ToChild()
	if !ok {
		panic("path too short to create leaf node")
	}

	return index, &Node{
		nType: nodeTypeLeaf,
		path:  childPath,
		value: value,
	}
}

func newBranchNode(path NibblePath, children [CHILDREN_COUNT]*Node) *Node {
	return &Node{
		nType:    nodeTypeBranch,
		path:     path,
		children: children,
	}
}

// Insert inserts key and value into mpt.
// Note, key should be in the same length and inserted in order.
func (n *Node) Insert(key, value []byte) {
	if len(key) == 0 {
		panic("key is empty")
	}

	path := NewNibblePath(key)

	n.insert(path, value)
}

func (n *Node) insert(path NibblePath, value []byte) {
	switch n.nType {
	case nodeTypeBranch:
		prefix, path1, path2 := n.path.CommonPrefix(&path)
		if path2.Length() == 0 {
			panic("key length inconsistent")
		}

		if path1.Length() == 0 {
			childIndex, leaf := mustNewLeafNode(path2, value)
			if n.children[childIndex] == nil {
				n.children[childIndex] = leaf
			} else {
				n.children[childIndex].insert(leaf.path, value)
			}
		} else {
			childIndex1, path1, _ := path1.ToChild()
			if path1.Length() > 0 {
				panic("key not inserted in sequence")
			}
			child1 := newBranchNode(path1, n.children)

			childIndex2, child2 := mustNewLeafNode(path2, value)

			// update current branch node
			n.path = prefix
			n.children = [CHILDREN_COUNT]*Node{}
			n.children[childIndex1] = child1
			n.children[childIndex2] = child2
		}
	case nodeTypeLeaf:
		if n.path.Length() != path.Length() {
			panic("key length inconsistent")
		}

		// find common prefix for branch node
		prefix, path1, path2 := n.path.CommonPrefix(&path)
		if path1.Length() == 0 {
			panic("duplicate key inserted")
		}

		// construct 2 new leaf nodes
		childIndex1, leaf1 := mustNewLeafNode(path1, n.value)
		childIndex2, leaf2 := mustNewLeafNode(path2, value)

		// change current leaf node to branch node
		n.nType = nodeTypeBranch
		n.path = prefix
		n.value = nil
		n.children[childIndex1] = leaf1
		n.children[childIndex2] = leaf2
	default:
		panic("invalid node type")
	}
}

func (n *Node) Hash() common.Hash {
	switch n.nType {
	case nodeTypeBranch:
		if n.children[0] == nil {
			return KECCAKE_EMPTY
		}

		if n.children[1] == nil {
			return n.children[0].computeMerkle()
		}

		return n.computeMerkle()
	case nodeTypeLeaf:
		panic("root should be a branch node")
	default:
		panic("invalid node type")
	}
}

func (n *Node) computeMerkle() common.Hash {
	if merkle := n.merkle.Load(); merkle != nil {
		return merkle.(common.Hash)
	}

	childrenHashes := n.computeChildrenHashes()
	nodeMerkle := computeNodeMerkle(childrenHashes, n.value)
	merkle := n.path.ComputeMerkle(nodeMerkle)

	n.merkle.Store(merkle)

	return merkle
}

func (n *Node) computeChildrenHashes() (childrenHashes [CHILDREN_COUNT]common.Hash) {
	if n.nType == nodeTypeLeaf {
		return LEAF_NODE_CHILDREN
	}

	for i, v := range n.children {
		if v == nil {
			childrenHashes[i] = KECCAKE_EMPTY
		} else {
			childrenHashes[i] = v.computeMerkle()
		}
	}

	return
}

func computeNodeMerkle(children [CHILDREN_COUNT]common.Hash, value []byte) common.Hash {
	var buffer []byte

	buffer = append(buffer, 'n')

	for _, v := range children {
		buffer = append(buffer, v.Bytes()...)
	}

	if len(value) > 0 {
		buffer = append(buffer, 'v')
		buffer = append(buffer, value...)
	}

	return crypto.Keccak256Hash(buffer)
}
