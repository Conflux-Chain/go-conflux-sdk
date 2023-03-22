package types

import (
	"encoding/json"
	"testing"

	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/stretchr/testify/assert"
)

func TestMarshalCfxFilterChanges(t *testing.T) {
	// [{"address":"cfxtest:aang4d91rejdbpgmgtmspdyefxkubj2bbywrwm9j3z","topics":null,"data":"0x"},{"revertTo":"0x1"}]
	logs := []SubscriptionLog{
		{Log: &Log{Address: cfxaddress.MustNew("cfxtest:aang4d91rejdbpgmgtmspdyefxkubj2bbywrwm9j3z")}},
		{ChainReorg: &ChainReorg{RevertTo: NewBigInt(1)}},
	}
	cfc := CfxFilterChanges{
		Type: "log",
		Logs: logs,
	}
	j1, _ := json.Marshal(logs)
	j2, _ := json.Marshal(cfc)
	assert.Equal(t, j1, j2)

	hashes := []Hash{"0x564cdd312483a5007740ebca3210bbd53ef390045170539b56da0ae42f8c3f9c"}
	cfc = CfxFilterChanges{
		Type:   "hash",
		Hashes: hashes,
	}
	j1, _ = json.Marshal(hashes)
	j2, _ = json.Marshal(cfc)
	assert.Equal(t, j1, j2)
}

func TestUnmarshalCfxFilterChanges(t *testing.T) {
	logs := []SubscriptionLog{
		{Log: &Log{Address: cfxaddress.MustNew("cfxtest:aang4d91rejdbpgmgtmspdyefxkubj2bbywrwm9j3z")}},
		{ChainReorg: &ChainReorg{RevertTo: NewBigInt(1)}},
	}
	j1, _ := json.Marshal(logs)

	var cfc CfxFilterChanges
	err := json.Unmarshal(j1, &cfc)
	assert.NoError(t, err)
	assert.Equal(t, "log", cfc.Type)
	j2, _ := json.Marshal(cfc)
	assert.Equal(t, j1, j2)

	hashes := []Hash{"0x564cdd312483a5007740ebca3210bbd53ef390045170539b56da0ae42f8c3f9c"}
	j1, _ = json.Marshal(hashes)
	err = json.Unmarshal(j1, &cfc)
	assert.NoError(t, err)
	assert.Equal(t, "hash", cfc.Type)
	j2, _ = json.Marshal(cfc)
	assert.Equal(t, j1, j2)

	empty := []string{}
	j1, _ = json.Marshal(empty)
	err = json.Unmarshal(j1, &cfc)
	assert.NoError(t, err)
	assert.Equal(t, "empty", cfc.Type)
	j2, _ = json.Marshal(cfc)
	assert.Equal(t, j1, j2)
}
