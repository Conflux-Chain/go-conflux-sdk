package types

import (
	"encoding/hex"
	"encoding/json"
	"testing"
)

func TestDecodeTransaction(t *testing.T) {
	rawData := `0xf84dc901010180010101010101a0ca43b3f84e4afefcc6946d2953a0391774bd3c692015f64e51895b4a93fcba31a036953d461a84e15367a463d4f7624970ad4c3833013f23103f2ea90a202e4aea`
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

	expect := `{"UnsignedTransaction":{"From":null,"Nonce":"0x1","GasPrice":"0x1","Gas":"0x1","Value":"0x1","StorageLimit":"0x1","EpochHeight":"0x1","ChainID":"0x1","To":null,"Data":"0x01"},"V":1,"R":"0xca43b3f84e4afefcc6946d2953a0391774bd3c692015f64e51895b4a93fcba31","S":"0x36953d461a84e15367a463d4f7624970ad4c3833013f23103f2ea90a202e4aea"}`

	if string(b) != expect {
		t.Fatalf("expect: %s, got: %s", expect, string(b))
	}
}

func TestEncodeTransaction(t *testing.T) {
	jsonStr := `{"UnsignedTransaction":{"From":"CFXTEST:TYPE.USER:AAR7X4R8MKRNW39GGS8RZ40J1ZNWH5MRRPUFPR2U76","Nonce":"0x1","GasPrice":"0x1","Gas":"0x1","Value":"0x1","StorageLimit":"0x1","EpochHeight":"0x1","ChainID":"0x1","To":null,"Data":"0x01"},"V":1,"R":"0xca43b3f84e4afefcc6946d2953a0391774bd3c692015f64e51895b4a93fcba31","S":"0x36953d461a84e15367a463d4f7624970ad4c3833013f23103f2ea90a202e4aea"}`
	tx := SignedTransaction{}
	err := json.Unmarshal([]byte(jsonStr), &tx)
	if err != nil {
		t.Fatal(err)
	}

	b, err := tx.Encode()
	if err != nil {
		t.Fatal(err)
	}

	expect := `f84dc901010180010101010101a0ca43b3f84e4afefcc6946d2953a0391774bd3c692015f64e51895b4a93fcba31a036953d461a84e15367a463d4f7624970ad4c3833013f23103f2ea90a202e4aea`
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
