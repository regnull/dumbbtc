package ecc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFieldElementPoint(t *testing.T) {
	assert := assert.New(t)

	prime := 223

	a := NewFieldElement(0, prime)
	b := NewFieldElement(7, prime)

	validPoints := []struct {
		x int
		y int
	}{
		{192, 105},
		{17, 56},
		{1, 193},
	}

	for _, pp := range validPoints {
		x := NewFieldElement(pp.x, prime)
		y := NewFieldElement(pp.y, prime)
		assert.NotNil(NewFieldElementPoint(x, y, a, b))
	}

	invalidPoints := []struct {
		x int
		y int
	}{
		{200, 119},
		{42, 99},
	}

	for _, pp := range invalidPoints {
		x := NewFieldElement(pp.x, prime)
		y := NewFieldElement(pp.y, prime)
		assert.Panics(func() {
			NewFieldElementPoint(x, y, a, b)
		})
	}
}

func TestFieldElementPointInf(t *testing.T) {
	assert := assert.New(t)

	prime := 223

	a := NewFieldElement(0, prime)
	b := NewFieldElement(7, prime)
	p := NewFieldElementPoint(NewFieldElement(192, prime), NewFieldElement(105, prime), a, b)
	assert.False(p.IsInf())

	inf := NewFieldElementInf(a, b)
	assert.True(inf.IsInf())
}

func TestFieldElementAdd(t *testing.T) {
	assert := assert.New(t)

	x := (56 - 105) % 223
	_ = x

	prime := 223
	a := NewFieldElement(0, prime)
	b := NewFieldElement(7, prime)
	p1 := NewFieldElementPoint(NewFieldElement(192, prime), NewFieldElement(105, prime), a, b)
	p2 := NewFieldElementPoint(NewFieldElement(17, prime), NewFieldElement(56, prime), a, b)
	p3 := p1.Add(p2)
	assert.Equal(NewFieldElementPoint(NewFieldElement(170, prime), NewFieldElement(142, prime), a, b), p3)
}
