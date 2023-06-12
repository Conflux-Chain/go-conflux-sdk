package contract

import (
	"fmt"
	"io"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/pkg/errors"
)

const (
	LogSpaceNative   uint8 = 1
	LogSpaceEthereum uint8 = 2
)

// EncodeRLP implements the rlp.Encoder interface.
func (log TypesTxLog) EncodeRLP(w io.Writer) error {
	switch log.Space {
	case LogSpaceNative:
		return rlp.Encode(w, []interface{}{log.Addr, log.Topics, log.Data})
	case LogSpaceEthereum:
		return rlp.Encode(w, []interface{}{log.Addr, log.Topics, log.Data, log.Space})
	default:
		return errors.Errorf("invalid log space %v", log.Space)
	}
}

func (proof TypesReceiptProof) ABIEncode() []byte {
	abi, err := LightNodeMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("Failed to get ABI for LightNode, err = %v", err.Error()))
	}

	encoded, err := abi.Pack("verifyReceiptProof", proof)
	if err != nil {
		panic(fmt.Sprintf("Failed to pack receipt proof, err = %v", err.Error()))
	}

	return encoded[4:]
}
