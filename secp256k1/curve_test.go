package secp256k1

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"teralyt.com/dumbbtc/ecc"
)

func TestCurve(t *testing.T) {
	assert := assert.New(t)

	x := ecc.NewFieldElement(gx, p)
	y := ecc.NewFieldElement(gy, p)
	point := ecc.NewPoint(x, y, ecc.NewFieldElement(a, p), ecc.NewFieldElement(b, p))
	p1 := point.Mul(n)
	assert.True(p1.IsInf())
}
