package secp256k1

import (
	"math/big"

	"teralyt.com/dumbbtc/ecc"
)

type S256Point struct {
	*ecc.Point
}

func NewS256Point(x, y *S256Element) *S256Point {
	return &S256Point{ecc.NewPoint(x.FieldElement, y.FieldElement,
		NewS256Element(CurveParams.A).FieldElement, NewS256Element(CurveParams.B).FieldElement)}
}

// func (p *S256Point) VerifySignature(z *big.Int, sig *Signature) bool {
// 	s := NewS256Element(sig.s)
// 	nMinusTwo := new(big.Int).Sub(CurveParams.N, big.NewInt(2))
// 	sInv := s.Pow(nMinusTwo)
// 	u := NewS256Element(z).Mul(sInv)
// 	v := NewS256Element(sig.r).Mul(sInv)
// 	total := CurveParams.G.Mul(u.Num()).Add(p.Mul(v.Num()))
// 	return total.X().Equal(NewS256Element(sig.r).FieldElement)
// }

func (p *S256Point) Mul(factor *big.Int) *S256Point {
	f := new(big.Int).Mod(factor, CurveParams.N)
	p1 := p.Point.Mul(f)
	return &S256Point{p1}
}
