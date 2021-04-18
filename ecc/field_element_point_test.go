package ecc

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFieldElementPoint(t *testing.T) {
	assert := assert.New(t)

	prime := big.NewInt(223)

	a := NewFieldElement(big.NewInt(0), prime)
	b := NewFieldElement(big.NewInt(7), prime)

	validPoints := []struct {
		x int64
		y int64
	}{
		{192, 105},
		{17, 56},
		{1, 193},
	}

	for _, pp := range validPoints {
		x := NewFieldElement(big.NewInt(pp.x), prime)
		y := NewFieldElement(big.NewInt(pp.y), prime)
		assert.NotNil(NewPoint(x, y, a, b))
	}

	invalidPoints := []struct {
		x int64
		y int64
	}{
		{200, 119},
		{42, 99},
	}

	for _, pp := range invalidPoints {
		x := NewFieldElement(big.NewInt(pp.x), prime)
		y := NewFieldElement(big.NewInt(pp.y), prime)
		assert.Panics(func() {
			NewPoint(x, y, a, b)
		})
	}
}

func TestFieldElementPointotherFieldElement(t *testing.T) {
	assert := assert.New(t)

	prime := big.NewInt(223)

	a := NewFieldElement(big.NewInt(0), prime)
	b := NewFieldElement(big.NewInt(7), prime)
	p := NewPoint(NewFieldElement(big.NewInt(192), prime),
		NewFieldElement(big.NewInt(105), prime), a, b)
	assert.False(p.IsInf())

	inf := NewInf(a, b)
	assert.True(inf.IsInf())
}

type testVal struct {
	x, y int64
}

func TestFieldElementAdd(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		one, two, res testVal
	}{
		{testVal{192, 105}, testVal{17, 56}, testVal{170, 142}},
		{testVal{170, 142}, testVal{60, 139}, testVal{220, 181}},
		{testVal{47, 71}, testVal{17, 56}, testVal{215, 68}},
		{testVal{143, 98}, testVal{76, 66}, testVal{47, 71}},
	}

	prime := big.NewInt(223)
	for _, test := range tests {
		a := NewFieldElement(big.NewInt(0), prime)
		b := NewFieldElement(big.NewInt(7), prime)
		p1 := NewPoint(NewFieldElement(big.NewInt(test.one.x), prime),
			NewFieldElement(big.NewInt(test.one.y), prime), a, b)
		p2 := NewPoint(NewFieldElement(big.NewInt(test.two.x), prime),
			NewFieldElement(big.NewInt(test.two.y), prime), a, b)
		p3 := p1.Add(p2)
		assert.Equal(NewPoint(NewFieldElement(big.NewInt(test.res.x), prime),
			NewFieldElement(big.NewInt(test.res.y), prime), a, b), p3)
	}

}
