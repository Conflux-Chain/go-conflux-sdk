package light

import (
	"github.com/Conflux-Chain/go-conflux-sdk/light/mpt"
	"github.com/Conflux-Chain/go-conflux-sdk/light/primitives"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func CreateTransactionsMPT(txs []types.WrapTransaction) *mpt.Node {
	var root mpt.Node

	keyLen := mpt.MinReprBytes(len(txs))

	for i, v := range txs {
		key := mpt.ToIndexBytes(i, keyLen)

		if v.NativeTransaction != nil {
			root.Insert(key, hexutil.MustDecode(v.NativeTransaction.Hash.String()))
		} else {
			root.Insert(key, v.EthTransaction.Hash().Bytes())
		}
	}

	return &root
}

func CreateReceiptsMPT(epochReceipts [][]types.TransactionReceipt) ([]*mpt.Node, *mpt.Node) {
	var subtrees []*mpt.Node

	for _, blockReceipts := range epochReceipts {
		var root mpt.Node

		keyLen := mpt.MinReprBytes(len(blockReceipts))

		for i, v := range blockReceipts {
			key := mpt.IndexToKey(i, keyLen)
			value := primitives.MustRLPEncodeReceipt(&v)
			root.Insert(key, value)
		}

		subtrees = append(subtrees, &root)
	}

	var root mpt.Node
	keyLen := mpt.MinReprBytes(len(subtrees))
	for i, v := range subtrees {
		key := mpt.IndexToKey(i, keyLen)
		value := v.Hash().Bytes()
		root.Insert(key, value)
	}

	return subtrees, &root
}
