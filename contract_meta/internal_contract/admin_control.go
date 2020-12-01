package internalcontract

import (
	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

// AdminControl contract
type AdminControl struct {
	sdk.Contract
}

var adminControl *AdminControl

// NewAdminControl gets the AdminControl contract object
func NewAdminControl(client sdk.ClientOperator) *AdminControl {

	if adminControl == nil {
		abi := getAdminControlAbi()
		address := getAdminControlAddress()
		contract, _ := sdk.NewContract([]byte(abi), client, &address)
		adminControl = &AdminControl{Contract: *contract}
	}
	return adminControl
}

// Destroy destroies contract `contractAddr`.
func (ac *AdminControl) Destroy(option *types.ContractMethodSendOption, contractAddr types.Address) (*types.Hash, error) {
	return ac.SendTransaction(option, "destroy", contractAddr.ToCommonAddress())
}

// GetAdmin returns admin of specific contract
func (ac *AdminControl) GetAdmin(option *types.ContractMethodCallOption, contractAddr types.Address) (result *types.Address, err error) {
	var tmp *common.Address = &common.Address{}
	err = ac.Call(option, tmp, "getAdmin", contractAddr.ToCommonAddress())
	if err != nil {
		return nil, err
	}
	return types.NewAddressFromCommon(*tmp), nil
}

// SetAdmin sets the administrator of contract `contractAddr` to `newAdmin`.
func (ac *AdminControl) SetAdmin(option *types.ContractMethodSendOption, contractAddr types.Address, newAdmin types.Address) (*types.Hash, error) {
	return ac.SendTransaction(option, "setAdmin", contractAddr.ToCommonAddress(), newAdmin.ToCommonAddress())
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

func getAdminControlAddress() types.Address {
	return types.Address("0x0888000000000000000000000000000000000000")
}
