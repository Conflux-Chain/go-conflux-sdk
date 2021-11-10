// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package client

import (
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/interfaces"
)

// Client represents a client to interact with Conflux blockchain.
type Client struct {
	*ClientCore
	*ClientHelper

	*RpcCfxClient
	debug  *RpcDebugClient
	pubsub *RpcPubsubClient
	trace  *RpcTraceClient
	//TODO
	//  pos    *RpcPosClient
}

// NewClient creates an instance of Client with specified conflux node url, it will create account manager if option.KeystorePath not empty.
func NewClient(nodeURL string) (Client, error) {
	core, err := NewClientCore(nodeURL)
	if err != nil {
		return Client{}, err
	}
	client := Client{}
	client.ClientCore = core
	client.RpcCfxClient = &RpcCfxClient{core}
	client.debug = &RpcDebugClient{core}
	client.pubsub = &RpcPubsubClient{core}
	client.trace = &RpcTraceClient{core}
	// TODO
	// client.pos = &RpcPosClient{&core}

	client.ClientHelper = &ClientHelper{&client}
	return client, nil
}

func MustNewClient(nodeURL string) *Client {
	client, err := NewClient(nodeURL)
	if err != nil {
		panic(err)
	}
	return &client
}

func (c *Client) SetRetry(retryCount int, retryInterval time.Duration) *Client {
	c.ClientCore.SetRetry(retryCount, retryInterval)
	return c
}

func (c *Client) SetRequestTimeout(timeout time.Duration) *Client {
	c.ClientCore.SetRequestTimeout(timeout)
	return c
}

func (c *Client) Cfx() interfaces.RpcCfxCaller {
	return c.RpcCfxClient
}
func (c *Client) Debug() interfaces.RpcDebugCaller {
	return c.debug
}

// TODO
func (c *Client) Pos() interfaces.RpcPosCaller {
	// return c.pos
	return nil
}
func (c *Client) Pubsub() interfaces.RpcPubsubCaller {
	return c.pubsub
}
func (c *Client) Trace() interfaces.RpcTraceCaller {
	return c.trace
}

// GetNetworkID returns networkID of connecting conflux node
func (c *Client) GetNetworkID() (uint32, error) {
	return c.ClientCore.getNetworkId()
}
