package secp256k1

import (
	"fmt"
)

type PublicKey struct {
	pt *S256Point
}

func (pk *PublicKey) String() string {
	return fmt.Sprintf("%064x,%064x", pk.pt.X().Num(), pk.pt.Y().Num())
}
