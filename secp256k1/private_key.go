package secp256k1

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type PrivateKey struct {
	secret    *big.Int
	publicKey *PublicKey
}

func NewPrivateKey(secret *big.Int) *PrivateKey {
	return &PrivateKey{
		secret:    secret,
		publicKey: &PublicKey{pt: CurveParams.G.Mul(secret)},
	}
}

func NewRandomPrivateKey() *PrivateKey {
	secret := generateRandomPoint()
	return &PrivateKey{
		secret:    secret,
		publicKey: &PublicKey{pt: CurveParams.G.Mul(secret)},
	}
}

func (key *PrivateKey) String() string {
	return fmt.Sprintf("%064x", key.secret)
}

func (key *PrivateKey) GetPublicKey() *PublicKey {
	return key.publicKey
}

func (key *PrivateKey) Sign(z *big.Int) *Signature {
	k := generateRandomPoint()
	r := CurveParams.G.Mul(k).X().Num()
	kInv := Inverse(k)
	s := new(big.Int).Mul(
		(new(big.Int).Add(z,
			new(big.Int).Mul(r, key.secret))), kInv)
	s = s.Mod(s, CurveParams.N)
	// ?????
	// if new(big.Int).Div(CurveParams.N, big.NewInt(2)).Cmp(s) < 0 {
	// 	s.Sub(CurveParams.N, s)
	// }
	return NewSignature(r, s)
}

func generateRandomPoint() *big.Int {
	for {
		b := make([]byte, 32)
		_, err := rand.Read(b)
		if err != nil {
			panic("failed to generate random")
		}
		k := new(big.Int).SetBytes(b)
		if k.Cmp(CurveParams.N) < 0 {
			return k
		}
	}
}
