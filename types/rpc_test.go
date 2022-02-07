package types

import (
	"encoding/json"
	"testing"
)

func TestJsonStatus(t *testing.T) {
	j := `{"bestHash": "0x3370751173cae9b37ec171f5fc234d58c597138336a23d26ce9b34fa907244af","blockNumber": "0x4b983fe","chainId": "0x1","epochNumber": "0x3c097d3","ethereumSpaceChainId": "0x47","latestCheckpoint": "0x3bf2ae0","latestConfirmed": "0x3c09784","latestFinalized": "0x3c0954c","latestState": "0x3c097cf","networkId": "0x1","pendingTxNumber": "0x0"}`
	s := Status{}
	e := json.Unmarshal([]byte(j), &s)
	if e != nil {
		t.Fatal(e)
	}
	// fmt.Printf("%+v\n", s)
}
