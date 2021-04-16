package ecc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	assert := assert.New(t)
	a := NewFieldElement(7, 13)
	b := NewFieldElement(6, 13)

	assert.False(a.Equal(b))
	assert.True(a.Equal(a))
}

func TestAdd(t *testing.T) {
	assert := assert.New(t)
	a := NewFieldElement(44, 57)
	b := NewFieldElement(33, 57)

	c := a.Add(b)
	assert.Equal(NewFieldElement(20, 57), c)
}

func TestMultiply(t *testing.T) {
	assert := assert.New(t)

	a := NewFieldElement(3, 13)
	b := NewFieldElement(12, 13)
	assert.Equal(NewFieldElement(10, 13), a.Multiply(b))
}

func TestPow(t *testing.T) {
	assert := assert.New(t)

	a := NewFieldElement(3, 13)
	assert.Equal(NewFieldElement(1, 13), a.Pow(3))
}

func TestDivive(t *testing.T) {
	assert := assert.New(t)

	a := NewFieldElement(2, 19)
	b := NewFieldElement(7, 19)
	assert.Equal(NewFieldElement(3, 19), a.Divide(b))
}
