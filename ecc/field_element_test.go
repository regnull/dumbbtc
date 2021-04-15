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
