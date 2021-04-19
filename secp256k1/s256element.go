package secp256k1

import (
	"fmt"
	"math/big"

	"teralyt.com/dumbbtc/ecc"
)

type S256Element struct {
	*ecc.FieldElement
}

func NewS256Element(num *big.Int) *S256Element {
	return &S256Element{ecc.NewFieldElement(num, CurveParams.P)}
}

func (e *S256Element) String() string {
	return fmt.Sprintf("%064x", e.FieldElement.Num())
}
