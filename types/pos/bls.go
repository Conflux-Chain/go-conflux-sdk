package postypes

import (
	"crypto/sha256"
	"math/big"

	"github.com/Conflux-Chain/go-conflux-sdk/light/bcs"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	// "github.com/ethereum/go-ethereum/crypto/bls12381"
	bls12381 "github.com/kilic/bls12-381"
	"github.com/pkg/errors"
)

var (
	bcsPrefix []byte = hexutil.MustDecode("0xcd510d1ab583c33b54fa949014601df0664857c18c4cfb228c862dd869df1b62")

	hashToField_p *big.Int
	DST           []byte = []byte("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_NUL_")
)

const (
	L                = 64
	H_IN_CHUNK_SIZE  = 64
	H_OUT_CHUNK_SIZE = 32
)

func init() {
	p, ok := new(big.Int).SetString("4002409555221667393417789825735904156556882819939007885332058136124031650490837864442687629129015664037894272559787", 10)
	if !ok {
		panic("invalid p value")
	}

	hashToField_p = p
}

type Committee struct {
	state            *EpochState
	uncompressedKeys map[common.Hash]hexutil.Bytes // pos account => uncompressed BLS public key
}

func (c *Committee) GetPublicKey(account common.Hash) (pubKey []byte, ok bool) {
	pubKey, ok = c.uncompressedKeys[account]
	return
}

func (info *LedgerInfoWithSignatures) NextCommittee() (Committee, bool) {
	if info.LedgerInfo.CommitInfo.NextEpochState == nil {
		return Committee{}, false
	}

	return Committee{
		state:            info.LedgerInfo.CommitInfo.NextEpochState,
		uncompressedKeys: info.NextEpochValidators,
	}, true
}

func (info *LedgerInfoWithSignatures) EncodeBCS() []byte {
	encoded := bcs.MustEncodeToBytes(info.LedgerInfo)
	encoded = append(bcsPrefix, encoded...)
	return encoded
}

func (info *LedgerInfoWithSignatures) Verify(committee Committee) (bool, error) {
	if len(info.Signatures) == 0 {
		return false, nil
	}

	if info.LedgerInfo.CommitInfo.Epoch != committee.state.Epoch {
		return false, errors.Errorf("Epoch mismatch, ledger: %v, committee: %v",
			uint64(info.LedgerInfo.CommitInfo.Epoch), uint64(committee.state.Epoch))
	}

	bcsEncoded := bcs.MustEncodeToBytes(info.LedgerInfo)
	bcsEncoded = append(bcsPrefix, bcsEncoded...)
	hash, err := hashToCurve(bcsEncoded)
	if err != nil {
		return false, errors.WithMessage(err, "Failed to hash message into G2 point")
	}

	var votes hexutil.Uint64

	for account := range info.Signatures {
		votes += committee.state.Verifier.AddressToValidatorInfo[account].VotingPower
	}

	if votes < committee.state.Verifier.QuorumVotingPower {
		return false, errors.Errorf("Votes not enough, expected: %v, actual: %v",
			uint64(committee.state.Verifier.QuorumVotingPower), uint64(votes))
	}

	accPubKey, err := info.AggregatedPublicKey(committee)
	if err != nil {
		return false, errors.WithMessage(err, "Failed to aggregate public keys")
	}

	verified, err := verifyBLS(info.AggregatedSignature, accPubKey, hash)
	if err != nil {
		return false, errors.WithMessage(err, "Failed to verify BLS signature")
	}

	return verified, nil
}

func (info *LedgerInfoWithSignatures) AggregatedPublicKey(committee Committee) (*bls12381.PointG1, error) {
	g1 := bls12381.NewG1()
	acc := g1.Zero()

	for _, account := range info.ValidatorsSorted() {
		publicKey, ok := committee.uncompressedKeys[account]
		if !ok {
			return nil, errors.Errorf("PoS account %v not found in committee", account)
		}

		cur, err := g1.FromBytes(publicKey)
		if err != nil {
			return nil, errors.WithMessage(err, "Failed to decode public key to G1 point")
		}

		g1.Add(acc, acc, cur)
	}

	return acc, nil
}

func verifyBLS(signature []byte, publicKey *bls12381.PointG1, hash *bls12381.PointG2) (bool, error) {
	engine := bls12381.NewEngine()

	// signature is uncompressed BLS signature in 192 bytes
	sigPointG2, err := engine.G2.FromBytes(signature)
	if err != nil {
		return false, errors.WithMessage(err, "Failed to decode signature to G2 point")
	}

	return engine.
		AddPair(publicKey, hash).
		AddPairInv(engine.G1.One(), sigPointG2).
		Check(), nil
}

func hashToCurve(msg []byte) (*bls12381.PointG2, error) {
	fe := hashToField(msg, 2, 2, DST)

	p0, err := mapToCurve(fe[0])
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to do mapToCurve for p0")
	}

	p1, err := mapToCurve(fe[1])
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to do mapToCurve for p1")
	}

	g := bls12381.NewG2()
	r := g.New()
	g.Add(r, p0, p1)

	return r, nil
}

func mapToCurve(input []byte) (*bls12381.PointG2, error) {
	// if len(input) != 128 {
	// 	panic("invalid length")
	// }

	fe := make([]byte, 96)
	c0 := mustDecodeBLS12381FieldElement(input[:64])
	copy(fe[48:], c0)
	c1 := mustDecodeBLS12381FieldElement(input[64:])
	copy(fe[:48], c1)

	g := bls12381.NewG2()
	r, err := g.MapToCurve(fe)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to do mapToCurve")
	}

	return r, nil
}

func mustDecodeBLS12381FieldElement(in []byte) []byte {
	// if len(in) != 64 {
	// 	panic("invalid field element length")
	// }

	// check top bytes
	for i := 0; i < 16; i++ {
		if in[i] != byte(0x00) {
			panic("invalid field element prefix")
		}
	}

	out := make([]byte, 48)
	copy(out[:], in[16:])

	return out
}

func hashToField(msg []byte, count int, degree int, dst []byte) [][]byte {
	expandedLen := L * degree * count
	expanded := expandMessageXmd(msg, expandedLen, dst)

	for i := 0; i < expandedLen; i += L {
		chunk := expanded[i : i+L]
		chunkBig := new(big.Int).SetBytes(chunk)
		mod := new(big.Int).Mod(chunkBig, hashToField_p)
		copy(expanded[i:i+L], big2OSP(mod, L))
	}

	return [][]byte{
		expanded[:expandedLen/2],
		expanded[expandedLen/2:],
	}
}

func expandMessageXmd(msg []byte, outputLen int, dst []byte) []byte {
	// if outputLen > 65535 || outputLen > 255*H_OUT_CHUNK_SIZE {
	// 	panic("outputLen too large")
	// }

	// if len(dst) > 255 {
	// 	panic("DST length too large")
	// }

	suffix := append(dst, i2OSP(len(dst), 1)...)

	ell := outputLen / H_OUT_CHUNK_SIZE
	if outputLen%H_OUT_CHUNK_SIZE != 0 {
		ell++
	}

	var input0 []byte
	input0 = append(input0, i2OSP(0, H_IN_CHUNK_SIZE)...)
	input0 = append(input0, msg...)
	input0 = append(input0, i2OSP(outputLen, 2)...)
	input0 = append(input0, i2OSP(0, 1)...)
	input0 = append(input0, suffix...)

	b := [][32]byte{
		sha256.Sum256(input0),
	}

	var input1 []byte
	input1 = append(input1, b[0][:]...)
	input1 = append(input1, i2OSP(1, 1)...)
	input1 = append(input1, suffix...)
	b = append(b, sha256.Sum256(input1))

	for i := 2; i <= ell; i++ {
		var input []byte
		input = append(input, xor(b[0][:], b[i-1][:])...)
		input = append(input, i2OSP(i, 1)...)
		input = append(input, suffix...)
		b = append(b, sha256.Sum256(input))
	}

	var flatten []byte
	for _, v := range b[1:] {
		flatten = append(flatten, v[:]...)
	}

	if len(flatten) <= outputLen {
		return flatten
	}

	return flatten[:outputLen]
}

func xor(b1, b2 []byte) []byte {
	var result []byte

	for i, v := range b1 {
		result = append(result, v^b2[i])
	}

	return result
}

func i2OSP(x, len int) []byte {
	result := make([]byte, len)
	index := len - 1

	for x > 0 {
		result[index] = byte(x & 0xFF)
		index--
		x >>= 8
	}

	return result
}

func big2OSP(x *big.Int, len int) []byte {
	result := make([]byte, len)
	index := len - 1
	for x.Sign() > 0 {
		val := new(big.Int).And(x, big.NewInt(0xFF))

		result[index] = byte(val.Uint64())
		index--
		x.Div(x, big.NewInt(256))
	}

	return result
}
