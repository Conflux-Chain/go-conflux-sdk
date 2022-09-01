package bulk

import (
	"fmt"
	"math/big"
	"testing"

	client "github.com/Conflux-Chain/go-conflux-sdk"
	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/status-im/keycard-go/hexutils"
	"github.com/stretchr/testify/assert"
)

func _initClinetForTest() *sdk.Client {
	_client, err := client.NewClient("https://test.confluxrpc.com", client.ClientOption{
		KeystorePath: "keystore",
	})
	if err != nil {
		panic(err)
	}
	if len(_client.AccountManager.List()) == 0 {
		_client.AccountManager.Create("")
	}
	return _client
}

// nil, nil and will err, nil, nil and will err, nil => 0, nil, 1, nil, 2
// nil, 3 , nil, nil, nil => 0, 3, 1, 2, 4
func TestNonceCorrectWhenBulkSendPopulate(t *testing.T) {

	bulkSender := NewBulkSender(*_initClinetForTest())

	user := cfxaddress.MustNew("cfxtest:aaskvgxcfej371g4ecepx9an78ngpejvcekupe69t3")
	// value := types.NewBigInt(10000)
	usdt := cfxaddress.MustNew("cfxtest:acepe88unk7fvs18436178up33hb4zkuf62a9dk1gv")
	dtatOfTransfer1000 := hexutils.HexToBytes("32f289cf00000000000000000000000088c27bd05a7a58bafed6797efa0cce4e1d55302f")

	bulkSender.
		AppendTransaction(&types.UnsignedTransaction{types.UnsignedTransactionBase{From: &user}, &user, nil}).                // correct
		AppendTransaction(&types.UnsignedTransaction{types.UnsignedTransactionBase{From: &user}, &usdt, dtatOfTransfer1000}). // fail
		AppendTransaction(&types.UnsignedTransaction{types.UnsignedTransactionBase{From: &user}, &user, nil}).                // correct
		AppendTransaction(&types.UnsignedTransaction{types.UnsignedTransactionBase{From: &user}, &usdt, dtatOfTransfer1000}). // fail
		AppendTransaction(&types.UnsignedTransaction{types.UnsignedTransactionBase{From: &user}, &user, nil})                 // correct
	populated, err := bulkSender.PopulateTransactions(false)
	fmt.Printf("%v\n", utils.PrettyJSON(populated))
	fmt.Printf("error %+v\n", err)
	assert.True(t, err != nil)
	assert.True(t, populated != nil)
	assert.True(t, populated[0].Nonce.ToInt().Cmp(big.NewInt(0)) == 0)
	assert.True(t, populated[1].Nonce == nil)
	assert.True(t, populated[2].Nonce.ToInt().Cmp(big.NewInt(1)) == 0)
	assert.True(t, populated[3].Nonce == nil)
	assert.True(t, populated[4].Nonce.ToInt().Cmp(big.NewInt(2)) == 0)

	bulkSender.Clear()
	bulkSender.
		AppendTransaction(&types.UnsignedTransaction{types.UnsignedTransactionBase{From: &user}, &user, nil}).
		AppendTransaction(&types.UnsignedTransaction{types.UnsignedTransactionBase{From: &user, Nonce: types.NewBigInt(3)}, &usdt, nil}).
		AppendTransaction(&types.UnsignedTransaction{types.UnsignedTransactionBase{From: &user}, &user, nil}).
		AppendTransaction(&types.UnsignedTransaction{types.UnsignedTransactionBase{From: &user}, &user, nil}).
		AppendTransaction(&types.UnsignedTransaction{types.UnsignedTransactionBase{From: &user}, &user, nil})
	populated, err = bulkSender.PopulateTransactions(false)
	fmt.Printf("%v\n", utils.PrettyJSON(populated))
	fmt.Printf("error %+v\n", err)
	assert.True(t, err != nil)
	assert.True(t, populated != nil)
	assert.True(t, populated[0].Nonce.ToInt().Cmp(big.NewInt(0)) == 0)
	assert.True(t, populated[1].Nonce.ToInt().Cmp(big.NewInt(3)) == 0)
	assert.True(t, populated[2].Nonce.ToInt().Cmp(big.NewInt(1)) == 0)
	assert.True(t, populated[3].Nonce.ToInt().Cmp(big.NewInt(2)) == 0)
	assert.True(t, populated[4].Nonce.ToInt().Cmp(big.NewInt(4)) == 0)
}
