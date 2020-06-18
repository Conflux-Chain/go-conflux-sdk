---
id: go_sdk_zh_cn
title: Golang SDK
custom_edit_url: https://github.com/Conflux-Chain/go-conflux-sdk/edit/master/api_zh_cn.md
keywords:
  - conflux
  - go
  - sdk
---
# API Reference
## 简介
go conflux sdk模块是一个包的集合，其中包含conflux生态系统的具体功能。

- SDK包用于与conflux链、账户管理器和运行智能合约进行交互
- utils包含有对Dapp开发人员有用的帮助功能。

## 安装
您可以直接获取Conflux Golang API或使用go模块，如下所示
```
go get github.com/Conflux-Chain/go-conflux-sdk
```
您还可以将Conflux Golang API添加到vendor文件夹中。
```
govendor fetch github.com/Conflux-Chain/go-conflux-sdk
```

之后，您需要创建一个带有节点 url 地址和账户管理器实例的客户端实例。
```go
url:= "http://testnet-jsonrpc.conflux-chain.org:12537"
client, err := sdk.NewClient(url)
if err != nil {
	fmt.Println("new client error:", err)
	return
}
am := sdk.NewAccountManager("./keystore")
client.SetAccountManager(am)
```
## sdk包
```
import "github.com/Conflux-Chain/go-conflux-sdk"
```


### type AccountManager 函数

```go
type AccountManager struct {
}
```

账户管理器AccountManager管理conflux的账户

#### func NewAccountManager函数

```go
func NewAccountManager(keydir string) *AccountManager
```
新账户管理器（NewAccountManager）基于密钥库目录“keydir”创建账户管理器（AccountManager）的实例。

#### func (*AccountManager) Create 函数

```go
func (m *AccountManager) Create(passphrase string) (types.Address, error)
```
关键字Create创建一个新帐户并将密钥（keystore）文件放入密钥（keystore）目录

#### func (*AccountManager) Delete函数

```go
func (m *AccountManager) Delete(address types.Address, passphrase string) error
```
关键字Delete删除指定的帐户并从密钥（keystore）目录中删除密钥（keystore）文件。

#### func (*AccountManager) GetDefault函数

```go
func (m *AccountManager) GetDefault() (*types.Address, error)
```
关键字GetDefault返回密钥（keystore）目录中的第一个帐户

#### func (*AccountManager) Import函数

```go
func (m *AccountManager) Import(keyFile, passphrase, newPassphrase string) (types.Address, error)
```
关键字Import将帐户从外部密钥文件导入密钥库目录。如果帐户已存在，则返回错误error提示。

#### func (*AccountManager) List函数

```go
func (m *AccountManager) List() []types.Address
```
List关键字列出keystore目录中的所有帐户。

#### func (*AccountManager) Lock函数

```go
func (m *AccountManager) Lock(address types.Address) error
```
Lock关键字锁定指定的帐户。

#### func (*AccountManager) Sign函数

```go
func (m *AccountManager) Sign(tx types.UnsignedTransaction, passphrase string) (v byte, r, s []byte, err error)
```
Sign关键字通过密码对tx进行签名并返回签名

#### func (*AccountManager) SignAndEcodeTransactionWithPassphrase函数

```go
func (m *AccountManager) SignAndEcodeTransactionWithPassphrase(tx types.UnsignedTransaction, passphrase string) ([]byte, error)
```
关键字SignAndEcodeTransactionWithPassphrase使用给定的密码对tx进行签名并返回其RLP编码的数据。


#### func (*AccountManager) SignTransaction函数

```go
func (m *AccountManager) SignTransaction(tx types.UnsignedTransaction) ([]byte, error)
```
关键字SignTransaction对tx进行签名并返回其RLP编码的数据。

#### func (*AccountManager) SignTransactionWithPassphrase函数

```go
func (m *AccountManager) SignTransactionWithPassphrase(tx types.UnsignedTransaction, passphrase string) (*types.SignedTransaction, error)
```
关键字SignTransactionWithPassphrase使用给定的密码对tx进行签名，并返回带有签名的交易

#### func (*AccountManager) TimedUnlock

```go
func (m *AccountManager) TimedUnlock(address types.Address, passphrase string, timeout time.Duration) error
```
关键字TimedUnlock解锁一段时间内的指定帐户。

#### func (*AccountManager) TimedUnlockDefault函数

```go
func (m *AccountManager) TimedUnlockDefault(passphrase string, timeout time.Duration) error
```
关键字TimedUnlockDefault解锁一段时间内的默认帐户。

#### func (*AccountManager) Unlock函数

```go
func (m *AccountManager) Unlock(address types.Address, passphrase string) error
```
关键字unlock无限期地解锁指定的帐户。

#### func (*AccountManager) UnlockDefault函数

```go
func (m *AccountManager) UnlockDefault(passphrase string) error
```
关键字UnlockDefault无限期地解锁默认帐户。

#### func (*AccountManager) Update函数

```go
func (m *AccountManager) Update(address types.Address, passphrase, newPassphrase string) error
```
关键字Update更新指定帐户的密码短语。

### type Client函数

```go
type Client struct {
}
```

Client表示要与Conflux区块链交互的客户端。

#### func NewClient函数

```go
func NewClient(nodeURL string) (*Client, error)
```
new Client使用指定的conflux节点url地址创建一个新的Client实例。

#### func NewClientWithRPCRequester函数

```go
func NewClientWithRPCRequester(rpcRequester rpcRequester) (*Client, error)
```
关键字NewClientWithRPCRequester使用指定的rpcRequester创建客户端

#### func NewClientWithRetry函数

```go
func NewClientWithRetry(nodeURL string, retryCount int, retryInterval time.Duration) (*Client, error)
```
关键字NewClientWithRetry使用指定的conflux节点url地址和重试选项创建可重试的新客户端实例。
如果通过0，则重试间隔将设置为1秒

#### func (*Client) ApplyUnsignedTransactionDefault函数

```go
func (client *Client) ApplyUnsignedTransactionDefault(tx *types.UnsignedTransaction) error
```
ApplyUnsignedTransactionDefault将从conflux节点获取的值设置为空字节段

#### func (*Client) BatchCallRPC函数

```go
func (client *Client) BatchCallRPC(b []rpc.BatchElem) error
```
"BatchCallRPC将所有给定的请求作为单个批发送，并等待服务器返回所有请求的响应。



与Call相反，BatchCall只返回I/O错误。任何请求的特定错误都会通过相应BatchElem的error字段报告。



请注意，批处理调用不能在服务器端以原子方式执行。"

#### func (*Client) BatchGetBlockConfirmationRisk函数

```go
func (client *Client) BatchGetBlockConfirmationRisk(blockhashes []types.Hash) (map[types.Hash]*big.Float, error)
```
关键字BatchGetBlockConfirmationRisk通过区块哈希批量获取确认风险信息

#### func (*Client) BatchGetBlockSummarys函数

```go
func (client *Client) BatchGetBlockSummarys(blockhashes []types.Hash) (map[types.Hash]*types.BlockSummary, error)
```
关键字BatchGetBlockSummarys按区块哈希批量请求区块摘要信息

#### func (*Client) BatchGetRawBlockConfirmationRisk函数

```go
func (client *Client) BatchGetRawBlockConfirmationRisk(blockhashes []types.Hash) (map[types.Hash]*big.Int, error)
```
BatchGetRawBlockConfirmationRisk按区块哈希批量请求原始确认风险信息

#### func (*Client) BatchGetTxByHashes函数

```go
func (client *Client) BatchGetTxByHashes(txhashes []types.Hash) (map[types.Hash]*types.Transaction, error)
```
BatchGetTxByHash通过交易哈希（TxHash）批量请求交易信息

#### func (*Client) Call函数

```go
func (client *Client) Call(request types.CallRequest, epoch *types.Epoch) (*string, error)
```
Call在特定的纪元（epoch）执行调用交易“请求”的消息，它直接在节点的VM中执行，但从未被挖掘到区块链中并返回合约执行结果。

#### func (*Client) CallRPC函数

```go
func (client *Client) CallRPC(result interface{}, method string, args ...interface{}) error
```
call RPC使用给定的参数执行JSON-RPC调用，如果没有发生错误，则将其解编为结果
结果必须是一个指针，以便json包可以解编到其中。您还可以传递空值（nil），在这种情况下，结果将被忽略。

#### func (*Client) Close函数

```go
func (client *Client) Close()
```
关键字Close关闭客户端，中止任何正在运行的请求。

#### func (*Client) CreateUnsignedTransaction函数

```go
func (client *Client) CreateUnsignedTransaction(from types.Address, to types.Address, amount *hexutil.Big, data []byte) (*types.UnsignedTransaction, error)
```
关键字CreateUnsignedTransaction通过参数创建未签名的交易，其他字段将被设置为从conflux节点获取的值。

#### func (*Client) Debug函数

```go
func (client *Client) Debug(method string, args ...interface{}) (interface{}, error)
```
关键字debug调用Conflux调试API。

#### func (*Client) DeployContract函数

```go
func (client *Client) DeployContract(option *types.ContractDeployOption, abiJSON []byte,
	bytecode []byte, constroctorParams ...interface{}) *ContractDeployResult
```
关键字DeployContract通过abiJSON、字节码和构造函数参数（constructor params）部署合约。
它返回一个ContractDeployState实例，该实例包含3个通道，用于在状态更改时通知的。

#### func (*Client) EstimateGasAndCollateral函数

```go
func (client *Client) EstimateGasAndCollateral(request types.CallRequest) (*types.Estimate, error)
```
关键字EstimateGasAndcollateral执行消息调用“请求”，并返回已经使用的gas数量和抵押品存量

#### func (*Client) GetBalance函数

```go
func (client *Client) GetBalance(address types.Address, epoch ...*types.Epoch) (*big.Int, error)
```
关键字GetBalance返回指定地址在纪元（epoch）的余额。

#### func (*Client) GetBestBlockHash函数

```go
func (client *Client) GetBestBlockHash() (types.Hash, error)
```
关键字GetBestBlockHash返回当前最佳区块哈希。

#### func (*Client) GetBlockByEpoch函数

```go
func (client *Client) GetBlockByEpoch(epoch *types.Epoch) (*types.Block, error)
```
关键GetBlockByEpoch返回指定纪元（epoch）的区块。如果纪元（epoch）无效，则返回具体错误。

#### func (*Client) GetBlockByHash函数

```go
func (client *Client) GetBlockByHash(blockHash types.Hash) (*types.Block, error)
```
关键字GetBlockByHash返回指定区块哈希值的区块。如果找不到区块，则返回空值nil。

#### func (*Client) GetBlockConfirmationRisk函数

```go
func (client *Client) GetBlockConfirmationRisk(blockHash types.Hash) (*big.Float, error)
```
关键字GetBlockConfirmationRisk表示所在纪元的轴心块成为正常块的概率。

（原始确认风险系数/（2^256-1））

#### func (*Client) GetBlockSummaryByEpoch函数

```go
func (client *Client) GetBlockSummaryByEpoch(epoch *types.Epoch) (*types.BlockSummary, error)
```
关键字GetBlockSummaryByEpoch返回指定纪元（epoch）的区块摘要。如果纪元（epoch）无效，则返回具体错误。

#### func (*Client) GetBlockSummaryByHash函数

```go
func (client *Client) GetBlockSummaryByHash(blockHash types.Hash) (*types.BlockSummary, error)
```
关键字GetBlockSummaryByHash返回指定区块哈希值的区块摘要。如果区块没有找到，则返回空值nil

#### func (*Client) GetBlocksByEpoch函数

```go
func (client *Client) GetBlocksByEpoch(epoch *types.Epoch) ([]types.Hash, error)
```
关键字GetBlocksByEpoch返回指定纪元（epoch）的区块哈希值。

#### func (*Client) GetCode函数

```go
func (client *Client) GetCode(address types.Address, epoch ...*types.Epoch) (string, error)
```
关键字GetCode为纪元（epoch）中指定的地址返回十六进制格式的字节码。

#### func (*Client) GetContract函数

```go
func (client *Client) GetContract(abiJSON []byte, deployedAt *types.Address) (*Contract, error)
```
关键字GetContract根据abi json和它的部署地址创建一个合约实例

#### func (*Client) GetEpochNumber函数

```go
func (client *Client) GetEpochNumber(epoch ...*types.Epoch) (*big.Int, error)
```
关键字GetEpochNumber返回最高或指定的纪元号。

#### func (*Client) GetGasPrice函数

```go
func (client *Client) GetGasPrice() (*big.Int, error)
```
关键字GetGasPrice返回最近的平均gas价格。

#### func (*Client) GetLogs函数

```go
func (client *Client) GetLogs(filter types.LogFilter) ([]types.Log, error)
```
关键字GetLogs返回与指定过滤器匹配的日志。

#### func (*Client) GetNextNonce

```go
func (client *Client) GetNextNonce(address types.Address, epoch *types.Epoch) (*big.Int, error)
```
关键字GetNextNonce返回地址的下一个交易随机数

#### func (*Client) GetNodeURL函数

```go
func (client *Client) GetNodeURL() string
```
关键字GetNodeURL返回节点url

#### func (*Client) GetRawBlockConfirmationRisk函数

```go
func (client *Client) GetRawBlockConfirmationRisk(blockhash types.Hash) (*big.Int, error)
```
关键字GetRawBlockConfirmationRisk表示块所在纪元（epoch）的轴心块变为正常块的风险系数。

#### func (*Client) GetStatus函数

```go
func (client *Client) GetStatus() (*types.Status, error)
```
关键字GetStatus返回连接conflux节点的链的ID

#### func (*Client) GetTransactionByHash函数

```go
func (client *Client) GetTransactionByHash(txHash types.Hash) (*types.Transaction, error)
```
关键字GetTransactionByHash返回指定交易哈希值的交易。如果找不到交易，则返回空值nil。

#### func (*Client) GetTransactionReceipt函数

```go
func (client *Client) GetTransactionReceipt(txHash types.Hash) (*types.TransactionReceipt, error)
```
关键字GetTransactionReceipt返回指定交易哈希值的收据。如果找不到收据，则返回空值nil。

#### func (*Client) SendRawTransaction函数

```go
func (client *Client) SendRawTransaction(rawData []byte) (types.Hash, error)
```
关键字SendRawTransaction发送已签名的交易并返回其哈希值。

#### func (*Client) SendTransaction函数

```go
func (client *Client) SendTransaction(tx *types.UnsignedTransaction) (types.Hash, error)
```
关键字SendTransaction将交易签名并发送到conflux节点并返回交易哈希值。

#### func (*Client) SetAccountManager函数

```go
func (client *Client) SetAccountManager(accountManager AccountManagerOperator)
```
关键字SetAccountManager为签名的交易设置帐户管理器

#### func (*Client) SignEncodedTransactionAndSend函数

```go
func (client *Client) SignEncodedTransactionAndSend(encodedTx []byte, v byte, r, s []byte) (*types.Transaction, error)
```
关键字SignEncodedTransactionAndSend通过“r，s，v”签名对RLP编码(译者注：Recursive Length Prefix，递归的长度前缀,一种编码规则，后同不翻)
的交易“encodedTx”
进行签名并将其发送到节点，并返回响应的交易。

### type Contract函数

```go
type Contract struct {
	ABI     abi.ABI
	Client  ClientOperator
	Address *types.Address
}
```

关键字Contract代表智能合约。您可以方便地通过客户端创建智能合约.
使用上文提到的GetContract或者Client.DeployContract来部署合约

#### func  NewContract函数

```go
func NewContract(abiJSON []byte, client ClientOperator, address *types.Address) (*Contract, error)
```
关键字NewContract通过abi创建智能合约并且部署地址

#### func (*Contract) Call函数

```go
func (contract *Contract) Call(option *types.ContractMethodCallOption, resultPtr interface{}, method string, args ...interface{}) error
```
关键字Call使用参数调用智能合约方法，并将执行的结果填充到“resultPtr”。
resultPtr应该是输出结构类型方法的指针
请参考 https://github.com/Conflux-Chain/go-conflux-sdk/blob/master/README.md获取solidity类型到go类型的映射

#### func (*Contract) DecodeEvent函数

```go
func (contract *Contract) DecodeEvent(out interface{}, event string, log types.LogEntry) error
```
关键字DecodeEvent将检索到的日志解压缩到提供的输出结构中。
请参考https://github.com/Conflux-Chain/go-Conflux-sdk/blob/master/README.md获取solidity类型到go类型的映射

#### func (*Contract) GetData函数

```go
func (contract *Contract) GetData(method string, args ...interface{}) ([]byte, error)
```
关键字GetData打包给定的方法名使之与合约的ABI一致。方法调用的数据将由method_id, args0, arg1, ...argN组成。
Method id由4个字节组成，参数都是32个字节。Method id是从字符串签名的方法的哈希值的前4个字节创建的。
（签名=baz（uint32，string32））

请参考https://github.com/Conflux-Chain/go-Conflux-sdk/blob/master/README.md获取solidity类型到go类型的映射


#### func (*Contract) SendTransaction函数

```go
func (contract *Contract) SendTransaction(option *types.ContractMethodSendOption, method string, args ...interface{}) (*types.Hash, error)
```
关键字SendTransaction使用参数（args）将交易发送到智能合约方法并返回其交易哈希值



请参考https://github.com/Conflux-Chain/go-Conflux-sdk/blob/master/README.md获取solidity类型到go类型的映射

### type ContractDeployResult函数

```go
type ContractDeployResult struct {
	//DoneChannel channel for notifying when contract deployed done
	DoneChannel      <-chan struct{}
	TransactionHash  *types.Hash
	Error            error
	DeployedContract *Contract
}
```

当部署智能合约时，关键字ContractDeployResult用于状态改变的通知
## package utils函数
```
import "github.com/Conflux-Chain/go-conflux-sdk/utils"
```


#### func CalcBlockConfirmationRisk函数

```go
func CalcBlockConfirmationRisk(rawConfirmationRisk *big.Int) *big.Float
```
关键字CalcBlockConfirmationRisk计算区块恢复率

#### func Keccak256函数

```go
func Keccak256(hexStr string) (string, error)
```
关键字Keccak256通过keccak256算法计算16进制字符串哈希值并且返回其哈希值

#### func PrivateKeyToPublicKey函数

```go
func PrivateKeyToPublicKey(privateKey string) string
```
关键字PrivateKeyToPublicKey根据私钥计算公钥

#### func PublicKeyToAddress函数

```go
func PublicKeyToAddress(publicKey string) types.Address
```
关键字PublicKeyToAddress根据公钥生成地址

conflux的账户地址以“0x1”开头

#### func ToCfxGeneralAddress函数

```go
func ToCfxGeneralAddress(address common.Address) types.Address
```
关键字ToCfxGeneralAddress将普通地址转换为conflux 格式常规地址，
其十六进制字符串以“0x1”开头