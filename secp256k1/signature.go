package secp256k1

import (
	"fmt"
	"math/big"
)

type Signature struct {
	r *big.Int
	s *big.Int
}

func NewSignature(r, s *big.Int) *Signature {
	return &Signature{r: r, s: s}
}

func (s *Signature) String() string {
	return fmt.Sprintf("Signature(%x, %x)", s.r, s.s)
}

func (sig *Signature) Verify(z *big.Int, p *PublicKey) bool {
	sInv := Inverse(sig.s)
	u := new(big.Int).Mul(z, sInv)
	u.Mod(u, CurveParams.N)
	v := new(big.Int).Mul(sig.r, sInv)
	v.Mod(v, CurveParams.N)

	total := CurveParams.G.Mul(u).Add(p.pt.Mul(v).Point)
	return total.X().Equal(NewS256Element(sig.r).FieldElement)
}
