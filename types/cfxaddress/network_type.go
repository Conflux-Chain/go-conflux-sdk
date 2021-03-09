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

// NetworkType reprents network type mapped with network-id
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

// NewNetowrkType creates network type by string
func NewNetowrkType(netType string) (NetworkType, error) {
	if netType == NetworkTypeMainnetPrefix.String() || netType == NetworkTypeTestNetPrefix.String() {
		return NetworkType(netType), nil
	}
	_, err := getIDWhenBeginWithNet(netType)
	if err != nil {
		return "", err
	}
	return NetworkType(netType), nil
}

// NewNetworkTypeByID creates network type by network ID
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

// ToNetworkID returns network ID
func (n NetworkType) ToNetworkID() (uint32, error) {
	switch n {
	case NetworkTypeMainnetPrefix:
		return NetowrkTypeMainnetID, nil
	case NetworkTypeTestNetPrefix:
		return NetworkTypeTestnetID, nil
	default:
		return getIDWhenBeginWithNet(string(n))
	}
}

func getIDWhenBeginWithNet(netIDStr string) (uint32, error) {
	if len(netIDStr) < 3 {
		return 0, errors.Errorf("Invalid network: %v", netIDStr)
	}

	if netIDStr[0:3] != "net" {
		return 0, errors.Errorf("Invalid network: %v", netIDStr)
	}

	netID, err := strconv.ParseUint(string(netIDStr[3:]), 0, 32)
	if err != nil {
		return 0, err
	}
	return uint32(netID), nil
}
