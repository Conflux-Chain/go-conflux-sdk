package context

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"path"
	"time"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/contracts"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	address "github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// JSONFmt ...
func JSONFmt(v interface{}) string {
	j, e := json.Marshal(v)
	if e != nil {
		panic(e)
	}
	var str bytes.Buffer
	_ = json.Indent(&str, j, "", "    ")
	return str.String()
}

// WaitPacked ...
func WaitPacked(client sdk.RpcCaller, txhash types.Hash) *types.TransactionReceipt {
	fmt.Printf("wait for transaction %v be packed\n", txhash)
	var tx *types.TransactionReceipt
	for {
		time.Sleep(time.Duration(1) * time.Second)
		var err error
		tx, err = client.GetTransactionReceipt(txhash)
		PanicIfErrf(err, "failed to get transaction receipt of %v", txhash)
		fmt.Print(".")
		if tx != nil {
			fmt.Printf("transaction is packed:%+v\n\n", JSONFmt(tx))
			break
		}
	}
	return tx
}

// GetNextNonceAndIncrease ...
func GetNextNonceAndIncrease() *hexutil.Big {
	client := config.GetClient()

	if nextNonce == nil {
		var err error
		if nextNonce, err = client.GetNextNonce(*defaultAccount, nil); err != nil {
			panic(err)
		}
	}

	currentNonce := types.NewBigIntByRaw(nextNonce.ToInt())
	nextNonce = types.NewBigIntByRaw(big.NewInt(1).Add(nextNonce.ToInt(), big.NewInt(1)))
	return currentNonce
}

// CreateSignedTx ...
func CreateSignedTx(client sdk.SignableRpcCaller) []byte {

	to := cfxaddress.MustNewFromHex("0x10f4bcf113e0b896d9b34294fd3da86b4adf0302")
	to.CompleteByClient(client)
	unSignedTx := types.UnsignedTransaction{
		UnsignedTransactionBase: types.UnsignedTransactionBase{
			From:  defaultAccount,
			Value: types.NewBigInt(100),
			Nonce: GetNextNonceAndIncrease(),
		},
		To: &to,
	}
	err := client.PopulateTransaction(&unSignedTx)
	if err != nil {
		panic(err)
	}

	signedTx, err := client.GetWallet().SignTransactionWithPassphraseAndEcode(unSignedTx, "hello")
	if err != nil {
		panic(err)
	}
	fmt.Printf("signed tx %v result:\n0x%x\n\n", JSONFmt(unSignedTx), signedTx)
	return signedTx
}

// DeployNewErc20 ...
func DeployNewErc20() sdk.Contractor {
	abiFilePath := path.Join(currentDir, "contract/erc20.abi")
	bytecodeFilePath := path.Join(currentDir, "contract/erc20.bytecode")
	contract, _ := DeployContractWithConstroctor(abiFilePath, bytecodeFilePath, big.NewInt(100000), "biu", uint8(10), "BIU")
	return contract
}

// DeployIfNotExist ...
func DeployIfNotExist(contractAddress types.Address, abiFilePath string, bytecodeFilePath string, force bool) (sdk.Contractor, *types.Hash) {
	client := config.GetClient()

	isContract := contractAddress.GetAddressType() == address.AddressTypeContract
	isCodeExist := false

	if isContract {
		code, err := client.GetCode(contractAddress)
		// fmt.Printf("err: %v,code:%v\n", err, len(code))
		if err == nil && len(code) > 0 {
			isCodeExist = true
		}
	}

	fmt.Printf("%v isAddress:%v, isCodeExist:%v\n", contractAddress, isContract, isCodeExist)
	if !force && isContract && isCodeExist {
		abi, err := ioutil.ReadFile(abiFilePath)
		if err != nil {
			panic(err)
		}
		contract, err := contracts.NewContract(client, abi, &contractAddress)
		if err != nil {
			panic(err)
		}
		return contract, nil
	}

	contract, txhash := DeployContractWithConstroctor(abiFilePath, bytecodeFilePath, big.NewInt(100000), "biu", uint8(10), "BIU")
	return contract, txhash
}

// DeployContractWithConstroctor ...
func DeployContractWithConstroctor(abiFile string, bytecodeFile string, params ...interface{}) (sdk.Contractor, *types.Hash) {

	fmt.Println("start deploy contract with construcotr")
	abi, err := ioutil.ReadFile(abiFile)
	if err != nil {
		panic(err)
	}

	bytecodeHexStr, err := ioutil.ReadFile(bytecodeFile)
	if err != nil {
		panic(err)
	}

	bytecode, err := hex.DecodeString(string(bytecodeHexStr))
	if err != nil {
		panic(err)
	}

	option := types.ContractDeployOption{}
	option.Nonce = GetNextNonceAndIncrease()
	result := contracts.DeployContract(config.GetClient(), &option, abi, bytecode, params...)

	_ = <-result.DoneChannel
	if result.Error != nil {
		panic(result.Error)
	}
	contract := result.DeployedContract
	fmt.Printf("deploy contract with abi: %v, bytecode: %v done\ncontract address: %+v\ntxhash:%v\n\n", abiFile, bytecodeFile, contract.Address(), result.TransactionHash)

	return contract, result.TransactionHash
}

// PanicIfErrf ...
func PanicIfErrf(err error, msg string, values ...interface{}) {
	if err != nil {
		fmt.Printf(msg, values...)
		fmt.Println()
		fmt.Printf("err stack:%+v", err)
		fmt.Println()
		panic(err)
	}
}
