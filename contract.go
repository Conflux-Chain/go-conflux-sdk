// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package sdk

import (
	"errors"
	"fmt"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

// Contract ...
type Contract struct {
	ABI     abi.ABI
	Client  *Client
	Address types.Address
}

// GetData ...
func (c *Contract) GetData(method string, args ...interface{}) (*[]byte, error) {
	packed, err := c.ABI.Pack(method, args)
	if err != nil {
		msg := fmt.Sprintf("encode method %+v with args %+v error", method, args)
		return nil, types.WrapError(err, msg)
	}

	return &packed, nil
}

// Call ...
func (c *Contract) Call(callRequest types.CallRequest, method string, args ...interface{}) (interface{}, error) {
	return nil, errors.New("not implement")

	// data, err := c.GetData(method, args)
	// if err != nil {
	// 	msg := fmt.Sprintf("get data of method %+v with args %+v error", method, args)
	// 	return nil, types.WrapError(err, msg)
	// }
	// c.Client.Call()
	// return data, nil
}

// SendTransaction ...
func (c *Contract) SendTransaction(callRequest types.CallRequest, method string, args ...interface{}) (*types.Hash, error) {
	return nil, errors.New("not implement")

	// data, err := c.GetData(method, args)
	// if err != nil {
	// 	return nil, err
	// }
}
