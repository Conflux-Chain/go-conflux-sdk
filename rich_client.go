// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	richtypes "github.com/Conflux-Chain/go-conflux-sdk/types/richclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// RichClient contains client, cfx-scan-backend service and contract-manager service
//
// cfx-scan-backend service and contract-manager service are centralized service for best query performance
type RichClient struct {
	CfxScanBackend  *ScanServer
	ContractManager *ScanServer
	Client          *Client
}

// ScanServer represents a centralized server
type ScanServer struct {
	Scheme        string
	HostName      string
	HTTPRequester HTTPRequester
}

const (
	accountBalancesPath    = "/account/token/list"
	accountTokenTxListPath = "/future/transfer/list"
	txListPath             = "/future/transaction/list"
	contractQueryPath      = "/contract/query"
)

// NewRichClient create new rich client
func NewRichClient(client *Client) *RichClient {

	cfxScanBackend := &ScanServer{
		Scheme:        "http",
		HostName:      "101.201.103.131:8885", //"testnet-jsonrpc.conflux-chain.org:18084",
		HTTPRequester: &http.Client{},
	}

	contractManager := &ScanServer{
		Scheme:        "http",
		HostName:      "101.201.103.131:8886/api", //"13.75.69.106:8886",
		HTTPRequester: &http.Client{},
	}

	richClient := RichClient{
		cfxScanBackend,
		contractManager,
		client,
	}

	return &richClient
}

// URL returns url build by schema, host, path and params
func (s *ScanServer) URL(path string, params map[string]interface{}) string {
	q := url.Values{}
	for key, val := range params {
		q.Add(key, fmt.Sprintf("%+v", val))
	}
	encodedParams := q.Encode()
	result := fmt.Sprintf("%+v://%+v%+v?%+v", s.Scheme, s.HostName, path, encodedParams)
	return result
}

// Get sends a "Get" request and fill the unmarshaled value of field "Result" in response to unmarshaledResult
func (s *ScanServer) Get(path string, params map[string]interface{}, unmarshaledResult interface{}) error {
	client := s.HTTPRequester
	fmt.Println("request url:", s.URL(path, params))
	rspBytes, err := client.Get(s.URL(path, params))
	if err != nil {
		return err
	}

	defer func() {
		err := rspBytes.Body.Close()
		if err != nil {
			fmt.Println("close rsp error", err)
		}
	}()

	body, err := ioutil.ReadAll(rspBytes.Body)
	if err != nil {
		return err
	}
	// fmt.Printf("body:%+v\n\n", string(body))

	var rsp richtypes.Response
	err = json.Unmarshal(body, &rsp)
	if err != nil {
		return err
	}
	// fmt.Printf("unmarshaled body: %+v\n\n", rsp)

	if rsp.Code != 0 {
		msg := fmt.Sprintf("code:%+v, message:%+v", rsp.Code, rsp.Message)
		return errors.New(msg)
	}

	rstBytes, err := json.Marshal(rsp.Result)
	if err != nil {
		return err
	}
	// fmt.Printf("marshaled result: %+v\n\n", string(rstBytes))

	err = json.Unmarshal(rstBytes, unmarshaledResult)
	if err != nil {
		return err
	}
	// fmt.Printf("unmarshaled result: %+v\n\n", unmarshaledResult)
	return nil
}

// GetAccountTokenTransfers returns address releated transactions
func (rc *RichClient) GetAccountTokenTransfers(address types.Address, tokenIdentifier *types.Address, pageNumber, pageSize uint) (*richtypes.TokenTransferEventList, error) {
	params := make(map[string]interface{})
	params["address"] = address
	params["page"] = pageNumber
	params["pageSize"] = pageSize
	params["txType"] = "all"

	var tteList *richtypes.TokenTransferEventList
	// when tokenIdentifier is not nil return transfer events of the token
	if tokenIdentifier != nil {
		var tts richtypes.TokenTransferEventList
		params["contractAddress"] = *tokenIdentifier
		err := rc.CfxScanBackend.Get(accountTokenTxListPath, params, &tts)
		if err != nil {
			msg := fmt.Sprintf("get result of CfxScanBackend server and path {%+v}, params: {%+v} error", accountTokenTxListPath, params)
			return nil, types.WrapError(err, msg)
		}
		// return &tts, nil
		tteList = &tts
	} else {
		// when tokenIdentifier is nil return transaction of main coin
		var txs richtypes.TransactionList
		err := rc.CfxScanBackend.Get(txListPath, params, &txs)
		if err != nil {
			msg := fmt.Sprintf("get result of CfxScanBackend server and path {%+v}, params: {%+v} error", txListPath, params)
			return nil, types.WrapError(err, msg)
		}
		fmt.Printf("txs length: %v\n\n", len(txs.List))
		tteList = txs.ToTokenTransferEventList()
	}

	fmt.Printf("ttelist length: %v\n\n", len(tteList.List))
	// get epoch number and revert rate of every transaction
	var wg sync.WaitGroup
	wg.Add(int(len(tteList.List)))

	errorStrs := []string{}

	for i := range tteList.List {
		go func(_tte *richtypes.TokenTransferEvent) {
			defer wg.Done()

			tx, err := rc.Client.GetTransactionByHash(_tte.TransactionHash)
			if err != nil {
				errMsg := fmt.Sprintf("get transaction by hash %+v error: %+v", _tte.TransactionHash, err.Error())
				errorStrs = append(errorStrs, errMsg)
				return
			}

			block, err := rc.Client.GetBlockByHash(*tx.BlockHash)
			if err != nil {
				errMsg := fmt.Sprintf("get block by hash %+v error: %+v", tx.BlockHash, err.Error())
				errorStrs = append(errorStrs, errMsg)
				return
			}

			_tte.BlockHash = block.Hash
			_tte.RevertRate = 0
		}(&tteList.List[i])
	}
	wg.Wait()

	if len(errorStrs) > 0 {
		joinedErr := strings.Join(errorStrs, "\n")
		return nil, errors.New(joinedErr)
	}

	return tteList, nil
}

// CreateSendTokenTransaction create unsigned transaction according to input params
func (rc *RichClient) CreateSendTokenTransaction(from types.Address, to types.Address, amount *hexutil.Big, tokenIdentifier *types.Address) (*types.UnsignedTransaction, error) {
	if tokenIdentifier == nil {
		tx, err := rc.Client.CreateUnsignedTransaction(from, to, amount, nil)
		if err != nil {
			msg := fmt.Sprintf("Create Unsigned Transaction by from {%+v}, to {%+v}, amount {%+v} error", from, to, amount)
			return nil, types.WrapError(err, msg)
		}
		return tx, nil
	}

	params := make(map[string]interface{})
	params["address"] = tokenIdentifier
	params["fields"] = "abi,typeCode"

	var cInfo richtypes.Contract
	err := rc.ContractManager.Get(contractQueryPath, params, &cInfo)
	if err != nil {
		msg := fmt.Sprintf("get and unmarsal data from contract manager server with path {%+v}, paramas {%+v} error", contractQueryPath, params)
		return nil, types.WrapError(err, msg)
	}

	contract, err := rc.Client.GetContract(cInfo.ABI, &to)
	if err != nil {
		msg := fmt.Sprintf("get contract by ABI {%+v}, to {%+v} error", cInfo.ABI, to)
		return nil, types.WrapError(err, msg)
	}

	data, err := rc.getDataForTransToken(cInfo.GetContractType(), contract, to, amount)
	if err != nil {
		msg := fmt.Sprintf("get data for transfer token method error, contract type {%+v} ", cInfo.GetContractType())
		return nil, types.WrapError(err, msg)
	}

	tx, err := rc.Client.CreateUnsignedTransaction(from, to, nil, data)
	if err != nil {
		msg := fmt.Sprintf("create transaction with params {from: %+v, to: %+v, data: %+v} error ", from, to, data)
		return nil, types.WrapError(err, msg)
	}
	return tx, nil
}

func (rc *RichClient) getDataForTransToken(contractType richtypes.ContractType, contract Contractor, to types.Address, amount *hexutil.Big) (*[]byte, error) {
	var data *[]byte
	var err error

	// erc20 or fanscoin method signature are transfer(address,uint256)
	if contractType == richtypes.ERC20 || contractType == richtypes.FANSCOIN {
		data, err = contract.GetData("transfer", common.HexToAddress(string(to)), amount.ToInt())
		if err != nil {
			msg := fmt.Sprintf("get data of contract {%+v}, method {%+v}, params {to: %+v, amount: %+v} error ", contract, "transfer", to, amount)
			return nil, types.WrapError(err, msg)
		}
		return data, nil
	}

	// erc721 send by token_id
	//
	// if cInfo.ContractType == scantypes.ERC721 {
	// 	data, err = contract.GetData()
	// }

	// erc777 method signature is send(address,uint256,bytes)
	if contractType == richtypes.ERC777 {
		data, err = contract.GetData("send", common.HexToAddress(string(to)), amount.ToInt(), []byte{})
		if err != nil {
			msg := fmt.Sprintf("get data of contract {%+v}, method {%+v}, params {to: %+v, amount: %+v} error ", contract, "send", to, amount)
			return nil, types.WrapError(err, msg)
		}
		return data, nil
	}

	// if cInfo.ContractType == scantypes.DEX {
	// 	data, err = contract.GetData()
	// }
	msg := fmt.Sprintf("Do not support build data for transfer token function of contract type %+v", contractType)
	err = errors.New(msg)
	return nil, err
}

// GetTokenByIdentifier get token detail infomation by token identifier
func (rc *RichClient) GetTokenByIdentifier(tokenIdentifier types.Address) (*richtypes.Contract, error) {
	params := make(map[string]interface{})
	params["address"] = tokenIdentifier
	var contract richtypes.Contract
	err := rc.ContractManager.Get(contractQueryPath, params, &contract)
	if err != nil {
		msg := fmt.Sprintf("get and unmarshal result of ContractManager server and path {%+v}, params: {%+v} error", contractQueryPath, params)
		return nil, types.WrapError(err, msg)
	}
	return &contract, nil
}

// GetAccountTokens returns coin balance and all token balances of specific address
func (rc *RichClient) GetAccountTokens(account types.Address) (*richtypes.TokenWithBlanceList, error) {
	params := make(map[string]interface{})
	params["address"] = account

	var tbs richtypes.TokenWithBlanceList
	err := rc.ContractManager.Get(accountBalancesPath, params, &tbs)
	if err != nil {
		msg := fmt.Sprintf("get and unmarshal result of ContractManager server and path {%+v}, params: {%+v} error", accountBalancesPath, params)
		return nil, types.WrapError(err, msg)
	}
	return &tbs, nil
}
