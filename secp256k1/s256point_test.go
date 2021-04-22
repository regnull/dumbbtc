package secp256k1

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SignAndVerify(t *testing.T) {
	assert := assert.New(t)

	e := new(big.Int).SetBytes(hash256([]byte("my secret")))
	z := new(big.Int).SetBytes(hash256([]byte("my message")))
	fmt.Printf("%064x\n", z)

	key := NewPrivateKey(e)
	public := key.GetPublicKey()
	sig := key.Sign(z)
	assert.True(sig.Verify(z, public))
}

func Test_Verify(t *testing.T) {
	assert := assert.New(t)

	z, _ := new(big.Int).SetString("ec208baa0fc1c19f708a9ca96fdeff3ac3f230bb4a7ba4aede4942ad003c0f60", 16)
	r, _ := new(big.Int).SetString("ac8d1c87e51d0d441be8b3dd5b05c8795b48875dffe00b7ffcfac23010d3a395", 16)
	s, _ := new(big.Int).SetString("68342ceff8935ededd102dd876ffd6ba72d6a427a3edb13d26eb0781cb423c4", 16)
	px, _ := new(big.Int).SetString("887387e452b8eacc4acfde10d9aaf7f6d9a0f975aabb10d006e4da568744d06c", 16)
	py, _ := new(big.Int).SetString("61de6d95231cd89026e286df3b6ae4a894a3378e393e93a0f45b666329a0ae34", 16)
	public := &PublicKey{
		pt: NewS256Point(
			NewS256Element(px),
			NewS256Element(py))}
	sig := &Signature{
		r: r,
		s: s,
	}
	assert.True(sig.Verify(z, public))
}

func hash256(data []byte) []byte {
	s1 := sha256.Sum256(data)
	s2 := s1[:]
	s3 := sha256.Sum256(s2)
	return s3[:]
}
