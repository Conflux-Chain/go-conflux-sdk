package internalcontract

import (
	"math/big"
	"sync"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// Sponsor represents SponsorWhitelistControl contract
type Sponsor struct {
	sdk.Contract
}

var sponsor *Sponsor
var sponsorMu sync.Mutex

// NewSponsor gets the SponsorWhitelistControl contract object
func NewSponsor(client sdk.ClientOperator) (s Sponsor, err error) {
	if sponsor == nil {
		sponsorMu.Lock()
		defer sponsorMu.Unlock()
		abi := getSponsorAbi()
		address, e := getSponsorHexAddress(client)
		if e != nil {
			return s, errors.Wrap(e, "failed to get sponsor contract address")
		}
		contract, e := sdk.NewContract([]byte(abi), client, &address)
		if e != nil {
			return s, errors.Wrap(e, "failed to new sponsor contract")
		}
		sponsor = &Sponsor{Contract: *contract}
	}
	return *sponsor, nil
}

// GetSponsorForGas gets gas sponsor address of specific contract
func (s *Sponsor) GetSponsorForGas(option *types.ContractMethodCallOption, contractAddr types.Address) (address types.Address, err error) {
	var tmp *common.Address = &common.Address{}
	err = s.Call(option, tmp, "getSponsorForGas", contractAddr.MustGetCommonAddress())
	if err != nil {
		return
	}
	address = cfxaddress.MustNewFromCommon(*tmp)
	err = address.CompleteByClient(s.Client)
	return
}

// GetSponsoredBalanceForGas gets current Sponsored Balance for gas
func (s *Sponsor) GetSponsoredBalanceForGas(option *types.ContractMethodCallOption, contractAddr types.Address) (*big.Int, error) {
	balance := big.NewInt(0)
	err := s.Call(option, &balance, "getSponsoredBalanceForGas", contractAddr.MustGetCommonAddress())
	if err != nil {
		return nil, err
	}
	return balance, nil
}

// GetSponsoredGasFeeUpperBound gets current Sponsored Gas fee upper bound
func (s *Sponsor) GetSponsoredGasFeeUpperBound(option *types.ContractMethodCallOption, contractAddr types.Address) (*big.Int, error) {
	gasFee := big.NewInt(0)
	err := s.Call(option, &gasFee, "getSponsoredGasFeeUpperBound", contractAddr.MustGetCommonAddress())
	if err != nil {
		return nil, err
	}
	return gasFee, nil
}

// GetSponsorForCollateral gets collateral sponsor address
func (s *Sponsor) GetSponsorForCollateral(option *types.ContractMethodCallOption, contractAddr types.Address) (address types.Address, err error) {
	var sponsor *common.Address = &common.Address{}
	err = s.Call(option, sponsor, "getSponsorForCollateral", contractAddr.MustGetCommonAddress())
	if err != nil {
		return
	}
	address = cfxaddress.MustNewFromCommon(*sponsor)
	err = address.CompleteByClient(s.Client)
	return
}

// GetSponsoredBalanceForCollateral gets current Sponsored Balance for collateral
func (s *Sponsor) GetSponsoredBalanceForCollateral(option *types.ContractMethodCallOption, contractAddr types.Address) (*big.Int, error) {
	sponsoredBalance := big.NewInt(0)
	err := s.Call(option, &sponsoredBalance, "getSponsoredBalanceForCollateral", contractAddr.MustGetCommonAddress())
	if err != nil {
		return nil, err
	}
	return sponsoredBalance, nil
}

// IsWhitelisted checks if a user is in a contract's whitelist
func (s *Sponsor) IsWhitelisted(option *types.ContractMethodCallOption, contractAddr types.Address, userAddr types.Address) (bool, error) {
	sponsoredBalance := false
	err := s.Call(option, &sponsoredBalance, "isWhitelisted", contractAddr.MustGetCommonAddress(), userAddr.MustGetCommonAddress())
	if err != nil {
		return false, err
	}
	return sponsoredBalance, nil
}

// IsAllWhitelisted checks if all users are in a contract's whitelist
func (s *Sponsor) IsAllWhitelisted(option *types.ContractMethodCallOption, contractAddr types.Address) (bool, error) {
	result := false
	err := s.Call(option, &result, "isAllWhitelisted", contractAddr.MustGetCommonAddress())
	if err != nil {
		return false, err
	}
	return result, nil
}

// AddPrivilegeByAdmin for admin adds user to whitelist
func (s *Sponsor) AddPrivilegeByAdmin(option *types.ContractMethodSendOption, contractAddr types.Address, userAddresses []types.Address) (types.Hash, error) {
	userAddressesComm := make([]common.Address, 0)
	for _, v := range userAddresses {
		userAddressesComm = append(userAddressesComm, v.MustGetCommonAddress())
	}
	return s.SendTransaction(option, "addPrivilegeByAdmin", contractAddr.MustGetCommonAddress(), userAddressesComm)
}

// RemovePrivilegeByAdmin for admin removes user from whitelist
func (s *Sponsor) RemovePrivilegeByAdmin(option *types.ContractMethodSendOption, contractAddr types.Address, userAddresses []types.Address) (types.Hash, error) {
	userAddressesComm := make([]common.Address, 0)
	for _, v := range userAddresses {
		userAddressesComm = append(userAddressesComm, v.MustGetCommonAddress())
	}
	return s.SendTransaction(option, "removePrivilegeByAdmin", contractAddr.MustGetCommonAddress(), userAddressesComm)
}

// SetSponsorForGas for someone sponsor the gas cost for contract `contractAddr` with an `upper_bound` for a single transaction, it is payable
func (s *Sponsor) SetSponsorForGas(option *types.ContractMethodSendOption, contractAddr types.Address, upperBound *big.Int) (types.Hash, error) {
	return s.SendTransaction(option, "setSponsorForGas", contractAddr.MustGetCommonAddress(), upperBound)
}

// SetSponsorForCollateral for someone sponsor the storage collateral for contract `contractAddr`, it is payable
func (s *Sponsor) SetSponsorForCollateral(option *types.ContractMethodSendOption, contractAddr types.Address) (types.Hash, error) {
	return s.SendTransaction(option, "setSponsorForCollateral", contractAddr.MustGetCommonAddress())
}

func getSponsorAbi() string {
	return `[
        {
            "inputs": [
                {
                    "internalType": "address[]",
                    "name": "",
                    "type": "address[]"
                }
            ],
            "name": "addPrivilege",
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
                },
                {
                    "internalType": "address[]",
                    "name": "addresses",
                    "type": "address[]"
                }
            ],
            "name": "addPrivilegeByAdmin",
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
            "name": "getSponsorForCollateral",
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
                }
            ],
            "name": "getSponsorForGas",
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
                }
            ],
            "name": "getSponsoredBalanceForCollateral",
            "outputs": [
                {
                    "internalType": "uint256",
                    "name": "",
                    "type": "uint256"
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
                }
            ],
            "name": "getSponsoredBalanceForGas",
            "outputs": [
                {
                    "internalType": "uint256",
                    "name": "",
                    "type": "uint256"
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
                }
            ],
            "name": "getSponsoredGasFeeUpperBound",
            "outputs": [
                {
                    "internalType": "uint256",
                    "name": "",
                    "type": "uint256"
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
                }
            ],
            "name": "isAllWhitelisted",
            "outputs": [
                {
                    "internalType": "bool",
                    "name": "",
                    "type": "bool"
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
                    "name": "user",
                    "type": "address"
                }
            ],
            "name": "isWhitelisted",
            "outputs": [
                {
                    "internalType": "bool",
                    "name": "",
                    "type": "bool"
                }
            ],
            "stateMutability": "view",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "address[]",
                    "name": "",
                    "type": "address[]"
                }
            ],
            "name": "removePrivilege",
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
                },
                {
                    "internalType": "address[]",
                    "name": "addresses",
                    "type": "address[]"
                }
            ],
            "name": "removePrivilegeByAdmin",
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
            "name": "setSponsorForCollateral",
            "outputs": [],
            "stateMutability": "payable",
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
                    "internalType": "uint256",
                    "name": "upperBound",
                    "type": "uint256"
                }
            ],
            "name": "setSponsorForGas",
            "outputs": [],
            "stateMutability": "payable",
            "type": "function"
        }
    ]
	`
}

func getSponsorHexAddress(client sdk.ClientOperator) (types.Address, error) {
	addr := cfxaddress.MustNewFromHex("0x0888000000000000000000000000000000000001")
	err := addr.CompleteByClient(client)
	return addr, err
}
