package cfxaddress

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"

	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// Address represents
type Address struct {
	NetworkType NetworkType
	AddressType AddressType
	Body        Body
	Checksum    Checksum
}

// String returns verbose string of address
func (a Address) String() string {
	return strings.ToUpper(fmt.Sprintf("%v:%v:%v%v", a.NetworkType, a.AddressType, a.Body, a.Checksum))
}

// ShortString returns short string of address which doesn't include address type
func (a Address) ShortString() string {
	return fmt.Sprintf("%v:%v%v", a.NetworkType, a.Body, a.Checksum)
}

// Equals reports whether a and target are equal
func (a *Address) Equals(target *Address) bool {
	return reflect.DeepEqual(a, target)
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

// MustNewFromHex creates address by hex address string with networkID and panic if error
func MustNewFromHex(hexAddressStr string, networkID ...uint32) (val Address) {
	addr, err := NewFromHex(hexAddressStr, get1stNetworkIDIfy(networkID))
	utils.PanicIfErrf(err, "input hex address:%v, networkID:%v", hexAddressStr, networkID)
	return addr
}

// NewFromCommon creates an address from common.Address with networkID
func NewFromCommon(commonAddress common.Address, networkID ...uint32) (val Address, err error) {
	return NewFromBytes(commonAddress.Bytes(), networkID...)
}

// NewFromBytes creates an address from hexAddress byte slice with networkID
func NewFromBytes(hexAddress []byte, networkID ...uint32) (val Address, err error) {
	val.NetworkType = NewNetworkTypeByID(get1stNetworkIDIfy(networkID))
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

// MustNewFromCommon creates an address from common.Address with networkID and panic if error
func MustNewFromCommon(commonAddress common.Address, networkID ...uint32) (address Address) {
	addr, err := NewFromCommon(commonAddress, get1stNetworkIDIfy(networkID))
	utils.PanicIfErrf(err, "input common address:%x, networkID:%v", commonAddress, networkID)
	return addr
}

// ToHex returns hex address and networkID
func (a *Address) ToHex() (hexAddressStr string, networkID uint32, err error) {
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

// ToCommon converts address to common.Address
func (a *Address) ToCommon() (address common.Address, networkID uint32, err error) {
	hexAddr, networkID, err := a.ToHex()
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

// MustGetHexAddress returns hex format address and panic if error
func (a *Address) MustGetHexAddress() string {
	addr, _, err := a.ToHex()
	utils.PanicIfErrf(err, "failed to get hex address of %v", a)
	return addr
}

// MustGetCommonAddress returns common address and panic if error
func (a *Address) MustGetCommonAddress() common.Address {
	addr, _, err := a.ToCommon()
	utils.PanicIfErrf(err, "failed to get common address of %v", a)
	return addr
}

// CompleteByClient will set networkID by client.GetNetworkID() if a.networkID not be 0
func (a *Address) CompleteByClient(client NetworkIDGetter) error {
	networkID, err := client.GetNetworkID()
	if err != nil {
		return errors.Wrapf(err, "failed to get networkID")
	}
	a.CompleteByNetworkID(networkID)
	return nil
}

// CompleteByNetworkID will set networkID if a.networkID not be 0
func (a *Address) CompleteByNetworkID(networkID uint32) error {
	if a == nil {
		return nil
	}

	id, err := a.NetworkType.ToNetworkID()
	if err != nil || id == 0 {
		a.NetworkType = NewNetworkTypeByID(networkID)
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

// MarshalText implements the encoding.TextMarshaler interface.
func (a Address) MarshalText() ([]byte, error) {
	// fmt.Println("marshal text for epoch")
	return []byte(a.String()), nil
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

// NetworkIDGetter is a interface for obtaining networkID
type NetworkIDGetter interface {
	GetNetworkID() (uint32, error)
}
