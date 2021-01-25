package cfxaddress

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

/*
Network-prefix:
    match network-id:
        case 1029: "cfx"
        case 1:    "cfxtest"
        case n:    "net[n]"
Examples of valid network-prefixes: "cfx", "cfxtest", "net17"
Examples of invalid network-prefixes: "bch", "conflux", "net1", "net1029"
*/

// NetworkType ...
type NetworkType string

func (n NetworkType) String() string {
	return string(n)
}

const (
	NetworkTypeMainnetPrefix NetworkType = "cfx"
	NetworkTypeTestNetPrefix NetworkType = "cfxtest"

	NetowrkTypeMainnetID uint32 = 1029
	NetworkTypeTestnetID uint32 = 1
)

// NewNetowrkType ...
func NewNetowrkType(nt string) (NetworkType, error) {
	if nt == NetworkTypeMainnetPrefix.String() || nt == NetworkTypeTestNetPrefix.String() {
		return NetworkType(nt), nil
	}
	if nt[0:3] == "net" {
		chainID, err := strconv.Atoi(nt[3:])
		if err != nil {
			return "", errors.Wrapf(err, "chainID %v is not uint32", chainID)
		}
		if chainID >= (1 << 32) {
			return "", errors.Errorf("NetworkID %v not in range 0~0xffffffff", chainID)
		}
		return NetworkType(nt), nil
	}
	return "", errors.New("invalid network type")
}

// NewNetworkTypeByID ...
func NewNetworkTypeByID(networkID uint32) NetworkType {
	var nt NetworkType
	switch networkID {
	case NetowrkTypeMainnetID:
		nt = NetworkTypeMainnetPrefix
	case NetworkTypeTestnetID:
		nt = NetworkTypeTestNetPrefix
	default:
		nt = NetworkType(fmt.Sprintf("net%v", networkID))
	}
	return nt
}

// ToNetworkID ...
func (n NetworkType) ToNetworkID() (uint32, error) {
	switch n {
	case NetworkTypeMainnetPrefix:
		return NetowrkTypeMainnetID, nil
	case NetworkTypeTestNetPrefix:
		return NetworkTypeTestnetID, nil
	default:
		if n[0:3] == "net" {
			netID, err := strconv.Atoi(string(n[3:]))
			if err != nil {
				return 0, err
			}
			if netID >= (1 << 32) {
				return 0, errors.Errorf("NetworkID %v not in range 0~0xffffffff", netID)
			}
			return uint32(netID), nil
		}
		return 0, errors.New("Invalid network")
	}
}
