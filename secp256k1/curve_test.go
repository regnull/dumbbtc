package secp256k1

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"teralyt.com/dumbbtc/ecc"
)

func TestCurve(t *testing.T) {
	assert := assert.New(t)

	x := ecc.NewFieldElement(CurveParams.Gx, CurveParams.P)
	y := ecc.NewFieldElement(CurveParams.Gy, CurveParams.P)
	point := ecc.NewPoint(x, y, ecc.NewFieldElement(CurveParams.A, CurveParams.P), ecc.NewFieldElement(CurveParams.B, CurveParams.P))
	p1 := point.Mul(CurveParams.N)
	assert.True(p1.IsInf())
}

func TestInverse(t *testing.T) {
	assert := assert.New(t)

	x := big.NewInt(1234567890)
	xInv := Inverse(x)
	m := new(big.Int).Mul(x, xInv)
	m.Mod(m, CurveParams.N)
	assert.Equal(big.NewInt(1), m)
}
