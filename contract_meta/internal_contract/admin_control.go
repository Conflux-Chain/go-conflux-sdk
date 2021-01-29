package internalcontract

import (
	"fmt"
	"sync"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	address "github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// AdminControl contract
type AdminControl struct {
	sdk.Contract
}

var adminControl *AdminControl
var adminControlMu sync.Mutex

// NewAdminControl gets the AdminControl contract object
func NewAdminControl(client sdk.ClientOperator) (ac AdminControl, err error) {
	if adminControl == nil {
		adminControlMu.Lock()
		defer adminControlMu.Unlock()
		abi := getAdminControlAbi()
		address, e := getAdminControlAddress(client)
		if e != nil {
			return ac, errors.Wrap(e, "failed to get admin control contract address")
		}
		contract, e := sdk.NewContract([]byte(abi), client, &address)
		if e != nil {
			return ac, errors.Wrap(e, "failed to new admin control contract")
		}
		adminControl = &AdminControl{Contract: *contract}
	}
	return *adminControl, nil
}

// Destroy destroies contract `contractAddr`.
func (ac *AdminControl) Destroy(option *types.ContractMethodSendOption, contractAddr types.Address) (types.Hash, error) {
	return ac.SendTransaction(option, "destroy", contractAddr.MustGetCommonAddress())
}

// GetAdmin returns admin of specific contract
func (ac *AdminControl) GetAdmin(option *types.ContractMethodCallOption, contractAddr types.Address) (types.Address, error) {
	empty := cfxaddress.Address{}

	var tmp *common.Address = &common.Address{}
	fmt.Printf("contract addr common address: %x\n", contractAddr.MustGetCommonAddress())
	err := ac.Call(option, tmp, "getAdmin", contractAddr.MustGetCommonAddress())
	if err != nil {
		return empty, errors.Wrap(err, "failed to call getAdmin")
	}

	addr, err := address.NewFromCommon(*tmp)
	if err != nil {
		return empty, errors.Wrapf(err, "failed to new address from common %v", *tmp)
	}
	err = addr.CompleteByClient(ac.Client)
	return addr, errors.Wrapf(err, "failed to complete network type")
}

// SetAdmin sets the administrator of contract `contractAddr` to `newAdmin`.
func (ac *AdminControl) SetAdmin(option *types.ContractMethodSendOption, contractAddr types.Address, newAdmin types.Address) (types.Hash, error) {
	return ac.SendTransaction(option, "setAdmin", contractAddr.MustGetCommonAddress(), newAdmin.MustGetCommonAddress())
}

func getAdminControlAbi() string {
	return `
[
        {
            "inputs": [
                {
                    "internalType": "address",
                    "name": "contractAddr",
                    "type": "address"
                }
            ],
            "name": "destroy",
            "outputs": [],
            "stateMutability": "nonpayable",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "address",
                    "name": "contractAddr",
                    "type": "address"
                }
            ],
            "name": "getAdmin",
            "outputs": [
                {
                    "internalType": "address",
                    "name": "",
                    "type": "address"
                }
            ],
            "stateMutability": "view",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "address",
                    "name": "contractAddr",
                    "type": "address"
                },
                {
                    "internalType": "address",
                    "name": "newAdmin",
                    "type": "address"
                }
            ],
            "name": "setAdmin",
            "outputs": [],
            "stateMutability": "nonpayable",
            "type": "function"
        }
    ]
`
}

func getAdminControlAddress(client sdk.ClientOperator) (types.Address, error) {
	addr := cfxaddress.MustNewFromHex("0x0888000000000000000000000000000000000000")
	err := addr.CompleteByClient(client)
	return addr, err
}
