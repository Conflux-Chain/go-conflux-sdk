package types

import (
	"github.com/Conflux-Chain/go-conflux-sdk/utils"
)

type CfxFilterChanges struct {
	Type   string
	Logs   []*SubscriptionLog
	Hashes []Hash
}

func (u CfxFilterChanges) MarshalJSON() ([]byte, error) {
	switch u.Type {
	case "log":
		return utils.JsonMarshal(u.Logs)
	case "hash":
		return utils.JsonMarshal(u.Hashes)
	}
	return []byte(`[]`), nil
}

func (u *CfxFilterChanges) UnmarshalJSON(data []byte) error {
	if string(data) == `[]` {
		u.Type = "empty"
		return nil
	}

	logs := []*SubscriptionLog{}
	err := utils.JsonUnmarshal(data, &logs)
	if err == nil {
		u.Logs = logs
		u.Type = "log"
		return nil
	}
	hashes := []Hash{}
	err = utils.JsonUnmarshal(data, &hashes)
	if err == nil {
		u.Hashes = hashes
		u.Type = "hash"
		return nil
	}
	return err
}
