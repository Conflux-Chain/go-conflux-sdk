package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	address "github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
)

var am *sdk.AccountManager

func init() {
	initAccountManager()
	fmt.Println("init account manager done")
}

func main() {
	listAccount()
	creatAccount()
	importAccount()
	hasChecksumChecked()
	updateAccount()
	deleteAccount()
	signTx()
	decodeRawTx()
}

func initAccountManager() *sdk.AccountManager {
	keydir := "./keystore"
	am = sdk.NewAccountManager(keydir, 1)
	return am
}

func listAccount() {
	fmt.Printf("account list: %+v\n\n", am.List())
}

func creatAccount() {
	am := initAccountManager()
	addr, err := am.Create("hello")
	if err != nil {
		fmt.Println("create account error", err)
		return
	}
	fmt.Println("creat account:", addr)
}

func importAccount() {
	am := initAccountManager()
	dir, _ := os.Getwd()
	keydir := dir + "/keystore_tmp"
	files, err := ioutil.ReadDir(keydir)
	if err != nil {
		panic(err)
	}
	if len(files) == 0 {
		panic("no files in directory:" + dir)
	}

	addr, err := am.Import(keydir+"/"+files[0].Name(), "hello", "hello")
	if err != nil {
		fmt.Println("import account error:", err)
		return
	}
	fmt.Println("import account done:", addr)
}

func updateAccount() {
	address := cfxaddress.MustNewFromHex("0x14b899ed1cd49da2c11093606465baa102662ab5", 1)
	err := am.Update(address, "hello", "hello world")
	if err != nil {
		fmt.Printf("update address error: %v \n\n", err)
		return
	}
	fmt.Printf("update address %v done\n\n", address)
}

func deleteAccount() {
	address := cfxaddress.MustNewFromHex("0x14b899ed1cd49da2c11093606465baa102662ab5", 1)
	err := am.Delete(address, "hello world")
	if err != nil {
		fmt.Printf("delete address error: %v \n\n", err)
		return
	}
	fmt.Printf("delete address %v done\n\n", address)
}

func signTx() []byte {
	am := initAccountManager()

	from := cfxaddress.MustNewFromHex("0x1ceb7b1c5252ae3eaaf19d3a785cfbea48cc37f7", 1)
	to := cfxaddress.MustNewFromHex("0x10f4bcf113e0b896d9b34294fd3da86b4adf0302", 1)
	unSignedTx := types.UnsignedTransaction{
		UnsignedTransactionBase: types.UnsignedTransactionBase{
			From:        &from,
			Value:       types.NewBigInt(100),
			Gas:         types.NewBigInt(21000),
			GasPrice:    types.NewBigInt(100000000),
			Nonce:       types.NewBigInt(1),
			EpochHeight: types.NewUint64(1),
			ChainID:     types.NewUint(1),
		},
		To: &to,
	}

	signedTx, err := am.SignAndEcodeTransactionWithPassphrase(unSignedTx, "hello")
	if err != nil {
		fmt.Printf("signed tx %+v error:%v\n\n", unSignedTx, err)
		return nil
	}
	fmt.Printf("signed tx %+v result:\n0x%x\n\n", unSignedTx, signedTx)
	return signedTx
}

func decodeRawTx() {
	rawTx, _ := hex.DecodeString("f867e3018405f5e1008252089410f4bcf113e0b896d9b34294fd3da86b4adf0302648001018080a072aa2777c4b8cee3829ea3fb9658276e40cc4234eb05f176459042e48f69428da07a9bbee20b9a219907c91b562b64ee2e9456d2f280c31ce98736d0feb5e47372")
	tx := new(types.SignedTransaction)
	err := tx.Decode(rawTx, 1)
	if err != nil {
		fmt.Printf("decoded rawTx error: %+v\n\n", err)
		return
	}
	fmt.Printf("decoded rawTx %x result: %+v\n\n", rawTx, tx)
}

func hasChecksumChecked() {
	_, err := am.Export(address.MustNewFromHex("0x14b899ed1cd49da2c11093606465baa102662ab5", 1), "hello")
	if err != nil {
		panic(err)
	}

	_, err = am.Export(address.MustNewFromHex("0x14b899eD1cD49Da2c11093606465Baa102662ab5", 1), "hello")
	if err != nil {
		panic(err)
	}
}
