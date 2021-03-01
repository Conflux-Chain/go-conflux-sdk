package sdk

import (
	"reflect"
	"testing"

	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/accounts"
)

func TestGetAccountMatchAddress(t *testing.T) {
	am := NewAccountManager("./tmp/keystore", 1)
	testNum := 10

	accounts := []accounts.Account{}
	existLen := len(am.List())
	if existLen < testNum {
		for i := 0; i < testNum-existLen; i++ {
			_, e := am.Create("123")
			if e != nil {
				t.Fatalf("failed to create account %v", e.Error())
			}
		}
	}

	addresses := am.List()
	for i := 0; i < testNum; i++ {
		account, _ := am.account(addresses[i])
		accounts = append(accounts, account)
	}

	for i := 0; i < testNum; i++ {
		if !reflect.DeepEqual(accounts[i].Address.Bytes()[1:], addresses[i].MustGetCommonAddress().Bytes()[1:]) {
			t.Fatalf("%v expect %x, actual %x", i, addresses[i].MustGetCommonAddress().Bytes()[1:], accounts[i].Address.Bytes()[1:])
		}
	}

	account, err := am.account(cfxaddress.MustNewFromHex("0x15359711FDDfe27c6009F63C9E0A5d26cC78ED44"))
	if err == nil {
		t.Fatalf("expect <not found> error, actual %v", account.Address)
	}
}
