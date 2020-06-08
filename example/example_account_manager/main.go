package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
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
	updateAccount()
	deleteAccount()
	signTx()
	decodeRawTx()
}

func initAccountManager() *sdk.AccountManager {
	keydir := "./keystore"
	am = sdk.NewAccountManager(keydir)
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
	address := types.Address("0x14b899ed1cd49da2c11093606465baa102662ab5")
	err := am.Update(address, "hello", "hello world")
	if err != nil {
		fmt.Printf("update address error: %v \n\n", err)
		return
	}
	fmt.Printf("update address %v done\n\n", address)
}

func deleteAccount() {
	address := types.Address("0x14b899ed1cd49da2c11093606465baa102662ab5")
	err := am.Delete(address, "hello world")
	if err != nil {
		fmt.Printf("delete address error: %v \n\n", err)
		return
	}
	fmt.Printf("delete address %v done\n\n", address)
}

func signTx() []byte {
	am := initAccountManager()
	unSignedTx := types.UnsignedTransaction{
		UnsignedTransactionBase: types.UnsignedTransactionBase{
			From:        types.NewAddress("0x1ceb7b1c5252ae3eaaf19d3a785cfbea48cc37f7"),
			Value:       types.NewBigInt(100),
			Gas:         types.NewBigInt(21000),
			GasPrice:    types.NewBigInt(100000000),
			Nonce:       types.NewBigInt(1),
			EpochHeight: types.NewBigInt(1),
		},
		To: types.NewAddress("0x10f4bcf113e0b896d9b34294fd3da86b4adf0302"),
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
	rawTx, _ := hex.DecodeString("f866e280018252089410f4bcf113e0b896d9b34294fd3da86b4adf0302648083025d93808080a0fab5b7bfe91ebb7367e92cff27ab153153eb80e62e28852a9c4a0cc2133a4161a048aee7e729790c72a1d75329372439b36e923ef15dbf753e66a909469d2d1a2d")
	tx := new(types.SignedTransaction)
	err := tx.Decode(rawTx)
	if err != nil {
		fmt.Printf("decoded rawTx error: %+v\n\n", err)
		return
	}
	fmt.Printf("decoded rawTx %x result: %+v\n\n", rawTx, tx)
}
