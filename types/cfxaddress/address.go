package cfxaddress

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// Address represents
type Address struct {
	networkType NetworkType
	addressType AddressType
	body        Body
	checksum    Checksum

	// cache
	hex       []byte
	networkID uint32
}

// String returns verbose base32 string of address
func (a Address) String() string {
	return a.MustGetBase32Address()
}

// Equals reports whether a and target are equal
func (a *Address) Equals(target *Address) bool {
	return reflect.DeepEqual(a, target)
}

// New create conflux address by base32 string or hex40 string, if base32OrHex is base32 and networkID is passed it will create cfx Address use networkID.
func New(base32OrHex string, networkID ...uint32) (Address, error) {
	hexPattern := `(?i)^0x[a-f0-9]{40}$`
	base32Pattern := `(?i)^(cfx|cfxtest|net\d+):(type\.user:|type\.builtin:|type\.contract:|type\.null:|)\w{42}$`
	_networkID := uint32(0)
	if len(networkID) > 0 {
		_networkID = networkID[0]
	}

	if ok, _ := regexp.Match(hexPattern, []byte(base32OrHex)); ok {

		addr, err := NewFromHex(base32OrHex, _networkID)
		if err != nil {
			return Address{}, errors.Wrapf(err, "Failed to new address from hex %v networkID %v", base32OrHex, _networkID)
		}
		return addr, nil
	}

	if ok, _ := regexp.Match(base32Pattern, []byte(base32OrHex)); ok {
		addr, err := NewFromBase32(base32OrHex)
		if err != nil {
			return Address{}, errors.Wrapf(err, "Failed to new address from base32 string %v", base32OrHex)
		}
		if _networkID != 0 && addr.GetNetworkID() != _networkID {
			addr, err = NewFromHex(addr.GetHexAddress(), _networkID)
			if err != nil {
				return Address{}, errors.Wrapf(err, "Failed to new address from hex %v networkID %v", addr.GetHexAddress(), _networkID)
			}
		}
		return addr, nil
	}
	return Address{}, errors.Errorf("Input %v need be base32 string or hex40 string,", base32OrHex)
}

// MustNew create conflux address by base32 string or hex40 string, if base32OrHex is base32 and networkID is setted it will check if networkID match,
// it will painc if error occured.
func MustNew(base32OrHex string, networkID ...uint32) Address {
	address, err := New(base32OrHex, networkID...)
	if err != nil {
		panic(err)
	}
	return address
}

// NewFromBase32 creates address by base32 string
func NewFromBase32(base32Str string) (cfxAddress Address, err error) {
	if strings.ToLower(base32Str) != base32Str && strings.ToUpper(base32Str) != base32Str {
		return cfxAddress, errors.Errorf("not support base32 string with mix lowercase and uppercase %v", base32Str)
	}
	base32Str = strings.ToLower(base32Str)

	parts := strings.Split(base32Str, ":")
	if len(parts) < 2 || len(parts) > 3 {
		return cfxAddress, errors.Errorf("base32 string %v is invalid format", base32Str)
	}

	cfxAddress.networkType, err = NewNetowrkType(parts[0])
	if err != nil {
		return cfxAddress, errors.Wrapf(err, "failed to get network type of %v", parts[0])
	}

	bodyWithChecksum := parts[len(parts)-1]
	if len(bodyWithChecksum) < 8 {
		return cfxAddress, errors.Errorf("body with checksum length chould not length than 8, actual len(%v)=%v", bodyWithChecksum, len(bodyWithChecksum))
	}
	bodyStr := bodyWithChecksum[0 : len(bodyWithChecksum)-8]

	cfxAddress.body, err = NewBodyByString(bodyStr)
	if err != nil {
		return cfxAddress, errors.Wrapf(err, "failed to create body by %v", bodyStr)
	}

	_, hexAddress, err := cfxAddress.body.ToHexAddress()
	if err != nil {
		return cfxAddress, errors.Wrapf(err, "failed to get hex address by body %v", cfxAddress.body)
	}

	cfxAddress.addressType, err = CalcAddressType(hexAddress)
	if err != nil {
		return cfxAddress, errors.Wrapf(err, "failed to calc address type of %v", hexAddress)
	}

	if len(parts) == 3 && strings.ToLower(parts[1]) != cfxAddress.addressType.String() {
		return cfxAddress, errors.Errorf("invalid address type, expect %v actual %v", cfxAddress.addressType, parts[1])

	}

	cfxAddress.checksum, err = CalcChecksum(cfxAddress.networkType, cfxAddress.body)
	if err != nil {
		return cfxAddress, errors.Wrapf(err, "failed to calc checksum by network type %v and body %x", cfxAddress.networkType, cfxAddress.body)
	}

	expectChk := cfxAddress.checksum.String()
	actualChk := bodyWithChecksum[len(bodyWithChecksum)-8:]
	if expectChk != actualChk {
		err = errors.Errorf("invalid checksum, expect %v actual %v", expectChk, actualChk)
	}

	if err := cfxAddress.setCache(); err != nil {
		err = errors.Wrapf(err, "failed to set cache")
	}

	return
}

// NewFromHex creates address by hex address string with networkID
// If not pass networkID, it will be auto completed when it could be obtained form context.
func NewFromHex(hexAddressStr string, networkID ...uint32) (val Address, err error) {
	if hexAddressStr[0:2] == "0x" {
		hexAddressStr = hexAddressStr[2:]
	}

	hexAddress, err := hex.DecodeString(hexAddressStr)
	if err != nil {
		return val, errors.Wrapf(err, "failed to decode address string %x to hex", hexAddressStr)
	}

	return NewFromBytes(hexAddress, networkID...)
}

// NewFromCommon creates an address from common.Address with networkID
func NewFromCommon(commonAddress common.Address, networkID ...uint32) (val Address, err error) {
	return NewFromBytes(commonAddress.Bytes(), networkID...)
}

// NewFromBytes creates an address from hexAddress byte slice with networkID
func NewFromBytes(hexAddress []byte, networkID ...uint32) (val Address, err error) {
	val.networkType = NewNetworkTypeByID(get1stNetworkIDIfy(networkID))
	val.addressType, err = CalcAddressType(hexAddress)

	if err != nil {
		return val, errors.Wrapf(err, "failed to calculate address type of %x", hexAddress)
	}

	versionByte, err := CalcVersionByte(hexAddress)
	if err != nil {
		return val, errors.Wrapf(err, "failed to calculate version type of %x", hexAddress)
	}

	val.body, err = NewBodyByHexAddress(versionByte, hexAddress)
	if err != nil {
		return val, errors.Wrapf(err, "failed to create body by version byte %v and hex address %x", versionByte, hexAddress)
	}

	val.checksum, err = CalcChecksum(val.networkType, val.body)
	if err != nil {
		return val, errors.Wrapf(err, "failed to calc checksum by network type %v and body %x", val.networkType, val.body)
	}

	if err = val.setCache(); err != nil {
		err = errors.Wrapf(err, "failed to set cache")
		return
	}

	return val, nil
}

// MustNewFromBase32 creates address by base32 string and panic if error
func MustNewFromBase32(base32Str string) (address Address) {
	address, err := NewFromBase32(base32Str)
	if err != nil {
		utils.PanicIfErrf(err, "input base32 string: %v", base32Str)
	}
	return
}

// MustNewFromHex creates address by hex address string with networkID and panic if error
func MustNewFromHex(hexAddressStr string, networkID ...uint32) (val Address) {
	addr, err := NewFromHex(hexAddressStr, get1stNetworkIDIfy(networkID))
	utils.PanicIfErrf(err, "input hex address:%v, networkID:%v", hexAddressStr, networkID)
	return addr
}

// MustNewFromCommon creates an address from common.Address with networkID and panic if error
func MustNewFromCommon(commonAddress common.Address, networkID ...uint32) (address Address) {
	addr, err := NewFromCommon(commonAddress, get1stNetworkIDIfy(networkID))
	utils.PanicIfErrf(err, "input common address:%x, networkID:%v", commonAddress, networkID)
	return addr
}

// MustNewFromBytes creates an address from hexAddress byte slice with networkID and panic if error
func MustNewFromBytes(hexAddress []byte, networkID ...uint32) (address Address) {
	addr, err := NewFromBytes(hexAddress, get1stNetworkIDIfy(networkID))
	utils.PanicIfErrf(err, "input common address:%x, networkID:%v", hexAddress, networkID)
	return addr
}

// ToHex returns hex address and networkID
func (a *Address) ToHex() (hexAddressStr string, networkID uint32) {
	return "0x" + hex.EncodeToString(a.hex), a.networkID
}

// ToCommon returns common.Address and networkID
func (a *Address) ToCommon() (address common.Address, networkID uint32, err error) {
	if len(a.hex) > 42 {
		err = errors.Errorf("could not convert %v to common address which length large than 42", a.hex)
		return
	}
	address = common.BytesToAddress(a.hex)
	return
}

// MustGetBase32Address returns base32 string of address which doesn't include address type
func (a *Address) MustGetBase32Address() string {
	return strings.ToLower(fmt.Sprintf("%v:%v%v", a.networkType, a.body, a.checksum))
}

// MustGetVerboseBase32Address returns base32 string of address with address type
func (a *Address) MustGetVerboseBase32Address() string {
	return strings.ToUpper(fmt.Sprintf("%v:%v:%v%v", a.networkType, a.addressType, a.body, a.checksum))
}

// GetHexAddress returns hex format address and panic if error
func (a *Address) GetHexAddress() string {
	addr, _ := a.ToHex()
	return addr
}

// GetNetworkID returns networkID and panic if error
func (a *Address) GetNetworkID() uint32 {
	id, err := a.networkType.ToNetworkID()
	utils.PanicIfErrf(err, "failed to get networkID of %v", a)
	return id
}

// MustGetCommonAddress returns common address and panic if error
func (a *Address) MustGetCommonAddress() common.Address {
	addr, _, err := a.ToCommon()
	utils.PanicIfErrf(err, "failed to get common address of %v", a)
	return addr
}

// GetNetworkType returns network type
func (a *Address) GetNetworkType() NetworkType {
	return a.networkType
}

// GetAddressType returuns address type
func (a *Address) GetAddressType() AddressType {
	return a.addressType
}

// GetBody returns body
func (a *Address) GetBody() Body {
	return a.body
}

// GetChecksum returns checksum
func (a *Address) GetChecksum() Checksum {
	return a.checksum
}

// CompleteByClient will set networkID by client.GetNetworkID() if a.networkID not be 0
func (a *Address) CompleteByClient(client networkIDGetter) error {
	networkID, err := client.GetNetworkID()
	if err != nil {
		return errors.Wrapf(err, "failed to get networkID")
	}
	a.CompleteByNetworkID(networkID)
	return nil
}

// CompleteByNetworkID will set networkID if current networkID isn't 0
func (a *Address) CompleteByNetworkID(networkID uint32) error {
	if a == nil {
		return nil
	}

	id, err := a.networkType.ToNetworkID()
	if err != nil || id == 0 {
		a.networkType = NewNetworkTypeByID(networkID)
		a.checksum, err = CalcChecksum(a.networkType, a.body)
		if err != nil {
			return errors.Wrapf(err, "failed to calc checksum by network type %v and body %v", a.networkType, a.body)
		}
	}
	return nil
}

// IsValid return true if address is valid
func (a *Address) IsValid() bool {
	return a.addressType == AddressTypeNull ||
		a.addressType == AddressTypeContract ||
		a.addressType == AddressTypeUser ||
		a.addressType == AddressTypeBuiltin
}

// MarshalText implements the encoding.TextMarshaler interface.
func (a Address) MarshalText() ([]byte, error) {
	// fmt.Println("marshal text for epoch")
	return []byte(a.String()), nil
}

func (a *Address) UnmarshalText(data []byte) error {
	data = append([]byte{0}, data...)
	data = append(data, 0)
	return a.UnmarshalJSON(data)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (a *Address) UnmarshalJSON(data []byte) error {
	// fmt.Println("json unmarshal address")
	if string(data) == "null" {
		return nil
	}

	data = data[1 : len(data)-1]

	addr, err := NewFromBase32(string(data))
	if err != nil {
		return errors.Wrapf(err, "failed to create address from base32 string %v", string(data))
	}
	*a = addr
	return nil
}

func get1stNetworkIDIfy(networkID []uint32) uint32 {
	if len(networkID) > 0 {
		return networkID[0]
	}
	return 0
}

func (a *Address) setCache() error {
	var hexAddress []byte
	_, hexAddress, err := a.body.ToHexAddress()
	if err != nil {
		return errors.Wrapf(err, "failed convert %v to hex address", a.body)
	}
	a.hex = hexAddress

	networkID, err := a.networkType.ToNetworkID()
	if err != nil {
		return errors.Wrapf(err, "failed to get networkID of %v", networkID)
	}
	a.networkID = networkID
	return nil
}

// networkIDGetter is a interface for obtaining networkID
type networkIDGetter interface {
	GetNetworkID() (uint32, error)
}
