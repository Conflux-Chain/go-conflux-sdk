package cfxaddress

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// Address ...
type Address struct {
	NetworkType NetworkType
	AddressType AddressType
	Body        Body
	Checksum    Checksum
}

// String ...
func (a Address) String() string {
	return strings.ToUpper(fmt.Sprintf("%v:%v:%v%v", a.NetworkType, a.AddressType, a.Body, a.Checksum))
}

// ShortString ...
func (a Address) ShortString() string {
	return fmt.Sprintf("%v:%v%v", a.NetworkType, a.Body, a.Checksum)
}

// Equals ...
func (a *Address) Equals(target *Address) bool {
	return reflect.DeepEqual(a, target)
}

// NewAddressFromBase32 ...
func NewAddressFromBase32(base32Str string) (cfxAddress Address, err error) {
	if strings.ToLower(base32Str) != base32Str && strings.ToUpper(base32Str) != base32Str {
		return cfxAddress, errors.Errorf("not support base32 string with mix lowercase and uppercase %v", base32Str)
	}
	base32Str = strings.ToLower(base32Str)

	parts := strings.Split(base32Str, ":")
	if len(parts) < 2 || len(parts) > 3 {
		return cfxAddress, errors.Errorf("base32 string %v is invalid format", base32Str)
	}

	cfxAddress.NetworkType, err = NewNetowrkType(parts[0])
	if err != nil {
		return cfxAddress, errors.Wrapf(err, "failed to get network type of %v", parts[0])
	}

	bodyWithChecksum := parts[len(parts)-1]
	if len(bodyWithChecksum) < 8 {
		return cfxAddress, errors.Errorf("body with checksum length chould not length than 8, actual len(%v)=%v", bodyWithChecksum, len(bodyWithChecksum))
	}
	bodyStr := bodyWithChecksum[0 : len(bodyWithChecksum)-8]

	cfxAddress.Body, err = NewBodyByString(bodyStr)
	if err != nil {
		return cfxAddress, errors.Wrapf(err, "failed to create body by %v", bodyStr)
	}

	_, hexAddress, err := cfxAddress.Body.ToHexAddress()
	if err != nil {
		return cfxAddress, errors.Wrapf(err, "failed to get hex address by body %v", cfxAddress.Body)
	}

	cfxAddress.AddressType, err = CalcAddressType(hexAddress)
	if err != nil {
		return cfxAddress, errors.Wrapf(err, "failed to calc address type of %v", hexAddress)
	}

	if len(parts) == 3 && strings.ToLower(parts[1]) != cfxAddress.AddressType.String() {
		return cfxAddress, errors.Errorf("invalid address type, expect %v actual %v", cfxAddress.AddressType, parts[1])

	}

	cfxAddress.Checksum, err = CalcChecksum(cfxAddress.NetworkType, cfxAddress.Body)
	if err != nil {
		return cfxAddress, errors.Wrapf(err, "failed to calc checksum by network type %v and body %x", cfxAddress.NetworkType, cfxAddress.Body)
	}

	expectChk := cfxAddress.Checksum.String()
	actualChk := bodyWithChecksum[len(bodyWithChecksum)-8:]
	if expectChk != actualChk {
		err = errors.Errorf("invalid checksum, expect %v actual %v", expectChk, actualChk)
	}
	return
}

// NewAddressFromHex encode hex address with networkID to base32 address according to CIP37
// INPUT: an addr (20-byte conflux-hex-address), a network-id (4 bytes)
// OUTPUT: a conflux-base32-address
func NewAddressFromHex(hexAddressStr string, chainID ...uint32) (val Address, err error) {
	if hexAddressStr[0:2] == "0x" {
		hexAddressStr = hexAddressStr[2:]
	}

	hexAddress, err := hex.DecodeString(hexAddressStr)
	if err != nil {
		return val, errors.Wrapf(err, "failed to decode address string %x to hex", hexAddressStr)
	}

	return newAddressFromBytes(hexAddress, chainID...)
}

// MustNewAddressFromHex ...
func MustNewAddressFromHex(hexAddressStr string, chainID ...uint32) (val Address) {
	addr, err := NewAddressFromHex(hexAddressStr, get1stChainIDIfy(chainID))
	PanicIfErrf(err, "input hex address:%v, chainID:%v", hexAddressStr, chainID)
	return addr
}

// NewAddressFromCommon creates an address from common.Address
func NewAddressFromCommon(commonAddress common.Address, chainID ...uint32) (val Address, err error) {
	return newAddressFromBytes(commonAddress.Bytes(), chainID...)
}

func newAddressFromBytes(hexAddress []byte, chainID ...uint32) (val Address, err error) {
	val.NetworkType = NewNetworkTypeByID(get1stChainIDIfy(chainID))
	val.AddressType, err = CalcAddressType(hexAddress)

	if err != nil {
		return val, errors.Wrapf(err, "failed to calculate address type of %x", hexAddress)
	}

	versionByte, err := CalcVersionByte(hexAddress)
	if err != nil {
		return val, errors.Wrapf(err, "failed to calculate version type of %x", hexAddress)
	}

	val.Body, err = NewBodyByHexAddress(versionByte, hexAddress)
	if err != nil {
		return val, errors.Wrapf(err, "failed to create body by version byte %v and hex address %x", versionByte, hexAddress)
	}

	val.Checksum, err = CalcChecksum(val.NetworkType, val.Body)
	if err != nil {
		return val, errors.Wrapf(err, "failed to calc checksum by network type %v and body %x", val.NetworkType, val.Body)
	}
	return val, nil
}

// MustNewAddressFromCommon ...
func MustNewAddressFromCommon(commonAddress common.Address, chainID ...uint32) (address Address) {
	addr, err := NewAddressFromCommon(commonAddress, get1stChainIDIfy(chainID))
	PanicIfErrf(err, "input common address:%x, chainID:%v", commonAddress, chainID)
	return addr
}

// ToHexAddress ...
func (a *Address) ToHexAddress() (hexAddressStr string, networkID uint32, err error) {
	// verify checksum
	var actualChecksum Checksum
	actualChecksum, err = CalcChecksum(a.NetworkType, a.Body)
	if err != nil {
		return
	}

	if actualChecksum != a.Checksum {
		err = errors.Errorf("invalid checksum, expect %v actual %v", a.Checksum, actualChecksum)
		return
	}

	var hexAddress []byte
	_, hexAddress, err = a.Body.ToHexAddress()
	if err != nil {
		return
	}
	hexAddressStr = hex.EncodeToString(hexAddress)

	networkID, err = a.NetworkType.ToNetworkID()
	if err != nil {
		return
	}
	return
}

// ToCommonAddress converts address to common.Address
func (a *Address) ToCommonAddress() (address common.Address, chainID uint32, err error) {
	hexAddr, chainID, err := a.ToHexAddress()
	if err != nil {
		err = errors.Wrap(err, "failed to get hex address")
		return
	}

	if len(hexAddr) > 42 {
		err = errors.Errorf("could not convert %v to common address which length large than 42", hexAddr)
		return
	}
	address = common.HexToAddress(hexAddr)
	return
}

// MustGetHexAddress ...
func (a *Address) MustGetHexAddress() string {
	addr, _, err := a.ToHexAddress()
	PanicIfErrf(err, "failed to get hex address of %v", a)
	return addr
}

// MustGetCommonAddress ...
func (a *Address) MustGetCommonAddress() common.Address {
	addr, _, err := a.ToCommonAddress()
	PanicIfErrf(err, "failed to get common address of %v", a)
	return addr
}

// CompleteAddressByClient ...
func (a *Address) CompleteAddressByClient(client ChainIDGetter) error {
	chainID, err := client.GetChainID()
	if err != nil {
		return errors.Wrapf(err, "failed to get chainID")
	}
	a.CompleteAddressByChainID(chainID)
	return nil
}

// CompleteAddressByChainID ...
func (a *Address) CompleteAddressByChainID(chainID uint32) error {
	if a == nil {
		return nil
	}

	id, err := a.NetworkType.ToNetworkID()
	if err != nil || id == 0 {
		a.NetworkType = NewNetworkTypeByID(chainID)
		a.Checksum, err = CalcChecksum(a.NetworkType, a.Body)
		if err != nil {
			return errors.Wrapf(err, "failed to calc checksum by network type %v and body %v", a.NetworkType, a.Body)
		}
	}
	return nil
}

// IsValid return true if address is valid
func (a *Address) IsValid() bool {
	return a.AddressType == AddressTypeNull ||
		a.AddressType == AddressTypeContract ||
		a.AddressType == AddressTypeUser ||
		a.AddressType == AddressTypeBuiltin
}

func (a Address) MarshalText() ([]byte, error) {
	// fmt.Println("marshal text for epoch")
	return []byte(a.String()), nil
}

func (a *Address) UnmarshalJSON(data []byte) error {
	// fmt.Println("json unmarshal address")
	if string(data) == "null" {
		return nil
	}

	data = data[1 : len(data)-1]

	addr, err := NewAddressFromBase32(string(data))
	if err != nil {
		return errors.Wrapf(err, "failed to create address from base32 string %v", string(data))
	}
	*a = addr
	return nil
}

// PanicIfErrf ...
func PanicIfErrf(err error, msg string, args ...interface{}) {
	if err != nil {
		fmt.Printf(msg, args...)
		fmt.Println()
		panic(err)
	}
}

func PanicIfErr(err error, msg string) {
	if err != nil {
		fmt.Printf(msg)
		fmt.Println()
		panic(err)
	}
}

func get1stChainIDIfy(chainID []uint32) uint32 {
	if len(chainID) > 0 {
		return chainID[0]
	}
	return 0
}

// ChainIDGetter ...
type ChainIDGetter interface {
	GetChainID() (uint32, error)
}
