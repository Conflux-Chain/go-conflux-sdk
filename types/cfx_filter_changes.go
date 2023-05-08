package types

import (
	"encoding/json"
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
	err := json.Unmarshal(data, &logs)
	if err == nil {
		u.Logs = logs
		u.Type = "log"
		return nil
	}
	hashes := []Hash{}
	err = json.Unmarshal(data, &hashes)
	if err == nil {
		u.Hashes = hashes
		u.Type = "hash"
		return nil
	}
	return err
}
