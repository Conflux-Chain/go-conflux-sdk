package main

import (
	"fmt"
	"math/big"
	"time"

	internalcontract "github.com/Conflux-Chain/go-conflux-sdk/contract_meta/internal_contract"
	"github.com/Conflux-Chain/go-conflux-sdk/example/context"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func testSponsor() {
	sponsor, err := internalcontract.NewSponsor(client)
	context.PanicIfErrf(err, "failed to new sponsor")
	// new erc20 contract
	contract := context.DeployNewErc20()
	contractAddr := *contract.Address
	time.Sleep(time.Second * 5)

	// SetSponsorForGas
	txhash, err := sponsor.SetSponsorForGas(&types.ContractMethodSendOption{
		Value: types.NewBigInt(2e18),
		Nonce: context.GetNextNonceAndIncrease(),
	}, *contract.Address, big.NewInt(1e9))
	context.PanicIfErrf(err, "SetSponsorForGas panic")
	receipt, err := client.WaitForTransationReceipt(txhash, time.Second)
	fmt.Printf("SetSponsorForGas is success?%v\n", receipt.OutcomeStatus == hexutil.Uint64(0))

	// GetSponsorForGas
	sponsorUser, err := sponsor.GetSponsorForGas(nil, contractAddr)
	context.PanicIfErrf(err, "GetSponsorForGas panic")
	fmt.Printf("sponsorUser is %v\n", sponsorUser)

	// GetSponsoredBalanceForGas
	sponsorBalaceForGas, err := sponsor.GetSponsoredBalanceForGas(nil, contractAddr)
	context.PanicIfErrf(err, "GetSponsoredBalanceForGas panic")
	fmt.Printf("sponsorBalaceForGas is %v\n", sponsorBalaceForGas)

	// GetSponsoredGasFeeUpperBound
	sponsorBalaceForGasUpperBound, err := sponsor.GetSponsoredGasFeeUpperBound(nil, contractAddr)
	context.PanicIfErrf(err, "GetSponsoredGasFeeUpperBound panic")
	fmt.Printf("sponsorBalaceForGasUpperBound is %v\n", sponsorBalaceForGasUpperBound)

	// SetSponsorForCollateral
	txhash, err = sponsor.SetSponsorForCollateral(&types.ContractMethodSendOption{
		Value: types.NewBigInt(2e18),
		Nonce: context.GetNextNonceAndIncrease(),
	}, *contract.Address)
	context.PanicIfErrf(err, "SetSponsorForCollateral panic")
	receipt, err = client.WaitForTransationReceipt(txhash, time.Second)
	fmt.Printf("SetSponsorForCollateral is success? %v\n", receipt.OutcomeStatus == hexutil.Uint64(0))

	// GetSponsorForCollateral
	sponsorForCollateral, err := sponsor.GetSponsorForCollateral(nil, contractAddr)
	context.PanicIfErrf(err, "GetSponsorForCollateral panic")
	fmt.Printf("sponsorForCollateral is %v\n", sponsorForCollateral)

	// GetSponsoredBalanceForCollateral
	sponsoredBalanceForCollateral, err := sponsor.GetSponsoredBalanceForCollateral(nil, contractAddr)
	context.PanicIfErrf(err, "GetSponsoredBalanceForCollateral panic")
	fmt.Printf("sponsorForCollateral is %v\n", sponsoredBalanceForCollateral)

	// AddPrivilegeByAdmin
	txhash, err = sponsor.AddPrivilegeByAdmin(&types.ContractMethodSendOption{
		Nonce: context.GetNextNonceAndIncrease()},
		*contract.Address, []types.Address{cfxaddress.MustNewFromHex("0x15294fd6b3452e657ac2424391d08250340970d4"), cfxaddress.MustNewFromHex("0x1cfa93e2e549c27a84b2121c2da532e18353ec5b")})
	context.PanicIfErrf(err, "AddPrivilegeByAdmin panic")
	receipt, err = client.WaitForTransationReceipt(txhash, time.Second)
	fmt.Printf("AddPrivilegeByAdmin is success? %v\n", receipt.OutcomeStatus == hexutil.Uint64(0))

	// IsAllWhitelisted
	isAllWhitelisted, err := sponsor.IsAllWhitelisted(nil, contractAddr)
	context.PanicIfErrf(err, "IsAllWhitelisted panic")
	fmt.Printf("isAllWhitelisted shold be false: %v\n", isAllWhitelisted)

	// IsWhitelisted
	isWhitelisted, err := sponsor.IsWhitelisted(nil, contractAddr, cfxaddress.MustNewFromHex("0x15294fd6b3452e657ac2424391d08250340970d4"))
	context.PanicIfErrf(err, "IsWhitelisted panic")
	fmt.Printf("isWhitelisted should be true: %v\n", isWhitelisted)

	// RemovePrivilegeByAdmin
	txhash, err = sponsor.RemovePrivilegeByAdmin(&types.ContractMethodSendOption{
		Nonce: context.GetNextNonceAndIncrease()},
		*contract.Address, []types.Address{cfxaddress.MustNewFromHex("0x15294fd6b3452e657ac2424391d08250340970d4")})
	context.PanicIfErrf(err, "RemovePrivilegeByAdmin panic")

	receipt, err = client.WaitForTransationReceipt(txhash, time.Second)
	context.PanicIfErrf(err, "failed to get tx receipt of %v ", txhash)

	fmt.Printf("RemovePrivilegeByAdmin is success? %v\n", receipt.OutcomeStatus == hexutil.Uint64(0))

	// IsWhitelisted
	isWhitelisted, err = sponsor.IsWhitelisted(nil, contractAddr, cfxaddress.MustNewFromHex("0x15294fd6b3452e657ac2424391d08250340970d4"))
	context.PanicIfErrf(err, "IsWhitelisted panic")
	fmt.Printf("isWhitelisted should be false: %v\n", isWhitelisted)
}
