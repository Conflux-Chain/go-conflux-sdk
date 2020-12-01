package main

import (
	"fmt"

	internalContract "github.com/Conflux-Chain/go-conflux-sdk/contract_meta/internal_contract"
	"github.com/Conflux-Chain/go-conflux-sdk/example/context"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
)

func testAdmin() {
	adminControl := internalContract.NewAdminControl(client)
	// get admin
	admin, err := adminControl.GetAdmin(nil, config.ERC20Address)
	if err != nil {
		fmt.Printf("get admin of %v error: %v\n", config.ERC20Address, err)
		return
	} else {
		fmt.Printf("admin of %v is %v\n", config.ERC20Address, admin)
	}

	if *defaultAccount != *admin {
		panic("admin is not " + defaultAccount.String() + "\n")
	}

	// destory
	txhash, err := adminControl.Destroy(nil, config.ERC20Address)
	if err != nil {
		fmt.Printf("detory error %v\n", err)
		return
	}

	context.WaitPacked(client, *txhash)
	code, err := client.GetCode(config.ERC20Address)
	if err != nil {
		fmt.Printf("destory done")
		return
	}
	fmt.Printf("destory error, contract code still exist: %v\n", code)

	// set admin
	config = context.PrepareForClientExample()

	txhash, err = adminControl.SetAdmin(nil, config.ERC20Address, types.Address("0x0000000000000000000000000000000000000000"))
	if err != nil {
		fmt.Printf("set admin error %v\n", err)
		return
	}

	context.WaitPacked(client, *txhash)
	admin, _ = adminControl.GetAdmin(nil, config.ERC20Address)
	fmt.Printf("set admin done, new admin of %v is %v\n", config.ERC20Address, admin)
}
