package sdk

import (
	"reflect"
	"testing"

	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/accounts"
)

func TestGetAccountMatchAddress(t *testing.T) {
	am := NewAccountManager("./tmp/keystore", 1)

	accounts := []accounts.Account{}

	if len(am.List()) < 10 {
		for i := 0; i < 10-len(am.List()); i++ {
			am.Create("123")
		}
	}

	addresses := am.List()

	for i := 0; i < 10; i++ {
		account, _ := am.account(addresses[i])
		accounts = append(accounts, account)
	}

	for i := 0; i < 10; i++ {
		if !reflect.DeepEqual(accounts[i].Address.Bytes()[1:], addresses[i].MustGetCommonAddress().Bytes()[1:]) {
			t.Fatalf("%v expect %x, actual %x", i, addresses[i].MustGetCommonAddress().Bytes()[1:], accounts[i].Address.Bytes()[1:])
		}
	}

	account, err := am.account(cfxaddress.MustNewFromHex("0x15359711FDDfe27c6009F63C9E0A5d26cC78ED44"))
	if err == nil {
		t.Fatalf("expect <not found> error, actual %v", account.Address)
	}
}
