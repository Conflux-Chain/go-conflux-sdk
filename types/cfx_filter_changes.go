package types

import (
	"encoding/json"
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
		return json.Marshal(u.Logs)
	case "hash":
		return json.Marshal(u.Hashes)
	}
	return []byte(`[]`), nil
}

func (u *CfxFilterChanges) UnmarshalJSON(data []byte) error {
	if string(data) == `[]` {
		u.Type = "empty"
		return nil
	}

	logs := []*SubscriptionLog{}
	err := utils.JSONUnmarshal(data, &logs)
	if err == nil {
		u.Logs = logs
		u.Type = "log"
		return nil
	}
	hashes := []Hash{}
	err = utils.JSONUnmarshal(data, &hashes)
	if err == nil {
		u.Hashes = hashes
		u.Type = "hash"
		return nil
	}
	return err
}
