// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	utx := UnsignedTransaction{
		UnsignedTransactionBase: UnsignedTransactionBase{
			From: NewAddress("0x1cad0b19bb29d4674531d6f115237e16afce377c"),

			Nonce:    16,
			GasPrice: NewBigInt(32),
			Gas:      64,
			Value:    NewBigInt(128),

			StorageLimit: NewBigInt(256),
			EpochHeight:  NewBigInt(512),
			ChainID:      1024,
		},
		To:   NewAddress("0x1cad0b19bb29d4674531d6f115237e16afce377d"),
		Data: []byte{1, 2, 3},
	}
	expect := []byte{231, 16, 32, 64, 148, 28, 173, 11, 25, 187, 41, 212, 103, 69, 49, 214, 241, 21, 35, 126, 22, 175, 206, 55, 125, 129, 128, 130, 1, 0, 130, 2, 0, 130, 4, 0, 131, 1, 2, 3}
	// oldAPIActual, _ := rlp.EncodeToBytes(utx.getRlpNeedFields())
	newAPIActual, _ := utx.Encode()

	// if !reflect.DeepEqual(expect, oldAPIActual) {
	// 	t.Errorf("expect is %v, old actual is %v", expect, oldAPIActual)
	// }
	if !reflect.DeepEqual(expect, newAPIActual) {
		t.Errorf("expect is %v, new actual is %v", expect, newAPIActual)
	}
}

func TestEncodeWithSignature(t *testing.T) {
	utx := UnsignedTransaction{
		UnsignedTransactionBase: UnsignedTransactionBase{
			From:     NewAddress("0x1cad0b19bb29d4674531d6f115237e16afce377c"),
			Nonce:    16,
			GasPrice: NewBigInt(32),
			Gas:      64,
			Value:    NewBigInt(128),

			StorageLimit: NewBigInt(256),
			EpochHeight:  NewBigInt(512),
			ChainID:      1024,
		},
		To:   NewAddress("0x1cad0b19bb29d4674531d6f115237e16afce377d"),
		Data: []byte{1, 2, 3},
	}
	v := byte(27)
	r := []byte{1, 2, 3, 4, 5}
	s := []byte{0xa, 0xb, 0xc}

	expect := []byte{243, 231, 16, 32, 64, 148, 28, 173, 11, 25, 187, 41, 212, 103, 69, 49, 214, 241, 21, 35, 126, 22, 175, 206, 55, 125, 129, 128, 130, 1, 0, 130, 2, 0, 130, 4, 0, 131, 1, 2, 3, 27, 133, 1, 2, 3, 4, 5, 131, 10, 11, 12}
	// oldAPIActual, _ := rlp.EncodeToBytes([]interface{}{utx.getRlpNeedFields(), v, r, s})
	newAPIActual, _ := utx.EncodeWithSignature(v, r, s)

	// if !reflect.DeepEqual(expect, oldAPIActual) {
	// 	t.Errorf("expect is %v, old actual is %v", expect, oldAPIActual)
	// }
	if !reflect.DeepEqual(expect, newAPIActual) {
		t.Errorf("expect is %v, new actual is %v", expect, newAPIActual)
	}
}

func TestDecodeRlpToUnsignTransction(t *testing.T) {

	rlp := []byte{231, 16, 32, 64, 148, 28, 173, 11, 25, 187, 41, 212, 103, 69, 49, 214, 241, 21, 35, 126, 22, 175, 206, 55, 125, 129, 128, 130, 1, 0, 130, 2, 0, 130, 4, 0, 131, 1, 2, 3}
	expect := &UnsignedTransaction{
		UnsignedTransactionBase: UnsignedTransactionBase{
			From:     nil,
			Nonce:    16,
			GasPrice: NewBigInt(32),
			Gas:      64,
			Value:    NewBigInt(128),

			StorageLimit: NewBigInt(256),
			EpochHeight:  NewBigInt(512),
			ChainID:      1024,
		},
		To:   NewAddress("0x1cad0b19bb29d4674531d6f115237e16afce377d"),
		Data: []byte{1, 2, 3},
	}
	// t.Errorf("%+v", expect)
	actual := new(UnsignedTransaction)
	actual.Decode(rlp)

	jexpect, _ := json.Marshal(expect)
	jactual, _ := json.Marshal(actual)
	if !reflect.DeepEqual(jexpect, jactual) {
		t.Errorf("\njson of expect is %+v,\njson of acutal is %+v", jexpect, jactual)
	} else {
		return
	}

	if string(*expect.From) != string(*actual.From) {
		t.Errorf("expect from is %+v, actual from is %+v", expect.From, actual.From)
	}
	if string(*expect.To) != string(*actual.To) {
		t.Errorf("expect to is %+v, actual to is %+v", *expect.To, *actual.To)
	}
	if expect.Nonce != actual.Nonce {
		t.Errorf("expect Nonce is %+v, actual Nonce is %+v", expect.Nonce, actual.Nonce)
	}
	if expect.GasPrice.ToInt().Cmp(actual.GasPrice.ToInt()) != 0 {
		t.Errorf("expect GasPrice is %+v, actual GasPrice is %+v", expect.GasPrice, actual.GasPrice)
	}
	if expect.Gas != actual.Gas {
		t.Errorf("expect Gas is %+v, actual Gas is %+v", expect.Gas, actual.Gas)
	}
	if expect.Value.ToInt().Cmp(actual.Value.ToInt()) != 0 {
		t.Errorf("expect Value is %+v, actual Value is %+v", expect.Value, actual.Value)
	}
	if !reflect.DeepEqual(expect.Data, actual.Data) {
		t.Errorf("expect Data is %+v, actual Data is %+v", expect.Data, actual.Data)
	}
	if expect.StorageLimit.ToInt().Cmp(actual.StorageLimit.ToInt()) != 0 {
		t.Errorf("expect StorageLimit is %+v, actual StorageLimit is %+v", expect.StorageLimit, actual.StorageLimit)
	}
	if expect.EpochHeight.ToInt().Cmp(actual.EpochHeight.ToInt()) != 0 {
		t.Errorf("expect EpochHeight is %+v, actual EpochHeight is %+v", expect.EpochHeight, actual.EpochHeight)
	}
	if expect.ChainID != actual.ChainID {
		t.Errorf("expect ChainID is %+v, actual ChainID is %+v", expect.ChainID, actual.ChainID)
	}
}
