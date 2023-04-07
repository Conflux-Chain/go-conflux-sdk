package types

import (
	"encoding/json"
	"testing"

	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/stretchr/testify/assert"
)

func TestMarshalCfxFilterChanges(t *testing.T) {
	// [{"address":"cfxtest:aang4d91rejdbpgmgtmspdyefxkubj2bbywrwm9j3z","topics":null,"data":"0x"},{"revertTo":"0x1"}]
	logs := []CfxFilterLog{
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

func TestUnmarshalCfxFilterChanges2(t *testing.T) {
	str := `[{"Log":{"address":"CFXTEST:TYPE.CONTRACT:ACHS3NEHAE0J6KSVY1BHRFFSH1RTFRW1F6W1KZV46T","blockHash":"0xcda21448052710d94f97664f4dfdd9a00d73e1cfb9c99f27a6f27ade0d3f43bf","data":"0x00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","epochNumber":"0x7087a52","logIndex":"0x0","topics":["0x2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d","0x000000000000000000000000187f1d870c7da2a5790c16ab6ee02279e0401c95","0x000000000000000000000000187f1d870c7da2a5790c16ab6ee02279e0401c95"],"transactionHash":"0x193e135c59c681987c22779a912955664e2c5e35c0f95f50cc897e756eab83d9","transactionIndex":"0x0","transactionLogIndex":"0x0"}},{"Log":{"address":"CFXTEST:TYPE.CONTRACT:ACHS3NEHAE0J6KSVY1BHRFFSH1RTFRW1F6W1KZV46T","blockHash":"0xcda21448052710d94f97664f4dfdd9a00d73e1cfb9c99f27a6f27ade0d3f43bf","data":"0x0000000000000000000000000000000000000000000000000000000000000001","epochNumber":"0x7087a52","logIndex":"0x1","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x0000000000000000000000000000000000000000000000000000000000000000","0x000000000000000000000000187f1d870c7da2a5790c16ab6ee02279e0401c95"],"transactionHash":"0x193e135c59c681987c22779a912955664e2c5e35c0f95f50cc897e756eab83d9","transactionIndex":"0x0","transactionLogIndex":"0x1"}},{"Log":{"address":"CFXTEST:TYPE.CONTRACT:ACHS3NEHAE0J6KSVY1BHRFFSH1RTFRW1F6W1KZV46T","blockHash":"0xcda21448052710d94f97664f4dfdd9a00d73e1cfb9c99f27a6f27ade0d3f43bf","data":"0x0000000000000000000000000000000000000000000000000000000000000001","epochNumber":"0x7087a52","logIndex":"0x2","topics":["0x68051bc50b1ef1654bf1e6204b5f8fa9badcd038e00fa5b43f21f898fc2728ca","0x000000000000000000000000187f1d870c7da2a5790c16ab6ee02279e0401c95","0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"],"transactionHash":"0x193e135c59c681987c22779a912955664e2c5e35c0f95f50cc897e756eab83d9","transactionIndex":"0x0","transactionLogIndex":"0x2"}}]`

	var cfc CfxFilterChanges
	err := json.Unmarshal([]byte(str), &cfc)

	assert.NoError(t, err)
	assert.NotNil(t, cfc.Logs[0].Log.Topics)
}
