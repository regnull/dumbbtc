package secp256k1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PrivateKey_NewRandom(t *testing.T) {
	assert := assert.New(t)

	k := NewRandomPrivateKey()
	s := k.String()
	assert.True(len(s) == 64)
}

func Test_PrivateKey_Sign(t *testing.T) {
	assert := assert.New(t)

	k := NewRandomPrivateKey()
	z := generateRandomPoint()
	sig := k.Sign(z)
	fmt.Printf("%s\n", sig.String())
	assert.NotNil(sig)
}

func Test_PrivateKey_GetPublicKey(t *testing.T) {
	assert := assert.New(t)

	k := NewRandomPrivateKey()
	fmt.Printf("%s\n", k.String())
	fmt.Printf("%s\n", k.GetPublicKey())
	assert.NotNil(k.GetPublicKey())
}
