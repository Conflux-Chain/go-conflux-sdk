package contract

import (
	"fmt"
)

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
