Conflux Address
---------------------

Conflux Address is base32 encoding defined in [CIP-37]((https://github.com/Conflux-Chain/CIPs/blob/master/CIPs/cip-37.md)), see details from 
 [Conflux Addresses Introduction](https://developer.confluxnetwork.org/introduction/en/conflux_basics#address)

## Usage

### Create an Address
There are four ways to create an address instance. In each way a MustXxx method is provided to realize quick creation. If an error occurs when using this method, panic will take place. You can also check whether an address is valid through the returns of the following methods:

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

`Note:` The networkId configuration is optional. If it is not set manually, it will be set automatically when sending RPC requests in client.