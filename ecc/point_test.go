package ecc

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPoint(t *testing.T) {
	assert := assert.New(t)

	p := NewPoint(NewInt64(-1), NewInt64(-1), NewInt64(5), NewInt64(7))
	assert.EqualValues(NewInt64(-1), p.x)
	assert.EqualValues(NewInt64(-1), p.y)
	assert.EqualValues(NewInt64(5), p.a)
	assert.EqualValues(NewInt64(7), p.b)
}

func TestPointInf(t *testing.T) {
	assert := assert.New(t)

	p := NewPoint(NewInt64(-1), NewInt64(-1), NewInt64(5), NewInt64(7))
	assert.False(p.IsInf())

	inf := NewInf(NewInt64(5), NewInt64(7))
	assert.True(inf.IsInf())
}

func TestPointEqual(t *testing.T) {
	assert := assert.New(t)

	a := NewPoint(NewInt64(-1), NewInt64(-1), NewInt64(5), NewInt64(7))
	b := NewPoint(NewInt64(-1), NewInt64(1), NewInt64(5), NewInt64(7))
	assert.False(a.Equal(b))
	assert.True(a.Equal(a))
}

func TestPointAdd(t *testing.T) {
	assert := assert.New(t)

	a := NewPoint(NewInt64(-1), NewInt64(-1), NewInt64(5), NewInt64(7))
	b := NewPoint(NewInt64(-1), NewInt64(1), NewInt64(5), NewInt64(7))
	inf := NewInf(NewInt64(5), NewInt64(7))

	a1 := a.Add(inf)
	assert.Equal(a, a1)

	b1 := b.Add(inf)
	assert.Equal(b, b1)

	inf1 := a.Add(b)
	assert.Equal(inf, inf1)

	c := NewPoint(NewInt64(2), NewInt64(5), NewInt64(5), NewInt64(7))
	assert.Equal(NewPoint(NewInt64(3), NewInt64(-7), NewInt64(5), NewInt64(7)), c.Add(a))
}

func TestPointMul(t *testing.T) {
	assert := assert.New(t)

	a := NewPoint(NewInt64(-1), NewInt64(-1), NewInt64(5), NewInt64(7))
	a3 := a.Add(a).Add(a)

	at3 := a.Mul(big.NewInt(3))
	assert.Equal(a3, at3)
}
