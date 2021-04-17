package ecc

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FieldElement_Equal(t *testing.T) {
	assert := assert.New(t)
	a := NewFieldElementFromInt(7, 13)
	b := NewFieldElementFromInt(6, 13)

	assert.False(a.Equal(b))
	assert.True(a.Equal(a))
}

func Test_FieldElement_Add(t *testing.T) {
	assert := assert.New(t)

	prime := int64(57)
	tests := []struct {
		one, two int64
		res      int64
	}{
		{44, 33, 20},
		{55, 54, 52},
	}

	for _, test := range tests {
		a := NewFieldElementFromInt(test.one, prime)
		b := NewFieldElementFromInt(test.two, prime)
		res := a.Add(b)
		expected := NewFieldElementFromInt(test.res, prime)
		assert.Equal(expected, res)
	}
}

func Test_FieldElement_Sub(t *testing.T) {
	assert := assert.New(t)

	prime := int64(57)
	tests := []struct {
		one, two int64
		res      int64
	}{
		{55, 2, 53},
		{0, 10, 47},
		{3, 4, 56},
	}
	for _, test := range tests {
		a := NewFieldElementFromInt(test.one, prime)
		b := NewFieldElementFromInt(test.two, prime)
		res := a.Sub(b)
		expected := NewFieldElementFromInt(test.res, prime)
		assert.Equal(expected, res)
	}
}

func Test_FieldElement_Mul(t *testing.T) {
	assert := assert.New(t)

	prime := int64(97)
	tests := []struct {
		one, two int64
		res      int64
	}{
		{95, 45, 7},
		{7, 31, 23},
		{17, 13, 27},
		{27, 19, 28},
		{28, 44, 68},
		{8, 20, 63},
	}
	for _, test := range tests {
		a := NewFieldElementFromInt(test.one, prime)
		b := NewFieldElementFromInt(test.two, prime)
		res := a.Mul(b)
		expected := NewFieldElementFromInt(test.res, prime)
		assert.Equal(expected, res)
	}

	prime = int64(31)
	tests = []struct {
		one, two int64
		res      int64
	}{
		{4, 11, 13},
	}
	for _, test := range tests {
		a := NewFieldElementFromInt(test.one, prime)
		b := NewFieldElementFromInt(test.two, prime)
		res := a.Mul(b)
		expected := NewFieldElementFromInt(test.res, prime)
		assert.Equal(expected, res)
	}
}

func Test_FieldElement_Pow(t *testing.T) {
	assert := assert.New(t)

	prime := int64(97)
	tests := []struct {
		one, two int64
		res      int64
	}{
		{12, 7, 8},
		{77, 49, 20},
		{11, -3, 79},
	}
	for _, test := range tests {
		a := NewFieldElementFromInt(test.one, prime)
		b := big.NewInt(test.two)
		res := a.Pow(b)
		expected := NewFieldElementFromInt(test.res, prime)
		assert.Equal(expected, res)
	}

	prime = int64(31)
	tests = []struct {
		one, two int64
		res      int64
	}{
		{17, -3, 29},
		{4, -4, 4},
	}
	for _, test := range tests {
		a := NewFieldElementFromInt(test.one, prime)
		b := big.NewInt(test.two)
		res := a.Pow(b)
		expected := NewFieldElementFromInt(test.res, prime)
		assert.Equal(expected, res)
	}
}

func Test_FieldElement_Div(t *testing.T) {
	assert := assert.New(t)

	prime := int64(31)
	tests := []struct {
		one, two int64
		res      int64
	}{
		{3, 24, 4},
	}
	for _, test := range tests {
		a := NewFieldElementFromInt(test.one, prime)
		b := NewFieldElementFromInt(test.two, prime)
		res := a.Div(b)
		expected := NewFieldElementFromInt(test.res, prime)
		assert.Equal(expected, res)
	}
}
