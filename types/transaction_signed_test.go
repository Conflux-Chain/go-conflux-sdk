package types

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/status-im/keycard-go/hexutils"
	"github.com/stretchr/testify/assert"
)

func TestDecodeTransaction(t *testing.T) {
	rawData := `0xf852ce010183021000800183038e9b800101a0ca43b3f84e4afefcc6946d2953a0391774bd3c692015f64e51895b4a93fcba31a036953d461a84e15367a463d4f7624970ad4c3833013f23103f2ea90a202e4aea`
	tx := SignedTransaction{}
	b := mustDecodeStringToBytes(t, rawData)
	err := tx.Decode(b, 1037)
	if err != nil {
		t.Fatal(err)
	}

	b, err = json.Marshal(tx)
	if err != nil {
		t.Fatal(err)
	}

	expect := `{"UnsignedTransaction":{"From":null,"Nonce":"0x1","GasPrice":"0x1","Gas":"0x21000","Value":"0x1","EpochHeight":"0x38e9b","ChainID":null,"To":null,"Data":"0x01"},"V":1,"R":"0xca43b3f84e4afefcc6946d2953a0391774bd3c692015f64e51895b4a93fcba31","S":"0x36953d461a84e15367a463d4f7624970ad4c3833013f23103f2ea90a202e4aea"}0xc0ebdb1465edbffd0491b5ad9080745b1adbbe0b4d19d74959734e8333b5cce4`

	if string(b) != expect {
		t.Fatalf("expect: %s, got: %s", expect, string(b))
	}
}

func TestEncodeTransaction(t *testing.T) {
	jsonStr := `{"UnsignedTransaction":{"From":"CFXTEST:TYPE.USER:AAR7X4R8MKRNW39GGS8RZ40J1ZNWH5MRRPUFPR2U76","Nonce":"0x1","GasPrice":"0x1","Gas":"0x21000","Value":"0x1","StorageLimit":"0x1","EpochHeight":"0x38e9b","ChainID":"0x0","To":null,"Data":"0x01"},"V":1,"R":"0xca43b3f84e4afefcc6946d2953a0391774bd3c692015f64e51895b4a93fcba31","S":"0x36953d461a84e15367a463d4f7624970ad4c3833013f23103f2ea90a202e4aea"}`
	tx := SignedTransaction{}
	err := json.Unmarshal([]byte(jsonStr), &tx)
	if err != nil {
		t.Fatal(err)
	}

	b, err := tx.Encode()
	if err != nil {
		t.Fatal(err)
	}

	expect := `f852ce010183021000800183038e9b800101a0ca43b3f84e4afefcc6946d2953a0391774bd3c692015f64e51895b4a93fcba31a036953d461a84e15367a463d4f7624970ad4c3833013f23103f2ea90a202e4aea`
	if hex.EncodeToString(b) != expect {
		t.Fatalf("expect %s, but got %s", expect, hex.EncodeToString(b))
	}
}

func mustDecodeStringToBytes(t *testing.T, s string) []byte {
	b, e := hex.DecodeString(s[2:])
	if e != nil {
		t.Fatal(e)
	}
	return b
}

/*
  "result": {
    "hash": "0xbe47fb886dd24dba30cb25527b866b3139fb020ac344d5dc910b0140cce05478",
    "nonce": "0x4d073",
    "blockHash": "0xd530f560436892a034ccd3516de50a8745d7607b834d483d75097ad13e1bfb0a",
    "transactionIndex": "0x0",
    "from": "cfxtest:aaph8hphbv84fkn3bunm051aek68aua6wy0tg08xnd",
    "to": "cfxtest:achs3nehae0j6ksvy1bhrffsh1rtfrw1f6w1kzv46t",
    "value": "0x1",
    "gasPrice": "0x3b9aca00",
    "gas": "0xb316",
    "contractCreated": null,
    "data": "0xd0e30db0",
    "storageLimit": "0x0",
    "epochHeight": "0x6f26e7b",
    "chainId": "0x1",
    "status": "0x0",
    "v": "0x1",
    "r": "0x8a0a8568bab8a8cd0105e9caeb0fc0d5aa4a568cea123ce41a0299e9f18ebccf",
    "s": "0x7831351ea915b82625f728239eef279aaecb74a90abd066cf85e2f5d3e3a5f40"
  }
*/
func _TestTransactionHashShouldBeSignedTxHash(t *testing.T) {
	to := cfxaddress.MustNew("cfxtest:achs3nehae0j6ksvy1bhrffsh1rtfrw1f6w1kzv46t")
	utx := UnsignedTransaction{
		UnsignedTransactionBase: UnsignedTransactionBase{
			Nonce:    NewBigInt(0x4d073),
			Value:    NewBigInt(1),
			GasPrice: NewBigInt(0x3b9aca00),
			Gas:      NewBigInt(0xb316),
			// StorageLimit: NewUint64(0x0),
			EpochHeight: NewUint64(0x6f26e7b),
			ChainID:     NewUint(1),
		},
		To:   &to,
		Data: hexutils.HexToBytes("d0e30db0"),
	}
	fmt.Printf("utx %+v\n", utx)

	signedTx := SignedTransaction{
		UnsignedTransaction: utx,
		V:                   1,
		R:                   hexutils.HexToBytes("8a0a8568bab8a8cd0105e9caeb0fc0d5aa4a568cea123ce41a0299e9f18ebccf"),
		S:                   hexutils.HexToBytes("7831351ea915b82625f728239eef279aaecb74a90abd066cf85e2f5d3e3a5f40"),
	}

	utxHash, err := utx.Hash()
	assert.NoError(t, err)

	signedTxhash, err := signedTx.Hash()
	assert.NoError(t, err)

	assert.NotEqual(t, "0xbe47fb886dd24dba30cb25527b866b3139fb020ac344d5dc910b0140cce05478", hexutil.Encode(utxHash))
	assert.Equal(t, "0xbe47fb886dd24dba30cb25527b866b3139fb020ac344d5dc910b0140cce05478", hexutil.Encode(signedTxhash))
}
