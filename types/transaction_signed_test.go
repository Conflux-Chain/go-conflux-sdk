package types

import (
	"encoding/hex"
	"fmt"
	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"testing"

	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/status-im/keycard-go/hexutils"
	"github.com/stretchr/testify/assert"
)

func TestEncodeSignedTransaction(t *testing.T) {

	table := []struct {
		raw    string
		expect string
	}{
		{
			raw:    `{"UnsignedTransaction":{"From":null,"Nonce":"0x1","GasPrice":"0x1","Gas":"0x1","Value":"0x1","StorageLimit":"0x1","EpochHeight":"0x1","ChainID":"0x1","AccessList":null,"MaxPriorityFeePerGas":null,"MaxFeePerGas":null,"Type":null,"To":null,"Data":"0x01"},"V":1,"R":"0xca43b3f84e4afefcc6946d2953a0391774bd3c692015f64e51895b4a93fcba31","S":"0x36953d461a84e15367a463d4f7624970ad4c3833013f23103f2ea90a202e4aea"}`,
			expect: `f84dc901010180010101010101a0ca43b3f84e4afefcc6946d2953a0391774bd3c692015f64e51895b4a93fcba31a036953d461a84e15367a463d4f7624970ad4c3833013f23103f2ea90a202e4aea`,
		},
		{
			raw:    `{"UnsignedTransaction":{"From":null,"Nonce":"0x64","GasPrice":"0x64","Gas":"0x64","Value":"0x64","StorageLimit":"0x64","EpochHeight":"0x64","ChainID":"0x64","AccessList":[{"address":"net0:aapztdhx26tm0wgtueghvrj1nzsk651yjejbamakm5","storageKeys":["0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"]}],"MaxPriorityFeePerGas":null,"MaxFeePerGas":null,"Type":1,"To":"net0:aapztdhx26tm0wgtueghvrj1nzsk651yjejbamakm5","Data":"0x48656c6c6f2c20576f726c64"},"V":0,"R":"0x01","S":"0x01"}`,
			expect: `63667801f868f8636464649419578cf3c71eab48cf810c78b5175d5c9e6ef441646464648c48656c6c6f2c20576f726c64f838f79419578cf3c71eab48cf810c78b5175d5c9e6ef441e1a01234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef800101`,
		},
		{
			raw:    `{"UnsignedTransaction":{"From":null,"Nonce":"0x64","GasPrice":null,"Gas":"0x64","Value":"0x64","StorageLimit":"0x64","EpochHeight":"0x64","ChainID":"0x64","AccessList":[{"address":"net0:aapztdhx26tm0wgtueghvrj1nzsk651yjejbamakm5","storageKeys":["0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"]}],"MaxPriorityFeePerGas":"0x64","MaxFeePerGas":"0x64","Type":2,"To":"net0:aapztdhx26tm0wgtueghvrj1nzsk651yjejbamakm5","Data":"0x48656c6c6f2c20576f726c64"},"V":0,"R":"0x01","S":"0x01"}`,
			expect: `63667802f869f864646464649419578cf3c71eab48cf810c78b5175d5c9e6ef441646464648c48656c6c6f2c20576f726c64f838f79419578cf3c71eab48cf810c78b5175d5c9e6ef441e1a01234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef800101`,
		},
	}

	for _, item := range table {
		tx := SignedTransaction{}
		err := utils.JsonUnmarshal([]byte(item.raw), &tx)
		if err != nil {
			t.Fatal(err)
		}

		b, err := tx.Encode()
		if err != nil {
			t.Fatal(err)
		}

		if hex.EncodeToString(b) != item.expect {
			t.Fatalf("expect %s, but got %s", item.expect, hex.EncodeToString(b))
		}
	}

}
func TestDecodeSignedTransaction(t *testing.T) {

	table := []struct {
		raw    string
		expect string
	}{
		{
			raw:    `0xf84dc901010180010101010101a0ca43b3f84e4afefcc6946d2953a0391774bd3c692015f64e51895b4a93fcba31a036953d461a84e15367a463d4f7624970ad4c3833013f23103f2ea90a202e4aea`,
			expect: `{"UnsignedTransaction":{"From":null,"Nonce":"0x1","GasPrice":"0x1","Gas":"0x1","Value":"0x1","StorageLimit":"0x1","EpochHeight":"0x1","ChainID":"0x1","AccessList":null,"MaxPriorityFeePerGas":null,"MaxFeePerGas":null,"Type":null,"To":null,"Data":"0x01"},"V":1,"R":"0xca43b3f84e4afefcc6946d2953a0391774bd3c692015f64e51895b4a93fcba31","S":"0x36953d461a84e15367a463d4f7624970ad4c3833013f23103f2ea90a202e4aea"}`,
		},
		{
			raw:    `0x63667801f868f8636464649419578cf3c71eab48cf810c78b5175d5c9e6ef441646464648c48656c6c6f2c20576f726c64f838f79419578cf3c71eab48cf810c78b5175d5c9e6ef441e1a01234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef800101`,
			expect: `{"UnsignedTransaction":{"From":null,"Nonce":"0x64","GasPrice":"0x64","Gas":"0x64","Value":"0x64","StorageLimit":"0x64","EpochHeight":"0x64","ChainID":"0x64","AccessList":[{"address":"net0:aapztdhx26tm0wgtueghvrj1nzsk651yjejbamakm5","storageKeys":["0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"]}],"MaxPriorityFeePerGas":null,"MaxFeePerGas":null,"Type":1,"To":"net0:aapztdhx26tm0wgtueghvrj1nzsk651yjejbamakm5","Data":"0x48656c6c6f2c20576f726c64"},"V":0,"R":"0x01","S":"0x01"}`,
		},
		{
			raw:    `0x63667802f869f864646464649419578cf3c71eab48cf810c78b5175d5c9e6ef441646464648c48656c6c6f2c20576f726c64f838f79419578cf3c71eab48cf810c78b5175d5c9e6ef441e1a01234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef800101`,
			expect: `{"UnsignedTransaction":{"From":null,"Nonce":"0x64","GasPrice":null,"Gas":"0x64","Value":"0x64","StorageLimit":"0x64","EpochHeight":"0x64","ChainID":"0x64","AccessList":[{"address":"net0:aapztdhx26tm0wgtueghvrj1nzsk651yjejbamakm5","storageKeys":["0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"]}],"MaxPriorityFeePerGas":"0x64","MaxFeePerGas":"0x64","Type":2,"To":"net0:aapztdhx26tm0wgtueghvrj1nzsk651yjejbamakm5","Data":"0x48656c6c6f2c20576f726c64"},"V":0,"R":"0x01","S":"0x01"}`,
		},
	}

	for i, item := range table {
		// rawData := `0xf84dc901010180010101010101a0ca43b3f84e4afefcc6946d2953a0391774bd3c692015f64e51895b4a93fcba31a036953d461a84e15367a463d4f7624970ad4c3833013f23103f2ea90a202e4aea`
		tx := SignedTransaction{}
		b := mustDecodeStringToBytes(t, item.raw)
		err := tx.Decode(b, 0)
		if err != nil {
			t.Fatal(err)
		}

		b, err = utils.JsonMarshal(tx)
		if err != nil {
			t.Fatal(err)
		}

		// expect := `{"UnsignedTransaction":{"From":null,"Nonce":"0x1","GasPrice":"0x1","Gas":"0x1","Value":"0x1","StorageLimit":"0x1","EpochHeight":"0x1","ChainID":"0x1","AccessList":null,"MaxPriorityFeePerGas":null,"MaxFeePerGas":null,"Type":null,"To":null,"Data":"0x01"},"V":1,"R":"0xca43b3f84e4afefcc6946d2953a0391774bd3c692015f64e51895b4a93fcba31","S":"0x36953d461a84e15367a463d4f7624970ad4c3833013f23103f2ea90a202e4aea"}`

		if string(b) != item.expect {
			t.Fatalf("item %d expect: %s, got: %s", i, item.expect, string(b))
		}
	}

}

func TestHashSignedTransaction(t *testing.T) {
	to := cfxaddress.MustNew("cfxtest:achs3nehae0j6ksvy1bhrffsh1rtfrw1f6w1kzv46t")
	utx := UnsignedTransaction{
		UnsignedTransactionBase: UnsignedTransactionBase{
			Nonce:        NewBigInt(0x4d073),
			Value:        NewBigInt(1),
			GasPrice:     NewBigInt(0x3b9aca00),
			Gas:          NewBigInt(0xb316),
			StorageLimit: NewUint64(0x0),
			EpochHeight:  NewUint64(0x6f26e7b),
			ChainID:      NewUint(1),
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

func TestHashUnsignedTransaction(t *testing.T) {

	table := []struct {
		tx         string
		rlpExpect  string
		hashExpect string
	}{
		{
			tx:         `{"From":null,"Nonce":"0x64","GasPrice":"0x64","Gas":"0x64","Value":"0x64","StorageLimit":"0x64","EpochHeight":"0x64","ChainID":"0x64","To":"net0:aapztdhx26tm0wgtueghvrj1nzsk651yjejbamakm5","Data":"0x48656c6c6f2c20576f726c64"}`,
			rlpExpect:  `e96464649419578cf3c71eab48cf810c78b5175d5c9e6ef441646464648c48656c6c6f2c20576f726c64`,
			hashExpect: `5487fa843420144fd78f19bb86e9da81040e50423ab3ec2818ad4b6c86fcecc2`,
		},
		{
			tx:         `{"From":null,"Nonce":"0x64","GasPrice":"0x64","Gas":"0x64","Value":"0x64","StorageLimit":"0x64","EpochHeight":"0x64","ChainID":"0x64","AccessList":[{"address":"net0:aapztdhx26tm0wgtueghvrj1nzsk651yjejbamakm5","storageKeys":["0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"]}],"MaxPriorityFeePerGas":null,"MaxFeePerGas":null,"Type":1,"To":"net0:aapztdhx26tm0wgtueghvrj1nzsk651yjejbamakm5","Data":"0x48656c6c6f2c20576f726c64"}`,
			rlpExpect:  `63667801f8636464649419578cf3c71eab48cf810c78b5175d5c9e6ef441646464648c48656c6c6f2c20576f726c64f838f79419578cf3c71eab48cf810c78b5175d5c9e6ef441e1a01234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef`,
			hashExpect: `690d58e271b90254e7954147846d5de0f76f3649510bb58a5f26e4fef8d601ba`,
		},
		{
			tx:         `{"From":null,"Nonce":"0x64","GasPrice":null,"Gas":"0x64","Value":"0x64","StorageLimit":"0x64","EpochHeight":"0x64","ChainID":"0x64","AccessList":[{"address":"net0:aapztdhx26tm0wgtueghvrj1nzsk651yjejbamakm5","storageKeys":["0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"]}],"MaxPriorityFeePerGas":"0x64","MaxFeePerGas":"0x64","Type":2,"To":"net0:aapztdhx26tm0wgtueghvrj1nzsk651yjejbamakm5","Data":"0x48656c6c6f2c20576f726c64"}`,
			rlpExpect:  `63667802f864646464649419578cf3c71eab48cf810c78b5175d5c9e6ef441646464648c48656c6c6f2c20576f726c64f838f79419578cf3c71eab48cf810c78b5175d5c9e6ef441e1a01234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef`,
			hashExpect: `3da56dbe2b76c41135c2429f3035cd79b1abb68902cf588075c30d4912e71cf3`,
		},
	}

	for _, item := range table {
		tx := UnsignedTransaction{}
		err := utils.JsonUnmarshal([]byte(item.tx), &tx)
		if err != nil {
			t.Fatal(err)
		}

		rlp, err := tx.Encode()
		if err != nil {
			t.Fatal(err)
		}

		hash, err := tx.Hash()
		if err != nil {
			t.Fatal(err)
		}

		if hex.EncodeToString(rlp) != item.rlpExpect {
			t.Fatalf("expect %s, but got %s", item.rlpExpect, hex.EncodeToString(rlp))
		}

		if hex.EncodeToString(hash) != item.hashExpect {
			t.Fatalf("expect %s, but got %s", item.hashExpect, hex.EncodeToString(hash))
		}
	}
}

func TestDecodeUnsignedTransaction(t *testing.T) {
	table := []struct {
		raw    string
		expect string
	}{
		{
			raw:    `0xe96464649419578cf3c71eab48cf810c78b5175d5c9e6ef441646464648c48656c6c6f2c20576f726c64`,
			expect: `{"From":null,"Nonce":"0x64","GasPrice":"0x64","Gas":"0x64","Value":"0x64","StorageLimit":"0x64","EpochHeight":"0x64","ChainID":"0x64","AccessList":null,"MaxPriorityFeePerGas":null,"MaxFeePerGas":null,"Type":null,"To":"net0:aapztdhx26tm0wgtueghvrj1nzsk651yjejbamakm5","Data":"0x48656c6c6f2c20576f726c64"}`,
		},
		{
			raw:    `0x63667801f8636464649419578cf3c71eab48cf810c78b5175d5c9e6ef441646464648c48656c6c6f2c20576f726c64f838f79419578cf3c71eab48cf810c78b5175d5c9e6ef441e1a01234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef`,
			expect: `{"From":null,"Nonce":"0x64","GasPrice":"0x64","Gas":"0x64","Value":"0x64","StorageLimit":"0x64","EpochHeight":"0x64","ChainID":"0x64","AccessList":[{"address":"net0:aapztdhx26tm0wgtueghvrj1nzsk651yjejbamakm5","storageKeys":["0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"]}],"MaxPriorityFeePerGas":null,"MaxFeePerGas":null,"Type":1,"To":"net0:aapztdhx26tm0wgtueghvrj1nzsk651yjejbamakm5","Data":"0x48656c6c6f2c20576f726c64"}`,
		},
		{
			raw:    `0x63667802f864646464649419578cf3c71eab48cf810c78b5175d5c9e6ef441646464648c48656c6c6f2c20576f726c64f838f79419578cf3c71eab48cf810c78b5175d5c9e6ef441e1a01234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef`,
			expect: `{"From":null,"Nonce":"0x64","GasPrice":null,"Gas":"0x64","Value":"0x64","StorageLimit":"0x64","EpochHeight":"0x64","ChainID":"0x64","AccessList":[{"address":"net0:aapztdhx26tm0wgtueghvrj1nzsk651yjejbamakm5","storageKeys":["0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"]}],"MaxPriorityFeePerGas":"0x64","MaxFeePerGas":"0x64","Type":2,"To":"net0:aapztdhx26tm0wgtueghvrj1nzsk651yjejbamakm5","Data":"0x48656c6c6f2c20576f726c64"}`,
		},
	}

	for i, item := range table {
		tx := UnsignedTransaction{}
		b := mustDecodeStringToBytes(t, item.raw)
		err := tx.Decode(b, 0)
		if err != nil {
			t.Fatal(err)
		}

		b, err = utils.JsonMarshal(tx)
		if err != nil {
			t.Fatal(err)
		}

		if string(b) != item.expect {
			t.Fatalf("item %d expect: %s, got: %s", i, item.expect, string(b))
		}
	}
}

func mustDecodeStringToBytes(t *testing.T, s string) []byte {
	b, e := hex.DecodeString(s[2:])
	if e != nil {
		t.Fatal(e)
	}
	return b
}

func TestTemp(t *testing.T) {
	accessList := AccessList([]AccessTuple{
		{
			Address: cfxaddress.MustNew("0x19578CF3c71eaB48cF810c78B5175d5c9E6Ef441"),
			StorageKeys: []common.Hash{
				common.HexToHash("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"),
			},
		},
	})
	to := cfxaddress.MustNew("0x19578CF3c71eaB48cF810c78B5175d5c9E6Ef441")
	tx := SignedTransaction{
		UnsignedTransaction: UnsignedTransaction{
			UnsignedTransactionBase: UnsignedTransactionBase{
				Type:         TRANSACTION_TYPE_2930.Ptr(),
				Nonce:        NewBigInt(100),
				GasPrice:     NewBigInt(100),
				Gas:          NewBigInt(100),
				Value:        NewBigInt(100),
				StorageLimit: NewUint64(100),
				EpochHeight:  NewUint64(100),
				ChainID:      NewUint(100),
				AccessList:   accessList,
			},
			To:   &to,
			Data: []byte("Hello, World"),
		},
		R: NewBytes([]byte{1}),
		S: NewBytes([]byte{1}),
		V: byte(0),
	}
	j, _ := utils.JsonMarshal(tx)
	fmt.Println(string(j))

	tx = SignedTransaction{
		UnsignedTransaction: UnsignedTransaction{
			UnsignedTransactionBase: UnsignedTransactionBase{
				Type:                 TRANSACTION_TYPE_1559.Ptr(),
				Nonce:                NewBigInt(100),
				MaxPriorityFeePerGas: NewBigInt(100),
				MaxFeePerGas:         NewBigInt(100),
				Gas:                  NewBigInt(100),
				Value:                NewBigInt(100),
				StorageLimit:         NewUint64(100),
				EpochHeight:          NewUint64(100),
				ChainID:              NewUint(100),
				AccessList:           accessList,
			},
			To:   &to,
			Data: []byte("Hello, World"),
		},
		R: NewBytes([]byte{1}),
		S: NewBytes([]byte{1}),
		V: byte(0),
	}
	j, _ = utils.JsonMarshal(tx)
	fmt.Println(string(j))
}

func TestTemp2(t *testing.T) {
	accessList := AccessList([]AccessTuple{
		{
			Address: cfxaddress.MustNew("0x19578CF3c71eaB48cF810c78B5175d5c9E6Ef441"),
			StorageKeys: []common.Hash{
				common.HexToHash("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"),
			},
		},
	})
	to := cfxaddress.MustNew("0x19578CF3c71eaB48cF810c78B5175d5c9E6Ef441")
	tx := SignedTransaction{
		UnsignedTransaction: UnsignedTransaction{
			UnsignedTransactionBase: UnsignedTransactionBase{
				Type:         TRANSACTION_TYPE_2930.Ptr(),
				Nonce:        NewBigInt(100),
				GasPrice:     NewBigInt(100),
				Gas:          NewBigInt(100),
				Value:        NewBigInt(100),
				StorageLimit: NewUint64(100),
				EpochHeight:  NewUint64(100),
				ChainID:      NewUint(100),
				AccessList:   accessList,
			},
			To:   &to,
			Data: []byte("Hello, World"),
		},
		R: NewBytes([]byte{1}),
		S: NewBytes([]byte{1}),
		V: byte(0),
	}
	j, _ := utils.JsonMarshal(tx.UnsignedTransaction)
	fmt.Println(string(j))

	tx = SignedTransaction{
		UnsignedTransaction: UnsignedTransaction{
			UnsignedTransactionBase: UnsignedTransactionBase{
				Type:                 TRANSACTION_TYPE_1559.Ptr(),
				Nonce:                NewBigInt(100),
				MaxPriorityFeePerGas: NewBigInt(100),
				MaxFeePerGas:         NewBigInt(100),
				Gas:                  NewBigInt(100),
				Value:                NewBigInt(100),
				StorageLimit:         NewUint64(100),
				EpochHeight:          NewUint64(100),
				ChainID:              NewUint(100),
				AccessList:           accessList,
			},
			To:   &to,
			Data: []byte("Hello, World"),
		},
		R: NewBytes([]byte{1}),
		S: NewBytes([]byte{1}),
		V: byte(0),
	}
	j, _ = utils.JsonMarshal(tx.UnsignedTransaction)
	fmt.Println(string(j))
}
