package types

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/rlp"
)

func TestRLPMarshalTransaction(t *testing.T) {
	txJson := `{"hash":"0x4016c5b1675182700ef67b9df90c13ddf2e774b12385af63ba43576039b13f8a","nonce":"0x0","blockHash":"0xce70d3387731e0d09605c37c322070ba30223b327b8790b15086847241cb857d","transactionIndex":"0x0","from":"cfx:aam9edem02j9tmzavu81y4g0ygy06cyeg2vx3zwe8y","to":"cfx:acb59fk6vryh8dj5vyvehj9apzhpd72rdpwsc651kz","value":"0x0","gasPrice":"0x1","gas":"0x4c4b40","contractCreated":null,"data":"0xfebe49090000000000000000000000000000000000000000000000000000000000000000000000000000000000000000162788589c8e386863f217faef78840919fb2854","storageLimit":"0x40","epochHeight":"0x5","chainId":"0x405","status":"0x0","v":"0x0","r":"0x7c15d3f0e517a66ad96a79cfd24b00439de725d9cbbcb8a1d0107cd5c16c8751","s":"0x27a0faab73d4ab512fa97d9e52ef3b3d2bea2a39d58c8d3d65dac59abbe6fee5"}`

	var tx Transaction
	err := json.Unmarshal([]byte(txJson), &tx)
	fatalIfErr(t, err)
	// RLP marshal tx to bytes
	dBytes, err := rlp.EncodeToBytes(tx)
	fatalIfErr(t, err)
	// RLP unmarshal bytes back to transaction
	var tx2 Transaction
	err = rlp.DecodeBytes(dBytes, &tx2)
	fatalIfErr(t, err)
	// Json marshal tx
	jBytes1, err := json.Marshal(tx)
	fatalIfErr(t, err)
	txJsonStr := string(jBytes1)
	// Json marshal tx2
	jBytes2, err := json.Marshal(tx2)
	fatalIfErr(t, err)
	txJsonStr2 := string(jBytes2)

	if txJsonStr2 != txJsonStr {
		t.Fatalf("expect %#v, actual %#v", txJsonStr, txJsonStr2)
	}
}

func TestRLPMarshalTransactionReceipt(t *testing.T) {
	trJson := `{"transactionHash":"0xa2c678cc97e07ce060b71f87ac65e68d482abf8e1a93b7d1bc425504c4584ca7","index":"0x1","blockHash":"0x935bc7ebab4e6256ade6327682de8bee6b15dcee5c83f7772ae6a945510fe940","epochNumber":"0xf7cf0d","from":"cfx:aaspca00u76ew2eharzsy355svm3t6yw4pxsyu7vka","to":"cfx:acam64yj323zd4t1fhybxh3jsg7hu4012yz9kakxs9","gasUsed":"0x22144","gasFee":"0xd4fe90","contractCreated":null,"logs":[{"address":"cfx:acg158kvr8zanb1bs048ryb6rtrhr283ma70vz70tx","topics":["0x2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d","0x00000000000000000000000080ae6a88ce3351e9f729e8199f2871ba786ad7c5","0x0000000000000000000000008d545118d91c027c805c552f63a5c00a20ae6aca"],"data":"0x00000000000000000000000000000000000000000000003b16c9e8eeb7c800000000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","blockHash":null,"epochNumber":null,"transactionHash":null,"transactionIndex":null,"logIndex":null,"transactionLogIndex":null},{"address":"cfx:acg158kvr8zanb1bs048ryb6rtrhr283ma70vz70tx","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x0000000000000000000000000000000000000000000000000000000000000000","0x0000000000000000000000008d545118d91c027c805c552f63a5c00a20ae6aca"],"data":"0x00000000000000000000000000000000000000000000003b16c9e8eeb7c80000","blockHash":null,"epochNumber":null,"transactionHash":null,"transactionIndex":null,"logIndex":null,"transactionLogIndex":null},{"address":"cfx:acg158kvr8zanb1bs048ryb6rtrhr283ma70vz70tx","topics":["0x68051bc50b1ef1654bf1e6204b5f8fa9badcd038e00fa5b43f21f898fc2728ca","0x0000000000000000000000008d545118d91c027c805c552f63a5c00a20ae6aca","0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"],"data":"0x00000000000000000000000000000000000000000000003b16c9e8eeb7c80000","blockHash":null,"epochNumber":null,"transactionHash":null,"transactionIndex":null,"logIndex":null,"transactionLogIndex":null},{"address":"cfx:acf2rcsh8payyxpg6xj7b0ztswwh81ute60tsw35j7","topics":["0x06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987","0x0000000000000000000000008d545118d91c027c805c552f63a5c00a20ae6aca","0x0000000000000000000000008d545118d91c027c805c552f63a5c00a20ae6aca","0x0000000000000000000000001cc102d68778496087036aea677b745597f292d3"],"data":"0x000000000000000000000000000000000000000000000019ceb8990635656bbf0000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","blockHash":null,"epochNumber":null,"transactionHash":null,"transactionIndex":null,"logIndex":null,"transactionLogIndex":null},{"address":"cfx:acf2rcsh8payyxpg6xj7b0ztswwh81ute60tsw35j7","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x0000000000000000000000008d545118d91c027c805c552f63a5c00a20ae6aca","0x0000000000000000000000001cc102d68778496087036aea677b745597f292d3"],"data":"0x000000000000000000000000000000000000000000000019ceb8990635656bbf","blockHash":null,"epochNumber":null,"transactionHash":null,"transactionIndex":null,"logIndex":null,"transactionLogIndex":null},{"address":"cfx:acgzjyj25esae9eanvmw827f2afcbnxm3jr3784tyu","topics":["0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1"],"data":"0x000000000000000000000000000000000000000000003a4a9c46eba30a4fb89e00000000000000000000000000000000000000000000854b7ce25c8407e77655","blockHash":null,"epochNumber":null,"transactionHash":null,"transactionIndex":null,"logIndex":null,"transactionLogIndex":null},{"address":"cfx:acgzjyj25esae9eanvmw827f2afcbnxm3jr3784tyu","topics":["0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822","0x00000000000000000000000080ae6a88ce3351e9f729e8199f2871ba786ad7c5","0x0000000000000000000000001cc102d68778496087036aea677b745597f292d3"],"data":"0x000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003b16c9e8eeb7c80000000000000000000000000000000000000000000000000019ceb8990635656bbf0000000000000000000000000000000000000000000000000000000000000000","blockHash":null,"epochNumber":null,"transactionHash":null,"transactionIndex":null,"logIndex":null,"transactionLogIndex":null}],"logsBloom":"0x0020000000000000000000008000000000000808000000000000000000000000804000000000200040080000009000000000000000000000000000000008000000000000000010000000000a00000020000000000000000000000000000000000000000002004000000000000000080004000040000000000000001000000000000008000020000000001100000002000000000000000018000000400000000000000000012000000000000000000000000000000000000000000000000000000000000200000010000004000000000000000000000000100000a000000020000000000000000400000000000000000000000000100000000000000000040000","stateRoot":"0x2b35e444ba6118f85447efc1c504b3526229c99451fd172ae35a49fe5d7127bd","outcomeStatus":"0x0","txExecErrorMsg":null,"gasCoveredBySponsor":true,"storageCoveredBySponsor":true,"storageCollateralized":"0x0","storageReleased":[]}`

	var tr TransactionReceipt
	err := json.Unmarshal([]byte(trJson), &tr)
	fatalIfErr(t, err)
	// RLP marshal transaction receipt to bytes
	dBytes, err := rlp.EncodeToBytes(tr)
	fatalIfErr(t, err)
	// RLP unmarshal bytes back to transaction receipt
	var tr2 TransactionReceipt
	err = rlp.DecodeBytes(dBytes, &tr2)
	fatalIfErr(t, err)
	// Json marshal tr
	jBytes1, err := json.Marshal(tr)
	fatalIfErr(t, err)
	trJsonStr := string(jBytes1)
	// Json marshal tr2
	jBytes2, err := json.Marshal(tr2)
	fatalIfErr(t, err)
	trJsonStr2 := string(jBytes2)

	if trJsonStr2 != trJsonStr {
		t.Fatalf("expect %#v, actual %#v", trJsonStr, trJsonStr2)
	}
}

func TestMarshalTransactionStatus(t *testing.T) {
	testMarshalTransactionStatus(t, TransactionStatus{
		packedOrReady: "packed",
	}, "packed", []byte("\"packed\""))

	testMarshalTransactionStatus(t, TransactionStatus{
		pending: pending{"futureNonce"},
	}, "futureNonce", []byte("{\"pending\":\"futureNonce\"}"))

}
func testMarshalTransactionStatus(t *testing.T, originTxStatus TransactionStatus, expectString string, expectJson []byte) {
	if originTxStatus.String() != expectString {
		t.Fatalf("expect string %#v, actual %#v", expectString, originTxStatus.String())
	}

	actualJson, _ := json.Marshal(originTxStatus)
	if !reflect.DeepEqual(actualJson, expectJson) {
		t.Fatalf("expect json %#v, actual %#v", string(expectJson), string(actualJson))
	}
}

func TestUnmarshalTransactionStatus(t *testing.T) {
	testUnmarshalTransactionStatus(t, []byte("\"packed\""), TransactionStatus{
		packedOrReady: "packed",
	})
	testUnmarshalTransactionStatus(t, []byte("{\"pending\": \"futureNonce\"}"), TransactionStatus{
		pending: pending{"futureNonce"},
	})
}

func testUnmarshalTransactionStatus(t *testing.T, originTxStatusJson []byte, expectTxStatus TransactionStatus) {
	actualTxStatus := TransactionStatus{}
	json.Unmarshal(originTxStatusJson, &actualTxStatus)
	if !reflect.DeepEqual(actualTxStatus, expectTxStatus) {
		t.Fatalf("expect %#v, actual %#v", expectTxStatus, actualTxStatus)
	}
}

func fatalIfErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
