package types

import (
	"encoding/json"
	"fmt"
)

/*
#[derive(Debug, PartialEq)]
pub enum CfxFilterChanges {
    /// New logs.
    Logs(Vec<CfxFilterLog>),
    /// New hashes (block or transactions)
    Hashes(Vec<H256>),
    /// Empty result
    Empty,
}
*/

type CfxFilterChanges struct {
	Type   string
	Logs   []CfxFilterLog
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
	fmt.Printf("received data string: %s\n", data)
	if string(data) == `[]` {
		u.Type = "empty"
		return nil
	}

	logs := []CfxFilterLog{}
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
