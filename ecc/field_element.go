package ecc

import (
	"fmt"
	"math/big"
)

type FieldElement struct {
	num   *big.Int
	prime *big.Int
}

var zero = big.NewInt(0)

func NewFieldElement(num, prime *big.Int) *FieldElement {
	if num.Cmp(prime) >= 0 || num.Cmp(zero) < 0 {
		panic("bad number")
	}
	return &FieldElement{num: num, prime: prime}
}

func NewFieldElementFromInt(num, prime int64) *FieldElement {
	return NewFieldElement(big.NewInt(num), big.NewInt(prime))
}

func (fe *FieldElement) String() string {
	return fmt.Sprintf("FieldElement_%s (%s)", fe.num.String(), fe.prime.String())
}

func (fe *FieldElement) Equal(other *FieldElement) bool {
	if other == nil {
		return false
	}

	return fe.num.Cmp(other.num) == 0 && fe.prime.Cmp(other.prime) == 0
}

func (fe *FieldElement) NotEqual(other *FieldElement) bool {
	return !fe.Equal(other)
}

func (fe *FieldElement) Add(other *FieldElement) *FieldElement {
	if fe.prime.Cmp(other.prime) != 0 {
		panic("bad number")
	}

	r := new(big.Int)
	r.Add(fe.num, other.num)
	r.Mod(r, fe.prime)
	return NewFieldElement(r, fe.prime)
}

func (fe *FieldElement) Sub(other *FieldElement) *FieldElement {
	if fe.prime.Cmp(other.prime) != 0 {
		panic("bad number")
	}

	r := new(big.Int)
	r.Sub(fe.num, other.num)
	r.Mod(r, fe.prime)
	return NewFieldElement(r, fe.prime)
}

func (fe *FieldElement) Mul(other *FieldElement) *FieldElement {
	if fe.prime.Cmp(other.prime) != 0 {
		panic("bad number")
	}
	r := new(big.Int)
	r.Mul(fe.num, other.num)
	r.Mod(r, fe.prime)
	return NewFieldElement(r, fe.prime)
}

func (fe *FieldElement) Pow(exp *big.Int) *FieldElement {
	n := new(big.Int)
	n.Mod(exp, new(big.Int).Sub(fe.prime, big.NewInt(1)))
	r := new(big.Int)
	r.Exp(fe.num, n, fe.prime)
	return NewFieldElement(r, fe.prime)
}

func (fe *FieldElement) Div(other *FieldElement) *FieldElement {
	if fe.prime.Cmp(other.prime) != 0 {
		panic("bad number")
	}

	return fe.Mul(other.Pow(new(big.Int).Sub(other.prime, big.NewInt(2))))
}
