package postypes

import "github.com/ethereum/go-ethereum/common/hexutil"

// BLS related data structures that supports BCS serialization.

type ConsensusSignature [192]byte

// MarshalText implements the encoding.TextMarshaler interface to return hex representation of s.
func (s ConsensusSignature) MarshalText() (text []byte, err error) {
	return hexutil.Bytes(s[:]).MarshalText()
}

// UnmarshalText implements the encoding.TextUnmarshaler interface to parse a signature in hex syntax.
func (s *ConsensusSignature) UnmarshalText(input []byte) error {
	return hexutil.UnmarshalFixedUnprefixedText("ConsensusSignature", input, s[:])
}

// String implements the fmt.Stringer interface to return hex value with 0x prefix.
func (s ConsensusSignature) String() string {
	return hexutil.Encode(s[:])
}

type ConsensusPublicKey [48]byte

// MarshalText implements the encoding.TextMarshaler interface to return hex representation of k.
func (k ConsensusPublicKey) MarshalText() (text []byte, err error) {
	return hexutil.Bytes(k[:]).MarshalText()
}

// UnmarshalText implements the encoding.TextUnmarshaler interface to parse a public key in hex syntax.
func (k *ConsensusPublicKey) UnmarshalText(input []byte) error {
	return hexutil.UnmarshalFixedUnprefixedText("ConsensusPublicKey", input, k[:])
}

// String implements the fmt.Stringer interface to return hex value with 0x prefix.
func (k ConsensusPublicKey) String() string {
	return hexutil.Encode(k[:])
}

type ConsensusVRFPublicKey [33]byte

// MarshalText implements the encoding.TextMarshaler interface to return hex representation of k.
func (k ConsensusVRFPublicKey) MarshalText() (text []byte, err error) {
	return hexutil.Bytes(k[:]).MarshalText()
}

// UnmarshalText implements the encoding.TextUnmarshaler interface to parse a VRF public key in hex syntax.
func (k *ConsensusVRFPublicKey) UnmarshalText(input []byte) error {
	return hexutil.UnmarshalFixedUnprefixedText("ConsensusVRFPublicKey", input, k[:])
}

// String implements the fmt.Stringer interface to return hex value with 0x prefix.
func (k ConsensusVRFPublicKey) String() string {
	return hexutil.Encode(k[:])
}
