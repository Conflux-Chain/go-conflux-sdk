package cfxaddress

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestCfxAddress(t *testing.T) {
	verify(t, "85d80245dc02f5a89589e1f19c5c718e405b56cd", 1029, "cfx:acc7uawf5ubtnmezvhu9dhc6sghea0403y2dgpyfjp")
	verify(t, "85d80245dc02f5a89589e1f19c5c718e405b56cd", 1, "cfxtest:acc7uawf5ubtnmezvhu9dhc6sghea0403ywjz6wtpg")
	verify(t, "85d80245dc02f5a89589e1f19c5c718e405b56cd", 1, "cfxtest:type.contract:acc7uawf5ubtnmezvhu9dhc6sghea0403ywjz6wtpg")

	verify(t, "1a2f80341409639ea6a35bbcab8299066109aa55", 1029, "cfx:aarc9abycue0hhzgyrr53m6cxedgccrmmyybjgh4xg")
	verify(t, "1a2f80341409639ea6a35bbcab8299066109aa55", 1, "cfxtest:aarc9abycue0hhzgyrr53m6cxedgccrmmy8m50bu1p")
	verify(t, "1a2f80341409639ea6a35bbcab8299066109aa55", 1, "cfxtest:type.user:aarc9abycue0hhzgyrr53m6cxedgccrmmy8m50bu1p")

	verify(t, "19c742cec42b9e4eff3b84cdedcde2f58a36f44f", 1029, "cfx:aap6su0s2uz36x19hscp55sr6n42yr1yk6r2rx2eh7")
	verify(t, "19c742cec42b9e4eff3b84cdedcde2f58a36f44f", 1, "cfxtest:aap6su0s2uz36x19hscp55sr6n42yr1yk6hx8d8sd1")
	verify(t, "19c742cec42b9e4eff3b84cdedcde2f58a36f44f", 1, "cfxtest:type.user:aap6su0s2uz36x19hscp55sr6n42yr1yk6hx8d8sd1")

	verify(t, "84980a94d94f54ac335109393c08c866a21b1b0e", 1029, "cfx:acckucyy5fhzknbxmeexwtaj3bxmeg25b2b50pta6v")
	verify(t, "84980a94d94f54ac335109393c08c866a21b1b0e", 1, "cfxtest:acckucyy5fhzknbxmeexwtaj3bxmeg25b2nuf6km25")
	verify(t, "84980a94d94f54ac335109393c08c866a21b1b0e", 1, "cfxtest:type.contract:acckucyy5fhzknbxmeexwtaj3bxmeg25b2nuf6km25")

	verify(t, "1cdf3969a428a750b89b33cf93c96560e2bd17d1", 1029, "cfx:aasr8snkyuymsyf2xp369e8kpzusftj14ec1n0vxj1")
	verify(t, "1cdf3969a428a750b89b33cf93c96560e2bd17d1", 1, "cfxtest:aasr8snkyuymsyf2xp369e8kpzusftj14ej62g13p7")
	verify(t, "1cdf3969a428a750b89b33cf93c96560e2bd17d1", 1, "cfxtest:type.user:aasr8snkyuymsyf2xp369e8kpzusftj14ej62g13p7")

	verify(t, "0888000000000000000000000000000000000002", 1029, "cfx:aaejuaaaaaaaaaaaaaaaaaaaaaaaaaaaajrwuc9jnb")
	verify(t, "0888000000000000000000000000000000000002", 1, "cfxtest:aaejuaaaaaaaaaaaaaaaaaaaaaaaaaaaajh3dw3ctn")
	verify(t, "0888000000000000000000000000000000000002", 1, "cfxtest:type.builtin:aaejuaaaaaaaaaaaaaaaaaaaaaaaaaaaajh3dw3ctn")
}

func TestMarshalJSON(t *testing.T) {
	cfxAddressFromHex, e := NewFromHex("1cdf3969a428a750b89b33cf93c96560e2bd17d1", 1029)
	fatalIfErr(t, e)
	j, e := json.Marshal(cfxAddressFromHex)
	// encoding.TextMarshaler
	fatalIfErr(t, e)
	expect := "\"CFX:TYPE.USER:AASR8SNKYUYMSYF2XP369E8KPZUSFTJ14EC1N0VXJ1\""
	if string(j) != expect {
		t.Fatalf("expect %#v, actual %#v", expect, string(j))
	}
}

func TestUnmarshalJSON(t *testing.T) {
	fmt.Println("start")
	var actual Address
	err := json.Unmarshal([]byte("null"), &actual)
	fatalIfErr(t, err)
	err = json.Unmarshal([]byte("\"CFX:TYPE.USER:AASR8SNKYUYMSYF2XP369E8KPZUSFTJ14EC1N0VXJ1\""), &actual)
	fatalIfErr(t, err)
	expect, err := NewFromHex("1cdf3969a428a750b89b33cf93c96560e2bd17d1", 1029)
	fatalIfErr(t, err)
	if !reflect.DeepEqual(actual, expect) {
		t.Fatalf("expect %#v, actual %#v", expect, actual)
	}

}

func verify(t *testing.T, hexAddressStr string, networkID uint32, base32Address string) {
	cfxAddressFromHex, err := NewFromHex(hexAddressStr, networkID)
	fatalIfErr(t, err)

	// fmt.Printf("cfxAddressFromHex %v\n", cfxAddressFromHex)
	cfxAddressFromBase32, err := NewFromBase32(base32Address)
	fatalIfErr(t, err)

	if !reflect.DeepEqual(cfxAddressFromHex, cfxAddressFromBase32) {
		t.Fatalf("expect %v, actual %v", cfxAddressFromHex.ShortString(), cfxAddressFromBase32.ShortString())
	}
}

func fatalIfErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
