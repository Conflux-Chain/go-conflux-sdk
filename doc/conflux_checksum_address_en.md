Conflux base32Check addresses
===
As a new public chain, Conflux realizes high performance as well as compatibility with Ethereum. Conflux adopts address format compatible with Ethereum addresses, and thus is compatible with Ethereum Virtual Machine (EVM).
The advantage of the compatibility between Conflux and Ethereum is obvious:  it reduces the cost and difficulty of cross-chain migration. But there are also some problems. Since the addresses on Conflux and Ethereum are similar, users may loss their assets when performing cross-chain transactions using ShuttleFlow if they transfer to a mistake address, which is a serious problem. To improve user experience and reduce address mistakes when users use cross-chain functions, Conflux introduces a new address format: base32Check in [CIP37](https://github.com/Conflux-Chain/CIPs/blob/master/CIPs/cip-37.md).


### Before CIP37 
At first, Conflux adopts the address format similar with Ethereum, which is a hex40 address (hex code with a length of 40 bits). The difference is that Conflux differentiate the addresses with different starts: 0x1 for ordinary individual addresses, 0x8 for smart contracts and 0x0 for in-built contracts.

Only hex40 addresses with these three starts are available on Conflux. Some Ethereum addresses (with a 0x1 start) can be used as Conflux addresses, while a Conflux address has a 1/16 chance of being used as an Ethereum address.

Currently, there are three kinds of addresses:

* Ordinary addresses: `0x1`386b4185a223ef49592233b69291bbe5a80c527
* Smart contract addresses: `0x8`269f0add11b4915d78791470d091d25cff73ee5
* In-built contract addresses: `0x0`888000000000000000000000000000000000002

Because the addresses are not completely compatible on Conflux and Ethereum, users will loss assets when they use a wrong address. Ethereum has introduced a regulation with a checksum in [EIP55](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-55.md) to change the characters meeting the requirement into the upper case in order to prevent transferring to wrong addresses. Conflux also introduces regulations to change checksums.

* Non-chechsum address: 0x1386`b`4185`a`223`ef`49592233b69291bbe5a80`c`527
* Chechsum address: 0x1386`B`4185`A`223`EF`49592233b69291bbe5a80`C`527

### CIP37 Address
In order to solve the problems of mistakenly using wrong addresses, we introduces a brand new base32 checksum address format in [CIP37](https://github.com/Conflux-Chain/CIPs/blob/master/CIPs/cip-37.md). Besides checksum, the new addresses also include information such as network, type.

Old address vs new address:

* hex40 address: `0x1`386b4185a223ef49592233b69291bbe5a80c527
* base32 address: cfx:aak2rra2njvd77ezwjvx04kkds9fzagfe6ku8scz91

The new addresses use customized base32 code address. Currently applied characters are: `abcdefghjkmnprstuvwxyz0123456789` (i, l, o, q removed).

In new format addresses, network types are included. Up to now there are three types: cfx，cfxtest，net[n]

* cfx:aak2rra2njvd77ezwjvx04kkds9fzagfe6ku8scz91
* cfxtest:aak2rra2njvd77ezwjvx04kkds9fzagfe6d5r8e957
* net1921:aak2rra2njvd77ezwjvx04kkds9fzagfe65k87kwdf

Meanwhile, new addresses also include address type information, currently four types (types are usually in upper case):

* user: CFX:TYPE.USER:AAK2RRA2NJVD77EZWJVX04KKDS9FZAGFE6KU8SCZ91
* contract: CFX:TYPE.CONTRACT:ACB2RRA2NJVD77EZWJVX04KKDS9FZAGFE640XW9UAE
* builtin: CFX:TYPE.BUILTIN:AAEJUAAAAAAAAAAAAAAAAAAAAAAAAAAAAJRWUC9JNB
* null: CFX:TYPE.NULL:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA0SFBNJM2

The two address formats (hex40 and base32) are convertible to each other. They are the same if converted to byte arrays. However, when converting hex40 addresses (starting with 0x) into base32check addresses, the network ID information is also required.

### Conflux fullnode RPC
From v1.1.1, Conflux-rust will apply the new address format. If returns include address information, it will be in the new format.

If you use hex40 addresses to call RPC, it will return with an error:
```js
{
    "code": -32602,
    "message": "Invalid params: Invalid base32 address: zero or multiple prefixes."
}
```

If you use a wrong network type (eg. use a testnet address for the mainnet PRC), it will return with an error:
```js
{
    "code": -32602,
    "message": "Invalid parameters: address",
    "data": "\"network prefix unexpected: ours cfx, got cfxtest\""
}
```

### go-conflux-sdk

#### Address
##### Create an Address
There are four ways to create an address instance. In each way a MustXxx method is provided to realize quick creation. If error occurs when using this method, panic will take place. You can also check whether an address is valid through the returns of the following methods:

- Create via base32 string
```go
func NewFromBase32(base32Str string) (cfxAddress Address, err error)
func MustNewFromBase32(base32Str string) (address Address)
```
- Create via hex40 string
```go
func NewFromHex(hexAddressStr string, networkID ...uint32) (val Address, err error)
func MustNewFromHex(hexAddressStr string, networkID ...uint32) (val Address)
```
- Create via common.Address
```go
func NewFromCommon(commonAddress common.Address, networkID ...uint32) (val Address, err error)
func MustNewFromCommon(commonAddress common.Address, networkID ...uint32) (address Address)
```
- Create via byte array
```go
func NewFromBytes(hexAddress []byte, networkID ...uint32) (val Address, err error)
func MustNewFromBytes(hexAddress []byte, networkID ...uint32) (address Address)
```

`Note:` The networkId configuration is optional. If it is not set manually, it will be set automatically when sending RPC request in client.

##### Address instance method
```go
// ToHex returns hex address string and networkID
func (a *Address) ToHex() (hexAddressStr string, networkID uint32)

// ToCommon returns common.Address and networkID
func (a *Address) ToCommon() (address common.Address, networkID uint32, err error) 

// MustGetBase32Address returns base32 string of address which doesn't include address type
func (a *Address) MustGetBase32Address() string

// MustGetVerboseBase32Address returns base32 string of address with address type
func (a *Address) MustGetVerboseBase32Address() string

// MustGetHexAddress returns hex format address and panic if error
func (a *Address) GetHexAddress() string

// MustGetCommonAddress returns common address and panic if error
func (a *Address) MustGetCommonAddress() common.Address

// MustGetNetworkID returns networkID and panic if error
func (a *Address) GetNetworkID() uint32

// GetNetworkType returns network type
func (a *Address) GetNetworkType() NetworkType 

// GetAddressType returuns address type
func (a *Address) GetAddressType() AddressType 

// GetBody returns body
func (a *Address) GetBody() Body

// GetChecksum returns checksum
func (a *Address) GetChecksum() Checksum

// CompleteByClient will set networkID by client.GetNetworkID() if a.networkID not be 0
func (a *Address) CompleteByClient(client NetworkIDGetter) error

// CompleteByNetworkID will set networkID if current networkID isn't 0
func (a *Address) CompleteByNetworkID(networkID uint32) error

// IsValid return true if address is valid
func (a *Address) IsValid() bool

// MarshalText implements the encoding.TextMarshaler interface.
func (a Address) MarshalText() ([]byte, error)

// UnmarshalJSON implements the json.Unmarshaler interface.
func (a *Address) UnmarshalJSON(data []byte) error
```

#### NetworkID
In the Client class, new method `GetNetworkID`is added to get NetowrkID
```
func (client *Client) GetNetworkID() (uint32, error)
```

#### AccountManager
AccountManager instantiation requires to send networkId
```
func NewAccountManager(keydir string, networkID uint32) *AccountManager
```

#### Contract
Solidity is still using hex40 addresses, so when contract objects are interacting with contracts, common.Address is still sueful.

##### Event Decode
Event decoding will continue to use common.Address

##### Internal contract
The address-related parameters in AdminControl, Sponsor and Staking of in-built contracts will be in base32check format. SDK will help to convert the format.