package secp256k1

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_S256Element_String(t *testing.T) {
	assert := assert.New(t)

	e := NewS256Element(big.NewInt(1234567890))
	assert.EqualValues("00000000000000000000000000000000000000000000000000000000499602d2", e.String())
}
