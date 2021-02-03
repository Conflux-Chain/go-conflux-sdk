package cfxaddress

// FormatAddressStrToHex format hex or base32 address to hex string
func FormatAddressStrToHex(address string) string {
	if address == "" || address[0:2] == "0x" {
		return address
	}
	cfxAddr := MustNewFromBase32(address)
	return "0x" + cfxAddr.GetHexAddress()
}
