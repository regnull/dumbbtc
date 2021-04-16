package ecc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPoint(t *testing.T) {
	assert := assert.New(t)

	p := NewPoint(-1, -1, 5, 7)
	assert.EqualValues(-1, p.x)
	assert.EqualValues(-1, p.y)
	assert.EqualValues(5, p.a)
	assert.EqualValues(7, p.b)
}

func TestPointInf(t *testing.T) {
	assert := assert.New(t)

	p := NewPoint(-1, -1, 5, 7)
	assert.False(p.IsInf())

	inf := NewInf(5, 7)
	assert.True(inf.IsInf())
}

func TestPointEqual(t *testing.T) {
	assert := assert.New(t)

	a := NewPoint(-1, -1, 5, 7)
	b := NewPoint(-1, 1, 5, 7)
	assert.False(a.Equal(b))
	assert.True(a.Equal(a))
}

func TestPointAdd(t *testing.T) {
	assert := assert.New(t)

	a := NewPoint(-1, -1, 5, 7)
	b := NewPoint(-1, 1, 5, 7)
	inf := NewInf(5, 7)

	a1 := a.Add(inf)
	assert.Equal(a, a1)

	b1 := b.Add(inf)
	assert.Equal(b, b1)

	inf1 := a.Add(b)
	assert.Equal(inf, inf1)
}
